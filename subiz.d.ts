import * as $protobuf from "protobufjs";

export namespace account {

    interface IAgent {
        ctx?: (common.IContext|null);
        id?: (string|null);
        account_id?: (string|null);
        fullname?: (string|null);
        email?: (string|null);
        emails?: (string[]|null);
        phones?: (string[]|null);
        phone?: (string|null);
        is_owner?: (boolean|null);
        job_title?: (string|null);
        gender?: (string|null);
        avatar_url?: (string|null);
        lang?: (string|null);
        location?: (string|null);
        timezone?: (string|null);
        encrypted_password?: (string|null);
        joined?: (number|Long|null);
        invited_by?: (string|null);
        state?: (string|null);
        password_changed?: (number|Long|null);
        seen?: (number|Long|null);
        modified?: (number|Long|null);
        method?: (auth.IMethod|null);
        account?: (account.IAccount|null);
        country_code?: (string|null);
        v3_state?: (number|null);
        v3_hashed_password?: (string|null);
    }

    class Agent implements IAgent {
        constructor(p?: account.IAgent);
        public ctx?: (common.IContext|null);
        public id: string;
        public account_id: string;
        public fullname: string;
        public email: string;
        public emails: string[];
        public phones: string[];
        public phone: string;
        public is_owner: boolean;
        public job_title: string;
        public gender: string;
        public avatar_url: string;
        public lang: string;
        public location: string;
        public timezone: string;
        public encrypted_password: string;
        public joined: (number|Long);
        public invited_by: string;
        public state: string;
        public password_changed: (number|Long);
        public seen: (number|Long);
        public modified: (number|Long);
        public method?: (auth.IMethod|null);
        public account?: (account.IAccount|null);
        public country_code: string;
        public v3_state: number;
        public v3_hashed_password: string;
    }

    namespace Agent {

        enum Gender {
            unset = 0,
            male = 1,
            female = 2,
            bisexual = 3,
            asexual = 4
        }

        enum AgentState {
            pending = 0,
            active = 1,
            inactive = 2,
            deleted = 3
        }
    }

    interface IInvitation {
        ctx?: (common.IContext|null);
        id?: (string|null);
        account_id?: (string|null);
        from_id?: (string|null);
        email?: (string|null);
        agent_id?: (string|null);
        sent?: (number|Long|null);
        replied?: (number|Long|null);
        agent_fullname?: (string|null);
        agent_job_title?: (string|null);
        token?: (string|null);
    }

    class Invitation implements IInvitation {
        constructor(p?: account.IInvitation);
        public ctx?: (common.IContext|null);
        public id: string;
        public account_id: string;
        public from_id: string;
        public email: string;
        public agent_id: string;
        public sent: (number|Long);
        public replied: (number|Long);
        public agent_fullname: string;
        public agent_job_title: string;
        public token: string;
    }

    interface IAgentGroup {
        ctx?: (common.IContext|null);
        id?: (string|null);
        account_id?: (string|null);
        name?: (string|null);
        logo_url?: (string|null);
        members?: (account.IAgent[]|null);
        created?: (number|Long|null);
        modified?: (number|Long|null);
        members_count?: (number|null);
    }

    class AgentGroup implements IAgentGroup {
        constructor(p?: account.IAgentGroup);
        public ctx?: (common.IContext|null);
        public id: string;
        public account_id: string;
        public name: string;
        public logo_url: string;
        public members: account.IAgent[];
        public created: (number|Long);
        public modified: (number|Long);
        public members_count: number;
    }

    interface IResetPasswordRequest {
        ctx?: (common.IContext|null);
        email?: (string|null);
    }

    class ResetPasswordRequest implements IResetPasswordRequest {
        constructor(p?: account.IResetPasswordRequest);
        public ctx?: (common.IContext|null);
        public email: string;
    }

    interface IAgentPerm {
        ctx?: (common.IContext|null);
        account_id?: (string|null);
        agent_id?: (string|null);
        method?: (auth.IMethod|null);
    }

    class AgentPerm implements IAgentPerm {
        constructor(p?: account.IAgentPerm);
        public ctx?: (common.IContext|null);
        public account_id: string;
        public agent_id: string;
        public method?: (auth.IMethod|null);
    }

    interface IAccount {
        ctx?: (common.IContext|null);
        id?: (string|null);
        name?: (string|null);
        logo_url?: (string|null);
        owner_id?: (string|null);
        state?: (string|null);
        created?: (number|Long|null);
        confirmed?: (number|Long|null);
        modified?: (number|Long|null);
        slogan?: (string|null);
        referer_id?: (string|null);
        city?: (string|null);
        zip_code?: (number|null);
        tax_id?: (string|null);
        facebook?: (string|null);
        twitter?: (string|null);
        phone?: (string|null);
        address?: (string|null);
        url?: (string|null);
        lang?: (string|null);
        referer_from?: (string|null);
        timezone?: (string|null);
        limit?: (payment.ILimit|null);
        country?: (string|null);
        v3_state?: (number|null);
    }

    class Account implements IAccount {
        constructor(p?: account.IAccount);
        public ctx?: (common.IContext|null);
        public id: string;
        public name: string;
        public logo_url: string;
        public owner_id: string;
        public state: string;
        public created: (number|Long);
        public confirmed: (number|Long);
        public modified: (number|Long);
        public slogan: string;
        public referer_id: string;
        public city: string;
        public zip_code: number;
        public tax_id: string;
        public facebook: string;
        public twitter: string;
        public phone: string;
        public address: string;
        public url: string;
        public lang: string;
        public referer_from: string;
        public timezone: string;
        public limit?: (payment.ILimit|null);
        public country: string;
        public v3_state: number;
    }

    namespace Account {

        enum State {
            pending = 0,
            activated = 1,
            locked = 2,
            deleted = 3
        }
    }

    interface IGroupMember {
        ctx?: (common.IContext|null);
        account_id?: (string|null);
        group_id?: (string|null);
        agent_id?: (string|null);
    }

    class GroupMember implements IGroupMember {
        constructor(p?: account.IGroupMember);
        public ctx?: (common.IContext|null);
        public account_id: string;
        public group_id: string;
        public agent_id: string;
    }

    interface ICreateAccountRequest {
        ctx?: (common.IContext|null);
        fullname?: (string|null);
        email?: (string|null);
        lang?: (string|null);
        password?: (string|null);
        account_name?: (string|null);
        account_url?: (string|null);
        avatar_url?: (string|null);
        timezone?: (string|null);
        phone?: (string|null);
        country_code?: (string|null);
    }

    class CreateAccountRequest implements ICreateAccountRequest {
        constructor(p?: account.ICreateAccountRequest);
        public ctx?: (common.IContext|null);
        public fullname: string;
        public email: string;
        public lang: string;
        public password: string;
        public account_name: string;
        public account_url: string;
        public avatar_url: string;
        public timezone: string;
        public phone: string;
        public country_code: string;
    }

    interface ILoginRequest {
        ctx?: (common.IContext|null);
        email?: (string|null);
        password?: (string|null);
    }

    class LoginRequest implements ILoginRequest {
        constructor(p?: account.ILoginRequest);
        public ctx?: (common.IContext|null);
        public email: string;
        public password: string;
    }

    interface IAgents {
        ctx?: (common.IContext|null);
        Agents?: (account.IAgent[]|null);
    }

    class Agents implements IAgents {
        constructor(p?: account.IAgents);
        public ctx?: (common.IContext|null);
        public Agents: account.IAgent[];
    }

    interface INewPassword {
        ctx?: (common.IContext|null);
        token?: (string|null);
        new_password?: (string|null);
        old_password?: (string|null);
        email?: (string|null);
    }

    class NewPassword implements INewPassword {
        constructor(p?: account.INewPassword);
        public ctx?: (common.IContext|null);
        public token: string;
        public new_password: string;
        public old_password: string;
        public email: string;
    }

    interface IAgentGroups {
        ctx?: (common.IContext|null);
        Groups?: (account.IAgentGroup[]|null);
    }

    class AgentGroups implements IAgentGroups {
        constructor(p?: account.IAgentGroups);
        public ctx?: (common.IContext|null);
        public Groups: account.IAgentGroup[];
    }

    interface IToken {
        ctx?: (common.IContext|null);
        is_set?: (boolean|null);
        token?: (string|null);
        account_id?: (string|null);
    }

    class Token implements IToken {
        constructor(p?: account.IToken);
        public ctx?: (common.IContext|null);
        public is_set: boolean;
        public token: string;
        public account_id: string;
    }

    interface IAccountV3 {
        ctx?: (common.IContext|null);
        account?: (account.IAccount|null);
        owner?: (account.IAgent|null);
        subscription?: (payment.ISubscription|null);
    }

    class AccountV3 implements IAccountV3 {
        constructor(p?: account.IAccountV3);
        public ctx?: (common.IContext|null);
        public account?: (account.IAccount|null);
        public owner?: (account.IAgent|null);
        public subscription?: (payment.ISubscription|null);
    }

    enum Event {
        AccountRequested = 1000,
        AccountSynced = 1001,
        AccountV3Synced = 1002,
        AgentGroupDeleted = 0,
        AgentLeftGroup = 1,
        AgentJoinedGroup = 2,
        AgentGroupUpserted = 4,
        AgentUpserted = 6,
        AgentPermissionUpdated = 9,
        AccountUpserted = 14,
        AccountPlanUpdated = 16,
        AccountStateUpdated = 17,
        AccountConfirmRequest = 19,
        HandleExpiredInvitation = 20,
        AccountConfirmRequested = 21,
        AccountConfirmSuccessEmailRequested = 25,
        AccountResetPasswordEmail = 33,
        AccountPasswordChangedEmailRequested = 34,
        AccountInviteEmail = 22,
        AccountDeleted = 24,
        AccountCreated = 45,
        AccountActivated = 46,
        AccountInfoUpdated = 47,
        AgentRejected = 10,
        AgentAccepted = 11,
        AgentInvited = 13,
        AgentDeleted = 15,
        AgentActivated = 50,
        AgentDeactivated = 51,
        AgentInfoUpdated = 59,
        AccountV3Created = 60,
        AccountPaymentV3Synced = 63
    }

    interface IConfirmEmail {
        ctx?: (common.IContext|null);
        to?: (string|null);
        account_id?: (string|null);
        owner_id?: (string|null);
        token?: (string|null);
        expired_in?: (number|Long|null);
        account_name?: (string|null);
        lang?: (lang.L|null);
        owner_name?: (string|null);
        from?: (string|null);
    }

    class ConfirmEmail implements IConfirmEmail {
        constructor(p?: account.IConfirmEmail);
        public ctx?: (common.IContext|null);
        public to: string;
        public account_id: string;
        public owner_id: string;
        public token: string;
        public expired_in: (number|Long);
        public account_name: string;
        public lang: lang.L;
        public owner_name: string;
        public from: string;
    }

    interface IInviteEmail {
        ctx?: (common.IContext|null);
        to?: (string|null);
        account_id?: (string|null);
        sender_id?: (string|null);
        expired_in?: (number|Long|null);
        token?: (string|null);
        lang?: (lang.L|null);
        fullname?: (string|null);
        sender_name?: (string|null);
        account_name?: (string|null);
        account_logo_url?: (string|null);
        from?: (string|null);
    }

    class InviteEmail implements IInviteEmail {
        constructor(p?: account.IInviteEmail);
        public ctx?: (common.IContext|null);
        public to: string;
        public account_id: string;
        public sender_id: string;
        public expired_in: (number|Long);
        public token: string;
        public lang: lang.L;
        public fullname: string;
        public sender_name: string;
        public account_name: string;
        public account_logo_url: string;
        public from: string;
    }

    interface IResetPasswordEmail {
        ctx?: (common.IContext|null);
        from?: (string|null);
        to?: (string|null);
        expired_in?: (number|Long|null);
        token?: (string|null);
        account_id?: (string|null);
        agent_id?: (string|null);
        agent_name?: (string|null);
        lang?: (lang.L|null);
    }

    class ResetPasswordEmail implements IResetPasswordEmail {
        constructor(p?: account.IResetPasswordEmail);
        public ctx?: (common.IContext|null);
        public from: string;
        public to: string;
        public expired_in: (number|Long);
        public token: string;
        public account_id: string;
        public agent_id: string;
        public agent_name: string;
        public lang: lang.L;
    }

    interface IPasswordChangedEmail {
        ctx?: (common.IContext|null);
        to?: (string|null);
        account_id?: (string|null);
        agent_id?: (string|null);
        agent_name?: (string|null);
        lang?: (lang.L|null);
        from?: (string|null);
    }

    class PasswordChangedEmail implements IPasswordChangedEmail {
        constructor(p?: account.IPasswordChangedEmail);
        public ctx?: (common.IContext|null);
        public to: string;
        public account_id: string;
        public agent_id: string;
        public agent_name: string;
        public lang: lang.L;
        public from: string;
    }

    interface IAccountConfirmSuccessEmail {
        ctx?: (common.IContext|null);
        to?: (string|null);
        account_id?: (string|null);
        agent_id?: (string|null);
        agent_name?: (string|null);
        lang?: (lang.L|null);
        from?: (string|null);
        from_time?: (number|Long|null);
        to_time?: (number|Long|null);
    }

    class AccountConfirmSuccessEmail implements IAccountConfirmSuccessEmail {
        constructor(p?: account.IAccountConfirmSuccessEmail);
        public ctx?: (common.IContext|null);
        public to: string;
        public account_id: string;
        public agent_id: string;
        public agent_name: string;
        public lang: lang.L;
        public from: string;
        public from_time: (number|Long);
        public to_time: (number|Long);
    }

    class AccountMgr extends $protobuf.rpc.Service {
        constructor(rpcImpl: $protobuf.RPCImpl, requestDelimited?: boolean, responseDelimited?: boolean);
        public createGroup(request: account.IAgentGroup, callback: account.AccountMgr.CreateGroupCallback): void;
        public createGroup(request: account.IAgentGroup): Promise<account.AgentGroup>;
        public updateGroup(request: account.IAgentGroup, callback: account.AccountMgr.UpdateGroupCallback): void;
        public updateGroup(request: account.IAgentGroup): Promise<account.AgentGroup>;
        public getGroup(request: common.IId, callback: account.AccountMgr.GetGroupCallback): void;
        public getGroup(request: common.IId): Promise<account.AgentGroup>;
        public getAgentPermission(request: common.IId, callback: account.AccountMgr.GetAgentPermissionCallback): void;
        public getAgentPermission(request: common.IId): Promise<auth.Method>;
        public updateAgentPermission(request: account.IAgentPerm, callback: account.AccountMgr.UpdateAgentPermissionCallback): void;
        public updateAgentPermission(request: account.IAgentPerm): Promise<auth.Method>;
        public requestResetPassword(request: account.IResetPasswordRequest, callback: account.AccountMgr.RequestResetPasswordCallback): void;
        public requestResetPassword(request: account.IResetPasswordRequest): Promise<common.Empty>;
        public updatePassword(request: account.INewPassword, callback: account.AccountMgr.UpdatePasswordCallback): void;
        public updatePassword(request: account.INewPassword): Promise<account.Agent>;
        public updateAgent(request: account.IAgent, callback: account.AccountMgr.UpdateAgentCallback): void;
        public updateAgent(request: account.IAgent): Promise<account.Agent>;
        public deleteAgent(request: common.IId, callback: account.AccountMgr.DeleteAgentCallback): void;
        public deleteAgent(request: common.IId): Promise<account.Agent>;
        public acceptInvitation(request: account.INewPassword, callback: account.AccountMgr.AcceptInvitationCallback): void;
        public acceptInvitation(request: account.INewPassword): Promise<account.Agent>;
        public getInvitation(request: account.IToken, callback: account.AccountMgr.GetInvitationCallback): void;
        public getInvitation(request: account.IToken): Promise<account.Agent>;
        public inviteAgent(request: account.IAgent, callback: account.AccountMgr.InviteAgentCallback): void;
        public inviteAgent(request: account.IAgent): Promise<account.Agent>;
        public getAgent(request: common.IId, callback: account.AccountMgr.GetAgentCallback): void;
        public getAgent(request: common.IId): Promise<account.Agent>;
        public confirmAccount(request: account.IToken, callback: account.AccountMgr.ConfirmAccountCallback): void;
        public confirmAccount(request: account.IToken): Promise<account.Account>;
        public getAccount(request: common.IId, callback: account.AccountMgr.GetAccountCallback): void;
        public getAccount(request: common.IId): Promise<account.Account>;
        public updateAccount(request: account.IAccount, callback: account.AccountMgr.UpdateAccountCallback): void;
        public updateAccount(request: account.IAccount): Promise<account.Account>;
        public createAccount(request: account.ICreateAccountRequest, callback: account.AccountMgr.CreateAccountCallback): void;
        public createAccount(request: account.ICreateAccountRequest): Promise<account.Account>;
        public listAgents(request: common.IId, callback: account.AccountMgr.ListAgentsCallback): void;
        public listAgents(request: common.IId): Promise<account.Agents>;
        public checkLogin(request: account.ILoginRequest, callback: account.AccountMgr.CheckLoginCallback): void;
        public checkLogin(request: account.ILoginRequest): Promise<common.Id>;
        public deleteGroup(request: common.IId, callback: account.AccountMgr.DeleteGroupCallback): void;
        public deleteGroup(request: common.IId): Promise<common.Empty>;
        public listGroups(request: common.IId, callback: account.AccountMgr.ListGroupsCallback): void;
        public listGroups(request: common.IId): Promise<account.AgentGroups>;
        public addAgentToGroup(request: account.IGroupMember, callback: account.AccountMgr.AddAgentToGroupCallback): void;
        public addAgentToGroup(request: account.IGroupMember): Promise<common.Empty>;
        public removeAgentFromGroup(request: account.IGroupMember, callback: account.AccountMgr.RemoveAgentFromGroupCallback): void;
        public removeAgentFromGroup(request: account.IGroupMember): Promise<common.Empty>;
    }

    namespace AccountMgr {

        type CreateGroupCallback = (error: (Error|null), response?: account.AgentGroup) => void;

        type UpdateGroupCallback = (error: (Error|null), response?: account.AgentGroup) => void;

        type GetGroupCallback = (error: (Error|null), response?: account.AgentGroup) => void;

        type GetAgentPermissionCallback = (error: (Error|null), response?: auth.Method) => void;

        type UpdateAgentPermissionCallback = (error: (Error|null), response?: auth.Method) => void;

        type RequestResetPasswordCallback = (error: (Error|null), response?: common.Empty) => void;

        type UpdatePasswordCallback = (error: (Error|null), response?: account.Agent) => void;

        type UpdateAgentCallback = (error: (Error|null), response?: account.Agent) => void;

        type DeleteAgentCallback = (error: (Error|null), response?: account.Agent) => void;

        type AcceptInvitationCallback = (error: (Error|null), response?: account.Agent) => void;

        type GetInvitationCallback = (error: (Error|null), response?: account.Agent) => void;

        type InviteAgentCallback = (error: (Error|null), response?: account.Agent) => void;

        type GetAgentCallback = (error: (Error|null), response?: account.Agent) => void;

        type ConfirmAccountCallback = (error: (Error|null), response?: account.Account) => void;

        type GetAccountCallback = (error: (Error|null), response?: account.Account) => void;

        type UpdateAccountCallback = (error: (Error|null), response?: account.Account) => void;

        type CreateAccountCallback = (error: (Error|null), response?: account.Account) => void;

        type ListAgentsCallback = (error: (Error|null), response?: account.Agents) => void;

        type CheckLoginCallback = (error: (Error|null), response?: common.Id) => void;

        type DeleteGroupCallback = (error: (Error|null), response?: common.Empty) => void;

        type ListGroupsCallback = (error: (Error|null), response?: account.AgentGroups) => void;

        type AddAgentToGroupCallback = (error: (Error|null), response?: common.Empty) => void;

        type RemoveAgentFromGroupCallback = (error: (Error|null), response?: common.Empty) => void;
    }
}

export namespace api {

    interface IEmpty {
    }

    class Empty implements IEmpty {
        constructor(p?: api.IEmpty);
    }
}

export namespace auth {

    interface ICredential {
        account_id?: (string|null);
        issuer?: (string|null);
        type?: (auth.Type|null);
        method?: (auth.IMethod|null);
        client_id?: (string|null);
        client_type?: (auth.Type|null);
        client_account_id?: (string|null);
        scopes?: (string[]|null);
        avatar_url?: (string|null);
        name?: (string|null);
        email?: (string|null);
    }

    class Credential implements ICredential {
        constructor(p?: auth.ICredential);
        public account_id: string;
        public issuer: string;
        public type: auth.Type;
        public method?: (auth.IMethod|null);
        public client_id: string;
        public client_type: auth.Type;
        public client_account_id: string;
        public scopes: string[];
        public avatar_url: string;
        public name: string;
        public email: string;
    }

    interface IScope {
        id?: (string|null);
        name?: (string|null);
        logo_url?: (string|null);
        title?: (string|null);
        description?: (string|null);
        method?: (auth.IMethod|null);
        event?: (string[]|null);
    }

    class Scope implements IScope {
        constructor(p?: auth.IScope);
        public id: string;
        public name: string;
        public logo_url: string;
        public title: string;
        public description: string;
        public method?: (auth.IMethod|null);
        public event: string[];
    }

    interface IMethod {
        ping?: (boolean|null);
        update_trigger?: (boolean|null);
        delete_trigger?: (boolean|null);
        create_trigger?: (boolean|null);
        read_trigger?: (boolean|null);
        read_segmentation?: (boolean|null);
        update_segmentation?: (boolean|null);
        delete_segmentation?: (boolean|null);
        create_segmentation?: (boolean|null);
        invite_agent?: (boolean|null);
        update_agent?: (boolean|null);
        update_agents?: (boolean|null);
        read_agent?: (boolean|null);
        read_agents?: (boolean|null);
        reset_password?: (boolean|null);
        update_agents_permission?: (boolean|null);
        read_agent_permission?: (boolean|null);
        update_agents_state?: (boolean|null);
        read_account?: (boolean|null);
        create_agent_group?: (boolean|null);
        delete_agent_group?: (boolean|null);
        read_agent_group?: (boolean|null);
        update_agent_group?: (boolean|null);
        update_plan?: (boolean|null);
        update_account_infomation?: (boolean|null);
        read_client?: (boolean|null);
        update_client?: (boolean|null);
        delete_client?: (boolean|null);
        create_client?: (boolean|null);
        read_rule?: (boolean|null);
        create_rule?: (boolean|null);
        delete_rule?: (boolean|null);
        update_rule?: (boolean|null);
        start_conversation?: (boolean|null);
        read_conversation?: (boolean|null);
        export_conversations?: (boolean|null);
        read_teammates_conversations?: (boolean|null);
        send_message?: (boolean|null);
        integrate_connector?: (boolean|null);
        read_user_email?: (boolean|null);
        read_user_facebook_id?: (boolean|null);
        read_user_phones?: (boolean|null);
        read_user_widget_setting?: (boolean|null);
        read_tag?: (boolean|null);
        update_tag?: (boolean|null);
        delete_tag?: (boolean|null);
        update_widget_setting?: (boolean|null);
        create_whitelist_domain?: (boolean|null);
        create_whitelist_ip?: (boolean|null);
        create_whitelist_user?: (boolean|null);
        delete_whitelist_domain?: (boolean|null);
        delete_whitelist_ip?: (boolean|null);
        delete_whitelist_user?: (boolean|null);
        read_whitelist_ip?: (boolean|null);
        read_whitelist_user?: (boolean|null);
        read_whitelist_domain?: (boolean|null);
        purchase_service?: (boolean|null);
        update_payment_method?: (boolean|null);
        add_credit?: (boolean|null);
        update_billing_cycle?: (boolean|null);
        read_invoices?: (boolean|null);
        read_subscriptions?: (boolean|null);
        read_attribute?: (boolean|null);
        create_attribute?: (boolean|null);
        update_attribute?: (boolean|null);
        delete_attribute?: (boolean|null);
    }

    class Method implements IMethod {
        constructor(p?: auth.IMethod);
        public ping: boolean;
        public update_trigger: boolean;
        public delete_trigger: boolean;
        public create_trigger: boolean;
        public read_trigger: boolean;
        public read_segmentation: boolean;
        public update_segmentation: boolean;
        public delete_segmentation: boolean;
        public create_segmentation: boolean;
        public invite_agent: boolean;
        public update_agent: boolean;
        public update_agents: boolean;
        public read_agent: boolean;
        public read_agents: boolean;
        public reset_password: boolean;
        public update_agents_permission: boolean;
        public read_agent_permission: boolean;
        public update_agents_state: boolean;
        public read_account: boolean;
        public create_agent_group: boolean;
        public delete_agent_group: boolean;
        public read_agent_group: boolean;
        public update_agent_group: boolean;
        public update_plan: boolean;
        public update_account_infomation: boolean;
        public read_client: boolean;
        public update_client: boolean;
        public delete_client: boolean;
        public create_client: boolean;
        public read_rule: boolean;
        public create_rule: boolean;
        public delete_rule: boolean;
        public update_rule: boolean;
        public start_conversation: boolean;
        public read_conversation: boolean;
        public export_conversations: boolean;
        public read_teammates_conversations: boolean;
        public send_message: boolean;
        public integrate_connector: boolean;
        public read_user_email: boolean;
        public read_user_facebook_id: boolean;
        public read_user_phones: boolean;
        public read_user_widget_setting: boolean;
        public read_tag: boolean;
        public update_tag: boolean;
        public delete_tag: boolean;
        public update_widget_setting: boolean;
        public create_whitelist_domain: boolean;
        public create_whitelist_ip: boolean;
        public create_whitelist_user: boolean;
        public delete_whitelist_domain: boolean;
        public delete_whitelist_ip: boolean;
        public delete_whitelist_user: boolean;
        public read_whitelist_ip: boolean;
        public read_whitelist_user: boolean;
        public read_whitelist_domain: boolean;
        public purchase_service: boolean;
        public update_payment_method: boolean;
        public add_credit: boolean;
        public update_billing_cycle: boolean;
        public read_invoices: boolean;
        public read_subscriptions: boolean;
        public read_attribute: boolean;
        public create_attribute: boolean;
        public update_attribute: boolean;
        public delete_attribute: boolean;
    }

    interface IUserAuth {
        user_id?: (string|null);
        method?: (auth.IMethod|null);
    }

    class UserAuth implements IUserAuth {
        constructor(p?: auth.IUserAuth);
        public user_id: string;
        public method?: (auth.IMethod|null);
    }

    interface IPasswordCredential {
        username?: (string|null);
        password?: (string|null);
    }

    class PasswordCredential implements IPasswordCredential {
        constructor(p?: auth.IPasswordCredential);
        public username: string;
        public password: string;
    }

    interface ISuperPasswordCredential {
        subiz_username?: (string|null);
        subiz_password?: (string|null);
        subiz_token?: (string|null);
        issuer_id?: (string|null);
        account_id?: (string|null);
        expired?: (number|null);
        scopes?: (string[]|null);
    }

    class SuperPasswordCredential implements ISuperPasswordCredential {
        constructor(p?: auth.ISuperPasswordCredential);
        public subiz_username: string;
        public subiz_password: string;
        public subiz_token: string;
        public issuer_id: string;
        public account_id: string;
        public expired: number;
        public scopes: string[];
    }

    interface IAuthCookie {
        user_id?: (string|null);
        account_id?: (string|null);
        expired?: (number|null);
        issued?: (number|null);
        type?: (string|null);
    }

    class AuthCookie implements IAuthCookie {
        constructor(p?: auth.IAuthCookie);
        public user_id: string;
        public account_id: string;
        public expired: number;
        public issued: number;
        public type: string;
    }

    interface IOauthAccessToken {
        access_token?: (string|null);
        token_type?: (string|null);
        expires_in?: (number|Long|null);
        refresh_token?: (string|null);
        scope?: (string|null);
        state?: (string|null);
        error?: (string|null);
        error_description?: (string|null);
        error_uri?: (string|null);
    }

    class OauthAccessToken implements IOauthAccessToken {
        constructor(p?: auth.IOauthAccessToken);
        public access_token: string;
        public token_type: string;
        public expires_in: (number|Long);
        public refresh_token: string;
        public scope: string;
        public state: string;
        public error: string;
        public error_description: string;
        public error_uri: string;
    }

    interface IAccessToken {
        issuer_id?: (string|null);
        issuer_type?: (string|null);
        client_id?: (string|null);
        client_type?: (string|null);
        account_id?: (string|null);
        client_account_id?: (string|null);
        expired?: (number|null);
        scopes?: (string[]|null);
    }

    class AccessToken implements IAccessToken {
        constructor(p?: auth.IAccessToken);
        public issuer_id: string;
        public issuer_type: string;
        public client_id: string;
        public client_type: string;
        public account_id: string;
        public client_account_id: string;
        public expired: number;
        public scopes: string[];
    }

    interface ICookieExpire {
        user_id?: (string|null);
        account_id?: (string|null);
        expired_token?: (string|null);
        before_time?: (number|Long|null);
        except_token?: (string|null);
    }

    class CookieExpire implements ICookieExpire {
        constructor(p?: auth.ICookieExpire);
        public user_id: string;
        public account_id: string;
        public expired_token: string;
        public before_time: (number|Long);
        public except_token: string;
    }

    enum Type {
        unknown = 0,
        user = 1,
        agent = 2,
        subiz = 3,
        app = 5,
        rest = 8,
        connector = 6,
        bot = 7,
        dashboard = 10
    }

    interface IUser {
        id?: (string|null);
        account_id?: (string|null);
        email?: (string|null);
        encrypted_password?: (string|null);
        is_active?: (boolean|null);
        upserted?: (number|Long|null);
    }

    class User implements IUser {
        constructor(p?: auth.IUser);
        public id: string;
        public account_id: string;
        public email: string;
        public encrypted_password: string;
        public is_active: boolean;
        public upserted: (number|Long);
    }

    enum AuthorizationType {
        client_auth = 0,
        channel_auth = 1
    }

    enum Event {
        AUTH = 0,
        AuthExpireCookie = 4
    }
}

export namespace client {

    interface IAllType {
        client?: (client.IClients|null);
    }

    class AllType implements IAllType {
        constructor(p?: client.IAllType);
        public client?: (client.IClients|null);
    }

    interface IClients {
        ctx?: (common.IContext|null);
        clients?: (client.IClient[]|null);
    }

    class Clients implements IClients {
        constructor(p?: client.IClients);
        public ctx?: (common.IContext|null);
        public clients: client.IClient[];
    }

    interface IClient {
        ctx?: (common.IContext|null);
        id?: (string|null);
        secret?: (string|null);
        logo_url?: (string|null);
        account_id?: (string|null);
        is_verified?: (boolean|null);
        verified?: (number|Long|null);
        redirect_url?: (string|null);
        type?: (string|null);
        name?: (string|null);
        version?: (string|null);
        is_enabled?: (boolean|null);
        created?: (number|Long|null);
        modified?: (number|Long|null);
        webhook_url?: (string|null);
        events?: (string[]|null);
        channel_type?: (string|null);
        availability_url?: (string|null);
        ping_url?: (string|null);
    }

    class Client implements IClient {
        constructor(p?: client.IClient);
        public ctx?: (common.IContext|null);
        public id: string;
        public secret: string;
        public logo_url: string;
        public account_id: string;
        public is_verified: boolean;
        public verified: (number|Long);
        public redirect_url: string;
        public type: string;
        public name: string;
        public version: string;
        public is_enabled: boolean;
        public created: (number|Long);
        public modified: (number|Long);
        public webhook_url: string;
        public events: string[];
        public channel_type: string;
        public availability_url: string;
        public ping_url: string;
    }

    namespace Client {

        enum Type {
            app = 0,
            connector = 1,
            bot = 3
        }

        enum ChannelType {
            subiz = 0,
            email = 1,
            facebook = 2,
            viber = 3
        }
    }

    enum Event {
        ClientCreateRequested = 20,
        ClientUpdateRequested = 21,
        ClientDeleteRequested = 22,
        ClientReadRequested = 23,
        ClientListRequested = 24,
        ClientUpserted = 25,
        ClientDeleted = 26,
        ClientAuthorized = 12,
        ClientRequested = 30,
        ClientSynced = 31,
        ConnectorAuthorized = 13
    }

    interface IAuthorization {
        ctx?: (common.IContext|null);
        account_id?: (string|null);
        issuer?: (string|null);
        type?: (auth.Type|null);
        method?: (auth.IMethod|null);
        client_id?: (string|null);
        client_type?: (auth.Type|null);
        client_account_id?: (string|null);
        scopes?: (string[]|null);
        integration_id?: (string|null);
    }

    class Authorization implements IAuthorization {
        constructor(p?: client.IAuthorization);
        public ctx?: (common.IContext|null);
        public account_id: string;
        public issuer: string;
        public type: auth.Type;
        public method?: (auth.IMethod|null);
        public client_id: string;
        public client_type: auth.Type;
        public client_account_id: string;
        public scopes: string[];
        public integration_id: string;
    }
}

export namespace common {

    interface IEmpty {
        ctx?: (common.IContext|null);
    }

    class Empty implements IEmpty {
        constructor(p?: common.IEmpty);
        public ctx?: (common.IContext|null);
    }

    interface IId {
        ctx?: (common.IContext|null);
        account_id?: (string|null);
        id?: (string|null);
    }

    class Id implements IId {
        constructor(p?: common.IId);
        public ctx?: (common.IContext|null);
        public account_id: string;
        public id: string;
    }

    interface IIds {
        ctx?: (common.IContext|null);
        account_id?: (string|null);
        ids?: (string[]|null);
    }

    class Ids implements IIds {
        constructor(p?: common.IIds);
        public ctx?: (common.IContext|null);
        public account_id: string;
        public ids: string[];
    }

    interface IContext {
        event_id?: (string|null);
        state?: (Uint8Array|null);
        node?: (string|null);
        reply_topic?: (string|null);
        credential?: (auth.ICredential|null);
        tracing?: (Uint8Array|null);
        reply_key?: (string|null);
        by_device?: (common.IDevice|null);
        topic?: (string|null);
        partition?: (number|null);
        offset?: (number|Long|null);
        term?: (number|Long|null);
        router_topic?: (string|null);
    }

    class Context implements IContext {
        constructor(p?: common.IContext);
        public event_id: string;
        public state: Uint8Array;
        public node: string;
        public reply_topic: string;
        public credential?: (auth.ICredential|null);
        public tracing: Uint8Array;
        public reply_key: string;
        public by_device?: (common.IDevice|null);
        public topic: string;
        public partition: number;
        public offset: (number|Long);
        public term: (number|Long);
        public router_topic: string;
    }

    interface IReply {
        ctx?: (common.IContext|null);
        event_id?: (string|null);
        state?: (Uint8Array|null);
        service?: (string|null);
        service_id?: (string|null);
        err?: (boolean|null);
        err_description?: (string|null);
        err_code?: (lang.T|null);
        err_class?: (number|null);
        err_hash?: (string|null);
        payload?: (Uint8Array|null);
        published?: (number|Long|null);
    }

    class Reply implements IReply {
        constructor(p?: common.IReply);
        public ctx?: (common.IContext|null);
        public event_id: string;
        public state: Uint8Array;
        public service: string;
        public service_id: string;
        public err: boolean;
        public err_description: string;
        public err_code: lang.T;
        public err_class: number;
        public err_hash: string;
        public payload: Uint8Array;
        public published: (number|Long);
    }

    interface IError {
        error?: (string|null);
        request_id?: (string|null);
        description?: (string|null);
        hash?: (string|null);
        debug?: (string|null);
    }

    class Error implements IError {
        constructor(p?: common.IError);
        public error: string;
        public request_id: string;
        public description: string;
        public hash: string;
        public debug: string;
    }

    interface IDevice {
        ip?: (string|null);
        user_agent?: (string|null);
        screen_resolution?: (string|null);
        timezone?: (string|null);
        language?: (string|null);
        referrer?: (string|null);
        type?: (string|null);
    }

    class Device implements IDevice {
        constructor(p?: common.IDevice);
        public ip: string;
        public user_agent: string;
        public screen_resolution: string;
        public timezone: string;
        public language: string;
        public referrer: string;
        public type: string;
    }

    namespace Device {

        enum DeviceType {
            unknown = 0,
            mobile = 1,
            tablet = 2,
            desktop = 3
        }
    }

    enum Event {
        Send_ = 0,
        ApiReply = 2
    }

    interface IPingRequest {
        message?: (string|null);
    }

    class PingRequest implements IPingRequest {
        constructor(p?: common.IPingRequest);
        public message: string;
    }

    interface IPong {
        message?: (string|null);
    }

    class Pong implements IPong {
        constructor(p?: common.IPong);
        public message: string;
    }

    class Agent extends $protobuf.rpc.Service {
        constructor(rpcImpl: $protobuf.RPCImpl, requestDelimited?: boolean, responseDelimited?: boolean);
        public ping(request: common.IPingRequest, callback: common.Agent.PingCallback): void;
        public ping(request: common.IPingRequest): Promise<common.Pong>;
    }

    namespace Agent {

        type PingCallback = (error: (Error|null), response?: common.Pong) => void;
    }
}

export namespace content {

    interface IAllType {
        content?: (content.IContent|null);
        contents?: (content.IContents|null);
        lr?: (content.IListRequest|null);
    }

    class AllType implements IAllType {
        constructor(p?: content.IAllType);
        public content?: (content.IContent|null);
        public contents?: (content.IContents|null);
        public lr?: (content.IListRequest|null);
    }

    class MyServer extends $protobuf.rpc.Service {
        constructor(rpcImpl: $protobuf.RPCImpl, requestDelimited?: boolean, responseDelimited?: boolean);
        public do(request: content.IAllType, callback: content.MyServer.DoCallback): void;
        public do(request: content.IAllType): Promise<content.AllType>;
    }

    namespace MyServer {

        type DoCallback = (error: (Error|null), response?: content.AllType) => void;
    }

    interface IKeyValue {
        key?: (string|null);
        value?: (string|null);
    }

    class KeyValue implements IKeyValue {
        constructor(p?: content.IKeyValue);
        public key: string;
        public value: string;
    }

    interface IContent {
        ctx?: (common.IContext|null);
        sbid?: (string|null);
        id?: (string|null);
        account_id?: (string|null);
        description?: (string|null);
        title?: (string|null);
        url?: (string|null);
        labels?: (string[]|null);
        availability?: (string|null);
        price?: (number|null);
        currency?: (string|null);
        sale_price?: (number|null);
        fields?: (content.IKeyValue[]|null);
        categories?: (string[]|null);
        relates?: (string[]|null);
        attachment_urls?: (string[]|null);
        type?: (string|null);
    }

    class Content implements IContent {
        constructor(p?: content.IContent);
        public ctx?: (common.IContext|null);
        public sbid: string;
        public id: string;
        public account_id: string;
        public description: string;
        public title: string;
        public url: string;
        public labels: string[];
        public availability: string;
        public price: number;
        public currency: string;
        public sale_price: number;
        public fields: content.IKeyValue[];
        public categories: string[];
        public relates: string[];
        public attachment_urls: string[];
        public type: string;
    }

    interface IContents {
        ctx?: (common.IContext|null);
        contents?: (content.IContent[]|null);
    }

    class Contents implements IContents {
        constructor(p?: content.IContents);
        public ctx?: (common.IContext|null);
        public contents: content.IContent[];
    }

    interface IListRequest {
        account_id?: (string|null);
        anchor?: (string|null);
        category?: (string|null);
        limit?: (number|null);
        label?: (string|null);
        query?: (string|null);
    }

    class ListRequest implements IListRequest {
        constructor(p?: content.IListRequest);
        public account_id: string;
        public anchor: string;
        public category: string;
        public limit: number;
        public label: string;
        public query: string;
    }
}

export namespace conversation {

    interface IRule {
        ctx?: (common.IContext|null);
        id?: (string|null);
        account_id?: (string|null);
        priority?: (number|null);
        strategy?: (string|null);
        assign_tos?: (string[]|null);
        conditions?: (conversation.ICondition[]|null);
        enabled?: (boolean|null);
        created?: (number|Long|null);
        modified?: (number|Long|null);
        name?: (string|null);
        description?: (string|null);
    }

    class Rule implements IRule {
        constructor(p?: conversation.IRule);
        public ctx?: (common.IContext|null);
        public id: string;
        public account_id: string;
        public priority: number;
        public strategy: string;
        public assign_tos: string[];
        public conditions: conversation.ICondition[];
        public enabled: boolean;
        public created: (number|Long);
        public modified: (number|Long);
        public name: string;
        public description: string;
    }

    namespace Rule {

        enum AssignStrategy {
            all_agents = 0,
            agentgroup = 2,
            agents = 3,
            most_recent = 4,
            roundrobin_all_agents = 5,
            roundrobin_agents = 6
        }
    }

    interface IAvailabilityCheckRequest {
        ctx?: (common.IContext|null);
        integration_id?: (string|null);
        user?: (user.IUser|null);
        account_id?: (string|null);
        user_id?: (string|null);
    }

    class AvailabilityCheckRequest implements IAvailabilityCheckRequest {
        constructor(p?: conversation.IAvailabilityCheckRequest);
        public ctx?: (common.IContext|null);
        public integration_id: string;
        public user?: (user.IUser|null);
        public account_id: string;
        public user_id: string;
    }

    interface IAvailabilityCheckResult {
        ctx?: (common.IContext|null);
        availability?: (boolean|null);
        reason?: (string|null);
        payload?: (string|null);
    }

    class AvailabilityCheckResult implements IAvailabilityCheckResult {
        constructor(p?: conversation.IAvailabilityCheckResult);
        public ctx?: (common.IContext|null);
        public availability: boolean;
        public reason: string;
        public payload: string;
    }

    interface ICondition {
        join?: (string|null);
        key?: (string|null);
        operator?: (string|null);
        value?: (string|null);
    }

    class Condition implements ICondition {
        constructor(p?: conversation.ICondition);
        public join: string;
        public key: string;
        public operator: string;
        public value: string;
    }

    namespace Condition {

        enum JoinOperator {
            none = 0,
            and = 1,
            or = 2
        }
    }

    interface IRoute {
        ctx?: (common.IContext|null);
        rules?: (conversation.IRule[]|null);
    }

    class Route implements IRoute {
        constructor(p?: conversation.IRoute);
        public ctx?: (common.IContext|null);
        public rules: conversation.IRule[];
    }

    interface IRouteResult {
        rule_id?: (string|null);
        strategy?: (string|null);
        agent_ids?: (string[]|null);
    }

    class RouteResult implements IRouteResult {
        constructor(p?: conversation.IRouteResult);
        public rule_id: string;
        public strategy: string;
        public agent_ids: string[];
    }

    interface IMemberSeen {
        member_id?: (string|null);
        message_id?: (string|null);
    }

    class MemberSeen implements IMemberSeen {
        constructor(p?: conversation.IMemberSeen);
        public member_id: string;
        public message_id: string;
    }

    interface IMemberV3 {
        ctx?: (common.IContext|null);
        id?: (string|null);
        conversation_id?: (string|null);
        account_id?: (string|null);
        name?: (string|null);
        avatar_url?: (string|null);
        type?: (string|null);
    }

    class MemberV3 implements IMemberV3 {
        constructor(p?: conversation.IMemberV3);
        public ctx?: (common.IContext|null);
        public id: string;
        public conversation_id: string;
        public account_id: string;
        public name: string;
        public avatar_url: string;
        public type: string;
    }

    interface IMember {
        type?: (string|null);
        subiz_id?: (string|null);
        id?: (string|null);
        name?: (string|null);
        avatar_url?: (string|null);
        left?: (boolean|null);
        conversation_id?: (string|null);
    }

    class Member implements IMember {
        constructor(p?: conversation.IMember);
        public type: string;
        public subiz_id: string;
        public id: string;
        public name: string;
        public avatar_url: string;
        public left: boolean;
        public conversation_id: string;
    }

    interface IConversation {
        ctx?: (common.IContext|null);
        id?: (string|null);
        account_id?: (string|null);
        created?: (number|Long|null);
        closed?: (number|Long|null);
        members?: (conversation.IMember[]|null);
        tags?: (conversation.ITag[]|null);
        state?: (string|null);
        request?: (conversation.IStartRequest|null);
        accepted?: (number|Long|null);
        channel_type?: (string|null);
        integration?: (conversation.IIntegration|null);
        actived?: (number|Long|null);
        last_message_id?: (string|null);
        last_message?: (conversation.IMessage|null);
        response_sec?: (number|Long|null);
    }

    class Conversation implements IConversation {
        constructor(p?: conversation.IConversation);
        public ctx?: (common.IContext|null);
        public id: string;
        public account_id: string;
        public created: (number|Long);
        public closed: (number|Long);
        public members: conversation.IMember[];
        public tags: conversation.ITag[];
        public state: string;
        public request?: (conversation.IStartRequest|null);
        public accepted: (number|Long);
        public channel_type: string;
        public integration?: (conversation.IIntegration|null);
        public actived: (number|Long);
        public last_message_id: string;
        public last_message?: (conversation.IMessage|null);
        public response_sec: (number|Long);
    }

    namespace Conversation {

        enum State {
            conversation_none = 0,
            unassigned = 2,
            active = 6,
            ended = 8,
            pending = 9
        }
    }

    interface IUserConversation {
        account_id?: (string|null);
        state?: (string|null);
        user_id?: (string|null);
        convo_id?: (string|null);
        last_seen_event?: (string|null);
    }

    class UserConversation implements IUserConversation {
        constructor(p?: conversation.IUserConversation);
        public account_id: string;
        public state: string;
        public user_id: string;
        public convo_id: string;
        public last_seen_event: string;
    }

    interface IStartRequest {
        ctx?: (common.IContext|null);
        id?: (string|null);
        account_id?: (string|null);
        channel_type?: (string|null);
        from?: (string|null);
        to?: (string[]|null);
        page_url?: (string|null);
        page_title?: (string|null);
        message?: (string|null);
        browser_language?: (string|null);
        language?: (string|null);
        device_type?: (string|null);
        created?: (number|Long|null);
        conversation_id?: (string|null);
        ip?: (string|null);
        country?: (string|null);
        country_code?: (string|null);
        city?: (string|null);
        timezone?: (string|null);
        starter_id?: (string|null);
        starter_type?: (string|null);
        agent_ids?: (string[]|null);
        user?: (user.IUser|null);
        integration_id?: (string|null);
    }

    class StartRequest implements IStartRequest {
        constructor(p?: conversation.IStartRequest);
        public ctx?: (common.IContext|null);
        public id: string;
        public account_id: string;
        public channel_type: string;
        public from: string;
        public to: string[];
        public page_url: string;
        public page_title: string;
        public message: string;
        public browser_language: string;
        public language: string;
        public device_type: string;
        public created: (number|Long);
        public conversation_id: string;
        public ip: string;
        public country: string;
        public country_code: string;
        public city: string;
        public timezone: string;
        public starter_id: string;
        public starter_type: string;
        public agent_ids: string[];
        public user?: (user.IUser|null);
        public integration_id: string;
    }

    interface IConversations {
        ctx?: (common.IContext|null);
        conversations?: (conversation.IConversation[]|null);
        anchor?: (string|null);
        user_ids?: (string[]|null);
    }

    class Conversations implements IConversations {
        constructor(p?: conversation.IConversations);
        public ctx?: (common.IContext|null);
        public conversations: conversation.IConversation[];
        public anchor: string;
        public user_ids: string[];
    }

    interface ISearch {
        ctx?: (common.IContext|null);
        account_id?: (string|null);
        keyword?: (string|null);
        limit?: (number|null);
        before_id?: (string|null);
        after_id?: (string|null);
    }

    class Search implements ISearch {
        constructor(p?: conversation.ISearch);
        public ctx?: (common.IContext|null);
        public account_id: string;
        public keyword: string;
        public limit: number;
        public before_id: string;
        public after_id: string;
    }

    interface IListConversationsRequest {
        ctx?: (common.IContext|null);
        account_id?: (string|null);
        state?: (string|null);
        limit?: (number|null);
        anchor?: (string|null);
        member_id?: (string|null);
        group_by?: (string|null);
        integration_id?: (string|null);
    }

    class ListConversationsRequest implements IListConversationsRequest {
        constructor(p?: conversation.IListConversationsRequest);
        public ctx?: (common.IContext|null);
        public account_id: string;
        public state: string;
        public limit: number;
        public anchor: string;
        public member_id: string;
        public group_by: string;
        public integration_id: string;
    }

    interface IListEventsRequest {
        ctx?: (common.IContext|null);
        account_id?: (string|null);
        conversation_id?: (string|null);
        start_id?: (string|null);
        limit?: (number|null);
    }

    class ListEventsRequest implements IListEventsRequest {
        constructor(p?: conversation.IListEventsRequest);
        public ctx?: (common.IContext|null);
        public account_id: string;
        public conversation_id: string;
        public start_id: string;
        public limit: number;
    }

    interface IListConversationsByUserRequest {
        account_id?: (string|null);
        channel_id?: (string|null);
        user_id?: (string|null);
        start_id?: (string|null);
        limit?: (number|null);
    }

    class ListConversationsByUserRequest implements IListConversationsByUserRequest {
        constructor(p?: conversation.IListConversationsByUserRequest);
        public account_id: string;
        public channel_id: string;
        public user_id: string;
        public start_id: string;
        public limit: number;
    }

    enum Event {
        ConversationAssigned = 0,
        ConversationWaiting = 2,
        ConversationStartRequested = 3,
        ConversationAccepted = 4,
        ConversationDropped = 5,
        ConversationEventCreated = 6,
        ConversationJoinRequested = 7,
        ConversationMessageRequested = 8,
        ConversationLeaveRequested = 9,
        ConversationCloseRequested = 10,
        ConversationTagRequested = 11,
        ConversationUntagRequested = 12,
        ConversationReadRequested = 13,
        ConversationListRequested = 14,
        ConversationAcceptRequested = 15,
        ConversationUpdateRuleRequested = 20,
        ConversationCreateRuleRequested = 21,
        ConversationDeleteRuleRequested = 22,
        ConversationReadRuleRequested = 23,
        ConversationListRuleRequested = 24,
        ConversationUserRequestReply = 51,
        ConversationLimitUpdated = 52,
        ConversationRequestWaitTimeout = 60,
        ConversationListEventsRequested = 61,
        ChannelDeintegrateRequested = 65,
        ChannelIntegrateRequested = 66,
        ChannelIntegrationListRequested = 67,
        ConnectorUpsertRequested = 68,
        ConnectorListRequested = 69,
        CannedResponseCreateRequested = 80,
        CannedResponseUpdateRequested = 81,
        CannedResponseDeleteRequested = 82,
        CannedResponseReadRequested = 83,
        CannedResponseListRequested = 84,
        TagCreateRequested = 85,
        TagUpdateRequested = 86,
        TagReadRequested = 87,
        TagDeleteRequested = 88,
        TagListRequested = 89,
        TagCreated = 95,
        ConversationUpserted = 97,
        ConversationMessageSent = 99,
        ConversationMessageAckRequested = 90,
        ConversationMessageReceiveRequested = 91,
        ConversationMessageSeeRequested = 92,
        ChannelIntegrationAvailabilityCheck = 93,
        ConversationMessageSearchRequest = 94,
        ConversationRequested = 100,
        ConversationSynced = 101,
        ConversationV3Synced = 102
    }

    interface IRequestState {
        account_id?: (string|null);
        conversation_id?: (string|null);
        channel_id?: (string|null);
    }

    class RequestState implements IRequestState {
        constructor(p?: conversation.IRequestState);
        public account_id: string;
        public conversation_id: string;
        public channel_id: string;
    }

    interface IReaction {
        name?: (string|null);
        count?: (number|null);
        users?: (string[]|null);
    }

    class Reaction implements IReaction {
        constructor(p?: conversation.IReaction);
        public name: string;
        public count: number;
        public users: string[];
    }

    interface IEsMessage {
        id?: (string|null);
        account_id?: (string|null);
        conversation_id?: (string|null);
        member_ids?: (string[]|null);
        text?: (string|null);
        attachments?: (string[]|null);
        fields?: (string[]|null);
    }

    class EsMessage implements IEsMessage {
        constructor(p?: conversation.IEsMessage);
        public id: string;
        public account_id: string;
        public conversation_id: string;
        public member_ids: string[];
        public text: string;
        public attachments: string[];
        public fields: string[];
    }

    interface IMessage {
        ctx?: (common.IContext|null);
        account_id?: (string|null);
        conversation_id?: (string|null);
        id?: (string|null);
        text?: (string|null);
        format?: (string|null);
        attachments?: (conversation.IAttachment[]|null);
        reactions?: (conversation.IReaction[]|null);
        fields?: (conversation.IField[]|null);
        computed?: (conversation.IComputed|null);
        integration_id?: (string|null);
    }

    class Message implements IMessage {
        constructor(p?: conversation.IMessage);
        public ctx?: (common.IContext|null);
        public account_id: string;
        public conversation_id: string;
        public id: string;
        public text: string;
        public format: string;
        public attachments: conversation.IAttachment[];
        public reactions: conversation.IReaction[];
        public fields: conversation.IField[];
        public computed?: (conversation.IComputed|null);
        public integration_id: string;
    }

    interface IComputed {
        seen?: (conversation.ISeen[]|null);
        ack?: (conversation.IAck[]|null);
        received?: (conversation.IReceived[]|null);
    }

    class Computed implements IComputed {
        constructor(p?: conversation.IComputed);
        public seen: conversation.ISeen[];
        public ack: conversation.IAck[];
        public received: conversation.IReceived[];
    }

    interface ISeen {
        member_id?: (string|null);
        at?: (number|Long|null);
    }

    class Seen implements ISeen {
        constructor(p?: conversation.ISeen);
        public member_id: string;
        public at: (number|Long);
    }

    interface IAck {
        member_id?: (string|null);
        error?: (string|null);
        at?: (number|Long|null);
    }

    class Ack implements IAck {
        constructor(p?: conversation.IAck);
        public member_id: string;
        public error: string;
        public at: (number|Long);
    }

    interface IReceived {
        member_id?: (string|null);
        at?: (number|Long|null);
    }

    class Received implements IReceived {
        constructor(p?: conversation.IReceived);
        public member_id: string;
        public at: (number|Long);
    }

    interface IField {
        value?: (string|null);
        key?: (string|null);
    }

    class Field implements IField {
        constructor(p?: conversation.IField);
        public value: string;
        public key: string;
    }

    interface IButton {
        type?: (string|null);
        id?: (string|null);
        title?: (string|null);
        payload?: (string|null);
        image_url?: (string|null);
        content_id?: (string|null);
        url?: (string|null);
    }

    class Button implements IButton {
        constructor(p?: conversation.IButton);
        public type: string;
        public id: string;
        public title: string;
        public payload: string;
        public image_url: string;
        public content_id: string;
        public url: string;
    }

    namespace Button {

        enum ButtonType {
            url_button = 2,
            postback_button = 3,
            event_button = 4
        }
    }

    interface IAskInfomationAnswer {
        message_id?: (string|null);
        answer?: (string|null);
    }

    class AskInfomationAnswer implements IAskInfomationAnswer {
        constructor(p?: conversation.IAskInfomationAnswer);
        public message_id: string;
        public answer: string;
    }

    interface IAskInfomation {
        question?: (string|null);
        input_type?: (string|null);
        key?: (string|null);
        answer?: (string|null);
        answered?: (number|Long|null);
    }

    class AskInfomation implements IAskInfomation {
        constructor(p?: conversation.IAskInfomation);
        public question: string;
        public input_type: string;
        public key: string;
        public answer: string;
        public answered: (number|Long);
    }

    namespace AskInfomation {

        enum InputType {
            phone = 0,
            email = 1,
            text = 2,
            password = 3,
            number = 4,
            date = 5,
            color = 6,
            location = 7,
            time = 8,
            url = 9
        }
    }

    interface IGenericElementTemplate {
        title?: (string|null);
        image_url?: (string|null);
        subtitle?: (string|null);
        default_action?: (conversation.IButton|null);
        buttons?: (conversation.IButton[]|null);
    }

    class GenericElementTemplate implements IGenericElementTemplate {
        constructor(p?: conversation.IGenericElementTemplate);
        public title: string;
        public image_url: string;
        public subtitle: string;
        public default_action?: (conversation.IButton|null);
        public buttons: conversation.IButton[];
    }

    interface IAttachment {
        type?: (string|null);
        mimetype?: (string|null);
        url?: (string|null);
        thumbnail_url?: (string|null);
        name?: (string|null);
        description?: (string|null);
        length?: (number|null);
        size?: (number|null);
        elements?: (conversation.IGenericElementTemplate[]|null);
        title?: (string|null);
        color?: (string|null);
        pretext?: (string|null);
        buttons?: (conversation.IButton[]|null);
        ask_info?: (conversation.IAskInfomation|null);
        ask_info_answer?: (conversation.IAskInfomationAnswer|null);
    }

    class Attachment implements IAttachment {
        constructor(p?: conversation.IAttachment);
        public type: string;
        public mimetype: string;
        public url: string;
        public thumbnail_url: string;
        public name: string;
        public description: string;
        public length: number;
        public size: number;
        public elements: conversation.IGenericElementTemplate[];
        public title: string;
        public color: string;
        public pretext: string;
        public buttons: conversation.IButton[];
        public ask_info?: (conversation.IAskInfomation|null);
        public ask_info_answer?: (conversation.IAskInfomationAnswer|null);
    }

    namespace Attachment {

        enum AttachmentType {
            file = 2,
            generic = 3,
            preview = 4,
            button = 5,
            input = 6,
            ask_info_form = 7,
            ask_info_form_answer = 8
        }
    }

    interface ITag {
        ctx?: (common.IContext|null);
        id?: (string|null);
        account_id?: (string|null);
        title?: (string|null);
        created?: (number|Long|null);
        modified?: (number|Long|null);
        creator_id?: (string|null);
        color?: (string|null);
    }

    class Tag implements ITag {
        constructor(p?: conversation.ITag);
        public ctx?: (common.IContext|null);
        public id: string;
        public account_id: string;
        public title: string;
        public created: (number|Long);
        public modified: (number|Long);
        public creator_id: string;
        public color: string;
    }

    interface ICannedResponse {
        ctx?: (common.IContext|null);
        id?: (string|null);
        account_id?: (string|null);
        text?: (string|null);
        keys?: (string[]|null);
        created?: (number|Long|null);
        modified?: (number|Long|null);
        creator?: (string|null);
    }

    class CannedResponse implements ICannedResponse {
        constructor(p?: conversation.ICannedResponse);
        public ctx?: (common.IContext|null);
        public id: string;
        public account_id: string;
        public text: string;
        public keys: string[];
        public created: (number|Long);
        public modified: (number|Long);
        public creator: string;
    }

    interface ICannedResponses {
        ctx?: (common.IContext|null);
        responses?: (conversation.ICannedResponse[]|null);
    }

    class CannedResponses implements ICannedResponses {
        constructor(p?: conversation.ICannedResponses);
        public ctx?: (common.IContext|null);
        public responses: conversation.ICannedResponse[];
    }

    interface ITags {
        ctx?: (common.IContext|null);
        tags?: (conversation.ITag[]|null);
    }

    class Tags implements ITags {
        constructor(p?: conversation.ITags);
        public ctx?: (common.IContext|null);
        public tags: conversation.ITag[];
    }

    interface IPostback {
        message?: (conversation.IMessage|null);
        button?: (conversation.IButton|null);
    }

    class Postback implements IPostback {
        constructor(p?: conversation.IPostback);
        public message?: (conversation.IMessage|null);
        public button?: (conversation.IButton|null);
    }

    interface IUseConnector {
        account_id?: (string|null);
        connector_id?: (string|null);
        state?: (string|null);
        updated?: (string|null);
        by?: (string|null);
    }

    class UseConnector implements IUseConnector {
        constructor(p?: conversation.IUseConnector);
        public account_id: string;
        public connector_id: string;
        public state: string;
        public updated: string;
        public by: string;
    }

    namespace UseConnector {

        enum State {
            disabled = 0,
            activated = 1
        }
    }

    interface IIntegrations {
        ctx?: (common.IContext|null);
        account_id?: (string|null);
        integrations?: (conversation.IIntegration[]|null);
    }

    class Integrations implements IIntegrations {
        constructor(p?: conversation.IIntegrations);
        public ctx?: (common.IContext|null);
        public account_id: string;
        public integrations: conversation.IIntegration[];
    }

    interface IIntegration {
        ctx?: (common.IContext|null);
        account_id?: (string|null);
        connector_id?: (string|null);
        logo_url?: (string|null);
        name?: (string|null);
        connector_type?: (string|null);
        integrated?: (number|Long|null);
        updated?: (number|Long|null);
        state?: (string|null);
        id?: (string|null);
    }

    class Integration implements IIntegration {
        constructor(p?: conversation.IIntegration);
        public ctx?: (common.IContext|null);
        public account_id: string;
        public connector_id: string;
        public logo_url: string;
        public name: string;
        public connector_type: string;
        public integrated: (number|Long);
        public updated: (number|Long);
        public state: string;
        public id: string;
    }

    namespace Integration {

        enum State {
            disabled = 0,
            activated = 1
        }
    }

    interface ISearchMessageRequest {
        ctx?: (common.IContext|null);
        account_id?: (string|null);
        conversation_id?: (string|null);
        user_ids?: (string[]|null);
        query?: (string|null);
        anchor?: (string|null);
        limit?: (number|null);
    }

    class SearchMessageRequest implements ISearchMessageRequest {
        constructor(p?: conversation.ISearchMessageRequest);
        public ctx?: (common.IContext|null);
        public account_id: string;
        public conversation_id: string;
        public user_ids: string[];
        public query: string;
        public anchor: string;
        public limit: number;
    }

    class RuleMgr extends $protobuf.rpc.Service {
        constructor(rpcImpl: $protobuf.RPCImpl, requestDelimited?: boolean, responseDelimited?: boolean);
        public updateRule(request: conversation.IRule, callback: conversation.RuleMgr.UpdateRuleCallback): void;
        public updateRule(request: conversation.IRule): Promise<conversation.Rule>;
        public createRule(request: conversation.IRule, callback: conversation.RuleMgr.CreateRuleCallback): void;
        public createRule(request: conversation.IRule): Promise<conversation.Rule>;
        public deleteRule(request: common.IId, callback: conversation.RuleMgr.DeleteRuleCallback): void;
        public deleteRule(request: common.IId): Promise<common.Empty>;
        public readRule(request: common.IId, callback: conversation.RuleMgr.ReadRuleCallback): void;
        public readRule(request: common.IId): Promise<conversation.Rule>;
        public listRules(request: common.IId, callback: conversation.RuleMgr.ListRulesCallback): void;
        public listRules(request: common.IId): Promise<conversation.Route>;
    }

    namespace RuleMgr {

        type UpdateRuleCallback = (error: (Error|null), response?: conversation.Rule) => void;

        type CreateRuleCallback = (error: (Error|null), response?: conversation.Rule) => void;

        type DeleteRuleCallback = (error: (Error|null), response?: common.Empty) => void;

        type ReadRuleCallback = (error: (Error|null), response?: conversation.Rule) => void;

        type ListRulesCallback = (error: (Error|null), response?: conversation.Route) => void;
    }

    interface IMessageId {
        conversation_id?: (string|null);
        message_id?: (string|null);
    }

    class MessageId implements IMessageId {
        constructor(p?: conversation.IMessageId);
        public conversation_id: string;
        public message_id: string;
    }

    interface IMessageAck {
        conversation_id?: (string|null);
        message_id?: (string|null);
        error?: (string|null);
    }

    class MessageAck implements IMessageAck {
        constructor(p?: conversation.IMessageAck);
        public conversation_id: string;
        public message_id: string;
        public error: string;
    }

    class ConversationMgr extends $protobuf.rpc.Service {
        constructor(rpcImpl: $protobuf.RPCImpl, requestDelimited?: boolean, responseDelimited?: boolean);
        public seenMessage(request: conversation.IMessageId, callback: conversation.ConversationMgr.SeenMessageCallback): void;
        public seenMessage(request: conversation.IMessageId): Promise<common.Empty>;
        public receiveMessage(request: conversation.IMessageId, callback: conversation.ConversationMgr.ReceiveMessageCallback): void;
        public receiveMessage(request: conversation.IMessageId): Promise<common.Empty>;
        public ackMessage(request: conversation.IMessageAck, callback: conversation.ConversationMgr.AckMessageCallback): void;
        public ackMessage(request: conversation.IMessageAck): Promise<common.Empty>;
        public startConversation(request: conversation.IStartRequest, callback: conversation.ConversationMgr.StartConversationCallback): void;
        public startConversation(request: conversation.IStartRequest): Promise<conversation.Conversation>;
        public endConversation(request: common.IId, callback: conversation.ConversationMgr.EndConversationCallback): void;
        public endConversation(request: common.IId): Promise<conversation.Conversation>;
        public getConversation(request: common.IId, callback: conversation.ConversationMgr.GetConversationCallback): void;
        public getConversation(request: common.IId): Promise<conversation.Conversation>;
        public listConversations(request: conversation.IListConversationsRequest, callback: conversation.ConversationMgr.ListConversationsCallback): void;
        public listConversations(request: conversation.IListConversationsRequest): Promise<conversation.Conversations>;
        public acceptConversation(request: common.IId, callback: conversation.ConversationMgr.AcceptConversationCallback): void;
        public acceptConversation(request: common.IId): Promise<conversation.Conversation>;
        public tagConversation(request: conversation.ITagRequest, callback: conversation.ConversationMgr.TagConversationCallback): void;
        public tagConversation(request: conversation.ITagRequest): Promise<common.Empty>;
        public untagConversation(request: conversation.ITagRequest, callback: conversation.ConversationMgr.UntagConversationCallback): void;
        public untagConversation(request: conversation.ITagRequest): Promise<common.Empty>;
        public joinConversation(request: conversation.IMember, callback: conversation.ConversationMgr.JoinConversationCallback): void;
        public joinConversation(request: conversation.IMember): Promise<common.Empty>;
        public leftConversation(request: conversation.IMember, callback: conversation.ConversationMgr.LeftConversationCallback): void;
        public leftConversation(request: conversation.IMember): Promise<common.Empty>;
        public typing(request: common.IId, callback: conversation.ConversationMgr.TypingCallback): void;
        public typing(request: common.IId): Promise<common.Empty>;
    }

    namespace ConversationMgr {

        type SeenMessageCallback = (error: (Error|null), response?: common.Empty) => void;

        type ReceiveMessageCallback = (error: (Error|null), response?: common.Empty) => void;

        type AckMessageCallback = (error: (Error|null), response?: common.Empty) => void;

        type StartConversationCallback = (error: (Error|null), response?: conversation.Conversation) => void;

        type EndConversationCallback = (error: (Error|null), response?: conversation.Conversation) => void;

        type GetConversationCallback = (error: (Error|null), response?: conversation.Conversation) => void;

        type ListConversationsCallback = (error: (Error|null), response?: conversation.Conversations) => void;

        type AcceptConversationCallback = (error: (Error|null), response?: conversation.Conversation) => void;

        type TagConversationCallback = (error: (Error|null), response?: common.Empty) => void;

        type UntagConversationCallback = (error: (Error|null), response?: common.Empty) => void;

        type JoinConversationCallback = (error: (Error|null), response?: common.Empty) => void;

        type LeftConversationCallback = (error: (Error|null), response?: common.Empty) => void;

        type TypingCallback = (error: (Error|null), response?: common.Empty) => void;
    }

    class CannedResponseMgr extends $protobuf.rpc.Service {
        constructor(rpcImpl: $protobuf.RPCImpl, requestDelimited?: boolean, responseDelimited?: boolean);
        public createCannedResponse(request: conversation.ICannedResponse, callback: conversation.CannedResponseMgr.CreateCannedResponseCallback): void;
        public createCannedResponse(request: conversation.ICannedResponse): Promise<conversation.CannedResponse>;
        public updateCannedResponse(request: conversation.ICannedResponse, callback: conversation.CannedResponseMgr.UpdateCannedResponseCallback): void;
        public updateCannedResponse(request: conversation.ICannedResponse): Promise<conversation.CannedResponse>;
        public listCannedResponses(request: common.IId, callback: conversation.CannedResponseMgr.ListCannedResponsesCallback): void;
        public listCannedResponses(request: common.IId): Promise<conversation.CannedResponses>;
        public deleteCannedResponse(request: common.IId, callback: conversation.CannedResponseMgr.DeleteCannedResponseCallback): void;
        public deleteCannedResponse(request: common.IId): Promise<common.Empty>;
        public getCannedResponse(request: common.IId, callback: conversation.CannedResponseMgr.GetCannedResponseCallback): void;
        public getCannedResponse(request: common.IId): Promise<conversation.CannedResponse>;
    }

    namespace CannedResponseMgr {

        type CreateCannedResponseCallback = (error: (Error|null), response?: conversation.CannedResponse) => void;

        type UpdateCannedResponseCallback = (error: (Error|null), response?: conversation.CannedResponse) => void;

        type ListCannedResponsesCallback = (error: (Error|null), response?: conversation.CannedResponses) => void;

        type DeleteCannedResponseCallback = (error: (Error|null), response?: common.Empty) => void;

        type GetCannedResponseCallback = (error: (Error|null), response?: conversation.CannedResponse) => void;
    }

    class TagMgr extends $protobuf.rpc.Service {
        constructor(rpcImpl: $protobuf.RPCImpl, requestDelimited?: boolean, responseDelimited?: boolean);
        public createTag(request: conversation.ITag, callback: conversation.TagMgr.CreateTagCallback): void;
        public createTag(request: conversation.ITag): Promise<conversation.Tag>;
        public updateTag(request: conversation.ITag, callback: conversation.TagMgr.UpdateTagCallback): void;
        public updateTag(request: conversation.ITag): Promise<conversation.Tag>;
        public deleteTag(request: common.IId, callback: conversation.TagMgr.DeleteTagCallback): void;
        public deleteTag(request: common.IId): Promise<common.Empty>;
        public listTags(request: common.IId, callback: conversation.TagMgr.ListTagsCallback): void;
        public listTags(request: common.IId): Promise<conversation.Tags>;
        public getTag(request: common.IId, callback: conversation.TagMgr.GetTagCallback): void;
        public getTag(request: common.IId): Promise<conversation.Tag>;
    }

    namespace TagMgr {

        type CreateTagCallback = (error: (Error|null), response?: conversation.Tag) => void;

        type UpdateTagCallback = (error: (Error|null), response?: conversation.Tag) => void;

        type DeleteTagCallback = (error: (Error|null), response?: common.Empty) => void;

        type ListTagsCallback = (error: (Error|null), response?: conversation.Tags) => void;

        type GetTagCallback = (error: (Error|null), response?: conversation.Tag) => void;
    }

    class IntegrationMgr extends $protobuf.rpc.Service {
        constructor(rpcImpl: $protobuf.RPCImpl, requestDelimited?: boolean, responseDelimited?: boolean);
        public deintegrate(request: common.IId, callback: conversation.IntegrationMgr.DeintegrateCallback): void;
        public deintegrate(request: common.IId): Promise<common.Empty>;
        public listIntegrations(request: common.IId, callback: conversation.IntegrationMgr.ListIntegrationsCallback): void;
        public listIntegrations(request: common.IId): Promise<conversation.Integrations>;
        public integrate(request: conversation.IIntegration, callback: conversation.IntegrationMgr.IntegrateCallback): void;
        public integrate(request: conversation.IIntegration): Promise<conversation.Integration>;
        public checkAvailability(request: conversation.IAvailabilityCheckRequest, callback: conversation.IntegrationMgr.CheckAvailabilityCallback): void;
        public checkAvailability(request: conversation.IAvailabilityCheckRequest): Promise<conversation.AvailabilityCheckResult>;
    }

    namespace IntegrationMgr {

        type DeintegrateCallback = (error: (Error|null), response?: common.Empty) => void;

        type ListIntegrationsCallback = (error: (Error|null), response?: conversation.Integrations) => void;

        type IntegrateCallback = (error: (Error|null), response?: conversation.Integration) => void;

        type CheckAvailabilityCallback = (error: (Error|null), response?: conversation.AvailabilityCheckResult) => void;
    }

    interface ITagRequest {
        ctx?: (common.IContext|null);
        account_id?: (string|null);
        conversation_id?: (string|null);
        id?: (string|null);
    }

    class TagRequest implements ITagRequest {
        constructor(p?: conversation.ITagRequest);
        public ctx?: (common.IContext|null);
        public account_id: string;
        public conversation_id: string;
        public id: string;
    }

    interface ICountByAgentsRequest {
        integration_id?: (string|null);
        agent_ids?: (string[]|null);
        from?: (number|null);
        to?: (number|null);
        range?: (string|null);
    }

    class CountByAgentsRequest implements ICountByAgentsRequest {
        constructor(p?: conversation.ICountByAgentsRequest);
        public integration_id: string;
        public agent_ids: string[];
        public from: number;
        public to: number;
        public range: string;
    }

    namespace CountByAgentsRequest {

        enum Range {
            hour = 0,
            day = 1
        }
    }

    interface ICountByTagsRequest {
        integration_id?: (string|null);
        tag_ids?: (string[]|null);
        from?: (number|null);
        to?: (number|null);
        range?: (string|null);
    }

    class CountByTagsRequest implements ICountByTagsRequest {
        constructor(p?: conversation.ICountByTagsRequest);
        public integration_id: string;
        public tag_ids: string[];
        public from: number;
        public to: number;
        public range: string;
    }

    namespace CountByTagsRequest {

        enum Range {
            hour = 0,
            day = 1
        }
    }

    interface ICountByAgent {
        agent_id?: (string|null);
        data?: ((number|Long)[]|null);
    }

    class CountByAgent implements ICountByAgent {
        constructor(p?: conversation.ICountByAgent);
        public agent_id: string;
        public data: (number|Long)[];
    }

    interface ICountByTag {
        tag_id?: (string|null);
        data?: ((number|Long)[]|null);
    }

    class CountByTag implements ICountByTag {
        constructor(p?: conversation.ICountByTag);
        public tag_id: string;
        public data: (number|Long)[];
    }

    interface ICountByTagsResponse {
        data?: (conversation.ICountByTag[]|null);
    }

    class CountByTagsResponse implements ICountByTagsResponse {
        constructor(p?: conversation.ICountByTagsResponse);
        public data: conversation.ICountByTag[];
    }

    interface ICountByAgentsResponse {
        data?: (conversation.ICountByAgent[]|null);
    }

    class CountByAgentsResponse implements ICountByAgentsResponse {
        constructor(p?: conversation.ICountByAgentsResponse);
        public data: conversation.ICountByAgent[];
    }

    interface IAvgResponseTimeRequest {
        integration_id?: (string|null);
        from?: (number|null);
        to?: (number|null);
    }

    class AvgResponseTimeRequest implements IAvgResponseTimeRequest {
        constructor(p?: conversation.IAvgResponseTimeRequest);
        public integration_id: string;
        public from: number;
        public to: number;
    }

    interface IAvgResponseTimeResponse {
        avg_response_sec?: (number|null);
    }

    class AvgResponseTimeResponse implements IAvgResponseTimeResponse {
        constructor(p?: conversation.IAvgResponseTimeResponse);
        public avg_response_sec: number;
    }

    interface ITotalConversationResponse {
        total_conversation?: (number|Long|null);
    }

    class TotalConversationResponse implements ITotalConversationResponse {
        constructor(p?: conversation.ITotalConversationResponse);
        public total_conversation: (number|Long);
    }

    class ConversationReporter extends $protobuf.rpc.Service {
        constructor(rpcImpl: $protobuf.RPCImpl, requestDelimited?: boolean, responseDelimited?: boolean);
        public countConversationsByAgents(request: conversation.ICountByAgentsRequest, callback: conversation.ConversationReporter.CountConversationsByAgentsCallback): void;
        public countConversationsByAgents(request: conversation.ICountByAgentsRequest): Promise<conversation.CountByAgentsResponse>;
        public countConversationsByTags(request: conversation.ICountByTagsRequest, callback: conversation.ConversationReporter.CountConversationsByTagsCallback): void;
        public countConversationsByTags(request: conversation.ICountByTagsRequest): Promise<conversation.CountByTagsResponse>;
        public getAvgResponseTimes(request: conversation.IAvgResponseTimeRequest, callback: conversation.ConversationReporter.GetAvgResponseTimesCallback): void;
        public getAvgResponseTimes(request: conversation.IAvgResponseTimeRequest): Promise<conversation.AvgResponseTimeResponse>;
        public getTotalConversation(request: conversation.IAvgResponseTimeRequest, callback: conversation.ConversationReporter.GetTotalConversationCallback): void;
        public getTotalConversation(request: conversation.IAvgResponseTimeRequest): Promise<conversation.TotalConversationResponse>;
    }

    namespace ConversationReporter {

        type CountConversationsByAgentsCallback = (error: (Error|null), response?: conversation.CountByAgentsResponse) => void;

        type CountConversationsByTagsCallback = (error: (Error|null), response?: conversation.CountByTagsResponse) => void;

        type GetAvgResponseTimesCallback = (error: (Error|null), response?: conversation.AvgResponseTimeResponse) => void;

        type GetTotalConversationCallback = (error: (Error|null), response?: conversation.TotalConversationResponse) => void;
    }
}

export namespace dashboard {

    interface ISessionCookie {
        refresh_token?: (string|null);
        expired_at?: (number|Long|null);
        issued_at?: (number|Long|null);
        type?: (string|null);
        email?: (string|null);
        remember_me?: (boolean|null);
    }

    class SessionCookie implements ISessionCookie {
        constructor(p?: dashboard.ISessionCookie);
        public refresh_token: string;
        public expired_at: (number|Long);
        public issued_at: (number|Long);
        public type: string;
        public email: string;
        public remember_me: boolean;
    }

    interface IGlobal {
        ctx?: (common.IContext|null);
        name?: (string|null);
        data?: (string|null);
        pk?: (string|null);
    }

    class Global implements IGlobal {
        constructor(p?: dashboard.IGlobal);
        public ctx?: (common.IContext|null);
        public name: string;
        public data: string;
        public pk: string;
    }

    enum Event {
        DashboardVersionUpdated = 0,
        DashboardLanguageUpdated = 5,
        DashboardLanguageUpdateRequested = 7
    }
}

export namespace email {

    interface IEmail {
        ctx?: (common.IContext|null);
        from?: (string|null);
        subject?: (string|null);
        body?: (string|null);
        header?: ({ [k: string]: string }|null);
        attachments?: (email.IAttachment[]|null);
        to?: (string[]|null);
        cc?: (string[]|null);
        bcc?: (string[]|null);
    }

    class Email implements IEmail {
        constructor(p?: email.IEmail);
        public ctx?: (common.IContext|null);
        public from: string;
        public subject: string;
        public body: string;
        public header: { [k: string]: string };
        public attachments: email.IAttachment[];
        public to: string[];
        public cc: string[];
        public bcc: string[];
    }

    interface IAttachment {
        url?: (string|null);
        name?: (string|null);
        mimetype?: (string|null);
    }

    class Attachment implements IAttachment {
        constructor(p?: email.IAttachment);
        public url: string;
        public name: string;
        public mimetype: string;
    }

    enum Event {
        Email_SendRequest = 0
    }
}

export namespace event {

    interface IRawEventCreatedPayload {
        ctx?: (common.IContext|null);
        subs?: (string[]|null);
        target_topic?: (string|null);
        payload?: (string|null);
        target_key?: (string|null);
        payloads?: (string[]|null);
        topic?: (string|null);
        router_topic?: (string|null);
        sub?: (string|null);
    }

    class RawEventCreatedPayload implements IRawEventCreatedPayload {
        constructor(p?: event.IRawEventCreatedPayload);
        public ctx?: (common.IContext|null);
        public subs: string[];
        public target_topic: string;
        public payload: string;
        public target_key: string;
        public payloads: string[];
        public topic: string;
        public router_topic: string;
        public sub: string;
    }

    interface IPubSubMessage {
        ctx?: (common.IContext|null);
        sub?: (event.ISubscription|null);
        payload?: (Uint8Array|null);
    }

    class PubSubMessage implements IPubSubMessage {
        constructor(p?: event.IPubSubMessage);
        public ctx?: (common.IContext|null);
        public sub?: (event.ISubscription|null);
        public payload: Uint8Array;
    }

    interface IUnsubscribeMessage {
        ctx?: (common.IContext|null);
        topic?: (string|null);
        sub_id?: (string|null);
    }

    class UnsubscribeMessage implements IUnsubscribeMessage {
        constructor(p?: event.IUnsubscribeMessage);
        public ctx?: (common.IContext|null);
        public topic: string;
        public sub_id: string;
    }

    enum Event {
        Sub = 0,
        RawEventCreated = 3,
        Subscribe = 4,
        SubscribeReply = 6,
        Unsubscribe = 5,
        UnsubscribeReply = 7,
        EventSync = 8,
        EventCreated = 9
    }

    interface IRawEvents {
        ctx?: (common.IContext|null);
        events?: (event.IRawEvent[]|null);
        total?: (number|Long|null);
        anchor?: (string|null);
    }

    class RawEvents implements IRawEvents {
        constructor(p?: event.IRawEvents);
        public ctx?: (common.IContext|null);
        public events: event.IRawEvent[];
        public total: (number|Long);
        public anchor: string;
    }

    interface IBy {
        device?: (common.IDevice|null);
        id?: (string|null);
        client_id?: (string|null);
    }

    class By implements IBy {
        constructor(p?: event.IBy);
        public device?: (common.IDevice|null);
        public id: string;
        public client_id: string;
    }

    interface IRawEvent {
        ctx?: (common.IContext|null);
        id?: (string|null);
        account_id?: (string|null);
        created?: (number|Long|null);
        type?: (string|null);
        topics?: (string[]|null);
        by?: (event.IBy|null);
        object?: (string|null);
        conversation_id?: (string|null);
        data?: (event.RawEvent.IData|null);
    }

    class RawEvent implements IRawEvent {
        constructor(p?: event.IRawEvent);
        public ctx?: (common.IContext|null);
        public id: string;
        public account_id: string;
        public created: (number|Long);
        public type: string;
        public topics: string[];
        public by?: (event.IBy|null);
        public object: string;
        public conversation_id: string;
        public data?: (event.RawEvent.IData|null);
    }

    namespace RawEvent {

        interface IData {
            account?: (account.IAccount|null);
            agent?: (account.IAgent|null);
            message?: (conversation.IMessage|null);
            conversation?: (conversation.IConversation|null);
            postback?: (conversation.IPostback|null);
            content?: (content.IContent|null);
            topic?: (user.ITopic|null);
            presence?: (user.IPresence|null);
            user?: (user.IUser|null);
            unread_topic?: (user.IUnreadTopic|null);
            my_user?: (user.IMyUser|null);
            notification?: (notibox.INotification|null);
            notibox?: (notibox.IBox|null);
            agent_perm?: (account.IAgentPerm|null);
            group_member?: (account.IGroupMember|null);
            group?: (account.IAgentGroup|null);
            limit?: (payment.ILimit|null);
            user_attribute?: (user.IAttributeData|null);
        }

        class Data implements IData {
            constructor(p?: event.RawEvent.IData);
            public account?: (account.IAccount|null);
            public agent?: (account.IAgent|null);
            public message?: (conversation.IMessage|null);
            public conversation?: (conversation.IConversation|null);
            public postback?: (conversation.IPostback|null);
            public content?: (content.IContent|null);
            public topic?: (user.ITopic|null);
            public presence?: (user.IPresence|null);
            public user?: (user.IUser|null);
            public unread_topic?: (user.IUnreadTopic|null);
            public my_user?: (user.IMyUser|null);
            public notification?: (notibox.INotification|null);
            public notibox?: (notibox.IBox|null);
            public agent_perm?: (account.IAgentPerm|null);
            public group_member?: (account.IGroupMember|null);
            public group?: (account.IAgentGroup|null);
            public limit?: (payment.ILimit|null);
            public user_attribute?: (user.IAttributeData|null);
        }
    }

    enum AccountType {
        a = 0,
        account_deleted = 60,
        account_info_updated = 61,
        group_created = 62,
        group_deleted = 63,
        group_info_updated = 64,
        group_joined = 65,
        group_left = 66,
        invitation_accepted = 68,
        agent_deleted = 69,
        agent_info_updated = 70,
        agent_permission_updated = 71,
        limit_updated = 72
    }

    enum UserType {
        u = 0,
        my_user_upserted = 46,
        user_info_updated = 48
    }

    enum EventType {
        e = 0,
        presence_updated = 38,
        content_viewed = 39,
        content_searched = 40,
        content_addedtocart = 41,
        content_checkedout = 42,
        content_purchased = 43,
        topic_read = 44,
        subscribed_topic_updated = 45,
        user_topic_updated = 50,
        user_attribute_created = 51,
        user_attribute_updated = 52,
        user_attribute_deleted = 55
    }

    enum ConvoType {
        c = 0,
        conversation_updated = 9,
        message_sent = 10,
        conversation_state_updated = 11,
        conversation_joined = 2,
        conversation_left = 4,
        conversation_tagged = 6,
        conversation_untagged = 7,
        channel_deintegrated = 20,
        channel_integrated = 21,
        message_seen = 30,
        message_acked = 31,
        message_received = 32,
        conversation_member_typing = 33,
        conversation_postbacked = 34,
        conversation_unassigned = 35,
        conversation_assigned = 36,
        conversation_pending = 81
    }

    enum NotiboxType {
        n = 0,
        notification_upserted = 52,
        notibox_upserted = 53
    }

    enum SubscriberType {
        st_user = 0,
        st_channel = 1
    }

    interface ISubscription {
        ctx?: (common.IContext|null);
        topic?: (string|null);
        sub_id?: (string|null);
        target_topic?: (string|null);
        target_key?: (string|null);
        ttls?: (number|Long|null);
        router_topic?: (string|null);
        target_partition?: (number|null);
    }

    class Subscription implements ISubscription {
        constructor(p?: event.ISubscription);
        public ctx?: (common.IContext|null);
        public topic: string;
        public sub_id: string;
        public target_topic: string;
        public target_key: string;
        public ttls: (number|Long);
        public router_topic: string;
        public target_partition: number;
    }

    enum SubPrefix {
        Webhook = 0,
        Websocket = 1
    }

    enum Object {
        none = 0,
        account = 1,
        agent = 2,
        message = 3,
        conversation = 4,
        postback = 5,
        content = 6,
        topic = 7,
        presence = 8,
        user = 10,
        unread_topic = 11,
        my_user = 12,
        notification = 14,
        notibox = 15,
        agent_perm = 16,
        group_member = 17,
        group = 18,
        integration = 30,
        limit = 19
    }

    interface IListEventsRequest {
        ctx?: (common.IContext|null);
        account_id?: (string|null);
        user_id?: (string|null);
        query?: (string|null);
        anchor?: (string|null);
        limit?: (number|null);
        category?: (string|null);
    }

    class ListEventsRequest implements IListEventsRequest {
        constructor(p?: event.IListEventsRequest);
        public ctx?: (common.IContext|null);
        public account_id: string;
        public user_id: string;
        public query: string;
        public anchor: string;
        public limit: number;
        public category: string;
    }

    interface IUserEvent {
        ctx?: (common.IContext|null);
        account_id?: (string|null);
        user_id?: (string|null);
        event?: (event.IRawEvent|null);
    }

    class UserEvent implements IUserEvent {
        constructor(p?: event.IUserEvent);
        public ctx?: (common.IContext|null);
        public account_id: string;
        public user_id: string;
        public event?: (event.IRawEvent|null);
    }

    interface ISubscriptionResponse {
        ctx?: (common.IContext|null);
        status?: (boolean|null);
    }

    class SubscriptionResponse implements ISubscriptionResponse {
        constructor(p?: event.ISubscriptionResponse);
        public ctx?: (common.IContext|null);
        public status: boolean;
    }

    class Publisher extends $protobuf.rpc.Service {
        constructor(rpcImpl: $protobuf.RPCImpl, requestDelimited?: boolean, responseDelimited?: boolean);
        public subscribe(request: event.ISubscription, callback: event.Publisher.SubscribeCallback): void;
        public subscribe(request: event.ISubscription): Promise<event.SubscriptionResponse>;
        public unsubscribe(request: event.ISubscription, callback: event.Publisher.UnsubscribeCallback): void;
        public unsubscribe(request: event.ISubscription): Promise<event.SubscriptionResponse>;
    }

    namespace Publisher {

        type SubscribeCallback = (error: (Error|null), response?: event.SubscriptionResponse) => void;

        type UnsubscribeCallback = (error: (Error|null), response?: event.SubscriptionResponse) => void;
    }

    class EventMgr extends $protobuf.rpc.Service {
        constructor(rpcImpl: $protobuf.RPCImpl, requestDelimited?: boolean, responseDelimited?: boolean);
        public searchEvents(request: event.IListEventsRequest, callback: event.EventMgr.SearchEventsCallback): void;
        public searchEvents(request: event.IListEventsRequest): Promise<event.RawEvents>;
        public createEvent(request: event.IUserEvent, callback: event.EventMgr.CreateEventCallback): void;
        public createEvent(request: event.IUserEvent): Promise<event.RawEvent>;
        public subscribe(request: user.ISubscribeRequest, callback: event.EventMgr.SubscribeCallback): void;
        public subscribe(request: user.ISubscribeRequest): Promise<common.Empty>;
        public unsubscribe(request: user.ISubscribeRequest, callback: event.EventMgr.UnsubscribeCallback): void;
        public unsubscribe(request: user.ISubscribeRequest): Promise<common.Empty>;
        public readTopic(request: user.IReadTopicRequest, callback: event.EventMgr.ReadTopicCallback): void;
        public readTopic(request: user.IReadTopicRequest): Promise<common.Empty>;
        public searchTopics(request: user.IListTopicsRequest, callback: event.EventMgr.SearchTopicsCallback): void;
        public searchTopics(request: user.IListTopicsRequest): Promise<user.ListTopicsResult>;
    }

    namespace EventMgr {

        type SearchEventsCallback = (error: (Error|null), response?: event.RawEvents) => void;

        type CreateEventCallback = (error: (Error|null), response?: event.RawEvent) => void;

        type SubscribeCallback = (error: (Error|null), response?: common.Empty) => void;

        type UnsubscribeCallback = (error: (Error|null), response?: common.Empty) => void;

        type ReadTopicCallback = (error: (Error|null), response?: common.Empty) => void;

        type SearchTopicsCallback = (error: (Error|null), response?: user.ListTopicsResult) => void;
    }

    class ConversationEventReader extends $protobuf.rpc.Service {
        constructor(rpcImpl: $protobuf.RPCImpl, requestDelimited?: boolean, responseDelimited?: boolean);
        public sendMessage(request: event.IRawEvent, callback: event.ConversationEventReader.SendMessageCallback): void;
        public sendMessage(request: event.IRawEvent): Promise<event.RawEvent>;
        public listEvents(request: conversation.IListEventsRequest, callback: event.ConversationEventReader.ListEventsCallback): void;
        public listEvents(request: conversation.IListEventsRequest): Promise<event.RawEvents>;
        public searchEvents(request: conversation.ISearchMessageRequest, callback: event.ConversationEventReader.SearchEventsCallback): void;
        public searchEvents(request: conversation.ISearchMessageRequest): Promise<event.RawEvents>;
    }

    namespace ConversationEventReader {

        type SendMessageCallback = (error: (Error|null), response?: event.RawEvent) => void;

        type ListEventsCallback = (error: (Error|null), response?: event.RawEvents) => void;

        type SearchEventsCallback = (error: (Error|null), response?: event.RawEvents) => void;
    }

    interface IAutomationEvent {
        ctx?: (common.IContext|null);
        account_id?: (string|null);
        user_id?: (string|null);
        automation?: (user.IAutomation|null);
        session?: (user.ISession|null);
        user?: (user.IUser|null);
        event?: (event.IRawEvent|null);
    }

    class AutomationEvent implements IAutomationEvent {
        constructor(p?: event.IAutomationEvent);
        public ctx?: (common.IContext|null);
        public account_id: string;
        public user_id: string;
        public automation?: (user.IAutomation|null);
        public session?: (user.ISession|null);
        public user?: (user.IUser|null);
        public event?: (event.IRawEvent|null);
    }
}

export namespace fabikon {

    interface IFacebookPage {
        account_id?: (string|null);
        id?: (string|null);
        created?: (number|Long|null);
        picture_url?: (string|null);
        name?: (string|null);
        access_token?: (string|null);
    }

    class FacebookPage implements IFacebookPage {
        constructor(p?: fabikon.IFacebookPage);
        public account_id: string;
        public id: string;
        public created: (number|Long);
        public picture_url: string;
        public name: string;
        public access_token: string;
    }

    enum Event {
        FabikonRequested = 1000,
        FabikonSynced = 1001,
        FabikonJob = 2,
        FabikonCreateAccountRequested = 0
    }

    interface IHttpChunk {
        id?: (string|null);
        chunk_id?: (number|null);
        data?: (Uint8Array|null);
        chunk_size?: (number|null);
    }

    class HttpChunk implements IHttpChunk {
        constructor(p?: fabikon.IHttpChunk);
        public id: string;
        public chunk_id: number;
        public data: Uint8Array;
        public chunk_size: number;
    }

    interface IJob {
        topic?: (string|null);
        partition?: (number|null);
        offset?: (number|Long|null);
        type?: (string|null);
        request_id?: (string|null);
    }

    class Job implements IJob {
        constructor(p?: fabikon.IJob);
        public topic: string;
        public partition: number;
        public offset: (number|Long);
        public type: string;
        public request_id: string;
    }

    enum JobType {
        facebook_hook = 0,
        subiz_event = 3
    }

    interface IFbWebhookEvent {
        object?: (string|null);
        entry?: (fabikon.IFbEntry[]|null);
    }

    class FbWebhookEvent implements IFbWebhookEvent {
        constructor(p?: fabikon.IFbWebhookEvent);
        public object: string;
        public entry: fabikon.IFbEntry[];
    }

    interface IFbEntry {
        id?: (string|null);
        time?: (number|Long|null);
        messaging?: (fabikon.IFbMessaging[]|null);
    }

    class FbEntry implements IFbEntry {
        constructor(p?: fabikon.IFbEntry);
        public id: string;
        public time: (number|Long);
        public messaging: fabikon.IFbMessaging[];
    }

    interface IFbMessaging {
        sender?: (fabikon.IFbSender|null);
        recipient?: (fabikon.IFbRecipient|null);
        timestamp?: (number|Long|null);
        message?: (fabikon.IFbMessage|null);
        delivery?: (fabikon.IFbDelivery|null);
        message_type?: (string|null);
    }

    class FbMessaging implements IFbMessaging {
        constructor(p?: fabikon.IFbMessaging);
        public sender?: (fabikon.IFbSender|null);
        public recipient?: (fabikon.IFbRecipient|null);
        public timestamp: (number|Long);
        public message?: (fabikon.IFbMessage|null);
        public delivery?: (fabikon.IFbDelivery|null);
        public message_type: string;
    }

    interface IFbMessage {
        mid?: (string|null);
        text?: (string|null);
        attachments?: (fabikon.IFbAttachment[]|null);
        read?: (fabikon.IFbRead|null);
        attachment?: (fabikon.IFbAttachment|null);
    }

    class FbMessage implements IFbMessage {
        constructor(p?: fabikon.IFbMessage);
        public mid: string;
        public text: string;
        public attachments: fabikon.IFbAttachment[];
        public read?: (fabikon.IFbRead|null);
        public attachment?: (fabikon.IFbAttachment|null);
    }

    interface IFbSender {
        id?: (string|null);
    }

    class FbSender implements IFbSender {
        constructor(p?: fabikon.IFbSender);
        public id: string;
    }

    interface IFbRecipient {
        id?: (string|null);
    }

    class FbRecipient implements IFbRecipient {
        constructor(p?: fabikon.IFbRecipient);
        public id: string;
    }

    interface IFbAttachment {
        type?: (string|null);
        payload?: (fabikon.IFbPayload|null);
        title?: (string|null);
        URL?: (string|null);
    }

    class FbAttachment implements IFbAttachment {
        constructor(p?: fabikon.IFbAttachment);
        public type: string;
        public payload?: (fabikon.IFbPayload|null);
        public title: string;
        public URL: string;
    }

    interface IFbPayload {
        url?: (string|null);
        is_reuseable?: (boolean|null);
    }

    class FbPayload implements IFbPayload {
        constructor(p?: fabikon.IFbPayload);
        public url: string;
        public is_reuseable: boolean;
    }

    interface IFbRead {
        watermark?: (number|Long|null);
        seq?: (number|null);
    }

    class FbRead implements IFbRead {
        constructor(p?: fabikon.IFbRead);
        public watermark: (number|Long);
        public seq: number;
    }

    interface IFbDelivery {
        mids?: (string[]|null);
        watermark?: (number|Long|null);
        seq?: (number|null);
    }

    class FbDelivery implements IFbDelivery {
        constructor(p?: fabikon.IFbDelivery);
        public mids: string[];
        public watermark: (number|Long);
        public seq: number;
    }

    interface IFbSendResponse {
        recipient_id?: (string|null);
        message_id?: (string|null);
    }

    class FbSendResponse implements IFbSendResponse {
        constructor(p?: fabikon.IFbSendResponse);
        public recipient_id: string;
        public message_id: string;
    }

    interface IConversation {
        account_id?: (string|null);
        id?: (string|null);
        page_id?: (string|null);
        created?: (number|Long|null);
        fbuserid?: (string|null);
        sbuserid?: (string|null);
    }

    class Conversation implements IConversation {
        constructor(p?: fabikon.IConversation);
        public account_id: string;
        public id: string;
        public page_id: string;
        public created: (number|Long);
        public fbuserid: string;
        public sbuserid: string;
    }

    interface ICurrentConvo {
        account_id?: (string|null);
        convo_id?: (string|null);
        page_id?: (string|null);
        fbuser_id?: (string|null);
        sbuser_id?: (string|null);
    }

    class CurrentConvo implements ICurrentConvo {
        constructor(p?: fabikon.ICurrentConvo);
        public account_id: string;
        public convo_id: string;
        public page_id: string;
        public fbuser_id: string;
        public sbuser_id: string;
    }

    interface IUserAvail {
        account_id?: (string|null);
        fbuser_id?: (string|null);
        availability?: (boolean|null);
        updated?: (number|Long|null);
    }

    class UserAvail implements IUserAvail {
        constructor(p?: fabikon.IUserAvail);
        public account_id: string;
        public fbuser_id: string;
        public availability: boolean;
        public updated: (number|Long);
    }

    interface IFbPagePicture {
        data?: (fabikon.IFbPagePictureData|null);
    }

    class FbPagePicture implements IFbPagePicture {
        constructor(p?: fabikon.IFbPagePicture);
        public data?: (fabikon.IFbPagePictureData|null);
    }

    interface IFbPagePictureData {
        url?: (string|null);
    }

    class FbPagePictureData implements IFbPagePictureData {
        constructor(p?: fabikon.IFbPagePictureData);
        public url: string;
    }

    interface IFbPageEntry {
        name?: (string|null);
        picture?: (fabikon.IFbPagePicture|null);
        id?: (string|null);
        access_token?: (string|null);
    }

    class FbPageEntry implements IFbPageEntry {
        constructor(p?: fabikon.IFbPageEntry);
        public name: string;
        public picture?: (fabikon.IFbPagePicture|null);
        public id: string;
        public access_token: string;
    }

    interface IFbPageRet {
        data?: (fabikon.IFbPageEntry[]|null);
    }

    class FbPageRet implements IFbPageRet {
        constructor(p?: fabikon.IFbPageRet);
        public data: fabikon.IFbPageEntry[];
    }

    interface IFacebookUser {
        first_name?: (string|null);
        last_name?: (string|null);
        profile_pic?: (string|null);
        locale?: (string|null);
        timezone?: (number|null);
        gender?: (string|null);
    }

    class FacebookUser implements IFacebookUser {
        constructor(p?: fabikon.IFacebookUser);
        public first_name: string;
        public last_name: string;
        public profile_pic: string;
        public locale: string;
        public timezone: number;
        public gender: string;
    }

    interface IFb2SbEvent {
        fbmid?: (string|null);
        account_id?: (string|null);
        conversation_id?: (string|null);
        sbmid?: (string|null);
        page_id?: (string|null);
    }

    class Fb2SbEvent implements IFb2SbEvent {
        constructor(p?: fabikon.IFb2SbEvent);
        public fbmid: string;
        public account_id: string;
        public conversation_id: string;
        public sbmid: string;
        public page_id: string;
    }

    interface ILongLivedAccessToken {
        access_token?: (string|null);
        token_type?: (string|null);
        expires_in?: (number|null);
    }

    class LongLivedAccessToken implements ILongLivedAccessToken {
        constructor(p?: fabikon.ILongLivedAccessToken);
        public access_token: string;
        public token_type: string;
        public expires_in: number;
    }

    interface ISubscribeRet {
        success?: (boolean|null);
    }

    class SubscribeRet implements ISubscribeRet {
        constructor(p?: fabikon.ISubscribeRet);
        public success: boolean;
    }
}

export namespace file {

    interface IAllType {
        fh?: (file.IFileHeader|null);
        ps?: (file.IPresignResult|null);
        f?: (file.IFile|null);
    }

    class AllType implements IAllType {
        constructor(p?: file.IAllType);
        public fh?: (file.IFileHeader|null);
        public ps?: (file.IPresignResult|null);
        public f?: (file.IFile|null);
    }

    class MyServer extends $protobuf.rpc.Service {
        constructor(rpcImpl: $protobuf.RPCImpl, requestDelimited?: boolean, responseDelimited?: boolean);
        public do(request: file.IAllType, callback: file.MyServer.DoCallback): void;
        public do(request: file.IAllType): Promise<file.AllType>;
    }

    namespace MyServer {

        type DoCallback = (error: (Error|null), response?: file.AllType) => void;
    }

    class FileMgr extends $protobuf.rpc.Service {
        constructor(rpcImpl: $protobuf.RPCImpl, requestDelimited?: boolean, responseDelimited?: boolean);
        public presign(request: file.IFileHeader, callback: file.FileMgr.PresignCallback): void;
        public presign(request: file.IFileHeader): Promise<file.PresignResult>;
        public read(request: file.IFileRequest, callback: file.FileMgr.ReadCallback): void;
        public read(request: file.IFileRequest): Promise<file.File>;
        public uploaded(request: file.IFileRequest, callback: file.FileMgr.UploadedCallback): void;
        public uploaded(request: file.IFileRequest): Promise<file.File>;
    }

    namespace FileMgr {

        type PresignCallback = (error: (Error|null), response?: file.PresignResult) => void;

        type ReadCallback = (error: (Error|null), response?: file.File) => void;

        type UploadedCallback = (error: (Error|null), response?: file.File) => void;
    }

    interface IFileHeader {
        ctx?: (common.IContext|null);
        account_id?: (string|null);
        name?: (string|null);
        type?: (string|null);
        size?: (number|Long|null);
        md5?: (string|null);
        description?: (string|null);
    }

    class FileHeader implements IFileHeader {
        constructor(p?: file.IFileHeader);
        public ctx?: (common.IContext|null);
        public account_id: string;
        public name: string;
        public type: string;
        public size: (number|Long);
        public md5: string;
        public description: string;
    }

    interface IPresignResult {
        ctx?: (common.IContext|null);
        account_id?: (string|null);
        url?: (string|null);
        id?: (string|null);
        signed_url?: (string|null);
    }

    class PresignResult implements IPresignResult {
        constructor(p?: file.IPresignResult);
        public ctx?: (common.IContext|null);
        public account_id: string;
        public url: string;
        public id: string;
        public signed_url: string;
    }

    interface IFileRequest {
        ctx?: (common.IContext|null);
        account_id?: (string|null);
        id?: (string|null);
    }

    class FileRequest implements IFileRequest {
        constructor(p?: file.IFileRequest);
        public ctx?: (common.IContext|null);
        public account_id: string;
        public id: string;
    }

    interface IFile {
        ctx?: (common.IContext|null);
        account_id?: (string|null);
        name?: (string|null);
        type?: (string|null);
        size?: (number|Long|null);
        md5?: (string|null);
        description?: (string|null);
        created?: (number|Long|null);
        url?: (string|null);
        creator?: (string|null);
        id?: (string|null);
    }

    class File implements IFile {
        constructor(p?: file.IFile);
        public ctx?: (common.IContext|null);
        public account_id: string;
        public name: string;
        public type: string;
        public size: (number|Long);
        public md5: string;
        public description: string;
        public created: (number|Long);
        public url: string;
        public creator: string;
        public id: string;
    }

    enum Event {
        FilePresignRequested = 0,
        FileCreated = 3,
        FileInfoRequested = 4,
        FileUploaded = 6,
        FileRequested = 10
    }
}

export namespace kv {

    interface IKey {
        partition?: (string|null);
        key?: (string|null);
    }

    class Key implements IKey {
        constructor(p?: kv.IKey);
        public partition: string;
        public key: string;
    }

    interface IValue {
        partition?: (string|null);
        key?: (string|null);
        bytes?: (Uint8Array|null);
        value?: (string|null);
    }

    class Value implements IValue {
        constructor(p?: kv.IValue);
        public partition: string;
        public key: string;
        public bytes: Uint8Array;
        public value: string;
    }

    interface IBool {
        value?: (boolean|null);
    }

    class Bool implements IBool {
        constructor(p?: kv.IBool);
        public value: boolean;
    }

    class KV extends $protobuf.rpc.Service {
        constructor(rpcImpl: $protobuf.RPCImpl, requestDelimited?: boolean, responseDelimited?: boolean);
        public set(request: kv.IValue, callback: kv.KV.SetCallback): void;
        public set(request: kv.IValue): Promise<kv.Value>;
        public get(request: kv.IKey, callback: kv.KV.GetCallback): void;
        public get(request: kv.IKey): Promise<kv.Value>;
        public has(request: kv.IKey, callback: kv.KV.HasCallback): void;
        public has(request: kv.IKey): Promise<kv.Bool>;
    }

    namespace KV {

        type SetCallback = (error: (Error|null), response?: kv.Value) => void;

        type GetCallback = (error: (Error|null), response?: kv.Value) => void;

        type HasCallback = (error: (Error|null), response?: kv.Bool) => void;
    }
}

export namespace lang {

    enum T {
        undefined = 0,
        user_has_already_in_conversation = 1,
        conversation_closed = 2,
        invalid_invite = 3,
        invalid_agent = 4,
        member_is_not_in_conversation = 5,
        conversation_not_found = 6,
        internal_error = 30,
        invalid_input = 22,
        invalid_form = 20,
        access_token_expired = 21,
        credential_not_set = 7,
        wrong_account_in_credential = 8,
        access_deny = 9,
        wrong_user_in_credential = 10,
        unable_to_send_message = 11,
        topic_is_empty = 12,
        invalid_credential = 13,
        invalid_left = 14,
        invalid_json = 15,
        invalid_protobuf = 16,
        invalid_password = 17,
        wrong_password = 18,
        invalid_agent_state = 19,
        unable_to_lock = 40,
        empty = 41,
        wrong_type = 42,
        invalid_kafka_topic = 43,
        database_error = 44,
        timeout = 45,
        websocket_error = 46,
        kafka_error = 47,
        invalid_token = 48,
        account_not_found = 49,
        agent_not_found = 50,
        invalid_email = 60,
        plan_not_found = 61,
        agent_group_not_found = 62,
        empty_client_type = 63,
        empty_url = 64,
        empty_name = 65,
        client_not_found = 66,
        empty_account = 70,
        invalid_conversation_state = 71,
        invalid_message_id = 80,
        invalid_mask = 81,
        randomize_error = 82,
        duplicated_message_received_error = 83,
        network_error = 84,
        rsa_key_not_found = 85,
        jwt_sign_error = 86,
        env_config_error = 87,
        scrypt_error = 90,
        challenge_not_found = 91,
        wrong_answer = 92,
        server_listen_error = 93,
        scrypt_file_not_found = 94,
        invalid_topic = 95,
        file_not_found = 99,
        user_not_found = 100,
        empty_md5 = 101,
        base_convert_error = 102,
        s3_error = 103,
        aws_credential_error = 104,
        aws_send_error = 105,
        elasticsearch_error = 200,
        invalid_template = 203,
        sendgrid_send_error = 204,
        whitelist_domain_not_found = 205,
        blacklist_ip_not_found = 206,
        blocked_user_not_found = 207,
        invalid_content_type = 210,
        parse_file_error = 211,
        invalid_integration_id = 220,
        invalid_integration = 221,
        webhook_not_found = 222,
        tempfile_error = 223,
        write_file_error = 224,
        close_file_error = 225,
        execute_shell_error = 226,
        invalid_css = 227,
        invalid_hmac = 228,
        consul_error = 230,
        maxminddb_err = 231,
        invalid_condition_key = 232,
        invalid_po_file = 233,
        integration_not_found = 234,
        webhook_is_disabled = 235,
        agent_not_activated = 240,
        empty_message = 242,
        message_too_large = 243,
        unknown_message_format = 244,
        too_many_attachments = 245,
        too_many_fields = 246,
        attachment_too_large = 247,
        invalid_end = 248,
        invalid_ack = 249,
        message_not_found = 250,
        invalid_apikey = 255,
        invalid_access_token = 256,
        field_too_long = 257,
        invalid_join = 258,
        automation_not_found = 259,
        invalid_automation_condition = 260,
        automation_cache_miss = 261,
        segment_loop_duplicated = 262,
        segment_loop_stale = 263,
        segment_not_found = 264,
        domain_is_blocked = 265,
        user_is_blocked = 266,
        ip_is_blocked = 267,
        stripe_error = 270,
        missing_stripe_token = 271,
        missing_stripe_customer_id = 272,
        payment_method_not_found = 273,
        not_enough_money = 274,
        invalid_account = 275,
        invalid_payment_method = 276,
        invalid_subscription = 278,
        invalid_plan = 279,
        invoice_not_found = 280,
        billing_not_found = 281,
        exchange_rate_not_found = 282,
        invalid_join_free = 283,
        invalid_active_free = 284,
        kv_not_found = 285,
        attribute_not_found = 286,
        attribute_key_not_found = 287,
        attribute_max_allowed = 288,
        invalid_attribute = 289
    }

    enum L {
        en = 0,
        vn = 1,
        cn = 3,
        fr = 5
    }
}

export namespace logan {

    interface IDebug {
        mem_stats?: (logan.IMemStats|null);
        num_cpu?: (number|null);
        num_goroutine?: (number|null);
        stack_trace?: (Uint8Array|null);
        hostname?: (string|null);
        kafka?: (logan.IKafkaInfo|null);
    }

    class Debug implements IDebug {
        constructor(p?: logan.IDebug);
        public mem_stats?: (logan.IMemStats|null);
        public num_cpu: number;
        public num_goroutine: number;
        public stack_trace: Uint8Array;
        public hostname: string;
        public kafka?: (logan.IKafkaInfo|null);
    }

    interface IKafkaInfo {
        topic?: (string|null);
        partition?: (number|null);
        offset?: (number|Long|null);
    }

    class KafkaInfo implements IKafkaInfo {
        constructor(p?: logan.IKafkaInfo);
        public topic: string;
        public partition: number;
        public offset: (number|Long);
    }

    interface ILog {
        ctx?: (common.IContext|null);
        trace_id?: (string|null);
        created?: (number|Long|null);
        level?: (string|null);
        tags?: (string[]|null);
        debug?: (logan.IDebug|null);
        message?: (Uint8Array|null);
    }

    class Log implements ILog {
        constructor(p?: logan.ILog);
        public ctx?: (common.IContext|null);
        public trace_id: string;
        public created: (number|Long);
        public level: string;
        public tags: string[];
        public debug?: (logan.IDebug|null);
        public message: Uint8Array;
    }

    enum Level {
        debug = 0,
        info = 1,
        notice = 2,
        warning = 3,
        error = 4,
        critical = 5,
        alert = 6,
        emergency = 7,
        panic = 8,
        fatal = 9
    }

    enum Event {
        LogLogRequested = 0,
        LogRequested = 1000,
        LogSynced = 1001
    }

    interface IMemStats {
        alloc?: (number|Long|null);
        total_alloc?: (number|Long|null);
        sys?: (number|Long|null);
        lookups?: (number|Long|null);
        mallocs?: (number|Long|null);
        frees?: (number|Long|null);
        heap_alloc?: (number|Long|null);
        heap_sys?: (number|Long|null);
        heap_idle?: (number|Long|null);
        heap_inuse?: (number|Long|null);
        heap_released?: (number|Long|null);
        heap_objects?: (number|Long|null);
        stack_inuse?: (number|Long|null);
        stack_sys?: (number|Long|null);
        m_span_inuse?: (number|Long|null);
        m_span_sys?: (number|Long|null);
        m_cache_inuse?: (number|Long|null);
        m_cache_sys?: (number|Long|null);
        buck_hash_sys?: (number|Long|null);
        gc_sys?: (number|Long|null);
        other_sys?: (number|Long|null);
        next_gc?: (number|Long|null);
        last_gc?: (number|Long|null);
        pause_total_ns?: (number|Long|null);
        num_gc?: (number|null);
        num_forced_gc?: (number|null);
        gc_cpu_fraction?: (number|null);
    }

    class MemStats implements IMemStats {
        constructor(p?: logan.IMemStats);
        public alloc: (number|Long);
        public total_alloc: (number|Long);
        public sys: (number|Long);
        public lookups: (number|Long);
        public mallocs: (number|Long);
        public frees: (number|Long);
        public heap_alloc: (number|Long);
        public heap_sys: (number|Long);
        public heap_idle: (number|Long);
        public heap_inuse: (number|Long);
        public heap_released: (number|Long);
        public heap_objects: (number|Long);
        public stack_inuse: (number|Long);
        public stack_sys: (number|Long);
        public m_span_inuse: (number|Long);
        public m_span_sys: (number|Long);
        public m_cache_inuse: (number|Long);
        public m_cache_sys: (number|Long);
        public buck_hash_sys: (number|Long);
        public gc_sys: (number|Long);
        public other_sys: (number|Long);
        public next_gc: (number|Long);
        public last_gc: (number|Long);
        public pause_total_ns: (number|Long);
        public num_gc: number;
        public num_forced_gc: number;
        public gc_cpu_fraction: number;
    }
}

export namespace mailkon {

    interface IAddress {
        account_id?: (string|null);
        address?: (string|null);
        updated?: (number|Long|null);
    }

    class Address implements IAddress {
        constructor(p?: mailkon.IAddress);
        public account_id: string;
        public address: string;
        public updated: (number|Long);
    }

    interface IDomain {
        account_id?: (string|null);
        domain?: (string|null);
        updated?: (number|Long|null);
        dnstype?: (string|null);
        data?: (string|null);
    }

    class Domain implements IDomain {
        constructor(p?: mailkon.IDomain);
        public account_id: string;
        public domain: string;
        public updated: (number|Long);
        public dnstype: string;
        public data: string;
    }

    interface IMessage {
        message_id?: (string|null);
        account_id?: (string|null);
        conversation_id?: (string|null);
        from_addr?: (string|null);
        to_addr?: (string|null);
        subject?: (string|null);
        body?: (string|null);
        header?: (string|null);
        attachments?: (string[]|null);
        created?: (number|Long|null);
    }

    class Message implements IMessage {
        constructor(p?: mailkon.IMessage);
        public message_id: string;
        public account_id: string;
        public conversation_id: string;
        public from_addr: string;
        public to_addr: string;
        public subject: string;
        public body: string;
        public header: string;
        public attachments: string[];
        public created: (number|Long);
    }

    interface IUser {
        account_id?: (string|null);
        user_id?: (string|null);
        email_address?: (string|null);
    }

    class User implements IUser {
        constructor(p?: mailkon.IUser);
        public account_id: string;
        public user_id: string;
        public email_address: string;
    }

    enum Event {
        MaikonRequested = 1000,
        MailkonSynced = 1001,
        MailkonJob = 2,
        MailkonCreateAccountRequested = 0
    }

    interface ISendgridEvent {
        event_id?: (string|null);
        account_id?: (string|null);
        conversation_id?: (string|null);
        message_id?: (string|null);
        data?: (string[]|null);
        full_message_id?: (string|null);
        subject?: (string|null);
    }

    class SendgridEvent implements ISendgridEvent {
        constructor(p?: mailkon.ISendgridEvent);
        public event_id: string;
        public account_id: string;
        public conversation_id: string;
        public message_id: string;
        public data: string[];
        public full_message_id: string;
        public subject: string;
    }

    interface IUserAvail {
        account_id?: (string|null);
        email_address?: (string|null);
        availability?: (boolean|null);
        updated?: (number|Long|null);
    }

    class UserAvail implements IUserAvail {
        constructor(p?: mailkon.IUserAvail);
        public account_id: string;
        public email_address: string;
        public availability: boolean;
        public updated: (number|Long);
    }

    interface IHttpChunk {
        id?: (string|null);
        chunk_id?: (number|null);
        data?: (Uint8Array|null);
        chunk_size?: (number|null);
    }

    class HttpChunk implements IHttpChunk {
        constructor(p?: mailkon.IHttpChunk);
        public id: string;
        public chunk_id: number;
        public data: Uint8Array;
        public chunk_size: number;
    }

    interface IJob {
        topic?: (string|null);
        partition?: (number|null);
        offset?: (number|Long|null);
        content_type?: (string|null);
        type?: (string|null);
        request_id?: (string|null);
    }

    class Job implements IJob {
        constructor(p?: mailkon.IJob);
        public topic: string;
        public partition: number;
        public offset: (number|Long);
        public content_type: string;
        public type: string;
        public request_id: string;
    }

    enum JobType {
        sendgrid_inbound = 0,
        sendgrid_event = 1,
        subiz_event = 3
    }
}

export namespace noti5 {

    enum Type {
        new_message = 0,
        new_conversation = 1,
        nothing = 2
    }

    interface ISetting {
        ctx?: (common.IContext|null);
        account_id?: (string|null);
        agent_id?: (string|null);
        web_type?: (string|null);
        mobile_type?: (string|null);
        email_type?: (string|null);
        email_after?: (number|null);
    }

    class Setting implements ISetting {
        constructor(p?: noti5.ISetting);
        public ctx?: (common.IContext|null);
        public account_id: string;
        public agent_id: string;
        public web_type: string;
        public mobile_type: string;
        public email_type: string;
        public email_after: number;
    }

    class Noti5Service extends $protobuf.rpc.Service {
        constructor(rpcImpl: $protobuf.RPCImpl, requestDelimited?: boolean, responseDelimited?: boolean);
        public readNotificationSetting(request: common.IEmpty, callback: noti5.Noti5Service.ReadNotificationSettingCallback): void;
        public readNotificationSetting(request: common.IEmpty): Promise<noti5.Setting>;
        public updateNotificationSetting(request: noti5.ISetting, callback: noti5.Noti5Service.UpdateNotificationSettingCallback): void;
        public updateNotificationSetting(request: noti5.ISetting): Promise<common.Empty>;
    }

    namespace Noti5Service {

        type ReadNotificationSettingCallback = (error: (Error|null), response?: noti5.Setting) => void;

        type UpdateNotificationSettingCallback = (error: (Error|null), response?: common.Empty) => void;
    }
}

export namespace notibox {

    interface INotification {
        ctx?: (common.IContext|null);
        box?: (string|null);
        topic?: (string|null);
        type?: (string|null);
        data?: (string|null);
        created?: (number|Long|null);
        read?: (number|Long|null);
        view?: (boolean|null);
    }

    class Notification implements INotification {
        constructor(p?: notibox.INotification);
        public ctx?: (common.IContext|null);
        public box: string;
        public topic: string;
        public type: string;
        public data: string;
        public created: (number|Long);
        public read: (number|Long);
        public view: boolean;
    }

    enum Type {
        account_created = 0,
        trial_almost_expired = 2,
        trial_expired = 3,
        system_maintainance_scheduled_1 = 4,
        system_maintainance_scheduled_2 = 5,
        system_maintainance_completed = 6,
        agent_activated = 7,
        conversation_unassigned = 8,
        agent_permission_updated = 9
    }

    interface IAddNotificationRequest {
        ctx?: (common.IContext|null);
        boxs?: (string[]|null);
        notification?: (notibox.INotification|null);
    }

    class AddNotificationRequest implements IAddNotificationRequest {
        constructor(p?: notibox.IAddNotificationRequest);
        public ctx?: (common.IContext|null);
        public boxs: string[];
        public notification?: (notibox.INotification|null);
    }

    interface IReadNotification {
        ctx?: (common.IContext|null);
        box?: (string|null);
        topics?: (string[]|null);
        read?: (boolean|null);
    }

    class ReadNotification implements IReadNotification {
        constructor(p?: notibox.IReadNotification);
        public ctx?: (common.IContext|null);
        public box: string;
        public topics: string[];
        public read: boolean;
    }

    interface IListRequest {
        ctx?: (common.IContext|null);
        box?: (string|null);
        anchor?: (string|null);
        limit?: (number|null);
    }

    class ListRequest implements IListRequest {
        constructor(p?: notibox.IListRequest);
        public ctx?: (common.IContext|null);
        public box: string;
        public anchor: string;
        public limit: number;
    }

    interface INotiboxRequest {
        ctx?: (common.IContext|null);
        box?: (string|null);
        topic?: (string|null);
    }

    class NotiboxRequest implements INotiboxRequest {
        constructor(p?: notibox.INotiboxRequest);
        public ctx?: (common.IContext|null);
        public box: string;
        public topic: string;
    }

    interface INotifications {
        ctx?: (common.IContext|null);
        notifications?: (notibox.INotification[]|null);
        anchor?: (string|null);
    }

    class Notifications implements INotifications {
        constructor(p?: notibox.INotifications);
        public ctx?: (common.IContext|null);
        public notifications: notibox.INotification[];
        public anchor: string;
    }

    interface IBox {
        ctx?: (common.IContext|null);
        box?: (string|null);
        new_count?: (number|Long|null);
    }

    class Box implements IBox {
        constructor(p?: notibox.IBox);
        public ctx?: (common.IContext|null);
        public box: string;
        public new_count: (number|Long);
    }

    interface IAllType {
        nts?: (notibox.INotifications|null);
        listr?: (notibox.IListRequest|null);
        rno?: (notibox.IReadNotification|null);
        anorfr?: (notibox.IAddNotificationRequest|null);
        noti?: (notibox.INotification|null);
        box?: (notibox.IBox|null);
    }

    class AllType implements IAllType {
        constructor(p?: notibox.IAllType);
        public nts?: (notibox.INotifications|null);
        public listr?: (notibox.IListRequest|null);
        public rno?: (notibox.IReadNotification|null);
        public anorfr?: (notibox.IAddNotificationRequest|null);
        public noti?: (notibox.INotification|null);
        public box?: (notibox.IBox|null);
    }

    class YourService extends $protobuf.rpc.Service {
        constructor(rpcImpl: $protobuf.RPCImpl, requestDelimited?: boolean, responseDelimited?: boolean);
        public do(request: notibox.IAllType, callback: notibox.YourService.DoCallback): void;
        public do(request: notibox.IAllType): Promise<notibox.AllType>;
    }

    namespace YourService {

        type DoCallback = (error: (Error|null), response?: notibox.AllType) => void;
    }

    class NotiboxService extends $protobuf.rpc.Service {
        constructor(rpcImpl: $protobuf.RPCImpl, requestDelimited?: boolean, responseDelimited?: boolean);
        public upsert(request: notibox.INotification, callback: notibox.NotiboxService.UpsertCallback): void;
        public upsert(request: notibox.INotification): Promise<notibox.Notification>;
        public readTopic(request: notibox.INotiboxRequest, callback: notibox.NotiboxService.ReadTopicCallback): void;
        public readTopic(request: notibox.INotiboxRequest): Promise<common.Empty>;
        public unreadTopic(request: notibox.INotiboxRequest, callback: notibox.NotiboxService.UnreadTopicCallback): void;
        public unreadTopic(request: notibox.INotiboxRequest): Promise<common.Empty>;
        public listNotis(request: notibox.IListRequest, callback: notibox.NotiboxService.ListNotisCallback): void;
        public listNotis(request: notibox.IListRequest): Promise<notibox.Notifications>;
        public decreaseNew(request: notibox.INotiboxRequest, callback: notibox.NotiboxService.DecreaseNewCallback): void;
        public decreaseNew(request: notibox.INotiboxRequest): Promise<notibox.Box>;
        public readBox(request: notibox.INotiboxRequest, callback: notibox.NotiboxService.ReadBoxCallback): void;
        public readBox(request: notibox.INotiboxRequest): Promise<notibox.Box>;
    }

    namespace NotiboxService {

        type UpsertCallback = (error: (Error|null), response?: notibox.Notification) => void;

        type ReadTopicCallback = (error: (Error|null), response?: common.Empty) => void;

        type UnreadTopicCallback = (error: (Error|null), response?: common.Empty) => void;

        type ListNotisCallback = (error: (Error|null), response?: notibox.Notifications) => void;

        type DecreaseNewCallback = (error: (Error|null), response?: notibox.Box) => void;

        type ReadBoxCallback = (error: (Error|null), response?: notibox.Box) => void;
    }

    enum Event {
        NotiboxRequested = 0,
        NotiUpsertRequested = 1,
        NotiReadRequested = 2,
        NotiUnreadRequested = 4,
        NotiListRequested = 5,
        NotiDecreaseNewRequested = 6,
        NotiNewReadRequested = 7,
        NotiboxSynced = 101
    }
}

export namespace payment {

    enum Currency {
        usd = 0,
        vnd = 1,
        brl = 2
    }

    interface IStripe {
        card_last4?: (string|null);
        customer_id?: (string|null);
        token?: (string|null);
    }

    class Stripe implements IStripe {
        constructor(p?: payment.IStripe);
        public card_last4: string;
        public customer_id: string;
        public token: string;
    }

    interface IPaymentMethods {
        payment_methods?: (payment.IPaymentMethod[]|null);
    }

    class PaymentMethods implements IPaymentMethods {
        constructor(p?: payment.IPaymentMethods);
        public payment_methods: payment.IPaymentMethod[];
    }

    interface IPaymentMethod {
        ctx?: (common.IContext|null);
        type?: (string|null);
        id?: (string|null);
        account_id?: (string|null);
        state?: (string|null);
        created?: (number|Long|null);
        stripe?: (payment.IStripe|null);
        failed_message?: (string|null);
        charged?: (number|Long|null);
    }

    class PaymentMethod implements IPaymentMethod {
        constructor(p?: payment.IPaymentMethod);
        public ctx?: (common.IContext|null);
        public type: string;
        public id: string;
        public account_id: string;
        public state: string;
        public created: (number|Long);
        public stripe?: (payment.IStripe|null);
        public failed_message: string;
        public charged: (number|Long);
    }

    namespace PaymentMethod {

        enum Type {
            bank_transfer = 0,
            credit_card = 1
        }

        enum State {
            active = 0,
            failed = 1
        }
    }

    interface ILimit {
        ctx?: (common.IContext|null);
        account_id?: (string|null);
        max_automations?: (number|null);
        max_segments?: (number|null);
        max_agents?: (number|null);
        can_buy_agent?: (boolean|null);
        automation_webhook_quota?: (number|null);
        automation_email_quota?: (number|null);
        automation_message_quota?: (number|null);
        max_rules?: (number|null);
    }

    class Limit implements ILimit {
        constructor(p?: payment.ILimit);
        public ctx?: (common.IContext|null);
        public account_id: string;
        public max_automations: number;
        public max_segments: number;
        public max_agents: number;
        public can_buy_agent: boolean;
        public automation_webhook_quota: number;
        public automation_email_quota: number;
        public automation_message_quota: number;
        public max_rules: number;
    }

    interface IPlans {
        plans?: (payment.IPlan[]|null);
    }

    class Plans implements IPlans {
        constructor(p?: payment.IPlans);
        public plans: payment.IPlan[];
    }

    interface IPlan {
        name?: (string|null);
        limit?: (payment.ILimit|null);
        price?: (number|null);
        level?: (number|null);
        can_buy_agent?: (boolean|null);
        can_buy?: (boolean|null);
        has_start_time?: (boolean|null);
    }

    class Plan implements IPlan {
        constructor(p?: payment.IPlan);
        public name: string;
        public limit?: (payment.ILimit|null);
        public price: number;
        public level: number;
        public can_buy_agent: boolean;
        public can_buy: boolean;
        public has_start_time: boolean;
    }

    namespace Plan {

        enum Type {
            trial = 0,
            free = 1,
            standard = 2,
            advanced = 3
        }
    }

    interface ISubscription {
        ctx?: (common.IContext|null);
        account_id?: (string|null);
        created?: (number|Long|null);
        promotion_code?: (string|null);
        name?: (string|null);
        started?: (number|Long|null);
        due_date?: (number|Long|null);
        auto_renew?: (boolean|null);
        billing_cycle_month?: (number|null);
        next_billing_cycle_month?: (number|null);
        plan?: (string|null);
        addons?: (payment.IAddon[]|null);
        credit?: (number|null);
        notes?: (payment.INote[]|null);
        referral_by?: (string|null);
        customer?: (payment.ICustomer|null);
        primary_payment_method?: (string|null);
        limit?: (payment.ILimit|null);
        v3_state?: (number|null);
    }

    class Subscription implements ISubscription {
        constructor(p?: payment.ISubscription);
        public ctx?: (common.IContext|null);
        public account_id: string;
        public created: (number|Long);
        public promotion_code: string;
        public name: string;
        public started: (number|Long);
        public due_date: (number|Long);
        public auto_renew: boolean;
        public billing_cycle_month: number;
        public next_billing_cycle_month: number;
        public plan: string;
        public addons: payment.IAddon[];
        public credit: number;
        public notes: payment.INote[];
        public referral_by: string;
        public customer?: (payment.ICustomer|null);
        public primary_payment_method: string;
        public limit?: (payment.ILimit|null);
        public v3_state: number;
    }

    interface IBill {
        ctx?: (common.IContext|null);
        id?: (string|null);
        account_id?: (string|null);
        amount?: (number|null);
        invoice_ids?: (string[]|null);
        created?: (number|Long|null);
        customer_info?: (payment.IContact|null);
        payment_method?: (string|null);
        year?: (number|null);
        description?: (string|null);
    }

    class Bill implements IBill {
        constructor(p?: payment.IBill);
        public ctx?: (common.IContext|null);
        public id: string;
        public account_id: string;
        public amount: number;
        public invoice_ids: string[];
        public created: (number|Long);
        public customer_info?: (payment.IContact|null);
        public payment_method: string;
        public year: number;
        public description: string;
    }

    interface INote {
        message?: (string|null);
        creator?: (string|null);
        created?: (number|Long|null);
    }

    class Note implements INote {
        constructor(p?: payment.INote);
        public message: string;
        public creator: string;
        public created: (number|Long);
    }

    interface IInvoices {
        invoices?: (payment.IInvoice[]|null);
    }

    class Invoices implements IInvoices {
        constructor(p?: payment.IInvoices);
        public invoices: payment.IInvoice[];
    }

    interface IInvoice {
        ctx?: (common.IContext|null);
        account_id?: (string|null);
        id?: (string|null);
        amount_due?: (number|null);
        promotion_code?: (string|null);
        description?: (string|null);
        billing_info?: (payment.IBillingInfo|null);
        due_date?: (number|Long|null);
        state?: (string|null);
        created?: (number|Long|null);
        items?: (payment.IInvoiceItem[]|null);
        subtotal?: (number|null);
        tax_percent?: (number|null);
        tax?: (number|null);
        total?: (number|null);
        updated?: (number|Long|null);
        year?: (number|null);
        notes?: (payment.INote[]|null);
        bills?: (string[]|null);
        payment_made?: (number|null);
        current_sub?: (payment.ISubscription|null);
        current_plan?: (payment.IPlan|null);
    }

    class Invoice implements IInvoice {
        constructor(p?: payment.IInvoice);
        public ctx?: (common.IContext|null);
        public account_id: string;
        public id: string;
        public amount_due: number;
        public promotion_code: string;
        public description: string;
        public billing_info?: (payment.IBillingInfo|null);
        public due_date: (number|Long);
        public state: string;
        public created: (number|Long);
        public items: payment.IInvoiceItem[];
        public subtotal: number;
        public tax_percent: number;
        public tax: number;
        public total: number;
        public updated: (number|Long);
        public year: number;
        public notes: payment.INote[];
        public bills: string[];
        public payment_made: number;
        public current_sub?: (payment.ISubscription|null);
        public current_plan?: (payment.IPlan|null);
    }

    namespace Invoice {

        enum State {
            draft = 0,
            open = 1,
            overdue = 2,
            paid = 3,
            voided = 4,
            queueing = 5
        }
    }

    interface IAgentInvoiceItem {
        plan?: (string|null);
        day_left?: (number|null);
        agent_count?: (number|null);
    }

    class AgentInvoiceItem implements IAgentInvoiceItem {
        constructor(p?: payment.IAgentInvoiceItem);
        public plan: string;
        public day_left: number;
        public agent_count: number;
    }

    interface IRenewInvoiceItem {
        plan?: (string|null);
        billing_cycle_month?: (number|null);
        agent_count?: (number|null);
        from_time?: (number|Long|null);
    }

    class RenewInvoiceItem implements IRenewInvoiceItem {
        constructor(p?: payment.IRenewInvoiceItem);
        public plan: string;
        public billing_cycle_month: number;
        public agent_count: number;
        public from_time: (number|Long);
    }

    interface IPlanInvoiceItem {
        agent_count?: (number|null);
        billing_cycle_month?: (number|null);
        old_plan?: (string|null);
        new_plan?: (string|null);
        save_percentage?: (number|null);
        started?: (number|Long|null);
        day_left?: (number|null);
    }

    class PlanInvoiceItem implements IPlanInvoiceItem {
        constructor(p?: payment.IPlanInvoiceItem);
        public agent_count: number;
        public billing_cycle_month: number;
        public old_plan: string;
        public new_plan: string;
        public save_percentage: number;
        public started: (number|Long);
        public day_left: number;
    }

    interface IInvoiceItem {
        description?: (string|null);
        invoice_id?: (string|null);
        quantity?: (number|null);
        price?: (number|null);
        data?: (payment.InvoiceItem.IData|null);
    }

    class InvoiceItem implements IInvoiceItem {
        constructor(p?: payment.IInvoiceItem);
        public description: string;
        public invoice_id: string;
        public quantity: number;
        public price: number;
        public data?: (payment.InvoiceItem.IData|null);
    }

    namespace InvoiceItem {

        interface IData {
            renew?: (payment.IRenewInvoiceItem|null);
            agent?: (payment.IAgentInvoiceItem|null);
            plan?: (payment.IPlanInvoiceItem|null);
        }

        class Data implements IData {
            constructor(p?: payment.InvoiceItem.IData);
            public renew?: (payment.IRenewInvoiceItem|null);
            public agent?: (payment.IAgentInvoiceItem|null);
            public plan?: (payment.IPlanInvoiceItem|null);
        }
    }

    interface IBillingInfo {
        name?: (string|null);
        address?: (string|null);
        vat?: (string|null);
        country_code?: (string|null);
    }

    class BillingInfo implements IBillingInfo {
        constructor(p?: payment.IBillingInfo);
        public name: string;
        public address: string;
        public vat: string;
        public country_code: string;
    }

    interface IContact {
        ctx?: (common.IContext|null);
        name?: (string|null);
        email?: (string|null);
        phone?: (string|null);
        job_title?: (string|null);
        title?: (string|null);
        primary?: (boolean|null);
    }

    class Contact implements IContact {
        constructor(p?: payment.IContact);
        public ctx?: (common.IContext|null);
        public name: string;
        public email: string;
        public phone: string;
        public job_title: string;
        public title: string;
        public primary: boolean;
    }

    namespace Contact {

        enum Title {
            mr = 0,
            ms = 1,
            mrs = 2,
            dr = 3
        }
    }

    interface ICustomer {
        id?: (string|null);
        account_id?: (string|null);
        contacts?: (payment.IContact[]|null);
        created?: (number|Long|null);
        billing_info?: (payment.IBillingInfo|null);
    }

    class Customer implements ICustomer {
        constructor(p?: payment.ICustomer);
        public id: string;
        public account_id: string;
        public contacts: payment.IContact[];
        public created: (number|Long);
        public billing_info?: (payment.IBillingInfo|null);
    }

    interface IFixedAmountPromotionCode {
        amount?: (number|null);
    }

    class FixedAmountPromotionCode implements IFixedAmountPromotionCode {
        constructor(p?: payment.IFixedAmountPromotionCode);
        public amount: number;
    }

    interface IPercentPromotionCode {
        percent?: (number|null);
    }

    class PercentPromotionCode implements IPercentPromotionCode {
        constructor(p?: payment.IPercentPromotionCode);
        public percent: number;
    }

    interface ICreditCode {
        credit?: (number|null);
    }

    class CreditCode implements ICreditCode {
        constructor(p?: payment.ICreditCode);
        public credit: number;
    }

    interface IReferralCreditCode {
        referrer_id?: (string|null);
        credit?: (number|null);
    }

    class ReferralCreditCode implements IReferralCreditCode {
        constructor(p?: payment.IReferralCreditCode);
        public referrer_id: string;
        public credit: number;
    }

    interface IPromotionCode {
        ctx?: (common.IContext|null);
        description?: (string|null);
        type?: (string|null);
        redeem_count?: (number|null);
        creator?: (string|null);
        created?: (number|Long|null);
        code?: (string|null);
        data?: (payment.PromotionCode.IData|null);
        start?: (number|Long|null);
        end?: (number|Long|null);
        for_plan?: (string|null);
        for_account_id?: (string|null);
        max_redemptions?: (number|null);
        addon?: (string|null);
        for_items?: (string[]|null);
        min_amount?: (number|null);
        max_amount?: (number|null);
    }

    class PromotionCode implements IPromotionCode {
        constructor(p?: payment.IPromotionCode);
        public ctx?: (common.IContext|null);
        public description: string;
        public type: string;
        public redeem_count: number;
        public creator: string;
        public created: (number|Long);
        public code: string;
        public data?: (payment.PromotionCode.IData|null);
        public start: (number|Long);
        public end: (number|Long);
        public for_plan: string;
        public for_account_id: string;
        public max_redemptions: number;
        public addon: string;
        public for_items: string[];
        public min_amount: number;
        public max_amount: number;
    }

    namespace PromotionCode {

        enum Type {
            fixed_amount_promotion_code = 0,
            percent_promotion_code = 1,
            credit_code = 2,
            referral_credit_code = 3
        }

        interface IData {
            fixed_amount?: (payment.IFixedAmountPromotionCode|null);
            percent?: (payment.IPercentPromotionCode|null);
            credit?: (payment.ICreditCode|null);
            referral?: (payment.IReferralCreditCode|null);
        }

        class Data implements IData {
            constructor(p?: payment.PromotionCode.IData);
            public fixed_amount?: (payment.IFixedAmountPromotionCode|null);
            public percent?: (payment.IPercentPromotionCode|null);
            public credit?: (payment.ICreditCode|null);
            public referral?: (payment.IReferralCreditCode|null);
        }
    }

    interface IExchangeRate {
        from_currency?: (string|null);
        to_currency?: (string|null);
        exchange_rate?: (number|null);
        exchange_time?: (number|Long|null);
    }

    class ExchangeRate implements IExchangeRate {
        constructor(p?: payment.IExchangeRate);
        public from_currency: string;
        public to_currency: string;
        public exchange_rate: number;
        public exchange_time: (number|Long);
    }

    interface ILog {
        ctx?: (common.IContext|null);
        user?: (string|null);
        id?: (string|null);
        action?: (string|null);
        created?: (number|Long|null);
        description?: (string|null);
        account_id?: (string|null);
        month?: (number|null);
    }

    class Log implements ILog {
        constructor(p?: payment.ILog);
        public ctx?: (common.IContext|null);
        public user: string;
        public id: string;
        public action: string;
        public created: (number|Long);
        public description: string;
        public account_id: string;
        public month: number;
    }

    namespace Log {

        enum Action {
            create_invoice = 0,
            change_invoice_status = 1,
            create_discount = 2,
            delete_discount = 3,
            redeem_discount = 4,
            add_credit = 5,
            redeem_credit = 6,
            delete_account = 7,
            change_plan = 8,
            renew_subscription = 10,
            click_subscribe_button = 11,
            pay_for_referrer = 12,
            add_money_for_referrer = 13,
            pay_invoice = 14
        }
    }

    interface IAddon {
        type?: (string|null);
        name?: (string|null);
        price?: (number|null);
        currency?: (payment.Currency|null);
        charge_type?: (string|null);
        period?: (number|null);
        period_unit?: (string|null);
        quantity?: (number|null);
        created?: (number|Long|null);
    }

    class Addon implements IAddon {
        constructor(p?: payment.IAddon);
        public type: string;
        public name: string;
        public price: number;
        public currency: payment.Currency;
        public charge_type: string;
        public period: number;
        public period_unit: string;
        public quantity: number;
        public created: (number|Long);
    }

    namespace Addon {

        enum Type {
            credit = 0,
            agent = 1
        }

        enum ChargeType {
            one_time = 0,
            recurring = 1
        }

        enum PeriodUnit {
            day = 0,
            week = 1,
            month = 2,
            year = 3
        }
    }

    enum Event {
        PaymentSynced = 0,
        LimitUpdated = 1,
        PaymentRequested = 4,
        PaymentRenewCycle = 5,
        InvoiceCreatedEmailSend = 6,
        PaymentV3Synced = 8,
        SubscriptionUpdated = 14,
        InvoiceUpdated = 10,
        PaymentMethodUpdated = 11,
        BillingUpdated = 12,
        LogUpdated = 13,
        PromotionCodeUpdated = 15
    }

    interface IPaymentRenewCycleRequested {
        ctx?: (common.IContext|null);
        account_id?: (string|null);
        sub?: (payment.ISubscription|null);
        cycle_id?: (string|null);
    }

    class PaymentRenewCycleRequested implements IPaymentRenewCycleRequested {
        constructor(p?: payment.IPaymentRenewCycleRequested);
        public ctx?: (common.IContext|null);
        public account_id: string;
        public sub?: (payment.ISubscription|null);
        public cycle_id: string;
    }

    class SubizInternalPaymentMgr extends $protobuf.rpc.Service {
        constructor(rpcImpl: $protobuf.RPCImpl, requestDelimited?: boolean, responseDelimited?: boolean);
        public createBill(request: payment.IBill, callback: payment.SubizInternalPaymentMgr.CreateBillCallback): void;
        public createBill(request: payment.IBill): Promise<payment.Bill>;
        public updateExchangeRate(request: payment.IExchangeRate, callback: payment.SubizInternalPaymentMgr.UpdateExchangeRateCallback): void;
        public updateExchangeRate(request: payment.IExchangeRate): Promise<payment.ExchangeRate>;
    }

    namespace SubizInternalPaymentMgr {

        type CreateBillCallback = (error: (Error|null), response?: payment.Bill) => void;

        type UpdateExchangeRateCallback = (error: (Error|null), response?: payment.ExchangeRate) => void;
    }

    class PaymentMgr extends $protobuf.rpc.Service {
        constructor(rpcImpl: $protobuf.RPCImpl, requestDelimited?: boolean, responseDelimited?: boolean);
        public purchase(request: payment.ISubscription, callback: payment.PaymentMgr.PurchaseCallback): void;
        public purchase(request: payment.ISubscription): Promise<payment.Subscription>;
        public updateSubscription(request: payment.ISubscription, callback: payment.PaymentMgr.UpdateSubscriptionCallback): void;
        public updateSubscription(request: payment.ISubscription): Promise<payment.Subscription>;
        public getSubscription(request: common.IEmpty, callback: payment.PaymentMgr.GetSubscriptionCallback): void;
        public getSubscription(request: common.IEmpty): Promise<payment.Subscription>;
        public getPromotionCode(request: payment.IString, callback: payment.PaymentMgr.GetPromotionCodeCallback): void;
        public getPromotionCode(request: payment.IString): Promise<payment.PromotionCode>;
        public addPaymentMethod(request: payment.IPaymentMethod, callback: payment.PaymentMgr.AddPaymentMethodCallback): void;
        public addPaymentMethod(request: payment.IPaymentMethod): Promise<payment.PaymentMethod>;
        public updatePaymentMethod(request: payment.IPaymentMethod, callback: payment.PaymentMgr.UpdatePaymentMethodCallback): void;
        public updatePaymentMethod(request: payment.IPaymentMethod): Promise<payment.PaymentMethod>;
        public deletePaymentMethod(request: common.IId, callback: payment.PaymentMgr.DeletePaymentMethodCallback): void;
        public deletePaymentMethod(request: common.IId): Promise<common.Empty>;
        public listPaymentMethods(request: common.IEmpty, callback: payment.PaymentMgr.ListPaymentMethodsCallback): void;
        public listPaymentMethods(request: common.IEmpty): Promise<payment.PaymentMethods>;
        public pay(request: payment.IPayRequest, callback: payment.PaymentMgr.PayCallback): void;
        public pay(request: payment.IPayRequest): Promise<payment.Bill>;
        public listInvoices(request: common.IId, callback: payment.PaymentMgr.ListInvoicesCallback): void;
        public listInvoices(request: common.IId): Promise<payment.Invoices>;
        public createInvoice(request: payment.IInvoice, callback: payment.PaymentMgr.CreateInvoiceCallback): void;
        public createInvoice(request: payment.IInvoice): Promise<payment.Invoice>;
        public listPlans(request: common.IEmpty, callback: payment.PaymentMgr.ListPlansCallback): void;
        public listPlans(request: common.IEmpty): Promise<payment.Plans>;
        public exportInvoice(request: common.IId, callback: payment.PaymentMgr.ExportInvoiceCallback): void;
        public exportInvoice(request: common.IId): Promise<payment.String>;
        public getExchangeRate(request: payment.IExchangeRate, callback: payment.PaymentMgr.GetExchangeRateCallback): void;
        public getExchangeRate(request: payment.IExchangeRate): Promise<payment.ExchangeRate>;
        public transferMoney(request: payment.IPayRequest, callback: payment.PaymentMgr.TransferMoneyCallback): void;
        public transferMoney(request: payment.IPayRequest): Promise<payment.Bill>;
    }

    namespace PaymentMgr {

        type PurchaseCallback = (error: (Error|null), response?: payment.Subscription) => void;

        type UpdateSubscriptionCallback = (error: (Error|null), response?: payment.Subscription) => void;

        type GetSubscriptionCallback = (error: (Error|null), response?: payment.Subscription) => void;

        type GetPromotionCodeCallback = (error: (Error|null), response?: payment.PromotionCode) => void;

        type AddPaymentMethodCallback = (error: (Error|null), response?: payment.PaymentMethod) => void;

        type UpdatePaymentMethodCallback = (error: (Error|null), response?: payment.PaymentMethod) => void;

        type DeletePaymentMethodCallback = (error: (Error|null), response?: common.Empty) => void;

        type ListPaymentMethodsCallback = (error: (Error|null), response?: payment.PaymentMethods) => void;

        type PayCallback = (error: (Error|null), response?: payment.Bill) => void;

        type ListInvoicesCallback = (error: (Error|null), response?: payment.Invoices) => void;

        type CreateInvoiceCallback = (error: (Error|null), response?: payment.Invoice) => void;

        type ListPlansCallback = (error: (Error|null), response?: payment.Plans) => void;

        type ExportInvoiceCallback = (error: (Error|null), response?: payment.String) => void;

        type GetExchangeRateCallback = (error: (Error|null), response?: payment.ExchangeRate) => void;

        type TransferMoneyCallback = (error: (Error|null), response?: payment.Bill) => void;
    }

    interface IString {
        str?: (string|null);
    }

    class String implements IString {
        constructor(p?: payment.IString);
        public str: string;
    }

    interface IPayRequest {
        account_id?: (string|null);
        invoice_ids?: (string[]|null);
        description?: (string|null);
        CustomerInfo?: (payment.IContact|null);
        amount?: (number|null);
    }

    class PayRequest implements IPayRequest {
        constructor(p?: payment.IPayRequest);
        public account_id: string;
        public invoice_ids: string[];
        public description: string;
        public CustomerInfo?: (payment.IContact|null);
        public amount: number;
    }

    interface IESubscription {
        sub?: (payment.ISubscription|null);
        err?: (string|null);
    }

    class ESubscription implements IESubscription {
        constructor(p?: payment.IESubscription);
        public sub?: (payment.ISubscription|null);
        public err: string;
    }

    interface IEInvoice {
        inv?: (payment.IInvoice|null);
        err?: (string|null);
    }

    class EInvoice implements IEInvoice {
        constructor(p?: payment.IEInvoice);
        public inv?: (payment.IInvoice|null);
        public err: string;
    }

    interface IInvoiceCreatedEmail {
        ctx?: (common.IContext|null);
        to?: (string|null);
        account_id?: (string|null);
        billing_name?: (string|null);
        invoice_id?: (string|null);
        created?: (number|Long|null);
        lang?: (lang.L|null);
        from?: (string|null);
    }

    class InvoiceCreatedEmail implements IInvoiceCreatedEmail {
        constructor(p?: payment.IInvoiceCreatedEmail);
        public ctx?: (common.IContext|null);
        public to: string;
        public account_id: string;
        public billing_name: string;
        public invoice_id: string;
        public created: (number|Long);
        public lang: lang.L;
        public from: string;
    }
}

export namespace pubsub {

    interface ISubscription {
        ctx?: (common.IContext|null);
        user_id?: (string|null);
        topic?: (string|null);
        sub_id?: (string|null);
        target_topic?: (string|null);
        target_key?: (string|null);
        ttls?: (number|Long|null);
        router_topic?: (string|null);
        target_partition?: (number|null);
        topics?: (string[]|null);
    }

    class Subscription implements ISubscription {
        constructor(p?: pubsub.ISubscription);
        public ctx?: (common.IContext|null);
        public user_id: string;
        public topic: string;
        public sub_id: string;
        public target_topic: string;
        public target_key: string;
        public ttls: (number|Long);
        public router_topic: string;
        public target_partition: number;
        public topics: string[];
    }

    interface IPublishMessage {
        ctx?: (common.IContext|null);
        topics?: (string[]|null);
        payload?: (Uint8Array|null);
        user_ids_filter?: (Uint8Array|null);
        neg_user_ids_filter?: (Uint8Array|null);
    }

    class PublishMessage implements IPublishMessage {
        constructor(p?: pubsub.IPublishMessage);
        public ctx?: (common.IContext|null);
        public topics: string[];
        public payload: Uint8Array;
        public user_ids_filter: Uint8Array;
        public neg_user_ids_filter: Uint8Array;
    }

    class Pubsub extends $protobuf.rpc.Service {
        constructor(rpcImpl: $protobuf.RPCImpl, requestDelimited?: boolean, responseDelimited?: boolean);
        public subscribe(request: pubsub.ISubscription, callback: pubsub.Pubsub.SubscribeCallback): void;
        public subscribe(request: pubsub.ISubscription): Promise<common.Empty>;
        public unsubscribe(request: pubsub.ISubscription, callback: pubsub.Pubsub.UnsubscribeCallback): void;
        public unsubscribe(request: pubsub.ISubscription): Promise<common.Empty>;
        public publish(request: pubsub.IPublishMessage, callback: pubsub.Pubsub.PublishCallback): void;
        public publish(request: pubsub.IPublishMessage): Promise<common.Empty>;
    }

    namespace Pubsub {

        type SubscribeCallback = (error: (Error|null), response?: common.Empty) => void;

        type UnsubscribeCallback = (error: (Error|null), response?: common.Empty) => void;

        type PublishCallback = (error: (Error|null), response?: common.Empty) => void;
    }
}

export namespace scheduler {

    interface ITask {
        id?: (string|null);
        callback_time?: (number|Long|null);
        topic?: (string|null);
        data?: (Uint8Array|null);
        key?: (string|null);
        called?: (number|Long|null);
        sec?: (number|Long|null);
        par?: (string|null);
    }

    class Task implements ITask {
        constructor(p?: scheduler.ITask);
        public id: string;
        public callback_time: (number|Long);
        public topic: string;
        public data: Uint8Array;
        public key: string;
        public called: (number|Long);
        public sec: (number|Long);
        public par: string;
    }

    interface IId {
        id?: (string|null);
        callback_time?: (number|Long|null);
    }

    class Id implements IId {
        constructor(p?: scheduler.IId);
        public id: string;
        public callback_time: (number|Long);
    }

    enum Event {
        SchedulerRequested = 0,
        SchedulerCancelled = 1
    }
}

export namespace template {

    interface ITemplate {
        Id?: (string|null);
        Language?: (string|null);
        Type?: (template.Type|null);
        Email?: (template.IEmailTemplate|null);
        WebPush?: (template.IWebPushTemplate|null);
        Notification?: (template.INotificationTemplate|null);
    }

    class Template implements ITemplate {
        constructor(p?: template.ITemplate);
        public Id: string;
        public Language: string;
        public Type: template.Type;
        public Email?: (template.IEmailTemplate|null);
        public WebPush?: (template.IWebPushTemplate|null);
        public Notification?: (template.INotificationTemplate|null);
    }

    enum Type {
        Email = 0,
        WebPush = 1,
        Notification = 2
    }

    interface IEmailTemplate {
        Subject?: (string|null);
        Body?: (string|null);
    }

    class EmailTemplate implements IEmailTemplate {
        constructor(p?: template.IEmailTemplate);
        public Subject: string;
        public Body: string;
    }

    interface IWebPushTemplate {
        Title?: (string|null);
        Body?: (string|null);
        Icon?: (string|null);
    }

    class WebPushTemplate implements IWebPushTemplate {
        constructor(p?: template.IWebPushTemplate);
        public Title: string;
        public Body: string;
        public Icon: string;
    }

    interface INotificationTemplate {
        Image?: (string|null);
        Footer?: (string|null);
        Body?: (string|null);
    }

    class NotificationTemplate implements INotificationTemplate {
        constructor(p?: template.INotificationTemplate);
        public Image: string;
        public Footer: string;
        public Body: string;
    }
}

export namespace user {

    interface IAllType {
        user?: (user.IUser|null);
        cr?: (user.ICreateRequest|null);
        fields?: (user.IField|null);
        ucr?: (user.IUserCreateResult|null);
        rpr?: (user.IReadTopicRequest|null);
        usersearchrequest?: (user.IUserSearchRequest|null);
        usersearchresult?: (user.IUserSearchResult|null);
        ms?: (user.IMaskResponse|null);
        ltr?: (user.IListTopicsResult|null);
        sg?: (user.ISegmentation|null);
        presn?: (user.IPresence|null);
        vi?: (user.IVisitor|null);
        vis?: (user.IVisitors|null);
        topic?: (user.ITopic|null);
        utopic?: (user.IUnreadTopic|null);
        myuser?: (user.IMyUser|null);
        automation?: (user.IAutomation|null);
        automations?: (user.IAutomations|null);
        session?: (user.ISession|null);
    }

    class AllType implements IAllType {
        constructor(p?: user.IAllType);
        public user?: (user.IUser|null);
        public cr?: (user.ICreateRequest|null);
        public fields?: (user.IField|null);
        public ucr?: (user.IUserCreateResult|null);
        public rpr?: (user.IReadTopicRequest|null);
        public usersearchrequest?: (user.IUserSearchRequest|null);
        public usersearchresult?: (user.IUserSearchResult|null);
        public ms?: (user.IMaskResponse|null);
        public ltr?: (user.IListTopicsResult|null);
        public sg?: (user.ISegmentation|null);
        public presn?: (user.IPresence|null);
        public vi?: (user.IVisitor|null);
        public vis?: (user.IVisitors|null);
        public topic?: (user.ITopic|null);
        public utopic?: (user.IUnreadTopic|null);
        public myuser?: (user.IMyUser|null);
        public automation?: (user.IAutomation|null);
        public automations?: (user.IAutomations|null);
        public session?: (user.ISession|null);
    }

    class MyServer extends $protobuf.rpc.Service {
        constructor(rpcImpl: $protobuf.RPCImpl, requestDelimited?: boolean, responseDelimited?: boolean);
        public do(request: user.IAllType, callback: user.MyServer.DoCallback): void;
        public do(request: user.IAllType): Promise<user.AllType>;
    }

    namespace MyServer {

        type DoCallback = (error: (Error|null), response?: user.AllType) => void;
    }

    interface IAddToMyRequest {
        ctx?: (common.IContext|null);
        user_id?: (string|null);
        agent_ids?: (string[]|null);
    }

    class AddToMyRequest implements IAddToMyRequest {
        constructor(p?: user.IAddToMyRequest);
        public ctx?: (common.IContext|null);
        public user_id: string;
        public agent_ids: string[];
    }

    interface IUserCreateResult {
        id?: (string|null);
        mask?: (string|null);
    }

    class UserCreateResult implements IUserCreateResult {
        constructor(p?: user.IUserCreateResult);
        public id: string;
        public mask: string;
    }

    interface IMyUser {
        ctx?: (common.IContext|null);
        agent_id?: (string|null);
        user?: (user.IUser|null);
        unread?: (number|null);
        updated?: (number|Long|null);
    }

    class MyUser implements IMyUser {
        constructor(p?: user.IMyUser);
        public ctx?: (common.IContext|null);
        public agent_id: string;
        public user?: (user.IUser|null);
        public unread: number;
        public updated: (number|Long);
    }

    enum AttributeType {
        text = 0,
        number = 1,
        boolean = 2,
        datetime = 3,
        list = 4
    }

    enum AttributeProtected {
        id = 0,
        account_id = 1,
        fullname = 2,
        emails = 3,
        phones = 4
    }

    interface IAttributeDefinition {
        ctx?: (common.IContext|null);
        account_id?: (string|null);
        name?: (string|null);
        description?: (string|null);
        type?: (string|null);
        key?: (string|null);
        updated?: (number|Long|null);
    }

    class AttributeDefinition implements IAttributeDefinition {
        constructor(p?: user.IAttributeDefinition);
        public ctx?: (common.IContext|null);
        public account_id: string;
        public name: string;
        public description: string;
        public type: string;
        public key: string;
        public updated: (number|Long);
    }

    interface IAttributeDefinitions {
        ctx?: (common.IContext|null);
        attributes?: (user.IAttributeDefinition[]|null);
    }

    class AttributeDefinitions implements IAttributeDefinitions {
        constructor(p?: user.IAttributeDefinitions);
        public ctx?: (common.IContext|null);
        public attributes: user.IAttributeDefinition[];
    }

    enum AttributeDataState {
        live = 0,
        deleted = 1
    }

    interface IAttributeData {
        ctx?: (common.IContext|null);
        account_id?: (string|null);
        user_id?: (string|null);
        key?: (string|null);
        value?: (string|null);
        state?: (string|null);
        created?: (number|Long|null);
        modified?: (number|Long|null);
    }

    class AttributeData implements IAttributeData {
        constructor(p?: user.IAttributeData);
        public ctx?: (common.IContext|null);
        public account_id: string;
        public user_id: string;
        public key: string;
        public value: string;
        public state: string;
        public created: (number|Long);
        public modified: (number|Long);
    }

    interface IUser {
        ctx?: (common.IContext|null);
        id?: (string|null);
        account_id?: (string|null);
        fullname?: (string|null);
        phones?: (string[]|null);
        emails?: (string[]|null);
        traces?: (user.ITrace[]|null);
        devices?: (user.IDevice[]|null);
        is_ban?: (boolean|null);
        avatar_url?: (string|null);
        attributes?: (user.IAttributeData[]|null);
        segments?: (string[]|null);
        labels?: (string[]|null);
        unsubscribed?: (boolean|null);
        marked_spam?: (boolean|null);
        hard_bounced?: (boolean|null);
        total_sessions?: (number|null);
        subiz_id?: (string|null);
        timezone?: (string|null);
        country_code?: (string|null);
        city?: (string|null);
        language?: (string|null);
        aliases?: (string[]|null);
        seen?: (number|Long|null);
        fields?: (user.IField[]|null);
        par?: (number|null);
        modified?: (number|Long|null);
        modified_week?: (number|null);
        activated?: (number|Long|null);
    }

    class User implements IUser {
        constructor(p?: user.IUser);
        public ctx?: (common.IContext|null);
        public id: string;
        public account_id: string;
        public fullname: string;
        public phones: string[];
        public emails: string[];
        public traces: user.ITrace[];
        public devices: user.IDevice[];
        public is_ban: boolean;
        public avatar_url: string;
        public attributes: user.IAttributeData[];
        public segments: string[];
        public labels: string[];
        public unsubscribed: boolean;
        public marked_spam: boolean;
        public hard_bounced: boolean;
        public total_sessions: number;
        public subiz_id: string;
        public timezone: string;
        public country_code: string;
        public city: string;
        public language: string;
        public aliases: string[];
        public seen: (number|Long);
        public fields: user.IField[];
        public par: number;
        public modified: (number|Long);
        public modified_week: number;
        public activated: (number|Long);
    }

    interface IUsers {
        users?: (user.IUser[]|null);
    }

    class Users implements IUsers {
        constructor(p?: user.IUsers);
        public users: user.IUser[];
    }

    interface IDevice {
        id?: (number|null);
        useragent_id?: (number|null);
        useragent?: (string|null);
        screen_resolution?: (string|null);
        language_id?: (number|null);
        language?: (string|null);
    }

    class Device implements IDevice {
        constructor(p?: user.IDevice);
        public id: number;
        public useragent_id: number;
        public useragent: string;
        public screen_resolution: string;
        public language_id: number;
        public language: string;
    }

    interface ITrace {
        id?: (string|null);
        ip?: (string|null);
        location_id?: (number|null);
        city_name?: (string|null);
        country_name?: (string|null);
        country_code?: (string|null);
        continent_code?: (string|null);
        continent_name?: (string|null);
        latitude?: (number|null);
        longitude?: (number|null);
        postal_code?: (string|null);
        timezone?: (string|null);
        isp?: (string|null);
    }

    class Trace implements ITrace {
        constructor(p?: user.ITrace);
        public id: string;
        public ip: string;
        public location_id: number;
        public city_name: string;
        public country_name: string;
        public country_code: string;
        public continent_code: string;
        public continent_name: string;
        public latitude: number;
        public longitude: number;
        public postal_code: string;
        public timezone: string;
        public isp: string;
    }

    interface IMergeRequest {
        ctx?: (common.IContext|null);
        id?: (string|null);
        recent_id?: (string|null);
    }

    class MergeRequest implements IMergeRequest {
        constructor(p?: user.IMergeRequest);
        public ctx?: (common.IContext|null);
        public id: string;
        public recent_id: string;
    }

    interface ICreateRequest {
        challenge_id?: (string|null);
        answer?: (string|null);
    }

    class CreateRequest implements ICreateRequest {
        constructor(p?: user.ICreateRequest);
        public challenge_id: string;
        public answer: string;
    }

    interface ITopic {
        ctx?: (common.IContext|null);
        account_id?: (string|null);
        topic?: (string|null);
        type?: (string|null);
        updated?: (number|Long|null);
        unread?: (number|null);
    }

    class Topic implements ITopic {
        constructor(p?: user.ITopic);
        public ctx?: (common.IContext|null);
        public account_id: string;
        public topic: string;
        public type: string;
        public updated: (number|Long);
        public unread: number;
    }

    interface IUnreadTopic {
        ctx?: (common.IContext|null);
        topic?: (string|null);
        agent_id?: (string|null);
        user_id?: (string|null);
        type?: (string|null);
        updated?: (number|Long|null);
        unread?: (number|null);
    }

    class UnreadTopic implements IUnreadTopic {
        constructor(p?: user.IUnreadTopic);
        public ctx?: (common.IContext|null);
        public topic: string;
        public agent_id: string;
        public user_id: string;
        public type: string;
        public updated: (number|Long);
        public unread: number;
    }

    interface IReadTopicRequest {
        ctx?: (common.IContext|null);
        topic?: (string|null);
        user_id?: (string|null);
        agent_id?: (string|null);
    }

    class ReadTopicRequest implements IReadTopicRequest {
        constructor(p?: user.IReadTopicRequest);
        public ctx?: (common.IContext|null);
        public topic: string;
        public user_id: string;
        public agent_id: string;
    }

    interface ISubscribeRequest {
        ctx?: (common.IContext|null);
        agent_id?: (string|null);
        topics?: (string[]|null);
    }

    class SubscribeRequest implements ISubscribeRequest {
        constructor(p?: user.ISubscribeRequest);
        public ctx?: (common.IContext|null);
        public agent_id: string;
        public topics: string[];
    }

    enum Event {
        UserReadRequested = 0,
        UserUpdateRequested = 2,
        UserCreateRequested = 3,
        UserSearchRequested = 4,
        UserEventCreateRequested = 5,
        UserEventSearchRequested = 7,
        UserTopicSearchRequested = 6,
        UserSegmentationCreateRequested = 10,
        UserSegmentationUpdateRequested = 11,
        UserSegmentationDeleteRequested = 12,
        UserSegmentationListRequested = 13,
        UserSegmentationReadRequested = 14,
        UserAddToMyListRequested = 20,
        UserEventTopicSubscribeRequested = 35,
        UserEventTopicUnsubscribeRequested = 36,
        UserReadTopicRequested = 41,
        UserSubizId = 42,
        UserPresenceReadRequested = 44,
        UserPreviewingReadRequested = 46,
        UserListTopRequested = 47,
        UserAutomationUpsertRequested = 50,
        UserAutomationDeleteRequested = 51,
        UserAutomationListRequested = 52,
        UserAutomationReadRequested = 53,
        AutomationAgentNotificationFired = 54,
        AutomationConversationMessageFired = 55,
        UserSessionUpdateRequested = 65,
        UserSessionCreateRequested = 66,
        UserSessionReadRequested = 67,
        SegmentationLoop = 68,
        AutomationSynced = 102,
        AutomationFired = 103,
        UserRequested = 100,
        UserSynced = 101,
        UserUpserted = 105,
        UserV3Synced = 106
    }

    interface ISubizIDRequest {
        ctx?: (common.IContext|null);
        subiz_id?: (string|null);
        account_id?: (string|null);
    }

    class SubizIDRequest implements ISubizIDRequest {
        constructor(p?: user.ISubizIDRequest);
        public ctx?: (common.IContext|null);
        public subiz_id: string;
        public account_id: string;
    }

    interface IMaskResponse {
        subiz_id?: (string|null);
        account_id?: (string|null);
        user_id?: (string|null);
        mask?: (string|null);
    }

    class MaskResponse implements IMaskResponse {
        constructor(p?: user.IMaskResponse);
        public subiz_id: string;
        public account_id: string;
        public user_id: string;
        public mask: string;
    }

    interface ISubizIDResponse {
        ctx?: (common.IContext|null);
        subiz_id?: (string|null);
        account_id?: (string|null);
        user_id?: (string|null);
    }

    class SubizIDResponse implements ISubizIDResponse {
        constructor(p?: user.ISubizIDResponse);
        public ctx?: (common.IContext|null);
        public subiz_id: string;
        public account_id: string;
        public user_id: string;
    }

    interface ISegmentations {
        ctx?: (common.IContext|null);
        segmentations?: (user.ISegmentation[]|null);
    }

    class Segmentations implements ISegmentations {
        constructor(p?: user.ISegmentations);
        public ctx?: (common.IContext|null);
        public segmentations: user.ISegmentation[];
    }

    interface ISegmentLoopState {
        ctx?: (common.IContext|null);
        account_id?: (string|null);
        user_par?: (number|null);
        loop_created?: (number|Long|null);
        loop_number?: (number|Long|null);
    }

    class SegmentLoopState implements ISegmentLoopState {
        constructor(p?: user.ISegmentLoopState);
        public ctx?: (common.IContext|null);
        public account_id: string;
        public user_par: number;
        public loop_created: (number|Long);
        public loop_number: (number|Long);
    }

    interface IUserSegmentCache {
        ctx?: (common.IContext|null);
        account_id?: (string|null);
        id?: (string|null);
        condition_id?: (string|null);
    }

    class UserSegmentCache implements IUserSegmentCache {
        constructor(p?: user.IUserSegmentCache);
        public ctx?: (common.IContext|null);
        public account_id: string;
        public id: string;
        public condition_id: string;
    }

    interface ISegmentation {
        ctx?: (common.IContext|null);
        account_id?: (string|null);
        id?: (string|null);
        name?: (string|null);
        description?: (string|null);
        user_count?: (number|Long|null);
        ran?: (number|Long|null);
        started_from?: (number|Long|null);
        created?: (number|Long|null);
        modified?: (number|Long|null);
        state?: (string|null);
        condition?: (user.ISegmentCondition|null);
        current_cursor?: (string|null);
    }

    class Segmentation implements ISegmentation {
        constructor(p?: user.ISegmentation);
        public ctx?: (common.IContext|null);
        public account_id: string;
        public id: string;
        public name: string;
        public description: string;
        public user_count: (number|Long);
        public ran: (number|Long);
        public started_from: (number|Long);
        public created: (number|Long);
        public modified: (number|Long);
        public state: string;
        public condition?: (user.ISegmentCondition|null);
        public current_cursor: string;
    }

    namespace Segmentation {

        enum State {
            active = 0,
            inactive = 1
        }
    }

    interface ISegmentTracking {
        ctx?: (common.IContext|null);
        user_par?: (number|null);
        account_id?: (string|null);
        loop_created?: (number|Long|null);
        loop_number?: (number|Long|null);
    }

    class SegmentTracking implements ISegmentTracking {
        constructor(p?: user.ISegmentTracking);
        public ctx?: (common.IContext|null);
        public user_par: number;
        public account_id: string;
        public loop_created: (number|Long);
        public loop_number: (number|Long);
    }

    interface ISegmentCondition {
        join?: (string|null);
        id?: (string|null);
        event_type?: (string|null);
        composition?: (string|null);
        key?: (string|null);
        conditions?: (user.ICondition[]|null);
        left_condition?: (user.ISegmentCondition|null);
        right_condition?: (user.ISegmentCondition|null);
    }

    class SegmentCondition implements ISegmentCondition {
        constructor(p?: user.ISegmentCondition);
        public join: string;
        public id: string;
        public event_type: string;
        public composition: string;
        public key: string;
        public conditions: user.ICondition[];
        public left_condition?: (user.ISegmentCondition|null);
        public right_condition?: (user.ISegmentCondition|null);
    }

    namespace SegmentCondition {

        enum JoinOperator {
            none = 0,
            and = 1,
            or = 2
        }

        enum Composition {
            true = 0,
            sum = 1,
            avg = 2,
            count = 3
        }
    }

    interface ICondition {
        join?: (string|null);
        key?: (string|null);
        operator?: (string|null);
        value?: (string|null);
    }

    class Condition implements ICondition {
        constructor(p?: user.ICondition);
        public join: string;
        public key: string;
        public operator: string;
        public value: string;
    }

    namespace Condition {

        enum JoinOperator {
            none = 0,
            and = 1,
            or = 2
        }
    }

    interface IUserSearchResult {
        ctx?: (common.IContext|null);
        total?: (number|Long|null);
        users?: (user.IUser[]|null);
        anchor?: (string|null);
        unreads?: (number[]|null);
    }

    class UserSearchResult implements IUserSearchResult {
        constructor(p?: user.IUserSearchResult);
        public ctx?: (common.IContext|null);
        public total: (number|Long);
        public users: user.IUser[];
        public anchor: string;
        public unreads: number[];
    }

    interface IUserSearchRequest {
        ctx?: (common.IContext|null);
        segmentation_id?: (string|null);
        query?: (string|null);
        anchor?: (string|null);
        limit?: (number|null);
        agent_id?: (string|null);
        unread?: (boolean|null);
    }

    class UserSearchRequest implements IUserSearchRequest {
        constructor(p?: user.IUserSearchRequest);
        public ctx?: (common.IContext|null);
        public segmentation_id: string;
        public query: string;
        public anchor: string;
        public limit: number;
        public agent_id: string;
        public unread: boolean;
    }

    interface IIndexEvent {
        id?: (string|null);
        account_id?: (string|null);
        user_id?: (string|null);
        created?: (number|Long|null);
        category?: (string|null);
        topics?: (string[]|null);
        object?: (string|null);
        text?: (string|null);
    }

    class IndexEvent implements IIndexEvent {
        constructor(p?: user.IIndexEvent);
        public id: string;
        public account_id: string;
        public user_id: string;
        public created: (number|Long);
        public category: string;
        public topics: string[];
        public object: string;
        public text: string;
    }

    interface IListTopicsRequest {
        ctx?: (common.IContext|null);
        user_id?: (string|null);
        agent_id?: (string|null);
        anchor?: (string|null);
        limit?: (number|null);
        unread?: (boolean|null);
    }

    class ListTopicsRequest implements IListTopicsRequest {
        constructor(p?: user.IListTopicsRequest);
        public ctx?: (common.IContext|null);
        public user_id: string;
        public agent_id: string;
        public anchor: string;
        public limit: number;
        public unread: boolean;
    }

    interface IListTopicsResult {
        ctx?: (common.IContext|null);
        topics?: (user.ITopic[]|null);
        anchor?: (string|null);
    }

    class ListTopicsResult implements IListTopicsResult {
        constructor(p?: user.IListTopicsResult);
        public ctx?: (common.IContext|null);
        public topics: user.ITopic[];
        public anchor: string;
    }

    interface IListNewsRequest {
        ctx?: (common.IContext|null);
        user_id?: (string|null);
        start_time?: (number|Long|null);
        limit?: (string|null);
    }

    class ListNewsRequest implements IListNewsRequest {
        constructor(p?: user.IListNewsRequest);
        public ctx?: (common.IContext|null);
        public user_id: string;
        public start_time: (number|Long);
        public limit: string;
    }

    interface IAddToMyList {
        ctx?: (common.IContext|null);
        agent_id?: (string|null);
        user_id?: (string|null);
    }

    class AddToMyList implements IAddToMyList {
        constructor(p?: user.IAddToMyList);
        public ctx?: (common.IContext|null);
        public agent_id: string;
        public user_id: string;
    }

    interface IField {
        name?: (string|null);
        account_id?: (string|null);
        user_id?: (string|null);
        setter?: (string|null);
        setter_type?: (string|null);
        updated?: (number|Long|null);
        data?: (string|null);
        id?: (string|null);
    }

    class Field implements IField {
        constructor(p?: user.IField);
        public name: string;
        public account_id: string;
        public user_id: string;
        public setter: string;
        public setter_type: string;
        public updated: (number|Long);
        public data: string;
        public id: string;
    }

    interface IPresence {
        ctx?: (common.IContext|null);
        account_id?: (string|null);
        user_id?: (string|null);
        pinged?: (number|Long|null);
        pinged_minute?: (number|Long|null);
    }

    class Presence implements IPresence {
        constructor(p?: user.IPresence);
        public ctx?: (common.IContext|null);
        public account_id: string;
        public user_id: string;
        public pinged: (number|Long);
        public pinged_minute: (number|Long);
    }

    interface IPresences {
        ctx?: (common.IContext|null);
        presences?: (user.IPresence[]|null);
    }

    class Presences implements IPresences {
        constructor(p?: user.IPresences);
        public ctx?: (common.IContext|null);
        public presences: user.IPresence[];
    }

    interface IVisitor {
        ctx?: (common.IContext|null);
        account_id?: (string|null);
        user?: (user.IUser|null);
        pinged?: (number|Long|null);
        page_url?: (string|null);
        page_viewed?: (number|Long|null);
        page_title?: (string|null);
    }

    class Visitor implements IVisitor {
        constructor(p?: user.IVisitor);
        public ctx?: (common.IContext|null);
        public account_id: string;
        public user?: (user.IUser|null);
        public pinged: (number|Long);
        public page_url: string;
        public page_viewed: (number|Long);
        public page_title: string;
    }

    interface IVisitors {
        ctx?: (common.IContext|null);
        visitors?: (user.IVisitor[]|null);
    }

    class Visitors implements IVisitors {
        constructor(p?: user.IVisitors);
        public ctx?: (common.IContext|null);
        public visitors: user.IVisitor[];
    }

    interface ILastView {
        account_id?: (string|null);
        user_id?: (string|null);
        url?: (string|null);
        ua?: (string|null);
        ip?: (string|null);
        created?: (number|Long|null);
        event_id?: (string|null);
        title?: (string|null);
    }

    class LastView implements ILastView {
        constructor(p?: user.ILastView);
        public account_id: string;
        public user_id: string;
        public url: string;
        public ua: string;
        public ip: string;
        public created: (number|Long);
        public event_id: string;
        public title: string;
    }

    interface IAutomation {
        ctx?: (common.IContext|null);
        account_id?: (string|null);
        id?: (string|null);
        channel?: (string|null);
        name?: (string|null);
        description?: (string|null);
        conditions?: (user.ICondition[]|null);
        created?: (number|Long|null);
        modified?: (number|Long|null);
        state?: (string|null);
        action_type?: (string|null);
        action_data?: (string|null);
        scope?: (string|null);
    }

    class Automation implements IAutomation {
        constructor(p?: user.IAutomation);
        public ctx?: (common.IContext|null);
        public account_id: string;
        public id: string;
        public channel: string;
        public name: string;
        public description: string;
        public conditions: user.ICondition[];
        public created: (number|Long);
        public modified: (number|Long);
        public state: string;
        public action_type: string;
        public action_data: string;
        public scope: string;
    }

    namespace Automation {

        enum State {
            active = 0,
            inactive = 1
        }

        enum ActionType {
            conversation_message = 0,
            agent_notification = 1,
            automation_invite_message = 4
        }

        enum Scope {
            conversation = 2,
            user = 3
        }
    }

    interface IAutomations {
        ctx?: (common.IContext|null);
        automations?: (user.IAutomation[]|null);
    }

    class Automations implements IAutomations {
        constructor(p?: user.IAutomations);
        public ctx?: (common.IContext|null);
        public automations: user.IAutomation[];
    }

    interface ISession {
        ctx?: (common.IContext|null);
        account_id?: (string|null);
        user_id?: (string|null);
        id?: (string|null);
        platform?: (string|null);
        referrer?: (string|null);
        search_engine?: (string|null);
        started?: (number|Long|null);
        tracked?: (number|Long|null);
        status?: (string|null);
        events_count?: (number|null);
        content_views_count?: (number|null);
        search_term?: (string|null);
    }

    class Session implements ISession {
        constructor(p?: user.ISession);
        public ctx?: (common.IContext|null);
        public account_id: string;
        public user_id: string;
        public id: string;
        public platform: string;
        public referrer: string;
        public search_engine: string;
        public started: (number|Long);
        public tracked: (number|Long);
        public status: string;
        public events_count: number;
        public content_views_count: number;
        public search_term: string;
    }

    namespace Session {

        enum Platform {
            web = 0,
            mobile = 2,
            desktop = 4
        }

        enum Status {
            open = 0,
            closed = 1
        }
    }

    interface IDeleteAttrRequest {
        ctx?: (common.IContext|null);
        key?: (string|null);
    }

    class DeleteAttrRequest implements IDeleteAttrRequest {
        constructor(p?: user.IDeleteAttrRequest);
        public ctx?: (common.IContext|null);
        public key: string;
    }

    class SegmentationMgr extends $protobuf.rpc.Service {
        constructor(rpcImpl: $protobuf.RPCImpl, requestDelimited?: boolean, responseDelimited?: boolean);
        public createSegment(request: user.ISegmentation, callback: user.SegmentationMgr.CreateSegmentCallback): void;
        public createSegment(request: user.ISegmentation): Promise<user.Segmentation>;
        public updateSegment(request: user.ISegmentation, callback: user.SegmentationMgr.UpdateSegmentCallback): void;
        public updateSegment(request: user.ISegmentation): Promise<user.Segmentation>;
        public listSegments(request: common.IId, callback: user.SegmentationMgr.ListSegmentsCallback): void;
        public listSegments(request: common.IId): Promise<user.Segmentations>;
        public deleteSegment(request: common.IId, callback: user.SegmentationMgr.DeleteSegmentCallback): void;
        public deleteSegment(request: common.IId): Promise<common.Empty>;
        public readSegment(request: common.IId, callback: user.SegmentationMgr.ReadSegmentCallback): void;
        public readSegment(request: common.IId): Promise<user.Segmentation>;
    }

    namespace SegmentationMgr {

        type CreateSegmentCallback = (error: (Error|null), response?: user.Segmentation) => void;

        type UpdateSegmentCallback = (error: (Error|null), response?: user.Segmentation) => void;

        type ListSegmentsCallback = (error: (Error|null), response?: user.Segmentations) => void;

        type DeleteSegmentCallback = (error: (Error|null), response?: common.Empty) => void;

        type ReadSegmentCallback = (error: (Error|null), response?: user.Segmentation) => void;
    }

    class VisitorMgr extends $protobuf.rpc.Service {
        constructor(rpcImpl: $protobuf.RPCImpl, requestDelimited?: boolean, responseDelimited?: boolean);
        public readPresence(request: common.IId, callback: user.VisitorMgr.ReadPresenceCallback): void;
        public readPresence(request: common.IId): Promise<user.Presence>;
        public readPresences(request: common.IIds, callback: user.VisitorMgr.ReadPresencesCallback): void;
        public readPresences(request: common.IIds): Promise<user.Presences>;
        public readPreview(request: common.IId, callback: user.VisitorMgr.ReadPreviewCallback): void;
        public readPreview(request: common.IId): Promise<user.LastView>;
        public listTopVisitors(request: common.IId, callback: user.VisitorMgr.ListTopVisitorsCallback): void;
        public listTopVisitors(request: common.IId): Promise<user.Visitors>;
    }

    namespace VisitorMgr {

        type ReadPresenceCallback = (error: (Error|null), response?: user.Presence) => void;

        type ReadPresencesCallback = (error: (Error|null), response?: user.Presences) => void;

        type ReadPreviewCallback = (error: (Error|null), response?: user.LastView) => void;

        type ListTopVisitorsCallback = (error: (Error|null), response?: user.Visitors) => void;
    }

    class UserMgr extends $protobuf.rpc.Service {
        constructor(rpcImpl: $protobuf.RPCImpl, requestDelimited?: boolean, responseDelimited?: boolean);
        public searchUsers(request: user.IUserSearchRequest, callback: user.UserMgr.SearchUsersCallback): void;
        public searchUsers(request: user.IUserSearchRequest): Promise<user.UserSearchResult>;
        public subizID(request: user.ISubizIDRequest, callback: user.UserMgr.SubizIDCallback): void;
        public subizID(request: user.ISubizIDRequest): Promise<user.SubizIDResponse>;
        public addToMy(request: user.IAddToMyRequest, callback: user.UserMgr.AddToMyCallback): void;
        public addToMy(request: user.IAddToMyRequest): Promise<common.Empty>;
        public createUser(request: user.IUser, callback: user.UserMgr.CreateUserCallback): void;
        public createUser(request: user.IUser): Promise<user.User>;
        public updateUser(request: user.IUser, callback: user.UserMgr.UpdateUserCallback): void;
        public updateUser(request: user.IUser): Promise<user.User>;
        public readUser(request: common.IId, callback: user.UserMgr.ReadUserCallback): void;
        public readUser(request: common.IId): Promise<user.User>;
        public removeFromMy(request: common.IId, callback: user.UserMgr.RemoveFromMyCallback): void;
        public removeFromMy(request: common.IId): Promise<common.Empty>;
    }

    namespace UserMgr {

        type SearchUsersCallback = (error: (Error|null), response?: user.UserSearchResult) => void;

        type SubizIDCallback = (error: (Error|null), response?: user.SubizIDResponse) => void;

        type AddToMyCallback = (error: (Error|null), response?: common.Empty) => void;

        type CreateUserCallback = (error: (Error|null), response?: user.User) => void;

        type UpdateUserCallback = (error: (Error|null), response?: user.User) => void;

        type ReadUserCallback = (error: (Error|null), response?: user.User) => void;

        type RemoveFromMyCallback = (error: (Error|null), response?: common.Empty) => void;
    }

    class AutomationMgr extends $protobuf.rpc.Service {
        constructor(rpcImpl: $protobuf.RPCImpl, requestDelimited?: boolean, responseDelimited?: boolean);
        public listAutomations(request: common.IId, callback: user.AutomationMgr.ListAutomationsCallback): void;
        public listAutomations(request: common.IId): Promise<user.Automations>;
        public updateAutomation(request: user.IAutomation, callback: user.AutomationMgr.UpdateAutomationCallback): void;
        public updateAutomation(request: user.IAutomation): Promise<user.Automation>;
        public deleteAutomation(request: common.IId, callback: user.AutomationMgr.DeleteAutomationCallback): void;
        public deleteAutomation(request: common.IId): Promise<common.Empty>;
        public readAutomation(request: common.IId, callback: user.AutomationMgr.ReadAutomationCallback): void;
        public readAutomation(request: common.IId): Promise<user.Automation>;
        public createAutomation(request: user.IAutomation, callback: user.AutomationMgr.CreateAutomationCallback): void;
        public createAutomation(request: user.IAutomation): Promise<user.Automation>;
    }

    namespace AutomationMgr {

        type ListAutomationsCallback = (error: (Error|null), response?: user.Automations) => void;

        type UpdateAutomationCallback = (error: (Error|null), response?: user.Automation) => void;

        type DeleteAutomationCallback = (error: (Error|null), response?: common.Empty) => void;

        type ReadAutomationCallback = (error: (Error|null), response?: user.Automation) => void;

        type CreateAutomationCallback = (error: (Error|null), response?: user.Automation) => void;
    }

    class SessionMgr extends $protobuf.rpc.Service {
        constructor(rpcImpl: $protobuf.RPCImpl, requestDelimited?: boolean, responseDelimited?: boolean);
        public createSession(request: user.ISession, callback: user.SessionMgr.CreateSessionCallback): void;
        public createSession(request: user.ISession): Promise<user.Session>;
        public readSession(request: user.ISession, callback: user.SessionMgr.ReadSessionCallback): void;
        public readSession(request: user.ISession): Promise<user.Session>;
        public updateSession(request: user.ISession, callback: user.SessionMgr.UpdateSessionCallback): void;
        public updateSession(request: user.ISession): Promise<user.Session>;
    }

    namespace SessionMgr {

        type CreateSessionCallback = (error: (Error|null), response?: user.Session) => void;

        type ReadSessionCallback = (error: (Error|null), response?: user.Session) => void;

        type UpdateSessionCallback = (error: (Error|null), response?: user.Session) => void;
    }

    class AttributeMgr extends $protobuf.rpc.Service {
        constructor(rpcImpl: $protobuf.RPCImpl, requestDelimited?: boolean, responseDelimited?: boolean);
        public listAttributeDefinitions(request: common.IEmpty, callback: user.AttributeMgr.ListAttributeDefinitionsCallback): void;
        public listAttributeDefinitions(request: common.IEmpty): Promise<user.AttributeDefinitions>;
        public createAttributeDefinition(request: user.IAttributeDefinition, callback: user.AttributeMgr.CreateAttributeDefinitionCallback): void;
        public createAttributeDefinition(request: user.IAttributeDefinition): Promise<user.AttributeDefinition>;
        public updateAttributeDefinition(request: user.IAttributeDefinition, callback: user.AttributeMgr.UpdateAttributeDefinitionCallback): void;
        public updateAttributeDefinition(request: user.IAttributeDefinition): Promise<user.AttributeDefinition>;
        public deleteAttributeDefinition(request: user.IDeleteAttrRequest, callback: user.AttributeMgr.DeleteAttributeDefinitionCallback): void;
        public deleteAttributeDefinition(request: user.IDeleteAttrRequest): Promise<common.Empty>;
    }

    namespace AttributeMgr {

        type ListAttributeDefinitionsCallback = (error: (Error|null), response?: user.AttributeDefinitions) => void;

        type CreateAttributeDefinitionCallback = (error: (Error|null), response?: user.AttributeDefinition) => void;

        type UpdateAttributeDefinitionCallback = (error: (Error|null), response?: user.AttributeDefinition) => void;

        type DeleteAttributeDefinitionCallback = (error: (Error|null), response?: common.Empty) => void;
    }

    interface IAutomationCheck {
        account_id?: (string|null);
        automation_id?: (string|null);
        user_id?: (string|null);
        event_id?: (string|null);
        scope?: (string|null);
    }

    class AutomationCheck implements IAutomationCheck {
        constructor(p?: user.IAutomationCheck);
        public account_id: string;
        public automation_id: string;
        public user_id: string;
        public event_id: string;
        public scope: string;
    }
}

export namespace webhook4 {

    interface IWebhook {
        ctx?: (common.IContext|null);
        account_id?: (string|null);
        client_id?: (string|null);
        url?: (string|null);
        secret?: (string|null);
        events?: (string[]|null);
        state?: (string|null);
        last_hook_at?: (number|Long|null);
        last_hook_id?: (number|Long|null);
        events_count?: (number|Long|null);
        last_hook_response?: (string|null);
        last_hook_status?: (number|null);
        last_hook_payload?: (string|null);
        backoffs_count?: (number|Long|null);
    }

    class Webhook implements IWebhook {
        constructor(p?: webhook4.IWebhook);
        public ctx?: (common.IContext|null);
        public account_id: string;
        public client_id: string;
        public url: string;
        public secret: string;
        public events: string[];
        public state: string;
        public last_hook_at: (number|Long);
        public last_hook_id: (number|Long);
        public events_count: (number|Long);
        public last_hook_response: string;
        public last_hook_status: number;
        public last_hook_payload: string;
        public backoffs_count: (number|Long);
    }

    namespace Webhook {

        enum State {
            normal = 0,
            backoff = 1,
            inactive = 2
        }
    }

    enum Event {
        Webhook4WebhookReadRequested = 3,
        Webhook4WebhookUpserted = 4,
        Webhook4WebhookVerifyRequested = 5,
        Webhook4Requested = 6,
        Webhook4ClearBackoffRequested = 7,
        Webhook4PurgeQueueRequested = 8,
        Webhook4ShardSend = 9
    }

    class WebhookService extends $protobuf.rpc.Service {
        constructor(rpcImpl: $protobuf.RPCImpl, requestDelimited?: boolean, responseDelimited?: boolean);
        public read(request: common.IId, callback: webhook4.WebhookService.ReadCallback): void;
        public read(request: common.IId): Promise<webhook4.Webhook>;
        public clearBackoff(request: common.IId, callback: webhook4.WebhookService.ClearBackoffCallback): void;
        public clearBackoff(request: common.IId): Promise<common.Empty>;
        public purgeQueue(request: common.IId, callback: webhook4.WebhookService.PurgeQueueCallback): void;
        public purgeQueue(request: common.IId): Promise<common.Empty>;
    }

    namespace WebhookService {

        type ReadCallback = (error: (Error|null), response?: webhook4.Webhook) => void;

        type ClearBackoffCallback = (error: (Error|null), response?: common.Empty) => void;

        type PurgeQueueCallback = (error: (Error|null), response?: common.Empty) => void;
    }
}

export namespace widget {

    interface IAllType {
        theme?: (widget.ITheme|null);
        sound?: (widget.ISound|null);
        setting?: (widget.ISetting|null);
        us?: (widget.IUserSetting|null);
    }

    class AllType implements IAllType {
        constructor(p?: widget.IAllType);
        public theme?: (widget.ITheme|null);
        public sound?: (widget.ISound|null);
        public setting?: (widget.ISetting|null);
        public us?: (widget.IUserSetting|null);
    }

    class MyServer extends $protobuf.rpc.Service {
        constructor(rpcImpl: $protobuf.RPCImpl, requestDelimited?: boolean, responseDelimited?: boolean);
        public do(request: widget.IAllType, callback: widget.MyServer.DoCallback): void;
        public do(request: widget.IAllType): Promise<widget.AllType>;
    }

    namespace MyServer {

        type DoCallback = (error: (Error|null), response?: widget.AllType) => void;
    }

    class WidgetService extends $protobuf.rpc.Service {
        constructor(rpcImpl: $protobuf.RPCImpl, requestDelimited?: boolean, responseDelimited?: boolean);
        public read(request: common.IId, callback: widget.WidgetService.ReadCallback): void;
        public read(request: common.IId): Promise<widget.Setting>;
        public update(request: widget.ISetting, callback: widget.WidgetService.UpdateCallback): void;
        public update(request: widget.ISetting): Promise<widget.Setting>;
        public readUserSetting(request: common.IId, callback: widget.WidgetService.ReadUserSettingCallback): void;
        public readUserSetting(request: common.IId): Promise<widget.UserSetting>;
        public updateUserSetting(request: widget.IUserSetting, callback: widget.WidgetService.UpdateUserSettingCallback): void;
        public updateUserSetting(request: widget.IUserSetting): Promise<widget.UserSetting>;
    }

    namespace WidgetService {

        type ReadCallback = (error: (Error|null), response?: widget.Setting) => void;

        type UpdateCallback = (error: (Error|null), response?: widget.Setting) => void;

        type ReadUserSettingCallback = (error: (Error|null), response?: widget.UserSetting) => void;

        type UpdateUserSettingCallback = (error: (Error|null), response?: widget.UserSetting) => void;
    }

    interface ITheme {
        account_id?: (string|null);
        custom_css?: (string|null);
        widget_position?: (string|null);
        window_mode?: (string|null);
    }

    class Theme implements ITheme {
        constructor(p?: widget.ITheme);
        public account_id: string;
        public custom_css: string;
        public widget_position: string;
        public window_mode: string;
    }

    namespace Theme {

        enum ButtonPosition {
            left = 0,
            right = 1
        }

        enum WindowMode {
            mini = 0,
            full = 1
        }
    }

    interface ISound {
        account_id?: (string|null);
        enabled?: (boolean|null);
        new_conversation?: (string|null);
        file_create?: (string|null);
        new_message?: (string|null);
        message_send_failed?: (string|null);
    }

    class Sound implements ISound {
        constructor(p?: widget.ISound);
        public account_id: string;
        public enabled: boolean;
        public new_conversation: string;
        public file_create: string;
        public new_message: string;
        public message_send_failed: string;
    }

    interface ISetting {
        ctx?: (common.IContext|null);
        account_id?: (string|null);
        widget_version?: (string|null);
        sound?: (widget.ISound|null);
        language?: (string|null);
        theme?: (widget.ITheme|null);
        replytime?: (number|null);
        agents?: (account.IAgent[]|null);
        agent_ids?: (string[]|null);
        language_url?: (string|null);
        custom_language_url?: (string|null);
        css_url?: (string|null);
        custom_css_url?: (string|null);
        custom_language?: (string|null);
    }

    class Setting implements ISetting {
        constructor(p?: widget.ISetting);
        public ctx?: (common.IContext|null);
        public account_id: string;
        public widget_version: string;
        public sound?: (widget.ISound|null);
        public language: string;
        public theme?: (widget.ITheme|null);
        public replytime: number;
        public agents: account.IAgent[];
        public agent_ids: string[];
        public language_url: string;
        public custom_language_url: string;
        public css_url: string;
        public custom_css_url: string;
        public custom_language: string;
    }

    interface IUserSetting {
        ctx?: (common.IContext|null);
        account?: (account.IAccount|null);
        account_id?: (string|null);
        user?: (user.IUser|null);
        user_id?: (string|null);
        sound_enabled?: (boolean|null);
        language?: (string|null);
        send_transcript?: (boolean|null);
        account_setting?: (widget.ISetting|null);
        desktop_notification?: (boolean|null);
    }

    class UserSetting implements IUserSetting {
        constructor(p?: widget.IUserSetting);
        public ctx?: (common.IContext|null);
        public account?: (account.IAccount|null);
        public account_id: string;
        public user?: (user.IUser|null);
        public user_id: string;
        public sound_enabled: boolean;
        public language: string;
        public send_transcript: boolean;
        public account_setting?: (widget.ISetting|null);
        public desktop_notification: boolean;
    }

    interface IGlobal {
        ctx?: (common.IContext|null);
        name?: (string|null);
        data?: (string|null);
        pk?: (string|null);
    }

    class Global implements IGlobal {
        constructor(p?: widget.IGlobal);
        public ctx?: (common.IContext|null);
        public name: string;
        public data: string;
        public pk: string;
    }

    enum Event {
        WidgetUserSettingReadRequested = 0,
        WidgetUserSettingUpdateRequested = 1,
        WidgetSettingReadRequested = 2,
        WidgetSettingUpdateRequested = 3,
        WidgetSettingCssVersionUpdated = 4,
        WidgetSettingLanguageUpdated = 5,
        WidgetSettingCssVersionUpdateRequested = 6,
        WidgetSettingLanguageUpdateRequested = 7,
        WidgetRequested = 8,
        WidgetSettingUpserted = 10,
        WidgetSynced = 100,
        WidgetV3Synced = 103
    }
}

export namespace ws {

    interface IAllType {
        sub?: (ws.ISubscribe|null);
    }

    class AllType implements IAllType {
        constructor(p?: ws.IAllType);
        public sub?: (ws.ISubscribe|null);
    }

    class MyServer extends $protobuf.rpc.Service {
        constructor(rpcImpl: $protobuf.RPCImpl, requestDelimited?: boolean, responseDelimited?: boolean);
        public do(request: ws.IAllType, callback: ws.MyServer.DoCallback): void;
        public do(request: ws.IAllType): Promise<ws.AllType>;
    }

    namespace MyServer {

        type DoCallback = (error: (Error|null), response?: ws.AllType) => void;
    }

    interface ISubscribe {
        events?: (string[]|null);
        connection_id?: (string|null);
    }

    class Subscribe implements ISubscribe {
        constructor(p?: ws.ISubscribe);
        public events: string[];
        public connection_id: string;
    }

    interface IListRequest {
        ctx?: (common.IContext|null);
        connection_id?: (string|null);
    }

    class ListRequest implements IListRequest {
        constructor(p?: ws.IListRequest);
        public ctx?: (common.IContext|null);
        public connection_id: string;
    }

    enum Event {
        WsSynced = 0,
        WsRequested = 1,
        WsSend = 10
    }

    interface IPayload {
        id?: (number|Long|null);
        message?: (string|null);
        error?: (string|null);
    }

    class Payload implements IPayload {
        constructor(p?: ws.IPayload);
        public id: (number|Long);
        public message: string;
        public error: string;
    }
}
