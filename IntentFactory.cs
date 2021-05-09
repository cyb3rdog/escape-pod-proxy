using System;
using System.Collections.Generic;
using System.Threading.Tasks;

namespace Cyb3rPod
{
    public class IntentFactory : Dictionary<string, string>, IDisposable
    {
        private EscapePodPlugin Plugin = null;
        public List<EscapePodIntent> EscapePodIntents { get; private set; }

        public IntentFactory(EscapePodPlugin escapePodPlugin)
        {
            this.Plugin = escapePodPlugin;
        }
        ~IntentFactory()
        {
            this.Dispose();
        }

        public async Task Initialze(string hostName = "")
        {
            try
            {
                if (!string.IsNullOrEmpty(hostName))
                    this.Plugin.Host = hostName;

                if (this.Plugin.IsEnabled)
                {
                    if (!this.Plugin.Client.IsConnected)
                        await this.Plugin.Client.Connect().ConfigureAwait(false);
                    this.EscapePodIntents = await this.SelectIntents().ConfigureAwait(false);
                }
            }
            catch { /* supress */ }
        }

        public async Task AddIntent(string blockId, string keywords)
        {
            EscapePodIntent newIntent = new EscapePodIntent(blockId, "intent_" + blockId, keywords);
            string intentOid = await this.Plugin.Client.InsertIntent(newIntent).ConfigureAwait(false);
            this.Add(blockId, intentOid);
        }

        public async Task RemoveIntent(string blockId)
        {
            if (!this.ContainsKey(blockId)) return;
            await this.Plugin.Client.DeleteIntent(this[blockId]).ConfigureAwait(false);
            this.Remove(blockId);
        }

        public async Task ClearIntents()
        {
            foreach (string intentOid in this.Values)
                try { await this.Plugin.Client.DeleteIntent(intentOid).ConfigureAwait(false); }
                catch { /* suppress */ }
        }

        public async Task<List<EscapePodIntent>> SelectIntents(string filter = "{}")
        {
            return await this.Plugin.Client.SelectIntents(filter).ConfigureAwait(false);
        }

        public void Dispose()
        {
            _ = this.ClearIntents().ConfigureAwait(false);
            this.Clear();
        }
    }
}
