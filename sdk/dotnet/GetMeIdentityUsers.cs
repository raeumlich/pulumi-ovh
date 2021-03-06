// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh
{
    public static class GetMeIdentityUsers
    {
        /// <summary>
        /// Use this data source to retrieve list of user logins of the account's identity users.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Ovh = Pulumi.Ovh;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var users = Output.Create(Ovh.GetMeIdentityUsers.InvokeAsync());
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetMeIdentityUsersResult> InvokeAsync(InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetMeIdentityUsersResult>("ovh:index/getMeIdentityUsers:getMeIdentityUsers", InvokeArgs.Empty, options.WithDefaults());
    }


    [OutputType]
    public sealed class GetMeIdentityUsersResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The list of the user's logins of all the identity users.
        /// </summary>
        public readonly ImmutableArray<string> Users;

        [OutputConstructor]
        private GetMeIdentityUsersResult(
            string id,

            ImmutableArray<string> users)
        {
            Id = id;
            Users = users;
        }
    }
}
