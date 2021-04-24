using Newtonsoft.Json.Linq;
using System;

namespace Cyb3rPod
{
    public class KeepAliveMessage
    {
        public DateTime LastKeepAlive;
        public KeepAliveMessage(string jsonMessage)
        {
            JObject json = JObject.Parse(jsonMessage);
            JToken token = json.SelectToken("time");
            //return json.Value<JObject>(key).ToString();
            string normalized = token.ToString().Substring(0, token.ToString().IndexOf(' ', 20));
            DateTime.TryParse(normalized, out this.LastKeepAlive); //json[key].ToString()
        }
    }
}