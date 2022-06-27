// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.Outputs
{

    [OutputType]
    public sealed class DBaaSLogsInputConfiguration
    {
        /// <summary>
        /// Flowgger configuration
        /// </summary>
        public readonly Outputs.DBaaSLogsInputConfigurationFlowgger? Flowgger;
        /// <summary>
        /// Logstash configuration
        /// </summary>
        public readonly Outputs.DBaaSLogsInputConfigurationLogstash? Logstash;

        [OutputConstructor]
        private DBaaSLogsInputConfiguration(
            Outputs.DBaaSLogsInputConfigurationFlowgger? flowgger,

            Outputs.DBaaSLogsInputConfigurationLogstash? logstash)
        {
            Flowgger = flowgger;
            Logstash = logstash;
        }
    }
}