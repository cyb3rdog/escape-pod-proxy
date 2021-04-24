using System;

namespace Cyb3rPod
{
    public class EscapePodException : Exception
    {
        public EscapePodException(string message) : base(message) { }
        public EscapePodException(string message, Exception innerException) : base(message, innerException) { }
    }

    public class EscapePodConnectionFailed : EscapePodException
    {
        public EscapePodConnectionFailed(string message) : base(message) { }
        public EscapePodConnectionFailed(string message, Exception innerException) : base(message, innerException) { }
    }

    public class EscapePodCommandFailed : EscapePodException
    {
        public EscapePodCommandFailed(string message) : base(message) { }
        public EscapePodCommandFailed(string message, Exception innerException) : base(message, innerException) { }
    }
}