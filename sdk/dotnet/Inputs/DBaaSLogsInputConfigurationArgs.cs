// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.Inputs
{

    public sealed class DBaaSLogsInputConfigurationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Flowgger configuration
        /// </summary>
        [Input("flowgger")]
        public Input<Inputs.DBaaSLogsInputConfigurationFlowggerArgs>? Flowgger { get; set; }

        /// <summary>
        /// Logstash configuration
        /// </summary>
        [Input("logstash")]
        public Input<Inputs.DBaaSLogsInputConfigurationLogstashArgs>? Logstash { get; set; }

        public DBaaSLogsInputConfigurationArgs()
        {
        }
    }
}
