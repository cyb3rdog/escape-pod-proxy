using Renci.SshNet;
using Renci.SshNet.Common;
using System;
using System.IO;
using System.Threading.Tasks;

namespace Cyb3rPod
{
    /// <summary>
    /// Deployment event args
    /// </summary>
    /// <seealso cref="System.EventArgs" />
    public class PluginDeploymentEventArgs : EventArgs
    {
        public string Description { get; }
        public double Progress { get; }

        /// <summary>
        /// Initializes a new instance of the <see cref="ProxyMessageEventArgs"/> class.
        /// </summary>
        /// <param name="eventType">Type of the event.</param>
        public PluginDeploymentEventArgs(string description, double progress = 0)
        {
            this.Description = description;
            this.Progress = progress;
        }
    }

    public class EscapePodPlugin : IDisposable
    {
        private string DeploymentStepDescription { get; set; }

        public bool IsEnabled { get; set; }
        public EscapePodClient Client { get; }

        public event EventHandler<PluginDeploymentEventArgs> PluginDeploying;

        public string Host
        {
            get { return this._hostName; }
            set { this._hostName = value; this.Client.HostName = value; }
        }


        public EscapePodPlugin(string hostName = "")
        {
            this.Client = new EscapePodClient(hostName);
            if (!string.IsNullOrEmpty(hostName)) this.Host = Host;
        }

        public void Dispose()
        {
            this.Client?.Dispose();
        }


        private void SetDeploymentStep(string stepDescription)
        {
            this.DeploymentStepDescription = stepDescription;
            if (this.PluginDeploying != null)
                this.PluginDeploying(this, new PluginDeploymentEventArgs(stepDescription));
        }

        public async Task DeployPlugin(string host, string user, string pass)
        {
            if (!string.IsNullOrEmpty(host)) this.Host = host;
            if (string.IsNullOrEmpty(this.Host)) throw new ArgumentException("EscapePod host cannot be empty", host);

            try
            {
                if (this.Client.IsConnected)
                    await this.Client.Disconnect().ConfigureAwait(false);

                this.SetDeploymentStep("Connecting to EscapePod...");
                System.Reflection.Assembly assembly = typeof(EscapePodPlugin).Assembly;
                ConnectionInfo sshConnection = new ConnectionInfo(host, user, new PasswordAuthenticationMethod(user, pass));
                await this.ShellExectue(sshConnection, "").ConfigureAwait(false);

                //using (Stream targeFile = File.OpenRead("Resources/cybervector-proxy.conf"))
                using (Stream targetFile = assembly.GetManifestResourceStream("Cyb3rPod.Resources.cybervector-proxy.conf"))
                {
                    this.SetDeploymentStep("Uploading Config Files...");
                    string targetPath = "/usr/local/escapepod/bin/cybervector-proxy.conf";
                    await this.UploadFile(sshConnection, targetFile, targetPath).ConfigureAwait(false);
                }

                //using (Stream targetFile = File.OpenRead("Resources/cybervector-proxy.sh"))
                using (Stream targetFile = assembly.GetManifestResourceStream("Cyb3rPod.Resources.cybervector-proxy.sh"))
                {
                    this.SetDeploymentStep("Uploading Config Files...");
                    string targetPath = "/usr/local/escapepod/bin/cybervector-proxy.sh";
                    await this.UploadFile(sshConnection, targetFile, targetPath).ConfigureAwait(false);
                }

                using (Stream targetFile = assembly.GetManifestResourceStream("Cyb3rPod.Resources.cybervector-proxy.service"))
                //using (Stream targetFile = File.OpenRead("Resources/cybervector-proxy.service"))
                {
                    this.SetDeploymentStep("Uploading Proxy Service...");
                    string targetPath = "/usr/local/escapepod/bin/cybervector-proxy.service";
                    await this.UploadFile(sshConnection, targetFile, targetPath).ConfigureAwait(false);
                }

                using (Stream targetFile = assembly.GetManifestResourceStream("Cyb3rPod.Resources.cybervector-proxy"))
                //using (Stream targetFile = File.OpenRead("Resources/cybervector-proxy"))
                {
                    this.SetDeploymentStep("Uploading Cyb3rVector Extension...");

                    try { await this.ShellExectue(sshConnection, "sudo systemctl stop cybervector-proxy").ConfigureAwait(false); }
                    catch { /* supress */ }

                    string targetPath = "/usr/local/escapepod/bin/cybervector-proxy";
                    await this.UploadFile(sshConnection, targetFile, targetPath).ConfigureAwait(false);
                }

                this.SetDeploymentStep("Starting up Cyb3rVector Extension...");
                string startCommand = "chmod 755 /usr/local/escapepod/bin/cybervector-proxy.sh && /usr/local/escapepod/bin/cybervector-proxy.sh && rm -rf /usr/local/escapepod/bin/cybervector-proxy.sh";
                await this.ShellExectue(sshConnection, startCommand).ConfigureAwait(false);
            }
            catch (Exception error)
            {
                throw new EscapePodConnectionFailed(error.Message);
            }
        }

        public async Task RestartServices(string host, string user, string pass)
        {
            if (!string.IsNullOrEmpty(host)) this.Host = host;
            if (string.IsNullOrEmpty(this.Host)) throw new ArgumentException("EscapePod host cannot be empty", host);

            try
            {
                if (this.Client.IsConnected)
                    await this.Client.Disconnect().ConfigureAwait(false);

                this.SetDeploymentStep("Connecting to EscapePod...");
                System.Reflection.Assembly assembly = typeof(EscapePodPlugin).Assembly;
                ConnectionInfo sshConnection = new ConnectionInfo(host, user, new PasswordAuthenticationMethod(user, pass));
                await this.ShellExectue(sshConnection, "").ConfigureAwait(false);

                this.SetDeploymentStep("Restarting EscapePod Service...");
                await this.ShellExectue(sshConnection, "sudo systemctl restart escape_pod");

                this.SetDeploymentStep("Restarting Cyb3rVector Extension...");
                await this.ShellExectue(sshConnection, "sudo systemctl restart cybervector-proxy");
            }
            catch (Exception error)
            {
                throw new EscapePodConnectionFailed(error.Message);
            }
        }

        private Task<bool> UploadFile(ConnectionInfo connection, Stream file, string path, string description = "")
        {
            Task<bool> task = new Task<bool>(() =>
            {
                if (string.IsNullOrEmpty(description)) description = this.DeploymentStepDescription;
                using (ScpClient client = new ScpClient(connection))
                {
                    bool uploadResult = false;
                    client.Connect();
                    client.Uploading += (object sender, ScpUploadEventArgs args) =>
                    {
                        if ((this.PluginDeploying != null) && (args.Size != 0))
                            this.PluginDeploying(this, new PluginDeploymentEventArgs(description, (double)args.Uploaded / (double)args.Size * 100.0));
                        uploadResult = (args.Size == args.Uploaded);
                    };
                    client.Upload(file, path);
                    return uploadResult;
                }
            });
            task.Start();
            return task;
        }

        private Task<Stream> DownloadFile(ConnectionInfo connection, string path, string description = "")
        {
            Task<Stream> task = new Task<Stream>(() =>
            {
                if (string.IsNullOrEmpty(description)) description = this.DeploymentStepDescription;
                using (ScpClient client = new ScpClient(connection))
                {
                    bool uploadResult = false;
                    client.Connect();
                    client.Downloading += (object sender, ScpDownloadEventArgs args) =>
                    {
                        if ((this.PluginDeploying != null) && (args.Size != 0))
                            this.PluginDeploying(this, new PluginDeploymentEventArgs(description, (double)args.Downloaded / (double)args.Size * 100.0));
                        uploadResult = (args.Size == args.Downloaded);
                    };
                    MemoryStream outputStream = new MemoryStream();
                    client.Download(path, outputStream);
                    return outputStream;
                }
            });
            task.Start();
            return task;
        }

        private Task<string> ShellExectue(ConnectionInfo connection, string command, string description = "")
        {
            Task<string> task = new Task<string>(() =>
            {
                if (string.IsNullOrEmpty(description)) description = this.DeploymentStepDescription;
                using (SshClient ssh = new SshClient(connection))
                {
                    if (this.PluginDeploying != null)
                        this.PluginDeploying(this, new PluginDeploymentEventArgs(description));

                    ssh.Connect();

                    if (string.IsNullOrEmpty(command)) return null;

                    SshCommand sshCommand = ssh.CreateCommand(command);
                    sshCommand.Execute();
                    if (sshCommand.ExitStatus != 0)
                        throw new EscapePodException(sshCommand.Error);
                    else
                        return sshCommand.Result;
                }
            });
            task.Start();
            return task;
        }

        private string _hostName = "escapepod.local";
    }
}