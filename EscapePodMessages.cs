using Cybervector;
using Newtonsoft.Json.Linq;
using System;

namespace Cyb3rPod
{
    public class SubscribedMessage
    {
        public string Uuid;
        public string Message;

        public SubscribedMessage(string jsonMessage)
        {
            JObject json = JObject.Parse(jsonMessage);
            JToken uuid = json.SelectToken("uuid");
            this.Uuid = (uuid != null) ? uuid.ToString() : "";
            JToken message = json.SelectToken("uuid");
            this.Message = (message != null) ? message.ToString() : "";
        }
    }

    public class KeepAliveMessage
    {
        public DateTime LastKeepAlive;
        public KeepAliveMessage(string jsonMessage)
        {
            JObject json = JObject.Parse(jsonMessage);
            JToken token = json.SelectToken("time");
            string normalized = token.ToString().Substring(0, token.ToString().IndexOf(' ', 20));
            DateTime.TryParse(normalized, out this.LastKeepAlive); //json[key].ToString()
        }
    }

    public class StatusMessage
    {
        public string Version;
        public int Subscribed;

        public StatusMessage(StatusResponse response)
        {
            if (response == null) return;

            this.Version = response.Version;
            this.Subscribed = response.Subscribed;
        }
    }
}