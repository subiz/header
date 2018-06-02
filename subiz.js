/*eslint-disable block-scoped-var, no-redeclare, no-control-regex, no-prototype-builtins*/
import * as $protobuf from "protobufjs/minimal";

// Common aliases
const $util = $protobuf.util;

// Exported root namespace
const $root = $protobuf.roots["default"] || ($protobuf.roots["default"] = {});

export const account = $root.account = (() => {

    /**
     * Namespace account.
     * @exports account
     * @namespace
     */
    const account = {};

    account.Agent = (function() {

        /**
         * Properties of an Agent.
         * @memberof account
         * @interface IAgent
         * @property {common.IContext|null} [ctx] Agent ctx
         * @property {string|null} [id] Agent id
         * @property {string|null} [account_id] Agent account_id
         * @property {string|null} [fullname] Agent fullname
         * @property {string|null} [email] Agent email
         * @property {Array.<string>|null} [emails] Agent emails
         * @property {Array.<string>|null} [phones] Agent phones
         * @property {string|null} [phone] Agent phone
         * @property {boolean|null} [is_owner] Agent is_owner
         * @property {string|null} [job_title] Agent job_title
         * @property {string|null} [gender] Agent gender
         * @property {string|null} [avatar_url] Agent avatar_url
         * @property {string|null} [lang] Agent lang
         * @property {string|null} [location] Agent location
         * @property {string|null} [timezone] Agent timezone
         * @property {string|null} [encrypted_password] Agent encrypted_password
         * @property {number|Long|null} [joined] Agent joined
         * @property {string|null} [invited_by] Agent invited_by
         * @property {string|null} [state] Agent state
         * @property {number|Long|null} [password_changed] Agent password_changed
         * @property {number|Long|null} [seen] Agent seen
         * @property {number|Long|null} [modified] Agent modified
         * @property {auth.IMethod|null} [method] Agent method
         * @property {account.IAccount|null} [account] Agent account
         * @property {string|null} [country_code] Agent country_code
         * @property {number|null} [v3_state] Agent v3_state
         * @property {string|null} [v3_hashed_password] Agent v3_hashed_password
         */

        /**
         * Constructs a new Agent.
         * @memberof account
         * @classdesc Represents an Agent.
         * @implements IAgent
         * @constructor
         * @param {account.IAgent=} [p] Properties to set
         */
        function Agent(p) {
            this.emails = [];
            this.phones = [];
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * Agent ctx.
         * @member {common.IContext|null|undefined} ctx
         * @memberof account.Agent
         * @instance
         */
        Agent.prototype.ctx = null;

        /**
         * Agent id.
         * @member {string} id
         * @memberof account.Agent
         * @instance
         */
        Agent.prototype.id = "";

        /**
         * Agent account_id.
         * @member {string} account_id
         * @memberof account.Agent
         * @instance
         */
        Agent.prototype.account_id = "";

        /**
         * Agent fullname.
         * @member {string} fullname
         * @memberof account.Agent
         * @instance
         */
        Agent.prototype.fullname = "";

        /**
         * Agent email.
         * @member {string} email
         * @memberof account.Agent
         * @instance
         */
        Agent.prototype.email = "";

        /**
         * Agent emails.
         * @member {Array.<string>} emails
         * @memberof account.Agent
         * @instance
         */
        Agent.prototype.emails = $util.emptyArray;

        /**
         * Agent phones.
         * @member {Array.<string>} phones
         * @memberof account.Agent
         * @instance
         */
        Agent.prototype.phones = $util.emptyArray;

        /**
         * Agent phone.
         * @member {string} phone
         * @memberof account.Agent
         * @instance
         */
        Agent.prototype.phone = "";

        /**
         * Agent is_owner.
         * @member {boolean} is_owner
         * @memberof account.Agent
         * @instance
         */
        Agent.prototype.is_owner = false;

        /**
         * Agent job_title.
         * @member {string} job_title
         * @memberof account.Agent
         * @instance
         */
        Agent.prototype.job_title = "";

        /**
         * Agent gender.
         * @member {string} gender
         * @memberof account.Agent
         * @instance
         */
        Agent.prototype.gender = "";

        /**
         * Agent avatar_url.
         * @member {string} avatar_url
         * @memberof account.Agent
         * @instance
         */
        Agent.prototype.avatar_url = "";

        /**
         * Agent lang.
         * @member {string} lang
         * @memberof account.Agent
         * @instance
         */
        Agent.prototype.lang = "";

        /**
         * Agent location.
         * @member {string} location
         * @memberof account.Agent
         * @instance
         */
        Agent.prototype.location = "";

        /**
         * Agent timezone.
         * @member {string} timezone
         * @memberof account.Agent
         * @instance
         */
        Agent.prototype.timezone = "";

        /**
         * Agent encrypted_password.
         * @member {string} encrypted_password
         * @memberof account.Agent
         * @instance
         */
        Agent.prototype.encrypted_password = "";

        /**
         * Agent joined.
         * @member {number|Long} joined
         * @memberof account.Agent
         * @instance
         */
        Agent.prototype.joined = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

        /**
         * Agent invited_by.
         * @member {string} invited_by
         * @memberof account.Agent
         * @instance
         */
        Agent.prototype.invited_by = "";

        /**
         * Agent state.
         * @member {string} state
         * @memberof account.Agent
         * @instance
         */
        Agent.prototype.state = "";

        /**
         * Agent password_changed.
         * @member {number|Long} password_changed
         * @memberof account.Agent
         * @instance
         */
        Agent.prototype.password_changed = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

        /**
         * Agent seen.
         * @member {number|Long} seen
         * @memberof account.Agent
         * @instance
         */
        Agent.prototype.seen = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

        /**
         * Agent modified.
         * @member {number|Long} modified
         * @memberof account.Agent
         * @instance
         */
        Agent.prototype.modified = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

        /**
         * Agent method.
         * @member {auth.IMethod|null|undefined} method
         * @memberof account.Agent
         * @instance
         */
        Agent.prototype.method = null;

        /**
         * Agent account.
         * @member {account.IAccount|null|undefined} account
         * @memberof account.Agent
         * @instance
         */
        Agent.prototype.account = null;

        /**
         * Agent country_code.
         * @member {string} country_code
         * @memberof account.Agent
         * @instance
         */
        Agent.prototype.country_code = "";

        /**
         * Agent v3_state.
         * @member {number} v3_state
         * @memberof account.Agent
         * @instance
         */
        Agent.prototype.v3_state = 0;

        /**
         * Agent v3_hashed_password.
         * @member {string} v3_hashed_password
         * @memberof account.Agent
         * @instance
         */
        Agent.prototype.v3_hashed_password = "";

        /**
         * Gender enum.
         * @name account.Agent.Gender
         * @enum {string}
         * @property {number} unset=0 unset value
         * @property {number} male=1 male value
         * @property {number} female=2 female value
         * @property {number} bisexual=3 bisexual value
         * @property {number} asexual=4 asexual value
         */
        Agent.Gender = (function() {
            const valuesById = {}, values = Object.create(valuesById);
            values[valuesById[0] = "unset"] = 0;
            values[valuesById[1] = "male"] = 1;
            values[valuesById[2] = "female"] = 2;
            values[valuesById[3] = "bisexual"] = 3;
            values[valuesById[4] = "asexual"] = 4;
            return values;
        })();

        /**
         * AgentState enum.
         * @name account.Agent.AgentState
         * @enum {string}
         * @property {number} pending=0 pending value
         * @property {number} active=1 active value
         * @property {number} inactive=2 inactive value
         * @property {number} deleted=3 deleted value
         */
        Agent.AgentState = (function() {
            const valuesById = {}, values = Object.create(valuesById);
            values[valuesById[0] = "pending"] = 0;
            values[valuesById[1] = "active"] = 1;
            values[valuesById[2] = "inactive"] = 2;
            values[valuesById[3] = "deleted"] = 3;
            return values;
        })();

        return Agent;
    })();

    account.Invitation = (function() {

        /**
         * Properties of an Invitation.
         * @memberof account
         * @interface IInvitation
         * @property {common.IContext|null} [ctx] Invitation ctx
         * @property {string|null} [id] Invitation id
         * @property {string|null} [account_id] Invitation account_id
         * @property {string|null} [from_id] Invitation from_id
         * @property {string|null} [email] Invitation email
         * @property {string|null} [agent_id] Invitation agent_id
         * @property {number|Long|null} [sent] Invitation sent
         * @property {number|Long|null} [replied] Invitation replied
         * @property {string|null} [agent_fullname] Invitation agent_fullname
         * @property {string|null} [agent_job_title] Invitation agent_job_title
         * @property {string|null} [token] Invitation token
         */

        /**
         * Constructs a new Invitation.
         * @memberof account
         * @classdesc Represents an Invitation.
         * @implements IInvitation
         * @constructor
         * @param {account.IInvitation=} [p] Properties to set
         */
        function Invitation(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * Invitation ctx.
         * @member {common.IContext|null|undefined} ctx
         * @memberof account.Invitation
         * @instance
         */
        Invitation.prototype.ctx = null;

        /**
         * Invitation id.
         * @member {string} id
         * @memberof account.Invitation
         * @instance
         */
        Invitation.prototype.id = "";

        /**
         * Invitation account_id.
         * @member {string} account_id
         * @memberof account.Invitation
         * @instance
         */
        Invitation.prototype.account_id = "";

        /**
         * Invitation from_id.
         * @member {string} from_id
         * @memberof account.Invitation
         * @instance
         */
        Invitation.prototype.from_id = "";

        /**
         * Invitation email.
         * @member {string} email
         * @memberof account.Invitation
         * @instance
         */
        Invitation.prototype.email = "";

        /**
         * Invitation agent_id.
         * @member {string} agent_id
         * @memberof account.Invitation
         * @instance
         */
        Invitation.prototype.agent_id = "";

        /**
         * Invitation sent.
         * @member {number|Long} sent
         * @memberof account.Invitation
         * @instance
         */
        Invitation.prototype.sent = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

        /**
         * Invitation replied.
         * @member {number|Long} replied
         * @memberof account.Invitation
         * @instance
         */
        Invitation.prototype.replied = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

        /**
         * Invitation agent_fullname.
         * @member {string} agent_fullname
         * @memberof account.Invitation
         * @instance
         */
        Invitation.prototype.agent_fullname = "";

        /**
         * Invitation agent_job_title.
         * @member {string} agent_job_title
         * @memberof account.Invitation
         * @instance
         */
        Invitation.prototype.agent_job_title = "";

        /**
         * Invitation token.
         * @member {string} token
         * @memberof account.Invitation
         * @instance
         */
        Invitation.prototype.token = "";

        return Invitation;
    })();

    account.AgentGroup = (function() {

        /**
         * Properties of an AgentGroup.
         * @memberof account
         * @interface IAgentGroup
         * @property {common.IContext|null} [ctx] AgentGroup ctx
         * @property {string|null} [id] AgentGroup id
         * @property {string|null} [account_id] AgentGroup account_id
         * @property {string|null} [name] AgentGroup name
         * @property {string|null} [logo_url] AgentGroup logo_url
         * @property {Array.<account.IAgent>|null} [members] AgentGroup members
         * @property {number|Long|null} [created] AgentGroup created
         * @property {number|Long|null} [modified] AgentGroup modified
         * @property {number|null} [members_count] AgentGroup members_count
         */

        /**
         * Constructs a new AgentGroup.
         * @memberof account
         * @classdesc Represents an AgentGroup.
         * @implements IAgentGroup
         * @constructor
         * @param {account.IAgentGroup=} [p] Properties to set
         */
        function AgentGroup(p) {
            this.members = [];
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * AgentGroup ctx.
         * @member {common.IContext|null|undefined} ctx
         * @memberof account.AgentGroup
         * @instance
         */
        AgentGroup.prototype.ctx = null;

        /**
         * AgentGroup id.
         * @member {string} id
         * @memberof account.AgentGroup
         * @instance
         */
        AgentGroup.prototype.id = "";

        /**
         * AgentGroup account_id.
         * @member {string} account_id
         * @memberof account.AgentGroup
         * @instance
         */
        AgentGroup.prototype.account_id = "";

        /**
         * AgentGroup name.
         * @member {string} name
         * @memberof account.AgentGroup
         * @instance
         */
        AgentGroup.prototype.name = "";

        /**
         * AgentGroup logo_url.
         * @member {string} logo_url
         * @memberof account.AgentGroup
         * @instance
         */
        AgentGroup.prototype.logo_url = "";

        /**
         * AgentGroup members.
         * @member {Array.<account.IAgent>} members
         * @memberof account.AgentGroup
         * @instance
         */
        AgentGroup.prototype.members = $util.emptyArray;

        /**
         * AgentGroup created.
         * @member {number|Long} created
         * @memberof account.AgentGroup
         * @instance
         */
        AgentGroup.prototype.created = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

        /**
         * AgentGroup modified.
         * @member {number|Long} modified
         * @memberof account.AgentGroup
         * @instance
         */
        AgentGroup.prototype.modified = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

        /**
         * AgentGroup members_count.
         * @member {number} members_count
         * @memberof account.AgentGroup
         * @instance
         */
        AgentGroup.prototype.members_count = 0;

        return AgentGroup;
    })();

    account.ResetPasswordRequest = (function() {

        /**
         * Properties of a ResetPasswordRequest.
         * @memberof account
         * @interface IResetPasswordRequest
         * @property {common.IContext|null} [ctx] ResetPasswordRequest ctx
         * @property {string|null} [email] ResetPasswordRequest email
         */

        /**
         * Constructs a new ResetPasswordRequest.
         * @memberof account
         * @classdesc Represents a ResetPasswordRequest.
         * @implements IResetPasswordRequest
         * @constructor
         * @param {account.IResetPasswordRequest=} [p] Properties to set
         */
        function ResetPasswordRequest(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * ResetPasswordRequest ctx.
         * @member {common.IContext|null|undefined} ctx
         * @memberof account.ResetPasswordRequest
         * @instance
         */
        ResetPasswordRequest.prototype.ctx = null;

        /**
         * ResetPasswordRequest email.
         * @member {string} email
         * @memberof account.ResetPasswordRequest
         * @instance
         */
        ResetPasswordRequest.prototype.email = "";

        return ResetPasswordRequest;
    })();

    account.AgentPerm = (function() {

        /**
         * Properties of an AgentPerm.
         * @memberof account
         * @interface IAgentPerm
         * @property {common.IContext|null} [ctx] AgentPerm ctx
         * @property {string|null} [account_id] AgentPerm account_id
         * @property {string|null} [agent_id] AgentPerm agent_id
         * @property {auth.IMethod|null} [method] AgentPerm method
         */

        /**
         * Constructs a new AgentPerm.
         * @memberof account
         * @classdesc Represents an AgentPerm.
         * @implements IAgentPerm
         * @constructor
         * @param {account.IAgentPerm=} [p] Properties to set
         */
        function AgentPerm(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * AgentPerm ctx.
         * @member {common.IContext|null|undefined} ctx
         * @memberof account.AgentPerm
         * @instance
         */
        AgentPerm.prototype.ctx = null;

        /**
         * AgentPerm account_id.
         * @member {string} account_id
         * @memberof account.AgentPerm
         * @instance
         */
        AgentPerm.prototype.account_id = "";

        /**
         * AgentPerm agent_id.
         * @member {string} agent_id
         * @memberof account.AgentPerm
         * @instance
         */
        AgentPerm.prototype.agent_id = "";

        /**
         * AgentPerm method.
         * @member {auth.IMethod|null|undefined} method
         * @memberof account.AgentPerm
         * @instance
         */
        AgentPerm.prototype.method = null;

        return AgentPerm;
    })();

    account.Account = (function() {

        /**
         * Properties of an Account.
         * @memberof account
         * @interface IAccount
         * @property {common.IContext|null} [ctx] Account ctx
         * @property {string|null} [id] Account id
         * @property {string|null} [name] Account name
         * @property {string|null} [logo_url] Account logo_url
         * @property {string|null} [owner_id] Account owner_id
         * @property {string|null} [state] Account state
         * @property {number|Long|null} [created] Account created
         * @property {number|Long|null} [confirmed] Account confirmed
         * @property {number|Long|null} [modified] Account modified
         * @property {string|null} [slogan] Account slogan
         * @property {string|null} [referer_id] Account referer_id
         * @property {string|null} [city] Account city
         * @property {number|null} [zip_code] Account zip_code
         * @property {string|null} [tax_id] Account tax_id
         * @property {string|null} [facebook] Account facebook
         * @property {string|null} [twitter] Account twitter
         * @property {string|null} [phone] Account phone
         * @property {string|null} [address] Account address
         * @property {string|null} [url] Account url
         * @property {string|null} [lang] Account lang
         * @property {string|null} [referer_from] Account referer_from
         * @property {string|null} [timezone] Account timezone
         * @property {payment.ILimit|null} [limit] Account limit
         * @property {string|null} [country] Account country
         * @property {number|null} [v3_state] Account v3_state
         */

        /**
         * Constructs a new Account.
         * @memberof account
         * @classdesc Represents an Account.
         * @implements IAccount
         * @constructor
         * @param {account.IAccount=} [p] Properties to set
         */
        function Account(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * Account ctx.
         * @member {common.IContext|null|undefined} ctx
         * @memberof account.Account
         * @instance
         */
        Account.prototype.ctx = null;

        /**
         * Account id.
         * @member {string} id
         * @memberof account.Account
         * @instance
         */
        Account.prototype.id = "";

        /**
         * Account name.
         * @member {string} name
         * @memberof account.Account
         * @instance
         */
        Account.prototype.name = "";

        /**
         * Account logo_url.
         * @member {string} logo_url
         * @memberof account.Account
         * @instance
         */
        Account.prototype.logo_url = "";

        /**
         * Account owner_id.
         * @member {string} owner_id
         * @memberof account.Account
         * @instance
         */
        Account.prototype.owner_id = "";

        /**
         * Account state.
         * @member {string} state
         * @memberof account.Account
         * @instance
         */
        Account.prototype.state = "";

        /**
         * Account created.
         * @member {number|Long} created
         * @memberof account.Account
         * @instance
         */
        Account.prototype.created = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

        /**
         * Account confirmed.
         * @member {number|Long} confirmed
         * @memberof account.Account
         * @instance
         */
        Account.prototype.confirmed = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

        /**
         * Account modified.
         * @member {number|Long} modified
         * @memberof account.Account
         * @instance
         */
        Account.prototype.modified = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

        /**
         * Account slogan.
         * @member {string} slogan
         * @memberof account.Account
         * @instance
         */
        Account.prototype.slogan = "";

        /**
         * Account referer_id.
         * @member {string} referer_id
         * @memberof account.Account
         * @instance
         */
        Account.prototype.referer_id = "";

        /**
         * Account city.
         * @member {string} city
         * @memberof account.Account
         * @instance
         */
        Account.prototype.city = "";

        /**
         * Account zip_code.
         * @member {number} zip_code
         * @memberof account.Account
         * @instance
         */
        Account.prototype.zip_code = 0;

        /**
         * Account tax_id.
         * @member {string} tax_id
         * @memberof account.Account
         * @instance
         */
        Account.prototype.tax_id = "";

        /**
         * Account facebook.
         * @member {string} facebook
         * @memberof account.Account
         * @instance
         */
        Account.prototype.facebook = "";

        /**
         * Account twitter.
         * @member {string} twitter
         * @memberof account.Account
         * @instance
         */
        Account.prototype.twitter = "";

        /**
         * Account phone.
         * @member {string} phone
         * @memberof account.Account
         * @instance
         */
        Account.prototype.phone = "";

        /**
         * Account address.
         * @member {string} address
         * @memberof account.Account
         * @instance
         */
        Account.prototype.address = "";

        /**
         * Account url.
         * @member {string} url
         * @memberof account.Account
         * @instance
         */
        Account.prototype.url = "";

        /**
         * Account lang.
         * @member {string} lang
         * @memberof account.Account
         * @instance
         */
        Account.prototype.lang = "";

        /**
         * Account referer_from.
         * @member {string} referer_from
         * @memberof account.Account
         * @instance
         */
        Account.prototype.referer_from = "";

        /**
         * Account timezone.
         * @member {string} timezone
         * @memberof account.Account
         * @instance
         */
        Account.prototype.timezone = "";

        /**
         * Account limit.
         * @member {payment.ILimit|null|undefined} limit
         * @memberof account.Account
         * @instance
         */
        Account.prototype.limit = null;

        /**
         * Account country.
         * @member {string} country
         * @memberof account.Account
         * @instance
         */
        Account.prototype.country = "";

        /**
         * Account v3_state.
         * @member {number} v3_state
         * @memberof account.Account
         * @instance
         */
        Account.prototype.v3_state = 0;

        /**
         * State enum.
         * @name account.Account.State
         * @enum {string}
         * @property {number} pending=0 pending value
         * @property {number} activated=1 activated value
         * @property {number} locked=2 locked value
         * @property {number} deleted=3 deleted value
         */
        Account.State = (function() {
            const valuesById = {}, values = Object.create(valuesById);
            values[valuesById[0] = "pending"] = 0;
            values[valuesById[1] = "activated"] = 1;
            values[valuesById[2] = "locked"] = 2;
            values[valuesById[3] = "deleted"] = 3;
            return values;
        })();

        return Account;
    })();

    account.GroupMember = (function() {

        /**
         * Properties of a GroupMember.
         * @memberof account
         * @interface IGroupMember
         * @property {common.IContext|null} [ctx] GroupMember ctx
         * @property {string|null} [account_id] GroupMember account_id
         * @property {string|null} [group_id] GroupMember group_id
         * @property {string|null} [agent_id] GroupMember agent_id
         */

        /**
         * Constructs a new GroupMember.
         * @memberof account
         * @classdesc Represents a GroupMember.
         * @implements IGroupMember
         * @constructor
         * @param {account.IGroupMember=} [p] Properties to set
         */
        function GroupMember(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * GroupMember ctx.
         * @member {common.IContext|null|undefined} ctx
         * @memberof account.GroupMember
         * @instance
         */
        GroupMember.prototype.ctx = null;

        /**
         * GroupMember account_id.
         * @member {string} account_id
         * @memberof account.GroupMember
         * @instance
         */
        GroupMember.prototype.account_id = "";

        /**
         * GroupMember group_id.
         * @member {string} group_id
         * @memberof account.GroupMember
         * @instance
         */
        GroupMember.prototype.group_id = "";

        /**
         * GroupMember agent_id.
         * @member {string} agent_id
         * @memberof account.GroupMember
         * @instance
         */
        GroupMember.prototype.agent_id = "";

        return GroupMember;
    })();

    account.CreateAccountRequest = (function() {

        /**
         * Properties of a CreateAccountRequest.
         * @memberof account
         * @interface ICreateAccountRequest
         * @property {common.IContext|null} [ctx] CreateAccountRequest ctx
         * @property {string|null} [fullname] CreateAccountRequest fullname
         * @property {string|null} [email] CreateAccountRequest email
         * @property {string|null} [lang] CreateAccountRequest lang
         * @property {string|null} [password] CreateAccountRequest password
         * @property {string|null} [account_name] CreateAccountRequest account_name
         * @property {string|null} [account_url] CreateAccountRequest account_url
         * @property {string|null} [avatar_url] CreateAccountRequest avatar_url
         * @property {string|null} [timezone] CreateAccountRequest timezone
         * @property {string|null} [phone] CreateAccountRequest phone
         * @property {string|null} [country_code] CreateAccountRequest country_code
         */

        /**
         * Constructs a new CreateAccountRequest.
         * @memberof account
         * @classdesc Represents a CreateAccountRequest.
         * @implements ICreateAccountRequest
         * @constructor
         * @param {account.ICreateAccountRequest=} [p] Properties to set
         */
        function CreateAccountRequest(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * CreateAccountRequest ctx.
         * @member {common.IContext|null|undefined} ctx
         * @memberof account.CreateAccountRequest
         * @instance
         */
        CreateAccountRequest.prototype.ctx = null;

        /**
         * CreateAccountRequest fullname.
         * @member {string} fullname
         * @memberof account.CreateAccountRequest
         * @instance
         */
        CreateAccountRequest.prototype.fullname = "";

        /**
         * CreateAccountRequest email.
         * @member {string} email
         * @memberof account.CreateAccountRequest
         * @instance
         */
        CreateAccountRequest.prototype.email = "";

        /**
         * CreateAccountRequest lang.
         * @member {string} lang
         * @memberof account.CreateAccountRequest
         * @instance
         */
        CreateAccountRequest.prototype.lang = "";

        /**
         * CreateAccountRequest password.
         * @member {string} password
         * @memberof account.CreateAccountRequest
         * @instance
         */
        CreateAccountRequest.prototype.password = "";

        /**
         * CreateAccountRequest account_name.
         * @member {string} account_name
         * @memberof account.CreateAccountRequest
         * @instance
         */
        CreateAccountRequest.prototype.account_name = "";

        /**
         * CreateAccountRequest account_url.
         * @member {string} account_url
         * @memberof account.CreateAccountRequest
         * @instance
         */
        CreateAccountRequest.prototype.account_url = "";

        /**
         * CreateAccountRequest avatar_url.
         * @member {string} avatar_url
         * @memberof account.CreateAccountRequest
         * @instance
         */
        CreateAccountRequest.prototype.avatar_url = "";

        /**
         * CreateAccountRequest timezone.
         * @member {string} timezone
         * @memberof account.CreateAccountRequest
         * @instance
         */
        CreateAccountRequest.prototype.timezone = "";

        /**
         * CreateAccountRequest phone.
         * @member {string} phone
         * @memberof account.CreateAccountRequest
         * @instance
         */
        CreateAccountRequest.prototype.phone = "";

        /**
         * CreateAccountRequest country_code.
         * @member {string} country_code
         * @memberof account.CreateAccountRequest
         * @instance
         */
        CreateAccountRequest.prototype.country_code = "";

        return CreateAccountRequest;
    })();

    account.LoginRequest = (function() {

        /**
         * Properties of a LoginRequest.
         * @memberof account
         * @interface ILoginRequest
         * @property {common.IContext|null} [ctx] LoginRequest ctx
         * @property {string|null} [email] LoginRequest email
         * @property {string|null} [password] LoginRequest password
         */

        /**
         * Constructs a new LoginRequest.
         * @memberof account
         * @classdesc Represents a LoginRequest.
         * @implements ILoginRequest
         * @constructor
         * @param {account.ILoginRequest=} [p] Properties to set
         */
        function LoginRequest(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * LoginRequest ctx.
         * @member {common.IContext|null|undefined} ctx
         * @memberof account.LoginRequest
         * @instance
         */
        LoginRequest.prototype.ctx = null;

        /**
         * LoginRequest email.
         * @member {string} email
         * @memberof account.LoginRequest
         * @instance
         */
        LoginRequest.prototype.email = "";

        /**
         * LoginRequest password.
         * @member {string} password
         * @memberof account.LoginRequest
         * @instance
         */
        LoginRequest.prototype.password = "";

        return LoginRequest;
    })();

    account.Agents = (function() {

        /**
         * Properties of an Agents.
         * @memberof account
         * @interface IAgents
         * @property {common.IContext|null} [ctx] Agents ctx
         * @property {Array.<account.IAgent>|null} [Agents] Agents Agents
         */

        /**
         * Constructs a new Agents.
         * @memberof account
         * @classdesc Represents an Agents.
         * @implements IAgents
         * @constructor
         * @param {account.IAgents=} [p] Properties to set
         */
        function Agents(p) {
            this.Agents = [];
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * Agents ctx.
         * @member {common.IContext|null|undefined} ctx
         * @memberof account.Agents
         * @instance
         */
        Agents.prototype.ctx = null;

        /**
         * Agents Agents.
         * @member {Array.<account.IAgent>} Agents
         * @memberof account.Agents
         * @instance
         */
        Agents.prototype.Agents = $util.emptyArray;

        return Agents;
    })();

    account.NewPassword = (function() {

        /**
         * Properties of a NewPassword.
         * @memberof account
         * @interface INewPassword
         * @property {common.IContext|null} [ctx] NewPassword ctx
         * @property {string|null} [token] NewPassword token
         * @property {string|null} [new_password] NewPassword new_password
         * @property {string|null} [old_password] NewPassword old_password
         * @property {string|null} [email] NewPassword email
         */

        /**
         * Constructs a new NewPassword.
         * @memberof account
         * @classdesc Represents a NewPassword.
         * @implements INewPassword
         * @constructor
         * @param {account.INewPassword=} [p] Properties to set
         */
        function NewPassword(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * NewPassword ctx.
         * @member {common.IContext|null|undefined} ctx
         * @memberof account.NewPassword
         * @instance
         */
        NewPassword.prototype.ctx = null;

        /**
         * NewPassword token.
         * @member {string} token
         * @memberof account.NewPassword
         * @instance
         */
        NewPassword.prototype.token = "";

        /**
         * NewPassword new_password.
         * @member {string} new_password
         * @memberof account.NewPassword
         * @instance
         */
        NewPassword.prototype.new_password = "";

        /**
         * NewPassword old_password.
         * @member {string} old_password
         * @memberof account.NewPassword
         * @instance
         */
        NewPassword.prototype.old_password = "";

        /**
         * NewPassword email.
         * @member {string} email
         * @memberof account.NewPassword
         * @instance
         */
        NewPassword.prototype.email = "";

        return NewPassword;
    })();

    account.AgentGroups = (function() {

        /**
         * Properties of an AgentGroups.
         * @memberof account
         * @interface IAgentGroups
         * @property {common.IContext|null} [ctx] AgentGroups ctx
         * @property {Array.<account.IAgentGroup>|null} [Groups] AgentGroups Groups
         */

        /**
         * Constructs a new AgentGroups.
         * @memberof account
         * @classdesc Represents an AgentGroups.
         * @implements IAgentGroups
         * @constructor
         * @param {account.IAgentGroups=} [p] Properties to set
         */
        function AgentGroups(p) {
            this.Groups = [];
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * AgentGroups ctx.
         * @member {common.IContext|null|undefined} ctx
         * @memberof account.AgentGroups
         * @instance
         */
        AgentGroups.prototype.ctx = null;

        /**
         * AgentGroups Groups.
         * @member {Array.<account.IAgentGroup>} Groups
         * @memberof account.AgentGroups
         * @instance
         */
        AgentGroups.prototype.Groups = $util.emptyArray;

        return AgentGroups;
    })();

    account.Token = (function() {

        /**
         * Properties of a Token.
         * @memberof account
         * @interface IToken
         * @property {common.IContext|null} [ctx] Token ctx
         * @property {boolean|null} [is_set] Token is_set
         * @property {string|null} [token] Token token
         * @property {string|null} [account_id] Token account_id
         */

        /**
         * Constructs a new Token.
         * @memberof account
         * @classdesc Represents a Token.
         * @implements IToken
         * @constructor
         * @param {account.IToken=} [p] Properties to set
         */
        function Token(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * Token ctx.
         * @member {common.IContext|null|undefined} ctx
         * @memberof account.Token
         * @instance
         */
        Token.prototype.ctx = null;

        /**
         * Token is_set.
         * @member {boolean} is_set
         * @memberof account.Token
         * @instance
         */
        Token.prototype.is_set = false;

        /**
         * Token token.
         * @member {string} token
         * @memberof account.Token
         * @instance
         */
        Token.prototype.token = "";

        /**
         * Token account_id.
         * @member {string} account_id
         * @memberof account.Token
         * @instance
         */
        Token.prototype.account_id = "";

        return Token;
    })();

    account.AccountV3 = (function() {

        /**
         * Properties of an AccountV3.
         * @memberof account
         * @interface IAccountV3
         * @property {common.IContext|null} [ctx] AccountV3 ctx
         * @property {account.IAccount|null} [account] AccountV3 account
         * @property {account.IAgent|null} [owner] AccountV3 owner
         * @property {payment.ISubscription|null} [subscription] AccountV3 subscription
         */

        /**
         * Constructs a new AccountV3.
         * @memberof account
         * @classdesc Represents an AccountV3.
         * @implements IAccountV3
         * @constructor
         * @param {account.IAccountV3=} [p] Properties to set
         */
        function AccountV3(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * AccountV3 ctx.
         * @member {common.IContext|null|undefined} ctx
         * @memberof account.AccountV3
         * @instance
         */
        AccountV3.prototype.ctx = null;

        /**
         * AccountV3 account.
         * @member {account.IAccount|null|undefined} account
         * @memberof account.AccountV3
         * @instance
         */
        AccountV3.prototype.account = null;

        /**
         * AccountV3 owner.
         * @member {account.IAgent|null|undefined} owner
         * @memberof account.AccountV3
         * @instance
         */
        AccountV3.prototype.owner = null;

        /**
         * AccountV3 subscription.
         * @member {payment.ISubscription|null|undefined} subscription
         * @memberof account.AccountV3
         * @instance
         */
        AccountV3.prototype.subscription = null;

        return AccountV3;
    })();

    /**
     * Event enum.
     * @name account.Event
     * @enum {string}
     * @property {number} AccountRequested=1000 AccountRequested value
     * @property {number} AccountSynced=1001 AccountSynced value
     * @property {number} AccountV3Synced=1002 AccountV3Synced value
     * @property {number} AgentGroupDeleted=0 AgentGroupDeleted value
     * @property {number} AgentLeftGroup=1 AgentLeftGroup value
     * @property {number} AgentJoinedGroup=2 AgentJoinedGroup value
     * @property {number} AgentGroupUpserted=4 AgentGroupUpserted value
     * @property {number} AgentUpserted=6 AgentUpserted value
     * @property {number} AgentPermissionUpdated=9 AgentPermissionUpdated value
     * @property {number} AccountUpserted=14 AccountUpserted value
     * @property {number} AccountPlanUpdated=16 AccountPlanUpdated value
     * @property {number} AccountStateUpdated=17 AccountStateUpdated value
     * @property {number} AccountConfirmRequest=19 AccountConfirmRequest value
     * @property {number} HandleExpiredInvitation=20 HandleExpiredInvitation value
     * @property {number} AccountConfirmRequested=21 AccountConfirmRequested value
     * @property {number} AccountConfirmSuccessEmailRequested=25 AccountConfirmSuccessEmailRequested value
     * @property {number} AccountResetPasswordEmail=33 AccountResetPasswordEmail value
     * @property {number} AccountPasswordChangedEmailRequested=34 AccountPasswordChangedEmailRequested value
     * @property {number} AccountInviteEmail=22 AccountInviteEmail value
     * @property {number} AccountDeleted=24 AccountDeleted value
     * @property {number} AccountCreated=45 AccountCreated value
     * @property {number} AccountActivated=46 AccountActivated value
     * @property {number} AccountInfoUpdated=47 AccountInfoUpdated value
     * @property {number} AgentRejected=10 AgentRejected value
     * @property {number} AgentAccepted=11 AgentAccepted value
     * @property {number} AgentInvited=13 AgentInvited value
     * @property {number} AgentDeleted=15 AgentDeleted value
     * @property {number} AgentActivated=50 AgentActivated value
     * @property {number} AgentDeactivated=51 AgentDeactivated value
     * @property {number} AgentInfoUpdated=59 AgentInfoUpdated value
     * @property {number} AccountV3Created=60 AccountV3Created value
     * @property {number} AccountPaymentV3Synced=63 AccountPaymentV3Synced value
     */
    account.Event = (function() {
        const valuesById = {}, values = Object.create(valuesById);
        values[valuesById[1000] = "AccountRequested"] = 1000;
        values[valuesById[1001] = "AccountSynced"] = 1001;
        values[valuesById[1002] = "AccountV3Synced"] = 1002;
        values[valuesById[0] = "AgentGroupDeleted"] = 0;
        values[valuesById[1] = "AgentLeftGroup"] = 1;
        values[valuesById[2] = "AgentJoinedGroup"] = 2;
        values[valuesById[4] = "AgentGroupUpserted"] = 4;
        values[valuesById[6] = "AgentUpserted"] = 6;
        values[valuesById[9] = "AgentPermissionUpdated"] = 9;
        values[valuesById[14] = "AccountUpserted"] = 14;
        values[valuesById[16] = "AccountPlanUpdated"] = 16;
        values[valuesById[17] = "AccountStateUpdated"] = 17;
        values[valuesById[19] = "AccountConfirmRequest"] = 19;
        values[valuesById[20] = "HandleExpiredInvitation"] = 20;
        values[valuesById[21] = "AccountConfirmRequested"] = 21;
        values[valuesById[25] = "AccountConfirmSuccessEmailRequested"] = 25;
        values[valuesById[33] = "AccountResetPasswordEmail"] = 33;
        values[valuesById[34] = "AccountPasswordChangedEmailRequested"] = 34;
        values[valuesById[22] = "AccountInviteEmail"] = 22;
        values[valuesById[24] = "AccountDeleted"] = 24;
        values[valuesById[45] = "AccountCreated"] = 45;
        values[valuesById[46] = "AccountActivated"] = 46;
        values[valuesById[47] = "AccountInfoUpdated"] = 47;
        values[valuesById[10] = "AgentRejected"] = 10;
        values[valuesById[11] = "AgentAccepted"] = 11;
        values[valuesById[13] = "AgentInvited"] = 13;
        values[valuesById[15] = "AgentDeleted"] = 15;
        values[valuesById[50] = "AgentActivated"] = 50;
        values[valuesById[51] = "AgentDeactivated"] = 51;
        values[valuesById[59] = "AgentInfoUpdated"] = 59;
        values[valuesById[60] = "AccountV3Created"] = 60;
        values[valuesById[63] = "AccountPaymentV3Synced"] = 63;
        return values;
    })();

    account.ConfirmEmail = (function() {

        /**
         * Properties of a ConfirmEmail.
         * @memberof account
         * @interface IConfirmEmail
         * @property {common.IContext|null} [ctx] ConfirmEmail ctx
         * @property {string|null} [to] ConfirmEmail to
         * @property {string|null} [account_id] ConfirmEmail account_id
         * @property {string|null} [owner_id] ConfirmEmail owner_id
         * @property {string|null} [token] ConfirmEmail token
         * @property {number|Long|null} [expired_in] ConfirmEmail expired_in
         * @property {string|null} [account_name] ConfirmEmail account_name
         * @property {lang.L|null} [lang] ConfirmEmail lang
         * @property {string|null} [owner_name] ConfirmEmail owner_name
         * @property {string|null} [from] ConfirmEmail from
         */

        /**
         * Constructs a new ConfirmEmail.
         * @memberof account
         * @classdesc Represents a ConfirmEmail.
         * @implements IConfirmEmail
         * @constructor
         * @param {account.IConfirmEmail=} [p] Properties to set
         */
        function ConfirmEmail(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * ConfirmEmail ctx.
         * @member {common.IContext|null|undefined} ctx
         * @memberof account.ConfirmEmail
         * @instance
         */
        ConfirmEmail.prototype.ctx = null;

        /**
         * ConfirmEmail to.
         * @member {string} to
         * @memberof account.ConfirmEmail
         * @instance
         */
        ConfirmEmail.prototype.to = "";

        /**
         * ConfirmEmail account_id.
         * @member {string} account_id
         * @memberof account.ConfirmEmail
         * @instance
         */
        ConfirmEmail.prototype.account_id = "";

        /**
         * ConfirmEmail owner_id.
         * @member {string} owner_id
         * @memberof account.ConfirmEmail
         * @instance
         */
        ConfirmEmail.prototype.owner_id = "";

        /**
         * ConfirmEmail token.
         * @member {string} token
         * @memberof account.ConfirmEmail
         * @instance
         */
        ConfirmEmail.prototype.token = "";

        /**
         * ConfirmEmail expired_in.
         * @member {number|Long} expired_in
         * @memberof account.ConfirmEmail
         * @instance
         */
        ConfirmEmail.prototype.expired_in = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

        /**
         * ConfirmEmail account_name.
         * @member {string} account_name
         * @memberof account.ConfirmEmail
         * @instance
         */
        ConfirmEmail.prototype.account_name = "";

        /**
         * ConfirmEmail lang.
         * @member {lang.L} lang
         * @memberof account.ConfirmEmail
         * @instance
         */
        ConfirmEmail.prototype.lang = 0;

        /**
         * ConfirmEmail owner_name.
         * @member {string} owner_name
         * @memberof account.ConfirmEmail
         * @instance
         */
        ConfirmEmail.prototype.owner_name = "";

        /**
         * ConfirmEmail from.
         * @member {string} from
         * @memberof account.ConfirmEmail
         * @instance
         */
        ConfirmEmail.prototype.from = "";

        return ConfirmEmail;
    })();

    account.InviteEmail = (function() {

        /**
         * Properties of an InviteEmail.
         * @memberof account
         * @interface IInviteEmail
         * @property {common.IContext|null} [ctx] InviteEmail ctx
         * @property {string|null} [to] InviteEmail to
         * @property {string|null} [account_id] InviteEmail account_id
         * @property {string|null} [sender_id] InviteEmail sender_id
         * @property {number|Long|null} [expired_in] InviteEmail expired_in
         * @property {string|null} [token] InviteEmail token
         * @property {lang.L|null} [lang] InviteEmail lang
         * @property {string|null} [fullname] InviteEmail fullname
         * @property {string|null} [sender_name] InviteEmail sender_name
         * @property {string|null} [account_name] InviteEmail account_name
         * @property {string|null} [account_logo_url] InviteEmail account_logo_url
         * @property {string|null} [from] InviteEmail from
         */

        /**
         * Constructs a new InviteEmail.
         * @memberof account
         * @classdesc Represents an InviteEmail.
         * @implements IInviteEmail
         * @constructor
         * @param {account.IInviteEmail=} [p] Properties to set
         */
        function InviteEmail(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * InviteEmail ctx.
         * @member {common.IContext|null|undefined} ctx
         * @memberof account.InviteEmail
         * @instance
         */
        InviteEmail.prototype.ctx = null;

        /**
         * InviteEmail to.
         * @member {string} to
         * @memberof account.InviteEmail
         * @instance
         */
        InviteEmail.prototype.to = "";

        /**
         * InviteEmail account_id.
         * @member {string} account_id
         * @memberof account.InviteEmail
         * @instance
         */
        InviteEmail.prototype.account_id = "";

        /**
         * InviteEmail sender_id.
         * @member {string} sender_id
         * @memberof account.InviteEmail
         * @instance
         */
        InviteEmail.prototype.sender_id = "";

        /**
         * InviteEmail expired_in.
         * @member {number|Long} expired_in
         * @memberof account.InviteEmail
         * @instance
         */
        InviteEmail.prototype.expired_in = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

        /**
         * InviteEmail token.
         * @member {string} token
         * @memberof account.InviteEmail
         * @instance
         */
        InviteEmail.prototype.token = "";

        /**
         * InviteEmail lang.
         * @member {lang.L} lang
         * @memberof account.InviteEmail
         * @instance
         */
        InviteEmail.prototype.lang = 0;

        /**
         * InviteEmail fullname.
         * @member {string} fullname
         * @memberof account.InviteEmail
         * @instance
         */
        InviteEmail.prototype.fullname = "";

        /**
         * InviteEmail sender_name.
         * @member {string} sender_name
         * @memberof account.InviteEmail
         * @instance
         */
        InviteEmail.prototype.sender_name = "";

        /**
         * InviteEmail account_name.
         * @member {string} account_name
         * @memberof account.InviteEmail
         * @instance
         */
        InviteEmail.prototype.account_name = "";

        /**
         * InviteEmail account_logo_url.
         * @member {string} account_logo_url
         * @memberof account.InviteEmail
         * @instance
         */
        InviteEmail.prototype.account_logo_url = "";

        /**
         * InviteEmail from.
         * @member {string} from
         * @memberof account.InviteEmail
         * @instance
         */
        InviteEmail.prototype.from = "";

        return InviteEmail;
    })();

    account.ResetPasswordEmail = (function() {

        /**
         * Properties of a ResetPasswordEmail.
         * @memberof account
         * @interface IResetPasswordEmail
         * @property {common.IContext|null} [ctx] ResetPasswordEmail ctx
         * @property {string|null} [from] ResetPasswordEmail from
         * @property {string|null} [to] ResetPasswordEmail to
         * @property {number|Long|null} [expired_in] ResetPasswordEmail expired_in
         * @property {string|null} [token] ResetPasswordEmail token
         * @property {string|null} [account_id] ResetPasswordEmail account_id
         * @property {string|null} [agent_id] ResetPasswordEmail agent_id
         * @property {string|null} [agent_name] ResetPasswordEmail agent_name
         * @property {lang.L|null} [lang] ResetPasswordEmail lang
         */

        /**
         * Constructs a new ResetPasswordEmail.
         * @memberof account
         * @classdesc Represents a ResetPasswordEmail.
         * @implements IResetPasswordEmail
         * @constructor
         * @param {account.IResetPasswordEmail=} [p] Properties to set
         */
        function ResetPasswordEmail(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * ResetPasswordEmail ctx.
         * @member {common.IContext|null|undefined} ctx
         * @memberof account.ResetPasswordEmail
         * @instance
         */
        ResetPasswordEmail.prototype.ctx = null;

        /**
         * ResetPasswordEmail from.
         * @member {string} from
         * @memberof account.ResetPasswordEmail
         * @instance
         */
        ResetPasswordEmail.prototype.from = "";

        /**
         * ResetPasswordEmail to.
         * @member {string} to
         * @memberof account.ResetPasswordEmail
         * @instance
         */
        ResetPasswordEmail.prototype.to = "";

        /**
         * ResetPasswordEmail expired_in.
         * @member {number|Long} expired_in
         * @memberof account.ResetPasswordEmail
         * @instance
         */
        ResetPasswordEmail.prototype.expired_in = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

        /**
         * ResetPasswordEmail token.
         * @member {string} token
         * @memberof account.ResetPasswordEmail
         * @instance
         */
        ResetPasswordEmail.prototype.token = "";

        /**
         * ResetPasswordEmail account_id.
         * @member {string} account_id
         * @memberof account.ResetPasswordEmail
         * @instance
         */
        ResetPasswordEmail.prototype.account_id = "";

        /**
         * ResetPasswordEmail agent_id.
         * @member {string} agent_id
         * @memberof account.ResetPasswordEmail
         * @instance
         */
        ResetPasswordEmail.prototype.agent_id = "";

        /**
         * ResetPasswordEmail agent_name.
         * @member {string} agent_name
         * @memberof account.ResetPasswordEmail
         * @instance
         */
        ResetPasswordEmail.prototype.agent_name = "";

        /**
         * ResetPasswordEmail lang.
         * @member {lang.L} lang
         * @memberof account.ResetPasswordEmail
         * @instance
         */
        ResetPasswordEmail.prototype.lang = 0;

        return ResetPasswordEmail;
    })();

    account.PasswordChangedEmail = (function() {

        /**
         * Properties of a PasswordChangedEmail.
         * @memberof account
         * @interface IPasswordChangedEmail
         * @property {common.IContext|null} [ctx] PasswordChangedEmail ctx
         * @property {string|null} [to] PasswordChangedEmail to
         * @property {string|null} [account_id] PasswordChangedEmail account_id
         * @property {string|null} [agent_id] PasswordChangedEmail agent_id
         * @property {string|null} [agent_name] PasswordChangedEmail agent_name
         * @property {lang.L|null} [lang] PasswordChangedEmail lang
         * @property {string|null} [from] PasswordChangedEmail from
         */

        /**
         * Constructs a new PasswordChangedEmail.
         * @memberof account
         * @classdesc Represents a PasswordChangedEmail.
         * @implements IPasswordChangedEmail
         * @constructor
         * @param {account.IPasswordChangedEmail=} [p] Properties to set
         */
        function PasswordChangedEmail(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * PasswordChangedEmail ctx.
         * @member {common.IContext|null|undefined} ctx
         * @memberof account.PasswordChangedEmail
         * @instance
         */
        PasswordChangedEmail.prototype.ctx = null;

        /**
         * PasswordChangedEmail to.
         * @member {string} to
         * @memberof account.PasswordChangedEmail
         * @instance
         */
        PasswordChangedEmail.prototype.to = "";

        /**
         * PasswordChangedEmail account_id.
         * @member {string} account_id
         * @memberof account.PasswordChangedEmail
         * @instance
         */
        PasswordChangedEmail.prototype.account_id = "";

        /**
         * PasswordChangedEmail agent_id.
         * @member {string} agent_id
         * @memberof account.PasswordChangedEmail
         * @instance
         */
        PasswordChangedEmail.prototype.agent_id = "";

        /**
         * PasswordChangedEmail agent_name.
         * @member {string} agent_name
         * @memberof account.PasswordChangedEmail
         * @instance
         */
        PasswordChangedEmail.prototype.agent_name = "";

        /**
         * PasswordChangedEmail lang.
         * @member {lang.L} lang
         * @memberof account.PasswordChangedEmail
         * @instance
         */
        PasswordChangedEmail.prototype.lang = 0;

        /**
         * PasswordChangedEmail from.
         * @member {string} from
         * @memberof account.PasswordChangedEmail
         * @instance
         */
        PasswordChangedEmail.prototype.from = "";

        return PasswordChangedEmail;
    })();

    account.AccountConfirmSuccessEmail = (function() {

        /**
         * Properties of an AccountConfirmSuccessEmail.
         * @memberof account
         * @interface IAccountConfirmSuccessEmail
         * @property {common.IContext|null} [ctx] AccountConfirmSuccessEmail ctx
         * @property {string|null} [to] AccountConfirmSuccessEmail to
         * @property {string|null} [account_id] AccountConfirmSuccessEmail account_id
         * @property {string|null} [agent_id] AccountConfirmSuccessEmail agent_id
         * @property {string|null} [agent_name] AccountConfirmSuccessEmail agent_name
         * @property {lang.L|null} [lang] AccountConfirmSuccessEmail lang
         * @property {string|null} [from] AccountConfirmSuccessEmail from
         * @property {number|Long|null} [from_time] AccountConfirmSuccessEmail from_time
         * @property {number|Long|null} [to_time] AccountConfirmSuccessEmail to_time
         */

        /**
         * Constructs a new AccountConfirmSuccessEmail.
         * @memberof account
         * @classdesc Represents an AccountConfirmSuccessEmail.
         * @implements IAccountConfirmSuccessEmail
         * @constructor
         * @param {account.IAccountConfirmSuccessEmail=} [p] Properties to set
         */
        function AccountConfirmSuccessEmail(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * AccountConfirmSuccessEmail ctx.
         * @member {common.IContext|null|undefined} ctx
         * @memberof account.AccountConfirmSuccessEmail
         * @instance
         */
        AccountConfirmSuccessEmail.prototype.ctx = null;

        /**
         * AccountConfirmSuccessEmail to.
         * @member {string} to
         * @memberof account.AccountConfirmSuccessEmail
         * @instance
         */
        AccountConfirmSuccessEmail.prototype.to = "";

        /**
         * AccountConfirmSuccessEmail account_id.
         * @member {string} account_id
         * @memberof account.AccountConfirmSuccessEmail
         * @instance
         */
        AccountConfirmSuccessEmail.prototype.account_id = "";

        /**
         * AccountConfirmSuccessEmail agent_id.
         * @member {string} agent_id
         * @memberof account.AccountConfirmSuccessEmail
         * @instance
         */
        AccountConfirmSuccessEmail.prototype.agent_id = "";

        /**
         * AccountConfirmSuccessEmail agent_name.
         * @member {string} agent_name
         * @memberof account.AccountConfirmSuccessEmail
         * @instance
         */
        AccountConfirmSuccessEmail.prototype.agent_name = "";

        /**
         * AccountConfirmSuccessEmail lang.
         * @member {lang.L} lang
         * @memberof account.AccountConfirmSuccessEmail
         * @instance
         */
        AccountConfirmSuccessEmail.prototype.lang = 0;

        /**
         * AccountConfirmSuccessEmail from.
         * @member {string} from
         * @memberof account.AccountConfirmSuccessEmail
         * @instance
         */
        AccountConfirmSuccessEmail.prototype.from = "";

        /**
         * AccountConfirmSuccessEmail from_time.
         * @member {number|Long} from_time
         * @memberof account.AccountConfirmSuccessEmail
         * @instance
         */
        AccountConfirmSuccessEmail.prototype.from_time = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

        /**
         * AccountConfirmSuccessEmail to_time.
         * @member {number|Long} to_time
         * @memberof account.AccountConfirmSuccessEmail
         * @instance
         */
        AccountConfirmSuccessEmail.prototype.to_time = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

        return AccountConfirmSuccessEmail;
    })();

    account.AccountMgr = (function() {

        /**
         * Constructs a new AccountMgr service.
         * @memberof account
         * @classdesc Represents an AccountMgr
         * @extends $protobuf.rpc.Service
         * @constructor
         * @param {$protobuf.RPCImpl} rpcImpl RPC implementation
         * @param {boolean} [requestDelimited=false] Whether requests are length-delimited
         * @param {boolean} [responseDelimited=false] Whether responses are length-delimited
         */
        function AccountMgr(rpcImpl, requestDelimited, responseDelimited) {
            $protobuf.rpc.Service.call(this, rpcImpl, requestDelimited, responseDelimited);
        }

        (AccountMgr.prototype = Object.create($protobuf.rpc.Service.prototype)).constructor = AccountMgr;

        /**
         * Callback as used by {@link account.AccountMgr#createGroup}.
         * @memberof account.AccountMgr
         * @typedef CreateGroupCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {account.AgentGroup} [response] AgentGroup
         */

        /**
         * Calls CreateGroup.
         * @function createGroup
         * @memberof account.AccountMgr
         * @instance
         * @param {account.IAgentGroup} request AgentGroup message or plain object
         * @param {account.AccountMgr.CreateGroupCallback} callback Node-style callback called with the error, if any, and AgentGroup
         * @returns {undefined}
         * @variation 1
         */
        AccountMgr.prototype.createGroup = function createGroup(request, callback) {
            return this.rpcCall(createGroup, $root.account.AgentGroup, $root.account.AgentGroup, request, callback);
        };

        /**
         * Calls CreateGroup.
         * @function createGroup
         * @memberof account.AccountMgr
         * @instance
         * @param {account.IAgentGroup} request AgentGroup message or plain object
         * @returns {Promise<account.AgentGroup>} Promise
         * @variation 2
         */

        /**
         * Callback as used by {@link account.AccountMgr#updateGroup}.
         * @memberof account.AccountMgr
         * @typedef UpdateGroupCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {account.AgentGroup} [response] AgentGroup
         */

        /**
         * Calls UpdateGroup.
         * @function updateGroup
         * @memberof account.AccountMgr
         * @instance
         * @param {account.IAgentGroup} request AgentGroup message or plain object
         * @param {account.AccountMgr.UpdateGroupCallback} callback Node-style callback called with the error, if any, and AgentGroup
         * @returns {undefined}
         * @variation 1
         */
        AccountMgr.prototype.updateGroup = function updateGroup(request, callback) {
            return this.rpcCall(updateGroup, $root.account.AgentGroup, $root.account.AgentGroup, request, callback);
        };

        /**
         * Calls UpdateGroup.
         * @function updateGroup
         * @memberof account.AccountMgr
         * @instance
         * @param {account.IAgentGroup} request AgentGroup message or plain object
         * @returns {Promise<account.AgentGroup>} Promise
         * @variation 2
         */

        /**
         * Callback as used by {@link account.AccountMgr#getGroup}.
         * @memberof account.AccountMgr
         * @typedef GetGroupCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {account.AgentGroup} [response] AgentGroup
         */

        /**
         * Calls GetGroup.
         * @function getGroup
         * @memberof account.AccountMgr
         * @instance
         * @param {common.IId} request Id message or plain object
         * @param {account.AccountMgr.GetGroupCallback} callback Node-style callback called with the error, if any, and AgentGroup
         * @returns {undefined}
         * @variation 1
         */
        AccountMgr.prototype.getGroup = function getGroup(request, callback) {
            return this.rpcCall(getGroup, $root.common.Id, $root.account.AgentGroup, request, callback);
        };

        /**
         * Calls GetGroup.
         * @function getGroup
         * @memberof account.AccountMgr
         * @instance
         * @param {common.IId} request Id message or plain object
         * @returns {Promise<account.AgentGroup>} Promise
         * @variation 2
         */

        /**
         * Callback as used by {@link account.AccountMgr#getAgentPermission}.
         * @memberof account.AccountMgr
         * @typedef GetAgentPermissionCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {auth.Method} [response] Method
         */

        /**
         * Calls GetAgentPermission.
         * @function getAgentPermission
         * @memberof account.AccountMgr
         * @instance
         * @param {common.IId} request Id message or plain object
         * @param {account.AccountMgr.GetAgentPermissionCallback} callback Node-style callback called with the error, if any, and Method
         * @returns {undefined}
         * @variation 1
         */
        AccountMgr.prototype.getAgentPermission = function getAgentPermission(request, callback) {
            return this.rpcCall(getAgentPermission, $root.common.Id, $root.auth.Method, request, callback);
        };

        /**
         * Calls GetAgentPermission.
         * @function getAgentPermission
         * @memberof account.AccountMgr
         * @instance
         * @param {common.IId} request Id message or plain object
         * @returns {Promise<auth.Method>} Promise
         * @variation 2
         */

        /**
         * Callback as used by {@link account.AccountMgr#updateAgentPermission}.
         * @memberof account.AccountMgr
         * @typedef UpdateAgentPermissionCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {auth.Method} [response] Method
         */

        /**
         * Calls UpdateAgentPermission.
         * @function updateAgentPermission
         * @memberof account.AccountMgr
         * @instance
         * @param {account.IAgentPerm} request AgentPerm message or plain object
         * @param {account.AccountMgr.UpdateAgentPermissionCallback} callback Node-style callback called with the error, if any, and Method
         * @returns {undefined}
         * @variation 1
         */
        AccountMgr.prototype.updateAgentPermission = function updateAgentPermission(request, callback) {
            return this.rpcCall(updateAgentPermission, $root.account.AgentPerm, $root.auth.Method, request, callback);
        };

        /**
         * Calls UpdateAgentPermission.
         * @function updateAgentPermission
         * @memberof account.AccountMgr
         * @instance
         * @param {account.IAgentPerm} request AgentPerm message or plain object
         * @returns {Promise<auth.Method>} Promise
         * @variation 2
         */

        /**
         * Callback as used by {@link account.AccountMgr#requestResetPassword}.
         * @memberof account.AccountMgr
         * @typedef RequestResetPasswordCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {common.Empty} [response] Empty
         */

        /**
         * Calls RequestResetPassword.
         * @function requestResetPassword
         * @memberof account.AccountMgr
         * @instance
         * @param {account.IResetPasswordRequest} request ResetPasswordRequest message or plain object
         * @param {account.AccountMgr.RequestResetPasswordCallback} callback Node-style callback called with the error, if any, and Empty
         * @returns {undefined}
         * @variation 1
         */
        AccountMgr.prototype.requestResetPassword = function requestResetPassword(request, callback) {
            return this.rpcCall(requestResetPassword, $root.account.ResetPasswordRequest, $root.common.Empty, request, callback);
        };

        /**
         * Calls RequestResetPassword.
         * @function requestResetPassword
         * @memberof account.AccountMgr
         * @instance
         * @param {account.IResetPasswordRequest} request ResetPasswordRequest message or plain object
         * @returns {Promise<common.Empty>} Promise
         * @variation 2
         */

        /**
         * Callback as used by {@link account.AccountMgr#updatePassword}.
         * @memberof account.AccountMgr
         * @typedef UpdatePasswordCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {account.Agent} [response] Agent
         */

        /**
         * Calls UpdatePassword.
         * @function updatePassword
         * @memberof account.AccountMgr
         * @instance
         * @param {account.INewPassword} request NewPassword message or plain object
         * @param {account.AccountMgr.UpdatePasswordCallback} callback Node-style callback called with the error, if any, and Agent
         * @returns {undefined}
         * @variation 1
         */
        AccountMgr.prototype.updatePassword = function updatePassword(request, callback) {
            return this.rpcCall(updatePassword, $root.account.NewPassword, $root.account.Agent, request, callback);
        };

        /**
         * Calls UpdatePassword.
         * @function updatePassword
         * @memberof account.AccountMgr
         * @instance
         * @param {account.INewPassword} request NewPassword message or plain object
         * @returns {Promise<account.Agent>} Promise
         * @variation 2
         */

        /**
         * Callback as used by {@link account.AccountMgr#updateAgent}.
         * @memberof account.AccountMgr
         * @typedef UpdateAgentCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {account.Agent} [response] Agent
         */

        /**
         * Calls UpdateAgent.
         * @function updateAgent
         * @memberof account.AccountMgr
         * @instance
         * @param {account.IAgent} request Agent message or plain object
         * @param {account.AccountMgr.UpdateAgentCallback} callback Node-style callback called with the error, if any, and Agent
         * @returns {undefined}
         * @variation 1
         */
        AccountMgr.prototype.updateAgent = function updateAgent(request, callback) {
            return this.rpcCall(updateAgent, $root.account.Agent, $root.account.Agent, request, callback);
        };

        /**
         * Calls UpdateAgent.
         * @function updateAgent
         * @memberof account.AccountMgr
         * @instance
         * @param {account.IAgent} request Agent message or plain object
         * @returns {Promise<account.Agent>} Promise
         * @variation 2
         */

        /**
         * Callback as used by {@link account.AccountMgr#deleteAgent}.
         * @memberof account.AccountMgr
         * @typedef DeleteAgentCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {account.Agent} [response] Agent
         */

        /**
         * Calls DeleteAgent.
         * @function deleteAgent
         * @memberof account.AccountMgr
         * @instance
         * @param {common.IId} request Id message or plain object
         * @param {account.AccountMgr.DeleteAgentCallback} callback Node-style callback called with the error, if any, and Agent
         * @returns {undefined}
         * @variation 1
         */
        AccountMgr.prototype.deleteAgent = function deleteAgent(request, callback) {
            return this.rpcCall(deleteAgent, $root.common.Id, $root.account.Agent, request, callback);
        };

        /**
         * Calls DeleteAgent.
         * @function deleteAgent
         * @memberof account.AccountMgr
         * @instance
         * @param {common.IId} request Id message or plain object
         * @returns {Promise<account.Agent>} Promise
         * @variation 2
         */

        /**
         * Callback as used by {@link account.AccountMgr#acceptInvitation}.
         * @memberof account.AccountMgr
         * @typedef AcceptInvitationCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {account.Agent} [response] Agent
         */

        /**
         * Calls AcceptInvitation.
         * @function acceptInvitation
         * @memberof account.AccountMgr
         * @instance
         * @param {account.INewPassword} request NewPassword message or plain object
         * @param {account.AccountMgr.AcceptInvitationCallback} callback Node-style callback called with the error, if any, and Agent
         * @returns {undefined}
         * @variation 1
         */
        AccountMgr.prototype.acceptInvitation = function acceptInvitation(request, callback) {
            return this.rpcCall(acceptInvitation, $root.account.NewPassword, $root.account.Agent, request, callback);
        };

        /**
         * Calls AcceptInvitation.
         * @function acceptInvitation
         * @memberof account.AccountMgr
         * @instance
         * @param {account.INewPassword} request NewPassword message or plain object
         * @returns {Promise<account.Agent>} Promise
         * @variation 2
         */

        /**
         * Callback as used by {@link account.AccountMgr#getInvitation}.
         * @memberof account.AccountMgr
         * @typedef GetInvitationCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {account.Agent} [response] Agent
         */

        /**
         * Calls GetInvitation.
         * @function getInvitation
         * @memberof account.AccountMgr
         * @instance
         * @param {account.IToken} request Token message or plain object
         * @param {account.AccountMgr.GetInvitationCallback} callback Node-style callback called with the error, if any, and Agent
         * @returns {undefined}
         * @variation 1
         */
        AccountMgr.prototype.getInvitation = function getInvitation(request, callback) {
            return this.rpcCall(getInvitation, $root.account.Token, $root.account.Agent, request, callback);
        };

        /**
         * Calls GetInvitation.
         * @function getInvitation
         * @memberof account.AccountMgr
         * @instance
         * @param {account.IToken} request Token message or plain object
         * @returns {Promise<account.Agent>} Promise
         * @variation 2
         */

        /**
         * Callback as used by {@link account.AccountMgr#inviteAgent}.
         * @memberof account.AccountMgr
         * @typedef InviteAgentCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {account.Agent} [response] Agent
         */

        /**
         * Calls InviteAgent.
         * @function inviteAgent
         * @memberof account.AccountMgr
         * @instance
         * @param {account.IAgent} request Agent message or plain object
         * @param {account.AccountMgr.InviteAgentCallback} callback Node-style callback called with the error, if any, and Agent
         * @returns {undefined}
         * @variation 1
         */
        AccountMgr.prototype.inviteAgent = function inviteAgent(request, callback) {
            return this.rpcCall(inviteAgent, $root.account.Agent, $root.account.Agent, request, callback);
        };

        /**
         * Calls InviteAgent.
         * @function inviteAgent
         * @memberof account.AccountMgr
         * @instance
         * @param {account.IAgent} request Agent message or plain object
         * @returns {Promise<account.Agent>} Promise
         * @variation 2
         */

        /**
         * Callback as used by {@link account.AccountMgr#getAgent}.
         * @memberof account.AccountMgr
         * @typedef GetAgentCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {account.Agent} [response] Agent
         */

        /**
         * Calls GetAgent.
         * @function getAgent
         * @memberof account.AccountMgr
         * @instance
         * @param {common.IId} request Id message or plain object
         * @param {account.AccountMgr.GetAgentCallback} callback Node-style callback called with the error, if any, and Agent
         * @returns {undefined}
         * @variation 1
         */
        AccountMgr.prototype.getAgent = function getAgent(request, callback) {
            return this.rpcCall(getAgent, $root.common.Id, $root.account.Agent, request, callback);
        };

        /**
         * Calls GetAgent.
         * @function getAgent
         * @memberof account.AccountMgr
         * @instance
         * @param {common.IId} request Id message or plain object
         * @returns {Promise<account.Agent>} Promise
         * @variation 2
         */

        /**
         * Callback as used by {@link account.AccountMgr#confirmAccount}.
         * @memberof account.AccountMgr
         * @typedef ConfirmAccountCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {account.Account} [response] Account
         */

        /**
         * Calls ConfirmAccount.
         * @function confirmAccount
         * @memberof account.AccountMgr
         * @instance
         * @param {account.IToken} request Token message or plain object
         * @param {account.AccountMgr.ConfirmAccountCallback} callback Node-style callback called with the error, if any, and Account
         * @returns {undefined}
         * @variation 1
         */
        AccountMgr.prototype.confirmAccount = function confirmAccount(request, callback) {
            return this.rpcCall(confirmAccount, $root.account.Token, $root.account.Account, request, callback);
        };

        /**
         * Calls ConfirmAccount.
         * @function confirmAccount
         * @memberof account.AccountMgr
         * @instance
         * @param {account.IToken} request Token message or plain object
         * @returns {Promise<account.Account>} Promise
         * @variation 2
         */

        /**
         * Callback as used by {@link account.AccountMgr#getAccount}.
         * @memberof account.AccountMgr
         * @typedef GetAccountCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {account.Account} [response] Account
         */

        /**
         * Calls GetAccount.
         * @function getAccount
         * @memberof account.AccountMgr
         * @instance
         * @param {common.IId} request Id message or plain object
         * @param {account.AccountMgr.GetAccountCallback} callback Node-style callback called with the error, if any, and Account
         * @returns {undefined}
         * @variation 1
         */
        AccountMgr.prototype.getAccount = function getAccount(request, callback) {
            return this.rpcCall(getAccount, $root.common.Id, $root.account.Account, request, callback);
        };

        /**
         * Calls GetAccount.
         * @function getAccount
         * @memberof account.AccountMgr
         * @instance
         * @param {common.IId} request Id message or plain object
         * @returns {Promise<account.Account>} Promise
         * @variation 2
         */

        /**
         * Callback as used by {@link account.AccountMgr#updateAccount}.
         * @memberof account.AccountMgr
         * @typedef UpdateAccountCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {account.Account} [response] Account
         */

        /**
         * Calls UpdateAccount.
         * @function updateAccount
         * @memberof account.AccountMgr
         * @instance
         * @param {account.IAccount} request Account message or plain object
         * @param {account.AccountMgr.UpdateAccountCallback} callback Node-style callback called with the error, if any, and Account
         * @returns {undefined}
         * @variation 1
         */
        AccountMgr.prototype.updateAccount = function updateAccount(request, callback) {
            return this.rpcCall(updateAccount, $root.account.Account, $root.account.Account, request, callback);
        };

        /**
         * Calls UpdateAccount.
         * @function updateAccount
         * @memberof account.AccountMgr
         * @instance
         * @param {account.IAccount} request Account message or plain object
         * @returns {Promise<account.Account>} Promise
         * @variation 2
         */

        /**
         * Callback as used by {@link account.AccountMgr#createAccount}.
         * @memberof account.AccountMgr
         * @typedef CreateAccountCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {account.Account} [response] Account
         */

        /**
         * Calls CreateAccount.
         * @function createAccount
         * @memberof account.AccountMgr
         * @instance
         * @param {account.ICreateAccountRequest} request CreateAccountRequest message or plain object
         * @param {account.AccountMgr.CreateAccountCallback} callback Node-style callback called with the error, if any, and Account
         * @returns {undefined}
         * @variation 1
         */
        AccountMgr.prototype.createAccount = function createAccount(request, callback) {
            return this.rpcCall(createAccount, $root.account.CreateAccountRequest, $root.account.Account, request, callback);
        };

        /**
         * Calls CreateAccount.
         * @function createAccount
         * @memberof account.AccountMgr
         * @instance
         * @param {account.ICreateAccountRequest} request CreateAccountRequest message or plain object
         * @returns {Promise<account.Account>} Promise
         * @variation 2
         */

        /**
         * Callback as used by {@link account.AccountMgr#listAgents}.
         * @memberof account.AccountMgr
         * @typedef ListAgentsCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {account.Agents} [response] Agents
         */

        /**
         * Calls ListAgents.
         * @function listAgents
         * @memberof account.AccountMgr
         * @instance
         * @param {common.IId} request Id message or plain object
         * @param {account.AccountMgr.ListAgentsCallback} callback Node-style callback called with the error, if any, and Agents
         * @returns {undefined}
         * @variation 1
         */
        AccountMgr.prototype.listAgents = function listAgents(request, callback) {
            return this.rpcCall(listAgents, $root.common.Id, $root.account.Agents, request, callback);
        };

        /**
         * Calls ListAgents.
         * @function listAgents
         * @memberof account.AccountMgr
         * @instance
         * @param {common.IId} request Id message or plain object
         * @returns {Promise<account.Agents>} Promise
         * @variation 2
         */

        /**
         * Callback as used by {@link account.AccountMgr#checkLogin}.
         * @memberof account.AccountMgr
         * @typedef CheckLoginCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {common.Id} [response] Id
         */

        /**
         * Calls CheckLogin.
         * @function checkLogin
         * @memberof account.AccountMgr
         * @instance
         * @param {account.ILoginRequest} request LoginRequest message or plain object
         * @param {account.AccountMgr.CheckLoginCallback} callback Node-style callback called with the error, if any, and Id
         * @returns {undefined}
         * @variation 1
         */
        AccountMgr.prototype.checkLogin = function checkLogin(request, callback) {
            return this.rpcCall(checkLogin, $root.account.LoginRequest, $root.common.Id, request, callback);
        };

        /**
         * Calls CheckLogin.
         * @function checkLogin
         * @memberof account.AccountMgr
         * @instance
         * @param {account.ILoginRequest} request LoginRequest message or plain object
         * @returns {Promise<common.Id>} Promise
         * @variation 2
         */

        /**
         * Callback as used by {@link account.AccountMgr#deleteGroup}.
         * @memberof account.AccountMgr
         * @typedef DeleteGroupCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {common.Empty} [response] Empty
         */

        /**
         * Calls DeleteGroup.
         * @function deleteGroup
         * @memberof account.AccountMgr
         * @instance
         * @param {common.IId} request Id message or plain object
         * @param {account.AccountMgr.DeleteGroupCallback} callback Node-style callback called with the error, if any, and Empty
         * @returns {undefined}
         * @variation 1
         */
        AccountMgr.prototype.deleteGroup = function deleteGroup(request, callback) {
            return this.rpcCall(deleteGroup, $root.common.Id, $root.common.Empty, request, callback);
        };

        /**
         * Calls DeleteGroup.
         * @function deleteGroup
         * @memberof account.AccountMgr
         * @instance
         * @param {common.IId} request Id message or plain object
         * @returns {Promise<common.Empty>} Promise
         * @variation 2
         */

        /**
         * Callback as used by {@link account.AccountMgr#listGroups}.
         * @memberof account.AccountMgr
         * @typedef ListGroupsCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {account.AgentGroups} [response] AgentGroups
         */

        /**
         * Calls ListGroups.
         * @function listGroups
         * @memberof account.AccountMgr
         * @instance
         * @param {common.IId} request Id message or plain object
         * @param {account.AccountMgr.ListGroupsCallback} callback Node-style callback called with the error, if any, and AgentGroups
         * @returns {undefined}
         * @variation 1
         */
        AccountMgr.prototype.listGroups = function listGroups(request, callback) {
            return this.rpcCall(listGroups, $root.common.Id, $root.account.AgentGroups, request, callback);
        };

        /**
         * Calls ListGroups.
         * @function listGroups
         * @memberof account.AccountMgr
         * @instance
         * @param {common.IId} request Id message or plain object
         * @returns {Promise<account.AgentGroups>} Promise
         * @variation 2
         */

        /**
         * Callback as used by {@link account.AccountMgr#addAgentToGroup}.
         * @memberof account.AccountMgr
         * @typedef AddAgentToGroupCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {common.Empty} [response] Empty
         */

        /**
         * Calls AddAgentToGroup.
         * @function addAgentToGroup
         * @memberof account.AccountMgr
         * @instance
         * @param {account.IGroupMember} request GroupMember message or plain object
         * @param {account.AccountMgr.AddAgentToGroupCallback} callback Node-style callback called with the error, if any, and Empty
         * @returns {undefined}
         * @variation 1
         */
        AccountMgr.prototype.addAgentToGroup = function addAgentToGroup(request, callback) {
            return this.rpcCall(addAgentToGroup, $root.account.GroupMember, $root.common.Empty, request, callback);
        };

        /**
         * Calls AddAgentToGroup.
         * @function addAgentToGroup
         * @memberof account.AccountMgr
         * @instance
         * @param {account.IGroupMember} request GroupMember message or plain object
         * @returns {Promise<common.Empty>} Promise
         * @variation 2
         */

        /**
         * Callback as used by {@link account.AccountMgr#removeAgentFromGroup}.
         * @memberof account.AccountMgr
         * @typedef RemoveAgentFromGroupCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {common.Empty} [response] Empty
         */

        /**
         * Calls RemoveAgentFromGroup.
         * @function removeAgentFromGroup
         * @memberof account.AccountMgr
         * @instance
         * @param {account.IGroupMember} request GroupMember message or plain object
         * @param {account.AccountMgr.RemoveAgentFromGroupCallback} callback Node-style callback called with the error, if any, and Empty
         * @returns {undefined}
         * @variation 1
         */
        AccountMgr.prototype.removeAgentFromGroup = function removeAgentFromGroup(request, callback) {
            return this.rpcCall(removeAgentFromGroup, $root.account.GroupMember, $root.common.Empty, request, callback);
        };

        /**
         * Calls RemoveAgentFromGroup.
         * @function removeAgentFromGroup
         * @memberof account.AccountMgr
         * @instance
         * @param {account.IGroupMember} request GroupMember message or plain object
         * @returns {Promise<common.Empty>} Promise
         * @variation 2
         */

        return AccountMgr;
    })();

    return account;
})();

export const api = $root.api = (() => {

    /**
     * Namespace api.
     * @exports api
     * @namespace
     */
    const api = {};

    api.Empty = (function() {

        /**
         * Properties of an Empty.
         * @memberof api
         * @interface IEmpty
         */

        /**
         * Constructs a new Empty.
         * @memberof api
         * @classdesc Represents an Empty.
         * @implements IEmpty
         * @constructor
         * @param {api.IEmpty=} [p] Properties to set
         */
        function Empty(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        return Empty;
    })();

    return api;
})();

export const auth = $root.auth = (() => {

    /**
     * Namespace auth.
     * @exports auth
     * @namespace
     */
    const auth = {};

    auth.Credential = (function() {

        /**
         * Properties of a Credential.
         * @memberof auth
         * @interface ICredential
         * @property {string|null} [account_id] Credential account_id
         * @property {string|null} [issuer] Credential issuer
         * @property {auth.Type|null} [type] Credential type
         * @property {auth.IMethod|null} [method] Credential method
         * @property {string|null} [client_id] Credential client_id
         * @property {auth.Type|null} [client_type] Credential client_type
         * @property {string|null} [client_account_id] Credential client_account_id
         * @property {Array.<string>|null} [scopes] Credential scopes
         * @property {string|null} [avatar_url] Credential avatar_url
         * @property {string|null} [name] Credential name
         * @property {string|null} [email] Credential email
         */

        /**
         * Constructs a new Credential.
         * @memberof auth
         * @classdesc Represents a Credential.
         * @implements ICredential
         * @constructor
         * @param {auth.ICredential=} [p] Properties to set
         */
        function Credential(p) {
            this.scopes = [];
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * Credential account_id.
         * @member {string} account_id
         * @memberof auth.Credential
         * @instance
         */
        Credential.prototype.account_id = "";

        /**
         * Credential issuer.
         * @member {string} issuer
         * @memberof auth.Credential
         * @instance
         */
        Credential.prototype.issuer = "";

        /**
         * Credential type.
         * @member {auth.Type} type
         * @memberof auth.Credential
         * @instance
         */
        Credential.prototype.type = 0;

        /**
         * Credential method.
         * @member {auth.IMethod|null|undefined} method
         * @memberof auth.Credential
         * @instance
         */
        Credential.prototype.method = null;

        /**
         * Credential client_id.
         * @member {string} client_id
         * @memberof auth.Credential
         * @instance
         */
        Credential.prototype.client_id = "";

        /**
         * Credential client_type.
         * @member {auth.Type} client_type
         * @memberof auth.Credential
         * @instance
         */
        Credential.prototype.client_type = 0;

        /**
         * Credential client_account_id.
         * @member {string} client_account_id
         * @memberof auth.Credential
         * @instance
         */
        Credential.prototype.client_account_id = "";

        /**
         * Credential scopes.
         * @member {Array.<string>} scopes
         * @memberof auth.Credential
         * @instance
         */
        Credential.prototype.scopes = $util.emptyArray;

        /**
         * Credential avatar_url.
         * @member {string} avatar_url
         * @memberof auth.Credential
         * @instance
         */
        Credential.prototype.avatar_url = "";

        /**
         * Credential name.
         * @member {string} name
         * @memberof auth.Credential
         * @instance
         */
        Credential.prototype.name = "";

        /**
         * Credential email.
         * @member {string} email
         * @memberof auth.Credential
         * @instance
         */
        Credential.prototype.email = "";

        return Credential;
    })();

    auth.Scope = (function() {

        /**
         * Properties of a Scope.
         * @memberof auth
         * @interface IScope
         * @property {string|null} [id] Scope id
         * @property {string|null} [name] Scope name
         * @property {string|null} [logo_url] Scope logo_url
         * @property {string|null} [title] Scope title
         * @property {string|null} [description] Scope description
         * @property {auth.IMethod|null} [method] Scope method
         * @property {Array.<string>|null} [event] Scope event
         */

        /**
         * Constructs a new Scope.
         * @memberof auth
         * @classdesc Represents a Scope.
         * @implements IScope
         * @constructor
         * @param {auth.IScope=} [p] Properties to set
         */
        function Scope(p) {
            this.event = [];
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * Scope id.
         * @member {string} id
         * @memberof auth.Scope
         * @instance
         */
        Scope.prototype.id = "";

        /**
         * Scope name.
         * @member {string} name
         * @memberof auth.Scope
         * @instance
         */
        Scope.prototype.name = "";

        /**
         * Scope logo_url.
         * @member {string} logo_url
         * @memberof auth.Scope
         * @instance
         */
        Scope.prototype.logo_url = "";

        /**
         * Scope title.
         * @member {string} title
         * @memberof auth.Scope
         * @instance
         */
        Scope.prototype.title = "";

        /**
         * Scope description.
         * @member {string} description
         * @memberof auth.Scope
         * @instance
         */
        Scope.prototype.description = "";

        /**
         * Scope method.
         * @member {auth.IMethod|null|undefined} method
         * @memberof auth.Scope
         * @instance
         */
        Scope.prototype.method = null;

        /**
         * Scope event.
         * @member {Array.<string>} event
         * @memberof auth.Scope
         * @instance
         */
        Scope.prototype.event = $util.emptyArray;

        return Scope;
    })();

    auth.Method = (function() {

        /**
         * Properties of a Method.
         * @memberof auth
         * @interface IMethod
         * @property {boolean|null} [ping] Method ping
         * @property {boolean|null} [update_trigger] Method update_trigger
         * @property {boolean|null} [delete_trigger] Method delete_trigger
         * @property {boolean|null} [create_trigger] Method create_trigger
         * @property {boolean|null} [read_trigger] Method read_trigger
         * @property {boolean|null} [read_segmentation] Method read_segmentation
         * @property {boolean|null} [update_segmentation] Method update_segmentation
         * @property {boolean|null} [delete_segmentation] Method delete_segmentation
         * @property {boolean|null} [create_segmentation] Method create_segmentation
         * @property {boolean|null} [invite_agent] Method invite_agent
         * @property {boolean|null} [update_agent] Method update_agent
         * @property {boolean|null} [update_agents] Method update_agents
         * @property {boolean|null} [read_agent] Method read_agent
         * @property {boolean|null} [read_agents] Method read_agents
         * @property {boolean|null} [reset_password] Method reset_password
         * @property {boolean|null} [update_agents_permission] Method update_agents_permission
         * @property {boolean|null} [read_agent_permission] Method read_agent_permission
         * @property {boolean|null} [update_agents_state] Method update_agents_state
         * @property {boolean|null} [read_account] Method read_account
         * @property {boolean|null} [create_agent_group] Method create_agent_group
         * @property {boolean|null} [delete_agent_group] Method delete_agent_group
         * @property {boolean|null} [read_agent_group] Method read_agent_group
         * @property {boolean|null} [update_agent_group] Method update_agent_group
         * @property {boolean|null} [update_plan] Method update_plan
         * @property {boolean|null} [update_account_infomation] Method update_account_infomation
         * @property {boolean|null} [read_client] Method read_client
         * @property {boolean|null} [update_client] Method update_client
         * @property {boolean|null} [delete_client] Method delete_client
         * @property {boolean|null} [create_client] Method create_client
         * @property {boolean|null} [read_rule] Method read_rule
         * @property {boolean|null} [create_rule] Method create_rule
         * @property {boolean|null} [delete_rule] Method delete_rule
         * @property {boolean|null} [update_rule] Method update_rule
         * @property {boolean|null} [start_conversation] Method start_conversation
         * @property {boolean|null} [read_conversation] Method read_conversation
         * @property {boolean|null} [export_conversations] Method export_conversations
         * @property {boolean|null} [read_teammates_conversations] Method read_teammates_conversations
         * @property {boolean|null} [send_message] Method send_message
         * @property {boolean|null} [integrate_connector] Method integrate_connector
         * @property {boolean|null} [read_user_email] Method read_user_email
         * @property {boolean|null} [read_user_facebook_id] Method read_user_facebook_id
         * @property {boolean|null} [read_user_phones] Method read_user_phones
         * @property {boolean|null} [read_user_widget_setting] Method read_user_widget_setting
         * @property {boolean|null} [read_tag] Method read_tag
         * @property {boolean|null} [update_tag] Method update_tag
         * @property {boolean|null} [delete_tag] Method delete_tag
         * @property {boolean|null} [update_widget_setting] Method update_widget_setting
         * @property {boolean|null} [create_whitelist_domain] Method create_whitelist_domain
         * @property {boolean|null} [create_whitelist_ip] Method create_whitelist_ip
         * @property {boolean|null} [create_whitelist_user] Method create_whitelist_user
         * @property {boolean|null} [delete_whitelist_domain] Method delete_whitelist_domain
         * @property {boolean|null} [delete_whitelist_ip] Method delete_whitelist_ip
         * @property {boolean|null} [delete_whitelist_user] Method delete_whitelist_user
         * @property {boolean|null} [read_whitelist_ip] Method read_whitelist_ip
         * @property {boolean|null} [read_whitelist_user] Method read_whitelist_user
         * @property {boolean|null} [read_whitelist_domain] Method read_whitelist_domain
         * @property {boolean|null} [purchase_service] Method purchase_service
         * @property {boolean|null} [update_payment_method] Method update_payment_method
         * @property {boolean|null} [add_credit] Method add_credit
         * @property {boolean|null} [update_billing_cycle] Method update_billing_cycle
         * @property {boolean|null} [read_invoices] Method read_invoices
         * @property {boolean|null} [read_subscriptions] Method read_subscriptions
         * @property {boolean|null} [read_attribute] Method read_attribute
         * @property {boolean|null} [create_attribute] Method create_attribute
         * @property {boolean|null} [update_attribute] Method update_attribute
         * @property {boolean|null} [delete_attribute] Method delete_attribute
         */

        /**
         * Constructs a new Method.
         * @memberof auth
         * @classdesc Represents a Method.
         * @implements IMethod
         * @constructor
         * @param {auth.IMethod=} [p] Properties to set
         */
        function Method(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * Method ping.
         * @member {boolean} ping
         * @memberof auth.Method
         * @instance
         */
        Method.prototype.ping = false;

        /**
         * Method update_trigger.
         * @member {boolean} update_trigger
         * @memberof auth.Method
         * @instance
         */
        Method.prototype.update_trigger = false;

        /**
         * Method delete_trigger.
         * @member {boolean} delete_trigger
         * @memberof auth.Method
         * @instance
         */
        Method.prototype.delete_trigger = false;

        /**
         * Method create_trigger.
         * @member {boolean} create_trigger
         * @memberof auth.Method
         * @instance
         */
        Method.prototype.create_trigger = false;

        /**
         * Method read_trigger.
         * @member {boolean} read_trigger
         * @memberof auth.Method
         * @instance
         */
        Method.prototype.read_trigger = false;

        /**
         * Method read_segmentation.
         * @member {boolean} read_segmentation
         * @memberof auth.Method
         * @instance
         */
        Method.prototype.read_segmentation = false;

        /**
         * Method update_segmentation.
         * @member {boolean} update_segmentation
         * @memberof auth.Method
         * @instance
         */
        Method.prototype.update_segmentation = false;

        /**
         * Method delete_segmentation.
         * @member {boolean} delete_segmentation
         * @memberof auth.Method
         * @instance
         */
        Method.prototype.delete_segmentation = false;

        /**
         * Method create_segmentation.
         * @member {boolean} create_segmentation
         * @memberof auth.Method
         * @instance
         */
        Method.prototype.create_segmentation = false;

        /**
         * Method invite_agent.
         * @member {boolean} invite_agent
         * @memberof auth.Method
         * @instance
         */
        Method.prototype.invite_agent = false;

        /**
         * Method update_agent.
         * @member {boolean} update_agent
         * @memberof auth.Method
         * @instance
         */
        Method.prototype.update_agent = false;

        /**
         * Method update_agents.
         * @member {boolean} update_agents
         * @memberof auth.Method
         * @instance
         */
        Method.prototype.update_agents = false;

        /**
         * Method read_agent.
         * @member {boolean} read_agent
         * @memberof auth.Method
         * @instance
         */
        Method.prototype.read_agent = false;

        /**
         * Method read_agents.
         * @member {boolean} read_agents
         * @memberof auth.Method
         * @instance
         */
        Method.prototype.read_agents = false;

        /**
         * Method reset_password.
         * @member {boolean} reset_password
         * @memberof auth.Method
         * @instance
         */
        Method.prototype.reset_password = false;

        /**
         * Method update_agents_permission.
         * @member {boolean} update_agents_permission
         * @memberof auth.Method
         * @instance
         */
        Method.prototype.update_agents_permission = false;

        /**
         * Method read_agent_permission.
         * @member {boolean} read_agent_permission
         * @memberof auth.Method
         * @instance
         */
        Method.prototype.read_agent_permission = false;

        /**
         * Method update_agents_state.
         * @member {boolean} update_agents_state
         * @memberof auth.Method
         * @instance
         */
        Method.prototype.update_agents_state = false;

        /**
         * Method read_account.
         * @member {boolean} read_account
         * @memberof auth.Method
         * @instance
         */
        Method.prototype.read_account = false;

        /**
         * Method create_agent_group.
         * @member {boolean} create_agent_group
         * @memberof auth.Method
         * @instance
         */
        Method.prototype.create_agent_group = false;

        /**
         * Method delete_agent_group.
         * @member {boolean} delete_agent_group
         * @memberof auth.Method
         * @instance
         */
        Method.prototype.delete_agent_group = false;

        /**
         * Method read_agent_group.
         * @member {boolean} read_agent_group
         * @memberof auth.Method
         * @instance
         */
        Method.prototype.read_agent_group = false;

        /**
         * Method update_agent_group.
         * @member {boolean} update_agent_group
         * @memberof auth.Method
         * @instance
         */
        Method.prototype.update_agent_group = false;

        /**
         * Method update_plan.
         * @member {boolean} update_plan
         * @memberof auth.Method
         * @instance
         */
        Method.prototype.update_plan = false;

        /**
         * Method update_account_infomation.
         * @member {boolean} update_account_infomation
         * @memberof auth.Method
         * @instance
         */
        Method.prototype.update_account_infomation = false;

        /**
         * Method read_client.
         * @member {boolean} read_client
         * @memberof auth.Method
         * @instance
         */
        Method.prototype.read_client = false;

        /**
         * Method update_client.
         * @member {boolean} update_client
         * @memberof auth.Method
         * @instance
         */
        Method.prototype.update_client = false;

        /**
         * Method delete_client.
         * @member {boolean} delete_client
         * @memberof auth.Method
         * @instance
         */
        Method.prototype.delete_client = false;

        /**
         * Method create_client.
         * @member {boolean} create_client
         * @memberof auth.Method
         * @instance
         */
        Method.prototype.create_client = false;

        /**
         * Method read_rule.
         * @member {boolean} read_rule
         * @memberof auth.Method
         * @instance
         */
        Method.prototype.read_rule = false;

        /**
         * Method create_rule.
         * @member {boolean} create_rule
         * @memberof auth.Method
         * @instance
         */
        Method.prototype.create_rule = false;

        /**
         * Method delete_rule.
         * @member {boolean} delete_rule
         * @memberof auth.Method
         * @instance
         */
        Method.prototype.delete_rule = false;

        /**
         * Method update_rule.
         * @member {boolean} update_rule
         * @memberof auth.Method
         * @instance
         */
        Method.prototype.update_rule = false;

        /**
         * Method start_conversation.
         * @member {boolean} start_conversation
         * @memberof auth.Method
         * @instance
         */
        Method.prototype.start_conversation = false;

        /**
         * Method read_conversation.
         * @member {boolean} read_conversation
         * @memberof auth.Method
         * @instance
         */
        Method.prototype.read_conversation = false;

        /**
         * Method export_conversations.
         * @member {boolean} export_conversations
         * @memberof auth.Method
         * @instance
         */
        Method.prototype.export_conversations = false;

        /**
         * Method read_teammates_conversations.
         * @member {boolean} read_teammates_conversations
         * @memberof auth.Method
         * @instance
         */
        Method.prototype.read_teammates_conversations = false;

        /**
         * Method send_message.
         * @member {boolean} send_message
         * @memberof auth.Method
         * @instance
         */
        Method.prototype.send_message = false;

        /**
         * Method integrate_connector.
         * @member {boolean} integrate_connector
         * @memberof auth.Method
         * @instance
         */
        Method.prototype.integrate_connector = false;

        /**
         * Method read_user_email.
         * @member {boolean} read_user_email
         * @memberof auth.Method
         * @instance
         */
        Method.prototype.read_user_email = false;

        /**
         * Method read_user_facebook_id.
         * @member {boolean} read_user_facebook_id
         * @memberof auth.Method
         * @instance
         */
        Method.prototype.read_user_facebook_id = false;

        /**
         * Method read_user_phones.
         * @member {boolean} read_user_phones
         * @memberof auth.Method
         * @instance
         */
        Method.prototype.read_user_phones = false;

        /**
         * Method read_user_widget_setting.
         * @member {boolean} read_user_widget_setting
         * @memberof auth.Method
         * @instance
         */
        Method.prototype.read_user_widget_setting = false;

        /**
         * Method read_tag.
         * @member {boolean} read_tag
         * @memberof auth.Method
         * @instance
         */
        Method.prototype.read_tag = false;

        /**
         * Method update_tag.
         * @member {boolean} update_tag
         * @memberof auth.Method
         * @instance
         */
        Method.prototype.update_tag = false;

        /**
         * Method delete_tag.
         * @member {boolean} delete_tag
         * @memberof auth.Method
         * @instance
         */
        Method.prototype.delete_tag = false;

        /**
         * Method update_widget_setting.
         * @member {boolean} update_widget_setting
         * @memberof auth.Method
         * @instance
         */
        Method.prototype.update_widget_setting = false;

        /**
         * Method create_whitelist_domain.
         * @member {boolean} create_whitelist_domain
         * @memberof auth.Method
         * @instance
         */
        Method.prototype.create_whitelist_domain = false;

        /**
         * Method create_whitelist_ip.
         * @member {boolean} create_whitelist_ip
         * @memberof auth.Method
         * @instance
         */
        Method.prototype.create_whitelist_ip = false;

        /**
         * Method create_whitelist_user.
         * @member {boolean} create_whitelist_user
         * @memberof auth.Method
         * @instance
         */
        Method.prototype.create_whitelist_user = false;

        /**
         * Method delete_whitelist_domain.
         * @member {boolean} delete_whitelist_domain
         * @memberof auth.Method
         * @instance
         */
        Method.prototype.delete_whitelist_domain = false;

        /**
         * Method delete_whitelist_ip.
         * @member {boolean} delete_whitelist_ip
         * @memberof auth.Method
         * @instance
         */
        Method.prototype.delete_whitelist_ip = false;

        /**
         * Method delete_whitelist_user.
         * @member {boolean} delete_whitelist_user
         * @memberof auth.Method
         * @instance
         */
        Method.prototype.delete_whitelist_user = false;

        /**
         * Method read_whitelist_ip.
         * @member {boolean} read_whitelist_ip
         * @memberof auth.Method
         * @instance
         */
        Method.prototype.read_whitelist_ip = false;

        /**
         * Method read_whitelist_user.
         * @member {boolean} read_whitelist_user
         * @memberof auth.Method
         * @instance
         */
        Method.prototype.read_whitelist_user = false;

        /**
         * Method read_whitelist_domain.
         * @member {boolean} read_whitelist_domain
         * @memberof auth.Method
         * @instance
         */
        Method.prototype.read_whitelist_domain = false;

        /**
         * Method purchase_service.
         * @member {boolean} purchase_service
         * @memberof auth.Method
         * @instance
         */
        Method.prototype.purchase_service = false;

        /**
         * Method update_payment_method.
         * @member {boolean} update_payment_method
         * @memberof auth.Method
         * @instance
         */
        Method.prototype.update_payment_method = false;

        /**
         * Method add_credit.
         * @member {boolean} add_credit
         * @memberof auth.Method
         * @instance
         */
        Method.prototype.add_credit = false;

        /**
         * Method update_billing_cycle.
         * @member {boolean} update_billing_cycle
         * @memberof auth.Method
         * @instance
         */
        Method.prototype.update_billing_cycle = false;

        /**
         * Method read_invoices.
         * @member {boolean} read_invoices
         * @memberof auth.Method
         * @instance
         */
        Method.prototype.read_invoices = false;

        /**
         * Method read_subscriptions.
         * @member {boolean} read_subscriptions
         * @memberof auth.Method
         * @instance
         */
        Method.prototype.read_subscriptions = false;

        /**
         * Method read_attribute.
         * @member {boolean} read_attribute
         * @memberof auth.Method
         * @instance
         */
        Method.prototype.read_attribute = false;

        /**
         * Method create_attribute.
         * @member {boolean} create_attribute
         * @memberof auth.Method
         * @instance
         */
        Method.prototype.create_attribute = false;

        /**
         * Method update_attribute.
         * @member {boolean} update_attribute
         * @memberof auth.Method
         * @instance
         */
        Method.prototype.update_attribute = false;

        /**
         * Method delete_attribute.
         * @member {boolean} delete_attribute
         * @memberof auth.Method
         * @instance
         */
        Method.prototype.delete_attribute = false;

        return Method;
    })();

    auth.UserAuth = (function() {

        /**
         * Properties of a UserAuth.
         * @memberof auth
         * @interface IUserAuth
         * @property {string|null} [user_id] UserAuth user_id
         * @property {auth.IMethod|null} [method] UserAuth method
         */

        /**
         * Constructs a new UserAuth.
         * @memberof auth
         * @classdesc Represents a UserAuth.
         * @implements IUserAuth
         * @constructor
         * @param {auth.IUserAuth=} [p] Properties to set
         */
        function UserAuth(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * UserAuth user_id.
         * @member {string} user_id
         * @memberof auth.UserAuth
         * @instance
         */
        UserAuth.prototype.user_id = "";

        /**
         * UserAuth method.
         * @member {auth.IMethod|null|undefined} method
         * @memberof auth.UserAuth
         * @instance
         */
        UserAuth.prototype.method = null;

        return UserAuth;
    })();

    auth.PasswordCredential = (function() {

        /**
         * Properties of a PasswordCredential.
         * @memberof auth
         * @interface IPasswordCredential
         * @property {string|null} [username] PasswordCredential username
         * @property {string|null} [password] PasswordCredential password
         */

        /**
         * Constructs a new PasswordCredential.
         * @memberof auth
         * @classdesc Represents a PasswordCredential.
         * @implements IPasswordCredential
         * @constructor
         * @param {auth.IPasswordCredential=} [p] Properties to set
         */
        function PasswordCredential(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * PasswordCredential username.
         * @member {string} username
         * @memberof auth.PasswordCredential
         * @instance
         */
        PasswordCredential.prototype.username = "";

        /**
         * PasswordCredential password.
         * @member {string} password
         * @memberof auth.PasswordCredential
         * @instance
         */
        PasswordCredential.prototype.password = "";

        return PasswordCredential;
    })();

    auth.SuperPasswordCredential = (function() {

        /**
         * Properties of a SuperPasswordCredential.
         * @memberof auth
         * @interface ISuperPasswordCredential
         * @property {string|null} [subiz_username] SuperPasswordCredential subiz_username
         * @property {string|null} [subiz_password] SuperPasswordCredential subiz_password
         * @property {string|null} [subiz_token] SuperPasswordCredential subiz_token
         * @property {string|null} [issuer_id] SuperPasswordCredential issuer_id
         * @property {string|null} [account_id] SuperPasswordCredential account_id
         * @property {number|null} [expired] SuperPasswordCredential expired
         * @property {Array.<string>|null} [scopes] SuperPasswordCredential scopes
         */

        /**
         * Constructs a new SuperPasswordCredential.
         * @memberof auth
         * @classdesc Represents a SuperPasswordCredential.
         * @implements ISuperPasswordCredential
         * @constructor
         * @param {auth.ISuperPasswordCredential=} [p] Properties to set
         */
        function SuperPasswordCredential(p) {
            this.scopes = [];
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * SuperPasswordCredential subiz_username.
         * @member {string} subiz_username
         * @memberof auth.SuperPasswordCredential
         * @instance
         */
        SuperPasswordCredential.prototype.subiz_username = "";

        /**
         * SuperPasswordCredential subiz_password.
         * @member {string} subiz_password
         * @memberof auth.SuperPasswordCredential
         * @instance
         */
        SuperPasswordCredential.prototype.subiz_password = "";

        /**
         * SuperPasswordCredential subiz_token.
         * @member {string} subiz_token
         * @memberof auth.SuperPasswordCredential
         * @instance
         */
        SuperPasswordCredential.prototype.subiz_token = "";

        /**
         * SuperPasswordCredential issuer_id.
         * @member {string} issuer_id
         * @memberof auth.SuperPasswordCredential
         * @instance
         */
        SuperPasswordCredential.prototype.issuer_id = "";

        /**
         * SuperPasswordCredential account_id.
         * @member {string} account_id
         * @memberof auth.SuperPasswordCredential
         * @instance
         */
        SuperPasswordCredential.prototype.account_id = "";

        /**
         * SuperPasswordCredential expired.
         * @member {number} expired
         * @memberof auth.SuperPasswordCredential
         * @instance
         */
        SuperPasswordCredential.prototype.expired = 0;

        /**
         * SuperPasswordCredential scopes.
         * @member {Array.<string>} scopes
         * @memberof auth.SuperPasswordCredential
         * @instance
         */
        SuperPasswordCredential.prototype.scopes = $util.emptyArray;

        return SuperPasswordCredential;
    })();

    auth.AuthCookie = (function() {

        /**
         * Properties of an AuthCookie.
         * @memberof auth
         * @interface IAuthCookie
         * @property {string|null} [user_id] AuthCookie user_id
         * @property {string|null} [account_id] AuthCookie account_id
         * @property {number|null} [expired] AuthCookie expired
         * @property {number|null} [issued] AuthCookie issued
         * @property {string|null} [type] AuthCookie type
         */

        /**
         * Constructs a new AuthCookie.
         * @memberof auth
         * @classdesc Represents an AuthCookie.
         * @implements IAuthCookie
         * @constructor
         * @param {auth.IAuthCookie=} [p] Properties to set
         */
        function AuthCookie(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * AuthCookie user_id.
         * @member {string} user_id
         * @memberof auth.AuthCookie
         * @instance
         */
        AuthCookie.prototype.user_id = "";

        /**
         * AuthCookie account_id.
         * @member {string} account_id
         * @memberof auth.AuthCookie
         * @instance
         */
        AuthCookie.prototype.account_id = "";

        /**
         * AuthCookie expired.
         * @member {number} expired
         * @memberof auth.AuthCookie
         * @instance
         */
        AuthCookie.prototype.expired = 0;

        /**
         * AuthCookie issued.
         * @member {number} issued
         * @memberof auth.AuthCookie
         * @instance
         */
        AuthCookie.prototype.issued = 0;

        /**
         * AuthCookie type.
         * @member {string} type
         * @memberof auth.AuthCookie
         * @instance
         */
        AuthCookie.prototype.type = "";

        return AuthCookie;
    })();

    auth.OauthAccessToken = (function() {

        /**
         * Properties of an OauthAccessToken.
         * @memberof auth
         * @interface IOauthAccessToken
         * @property {string|null} [access_token] OauthAccessToken access_token
         * @property {string|null} [token_type] OauthAccessToken token_type
         * @property {number|Long|null} [expires_in] OauthAccessToken expires_in
         * @property {string|null} [refresh_token] OauthAccessToken refresh_token
         * @property {string|null} [scope] OauthAccessToken scope
         * @property {string|null} [state] OauthAccessToken state
         * @property {string|null} [error] OauthAccessToken error
         * @property {string|null} [error_description] OauthAccessToken error_description
         * @property {string|null} [error_uri] OauthAccessToken error_uri
         */

        /**
         * Constructs a new OauthAccessToken.
         * @memberof auth
         * @classdesc Represents an OauthAccessToken.
         * @implements IOauthAccessToken
         * @constructor
         * @param {auth.IOauthAccessToken=} [p] Properties to set
         */
        function OauthAccessToken(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * OauthAccessToken access_token.
         * @member {string} access_token
         * @memberof auth.OauthAccessToken
         * @instance
         */
        OauthAccessToken.prototype.access_token = "";

        /**
         * OauthAccessToken token_type.
         * @member {string} token_type
         * @memberof auth.OauthAccessToken
         * @instance
         */
        OauthAccessToken.prototype.token_type = "";

        /**
         * OauthAccessToken expires_in.
         * @member {number|Long} expires_in
         * @memberof auth.OauthAccessToken
         * @instance
         */
        OauthAccessToken.prototype.expires_in = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

        /**
         * OauthAccessToken refresh_token.
         * @member {string} refresh_token
         * @memberof auth.OauthAccessToken
         * @instance
         */
        OauthAccessToken.prototype.refresh_token = "";

        /**
         * OauthAccessToken scope.
         * @member {string} scope
         * @memberof auth.OauthAccessToken
         * @instance
         */
        OauthAccessToken.prototype.scope = "";

        /**
         * OauthAccessToken state.
         * @member {string} state
         * @memberof auth.OauthAccessToken
         * @instance
         */
        OauthAccessToken.prototype.state = "";

        /**
         * OauthAccessToken error.
         * @member {string} error
         * @memberof auth.OauthAccessToken
         * @instance
         */
        OauthAccessToken.prototype.error = "";

        /**
         * OauthAccessToken error_description.
         * @member {string} error_description
         * @memberof auth.OauthAccessToken
         * @instance
         */
        OauthAccessToken.prototype.error_description = "";

        /**
         * OauthAccessToken error_uri.
         * @member {string} error_uri
         * @memberof auth.OauthAccessToken
         * @instance
         */
        OauthAccessToken.prototype.error_uri = "";

        return OauthAccessToken;
    })();

    auth.AccessToken = (function() {

        /**
         * Properties of an AccessToken.
         * @memberof auth
         * @interface IAccessToken
         * @property {string|null} [issuer_id] AccessToken issuer_id
         * @property {string|null} [issuer_type] AccessToken issuer_type
         * @property {string|null} [client_id] AccessToken client_id
         * @property {string|null} [client_type] AccessToken client_type
         * @property {string|null} [account_id] AccessToken account_id
         * @property {string|null} [client_account_id] AccessToken client_account_id
         * @property {number|null} [expired] AccessToken expired
         * @property {Array.<string>|null} [scopes] AccessToken scopes
         */

        /**
         * Constructs a new AccessToken.
         * @memberof auth
         * @classdesc Represents an AccessToken.
         * @implements IAccessToken
         * @constructor
         * @param {auth.IAccessToken=} [p] Properties to set
         */
        function AccessToken(p) {
            this.scopes = [];
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * AccessToken issuer_id.
         * @member {string} issuer_id
         * @memberof auth.AccessToken
         * @instance
         */
        AccessToken.prototype.issuer_id = "";

        /**
         * AccessToken issuer_type.
         * @member {string} issuer_type
         * @memberof auth.AccessToken
         * @instance
         */
        AccessToken.prototype.issuer_type = "";

        /**
         * AccessToken client_id.
         * @member {string} client_id
         * @memberof auth.AccessToken
         * @instance
         */
        AccessToken.prototype.client_id = "";

        /**
         * AccessToken client_type.
         * @member {string} client_type
         * @memberof auth.AccessToken
         * @instance
         */
        AccessToken.prototype.client_type = "";

        /**
         * AccessToken account_id.
         * @member {string} account_id
         * @memberof auth.AccessToken
         * @instance
         */
        AccessToken.prototype.account_id = "";

        /**
         * AccessToken client_account_id.
         * @member {string} client_account_id
         * @memberof auth.AccessToken
         * @instance
         */
        AccessToken.prototype.client_account_id = "";

        /**
         * AccessToken expired.
         * @member {number} expired
         * @memberof auth.AccessToken
         * @instance
         */
        AccessToken.prototype.expired = 0;

        /**
         * AccessToken scopes.
         * @member {Array.<string>} scopes
         * @memberof auth.AccessToken
         * @instance
         */
        AccessToken.prototype.scopes = $util.emptyArray;

        return AccessToken;
    })();

    auth.CookieExpire = (function() {

        /**
         * Properties of a CookieExpire.
         * @memberof auth
         * @interface ICookieExpire
         * @property {string|null} [user_id] CookieExpire user_id
         * @property {string|null} [account_id] CookieExpire account_id
         * @property {string|null} [expired_token] CookieExpire expired_token
         * @property {number|Long|null} [before_time] CookieExpire before_time
         * @property {string|null} [except_token] CookieExpire except_token
         */

        /**
         * Constructs a new CookieExpire.
         * @memberof auth
         * @classdesc Represents a CookieExpire.
         * @implements ICookieExpire
         * @constructor
         * @param {auth.ICookieExpire=} [p] Properties to set
         */
        function CookieExpire(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * CookieExpire user_id.
         * @member {string} user_id
         * @memberof auth.CookieExpire
         * @instance
         */
        CookieExpire.prototype.user_id = "";

        /**
         * CookieExpire account_id.
         * @member {string} account_id
         * @memberof auth.CookieExpire
         * @instance
         */
        CookieExpire.prototype.account_id = "";

        /**
         * CookieExpire expired_token.
         * @member {string} expired_token
         * @memberof auth.CookieExpire
         * @instance
         */
        CookieExpire.prototype.expired_token = "";

        /**
         * CookieExpire before_time.
         * @member {number|Long} before_time
         * @memberof auth.CookieExpire
         * @instance
         */
        CookieExpire.prototype.before_time = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

        /**
         * CookieExpire except_token.
         * @member {string} except_token
         * @memberof auth.CookieExpire
         * @instance
         */
        CookieExpire.prototype.except_token = "";

        return CookieExpire;
    })();

    /**
     * Type enum.
     * @name auth.Type
     * @enum {string}
     * @property {number} unknown=0 unknown value
     * @property {number} user=1 user value
     * @property {number} agent=2 agent value
     * @property {number} subiz=3 subiz value
     * @property {number} app=5 app value
     * @property {number} rest=8 rest value
     * @property {number} connector=6 connector value
     * @property {number} bot=7 bot value
     * @property {number} dashboard=10 dashboard value
     */
    auth.Type = (function() {
        const valuesById = {}, values = Object.create(valuesById);
        values[valuesById[0] = "unknown"] = 0;
        values[valuesById[1] = "user"] = 1;
        values[valuesById[2] = "agent"] = 2;
        values[valuesById[3] = "subiz"] = 3;
        values[valuesById[5] = "app"] = 5;
        values[valuesById[8] = "rest"] = 8;
        values[valuesById[6] = "connector"] = 6;
        values[valuesById[7] = "bot"] = 7;
        values[valuesById[10] = "dashboard"] = 10;
        return values;
    })();

    auth.User = (function() {

        /**
         * Properties of a User.
         * @memberof auth
         * @interface IUser
         * @property {string|null} [id] User id
         * @property {string|null} [account_id] User account_id
         * @property {string|null} [email] User email
         * @property {string|null} [encrypted_password] User encrypted_password
         * @property {boolean|null} [is_active] User is_active
         * @property {number|Long|null} [upserted] User upserted
         */

        /**
         * Constructs a new User.
         * @memberof auth
         * @classdesc Represents a User.
         * @implements IUser
         * @constructor
         * @param {auth.IUser=} [p] Properties to set
         */
        function User(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * User id.
         * @member {string} id
         * @memberof auth.User
         * @instance
         */
        User.prototype.id = "";

        /**
         * User account_id.
         * @member {string} account_id
         * @memberof auth.User
         * @instance
         */
        User.prototype.account_id = "";

        /**
         * User email.
         * @member {string} email
         * @memberof auth.User
         * @instance
         */
        User.prototype.email = "";

        /**
         * User encrypted_password.
         * @member {string} encrypted_password
         * @memberof auth.User
         * @instance
         */
        User.prototype.encrypted_password = "";

        /**
         * User is_active.
         * @member {boolean} is_active
         * @memberof auth.User
         * @instance
         */
        User.prototype.is_active = false;

        /**
         * User upserted.
         * @member {number|Long} upserted
         * @memberof auth.User
         * @instance
         */
        User.prototype.upserted = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

        return User;
    })();

    /**
     * AuthorizationType enum.
     * @name auth.AuthorizationType
     * @enum {string}
     * @property {number} client_auth=0 client_auth value
     * @property {number} channel_auth=1 channel_auth value
     */
    auth.AuthorizationType = (function() {
        const valuesById = {}, values = Object.create(valuesById);
        values[valuesById[0] = "client_auth"] = 0;
        values[valuesById[1] = "channel_auth"] = 1;
        return values;
    })();

    /**
     * Event enum.
     * @name auth.Event
     * @enum {string}
     * @property {number} AUTH=0 AUTH value
     * @property {number} AuthExpireCookie=4 AuthExpireCookie value
     */
    auth.Event = (function() {
        const valuesById = {}, values = Object.create(valuesById);
        values[valuesById[0] = "AUTH"] = 0;
        values[valuesById[4] = "AuthExpireCookie"] = 4;
        return values;
    })();

    return auth;
})();

export const client = $root.client = (() => {

    /**
     * Namespace client.
     * @exports client
     * @namespace
     */
    const client = {};

    client.AllType = (function() {

        /**
         * Properties of an AllType.
         * @memberof client
         * @interface IAllType
         * @property {client.IClients|null} [client] AllType client
         */

        /**
         * Constructs a new AllType.
         * @memberof client
         * @classdesc Represents an AllType.
         * @implements IAllType
         * @constructor
         * @param {client.IAllType=} [p] Properties to set
         */
        function AllType(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * AllType client.
         * @member {client.IClients|null|undefined} client
         * @memberof client.AllType
         * @instance
         */
        AllType.prototype.client = null;

        return AllType;
    })();

    client.Clients = (function() {

        /**
         * Properties of a Clients.
         * @memberof client
         * @interface IClients
         * @property {common.IContext|null} [ctx] Clients ctx
         * @property {Array.<client.IClient>|null} [clients] Clients clients
         */

        /**
         * Constructs a new Clients.
         * @memberof client
         * @classdesc Represents a Clients.
         * @implements IClients
         * @constructor
         * @param {client.IClients=} [p] Properties to set
         */
        function Clients(p) {
            this.clients = [];
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * Clients ctx.
         * @member {common.IContext|null|undefined} ctx
         * @memberof client.Clients
         * @instance
         */
        Clients.prototype.ctx = null;

        /**
         * Clients clients.
         * @member {Array.<client.IClient>} clients
         * @memberof client.Clients
         * @instance
         */
        Clients.prototype.clients = $util.emptyArray;

        return Clients;
    })();

    client.Client = (function() {

        /**
         * Properties of a Client.
         * @memberof client
         * @interface IClient
         * @property {common.IContext|null} [ctx] Client ctx
         * @property {string|null} [id] Client id
         * @property {string|null} [secret] Client secret
         * @property {string|null} [logo_url] Client logo_url
         * @property {string|null} [account_id] Client account_id
         * @property {boolean|null} [is_verified] Client is_verified
         * @property {number|Long|null} [verified] Client verified
         * @property {string|null} [redirect_url] Client redirect_url
         * @property {string|null} [type] Client type
         * @property {string|null} [name] Client name
         * @property {string|null} [version] Client version
         * @property {boolean|null} [is_enabled] Client is_enabled
         * @property {number|Long|null} [created] Client created
         * @property {number|Long|null} [modified] Client modified
         * @property {string|null} [webhook_url] Client webhook_url
         * @property {Array.<string>|null} [events] Client events
         * @property {string|null} [channel_type] Client channel_type
         * @property {string|null} [availability_url] Client availability_url
         * @property {string|null} [ping_url] Client ping_url
         */

        /**
         * Constructs a new Client.
         * @memberof client
         * @classdesc Represents a Client.
         * @implements IClient
         * @constructor
         * @param {client.IClient=} [p] Properties to set
         */
        function Client(p) {
            this.events = [];
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * Client ctx.
         * @member {common.IContext|null|undefined} ctx
         * @memberof client.Client
         * @instance
         */
        Client.prototype.ctx = null;

        /**
         * Client id.
         * @member {string} id
         * @memberof client.Client
         * @instance
         */
        Client.prototype.id = "";

        /**
         * Client secret.
         * @member {string} secret
         * @memberof client.Client
         * @instance
         */
        Client.prototype.secret = "";

        /**
         * Client logo_url.
         * @member {string} logo_url
         * @memberof client.Client
         * @instance
         */
        Client.prototype.logo_url = "";

        /**
         * Client account_id.
         * @member {string} account_id
         * @memberof client.Client
         * @instance
         */
        Client.prototype.account_id = "";

        /**
         * Client is_verified.
         * @member {boolean} is_verified
         * @memberof client.Client
         * @instance
         */
        Client.prototype.is_verified = false;

        /**
         * Client verified.
         * @member {number|Long} verified
         * @memberof client.Client
         * @instance
         */
        Client.prototype.verified = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

        /**
         * Client redirect_url.
         * @member {string} redirect_url
         * @memberof client.Client
         * @instance
         */
        Client.prototype.redirect_url = "";

        /**
         * Client type.
         * @member {string} type
         * @memberof client.Client
         * @instance
         */
        Client.prototype.type = "";

        /**
         * Client name.
         * @member {string} name
         * @memberof client.Client
         * @instance
         */
        Client.prototype.name = "";

        /**
         * Client version.
         * @member {string} version
         * @memberof client.Client
         * @instance
         */
        Client.prototype.version = "";

        /**
         * Client is_enabled.
         * @member {boolean} is_enabled
         * @memberof client.Client
         * @instance
         */
        Client.prototype.is_enabled = false;

        /**
         * Client created.
         * @member {number|Long} created
         * @memberof client.Client
         * @instance
         */
        Client.prototype.created = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

        /**
         * Client modified.
         * @member {number|Long} modified
         * @memberof client.Client
         * @instance
         */
        Client.prototype.modified = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

        /**
         * Client webhook_url.
         * @member {string} webhook_url
         * @memberof client.Client
         * @instance
         */
        Client.prototype.webhook_url = "";

        /**
         * Client events.
         * @member {Array.<string>} events
         * @memberof client.Client
         * @instance
         */
        Client.prototype.events = $util.emptyArray;

        /**
         * Client channel_type.
         * @member {string} channel_type
         * @memberof client.Client
         * @instance
         */
        Client.prototype.channel_type = "";

        /**
         * Client availability_url.
         * @member {string} availability_url
         * @memberof client.Client
         * @instance
         */
        Client.prototype.availability_url = "";

        /**
         * Client ping_url.
         * @member {string} ping_url
         * @memberof client.Client
         * @instance
         */
        Client.prototype.ping_url = "";

        /**
         * Type enum.
         * @name client.Client.Type
         * @enum {string}
         * @property {number} app=0 app value
         * @property {number} connector=1 connector value
         * @property {number} bot=3 bot value
         */
        Client.Type = (function() {
            const valuesById = {}, values = Object.create(valuesById);
            values[valuesById[0] = "app"] = 0;
            values[valuesById[1] = "connector"] = 1;
            values[valuesById[3] = "bot"] = 3;
            return values;
        })();

        /**
         * ChannelType enum.
         * @name client.Client.ChannelType
         * @enum {string}
         * @property {number} subiz=0 subiz value
         * @property {number} email=1 email value
         * @property {number} facebook=2 facebook value
         * @property {number} viber=3 viber value
         */
        Client.ChannelType = (function() {
            const valuesById = {}, values = Object.create(valuesById);
            values[valuesById[0] = "subiz"] = 0;
            values[valuesById[1] = "email"] = 1;
            values[valuesById[2] = "facebook"] = 2;
            values[valuesById[3] = "viber"] = 3;
            return values;
        })();

        return Client;
    })();

    /**
     * Event enum.
     * @name client.Event
     * @enum {string}
     * @property {number} ClientCreateRequested=20 ClientCreateRequested value
     * @property {number} ClientUpdateRequested=21 ClientUpdateRequested value
     * @property {number} ClientDeleteRequested=22 ClientDeleteRequested value
     * @property {number} ClientReadRequested=23 ClientReadRequested value
     * @property {number} ClientListRequested=24 ClientListRequested value
     * @property {number} ClientUpserted=25 ClientUpserted value
     * @property {number} ClientDeleted=26 ClientDeleted value
     * @property {number} ClientAuthorized=12 ClientAuthorized value
     * @property {number} ClientRequested=30 ClientRequested value
     * @property {number} ClientSynced=31 ClientSynced value
     * @property {number} ConnectorAuthorized=13 ConnectorAuthorized value
     */
    client.Event = (function() {
        const valuesById = {}, values = Object.create(valuesById);
        values[valuesById[20] = "ClientCreateRequested"] = 20;
        values[valuesById[21] = "ClientUpdateRequested"] = 21;
        values[valuesById[22] = "ClientDeleteRequested"] = 22;
        values[valuesById[23] = "ClientReadRequested"] = 23;
        values[valuesById[24] = "ClientListRequested"] = 24;
        values[valuesById[25] = "ClientUpserted"] = 25;
        values[valuesById[26] = "ClientDeleted"] = 26;
        values[valuesById[12] = "ClientAuthorized"] = 12;
        values[valuesById[30] = "ClientRequested"] = 30;
        values[valuesById[31] = "ClientSynced"] = 31;
        values[valuesById[13] = "ConnectorAuthorized"] = 13;
        return values;
    })();

    client.Authorization = (function() {

        /**
         * Properties of an Authorization.
         * @memberof client
         * @interface IAuthorization
         * @property {common.IContext|null} [ctx] Authorization ctx
         * @property {string|null} [account_id] Authorization account_id
         * @property {string|null} [issuer] Authorization issuer
         * @property {auth.Type|null} [type] Authorization type
         * @property {auth.IMethod|null} [method] Authorization method
         * @property {string|null} [client_id] Authorization client_id
         * @property {auth.Type|null} [client_type] Authorization client_type
         * @property {string|null} [client_account_id] Authorization client_account_id
         * @property {Array.<string>|null} [scopes] Authorization scopes
         * @property {string|null} [integration_id] Authorization integration_id
         */

        /**
         * Constructs a new Authorization.
         * @memberof client
         * @classdesc Represents an Authorization.
         * @implements IAuthorization
         * @constructor
         * @param {client.IAuthorization=} [p] Properties to set
         */
        function Authorization(p) {
            this.scopes = [];
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * Authorization ctx.
         * @member {common.IContext|null|undefined} ctx
         * @memberof client.Authorization
         * @instance
         */
        Authorization.prototype.ctx = null;

        /**
         * Authorization account_id.
         * @member {string} account_id
         * @memberof client.Authorization
         * @instance
         */
        Authorization.prototype.account_id = "";

        /**
         * Authorization issuer.
         * @member {string} issuer
         * @memberof client.Authorization
         * @instance
         */
        Authorization.prototype.issuer = "";

        /**
         * Authorization type.
         * @member {auth.Type} type
         * @memberof client.Authorization
         * @instance
         */
        Authorization.prototype.type = 0;

        /**
         * Authorization method.
         * @member {auth.IMethod|null|undefined} method
         * @memberof client.Authorization
         * @instance
         */
        Authorization.prototype.method = null;

        /**
         * Authorization client_id.
         * @member {string} client_id
         * @memberof client.Authorization
         * @instance
         */
        Authorization.prototype.client_id = "";

        /**
         * Authorization client_type.
         * @member {auth.Type} client_type
         * @memberof client.Authorization
         * @instance
         */
        Authorization.prototype.client_type = 0;

        /**
         * Authorization client_account_id.
         * @member {string} client_account_id
         * @memberof client.Authorization
         * @instance
         */
        Authorization.prototype.client_account_id = "";

        /**
         * Authorization scopes.
         * @member {Array.<string>} scopes
         * @memberof client.Authorization
         * @instance
         */
        Authorization.prototype.scopes = $util.emptyArray;

        /**
         * Authorization integration_id.
         * @member {string} integration_id
         * @memberof client.Authorization
         * @instance
         */
        Authorization.prototype.integration_id = "";

        return Authorization;
    })();

    return client;
})();

export const common = $root.common = (() => {

    /**
     * Namespace common.
     * @exports common
     * @namespace
     */
    const common = {};

    common.Empty = (function() {

        /**
         * Properties of an Empty.
         * @memberof common
         * @interface IEmpty
         * @property {common.IContext|null} [ctx] Empty ctx
         */

        /**
         * Constructs a new Empty.
         * @memberof common
         * @classdesc Represents an Empty.
         * @implements IEmpty
         * @constructor
         * @param {common.IEmpty=} [p] Properties to set
         */
        function Empty(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * Empty ctx.
         * @member {common.IContext|null|undefined} ctx
         * @memberof common.Empty
         * @instance
         */
        Empty.prototype.ctx = null;

        return Empty;
    })();

    common.Id = (function() {

        /**
         * Properties of an Id.
         * @memberof common
         * @interface IId
         * @property {common.IContext|null} [ctx] Id ctx
         * @property {string|null} [account_id] Id account_id
         * @property {string|null} [id] Id id
         */

        /**
         * Constructs a new Id.
         * @memberof common
         * @classdesc Represents an Id.
         * @implements IId
         * @constructor
         * @param {common.IId=} [p] Properties to set
         */
        function Id(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * Id ctx.
         * @member {common.IContext|null|undefined} ctx
         * @memberof common.Id
         * @instance
         */
        Id.prototype.ctx = null;

        /**
         * Id account_id.
         * @member {string} account_id
         * @memberof common.Id
         * @instance
         */
        Id.prototype.account_id = "";

        /**
         * Id id.
         * @member {string} id
         * @memberof common.Id
         * @instance
         */
        Id.prototype.id = "";

        return Id;
    })();

    common.Ids = (function() {

        /**
         * Properties of an Ids.
         * @memberof common
         * @interface IIds
         * @property {common.IContext|null} [ctx] Ids ctx
         * @property {string|null} [account_id] Ids account_id
         * @property {Array.<string>|null} [ids] Ids ids
         */

        /**
         * Constructs a new Ids.
         * @memberof common
         * @classdesc Represents an Ids.
         * @implements IIds
         * @constructor
         * @param {common.IIds=} [p] Properties to set
         */
        function Ids(p) {
            this.ids = [];
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * Ids ctx.
         * @member {common.IContext|null|undefined} ctx
         * @memberof common.Ids
         * @instance
         */
        Ids.prototype.ctx = null;

        /**
         * Ids account_id.
         * @member {string} account_id
         * @memberof common.Ids
         * @instance
         */
        Ids.prototype.account_id = "";

        /**
         * Ids ids.
         * @member {Array.<string>} ids
         * @memberof common.Ids
         * @instance
         */
        Ids.prototype.ids = $util.emptyArray;

        return Ids;
    })();

    common.Context = (function() {

        /**
         * Properties of a Context.
         * @memberof common
         * @interface IContext
         * @property {string|null} [event_id] Context event_id
         * @property {Uint8Array|null} [state] Context state
         * @property {string|null} [node] Context node
         * @property {string|null} [reply_topic] Context reply_topic
         * @property {auth.ICredential|null} [credential] Context credential
         * @property {Uint8Array|null} [tracing] Context tracing
         * @property {string|null} [reply_key] Context reply_key
         * @property {common.IDevice|null} [by_device] Context by_device
         * @property {string|null} [topic] Context topic
         * @property {number|null} [partition] Context partition
         * @property {number|Long|null} [offset] Context offset
         * @property {number|Long|null} [term] Context term
         * @property {string|null} [router_topic] Context router_topic
         */

        /**
         * Constructs a new Context.
         * @memberof common
         * @classdesc Represents a Context.
         * @implements IContext
         * @constructor
         * @param {common.IContext=} [p] Properties to set
         */
        function Context(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * Context event_id.
         * @member {string} event_id
         * @memberof common.Context
         * @instance
         */
        Context.prototype.event_id = "";

        /**
         * Context state.
         * @member {Uint8Array} state
         * @memberof common.Context
         * @instance
         */
        Context.prototype.state = $util.newBuffer([]);

        /**
         * Context node.
         * @member {string} node
         * @memberof common.Context
         * @instance
         */
        Context.prototype.node = "";

        /**
         * Context reply_topic.
         * @member {string} reply_topic
         * @memberof common.Context
         * @instance
         */
        Context.prototype.reply_topic = "";

        /**
         * Context credential.
         * @member {auth.ICredential|null|undefined} credential
         * @memberof common.Context
         * @instance
         */
        Context.prototype.credential = null;

        /**
         * Context tracing.
         * @member {Uint8Array} tracing
         * @memberof common.Context
         * @instance
         */
        Context.prototype.tracing = $util.newBuffer([]);

        /**
         * Context reply_key.
         * @member {string} reply_key
         * @memberof common.Context
         * @instance
         */
        Context.prototype.reply_key = "";

        /**
         * Context by_device.
         * @member {common.IDevice|null|undefined} by_device
         * @memberof common.Context
         * @instance
         */
        Context.prototype.by_device = null;

        /**
         * Context topic.
         * @member {string} topic
         * @memberof common.Context
         * @instance
         */
        Context.prototype.topic = "";

        /**
         * Context partition.
         * @member {number} partition
         * @memberof common.Context
         * @instance
         */
        Context.prototype.partition = 0;

        /**
         * Context offset.
         * @member {number|Long} offset
         * @memberof common.Context
         * @instance
         */
        Context.prototype.offset = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

        /**
         * Context term.
         * @member {number|Long} term
         * @memberof common.Context
         * @instance
         */
        Context.prototype.term = $util.Long ? $util.Long.fromBits(0,0,true) : 0;

        /**
         * Context router_topic.
         * @member {string} router_topic
         * @memberof common.Context
         * @instance
         */
        Context.prototype.router_topic = "";

        return Context;
    })();

    common.Reply = (function() {

        /**
         * Properties of a Reply.
         * @memberof common
         * @interface IReply
         * @property {common.IContext|null} [ctx] Reply ctx
         * @property {string|null} [event_id] Reply event_id
         * @property {Uint8Array|null} [state] Reply state
         * @property {string|null} [service] Reply service
         * @property {string|null} [service_id] Reply service_id
         * @property {boolean|null} [err] Reply err
         * @property {string|null} [err_description] Reply err_description
         * @property {lang.T|null} [err_code] Reply err_code
         * @property {number|null} [err_class] Reply err_class
         * @property {string|null} [err_hash] Reply err_hash
         * @property {Uint8Array|null} [payload] Reply payload
         * @property {number|Long|null} [published] Reply published
         */

        /**
         * Constructs a new Reply.
         * @memberof common
         * @classdesc Represents a Reply.
         * @implements IReply
         * @constructor
         * @param {common.IReply=} [p] Properties to set
         */
        function Reply(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * Reply ctx.
         * @member {common.IContext|null|undefined} ctx
         * @memberof common.Reply
         * @instance
         */
        Reply.prototype.ctx = null;

        /**
         * Reply event_id.
         * @member {string} event_id
         * @memberof common.Reply
         * @instance
         */
        Reply.prototype.event_id = "";

        /**
         * Reply state.
         * @member {Uint8Array} state
         * @memberof common.Reply
         * @instance
         */
        Reply.prototype.state = $util.newBuffer([]);

        /**
         * Reply service.
         * @member {string} service
         * @memberof common.Reply
         * @instance
         */
        Reply.prototype.service = "";

        /**
         * Reply service_id.
         * @member {string} service_id
         * @memberof common.Reply
         * @instance
         */
        Reply.prototype.service_id = "";

        /**
         * Reply err.
         * @member {boolean} err
         * @memberof common.Reply
         * @instance
         */
        Reply.prototype.err = false;

        /**
         * Reply err_description.
         * @member {string} err_description
         * @memberof common.Reply
         * @instance
         */
        Reply.prototype.err_description = "";

        /**
         * Reply err_code.
         * @member {lang.T} err_code
         * @memberof common.Reply
         * @instance
         */
        Reply.prototype.err_code = 0;

        /**
         * Reply err_class.
         * @member {number} err_class
         * @memberof common.Reply
         * @instance
         */
        Reply.prototype.err_class = 0;

        /**
         * Reply err_hash.
         * @member {string} err_hash
         * @memberof common.Reply
         * @instance
         */
        Reply.prototype.err_hash = "";

        /**
         * Reply payload.
         * @member {Uint8Array} payload
         * @memberof common.Reply
         * @instance
         */
        Reply.prototype.payload = $util.newBuffer([]);

        /**
         * Reply published.
         * @member {number|Long} published
         * @memberof common.Reply
         * @instance
         */
        Reply.prototype.published = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

        return Reply;
    })();

    common.Error = (function() {

        /**
         * Properties of an Error.
         * @memberof common
         * @interface IError
         * @property {string|null} [error] Error error
         * @property {string|null} [request_id] Error request_id
         * @property {string|null} [description] Error description
         * @property {string|null} [hash] Error hash
         * @property {string|null} [debug] Error debug
         */

        /**
         * Constructs a new Error.
         * @memberof common
         * @classdesc Represents an Error.
         * @implements IError
         * @constructor
         * @param {common.IError=} [p] Properties to set
         */
        function Error(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * Error error.
         * @member {string} error
         * @memberof common.Error
         * @instance
         */
        Error.prototype.error = "";

        /**
         * Error request_id.
         * @member {string} request_id
         * @memberof common.Error
         * @instance
         */
        Error.prototype.request_id = "";

        /**
         * Error description.
         * @member {string} description
         * @memberof common.Error
         * @instance
         */
        Error.prototype.description = "";

        /**
         * Error hash.
         * @member {string} hash
         * @memberof common.Error
         * @instance
         */
        Error.prototype.hash = "";

        /**
         * Error debug.
         * @member {string} debug
         * @memberof common.Error
         * @instance
         */
        Error.prototype.debug = "";

        return Error;
    })();

    common.Device = (function() {

        /**
         * Properties of a Device.
         * @memberof common
         * @interface IDevice
         * @property {string|null} [ip] Device ip
         * @property {string|null} [user_agent] Device user_agent
         * @property {string|null} [screen_resolution] Device screen_resolution
         * @property {string|null} [timezone] Device timezone
         * @property {string|null} [language] Device language
         * @property {string|null} [referrer] Device referrer
         * @property {string|null} [type] Device type
         */

        /**
         * Constructs a new Device.
         * @memberof common
         * @classdesc Represents a Device.
         * @implements IDevice
         * @constructor
         * @param {common.IDevice=} [p] Properties to set
         */
        function Device(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * Device ip.
         * @member {string} ip
         * @memberof common.Device
         * @instance
         */
        Device.prototype.ip = "";

        /**
         * Device user_agent.
         * @member {string} user_agent
         * @memberof common.Device
         * @instance
         */
        Device.prototype.user_agent = "";

        /**
         * Device screen_resolution.
         * @member {string} screen_resolution
         * @memberof common.Device
         * @instance
         */
        Device.prototype.screen_resolution = "";

        /**
         * Device timezone.
         * @member {string} timezone
         * @memberof common.Device
         * @instance
         */
        Device.prototype.timezone = "";

        /**
         * Device language.
         * @member {string} language
         * @memberof common.Device
         * @instance
         */
        Device.prototype.language = "";

        /**
         * Device referrer.
         * @member {string} referrer
         * @memberof common.Device
         * @instance
         */
        Device.prototype.referrer = "";

        /**
         * Device type.
         * @member {string} type
         * @memberof common.Device
         * @instance
         */
        Device.prototype.type = "";

        /**
         * DeviceType enum.
         * @name common.Device.DeviceType
         * @enum {string}
         * @property {number} unknown=0 unknown value
         * @property {number} mobile=1 mobile value
         * @property {number} tablet=2 tablet value
         * @property {number} desktop=3 desktop value
         */
        Device.DeviceType = (function() {
            const valuesById = {}, values = Object.create(valuesById);
            values[valuesById[0] = "unknown"] = 0;
            values[valuesById[1] = "mobile"] = 1;
            values[valuesById[2] = "tablet"] = 2;
            values[valuesById[3] = "desktop"] = 3;
            return values;
        })();

        return Device;
    })();

    /**
     * Event enum.
     * @name common.Event
     * @enum {string}
     * @property {number} Send_=0 Send_ value
     * @property {number} ApiReply=2 ApiReply value
     */
    common.Event = (function() {
        const valuesById = {}, values = Object.create(valuesById);
        values[valuesById[0] = "Send_"] = 0;
        values[valuesById[2] = "ApiReply"] = 2;
        return values;
    })();

    common.PingRequest = (function() {

        /**
         * Properties of a PingRequest.
         * @memberof common
         * @interface IPingRequest
         * @property {string|null} [message] PingRequest message
         */

        /**
         * Constructs a new PingRequest.
         * @memberof common
         * @classdesc Represents a PingRequest.
         * @implements IPingRequest
         * @constructor
         * @param {common.IPingRequest=} [p] Properties to set
         */
        function PingRequest(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * PingRequest message.
         * @member {string} message
         * @memberof common.PingRequest
         * @instance
         */
        PingRequest.prototype.message = "";

        return PingRequest;
    })();

    common.Pong = (function() {

        /**
         * Properties of a Pong.
         * @memberof common
         * @interface IPong
         * @property {string|null} [message] Pong message
         */

        /**
         * Constructs a new Pong.
         * @memberof common
         * @classdesc Represents a Pong.
         * @implements IPong
         * @constructor
         * @param {common.IPong=} [p] Properties to set
         */
        function Pong(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * Pong message.
         * @member {string} message
         * @memberof common.Pong
         * @instance
         */
        Pong.prototype.message = "";

        return Pong;
    })();

    common.Agent = (function() {

        /**
         * Constructs a new Agent service.
         * @memberof common
         * @classdesc Represents an Agent
         * @extends $protobuf.rpc.Service
         * @constructor
         * @param {$protobuf.RPCImpl} rpcImpl RPC implementation
         * @param {boolean} [requestDelimited=false] Whether requests are length-delimited
         * @param {boolean} [responseDelimited=false] Whether responses are length-delimited
         */
        function Agent(rpcImpl, requestDelimited, responseDelimited) {
            $protobuf.rpc.Service.call(this, rpcImpl, requestDelimited, responseDelimited);
        }

        (Agent.prototype = Object.create($protobuf.rpc.Service.prototype)).constructor = Agent;

        /**
         * Callback as used by {@link common.Agent#ping}.
         * @memberof common.Agent
         * @typedef PingCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {common.Pong} [response] Pong
         */

        /**
         * Calls Ping.
         * @function ping
         * @memberof common.Agent
         * @instance
         * @param {common.IPingRequest} request PingRequest message or plain object
         * @param {common.Agent.PingCallback} callback Node-style callback called with the error, if any, and Pong
         * @returns {undefined}
         * @variation 1
         */
        Agent.prototype.ping = function ping(request, callback) {
            return this.rpcCall(ping, $root.common.PingRequest, $root.common.Pong, request, callback);
        };

        /**
         * Calls Ping.
         * @function ping
         * @memberof common.Agent
         * @instance
         * @param {common.IPingRequest} request PingRequest message or plain object
         * @returns {Promise<common.Pong>} Promise
         * @variation 2
         */

        return Agent;
    })();

    return common;
})();

export const content = $root.content = (() => {

    /**
     * Namespace content.
     * @exports content
     * @namespace
     */
    const content = {};

    content.AllType = (function() {

        /**
         * Properties of an AllType.
         * @memberof content
         * @interface IAllType
         * @property {content.IContent|null} [content] AllType content
         * @property {content.IContents|null} [contents] AllType contents
         * @property {content.IListRequest|null} [lr] AllType lr
         */

        /**
         * Constructs a new AllType.
         * @memberof content
         * @classdesc Represents an AllType.
         * @implements IAllType
         * @constructor
         * @param {content.IAllType=} [p] Properties to set
         */
        function AllType(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * AllType content.
         * @member {content.IContent|null|undefined} content
         * @memberof content.AllType
         * @instance
         */
        AllType.prototype.content = null;

        /**
         * AllType contents.
         * @member {content.IContents|null|undefined} contents
         * @memberof content.AllType
         * @instance
         */
        AllType.prototype.contents = null;

        /**
         * AllType lr.
         * @member {content.IListRequest|null|undefined} lr
         * @memberof content.AllType
         * @instance
         */
        AllType.prototype.lr = null;

        return AllType;
    })();

    content.MyServer = (function() {

        /**
         * Constructs a new MyServer service.
         * @memberof content
         * @classdesc Represents a MyServer
         * @extends $protobuf.rpc.Service
         * @constructor
         * @param {$protobuf.RPCImpl} rpcImpl RPC implementation
         * @param {boolean} [requestDelimited=false] Whether requests are length-delimited
         * @param {boolean} [responseDelimited=false] Whether responses are length-delimited
         */
        function MyServer(rpcImpl, requestDelimited, responseDelimited) {
            $protobuf.rpc.Service.call(this, rpcImpl, requestDelimited, responseDelimited);
        }

        (MyServer.prototype = Object.create($protobuf.rpc.Service.prototype)).constructor = MyServer;

        /**
         * Callback as used by {@link content.MyServer#do_}.
         * @memberof content.MyServer
         * @typedef DoCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {content.AllType} [response] AllType
         */

        /**
         * Calls Do.
         * @function do
         * @memberof content.MyServer
         * @instance
         * @param {content.IAllType} request AllType message or plain object
         * @param {content.MyServer.DoCallback} callback Node-style callback called with the error, if any, and AllType
         * @returns {undefined}
         * @variation 1
         */
        MyServer.prototype["do"] = function do_(request, callback) {
            return this.rpcCall(do_, $root.content.AllType, $root.content.AllType, request, callback);
        };

        /**
         * Calls Do.
         * @function do
         * @memberof content.MyServer
         * @instance
         * @param {content.IAllType} request AllType message or plain object
         * @returns {Promise<content.AllType>} Promise
         * @variation 2
         */

        return MyServer;
    })();

    content.KeyValue = (function() {

        /**
         * Properties of a KeyValue.
         * @memberof content
         * @interface IKeyValue
         * @property {string|null} [key] KeyValue key
         * @property {string|null} [value] KeyValue value
         */

        /**
         * Constructs a new KeyValue.
         * @memberof content
         * @classdesc Represents a KeyValue.
         * @implements IKeyValue
         * @constructor
         * @param {content.IKeyValue=} [p] Properties to set
         */
        function KeyValue(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * KeyValue key.
         * @member {string} key
         * @memberof content.KeyValue
         * @instance
         */
        KeyValue.prototype.key = "";

        /**
         * KeyValue value.
         * @member {string} value
         * @memberof content.KeyValue
         * @instance
         */
        KeyValue.prototype.value = "";

        return KeyValue;
    })();

    content.Content = (function() {

        /**
         * Properties of a Content.
         * @memberof content
         * @interface IContent
         * @property {common.IContext|null} [ctx] Content ctx
         * @property {string|null} [sbid] Content sbid
         * @property {string|null} [id] Content id
         * @property {string|null} [account_id] Content account_id
         * @property {string|null} [description] Content description
         * @property {string|null} [title] Content title
         * @property {string|null} [url] Content url
         * @property {Array.<string>|null} [labels] Content labels
         * @property {string|null} [availability] Content availability
         * @property {number|null} [price] Content price
         * @property {string|null} [currency] Content currency
         * @property {number|null} [sale_price] Content sale_price
         * @property {Array.<content.IKeyValue>|null} [fields] Content fields
         * @property {Array.<string>|null} [categories] Content categories
         * @property {Array.<string>|null} [relates] Content relates
         * @property {Array.<string>|null} [attachment_urls] Content attachment_urls
         * @property {string|null} [type] Content type
         */

        /**
         * Constructs a new Content.
         * @memberof content
         * @classdesc Represents a Content.
         * @implements IContent
         * @constructor
         * @param {content.IContent=} [p] Properties to set
         */
        function Content(p) {
            this.labels = [];
            this.fields = [];
            this.categories = [];
            this.relates = [];
            this.attachment_urls = [];
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * Content ctx.
         * @member {common.IContext|null|undefined} ctx
         * @memberof content.Content
         * @instance
         */
        Content.prototype.ctx = null;

        /**
         * Content sbid.
         * @member {string} sbid
         * @memberof content.Content
         * @instance
         */
        Content.prototype.sbid = "";

        /**
         * Content id.
         * @member {string} id
         * @memberof content.Content
         * @instance
         */
        Content.prototype.id = "";

        /**
         * Content account_id.
         * @member {string} account_id
         * @memberof content.Content
         * @instance
         */
        Content.prototype.account_id = "";

        /**
         * Content description.
         * @member {string} description
         * @memberof content.Content
         * @instance
         */
        Content.prototype.description = "";

        /**
         * Content title.
         * @member {string} title
         * @memberof content.Content
         * @instance
         */
        Content.prototype.title = "";

        /**
         * Content url.
         * @member {string} url
         * @memberof content.Content
         * @instance
         */
        Content.prototype.url = "";

        /**
         * Content labels.
         * @member {Array.<string>} labels
         * @memberof content.Content
         * @instance
         */
        Content.prototype.labels = $util.emptyArray;

        /**
         * Content availability.
         * @member {string} availability
         * @memberof content.Content
         * @instance
         */
        Content.prototype.availability = "";

        /**
         * Content price.
         * @member {number} price
         * @memberof content.Content
         * @instance
         */
        Content.prototype.price = 0;

        /**
         * Content currency.
         * @member {string} currency
         * @memberof content.Content
         * @instance
         */
        Content.prototype.currency = "";

        /**
         * Content sale_price.
         * @member {number} sale_price
         * @memberof content.Content
         * @instance
         */
        Content.prototype.sale_price = 0;

        /**
         * Content fields.
         * @member {Array.<content.IKeyValue>} fields
         * @memberof content.Content
         * @instance
         */
        Content.prototype.fields = $util.emptyArray;

        /**
         * Content categories.
         * @member {Array.<string>} categories
         * @memberof content.Content
         * @instance
         */
        Content.prototype.categories = $util.emptyArray;

        /**
         * Content relates.
         * @member {Array.<string>} relates
         * @memberof content.Content
         * @instance
         */
        Content.prototype.relates = $util.emptyArray;

        /**
         * Content attachment_urls.
         * @member {Array.<string>} attachment_urls
         * @memberof content.Content
         * @instance
         */
        Content.prototype.attachment_urls = $util.emptyArray;

        /**
         * Content type.
         * @member {string} type
         * @memberof content.Content
         * @instance
         */
        Content.prototype.type = "";

        return Content;
    })();

    content.Contents = (function() {

        /**
         * Properties of a Contents.
         * @memberof content
         * @interface IContents
         * @property {common.IContext|null} [ctx] Contents ctx
         * @property {Array.<content.IContent>|null} [contents] Contents contents
         */

        /**
         * Constructs a new Contents.
         * @memberof content
         * @classdesc Represents a Contents.
         * @implements IContents
         * @constructor
         * @param {content.IContents=} [p] Properties to set
         */
        function Contents(p) {
            this.contents = [];
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * Contents ctx.
         * @member {common.IContext|null|undefined} ctx
         * @memberof content.Contents
         * @instance
         */
        Contents.prototype.ctx = null;

        /**
         * Contents contents.
         * @member {Array.<content.IContent>} contents
         * @memberof content.Contents
         * @instance
         */
        Contents.prototype.contents = $util.emptyArray;

        return Contents;
    })();

    content.ListRequest = (function() {

        /**
         * Properties of a ListRequest.
         * @memberof content
         * @interface IListRequest
         * @property {string|null} [account_id] ListRequest account_id
         * @property {string|null} [anchor] ListRequest anchor
         * @property {string|null} [category] ListRequest category
         * @property {number|null} [limit] ListRequest limit
         * @property {string|null} [label] ListRequest label
         * @property {string|null} [query] ListRequest query
         */

        /**
         * Constructs a new ListRequest.
         * @memberof content
         * @classdesc Represents a ListRequest.
         * @implements IListRequest
         * @constructor
         * @param {content.IListRequest=} [p] Properties to set
         */
        function ListRequest(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * ListRequest account_id.
         * @member {string} account_id
         * @memberof content.ListRequest
         * @instance
         */
        ListRequest.prototype.account_id = "";

        /**
         * ListRequest anchor.
         * @member {string} anchor
         * @memberof content.ListRequest
         * @instance
         */
        ListRequest.prototype.anchor = "";

        /**
         * ListRequest category.
         * @member {string} category
         * @memberof content.ListRequest
         * @instance
         */
        ListRequest.prototype.category = "";

        /**
         * ListRequest limit.
         * @member {number} limit
         * @memberof content.ListRequest
         * @instance
         */
        ListRequest.prototype.limit = 0;

        /**
         * ListRequest label.
         * @member {string} label
         * @memberof content.ListRequest
         * @instance
         */
        ListRequest.prototype.label = "";

        /**
         * ListRequest query.
         * @member {string} query
         * @memberof content.ListRequest
         * @instance
         */
        ListRequest.prototype.query = "";

        return ListRequest;
    })();

    return content;
})();

export const conversation = $root.conversation = (() => {

    /**
     * Namespace conversation.
     * @exports conversation
     * @namespace
     */
    const conversation = {};

    conversation.Rule = (function() {

        /**
         * Properties of a Rule.
         * @memberof conversation
         * @interface IRule
         * @property {common.IContext|null} [ctx] Rule ctx
         * @property {string|null} [id] Rule id
         * @property {string|null} [account_id] Rule account_id
         * @property {number|null} [priority] Rule priority
         * @property {string|null} [strategy] Rule strategy
         * @property {Array.<string>|null} [assign_tos] Rule assign_tos
         * @property {Array.<conversation.ICondition>|null} [conditions] Rule conditions
         * @property {boolean|null} [enabled] Rule enabled
         * @property {number|Long|null} [created] Rule created
         * @property {number|Long|null} [modified] Rule modified
         * @property {string|null} [name] Rule name
         * @property {string|null} [description] Rule description
         */

        /**
         * Constructs a new Rule.
         * @memberof conversation
         * @classdesc Represents a Rule.
         * @implements IRule
         * @constructor
         * @param {conversation.IRule=} [p] Properties to set
         */
        function Rule(p) {
            this.assign_tos = [];
            this.conditions = [];
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * Rule ctx.
         * @member {common.IContext|null|undefined} ctx
         * @memberof conversation.Rule
         * @instance
         */
        Rule.prototype.ctx = null;

        /**
         * Rule id.
         * @member {string} id
         * @memberof conversation.Rule
         * @instance
         */
        Rule.prototype.id = "";

        /**
         * Rule account_id.
         * @member {string} account_id
         * @memberof conversation.Rule
         * @instance
         */
        Rule.prototype.account_id = "";

        /**
         * Rule priority.
         * @member {number} priority
         * @memberof conversation.Rule
         * @instance
         */
        Rule.prototype.priority = 0;

        /**
         * Rule strategy.
         * @member {string} strategy
         * @memberof conversation.Rule
         * @instance
         */
        Rule.prototype.strategy = "";

        /**
         * Rule assign_tos.
         * @member {Array.<string>} assign_tos
         * @memberof conversation.Rule
         * @instance
         */
        Rule.prototype.assign_tos = $util.emptyArray;

        /**
         * Rule conditions.
         * @member {Array.<conversation.ICondition>} conditions
         * @memberof conversation.Rule
         * @instance
         */
        Rule.prototype.conditions = $util.emptyArray;

        /**
         * Rule enabled.
         * @member {boolean} enabled
         * @memberof conversation.Rule
         * @instance
         */
        Rule.prototype.enabled = false;

        /**
         * Rule created.
         * @member {number|Long} created
         * @memberof conversation.Rule
         * @instance
         */
        Rule.prototype.created = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

        /**
         * Rule modified.
         * @member {number|Long} modified
         * @memberof conversation.Rule
         * @instance
         */
        Rule.prototype.modified = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

        /**
         * Rule name.
         * @member {string} name
         * @memberof conversation.Rule
         * @instance
         */
        Rule.prototype.name = "";

        /**
         * Rule description.
         * @member {string} description
         * @memberof conversation.Rule
         * @instance
         */
        Rule.prototype.description = "";

        /**
         * AssignStrategy enum.
         * @name conversation.Rule.AssignStrategy
         * @enum {string}
         * @property {number} all_agents=0 all_agents value
         * @property {number} agentgroup=2 agentgroup value
         * @property {number} agents=3 agents value
         * @property {number} most_recent=4 most_recent value
         * @property {number} roundrobin_all_agents=5 roundrobin_all_agents value
         * @property {number} roundrobin_agents=6 roundrobin_agents value
         */
        Rule.AssignStrategy = (function() {
            const valuesById = {}, values = Object.create(valuesById);
            values[valuesById[0] = "all_agents"] = 0;
            values[valuesById[2] = "agentgroup"] = 2;
            values[valuesById[3] = "agents"] = 3;
            values[valuesById[4] = "most_recent"] = 4;
            values[valuesById[5] = "roundrobin_all_agents"] = 5;
            values[valuesById[6] = "roundrobin_agents"] = 6;
            return values;
        })();

        return Rule;
    })();

    conversation.AvailabilityCheckRequest = (function() {

        /**
         * Properties of an AvailabilityCheckRequest.
         * @memberof conversation
         * @interface IAvailabilityCheckRequest
         * @property {common.IContext|null} [ctx] AvailabilityCheckRequest ctx
         * @property {string|null} [integration_id] AvailabilityCheckRequest integration_id
         * @property {user.IUser|null} [user] AvailabilityCheckRequest user
         * @property {string|null} [account_id] AvailabilityCheckRequest account_id
         * @property {string|null} [user_id] AvailabilityCheckRequest user_id
         */

        /**
         * Constructs a new AvailabilityCheckRequest.
         * @memberof conversation
         * @classdesc Represents an AvailabilityCheckRequest.
         * @implements IAvailabilityCheckRequest
         * @constructor
         * @param {conversation.IAvailabilityCheckRequest=} [p] Properties to set
         */
        function AvailabilityCheckRequest(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * AvailabilityCheckRequest ctx.
         * @member {common.IContext|null|undefined} ctx
         * @memberof conversation.AvailabilityCheckRequest
         * @instance
         */
        AvailabilityCheckRequest.prototype.ctx = null;

        /**
         * AvailabilityCheckRequest integration_id.
         * @member {string} integration_id
         * @memberof conversation.AvailabilityCheckRequest
         * @instance
         */
        AvailabilityCheckRequest.prototype.integration_id = "";

        /**
         * AvailabilityCheckRequest user.
         * @member {user.IUser|null|undefined} user
         * @memberof conversation.AvailabilityCheckRequest
         * @instance
         */
        AvailabilityCheckRequest.prototype.user = null;

        /**
         * AvailabilityCheckRequest account_id.
         * @member {string} account_id
         * @memberof conversation.AvailabilityCheckRequest
         * @instance
         */
        AvailabilityCheckRequest.prototype.account_id = "";

        /**
         * AvailabilityCheckRequest user_id.
         * @member {string} user_id
         * @memberof conversation.AvailabilityCheckRequest
         * @instance
         */
        AvailabilityCheckRequest.prototype.user_id = "";

        return AvailabilityCheckRequest;
    })();

    conversation.AvailabilityCheckResult = (function() {

        /**
         * Properties of an AvailabilityCheckResult.
         * @memberof conversation
         * @interface IAvailabilityCheckResult
         * @property {common.IContext|null} [ctx] AvailabilityCheckResult ctx
         * @property {boolean|null} [availability] AvailabilityCheckResult availability
         * @property {string|null} [reason] AvailabilityCheckResult reason
         * @property {string|null} [payload] AvailabilityCheckResult payload
         */

        /**
         * Constructs a new AvailabilityCheckResult.
         * @memberof conversation
         * @classdesc Represents an AvailabilityCheckResult.
         * @implements IAvailabilityCheckResult
         * @constructor
         * @param {conversation.IAvailabilityCheckResult=} [p] Properties to set
         */
        function AvailabilityCheckResult(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * AvailabilityCheckResult ctx.
         * @member {common.IContext|null|undefined} ctx
         * @memberof conversation.AvailabilityCheckResult
         * @instance
         */
        AvailabilityCheckResult.prototype.ctx = null;

        /**
         * AvailabilityCheckResult availability.
         * @member {boolean} availability
         * @memberof conversation.AvailabilityCheckResult
         * @instance
         */
        AvailabilityCheckResult.prototype.availability = false;

        /**
         * AvailabilityCheckResult reason.
         * @member {string} reason
         * @memberof conversation.AvailabilityCheckResult
         * @instance
         */
        AvailabilityCheckResult.prototype.reason = "";

        /**
         * AvailabilityCheckResult payload.
         * @member {string} payload
         * @memberof conversation.AvailabilityCheckResult
         * @instance
         */
        AvailabilityCheckResult.prototype.payload = "";

        return AvailabilityCheckResult;
    })();

    conversation.Condition = (function() {

        /**
         * Properties of a Condition.
         * @memberof conversation
         * @interface ICondition
         * @property {string|null} [join] Condition join
         * @property {string|null} [key] Condition key
         * @property {string|null} [operator] Condition operator
         * @property {string|null} [value] Condition value
         */

        /**
         * Constructs a new Condition.
         * @memberof conversation
         * @classdesc Represents a Condition.
         * @implements ICondition
         * @constructor
         * @param {conversation.ICondition=} [p] Properties to set
         */
        function Condition(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * Condition join.
         * @member {string} join
         * @memberof conversation.Condition
         * @instance
         */
        Condition.prototype.join = "";

        /**
         * Condition key.
         * @member {string} key
         * @memberof conversation.Condition
         * @instance
         */
        Condition.prototype.key = "";

        /**
         * Condition operator.
         * @member {string} operator
         * @memberof conversation.Condition
         * @instance
         */
        Condition.prototype.operator = "";

        /**
         * Condition value.
         * @member {string} value
         * @memberof conversation.Condition
         * @instance
         */
        Condition.prototype.value = "";

        /**
         * JoinOperator enum.
         * @name conversation.Condition.JoinOperator
         * @enum {string}
         * @property {number} none=0 none value
         * @property {number} and=1 and value
         * @property {number} or=2 or value
         */
        Condition.JoinOperator = (function() {
            const valuesById = {}, values = Object.create(valuesById);
            values[valuesById[0] = "none"] = 0;
            values[valuesById[1] = "and"] = 1;
            values[valuesById[2] = "or"] = 2;
            return values;
        })();

        return Condition;
    })();

    conversation.Route = (function() {

        /**
         * Properties of a Route.
         * @memberof conversation
         * @interface IRoute
         * @property {common.IContext|null} [ctx] Route ctx
         * @property {Array.<conversation.IRule>|null} [rules] Route rules
         */

        /**
         * Constructs a new Route.
         * @memberof conversation
         * @classdesc Represents a Route.
         * @implements IRoute
         * @constructor
         * @param {conversation.IRoute=} [p] Properties to set
         */
        function Route(p) {
            this.rules = [];
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * Route ctx.
         * @member {common.IContext|null|undefined} ctx
         * @memberof conversation.Route
         * @instance
         */
        Route.prototype.ctx = null;

        /**
         * Route rules.
         * @member {Array.<conversation.IRule>} rules
         * @memberof conversation.Route
         * @instance
         */
        Route.prototype.rules = $util.emptyArray;

        return Route;
    })();

    conversation.RouteResult = (function() {

        /**
         * Properties of a RouteResult.
         * @memberof conversation
         * @interface IRouteResult
         * @property {string|null} [rule_id] RouteResult rule_id
         * @property {string|null} [strategy] RouteResult strategy
         * @property {Array.<string>|null} [agent_ids] RouteResult agent_ids
         * @property {string|null} [group_id] RouteResult group_id
         */

        /**
         * Constructs a new RouteResult.
         * @memberof conversation
         * @classdesc Represents a RouteResult.
         * @implements IRouteResult
         * @constructor
         * @param {conversation.IRouteResult=} [p] Properties to set
         */
        function RouteResult(p) {
            this.agent_ids = [];
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * RouteResult rule_id.
         * @member {string} rule_id
         * @memberof conversation.RouteResult
         * @instance
         */
        RouteResult.prototype.rule_id = "";

        /**
         * RouteResult strategy.
         * @member {string} strategy
         * @memberof conversation.RouteResult
         * @instance
         */
        RouteResult.prototype.strategy = "";

        /**
         * RouteResult agent_ids.
         * @member {Array.<string>} agent_ids
         * @memberof conversation.RouteResult
         * @instance
         */
        RouteResult.prototype.agent_ids = $util.emptyArray;

        /**
         * RouteResult group_id.
         * @member {string} group_id
         * @memberof conversation.RouteResult
         * @instance
         */
        RouteResult.prototype.group_id = "";

        return RouteResult;
    })();

    conversation.MemberSeen = (function() {

        /**
         * Properties of a MemberSeen.
         * @memberof conversation
         * @interface IMemberSeen
         * @property {string|null} [member_id] MemberSeen member_id
         * @property {string|null} [message_id] MemberSeen message_id
         */

        /**
         * Constructs a new MemberSeen.
         * @memberof conversation
         * @classdesc Represents a MemberSeen.
         * @implements IMemberSeen
         * @constructor
         * @param {conversation.IMemberSeen=} [p] Properties to set
         */
        function MemberSeen(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * MemberSeen member_id.
         * @member {string} member_id
         * @memberof conversation.MemberSeen
         * @instance
         */
        MemberSeen.prototype.member_id = "";

        /**
         * MemberSeen message_id.
         * @member {string} message_id
         * @memberof conversation.MemberSeen
         * @instance
         */
        MemberSeen.prototype.message_id = "";

        return MemberSeen;
    })();

    conversation.MemberV3 = (function() {

        /**
         * Properties of a MemberV3.
         * @memberof conversation
         * @interface IMemberV3
         * @property {common.IContext|null} [ctx] MemberV3 ctx
         * @property {string|null} [id] MemberV3 id
         * @property {string|null} [conversation_id] MemberV3 conversation_id
         * @property {string|null} [account_id] MemberV3 account_id
         * @property {string|null} [name] MemberV3 name
         * @property {string|null} [avatar_url] MemberV3 avatar_url
         * @property {string|null} [type] MemberV3 type
         */

        /**
         * Constructs a new MemberV3.
         * @memberof conversation
         * @classdesc Represents a MemberV3.
         * @implements IMemberV3
         * @constructor
         * @param {conversation.IMemberV3=} [p] Properties to set
         */
        function MemberV3(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * MemberV3 ctx.
         * @member {common.IContext|null|undefined} ctx
         * @memberof conversation.MemberV3
         * @instance
         */
        MemberV3.prototype.ctx = null;

        /**
         * MemberV3 id.
         * @member {string} id
         * @memberof conversation.MemberV3
         * @instance
         */
        MemberV3.prototype.id = "";

        /**
         * MemberV3 conversation_id.
         * @member {string} conversation_id
         * @memberof conversation.MemberV3
         * @instance
         */
        MemberV3.prototype.conversation_id = "";

        /**
         * MemberV3 account_id.
         * @member {string} account_id
         * @memberof conversation.MemberV3
         * @instance
         */
        MemberV3.prototype.account_id = "";

        /**
         * MemberV3 name.
         * @member {string} name
         * @memberof conversation.MemberV3
         * @instance
         */
        MemberV3.prototype.name = "";

        /**
         * MemberV3 avatar_url.
         * @member {string} avatar_url
         * @memberof conversation.MemberV3
         * @instance
         */
        MemberV3.prototype.avatar_url = "";

        /**
         * MemberV3 type.
         * @member {string} type
         * @memberof conversation.MemberV3
         * @instance
         */
        MemberV3.prototype.type = "";

        return MemberV3;
    })();

    conversation.Member = (function() {

        /**
         * Properties of a Member.
         * @memberof conversation
         * @interface IMember
         * @property {string|null} [type] Member type
         * @property {string|null} [subiz_id] Member subiz_id
         * @property {string|null} [id] Member id
         * @property {string|null} [name] Member name
         * @property {string|null} [avatar_url] Member avatar_url
         * @property {boolean|null} [left] Member left
         * @property {string|null} [conversation_id] Member conversation_id
         */

        /**
         * Constructs a new Member.
         * @memberof conversation
         * @classdesc Represents a Member.
         * @implements IMember
         * @constructor
         * @param {conversation.IMember=} [p] Properties to set
         */
        function Member(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * Member type.
         * @member {string} type
         * @memberof conversation.Member
         * @instance
         */
        Member.prototype.type = "";

        /**
         * Member subiz_id.
         * @member {string} subiz_id
         * @memberof conversation.Member
         * @instance
         */
        Member.prototype.subiz_id = "";

        /**
         * Member id.
         * @member {string} id
         * @memberof conversation.Member
         * @instance
         */
        Member.prototype.id = "";

        /**
         * Member name.
         * @member {string} name
         * @memberof conversation.Member
         * @instance
         */
        Member.prototype.name = "";

        /**
         * Member avatar_url.
         * @member {string} avatar_url
         * @memberof conversation.Member
         * @instance
         */
        Member.prototype.avatar_url = "";

        /**
         * Member left.
         * @member {boolean} left
         * @memberof conversation.Member
         * @instance
         */
        Member.prototype.left = false;

        /**
         * Member conversation_id.
         * @member {string} conversation_id
         * @memberof conversation.Member
         * @instance
         */
        Member.prototype.conversation_id = "";

        return Member;
    })();

    conversation.Conversation = (function() {

        /**
         * Properties of a Conversation.
         * @memberof conversation
         * @interface IConversation
         * @property {common.IContext|null} [ctx] Conversation ctx
         * @property {string|null} [id] Conversation id
         * @property {string|null} [account_id] Conversation account_id
         * @property {number|Long|null} [created] Conversation created
         * @property {number|Long|null} [closed] Conversation closed
         * @property {Array.<conversation.IMember>|null} [members] Conversation members
         * @property {Array.<conversation.ITag>|null} [tags] Conversation tags
         * @property {string|null} [state] Conversation state
         * @property {conversation.IStartRequest|null} [request] Conversation request
         * @property {number|Long|null} [accepted] Conversation accepted
         * @property {string|null} [channel_type] Conversation channel_type
         * @property {conversation.IIntegration|null} [integration] Conversation integration
         * @property {number|Long|null} [actived] Conversation actived
         * @property {string|null} [last_message_id] Conversation last_message_id
         * @property {conversation.IMessage|null} [last_message] Conversation last_message
         * @property {number|Long|null} [response_sec] Conversation response_sec
         */

        /**
         * Constructs a new Conversation.
         * @memberof conversation
         * @classdesc Represents a Conversation.
         * @implements IConversation
         * @constructor
         * @param {conversation.IConversation=} [p] Properties to set
         */
        function Conversation(p) {
            this.members = [];
            this.tags = [];
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * Conversation ctx.
         * @member {common.IContext|null|undefined} ctx
         * @memberof conversation.Conversation
         * @instance
         */
        Conversation.prototype.ctx = null;

        /**
         * Conversation id.
         * @member {string} id
         * @memberof conversation.Conversation
         * @instance
         */
        Conversation.prototype.id = "";

        /**
         * Conversation account_id.
         * @member {string} account_id
         * @memberof conversation.Conversation
         * @instance
         */
        Conversation.prototype.account_id = "";

        /**
         * Conversation created.
         * @member {number|Long} created
         * @memberof conversation.Conversation
         * @instance
         */
        Conversation.prototype.created = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

        /**
         * Conversation closed.
         * @member {number|Long} closed
         * @memberof conversation.Conversation
         * @instance
         */
        Conversation.prototype.closed = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

        /**
         * Conversation members.
         * @member {Array.<conversation.IMember>} members
         * @memberof conversation.Conversation
         * @instance
         */
        Conversation.prototype.members = $util.emptyArray;

        /**
         * Conversation tags.
         * @member {Array.<conversation.ITag>} tags
         * @memberof conversation.Conversation
         * @instance
         */
        Conversation.prototype.tags = $util.emptyArray;

        /**
         * Conversation state.
         * @member {string} state
         * @memberof conversation.Conversation
         * @instance
         */
        Conversation.prototype.state = "";

        /**
         * Conversation request.
         * @member {conversation.IStartRequest|null|undefined} request
         * @memberof conversation.Conversation
         * @instance
         */
        Conversation.prototype.request = null;

        /**
         * Conversation accepted.
         * @member {number|Long} accepted
         * @memberof conversation.Conversation
         * @instance
         */
        Conversation.prototype.accepted = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

        /**
         * Conversation channel_type.
         * @member {string} channel_type
         * @memberof conversation.Conversation
         * @instance
         */
        Conversation.prototype.channel_type = "";

        /**
         * Conversation integration.
         * @member {conversation.IIntegration|null|undefined} integration
         * @memberof conversation.Conversation
         * @instance
         */
        Conversation.prototype.integration = null;

        /**
         * Conversation actived.
         * @member {number|Long} actived
         * @memberof conversation.Conversation
         * @instance
         */
        Conversation.prototype.actived = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

        /**
         * Conversation last_message_id.
         * @member {string} last_message_id
         * @memberof conversation.Conversation
         * @instance
         */
        Conversation.prototype.last_message_id = "";

        /**
         * Conversation last_message.
         * @member {conversation.IMessage|null|undefined} last_message
         * @memberof conversation.Conversation
         * @instance
         */
        Conversation.prototype.last_message = null;

        /**
         * Conversation response_sec.
         * @member {number|Long} response_sec
         * @memberof conversation.Conversation
         * @instance
         */
        Conversation.prototype.response_sec = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

        /**
         * State enum.
         * @name conversation.Conversation.State
         * @enum {string}
         * @property {number} conversation_none=0 conversation_none value
         * @property {number} unassigned=2 unassigned value
         * @property {number} active=6 active value
         * @property {number} ended=8 ended value
         * @property {number} pending=9 pending value
         */
        Conversation.State = (function() {
            const valuesById = {}, values = Object.create(valuesById);
            values[valuesById[0] = "conversation_none"] = 0;
            values[valuesById[2] = "unassigned"] = 2;
            values[valuesById[6] = "active"] = 6;
            values[valuesById[8] = "ended"] = 8;
            values[valuesById[9] = "pending"] = 9;
            return values;
        })();

        return Conversation;
    })();

    conversation.UserConversation = (function() {

        /**
         * Properties of a UserConversation.
         * @memberof conversation
         * @interface IUserConversation
         * @property {string|null} [account_id] UserConversation account_id
         * @property {string|null} [state] UserConversation state
         * @property {string|null} [user_id] UserConversation user_id
         * @property {string|null} [convo_id] UserConversation convo_id
         * @property {string|null} [last_seen_event] UserConversation last_seen_event
         */

        /**
         * Constructs a new UserConversation.
         * @memberof conversation
         * @classdesc Represents a UserConversation.
         * @implements IUserConversation
         * @constructor
         * @param {conversation.IUserConversation=} [p] Properties to set
         */
        function UserConversation(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * UserConversation account_id.
         * @member {string} account_id
         * @memberof conversation.UserConversation
         * @instance
         */
        UserConversation.prototype.account_id = "";

        /**
         * UserConversation state.
         * @member {string} state
         * @memberof conversation.UserConversation
         * @instance
         */
        UserConversation.prototype.state = "";

        /**
         * UserConversation user_id.
         * @member {string} user_id
         * @memberof conversation.UserConversation
         * @instance
         */
        UserConversation.prototype.user_id = "";

        /**
         * UserConversation convo_id.
         * @member {string} convo_id
         * @memberof conversation.UserConversation
         * @instance
         */
        UserConversation.prototype.convo_id = "";

        /**
         * UserConversation last_seen_event.
         * @member {string} last_seen_event
         * @memberof conversation.UserConversation
         * @instance
         */
        UserConversation.prototype.last_seen_event = "";

        return UserConversation;
    })();

    conversation.StartRequest = (function() {

        /**
         * Properties of a StartRequest.
         * @memberof conversation
         * @interface IStartRequest
         * @property {common.IContext|null} [ctx] StartRequest ctx
         * @property {string|null} [id] StartRequest id
         * @property {string|null} [account_id] StartRequest account_id
         * @property {string|null} [channel_type] StartRequest channel_type
         * @property {string|null} [from] StartRequest from
         * @property {Array.<string>|null} [to] StartRequest to
         * @property {string|null} [page_url] StartRequest page_url
         * @property {string|null} [page_title] StartRequest page_title
         * @property {string|null} [message] StartRequest message
         * @property {string|null} [browser_language] StartRequest browser_language
         * @property {string|null} [language] StartRequest language
         * @property {string|null} [device_type] StartRequest device_type
         * @property {number|Long|null} [created] StartRequest created
         * @property {string|null} [conversation_id] StartRequest conversation_id
         * @property {string|null} [ip] StartRequest ip
         * @property {string|null} [country] StartRequest country
         * @property {string|null} [country_code] StartRequest country_code
         * @property {string|null} [city] StartRequest city
         * @property {string|null} [timezone] StartRequest timezone
         * @property {string|null} [starter_id] StartRequest starter_id
         * @property {string|null} [starter_type] StartRequest starter_type
         * @property {Array.<string>|null} [agent_ids] StartRequest agent_ids
         * @property {user.IUser|null} [user] StartRequest user
         * @property {string|null} [integration_id] StartRequest integration_id
         */

        /**
         * Constructs a new StartRequest.
         * @memberof conversation
         * @classdesc Represents a StartRequest.
         * @implements IStartRequest
         * @constructor
         * @param {conversation.IStartRequest=} [p] Properties to set
         */
        function StartRequest(p) {
            this.to = [];
            this.agent_ids = [];
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * StartRequest ctx.
         * @member {common.IContext|null|undefined} ctx
         * @memberof conversation.StartRequest
         * @instance
         */
        StartRequest.prototype.ctx = null;

        /**
         * StartRequest id.
         * @member {string} id
         * @memberof conversation.StartRequest
         * @instance
         */
        StartRequest.prototype.id = "";

        /**
         * StartRequest account_id.
         * @member {string} account_id
         * @memberof conversation.StartRequest
         * @instance
         */
        StartRequest.prototype.account_id = "";

        /**
         * StartRequest channel_type.
         * @member {string} channel_type
         * @memberof conversation.StartRequest
         * @instance
         */
        StartRequest.prototype.channel_type = "";

        /**
         * StartRequest from.
         * @member {string} from
         * @memberof conversation.StartRequest
         * @instance
         */
        StartRequest.prototype.from = "";

        /**
         * StartRequest to.
         * @member {Array.<string>} to
         * @memberof conversation.StartRequest
         * @instance
         */
        StartRequest.prototype.to = $util.emptyArray;

        /**
         * StartRequest page_url.
         * @member {string} page_url
         * @memberof conversation.StartRequest
         * @instance
         */
        StartRequest.prototype.page_url = "";

        /**
         * StartRequest page_title.
         * @member {string} page_title
         * @memberof conversation.StartRequest
         * @instance
         */
        StartRequest.prototype.page_title = "";

        /**
         * StartRequest message.
         * @member {string} message
         * @memberof conversation.StartRequest
         * @instance
         */
        StartRequest.prototype.message = "";

        /**
         * StartRequest browser_language.
         * @member {string} browser_language
         * @memberof conversation.StartRequest
         * @instance
         */
        StartRequest.prototype.browser_language = "";

        /**
         * StartRequest language.
         * @member {string} language
         * @memberof conversation.StartRequest
         * @instance
         */
        StartRequest.prototype.language = "";

        /**
         * StartRequest device_type.
         * @member {string} device_type
         * @memberof conversation.StartRequest
         * @instance
         */
        StartRequest.prototype.device_type = "";

        /**
         * StartRequest created.
         * @member {number|Long} created
         * @memberof conversation.StartRequest
         * @instance
         */
        StartRequest.prototype.created = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

        /**
         * StartRequest conversation_id.
         * @member {string} conversation_id
         * @memberof conversation.StartRequest
         * @instance
         */
        StartRequest.prototype.conversation_id = "";

        /**
         * StartRequest ip.
         * @member {string} ip
         * @memberof conversation.StartRequest
         * @instance
         */
        StartRequest.prototype.ip = "";

        /**
         * StartRequest country.
         * @member {string} country
         * @memberof conversation.StartRequest
         * @instance
         */
        StartRequest.prototype.country = "";

        /**
         * StartRequest country_code.
         * @member {string} country_code
         * @memberof conversation.StartRequest
         * @instance
         */
        StartRequest.prototype.country_code = "";

        /**
         * StartRequest city.
         * @member {string} city
         * @memberof conversation.StartRequest
         * @instance
         */
        StartRequest.prototype.city = "";

        /**
         * StartRequest timezone.
         * @member {string} timezone
         * @memberof conversation.StartRequest
         * @instance
         */
        StartRequest.prototype.timezone = "";

        /**
         * StartRequest starter_id.
         * @member {string} starter_id
         * @memberof conversation.StartRequest
         * @instance
         */
        StartRequest.prototype.starter_id = "";

        /**
         * StartRequest starter_type.
         * @member {string} starter_type
         * @memberof conversation.StartRequest
         * @instance
         */
        StartRequest.prototype.starter_type = "";

        /**
         * StartRequest agent_ids.
         * @member {Array.<string>} agent_ids
         * @memberof conversation.StartRequest
         * @instance
         */
        StartRequest.prototype.agent_ids = $util.emptyArray;

        /**
         * StartRequest user.
         * @member {user.IUser|null|undefined} user
         * @memberof conversation.StartRequest
         * @instance
         */
        StartRequest.prototype.user = null;

        /**
         * StartRequest integration_id.
         * @member {string} integration_id
         * @memberof conversation.StartRequest
         * @instance
         */
        StartRequest.prototype.integration_id = "";

        return StartRequest;
    })();

    conversation.Conversations = (function() {

        /**
         * Properties of a Conversations.
         * @memberof conversation
         * @interface IConversations
         * @property {common.IContext|null} [ctx] Conversations ctx
         * @property {Array.<conversation.IConversation>|null} [conversations] Conversations conversations
         * @property {string|null} [anchor] Conversations anchor
         * @property {Array.<string>|null} [user_ids] Conversations user_ids
         */

        /**
         * Constructs a new Conversations.
         * @memberof conversation
         * @classdesc Represents a Conversations.
         * @implements IConversations
         * @constructor
         * @param {conversation.IConversations=} [p] Properties to set
         */
        function Conversations(p) {
            this.conversations = [];
            this.user_ids = [];
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * Conversations ctx.
         * @member {common.IContext|null|undefined} ctx
         * @memberof conversation.Conversations
         * @instance
         */
        Conversations.prototype.ctx = null;

        /**
         * Conversations conversations.
         * @member {Array.<conversation.IConversation>} conversations
         * @memberof conversation.Conversations
         * @instance
         */
        Conversations.prototype.conversations = $util.emptyArray;

        /**
         * Conversations anchor.
         * @member {string} anchor
         * @memberof conversation.Conversations
         * @instance
         */
        Conversations.prototype.anchor = "";

        /**
         * Conversations user_ids.
         * @member {Array.<string>} user_ids
         * @memberof conversation.Conversations
         * @instance
         */
        Conversations.prototype.user_ids = $util.emptyArray;

        return Conversations;
    })();

    conversation.Search = (function() {

        /**
         * Properties of a Search.
         * @memberof conversation
         * @interface ISearch
         * @property {common.IContext|null} [ctx] Search ctx
         * @property {string|null} [account_id] Search account_id
         * @property {string|null} [keyword] Search keyword
         * @property {number|null} [limit] Search limit
         * @property {string|null} [before_id] Search before_id
         * @property {string|null} [after_id] Search after_id
         */

        /**
         * Constructs a new Search.
         * @memberof conversation
         * @classdesc Represents a Search.
         * @implements ISearch
         * @constructor
         * @param {conversation.ISearch=} [p] Properties to set
         */
        function Search(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * Search ctx.
         * @member {common.IContext|null|undefined} ctx
         * @memberof conversation.Search
         * @instance
         */
        Search.prototype.ctx = null;

        /**
         * Search account_id.
         * @member {string} account_id
         * @memberof conversation.Search
         * @instance
         */
        Search.prototype.account_id = "";

        /**
         * Search keyword.
         * @member {string} keyword
         * @memberof conversation.Search
         * @instance
         */
        Search.prototype.keyword = "";

        /**
         * Search limit.
         * @member {number} limit
         * @memberof conversation.Search
         * @instance
         */
        Search.prototype.limit = 0;

        /**
         * Search before_id.
         * @member {string} before_id
         * @memberof conversation.Search
         * @instance
         */
        Search.prototype.before_id = "";

        /**
         * Search after_id.
         * @member {string} after_id
         * @memberof conversation.Search
         * @instance
         */
        Search.prototype.after_id = "";

        return Search;
    })();

    conversation.ListConversationsRequest = (function() {

        /**
         * Properties of a ListConversationsRequest.
         * @memberof conversation
         * @interface IListConversationsRequest
         * @property {common.IContext|null} [ctx] ListConversationsRequest ctx
         * @property {string|null} [account_id] ListConversationsRequest account_id
         * @property {string|null} [state] ListConversationsRequest state
         * @property {number|null} [limit] ListConversationsRequest limit
         * @property {string|null} [anchor] ListConversationsRequest anchor
         * @property {string|null} [member_id] ListConversationsRequest member_id
         * @property {string|null} [group_by] ListConversationsRequest group_by
         * @property {string|null} [integration_id] ListConversationsRequest integration_id
         */

        /**
         * Constructs a new ListConversationsRequest.
         * @memberof conversation
         * @classdesc Represents a ListConversationsRequest.
         * @implements IListConversationsRequest
         * @constructor
         * @param {conversation.IListConversationsRequest=} [p] Properties to set
         */
        function ListConversationsRequest(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * ListConversationsRequest ctx.
         * @member {common.IContext|null|undefined} ctx
         * @memberof conversation.ListConversationsRequest
         * @instance
         */
        ListConversationsRequest.prototype.ctx = null;

        /**
         * ListConversationsRequest account_id.
         * @member {string} account_id
         * @memberof conversation.ListConversationsRequest
         * @instance
         */
        ListConversationsRequest.prototype.account_id = "";

        /**
         * ListConversationsRequest state.
         * @member {string} state
         * @memberof conversation.ListConversationsRequest
         * @instance
         */
        ListConversationsRequest.prototype.state = "";

        /**
         * ListConversationsRequest limit.
         * @member {number} limit
         * @memberof conversation.ListConversationsRequest
         * @instance
         */
        ListConversationsRequest.prototype.limit = 0;

        /**
         * ListConversationsRequest anchor.
         * @member {string} anchor
         * @memberof conversation.ListConversationsRequest
         * @instance
         */
        ListConversationsRequest.prototype.anchor = "";

        /**
         * ListConversationsRequest member_id.
         * @member {string} member_id
         * @memberof conversation.ListConversationsRequest
         * @instance
         */
        ListConversationsRequest.prototype.member_id = "";

        /**
         * ListConversationsRequest group_by.
         * @member {string} group_by
         * @memberof conversation.ListConversationsRequest
         * @instance
         */
        ListConversationsRequest.prototype.group_by = "";

        /**
         * ListConversationsRequest integration_id.
         * @member {string} integration_id
         * @memberof conversation.ListConversationsRequest
         * @instance
         */
        ListConversationsRequest.prototype.integration_id = "";

        return ListConversationsRequest;
    })();

    conversation.ListEventsRequest = (function() {

        /**
         * Properties of a ListEventsRequest.
         * @memberof conversation
         * @interface IListEventsRequest
         * @property {common.IContext|null} [ctx] ListEventsRequest ctx
         * @property {string|null} [account_id] ListEventsRequest account_id
         * @property {string|null} [conversation_id] ListEventsRequest conversation_id
         * @property {string|null} [start_id] ListEventsRequest start_id
         * @property {number|null} [limit] ListEventsRequest limit
         */

        /**
         * Constructs a new ListEventsRequest.
         * @memberof conversation
         * @classdesc Represents a ListEventsRequest.
         * @implements IListEventsRequest
         * @constructor
         * @param {conversation.IListEventsRequest=} [p] Properties to set
         */
        function ListEventsRequest(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * ListEventsRequest ctx.
         * @member {common.IContext|null|undefined} ctx
         * @memberof conversation.ListEventsRequest
         * @instance
         */
        ListEventsRequest.prototype.ctx = null;

        /**
         * ListEventsRequest account_id.
         * @member {string} account_id
         * @memberof conversation.ListEventsRequest
         * @instance
         */
        ListEventsRequest.prototype.account_id = "";

        /**
         * ListEventsRequest conversation_id.
         * @member {string} conversation_id
         * @memberof conversation.ListEventsRequest
         * @instance
         */
        ListEventsRequest.prototype.conversation_id = "";

        /**
         * ListEventsRequest start_id.
         * @member {string} start_id
         * @memberof conversation.ListEventsRequest
         * @instance
         */
        ListEventsRequest.prototype.start_id = "";

        /**
         * ListEventsRequest limit.
         * @member {number} limit
         * @memberof conversation.ListEventsRequest
         * @instance
         */
        ListEventsRequest.prototype.limit = 0;

        return ListEventsRequest;
    })();

    conversation.ListConversationsByUserRequest = (function() {

        /**
         * Properties of a ListConversationsByUserRequest.
         * @memberof conversation
         * @interface IListConversationsByUserRequest
         * @property {string|null} [account_id] ListConversationsByUserRequest account_id
         * @property {string|null} [channel_id] ListConversationsByUserRequest channel_id
         * @property {string|null} [user_id] ListConversationsByUserRequest user_id
         * @property {string|null} [start_id] ListConversationsByUserRequest start_id
         * @property {number|null} [limit] ListConversationsByUserRequest limit
         */

        /**
         * Constructs a new ListConversationsByUserRequest.
         * @memberof conversation
         * @classdesc Represents a ListConversationsByUserRequest.
         * @implements IListConversationsByUserRequest
         * @constructor
         * @param {conversation.IListConversationsByUserRequest=} [p] Properties to set
         */
        function ListConversationsByUserRequest(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * ListConversationsByUserRequest account_id.
         * @member {string} account_id
         * @memberof conversation.ListConversationsByUserRequest
         * @instance
         */
        ListConversationsByUserRequest.prototype.account_id = "";

        /**
         * ListConversationsByUserRequest channel_id.
         * @member {string} channel_id
         * @memberof conversation.ListConversationsByUserRequest
         * @instance
         */
        ListConversationsByUserRequest.prototype.channel_id = "";

        /**
         * ListConversationsByUserRequest user_id.
         * @member {string} user_id
         * @memberof conversation.ListConversationsByUserRequest
         * @instance
         */
        ListConversationsByUserRequest.prototype.user_id = "";

        /**
         * ListConversationsByUserRequest start_id.
         * @member {string} start_id
         * @memberof conversation.ListConversationsByUserRequest
         * @instance
         */
        ListConversationsByUserRequest.prototype.start_id = "";

        /**
         * ListConversationsByUserRequest limit.
         * @member {number} limit
         * @memberof conversation.ListConversationsByUserRequest
         * @instance
         */
        ListConversationsByUserRequest.prototype.limit = 0;

        return ListConversationsByUserRequest;
    })();

    /**
     * Event enum.
     * @name conversation.Event
     * @enum {string}
     * @property {number} ConversationAssigned=0 ConversationAssigned value
     * @property {number} ConversationWaiting=2 ConversationWaiting value
     * @property {number} ConversationStartRequested=3 ConversationStartRequested value
     * @property {number} ConversationAccepted=4 ConversationAccepted value
     * @property {number} ConversationDropped=5 ConversationDropped value
     * @property {number} ConversationEventCreated=6 ConversationEventCreated value
     * @property {number} ConversationJoinRequested=7 ConversationJoinRequested value
     * @property {number} ConversationMessageRequested=8 ConversationMessageRequested value
     * @property {number} ConversationLeaveRequested=9 ConversationLeaveRequested value
     * @property {number} ConversationCloseRequested=10 ConversationCloseRequested value
     * @property {number} ConversationTagRequested=11 ConversationTagRequested value
     * @property {number} ConversationUntagRequested=12 ConversationUntagRequested value
     * @property {number} ConversationReadRequested=13 ConversationReadRequested value
     * @property {number} ConversationListRequested=14 ConversationListRequested value
     * @property {number} ConversationAcceptRequested=15 ConversationAcceptRequested value
     * @property {number} ConversationUpdateRuleRequested=20 ConversationUpdateRuleRequested value
     * @property {number} ConversationCreateRuleRequested=21 ConversationCreateRuleRequested value
     * @property {number} ConversationDeleteRuleRequested=22 ConversationDeleteRuleRequested value
     * @property {number} ConversationReadRuleRequested=23 ConversationReadRuleRequested value
     * @property {number} ConversationListRuleRequested=24 ConversationListRuleRequested value
     * @property {number} ConversationUserRequestReply=51 ConversationUserRequestReply value
     * @property {number} ConversationLimitUpdated=52 ConversationLimitUpdated value
     * @property {number} ConversationRequestWaitTimeout=60 ConversationRequestWaitTimeout value
     * @property {number} ConversationListEventsRequested=61 ConversationListEventsRequested value
     * @property {number} ChannelDeintegrateRequested=65 ChannelDeintegrateRequested value
     * @property {number} ChannelIntegrateRequested=66 ChannelIntegrateRequested value
     * @property {number} ChannelIntegrationListRequested=67 ChannelIntegrationListRequested value
     * @property {number} ConnectorUpsertRequested=68 ConnectorUpsertRequested value
     * @property {number} ConnectorListRequested=69 ConnectorListRequested value
     * @property {number} CannedResponseCreateRequested=80 CannedResponseCreateRequested value
     * @property {number} CannedResponseUpdateRequested=81 CannedResponseUpdateRequested value
     * @property {number} CannedResponseDeleteRequested=82 CannedResponseDeleteRequested value
     * @property {number} CannedResponseReadRequested=83 CannedResponseReadRequested value
     * @property {number} CannedResponseListRequested=84 CannedResponseListRequested value
     * @property {number} TagCreateRequested=85 TagCreateRequested value
     * @property {number} TagUpdateRequested=86 TagUpdateRequested value
     * @property {number} TagReadRequested=87 TagReadRequested value
     * @property {number} TagDeleteRequested=88 TagDeleteRequested value
     * @property {number} TagListRequested=89 TagListRequested value
     * @property {number} TagCreated=95 TagCreated value
     * @property {number} ConversationUpserted=97 ConversationUpserted value
     * @property {number} ConversationMessageSent=99 ConversationMessageSent value
     * @property {number} ConversationMessageAckRequested=90 ConversationMessageAckRequested value
     * @property {number} ConversationMessageReceiveRequested=91 ConversationMessageReceiveRequested value
     * @property {number} ConversationMessageSeeRequested=92 ConversationMessageSeeRequested value
     * @property {number} ChannelIntegrationAvailabilityCheck=93 ChannelIntegrationAvailabilityCheck value
     * @property {number} ConversationMessageSearchRequest=94 ConversationMessageSearchRequest value
     * @property {number} ConversationRequested=100 ConversationRequested value
     * @property {number} ConversationSynced=101 ConversationSynced value
     * @property {number} ConversationV3Synced=102 ConversationV3Synced value
     */
    conversation.Event = (function() {
        const valuesById = {}, values = Object.create(valuesById);
        values[valuesById[0] = "ConversationAssigned"] = 0;
        values[valuesById[2] = "ConversationWaiting"] = 2;
        values[valuesById[3] = "ConversationStartRequested"] = 3;
        values[valuesById[4] = "ConversationAccepted"] = 4;
        values[valuesById[5] = "ConversationDropped"] = 5;
        values[valuesById[6] = "ConversationEventCreated"] = 6;
        values[valuesById[7] = "ConversationJoinRequested"] = 7;
        values[valuesById[8] = "ConversationMessageRequested"] = 8;
        values[valuesById[9] = "ConversationLeaveRequested"] = 9;
        values[valuesById[10] = "ConversationCloseRequested"] = 10;
        values[valuesById[11] = "ConversationTagRequested"] = 11;
        values[valuesById[12] = "ConversationUntagRequested"] = 12;
        values[valuesById[13] = "ConversationReadRequested"] = 13;
        values[valuesById[14] = "ConversationListRequested"] = 14;
        values[valuesById[15] = "ConversationAcceptRequested"] = 15;
        values[valuesById[20] = "ConversationUpdateRuleRequested"] = 20;
        values[valuesById[21] = "ConversationCreateRuleRequested"] = 21;
        values[valuesById[22] = "ConversationDeleteRuleRequested"] = 22;
        values[valuesById[23] = "ConversationReadRuleRequested"] = 23;
        values[valuesById[24] = "ConversationListRuleRequested"] = 24;
        values[valuesById[51] = "ConversationUserRequestReply"] = 51;
        values[valuesById[52] = "ConversationLimitUpdated"] = 52;
        values[valuesById[60] = "ConversationRequestWaitTimeout"] = 60;
        values[valuesById[61] = "ConversationListEventsRequested"] = 61;
        values[valuesById[65] = "ChannelDeintegrateRequested"] = 65;
        values[valuesById[66] = "ChannelIntegrateRequested"] = 66;
        values[valuesById[67] = "ChannelIntegrationListRequested"] = 67;
        values[valuesById[68] = "ConnectorUpsertRequested"] = 68;
        values[valuesById[69] = "ConnectorListRequested"] = 69;
        values[valuesById[80] = "CannedResponseCreateRequested"] = 80;
        values[valuesById[81] = "CannedResponseUpdateRequested"] = 81;
        values[valuesById[82] = "CannedResponseDeleteRequested"] = 82;
        values[valuesById[83] = "CannedResponseReadRequested"] = 83;
        values[valuesById[84] = "CannedResponseListRequested"] = 84;
        values[valuesById[85] = "TagCreateRequested"] = 85;
        values[valuesById[86] = "TagUpdateRequested"] = 86;
        values[valuesById[87] = "TagReadRequested"] = 87;
        values[valuesById[88] = "TagDeleteRequested"] = 88;
        values[valuesById[89] = "TagListRequested"] = 89;
        values[valuesById[95] = "TagCreated"] = 95;
        values[valuesById[97] = "ConversationUpserted"] = 97;
        values[valuesById[99] = "ConversationMessageSent"] = 99;
        values[valuesById[90] = "ConversationMessageAckRequested"] = 90;
        values[valuesById[91] = "ConversationMessageReceiveRequested"] = 91;
        values[valuesById[92] = "ConversationMessageSeeRequested"] = 92;
        values[valuesById[93] = "ChannelIntegrationAvailabilityCheck"] = 93;
        values[valuesById[94] = "ConversationMessageSearchRequest"] = 94;
        values[valuesById[100] = "ConversationRequested"] = 100;
        values[valuesById[101] = "ConversationSynced"] = 101;
        values[valuesById[102] = "ConversationV3Synced"] = 102;
        return values;
    })();

    conversation.RequestState = (function() {

        /**
         * Properties of a RequestState.
         * @memberof conversation
         * @interface IRequestState
         * @property {string|null} [account_id] RequestState account_id
         * @property {string|null} [conversation_id] RequestState conversation_id
         * @property {string|null} [channel_id] RequestState channel_id
         */

        /**
         * Constructs a new RequestState.
         * @memberof conversation
         * @classdesc Represents a RequestState.
         * @implements IRequestState
         * @constructor
         * @param {conversation.IRequestState=} [p] Properties to set
         */
        function RequestState(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * RequestState account_id.
         * @member {string} account_id
         * @memberof conversation.RequestState
         * @instance
         */
        RequestState.prototype.account_id = "";

        /**
         * RequestState conversation_id.
         * @member {string} conversation_id
         * @memberof conversation.RequestState
         * @instance
         */
        RequestState.prototype.conversation_id = "";

        /**
         * RequestState channel_id.
         * @member {string} channel_id
         * @memberof conversation.RequestState
         * @instance
         */
        RequestState.prototype.channel_id = "";

        return RequestState;
    })();

    conversation.Reaction = (function() {

        /**
         * Properties of a Reaction.
         * @memberof conversation
         * @interface IReaction
         * @property {string|null} [name] Reaction name
         * @property {number|null} [count] Reaction count
         * @property {Array.<string>|null} [users] Reaction users
         */

        /**
         * Constructs a new Reaction.
         * @memberof conversation
         * @classdesc Represents a Reaction.
         * @implements IReaction
         * @constructor
         * @param {conversation.IReaction=} [p] Properties to set
         */
        function Reaction(p) {
            this.users = [];
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * Reaction name.
         * @member {string} name
         * @memberof conversation.Reaction
         * @instance
         */
        Reaction.prototype.name = "";

        /**
         * Reaction count.
         * @member {number} count
         * @memberof conversation.Reaction
         * @instance
         */
        Reaction.prototype.count = 0;

        /**
         * Reaction users.
         * @member {Array.<string>} users
         * @memberof conversation.Reaction
         * @instance
         */
        Reaction.prototype.users = $util.emptyArray;

        return Reaction;
    })();

    conversation.EsMessage = (function() {

        /**
         * Properties of an EsMessage.
         * @memberof conversation
         * @interface IEsMessage
         * @property {string|null} [id] EsMessage id
         * @property {string|null} [account_id] EsMessage account_id
         * @property {string|null} [conversation_id] EsMessage conversation_id
         * @property {Array.<string>|null} [member_ids] EsMessage member_ids
         * @property {string|null} [text] EsMessage text
         * @property {Array.<string>|null} [attachments] EsMessage attachments
         * @property {Array.<string>|null} [fields] EsMessage fields
         */

        /**
         * Constructs a new EsMessage.
         * @memberof conversation
         * @classdesc Represents an EsMessage.
         * @implements IEsMessage
         * @constructor
         * @param {conversation.IEsMessage=} [p] Properties to set
         */
        function EsMessage(p) {
            this.member_ids = [];
            this.attachments = [];
            this.fields = [];
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * EsMessage id.
         * @member {string} id
         * @memberof conversation.EsMessage
         * @instance
         */
        EsMessage.prototype.id = "";

        /**
         * EsMessage account_id.
         * @member {string} account_id
         * @memberof conversation.EsMessage
         * @instance
         */
        EsMessage.prototype.account_id = "";

        /**
         * EsMessage conversation_id.
         * @member {string} conversation_id
         * @memberof conversation.EsMessage
         * @instance
         */
        EsMessage.prototype.conversation_id = "";

        /**
         * EsMessage member_ids.
         * @member {Array.<string>} member_ids
         * @memberof conversation.EsMessage
         * @instance
         */
        EsMessage.prototype.member_ids = $util.emptyArray;

        /**
         * EsMessage text.
         * @member {string} text
         * @memberof conversation.EsMessage
         * @instance
         */
        EsMessage.prototype.text = "";

        /**
         * EsMessage attachments.
         * @member {Array.<string>} attachments
         * @memberof conversation.EsMessage
         * @instance
         */
        EsMessage.prototype.attachments = $util.emptyArray;

        /**
         * EsMessage fields.
         * @member {Array.<string>} fields
         * @memberof conversation.EsMessage
         * @instance
         */
        EsMessage.prototype.fields = $util.emptyArray;

        return EsMessage;
    })();

    conversation.Message = (function() {

        /**
         * Properties of a Message.
         * @memberof conversation
         * @interface IMessage
         * @property {common.IContext|null} [ctx] Message ctx
         * @property {string|null} [account_id] Message account_id
         * @property {string|null} [conversation_id] Message conversation_id
         * @property {string|null} [id] Message id
         * @property {string|null} [text] Message text
         * @property {string|null} [format] Message format
         * @property {Array.<conversation.IAttachment>|null} [attachments] Message attachments
         * @property {Array.<conversation.IReaction>|null} [reactions] Message reactions
         * @property {Array.<conversation.IField>|null} [fields] Message fields
         * @property {conversation.IComputed|null} [computed] Message computed
         * @property {string|null} [integration_id] Message integration_id
         */

        /**
         * Constructs a new Message.
         * @memberof conversation
         * @classdesc Represents a Message.
         * @implements IMessage
         * @constructor
         * @param {conversation.IMessage=} [p] Properties to set
         */
        function Message(p) {
            this.attachments = [];
            this.reactions = [];
            this.fields = [];
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * Message ctx.
         * @member {common.IContext|null|undefined} ctx
         * @memberof conversation.Message
         * @instance
         */
        Message.prototype.ctx = null;

        /**
         * Message account_id.
         * @member {string} account_id
         * @memberof conversation.Message
         * @instance
         */
        Message.prototype.account_id = "";

        /**
         * Message conversation_id.
         * @member {string} conversation_id
         * @memberof conversation.Message
         * @instance
         */
        Message.prototype.conversation_id = "";

        /**
         * Message id.
         * @member {string} id
         * @memberof conversation.Message
         * @instance
         */
        Message.prototype.id = "";

        /**
         * Message text.
         * @member {string} text
         * @memberof conversation.Message
         * @instance
         */
        Message.prototype.text = "";

        /**
         * Message format.
         * @member {string} format
         * @memberof conversation.Message
         * @instance
         */
        Message.prototype.format = "";

        /**
         * Message attachments.
         * @member {Array.<conversation.IAttachment>} attachments
         * @memberof conversation.Message
         * @instance
         */
        Message.prototype.attachments = $util.emptyArray;

        /**
         * Message reactions.
         * @member {Array.<conversation.IReaction>} reactions
         * @memberof conversation.Message
         * @instance
         */
        Message.prototype.reactions = $util.emptyArray;

        /**
         * Message fields.
         * @member {Array.<conversation.IField>} fields
         * @memberof conversation.Message
         * @instance
         */
        Message.prototype.fields = $util.emptyArray;

        /**
         * Message computed.
         * @member {conversation.IComputed|null|undefined} computed
         * @memberof conversation.Message
         * @instance
         */
        Message.prototype.computed = null;

        /**
         * Message integration_id.
         * @member {string} integration_id
         * @memberof conversation.Message
         * @instance
         */
        Message.prototype.integration_id = "";

        return Message;
    })();

    conversation.Computed = (function() {

        /**
         * Properties of a Computed.
         * @memberof conversation
         * @interface IComputed
         * @property {Array.<conversation.ISeen>|null} [seen] Computed seen
         * @property {Array.<conversation.IAck>|null} [ack] Computed ack
         * @property {Array.<conversation.IReceived>|null} [received] Computed received
         */

        /**
         * Constructs a new Computed.
         * @memberof conversation
         * @classdesc Represents a Computed.
         * @implements IComputed
         * @constructor
         * @param {conversation.IComputed=} [p] Properties to set
         */
        function Computed(p) {
            this.seen = [];
            this.ack = [];
            this.received = [];
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * Computed seen.
         * @member {Array.<conversation.ISeen>} seen
         * @memberof conversation.Computed
         * @instance
         */
        Computed.prototype.seen = $util.emptyArray;

        /**
         * Computed ack.
         * @member {Array.<conversation.IAck>} ack
         * @memberof conversation.Computed
         * @instance
         */
        Computed.prototype.ack = $util.emptyArray;

        /**
         * Computed received.
         * @member {Array.<conversation.IReceived>} received
         * @memberof conversation.Computed
         * @instance
         */
        Computed.prototype.received = $util.emptyArray;

        return Computed;
    })();

    conversation.Seen = (function() {

        /**
         * Properties of a Seen.
         * @memberof conversation
         * @interface ISeen
         * @property {string|null} [member_id] Seen member_id
         * @property {number|Long|null} [at] Seen at
         */

        /**
         * Constructs a new Seen.
         * @memberof conversation
         * @classdesc Represents a Seen.
         * @implements ISeen
         * @constructor
         * @param {conversation.ISeen=} [p] Properties to set
         */
        function Seen(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * Seen member_id.
         * @member {string} member_id
         * @memberof conversation.Seen
         * @instance
         */
        Seen.prototype.member_id = "";

        /**
         * Seen at.
         * @member {number|Long} at
         * @memberof conversation.Seen
         * @instance
         */
        Seen.prototype.at = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

        return Seen;
    })();

    conversation.Ack = (function() {

        /**
         * Properties of an Ack.
         * @memberof conversation
         * @interface IAck
         * @property {string|null} [member_id] Ack member_id
         * @property {string|null} [error] Ack error
         * @property {number|Long|null} [at] Ack at
         */

        /**
         * Constructs a new Ack.
         * @memberof conversation
         * @classdesc Represents an Ack.
         * @implements IAck
         * @constructor
         * @param {conversation.IAck=} [p] Properties to set
         */
        function Ack(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * Ack member_id.
         * @member {string} member_id
         * @memberof conversation.Ack
         * @instance
         */
        Ack.prototype.member_id = "";

        /**
         * Ack error.
         * @member {string} error
         * @memberof conversation.Ack
         * @instance
         */
        Ack.prototype.error = "";

        /**
         * Ack at.
         * @member {number|Long} at
         * @memberof conversation.Ack
         * @instance
         */
        Ack.prototype.at = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

        return Ack;
    })();

    conversation.Received = (function() {

        /**
         * Properties of a Received.
         * @memberof conversation
         * @interface IReceived
         * @property {string|null} [member_id] Received member_id
         * @property {number|Long|null} [at] Received at
         */

        /**
         * Constructs a new Received.
         * @memberof conversation
         * @classdesc Represents a Received.
         * @implements IReceived
         * @constructor
         * @param {conversation.IReceived=} [p] Properties to set
         */
        function Received(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * Received member_id.
         * @member {string} member_id
         * @memberof conversation.Received
         * @instance
         */
        Received.prototype.member_id = "";

        /**
         * Received at.
         * @member {number|Long} at
         * @memberof conversation.Received
         * @instance
         */
        Received.prototype.at = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

        return Received;
    })();

    conversation.Field = (function() {

        /**
         * Properties of a Field.
         * @memberof conversation
         * @interface IField
         * @property {string|null} [value] Field value
         * @property {string|null} [key] Field key
         */

        /**
         * Constructs a new Field.
         * @memberof conversation
         * @classdesc Represents a Field.
         * @implements IField
         * @constructor
         * @param {conversation.IField=} [p] Properties to set
         */
        function Field(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * Field value.
         * @member {string} value
         * @memberof conversation.Field
         * @instance
         */
        Field.prototype.value = "";

        /**
         * Field key.
         * @member {string} key
         * @memberof conversation.Field
         * @instance
         */
        Field.prototype.key = "";

        return Field;
    })();

    conversation.Button = (function() {

        /**
         * Properties of a Button.
         * @memberof conversation
         * @interface IButton
         * @property {string|null} [type] Button type
         * @property {string|null} [id] Button id
         * @property {string|null} [title] Button title
         * @property {string|null} [payload] Button payload
         * @property {string|null} [image_url] Button image_url
         * @property {string|null} [content_id] Button content_id
         * @property {string|null} [url] Button url
         */

        /**
         * Constructs a new Button.
         * @memberof conversation
         * @classdesc Represents a Button.
         * @implements IButton
         * @constructor
         * @param {conversation.IButton=} [p] Properties to set
         */
        function Button(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * Button type.
         * @member {string} type
         * @memberof conversation.Button
         * @instance
         */
        Button.prototype.type = "";

        /**
         * Button id.
         * @member {string} id
         * @memberof conversation.Button
         * @instance
         */
        Button.prototype.id = "";

        /**
         * Button title.
         * @member {string} title
         * @memberof conversation.Button
         * @instance
         */
        Button.prototype.title = "";

        /**
         * Button payload.
         * @member {string} payload
         * @memberof conversation.Button
         * @instance
         */
        Button.prototype.payload = "";

        /**
         * Button image_url.
         * @member {string} image_url
         * @memberof conversation.Button
         * @instance
         */
        Button.prototype.image_url = "";

        /**
         * Button content_id.
         * @member {string} content_id
         * @memberof conversation.Button
         * @instance
         */
        Button.prototype.content_id = "";

        /**
         * Button url.
         * @member {string} url
         * @memberof conversation.Button
         * @instance
         */
        Button.prototype.url = "";

        /**
         * ButtonType enum.
         * @name conversation.Button.ButtonType
         * @enum {string}
         * @property {number} url_button=2 url_button value
         * @property {number} postback_button=3 postback_button value
         * @property {number} event_button=4 event_button value
         */
        Button.ButtonType = (function() {
            const valuesById = {}, values = Object.create(valuesById);
            values[valuesById[2] = "url_button"] = 2;
            values[valuesById[3] = "postback_button"] = 3;
            values[valuesById[4] = "event_button"] = 4;
            return values;
        })();

        return Button;
    })();

    conversation.AskInfomationAnswer = (function() {

        /**
         * Properties of an AskInfomationAnswer.
         * @memberof conversation
         * @interface IAskInfomationAnswer
         * @property {string|null} [message_id] AskInfomationAnswer message_id
         * @property {string|null} [answer] AskInfomationAnswer answer
         */

        /**
         * Constructs a new AskInfomationAnswer.
         * @memberof conversation
         * @classdesc Represents an AskInfomationAnswer.
         * @implements IAskInfomationAnswer
         * @constructor
         * @param {conversation.IAskInfomationAnswer=} [p] Properties to set
         */
        function AskInfomationAnswer(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * AskInfomationAnswer message_id.
         * @member {string} message_id
         * @memberof conversation.AskInfomationAnswer
         * @instance
         */
        AskInfomationAnswer.prototype.message_id = "";

        /**
         * AskInfomationAnswer answer.
         * @member {string} answer
         * @memberof conversation.AskInfomationAnswer
         * @instance
         */
        AskInfomationAnswer.prototype.answer = "";

        return AskInfomationAnswer;
    })();

    conversation.AskInfomation = (function() {

        /**
         * Properties of an AskInfomation.
         * @memberof conversation
         * @interface IAskInfomation
         * @property {string|null} [question] AskInfomation question
         * @property {string|null} [input_type] AskInfomation input_type
         * @property {string|null} [key] AskInfomation key
         * @property {string|null} [answer] AskInfomation answer
         * @property {number|Long|null} [answered] AskInfomation answered
         */

        /**
         * Constructs a new AskInfomation.
         * @memberof conversation
         * @classdesc Represents an AskInfomation.
         * @implements IAskInfomation
         * @constructor
         * @param {conversation.IAskInfomation=} [p] Properties to set
         */
        function AskInfomation(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * AskInfomation question.
         * @member {string} question
         * @memberof conversation.AskInfomation
         * @instance
         */
        AskInfomation.prototype.question = "";

        /**
         * AskInfomation input_type.
         * @member {string} input_type
         * @memberof conversation.AskInfomation
         * @instance
         */
        AskInfomation.prototype.input_type = "";

        /**
         * AskInfomation key.
         * @member {string} key
         * @memberof conversation.AskInfomation
         * @instance
         */
        AskInfomation.prototype.key = "";

        /**
         * AskInfomation answer.
         * @member {string} answer
         * @memberof conversation.AskInfomation
         * @instance
         */
        AskInfomation.prototype.answer = "";

        /**
         * AskInfomation answered.
         * @member {number|Long} answered
         * @memberof conversation.AskInfomation
         * @instance
         */
        AskInfomation.prototype.answered = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

        /**
         * InputType enum.
         * @name conversation.AskInfomation.InputType
         * @enum {string}
         * @property {number} phone=0 phone value
         * @property {number} email=1 email value
         * @property {number} text=2 text value
         * @property {number} password=3 password value
         * @property {number} number=4 number value
         * @property {number} date=5 date value
         * @property {number} color=6 color value
         * @property {number} location=7 location value
         * @property {number} time=8 time value
         * @property {number} url=9 url value
         */
        AskInfomation.InputType = (function() {
            const valuesById = {}, values = Object.create(valuesById);
            values[valuesById[0] = "phone"] = 0;
            values[valuesById[1] = "email"] = 1;
            values[valuesById[2] = "text"] = 2;
            values[valuesById[3] = "password"] = 3;
            values[valuesById[4] = "number"] = 4;
            values[valuesById[5] = "date"] = 5;
            values[valuesById[6] = "color"] = 6;
            values[valuesById[7] = "location"] = 7;
            values[valuesById[8] = "time"] = 8;
            values[valuesById[9] = "url"] = 9;
            return values;
        })();

        return AskInfomation;
    })();

    conversation.GenericElementTemplate = (function() {

        /**
         * Properties of a GenericElementTemplate.
         * @memberof conversation
         * @interface IGenericElementTemplate
         * @property {string|null} [title] GenericElementTemplate title
         * @property {string|null} [image_url] GenericElementTemplate image_url
         * @property {string|null} [subtitle] GenericElementTemplate subtitle
         * @property {conversation.IButton|null} [default_action] GenericElementTemplate default_action
         * @property {Array.<conversation.IButton>|null} [buttons] GenericElementTemplate buttons
         */

        /**
         * Constructs a new GenericElementTemplate.
         * @memberof conversation
         * @classdesc Represents a GenericElementTemplate.
         * @implements IGenericElementTemplate
         * @constructor
         * @param {conversation.IGenericElementTemplate=} [p] Properties to set
         */
        function GenericElementTemplate(p) {
            this.buttons = [];
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * GenericElementTemplate title.
         * @member {string} title
         * @memberof conversation.GenericElementTemplate
         * @instance
         */
        GenericElementTemplate.prototype.title = "";

        /**
         * GenericElementTemplate image_url.
         * @member {string} image_url
         * @memberof conversation.GenericElementTemplate
         * @instance
         */
        GenericElementTemplate.prototype.image_url = "";

        /**
         * GenericElementTemplate subtitle.
         * @member {string} subtitle
         * @memberof conversation.GenericElementTemplate
         * @instance
         */
        GenericElementTemplate.prototype.subtitle = "";

        /**
         * GenericElementTemplate default_action.
         * @member {conversation.IButton|null|undefined} default_action
         * @memberof conversation.GenericElementTemplate
         * @instance
         */
        GenericElementTemplate.prototype.default_action = null;

        /**
         * GenericElementTemplate buttons.
         * @member {Array.<conversation.IButton>} buttons
         * @memberof conversation.GenericElementTemplate
         * @instance
         */
        GenericElementTemplate.prototype.buttons = $util.emptyArray;

        return GenericElementTemplate;
    })();

    conversation.Attachment = (function() {

        /**
         * Properties of an Attachment.
         * @memberof conversation
         * @interface IAttachment
         * @property {string|null} [type] Attachment type
         * @property {string|null} [mimetype] Attachment mimetype
         * @property {string|null} [url] Attachment url
         * @property {string|null} [thumbnail_url] Attachment thumbnail_url
         * @property {string|null} [name] Attachment name
         * @property {string|null} [description] Attachment description
         * @property {number|null} [length] Attachment length
         * @property {number|null} [size] Attachment size
         * @property {Array.<conversation.IGenericElementTemplate>|null} [elements] Attachment elements
         * @property {string|null} [title] Attachment title
         * @property {string|null} [color] Attachment color
         * @property {string|null} [pretext] Attachment pretext
         * @property {Array.<conversation.IButton>|null} [buttons] Attachment buttons
         * @property {conversation.IAskInfomation|null} [ask_info] Attachment ask_info
         * @property {conversation.IAskInfomationAnswer|null} [ask_info_answer] Attachment ask_info_answer
         */

        /**
         * Constructs a new Attachment.
         * @memberof conversation
         * @classdesc Represents an Attachment.
         * @implements IAttachment
         * @constructor
         * @param {conversation.IAttachment=} [p] Properties to set
         */
        function Attachment(p) {
            this.elements = [];
            this.buttons = [];
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * Attachment type.
         * @member {string} type
         * @memberof conversation.Attachment
         * @instance
         */
        Attachment.prototype.type = "";

        /**
         * Attachment mimetype.
         * @member {string} mimetype
         * @memberof conversation.Attachment
         * @instance
         */
        Attachment.prototype.mimetype = "";

        /**
         * Attachment url.
         * @member {string} url
         * @memberof conversation.Attachment
         * @instance
         */
        Attachment.prototype.url = "";

        /**
         * Attachment thumbnail_url.
         * @member {string} thumbnail_url
         * @memberof conversation.Attachment
         * @instance
         */
        Attachment.prototype.thumbnail_url = "";

        /**
         * Attachment name.
         * @member {string} name
         * @memberof conversation.Attachment
         * @instance
         */
        Attachment.prototype.name = "";

        /**
         * Attachment description.
         * @member {string} description
         * @memberof conversation.Attachment
         * @instance
         */
        Attachment.prototype.description = "";

        /**
         * Attachment length.
         * @member {number} length
         * @memberof conversation.Attachment
         * @instance
         */
        Attachment.prototype.length = 0;

        /**
         * Attachment size.
         * @member {number} size
         * @memberof conversation.Attachment
         * @instance
         */
        Attachment.prototype.size = 0;

        /**
         * Attachment elements.
         * @member {Array.<conversation.IGenericElementTemplate>} elements
         * @memberof conversation.Attachment
         * @instance
         */
        Attachment.prototype.elements = $util.emptyArray;

        /**
         * Attachment title.
         * @member {string} title
         * @memberof conversation.Attachment
         * @instance
         */
        Attachment.prototype.title = "";

        /**
         * Attachment color.
         * @member {string} color
         * @memberof conversation.Attachment
         * @instance
         */
        Attachment.prototype.color = "";

        /**
         * Attachment pretext.
         * @member {string} pretext
         * @memberof conversation.Attachment
         * @instance
         */
        Attachment.prototype.pretext = "";

        /**
         * Attachment buttons.
         * @member {Array.<conversation.IButton>} buttons
         * @memberof conversation.Attachment
         * @instance
         */
        Attachment.prototype.buttons = $util.emptyArray;

        /**
         * Attachment ask_info.
         * @member {conversation.IAskInfomation|null|undefined} ask_info
         * @memberof conversation.Attachment
         * @instance
         */
        Attachment.prototype.ask_info = null;

        /**
         * Attachment ask_info_answer.
         * @member {conversation.IAskInfomationAnswer|null|undefined} ask_info_answer
         * @memberof conversation.Attachment
         * @instance
         */
        Attachment.prototype.ask_info_answer = null;

        /**
         * AttachmentType enum.
         * @name conversation.Attachment.AttachmentType
         * @enum {string}
         * @property {number} file=2 file value
         * @property {number} generic=3 generic value
         * @property {number} preview=4 preview value
         * @property {number} button=5 button value
         * @property {number} input=6 input value
         * @property {number} ask_info_form=7 ask_info_form value
         * @property {number} ask_info_form_answer=8 ask_info_form_answer value
         */
        Attachment.AttachmentType = (function() {
            const valuesById = {}, values = Object.create(valuesById);
            values[valuesById[2] = "file"] = 2;
            values[valuesById[3] = "generic"] = 3;
            values[valuesById[4] = "preview"] = 4;
            values[valuesById[5] = "button"] = 5;
            values[valuesById[6] = "input"] = 6;
            values[valuesById[7] = "ask_info_form"] = 7;
            values[valuesById[8] = "ask_info_form_answer"] = 8;
            return values;
        })();

        return Attachment;
    })();

    conversation.Tag = (function() {

        /**
         * Properties of a Tag.
         * @memberof conversation
         * @interface ITag
         * @property {common.IContext|null} [ctx] Tag ctx
         * @property {string|null} [id] Tag id
         * @property {string|null} [account_id] Tag account_id
         * @property {string|null} [title] Tag title
         * @property {number|Long|null} [created] Tag created
         * @property {number|Long|null} [modified] Tag modified
         * @property {string|null} [creator_id] Tag creator_id
         * @property {string|null} [color] Tag color
         */

        /**
         * Constructs a new Tag.
         * @memberof conversation
         * @classdesc Represents a Tag.
         * @implements ITag
         * @constructor
         * @param {conversation.ITag=} [p] Properties to set
         */
        function Tag(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * Tag ctx.
         * @member {common.IContext|null|undefined} ctx
         * @memberof conversation.Tag
         * @instance
         */
        Tag.prototype.ctx = null;

        /**
         * Tag id.
         * @member {string} id
         * @memberof conversation.Tag
         * @instance
         */
        Tag.prototype.id = "";

        /**
         * Tag account_id.
         * @member {string} account_id
         * @memberof conversation.Tag
         * @instance
         */
        Tag.prototype.account_id = "";

        /**
         * Tag title.
         * @member {string} title
         * @memberof conversation.Tag
         * @instance
         */
        Tag.prototype.title = "";

        /**
         * Tag created.
         * @member {number|Long} created
         * @memberof conversation.Tag
         * @instance
         */
        Tag.prototype.created = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

        /**
         * Tag modified.
         * @member {number|Long} modified
         * @memberof conversation.Tag
         * @instance
         */
        Tag.prototype.modified = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

        /**
         * Tag creator_id.
         * @member {string} creator_id
         * @memberof conversation.Tag
         * @instance
         */
        Tag.prototype.creator_id = "";

        /**
         * Tag color.
         * @member {string} color
         * @memberof conversation.Tag
         * @instance
         */
        Tag.prototype.color = "";

        return Tag;
    })();

    conversation.CannedResponse = (function() {

        /**
         * Properties of a CannedResponse.
         * @memberof conversation
         * @interface ICannedResponse
         * @property {common.IContext|null} [ctx] CannedResponse ctx
         * @property {string|null} [id] CannedResponse id
         * @property {string|null} [account_id] CannedResponse account_id
         * @property {string|null} [text] CannedResponse text
         * @property {Array.<string>|null} [keys] CannedResponse keys
         * @property {number|Long|null} [created] CannedResponse created
         * @property {number|Long|null} [modified] CannedResponse modified
         * @property {string|null} [creator] CannedResponse creator
         */

        /**
         * Constructs a new CannedResponse.
         * @memberof conversation
         * @classdesc Represents a CannedResponse.
         * @implements ICannedResponse
         * @constructor
         * @param {conversation.ICannedResponse=} [p] Properties to set
         */
        function CannedResponse(p) {
            this.keys = [];
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * CannedResponse ctx.
         * @member {common.IContext|null|undefined} ctx
         * @memberof conversation.CannedResponse
         * @instance
         */
        CannedResponse.prototype.ctx = null;

        /**
         * CannedResponse id.
         * @member {string} id
         * @memberof conversation.CannedResponse
         * @instance
         */
        CannedResponse.prototype.id = "";

        /**
         * CannedResponse account_id.
         * @member {string} account_id
         * @memberof conversation.CannedResponse
         * @instance
         */
        CannedResponse.prototype.account_id = "";

        /**
         * CannedResponse text.
         * @member {string} text
         * @memberof conversation.CannedResponse
         * @instance
         */
        CannedResponse.prototype.text = "";

        /**
         * CannedResponse keys.
         * @member {Array.<string>} keys
         * @memberof conversation.CannedResponse
         * @instance
         */
        CannedResponse.prototype.keys = $util.emptyArray;

        /**
         * CannedResponse created.
         * @member {number|Long} created
         * @memberof conversation.CannedResponse
         * @instance
         */
        CannedResponse.prototype.created = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

        /**
         * CannedResponse modified.
         * @member {number|Long} modified
         * @memberof conversation.CannedResponse
         * @instance
         */
        CannedResponse.prototype.modified = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

        /**
         * CannedResponse creator.
         * @member {string} creator
         * @memberof conversation.CannedResponse
         * @instance
         */
        CannedResponse.prototype.creator = "";

        return CannedResponse;
    })();

    conversation.CannedResponses = (function() {

        /**
         * Properties of a CannedResponses.
         * @memberof conversation
         * @interface ICannedResponses
         * @property {common.IContext|null} [ctx] CannedResponses ctx
         * @property {Array.<conversation.ICannedResponse>|null} [responses] CannedResponses responses
         */

        /**
         * Constructs a new CannedResponses.
         * @memberof conversation
         * @classdesc Represents a CannedResponses.
         * @implements ICannedResponses
         * @constructor
         * @param {conversation.ICannedResponses=} [p] Properties to set
         */
        function CannedResponses(p) {
            this.responses = [];
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * CannedResponses ctx.
         * @member {common.IContext|null|undefined} ctx
         * @memberof conversation.CannedResponses
         * @instance
         */
        CannedResponses.prototype.ctx = null;

        /**
         * CannedResponses responses.
         * @member {Array.<conversation.ICannedResponse>} responses
         * @memberof conversation.CannedResponses
         * @instance
         */
        CannedResponses.prototype.responses = $util.emptyArray;

        return CannedResponses;
    })();

    conversation.Tags = (function() {

        /**
         * Properties of a Tags.
         * @memberof conversation
         * @interface ITags
         * @property {common.IContext|null} [ctx] Tags ctx
         * @property {Array.<conversation.ITag>|null} [tags] Tags tags
         */

        /**
         * Constructs a new Tags.
         * @memberof conversation
         * @classdesc Represents a Tags.
         * @implements ITags
         * @constructor
         * @param {conversation.ITags=} [p] Properties to set
         */
        function Tags(p) {
            this.tags = [];
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * Tags ctx.
         * @member {common.IContext|null|undefined} ctx
         * @memberof conversation.Tags
         * @instance
         */
        Tags.prototype.ctx = null;

        /**
         * Tags tags.
         * @member {Array.<conversation.ITag>} tags
         * @memberof conversation.Tags
         * @instance
         */
        Tags.prototype.tags = $util.emptyArray;

        return Tags;
    })();

    conversation.Postback = (function() {

        /**
         * Properties of a Postback.
         * @memberof conversation
         * @interface IPostback
         * @property {conversation.IMessage|null} [message] Postback message
         * @property {conversation.IButton|null} [button] Postback button
         */

        /**
         * Constructs a new Postback.
         * @memberof conversation
         * @classdesc Represents a Postback.
         * @implements IPostback
         * @constructor
         * @param {conversation.IPostback=} [p] Properties to set
         */
        function Postback(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * Postback message.
         * @member {conversation.IMessage|null|undefined} message
         * @memberof conversation.Postback
         * @instance
         */
        Postback.prototype.message = null;

        /**
         * Postback button.
         * @member {conversation.IButton|null|undefined} button
         * @memberof conversation.Postback
         * @instance
         */
        Postback.prototype.button = null;

        return Postback;
    })();

    conversation.UseConnector = (function() {

        /**
         * Properties of a UseConnector.
         * @memberof conversation
         * @interface IUseConnector
         * @property {string|null} [account_id] UseConnector account_id
         * @property {string|null} [connector_id] UseConnector connector_id
         * @property {string|null} [state] UseConnector state
         * @property {string|null} [updated] UseConnector updated
         * @property {string|null} [by] UseConnector by
         */

        /**
         * Constructs a new UseConnector.
         * @memberof conversation
         * @classdesc Represents a UseConnector.
         * @implements IUseConnector
         * @constructor
         * @param {conversation.IUseConnector=} [p] Properties to set
         */
        function UseConnector(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * UseConnector account_id.
         * @member {string} account_id
         * @memberof conversation.UseConnector
         * @instance
         */
        UseConnector.prototype.account_id = "";

        /**
         * UseConnector connector_id.
         * @member {string} connector_id
         * @memberof conversation.UseConnector
         * @instance
         */
        UseConnector.prototype.connector_id = "";

        /**
         * UseConnector state.
         * @member {string} state
         * @memberof conversation.UseConnector
         * @instance
         */
        UseConnector.prototype.state = "";

        /**
         * UseConnector updated.
         * @member {string} updated
         * @memberof conversation.UseConnector
         * @instance
         */
        UseConnector.prototype.updated = "";

        /**
         * UseConnector by.
         * @member {string} by
         * @memberof conversation.UseConnector
         * @instance
         */
        UseConnector.prototype.by = "";

        /**
         * State enum.
         * @name conversation.UseConnector.State
         * @enum {string}
         * @property {number} disabled=0 disabled value
         * @property {number} activated=1 activated value
         */
        UseConnector.State = (function() {
            const valuesById = {}, values = Object.create(valuesById);
            values[valuesById[0] = "disabled"] = 0;
            values[valuesById[1] = "activated"] = 1;
            return values;
        })();

        return UseConnector;
    })();

    conversation.Integrations = (function() {

        /**
         * Properties of an Integrations.
         * @memberof conversation
         * @interface IIntegrations
         * @property {common.IContext|null} [ctx] Integrations ctx
         * @property {string|null} [account_id] Integrations account_id
         * @property {Array.<conversation.IIntegration>|null} [integrations] Integrations integrations
         */

        /**
         * Constructs a new Integrations.
         * @memberof conversation
         * @classdesc Represents an Integrations.
         * @implements IIntegrations
         * @constructor
         * @param {conversation.IIntegrations=} [p] Properties to set
         */
        function Integrations(p) {
            this.integrations = [];
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * Integrations ctx.
         * @member {common.IContext|null|undefined} ctx
         * @memberof conversation.Integrations
         * @instance
         */
        Integrations.prototype.ctx = null;

        /**
         * Integrations account_id.
         * @member {string} account_id
         * @memberof conversation.Integrations
         * @instance
         */
        Integrations.prototype.account_id = "";

        /**
         * Integrations integrations.
         * @member {Array.<conversation.IIntegration>} integrations
         * @memberof conversation.Integrations
         * @instance
         */
        Integrations.prototype.integrations = $util.emptyArray;

        return Integrations;
    })();

    conversation.Integration = (function() {

        /**
         * Properties of an Integration.
         * @memberof conversation
         * @interface IIntegration
         * @property {common.IContext|null} [ctx] Integration ctx
         * @property {string|null} [account_id] Integration account_id
         * @property {string|null} [connector_id] Integration connector_id
         * @property {string|null} [logo_url] Integration logo_url
         * @property {string|null} [name] Integration name
         * @property {string|null} [connector_type] Integration connector_type
         * @property {number|Long|null} [integrated] Integration integrated
         * @property {number|Long|null} [updated] Integration updated
         * @property {string|null} [state] Integration state
         * @property {string|null} [id] Integration id
         */

        /**
         * Constructs a new Integration.
         * @memberof conversation
         * @classdesc Represents an Integration.
         * @implements IIntegration
         * @constructor
         * @param {conversation.IIntegration=} [p] Properties to set
         */
        function Integration(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * Integration ctx.
         * @member {common.IContext|null|undefined} ctx
         * @memberof conversation.Integration
         * @instance
         */
        Integration.prototype.ctx = null;

        /**
         * Integration account_id.
         * @member {string} account_id
         * @memberof conversation.Integration
         * @instance
         */
        Integration.prototype.account_id = "";

        /**
         * Integration connector_id.
         * @member {string} connector_id
         * @memberof conversation.Integration
         * @instance
         */
        Integration.prototype.connector_id = "";

        /**
         * Integration logo_url.
         * @member {string} logo_url
         * @memberof conversation.Integration
         * @instance
         */
        Integration.prototype.logo_url = "";

        /**
         * Integration name.
         * @member {string} name
         * @memberof conversation.Integration
         * @instance
         */
        Integration.prototype.name = "";

        /**
         * Integration connector_type.
         * @member {string} connector_type
         * @memberof conversation.Integration
         * @instance
         */
        Integration.prototype.connector_type = "";

        /**
         * Integration integrated.
         * @member {number|Long} integrated
         * @memberof conversation.Integration
         * @instance
         */
        Integration.prototype.integrated = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

        /**
         * Integration updated.
         * @member {number|Long} updated
         * @memberof conversation.Integration
         * @instance
         */
        Integration.prototype.updated = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

        /**
         * Integration state.
         * @member {string} state
         * @memberof conversation.Integration
         * @instance
         */
        Integration.prototype.state = "";

        /**
         * Integration id.
         * @member {string} id
         * @memberof conversation.Integration
         * @instance
         */
        Integration.prototype.id = "";

        /**
         * State enum.
         * @name conversation.Integration.State
         * @enum {string}
         * @property {number} disabled=0 disabled value
         * @property {number} activated=1 activated value
         */
        Integration.State = (function() {
            const valuesById = {}, values = Object.create(valuesById);
            values[valuesById[0] = "disabled"] = 0;
            values[valuesById[1] = "activated"] = 1;
            return values;
        })();

        return Integration;
    })();

    conversation.SearchMessageRequest = (function() {

        /**
         * Properties of a SearchMessageRequest.
         * @memberof conversation
         * @interface ISearchMessageRequest
         * @property {common.IContext|null} [ctx] SearchMessageRequest ctx
         * @property {string|null} [account_id] SearchMessageRequest account_id
         * @property {string|null} [conversation_id] SearchMessageRequest conversation_id
         * @property {Array.<string>|null} [user_ids] SearchMessageRequest user_ids
         * @property {string|null} [query] SearchMessageRequest query
         * @property {string|null} [anchor] SearchMessageRequest anchor
         * @property {number|null} [limit] SearchMessageRequest limit
         */

        /**
         * Constructs a new SearchMessageRequest.
         * @memberof conversation
         * @classdesc Represents a SearchMessageRequest.
         * @implements ISearchMessageRequest
         * @constructor
         * @param {conversation.ISearchMessageRequest=} [p] Properties to set
         */
        function SearchMessageRequest(p) {
            this.user_ids = [];
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * SearchMessageRequest ctx.
         * @member {common.IContext|null|undefined} ctx
         * @memberof conversation.SearchMessageRequest
         * @instance
         */
        SearchMessageRequest.prototype.ctx = null;

        /**
         * SearchMessageRequest account_id.
         * @member {string} account_id
         * @memberof conversation.SearchMessageRequest
         * @instance
         */
        SearchMessageRequest.prototype.account_id = "";

        /**
         * SearchMessageRequest conversation_id.
         * @member {string} conversation_id
         * @memberof conversation.SearchMessageRequest
         * @instance
         */
        SearchMessageRequest.prototype.conversation_id = "";

        /**
         * SearchMessageRequest user_ids.
         * @member {Array.<string>} user_ids
         * @memberof conversation.SearchMessageRequest
         * @instance
         */
        SearchMessageRequest.prototype.user_ids = $util.emptyArray;

        /**
         * SearchMessageRequest query.
         * @member {string} query
         * @memberof conversation.SearchMessageRequest
         * @instance
         */
        SearchMessageRequest.prototype.query = "";

        /**
         * SearchMessageRequest anchor.
         * @member {string} anchor
         * @memberof conversation.SearchMessageRequest
         * @instance
         */
        SearchMessageRequest.prototype.anchor = "";

        /**
         * SearchMessageRequest limit.
         * @member {number} limit
         * @memberof conversation.SearchMessageRequest
         * @instance
         */
        SearchMessageRequest.prototype.limit = 0;

        return SearchMessageRequest;
    })();

    conversation.RuleMgr = (function() {

        /**
         * Constructs a new RuleMgr service.
         * @memberof conversation
         * @classdesc Represents a RuleMgr
         * @extends $protobuf.rpc.Service
         * @constructor
         * @param {$protobuf.RPCImpl} rpcImpl RPC implementation
         * @param {boolean} [requestDelimited=false] Whether requests are length-delimited
         * @param {boolean} [responseDelimited=false] Whether responses are length-delimited
         */
        function RuleMgr(rpcImpl, requestDelimited, responseDelimited) {
            $protobuf.rpc.Service.call(this, rpcImpl, requestDelimited, responseDelimited);
        }

        (RuleMgr.prototype = Object.create($protobuf.rpc.Service.prototype)).constructor = RuleMgr;

        /**
         * Callback as used by {@link conversation.RuleMgr#updateRule}.
         * @memberof conversation.RuleMgr
         * @typedef UpdateRuleCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {conversation.Rule} [response] Rule
         */

        /**
         * Calls UpdateRule.
         * @function updateRule
         * @memberof conversation.RuleMgr
         * @instance
         * @param {conversation.IRule} request Rule message or plain object
         * @param {conversation.RuleMgr.UpdateRuleCallback} callback Node-style callback called with the error, if any, and Rule
         * @returns {undefined}
         * @variation 1
         */
        RuleMgr.prototype.updateRule = function updateRule(request, callback) {
            return this.rpcCall(updateRule, $root.conversation.Rule, $root.conversation.Rule, request, callback);
        };

        /**
         * Calls UpdateRule.
         * @function updateRule
         * @memberof conversation.RuleMgr
         * @instance
         * @param {conversation.IRule} request Rule message or plain object
         * @returns {Promise<conversation.Rule>} Promise
         * @variation 2
         */

        /**
         * Callback as used by {@link conversation.RuleMgr#createRule}.
         * @memberof conversation.RuleMgr
         * @typedef CreateRuleCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {conversation.Rule} [response] Rule
         */

        /**
         * Calls CreateRule.
         * @function createRule
         * @memberof conversation.RuleMgr
         * @instance
         * @param {conversation.IRule} request Rule message or plain object
         * @param {conversation.RuleMgr.CreateRuleCallback} callback Node-style callback called with the error, if any, and Rule
         * @returns {undefined}
         * @variation 1
         */
        RuleMgr.prototype.createRule = function createRule(request, callback) {
            return this.rpcCall(createRule, $root.conversation.Rule, $root.conversation.Rule, request, callback);
        };

        /**
         * Calls CreateRule.
         * @function createRule
         * @memberof conversation.RuleMgr
         * @instance
         * @param {conversation.IRule} request Rule message or plain object
         * @returns {Promise<conversation.Rule>} Promise
         * @variation 2
         */

        /**
         * Callback as used by {@link conversation.RuleMgr#deleteRule}.
         * @memberof conversation.RuleMgr
         * @typedef DeleteRuleCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {common.Empty} [response] Empty
         */

        /**
         * Calls DeleteRule.
         * @function deleteRule
         * @memberof conversation.RuleMgr
         * @instance
         * @param {common.IId} request Id message or plain object
         * @param {conversation.RuleMgr.DeleteRuleCallback} callback Node-style callback called with the error, if any, and Empty
         * @returns {undefined}
         * @variation 1
         */
        RuleMgr.prototype.deleteRule = function deleteRule(request, callback) {
            return this.rpcCall(deleteRule, $root.common.Id, $root.common.Empty, request, callback);
        };

        /**
         * Calls DeleteRule.
         * @function deleteRule
         * @memberof conversation.RuleMgr
         * @instance
         * @param {common.IId} request Id message or plain object
         * @returns {Promise<common.Empty>} Promise
         * @variation 2
         */

        /**
         * Callback as used by {@link conversation.RuleMgr#readRule}.
         * @memberof conversation.RuleMgr
         * @typedef ReadRuleCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {conversation.Rule} [response] Rule
         */

        /**
         * Calls ReadRule.
         * @function readRule
         * @memberof conversation.RuleMgr
         * @instance
         * @param {common.IId} request Id message or plain object
         * @param {conversation.RuleMgr.ReadRuleCallback} callback Node-style callback called with the error, if any, and Rule
         * @returns {undefined}
         * @variation 1
         */
        RuleMgr.prototype.readRule = function readRule(request, callback) {
            return this.rpcCall(readRule, $root.common.Id, $root.conversation.Rule, request, callback);
        };

        /**
         * Calls ReadRule.
         * @function readRule
         * @memberof conversation.RuleMgr
         * @instance
         * @param {common.IId} request Id message or plain object
         * @returns {Promise<conversation.Rule>} Promise
         * @variation 2
         */

        /**
         * Callback as used by {@link conversation.RuleMgr#listRules}.
         * @memberof conversation.RuleMgr
         * @typedef ListRulesCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {conversation.Route} [response] Route
         */

        /**
         * Calls ListRules.
         * @function listRules
         * @memberof conversation.RuleMgr
         * @instance
         * @param {common.IId} request Id message or plain object
         * @param {conversation.RuleMgr.ListRulesCallback} callback Node-style callback called with the error, if any, and Route
         * @returns {undefined}
         * @variation 1
         */
        RuleMgr.prototype.listRules = function listRules(request, callback) {
            return this.rpcCall(listRules, $root.common.Id, $root.conversation.Route, request, callback);
        };

        /**
         * Calls ListRules.
         * @function listRules
         * @memberof conversation.RuleMgr
         * @instance
         * @param {common.IId} request Id message or plain object
         * @returns {Promise<conversation.Route>} Promise
         * @variation 2
         */

        return RuleMgr;
    })();

    conversation.MessageId = (function() {

        /**
         * Properties of a MessageId.
         * @memberof conversation
         * @interface IMessageId
         * @property {string|null} [conversation_id] MessageId conversation_id
         * @property {string|null} [message_id] MessageId message_id
         */

        /**
         * Constructs a new MessageId.
         * @memberof conversation
         * @classdesc Represents a MessageId.
         * @implements IMessageId
         * @constructor
         * @param {conversation.IMessageId=} [p] Properties to set
         */
        function MessageId(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * MessageId conversation_id.
         * @member {string} conversation_id
         * @memberof conversation.MessageId
         * @instance
         */
        MessageId.prototype.conversation_id = "";

        /**
         * MessageId message_id.
         * @member {string} message_id
         * @memberof conversation.MessageId
         * @instance
         */
        MessageId.prototype.message_id = "";

        return MessageId;
    })();

    conversation.MessageAck = (function() {

        /**
         * Properties of a MessageAck.
         * @memberof conversation
         * @interface IMessageAck
         * @property {string|null} [conversation_id] MessageAck conversation_id
         * @property {string|null} [message_id] MessageAck message_id
         * @property {string|null} [error] MessageAck error
         */

        /**
         * Constructs a new MessageAck.
         * @memberof conversation
         * @classdesc Represents a MessageAck.
         * @implements IMessageAck
         * @constructor
         * @param {conversation.IMessageAck=} [p] Properties to set
         */
        function MessageAck(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * MessageAck conversation_id.
         * @member {string} conversation_id
         * @memberof conversation.MessageAck
         * @instance
         */
        MessageAck.prototype.conversation_id = "";

        /**
         * MessageAck message_id.
         * @member {string} message_id
         * @memberof conversation.MessageAck
         * @instance
         */
        MessageAck.prototype.message_id = "";

        /**
         * MessageAck error.
         * @member {string} error
         * @memberof conversation.MessageAck
         * @instance
         */
        MessageAck.prototype.error = "";

        return MessageAck;
    })();

    conversation.ConversationMgr = (function() {

        /**
         * Constructs a new ConversationMgr service.
         * @memberof conversation
         * @classdesc Represents a ConversationMgr
         * @extends $protobuf.rpc.Service
         * @constructor
         * @param {$protobuf.RPCImpl} rpcImpl RPC implementation
         * @param {boolean} [requestDelimited=false] Whether requests are length-delimited
         * @param {boolean} [responseDelimited=false] Whether responses are length-delimited
         */
        function ConversationMgr(rpcImpl, requestDelimited, responseDelimited) {
            $protobuf.rpc.Service.call(this, rpcImpl, requestDelimited, responseDelimited);
        }

        (ConversationMgr.prototype = Object.create($protobuf.rpc.Service.prototype)).constructor = ConversationMgr;

        /**
         * Callback as used by {@link conversation.ConversationMgr#seenMessage}.
         * @memberof conversation.ConversationMgr
         * @typedef SeenMessageCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {common.Empty} [response] Empty
         */

        /**
         * Calls SeenMessage.
         * @function seenMessage
         * @memberof conversation.ConversationMgr
         * @instance
         * @param {conversation.IMessageId} request MessageId message or plain object
         * @param {conversation.ConversationMgr.SeenMessageCallback} callback Node-style callback called with the error, if any, and Empty
         * @returns {undefined}
         * @variation 1
         */
        ConversationMgr.prototype.seenMessage = function seenMessage(request, callback) {
            return this.rpcCall(seenMessage, $root.conversation.MessageId, $root.common.Empty, request, callback);
        };

        /**
         * Calls SeenMessage.
         * @function seenMessage
         * @memberof conversation.ConversationMgr
         * @instance
         * @param {conversation.IMessageId} request MessageId message or plain object
         * @returns {Promise<common.Empty>} Promise
         * @variation 2
         */

        /**
         * Callback as used by {@link conversation.ConversationMgr#receiveMessage}.
         * @memberof conversation.ConversationMgr
         * @typedef ReceiveMessageCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {common.Empty} [response] Empty
         */

        /**
         * Calls ReceiveMessage.
         * @function receiveMessage
         * @memberof conversation.ConversationMgr
         * @instance
         * @param {conversation.IMessageId} request MessageId message or plain object
         * @param {conversation.ConversationMgr.ReceiveMessageCallback} callback Node-style callback called with the error, if any, and Empty
         * @returns {undefined}
         * @variation 1
         */
        ConversationMgr.prototype.receiveMessage = function receiveMessage(request, callback) {
            return this.rpcCall(receiveMessage, $root.conversation.MessageId, $root.common.Empty, request, callback);
        };

        /**
         * Calls ReceiveMessage.
         * @function receiveMessage
         * @memberof conversation.ConversationMgr
         * @instance
         * @param {conversation.IMessageId} request MessageId message or plain object
         * @returns {Promise<common.Empty>} Promise
         * @variation 2
         */

        /**
         * Callback as used by {@link conversation.ConversationMgr#ackMessage}.
         * @memberof conversation.ConversationMgr
         * @typedef AckMessageCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {common.Empty} [response] Empty
         */

        /**
         * Calls AckMessage.
         * @function ackMessage
         * @memberof conversation.ConversationMgr
         * @instance
         * @param {conversation.IMessageAck} request MessageAck message or plain object
         * @param {conversation.ConversationMgr.AckMessageCallback} callback Node-style callback called with the error, if any, and Empty
         * @returns {undefined}
         * @variation 1
         */
        ConversationMgr.prototype.ackMessage = function ackMessage(request, callback) {
            return this.rpcCall(ackMessage, $root.conversation.MessageAck, $root.common.Empty, request, callback);
        };

        /**
         * Calls AckMessage.
         * @function ackMessage
         * @memberof conversation.ConversationMgr
         * @instance
         * @param {conversation.IMessageAck} request MessageAck message or plain object
         * @returns {Promise<common.Empty>} Promise
         * @variation 2
         */

        /**
         * Callback as used by {@link conversation.ConversationMgr#startConversation}.
         * @memberof conversation.ConversationMgr
         * @typedef StartConversationCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {conversation.Conversation} [response] Conversation
         */

        /**
         * Calls StartConversation.
         * @function startConversation
         * @memberof conversation.ConversationMgr
         * @instance
         * @param {conversation.IStartRequest} request StartRequest message or plain object
         * @param {conversation.ConversationMgr.StartConversationCallback} callback Node-style callback called with the error, if any, and Conversation
         * @returns {undefined}
         * @variation 1
         */
        ConversationMgr.prototype.startConversation = function startConversation(request, callback) {
            return this.rpcCall(startConversation, $root.conversation.StartRequest, $root.conversation.Conversation, request, callback);
        };

        /**
         * Calls StartConversation.
         * @function startConversation
         * @memberof conversation.ConversationMgr
         * @instance
         * @param {conversation.IStartRequest} request StartRequest message or plain object
         * @returns {Promise<conversation.Conversation>} Promise
         * @variation 2
         */

        /**
         * Callback as used by {@link conversation.ConversationMgr#endConversation}.
         * @memberof conversation.ConversationMgr
         * @typedef EndConversationCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {conversation.Conversation} [response] Conversation
         */

        /**
         * Calls EndConversation.
         * @function endConversation
         * @memberof conversation.ConversationMgr
         * @instance
         * @param {common.IId} request Id message or plain object
         * @param {conversation.ConversationMgr.EndConversationCallback} callback Node-style callback called with the error, if any, and Conversation
         * @returns {undefined}
         * @variation 1
         */
        ConversationMgr.prototype.endConversation = function endConversation(request, callback) {
            return this.rpcCall(endConversation, $root.common.Id, $root.conversation.Conversation, request, callback);
        };

        /**
         * Calls EndConversation.
         * @function endConversation
         * @memberof conversation.ConversationMgr
         * @instance
         * @param {common.IId} request Id message or plain object
         * @returns {Promise<conversation.Conversation>} Promise
         * @variation 2
         */

        /**
         * Callback as used by {@link conversation.ConversationMgr#getConversation}.
         * @memberof conversation.ConversationMgr
         * @typedef GetConversationCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {conversation.Conversation} [response] Conversation
         */

        /**
         * Calls GetConversation.
         * @function getConversation
         * @memberof conversation.ConversationMgr
         * @instance
         * @param {common.IId} request Id message or plain object
         * @param {conversation.ConversationMgr.GetConversationCallback} callback Node-style callback called with the error, if any, and Conversation
         * @returns {undefined}
         * @variation 1
         */
        ConversationMgr.prototype.getConversation = function getConversation(request, callback) {
            return this.rpcCall(getConversation, $root.common.Id, $root.conversation.Conversation, request, callback);
        };

        /**
         * Calls GetConversation.
         * @function getConversation
         * @memberof conversation.ConversationMgr
         * @instance
         * @param {common.IId} request Id message or plain object
         * @returns {Promise<conversation.Conversation>} Promise
         * @variation 2
         */

        /**
         * Callback as used by {@link conversation.ConversationMgr#listConversations}.
         * @memberof conversation.ConversationMgr
         * @typedef ListConversationsCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {conversation.Conversations} [response] Conversations
         */

        /**
         * Calls ListConversations.
         * @function listConversations
         * @memberof conversation.ConversationMgr
         * @instance
         * @param {conversation.IListConversationsRequest} request ListConversationsRequest message or plain object
         * @param {conversation.ConversationMgr.ListConversationsCallback} callback Node-style callback called with the error, if any, and Conversations
         * @returns {undefined}
         * @variation 1
         */
        ConversationMgr.prototype.listConversations = function listConversations(request, callback) {
            return this.rpcCall(listConversations, $root.conversation.ListConversationsRequest, $root.conversation.Conversations, request, callback);
        };

        /**
         * Calls ListConversations.
         * @function listConversations
         * @memberof conversation.ConversationMgr
         * @instance
         * @param {conversation.IListConversationsRequest} request ListConversationsRequest message or plain object
         * @returns {Promise<conversation.Conversations>} Promise
         * @variation 2
         */

        /**
         * Callback as used by {@link conversation.ConversationMgr#acceptConversation}.
         * @memberof conversation.ConversationMgr
         * @typedef AcceptConversationCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {conversation.Conversation} [response] Conversation
         */

        /**
         * Calls AcceptConversation.
         * @function acceptConversation
         * @memberof conversation.ConversationMgr
         * @instance
         * @param {common.IId} request Id message or plain object
         * @param {conversation.ConversationMgr.AcceptConversationCallback} callback Node-style callback called with the error, if any, and Conversation
         * @returns {undefined}
         * @variation 1
         */
        ConversationMgr.prototype.acceptConversation = function acceptConversation(request, callback) {
            return this.rpcCall(acceptConversation, $root.common.Id, $root.conversation.Conversation, request, callback);
        };

        /**
         * Calls AcceptConversation.
         * @function acceptConversation
         * @memberof conversation.ConversationMgr
         * @instance
         * @param {common.IId} request Id message or plain object
         * @returns {Promise<conversation.Conversation>} Promise
         * @variation 2
         */

        /**
         * Callback as used by {@link conversation.ConversationMgr#tagConversation}.
         * @memberof conversation.ConversationMgr
         * @typedef TagConversationCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {common.Empty} [response] Empty
         */

        /**
         * Calls TagConversation.
         * @function tagConversation
         * @memberof conversation.ConversationMgr
         * @instance
         * @param {conversation.ITagRequest} request TagRequest message or plain object
         * @param {conversation.ConversationMgr.TagConversationCallback} callback Node-style callback called with the error, if any, and Empty
         * @returns {undefined}
         * @variation 1
         */
        ConversationMgr.prototype.tagConversation = function tagConversation(request, callback) {
            return this.rpcCall(tagConversation, $root.conversation.TagRequest, $root.common.Empty, request, callback);
        };

        /**
         * Calls TagConversation.
         * @function tagConversation
         * @memberof conversation.ConversationMgr
         * @instance
         * @param {conversation.ITagRequest} request TagRequest message or plain object
         * @returns {Promise<common.Empty>} Promise
         * @variation 2
         */

        /**
         * Callback as used by {@link conversation.ConversationMgr#untagConversation}.
         * @memberof conversation.ConversationMgr
         * @typedef UntagConversationCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {common.Empty} [response] Empty
         */

        /**
         * Calls UntagConversation.
         * @function untagConversation
         * @memberof conversation.ConversationMgr
         * @instance
         * @param {conversation.ITagRequest} request TagRequest message or plain object
         * @param {conversation.ConversationMgr.UntagConversationCallback} callback Node-style callback called with the error, if any, and Empty
         * @returns {undefined}
         * @variation 1
         */
        ConversationMgr.prototype.untagConversation = function untagConversation(request, callback) {
            return this.rpcCall(untagConversation, $root.conversation.TagRequest, $root.common.Empty, request, callback);
        };

        /**
         * Calls UntagConversation.
         * @function untagConversation
         * @memberof conversation.ConversationMgr
         * @instance
         * @param {conversation.ITagRequest} request TagRequest message or plain object
         * @returns {Promise<common.Empty>} Promise
         * @variation 2
         */

        /**
         * Callback as used by {@link conversation.ConversationMgr#joinConversation}.
         * @memberof conversation.ConversationMgr
         * @typedef JoinConversationCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {common.Empty} [response] Empty
         */

        /**
         * Calls JoinConversation.
         * @function joinConversation
         * @memberof conversation.ConversationMgr
         * @instance
         * @param {conversation.IMember} request Member message or plain object
         * @param {conversation.ConversationMgr.JoinConversationCallback} callback Node-style callback called with the error, if any, and Empty
         * @returns {undefined}
         * @variation 1
         */
        ConversationMgr.prototype.joinConversation = function joinConversation(request, callback) {
            return this.rpcCall(joinConversation, $root.conversation.Member, $root.common.Empty, request, callback);
        };

        /**
         * Calls JoinConversation.
         * @function joinConversation
         * @memberof conversation.ConversationMgr
         * @instance
         * @param {conversation.IMember} request Member message or plain object
         * @returns {Promise<common.Empty>} Promise
         * @variation 2
         */

        /**
         * Callback as used by {@link conversation.ConversationMgr#leftConversation}.
         * @memberof conversation.ConversationMgr
         * @typedef LeftConversationCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {common.Empty} [response] Empty
         */

        /**
         * Calls LeftConversation.
         * @function leftConversation
         * @memberof conversation.ConversationMgr
         * @instance
         * @param {conversation.IMember} request Member message or plain object
         * @param {conversation.ConversationMgr.LeftConversationCallback} callback Node-style callback called with the error, if any, and Empty
         * @returns {undefined}
         * @variation 1
         */
        ConversationMgr.prototype.leftConversation = function leftConversation(request, callback) {
            return this.rpcCall(leftConversation, $root.conversation.Member, $root.common.Empty, request, callback);
        };

        /**
         * Calls LeftConversation.
         * @function leftConversation
         * @memberof conversation.ConversationMgr
         * @instance
         * @param {conversation.IMember} request Member message or plain object
         * @returns {Promise<common.Empty>} Promise
         * @variation 2
         */

        /**
         * Callback as used by {@link conversation.ConversationMgr#typing}.
         * @memberof conversation.ConversationMgr
         * @typedef TypingCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {common.Empty} [response] Empty
         */

        /**
         * Calls Typing.
         * @function typing
         * @memberof conversation.ConversationMgr
         * @instance
         * @param {common.IId} request Id message or plain object
         * @param {conversation.ConversationMgr.TypingCallback} callback Node-style callback called with the error, if any, and Empty
         * @returns {undefined}
         * @variation 1
         */
        ConversationMgr.prototype.typing = function typing(request, callback) {
            return this.rpcCall(typing, $root.common.Id, $root.common.Empty, request, callback);
        };

        /**
         * Calls Typing.
         * @function typing
         * @memberof conversation.ConversationMgr
         * @instance
         * @param {common.IId} request Id message or plain object
         * @returns {Promise<common.Empty>} Promise
         * @variation 2
         */

        return ConversationMgr;
    })();

    conversation.CannedResponseMgr = (function() {

        /**
         * Constructs a new CannedResponseMgr service.
         * @memberof conversation
         * @classdesc Represents a CannedResponseMgr
         * @extends $protobuf.rpc.Service
         * @constructor
         * @param {$protobuf.RPCImpl} rpcImpl RPC implementation
         * @param {boolean} [requestDelimited=false] Whether requests are length-delimited
         * @param {boolean} [responseDelimited=false] Whether responses are length-delimited
         */
        function CannedResponseMgr(rpcImpl, requestDelimited, responseDelimited) {
            $protobuf.rpc.Service.call(this, rpcImpl, requestDelimited, responseDelimited);
        }

        (CannedResponseMgr.prototype = Object.create($protobuf.rpc.Service.prototype)).constructor = CannedResponseMgr;

        /**
         * Callback as used by {@link conversation.CannedResponseMgr#createCannedResponse}.
         * @memberof conversation.CannedResponseMgr
         * @typedef CreateCannedResponseCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {conversation.CannedResponse} [response] CannedResponse
         */

        /**
         * Calls CreateCannedResponse.
         * @function createCannedResponse
         * @memberof conversation.CannedResponseMgr
         * @instance
         * @param {conversation.ICannedResponse} request CannedResponse message or plain object
         * @param {conversation.CannedResponseMgr.CreateCannedResponseCallback} callback Node-style callback called with the error, if any, and CannedResponse
         * @returns {undefined}
         * @variation 1
         */
        CannedResponseMgr.prototype.createCannedResponse = function createCannedResponse(request, callback) {
            return this.rpcCall(createCannedResponse, $root.conversation.CannedResponse, $root.conversation.CannedResponse, request, callback);
        };

        /**
         * Calls CreateCannedResponse.
         * @function createCannedResponse
         * @memberof conversation.CannedResponseMgr
         * @instance
         * @param {conversation.ICannedResponse} request CannedResponse message or plain object
         * @returns {Promise<conversation.CannedResponse>} Promise
         * @variation 2
         */

        /**
         * Callback as used by {@link conversation.CannedResponseMgr#updateCannedResponse}.
         * @memberof conversation.CannedResponseMgr
         * @typedef UpdateCannedResponseCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {conversation.CannedResponse} [response] CannedResponse
         */

        /**
         * Calls UpdateCannedResponse.
         * @function updateCannedResponse
         * @memberof conversation.CannedResponseMgr
         * @instance
         * @param {conversation.ICannedResponse} request CannedResponse message or plain object
         * @param {conversation.CannedResponseMgr.UpdateCannedResponseCallback} callback Node-style callback called with the error, if any, and CannedResponse
         * @returns {undefined}
         * @variation 1
         */
        CannedResponseMgr.prototype.updateCannedResponse = function updateCannedResponse(request, callback) {
            return this.rpcCall(updateCannedResponse, $root.conversation.CannedResponse, $root.conversation.CannedResponse, request, callback);
        };

        /**
         * Calls UpdateCannedResponse.
         * @function updateCannedResponse
         * @memberof conversation.CannedResponseMgr
         * @instance
         * @param {conversation.ICannedResponse} request CannedResponse message or plain object
         * @returns {Promise<conversation.CannedResponse>} Promise
         * @variation 2
         */

        /**
         * Callback as used by {@link conversation.CannedResponseMgr#listCannedResponses}.
         * @memberof conversation.CannedResponseMgr
         * @typedef ListCannedResponsesCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {conversation.CannedResponses} [response] CannedResponses
         */

        /**
         * Calls ListCannedResponses.
         * @function listCannedResponses
         * @memberof conversation.CannedResponseMgr
         * @instance
         * @param {common.IId} request Id message or plain object
         * @param {conversation.CannedResponseMgr.ListCannedResponsesCallback} callback Node-style callback called with the error, if any, and CannedResponses
         * @returns {undefined}
         * @variation 1
         */
        CannedResponseMgr.prototype.listCannedResponses = function listCannedResponses(request, callback) {
            return this.rpcCall(listCannedResponses, $root.common.Id, $root.conversation.CannedResponses, request, callback);
        };

        /**
         * Calls ListCannedResponses.
         * @function listCannedResponses
         * @memberof conversation.CannedResponseMgr
         * @instance
         * @param {common.IId} request Id message or plain object
         * @returns {Promise<conversation.CannedResponses>} Promise
         * @variation 2
         */

        /**
         * Callback as used by {@link conversation.CannedResponseMgr#deleteCannedResponse}.
         * @memberof conversation.CannedResponseMgr
         * @typedef DeleteCannedResponseCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {common.Empty} [response] Empty
         */

        /**
         * Calls DeleteCannedResponse.
         * @function deleteCannedResponse
         * @memberof conversation.CannedResponseMgr
         * @instance
         * @param {common.IId} request Id message or plain object
         * @param {conversation.CannedResponseMgr.DeleteCannedResponseCallback} callback Node-style callback called with the error, if any, and Empty
         * @returns {undefined}
         * @variation 1
         */
        CannedResponseMgr.prototype.deleteCannedResponse = function deleteCannedResponse(request, callback) {
            return this.rpcCall(deleteCannedResponse, $root.common.Id, $root.common.Empty, request, callback);
        };

        /**
         * Calls DeleteCannedResponse.
         * @function deleteCannedResponse
         * @memberof conversation.CannedResponseMgr
         * @instance
         * @param {common.IId} request Id message or plain object
         * @returns {Promise<common.Empty>} Promise
         * @variation 2
         */

        /**
         * Callback as used by {@link conversation.CannedResponseMgr#getCannedResponse}.
         * @memberof conversation.CannedResponseMgr
         * @typedef GetCannedResponseCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {conversation.CannedResponse} [response] CannedResponse
         */

        /**
         * Calls GetCannedResponse.
         * @function getCannedResponse
         * @memberof conversation.CannedResponseMgr
         * @instance
         * @param {common.IId} request Id message or plain object
         * @param {conversation.CannedResponseMgr.GetCannedResponseCallback} callback Node-style callback called with the error, if any, and CannedResponse
         * @returns {undefined}
         * @variation 1
         */
        CannedResponseMgr.prototype.getCannedResponse = function getCannedResponse(request, callback) {
            return this.rpcCall(getCannedResponse, $root.common.Id, $root.conversation.CannedResponse, request, callback);
        };

        /**
         * Calls GetCannedResponse.
         * @function getCannedResponse
         * @memberof conversation.CannedResponseMgr
         * @instance
         * @param {common.IId} request Id message or plain object
         * @returns {Promise<conversation.CannedResponse>} Promise
         * @variation 2
         */

        return CannedResponseMgr;
    })();

    conversation.TagMgr = (function() {

        /**
         * Constructs a new TagMgr service.
         * @memberof conversation
         * @classdesc Represents a TagMgr
         * @extends $protobuf.rpc.Service
         * @constructor
         * @param {$protobuf.RPCImpl} rpcImpl RPC implementation
         * @param {boolean} [requestDelimited=false] Whether requests are length-delimited
         * @param {boolean} [responseDelimited=false] Whether responses are length-delimited
         */
        function TagMgr(rpcImpl, requestDelimited, responseDelimited) {
            $protobuf.rpc.Service.call(this, rpcImpl, requestDelimited, responseDelimited);
        }

        (TagMgr.prototype = Object.create($protobuf.rpc.Service.prototype)).constructor = TagMgr;

        /**
         * Callback as used by {@link conversation.TagMgr#createTag}.
         * @memberof conversation.TagMgr
         * @typedef CreateTagCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {conversation.Tag} [response] Tag
         */

        /**
         * Calls CreateTag.
         * @function createTag
         * @memberof conversation.TagMgr
         * @instance
         * @param {conversation.ITag} request Tag message or plain object
         * @param {conversation.TagMgr.CreateTagCallback} callback Node-style callback called with the error, if any, and Tag
         * @returns {undefined}
         * @variation 1
         */
        TagMgr.prototype.createTag = function createTag(request, callback) {
            return this.rpcCall(createTag, $root.conversation.Tag, $root.conversation.Tag, request, callback);
        };

        /**
         * Calls CreateTag.
         * @function createTag
         * @memberof conversation.TagMgr
         * @instance
         * @param {conversation.ITag} request Tag message or plain object
         * @returns {Promise<conversation.Tag>} Promise
         * @variation 2
         */

        /**
         * Callback as used by {@link conversation.TagMgr#updateTag}.
         * @memberof conversation.TagMgr
         * @typedef UpdateTagCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {conversation.Tag} [response] Tag
         */

        /**
         * Calls UpdateTag.
         * @function updateTag
         * @memberof conversation.TagMgr
         * @instance
         * @param {conversation.ITag} request Tag message or plain object
         * @param {conversation.TagMgr.UpdateTagCallback} callback Node-style callback called with the error, if any, and Tag
         * @returns {undefined}
         * @variation 1
         */
        TagMgr.prototype.updateTag = function updateTag(request, callback) {
            return this.rpcCall(updateTag, $root.conversation.Tag, $root.conversation.Tag, request, callback);
        };

        /**
         * Calls UpdateTag.
         * @function updateTag
         * @memberof conversation.TagMgr
         * @instance
         * @param {conversation.ITag} request Tag message or plain object
         * @returns {Promise<conversation.Tag>} Promise
         * @variation 2
         */

        /**
         * Callback as used by {@link conversation.TagMgr#deleteTag}.
         * @memberof conversation.TagMgr
         * @typedef DeleteTagCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {common.Empty} [response] Empty
         */

        /**
         * Calls DeleteTag.
         * @function deleteTag
         * @memberof conversation.TagMgr
         * @instance
         * @param {common.IId} request Id message or plain object
         * @param {conversation.TagMgr.DeleteTagCallback} callback Node-style callback called with the error, if any, and Empty
         * @returns {undefined}
         * @variation 1
         */
        TagMgr.prototype.deleteTag = function deleteTag(request, callback) {
            return this.rpcCall(deleteTag, $root.common.Id, $root.common.Empty, request, callback);
        };

        /**
         * Calls DeleteTag.
         * @function deleteTag
         * @memberof conversation.TagMgr
         * @instance
         * @param {common.IId} request Id message or plain object
         * @returns {Promise<common.Empty>} Promise
         * @variation 2
         */

        /**
         * Callback as used by {@link conversation.TagMgr#listTags}.
         * @memberof conversation.TagMgr
         * @typedef ListTagsCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {conversation.Tags} [response] Tags
         */

        /**
         * Calls ListTags.
         * @function listTags
         * @memberof conversation.TagMgr
         * @instance
         * @param {common.IId} request Id message or plain object
         * @param {conversation.TagMgr.ListTagsCallback} callback Node-style callback called with the error, if any, and Tags
         * @returns {undefined}
         * @variation 1
         */
        TagMgr.prototype.listTags = function listTags(request, callback) {
            return this.rpcCall(listTags, $root.common.Id, $root.conversation.Tags, request, callback);
        };

        /**
         * Calls ListTags.
         * @function listTags
         * @memberof conversation.TagMgr
         * @instance
         * @param {common.IId} request Id message or plain object
         * @returns {Promise<conversation.Tags>} Promise
         * @variation 2
         */

        /**
         * Callback as used by {@link conversation.TagMgr#getTag}.
         * @memberof conversation.TagMgr
         * @typedef GetTagCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {conversation.Tag} [response] Tag
         */

        /**
         * Calls GetTag.
         * @function getTag
         * @memberof conversation.TagMgr
         * @instance
         * @param {common.IId} request Id message or plain object
         * @param {conversation.TagMgr.GetTagCallback} callback Node-style callback called with the error, if any, and Tag
         * @returns {undefined}
         * @variation 1
         */
        TagMgr.prototype.getTag = function getTag(request, callback) {
            return this.rpcCall(getTag, $root.common.Id, $root.conversation.Tag, request, callback);
        };

        /**
         * Calls GetTag.
         * @function getTag
         * @memberof conversation.TagMgr
         * @instance
         * @param {common.IId} request Id message or plain object
         * @returns {Promise<conversation.Tag>} Promise
         * @variation 2
         */

        return TagMgr;
    })();

    conversation.IntegrationMgr = (function() {

        /**
         * Constructs a new IntegrationMgr service.
         * @memberof conversation
         * @classdesc Represents an IntegrationMgr
         * @extends $protobuf.rpc.Service
         * @constructor
         * @param {$protobuf.RPCImpl} rpcImpl RPC implementation
         * @param {boolean} [requestDelimited=false] Whether requests are length-delimited
         * @param {boolean} [responseDelimited=false] Whether responses are length-delimited
         */
        function IntegrationMgr(rpcImpl, requestDelimited, responseDelimited) {
            $protobuf.rpc.Service.call(this, rpcImpl, requestDelimited, responseDelimited);
        }

        (IntegrationMgr.prototype = Object.create($protobuf.rpc.Service.prototype)).constructor = IntegrationMgr;

        /**
         * Callback as used by {@link conversation.IntegrationMgr#deintegrate}.
         * @memberof conversation.IntegrationMgr
         * @typedef DeintegrateCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {common.Empty} [response] Empty
         */

        /**
         * Calls Deintegrate.
         * @function deintegrate
         * @memberof conversation.IntegrationMgr
         * @instance
         * @param {common.IId} request Id message or plain object
         * @param {conversation.IntegrationMgr.DeintegrateCallback} callback Node-style callback called with the error, if any, and Empty
         * @returns {undefined}
         * @variation 1
         */
        IntegrationMgr.prototype.deintegrate = function deintegrate(request, callback) {
            return this.rpcCall(deintegrate, $root.common.Id, $root.common.Empty, request, callback);
        };

        /**
         * Calls Deintegrate.
         * @function deintegrate
         * @memberof conversation.IntegrationMgr
         * @instance
         * @param {common.IId} request Id message or plain object
         * @returns {Promise<common.Empty>} Promise
         * @variation 2
         */

        /**
         * Callback as used by {@link conversation.IntegrationMgr#listIntegrations}.
         * @memberof conversation.IntegrationMgr
         * @typedef ListIntegrationsCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {conversation.Integrations} [response] Integrations
         */

        /**
         * Calls ListIntegrations.
         * @function listIntegrations
         * @memberof conversation.IntegrationMgr
         * @instance
         * @param {common.IId} request Id message or plain object
         * @param {conversation.IntegrationMgr.ListIntegrationsCallback} callback Node-style callback called with the error, if any, and Integrations
         * @returns {undefined}
         * @variation 1
         */
        IntegrationMgr.prototype.listIntegrations = function listIntegrations(request, callback) {
            return this.rpcCall(listIntegrations, $root.common.Id, $root.conversation.Integrations, request, callback);
        };

        /**
         * Calls ListIntegrations.
         * @function listIntegrations
         * @memberof conversation.IntegrationMgr
         * @instance
         * @param {common.IId} request Id message or plain object
         * @returns {Promise<conversation.Integrations>} Promise
         * @variation 2
         */

        /**
         * Callback as used by {@link conversation.IntegrationMgr#integrate}.
         * @memberof conversation.IntegrationMgr
         * @typedef IntegrateCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {conversation.Integration} [response] Integration
         */

        /**
         * Calls Integrate.
         * @function integrate
         * @memberof conversation.IntegrationMgr
         * @instance
         * @param {conversation.IIntegration} request Integration message or plain object
         * @param {conversation.IntegrationMgr.IntegrateCallback} callback Node-style callback called with the error, if any, and Integration
         * @returns {undefined}
         * @variation 1
         */
        IntegrationMgr.prototype.integrate = function integrate(request, callback) {
            return this.rpcCall(integrate, $root.conversation.Integration, $root.conversation.Integration, request, callback);
        };

        /**
         * Calls Integrate.
         * @function integrate
         * @memberof conversation.IntegrationMgr
         * @instance
         * @param {conversation.IIntegration} request Integration message or plain object
         * @returns {Promise<conversation.Integration>} Promise
         * @variation 2
         */

        /**
         * Callback as used by {@link conversation.IntegrationMgr#checkAvailability}.
         * @memberof conversation.IntegrationMgr
         * @typedef CheckAvailabilityCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {conversation.AvailabilityCheckResult} [response] AvailabilityCheckResult
         */

        /**
         * Calls CheckAvailability.
         * @function checkAvailability
         * @memberof conversation.IntegrationMgr
         * @instance
         * @param {conversation.IAvailabilityCheckRequest} request AvailabilityCheckRequest message or plain object
         * @param {conversation.IntegrationMgr.CheckAvailabilityCallback} callback Node-style callback called with the error, if any, and AvailabilityCheckResult
         * @returns {undefined}
         * @variation 1
         */
        IntegrationMgr.prototype.checkAvailability = function checkAvailability(request, callback) {
            return this.rpcCall(checkAvailability, $root.conversation.AvailabilityCheckRequest, $root.conversation.AvailabilityCheckResult, request, callback);
        };

        /**
         * Calls CheckAvailability.
         * @function checkAvailability
         * @memberof conversation.IntegrationMgr
         * @instance
         * @param {conversation.IAvailabilityCheckRequest} request AvailabilityCheckRequest message or plain object
         * @returns {Promise<conversation.AvailabilityCheckResult>} Promise
         * @variation 2
         */

        return IntegrationMgr;
    })();

    conversation.TagRequest = (function() {

        /**
         * Properties of a TagRequest.
         * @memberof conversation
         * @interface ITagRequest
         * @property {common.IContext|null} [ctx] TagRequest ctx
         * @property {string|null} [account_id] TagRequest account_id
         * @property {string|null} [conversation_id] TagRequest conversation_id
         * @property {string|null} [id] TagRequest id
         */

        /**
         * Constructs a new TagRequest.
         * @memberof conversation
         * @classdesc Represents a TagRequest.
         * @implements ITagRequest
         * @constructor
         * @param {conversation.ITagRequest=} [p] Properties to set
         */
        function TagRequest(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * TagRequest ctx.
         * @member {common.IContext|null|undefined} ctx
         * @memberof conversation.TagRequest
         * @instance
         */
        TagRequest.prototype.ctx = null;

        /**
         * TagRequest account_id.
         * @member {string} account_id
         * @memberof conversation.TagRequest
         * @instance
         */
        TagRequest.prototype.account_id = "";

        /**
         * TagRequest conversation_id.
         * @member {string} conversation_id
         * @memberof conversation.TagRequest
         * @instance
         */
        TagRequest.prototype.conversation_id = "";

        /**
         * TagRequest id.
         * @member {string} id
         * @memberof conversation.TagRequest
         * @instance
         */
        TagRequest.prototype.id = "";

        return TagRequest;
    })();

    conversation.CountByAgentRequest = (function() {

        /**
         * Properties of a CountByAgentRequest.
         * @memberof conversation
         * @interface ICountByAgentRequest
         * @property {string|null} [integration_id] CountByAgentRequest integration_id
         * @property {string|null} [agent_id] CountByAgentRequest agent_id
         * @property {number|null} [from] CountByAgentRequest from
         * @property {number|null} [to] CountByAgentRequest to
         * @property {string|null} [range] CountByAgentRequest range
         */

        /**
         * Constructs a new CountByAgentRequest.
         * @memberof conversation
         * @classdesc Represents a CountByAgentRequest.
         * @implements ICountByAgentRequest
         * @constructor
         * @param {conversation.ICountByAgentRequest=} [p] Properties to set
         */
        function CountByAgentRequest(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * CountByAgentRequest integration_id.
         * @member {string} integration_id
         * @memberof conversation.CountByAgentRequest
         * @instance
         */
        CountByAgentRequest.prototype.integration_id = "";

        /**
         * CountByAgentRequest agent_id.
         * @member {string} agent_id
         * @memberof conversation.CountByAgentRequest
         * @instance
         */
        CountByAgentRequest.prototype.agent_id = "";

        /**
         * CountByAgentRequest from.
         * @member {number} from
         * @memberof conversation.CountByAgentRequest
         * @instance
         */
        CountByAgentRequest.prototype.from = 0;

        /**
         * CountByAgentRequest to.
         * @member {number} to
         * @memberof conversation.CountByAgentRequest
         * @instance
         */
        CountByAgentRequest.prototype.to = 0;

        /**
         * CountByAgentRequest range.
         * @member {string} range
         * @memberof conversation.CountByAgentRequest
         * @instance
         */
        CountByAgentRequest.prototype.range = "";

        /**
         * Range enum.
         * @name conversation.CountByAgentRequest.Range
         * @enum {string}
         * @property {number} hour=0 hour value
         * @property {number} day=1 day value
         */
        CountByAgentRequest.Range = (function() {
            const valuesById = {}, values = Object.create(valuesById);
            values[valuesById[0] = "hour"] = 0;
            values[valuesById[1] = "day"] = 1;
            return values;
        })();

        return CountByAgentRequest;
    })();

    conversation.CountByGroupRequest = (function() {

        /**
         * Properties of a CountByGroupRequest.
         * @memberof conversation
         * @interface ICountByGroupRequest
         * @property {string|null} [integration_id] CountByGroupRequest integration_id
         * @property {string|null} [group_id] CountByGroupRequest group_id
         * @property {number|null} [from] CountByGroupRequest from
         * @property {number|null} [to] CountByGroupRequest to
         * @property {string|null} [range] CountByGroupRequest range
         */

        /**
         * Constructs a new CountByGroupRequest.
         * @memberof conversation
         * @classdesc Represents a CountByGroupRequest.
         * @implements ICountByGroupRequest
         * @constructor
         * @param {conversation.ICountByGroupRequest=} [p] Properties to set
         */
        function CountByGroupRequest(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * CountByGroupRequest integration_id.
         * @member {string} integration_id
         * @memberof conversation.CountByGroupRequest
         * @instance
         */
        CountByGroupRequest.prototype.integration_id = "";

        /**
         * CountByGroupRequest group_id.
         * @member {string} group_id
         * @memberof conversation.CountByGroupRequest
         * @instance
         */
        CountByGroupRequest.prototype.group_id = "";

        /**
         * CountByGroupRequest from.
         * @member {number} from
         * @memberof conversation.CountByGroupRequest
         * @instance
         */
        CountByGroupRequest.prototype.from = 0;

        /**
         * CountByGroupRequest to.
         * @member {number} to
         * @memberof conversation.CountByGroupRequest
         * @instance
         */
        CountByGroupRequest.prototype.to = 0;

        /**
         * CountByGroupRequest range.
         * @member {string} range
         * @memberof conversation.CountByGroupRequest
         * @instance
         */
        CountByGroupRequest.prototype.range = "";

        /**
         * Range enum.
         * @name conversation.CountByGroupRequest.Range
         * @enum {string}
         * @property {number} hour=0 hour value
         * @property {number} day=1 day value
         */
        CountByGroupRequest.Range = (function() {
            const valuesById = {}, values = Object.create(valuesById);
            values[valuesById[0] = "hour"] = 0;
            values[valuesById[1] = "day"] = 1;
            return values;
        })();

        return CountByGroupRequest;
    })();

    conversation.CountByAgentsRequest = (function() {

        /**
         * Properties of a CountByAgentsRequest.
         * @memberof conversation
         * @interface ICountByAgentsRequest
         * @property {string|null} [integration_id] CountByAgentsRequest integration_id
         * @property {Array.<string>|null} [agent_ids] CountByAgentsRequest agent_ids
         * @property {number|null} [from] CountByAgentsRequest from
         * @property {number|null} [to] CountByAgentsRequest to
         * @property {string|null} [range] CountByAgentsRequest range
         */

        /**
         * Constructs a new CountByAgentsRequest.
         * @memberof conversation
         * @classdesc Represents a CountByAgentsRequest.
         * @implements ICountByAgentsRequest
         * @constructor
         * @param {conversation.ICountByAgentsRequest=} [p] Properties to set
         */
        function CountByAgentsRequest(p) {
            this.agent_ids = [];
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * CountByAgentsRequest integration_id.
         * @member {string} integration_id
         * @memberof conversation.CountByAgentsRequest
         * @instance
         */
        CountByAgentsRequest.prototype.integration_id = "";

        /**
         * CountByAgentsRequest agent_ids.
         * @member {Array.<string>} agent_ids
         * @memberof conversation.CountByAgentsRequest
         * @instance
         */
        CountByAgentsRequest.prototype.agent_ids = $util.emptyArray;

        /**
         * CountByAgentsRequest from.
         * @member {number} from
         * @memberof conversation.CountByAgentsRequest
         * @instance
         */
        CountByAgentsRequest.prototype.from = 0;

        /**
         * CountByAgentsRequest to.
         * @member {number} to
         * @memberof conversation.CountByAgentsRequest
         * @instance
         */
        CountByAgentsRequest.prototype.to = 0;

        /**
         * CountByAgentsRequest range.
         * @member {string} range
         * @memberof conversation.CountByAgentsRequest
         * @instance
         */
        CountByAgentsRequest.prototype.range = "";

        /**
         * Range enum.
         * @name conversation.CountByAgentsRequest.Range
         * @enum {string}
         * @property {number} hour=0 hour value
         * @property {number} day=1 day value
         */
        CountByAgentsRequest.Range = (function() {
            const valuesById = {}, values = Object.create(valuesById);
            values[valuesById[0] = "hour"] = 0;
            values[valuesById[1] = "day"] = 1;
            return values;
        })();

        return CountByAgentsRequest;
    })();

    conversation.CountByTagsRequest = (function() {

        /**
         * Properties of a CountByTagsRequest.
         * @memberof conversation
         * @interface ICountByTagsRequest
         * @property {string|null} [integration_id] CountByTagsRequest integration_id
         * @property {Array.<string>|null} [tag_ids] CountByTagsRequest tag_ids
         * @property {number|null} [from] CountByTagsRequest from
         * @property {number|null} [to] CountByTagsRequest to
         * @property {string|null} [range] CountByTagsRequest range
         */

        /**
         * Constructs a new CountByTagsRequest.
         * @memberof conversation
         * @classdesc Represents a CountByTagsRequest.
         * @implements ICountByTagsRequest
         * @constructor
         * @param {conversation.ICountByTagsRequest=} [p] Properties to set
         */
        function CountByTagsRequest(p) {
            this.tag_ids = [];
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * CountByTagsRequest integration_id.
         * @member {string} integration_id
         * @memberof conversation.CountByTagsRequest
         * @instance
         */
        CountByTagsRequest.prototype.integration_id = "";

        /**
         * CountByTagsRequest tag_ids.
         * @member {Array.<string>} tag_ids
         * @memberof conversation.CountByTagsRequest
         * @instance
         */
        CountByTagsRequest.prototype.tag_ids = $util.emptyArray;

        /**
         * CountByTagsRequest from.
         * @member {number} from
         * @memberof conversation.CountByTagsRequest
         * @instance
         */
        CountByTagsRequest.prototype.from = 0;

        /**
         * CountByTagsRequest to.
         * @member {number} to
         * @memberof conversation.CountByTagsRequest
         * @instance
         */
        CountByTagsRequest.prototype.to = 0;

        /**
         * CountByTagsRequest range.
         * @member {string} range
         * @memberof conversation.CountByTagsRequest
         * @instance
         */
        CountByTagsRequest.prototype.range = "";

        /**
         * Range enum.
         * @name conversation.CountByTagsRequest.Range
         * @enum {string}
         * @property {number} hour=0 hour value
         * @property {number} day=1 day value
         */
        CountByTagsRequest.Range = (function() {
            const valuesById = {}, values = Object.create(valuesById);
            values[valuesById[0] = "hour"] = 0;
            values[valuesById[1] = "day"] = 1;
            return values;
        })();

        return CountByTagsRequest;
    })();

    conversation.CountByAgent = (function() {

        /**
         * Properties of a CountByAgent.
         * @memberof conversation
         * @interface ICountByAgent
         * @property {string|null} [agent_id] CountByAgent agent_id
         * @property {Array.<number|Long>|null} [data] CountByAgent data
         */

        /**
         * Constructs a new CountByAgent.
         * @memberof conversation
         * @classdesc Represents a CountByAgent.
         * @implements ICountByAgent
         * @constructor
         * @param {conversation.ICountByAgent=} [p] Properties to set
         */
        function CountByAgent(p) {
            this.data = [];
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * CountByAgent agent_id.
         * @member {string} agent_id
         * @memberof conversation.CountByAgent
         * @instance
         */
        CountByAgent.prototype.agent_id = "";

        /**
         * CountByAgent data.
         * @member {Array.<number|Long>} data
         * @memberof conversation.CountByAgent
         * @instance
         */
        CountByAgent.prototype.data = $util.emptyArray;

        return CountByAgent;
    })();

    conversation.CountByGroup = (function() {

        /**
         * Properties of a CountByGroup.
         * @memberof conversation
         * @interface ICountByGroup
         * @property {string|null} [group_id] CountByGroup group_id
         * @property {Array.<number|Long>|null} [data] CountByGroup data
         */

        /**
         * Constructs a new CountByGroup.
         * @memberof conversation
         * @classdesc Represents a CountByGroup.
         * @implements ICountByGroup
         * @constructor
         * @param {conversation.ICountByGroup=} [p] Properties to set
         */
        function CountByGroup(p) {
            this.data = [];
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * CountByGroup group_id.
         * @member {string} group_id
         * @memberof conversation.CountByGroup
         * @instance
         */
        CountByGroup.prototype.group_id = "";

        /**
         * CountByGroup data.
         * @member {Array.<number|Long>} data
         * @memberof conversation.CountByGroup
         * @instance
         */
        CountByGroup.prototype.data = $util.emptyArray;

        return CountByGroup;
    })();

    conversation.CountByTag = (function() {

        /**
         * Properties of a CountByTag.
         * @memberof conversation
         * @interface ICountByTag
         * @property {string|null} [tag_id] CountByTag tag_id
         * @property {Array.<number|Long>|null} [data] CountByTag data
         */

        /**
         * Constructs a new CountByTag.
         * @memberof conversation
         * @classdesc Represents a CountByTag.
         * @implements ICountByTag
         * @constructor
         * @param {conversation.ICountByTag=} [p] Properties to set
         */
        function CountByTag(p) {
            this.data = [];
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * CountByTag tag_id.
         * @member {string} tag_id
         * @memberof conversation.CountByTag
         * @instance
         */
        CountByTag.prototype.tag_id = "";

        /**
         * CountByTag data.
         * @member {Array.<number|Long>} data
         * @memberof conversation.CountByTag
         * @instance
         */
        CountByTag.prototype.data = $util.emptyArray;

        return CountByTag;
    })();

    conversation.CountByTagsResponse = (function() {

        /**
         * Properties of a CountByTagsResponse.
         * @memberof conversation
         * @interface ICountByTagsResponse
         * @property {Array.<conversation.ICountByTag>|null} [data] CountByTagsResponse data
         */

        /**
         * Constructs a new CountByTagsResponse.
         * @memberof conversation
         * @classdesc Represents a CountByTagsResponse.
         * @implements ICountByTagsResponse
         * @constructor
         * @param {conversation.ICountByTagsResponse=} [p] Properties to set
         */
        function CountByTagsResponse(p) {
            this.data = [];
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * CountByTagsResponse data.
         * @member {Array.<conversation.ICountByTag>} data
         * @memberof conversation.CountByTagsResponse
         * @instance
         */
        CountByTagsResponse.prototype.data = $util.emptyArray;

        return CountByTagsResponse;
    })();

    conversation.CountByAgentsResponse = (function() {

        /**
         * Properties of a CountByAgentsResponse.
         * @memberof conversation
         * @interface ICountByAgentsResponse
         * @property {Array.<conversation.ICountByAgent>|null} [data] CountByAgentsResponse data
         */

        /**
         * Constructs a new CountByAgentsResponse.
         * @memberof conversation
         * @classdesc Represents a CountByAgentsResponse.
         * @implements ICountByAgentsResponse
         * @constructor
         * @param {conversation.ICountByAgentsResponse=} [p] Properties to set
         */
        function CountByAgentsResponse(p) {
            this.data = [];
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * CountByAgentsResponse data.
         * @member {Array.<conversation.ICountByAgent>} data
         * @memberof conversation.CountByAgentsResponse
         * @instance
         */
        CountByAgentsResponse.prototype.data = $util.emptyArray;

        return CountByAgentsResponse;
    })();

    conversation.AvgResponseTimeRequest = (function() {

        /**
         * Properties of an AvgResponseTimeRequest.
         * @memberof conversation
         * @interface IAvgResponseTimeRequest
         * @property {string|null} [integration_id] AvgResponseTimeRequest integration_id
         * @property {number|null} [from] AvgResponseTimeRequest from
         * @property {number|null} [to] AvgResponseTimeRequest to
         */

        /**
         * Constructs a new AvgResponseTimeRequest.
         * @memberof conversation
         * @classdesc Represents an AvgResponseTimeRequest.
         * @implements IAvgResponseTimeRequest
         * @constructor
         * @param {conversation.IAvgResponseTimeRequest=} [p] Properties to set
         */
        function AvgResponseTimeRequest(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * AvgResponseTimeRequest integration_id.
         * @member {string} integration_id
         * @memberof conversation.AvgResponseTimeRequest
         * @instance
         */
        AvgResponseTimeRequest.prototype.integration_id = "";

        /**
         * AvgResponseTimeRequest from.
         * @member {number} from
         * @memberof conversation.AvgResponseTimeRequest
         * @instance
         */
        AvgResponseTimeRequest.prototype.from = 0;

        /**
         * AvgResponseTimeRequest to.
         * @member {number} to
         * @memberof conversation.AvgResponseTimeRequest
         * @instance
         */
        AvgResponseTimeRequest.prototype.to = 0;

        return AvgResponseTimeRequest;
    })();

    conversation.AvgResponseTimeResponse = (function() {

        /**
         * Properties of an AvgResponseTimeResponse.
         * @memberof conversation
         * @interface IAvgResponseTimeResponse
         * @property {number|null} [avg_response_sec] AvgResponseTimeResponse avg_response_sec
         */

        /**
         * Constructs a new AvgResponseTimeResponse.
         * @memberof conversation
         * @classdesc Represents an AvgResponseTimeResponse.
         * @implements IAvgResponseTimeResponse
         * @constructor
         * @param {conversation.IAvgResponseTimeResponse=} [p] Properties to set
         */
        function AvgResponseTimeResponse(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * AvgResponseTimeResponse avg_response_sec.
         * @member {number} avg_response_sec
         * @memberof conversation.AvgResponseTimeResponse
         * @instance
         */
        AvgResponseTimeResponse.prototype.avg_response_sec = 0;

        return AvgResponseTimeResponse;
    })();

    conversation.TotalConversationResponse = (function() {

        /**
         * Properties of a TotalConversationResponse.
         * @memberof conversation
         * @interface ITotalConversationResponse
         * @property {number|Long|null} [total_conversation] TotalConversationResponse total_conversation
         */

        /**
         * Constructs a new TotalConversationResponse.
         * @memberof conversation
         * @classdesc Represents a TotalConversationResponse.
         * @implements ITotalConversationResponse
         * @constructor
         * @param {conversation.ITotalConversationResponse=} [p] Properties to set
         */
        function TotalConversationResponse(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * TotalConversationResponse total_conversation.
         * @member {number|Long} total_conversation
         * @memberof conversation.TotalConversationResponse
         * @instance
         */
        TotalConversationResponse.prototype.total_conversation = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

        return TotalConversationResponse;
    })();

    conversation.ConversationReporter = (function() {

        /**
         * Constructs a new ConversationReporter service.
         * @memberof conversation
         * @classdesc Represents a ConversationReporter
         * @extends $protobuf.rpc.Service
         * @constructor
         * @param {$protobuf.RPCImpl} rpcImpl RPC implementation
         * @param {boolean} [requestDelimited=false] Whether requests are length-delimited
         * @param {boolean} [responseDelimited=false] Whether responses are length-delimited
         */
        function ConversationReporter(rpcImpl, requestDelimited, responseDelimited) {
            $protobuf.rpc.Service.call(this, rpcImpl, requestDelimited, responseDelimited);
        }

        (ConversationReporter.prototype = Object.create($protobuf.rpc.Service.prototype)).constructor = ConversationReporter;

        /**
         * Callback as used by {@link conversation.ConversationReporter#countConversationsByAgent}.
         * @memberof conversation.ConversationReporter
         * @typedef CountConversationsByAgentCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {conversation.CountByAgent} [response] CountByAgent
         */

        /**
         * Calls CountConversationsByAgent.
         * @function countConversationsByAgent
         * @memberof conversation.ConversationReporter
         * @instance
         * @param {conversation.ICountByAgentRequest} request CountByAgentRequest message or plain object
         * @param {conversation.ConversationReporter.CountConversationsByAgentCallback} callback Node-style callback called with the error, if any, and CountByAgent
         * @returns {undefined}
         * @variation 1
         */
        ConversationReporter.prototype.countConversationsByAgent = function countConversationsByAgent(request, callback) {
            return this.rpcCall(countConversationsByAgent, $root.conversation.CountByAgentRequest, $root.conversation.CountByAgent, request, callback);
        };

        /**
         * Calls CountConversationsByAgent.
         * @function countConversationsByAgent
         * @memberof conversation.ConversationReporter
         * @instance
         * @param {conversation.ICountByAgentRequest} request CountByAgentRequest message or plain object
         * @returns {Promise<conversation.CountByAgent>} Promise
         * @variation 2
         */

        /**
         * Callback as used by {@link conversation.ConversationReporter#countConversationsByAgents}.
         * @memberof conversation.ConversationReporter
         * @typedef CountConversationsByAgentsCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {conversation.CountByAgentsResponse} [response] CountByAgentsResponse
         */

        /**
         * Calls CountConversationsByAgents.
         * @function countConversationsByAgents
         * @memberof conversation.ConversationReporter
         * @instance
         * @param {conversation.ICountByAgentsRequest} request CountByAgentsRequest message or plain object
         * @param {conversation.ConversationReporter.CountConversationsByAgentsCallback} callback Node-style callback called with the error, if any, and CountByAgentsResponse
         * @returns {undefined}
         * @variation 1
         */
        ConversationReporter.prototype.countConversationsByAgents = function countConversationsByAgents(request, callback) {
            return this.rpcCall(countConversationsByAgents, $root.conversation.CountByAgentsRequest, $root.conversation.CountByAgentsResponse, request, callback);
        };

        /**
         * Calls CountConversationsByAgents.
         * @function countConversationsByAgents
         * @memberof conversation.ConversationReporter
         * @instance
         * @param {conversation.ICountByAgentsRequest} request CountByAgentsRequest message or plain object
         * @returns {Promise<conversation.CountByAgentsResponse>} Promise
         * @variation 2
         */

        /**
         * Callback as used by {@link conversation.ConversationReporter#countConversationsByGroup}.
         * @memberof conversation.ConversationReporter
         * @typedef CountConversationsByGroupCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {conversation.CountByGroup} [response] CountByGroup
         */

        /**
         * Calls CountConversationsByGroup.
         * @function countConversationsByGroup
         * @memberof conversation.ConversationReporter
         * @instance
         * @param {conversation.ICountByGroupRequest} request CountByGroupRequest message or plain object
         * @param {conversation.ConversationReporter.CountConversationsByGroupCallback} callback Node-style callback called with the error, if any, and CountByGroup
         * @returns {undefined}
         * @variation 1
         */
        ConversationReporter.prototype.countConversationsByGroup = function countConversationsByGroup(request, callback) {
            return this.rpcCall(countConversationsByGroup, $root.conversation.CountByGroupRequest, $root.conversation.CountByGroup, request, callback);
        };

        /**
         * Calls CountConversationsByGroup.
         * @function countConversationsByGroup
         * @memberof conversation.ConversationReporter
         * @instance
         * @param {conversation.ICountByGroupRequest} request CountByGroupRequest message or plain object
         * @returns {Promise<conversation.CountByGroup>} Promise
         * @variation 2
         */

        /**
         * Callback as used by {@link conversation.ConversationReporter#countConversationsByTags}.
         * @memberof conversation.ConversationReporter
         * @typedef CountConversationsByTagsCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {conversation.CountByTagsResponse} [response] CountByTagsResponse
         */

        /**
         * Calls CountConversationsByTags.
         * @function countConversationsByTags
         * @memberof conversation.ConversationReporter
         * @instance
         * @param {conversation.ICountByTagsRequest} request CountByTagsRequest message or plain object
         * @param {conversation.ConversationReporter.CountConversationsByTagsCallback} callback Node-style callback called with the error, if any, and CountByTagsResponse
         * @returns {undefined}
         * @variation 1
         */
        ConversationReporter.prototype.countConversationsByTags = function countConversationsByTags(request, callback) {
            return this.rpcCall(countConversationsByTags, $root.conversation.CountByTagsRequest, $root.conversation.CountByTagsResponse, request, callback);
        };

        /**
         * Calls CountConversationsByTags.
         * @function countConversationsByTags
         * @memberof conversation.ConversationReporter
         * @instance
         * @param {conversation.ICountByTagsRequest} request CountByTagsRequest message or plain object
         * @returns {Promise<conversation.CountByTagsResponse>} Promise
         * @variation 2
         */

        /**
         * Callback as used by {@link conversation.ConversationReporter#getAvgResponseTimes}.
         * @memberof conversation.ConversationReporter
         * @typedef GetAvgResponseTimesCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {conversation.AvgResponseTimeResponse} [response] AvgResponseTimeResponse
         */

        /**
         * Calls GetAvgResponseTimes.
         * @function getAvgResponseTimes
         * @memberof conversation.ConversationReporter
         * @instance
         * @param {conversation.IAvgResponseTimeRequest} request AvgResponseTimeRequest message or plain object
         * @param {conversation.ConversationReporter.GetAvgResponseTimesCallback} callback Node-style callback called with the error, if any, and AvgResponseTimeResponse
         * @returns {undefined}
         * @variation 1
         */
        ConversationReporter.prototype.getAvgResponseTimes = function getAvgResponseTimes(request, callback) {
            return this.rpcCall(getAvgResponseTimes, $root.conversation.AvgResponseTimeRequest, $root.conversation.AvgResponseTimeResponse, request, callback);
        };

        /**
         * Calls GetAvgResponseTimes.
         * @function getAvgResponseTimes
         * @memberof conversation.ConversationReporter
         * @instance
         * @param {conversation.IAvgResponseTimeRequest} request AvgResponseTimeRequest message or plain object
         * @returns {Promise<conversation.AvgResponseTimeResponse>} Promise
         * @variation 2
         */

        /**
         * Callback as used by {@link conversation.ConversationReporter#getTotalConversation}.
         * @memberof conversation.ConversationReporter
         * @typedef GetTotalConversationCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {conversation.TotalConversationResponse} [response] TotalConversationResponse
         */

        /**
         * Calls GetTotalConversation.
         * @function getTotalConversation
         * @memberof conversation.ConversationReporter
         * @instance
         * @param {conversation.IAvgResponseTimeRequest} request AvgResponseTimeRequest message or plain object
         * @param {conversation.ConversationReporter.GetTotalConversationCallback} callback Node-style callback called with the error, if any, and TotalConversationResponse
         * @returns {undefined}
         * @variation 1
         */
        ConversationReporter.prototype.getTotalConversation = function getTotalConversation(request, callback) {
            return this.rpcCall(getTotalConversation, $root.conversation.AvgResponseTimeRequest, $root.conversation.TotalConversationResponse, request, callback);
        };

        /**
         * Calls GetTotalConversation.
         * @function getTotalConversation
         * @memberof conversation.ConversationReporter
         * @instance
         * @param {conversation.IAvgResponseTimeRequest} request AvgResponseTimeRequest message or plain object
         * @returns {Promise<conversation.TotalConversationResponse>} Promise
         * @variation 2
         */

        return ConversationReporter;
    })();

    return conversation;
})();

export const dashboard = $root.dashboard = (() => {

    /**
     * Namespace dashboard.
     * @exports dashboard
     * @namespace
     */
    const dashboard = {};

    dashboard.SessionCookie = (function() {

        /**
         * Properties of a SessionCookie.
         * @memberof dashboard
         * @interface ISessionCookie
         * @property {string|null} [refresh_token] SessionCookie refresh_token
         * @property {number|Long|null} [expired_at] SessionCookie expired_at
         * @property {number|Long|null} [issued_at] SessionCookie issued_at
         * @property {string|null} [type] SessionCookie type
         * @property {string|null} [email] SessionCookie email
         * @property {boolean|null} [remember_me] SessionCookie remember_me
         */

        /**
         * Constructs a new SessionCookie.
         * @memberof dashboard
         * @classdesc Represents a SessionCookie.
         * @implements ISessionCookie
         * @constructor
         * @param {dashboard.ISessionCookie=} [p] Properties to set
         */
        function SessionCookie(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * SessionCookie refresh_token.
         * @member {string} refresh_token
         * @memberof dashboard.SessionCookie
         * @instance
         */
        SessionCookie.prototype.refresh_token = "";

        /**
         * SessionCookie expired_at.
         * @member {number|Long} expired_at
         * @memberof dashboard.SessionCookie
         * @instance
         */
        SessionCookie.prototype.expired_at = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

        /**
         * SessionCookie issued_at.
         * @member {number|Long} issued_at
         * @memberof dashboard.SessionCookie
         * @instance
         */
        SessionCookie.prototype.issued_at = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

        /**
         * SessionCookie type.
         * @member {string} type
         * @memberof dashboard.SessionCookie
         * @instance
         */
        SessionCookie.prototype.type = "";

        /**
         * SessionCookie email.
         * @member {string} email
         * @memberof dashboard.SessionCookie
         * @instance
         */
        SessionCookie.prototype.email = "";

        /**
         * SessionCookie remember_me.
         * @member {boolean} remember_me
         * @memberof dashboard.SessionCookie
         * @instance
         */
        SessionCookie.prototype.remember_me = false;

        return SessionCookie;
    })();

    dashboard.Global = (function() {

        /**
         * Properties of a Global.
         * @memberof dashboard
         * @interface IGlobal
         * @property {common.IContext|null} [ctx] Global ctx
         * @property {string|null} [name] Global name
         * @property {string|null} [data] Global data
         * @property {string|null} [pk] Global pk
         */

        /**
         * Constructs a new Global.
         * @memberof dashboard
         * @classdesc Represents a Global.
         * @implements IGlobal
         * @constructor
         * @param {dashboard.IGlobal=} [p] Properties to set
         */
        function Global(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * Global ctx.
         * @member {common.IContext|null|undefined} ctx
         * @memberof dashboard.Global
         * @instance
         */
        Global.prototype.ctx = null;

        /**
         * Global name.
         * @member {string} name
         * @memberof dashboard.Global
         * @instance
         */
        Global.prototype.name = "";

        /**
         * Global data.
         * @member {string} data
         * @memberof dashboard.Global
         * @instance
         */
        Global.prototype.data = "";

        /**
         * Global pk.
         * @member {string} pk
         * @memberof dashboard.Global
         * @instance
         */
        Global.prototype.pk = "";

        return Global;
    })();

    /**
     * Event enum.
     * @name dashboard.Event
     * @enum {string}
     * @property {number} DashboardVersionUpdated=0 DashboardVersionUpdated value
     * @property {number} DashboardLanguageUpdated=5 DashboardLanguageUpdated value
     * @property {number} DashboardLanguageUpdateRequested=7 DashboardLanguageUpdateRequested value
     */
    dashboard.Event = (function() {
        const valuesById = {}, values = Object.create(valuesById);
        values[valuesById[0] = "DashboardVersionUpdated"] = 0;
        values[valuesById[5] = "DashboardLanguageUpdated"] = 5;
        values[valuesById[7] = "DashboardLanguageUpdateRequested"] = 7;
        return values;
    })();

    return dashboard;
})();

export const email = $root.email = (() => {

    /**
     * Namespace email.
     * @exports email
     * @namespace
     */
    const email = {};

    email.Email = (function() {

        /**
         * Properties of an Email.
         * @memberof email
         * @interface IEmail
         * @property {common.IContext|null} [ctx] Email ctx
         * @property {string|null} [from] Email from
         * @property {string|null} [subject] Email subject
         * @property {string|null} [body] Email body
         * @property {Object.<string,string>|null} [header] Email header
         * @property {Array.<email.IAttachment>|null} [attachments] Email attachments
         * @property {Array.<string>|null} [to] Email to
         * @property {Array.<string>|null} [cc] Email cc
         * @property {Array.<string>|null} [bcc] Email bcc
         */

        /**
         * Constructs a new Email.
         * @memberof email
         * @classdesc Represents an Email.
         * @implements IEmail
         * @constructor
         * @param {email.IEmail=} [p] Properties to set
         */
        function Email(p) {
            this.header = {};
            this.attachments = [];
            this.to = [];
            this.cc = [];
            this.bcc = [];
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * Email ctx.
         * @member {common.IContext|null|undefined} ctx
         * @memberof email.Email
         * @instance
         */
        Email.prototype.ctx = null;

        /**
         * Email from.
         * @member {string} from
         * @memberof email.Email
         * @instance
         */
        Email.prototype.from = "";

        /**
         * Email subject.
         * @member {string} subject
         * @memberof email.Email
         * @instance
         */
        Email.prototype.subject = "";

        /**
         * Email body.
         * @member {string} body
         * @memberof email.Email
         * @instance
         */
        Email.prototype.body = "";

        /**
         * Email header.
         * @member {Object.<string,string>} header
         * @memberof email.Email
         * @instance
         */
        Email.prototype.header = $util.emptyObject;

        /**
         * Email attachments.
         * @member {Array.<email.IAttachment>} attachments
         * @memberof email.Email
         * @instance
         */
        Email.prototype.attachments = $util.emptyArray;

        /**
         * Email to.
         * @member {Array.<string>} to
         * @memberof email.Email
         * @instance
         */
        Email.prototype.to = $util.emptyArray;

        /**
         * Email cc.
         * @member {Array.<string>} cc
         * @memberof email.Email
         * @instance
         */
        Email.prototype.cc = $util.emptyArray;

        /**
         * Email bcc.
         * @member {Array.<string>} bcc
         * @memberof email.Email
         * @instance
         */
        Email.prototype.bcc = $util.emptyArray;

        return Email;
    })();

    email.Attachment = (function() {

        /**
         * Properties of an Attachment.
         * @memberof email
         * @interface IAttachment
         * @property {string|null} [url] Attachment url
         * @property {string|null} [name] Attachment name
         * @property {string|null} [mimetype] Attachment mimetype
         */

        /**
         * Constructs a new Attachment.
         * @memberof email
         * @classdesc Represents an Attachment.
         * @implements IAttachment
         * @constructor
         * @param {email.IAttachment=} [p] Properties to set
         */
        function Attachment(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * Attachment url.
         * @member {string} url
         * @memberof email.Attachment
         * @instance
         */
        Attachment.prototype.url = "";

        /**
         * Attachment name.
         * @member {string} name
         * @memberof email.Attachment
         * @instance
         */
        Attachment.prototype.name = "";

        /**
         * Attachment mimetype.
         * @member {string} mimetype
         * @memberof email.Attachment
         * @instance
         */
        Attachment.prototype.mimetype = "";

        return Attachment;
    })();

    /**
     * Event enum.
     * @name email.Event
     * @enum {string}
     * @property {number} Email_SendRequest=0 Email_SendRequest value
     */
    email.Event = (function() {
        const valuesById = {}, values = Object.create(valuesById);
        values[valuesById[0] = "Email_SendRequest"] = 0;
        return values;
    })();

    return email;
})();

export const event = $root.event = (() => {

    /**
     * Namespace event.
     * @exports event
     * @namespace
     */
    const event = {};

    event.RawEventCreatedPayload = (function() {

        /**
         * Properties of a RawEventCreatedPayload.
         * @memberof event
         * @interface IRawEventCreatedPayload
         * @property {common.IContext|null} [ctx] RawEventCreatedPayload ctx
         * @property {Array.<string>|null} [subs] RawEventCreatedPayload subs
         * @property {string|null} [target_topic] RawEventCreatedPayload target_topic
         * @property {string|null} [payload] RawEventCreatedPayload payload
         * @property {string|null} [target_key] RawEventCreatedPayload target_key
         * @property {Array.<string>|null} [payloads] RawEventCreatedPayload payloads
         * @property {string|null} [topic] RawEventCreatedPayload topic
         * @property {string|null} [router_topic] RawEventCreatedPayload router_topic
         * @property {string|null} [sub] RawEventCreatedPayload sub
         */

        /**
         * Constructs a new RawEventCreatedPayload.
         * @memberof event
         * @classdesc Represents a RawEventCreatedPayload.
         * @implements IRawEventCreatedPayload
         * @constructor
         * @param {event.IRawEventCreatedPayload=} [p] Properties to set
         */
        function RawEventCreatedPayload(p) {
            this.subs = [];
            this.payloads = [];
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * RawEventCreatedPayload ctx.
         * @member {common.IContext|null|undefined} ctx
         * @memberof event.RawEventCreatedPayload
         * @instance
         */
        RawEventCreatedPayload.prototype.ctx = null;

        /**
         * RawEventCreatedPayload subs.
         * @member {Array.<string>} subs
         * @memberof event.RawEventCreatedPayload
         * @instance
         */
        RawEventCreatedPayload.prototype.subs = $util.emptyArray;

        /**
         * RawEventCreatedPayload target_topic.
         * @member {string} target_topic
         * @memberof event.RawEventCreatedPayload
         * @instance
         */
        RawEventCreatedPayload.prototype.target_topic = "";

        /**
         * RawEventCreatedPayload payload.
         * @member {string} payload
         * @memberof event.RawEventCreatedPayload
         * @instance
         */
        RawEventCreatedPayload.prototype.payload = "";

        /**
         * RawEventCreatedPayload target_key.
         * @member {string} target_key
         * @memberof event.RawEventCreatedPayload
         * @instance
         */
        RawEventCreatedPayload.prototype.target_key = "";

        /**
         * RawEventCreatedPayload payloads.
         * @member {Array.<string>} payloads
         * @memberof event.RawEventCreatedPayload
         * @instance
         */
        RawEventCreatedPayload.prototype.payloads = $util.emptyArray;

        /**
         * RawEventCreatedPayload topic.
         * @member {string} topic
         * @memberof event.RawEventCreatedPayload
         * @instance
         */
        RawEventCreatedPayload.prototype.topic = "";

        /**
         * RawEventCreatedPayload router_topic.
         * @member {string} router_topic
         * @memberof event.RawEventCreatedPayload
         * @instance
         */
        RawEventCreatedPayload.prototype.router_topic = "";

        /**
         * RawEventCreatedPayload sub.
         * @member {string} sub
         * @memberof event.RawEventCreatedPayload
         * @instance
         */
        RawEventCreatedPayload.prototype.sub = "";

        return RawEventCreatedPayload;
    })();

    event.PubSubMessage = (function() {

        /**
         * Properties of a PubSubMessage.
         * @memberof event
         * @interface IPubSubMessage
         * @property {common.IContext|null} [ctx] PubSubMessage ctx
         * @property {event.ISubscription|null} [sub] PubSubMessage sub
         * @property {Uint8Array|null} [payload] PubSubMessage payload
         */

        /**
         * Constructs a new PubSubMessage.
         * @memberof event
         * @classdesc Represents a PubSubMessage.
         * @implements IPubSubMessage
         * @constructor
         * @param {event.IPubSubMessage=} [p] Properties to set
         */
        function PubSubMessage(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * PubSubMessage ctx.
         * @member {common.IContext|null|undefined} ctx
         * @memberof event.PubSubMessage
         * @instance
         */
        PubSubMessage.prototype.ctx = null;

        /**
         * PubSubMessage sub.
         * @member {event.ISubscription|null|undefined} sub
         * @memberof event.PubSubMessage
         * @instance
         */
        PubSubMessage.prototype.sub = null;

        /**
         * PubSubMessage payload.
         * @member {Uint8Array} payload
         * @memberof event.PubSubMessage
         * @instance
         */
        PubSubMessage.prototype.payload = $util.newBuffer([]);

        return PubSubMessage;
    })();

    event.UnsubscribeMessage = (function() {

        /**
         * Properties of an UnsubscribeMessage.
         * @memberof event
         * @interface IUnsubscribeMessage
         * @property {common.IContext|null} [ctx] UnsubscribeMessage ctx
         * @property {string|null} [topic] UnsubscribeMessage topic
         * @property {string|null} [sub_id] UnsubscribeMessage sub_id
         */

        /**
         * Constructs a new UnsubscribeMessage.
         * @memberof event
         * @classdesc Represents an UnsubscribeMessage.
         * @implements IUnsubscribeMessage
         * @constructor
         * @param {event.IUnsubscribeMessage=} [p] Properties to set
         */
        function UnsubscribeMessage(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * UnsubscribeMessage ctx.
         * @member {common.IContext|null|undefined} ctx
         * @memberof event.UnsubscribeMessage
         * @instance
         */
        UnsubscribeMessage.prototype.ctx = null;

        /**
         * UnsubscribeMessage topic.
         * @member {string} topic
         * @memberof event.UnsubscribeMessage
         * @instance
         */
        UnsubscribeMessage.prototype.topic = "";

        /**
         * UnsubscribeMessage sub_id.
         * @member {string} sub_id
         * @memberof event.UnsubscribeMessage
         * @instance
         */
        UnsubscribeMessage.prototype.sub_id = "";

        return UnsubscribeMessage;
    })();

    /**
     * Event enum.
     * @name event.Event
     * @enum {string}
     * @property {number} Sub=0 Sub value
     * @property {number} RawEventCreated=3 RawEventCreated value
     * @property {number} Subscribe=4 Subscribe value
     * @property {number} SubscribeReply=6 SubscribeReply value
     * @property {number} Unsubscribe=5 Unsubscribe value
     * @property {number} UnsubscribeReply=7 UnsubscribeReply value
     * @property {number} EventSync=8 EventSync value
     * @property {number} EventCreated=9 EventCreated value
     */
    event.Event = (function() {
        const valuesById = {}, values = Object.create(valuesById);
        values[valuesById[0] = "Sub"] = 0;
        values[valuesById[3] = "RawEventCreated"] = 3;
        values[valuesById[4] = "Subscribe"] = 4;
        values[valuesById[6] = "SubscribeReply"] = 6;
        values[valuesById[5] = "Unsubscribe"] = 5;
        values[valuesById[7] = "UnsubscribeReply"] = 7;
        values[valuesById[8] = "EventSync"] = 8;
        values[valuesById[9] = "EventCreated"] = 9;
        return values;
    })();

    event.RawEvents = (function() {

        /**
         * Properties of a RawEvents.
         * @memberof event
         * @interface IRawEvents
         * @property {common.IContext|null} [ctx] RawEvents ctx
         * @property {Array.<event.IRawEvent>|null} [events] RawEvents events
         * @property {number|Long|null} [total] RawEvents total
         * @property {string|null} [anchor] RawEvents anchor
         */

        /**
         * Constructs a new RawEvents.
         * @memberof event
         * @classdesc Represents a RawEvents.
         * @implements IRawEvents
         * @constructor
         * @param {event.IRawEvents=} [p] Properties to set
         */
        function RawEvents(p) {
            this.events = [];
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * RawEvents ctx.
         * @member {common.IContext|null|undefined} ctx
         * @memberof event.RawEvents
         * @instance
         */
        RawEvents.prototype.ctx = null;

        /**
         * RawEvents events.
         * @member {Array.<event.IRawEvent>} events
         * @memberof event.RawEvents
         * @instance
         */
        RawEvents.prototype.events = $util.emptyArray;

        /**
         * RawEvents total.
         * @member {number|Long} total
         * @memberof event.RawEvents
         * @instance
         */
        RawEvents.prototype.total = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

        /**
         * RawEvents anchor.
         * @member {string} anchor
         * @memberof event.RawEvents
         * @instance
         */
        RawEvents.prototype.anchor = "";

        return RawEvents;
    })();

    event.By = (function() {

        /**
         * Properties of a By.
         * @memberof event
         * @interface IBy
         * @property {common.IDevice|null} [device] By device
         * @property {string|null} [id] By id
         * @property {string|null} [client_id] By client_id
         */

        /**
         * Constructs a new By.
         * @memberof event
         * @classdesc Represents a By.
         * @implements IBy
         * @constructor
         * @param {event.IBy=} [p] Properties to set
         */
        function By(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * By device.
         * @member {common.IDevice|null|undefined} device
         * @memberof event.By
         * @instance
         */
        By.prototype.device = null;

        /**
         * By id.
         * @member {string} id
         * @memberof event.By
         * @instance
         */
        By.prototype.id = "";

        /**
         * By client_id.
         * @member {string} client_id
         * @memberof event.By
         * @instance
         */
        By.prototype.client_id = "";

        return By;
    })();

    event.RawEvent = (function() {

        /**
         * Properties of a RawEvent.
         * @memberof event
         * @interface IRawEvent
         * @property {common.IContext|null} [ctx] RawEvent ctx
         * @property {string|null} [id] RawEvent id
         * @property {string|null} [account_id] RawEvent account_id
         * @property {number|Long|null} [created] RawEvent created
         * @property {string|null} [type] RawEvent type
         * @property {Array.<string>|null} [topics] RawEvent topics
         * @property {event.IBy|null} [by] RawEvent by
         * @property {string|null} [object] RawEvent object
         * @property {string|null} [conversation_id] RawEvent conversation_id
         * @property {event.RawEvent.IData|null} [data] RawEvent data
         */

        /**
         * Constructs a new RawEvent.
         * @memberof event
         * @classdesc Represents a RawEvent.
         * @implements IRawEvent
         * @constructor
         * @param {event.IRawEvent=} [p] Properties to set
         */
        function RawEvent(p) {
            this.topics = [];
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * RawEvent ctx.
         * @member {common.IContext|null|undefined} ctx
         * @memberof event.RawEvent
         * @instance
         */
        RawEvent.prototype.ctx = null;

        /**
         * RawEvent id.
         * @member {string} id
         * @memberof event.RawEvent
         * @instance
         */
        RawEvent.prototype.id = "";

        /**
         * RawEvent account_id.
         * @member {string} account_id
         * @memberof event.RawEvent
         * @instance
         */
        RawEvent.prototype.account_id = "";

        /**
         * RawEvent created.
         * @member {number|Long} created
         * @memberof event.RawEvent
         * @instance
         */
        RawEvent.prototype.created = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

        /**
         * RawEvent type.
         * @member {string} type
         * @memberof event.RawEvent
         * @instance
         */
        RawEvent.prototype.type = "";

        /**
         * RawEvent topics.
         * @member {Array.<string>} topics
         * @memberof event.RawEvent
         * @instance
         */
        RawEvent.prototype.topics = $util.emptyArray;

        /**
         * RawEvent by.
         * @member {event.IBy|null|undefined} by
         * @memberof event.RawEvent
         * @instance
         */
        RawEvent.prototype.by = null;

        /**
         * RawEvent object.
         * @member {string} object
         * @memberof event.RawEvent
         * @instance
         */
        RawEvent.prototype.object = "";

        /**
         * RawEvent conversation_id.
         * @member {string} conversation_id
         * @memberof event.RawEvent
         * @instance
         */
        RawEvent.prototype.conversation_id = "";

        /**
         * RawEvent data.
         * @member {event.RawEvent.IData|null|undefined} data
         * @memberof event.RawEvent
         * @instance
         */
        RawEvent.prototype.data = null;

        RawEvent.Data = (function() {

            /**
             * Properties of a Data.
             * @memberof event.RawEvent
             * @interface IData
             * @property {account.IAccount|null} [account] Data account
             * @property {account.IAgent|null} [agent] Data agent
             * @property {conversation.IMessage|null} [message] Data message
             * @property {conversation.IConversation|null} [conversation] Data conversation
             * @property {conversation.IPostback|null} [postback] Data postback
             * @property {content.IContent|null} [content] Data content
             * @property {user.ITopic|null} [topic] Data topic
             * @property {user.IPresence|null} [presence] Data presence
             * @property {user.IUser|null} [user] Data user
             * @property {user.IUnreadTopic|null} [unread_topic] Data unread_topic
             * @property {user.IMyUser|null} [my_user] Data my_user
             * @property {notibox.INotification|null} [notification] Data notification
             * @property {notibox.IBox|null} [notibox] Data notibox
             * @property {account.IAgentPerm|null} [agent_perm] Data agent_perm
             * @property {account.IGroupMember|null} [group_member] Data group_member
             * @property {account.IAgentGroup|null} [group] Data group
             * @property {payment.ILimit|null} [limit] Data limit
             * @property {user.IAttributeData|null} [user_attribute] Data user_attribute
             */

            /**
             * Constructs a new Data.
             * @memberof event.RawEvent
             * @classdesc Represents a Data.
             * @implements IData
             * @constructor
             * @param {event.RawEvent.IData=} [p] Properties to set
             */
            function Data(p) {
                if (p)
                    for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                        if (p[ks[i]] != null)
                            this[ks[i]] = p[ks[i]];
            }

            /**
             * Data account.
             * @member {account.IAccount|null|undefined} account
             * @memberof event.RawEvent.Data
             * @instance
             */
            Data.prototype.account = null;

            /**
             * Data agent.
             * @member {account.IAgent|null|undefined} agent
             * @memberof event.RawEvent.Data
             * @instance
             */
            Data.prototype.agent = null;

            /**
             * Data message.
             * @member {conversation.IMessage|null|undefined} message
             * @memberof event.RawEvent.Data
             * @instance
             */
            Data.prototype.message = null;

            /**
             * Data conversation.
             * @member {conversation.IConversation|null|undefined} conversation
             * @memberof event.RawEvent.Data
             * @instance
             */
            Data.prototype.conversation = null;

            /**
             * Data postback.
             * @member {conversation.IPostback|null|undefined} postback
             * @memberof event.RawEvent.Data
             * @instance
             */
            Data.prototype.postback = null;

            /**
             * Data content.
             * @member {content.IContent|null|undefined} content
             * @memberof event.RawEvent.Data
             * @instance
             */
            Data.prototype.content = null;

            /**
             * Data topic.
             * @member {user.ITopic|null|undefined} topic
             * @memberof event.RawEvent.Data
             * @instance
             */
            Data.prototype.topic = null;

            /**
             * Data presence.
             * @member {user.IPresence|null|undefined} presence
             * @memberof event.RawEvent.Data
             * @instance
             */
            Data.prototype.presence = null;

            /**
             * Data user.
             * @member {user.IUser|null|undefined} user
             * @memberof event.RawEvent.Data
             * @instance
             */
            Data.prototype.user = null;

            /**
             * Data unread_topic.
             * @member {user.IUnreadTopic|null|undefined} unread_topic
             * @memberof event.RawEvent.Data
             * @instance
             */
            Data.prototype.unread_topic = null;

            /**
             * Data my_user.
             * @member {user.IMyUser|null|undefined} my_user
             * @memberof event.RawEvent.Data
             * @instance
             */
            Data.prototype.my_user = null;

            /**
             * Data notification.
             * @member {notibox.INotification|null|undefined} notification
             * @memberof event.RawEvent.Data
             * @instance
             */
            Data.prototype.notification = null;

            /**
             * Data notibox.
             * @member {notibox.IBox|null|undefined} notibox
             * @memberof event.RawEvent.Data
             * @instance
             */
            Data.prototype.notibox = null;

            /**
             * Data agent_perm.
             * @member {account.IAgentPerm|null|undefined} agent_perm
             * @memberof event.RawEvent.Data
             * @instance
             */
            Data.prototype.agent_perm = null;

            /**
             * Data group_member.
             * @member {account.IGroupMember|null|undefined} group_member
             * @memberof event.RawEvent.Data
             * @instance
             */
            Data.prototype.group_member = null;

            /**
             * Data group.
             * @member {account.IAgentGroup|null|undefined} group
             * @memberof event.RawEvent.Data
             * @instance
             */
            Data.prototype.group = null;

            /**
             * Data limit.
             * @member {payment.ILimit|null|undefined} limit
             * @memberof event.RawEvent.Data
             * @instance
             */
            Data.prototype.limit = null;

            /**
             * Data user_attribute.
             * @member {user.IAttributeData|null|undefined} user_attribute
             * @memberof event.RawEvent.Data
             * @instance
             */
            Data.prototype.user_attribute = null;

            return Data;
        })();

        return RawEvent;
    })();

    /**
     * AccountType enum.
     * @name event.AccountType
     * @enum {string}
     * @property {number} a=0 a value
     * @property {number} account_deleted=60 account_deleted value
     * @property {number} account_info_updated=61 account_info_updated value
     * @property {number} group_created=62 group_created value
     * @property {number} group_deleted=63 group_deleted value
     * @property {number} group_info_updated=64 group_info_updated value
     * @property {number} group_joined=65 group_joined value
     * @property {number} group_left=66 group_left value
     * @property {number} invitation_accepted=68 invitation_accepted value
     * @property {number} agent_deleted=69 agent_deleted value
     * @property {number} agent_info_updated=70 agent_info_updated value
     * @property {number} agent_permission_updated=71 agent_permission_updated value
     * @property {number} limit_updated=72 limit_updated value
     */
    event.AccountType = (function() {
        const valuesById = {}, values = Object.create(valuesById);
        values[valuesById[0] = "a"] = 0;
        values[valuesById[60] = "account_deleted"] = 60;
        values[valuesById[61] = "account_info_updated"] = 61;
        values[valuesById[62] = "group_created"] = 62;
        values[valuesById[63] = "group_deleted"] = 63;
        values[valuesById[64] = "group_info_updated"] = 64;
        values[valuesById[65] = "group_joined"] = 65;
        values[valuesById[66] = "group_left"] = 66;
        values[valuesById[68] = "invitation_accepted"] = 68;
        values[valuesById[69] = "agent_deleted"] = 69;
        values[valuesById[70] = "agent_info_updated"] = 70;
        values[valuesById[71] = "agent_permission_updated"] = 71;
        values[valuesById[72] = "limit_updated"] = 72;
        return values;
    })();

    /**
     * UserType enum.
     * @name event.UserType
     * @enum {string}
     * @property {number} u=0 u value
     * @property {number} my_user_upserted=46 my_user_upserted value
     * @property {number} user_info_updated=48 user_info_updated value
     */
    event.UserType = (function() {
        const valuesById = {}, values = Object.create(valuesById);
        values[valuesById[0] = "u"] = 0;
        values[valuesById[46] = "my_user_upserted"] = 46;
        values[valuesById[48] = "user_info_updated"] = 48;
        return values;
    })();

    /**
     * EventType enum.
     * @name event.EventType
     * @enum {string}
     * @property {number} e=0 e value
     * @property {number} presence_updated=38 presence_updated value
     * @property {number} content_viewed=39 content_viewed value
     * @property {number} content_searched=40 content_searched value
     * @property {number} content_addedtocart=41 content_addedtocart value
     * @property {number} content_checkedout=42 content_checkedout value
     * @property {number} content_purchased=43 content_purchased value
     * @property {number} topic_read=44 topic_read value
     * @property {number} subscribed_topic_updated=45 subscribed_topic_updated value
     * @property {number} user_topic_updated=50 user_topic_updated value
     * @property {number} user_attribute_created=51 user_attribute_created value
     * @property {number} user_attribute_updated=52 user_attribute_updated value
     * @property {number} user_attribute_deleted=55 user_attribute_deleted value
     */
    event.EventType = (function() {
        const valuesById = {}, values = Object.create(valuesById);
        values[valuesById[0] = "e"] = 0;
        values[valuesById[38] = "presence_updated"] = 38;
        values[valuesById[39] = "content_viewed"] = 39;
        values[valuesById[40] = "content_searched"] = 40;
        values[valuesById[41] = "content_addedtocart"] = 41;
        values[valuesById[42] = "content_checkedout"] = 42;
        values[valuesById[43] = "content_purchased"] = 43;
        values[valuesById[44] = "topic_read"] = 44;
        values[valuesById[45] = "subscribed_topic_updated"] = 45;
        values[valuesById[50] = "user_topic_updated"] = 50;
        values[valuesById[51] = "user_attribute_created"] = 51;
        values[valuesById[52] = "user_attribute_updated"] = 52;
        values[valuesById[55] = "user_attribute_deleted"] = 55;
        return values;
    })();

    /**
     * ConvoType enum.
     * @name event.ConvoType
     * @enum {string}
     * @property {number} c=0 c value
     * @property {number} conversation_updated=9 conversation_updated value
     * @property {number} message_sent=10 message_sent value
     * @property {number} conversation_state_updated=11 conversation_state_updated value
     * @property {number} conversation_joined=2 conversation_joined value
     * @property {number} conversation_left=4 conversation_left value
     * @property {number} conversation_tagged=6 conversation_tagged value
     * @property {number} conversation_untagged=7 conversation_untagged value
     * @property {number} channel_deintegrated=20 channel_deintegrated value
     * @property {number} channel_integrated=21 channel_integrated value
     * @property {number} message_seen=30 message_seen value
     * @property {number} message_acked=31 message_acked value
     * @property {number} message_received=32 message_received value
     * @property {number} conversation_member_typing=33 conversation_member_typing value
     * @property {number} conversation_postbacked=34 conversation_postbacked value
     * @property {number} conversation_unassigned=35 conversation_unassigned value
     * @property {number} conversation_assigned=36 conversation_assigned value
     * @property {number} conversation_pending=81 conversation_pending value
     */
    event.ConvoType = (function() {
        const valuesById = {}, values = Object.create(valuesById);
        values[valuesById[0] = "c"] = 0;
        values[valuesById[9] = "conversation_updated"] = 9;
        values[valuesById[10] = "message_sent"] = 10;
        values[valuesById[11] = "conversation_state_updated"] = 11;
        values[valuesById[2] = "conversation_joined"] = 2;
        values[valuesById[4] = "conversation_left"] = 4;
        values[valuesById[6] = "conversation_tagged"] = 6;
        values[valuesById[7] = "conversation_untagged"] = 7;
        values[valuesById[20] = "channel_deintegrated"] = 20;
        values[valuesById[21] = "channel_integrated"] = 21;
        values[valuesById[30] = "message_seen"] = 30;
        values[valuesById[31] = "message_acked"] = 31;
        values[valuesById[32] = "message_received"] = 32;
        values[valuesById[33] = "conversation_member_typing"] = 33;
        values[valuesById[34] = "conversation_postbacked"] = 34;
        values[valuesById[35] = "conversation_unassigned"] = 35;
        values[valuesById[36] = "conversation_assigned"] = 36;
        values[valuesById[81] = "conversation_pending"] = 81;
        return values;
    })();

    /**
     * NotiboxType enum.
     * @name event.NotiboxType
     * @enum {string}
     * @property {number} n=0 n value
     * @property {number} notification_upserted=52 notification_upserted value
     * @property {number} notibox_upserted=53 notibox_upserted value
     */
    event.NotiboxType = (function() {
        const valuesById = {}, values = Object.create(valuesById);
        values[valuesById[0] = "n"] = 0;
        values[valuesById[52] = "notification_upserted"] = 52;
        values[valuesById[53] = "notibox_upserted"] = 53;
        return values;
    })();

    /**
     * SubscriberType enum.
     * @name event.SubscriberType
     * @enum {string}
     * @property {number} st_user=0 st_user value
     * @property {number} st_channel=1 st_channel value
     */
    event.SubscriberType = (function() {
        const valuesById = {}, values = Object.create(valuesById);
        values[valuesById[0] = "st_user"] = 0;
        values[valuesById[1] = "st_channel"] = 1;
        return values;
    })();

    event.Subscription = (function() {

        /**
         * Properties of a Subscription.
         * @memberof event
         * @interface ISubscription
         * @property {common.IContext|null} [ctx] Subscription ctx
         * @property {string|null} [topic] Subscription topic
         * @property {string|null} [sub_id] Subscription sub_id
         * @property {string|null} [target_topic] Subscription target_topic
         * @property {string|null} [target_key] Subscription target_key
         * @property {number|Long|null} [ttls] Subscription ttls
         * @property {string|null} [router_topic] Subscription router_topic
         * @property {number|null} [target_partition] Subscription target_partition
         */

        /**
         * Constructs a new Subscription.
         * @memberof event
         * @classdesc Represents a Subscription.
         * @implements ISubscription
         * @constructor
         * @param {event.ISubscription=} [p] Properties to set
         */
        function Subscription(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * Subscription ctx.
         * @member {common.IContext|null|undefined} ctx
         * @memberof event.Subscription
         * @instance
         */
        Subscription.prototype.ctx = null;

        /**
         * Subscription topic.
         * @member {string} topic
         * @memberof event.Subscription
         * @instance
         */
        Subscription.prototype.topic = "";

        /**
         * Subscription sub_id.
         * @member {string} sub_id
         * @memberof event.Subscription
         * @instance
         */
        Subscription.prototype.sub_id = "";

        /**
         * Subscription target_topic.
         * @member {string} target_topic
         * @memberof event.Subscription
         * @instance
         */
        Subscription.prototype.target_topic = "";

        /**
         * Subscription target_key.
         * @member {string} target_key
         * @memberof event.Subscription
         * @instance
         */
        Subscription.prototype.target_key = "";

        /**
         * Subscription ttls.
         * @member {number|Long} ttls
         * @memberof event.Subscription
         * @instance
         */
        Subscription.prototype.ttls = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

        /**
         * Subscription router_topic.
         * @member {string} router_topic
         * @memberof event.Subscription
         * @instance
         */
        Subscription.prototype.router_topic = "";

        /**
         * Subscription target_partition.
         * @member {number} target_partition
         * @memberof event.Subscription
         * @instance
         */
        Subscription.prototype.target_partition = 0;

        return Subscription;
    })();

    /**
     * SubPrefix enum.
     * @name event.SubPrefix
     * @enum {string}
     * @property {number} Webhook=0 Webhook value
     * @property {number} Websocket=1 Websocket value
     */
    event.SubPrefix = (function() {
        const valuesById = {}, values = Object.create(valuesById);
        values[valuesById[0] = "Webhook"] = 0;
        values[valuesById[1] = "Websocket"] = 1;
        return values;
    })();

    /**
     * Object enum.
     * @name event.Object
     * @enum {string}
     * @property {number} none=0 none value
     * @property {number} account=1 account value
     * @property {number} agent=2 agent value
     * @property {number} message=3 message value
     * @property {number} conversation=4 conversation value
     * @property {number} postback=5 postback value
     * @property {number} content=6 content value
     * @property {number} topic=7 topic value
     * @property {number} presence=8 presence value
     * @property {number} user=10 user value
     * @property {number} unread_topic=11 unread_topic value
     * @property {number} my_user=12 my_user value
     * @property {number} notification=14 notification value
     * @property {number} notibox=15 notibox value
     * @property {number} agent_perm=16 agent_perm value
     * @property {number} group_member=17 group_member value
     * @property {number} group=18 group value
     * @property {number} integration=30 integration value
     * @property {number} limit=19 limit value
     */
    event.Object = (function() {
        const valuesById = {}, values = Object.create(valuesById);
        values[valuesById[0] = "none"] = 0;
        values[valuesById[1] = "account"] = 1;
        values[valuesById[2] = "agent"] = 2;
        values[valuesById[3] = "message"] = 3;
        values[valuesById[4] = "conversation"] = 4;
        values[valuesById[5] = "postback"] = 5;
        values[valuesById[6] = "content"] = 6;
        values[valuesById[7] = "topic"] = 7;
        values[valuesById[8] = "presence"] = 8;
        values[valuesById[10] = "user"] = 10;
        values[valuesById[11] = "unread_topic"] = 11;
        values[valuesById[12] = "my_user"] = 12;
        values[valuesById[14] = "notification"] = 14;
        values[valuesById[15] = "notibox"] = 15;
        values[valuesById[16] = "agent_perm"] = 16;
        values[valuesById[17] = "group_member"] = 17;
        values[valuesById[18] = "group"] = 18;
        values[valuesById[30] = "integration"] = 30;
        values[valuesById[19] = "limit"] = 19;
        return values;
    })();

    event.ListEventsRequest = (function() {

        /**
         * Properties of a ListEventsRequest.
         * @memberof event
         * @interface IListEventsRequest
         * @property {common.IContext|null} [ctx] ListEventsRequest ctx
         * @property {string|null} [account_id] ListEventsRequest account_id
         * @property {string|null} [user_id] ListEventsRequest user_id
         * @property {string|null} [query] ListEventsRequest query
         * @property {string|null} [anchor] ListEventsRequest anchor
         * @property {number|null} [limit] ListEventsRequest limit
         * @property {string|null} [category] ListEventsRequest category
         */

        /**
         * Constructs a new ListEventsRequest.
         * @memberof event
         * @classdesc Represents a ListEventsRequest.
         * @implements IListEventsRequest
         * @constructor
         * @param {event.IListEventsRequest=} [p] Properties to set
         */
        function ListEventsRequest(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * ListEventsRequest ctx.
         * @member {common.IContext|null|undefined} ctx
         * @memberof event.ListEventsRequest
         * @instance
         */
        ListEventsRequest.prototype.ctx = null;

        /**
         * ListEventsRequest account_id.
         * @member {string} account_id
         * @memberof event.ListEventsRequest
         * @instance
         */
        ListEventsRequest.prototype.account_id = "";

        /**
         * ListEventsRequest user_id.
         * @member {string} user_id
         * @memberof event.ListEventsRequest
         * @instance
         */
        ListEventsRequest.prototype.user_id = "";

        /**
         * ListEventsRequest query.
         * @member {string} query
         * @memberof event.ListEventsRequest
         * @instance
         */
        ListEventsRequest.prototype.query = "";

        /**
         * ListEventsRequest anchor.
         * @member {string} anchor
         * @memberof event.ListEventsRequest
         * @instance
         */
        ListEventsRequest.prototype.anchor = "";

        /**
         * ListEventsRequest limit.
         * @member {number} limit
         * @memberof event.ListEventsRequest
         * @instance
         */
        ListEventsRequest.prototype.limit = 0;

        /**
         * ListEventsRequest category.
         * @member {string} category
         * @memberof event.ListEventsRequest
         * @instance
         */
        ListEventsRequest.prototype.category = "";

        return ListEventsRequest;
    })();

    event.UserEvent = (function() {

        /**
         * Properties of a UserEvent.
         * @memberof event
         * @interface IUserEvent
         * @property {common.IContext|null} [ctx] UserEvent ctx
         * @property {string|null} [account_id] UserEvent account_id
         * @property {string|null} [user_id] UserEvent user_id
         * @property {event.IRawEvent|null} [event] UserEvent event
         */

        /**
         * Constructs a new UserEvent.
         * @memberof event
         * @classdesc Represents a UserEvent.
         * @implements IUserEvent
         * @constructor
         * @param {event.IUserEvent=} [p] Properties to set
         */
        function UserEvent(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * UserEvent ctx.
         * @member {common.IContext|null|undefined} ctx
         * @memberof event.UserEvent
         * @instance
         */
        UserEvent.prototype.ctx = null;

        /**
         * UserEvent account_id.
         * @member {string} account_id
         * @memberof event.UserEvent
         * @instance
         */
        UserEvent.prototype.account_id = "";

        /**
         * UserEvent user_id.
         * @member {string} user_id
         * @memberof event.UserEvent
         * @instance
         */
        UserEvent.prototype.user_id = "";

        /**
         * UserEvent event.
         * @member {event.IRawEvent|null|undefined} event
         * @memberof event.UserEvent
         * @instance
         */
        UserEvent.prototype.event = null;

        return UserEvent;
    })();

    event.SubscriptionResponse = (function() {

        /**
         * Properties of a SubscriptionResponse.
         * @memberof event
         * @interface ISubscriptionResponse
         * @property {common.IContext|null} [ctx] SubscriptionResponse ctx
         * @property {boolean|null} [status] SubscriptionResponse status
         */

        /**
         * Constructs a new SubscriptionResponse.
         * @memberof event
         * @classdesc Represents a SubscriptionResponse.
         * @implements ISubscriptionResponse
         * @constructor
         * @param {event.ISubscriptionResponse=} [p] Properties to set
         */
        function SubscriptionResponse(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * SubscriptionResponse ctx.
         * @member {common.IContext|null|undefined} ctx
         * @memberof event.SubscriptionResponse
         * @instance
         */
        SubscriptionResponse.prototype.ctx = null;

        /**
         * SubscriptionResponse status.
         * @member {boolean} status
         * @memberof event.SubscriptionResponse
         * @instance
         */
        SubscriptionResponse.prototype.status = false;

        return SubscriptionResponse;
    })();

    event.Publisher = (function() {

        /**
         * Constructs a new Publisher service.
         * @memberof event
         * @classdesc Represents a Publisher
         * @extends $protobuf.rpc.Service
         * @constructor
         * @param {$protobuf.RPCImpl} rpcImpl RPC implementation
         * @param {boolean} [requestDelimited=false] Whether requests are length-delimited
         * @param {boolean} [responseDelimited=false] Whether responses are length-delimited
         */
        function Publisher(rpcImpl, requestDelimited, responseDelimited) {
            $protobuf.rpc.Service.call(this, rpcImpl, requestDelimited, responseDelimited);
        }

        (Publisher.prototype = Object.create($protobuf.rpc.Service.prototype)).constructor = Publisher;

        /**
         * Callback as used by {@link event.Publisher#subscribe}.
         * @memberof event.Publisher
         * @typedef SubscribeCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {event.SubscriptionResponse} [response] SubscriptionResponse
         */

        /**
         * Calls Subscribe.
         * @function subscribe
         * @memberof event.Publisher
         * @instance
         * @param {event.ISubscription} request Subscription message or plain object
         * @param {event.Publisher.SubscribeCallback} callback Node-style callback called with the error, if any, and SubscriptionResponse
         * @returns {undefined}
         * @variation 1
         */
        Publisher.prototype.subscribe = function subscribe(request, callback) {
            return this.rpcCall(subscribe, $root.event.Subscription, $root.event.SubscriptionResponse, request, callback);
        };

        /**
         * Calls Subscribe.
         * @function subscribe
         * @memberof event.Publisher
         * @instance
         * @param {event.ISubscription} request Subscription message or plain object
         * @returns {Promise<event.SubscriptionResponse>} Promise
         * @variation 2
         */

        /**
         * Callback as used by {@link event.Publisher#unsubscribe}.
         * @memberof event.Publisher
         * @typedef UnsubscribeCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {event.SubscriptionResponse} [response] SubscriptionResponse
         */

        /**
         * Calls Unsubscribe.
         * @function unsubscribe
         * @memberof event.Publisher
         * @instance
         * @param {event.ISubscription} request Subscription message or plain object
         * @param {event.Publisher.UnsubscribeCallback} callback Node-style callback called with the error, if any, and SubscriptionResponse
         * @returns {undefined}
         * @variation 1
         */
        Publisher.prototype.unsubscribe = function unsubscribe(request, callback) {
            return this.rpcCall(unsubscribe, $root.event.Subscription, $root.event.SubscriptionResponse, request, callback);
        };

        /**
         * Calls Unsubscribe.
         * @function unsubscribe
         * @memberof event.Publisher
         * @instance
         * @param {event.ISubscription} request Subscription message or plain object
         * @returns {Promise<event.SubscriptionResponse>} Promise
         * @variation 2
         */

        return Publisher;
    })();

    event.EventMgr = (function() {

        /**
         * Constructs a new EventMgr service.
         * @memberof event
         * @classdesc Represents an EventMgr
         * @extends $protobuf.rpc.Service
         * @constructor
         * @param {$protobuf.RPCImpl} rpcImpl RPC implementation
         * @param {boolean} [requestDelimited=false] Whether requests are length-delimited
         * @param {boolean} [responseDelimited=false] Whether responses are length-delimited
         */
        function EventMgr(rpcImpl, requestDelimited, responseDelimited) {
            $protobuf.rpc.Service.call(this, rpcImpl, requestDelimited, responseDelimited);
        }

        (EventMgr.prototype = Object.create($protobuf.rpc.Service.prototype)).constructor = EventMgr;

        /**
         * Callback as used by {@link event.EventMgr#searchEvents}.
         * @memberof event.EventMgr
         * @typedef SearchEventsCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {event.RawEvents} [response] RawEvents
         */

        /**
         * Calls SearchEvents.
         * @function searchEvents
         * @memberof event.EventMgr
         * @instance
         * @param {event.IListEventsRequest} request ListEventsRequest message or plain object
         * @param {event.EventMgr.SearchEventsCallback} callback Node-style callback called with the error, if any, and RawEvents
         * @returns {undefined}
         * @variation 1
         */
        EventMgr.prototype.searchEvents = function searchEvents(request, callback) {
            return this.rpcCall(searchEvents, $root.event.ListEventsRequest, $root.event.RawEvents, request, callback);
        };

        /**
         * Calls SearchEvents.
         * @function searchEvents
         * @memberof event.EventMgr
         * @instance
         * @param {event.IListEventsRequest} request ListEventsRequest message or plain object
         * @returns {Promise<event.RawEvents>} Promise
         * @variation 2
         */

        /**
         * Callback as used by {@link event.EventMgr#createEvent}.
         * @memberof event.EventMgr
         * @typedef CreateEventCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {event.RawEvent} [response] RawEvent
         */

        /**
         * Calls CreateEvent.
         * @function createEvent
         * @memberof event.EventMgr
         * @instance
         * @param {event.IUserEvent} request UserEvent message or plain object
         * @param {event.EventMgr.CreateEventCallback} callback Node-style callback called with the error, if any, and RawEvent
         * @returns {undefined}
         * @variation 1
         */
        EventMgr.prototype.createEvent = function createEvent(request, callback) {
            return this.rpcCall(createEvent, $root.event.UserEvent, $root.event.RawEvent, request, callback);
        };

        /**
         * Calls CreateEvent.
         * @function createEvent
         * @memberof event.EventMgr
         * @instance
         * @param {event.IUserEvent} request UserEvent message or plain object
         * @returns {Promise<event.RawEvent>} Promise
         * @variation 2
         */

        /**
         * Callback as used by {@link event.EventMgr#subscribe}.
         * @memberof event.EventMgr
         * @typedef SubscribeCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {common.Empty} [response] Empty
         */

        /**
         * Calls Subscribe.
         * @function subscribe
         * @memberof event.EventMgr
         * @instance
         * @param {user.ISubscribeRequest} request SubscribeRequest message or plain object
         * @param {event.EventMgr.SubscribeCallback} callback Node-style callback called with the error, if any, and Empty
         * @returns {undefined}
         * @variation 1
         */
        EventMgr.prototype.subscribe = function subscribe(request, callback) {
            return this.rpcCall(subscribe, $root.user.SubscribeRequest, $root.common.Empty, request, callback);
        };

        /**
         * Calls Subscribe.
         * @function subscribe
         * @memberof event.EventMgr
         * @instance
         * @param {user.ISubscribeRequest} request SubscribeRequest message or plain object
         * @returns {Promise<common.Empty>} Promise
         * @variation 2
         */

        /**
         * Callback as used by {@link event.EventMgr#unsubscribe}.
         * @memberof event.EventMgr
         * @typedef UnsubscribeCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {common.Empty} [response] Empty
         */

        /**
         * Calls Unsubscribe.
         * @function unsubscribe
         * @memberof event.EventMgr
         * @instance
         * @param {user.ISubscribeRequest} request SubscribeRequest message or plain object
         * @param {event.EventMgr.UnsubscribeCallback} callback Node-style callback called with the error, if any, and Empty
         * @returns {undefined}
         * @variation 1
         */
        EventMgr.prototype.unsubscribe = function unsubscribe(request, callback) {
            return this.rpcCall(unsubscribe, $root.user.SubscribeRequest, $root.common.Empty, request, callback);
        };

        /**
         * Calls Unsubscribe.
         * @function unsubscribe
         * @memberof event.EventMgr
         * @instance
         * @param {user.ISubscribeRequest} request SubscribeRequest message or plain object
         * @returns {Promise<common.Empty>} Promise
         * @variation 2
         */

        /**
         * Callback as used by {@link event.EventMgr#readTopic}.
         * @memberof event.EventMgr
         * @typedef ReadTopicCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {common.Empty} [response] Empty
         */

        /**
         * Calls ReadTopic.
         * @function readTopic
         * @memberof event.EventMgr
         * @instance
         * @param {user.IReadTopicRequest} request ReadTopicRequest message or plain object
         * @param {event.EventMgr.ReadTopicCallback} callback Node-style callback called with the error, if any, and Empty
         * @returns {undefined}
         * @variation 1
         */
        EventMgr.prototype.readTopic = function readTopic(request, callback) {
            return this.rpcCall(readTopic, $root.user.ReadTopicRequest, $root.common.Empty, request, callback);
        };

        /**
         * Calls ReadTopic.
         * @function readTopic
         * @memberof event.EventMgr
         * @instance
         * @param {user.IReadTopicRequest} request ReadTopicRequest message or plain object
         * @returns {Promise<common.Empty>} Promise
         * @variation 2
         */

        /**
         * Callback as used by {@link event.EventMgr#searchTopics}.
         * @memberof event.EventMgr
         * @typedef SearchTopicsCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {user.ListTopicsResult} [response] ListTopicsResult
         */

        /**
         * Calls SearchTopics.
         * @function searchTopics
         * @memberof event.EventMgr
         * @instance
         * @param {user.IListTopicsRequest} request ListTopicsRequest message or plain object
         * @param {event.EventMgr.SearchTopicsCallback} callback Node-style callback called with the error, if any, and ListTopicsResult
         * @returns {undefined}
         * @variation 1
         */
        EventMgr.prototype.searchTopics = function searchTopics(request, callback) {
            return this.rpcCall(searchTopics, $root.user.ListTopicsRequest, $root.user.ListTopicsResult, request, callback);
        };

        /**
         * Calls SearchTopics.
         * @function searchTopics
         * @memberof event.EventMgr
         * @instance
         * @param {user.IListTopicsRequest} request ListTopicsRequest message or plain object
         * @returns {Promise<user.ListTopicsResult>} Promise
         * @variation 2
         */

        return EventMgr;
    })();

    event.ConversationEventReader = (function() {

        /**
         * Constructs a new ConversationEventReader service.
         * @memberof event
         * @classdesc Represents a ConversationEventReader
         * @extends $protobuf.rpc.Service
         * @constructor
         * @param {$protobuf.RPCImpl} rpcImpl RPC implementation
         * @param {boolean} [requestDelimited=false] Whether requests are length-delimited
         * @param {boolean} [responseDelimited=false] Whether responses are length-delimited
         */
        function ConversationEventReader(rpcImpl, requestDelimited, responseDelimited) {
            $protobuf.rpc.Service.call(this, rpcImpl, requestDelimited, responseDelimited);
        }

        (ConversationEventReader.prototype = Object.create($protobuf.rpc.Service.prototype)).constructor = ConversationEventReader;

        /**
         * Callback as used by {@link event.ConversationEventReader#sendMessage}.
         * @memberof event.ConversationEventReader
         * @typedef SendMessageCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {event.RawEvent} [response] RawEvent
         */

        /**
         * Calls SendMessage.
         * @function sendMessage
         * @memberof event.ConversationEventReader
         * @instance
         * @param {event.IRawEvent} request RawEvent message or plain object
         * @param {event.ConversationEventReader.SendMessageCallback} callback Node-style callback called with the error, if any, and RawEvent
         * @returns {undefined}
         * @variation 1
         */
        ConversationEventReader.prototype.sendMessage = function sendMessage(request, callback) {
            return this.rpcCall(sendMessage, $root.event.RawEvent, $root.event.RawEvent, request, callback);
        };

        /**
         * Calls SendMessage.
         * @function sendMessage
         * @memberof event.ConversationEventReader
         * @instance
         * @param {event.IRawEvent} request RawEvent message or plain object
         * @returns {Promise<event.RawEvent>} Promise
         * @variation 2
         */

        /**
         * Callback as used by {@link event.ConversationEventReader#listEvents}.
         * @memberof event.ConversationEventReader
         * @typedef ListEventsCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {event.RawEvents} [response] RawEvents
         */

        /**
         * Calls ListEvents.
         * @function listEvents
         * @memberof event.ConversationEventReader
         * @instance
         * @param {conversation.IListEventsRequest} request ListEventsRequest message or plain object
         * @param {event.ConversationEventReader.ListEventsCallback} callback Node-style callback called with the error, if any, and RawEvents
         * @returns {undefined}
         * @variation 1
         */
        ConversationEventReader.prototype.listEvents = function listEvents(request, callback) {
            return this.rpcCall(listEvents, $root.conversation.ListEventsRequest, $root.event.RawEvents, request, callback);
        };

        /**
         * Calls ListEvents.
         * @function listEvents
         * @memberof event.ConversationEventReader
         * @instance
         * @param {conversation.IListEventsRequest} request ListEventsRequest message or plain object
         * @returns {Promise<event.RawEvents>} Promise
         * @variation 2
         */

        /**
         * Callback as used by {@link event.ConversationEventReader#searchEvents}.
         * @memberof event.ConversationEventReader
         * @typedef SearchEventsCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {event.RawEvents} [response] RawEvents
         */

        /**
         * Calls SearchEvents.
         * @function searchEvents
         * @memberof event.ConversationEventReader
         * @instance
         * @param {conversation.ISearchMessageRequest} request SearchMessageRequest message or plain object
         * @param {event.ConversationEventReader.SearchEventsCallback} callback Node-style callback called with the error, if any, and RawEvents
         * @returns {undefined}
         * @variation 1
         */
        ConversationEventReader.prototype.searchEvents = function searchEvents(request, callback) {
            return this.rpcCall(searchEvents, $root.conversation.SearchMessageRequest, $root.event.RawEvents, request, callback);
        };

        /**
         * Calls SearchEvents.
         * @function searchEvents
         * @memberof event.ConversationEventReader
         * @instance
         * @param {conversation.ISearchMessageRequest} request SearchMessageRequest message or plain object
         * @returns {Promise<event.RawEvents>} Promise
         * @variation 2
         */

        return ConversationEventReader;
    })();

    event.AutomationEvent = (function() {

        /**
         * Properties of an AutomationEvent.
         * @memberof event
         * @interface IAutomationEvent
         * @property {common.IContext|null} [ctx] AutomationEvent ctx
         * @property {string|null} [account_id] AutomationEvent account_id
         * @property {string|null} [user_id] AutomationEvent user_id
         * @property {user.IAutomation|null} [automation] AutomationEvent automation
         * @property {user.ISession|null} [session] AutomationEvent session
         * @property {user.IUser|null} [user] AutomationEvent user
         * @property {event.IRawEvent|null} [event] AutomationEvent event
         */

        /**
         * Constructs a new AutomationEvent.
         * @memberof event
         * @classdesc Represents an AutomationEvent.
         * @implements IAutomationEvent
         * @constructor
         * @param {event.IAutomationEvent=} [p] Properties to set
         */
        function AutomationEvent(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * AutomationEvent ctx.
         * @member {common.IContext|null|undefined} ctx
         * @memberof event.AutomationEvent
         * @instance
         */
        AutomationEvent.prototype.ctx = null;

        /**
         * AutomationEvent account_id.
         * @member {string} account_id
         * @memberof event.AutomationEvent
         * @instance
         */
        AutomationEvent.prototype.account_id = "";

        /**
         * AutomationEvent user_id.
         * @member {string} user_id
         * @memberof event.AutomationEvent
         * @instance
         */
        AutomationEvent.prototype.user_id = "";

        /**
         * AutomationEvent automation.
         * @member {user.IAutomation|null|undefined} automation
         * @memberof event.AutomationEvent
         * @instance
         */
        AutomationEvent.prototype.automation = null;

        /**
         * AutomationEvent session.
         * @member {user.ISession|null|undefined} session
         * @memberof event.AutomationEvent
         * @instance
         */
        AutomationEvent.prototype.session = null;

        /**
         * AutomationEvent user.
         * @member {user.IUser|null|undefined} user
         * @memberof event.AutomationEvent
         * @instance
         */
        AutomationEvent.prototype.user = null;

        /**
         * AutomationEvent event.
         * @member {event.IRawEvent|null|undefined} event
         * @memberof event.AutomationEvent
         * @instance
         */
        AutomationEvent.prototype.event = null;

        return AutomationEvent;
    })();

    return event;
})();

export const fabikon = $root.fabikon = (() => {

    /**
     * Namespace fabikon.
     * @exports fabikon
     * @namespace
     */
    const fabikon = {};

    fabikon.FacebookPage = (function() {

        /**
         * Properties of a FacebookPage.
         * @memberof fabikon
         * @interface IFacebookPage
         * @property {string|null} [account_id] FacebookPage account_id
         * @property {string|null} [id] FacebookPage id
         * @property {number|Long|null} [created] FacebookPage created
         * @property {string|null} [picture_url] FacebookPage picture_url
         * @property {string|null} [name] FacebookPage name
         * @property {string|null} [access_token] FacebookPage access_token
         */

        /**
         * Constructs a new FacebookPage.
         * @memberof fabikon
         * @classdesc Represents a FacebookPage.
         * @implements IFacebookPage
         * @constructor
         * @param {fabikon.IFacebookPage=} [p] Properties to set
         */
        function FacebookPage(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * FacebookPage account_id.
         * @member {string} account_id
         * @memberof fabikon.FacebookPage
         * @instance
         */
        FacebookPage.prototype.account_id = "";

        /**
         * FacebookPage id.
         * @member {string} id
         * @memberof fabikon.FacebookPage
         * @instance
         */
        FacebookPage.prototype.id = "";

        /**
         * FacebookPage created.
         * @member {number|Long} created
         * @memberof fabikon.FacebookPage
         * @instance
         */
        FacebookPage.prototype.created = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

        /**
         * FacebookPage picture_url.
         * @member {string} picture_url
         * @memberof fabikon.FacebookPage
         * @instance
         */
        FacebookPage.prototype.picture_url = "";

        /**
         * FacebookPage name.
         * @member {string} name
         * @memberof fabikon.FacebookPage
         * @instance
         */
        FacebookPage.prototype.name = "";

        /**
         * FacebookPage access_token.
         * @member {string} access_token
         * @memberof fabikon.FacebookPage
         * @instance
         */
        FacebookPage.prototype.access_token = "";

        return FacebookPage;
    })();

    /**
     * Event enum.
     * @name fabikon.Event
     * @enum {string}
     * @property {number} FabikonRequested=1000 FabikonRequested value
     * @property {number} FabikonSynced=1001 FabikonSynced value
     * @property {number} FabikonJob=2 FabikonJob value
     * @property {number} FabikonCreateAccountRequested=0 FabikonCreateAccountRequested value
     */
    fabikon.Event = (function() {
        const valuesById = {}, values = Object.create(valuesById);
        values[valuesById[1000] = "FabikonRequested"] = 1000;
        values[valuesById[1001] = "FabikonSynced"] = 1001;
        values[valuesById[2] = "FabikonJob"] = 2;
        values[valuesById[0] = "FabikonCreateAccountRequested"] = 0;
        return values;
    })();

    fabikon.HttpChunk = (function() {

        /**
         * Properties of a HttpChunk.
         * @memberof fabikon
         * @interface IHttpChunk
         * @property {string|null} [id] HttpChunk id
         * @property {number|null} [chunk_id] HttpChunk chunk_id
         * @property {Uint8Array|null} [data] HttpChunk data
         * @property {number|null} [chunk_size] HttpChunk chunk_size
         */

        /**
         * Constructs a new HttpChunk.
         * @memberof fabikon
         * @classdesc Represents a HttpChunk.
         * @implements IHttpChunk
         * @constructor
         * @param {fabikon.IHttpChunk=} [p] Properties to set
         */
        function HttpChunk(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * HttpChunk id.
         * @member {string} id
         * @memberof fabikon.HttpChunk
         * @instance
         */
        HttpChunk.prototype.id = "";

        /**
         * HttpChunk chunk_id.
         * @member {number} chunk_id
         * @memberof fabikon.HttpChunk
         * @instance
         */
        HttpChunk.prototype.chunk_id = 0;

        /**
         * HttpChunk data.
         * @member {Uint8Array} data
         * @memberof fabikon.HttpChunk
         * @instance
         */
        HttpChunk.prototype.data = $util.newBuffer([]);

        /**
         * HttpChunk chunk_size.
         * @member {number} chunk_size
         * @memberof fabikon.HttpChunk
         * @instance
         */
        HttpChunk.prototype.chunk_size = 0;

        return HttpChunk;
    })();

    fabikon.Job = (function() {

        /**
         * Properties of a Job.
         * @memberof fabikon
         * @interface IJob
         * @property {string|null} [topic] Job topic
         * @property {number|null} [partition] Job partition
         * @property {number|Long|null} [offset] Job offset
         * @property {string|null} [type] Job type
         * @property {string|null} [request_id] Job request_id
         */

        /**
         * Constructs a new Job.
         * @memberof fabikon
         * @classdesc Represents a Job.
         * @implements IJob
         * @constructor
         * @param {fabikon.IJob=} [p] Properties to set
         */
        function Job(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * Job topic.
         * @member {string} topic
         * @memberof fabikon.Job
         * @instance
         */
        Job.prototype.topic = "";

        /**
         * Job partition.
         * @member {number} partition
         * @memberof fabikon.Job
         * @instance
         */
        Job.prototype.partition = 0;

        /**
         * Job offset.
         * @member {number|Long} offset
         * @memberof fabikon.Job
         * @instance
         */
        Job.prototype.offset = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

        /**
         * Job type.
         * @member {string} type
         * @memberof fabikon.Job
         * @instance
         */
        Job.prototype.type = "";

        /**
         * Job request_id.
         * @member {string} request_id
         * @memberof fabikon.Job
         * @instance
         */
        Job.prototype.request_id = "";

        return Job;
    })();

    /**
     * JobType enum.
     * @name fabikon.JobType
     * @enum {string}
     * @property {number} facebook_hook=0 facebook_hook value
     * @property {number} subiz_event=3 subiz_event value
     */
    fabikon.JobType = (function() {
        const valuesById = {}, values = Object.create(valuesById);
        values[valuesById[0] = "facebook_hook"] = 0;
        values[valuesById[3] = "subiz_event"] = 3;
        return values;
    })();

    fabikon.FbWebhookEvent = (function() {

        /**
         * Properties of a FbWebhookEvent.
         * @memberof fabikon
         * @interface IFbWebhookEvent
         * @property {string|null} [object] FbWebhookEvent object
         * @property {Array.<fabikon.IFbEntry>|null} [entry] FbWebhookEvent entry
         */

        /**
         * Constructs a new FbWebhookEvent.
         * @memberof fabikon
         * @classdesc Represents a FbWebhookEvent.
         * @implements IFbWebhookEvent
         * @constructor
         * @param {fabikon.IFbWebhookEvent=} [p] Properties to set
         */
        function FbWebhookEvent(p) {
            this.entry = [];
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * FbWebhookEvent object.
         * @member {string} object
         * @memberof fabikon.FbWebhookEvent
         * @instance
         */
        FbWebhookEvent.prototype.object = "";

        /**
         * FbWebhookEvent entry.
         * @member {Array.<fabikon.IFbEntry>} entry
         * @memberof fabikon.FbWebhookEvent
         * @instance
         */
        FbWebhookEvent.prototype.entry = $util.emptyArray;

        return FbWebhookEvent;
    })();

    fabikon.FbEntry = (function() {

        /**
         * Properties of a FbEntry.
         * @memberof fabikon
         * @interface IFbEntry
         * @property {string|null} [id] FbEntry id
         * @property {number|Long|null} [time] FbEntry time
         * @property {Array.<fabikon.IFbMessaging>|null} [messaging] FbEntry messaging
         */

        /**
         * Constructs a new FbEntry.
         * @memberof fabikon
         * @classdesc Represents a FbEntry.
         * @implements IFbEntry
         * @constructor
         * @param {fabikon.IFbEntry=} [p] Properties to set
         */
        function FbEntry(p) {
            this.messaging = [];
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * FbEntry id.
         * @member {string} id
         * @memberof fabikon.FbEntry
         * @instance
         */
        FbEntry.prototype.id = "";

        /**
         * FbEntry time.
         * @member {number|Long} time
         * @memberof fabikon.FbEntry
         * @instance
         */
        FbEntry.prototype.time = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

        /**
         * FbEntry messaging.
         * @member {Array.<fabikon.IFbMessaging>} messaging
         * @memberof fabikon.FbEntry
         * @instance
         */
        FbEntry.prototype.messaging = $util.emptyArray;

        return FbEntry;
    })();

    fabikon.FbMessaging = (function() {

        /**
         * Properties of a FbMessaging.
         * @memberof fabikon
         * @interface IFbMessaging
         * @property {fabikon.IFbSender|null} [sender] FbMessaging sender
         * @property {fabikon.IFbRecipient|null} [recipient] FbMessaging recipient
         * @property {number|Long|null} [timestamp] FbMessaging timestamp
         * @property {fabikon.IFbMessage|null} [message] FbMessaging message
         * @property {fabikon.IFbDelivery|null} [delivery] FbMessaging delivery
         * @property {string|null} [message_type] FbMessaging message_type
         */

        /**
         * Constructs a new FbMessaging.
         * @memberof fabikon
         * @classdesc Represents a FbMessaging.
         * @implements IFbMessaging
         * @constructor
         * @param {fabikon.IFbMessaging=} [p] Properties to set
         */
        function FbMessaging(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * FbMessaging sender.
         * @member {fabikon.IFbSender|null|undefined} sender
         * @memberof fabikon.FbMessaging
         * @instance
         */
        FbMessaging.prototype.sender = null;

        /**
         * FbMessaging recipient.
         * @member {fabikon.IFbRecipient|null|undefined} recipient
         * @memberof fabikon.FbMessaging
         * @instance
         */
        FbMessaging.prototype.recipient = null;

        /**
         * FbMessaging timestamp.
         * @member {number|Long} timestamp
         * @memberof fabikon.FbMessaging
         * @instance
         */
        FbMessaging.prototype.timestamp = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

        /**
         * FbMessaging message.
         * @member {fabikon.IFbMessage|null|undefined} message
         * @memberof fabikon.FbMessaging
         * @instance
         */
        FbMessaging.prototype.message = null;

        /**
         * FbMessaging delivery.
         * @member {fabikon.IFbDelivery|null|undefined} delivery
         * @memberof fabikon.FbMessaging
         * @instance
         */
        FbMessaging.prototype.delivery = null;

        /**
         * FbMessaging message_type.
         * @member {string} message_type
         * @memberof fabikon.FbMessaging
         * @instance
         */
        FbMessaging.prototype.message_type = "";

        return FbMessaging;
    })();

    fabikon.FbMessage = (function() {

        /**
         * Properties of a FbMessage.
         * @memberof fabikon
         * @interface IFbMessage
         * @property {string|null} [mid] FbMessage mid
         * @property {string|null} [text] FbMessage text
         * @property {Array.<fabikon.IFbAttachment>|null} [attachments] FbMessage attachments
         * @property {fabikon.IFbRead|null} [read] FbMessage read
         * @property {fabikon.IFbAttachment|null} [attachment] FbMessage attachment
         */

        /**
         * Constructs a new FbMessage.
         * @memberof fabikon
         * @classdesc Represents a FbMessage.
         * @implements IFbMessage
         * @constructor
         * @param {fabikon.IFbMessage=} [p] Properties to set
         */
        function FbMessage(p) {
            this.attachments = [];
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * FbMessage mid.
         * @member {string} mid
         * @memberof fabikon.FbMessage
         * @instance
         */
        FbMessage.prototype.mid = "";

        /**
         * FbMessage text.
         * @member {string} text
         * @memberof fabikon.FbMessage
         * @instance
         */
        FbMessage.prototype.text = "";

        /**
         * FbMessage attachments.
         * @member {Array.<fabikon.IFbAttachment>} attachments
         * @memberof fabikon.FbMessage
         * @instance
         */
        FbMessage.prototype.attachments = $util.emptyArray;

        /**
         * FbMessage read.
         * @member {fabikon.IFbRead|null|undefined} read
         * @memberof fabikon.FbMessage
         * @instance
         */
        FbMessage.prototype.read = null;

        /**
         * FbMessage attachment.
         * @member {fabikon.IFbAttachment|null|undefined} attachment
         * @memberof fabikon.FbMessage
         * @instance
         */
        FbMessage.prototype.attachment = null;

        return FbMessage;
    })();

    fabikon.FbSender = (function() {

        /**
         * Properties of a FbSender.
         * @memberof fabikon
         * @interface IFbSender
         * @property {string|null} [id] FbSender id
         */

        /**
         * Constructs a new FbSender.
         * @memberof fabikon
         * @classdesc Represents a FbSender.
         * @implements IFbSender
         * @constructor
         * @param {fabikon.IFbSender=} [p] Properties to set
         */
        function FbSender(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * FbSender id.
         * @member {string} id
         * @memberof fabikon.FbSender
         * @instance
         */
        FbSender.prototype.id = "";

        return FbSender;
    })();

    fabikon.FbRecipient = (function() {

        /**
         * Properties of a FbRecipient.
         * @memberof fabikon
         * @interface IFbRecipient
         * @property {string|null} [id] FbRecipient id
         */

        /**
         * Constructs a new FbRecipient.
         * @memberof fabikon
         * @classdesc Represents a FbRecipient.
         * @implements IFbRecipient
         * @constructor
         * @param {fabikon.IFbRecipient=} [p] Properties to set
         */
        function FbRecipient(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * FbRecipient id.
         * @member {string} id
         * @memberof fabikon.FbRecipient
         * @instance
         */
        FbRecipient.prototype.id = "";

        return FbRecipient;
    })();

    fabikon.FbAttachment = (function() {

        /**
         * Properties of a FbAttachment.
         * @memberof fabikon
         * @interface IFbAttachment
         * @property {string|null} [type] FbAttachment type
         * @property {fabikon.IFbPayload|null} [payload] FbAttachment payload
         * @property {string|null} [title] FbAttachment title
         * @property {string|null} [URL] FbAttachment URL
         */

        /**
         * Constructs a new FbAttachment.
         * @memberof fabikon
         * @classdesc Represents a FbAttachment.
         * @implements IFbAttachment
         * @constructor
         * @param {fabikon.IFbAttachment=} [p] Properties to set
         */
        function FbAttachment(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * FbAttachment type.
         * @member {string} type
         * @memberof fabikon.FbAttachment
         * @instance
         */
        FbAttachment.prototype.type = "";

        /**
         * FbAttachment payload.
         * @member {fabikon.IFbPayload|null|undefined} payload
         * @memberof fabikon.FbAttachment
         * @instance
         */
        FbAttachment.prototype.payload = null;

        /**
         * FbAttachment title.
         * @member {string} title
         * @memberof fabikon.FbAttachment
         * @instance
         */
        FbAttachment.prototype.title = "";

        /**
         * FbAttachment URL.
         * @member {string} URL
         * @memberof fabikon.FbAttachment
         * @instance
         */
        FbAttachment.prototype.URL = "";

        return FbAttachment;
    })();

    fabikon.FbPayload = (function() {

        /**
         * Properties of a FbPayload.
         * @memberof fabikon
         * @interface IFbPayload
         * @property {string|null} [url] FbPayload url
         * @property {boolean|null} [is_reuseable] FbPayload is_reuseable
         */

        /**
         * Constructs a new FbPayload.
         * @memberof fabikon
         * @classdesc Represents a FbPayload.
         * @implements IFbPayload
         * @constructor
         * @param {fabikon.IFbPayload=} [p] Properties to set
         */
        function FbPayload(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * FbPayload url.
         * @member {string} url
         * @memberof fabikon.FbPayload
         * @instance
         */
        FbPayload.prototype.url = "";

        /**
         * FbPayload is_reuseable.
         * @member {boolean} is_reuseable
         * @memberof fabikon.FbPayload
         * @instance
         */
        FbPayload.prototype.is_reuseable = false;

        return FbPayload;
    })();

    fabikon.FbRead = (function() {

        /**
         * Properties of a FbRead.
         * @memberof fabikon
         * @interface IFbRead
         * @property {number|Long|null} [watermark] FbRead watermark
         * @property {number|null} [seq] FbRead seq
         */

        /**
         * Constructs a new FbRead.
         * @memberof fabikon
         * @classdesc Represents a FbRead.
         * @implements IFbRead
         * @constructor
         * @param {fabikon.IFbRead=} [p] Properties to set
         */
        function FbRead(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * FbRead watermark.
         * @member {number|Long} watermark
         * @memberof fabikon.FbRead
         * @instance
         */
        FbRead.prototype.watermark = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

        /**
         * FbRead seq.
         * @member {number} seq
         * @memberof fabikon.FbRead
         * @instance
         */
        FbRead.prototype.seq = 0;

        return FbRead;
    })();

    fabikon.FbDelivery = (function() {

        /**
         * Properties of a FbDelivery.
         * @memberof fabikon
         * @interface IFbDelivery
         * @property {Array.<string>|null} [mids] FbDelivery mids
         * @property {number|Long|null} [watermark] FbDelivery watermark
         * @property {number|null} [seq] FbDelivery seq
         */

        /**
         * Constructs a new FbDelivery.
         * @memberof fabikon
         * @classdesc Represents a FbDelivery.
         * @implements IFbDelivery
         * @constructor
         * @param {fabikon.IFbDelivery=} [p] Properties to set
         */
        function FbDelivery(p) {
            this.mids = [];
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * FbDelivery mids.
         * @member {Array.<string>} mids
         * @memberof fabikon.FbDelivery
         * @instance
         */
        FbDelivery.prototype.mids = $util.emptyArray;

        /**
         * FbDelivery watermark.
         * @member {number|Long} watermark
         * @memberof fabikon.FbDelivery
         * @instance
         */
        FbDelivery.prototype.watermark = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

        /**
         * FbDelivery seq.
         * @member {number} seq
         * @memberof fabikon.FbDelivery
         * @instance
         */
        FbDelivery.prototype.seq = 0;

        return FbDelivery;
    })();

    fabikon.FbSendResponse = (function() {

        /**
         * Properties of a FbSendResponse.
         * @memberof fabikon
         * @interface IFbSendResponse
         * @property {string|null} [recipient_id] FbSendResponse recipient_id
         * @property {string|null} [message_id] FbSendResponse message_id
         */

        /**
         * Constructs a new FbSendResponse.
         * @memberof fabikon
         * @classdesc Represents a FbSendResponse.
         * @implements IFbSendResponse
         * @constructor
         * @param {fabikon.IFbSendResponse=} [p] Properties to set
         */
        function FbSendResponse(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * FbSendResponse recipient_id.
         * @member {string} recipient_id
         * @memberof fabikon.FbSendResponse
         * @instance
         */
        FbSendResponse.prototype.recipient_id = "";

        /**
         * FbSendResponse message_id.
         * @member {string} message_id
         * @memberof fabikon.FbSendResponse
         * @instance
         */
        FbSendResponse.prototype.message_id = "";

        return FbSendResponse;
    })();

    fabikon.Conversation = (function() {

        /**
         * Properties of a Conversation.
         * @memberof fabikon
         * @interface IConversation
         * @property {string|null} [account_id] Conversation account_id
         * @property {string|null} [id] Conversation id
         * @property {string|null} [page_id] Conversation page_id
         * @property {number|Long|null} [created] Conversation created
         * @property {string|null} [fbuserid] Conversation fbuserid
         * @property {string|null} [sbuserid] Conversation sbuserid
         */

        /**
         * Constructs a new Conversation.
         * @memberof fabikon
         * @classdesc Represents a Conversation.
         * @implements IConversation
         * @constructor
         * @param {fabikon.IConversation=} [p] Properties to set
         */
        function Conversation(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * Conversation account_id.
         * @member {string} account_id
         * @memberof fabikon.Conversation
         * @instance
         */
        Conversation.prototype.account_id = "";

        /**
         * Conversation id.
         * @member {string} id
         * @memberof fabikon.Conversation
         * @instance
         */
        Conversation.prototype.id = "";

        /**
         * Conversation page_id.
         * @member {string} page_id
         * @memberof fabikon.Conversation
         * @instance
         */
        Conversation.prototype.page_id = "";

        /**
         * Conversation created.
         * @member {number|Long} created
         * @memberof fabikon.Conversation
         * @instance
         */
        Conversation.prototype.created = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

        /**
         * Conversation fbuserid.
         * @member {string} fbuserid
         * @memberof fabikon.Conversation
         * @instance
         */
        Conversation.prototype.fbuserid = "";

        /**
         * Conversation sbuserid.
         * @member {string} sbuserid
         * @memberof fabikon.Conversation
         * @instance
         */
        Conversation.prototype.sbuserid = "";

        return Conversation;
    })();

    fabikon.CurrentConvo = (function() {

        /**
         * Properties of a CurrentConvo.
         * @memberof fabikon
         * @interface ICurrentConvo
         * @property {string|null} [account_id] CurrentConvo account_id
         * @property {string|null} [convo_id] CurrentConvo convo_id
         * @property {string|null} [page_id] CurrentConvo page_id
         * @property {string|null} [fbuser_id] CurrentConvo fbuser_id
         * @property {string|null} [sbuser_id] CurrentConvo sbuser_id
         */

        /**
         * Constructs a new CurrentConvo.
         * @memberof fabikon
         * @classdesc Represents a CurrentConvo.
         * @implements ICurrentConvo
         * @constructor
         * @param {fabikon.ICurrentConvo=} [p] Properties to set
         */
        function CurrentConvo(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * CurrentConvo account_id.
         * @member {string} account_id
         * @memberof fabikon.CurrentConvo
         * @instance
         */
        CurrentConvo.prototype.account_id = "";

        /**
         * CurrentConvo convo_id.
         * @member {string} convo_id
         * @memberof fabikon.CurrentConvo
         * @instance
         */
        CurrentConvo.prototype.convo_id = "";

        /**
         * CurrentConvo page_id.
         * @member {string} page_id
         * @memberof fabikon.CurrentConvo
         * @instance
         */
        CurrentConvo.prototype.page_id = "";

        /**
         * CurrentConvo fbuser_id.
         * @member {string} fbuser_id
         * @memberof fabikon.CurrentConvo
         * @instance
         */
        CurrentConvo.prototype.fbuser_id = "";

        /**
         * CurrentConvo sbuser_id.
         * @member {string} sbuser_id
         * @memberof fabikon.CurrentConvo
         * @instance
         */
        CurrentConvo.prototype.sbuser_id = "";

        return CurrentConvo;
    })();

    fabikon.UserAvail = (function() {

        /**
         * Properties of a UserAvail.
         * @memberof fabikon
         * @interface IUserAvail
         * @property {string|null} [account_id] UserAvail account_id
         * @property {string|null} [fbuser_id] UserAvail fbuser_id
         * @property {boolean|null} [availability] UserAvail availability
         * @property {number|Long|null} [updated] UserAvail updated
         */

        /**
         * Constructs a new UserAvail.
         * @memberof fabikon
         * @classdesc Represents a UserAvail.
         * @implements IUserAvail
         * @constructor
         * @param {fabikon.IUserAvail=} [p] Properties to set
         */
        function UserAvail(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * UserAvail account_id.
         * @member {string} account_id
         * @memberof fabikon.UserAvail
         * @instance
         */
        UserAvail.prototype.account_id = "";

        /**
         * UserAvail fbuser_id.
         * @member {string} fbuser_id
         * @memberof fabikon.UserAvail
         * @instance
         */
        UserAvail.prototype.fbuser_id = "";

        /**
         * UserAvail availability.
         * @member {boolean} availability
         * @memberof fabikon.UserAvail
         * @instance
         */
        UserAvail.prototype.availability = false;

        /**
         * UserAvail updated.
         * @member {number|Long} updated
         * @memberof fabikon.UserAvail
         * @instance
         */
        UserAvail.prototype.updated = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

        return UserAvail;
    })();

    fabikon.FbPagePicture = (function() {

        /**
         * Properties of a FbPagePicture.
         * @memberof fabikon
         * @interface IFbPagePicture
         * @property {fabikon.IFbPagePictureData|null} [data] FbPagePicture data
         */

        /**
         * Constructs a new FbPagePicture.
         * @memberof fabikon
         * @classdesc Represents a FbPagePicture.
         * @implements IFbPagePicture
         * @constructor
         * @param {fabikon.IFbPagePicture=} [p] Properties to set
         */
        function FbPagePicture(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * FbPagePicture data.
         * @member {fabikon.IFbPagePictureData|null|undefined} data
         * @memberof fabikon.FbPagePicture
         * @instance
         */
        FbPagePicture.prototype.data = null;

        return FbPagePicture;
    })();

    fabikon.FbPagePictureData = (function() {

        /**
         * Properties of a FbPagePictureData.
         * @memberof fabikon
         * @interface IFbPagePictureData
         * @property {string|null} [url] FbPagePictureData url
         */

        /**
         * Constructs a new FbPagePictureData.
         * @memberof fabikon
         * @classdesc Represents a FbPagePictureData.
         * @implements IFbPagePictureData
         * @constructor
         * @param {fabikon.IFbPagePictureData=} [p] Properties to set
         */
        function FbPagePictureData(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * FbPagePictureData url.
         * @member {string} url
         * @memberof fabikon.FbPagePictureData
         * @instance
         */
        FbPagePictureData.prototype.url = "";

        return FbPagePictureData;
    })();

    fabikon.FbPageEntry = (function() {

        /**
         * Properties of a FbPageEntry.
         * @memberof fabikon
         * @interface IFbPageEntry
         * @property {string|null} [name] FbPageEntry name
         * @property {fabikon.IFbPagePicture|null} [picture] FbPageEntry picture
         * @property {string|null} [id] FbPageEntry id
         * @property {string|null} [access_token] FbPageEntry access_token
         */

        /**
         * Constructs a new FbPageEntry.
         * @memberof fabikon
         * @classdesc Represents a FbPageEntry.
         * @implements IFbPageEntry
         * @constructor
         * @param {fabikon.IFbPageEntry=} [p] Properties to set
         */
        function FbPageEntry(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * FbPageEntry name.
         * @member {string} name
         * @memberof fabikon.FbPageEntry
         * @instance
         */
        FbPageEntry.prototype.name = "";

        /**
         * FbPageEntry picture.
         * @member {fabikon.IFbPagePicture|null|undefined} picture
         * @memberof fabikon.FbPageEntry
         * @instance
         */
        FbPageEntry.prototype.picture = null;

        /**
         * FbPageEntry id.
         * @member {string} id
         * @memberof fabikon.FbPageEntry
         * @instance
         */
        FbPageEntry.prototype.id = "";

        /**
         * FbPageEntry access_token.
         * @member {string} access_token
         * @memberof fabikon.FbPageEntry
         * @instance
         */
        FbPageEntry.prototype.access_token = "";

        return FbPageEntry;
    })();

    fabikon.FbPageRet = (function() {

        /**
         * Properties of a FbPageRet.
         * @memberof fabikon
         * @interface IFbPageRet
         * @property {Array.<fabikon.IFbPageEntry>|null} [data] FbPageRet data
         */

        /**
         * Constructs a new FbPageRet.
         * @memberof fabikon
         * @classdesc Represents a FbPageRet.
         * @implements IFbPageRet
         * @constructor
         * @param {fabikon.IFbPageRet=} [p] Properties to set
         */
        function FbPageRet(p) {
            this.data = [];
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * FbPageRet data.
         * @member {Array.<fabikon.IFbPageEntry>} data
         * @memberof fabikon.FbPageRet
         * @instance
         */
        FbPageRet.prototype.data = $util.emptyArray;

        return FbPageRet;
    })();

    fabikon.FacebookUser = (function() {

        /**
         * Properties of a FacebookUser.
         * @memberof fabikon
         * @interface IFacebookUser
         * @property {string|null} [first_name] FacebookUser first_name
         * @property {string|null} [last_name] FacebookUser last_name
         * @property {string|null} [profile_pic] FacebookUser profile_pic
         * @property {string|null} [locale] FacebookUser locale
         * @property {number|null} [timezone] FacebookUser timezone
         * @property {string|null} [gender] FacebookUser gender
         */

        /**
         * Constructs a new FacebookUser.
         * @memberof fabikon
         * @classdesc Represents a FacebookUser.
         * @implements IFacebookUser
         * @constructor
         * @param {fabikon.IFacebookUser=} [p] Properties to set
         */
        function FacebookUser(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * FacebookUser first_name.
         * @member {string} first_name
         * @memberof fabikon.FacebookUser
         * @instance
         */
        FacebookUser.prototype.first_name = "";

        /**
         * FacebookUser last_name.
         * @member {string} last_name
         * @memberof fabikon.FacebookUser
         * @instance
         */
        FacebookUser.prototype.last_name = "";

        /**
         * FacebookUser profile_pic.
         * @member {string} profile_pic
         * @memberof fabikon.FacebookUser
         * @instance
         */
        FacebookUser.prototype.profile_pic = "";

        /**
         * FacebookUser locale.
         * @member {string} locale
         * @memberof fabikon.FacebookUser
         * @instance
         */
        FacebookUser.prototype.locale = "";

        /**
         * FacebookUser timezone.
         * @member {number} timezone
         * @memberof fabikon.FacebookUser
         * @instance
         */
        FacebookUser.prototype.timezone = 0;

        /**
         * FacebookUser gender.
         * @member {string} gender
         * @memberof fabikon.FacebookUser
         * @instance
         */
        FacebookUser.prototype.gender = "";

        return FacebookUser;
    })();

    fabikon.Fb2SbEvent = (function() {

        /**
         * Properties of a Fb2SbEvent.
         * @memberof fabikon
         * @interface IFb2SbEvent
         * @property {string|null} [fbmid] Fb2SbEvent fbmid
         * @property {string|null} [account_id] Fb2SbEvent account_id
         * @property {string|null} [conversation_id] Fb2SbEvent conversation_id
         * @property {string|null} [sbmid] Fb2SbEvent sbmid
         * @property {string|null} [page_id] Fb2SbEvent page_id
         */

        /**
         * Constructs a new Fb2SbEvent.
         * @memberof fabikon
         * @classdesc Represents a Fb2SbEvent.
         * @implements IFb2SbEvent
         * @constructor
         * @param {fabikon.IFb2SbEvent=} [p] Properties to set
         */
        function Fb2SbEvent(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * Fb2SbEvent fbmid.
         * @member {string} fbmid
         * @memberof fabikon.Fb2SbEvent
         * @instance
         */
        Fb2SbEvent.prototype.fbmid = "";

        /**
         * Fb2SbEvent account_id.
         * @member {string} account_id
         * @memberof fabikon.Fb2SbEvent
         * @instance
         */
        Fb2SbEvent.prototype.account_id = "";

        /**
         * Fb2SbEvent conversation_id.
         * @member {string} conversation_id
         * @memberof fabikon.Fb2SbEvent
         * @instance
         */
        Fb2SbEvent.prototype.conversation_id = "";

        /**
         * Fb2SbEvent sbmid.
         * @member {string} sbmid
         * @memberof fabikon.Fb2SbEvent
         * @instance
         */
        Fb2SbEvent.prototype.sbmid = "";

        /**
         * Fb2SbEvent page_id.
         * @member {string} page_id
         * @memberof fabikon.Fb2SbEvent
         * @instance
         */
        Fb2SbEvent.prototype.page_id = "";

        return Fb2SbEvent;
    })();

    fabikon.LongLivedAccessToken = (function() {

        /**
         * Properties of a LongLivedAccessToken.
         * @memberof fabikon
         * @interface ILongLivedAccessToken
         * @property {string|null} [access_token] LongLivedAccessToken access_token
         * @property {string|null} [token_type] LongLivedAccessToken token_type
         * @property {number|null} [expires_in] LongLivedAccessToken expires_in
         */

        /**
         * Constructs a new LongLivedAccessToken.
         * @memberof fabikon
         * @classdesc Represents a LongLivedAccessToken.
         * @implements ILongLivedAccessToken
         * @constructor
         * @param {fabikon.ILongLivedAccessToken=} [p] Properties to set
         */
        function LongLivedAccessToken(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * LongLivedAccessToken access_token.
         * @member {string} access_token
         * @memberof fabikon.LongLivedAccessToken
         * @instance
         */
        LongLivedAccessToken.prototype.access_token = "";

        /**
         * LongLivedAccessToken token_type.
         * @member {string} token_type
         * @memberof fabikon.LongLivedAccessToken
         * @instance
         */
        LongLivedAccessToken.prototype.token_type = "";

        /**
         * LongLivedAccessToken expires_in.
         * @member {number} expires_in
         * @memberof fabikon.LongLivedAccessToken
         * @instance
         */
        LongLivedAccessToken.prototype.expires_in = 0;

        return LongLivedAccessToken;
    })();

    fabikon.SubscribeRet = (function() {

        /**
         * Properties of a SubscribeRet.
         * @memberof fabikon
         * @interface ISubscribeRet
         * @property {boolean|null} [success] SubscribeRet success
         */

        /**
         * Constructs a new SubscribeRet.
         * @memberof fabikon
         * @classdesc Represents a SubscribeRet.
         * @implements ISubscribeRet
         * @constructor
         * @param {fabikon.ISubscribeRet=} [p] Properties to set
         */
        function SubscribeRet(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * SubscribeRet success.
         * @member {boolean} success
         * @memberof fabikon.SubscribeRet
         * @instance
         */
        SubscribeRet.prototype.success = false;

        return SubscribeRet;
    })();

    return fabikon;
})();

export const file = $root.file = (() => {

    /**
     * Namespace file.
     * @exports file
     * @namespace
     */
    const file = {};

    file.AllType = (function() {

        /**
         * Properties of an AllType.
         * @memberof file
         * @interface IAllType
         * @property {file.IFileHeader|null} [fh] AllType fh
         * @property {file.IPresignResult|null} [ps] AllType ps
         * @property {file.IFile|null} [f] AllType f
         */

        /**
         * Constructs a new AllType.
         * @memberof file
         * @classdesc Represents an AllType.
         * @implements IAllType
         * @constructor
         * @param {file.IAllType=} [p] Properties to set
         */
        function AllType(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * AllType fh.
         * @member {file.IFileHeader|null|undefined} fh
         * @memberof file.AllType
         * @instance
         */
        AllType.prototype.fh = null;

        /**
         * AllType ps.
         * @member {file.IPresignResult|null|undefined} ps
         * @memberof file.AllType
         * @instance
         */
        AllType.prototype.ps = null;

        /**
         * AllType f.
         * @member {file.IFile|null|undefined} f
         * @memberof file.AllType
         * @instance
         */
        AllType.prototype.f = null;

        return AllType;
    })();

    file.MyServer = (function() {

        /**
         * Constructs a new MyServer service.
         * @memberof file
         * @classdesc Represents a MyServer
         * @extends $protobuf.rpc.Service
         * @constructor
         * @param {$protobuf.RPCImpl} rpcImpl RPC implementation
         * @param {boolean} [requestDelimited=false] Whether requests are length-delimited
         * @param {boolean} [responseDelimited=false] Whether responses are length-delimited
         */
        function MyServer(rpcImpl, requestDelimited, responseDelimited) {
            $protobuf.rpc.Service.call(this, rpcImpl, requestDelimited, responseDelimited);
        }

        (MyServer.prototype = Object.create($protobuf.rpc.Service.prototype)).constructor = MyServer;

        /**
         * Callback as used by {@link file.MyServer#do_}.
         * @memberof file.MyServer
         * @typedef DoCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {file.AllType} [response] AllType
         */

        /**
         * Calls Do.
         * @function do
         * @memberof file.MyServer
         * @instance
         * @param {file.IAllType} request AllType message or plain object
         * @param {file.MyServer.DoCallback} callback Node-style callback called with the error, if any, and AllType
         * @returns {undefined}
         * @variation 1
         */
        MyServer.prototype["do"] = function do_(request, callback) {
            return this.rpcCall(do_, $root.file.AllType, $root.file.AllType, request, callback);
        };

        /**
         * Calls Do.
         * @function do
         * @memberof file.MyServer
         * @instance
         * @param {file.IAllType} request AllType message or plain object
         * @returns {Promise<file.AllType>} Promise
         * @variation 2
         */

        return MyServer;
    })();

    file.FileMgr = (function() {

        /**
         * Constructs a new FileMgr service.
         * @memberof file
         * @classdesc Represents a FileMgr
         * @extends $protobuf.rpc.Service
         * @constructor
         * @param {$protobuf.RPCImpl} rpcImpl RPC implementation
         * @param {boolean} [requestDelimited=false] Whether requests are length-delimited
         * @param {boolean} [responseDelimited=false] Whether responses are length-delimited
         */
        function FileMgr(rpcImpl, requestDelimited, responseDelimited) {
            $protobuf.rpc.Service.call(this, rpcImpl, requestDelimited, responseDelimited);
        }

        (FileMgr.prototype = Object.create($protobuf.rpc.Service.prototype)).constructor = FileMgr;

        /**
         * Callback as used by {@link file.FileMgr#presign}.
         * @memberof file.FileMgr
         * @typedef PresignCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {file.PresignResult} [response] PresignResult
         */

        /**
         * Calls Presign.
         * @function presign
         * @memberof file.FileMgr
         * @instance
         * @param {file.IFileHeader} request FileHeader message or plain object
         * @param {file.FileMgr.PresignCallback} callback Node-style callback called with the error, if any, and PresignResult
         * @returns {undefined}
         * @variation 1
         */
        FileMgr.prototype.presign = function presign(request, callback) {
            return this.rpcCall(presign, $root.file.FileHeader, $root.file.PresignResult, request, callback);
        };

        /**
         * Calls Presign.
         * @function presign
         * @memberof file.FileMgr
         * @instance
         * @param {file.IFileHeader} request FileHeader message or plain object
         * @returns {Promise<file.PresignResult>} Promise
         * @variation 2
         */

        /**
         * Callback as used by {@link file.FileMgr#read}.
         * @memberof file.FileMgr
         * @typedef ReadCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {file.File} [response] File
         */

        /**
         * Calls Read.
         * @function read
         * @memberof file.FileMgr
         * @instance
         * @param {file.IFileRequest} request FileRequest message or plain object
         * @param {file.FileMgr.ReadCallback} callback Node-style callback called with the error, if any, and File
         * @returns {undefined}
         * @variation 1
         */
        FileMgr.prototype.read = function read(request, callback) {
            return this.rpcCall(read, $root.file.FileRequest, $root.file.File, request, callback);
        };

        /**
         * Calls Read.
         * @function read
         * @memberof file.FileMgr
         * @instance
         * @param {file.IFileRequest} request FileRequest message or plain object
         * @returns {Promise<file.File>} Promise
         * @variation 2
         */

        /**
         * Callback as used by {@link file.FileMgr#uploaded}.
         * @memberof file.FileMgr
         * @typedef UploadedCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {file.File} [response] File
         */

        /**
         * Calls Uploaded.
         * @function uploaded
         * @memberof file.FileMgr
         * @instance
         * @param {file.IFileRequest} request FileRequest message or plain object
         * @param {file.FileMgr.UploadedCallback} callback Node-style callback called with the error, if any, and File
         * @returns {undefined}
         * @variation 1
         */
        FileMgr.prototype.uploaded = function uploaded(request, callback) {
            return this.rpcCall(uploaded, $root.file.FileRequest, $root.file.File, request, callback);
        };

        /**
         * Calls Uploaded.
         * @function uploaded
         * @memberof file.FileMgr
         * @instance
         * @param {file.IFileRequest} request FileRequest message or plain object
         * @returns {Promise<file.File>} Promise
         * @variation 2
         */

        return FileMgr;
    })();

    file.FileHeader = (function() {

        /**
         * Properties of a FileHeader.
         * @memberof file
         * @interface IFileHeader
         * @property {common.IContext|null} [ctx] FileHeader ctx
         * @property {string|null} [account_id] FileHeader account_id
         * @property {string|null} [name] FileHeader name
         * @property {string|null} [type] FileHeader type
         * @property {number|Long|null} [size] FileHeader size
         * @property {string|null} [md5] FileHeader md5
         * @property {string|null} [description] FileHeader description
         */

        /**
         * Constructs a new FileHeader.
         * @memberof file
         * @classdesc Represents a FileHeader.
         * @implements IFileHeader
         * @constructor
         * @param {file.IFileHeader=} [p] Properties to set
         */
        function FileHeader(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * FileHeader ctx.
         * @member {common.IContext|null|undefined} ctx
         * @memberof file.FileHeader
         * @instance
         */
        FileHeader.prototype.ctx = null;

        /**
         * FileHeader account_id.
         * @member {string} account_id
         * @memberof file.FileHeader
         * @instance
         */
        FileHeader.prototype.account_id = "";

        /**
         * FileHeader name.
         * @member {string} name
         * @memberof file.FileHeader
         * @instance
         */
        FileHeader.prototype.name = "";

        /**
         * FileHeader type.
         * @member {string} type
         * @memberof file.FileHeader
         * @instance
         */
        FileHeader.prototype.type = "";

        /**
         * FileHeader size.
         * @member {number|Long} size
         * @memberof file.FileHeader
         * @instance
         */
        FileHeader.prototype.size = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

        /**
         * FileHeader md5.
         * @member {string} md5
         * @memberof file.FileHeader
         * @instance
         */
        FileHeader.prototype.md5 = "";

        /**
         * FileHeader description.
         * @member {string} description
         * @memberof file.FileHeader
         * @instance
         */
        FileHeader.prototype.description = "";

        return FileHeader;
    })();

    file.PresignResult = (function() {

        /**
         * Properties of a PresignResult.
         * @memberof file
         * @interface IPresignResult
         * @property {common.IContext|null} [ctx] PresignResult ctx
         * @property {string|null} [account_id] PresignResult account_id
         * @property {string|null} [url] PresignResult url
         * @property {string|null} [id] PresignResult id
         * @property {string|null} [signed_url] PresignResult signed_url
         */

        /**
         * Constructs a new PresignResult.
         * @memberof file
         * @classdesc Represents a PresignResult.
         * @implements IPresignResult
         * @constructor
         * @param {file.IPresignResult=} [p] Properties to set
         */
        function PresignResult(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * PresignResult ctx.
         * @member {common.IContext|null|undefined} ctx
         * @memberof file.PresignResult
         * @instance
         */
        PresignResult.prototype.ctx = null;

        /**
         * PresignResult account_id.
         * @member {string} account_id
         * @memberof file.PresignResult
         * @instance
         */
        PresignResult.prototype.account_id = "";

        /**
         * PresignResult url.
         * @member {string} url
         * @memberof file.PresignResult
         * @instance
         */
        PresignResult.prototype.url = "";

        /**
         * PresignResult id.
         * @member {string} id
         * @memberof file.PresignResult
         * @instance
         */
        PresignResult.prototype.id = "";

        /**
         * PresignResult signed_url.
         * @member {string} signed_url
         * @memberof file.PresignResult
         * @instance
         */
        PresignResult.prototype.signed_url = "";

        return PresignResult;
    })();

    file.FileRequest = (function() {

        /**
         * Properties of a FileRequest.
         * @memberof file
         * @interface IFileRequest
         * @property {common.IContext|null} [ctx] FileRequest ctx
         * @property {string|null} [account_id] FileRequest account_id
         * @property {string|null} [id] FileRequest id
         */

        /**
         * Constructs a new FileRequest.
         * @memberof file
         * @classdesc Represents a FileRequest.
         * @implements IFileRequest
         * @constructor
         * @param {file.IFileRequest=} [p] Properties to set
         */
        function FileRequest(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * FileRequest ctx.
         * @member {common.IContext|null|undefined} ctx
         * @memberof file.FileRequest
         * @instance
         */
        FileRequest.prototype.ctx = null;

        /**
         * FileRequest account_id.
         * @member {string} account_id
         * @memberof file.FileRequest
         * @instance
         */
        FileRequest.prototype.account_id = "";

        /**
         * FileRequest id.
         * @member {string} id
         * @memberof file.FileRequest
         * @instance
         */
        FileRequest.prototype.id = "";

        return FileRequest;
    })();

    file.File = (function() {

        /**
         * Properties of a File.
         * @memberof file
         * @interface IFile
         * @property {common.IContext|null} [ctx] File ctx
         * @property {string|null} [account_id] File account_id
         * @property {string|null} [name] File name
         * @property {string|null} [type] File type
         * @property {number|Long|null} [size] File size
         * @property {string|null} [md5] File md5
         * @property {string|null} [description] File description
         * @property {number|Long|null} [created] File created
         * @property {string|null} [url] File url
         * @property {string|null} [creator] File creator
         * @property {string|null} [id] File id
         */

        /**
         * Constructs a new File.
         * @memberof file
         * @classdesc Represents a File.
         * @implements IFile
         * @constructor
         * @param {file.IFile=} [p] Properties to set
         */
        function File(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * File ctx.
         * @member {common.IContext|null|undefined} ctx
         * @memberof file.File
         * @instance
         */
        File.prototype.ctx = null;

        /**
         * File account_id.
         * @member {string} account_id
         * @memberof file.File
         * @instance
         */
        File.prototype.account_id = "";

        /**
         * File name.
         * @member {string} name
         * @memberof file.File
         * @instance
         */
        File.prototype.name = "";

        /**
         * File type.
         * @member {string} type
         * @memberof file.File
         * @instance
         */
        File.prototype.type = "";

        /**
         * File size.
         * @member {number|Long} size
         * @memberof file.File
         * @instance
         */
        File.prototype.size = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

        /**
         * File md5.
         * @member {string} md5
         * @memberof file.File
         * @instance
         */
        File.prototype.md5 = "";

        /**
         * File description.
         * @member {string} description
         * @memberof file.File
         * @instance
         */
        File.prototype.description = "";

        /**
         * File created.
         * @member {number|Long} created
         * @memberof file.File
         * @instance
         */
        File.prototype.created = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

        /**
         * File url.
         * @member {string} url
         * @memberof file.File
         * @instance
         */
        File.prototype.url = "";

        /**
         * File creator.
         * @member {string} creator
         * @memberof file.File
         * @instance
         */
        File.prototype.creator = "";

        /**
         * File id.
         * @member {string} id
         * @memberof file.File
         * @instance
         */
        File.prototype.id = "";

        return File;
    })();

    /**
     * Event enum.
     * @name file.Event
     * @enum {string}
     * @property {number} FilePresignRequested=0 FilePresignRequested value
     * @property {number} FileCreated=3 FileCreated value
     * @property {number} FileInfoRequested=4 FileInfoRequested value
     * @property {number} FileUploaded=6 FileUploaded value
     * @property {number} FileRequested=10 FileRequested value
     */
    file.Event = (function() {
        const valuesById = {}, values = Object.create(valuesById);
        values[valuesById[0] = "FilePresignRequested"] = 0;
        values[valuesById[3] = "FileCreated"] = 3;
        values[valuesById[4] = "FileInfoRequested"] = 4;
        values[valuesById[6] = "FileUploaded"] = 6;
        values[valuesById[10] = "FileRequested"] = 10;
        return values;
    })();

    return file;
})();

export const kv = $root.kv = (() => {

    /**
     * Namespace kv.
     * @exports kv
     * @namespace
     */
    const kv = {};

    kv.Key = (function() {

        /**
         * Properties of a Key.
         * @memberof kv
         * @interface IKey
         * @property {string|null} [partition] Key partition
         * @property {string|null} [key] Key key
         */

        /**
         * Constructs a new Key.
         * @memberof kv
         * @classdesc Represents a Key.
         * @implements IKey
         * @constructor
         * @param {kv.IKey=} [p] Properties to set
         */
        function Key(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * Key partition.
         * @member {string} partition
         * @memberof kv.Key
         * @instance
         */
        Key.prototype.partition = "";

        /**
         * Key key.
         * @member {string} key
         * @memberof kv.Key
         * @instance
         */
        Key.prototype.key = "";

        return Key;
    })();

    kv.Value = (function() {

        /**
         * Properties of a Value.
         * @memberof kv
         * @interface IValue
         * @property {string|null} [partition] Value partition
         * @property {string|null} [key] Value key
         * @property {Uint8Array|null} [bytes] Value bytes
         * @property {string|null} [value] Value value
         */

        /**
         * Constructs a new Value.
         * @memberof kv
         * @classdesc Represents a Value.
         * @implements IValue
         * @constructor
         * @param {kv.IValue=} [p] Properties to set
         */
        function Value(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * Value partition.
         * @member {string} partition
         * @memberof kv.Value
         * @instance
         */
        Value.prototype.partition = "";

        /**
         * Value key.
         * @member {string} key
         * @memberof kv.Value
         * @instance
         */
        Value.prototype.key = "";

        /**
         * Value bytes.
         * @member {Uint8Array} bytes
         * @memberof kv.Value
         * @instance
         */
        Value.prototype.bytes = $util.newBuffer([]);

        /**
         * Value value.
         * @member {string} value
         * @memberof kv.Value
         * @instance
         */
        Value.prototype.value = "";

        return Value;
    })();

    kv.Bool = (function() {

        /**
         * Properties of a Bool.
         * @memberof kv
         * @interface IBool
         * @property {boolean|null} [value] Bool value
         */

        /**
         * Constructs a new Bool.
         * @memberof kv
         * @classdesc Represents a Bool.
         * @implements IBool
         * @constructor
         * @param {kv.IBool=} [p] Properties to set
         */
        function Bool(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * Bool value.
         * @member {boolean} value
         * @memberof kv.Bool
         * @instance
         */
        Bool.prototype.value = false;

        return Bool;
    })();

    kv.KV = (function() {

        /**
         * Constructs a new KV service.
         * @memberof kv
         * @classdesc Represents a KV
         * @extends $protobuf.rpc.Service
         * @constructor
         * @param {$protobuf.RPCImpl} rpcImpl RPC implementation
         * @param {boolean} [requestDelimited=false] Whether requests are length-delimited
         * @param {boolean} [responseDelimited=false] Whether responses are length-delimited
         */
        function KV(rpcImpl, requestDelimited, responseDelimited) {
            $protobuf.rpc.Service.call(this, rpcImpl, requestDelimited, responseDelimited);
        }

        (KV.prototype = Object.create($protobuf.rpc.Service.prototype)).constructor = KV;

        /**
         * Callback as used by {@link kv.KV#set}.
         * @memberof kv.KV
         * @typedef SetCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {kv.Value} [response] Value
         */

        /**
         * Calls Set.
         * @function set
         * @memberof kv.KV
         * @instance
         * @param {kv.IValue} request Value message or plain object
         * @param {kv.KV.SetCallback} callback Node-style callback called with the error, if any, and Value
         * @returns {undefined}
         * @variation 1
         */
        KV.prototype.set = function set(request, callback) {
            return this.rpcCall(set, $root.kv.Value, $root.kv.Value, request, callback);
        };

        /**
         * Calls Set.
         * @function set
         * @memberof kv.KV
         * @instance
         * @param {kv.IValue} request Value message or plain object
         * @returns {Promise<kv.Value>} Promise
         * @variation 2
         */

        /**
         * Callback as used by {@link kv.KV#get}.
         * @memberof kv.KV
         * @typedef GetCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {kv.Value} [response] Value
         */

        /**
         * Calls Get.
         * @function get
         * @memberof kv.KV
         * @instance
         * @param {kv.IKey} request Key message or plain object
         * @param {kv.KV.GetCallback} callback Node-style callback called with the error, if any, and Value
         * @returns {undefined}
         * @variation 1
         */
        KV.prototype.get = function get(request, callback) {
            return this.rpcCall(get, $root.kv.Key, $root.kv.Value, request, callback);
        };

        /**
         * Calls Get.
         * @function get
         * @memberof kv.KV
         * @instance
         * @param {kv.IKey} request Key message or plain object
         * @returns {Promise<kv.Value>} Promise
         * @variation 2
         */

        /**
         * Callback as used by {@link kv.KV#has}.
         * @memberof kv.KV
         * @typedef HasCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {kv.Bool} [response] Bool
         */

        /**
         * Calls Has.
         * @function has
         * @memberof kv.KV
         * @instance
         * @param {kv.IKey} request Key message or plain object
         * @param {kv.KV.HasCallback} callback Node-style callback called with the error, if any, and Bool
         * @returns {undefined}
         * @variation 1
         */
        KV.prototype.has = function has(request, callback) {
            return this.rpcCall(has, $root.kv.Key, $root.kv.Bool, request, callback);
        };

        /**
         * Calls Has.
         * @function has
         * @memberof kv.KV
         * @instance
         * @param {kv.IKey} request Key message or plain object
         * @returns {Promise<kv.Bool>} Promise
         * @variation 2
         */

        return KV;
    })();

    return kv;
})();

export const lang = $root.lang = (() => {

    /**
     * Namespace lang.
     * @exports lang
     * @namespace
     */
    const lang = {};

    /**
     * T enum.
     * @name lang.T
     * @enum {string}
     * @property {number} undefined=0 undefined value
     * @property {number} user_has_already_in_conversation=1 user_has_already_in_conversation value
     * @property {number} conversation_closed=2 conversation_closed value
     * @property {number} invalid_invite=3 invalid_invite value
     * @property {number} invalid_agent=4 invalid_agent value
     * @property {number} member_is_not_in_conversation=5 member_is_not_in_conversation value
     * @property {number} conversation_not_found=6 conversation_not_found value
     * @property {number} internal_error=30 internal_error value
     * @property {number} invalid_input=22 invalid_input value
     * @property {number} invalid_form=20 invalid_form value
     * @property {number} access_token_expired=21 access_token_expired value
     * @property {number} credential_not_set=7 credential_not_set value
     * @property {number} wrong_account_in_credential=8 wrong_account_in_credential value
     * @property {number} access_deny=9 access_deny value
     * @property {number} wrong_user_in_credential=10 wrong_user_in_credential value
     * @property {number} unable_to_send_message=11 unable_to_send_message value
     * @property {number} topic_is_empty=12 topic_is_empty value
     * @property {number} invalid_credential=13 invalid_credential value
     * @property {number} invalid_left=14 invalid_left value
     * @property {number} invalid_json=15 invalid_json value
     * @property {number} invalid_protobuf=16 invalid_protobuf value
     * @property {number} invalid_password=17 invalid_password value
     * @property {number} wrong_password=18 wrong_password value
     * @property {number} invalid_agent_state=19 invalid_agent_state value
     * @property {number} unable_to_lock=40 unable_to_lock value
     * @property {number} empty=41 empty value
     * @property {number} wrong_type=42 wrong_type value
     * @property {number} invalid_kafka_topic=43 invalid_kafka_topic value
     * @property {number} database_error=44 database_error value
     * @property {number} timeout=45 timeout value
     * @property {number} websocket_error=46 websocket_error value
     * @property {number} kafka_error=47 kafka_error value
     * @property {number} invalid_token=48 invalid_token value
     * @property {number} account_not_found=49 account_not_found value
     * @property {number} agent_not_found=50 agent_not_found value
     * @property {number} invalid_email=60 invalid_email value
     * @property {number} plan_not_found=61 plan_not_found value
     * @property {number} agent_group_not_found=62 agent_group_not_found value
     * @property {number} empty_client_type=63 empty_client_type value
     * @property {number} empty_url=64 empty_url value
     * @property {number} empty_name=65 empty_name value
     * @property {number} client_not_found=66 client_not_found value
     * @property {number} empty_account=70 empty_account value
     * @property {number} invalid_conversation_state=71 invalid_conversation_state value
     * @property {number} invalid_message_id=80 invalid_message_id value
     * @property {number} invalid_mask=81 invalid_mask value
     * @property {number} randomize_error=82 randomize_error value
     * @property {number} duplicated_message_received_error=83 duplicated_message_received_error value
     * @property {number} network_error=84 network_error value
     * @property {number} rsa_key_not_found=85 rsa_key_not_found value
     * @property {number} jwt_sign_error=86 jwt_sign_error value
     * @property {number} env_config_error=87 env_config_error value
     * @property {number} scrypt_error=90 scrypt_error value
     * @property {number} challenge_not_found=91 challenge_not_found value
     * @property {number} wrong_answer=92 wrong_answer value
     * @property {number} server_listen_error=93 server_listen_error value
     * @property {number} scrypt_file_not_found=94 scrypt_file_not_found value
     * @property {number} invalid_topic=95 invalid_topic value
     * @property {number} file_not_found=99 file_not_found value
     * @property {number} user_not_found=100 user_not_found value
     * @property {number} empty_md5=101 empty_md5 value
     * @property {number} base_convert_error=102 base_convert_error value
     * @property {number} s3_error=103 s3_error value
     * @property {number} aws_credential_error=104 aws_credential_error value
     * @property {number} aws_send_error=105 aws_send_error value
     * @property {number} elasticsearch_error=200 elasticsearch_error value
     * @property {number} invalid_template=203 invalid_template value
     * @property {number} sendgrid_send_error=204 sendgrid_send_error value
     * @property {number} whitelist_domain_not_found=205 whitelist_domain_not_found value
     * @property {number} blacklist_ip_not_found=206 blacklist_ip_not_found value
     * @property {number} blocked_user_not_found=207 blocked_user_not_found value
     * @property {number} invalid_content_type=210 invalid_content_type value
     * @property {number} parse_file_error=211 parse_file_error value
     * @property {number} invalid_integration_id=220 invalid_integration_id value
     * @property {number} invalid_integration=221 invalid_integration value
     * @property {number} webhook_not_found=222 webhook_not_found value
     * @property {number} tempfile_error=223 tempfile_error value
     * @property {number} write_file_error=224 write_file_error value
     * @property {number} close_file_error=225 close_file_error value
     * @property {number} execute_shell_error=226 execute_shell_error value
     * @property {number} invalid_css=227 invalid_css value
     * @property {number} invalid_hmac=228 invalid_hmac value
     * @property {number} consul_error=230 consul_error value
     * @property {number} maxminddb_err=231 maxminddb_err value
     * @property {number} invalid_condition_key=232 invalid_condition_key value
     * @property {number} invalid_po_file=233 invalid_po_file value
     * @property {number} integration_not_found=234 integration_not_found value
     * @property {number} webhook_is_disabled=235 webhook_is_disabled value
     * @property {number} agent_not_activated=240 agent_not_activated value
     * @property {number} empty_message=242 empty_message value
     * @property {number} message_too_large=243 message_too_large value
     * @property {number} unknown_message_format=244 unknown_message_format value
     * @property {number} too_many_attachments=245 too_many_attachments value
     * @property {number} too_many_fields=246 too_many_fields value
     * @property {number} attachment_too_large=247 attachment_too_large value
     * @property {number} invalid_end=248 invalid_end value
     * @property {number} invalid_ack=249 invalid_ack value
     * @property {number} message_not_found=250 message_not_found value
     * @property {number} invalid_apikey=255 invalid_apikey value
     * @property {number} invalid_access_token=256 invalid_access_token value
     * @property {number} field_too_long=257 field_too_long value
     * @property {number} invalid_join=258 invalid_join value
     * @property {number} automation_not_found=259 automation_not_found value
     * @property {number} invalid_automation_condition=260 invalid_automation_condition value
     * @property {number} automation_cache_miss=261 automation_cache_miss value
     * @property {number} segment_loop_duplicated=262 segment_loop_duplicated value
     * @property {number} segment_loop_stale=263 segment_loop_stale value
     * @property {number} segment_not_found=264 segment_not_found value
     * @property {number} domain_is_blocked=265 domain_is_blocked value
     * @property {number} user_is_blocked=266 user_is_blocked value
     * @property {number} ip_is_blocked=267 ip_is_blocked value
     * @property {number} stripe_error=270 stripe_error value
     * @property {number} missing_stripe_token=271 missing_stripe_token value
     * @property {number} missing_stripe_customer_id=272 missing_stripe_customer_id value
     * @property {number} payment_method_not_found=273 payment_method_not_found value
     * @property {number} not_enough_money=274 not_enough_money value
     * @property {number} invalid_account=275 invalid_account value
     * @property {number} invalid_payment_method=276 invalid_payment_method value
     * @property {number} invalid_subscription=278 invalid_subscription value
     * @property {number} invalid_plan=279 invalid_plan value
     * @property {number} invoice_not_found=280 invoice_not_found value
     * @property {number} billing_not_found=281 billing_not_found value
     * @property {number} exchange_rate_not_found=282 exchange_rate_not_found value
     * @property {number} invalid_join_free=283 invalid_join_free value
     * @property {number} invalid_active_free=284 invalid_active_free value
     * @property {number} kv_not_found=285 kv_not_found value
     * @property {number} attribute_not_found=286 attribute_not_found value
     * @property {number} attribute_key_not_found=287 attribute_key_not_found value
     * @property {number} attribute_max_allowed=288 attribute_max_allowed value
     * @property {number} invalid_attribute=289 invalid_attribute value
     */
    lang.T = (function() {
        const valuesById = {}, values = Object.create(valuesById);
        values[valuesById[0] = "undefined"] = 0;
        values[valuesById[1] = "user_has_already_in_conversation"] = 1;
        values[valuesById[2] = "conversation_closed"] = 2;
        values[valuesById[3] = "invalid_invite"] = 3;
        values[valuesById[4] = "invalid_agent"] = 4;
        values[valuesById[5] = "member_is_not_in_conversation"] = 5;
        values[valuesById[6] = "conversation_not_found"] = 6;
        values[valuesById[30] = "internal_error"] = 30;
        values[valuesById[22] = "invalid_input"] = 22;
        values[valuesById[20] = "invalid_form"] = 20;
        values[valuesById[21] = "access_token_expired"] = 21;
        values[valuesById[7] = "credential_not_set"] = 7;
        values[valuesById[8] = "wrong_account_in_credential"] = 8;
        values[valuesById[9] = "access_deny"] = 9;
        values[valuesById[10] = "wrong_user_in_credential"] = 10;
        values[valuesById[11] = "unable_to_send_message"] = 11;
        values[valuesById[12] = "topic_is_empty"] = 12;
        values[valuesById[13] = "invalid_credential"] = 13;
        values[valuesById[14] = "invalid_left"] = 14;
        values[valuesById[15] = "invalid_json"] = 15;
        values[valuesById[16] = "invalid_protobuf"] = 16;
        values[valuesById[17] = "invalid_password"] = 17;
        values[valuesById[18] = "wrong_password"] = 18;
        values[valuesById[19] = "invalid_agent_state"] = 19;
        values[valuesById[40] = "unable_to_lock"] = 40;
        values[valuesById[41] = "empty"] = 41;
        values[valuesById[42] = "wrong_type"] = 42;
        values[valuesById[43] = "invalid_kafka_topic"] = 43;
        values[valuesById[44] = "database_error"] = 44;
        values[valuesById[45] = "timeout"] = 45;
        values[valuesById[46] = "websocket_error"] = 46;
        values[valuesById[47] = "kafka_error"] = 47;
        values[valuesById[48] = "invalid_token"] = 48;
        values[valuesById[49] = "account_not_found"] = 49;
        values[valuesById[50] = "agent_not_found"] = 50;
        values[valuesById[60] = "invalid_email"] = 60;
        values[valuesById[61] = "plan_not_found"] = 61;
        values[valuesById[62] = "agent_group_not_found"] = 62;
        values[valuesById[63] = "empty_client_type"] = 63;
        values[valuesById[64] = "empty_url"] = 64;
        values[valuesById[65] = "empty_name"] = 65;
        values[valuesById[66] = "client_not_found"] = 66;
        values[valuesById[70] = "empty_account"] = 70;
        values[valuesById[71] = "invalid_conversation_state"] = 71;
        values[valuesById[80] = "invalid_message_id"] = 80;
        values[valuesById[81] = "invalid_mask"] = 81;
        values[valuesById[82] = "randomize_error"] = 82;
        values[valuesById[83] = "duplicated_message_received_error"] = 83;
        values[valuesById[84] = "network_error"] = 84;
        values[valuesById[85] = "rsa_key_not_found"] = 85;
        values[valuesById[86] = "jwt_sign_error"] = 86;
        values[valuesById[87] = "env_config_error"] = 87;
        values[valuesById[90] = "scrypt_error"] = 90;
        values[valuesById[91] = "challenge_not_found"] = 91;
        values[valuesById[92] = "wrong_answer"] = 92;
        values[valuesById[93] = "server_listen_error"] = 93;
        values[valuesById[94] = "scrypt_file_not_found"] = 94;
        values[valuesById[95] = "invalid_topic"] = 95;
        values[valuesById[99] = "file_not_found"] = 99;
        values[valuesById[100] = "user_not_found"] = 100;
        values[valuesById[101] = "empty_md5"] = 101;
        values[valuesById[102] = "base_convert_error"] = 102;
        values[valuesById[103] = "s3_error"] = 103;
        values[valuesById[104] = "aws_credential_error"] = 104;
        values[valuesById[105] = "aws_send_error"] = 105;
        values[valuesById[200] = "elasticsearch_error"] = 200;
        values[valuesById[203] = "invalid_template"] = 203;
        values[valuesById[204] = "sendgrid_send_error"] = 204;
        values[valuesById[205] = "whitelist_domain_not_found"] = 205;
        values[valuesById[206] = "blacklist_ip_not_found"] = 206;
        values[valuesById[207] = "blocked_user_not_found"] = 207;
        values[valuesById[210] = "invalid_content_type"] = 210;
        values[valuesById[211] = "parse_file_error"] = 211;
        values[valuesById[220] = "invalid_integration_id"] = 220;
        values[valuesById[221] = "invalid_integration"] = 221;
        values[valuesById[222] = "webhook_not_found"] = 222;
        values[valuesById[223] = "tempfile_error"] = 223;
        values[valuesById[224] = "write_file_error"] = 224;
        values[valuesById[225] = "close_file_error"] = 225;
        values[valuesById[226] = "execute_shell_error"] = 226;
        values[valuesById[227] = "invalid_css"] = 227;
        values[valuesById[228] = "invalid_hmac"] = 228;
        values[valuesById[230] = "consul_error"] = 230;
        values[valuesById[231] = "maxminddb_err"] = 231;
        values[valuesById[232] = "invalid_condition_key"] = 232;
        values[valuesById[233] = "invalid_po_file"] = 233;
        values[valuesById[234] = "integration_not_found"] = 234;
        values[valuesById[235] = "webhook_is_disabled"] = 235;
        values[valuesById[240] = "agent_not_activated"] = 240;
        values[valuesById[242] = "empty_message"] = 242;
        values[valuesById[243] = "message_too_large"] = 243;
        values[valuesById[244] = "unknown_message_format"] = 244;
        values[valuesById[245] = "too_many_attachments"] = 245;
        values[valuesById[246] = "too_many_fields"] = 246;
        values[valuesById[247] = "attachment_too_large"] = 247;
        values[valuesById[248] = "invalid_end"] = 248;
        values[valuesById[249] = "invalid_ack"] = 249;
        values[valuesById[250] = "message_not_found"] = 250;
        values[valuesById[255] = "invalid_apikey"] = 255;
        values[valuesById[256] = "invalid_access_token"] = 256;
        values[valuesById[257] = "field_too_long"] = 257;
        values[valuesById[258] = "invalid_join"] = 258;
        values[valuesById[259] = "automation_not_found"] = 259;
        values[valuesById[260] = "invalid_automation_condition"] = 260;
        values[valuesById[261] = "automation_cache_miss"] = 261;
        values[valuesById[262] = "segment_loop_duplicated"] = 262;
        values[valuesById[263] = "segment_loop_stale"] = 263;
        values[valuesById[264] = "segment_not_found"] = 264;
        values[valuesById[265] = "domain_is_blocked"] = 265;
        values[valuesById[266] = "user_is_blocked"] = 266;
        values[valuesById[267] = "ip_is_blocked"] = 267;
        values[valuesById[270] = "stripe_error"] = 270;
        values[valuesById[271] = "missing_stripe_token"] = 271;
        values[valuesById[272] = "missing_stripe_customer_id"] = 272;
        values[valuesById[273] = "payment_method_not_found"] = 273;
        values[valuesById[274] = "not_enough_money"] = 274;
        values[valuesById[275] = "invalid_account"] = 275;
        values[valuesById[276] = "invalid_payment_method"] = 276;
        values[valuesById[278] = "invalid_subscription"] = 278;
        values[valuesById[279] = "invalid_plan"] = 279;
        values[valuesById[280] = "invoice_not_found"] = 280;
        values[valuesById[281] = "billing_not_found"] = 281;
        values[valuesById[282] = "exchange_rate_not_found"] = 282;
        values[valuesById[283] = "invalid_join_free"] = 283;
        values[valuesById[284] = "invalid_active_free"] = 284;
        values[valuesById[285] = "kv_not_found"] = 285;
        values[valuesById[286] = "attribute_not_found"] = 286;
        values[valuesById[287] = "attribute_key_not_found"] = 287;
        values[valuesById[288] = "attribute_max_allowed"] = 288;
        values[valuesById[289] = "invalid_attribute"] = 289;
        return values;
    })();

    /**
     * L enum.
     * @name lang.L
     * @enum {string}
     * @property {number} en=0 en value
     * @property {number} vn=1 vn value
     * @property {number} cn=3 cn value
     * @property {number} fr=5 fr value
     */
    lang.L = (function() {
        const valuesById = {}, values = Object.create(valuesById);
        values[valuesById[0] = "en"] = 0;
        values[valuesById[1] = "vn"] = 1;
        values[valuesById[3] = "cn"] = 3;
        values[valuesById[5] = "fr"] = 5;
        return values;
    })();

    return lang;
})();

export const logan = $root.logan = (() => {

    /**
     * Namespace logan.
     * @exports logan
     * @namespace
     */
    const logan = {};

    logan.Debug = (function() {

        /**
         * Properties of a Debug.
         * @memberof logan
         * @interface IDebug
         * @property {logan.IMemStats|null} [mem_stats] Debug mem_stats
         * @property {number|null} [num_cpu] Debug num_cpu
         * @property {number|null} [num_goroutine] Debug num_goroutine
         * @property {Uint8Array|null} [stack_trace] Debug stack_trace
         * @property {string|null} [hostname] Debug hostname
         * @property {logan.IKafkaInfo|null} [kafka] Debug kafka
         */

        /**
         * Constructs a new Debug.
         * @memberof logan
         * @classdesc Represents a Debug.
         * @implements IDebug
         * @constructor
         * @param {logan.IDebug=} [p] Properties to set
         */
        function Debug(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * Debug mem_stats.
         * @member {logan.IMemStats|null|undefined} mem_stats
         * @memberof logan.Debug
         * @instance
         */
        Debug.prototype.mem_stats = null;

        /**
         * Debug num_cpu.
         * @member {number} num_cpu
         * @memberof logan.Debug
         * @instance
         */
        Debug.prototype.num_cpu = 0;

        /**
         * Debug num_goroutine.
         * @member {number} num_goroutine
         * @memberof logan.Debug
         * @instance
         */
        Debug.prototype.num_goroutine = 0;

        /**
         * Debug stack_trace.
         * @member {Uint8Array} stack_trace
         * @memberof logan.Debug
         * @instance
         */
        Debug.prototype.stack_trace = $util.newBuffer([]);

        /**
         * Debug hostname.
         * @member {string} hostname
         * @memberof logan.Debug
         * @instance
         */
        Debug.prototype.hostname = "";

        /**
         * Debug kafka.
         * @member {logan.IKafkaInfo|null|undefined} kafka
         * @memberof logan.Debug
         * @instance
         */
        Debug.prototype.kafka = null;

        return Debug;
    })();

    logan.KafkaInfo = (function() {

        /**
         * Properties of a KafkaInfo.
         * @memberof logan
         * @interface IKafkaInfo
         * @property {string|null} [topic] KafkaInfo topic
         * @property {number|null} [partition] KafkaInfo partition
         * @property {number|Long|null} [offset] KafkaInfo offset
         */

        /**
         * Constructs a new KafkaInfo.
         * @memberof logan
         * @classdesc Represents a KafkaInfo.
         * @implements IKafkaInfo
         * @constructor
         * @param {logan.IKafkaInfo=} [p] Properties to set
         */
        function KafkaInfo(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * KafkaInfo topic.
         * @member {string} topic
         * @memberof logan.KafkaInfo
         * @instance
         */
        KafkaInfo.prototype.topic = "";

        /**
         * KafkaInfo partition.
         * @member {number} partition
         * @memberof logan.KafkaInfo
         * @instance
         */
        KafkaInfo.prototype.partition = 0;

        /**
         * KafkaInfo offset.
         * @member {number|Long} offset
         * @memberof logan.KafkaInfo
         * @instance
         */
        KafkaInfo.prototype.offset = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

        return KafkaInfo;
    })();

    logan.Log = (function() {

        /**
         * Properties of a Log.
         * @memberof logan
         * @interface ILog
         * @property {common.IContext|null} [ctx] Log ctx
         * @property {string|null} [trace_id] Log trace_id
         * @property {number|Long|null} [created] Log created
         * @property {string|null} [level] Log level
         * @property {Array.<string>|null} [tags] Log tags
         * @property {logan.IDebug|null} [debug] Log debug
         * @property {Uint8Array|null} [message] Log message
         */

        /**
         * Constructs a new Log.
         * @memberof logan
         * @classdesc Represents a Log.
         * @implements ILog
         * @constructor
         * @param {logan.ILog=} [p] Properties to set
         */
        function Log(p) {
            this.tags = [];
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * Log ctx.
         * @member {common.IContext|null|undefined} ctx
         * @memberof logan.Log
         * @instance
         */
        Log.prototype.ctx = null;

        /**
         * Log trace_id.
         * @member {string} trace_id
         * @memberof logan.Log
         * @instance
         */
        Log.prototype.trace_id = "";

        /**
         * Log created.
         * @member {number|Long} created
         * @memberof logan.Log
         * @instance
         */
        Log.prototype.created = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

        /**
         * Log level.
         * @member {string} level
         * @memberof logan.Log
         * @instance
         */
        Log.prototype.level = "";

        /**
         * Log tags.
         * @member {Array.<string>} tags
         * @memberof logan.Log
         * @instance
         */
        Log.prototype.tags = $util.emptyArray;

        /**
         * Log debug.
         * @member {logan.IDebug|null|undefined} debug
         * @memberof logan.Log
         * @instance
         */
        Log.prototype.debug = null;

        /**
         * Log message.
         * @member {Uint8Array} message
         * @memberof logan.Log
         * @instance
         */
        Log.prototype.message = $util.newBuffer([]);

        return Log;
    })();

    /**
     * Level enum.
     * @name logan.Level
     * @enum {string}
     * @property {number} debug=0 debug value
     * @property {number} info=1 info value
     * @property {number} notice=2 notice value
     * @property {number} warning=3 warning value
     * @property {number} error=4 error value
     * @property {number} critical=5 critical value
     * @property {number} alert=6 alert value
     * @property {number} emergency=7 emergency value
     * @property {number} panic=8 panic value
     * @property {number} fatal=9 fatal value
     */
    logan.Level = (function() {
        const valuesById = {}, values = Object.create(valuesById);
        values[valuesById[0] = "debug"] = 0;
        values[valuesById[1] = "info"] = 1;
        values[valuesById[2] = "notice"] = 2;
        values[valuesById[3] = "warning"] = 3;
        values[valuesById[4] = "error"] = 4;
        values[valuesById[5] = "critical"] = 5;
        values[valuesById[6] = "alert"] = 6;
        values[valuesById[7] = "emergency"] = 7;
        values[valuesById[8] = "panic"] = 8;
        values[valuesById[9] = "fatal"] = 9;
        return values;
    })();

    /**
     * Event enum.
     * @name logan.Event
     * @enum {string}
     * @property {number} LogLogRequested=0 LogLogRequested value
     * @property {number} LogRequested=1000 LogRequested value
     * @property {number} LogSynced=1001 LogSynced value
     */
    logan.Event = (function() {
        const valuesById = {}, values = Object.create(valuesById);
        values[valuesById[0] = "LogLogRequested"] = 0;
        values[valuesById[1000] = "LogRequested"] = 1000;
        values[valuesById[1001] = "LogSynced"] = 1001;
        return values;
    })();

    logan.MemStats = (function() {

        /**
         * Properties of a MemStats.
         * @memberof logan
         * @interface IMemStats
         * @property {number|Long|null} [alloc] MemStats alloc
         * @property {number|Long|null} [total_alloc] MemStats total_alloc
         * @property {number|Long|null} [sys] MemStats sys
         * @property {number|Long|null} [lookups] MemStats lookups
         * @property {number|Long|null} [mallocs] MemStats mallocs
         * @property {number|Long|null} [frees] MemStats frees
         * @property {number|Long|null} [heap_alloc] MemStats heap_alloc
         * @property {number|Long|null} [heap_sys] MemStats heap_sys
         * @property {number|Long|null} [heap_idle] MemStats heap_idle
         * @property {number|Long|null} [heap_inuse] MemStats heap_inuse
         * @property {number|Long|null} [heap_released] MemStats heap_released
         * @property {number|Long|null} [heap_objects] MemStats heap_objects
         * @property {number|Long|null} [stack_inuse] MemStats stack_inuse
         * @property {number|Long|null} [stack_sys] MemStats stack_sys
         * @property {number|Long|null} [m_span_inuse] MemStats m_span_inuse
         * @property {number|Long|null} [m_span_sys] MemStats m_span_sys
         * @property {number|Long|null} [m_cache_inuse] MemStats m_cache_inuse
         * @property {number|Long|null} [m_cache_sys] MemStats m_cache_sys
         * @property {number|Long|null} [buck_hash_sys] MemStats buck_hash_sys
         * @property {number|Long|null} [gc_sys] MemStats gc_sys
         * @property {number|Long|null} [other_sys] MemStats other_sys
         * @property {number|Long|null} [next_gc] MemStats next_gc
         * @property {number|Long|null} [last_gc] MemStats last_gc
         * @property {number|Long|null} [pause_total_ns] MemStats pause_total_ns
         * @property {number|null} [num_gc] MemStats num_gc
         * @property {number|null} [num_forced_gc] MemStats num_forced_gc
         * @property {number|null} [gc_cpu_fraction] MemStats gc_cpu_fraction
         */

        /**
         * Constructs a new MemStats.
         * @memberof logan
         * @classdesc Represents a MemStats.
         * @implements IMemStats
         * @constructor
         * @param {logan.IMemStats=} [p] Properties to set
         */
        function MemStats(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * MemStats alloc.
         * @member {number|Long} alloc
         * @memberof logan.MemStats
         * @instance
         */
        MemStats.prototype.alloc = $util.Long ? $util.Long.fromBits(0,0,true) : 0;

        /**
         * MemStats total_alloc.
         * @member {number|Long} total_alloc
         * @memberof logan.MemStats
         * @instance
         */
        MemStats.prototype.total_alloc = $util.Long ? $util.Long.fromBits(0,0,true) : 0;

        /**
         * MemStats sys.
         * @member {number|Long} sys
         * @memberof logan.MemStats
         * @instance
         */
        MemStats.prototype.sys = $util.Long ? $util.Long.fromBits(0,0,true) : 0;

        /**
         * MemStats lookups.
         * @member {number|Long} lookups
         * @memberof logan.MemStats
         * @instance
         */
        MemStats.prototype.lookups = $util.Long ? $util.Long.fromBits(0,0,true) : 0;

        /**
         * MemStats mallocs.
         * @member {number|Long} mallocs
         * @memberof logan.MemStats
         * @instance
         */
        MemStats.prototype.mallocs = $util.Long ? $util.Long.fromBits(0,0,true) : 0;

        /**
         * MemStats frees.
         * @member {number|Long} frees
         * @memberof logan.MemStats
         * @instance
         */
        MemStats.prototype.frees = $util.Long ? $util.Long.fromBits(0,0,true) : 0;

        /**
         * MemStats heap_alloc.
         * @member {number|Long} heap_alloc
         * @memberof logan.MemStats
         * @instance
         */
        MemStats.prototype.heap_alloc = $util.Long ? $util.Long.fromBits(0,0,true) : 0;

        /**
         * MemStats heap_sys.
         * @member {number|Long} heap_sys
         * @memberof logan.MemStats
         * @instance
         */
        MemStats.prototype.heap_sys = $util.Long ? $util.Long.fromBits(0,0,true) : 0;

        /**
         * MemStats heap_idle.
         * @member {number|Long} heap_idle
         * @memberof logan.MemStats
         * @instance
         */
        MemStats.prototype.heap_idle = $util.Long ? $util.Long.fromBits(0,0,true) : 0;

        /**
         * MemStats heap_inuse.
         * @member {number|Long} heap_inuse
         * @memberof logan.MemStats
         * @instance
         */
        MemStats.prototype.heap_inuse = $util.Long ? $util.Long.fromBits(0,0,true) : 0;

        /**
         * MemStats heap_released.
         * @member {number|Long} heap_released
         * @memberof logan.MemStats
         * @instance
         */
        MemStats.prototype.heap_released = $util.Long ? $util.Long.fromBits(0,0,true) : 0;

        /**
         * MemStats heap_objects.
         * @member {number|Long} heap_objects
         * @memberof logan.MemStats
         * @instance
         */
        MemStats.prototype.heap_objects = $util.Long ? $util.Long.fromBits(0,0,true) : 0;

        /**
         * MemStats stack_inuse.
         * @member {number|Long} stack_inuse
         * @memberof logan.MemStats
         * @instance
         */
        MemStats.prototype.stack_inuse = $util.Long ? $util.Long.fromBits(0,0,true) : 0;

        /**
         * MemStats stack_sys.
         * @member {number|Long} stack_sys
         * @memberof logan.MemStats
         * @instance
         */
        MemStats.prototype.stack_sys = $util.Long ? $util.Long.fromBits(0,0,true) : 0;

        /**
         * MemStats m_span_inuse.
         * @member {number|Long} m_span_inuse
         * @memberof logan.MemStats
         * @instance
         */
        MemStats.prototype.m_span_inuse = $util.Long ? $util.Long.fromBits(0,0,true) : 0;

        /**
         * MemStats m_span_sys.
         * @member {number|Long} m_span_sys
         * @memberof logan.MemStats
         * @instance
         */
        MemStats.prototype.m_span_sys = $util.Long ? $util.Long.fromBits(0,0,true) : 0;

        /**
         * MemStats m_cache_inuse.
         * @member {number|Long} m_cache_inuse
         * @memberof logan.MemStats
         * @instance
         */
        MemStats.prototype.m_cache_inuse = $util.Long ? $util.Long.fromBits(0,0,true) : 0;

        /**
         * MemStats m_cache_sys.
         * @member {number|Long} m_cache_sys
         * @memberof logan.MemStats
         * @instance
         */
        MemStats.prototype.m_cache_sys = $util.Long ? $util.Long.fromBits(0,0,true) : 0;

        /**
         * MemStats buck_hash_sys.
         * @member {number|Long} buck_hash_sys
         * @memberof logan.MemStats
         * @instance
         */
        MemStats.prototype.buck_hash_sys = $util.Long ? $util.Long.fromBits(0,0,true) : 0;

        /**
         * MemStats gc_sys.
         * @member {number|Long} gc_sys
         * @memberof logan.MemStats
         * @instance
         */
        MemStats.prototype.gc_sys = $util.Long ? $util.Long.fromBits(0,0,true) : 0;

        /**
         * MemStats other_sys.
         * @member {number|Long} other_sys
         * @memberof logan.MemStats
         * @instance
         */
        MemStats.prototype.other_sys = $util.Long ? $util.Long.fromBits(0,0,true) : 0;

        /**
         * MemStats next_gc.
         * @member {number|Long} next_gc
         * @memberof logan.MemStats
         * @instance
         */
        MemStats.prototype.next_gc = $util.Long ? $util.Long.fromBits(0,0,true) : 0;

        /**
         * MemStats last_gc.
         * @member {number|Long} last_gc
         * @memberof logan.MemStats
         * @instance
         */
        MemStats.prototype.last_gc = $util.Long ? $util.Long.fromBits(0,0,true) : 0;

        /**
         * MemStats pause_total_ns.
         * @member {number|Long} pause_total_ns
         * @memberof logan.MemStats
         * @instance
         */
        MemStats.prototype.pause_total_ns = $util.Long ? $util.Long.fromBits(0,0,true) : 0;

        /**
         * MemStats num_gc.
         * @member {number} num_gc
         * @memberof logan.MemStats
         * @instance
         */
        MemStats.prototype.num_gc = 0;

        /**
         * MemStats num_forced_gc.
         * @member {number} num_forced_gc
         * @memberof logan.MemStats
         * @instance
         */
        MemStats.prototype.num_forced_gc = 0;

        /**
         * MemStats gc_cpu_fraction.
         * @member {number} gc_cpu_fraction
         * @memberof logan.MemStats
         * @instance
         */
        MemStats.prototype.gc_cpu_fraction = 0;

        return MemStats;
    })();

    return logan;
})();

export const mailkon = $root.mailkon = (() => {

    /**
     * Namespace mailkon.
     * @exports mailkon
     * @namespace
     */
    const mailkon = {};

    mailkon.Address = (function() {

        /**
         * Properties of an Address.
         * @memberof mailkon
         * @interface IAddress
         * @property {string|null} [account_id] Address account_id
         * @property {string|null} [address] Address address
         * @property {number|Long|null} [updated] Address updated
         */

        /**
         * Constructs a new Address.
         * @memberof mailkon
         * @classdesc Represents an Address.
         * @implements IAddress
         * @constructor
         * @param {mailkon.IAddress=} [p] Properties to set
         */
        function Address(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * Address account_id.
         * @member {string} account_id
         * @memberof mailkon.Address
         * @instance
         */
        Address.prototype.account_id = "";

        /**
         * Address address.
         * @member {string} address
         * @memberof mailkon.Address
         * @instance
         */
        Address.prototype.address = "";

        /**
         * Address updated.
         * @member {number|Long} updated
         * @memberof mailkon.Address
         * @instance
         */
        Address.prototype.updated = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

        return Address;
    })();

    mailkon.Domain = (function() {

        /**
         * Properties of a Domain.
         * @memberof mailkon
         * @interface IDomain
         * @property {string|null} [account_id] Domain account_id
         * @property {string|null} [domain] Domain domain
         * @property {number|Long|null} [updated] Domain updated
         * @property {string|null} [dnstype] Domain dnstype
         * @property {string|null} [data] Domain data
         */

        /**
         * Constructs a new Domain.
         * @memberof mailkon
         * @classdesc Represents a Domain.
         * @implements IDomain
         * @constructor
         * @param {mailkon.IDomain=} [p] Properties to set
         */
        function Domain(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * Domain account_id.
         * @member {string} account_id
         * @memberof mailkon.Domain
         * @instance
         */
        Domain.prototype.account_id = "";

        /**
         * Domain domain.
         * @member {string} domain
         * @memberof mailkon.Domain
         * @instance
         */
        Domain.prototype.domain = "";

        /**
         * Domain updated.
         * @member {number|Long} updated
         * @memberof mailkon.Domain
         * @instance
         */
        Domain.prototype.updated = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

        /**
         * Domain dnstype.
         * @member {string} dnstype
         * @memberof mailkon.Domain
         * @instance
         */
        Domain.prototype.dnstype = "";

        /**
         * Domain data.
         * @member {string} data
         * @memberof mailkon.Domain
         * @instance
         */
        Domain.prototype.data = "";

        return Domain;
    })();

    mailkon.Message = (function() {

        /**
         * Properties of a Message.
         * @memberof mailkon
         * @interface IMessage
         * @property {string|null} [message_id] Message message_id
         * @property {string|null} [account_id] Message account_id
         * @property {string|null} [conversation_id] Message conversation_id
         * @property {string|null} [from_addr] Message from_addr
         * @property {string|null} [to_addr] Message to_addr
         * @property {string|null} [subject] Message subject
         * @property {string|null} [body] Message body
         * @property {string|null} [header] Message header
         * @property {Array.<string>|null} [attachments] Message attachments
         * @property {number|Long|null} [created] Message created
         */

        /**
         * Constructs a new Message.
         * @memberof mailkon
         * @classdesc Represents a Message.
         * @implements IMessage
         * @constructor
         * @param {mailkon.IMessage=} [p] Properties to set
         */
        function Message(p) {
            this.attachments = [];
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * Message message_id.
         * @member {string} message_id
         * @memberof mailkon.Message
         * @instance
         */
        Message.prototype.message_id = "";

        /**
         * Message account_id.
         * @member {string} account_id
         * @memberof mailkon.Message
         * @instance
         */
        Message.prototype.account_id = "";

        /**
         * Message conversation_id.
         * @member {string} conversation_id
         * @memberof mailkon.Message
         * @instance
         */
        Message.prototype.conversation_id = "";

        /**
         * Message from_addr.
         * @member {string} from_addr
         * @memberof mailkon.Message
         * @instance
         */
        Message.prototype.from_addr = "";

        /**
         * Message to_addr.
         * @member {string} to_addr
         * @memberof mailkon.Message
         * @instance
         */
        Message.prototype.to_addr = "";

        /**
         * Message subject.
         * @member {string} subject
         * @memberof mailkon.Message
         * @instance
         */
        Message.prototype.subject = "";

        /**
         * Message body.
         * @member {string} body
         * @memberof mailkon.Message
         * @instance
         */
        Message.prototype.body = "";

        /**
         * Message header.
         * @member {string} header
         * @memberof mailkon.Message
         * @instance
         */
        Message.prototype.header = "";

        /**
         * Message attachments.
         * @member {Array.<string>} attachments
         * @memberof mailkon.Message
         * @instance
         */
        Message.prototype.attachments = $util.emptyArray;

        /**
         * Message created.
         * @member {number|Long} created
         * @memberof mailkon.Message
         * @instance
         */
        Message.prototype.created = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

        return Message;
    })();

    mailkon.User = (function() {

        /**
         * Properties of a User.
         * @memberof mailkon
         * @interface IUser
         * @property {string|null} [account_id] User account_id
         * @property {string|null} [user_id] User user_id
         * @property {string|null} [email_address] User email_address
         */

        /**
         * Constructs a new User.
         * @memberof mailkon
         * @classdesc Represents a User.
         * @implements IUser
         * @constructor
         * @param {mailkon.IUser=} [p] Properties to set
         */
        function User(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * User account_id.
         * @member {string} account_id
         * @memberof mailkon.User
         * @instance
         */
        User.prototype.account_id = "";

        /**
         * User user_id.
         * @member {string} user_id
         * @memberof mailkon.User
         * @instance
         */
        User.prototype.user_id = "";

        /**
         * User email_address.
         * @member {string} email_address
         * @memberof mailkon.User
         * @instance
         */
        User.prototype.email_address = "";

        return User;
    })();

    /**
     * Event enum.
     * @name mailkon.Event
     * @enum {string}
     * @property {number} MaikonRequested=1000 MaikonRequested value
     * @property {number} MailkonSynced=1001 MailkonSynced value
     * @property {number} MailkonJob=2 MailkonJob value
     * @property {number} MailkonCreateAccountRequested=0 MailkonCreateAccountRequested value
     */
    mailkon.Event = (function() {
        const valuesById = {}, values = Object.create(valuesById);
        values[valuesById[1000] = "MaikonRequested"] = 1000;
        values[valuesById[1001] = "MailkonSynced"] = 1001;
        values[valuesById[2] = "MailkonJob"] = 2;
        values[valuesById[0] = "MailkonCreateAccountRequested"] = 0;
        return values;
    })();

    mailkon.SendgridEvent = (function() {

        /**
         * Properties of a SendgridEvent.
         * @memberof mailkon
         * @interface ISendgridEvent
         * @property {string|null} [event_id] SendgridEvent event_id
         * @property {string|null} [account_id] SendgridEvent account_id
         * @property {string|null} [conversation_id] SendgridEvent conversation_id
         * @property {string|null} [message_id] SendgridEvent message_id
         * @property {Array.<string>|null} [data] SendgridEvent data
         * @property {string|null} [full_message_id] SendgridEvent full_message_id
         * @property {string|null} [subject] SendgridEvent subject
         */

        /**
         * Constructs a new SendgridEvent.
         * @memberof mailkon
         * @classdesc Represents a SendgridEvent.
         * @implements ISendgridEvent
         * @constructor
         * @param {mailkon.ISendgridEvent=} [p] Properties to set
         */
        function SendgridEvent(p) {
            this.data = [];
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * SendgridEvent event_id.
         * @member {string} event_id
         * @memberof mailkon.SendgridEvent
         * @instance
         */
        SendgridEvent.prototype.event_id = "";

        /**
         * SendgridEvent account_id.
         * @member {string} account_id
         * @memberof mailkon.SendgridEvent
         * @instance
         */
        SendgridEvent.prototype.account_id = "";

        /**
         * SendgridEvent conversation_id.
         * @member {string} conversation_id
         * @memberof mailkon.SendgridEvent
         * @instance
         */
        SendgridEvent.prototype.conversation_id = "";

        /**
         * SendgridEvent message_id.
         * @member {string} message_id
         * @memberof mailkon.SendgridEvent
         * @instance
         */
        SendgridEvent.prototype.message_id = "";

        /**
         * SendgridEvent data.
         * @member {Array.<string>} data
         * @memberof mailkon.SendgridEvent
         * @instance
         */
        SendgridEvent.prototype.data = $util.emptyArray;

        /**
         * SendgridEvent full_message_id.
         * @member {string} full_message_id
         * @memberof mailkon.SendgridEvent
         * @instance
         */
        SendgridEvent.prototype.full_message_id = "";

        /**
         * SendgridEvent subject.
         * @member {string} subject
         * @memberof mailkon.SendgridEvent
         * @instance
         */
        SendgridEvent.prototype.subject = "";

        return SendgridEvent;
    })();

    mailkon.UserAvail = (function() {

        /**
         * Properties of a UserAvail.
         * @memberof mailkon
         * @interface IUserAvail
         * @property {string|null} [account_id] UserAvail account_id
         * @property {string|null} [email_address] UserAvail email_address
         * @property {boolean|null} [availability] UserAvail availability
         * @property {number|Long|null} [updated] UserAvail updated
         */

        /**
         * Constructs a new UserAvail.
         * @memberof mailkon
         * @classdesc Represents a UserAvail.
         * @implements IUserAvail
         * @constructor
         * @param {mailkon.IUserAvail=} [p] Properties to set
         */
        function UserAvail(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * UserAvail account_id.
         * @member {string} account_id
         * @memberof mailkon.UserAvail
         * @instance
         */
        UserAvail.prototype.account_id = "";

        /**
         * UserAvail email_address.
         * @member {string} email_address
         * @memberof mailkon.UserAvail
         * @instance
         */
        UserAvail.prototype.email_address = "";

        /**
         * UserAvail availability.
         * @member {boolean} availability
         * @memberof mailkon.UserAvail
         * @instance
         */
        UserAvail.prototype.availability = false;

        /**
         * UserAvail updated.
         * @member {number|Long} updated
         * @memberof mailkon.UserAvail
         * @instance
         */
        UserAvail.prototype.updated = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

        return UserAvail;
    })();

    mailkon.HttpChunk = (function() {

        /**
         * Properties of a HttpChunk.
         * @memberof mailkon
         * @interface IHttpChunk
         * @property {string|null} [id] HttpChunk id
         * @property {number|null} [chunk_id] HttpChunk chunk_id
         * @property {Uint8Array|null} [data] HttpChunk data
         * @property {number|null} [chunk_size] HttpChunk chunk_size
         */

        /**
         * Constructs a new HttpChunk.
         * @memberof mailkon
         * @classdesc Represents a HttpChunk.
         * @implements IHttpChunk
         * @constructor
         * @param {mailkon.IHttpChunk=} [p] Properties to set
         */
        function HttpChunk(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * HttpChunk id.
         * @member {string} id
         * @memberof mailkon.HttpChunk
         * @instance
         */
        HttpChunk.prototype.id = "";

        /**
         * HttpChunk chunk_id.
         * @member {number} chunk_id
         * @memberof mailkon.HttpChunk
         * @instance
         */
        HttpChunk.prototype.chunk_id = 0;

        /**
         * HttpChunk data.
         * @member {Uint8Array} data
         * @memberof mailkon.HttpChunk
         * @instance
         */
        HttpChunk.prototype.data = $util.newBuffer([]);

        /**
         * HttpChunk chunk_size.
         * @member {number} chunk_size
         * @memberof mailkon.HttpChunk
         * @instance
         */
        HttpChunk.prototype.chunk_size = 0;

        return HttpChunk;
    })();

    mailkon.Job = (function() {

        /**
         * Properties of a Job.
         * @memberof mailkon
         * @interface IJob
         * @property {string|null} [topic] Job topic
         * @property {number|null} [partition] Job partition
         * @property {number|Long|null} [offset] Job offset
         * @property {string|null} [content_type] Job content_type
         * @property {string|null} [type] Job type
         * @property {string|null} [request_id] Job request_id
         */

        /**
         * Constructs a new Job.
         * @memberof mailkon
         * @classdesc Represents a Job.
         * @implements IJob
         * @constructor
         * @param {mailkon.IJob=} [p] Properties to set
         */
        function Job(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * Job topic.
         * @member {string} topic
         * @memberof mailkon.Job
         * @instance
         */
        Job.prototype.topic = "";

        /**
         * Job partition.
         * @member {number} partition
         * @memberof mailkon.Job
         * @instance
         */
        Job.prototype.partition = 0;

        /**
         * Job offset.
         * @member {number|Long} offset
         * @memberof mailkon.Job
         * @instance
         */
        Job.prototype.offset = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

        /**
         * Job content_type.
         * @member {string} content_type
         * @memberof mailkon.Job
         * @instance
         */
        Job.prototype.content_type = "";

        /**
         * Job type.
         * @member {string} type
         * @memberof mailkon.Job
         * @instance
         */
        Job.prototype.type = "";

        /**
         * Job request_id.
         * @member {string} request_id
         * @memberof mailkon.Job
         * @instance
         */
        Job.prototype.request_id = "";

        return Job;
    })();

    /**
     * JobType enum.
     * @name mailkon.JobType
     * @enum {string}
     * @property {number} sendgrid_inbound=0 sendgrid_inbound value
     * @property {number} sendgrid_event=1 sendgrid_event value
     * @property {number} subiz_event=3 subiz_event value
     */
    mailkon.JobType = (function() {
        const valuesById = {}, values = Object.create(valuesById);
        values[valuesById[0] = "sendgrid_inbound"] = 0;
        values[valuesById[1] = "sendgrid_event"] = 1;
        values[valuesById[3] = "subiz_event"] = 3;
        return values;
    })();

    return mailkon;
})();

export const noti5 = $root.noti5 = (() => {

    /**
     * Namespace noti5.
     * @exports noti5
     * @namespace
     */
    const noti5 = {};

    /**
     * Type enum.
     * @name noti5.Type
     * @enum {string}
     * @property {number} new_message=0 new_message value
     * @property {number} new_conversation=1 new_conversation value
     * @property {number} nothing=2 nothing value
     */
    noti5.Type = (function() {
        const valuesById = {}, values = Object.create(valuesById);
        values[valuesById[0] = "new_message"] = 0;
        values[valuesById[1] = "new_conversation"] = 1;
        values[valuesById[2] = "nothing"] = 2;
        return values;
    })();

    noti5.Setting = (function() {

        /**
         * Properties of a Setting.
         * @memberof noti5
         * @interface ISetting
         * @property {common.IContext|null} [ctx] Setting ctx
         * @property {string|null} [account_id] Setting account_id
         * @property {string|null} [agent_id] Setting agent_id
         * @property {string|null} [web_type] Setting web_type
         * @property {string|null} [mobile_type] Setting mobile_type
         * @property {string|null} [email_type] Setting email_type
         * @property {number|null} [email_after] Setting email_after
         */

        /**
         * Constructs a new Setting.
         * @memberof noti5
         * @classdesc Represents a Setting.
         * @implements ISetting
         * @constructor
         * @param {noti5.ISetting=} [p] Properties to set
         */
        function Setting(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * Setting ctx.
         * @member {common.IContext|null|undefined} ctx
         * @memberof noti5.Setting
         * @instance
         */
        Setting.prototype.ctx = null;

        /**
         * Setting account_id.
         * @member {string} account_id
         * @memberof noti5.Setting
         * @instance
         */
        Setting.prototype.account_id = "";

        /**
         * Setting agent_id.
         * @member {string} agent_id
         * @memberof noti5.Setting
         * @instance
         */
        Setting.prototype.agent_id = "";

        /**
         * Setting web_type.
         * @member {string} web_type
         * @memberof noti5.Setting
         * @instance
         */
        Setting.prototype.web_type = "";

        /**
         * Setting mobile_type.
         * @member {string} mobile_type
         * @memberof noti5.Setting
         * @instance
         */
        Setting.prototype.mobile_type = "";

        /**
         * Setting email_type.
         * @member {string} email_type
         * @memberof noti5.Setting
         * @instance
         */
        Setting.prototype.email_type = "";

        /**
         * Setting email_after.
         * @member {number} email_after
         * @memberof noti5.Setting
         * @instance
         */
        Setting.prototype.email_after = 0;

        return Setting;
    })();

    noti5.Noti5Service = (function() {

        /**
         * Constructs a new Noti5Service service.
         * @memberof noti5
         * @classdesc Represents a Noti5Service
         * @extends $protobuf.rpc.Service
         * @constructor
         * @param {$protobuf.RPCImpl} rpcImpl RPC implementation
         * @param {boolean} [requestDelimited=false] Whether requests are length-delimited
         * @param {boolean} [responseDelimited=false] Whether responses are length-delimited
         */
        function Noti5Service(rpcImpl, requestDelimited, responseDelimited) {
            $protobuf.rpc.Service.call(this, rpcImpl, requestDelimited, responseDelimited);
        }

        (Noti5Service.prototype = Object.create($protobuf.rpc.Service.prototype)).constructor = Noti5Service;

        /**
         * Callback as used by {@link noti5.Noti5Service#readNotificationSetting}.
         * @memberof noti5.Noti5Service
         * @typedef ReadNotificationSettingCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {noti5.Setting} [response] Setting
         */

        /**
         * Calls ReadNotificationSetting.
         * @function readNotificationSetting
         * @memberof noti5.Noti5Service
         * @instance
         * @param {common.IEmpty} request Empty message or plain object
         * @param {noti5.Noti5Service.ReadNotificationSettingCallback} callback Node-style callback called with the error, if any, and Setting
         * @returns {undefined}
         * @variation 1
         */
        Noti5Service.prototype.readNotificationSetting = function readNotificationSetting(request, callback) {
            return this.rpcCall(readNotificationSetting, $root.common.Empty, $root.noti5.Setting, request, callback);
        };

        /**
         * Calls ReadNotificationSetting.
         * @function readNotificationSetting
         * @memberof noti5.Noti5Service
         * @instance
         * @param {common.IEmpty} request Empty message or plain object
         * @returns {Promise<noti5.Setting>} Promise
         * @variation 2
         */

        /**
         * Callback as used by {@link noti5.Noti5Service#updateNotificationSetting}.
         * @memberof noti5.Noti5Service
         * @typedef UpdateNotificationSettingCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {common.Empty} [response] Empty
         */

        /**
         * Calls UpdateNotificationSetting.
         * @function updateNotificationSetting
         * @memberof noti5.Noti5Service
         * @instance
         * @param {noti5.ISetting} request Setting message or plain object
         * @param {noti5.Noti5Service.UpdateNotificationSettingCallback} callback Node-style callback called with the error, if any, and Empty
         * @returns {undefined}
         * @variation 1
         */
        Noti5Service.prototype.updateNotificationSetting = function updateNotificationSetting(request, callback) {
            return this.rpcCall(updateNotificationSetting, $root.noti5.Setting, $root.common.Empty, request, callback);
        };

        /**
         * Calls UpdateNotificationSetting.
         * @function updateNotificationSetting
         * @memberof noti5.Noti5Service
         * @instance
         * @param {noti5.ISetting} request Setting message or plain object
         * @returns {Promise<common.Empty>} Promise
         * @variation 2
         */

        return Noti5Service;
    })();

    return noti5;
})();

export const notibox = $root.notibox = (() => {

    /**
     * Namespace notibox.
     * @exports notibox
     * @namespace
     */
    const notibox = {};

    notibox.Notification = (function() {

        /**
         * Properties of a Notification.
         * @memberof notibox
         * @interface INotification
         * @property {common.IContext|null} [ctx] Notification ctx
         * @property {string|null} [box] Notification box
         * @property {string|null} [topic] Notification topic
         * @property {string|null} [type] Notification type
         * @property {string|null} [data] Notification data
         * @property {number|Long|null} [created] Notification created
         * @property {number|Long|null} [read] Notification read
         * @property {boolean|null} [view] Notification view
         */

        /**
         * Constructs a new Notification.
         * @memberof notibox
         * @classdesc Represents a Notification.
         * @implements INotification
         * @constructor
         * @param {notibox.INotification=} [p] Properties to set
         */
        function Notification(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * Notification ctx.
         * @member {common.IContext|null|undefined} ctx
         * @memberof notibox.Notification
         * @instance
         */
        Notification.prototype.ctx = null;

        /**
         * Notification box.
         * @member {string} box
         * @memberof notibox.Notification
         * @instance
         */
        Notification.prototype.box = "";

        /**
         * Notification topic.
         * @member {string} topic
         * @memberof notibox.Notification
         * @instance
         */
        Notification.prototype.topic = "";

        /**
         * Notification type.
         * @member {string} type
         * @memberof notibox.Notification
         * @instance
         */
        Notification.prototype.type = "";

        /**
         * Notification data.
         * @member {string} data
         * @memberof notibox.Notification
         * @instance
         */
        Notification.prototype.data = "";

        /**
         * Notification created.
         * @member {number|Long} created
         * @memberof notibox.Notification
         * @instance
         */
        Notification.prototype.created = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

        /**
         * Notification read.
         * @member {number|Long} read
         * @memberof notibox.Notification
         * @instance
         */
        Notification.prototype.read = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

        /**
         * Notification view.
         * @member {boolean} view
         * @memberof notibox.Notification
         * @instance
         */
        Notification.prototype.view = false;

        return Notification;
    })();

    /**
     * Type enum.
     * @name notibox.Type
     * @enum {string}
     * @property {number} account_created=0 account_created value
     * @property {number} trial_almost_expired=2 trial_almost_expired value
     * @property {number} trial_expired=3 trial_expired value
     * @property {number} system_maintainance_scheduled_1=4 system_maintainance_scheduled_1 value
     * @property {number} system_maintainance_scheduled_2=5 system_maintainance_scheduled_2 value
     * @property {number} system_maintainance_completed=6 system_maintainance_completed value
     * @property {number} agent_activated=7 agent_activated value
     * @property {number} conversation_unassigned=8 conversation_unassigned value
     * @property {number} agent_permission_updated=9 agent_permission_updated value
     */
    notibox.Type = (function() {
        const valuesById = {}, values = Object.create(valuesById);
        values[valuesById[0] = "account_created"] = 0;
        values[valuesById[2] = "trial_almost_expired"] = 2;
        values[valuesById[3] = "trial_expired"] = 3;
        values[valuesById[4] = "system_maintainance_scheduled_1"] = 4;
        values[valuesById[5] = "system_maintainance_scheduled_2"] = 5;
        values[valuesById[6] = "system_maintainance_completed"] = 6;
        values[valuesById[7] = "agent_activated"] = 7;
        values[valuesById[8] = "conversation_unassigned"] = 8;
        values[valuesById[9] = "agent_permission_updated"] = 9;
        return values;
    })();

    notibox.AddNotificationRequest = (function() {

        /**
         * Properties of an AddNotificationRequest.
         * @memberof notibox
         * @interface IAddNotificationRequest
         * @property {common.IContext|null} [ctx] AddNotificationRequest ctx
         * @property {Array.<string>|null} [boxs] AddNotificationRequest boxs
         * @property {notibox.INotification|null} [notification] AddNotificationRequest notification
         */

        /**
         * Constructs a new AddNotificationRequest.
         * @memberof notibox
         * @classdesc Represents an AddNotificationRequest.
         * @implements IAddNotificationRequest
         * @constructor
         * @param {notibox.IAddNotificationRequest=} [p] Properties to set
         */
        function AddNotificationRequest(p) {
            this.boxs = [];
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * AddNotificationRequest ctx.
         * @member {common.IContext|null|undefined} ctx
         * @memberof notibox.AddNotificationRequest
         * @instance
         */
        AddNotificationRequest.prototype.ctx = null;

        /**
         * AddNotificationRequest boxs.
         * @member {Array.<string>} boxs
         * @memberof notibox.AddNotificationRequest
         * @instance
         */
        AddNotificationRequest.prototype.boxs = $util.emptyArray;

        /**
         * AddNotificationRequest notification.
         * @member {notibox.INotification|null|undefined} notification
         * @memberof notibox.AddNotificationRequest
         * @instance
         */
        AddNotificationRequest.prototype.notification = null;

        return AddNotificationRequest;
    })();

    notibox.ReadNotification = (function() {

        /**
         * Properties of a ReadNotification.
         * @memberof notibox
         * @interface IReadNotification
         * @property {common.IContext|null} [ctx] ReadNotification ctx
         * @property {string|null} [box] ReadNotification box
         * @property {Array.<string>|null} [topics] ReadNotification topics
         * @property {boolean|null} [read] ReadNotification read
         */

        /**
         * Constructs a new ReadNotification.
         * @memberof notibox
         * @classdesc Represents a ReadNotification.
         * @implements IReadNotification
         * @constructor
         * @param {notibox.IReadNotification=} [p] Properties to set
         */
        function ReadNotification(p) {
            this.topics = [];
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * ReadNotification ctx.
         * @member {common.IContext|null|undefined} ctx
         * @memberof notibox.ReadNotification
         * @instance
         */
        ReadNotification.prototype.ctx = null;

        /**
         * ReadNotification box.
         * @member {string} box
         * @memberof notibox.ReadNotification
         * @instance
         */
        ReadNotification.prototype.box = "";

        /**
         * ReadNotification topics.
         * @member {Array.<string>} topics
         * @memberof notibox.ReadNotification
         * @instance
         */
        ReadNotification.prototype.topics = $util.emptyArray;

        /**
         * ReadNotification read.
         * @member {boolean} read
         * @memberof notibox.ReadNotification
         * @instance
         */
        ReadNotification.prototype.read = false;

        return ReadNotification;
    })();

    notibox.ListRequest = (function() {

        /**
         * Properties of a ListRequest.
         * @memberof notibox
         * @interface IListRequest
         * @property {common.IContext|null} [ctx] ListRequest ctx
         * @property {string|null} [box] ListRequest box
         * @property {string|null} [anchor] ListRequest anchor
         * @property {number|null} [limit] ListRequest limit
         */

        /**
         * Constructs a new ListRequest.
         * @memberof notibox
         * @classdesc Represents a ListRequest.
         * @implements IListRequest
         * @constructor
         * @param {notibox.IListRequest=} [p] Properties to set
         */
        function ListRequest(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * ListRequest ctx.
         * @member {common.IContext|null|undefined} ctx
         * @memberof notibox.ListRequest
         * @instance
         */
        ListRequest.prototype.ctx = null;

        /**
         * ListRequest box.
         * @member {string} box
         * @memberof notibox.ListRequest
         * @instance
         */
        ListRequest.prototype.box = "";

        /**
         * ListRequest anchor.
         * @member {string} anchor
         * @memberof notibox.ListRequest
         * @instance
         */
        ListRequest.prototype.anchor = "";

        /**
         * ListRequest limit.
         * @member {number} limit
         * @memberof notibox.ListRequest
         * @instance
         */
        ListRequest.prototype.limit = 0;

        return ListRequest;
    })();

    notibox.NotiboxRequest = (function() {

        /**
         * Properties of a NotiboxRequest.
         * @memberof notibox
         * @interface INotiboxRequest
         * @property {common.IContext|null} [ctx] NotiboxRequest ctx
         * @property {string|null} [box] NotiboxRequest box
         * @property {string|null} [topic] NotiboxRequest topic
         */

        /**
         * Constructs a new NotiboxRequest.
         * @memberof notibox
         * @classdesc Represents a NotiboxRequest.
         * @implements INotiboxRequest
         * @constructor
         * @param {notibox.INotiboxRequest=} [p] Properties to set
         */
        function NotiboxRequest(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * NotiboxRequest ctx.
         * @member {common.IContext|null|undefined} ctx
         * @memberof notibox.NotiboxRequest
         * @instance
         */
        NotiboxRequest.prototype.ctx = null;

        /**
         * NotiboxRequest box.
         * @member {string} box
         * @memberof notibox.NotiboxRequest
         * @instance
         */
        NotiboxRequest.prototype.box = "";

        /**
         * NotiboxRequest topic.
         * @member {string} topic
         * @memberof notibox.NotiboxRequest
         * @instance
         */
        NotiboxRequest.prototype.topic = "";

        return NotiboxRequest;
    })();

    notibox.Notifications = (function() {

        /**
         * Properties of a Notifications.
         * @memberof notibox
         * @interface INotifications
         * @property {common.IContext|null} [ctx] Notifications ctx
         * @property {Array.<notibox.INotification>|null} [notifications] Notifications notifications
         * @property {string|null} [anchor] Notifications anchor
         */

        /**
         * Constructs a new Notifications.
         * @memberof notibox
         * @classdesc Represents a Notifications.
         * @implements INotifications
         * @constructor
         * @param {notibox.INotifications=} [p] Properties to set
         */
        function Notifications(p) {
            this.notifications = [];
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * Notifications ctx.
         * @member {common.IContext|null|undefined} ctx
         * @memberof notibox.Notifications
         * @instance
         */
        Notifications.prototype.ctx = null;

        /**
         * Notifications notifications.
         * @member {Array.<notibox.INotification>} notifications
         * @memberof notibox.Notifications
         * @instance
         */
        Notifications.prototype.notifications = $util.emptyArray;

        /**
         * Notifications anchor.
         * @member {string} anchor
         * @memberof notibox.Notifications
         * @instance
         */
        Notifications.prototype.anchor = "";

        return Notifications;
    })();

    notibox.Box = (function() {

        /**
         * Properties of a Box.
         * @memberof notibox
         * @interface IBox
         * @property {common.IContext|null} [ctx] Box ctx
         * @property {string|null} [box] Box box
         * @property {number|Long|null} [new_count] Box new_count
         */

        /**
         * Constructs a new Box.
         * @memberof notibox
         * @classdesc Represents a Box.
         * @implements IBox
         * @constructor
         * @param {notibox.IBox=} [p] Properties to set
         */
        function Box(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * Box ctx.
         * @member {common.IContext|null|undefined} ctx
         * @memberof notibox.Box
         * @instance
         */
        Box.prototype.ctx = null;

        /**
         * Box box.
         * @member {string} box
         * @memberof notibox.Box
         * @instance
         */
        Box.prototype.box = "";

        /**
         * Box new_count.
         * @member {number|Long} new_count
         * @memberof notibox.Box
         * @instance
         */
        Box.prototype.new_count = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

        return Box;
    })();

    notibox.AllType = (function() {

        /**
         * Properties of an AllType.
         * @memberof notibox
         * @interface IAllType
         * @property {notibox.INotifications|null} [nts] AllType nts
         * @property {notibox.IListRequest|null} [listr] AllType listr
         * @property {notibox.IReadNotification|null} [rno] AllType rno
         * @property {notibox.IAddNotificationRequest|null} [anorfr] AllType anorfr
         * @property {notibox.INotification|null} [noti] AllType noti
         * @property {notibox.IBox|null} [box] AllType box
         */

        /**
         * Constructs a new AllType.
         * @memberof notibox
         * @classdesc Represents an AllType.
         * @implements IAllType
         * @constructor
         * @param {notibox.IAllType=} [p] Properties to set
         */
        function AllType(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * AllType nts.
         * @member {notibox.INotifications|null|undefined} nts
         * @memberof notibox.AllType
         * @instance
         */
        AllType.prototype.nts = null;

        /**
         * AllType listr.
         * @member {notibox.IListRequest|null|undefined} listr
         * @memberof notibox.AllType
         * @instance
         */
        AllType.prototype.listr = null;

        /**
         * AllType rno.
         * @member {notibox.IReadNotification|null|undefined} rno
         * @memberof notibox.AllType
         * @instance
         */
        AllType.prototype.rno = null;

        /**
         * AllType anorfr.
         * @member {notibox.IAddNotificationRequest|null|undefined} anorfr
         * @memberof notibox.AllType
         * @instance
         */
        AllType.prototype.anorfr = null;

        /**
         * AllType noti.
         * @member {notibox.INotification|null|undefined} noti
         * @memberof notibox.AllType
         * @instance
         */
        AllType.prototype.noti = null;

        /**
         * AllType box.
         * @member {notibox.IBox|null|undefined} box
         * @memberof notibox.AllType
         * @instance
         */
        AllType.prototype.box = null;

        return AllType;
    })();

    notibox.YourService = (function() {

        /**
         * Constructs a new YourService service.
         * @memberof notibox
         * @classdesc Represents a YourService
         * @extends $protobuf.rpc.Service
         * @constructor
         * @param {$protobuf.RPCImpl} rpcImpl RPC implementation
         * @param {boolean} [requestDelimited=false] Whether requests are length-delimited
         * @param {boolean} [responseDelimited=false] Whether responses are length-delimited
         */
        function YourService(rpcImpl, requestDelimited, responseDelimited) {
            $protobuf.rpc.Service.call(this, rpcImpl, requestDelimited, responseDelimited);
        }

        (YourService.prototype = Object.create($protobuf.rpc.Service.prototype)).constructor = YourService;

        /**
         * Callback as used by {@link notibox.YourService#do_}.
         * @memberof notibox.YourService
         * @typedef DoCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {notibox.AllType} [response] AllType
         */

        /**
         * Calls Do.
         * @function do
         * @memberof notibox.YourService
         * @instance
         * @param {notibox.IAllType} request AllType message or plain object
         * @param {notibox.YourService.DoCallback} callback Node-style callback called with the error, if any, and AllType
         * @returns {undefined}
         * @variation 1
         */
        YourService.prototype["do"] = function do_(request, callback) {
            return this.rpcCall(do_, $root.notibox.AllType, $root.notibox.AllType, request, callback);
        };

        /**
         * Calls Do.
         * @function do
         * @memberof notibox.YourService
         * @instance
         * @param {notibox.IAllType} request AllType message or plain object
         * @returns {Promise<notibox.AllType>} Promise
         * @variation 2
         */

        return YourService;
    })();

    notibox.NotiboxService = (function() {

        /**
         * Constructs a new NotiboxService service.
         * @memberof notibox
         * @classdesc Represents a NotiboxService
         * @extends $protobuf.rpc.Service
         * @constructor
         * @param {$protobuf.RPCImpl} rpcImpl RPC implementation
         * @param {boolean} [requestDelimited=false] Whether requests are length-delimited
         * @param {boolean} [responseDelimited=false] Whether responses are length-delimited
         */
        function NotiboxService(rpcImpl, requestDelimited, responseDelimited) {
            $protobuf.rpc.Service.call(this, rpcImpl, requestDelimited, responseDelimited);
        }

        (NotiboxService.prototype = Object.create($protobuf.rpc.Service.prototype)).constructor = NotiboxService;

        /**
         * Callback as used by {@link notibox.NotiboxService#upsert}.
         * @memberof notibox.NotiboxService
         * @typedef UpsertCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {notibox.Notification} [response] Notification
         */

        /**
         * Calls Upsert.
         * @function upsert
         * @memberof notibox.NotiboxService
         * @instance
         * @param {notibox.INotification} request Notification message or plain object
         * @param {notibox.NotiboxService.UpsertCallback} callback Node-style callback called with the error, if any, and Notification
         * @returns {undefined}
         * @variation 1
         */
        NotiboxService.prototype.upsert = function upsert(request, callback) {
            return this.rpcCall(upsert, $root.notibox.Notification, $root.notibox.Notification, request, callback);
        };

        /**
         * Calls Upsert.
         * @function upsert
         * @memberof notibox.NotiboxService
         * @instance
         * @param {notibox.INotification} request Notification message or plain object
         * @returns {Promise<notibox.Notification>} Promise
         * @variation 2
         */

        /**
         * Callback as used by {@link notibox.NotiboxService#readTopic}.
         * @memberof notibox.NotiboxService
         * @typedef ReadTopicCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {common.Empty} [response] Empty
         */

        /**
         * Calls ReadTopic.
         * @function readTopic
         * @memberof notibox.NotiboxService
         * @instance
         * @param {notibox.INotiboxRequest} request NotiboxRequest message or plain object
         * @param {notibox.NotiboxService.ReadTopicCallback} callback Node-style callback called with the error, if any, and Empty
         * @returns {undefined}
         * @variation 1
         */
        NotiboxService.prototype.readTopic = function readTopic(request, callback) {
            return this.rpcCall(readTopic, $root.notibox.NotiboxRequest, $root.common.Empty, request, callback);
        };

        /**
         * Calls ReadTopic.
         * @function readTopic
         * @memberof notibox.NotiboxService
         * @instance
         * @param {notibox.INotiboxRequest} request NotiboxRequest message or plain object
         * @returns {Promise<common.Empty>} Promise
         * @variation 2
         */

        /**
         * Callback as used by {@link notibox.NotiboxService#unreadTopic}.
         * @memberof notibox.NotiboxService
         * @typedef UnreadTopicCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {common.Empty} [response] Empty
         */

        /**
         * Calls UnreadTopic.
         * @function unreadTopic
         * @memberof notibox.NotiboxService
         * @instance
         * @param {notibox.INotiboxRequest} request NotiboxRequest message or plain object
         * @param {notibox.NotiboxService.UnreadTopicCallback} callback Node-style callback called with the error, if any, and Empty
         * @returns {undefined}
         * @variation 1
         */
        NotiboxService.prototype.unreadTopic = function unreadTopic(request, callback) {
            return this.rpcCall(unreadTopic, $root.notibox.NotiboxRequest, $root.common.Empty, request, callback);
        };

        /**
         * Calls UnreadTopic.
         * @function unreadTopic
         * @memberof notibox.NotiboxService
         * @instance
         * @param {notibox.INotiboxRequest} request NotiboxRequest message or plain object
         * @returns {Promise<common.Empty>} Promise
         * @variation 2
         */

        /**
         * Callback as used by {@link notibox.NotiboxService#listNotis}.
         * @memberof notibox.NotiboxService
         * @typedef ListNotisCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {notibox.Notifications} [response] Notifications
         */

        /**
         * Calls ListNotis.
         * @function listNotis
         * @memberof notibox.NotiboxService
         * @instance
         * @param {notibox.IListRequest} request ListRequest message or plain object
         * @param {notibox.NotiboxService.ListNotisCallback} callback Node-style callback called with the error, if any, and Notifications
         * @returns {undefined}
         * @variation 1
         */
        NotiboxService.prototype.listNotis = function listNotis(request, callback) {
            return this.rpcCall(listNotis, $root.notibox.ListRequest, $root.notibox.Notifications, request, callback);
        };

        /**
         * Calls ListNotis.
         * @function listNotis
         * @memberof notibox.NotiboxService
         * @instance
         * @param {notibox.IListRequest} request ListRequest message or plain object
         * @returns {Promise<notibox.Notifications>} Promise
         * @variation 2
         */

        /**
         * Callback as used by {@link notibox.NotiboxService#decreaseNew}.
         * @memberof notibox.NotiboxService
         * @typedef DecreaseNewCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {notibox.Box} [response] Box
         */

        /**
         * Calls DecreaseNew.
         * @function decreaseNew
         * @memberof notibox.NotiboxService
         * @instance
         * @param {notibox.INotiboxRequest} request NotiboxRequest message or plain object
         * @param {notibox.NotiboxService.DecreaseNewCallback} callback Node-style callback called with the error, if any, and Box
         * @returns {undefined}
         * @variation 1
         */
        NotiboxService.prototype.decreaseNew = function decreaseNew(request, callback) {
            return this.rpcCall(decreaseNew, $root.notibox.NotiboxRequest, $root.notibox.Box, request, callback);
        };

        /**
         * Calls DecreaseNew.
         * @function decreaseNew
         * @memberof notibox.NotiboxService
         * @instance
         * @param {notibox.INotiboxRequest} request NotiboxRequest message or plain object
         * @returns {Promise<notibox.Box>} Promise
         * @variation 2
         */

        /**
         * Callback as used by {@link notibox.NotiboxService#readBox}.
         * @memberof notibox.NotiboxService
         * @typedef ReadBoxCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {notibox.Box} [response] Box
         */

        /**
         * Calls ReadBox.
         * @function readBox
         * @memberof notibox.NotiboxService
         * @instance
         * @param {notibox.INotiboxRequest} request NotiboxRequest message or plain object
         * @param {notibox.NotiboxService.ReadBoxCallback} callback Node-style callback called with the error, if any, and Box
         * @returns {undefined}
         * @variation 1
         */
        NotiboxService.prototype.readBox = function readBox(request, callback) {
            return this.rpcCall(readBox, $root.notibox.NotiboxRequest, $root.notibox.Box, request, callback);
        };

        /**
         * Calls ReadBox.
         * @function readBox
         * @memberof notibox.NotiboxService
         * @instance
         * @param {notibox.INotiboxRequest} request NotiboxRequest message or plain object
         * @returns {Promise<notibox.Box>} Promise
         * @variation 2
         */

        return NotiboxService;
    })();

    /**
     * Event enum.
     * @name notibox.Event
     * @enum {string}
     * @property {number} NotiboxRequested=0 NotiboxRequested value
     * @property {number} NotiUpsertRequested=1 NotiUpsertRequested value
     * @property {number} NotiReadRequested=2 NotiReadRequested value
     * @property {number} NotiUnreadRequested=4 NotiUnreadRequested value
     * @property {number} NotiListRequested=5 NotiListRequested value
     * @property {number} NotiDecreaseNewRequested=6 NotiDecreaseNewRequested value
     * @property {number} NotiNewReadRequested=7 NotiNewReadRequested value
     * @property {number} NotiboxSynced=101 NotiboxSynced value
     */
    notibox.Event = (function() {
        const valuesById = {}, values = Object.create(valuesById);
        values[valuesById[0] = "NotiboxRequested"] = 0;
        values[valuesById[1] = "NotiUpsertRequested"] = 1;
        values[valuesById[2] = "NotiReadRequested"] = 2;
        values[valuesById[4] = "NotiUnreadRequested"] = 4;
        values[valuesById[5] = "NotiListRequested"] = 5;
        values[valuesById[6] = "NotiDecreaseNewRequested"] = 6;
        values[valuesById[7] = "NotiNewReadRequested"] = 7;
        values[valuesById[101] = "NotiboxSynced"] = 101;
        return values;
    })();

    return notibox;
})();

export const payment = $root.payment = (() => {

    /**
     * Namespace payment.
     * @exports payment
     * @namespace
     */
    const payment = {};

    /**
     * Currency enum.
     * @name payment.Currency
     * @enum {string}
     * @property {number} usd=0 usd value
     * @property {number} vnd=1 vnd value
     * @property {number} brl=2 brl value
     */
    payment.Currency = (function() {
        const valuesById = {}, values = Object.create(valuesById);
        values[valuesById[0] = "usd"] = 0;
        values[valuesById[1] = "vnd"] = 1;
        values[valuesById[2] = "brl"] = 2;
        return values;
    })();

    payment.Stripe = (function() {

        /**
         * Properties of a Stripe.
         * @memberof payment
         * @interface IStripe
         * @property {string|null} [card_last4] Stripe card_last4
         * @property {string|null} [customer_id] Stripe customer_id
         * @property {string|null} [token] Stripe token
         */

        /**
         * Constructs a new Stripe.
         * @memberof payment
         * @classdesc Represents a Stripe.
         * @implements IStripe
         * @constructor
         * @param {payment.IStripe=} [p] Properties to set
         */
        function Stripe(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * Stripe card_last4.
         * @member {string} card_last4
         * @memberof payment.Stripe
         * @instance
         */
        Stripe.prototype.card_last4 = "";

        /**
         * Stripe customer_id.
         * @member {string} customer_id
         * @memberof payment.Stripe
         * @instance
         */
        Stripe.prototype.customer_id = "";

        /**
         * Stripe token.
         * @member {string} token
         * @memberof payment.Stripe
         * @instance
         */
        Stripe.prototype.token = "";

        return Stripe;
    })();

    payment.PaymentMethods = (function() {

        /**
         * Properties of a PaymentMethods.
         * @memberof payment
         * @interface IPaymentMethods
         * @property {Array.<payment.IPaymentMethod>|null} [payment_methods] PaymentMethods payment_methods
         */

        /**
         * Constructs a new PaymentMethods.
         * @memberof payment
         * @classdesc Represents a PaymentMethods.
         * @implements IPaymentMethods
         * @constructor
         * @param {payment.IPaymentMethods=} [p] Properties to set
         */
        function PaymentMethods(p) {
            this.payment_methods = [];
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * PaymentMethods payment_methods.
         * @member {Array.<payment.IPaymentMethod>} payment_methods
         * @memberof payment.PaymentMethods
         * @instance
         */
        PaymentMethods.prototype.payment_methods = $util.emptyArray;

        return PaymentMethods;
    })();

    payment.PaymentMethod = (function() {

        /**
         * Properties of a PaymentMethod.
         * @memberof payment
         * @interface IPaymentMethod
         * @property {common.IContext|null} [ctx] PaymentMethod ctx
         * @property {string|null} [type] PaymentMethod type
         * @property {string|null} [id] PaymentMethod id
         * @property {string|null} [account_id] PaymentMethod account_id
         * @property {string|null} [state] PaymentMethod state
         * @property {number|Long|null} [created] PaymentMethod created
         * @property {payment.IStripe|null} [stripe] PaymentMethod stripe
         * @property {string|null} [failed_message] PaymentMethod failed_message
         * @property {number|Long|null} [charged] PaymentMethod charged
         */

        /**
         * Constructs a new PaymentMethod.
         * @memberof payment
         * @classdesc Represents a PaymentMethod.
         * @implements IPaymentMethod
         * @constructor
         * @param {payment.IPaymentMethod=} [p] Properties to set
         */
        function PaymentMethod(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * PaymentMethod ctx.
         * @member {common.IContext|null|undefined} ctx
         * @memberof payment.PaymentMethod
         * @instance
         */
        PaymentMethod.prototype.ctx = null;

        /**
         * PaymentMethod type.
         * @member {string} type
         * @memberof payment.PaymentMethod
         * @instance
         */
        PaymentMethod.prototype.type = "";

        /**
         * PaymentMethod id.
         * @member {string} id
         * @memberof payment.PaymentMethod
         * @instance
         */
        PaymentMethod.prototype.id = "";

        /**
         * PaymentMethod account_id.
         * @member {string} account_id
         * @memberof payment.PaymentMethod
         * @instance
         */
        PaymentMethod.prototype.account_id = "";

        /**
         * PaymentMethod state.
         * @member {string} state
         * @memberof payment.PaymentMethod
         * @instance
         */
        PaymentMethod.prototype.state = "";

        /**
         * PaymentMethod created.
         * @member {number|Long} created
         * @memberof payment.PaymentMethod
         * @instance
         */
        PaymentMethod.prototype.created = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

        /**
         * PaymentMethod stripe.
         * @member {payment.IStripe|null|undefined} stripe
         * @memberof payment.PaymentMethod
         * @instance
         */
        PaymentMethod.prototype.stripe = null;

        /**
         * PaymentMethod failed_message.
         * @member {string} failed_message
         * @memberof payment.PaymentMethod
         * @instance
         */
        PaymentMethod.prototype.failed_message = "";

        /**
         * PaymentMethod charged.
         * @member {number|Long} charged
         * @memberof payment.PaymentMethod
         * @instance
         */
        PaymentMethod.prototype.charged = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

        /**
         * Type enum.
         * @name payment.PaymentMethod.Type
         * @enum {string}
         * @property {number} bank_transfer=0 bank_transfer value
         * @property {number} credit_card=1 credit_card value
         */
        PaymentMethod.Type = (function() {
            const valuesById = {}, values = Object.create(valuesById);
            values[valuesById[0] = "bank_transfer"] = 0;
            values[valuesById[1] = "credit_card"] = 1;
            return values;
        })();

        /**
         * State enum.
         * @name payment.PaymentMethod.State
         * @enum {string}
         * @property {number} active=0 active value
         * @property {number} failed=1 failed value
         */
        PaymentMethod.State = (function() {
            const valuesById = {}, values = Object.create(valuesById);
            values[valuesById[0] = "active"] = 0;
            values[valuesById[1] = "failed"] = 1;
            return values;
        })();

        return PaymentMethod;
    })();

    payment.Limit = (function() {

        /**
         * Properties of a Limit.
         * @memberof payment
         * @interface ILimit
         * @property {common.IContext|null} [ctx] Limit ctx
         * @property {string|null} [account_id] Limit account_id
         * @property {number|null} [max_automations] Limit max_automations
         * @property {number|null} [max_segments] Limit max_segments
         * @property {number|null} [max_agents] Limit max_agents
         * @property {boolean|null} [can_buy_agent] Limit can_buy_agent
         * @property {number|null} [automation_webhook_quota] Limit automation_webhook_quota
         * @property {number|null} [automation_email_quota] Limit automation_email_quota
         * @property {number|null} [automation_message_quota] Limit automation_message_quota
         * @property {number|null} [max_rules] Limit max_rules
         */

        /**
         * Constructs a new Limit.
         * @memberof payment
         * @classdesc Represents a Limit.
         * @implements ILimit
         * @constructor
         * @param {payment.ILimit=} [p] Properties to set
         */
        function Limit(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * Limit ctx.
         * @member {common.IContext|null|undefined} ctx
         * @memberof payment.Limit
         * @instance
         */
        Limit.prototype.ctx = null;

        /**
         * Limit account_id.
         * @member {string} account_id
         * @memberof payment.Limit
         * @instance
         */
        Limit.prototype.account_id = "";

        /**
         * Limit max_automations.
         * @member {number} max_automations
         * @memberof payment.Limit
         * @instance
         */
        Limit.prototype.max_automations = 0;

        /**
         * Limit max_segments.
         * @member {number} max_segments
         * @memberof payment.Limit
         * @instance
         */
        Limit.prototype.max_segments = 0;

        /**
         * Limit max_agents.
         * @member {number} max_agents
         * @memberof payment.Limit
         * @instance
         */
        Limit.prototype.max_agents = 0;

        /**
         * Limit can_buy_agent.
         * @member {boolean} can_buy_agent
         * @memberof payment.Limit
         * @instance
         */
        Limit.prototype.can_buy_agent = false;

        /**
         * Limit automation_webhook_quota.
         * @member {number} automation_webhook_quota
         * @memberof payment.Limit
         * @instance
         */
        Limit.prototype.automation_webhook_quota = 0;

        /**
         * Limit automation_email_quota.
         * @member {number} automation_email_quota
         * @memberof payment.Limit
         * @instance
         */
        Limit.prototype.automation_email_quota = 0;

        /**
         * Limit automation_message_quota.
         * @member {number} automation_message_quota
         * @memberof payment.Limit
         * @instance
         */
        Limit.prototype.automation_message_quota = 0;

        /**
         * Limit max_rules.
         * @member {number} max_rules
         * @memberof payment.Limit
         * @instance
         */
        Limit.prototype.max_rules = 0;

        return Limit;
    })();

    payment.Plans = (function() {

        /**
         * Properties of a Plans.
         * @memberof payment
         * @interface IPlans
         * @property {Array.<payment.IPlan>|null} [plans] Plans plans
         */

        /**
         * Constructs a new Plans.
         * @memberof payment
         * @classdesc Represents a Plans.
         * @implements IPlans
         * @constructor
         * @param {payment.IPlans=} [p] Properties to set
         */
        function Plans(p) {
            this.plans = [];
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * Plans plans.
         * @member {Array.<payment.IPlan>} plans
         * @memberof payment.Plans
         * @instance
         */
        Plans.prototype.plans = $util.emptyArray;

        return Plans;
    })();

    payment.Plan = (function() {

        /**
         * Properties of a Plan.
         * @memberof payment
         * @interface IPlan
         * @property {string|null} [name] Plan name
         * @property {payment.ILimit|null} [limit] Plan limit
         * @property {number|null} [price] Plan price
         * @property {number|null} [level] Plan level
         * @property {boolean|null} [can_buy_agent] Plan can_buy_agent
         * @property {boolean|null} [can_buy] Plan can_buy
         * @property {boolean|null} [has_start_time] Plan has_start_time
         */

        /**
         * Constructs a new Plan.
         * @memberof payment
         * @classdesc Represents a Plan.
         * @implements IPlan
         * @constructor
         * @param {payment.IPlan=} [p] Properties to set
         */
        function Plan(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * Plan name.
         * @member {string} name
         * @memberof payment.Plan
         * @instance
         */
        Plan.prototype.name = "";

        /**
         * Plan limit.
         * @member {payment.ILimit|null|undefined} limit
         * @memberof payment.Plan
         * @instance
         */
        Plan.prototype.limit = null;

        /**
         * Plan price.
         * @member {number} price
         * @memberof payment.Plan
         * @instance
         */
        Plan.prototype.price = 0;

        /**
         * Plan level.
         * @member {number} level
         * @memberof payment.Plan
         * @instance
         */
        Plan.prototype.level = 0;

        /**
         * Plan can_buy_agent.
         * @member {boolean} can_buy_agent
         * @memberof payment.Plan
         * @instance
         */
        Plan.prototype.can_buy_agent = false;

        /**
         * Plan can_buy.
         * @member {boolean} can_buy
         * @memberof payment.Plan
         * @instance
         */
        Plan.prototype.can_buy = false;

        /**
         * Plan has_start_time.
         * @member {boolean} has_start_time
         * @memberof payment.Plan
         * @instance
         */
        Plan.prototype.has_start_time = false;

        /**
         * Type enum.
         * @name payment.Plan.Type
         * @enum {string}
         * @property {number} trial=0 trial value
         * @property {number} free=1 free value
         * @property {number} standard=2 standard value
         * @property {number} advanced=3 advanced value
         */
        Plan.Type = (function() {
            const valuesById = {}, values = Object.create(valuesById);
            values[valuesById[0] = "trial"] = 0;
            values[valuesById[1] = "free"] = 1;
            values[valuesById[2] = "standard"] = 2;
            values[valuesById[3] = "advanced"] = 3;
            return values;
        })();

        return Plan;
    })();

    payment.Subscription = (function() {

        /**
         * Properties of a Subscription.
         * @memberof payment
         * @interface ISubscription
         * @property {common.IContext|null} [ctx] Subscription ctx
         * @property {string|null} [account_id] Subscription account_id
         * @property {number|Long|null} [created] Subscription created
         * @property {string|null} [promotion_code] Subscription promotion_code
         * @property {string|null} [name] Subscription name
         * @property {number|Long|null} [started] Subscription started
         * @property {number|Long|null} [due_date] Subscription due_date
         * @property {boolean|null} [auto_renew] Subscription auto_renew
         * @property {number|null} [billing_cycle_month] Subscription billing_cycle_month
         * @property {number|null} [next_billing_cycle_month] Subscription next_billing_cycle_month
         * @property {string|null} [plan] Subscription plan
         * @property {Array.<payment.IAddon>|null} [addons] Subscription addons
         * @property {number|null} [credit] Subscription credit
         * @property {Array.<payment.INote>|null} [notes] Subscription notes
         * @property {string|null} [referral_by] Subscription referral_by
         * @property {payment.ICustomer|null} [customer] Subscription customer
         * @property {string|null} [primary_payment_method] Subscription primary_payment_method
         * @property {payment.ILimit|null} [limit] Subscription limit
         * @property {number|null} [v3_state] Subscription v3_state
         */

        /**
         * Constructs a new Subscription.
         * @memberof payment
         * @classdesc Represents a Subscription.
         * @implements ISubscription
         * @constructor
         * @param {payment.ISubscription=} [p] Properties to set
         */
        function Subscription(p) {
            this.addons = [];
            this.notes = [];
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * Subscription ctx.
         * @member {common.IContext|null|undefined} ctx
         * @memberof payment.Subscription
         * @instance
         */
        Subscription.prototype.ctx = null;

        /**
         * Subscription account_id.
         * @member {string} account_id
         * @memberof payment.Subscription
         * @instance
         */
        Subscription.prototype.account_id = "";

        /**
         * Subscription created.
         * @member {number|Long} created
         * @memberof payment.Subscription
         * @instance
         */
        Subscription.prototype.created = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

        /**
         * Subscription promotion_code.
         * @member {string} promotion_code
         * @memberof payment.Subscription
         * @instance
         */
        Subscription.prototype.promotion_code = "";

        /**
         * Subscription name.
         * @member {string} name
         * @memberof payment.Subscription
         * @instance
         */
        Subscription.prototype.name = "";

        /**
         * Subscription started.
         * @member {number|Long} started
         * @memberof payment.Subscription
         * @instance
         */
        Subscription.prototype.started = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

        /**
         * Subscription due_date.
         * @member {number|Long} due_date
         * @memberof payment.Subscription
         * @instance
         */
        Subscription.prototype.due_date = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

        /**
         * Subscription auto_renew.
         * @member {boolean} auto_renew
         * @memberof payment.Subscription
         * @instance
         */
        Subscription.prototype.auto_renew = false;

        /**
         * Subscription billing_cycle_month.
         * @member {number} billing_cycle_month
         * @memberof payment.Subscription
         * @instance
         */
        Subscription.prototype.billing_cycle_month = 0;

        /**
         * Subscription next_billing_cycle_month.
         * @member {number} next_billing_cycle_month
         * @memberof payment.Subscription
         * @instance
         */
        Subscription.prototype.next_billing_cycle_month = 0;

        /**
         * Subscription plan.
         * @member {string} plan
         * @memberof payment.Subscription
         * @instance
         */
        Subscription.prototype.plan = "";

        /**
         * Subscription addons.
         * @member {Array.<payment.IAddon>} addons
         * @memberof payment.Subscription
         * @instance
         */
        Subscription.prototype.addons = $util.emptyArray;

        /**
         * Subscription credit.
         * @member {number} credit
         * @memberof payment.Subscription
         * @instance
         */
        Subscription.prototype.credit = 0;

        /**
         * Subscription notes.
         * @member {Array.<payment.INote>} notes
         * @memberof payment.Subscription
         * @instance
         */
        Subscription.prototype.notes = $util.emptyArray;

        /**
         * Subscription referral_by.
         * @member {string} referral_by
         * @memberof payment.Subscription
         * @instance
         */
        Subscription.prototype.referral_by = "";

        /**
         * Subscription customer.
         * @member {payment.ICustomer|null|undefined} customer
         * @memberof payment.Subscription
         * @instance
         */
        Subscription.prototype.customer = null;

        /**
         * Subscription primary_payment_method.
         * @member {string} primary_payment_method
         * @memberof payment.Subscription
         * @instance
         */
        Subscription.prototype.primary_payment_method = "";

        /**
         * Subscription limit.
         * @member {payment.ILimit|null|undefined} limit
         * @memberof payment.Subscription
         * @instance
         */
        Subscription.prototype.limit = null;

        /**
         * Subscription v3_state.
         * @member {number} v3_state
         * @memberof payment.Subscription
         * @instance
         */
        Subscription.prototype.v3_state = 0;

        return Subscription;
    })();

    payment.Bill = (function() {

        /**
         * Properties of a Bill.
         * @memberof payment
         * @interface IBill
         * @property {common.IContext|null} [ctx] Bill ctx
         * @property {string|null} [id] Bill id
         * @property {string|null} [account_id] Bill account_id
         * @property {number|null} [amount] Bill amount
         * @property {Array.<string>|null} [invoice_ids] Bill invoice_ids
         * @property {number|Long|null} [created] Bill created
         * @property {payment.IContact|null} [customer_info] Bill customer_info
         * @property {string|null} [payment_method] Bill payment_method
         * @property {number|null} [year] Bill year
         * @property {string|null} [description] Bill description
         */

        /**
         * Constructs a new Bill.
         * @memberof payment
         * @classdesc Represents a Bill.
         * @implements IBill
         * @constructor
         * @param {payment.IBill=} [p] Properties to set
         */
        function Bill(p) {
            this.invoice_ids = [];
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * Bill ctx.
         * @member {common.IContext|null|undefined} ctx
         * @memberof payment.Bill
         * @instance
         */
        Bill.prototype.ctx = null;

        /**
         * Bill id.
         * @member {string} id
         * @memberof payment.Bill
         * @instance
         */
        Bill.prototype.id = "";

        /**
         * Bill account_id.
         * @member {string} account_id
         * @memberof payment.Bill
         * @instance
         */
        Bill.prototype.account_id = "";

        /**
         * Bill amount.
         * @member {number} amount
         * @memberof payment.Bill
         * @instance
         */
        Bill.prototype.amount = 0;

        /**
         * Bill invoice_ids.
         * @member {Array.<string>} invoice_ids
         * @memberof payment.Bill
         * @instance
         */
        Bill.prototype.invoice_ids = $util.emptyArray;

        /**
         * Bill created.
         * @member {number|Long} created
         * @memberof payment.Bill
         * @instance
         */
        Bill.prototype.created = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

        /**
         * Bill customer_info.
         * @member {payment.IContact|null|undefined} customer_info
         * @memberof payment.Bill
         * @instance
         */
        Bill.prototype.customer_info = null;

        /**
         * Bill payment_method.
         * @member {string} payment_method
         * @memberof payment.Bill
         * @instance
         */
        Bill.prototype.payment_method = "";

        /**
         * Bill year.
         * @member {number} year
         * @memberof payment.Bill
         * @instance
         */
        Bill.prototype.year = 0;

        /**
         * Bill description.
         * @member {string} description
         * @memberof payment.Bill
         * @instance
         */
        Bill.prototype.description = "";

        return Bill;
    })();

    payment.Note = (function() {

        /**
         * Properties of a Note.
         * @memberof payment
         * @interface INote
         * @property {string|null} [message] Note message
         * @property {string|null} [creator] Note creator
         * @property {number|Long|null} [created] Note created
         */

        /**
         * Constructs a new Note.
         * @memberof payment
         * @classdesc Represents a Note.
         * @implements INote
         * @constructor
         * @param {payment.INote=} [p] Properties to set
         */
        function Note(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * Note message.
         * @member {string} message
         * @memberof payment.Note
         * @instance
         */
        Note.prototype.message = "";

        /**
         * Note creator.
         * @member {string} creator
         * @memberof payment.Note
         * @instance
         */
        Note.prototype.creator = "";

        /**
         * Note created.
         * @member {number|Long} created
         * @memberof payment.Note
         * @instance
         */
        Note.prototype.created = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

        return Note;
    })();

    payment.Invoices = (function() {

        /**
         * Properties of an Invoices.
         * @memberof payment
         * @interface IInvoices
         * @property {Array.<payment.IInvoice>|null} [invoices] Invoices invoices
         */

        /**
         * Constructs a new Invoices.
         * @memberof payment
         * @classdesc Represents an Invoices.
         * @implements IInvoices
         * @constructor
         * @param {payment.IInvoices=} [p] Properties to set
         */
        function Invoices(p) {
            this.invoices = [];
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * Invoices invoices.
         * @member {Array.<payment.IInvoice>} invoices
         * @memberof payment.Invoices
         * @instance
         */
        Invoices.prototype.invoices = $util.emptyArray;

        return Invoices;
    })();

    payment.Invoice = (function() {

        /**
         * Properties of an Invoice.
         * @memberof payment
         * @interface IInvoice
         * @property {common.IContext|null} [ctx] Invoice ctx
         * @property {string|null} [account_id] Invoice account_id
         * @property {string|null} [id] Invoice id
         * @property {number|null} [amount_due] Invoice amount_due
         * @property {string|null} [promotion_code] Invoice promotion_code
         * @property {string|null} [description] Invoice description
         * @property {payment.IBillingInfo|null} [billing_info] Invoice billing_info
         * @property {number|Long|null} [due_date] Invoice due_date
         * @property {string|null} [state] Invoice state
         * @property {number|Long|null} [created] Invoice created
         * @property {Array.<payment.IInvoiceItem>|null} [items] Invoice items
         * @property {number|null} [subtotal] Invoice subtotal
         * @property {number|null} [tax_percent] Invoice tax_percent
         * @property {number|null} [tax] Invoice tax
         * @property {number|null} [total] Invoice total
         * @property {number|Long|null} [updated] Invoice updated
         * @property {number|null} [year] Invoice year
         * @property {Array.<payment.INote>|null} [notes] Invoice notes
         * @property {Array.<string>|null} [bills] Invoice bills
         * @property {number|null} [payment_made] Invoice payment_made
         * @property {payment.ISubscription|null} [current_sub] Invoice current_sub
         * @property {payment.IPlan|null} [current_plan] Invoice current_plan
         */

        /**
         * Constructs a new Invoice.
         * @memberof payment
         * @classdesc Represents an Invoice.
         * @implements IInvoice
         * @constructor
         * @param {payment.IInvoice=} [p] Properties to set
         */
        function Invoice(p) {
            this.items = [];
            this.notes = [];
            this.bills = [];
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * Invoice ctx.
         * @member {common.IContext|null|undefined} ctx
         * @memberof payment.Invoice
         * @instance
         */
        Invoice.prototype.ctx = null;

        /**
         * Invoice account_id.
         * @member {string} account_id
         * @memberof payment.Invoice
         * @instance
         */
        Invoice.prototype.account_id = "";

        /**
         * Invoice id.
         * @member {string} id
         * @memberof payment.Invoice
         * @instance
         */
        Invoice.prototype.id = "";

        /**
         * Invoice amount_due.
         * @member {number} amount_due
         * @memberof payment.Invoice
         * @instance
         */
        Invoice.prototype.amount_due = 0;

        /**
         * Invoice promotion_code.
         * @member {string} promotion_code
         * @memberof payment.Invoice
         * @instance
         */
        Invoice.prototype.promotion_code = "";

        /**
         * Invoice description.
         * @member {string} description
         * @memberof payment.Invoice
         * @instance
         */
        Invoice.prototype.description = "";

        /**
         * Invoice billing_info.
         * @member {payment.IBillingInfo|null|undefined} billing_info
         * @memberof payment.Invoice
         * @instance
         */
        Invoice.prototype.billing_info = null;

        /**
         * Invoice due_date.
         * @member {number|Long} due_date
         * @memberof payment.Invoice
         * @instance
         */
        Invoice.prototype.due_date = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

        /**
         * Invoice state.
         * @member {string} state
         * @memberof payment.Invoice
         * @instance
         */
        Invoice.prototype.state = "";

        /**
         * Invoice created.
         * @member {number|Long} created
         * @memberof payment.Invoice
         * @instance
         */
        Invoice.prototype.created = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

        /**
         * Invoice items.
         * @member {Array.<payment.IInvoiceItem>} items
         * @memberof payment.Invoice
         * @instance
         */
        Invoice.prototype.items = $util.emptyArray;

        /**
         * Invoice subtotal.
         * @member {number} subtotal
         * @memberof payment.Invoice
         * @instance
         */
        Invoice.prototype.subtotal = 0;

        /**
         * Invoice tax_percent.
         * @member {number} tax_percent
         * @memberof payment.Invoice
         * @instance
         */
        Invoice.prototype.tax_percent = 0;

        /**
         * Invoice tax.
         * @member {number} tax
         * @memberof payment.Invoice
         * @instance
         */
        Invoice.prototype.tax = 0;

        /**
         * Invoice total.
         * @member {number} total
         * @memberof payment.Invoice
         * @instance
         */
        Invoice.prototype.total = 0;

        /**
         * Invoice updated.
         * @member {number|Long} updated
         * @memberof payment.Invoice
         * @instance
         */
        Invoice.prototype.updated = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

        /**
         * Invoice year.
         * @member {number} year
         * @memberof payment.Invoice
         * @instance
         */
        Invoice.prototype.year = 0;

        /**
         * Invoice notes.
         * @member {Array.<payment.INote>} notes
         * @memberof payment.Invoice
         * @instance
         */
        Invoice.prototype.notes = $util.emptyArray;

        /**
         * Invoice bills.
         * @member {Array.<string>} bills
         * @memberof payment.Invoice
         * @instance
         */
        Invoice.prototype.bills = $util.emptyArray;

        /**
         * Invoice payment_made.
         * @member {number} payment_made
         * @memberof payment.Invoice
         * @instance
         */
        Invoice.prototype.payment_made = 0;

        /**
         * Invoice current_sub.
         * @member {payment.ISubscription|null|undefined} current_sub
         * @memberof payment.Invoice
         * @instance
         */
        Invoice.prototype.current_sub = null;

        /**
         * Invoice current_plan.
         * @member {payment.IPlan|null|undefined} current_plan
         * @memberof payment.Invoice
         * @instance
         */
        Invoice.prototype.current_plan = null;

        /**
         * State enum.
         * @name payment.Invoice.State
         * @enum {string}
         * @property {number} draft=0 draft value
         * @property {number} open=1 open value
         * @property {number} overdue=2 overdue value
         * @property {number} paid=3 paid value
         * @property {number} voided=4 voided value
         * @property {number} queueing=5 queueing value
         */
        Invoice.State = (function() {
            const valuesById = {}, values = Object.create(valuesById);
            values[valuesById[0] = "draft"] = 0;
            values[valuesById[1] = "open"] = 1;
            values[valuesById[2] = "overdue"] = 2;
            values[valuesById[3] = "paid"] = 3;
            values[valuesById[4] = "voided"] = 4;
            values[valuesById[5] = "queueing"] = 5;
            return values;
        })();

        return Invoice;
    })();

    payment.AgentInvoiceItem = (function() {

        /**
         * Properties of an AgentInvoiceItem.
         * @memberof payment
         * @interface IAgentInvoiceItem
         * @property {string|null} [plan] AgentInvoiceItem plan
         * @property {number|null} [day_left] AgentInvoiceItem day_left
         * @property {number|null} [agent_count] AgentInvoiceItem agent_count
         */

        /**
         * Constructs a new AgentInvoiceItem.
         * @memberof payment
         * @classdesc Represents an AgentInvoiceItem.
         * @implements IAgentInvoiceItem
         * @constructor
         * @param {payment.IAgentInvoiceItem=} [p] Properties to set
         */
        function AgentInvoiceItem(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * AgentInvoiceItem plan.
         * @member {string} plan
         * @memberof payment.AgentInvoiceItem
         * @instance
         */
        AgentInvoiceItem.prototype.plan = "";

        /**
         * AgentInvoiceItem day_left.
         * @member {number} day_left
         * @memberof payment.AgentInvoiceItem
         * @instance
         */
        AgentInvoiceItem.prototype.day_left = 0;

        /**
         * AgentInvoiceItem agent_count.
         * @member {number} agent_count
         * @memberof payment.AgentInvoiceItem
         * @instance
         */
        AgentInvoiceItem.prototype.agent_count = 0;

        return AgentInvoiceItem;
    })();

    payment.RenewInvoiceItem = (function() {

        /**
         * Properties of a RenewInvoiceItem.
         * @memberof payment
         * @interface IRenewInvoiceItem
         * @property {string|null} [plan] RenewInvoiceItem plan
         * @property {number|null} [billing_cycle_month] RenewInvoiceItem billing_cycle_month
         * @property {number|null} [agent_count] RenewInvoiceItem agent_count
         * @property {number|Long|null} [from_time] RenewInvoiceItem from_time
         */

        /**
         * Constructs a new RenewInvoiceItem.
         * @memberof payment
         * @classdesc Represents a RenewInvoiceItem.
         * @implements IRenewInvoiceItem
         * @constructor
         * @param {payment.IRenewInvoiceItem=} [p] Properties to set
         */
        function RenewInvoiceItem(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * RenewInvoiceItem plan.
         * @member {string} plan
         * @memberof payment.RenewInvoiceItem
         * @instance
         */
        RenewInvoiceItem.prototype.plan = "";

        /**
         * RenewInvoiceItem billing_cycle_month.
         * @member {number} billing_cycle_month
         * @memberof payment.RenewInvoiceItem
         * @instance
         */
        RenewInvoiceItem.prototype.billing_cycle_month = 0;

        /**
         * RenewInvoiceItem agent_count.
         * @member {number} agent_count
         * @memberof payment.RenewInvoiceItem
         * @instance
         */
        RenewInvoiceItem.prototype.agent_count = 0;

        /**
         * RenewInvoiceItem from_time.
         * @member {number|Long} from_time
         * @memberof payment.RenewInvoiceItem
         * @instance
         */
        RenewInvoiceItem.prototype.from_time = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

        return RenewInvoiceItem;
    })();

    payment.PlanInvoiceItem = (function() {

        /**
         * Properties of a PlanInvoiceItem.
         * @memberof payment
         * @interface IPlanInvoiceItem
         * @property {number|null} [agent_count] PlanInvoiceItem agent_count
         * @property {number|null} [billing_cycle_month] PlanInvoiceItem billing_cycle_month
         * @property {string|null} [old_plan] PlanInvoiceItem old_plan
         * @property {string|null} [new_plan] PlanInvoiceItem new_plan
         * @property {number|null} [save_percentage] PlanInvoiceItem save_percentage
         * @property {number|Long|null} [started] PlanInvoiceItem started
         * @property {number|null} [day_left] PlanInvoiceItem day_left
         */

        /**
         * Constructs a new PlanInvoiceItem.
         * @memberof payment
         * @classdesc Represents a PlanInvoiceItem.
         * @implements IPlanInvoiceItem
         * @constructor
         * @param {payment.IPlanInvoiceItem=} [p] Properties to set
         */
        function PlanInvoiceItem(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * PlanInvoiceItem agent_count.
         * @member {number} agent_count
         * @memberof payment.PlanInvoiceItem
         * @instance
         */
        PlanInvoiceItem.prototype.agent_count = 0;

        /**
         * PlanInvoiceItem billing_cycle_month.
         * @member {number} billing_cycle_month
         * @memberof payment.PlanInvoiceItem
         * @instance
         */
        PlanInvoiceItem.prototype.billing_cycle_month = 0;

        /**
         * PlanInvoiceItem old_plan.
         * @member {string} old_plan
         * @memberof payment.PlanInvoiceItem
         * @instance
         */
        PlanInvoiceItem.prototype.old_plan = "";

        /**
         * PlanInvoiceItem new_plan.
         * @member {string} new_plan
         * @memberof payment.PlanInvoiceItem
         * @instance
         */
        PlanInvoiceItem.prototype.new_plan = "";

        /**
         * PlanInvoiceItem save_percentage.
         * @member {number} save_percentage
         * @memberof payment.PlanInvoiceItem
         * @instance
         */
        PlanInvoiceItem.prototype.save_percentage = 0;

        /**
         * PlanInvoiceItem started.
         * @member {number|Long} started
         * @memberof payment.PlanInvoiceItem
         * @instance
         */
        PlanInvoiceItem.prototype.started = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

        /**
         * PlanInvoiceItem day_left.
         * @member {number} day_left
         * @memberof payment.PlanInvoiceItem
         * @instance
         */
        PlanInvoiceItem.prototype.day_left = 0;

        return PlanInvoiceItem;
    })();

    payment.InvoiceItem = (function() {

        /**
         * Properties of an InvoiceItem.
         * @memberof payment
         * @interface IInvoiceItem
         * @property {string|null} [description] InvoiceItem description
         * @property {string|null} [invoice_id] InvoiceItem invoice_id
         * @property {number|null} [quantity] InvoiceItem quantity
         * @property {number|null} [price] InvoiceItem price
         * @property {payment.InvoiceItem.IData|null} [data] InvoiceItem data
         */

        /**
         * Constructs a new InvoiceItem.
         * @memberof payment
         * @classdesc Represents an InvoiceItem.
         * @implements IInvoiceItem
         * @constructor
         * @param {payment.IInvoiceItem=} [p] Properties to set
         */
        function InvoiceItem(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * InvoiceItem description.
         * @member {string} description
         * @memberof payment.InvoiceItem
         * @instance
         */
        InvoiceItem.prototype.description = "";

        /**
         * InvoiceItem invoice_id.
         * @member {string} invoice_id
         * @memberof payment.InvoiceItem
         * @instance
         */
        InvoiceItem.prototype.invoice_id = "";

        /**
         * InvoiceItem quantity.
         * @member {number} quantity
         * @memberof payment.InvoiceItem
         * @instance
         */
        InvoiceItem.prototype.quantity = 0;

        /**
         * InvoiceItem price.
         * @member {number} price
         * @memberof payment.InvoiceItem
         * @instance
         */
        InvoiceItem.prototype.price = 0;

        /**
         * InvoiceItem data.
         * @member {payment.InvoiceItem.IData|null|undefined} data
         * @memberof payment.InvoiceItem
         * @instance
         */
        InvoiceItem.prototype.data = null;

        InvoiceItem.Data = (function() {

            /**
             * Properties of a Data.
             * @memberof payment.InvoiceItem
             * @interface IData
             * @property {payment.IRenewInvoiceItem|null} [renew] Data renew
             * @property {payment.IAgentInvoiceItem|null} [agent] Data agent
             * @property {payment.IPlanInvoiceItem|null} [plan] Data plan
             */

            /**
             * Constructs a new Data.
             * @memberof payment.InvoiceItem
             * @classdesc Represents a Data.
             * @implements IData
             * @constructor
             * @param {payment.InvoiceItem.IData=} [p] Properties to set
             */
            function Data(p) {
                if (p)
                    for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                        if (p[ks[i]] != null)
                            this[ks[i]] = p[ks[i]];
            }

            /**
             * Data renew.
             * @member {payment.IRenewInvoiceItem|null|undefined} renew
             * @memberof payment.InvoiceItem.Data
             * @instance
             */
            Data.prototype.renew = null;

            /**
             * Data agent.
             * @member {payment.IAgentInvoiceItem|null|undefined} agent
             * @memberof payment.InvoiceItem.Data
             * @instance
             */
            Data.prototype.agent = null;

            /**
             * Data plan.
             * @member {payment.IPlanInvoiceItem|null|undefined} plan
             * @memberof payment.InvoiceItem.Data
             * @instance
             */
            Data.prototype.plan = null;

            return Data;
        })();

        return InvoiceItem;
    })();

    payment.BillingInfo = (function() {

        /**
         * Properties of a BillingInfo.
         * @memberof payment
         * @interface IBillingInfo
         * @property {string|null} [name] BillingInfo name
         * @property {string|null} [address] BillingInfo address
         * @property {string|null} [vat] BillingInfo vat
         * @property {string|null} [country_code] BillingInfo country_code
         */

        /**
         * Constructs a new BillingInfo.
         * @memberof payment
         * @classdesc Represents a BillingInfo.
         * @implements IBillingInfo
         * @constructor
         * @param {payment.IBillingInfo=} [p] Properties to set
         */
        function BillingInfo(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * BillingInfo name.
         * @member {string} name
         * @memberof payment.BillingInfo
         * @instance
         */
        BillingInfo.prototype.name = "";

        /**
         * BillingInfo address.
         * @member {string} address
         * @memberof payment.BillingInfo
         * @instance
         */
        BillingInfo.prototype.address = "";

        /**
         * BillingInfo vat.
         * @member {string} vat
         * @memberof payment.BillingInfo
         * @instance
         */
        BillingInfo.prototype.vat = "";

        /**
         * BillingInfo country_code.
         * @member {string} country_code
         * @memberof payment.BillingInfo
         * @instance
         */
        BillingInfo.prototype.country_code = "";

        return BillingInfo;
    })();

    payment.Contact = (function() {

        /**
         * Properties of a Contact.
         * @memberof payment
         * @interface IContact
         * @property {common.IContext|null} [ctx] Contact ctx
         * @property {string|null} [name] Contact name
         * @property {string|null} [email] Contact email
         * @property {string|null} [phone] Contact phone
         * @property {string|null} [job_title] Contact job_title
         * @property {string|null} [title] Contact title
         * @property {boolean|null} [primary] Contact primary
         */

        /**
         * Constructs a new Contact.
         * @memberof payment
         * @classdesc Represents a Contact.
         * @implements IContact
         * @constructor
         * @param {payment.IContact=} [p] Properties to set
         */
        function Contact(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * Contact ctx.
         * @member {common.IContext|null|undefined} ctx
         * @memberof payment.Contact
         * @instance
         */
        Contact.prototype.ctx = null;

        /**
         * Contact name.
         * @member {string} name
         * @memberof payment.Contact
         * @instance
         */
        Contact.prototype.name = "";

        /**
         * Contact email.
         * @member {string} email
         * @memberof payment.Contact
         * @instance
         */
        Contact.prototype.email = "";

        /**
         * Contact phone.
         * @member {string} phone
         * @memberof payment.Contact
         * @instance
         */
        Contact.prototype.phone = "";

        /**
         * Contact job_title.
         * @member {string} job_title
         * @memberof payment.Contact
         * @instance
         */
        Contact.prototype.job_title = "";

        /**
         * Contact title.
         * @member {string} title
         * @memberof payment.Contact
         * @instance
         */
        Contact.prototype.title = "";

        /**
         * Contact primary.
         * @member {boolean} primary
         * @memberof payment.Contact
         * @instance
         */
        Contact.prototype.primary = false;

        /**
         * Title enum.
         * @name payment.Contact.Title
         * @enum {string}
         * @property {number} mr=0 mr value
         * @property {number} ms=1 ms value
         * @property {number} mrs=2 mrs value
         * @property {number} dr=3 dr value
         */
        Contact.Title = (function() {
            const valuesById = {}, values = Object.create(valuesById);
            values[valuesById[0] = "mr"] = 0;
            values[valuesById[1] = "ms"] = 1;
            values[valuesById[2] = "mrs"] = 2;
            values[valuesById[3] = "dr"] = 3;
            return values;
        })();

        return Contact;
    })();

    payment.Customer = (function() {

        /**
         * Properties of a Customer.
         * @memberof payment
         * @interface ICustomer
         * @property {string|null} [id] Customer id
         * @property {string|null} [account_id] Customer account_id
         * @property {Array.<payment.IContact>|null} [contacts] Customer contacts
         * @property {number|Long|null} [created] Customer created
         * @property {payment.IBillingInfo|null} [billing_info] Customer billing_info
         */

        /**
         * Constructs a new Customer.
         * @memberof payment
         * @classdesc Represents a Customer.
         * @implements ICustomer
         * @constructor
         * @param {payment.ICustomer=} [p] Properties to set
         */
        function Customer(p) {
            this.contacts = [];
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * Customer id.
         * @member {string} id
         * @memberof payment.Customer
         * @instance
         */
        Customer.prototype.id = "";

        /**
         * Customer account_id.
         * @member {string} account_id
         * @memberof payment.Customer
         * @instance
         */
        Customer.prototype.account_id = "";

        /**
         * Customer contacts.
         * @member {Array.<payment.IContact>} contacts
         * @memberof payment.Customer
         * @instance
         */
        Customer.prototype.contacts = $util.emptyArray;

        /**
         * Customer created.
         * @member {number|Long} created
         * @memberof payment.Customer
         * @instance
         */
        Customer.prototype.created = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

        /**
         * Customer billing_info.
         * @member {payment.IBillingInfo|null|undefined} billing_info
         * @memberof payment.Customer
         * @instance
         */
        Customer.prototype.billing_info = null;

        return Customer;
    })();

    payment.FixedAmountPromotionCode = (function() {

        /**
         * Properties of a FixedAmountPromotionCode.
         * @memberof payment
         * @interface IFixedAmountPromotionCode
         * @property {number|null} [amount] FixedAmountPromotionCode amount
         */

        /**
         * Constructs a new FixedAmountPromotionCode.
         * @memberof payment
         * @classdesc Represents a FixedAmountPromotionCode.
         * @implements IFixedAmountPromotionCode
         * @constructor
         * @param {payment.IFixedAmountPromotionCode=} [p] Properties to set
         */
        function FixedAmountPromotionCode(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * FixedAmountPromotionCode amount.
         * @member {number} amount
         * @memberof payment.FixedAmountPromotionCode
         * @instance
         */
        FixedAmountPromotionCode.prototype.amount = 0;

        return FixedAmountPromotionCode;
    })();

    payment.PercentPromotionCode = (function() {

        /**
         * Properties of a PercentPromotionCode.
         * @memberof payment
         * @interface IPercentPromotionCode
         * @property {number|null} [percent] PercentPromotionCode percent
         */

        /**
         * Constructs a new PercentPromotionCode.
         * @memberof payment
         * @classdesc Represents a PercentPromotionCode.
         * @implements IPercentPromotionCode
         * @constructor
         * @param {payment.IPercentPromotionCode=} [p] Properties to set
         */
        function PercentPromotionCode(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * PercentPromotionCode percent.
         * @member {number} percent
         * @memberof payment.PercentPromotionCode
         * @instance
         */
        PercentPromotionCode.prototype.percent = 0;

        return PercentPromotionCode;
    })();

    payment.CreditCode = (function() {

        /**
         * Properties of a CreditCode.
         * @memberof payment
         * @interface ICreditCode
         * @property {number|null} [credit] CreditCode credit
         */

        /**
         * Constructs a new CreditCode.
         * @memberof payment
         * @classdesc Represents a CreditCode.
         * @implements ICreditCode
         * @constructor
         * @param {payment.ICreditCode=} [p] Properties to set
         */
        function CreditCode(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * CreditCode credit.
         * @member {number} credit
         * @memberof payment.CreditCode
         * @instance
         */
        CreditCode.prototype.credit = 0;

        return CreditCode;
    })();

    payment.ReferralCreditCode = (function() {

        /**
         * Properties of a ReferralCreditCode.
         * @memberof payment
         * @interface IReferralCreditCode
         * @property {string|null} [referrer_id] ReferralCreditCode referrer_id
         * @property {number|null} [credit] ReferralCreditCode credit
         */

        /**
         * Constructs a new ReferralCreditCode.
         * @memberof payment
         * @classdesc Represents a ReferralCreditCode.
         * @implements IReferralCreditCode
         * @constructor
         * @param {payment.IReferralCreditCode=} [p] Properties to set
         */
        function ReferralCreditCode(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * ReferralCreditCode referrer_id.
         * @member {string} referrer_id
         * @memberof payment.ReferralCreditCode
         * @instance
         */
        ReferralCreditCode.prototype.referrer_id = "";

        /**
         * ReferralCreditCode credit.
         * @member {number} credit
         * @memberof payment.ReferralCreditCode
         * @instance
         */
        ReferralCreditCode.prototype.credit = 0;

        return ReferralCreditCode;
    })();

    payment.PromotionCode = (function() {

        /**
         * Properties of a PromotionCode.
         * @memberof payment
         * @interface IPromotionCode
         * @property {common.IContext|null} [ctx] PromotionCode ctx
         * @property {string|null} [description] PromotionCode description
         * @property {string|null} [type] PromotionCode type
         * @property {number|null} [redeem_count] PromotionCode redeem_count
         * @property {string|null} [creator] PromotionCode creator
         * @property {number|Long|null} [created] PromotionCode created
         * @property {string|null} [code] PromotionCode code
         * @property {payment.PromotionCode.IData|null} [data] PromotionCode data
         * @property {number|Long|null} [start] PromotionCode start
         * @property {number|Long|null} [end] PromotionCode end
         * @property {string|null} [for_plan] PromotionCode for_plan
         * @property {string|null} [for_account_id] PromotionCode for_account_id
         * @property {number|null} [max_redemptions] PromotionCode max_redemptions
         * @property {string|null} [addon] PromotionCode addon
         * @property {Array.<string>|null} [for_items] PromotionCode for_items
         * @property {number|null} [min_amount] PromotionCode min_amount
         * @property {number|null} [max_amount] PromotionCode max_amount
         */

        /**
         * Constructs a new PromotionCode.
         * @memberof payment
         * @classdesc Represents a PromotionCode.
         * @implements IPromotionCode
         * @constructor
         * @param {payment.IPromotionCode=} [p] Properties to set
         */
        function PromotionCode(p) {
            this.for_items = [];
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * PromotionCode ctx.
         * @member {common.IContext|null|undefined} ctx
         * @memberof payment.PromotionCode
         * @instance
         */
        PromotionCode.prototype.ctx = null;

        /**
         * PromotionCode description.
         * @member {string} description
         * @memberof payment.PromotionCode
         * @instance
         */
        PromotionCode.prototype.description = "";

        /**
         * PromotionCode type.
         * @member {string} type
         * @memberof payment.PromotionCode
         * @instance
         */
        PromotionCode.prototype.type = "";

        /**
         * PromotionCode redeem_count.
         * @member {number} redeem_count
         * @memberof payment.PromotionCode
         * @instance
         */
        PromotionCode.prototype.redeem_count = 0;

        /**
         * PromotionCode creator.
         * @member {string} creator
         * @memberof payment.PromotionCode
         * @instance
         */
        PromotionCode.prototype.creator = "";

        /**
         * PromotionCode created.
         * @member {number|Long} created
         * @memberof payment.PromotionCode
         * @instance
         */
        PromotionCode.prototype.created = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

        /**
         * PromotionCode code.
         * @member {string} code
         * @memberof payment.PromotionCode
         * @instance
         */
        PromotionCode.prototype.code = "";

        /**
         * PromotionCode data.
         * @member {payment.PromotionCode.IData|null|undefined} data
         * @memberof payment.PromotionCode
         * @instance
         */
        PromotionCode.prototype.data = null;

        /**
         * PromotionCode start.
         * @member {number|Long} start
         * @memberof payment.PromotionCode
         * @instance
         */
        PromotionCode.prototype.start = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

        /**
         * PromotionCode end.
         * @member {number|Long} end
         * @memberof payment.PromotionCode
         * @instance
         */
        PromotionCode.prototype.end = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

        /**
         * PromotionCode for_plan.
         * @member {string} for_plan
         * @memberof payment.PromotionCode
         * @instance
         */
        PromotionCode.prototype.for_plan = "";

        /**
         * PromotionCode for_account_id.
         * @member {string} for_account_id
         * @memberof payment.PromotionCode
         * @instance
         */
        PromotionCode.prototype.for_account_id = "";

        /**
         * PromotionCode max_redemptions.
         * @member {number} max_redemptions
         * @memberof payment.PromotionCode
         * @instance
         */
        PromotionCode.prototype.max_redemptions = 0;

        /**
         * PromotionCode addon.
         * @member {string} addon
         * @memberof payment.PromotionCode
         * @instance
         */
        PromotionCode.prototype.addon = "";

        /**
         * PromotionCode for_items.
         * @member {Array.<string>} for_items
         * @memberof payment.PromotionCode
         * @instance
         */
        PromotionCode.prototype.for_items = $util.emptyArray;

        /**
         * PromotionCode min_amount.
         * @member {number} min_amount
         * @memberof payment.PromotionCode
         * @instance
         */
        PromotionCode.prototype.min_amount = 0;

        /**
         * PromotionCode max_amount.
         * @member {number} max_amount
         * @memberof payment.PromotionCode
         * @instance
         */
        PromotionCode.prototype.max_amount = 0;

        /**
         * Type enum.
         * @name payment.PromotionCode.Type
         * @enum {string}
         * @property {number} fixed_amount_promotion_code=0 fixed_amount_promotion_code value
         * @property {number} percent_promotion_code=1 percent_promotion_code value
         * @property {number} credit_code=2 credit_code value
         * @property {number} referral_credit_code=3 referral_credit_code value
         */
        PromotionCode.Type = (function() {
            const valuesById = {}, values = Object.create(valuesById);
            values[valuesById[0] = "fixed_amount_promotion_code"] = 0;
            values[valuesById[1] = "percent_promotion_code"] = 1;
            values[valuesById[2] = "credit_code"] = 2;
            values[valuesById[3] = "referral_credit_code"] = 3;
            return values;
        })();

        PromotionCode.Data = (function() {

            /**
             * Properties of a Data.
             * @memberof payment.PromotionCode
             * @interface IData
             * @property {payment.IFixedAmountPromotionCode|null} [fixed_amount] Data fixed_amount
             * @property {payment.IPercentPromotionCode|null} [percent] Data percent
             * @property {payment.ICreditCode|null} [credit] Data credit
             * @property {payment.IReferralCreditCode|null} [referral] Data referral
             */

            /**
             * Constructs a new Data.
             * @memberof payment.PromotionCode
             * @classdesc Represents a Data.
             * @implements IData
             * @constructor
             * @param {payment.PromotionCode.IData=} [p] Properties to set
             */
            function Data(p) {
                if (p)
                    for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                        if (p[ks[i]] != null)
                            this[ks[i]] = p[ks[i]];
            }

            /**
             * Data fixed_amount.
             * @member {payment.IFixedAmountPromotionCode|null|undefined} fixed_amount
             * @memberof payment.PromotionCode.Data
             * @instance
             */
            Data.prototype.fixed_amount = null;

            /**
             * Data percent.
             * @member {payment.IPercentPromotionCode|null|undefined} percent
             * @memberof payment.PromotionCode.Data
             * @instance
             */
            Data.prototype.percent = null;

            /**
             * Data credit.
             * @member {payment.ICreditCode|null|undefined} credit
             * @memberof payment.PromotionCode.Data
             * @instance
             */
            Data.prototype.credit = null;

            /**
             * Data referral.
             * @member {payment.IReferralCreditCode|null|undefined} referral
             * @memberof payment.PromotionCode.Data
             * @instance
             */
            Data.prototype.referral = null;

            return Data;
        })();

        return PromotionCode;
    })();

    payment.ExchangeRate = (function() {

        /**
         * Properties of an ExchangeRate.
         * @memberof payment
         * @interface IExchangeRate
         * @property {string|null} [from_currency] ExchangeRate from_currency
         * @property {string|null} [to_currency] ExchangeRate to_currency
         * @property {number|null} [exchange_rate] ExchangeRate exchange_rate
         * @property {number|Long|null} [exchange_time] ExchangeRate exchange_time
         */

        /**
         * Constructs a new ExchangeRate.
         * @memberof payment
         * @classdesc Represents an ExchangeRate.
         * @implements IExchangeRate
         * @constructor
         * @param {payment.IExchangeRate=} [p] Properties to set
         */
        function ExchangeRate(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * ExchangeRate from_currency.
         * @member {string} from_currency
         * @memberof payment.ExchangeRate
         * @instance
         */
        ExchangeRate.prototype.from_currency = "";

        /**
         * ExchangeRate to_currency.
         * @member {string} to_currency
         * @memberof payment.ExchangeRate
         * @instance
         */
        ExchangeRate.prototype.to_currency = "";

        /**
         * ExchangeRate exchange_rate.
         * @member {number} exchange_rate
         * @memberof payment.ExchangeRate
         * @instance
         */
        ExchangeRate.prototype.exchange_rate = 0;

        /**
         * ExchangeRate exchange_time.
         * @member {number|Long} exchange_time
         * @memberof payment.ExchangeRate
         * @instance
         */
        ExchangeRate.prototype.exchange_time = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

        return ExchangeRate;
    })();

    payment.Log = (function() {

        /**
         * Properties of a Log.
         * @memberof payment
         * @interface ILog
         * @property {common.IContext|null} [ctx] Log ctx
         * @property {string|null} [user] Log user
         * @property {string|null} [id] Log id
         * @property {string|null} [action] Log action
         * @property {number|Long|null} [created] Log created
         * @property {string|null} [description] Log description
         * @property {string|null} [account_id] Log account_id
         * @property {number|null} [month] Log month
         */

        /**
         * Constructs a new Log.
         * @memberof payment
         * @classdesc Represents a Log.
         * @implements ILog
         * @constructor
         * @param {payment.ILog=} [p] Properties to set
         */
        function Log(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * Log ctx.
         * @member {common.IContext|null|undefined} ctx
         * @memberof payment.Log
         * @instance
         */
        Log.prototype.ctx = null;

        /**
         * Log user.
         * @member {string} user
         * @memberof payment.Log
         * @instance
         */
        Log.prototype.user = "";

        /**
         * Log id.
         * @member {string} id
         * @memberof payment.Log
         * @instance
         */
        Log.prototype.id = "";

        /**
         * Log action.
         * @member {string} action
         * @memberof payment.Log
         * @instance
         */
        Log.prototype.action = "";

        /**
         * Log created.
         * @member {number|Long} created
         * @memberof payment.Log
         * @instance
         */
        Log.prototype.created = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

        /**
         * Log description.
         * @member {string} description
         * @memberof payment.Log
         * @instance
         */
        Log.prototype.description = "";

        /**
         * Log account_id.
         * @member {string} account_id
         * @memberof payment.Log
         * @instance
         */
        Log.prototype.account_id = "";

        /**
         * Log month.
         * @member {number} month
         * @memberof payment.Log
         * @instance
         */
        Log.prototype.month = 0;

        /**
         * Action enum.
         * @name payment.Log.Action
         * @enum {string}
         * @property {number} create_invoice=0 create_invoice value
         * @property {number} change_invoice_status=1 change_invoice_status value
         * @property {number} create_discount=2 create_discount value
         * @property {number} delete_discount=3 delete_discount value
         * @property {number} redeem_discount=4 redeem_discount value
         * @property {number} add_credit=5 add_credit value
         * @property {number} redeem_credit=6 redeem_credit value
         * @property {number} delete_account=7 delete_account value
         * @property {number} change_plan=8 change_plan value
         * @property {number} renew_subscription=10 renew_subscription value
         * @property {number} click_subscribe_button=11 click_subscribe_button value
         * @property {number} pay_for_referrer=12 pay_for_referrer value
         * @property {number} add_money_for_referrer=13 add_money_for_referrer value
         * @property {number} pay_invoice=14 pay_invoice value
         */
        Log.Action = (function() {
            const valuesById = {}, values = Object.create(valuesById);
            values[valuesById[0] = "create_invoice"] = 0;
            values[valuesById[1] = "change_invoice_status"] = 1;
            values[valuesById[2] = "create_discount"] = 2;
            values[valuesById[3] = "delete_discount"] = 3;
            values[valuesById[4] = "redeem_discount"] = 4;
            values[valuesById[5] = "add_credit"] = 5;
            values[valuesById[6] = "redeem_credit"] = 6;
            values[valuesById[7] = "delete_account"] = 7;
            values[valuesById[8] = "change_plan"] = 8;
            values[valuesById[10] = "renew_subscription"] = 10;
            values[valuesById[11] = "click_subscribe_button"] = 11;
            values[valuesById[12] = "pay_for_referrer"] = 12;
            values[valuesById[13] = "add_money_for_referrer"] = 13;
            values[valuesById[14] = "pay_invoice"] = 14;
            return values;
        })();

        return Log;
    })();

    payment.Addon = (function() {

        /**
         * Properties of an Addon.
         * @memberof payment
         * @interface IAddon
         * @property {string|null} [type] Addon type
         * @property {string|null} [name] Addon name
         * @property {number|null} [price] Addon price
         * @property {payment.Currency|null} [currency] Addon currency
         * @property {string|null} [charge_type] Addon charge_type
         * @property {number|null} [period] Addon period
         * @property {string|null} [period_unit] Addon period_unit
         * @property {number|null} [quantity] Addon quantity
         * @property {number|Long|null} [created] Addon created
         */

        /**
         * Constructs a new Addon.
         * @memberof payment
         * @classdesc Represents an Addon.
         * @implements IAddon
         * @constructor
         * @param {payment.IAddon=} [p] Properties to set
         */
        function Addon(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * Addon type.
         * @member {string} type
         * @memberof payment.Addon
         * @instance
         */
        Addon.prototype.type = "";

        /**
         * Addon name.
         * @member {string} name
         * @memberof payment.Addon
         * @instance
         */
        Addon.prototype.name = "";

        /**
         * Addon price.
         * @member {number} price
         * @memberof payment.Addon
         * @instance
         */
        Addon.prototype.price = 0;

        /**
         * Addon currency.
         * @member {payment.Currency} currency
         * @memberof payment.Addon
         * @instance
         */
        Addon.prototype.currency = 0;

        /**
         * Addon charge_type.
         * @member {string} charge_type
         * @memberof payment.Addon
         * @instance
         */
        Addon.prototype.charge_type = "";

        /**
         * Addon period.
         * @member {number} period
         * @memberof payment.Addon
         * @instance
         */
        Addon.prototype.period = 0;

        /**
         * Addon period_unit.
         * @member {string} period_unit
         * @memberof payment.Addon
         * @instance
         */
        Addon.prototype.period_unit = "";

        /**
         * Addon quantity.
         * @member {number} quantity
         * @memberof payment.Addon
         * @instance
         */
        Addon.prototype.quantity = 0;

        /**
         * Addon created.
         * @member {number|Long} created
         * @memberof payment.Addon
         * @instance
         */
        Addon.prototype.created = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

        /**
         * Type enum.
         * @name payment.Addon.Type
         * @enum {string}
         * @property {number} credit=0 credit value
         * @property {number} agent=1 agent value
         */
        Addon.Type = (function() {
            const valuesById = {}, values = Object.create(valuesById);
            values[valuesById[0] = "credit"] = 0;
            values[valuesById[1] = "agent"] = 1;
            return values;
        })();

        /**
         * ChargeType enum.
         * @name payment.Addon.ChargeType
         * @enum {string}
         * @property {number} one_time=0 one_time value
         * @property {number} recurring=1 recurring value
         */
        Addon.ChargeType = (function() {
            const valuesById = {}, values = Object.create(valuesById);
            values[valuesById[0] = "one_time"] = 0;
            values[valuesById[1] = "recurring"] = 1;
            return values;
        })();

        /**
         * PeriodUnit enum.
         * @name payment.Addon.PeriodUnit
         * @enum {string}
         * @property {number} day=0 day value
         * @property {number} week=1 week value
         * @property {number} month=2 month value
         * @property {number} year=3 year value
         */
        Addon.PeriodUnit = (function() {
            const valuesById = {}, values = Object.create(valuesById);
            values[valuesById[0] = "day"] = 0;
            values[valuesById[1] = "week"] = 1;
            values[valuesById[2] = "month"] = 2;
            values[valuesById[3] = "year"] = 3;
            return values;
        })();

        return Addon;
    })();

    /**
     * Event enum.
     * @name payment.Event
     * @enum {string}
     * @property {number} PaymentSynced=0 PaymentSynced value
     * @property {number} LimitUpdated=1 LimitUpdated value
     * @property {number} PaymentRequested=4 PaymentRequested value
     * @property {number} PaymentRenewCycle=5 PaymentRenewCycle value
     * @property {number} InvoiceCreatedEmailSend=6 InvoiceCreatedEmailSend value
     * @property {number} PaymentV3Synced=8 PaymentV3Synced value
     * @property {number} SubscriptionUpdated=14 SubscriptionUpdated value
     * @property {number} InvoiceUpdated=10 InvoiceUpdated value
     * @property {number} PaymentMethodUpdated=11 PaymentMethodUpdated value
     * @property {number} BillingUpdated=12 BillingUpdated value
     * @property {number} LogUpdated=13 LogUpdated value
     * @property {number} PromotionCodeUpdated=15 PromotionCodeUpdated value
     */
    payment.Event = (function() {
        const valuesById = {}, values = Object.create(valuesById);
        values[valuesById[0] = "PaymentSynced"] = 0;
        values[valuesById[1] = "LimitUpdated"] = 1;
        values[valuesById[4] = "PaymentRequested"] = 4;
        values[valuesById[5] = "PaymentRenewCycle"] = 5;
        values[valuesById[6] = "InvoiceCreatedEmailSend"] = 6;
        values[valuesById[8] = "PaymentV3Synced"] = 8;
        values[valuesById[14] = "SubscriptionUpdated"] = 14;
        values[valuesById[10] = "InvoiceUpdated"] = 10;
        values[valuesById[11] = "PaymentMethodUpdated"] = 11;
        values[valuesById[12] = "BillingUpdated"] = 12;
        values[valuesById[13] = "LogUpdated"] = 13;
        values[valuesById[15] = "PromotionCodeUpdated"] = 15;
        return values;
    })();

    payment.PaymentRenewCycleRequested = (function() {

        /**
         * Properties of a PaymentRenewCycleRequested.
         * @memberof payment
         * @interface IPaymentRenewCycleRequested
         * @property {common.IContext|null} [ctx] PaymentRenewCycleRequested ctx
         * @property {string|null} [account_id] PaymentRenewCycleRequested account_id
         * @property {payment.ISubscription|null} [sub] PaymentRenewCycleRequested sub
         * @property {string|null} [cycle_id] PaymentRenewCycleRequested cycle_id
         */

        /**
         * Constructs a new PaymentRenewCycleRequested.
         * @memberof payment
         * @classdesc Represents a PaymentRenewCycleRequested.
         * @implements IPaymentRenewCycleRequested
         * @constructor
         * @param {payment.IPaymentRenewCycleRequested=} [p] Properties to set
         */
        function PaymentRenewCycleRequested(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * PaymentRenewCycleRequested ctx.
         * @member {common.IContext|null|undefined} ctx
         * @memberof payment.PaymentRenewCycleRequested
         * @instance
         */
        PaymentRenewCycleRequested.prototype.ctx = null;

        /**
         * PaymentRenewCycleRequested account_id.
         * @member {string} account_id
         * @memberof payment.PaymentRenewCycleRequested
         * @instance
         */
        PaymentRenewCycleRequested.prototype.account_id = "";

        /**
         * PaymentRenewCycleRequested sub.
         * @member {payment.ISubscription|null|undefined} sub
         * @memberof payment.PaymentRenewCycleRequested
         * @instance
         */
        PaymentRenewCycleRequested.prototype.sub = null;

        /**
         * PaymentRenewCycleRequested cycle_id.
         * @member {string} cycle_id
         * @memberof payment.PaymentRenewCycleRequested
         * @instance
         */
        PaymentRenewCycleRequested.prototype.cycle_id = "";

        return PaymentRenewCycleRequested;
    })();

    payment.SubizInternalPaymentMgr = (function() {

        /**
         * Constructs a new SubizInternalPaymentMgr service.
         * @memberof payment
         * @classdesc Represents a SubizInternalPaymentMgr
         * @extends $protobuf.rpc.Service
         * @constructor
         * @param {$protobuf.RPCImpl} rpcImpl RPC implementation
         * @param {boolean} [requestDelimited=false] Whether requests are length-delimited
         * @param {boolean} [responseDelimited=false] Whether responses are length-delimited
         */
        function SubizInternalPaymentMgr(rpcImpl, requestDelimited, responseDelimited) {
            $protobuf.rpc.Service.call(this, rpcImpl, requestDelimited, responseDelimited);
        }

        (SubizInternalPaymentMgr.prototype = Object.create($protobuf.rpc.Service.prototype)).constructor = SubizInternalPaymentMgr;

        /**
         * Callback as used by {@link payment.SubizInternalPaymentMgr#createBill}.
         * @memberof payment.SubizInternalPaymentMgr
         * @typedef CreateBillCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {payment.Bill} [response] Bill
         */

        /**
         * Calls CreateBill.
         * @function createBill
         * @memberof payment.SubizInternalPaymentMgr
         * @instance
         * @param {payment.IBill} request Bill message or plain object
         * @param {payment.SubizInternalPaymentMgr.CreateBillCallback} callback Node-style callback called with the error, if any, and Bill
         * @returns {undefined}
         * @variation 1
         */
        SubizInternalPaymentMgr.prototype.createBill = function createBill(request, callback) {
            return this.rpcCall(createBill, $root.payment.Bill, $root.payment.Bill, request, callback);
        };

        /**
         * Calls CreateBill.
         * @function createBill
         * @memberof payment.SubizInternalPaymentMgr
         * @instance
         * @param {payment.IBill} request Bill message or plain object
         * @returns {Promise<payment.Bill>} Promise
         * @variation 2
         */

        /**
         * Callback as used by {@link payment.SubizInternalPaymentMgr#updateExchangeRate}.
         * @memberof payment.SubizInternalPaymentMgr
         * @typedef UpdateExchangeRateCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {payment.ExchangeRate} [response] ExchangeRate
         */

        /**
         * Calls UpdateExchangeRate.
         * @function updateExchangeRate
         * @memberof payment.SubizInternalPaymentMgr
         * @instance
         * @param {payment.IExchangeRate} request ExchangeRate message or plain object
         * @param {payment.SubizInternalPaymentMgr.UpdateExchangeRateCallback} callback Node-style callback called with the error, if any, and ExchangeRate
         * @returns {undefined}
         * @variation 1
         */
        SubizInternalPaymentMgr.prototype.updateExchangeRate = function updateExchangeRate(request, callback) {
            return this.rpcCall(updateExchangeRate, $root.payment.ExchangeRate, $root.payment.ExchangeRate, request, callback);
        };

        /**
         * Calls UpdateExchangeRate.
         * @function updateExchangeRate
         * @memberof payment.SubizInternalPaymentMgr
         * @instance
         * @param {payment.IExchangeRate} request ExchangeRate message or plain object
         * @returns {Promise<payment.ExchangeRate>} Promise
         * @variation 2
         */

        return SubizInternalPaymentMgr;
    })();

    payment.PaymentMgr = (function() {

        /**
         * Constructs a new PaymentMgr service.
         * @memberof payment
         * @classdesc Represents a PaymentMgr
         * @extends $protobuf.rpc.Service
         * @constructor
         * @param {$protobuf.RPCImpl} rpcImpl RPC implementation
         * @param {boolean} [requestDelimited=false] Whether requests are length-delimited
         * @param {boolean} [responseDelimited=false] Whether responses are length-delimited
         */
        function PaymentMgr(rpcImpl, requestDelimited, responseDelimited) {
            $protobuf.rpc.Service.call(this, rpcImpl, requestDelimited, responseDelimited);
        }

        (PaymentMgr.prototype = Object.create($protobuf.rpc.Service.prototype)).constructor = PaymentMgr;

        /**
         * Callback as used by {@link payment.PaymentMgr#purchase}.
         * @memberof payment.PaymentMgr
         * @typedef PurchaseCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {payment.Subscription} [response] Subscription
         */

        /**
         * Calls Purchase.
         * @function purchase
         * @memberof payment.PaymentMgr
         * @instance
         * @param {payment.ISubscription} request Subscription message or plain object
         * @param {payment.PaymentMgr.PurchaseCallback} callback Node-style callback called with the error, if any, and Subscription
         * @returns {undefined}
         * @variation 1
         */
        PaymentMgr.prototype.purchase = function purchase(request, callback) {
            return this.rpcCall(purchase, $root.payment.Subscription, $root.payment.Subscription, request, callback);
        };

        /**
         * Calls Purchase.
         * @function purchase
         * @memberof payment.PaymentMgr
         * @instance
         * @param {payment.ISubscription} request Subscription message or plain object
         * @returns {Promise<payment.Subscription>} Promise
         * @variation 2
         */

        /**
         * Callback as used by {@link payment.PaymentMgr#updateSubscription}.
         * @memberof payment.PaymentMgr
         * @typedef UpdateSubscriptionCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {payment.Subscription} [response] Subscription
         */

        /**
         * Calls UpdateSubscription.
         * @function updateSubscription
         * @memberof payment.PaymentMgr
         * @instance
         * @param {payment.ISubscription} request Subscription message or plain object
         * @param {payment.PaymentMgr.UpdateSubscriptionCallback} callback Node-style callback called with the error, if any, and Subscription
         * @returns {undefined}
         * @variation 1
         */
        PaymentMgr.prototype.updateSubscription = function updateSubscription(request, callback) {
            return this.rpcCall(updateSubscription, $root.payment.Subscription, $root.payment.Subscription, request, callback);
        };

        /**
         * Calls UpdateSubscription.
         * @function updateSubscription
         * @memberof payment.PaymentMgr
         * @instance
         * @param {payment.ISubscription} request Subscription message or plain object
         * @returns {Promise<payment.Subscription>} Promise
         * @variation 2
         */

        /**
         * Callback as used by {@link payment.PaymentMgr#getSubscription}.
         * @memberof payment.PaymentMgr
         * @typedef GetSubscriptionCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {payment.Subscription} [response] Subscription
         */

        /**
         * Calls GetSubscription.
         * @function getSubscription
         * @memberof payment.PaymentMgr
         * @instance
         * @param {common.IEmpty} request Empty message or plain object
         * @param {payment.PaymentMgr.GetSubscriptionCallback} callback Node-style callback called with the error, if any, and Subscription
         * @returns {undefined}
         * @variation 1
         */
        PaymentMgr.prototype.getSubscription = function getSubscription(request, callback) {
            return this.rpcCall(getSubscription, $root.common.Empty, $root.payment.Subscription, request, callback);
        };

        /**
         * Calls GetSubscription.
         * @function getSubscription
         * @memberof payment.PaymentMgr
         * @instance
         * @param {common.IEmpty} request Empty message or plain object
         * @returns {Promise<payment.Subscription>} Promise
         * @variation 2
         */

        /**
         * Callback as used by {@link payment.PaymentMgr#getPromotionCode}.
         * @memberof payment.PaymentMgr
         * @typedef GetPromotionCodeCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {payment.PromotionCode} [response] PromotionCode
         */

        /**
         * Calls GetPromotionCode.
         * @function getPromotionCode
         * @memberof payment.PaymentMgr
         * @instance
         * @param {payment.IString} request String message or plain object
         * @param {payment.PaymentMgr.GetPromotionCodeCallback} callback Node-style callback called with the error, if any, and PromotionCode
         * @returns {undefined}
         * @variation 1
         */
        PaymentMgr.prototype.getPromotionCode = function getPromotionCode(request, callback) {
            return this.rpcCall(getPromotionCode, $root.payment.String, $root.payment.PromotionCode, request, callback);
        };

        /**
         * Calls GetPromotionCode.
         * @function getPromotionCode
         * @memberof payment.PaymentMgr
         * @instance
         * @param {payment.IString} request String message or plain object
         * @returns {Promise<payment.PromotionCode>} Promise
         * @variation 2
         */

        /**
         * Callback as used by {@link payment.PaymentMgr#addPaymentMethod}.
         * @memberof payment.PaymentMgr
         * @typedef AddPaymentMethodCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {payment.PaymentMethod} [response] PaymentMethod
         */

        /**
         * Calls AddPaymentMethod.
         * @function addPaymentMethod
         * @memberof payment.PaymentMgr
         * @instance
         * @param {payment.IPaymentMethod} request PaymentMethod message or plain object
         * @param {payment.PaymentMgr.AddPaymentMethodCallback} callback Node-style callback called with the error, if any, and PaymentMethod
         * @returns {undefined}
         * @variation 1
         */
        PaymentMgr.prototype.addPaymentMethod = function addPaymentMethod(request, callback) {
            return this.rpcCall(addPaymentMethod, $root.payment.PaymentMethod, $root.payment.PaymentMethod, request, callback);
        };

        /**
         * Calls AddPaymentMethod.
         * @function addPaymentMethod
         * @memberof payment.PaymentMgr
         * @instance
         * @param {payment.IPaymentMethod} request PaymentMethod message or plain object
         * @returns {Promise<payment.PaymentMethod>} Promise
         * @variation 2
         */

        /**
         * Callback as used by {@link payment.PaymentMgr#updatePaymentMethod}.
         * @memberof payment.PaymentMgr
         * @typedef UpdatePaymentMethodCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {payment.PaymentMethod} [response] PaymentMethod
         */

        /**
         * Calls UpdatePaymentMethod.
         * @function updatePaymentMethod
         * @memberof payment.PaymentMgr
         * @instance
         * @param {payment.IPaymentMethod} request PaymentMethod message or plain object
         * @param {payment.PaymentMgr.UpdatePaymentMethodCallback} callback Node-style callback called with the error, if any, and PaymentMethod
         * @returns {undefined}
         * @variation 1
         */
        PaymentMgr.prototype.updatePaymentMethod = function updatePaymentMethod(request, callback) {
            return this.rpcCall(updatePaymentMethod, $root.payment.PaymentMethod, $root.payment.PaymentMethod, request, callback);
        };

        /**
         * Calls UpdatePaymentMethod.
         * @function updatePaymentMethod
         * @memberof payment.PaymentMgr
         * @instance
         * @param {payment.IPaymentMethod} request PaymentMethod message or plain object
         * @returns {Promise<payment.PaymentMethod>} Promise
         * @variation 2
         */

        /**
         * Callback as used by {@link payment.PaymentMgr#deletePaymentMethod}.
         * @memberof payment.PaymentMgr
         * @typedef DeletePaymentMethodCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {common.Empty} [response] Empty
         */

        /**
         * Calls DeletePaymentMethod.
         * @function deletePaymentMethod
         * @memberof payment.PaymentMgr
         * @instance
         * @param {common.IId} request Id message or plain object
         * @param {payment.PaymentMgr.DeletePaymentMethodCallback} callback Node-style callback called with the error, if any, and Empty
         * @returns {undefined}
         * @variation 1
         */
        PaymentMgr.prototype.deletePaymentMethod = function deletePaymentMethod(request, callback) {
            return this.rpcCall(deletePaymentMethod, $root.common.Id, $root.common.Empty, request, callback);
        };

        /**
         * Calls DeletePaymentMethod.
         * @function deletePaymentMethod
         * @memberof payment.PaymentMgr
         * @instance
         * @param {common.IId} request Id message or plain object
         * @returns {Promise<common.Empty>} Promise
         * @variation 2
         */

        /**
         * Callback as used by {@link payment.PaymentMgr#listPaymentMethods}.
         * @memberof payment.PaymentMgr
         * @typedef ListPaymentMethodsCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {payment.PaymentMethods} [response] PaymentMethods
         */

        /**
         * Calls ListPaymentMethods.
         * @function listPaymentMethods
         * @memberof payment.PaymentMgr
         * @instance
         * @param {common.IEmpty} request Empty message or plain object
         * @param {payment.PaymentMgr.ListPaymentMethodsCallback} callback Node-style callback called with the error, if any, and PaymentMethods
         * @returns {undefined}
         * @variation 1
         */
        PaymentMgr.prototype.listPaymentMethods = function listPaymentMethods(request, callback) {
            return this.rpcCall(listPaymentMethods, $root.common.Empty, $root.payment.PaymentMethods, request, callback);
        };

        /**
         * Calls ListPaymentMethods.
         * @function listPaymentMethods
         * @memberof payment.PaymentMgr
         * @instance
         * @param {common.IEmpty} request Empty message or plain object
         * @returns {Promise<payment.PaymentMethods>} Promise
         * @variation 2
         */

        /**
         * Callback as used by {@link payment.PaymentMgr#pay}.
         * @memberof payment.PaymentMgr
         * @typedef PayCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {payment.Bill} [response] Bill
         */

        /**
         * Calls Pay.
         * @function pay
         * @memberof payment.PaymentMgr
         * @instance
         * @param {payment.IPayRequest} request PayRequest message or plain object
         * @param {payment.PaymentMgr.PayCallback} callback Node-style callback called with the error, if any, and Bill
         * @returns {undefined}
         * @variation 1
         */
        PaymentMgr.prototype.pay = function pay(request, callback) {
            return this.rpcCall(pay, $root.payment.PayRequest, $root.payment.Bill, request, callback);
        };

        /**
         * Calls Pay.
         * @function pay
         * @memberof payment.PaymentMgr
         * @instance
         * @param {payment.IPayRequest} request PayRequest message or plain object
         * @returns {Promise<payment.Bill>} Promise
         * @variation 2
         */

        /**
         * Callback as used by {@link payment.PaymentMgr#listInvoices}.
         * @memberof payment.PaymentMgr
         * @typedef ListInvoicesCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {payment.Invoices} [response] Invoices
         */

        /**
         * Calls ListInvoices.
         * @function listInvoices
         * @memberof payment.PaymentMgr
         * @instance
         * @param {common.IId} request Id message or plain object
         * @param {payment.PaymentMgr.ListInvoicesCallback} callback Node-style callback called with the error, if any, and Invoices
         * @returns {undefined}
         * @variation 1
         */
        PaymentMgr.prototype.listInvoices = function listInvoices(request, callback) {
            return this.rpcCall(listInvoices, $root.common.Id, $root.payment.Invoices, request, callback);
        };

        /**
         * Calls ListInvoices.
         * @function listInvoices
         * @memberof payment.PaymentMgr
         * @instance
         * @param {common.IId} request Id message or plain object
         * @returns {Promise<payment.Invoices>} Promise
         * @variation 2
         */

        /**
         * Callback as used by {@link payment.PaymentMgr#createInvoice}.
         * @memberof payment.PaymentMgr
         * @typedef CreateInvoiceCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {payment.Invoice} [response] Invoice
         */

        /**
         * Calls CreateInvoice.
         * @function createInvoice
         * @memberof payment.PaymentMgr
         * @instance
         * @param {payment.IInvoice} request Invoice message or plain object
         * @param {payment.PaymentMgr.CreateInvoiceCallback} callback Node-style callback called with the error, if any, and Invoice
         * @returns {undefined}
         * @variation 1
         */
        PaymentMgr.prototype.createInvoice = function createInvoice(request, callback) {
            return this.rpcCall(createInvoice, $root.payment.Invoice, $root.payment.Invoice, request, callback);
        };

        /**
         * Calls CreateInvoice.
         * @function createInvoice
         * @memberof payment.PaymentMgr
         * @instance
         * @param {payment.IInvoice} request Invoice message or plain object
         * @returns {Promise<payment.Invoice>} Promise
         * @variation 2
         */

        /**
         * Callback as used by {@link payment.PaymentMgr#listPlans}.
         * @memberof payment.PaymentMgr
         * @typedef ListPlansCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {payment.Plans} [response] Plans
         */

        /**
         * Calls ListPlans.
         * @function listPlans
         * @memberof payment.PaymentMgr
         * @instance
         * @param {common.IEmpty} request Empty message or plain object
         * @param {payment.PaymentMgr.ListPlansCallback} callback Node-style callback called with the error, if any, and Plans
         * @returns {undefined}
         * @variation 1
         */
        PaymentMgr.prototype.listPlans = function listPlans(request, callback) {
            return this.rpcCall(listPlans, $root.common.Empty, $root.payment.Plans, request, callback);
        };

        /**
         * Calls ListPlans.
         * @function listPlans
         * @memberof payment.PaymentMgr
         * @instance
         * @param {common.IEmpty} request Empty message or plain object
         * @returns {Promise<payment.Plans>} Promise
         * @variation 2
         */

        /**
         * Callback as used by {@link payment.PaymentMgr#exportInvoice}.
         * @memberof payment.PaymentMgr
         * @typedef ExportInvoiceCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {payment.String} [response] String
         */

        /**
         * Calls ExportInvoice.
         * @function exportInvoice
         * @memberof payment.PaymentMgr
         * @instance
         * @param {common.IId} request Id message or plain object
         * @param {payment.PaymentMgr.ExportInvoiceCallback} callback Node-style callback called with the error, if any, and String
         * @returns {undefined}
         * @variation 1
         */
        PaymentMgr.prototype.exportInvoice = function exportInvoice(request, callback) {
            return this.rpcCall(exportInvoice, $root.common.Id, $root.payment.String, request, callback);
        };

        /**
         * Calls ExportInvoice.
         * @function exportInvoice
         * @memberof payment.PaymentMgr
         * @instance
         * @param {common.IId} request Id message or plain object
         * @returns {Promise<payment.String>} Promise
         * @variation 2
         */

        /**
         * Callback as used by {@link payment.PaymentMgr#getExchangeRate}.
         * @memberof payment.PaymentMgr
         * @typedef GetExchangeRateCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {payment.ExchangeRate} [response] ExchangeRate
         */

        /**
         * Calls GetExchangeRate.
         * @function getExchangeRate
         * @memberof payment.PaymentMgr
         * @instance
         * @param {payment.IExchangeRate} request ExchangeRate message or plain object
         * @param {payment.PaymentMgr.GetExchangeRateCallback} callback Node-style callback called with the error, if any, and ExchangeRate
         * @returns {undefined}
         * @variation 1
         */
        PaymentMgr.prototype.getExchangeRate = function getExchangeRate(request, callback) {
            return this.rpcCall(getExchangeRate, $root.payment.ExchangeRate, $root.payment.ExchangeRate, request, callback);
        };

        /**
         * Calls GetExchangeRate.
         * @function getExchangeRate
         * @memberof payment.PaymentMgr
         * @instance
         * @param {payment.IExchangeRate} request ExchangeRate message or plain object
         * @returns {Promise<payment.ExchangeRate>} Promise
         * @variation 2
         */

        /**
         * Callback as used by {@link payment.PaymentMgr#transferMoney}.
         * @memberof payment.PaymentMgr
         * @typedef TransferMoneyCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {payment.Bill} [response] Bill
         */

        /**
         * Calls TransferMoney.
         * @function transferMoney
         * @memberof payment.PaymentMgr
         * @instance
         * @param {payment.IPayRequest} request PayRequest message or plain object
         * @param {payment.PaymentMgr.TransferMoneyCallback} callback Node-style callback called with the error, if any, and Bill
         * @returns {undefined}
         * @variation 1
         */
        PaymentMgr.prototype.transferMoney = function transferMoney(request, callback) {
            return this.rpcCall(transferMoney, $root.payment.PayRequest, $root.payment.Bill, request, callback);
        };

        /**
         * Calls TransferMoney.
         * @function transferMoney
         * @memberof payment.PaymentMgr
         * @instance
         * @param {payment.IPayRequest} request PayRequest message or plain object
         * @returns {Promise<payment.Bill>} Promise
         * @variation 2
         */

        return PaymentMgr;
    })();

    payment.String = (function() {

        /**
         * Properties of a String.
         * @memberof payment
         * @interface IString
         * @property {string|null} [str] String str
         */

        /**
         * Constructs a new String.
         * @memberof payment
         * @classdesc Represents a String.
         * @implements IString
         * @constructor
         * @param {payment.IString=} [p] Properties to set
         */
        function String(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * String str.
         * @member {string} str
         * @memberof payment.String
         * @instance
         */
        String.prototype.str = "";

        return String;
    })();

    payment.PayRequest = (function() {

        /**
         * Properties of a PayRequest.
         * @memberof payment
         * @interface IPayRequest
         * @property {string|null} [account_id] PayRequest account_id
         * @property {Array.<string>|null} [invoice_ids] PayRequest invoice_ids
         * @property {string|null} [description] PayRequest description
         * @property {payment.IContact|null} [CustomerInfo] PayRequest CustomerInfo
         * @property {number|null} [amount] PayRequest amount
         */

        /**
         * Constructs a new PayRequest.
         * @memberof payment
         * @classdesc Represents a PayRequest.
         * @implements IPayRequest
         * @constructor
         * @param {payment.IPayRequest=} [p] Properties to set
         */
        function PayRequest(p) {
            this.invoice_ids = [];
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * PayRequest account_id.
         * @member {string} account_id
         * @memberof payment.PayRequest
         * @instance
         */
        PayRequest.prototype.account_id = "";

        /**
         * PayRequest invoice_ids.
         * @member {Array.<string>} invoice_ids
         * @memberof payment.PayRequest
         * @instance
         */
        PayRequest.prototype.invoice_ids = $util.emptyArray;

        /**
         * PayRequest description.
         * @member {string} description
         * @memberof payment.PayRequest
         * @instance
         */
        PayRequest.prototype.description = "";

        /**
         * PayRequest CustomerInfo.
         * @member {payment.IContact|null|undefined} CustomerInfo
         * @memberof payment.PayRequest
         * @instance
         */
        PayRequest.prototype.CustomerInfo = null;

        /**
         * PayRequest amount.
         * @member {number} amount
         * @memberof payment.PayRequest
         * @instance
         */
        PayRequest.prototype.amount = 0;

        return PayRequest;
    })();

    payment.ESubscription = (function() {

        /**
         * Properties of a ESubscription.
         * @memberof payment
         * @interface IESubscription
         * @property {payment.ISubscription|null} [sub] ESubscription sub
         * @property {string|null} [err] ESubscription err
         */

        /**
         * Constructs a new ESubscription.
         * @memberof payment
         * @classdesc Represents a ESubscription.
         * @implements IESubscription
         * @constructor
         * @param {payment.IESubscription=} [p] Properties to set
         */
        function ESubscription(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * ESubscription sub.
         * @member {payment.ISubscription|null|undefined} sub
         * @memberof payment.ESubscription
         * @instance
         */
        ESubscription.prototype.sub = null;

        /**
         * ESubscription err.
         * @member {string} err
         * @memberof payment.ESubscription
         * @instance
         */
        ESubscription.prototype.err = "";

        return ESubscription;
    })();

    payment.EInvoice = (function() {

        /**
         * Properties of a EInvoice.
         * @memberof payment
         * @interface IEInvoice
         * @property {payment.IInvoice|null} [inv] EInvoice inv
         * @property {string|null} [err] EInvoice err
         */

        /**
         * Constructs a new EInvoice.
         * @memberof payment
         * @classdesc Represents a EInvoice.
         * @implements IEInvoice
         * @constructor
         * @param {payment.IEInvoice=} [p] Properties to set
         */
        function EInvoice(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * EInvoice inv.
         * @member {payment.IInvoice|null|undefined} inv
         * @memberof payment.EInvoice
         * @instance
         */
        EInvoice.prototype.inv = null;

        /**
         * EInvoice err.
         * @member {string} err
         * @memberof payment.EInvoice
         * @instance
         */
        EInvoice.prototype.err = "";

        return EInvoice;
    })();

    payment.InvoiceCreatedEmail = (function() {

        /**
         * Properties of an InvoiceCreatedEmail.
         * @memberof payment
         * @interface IInvoiceCreatedEmail
         * @property {common.IContext|null} [ctx] InvoiceCreatedEmail ctx
         * @property {string|null} [to] InvoiceCreatedEmail to
         * @property {string|null} [account_id] InvoiceCreatedEmail account_id
         * @property {string|null} [billing_name] InvoiceCreatedEmail billing_name
         * @property {string|null} [invoice_id] InvoiceCreatedEmail invoice_id
         * @property {number|Long|null} [created] InvoiceCreatedEmail created
         * @property {lang.L|null} [lang] InvoiceCreatedEmail lang
         * @property {string|null} [from] InvoiceCreatedEmail from
         */

        /**
         * Constructs a new InvoiceCreatedEmail.
         * @memberof payment
         * @classdesc Represents an InvoiceCreatedEmail.
         * @implements IInvoiceCreatedEmail
         * @constructor
         * @param {payment.IInvoiceCreatedEmail=} [p] Properties to set
         */
        function InvoiceCreatedEmail(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * InvoiceCreatedEmail ctx.
         * @member {common.IContext|null|undefined} ctx
         * @memberof payment.InvoiceCreatedEmail
         * @instance
         */
        InvoiceCreatedEmail.prototype.ctx = null;

        /**
         * InvoiceCreatedEmail to.
         * @member {string} to
         * @memberof payment.InvoiceCreatedEmail
         * @instance
         */
        InvoiceCreatedEmail.prototype.to = "";

        /**
         * InvoiceCreatedEmail account_id.
         * @member {string} account_id
         * @memberof payment.InvoiceCreatedEmail
         * @instance
         */
        InvoiceCreatedEmail.prototype.account_id = "";

        /**
         * InvoiceCreatedEmail billing_name.
         * @member {string} billing_name
         * @memberof payment.InvoiceCreatedEmail
         * @instance
         */
        InvoiceCreatedEmail.prototype.billing_name = "";

        /**
         * InvoiceCreatedEmail invoice_id.
         * @member {string} invoice_id
         * @memberof payment.InvoiceCreatedEmail
         * @instance
         */
        InvoiceCreatedEmail.prototype.invoice_id = "";

        /**
         * InvoiceCreatedEmail created.
         * @member {number|Long} created
         * @memberof payment.InvoiceCreatedEmail
         * @instance
         */
        InvoiceCreatedEmail.prototype.created = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

        /**
         * InvoiceCreatedEmail lang.
         * @member {lang.L} lang
         * @memberof payment.InvoiceCreatedEmail
         * @instance
         */
        InvoiceCreatedEmail.prototype.lang = 0;

        /**
         * InvoiceCreatedEmail from.
         * @member {string} from
         * @memberof payment.InvoiceCreatedEmail
         * @instance
         */
        InvoiceCreatedEmail.prototype.from = "";

        return InvoiceCreatedEmail;
    })();

    return payment;
})();

export const pubsub = $root.pubsub = (() => {

    /**
     * Namespace pubsub.
     * @exports pubsub
     * @namespace
     */
    const pubsub = {};

    pubsub.Subscription = (function() {

        /**
         * Properties of a Subscription.
         * @memberof pubsub
         * @interface ISubscription
         * @property {common.IContext|null} [ctx] Subscription ctx
         * @property {string|null} [user_id] Subscription user_id
         * @property {string|null} [topic] Subscription topic
         * @property {string|null} [sub_id] Subscription sub_id
         * @property {string|null} [target_topic] Subscription target_topic
         * @property {string|null} [target_key] Subscription target_key
         * @property {number|Long|null} [ttls] Subscription ttls
         * @property {string|null} [router_topic] Subscription router_topic
         * @property {number|null} [target_partition] Subscription target_partition
         * @property {Array.<string>|null} [topics] Subscription topics
         */

        /**
         * Constructs a new Subscription.
         * @memberof pubsub
         * @classdesc Represents a Subscription.
         * @implements ISubscription
         * @constructor
         * @param {pubsub.ISubscription=} [p] Properties to set
         */
        function Subscription(p) {
            this.topics = [];
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * Subscription ctx.
         * @member {common.IContext|null|undefined} ctx
         * @memberof pubsub.Subscription
         * @instance
         */
        Subscription.prototype.ctx = null;

        /**
         * Subscription user_id.
         * @member {string} user_id
         * @memberof pubsub.Subscription
         * @instance
         */
        Subscription.prototype.user_id = "";

        /**
         * Subscription topic.
         * @member {string} topic
         * @memberof pubsub.Subscription
         * @instance
         */
        Subscription.prototype.topic = "";

        /**
         * Subscription sub_id.
         * @member {string} sub_id
         * @memberof pubsub.Subscription
         * @instance
         */
        Subscription.prototype.sub_id = "";

        /**
         * Subscription target_topic.
         * @member {string} target_topic
         * @memberof pubsub.Subscription
         * @instance
         */
        Subscription.prototype.target_topic = "";

        /**
         * Subscription target_key.
         * @member {string} target_key
         * @memberof pubsub.Subscription
         * @instance
         */
        Subscription.prototype.target_key = "";

        /**
         * Subscription ttls.
         * @member {number|Long} ttls
         * @memberof pubsub.Subscription
         * @instance
         */
        Subscription.prototype.ttls = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

        /**
         * Subscription router_topic.
         * @member {string} router_topic
         * @memberof pubsub.Subscription
         * @instance
         */
        Subscription.prototype.router_topic = "";

        /**
         * Subscription target_partition.
         * @member {number} target_partition
         * @memberof pubsub.Subscription
         * @instance
         */
        Subscription.prototype.target_partition = 0;

        /**
         * Subscription topics.
         * @member {Array.<string>} topics
         * @memberof pubsub.Subscription
         * @instance
         */
        Subscription.prototype.topics = $util.emptyArray;

        return Subscription;
    })();

    pubsub.PublishMessage = (function() {

        /**
         * Properties of a PublishMessage.
         * @memberof pubsub
         * @interface IPublishMessage
         * @property {common.IContext|null} [ctx] PublishMessage ctx
         * @property {Array.<string>|null} [topics] PublishMessage topics
         * @property {Uint8Array|null} [payload] PublishMessage payload
         * @property {Uint8Array|null} [user_ids_filter] PublishMessage user_ids_filter
         * @property {number|null} [filter_item_bit_size] PublishMessage filter_item_bit_size
         * @property {number|null} [filter_hash_count] PublishMessage filter_hash_count
         * @property {Uint8Array|null} [neg_user_ids_filter] PublishMessage neg_user_ids_filter
         */

        /**
         * Constructs a new PublishMessage.
         * @memberof pubsub
         * @classdesc Represents a PublishMessage.
         * @implements IPublishMessage
         * @constructor
         * @param {pubsub.IPublishMessage=} [p] Properties to set
         */
        function PublishMessage(p) {
            this.topics = [];
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * PublishMessage ctx.
         * @member {common.IContext|null|undefined} ctx
         * @memberof pubsub.PublishMessage
         * @instance
         */
        PublishMessage.prototype.ctx = null;

        /**
         * PublishMessage topics.
         * @member {Array.<string>} topics
         * @memberof pubsub.PublishMessage
         * @instance
         */
        PublishMessage.prototype.topics = $util.emptyArray;

        /**
         * PublishMessage payload.
         * @member {Uint8Array} payload
         * @memberof pubsub.PublishMessage
         * @instance
         */
        PublishMessage.prototype.payload = $util.newBuffer([]);

        /**
         * PublishMessage user_ids_filter.
         * @member {Uint8Array} user_ids_filter
         * @memberof pubsub.PublishMessage
         * @instance
         */
        PublishMessage.prototype.user_ids_filter = $util.newBuffer([]);

        /**
         * PublishMessage filter_item_bit_size.
         * @member {number} filter_item_bit_size
         * @memberof pubsub.PublishMessage
         * @instance
         */
        PublishMessage.prototype.filter_item_bit_size = 0;

        /**
         * PublishMessage filter_hash_count.
         * @member {number} filter_hash_count
         * @memberof pubsub.PublishMessage
         * @instance
         */
        PublishMessage.prototype.filter_hash_count = 0;

        /**
         * PublishMessage neg_user_ids_filter.
         * @member {Uint8Array} neg_user_ids_filter
         * @memberof pubsub.PublishMessage
         * @instance
         */
        PublishMessage.prototype.neg_user_ids_filter = $util.newBuffer([]);

        return PublishMessage;
    })();

    pubsub.Pubsub = (function() {

        /**
         * Constructs a new Pubsub service.
         * @memberof pubsub
         * @classdesc Represents a Pubsub
         * @extends $protobuf.rpc.Service
         * @constructor
         * @param {$protobuf.RPCImpl} rpcImpl RPC implementation
         * @param {boolean} [requestDelimited=false] Whether requests are length-delimited
         * @param {boolean} [responseDelimited=false] Whether responses are length-delimited
         */
        function Pubsub(rpcImpl, requestDelimited, responseDelimited) {
            $protobuf.rpc.Service.call(this, rpcImpl, requestDelimited, responseDelimited);
        }

        (Pubsub.prototype = Object.create($protobuf.rpc.Service.prototype)).constructor = Pubsub;

        /**
         * Callback as used by {@link pubsub.Pubsub#subscribe}.
         * @memberof pubsub.Pubsub
         * @typedef SubscribeCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {common.Empty} [response] Empty
         */

        /**
         * Calls Subscribe.
         * @function subscribe
         * @memberof pubsub.Pubsub
         * @instance
         * @param {pubsub.ISubscription} request Subscription message or plain object
         * @param {pubsub.Pubsub.SubscribeCallback} callback Node-style callback called with the error, if any, and Empty
         * @returns {undefined}
         * @variation 1
         */
        Pubsub.prototype.subscribe = function subscribe(request, callback) {
            return this.rpcCall(subscribe, $root.pubsub.Subscription, $root.common.Empty, request, callback);
        };

        /**
         * Calls Subscribe.
         * @function subscribe
         * @memberof pubsub.Pubsub
         * @instance
         * @param {pubsub.ISubscription} request Subscription message or plain object
         * @returns {Promise<common.Empty>} Promise
         * @variation 2
         */

        /**
         * Callback as used by {@link pubsub.Pubsub#unsubscribe}.
         * @memberof pubsub.Pubsub
         * @typedef UnsubscribeCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {common.Empty} [response] Empty
         */

        /**
         * Calls Unsubscribe.
         * @function unsubscribe
         * @memberof pubsub.Pubsub
         * @instance
         * @param {pubsub.ISubscription} request Subscription message or plain object
         * @param {pubsub.Pubsub.UnsubscribeCallback} callback Node-style callback called with the error, if any, and Empty
         * @returns {undefined}
         * @variation 1
         */
        Pubsub.prototype.unsubscribe = function unsubscribe(request, callback) {
            return this.rpcCall(unsubscribe, $root.pubsub.Subscription, $root.common.Empty, request, callback);
        };

        /**
         * Calls Unsubscribe.
         * @function unsubscribe
         * @memberof pubsub.Pubsub
         * @instance
         * @param {pubsub.ISubscription} request Subscription message or plain object
         * @returns {Promise<common.Empty>} Promise
         * @variation 2
         */

        return Pubsub;
    })();

    /**
     * Event enum.
     * @name pubsub.Event
     * @enum {string}
     * @property {number} PubsubSynced=0 PubsubSynced value
     * @property {number} PubsubRequested=1 PubsubRequested value
     * @property {number} PubsubPublish=2 PubsubPublish value
     */
    pubsub.Event = (function() {
        const valuesById = {}, values = Object.create(valuesById);
        values[valuesById[0] = "PubsubSynced"] = 0;
        values[valuesById[1] = "PubsubRequested"] = 1;
        values[valuesById[2] = "PubsubPublish"] = 2;
        return values;
    })();

    return pubsub;
})();

export const scheduler = $root.scheduler = (() => {

    /**
     * Namespace scheduler.
     * @exports scheduler
     * @namespace
     */
    const scheduler = {};

    scheduler.Task = (function() {

        /**
         * Properties of a Task.
         * @memberof scheduler
         * @interface ITask
         * @property {string|null} [id] Task id
         * @property {number|Long|null} [callback_time] Task callback_time
         * @property {string|null} [topic] Task topic
         * @property {Uint8Array|null} [data] Task data
         * @property {string|null} [key] Task key
         * @property {number|Long|null} [called] Task called
         * @property {number|Long|null} [sec] Task sec
         * @property {string|null} [par] Task par
         */

        /**
         * Constructs a new Task.
         * @memberof scheduler
         * @classdesc Represents a Task.
         * @implements ITask
         * @constructor
         * @param {scheduler.ITask=} [p] Properties to set
         */
        function Task(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * Task id.
         * @member {string} id
         * @memberof scheduler.Task
         * @instance
         */
        Task.prototype.id = "";

        /**
         * Task callback_time.
         * @member {number|Long} callback_time
         * @memberof scheduler.Task
         * @instance
         */
        Task.prototype.callback_time = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

        /**
         * Task topic.
         * @member {string} topic
         * @memberof scheduler.Task
         * @instance
         */
        Task.prototype.topic = "";

        /**
         * Task data.
         * @member {Uint8Array} data
         * @memberof scheduler.Task
         * @instance
         */
        Task.prototype.data = $util.newBuffer([]);

        /**
         * Task key.
         * @member {string} key
         * @memberof scheduler.Task
         * @instance
         */
        Task.prototype.key = "";

        /**
         * Task called.
         * @member {number|Long} called
         * @memberof scheduler.Task
         * @instance
         */
        Task.prototype.called = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

        /**
         * Task sec.
         * @member {number|Long} sec
         * @memberof scheduler.Task
         * @instance
         */
        Task.prototype.sec = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

        /**
         * Task par.
         * @member {string} par
         * @memberof scheduler.Task
         * @instance
         */
        Task.prototype.par = "";

        return Task;
    })();

    scheduler.Id = (function() {

        /**
         * Properties of an Id.
         * @memberof scheduler
         * @interface IId
         * @property {string|null} [id] Id id
         * @property {number|Long|null} [callback_time] Id callback_time
         */

        /**
         * Constructs a new Id.
         * @memberof scheduler
         * @classdesc Represents an Id.
         * @implements IId
         * @constructor
         * @param {scheduler.IId=} [p] Properties to set
         */
        function Id(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * Id id.
         * @member {string} id
         * @memberof scheduler.Id
         * @instance
         */
        Id.prototype.id = "";

        /**
         * Id callback_time.
         * @member {number|Long} callback_time
         * @memberof scheduler.Id
         * @instance
         */
        Id.prototype.callback_time = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

        return Id;
    })();

    /**
     * Event enum.
     * @name scheduler.Event
     * @enum {string}
     * @property {number} SchedulerRequested=0 SchedulerRequested value
     * @property {number} SchedulerCancelled=1 SchedulerCancelled value
     */
    scheduler.Event = (function() {
        const valuesById = {}, values = Object.create(valuesById);
        values[valuesById[0] = "SchedulerRequested"] = 0;
        values[valuesById[1] = "SchedulerCancelled"] = 1;
        return values;
    })();

    return scheduler;
})();

export const template = $root.template = (() => {

    /**
     * Namespace template.
     * @exports template
     * @namespace
     */
    const template = {};

    template.Template = (function() {

        /**
         * Properties of a Template.
         * @memberof template
         * @interface ITemplate
         * @property {string|null} [Id] Template Id
         * @property {string|null} [Language] Template Language
         * @property {template.Type|null} [Type] Template Type
         * @property {template.IEmailTemplate|null} [Email] Template Email
         * @property {template.IWebPushTemplate|null} [WebPush] Template WebPush
         * @property {template.INotificationTemplate|null} [Notification] Template Notification
         */

        /**
         * Constructs a new Template.
         * @memberof template
         * @classdesc Represents a Template.
         * @implements ITemplate
         * @constructor
         * @param {template.ITemplate=} [p] Properties to set
         */
        function Template(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * Template Id.
         * @member {string} Id
         * @memberof template.Template
         * @instance
         */
        Template.prototype.Id = "";

        /**
         * Template Language.
         * @member {string} Language
         * @memberof template.Template
         * @instance
         */
        Template.prototype.Language = "";

        /**
         * Template Type.
         * @member {template.Type} Type
         * @memberof template.Template
         * @instance
         */
        Template.prototype.Type = 0;

        /**
         * Template Email.
         * @member {template.IEmailTemplate|null|undefined} Email
         * @memberof template.Template
         * @instance
         */
        Template.prototype.Email = null;

        /**
         * Template WebPush.
         * @member {template.IWebPushTemplate|null|undefined} WebPush
         * @memberof template.Template
         * @instance
         */
        Template.prototype.WebPush = null;

        /**
         * Template Notification.
         * @member {template.INotificationTemplate|null|undefined} Notification
         * @memberof template.Template
         * @instance
         */
        Template.prototype.Notification = null;

        return Template;
    })();

    /**
     * Type enum.
     * @name template.Type
     * @enum {string}
     * @property {number} Email=0 Email value
     * @property {number} WebPush=1 WebPush value
     * @property {number} Notification=2 Notification value
     */
    template.Type = (function() {
        const valuesById = {}, values = Object.create(valuesById);
        values[valuesById[0] = "Email"] = 0;
        values[valuesById[1] = "WebPush"] = 1;
        values[valuesById[2] = "Notification"] = 2;
        return values;
    })();

    template.EmailTemplate = (function() {

        /**
         * Properties of an EmailTemplate.
         * @memberof template
         * @interface IEmailTemplate
         * @property {string|null} [Subject] EmailTemplate Subject
         * @property {string|null} [Body] EmailTemplate Body
         */

        /**
         * Constructs a new EmailTemplate.
         * @memberof template
         * @classdesc Represents an EmailTemplate.
         * @implements IEmailTemplate
         * @constructor
         * @param {template.IEmailTemplate=} [p] Properties to set
         */
        function EmailTemplate(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * EmailTemplate Subject.
         * @member {string} Subject
         * @memberof template.EmailTemplate
         * @instance
         */
        EmailTemplate.prototype.Subject = "";

        /**
         * EmailTemplate Body.
         * @member {string} Body
         * @memberof template.EmailTemplate
         * @instance
         */
        EmailTemplate.prototype.Body = "";

        return EmailTemplate;
    })();

    template.WebPushTemplate = (function() {

        /**
         * Properties of a WebPushTemplate.
         * @memberof template
         * @interface IWebPushTemplate
         * @property {string|null} [Title] WebPushTemplate Title
         * @property {string|null} [Body] WebPushTemplate Body
         * @property {string|null} [Icon] WebPushTemplate Icon
         */

        /**
         * Constructs a new WebPushTemplate.
         * @memberof template
         * @classdesc Represents a WebPushTemplate.
         * @implements IWebPushTemplate
         * @constructor
         * @param {template.IWebPushTemplate=} [p] Properties to set
         */
        function WebPushTemplate(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * WebPushTemplate Title.
         * @member {string} Title
         * @memberof template.WebPushTemplate
         * @instance
         */
        WebPushTemplate.prototype.Title = "";

        /**
         * WebPushTemplate Body.
         * @member {string} Body
         * @memberof template.WebPushTemplate
         * @instance
         */
        WebPushTemplate.prototype.Body = "";

        /**
         * WebPushTemplate Icon.
         * @member {string} Icon
         * @memberof template.WebPushTemplate
         * @instance
         */
        WebPushTemplate.prototype.Icon = "";

        return WebPushTemplate;
    })();

    template.NotificationTemplate = (function() {

        /**
         * Properties of a NotificationTemplate.
         * @memberof template
         * @interface INotificationTemplate
         * @property {string|null} [Image] NotificationTemplate Image
         * @property {string|null} [Footer] NotificationTemplate Footer
         * @property {string|null} [Body] NotificationTemplate Body
         */

        /**
         * Constructs a new NotificationTemplate.
         * @memberof template
         * @classdesc Represents a NotificationTemplate.
         * @implements INotificationTemplate
         * @constructor
         * @param {template.INotificationTemplate=} [p] Properties to set
         */
        function NotificationTemplate(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * NotificationTemplate Image.
         * @member {string} Image
         * @memberof template.NotificationTemplate
         * @instance
         */
        NotificationTemplate.prototype.Image = "";

        /**
         * NotificationTemplate Footer.
         * @member {string} Footer
         * @memberof template.NotificationTemplate
         * @instance
         */
        NotificationTemplate.prototype.Footer = "";

        /**
         * NotificationTemplate Body.
         * @member {string} Body
         * @memberof template.NotificationTemplate
         * @instance
         */
        NotificationTemplate.prototype.Body = "";

        return NotificationTemplate;
    })();

    return template;
})();

export const user = $root.user = (() => {

    /**
     * Namespace user.
     * @exports user
     * @namespace
     */
    const user = {};

    user.AllType = (function() {

        /**
         * Properties of an AllType.
         * @memberof user
         * @interface IAllType
         * @property {user.IUser|null} [user] AllType user
         * @property {user.ICreateRequest|null} [cr] AllType cr
         * @property {user.IField|null} [fields] AllType fields
         * @property {user.IUserCreateResult|null} [ucr] AllType ucr
         * @property {user.IReadTopicRequest|null} [rpr] AllType rpr
         * @property {user.IUserSearchRequest|null} [usersearchrequest] AllType usersearchrequest
         * @property {user.IUserSearchResult|null} [usersearchresult] AllType usersearchresult
         * @property {user.IMaskResponse|null} [ms] AllType ms
         * @property {user.IListTopicsResult|null} [ltr] AllType ltr
         * @property {user.ISegmentation|null} [sg] AllType sg
         * @property {user.IPresence|null} [presn] AllType presn
         * @property {user.IVisitor|null} [vi] AllType vi
         * @property {user.IVisitors|null} [vis] AllType vis
         * @property {user.ITopic|null} [topic] AllType topic
         * @property {user.IUnreadTopic|null} [utopic] AllType utopic
         * @property {user.IMyUser|null} [myuser] AllType myuser
         * @property {user.IAutomation|null} [automation] AllType automation
         * @property {user.IAutomations|null} [automations] AllType automations
         * @property {user.ISession|null} [session] AllType session
         */

        /**
         * Constructs a new AllType.
         * @memberof user
         * @classdesc Represents an AllType.
         * @implements IAllType
         * @constructor
         * @param {user.IAllType=} [p] Properties to set
         */
        function AllType(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * AllType user.
         * @member {user.IUser|null|undefined} user
         * @memberof user.AllType
         * @instance
         */
        AllType.prototype.user = null;

        /**
         * AllType cr.
         * @member {user.ICreateRequest|null|undefined} cr
         * @memberof user.AllType
         * @instance
         */
        AllType.prototype.cr = null;

        /**
         * AllType fields.
         * @member {user.IField|null|undefined} fields
         * @memberof user.AllType
         * @instance
         */
        AllType.prototype.fields = null;

        /**
         * AllType ucr.
         * @member {user.IUserCreateResult|null|undefined} ucr
         * @memberof user.AllType
         * @instance
         */
        AllType.prototype.ucr = null;

        /**
         * AllType rpr.
         * @member {user.IReadTopicRequest|null|undefined} rpr
         * @memberof user.AllType
         * @instance
         */
        AllType.prototype.rpr = null;

        /**
         * AllType usersearchrequest.
         * @member {user.IUserSearchRequest|null|undefined} usersearchrequest
         * @memberof user.AllType
         * @instance
         */
        AllType.prototype.usersearchrequest = null;

        /**
         * AllType usersearchresult.
         * @member {user.IUserSearchResult|null|undefined} usersearchresult
         * @memberof user.AllType
         * @instance
         */
        AllType.prototype.usersearchresult = null;

        /**
         * AllType ms.
         * @member {user.IMaskResponse|null|undefined} ms
         * @memberof user.AllType
         * @instance
         */
        AllType.prototype.ms = null;

        /**
         * AllType ltr.
         * @member {user.IListTopicsResult|null|undefined} ltr
         * @memberof user.AllType
         * @instance
         */
        AllType.prototype.ltr = null;

        /**
         * AllType sg.
         * @member {user.ISegmentation|null|undefined} sg
         * @memberof user.AllType
         * @instance
         */
        AllType.prototype.sg = null;

        /**
         * AllType presn.
         * @member {user.IPresence|null|undefined} presn
         * @memberof user.AllType
         * @instance
         */
        AllType.prototype.presn = null;

        /**
         * AllType vi.
         * @member {user.IVisitor|null|undefined} vi
         * @memberof user.AllType
         * @instance
         */
        AllType.prototype.vi = null;

        /**
         * AllType vis.
         * @member {user.IVisitors|null|undefined} vis
         * @memberof user.AllType
         * @instance
         */
        AllType.prototype.vis = null;

        /**
         * AllType topic.
         * @member {user.ITopic|null|undefined} topic
         * @memberof user.AllType
         * @instance
         */
        AllType.prototype.topic = null;

        /**
         * AllType utopic.
         * @member {user.IUnreadTopic|null|undefined} utopic
         * @memberof user.AllType
         * @instance
         */
        AllType.prototype.utopic = null;

        /**
         * AllType myuser.
         * @member {user.IMyUser|null|undefined} myuser
         * @memberof user.AllType
         * @instance
         */
        AllType.prototype.myuser = null;

        /**
         * AllType automation.
         * @member {user.IAutomation|null|undefined} automation
         * @memberof user.AllType
         * @instance
         */
        AllType.prototype.automation = null;

        /**
         * AllType automations.
         * @member {user.IAutomations|null|undefined} automations
         * @memberof user.AllType
         * @instance
         */
        AllType.prototype.automations = null;

        /**
         * AllType session.
         * @member {user.ISession|null|undefined} session
         * @memberof user.AllType
         * @instance
         */
        AllType.prototype.session = null;

        return AllType;
    })();

    user.MyServer = (function() {

        /**
         * Constructs a new MyServer service.
         * @memberof user
         * @classdesc Represents a MyServer
         * @extends $protobuf.rpc.Service
         * @constructor
         * @param {$protobuf.RPCImpl} rpcImpl RPC implementation
         * @param {boolean} [requestDelimited=false] Whether requests are length-delimited
         * @param {boolean} [responseDelimited=false] Whether responses are length-delimited
         */
        function MyServer(rpcImpl, requestDelimited, responseDelimited) {
            $protobuf.rpc.Service.call(this, rpcImpl, requestDelimited, responseDelimited);
        }

        (MyServer.prototype = Object.create($protobuf.rpc.Service.prototype)).constructor = MyServer;

        /**
         * Callback as used by {@link user.MyServer#do_}.
         * @memberof user.MyServer
         * @typedef DoCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {user.AllType} [response] AllType
         */

        /**
         * Calls Do.
         * @function do
         * @memberof user.MyServer
         * @instance
         * @param {user.IAllType} request AllType message or plain object
         * @param {user.MyServer.DoCallback} callback Node-style callback called with the error, if any, and AllType
         * @returns {undefined}
         * @variation 1
         */
        MyServer.prototype["do"] = function do_(request, callback) {
            return this.rpcCall(do_, $root.user.AllType, $root.user.AllType, request, callback);
        };

        /**
         * Calls Do.
         * @function do
         * @memberof user.MyServer
         * @instance
         * @param {user.IAllType} request AllType message or plain object
         * @returns {Promise<user.AllType>} Promise
         * @variation 2
         */

        return MyServer;
    })();

    user.AddToMyRequest = (function() {

        /**
         * Properties of an AddToMyRequest.
         * @memberof user
         * @interface IAddToMyRequest
         * @property {common.IContext|null} [ctx] AddToMyRequest ctx
         * @property {string|null} [user_id] AddToMyRequest user_id
         * @property {Array.<string>|null} [agent_ids] AddToMyRequest agent_ids
         */

        /**
         * Constructs a new AddToMyRequest.
         * @memberof user
         * @classdesc Represents an AddToMyRequest.
         * @implements IAddToMyRequest
         * @constructor
         * @param {user.IAddToMyRequest=} [p] Properties to set
         */
        function AddToMyRequest(p) {
            this.agent_ids = [];
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * AddToMyRequest ctx.
         * @member {common.IContext|null|undefined} ctx
         * @memberof user.AddToMyRequest
         * @instance
         */
        AddToMyRequest.prototype.ctx = null;

        /**
         * AddToMyRequest user_id.
         * @member {string} user_id
         * @memberof user.AddToMyRequest
         * @instance
         */
        AddToMyRequest.prototype.user_id = "";

        /**
         * AddToMyRequest agent_ids.
         * @member {Array.<string>} agent_ids
         * @memberof user.AddToMyRequest
         * @instance
         */
        AddToMyRequest.prototype.agent_ids = $util.emptyArray;

        return AddToMyRequest;
    })();

    user.UserCreateResult = (function() {

        /**
         * Properties of a UserCreateResult.
         * @memberof user
         * @interface IUserCreateResult
         * @property {string|null} [id] UserCreateResult id
         * @property {string|null} [mask] UserCreateResult mask
         */

        /**
         * Constructs a new UserCreateResult.
         * @memberof user
         * @classdesc Represents a UserCreateResult.
         * @implements IUserCreateResult
         * @constructor
         * @param {user.IUserCreateResult=} [p] Properties to set
         */
        function UserCreateResult(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * UserCreateResult id.
         * @member {string} id
         * @memberof user.UserCreateResult
         * @instance
         */
        UserCreateResult.prototype.id = "";

        /**
         * UserCreateResult mask.
         * @member {string} mask
         * @memberof user.UserCreateResult
         * @instance
         */
        UserCreateResult.prototype.mask = "";

        return UserCreateResult;
    })();

    user.MyUser = (function() {

        /**
         * Properties of a MyUser.
         * @memberof user
         * @interface IMyUser
         * @property {common.IContext|null} [ctx] MyUser ctx
         * @property {string|null} [agent_id] MyUser agent_id
         * @property {user.IUser|null} [user] MyUser user
         * @property {number|null} [unread] MyUser unread
         * @property {number|Long|null} [updated] MyUser updated
         */

        /**
         * Constructs a new MyUser.
         * @memberof user
         * @classdesc Represents a MyUser.
         * @implements IMyUser
         * @constructor
         * @param {user.IMyUser=} [p] Properties to set
         */
        function MyUser(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * MyUser ctx.
         * @member {common.IContext|null|undefined} ctx
         * @memberof user.MyUser
         * @instance
         */
        MyUser.prototype.ctx = null;

        /**
         * MyUser agent_id.
         * @member {string} agent_id
         * @memberof user.MyUser
         * @instance
         */
        MyUser.prototype.agent_id = "";

        /**
         * MyUser user.
         * @member {user.IUser|null|undefined} user
         * @memberof user.MyUser
         * @instance
         */
        MyUser.prototype.user = null;

        /**
         * MyUser unread.
         * @member {number} unread
         * @memberof user.MyUser
         * @instance
         */
        MyUser.prototype.unread = 0;

        /**
         * MyUser updated.
         * @member {number|Long} updated
         * @memberof user.MyUser
         * @instance
         */
        MyUser.prototype.updated = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

        return MyUser;
    })();

    /**
     * AttributeType enum.
     * @name user.AttributeType
     * @enum {string}
     * @property {number} text=0 text value
     * @property {number} number=1 number value
     * @property {number} boolean=2 boolean value
     * @property {number} datetime=3 datetime value
     * @property {number} list=4 list value
     */
    user.AttributeType = (function() {
        const valuesById = {}, values = Object.create(valuesById);
        values[valuesById[0] = "text"] = 0;
        values[valuesById[1] = "number"] = 1;
        values[valuesById[2] = "boolean"] = 2;
        values[valuesById[3] = "datetime"] = 3;
        values[valuesById[4] = "list"] = 4;
        return values;
    })();

    /**
     * AttributeKind enum.
     * @name user.AttributeKind
     * @enum {string}
     * @property {number} system=0 system value
     * @property {number} default=1 default value
     * @property {number} custom=2 custom value
     */
    user.AttributeKind = (function() {
        const valuesById = {}, values = Object.create(valuesById);
        values[valuesById[0] = "system"] = 0;
        values[valuesById[1] = "default"] = 1;
        values[valuesById[2] = "custom"] = 2;
        return values;
    })();

    user.AttributeDefinition = (function() {

        /**
         * Properties of an AttributeDefinition.
         * @memberof user
         * @interface IAttributeDefinition
         * @property {common.IContext|null} [ctx] AttributeDefinition ctx
         * @property {string|null} [account_id] AttributeDefinition account_id
         * @property {string|null} [name] AttributeDefinition name
         * @property {string|null} [description] AttributeDefinition description
         * @property {string|null} [type] AttributeDefinition type
         * @property {Array.<string>|null} [list_items] AttributeDefinition list_items
         * @property {string|null} [key] AttributeDefinition key
         * @property {string|null} [kind] AttributeDefinition kind
         * @property {number|Long|null} [updated] AttributeDefinition updated
         */

        /**
         * Constructs a new AttributeDefinition.
         * @memberof user
         * @classdesc Represents an AttributeDefinition.
         * @implements IAttributeDefinition
         * @constructor
         * @param {user.IAttributeDefinition=} [p] Properties to set
         */
        function AttributeDefinition(p) {
            this.list_items = [];
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * AttributeDefinition ctx.
         * @member {common.IContext|null|undefined} ctx
         * @memberof user.AttributeDefinition
         * @instance
         */
        AttributeDefinition.prototype.ctx = null;

        /**
         * AttributeDefinition account_id.
         * @member {string} account_id
         * @memberof user.AttributeDefinition
         * @instance
         */
        AttributeDefinition.prototype.account_id = "";

        /**
         * AttributeDefinition name.
         * @member {string} name
         * @memberof user.AttributeDefinition
         * @instance
         */
        AttributeDefinition.prototype.name = "";

        /**
         * AttributeDefinition description.
         * @member {string} description
         * @memberof user.AttributeDefinition
         * @instance
         */
        AttributeDefinition.prototype.description = "";

        /**
         * AttributeDefinition type.
         * @member {string} type
         * @memberof user.AttributeDefinition
         * @instance
         */
        AttributeDefinition.prototype.type = "";

        /**
         * AttributeDefinition list_items.
         * @member {Array.<string>} list_items
         * @memberof user.AttributeDefinition
         * @instance
         */
        AttributeDefinition.prototype.list_items = $util.emptyArray;

        /**
         * AttributeDefinition key.
         * @member {string} key
         * @memberof user.AttributeDefinition
         * @instance
         */
        AttributeDefinition.prototype.key = "";

        /**
         * AttributeDefinition kind.
         * @member {string} kind
         * @memberof user.AttributeDefinition
         * @instance
         */
        AttributeDefinition.prototype.kind = "";

        /**
         * AttributeDefinition updated.
         * @member {number|Long} updated
         * @memberof user.AttributeDefinition
         * @instance
         */
        AttributeDefinition.prototype.updated = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

        return AttributeDefinition;
    })();

    user.AttributeDefinitions = (function() {

        /**
         * Properties of an AttributeDefinitions.
         * @memberof user
         * @interface IAttributeDefinitions
         * @property {common.IContext|null} [ctx] AttributeDefinitions ctx
         * @property {Array.<user.IAttributeDefinition>|null} [attributes] AttributeDefinitions attributes
         */

        /**
         * Constructs a new AttributeDefinitions.
         * @memberof user
         * @classdesc Represents an AttributeDefinitions.
         * @implements IAttributeDefinitions
         * @constructor
         * @param {user.IAttributeDefinitions=} [p] Properties to set
         */
        function AttributeDefinitions(p) {
            this.attributes = [];
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * AttributeDefinitions ctx.
         * @member {common.IContext|null|undefined} ctx
         * @memberof user.AttributeDefinitions
         * @instance
         */
        AttributeDefinitions.prototype.ctx = null;

        /**
         * AttributeDefinitions attributes.
         * @member {Array.<user.IAttributeDefinition>} attributes
         * @memberof user.AttributeDefinitions
         * @instance
         */
        AttributeDefinitions.prototype.attributes = $util.emptyArray;

        return AttributeDefinitions;
    })();

    /**
     * AttributeDataState enum.
     * @name user.AttributeDataState
     * @enum {string}
     * @property {number} live=0 live value
     * @property {number} deleted=1 deleted value
     */
    user.AttributeDataState = (function() {
        const valuesById = {}, values = Object.create(valuesById);
        values[valuesById[0] = "live"] = 0;
        values[valuesById[1] = "deleted"] = 1;
        return values;
    })();

    user.AttributeData = (function() {

        /**
         * Properties of an AttributeData.
         * @memberof user
         * @interface IAttributeData
         * @property {common.IContext|null} [ctx] AttributeData ctx
         * @property {string|null} [account_id] AttributeData account_id
         * @property {string|null} [user_id] AttributeData user_id
         * @property {string|null} [key] AttributeData key
         * @property {string|null} [value] AttributeData value
         * @property {string|null} [state] AttributeData state
         * @property {number|Long|null} [created] AttributeData created
         * @property {number|Long|null} [modified] AttributeData modified
         */

        /**
         * Constructs a new AttributeData.
         * @memberof user
         * @classdesc Represents an AttributeData.
         * @implements IAttributeData
         * @constructor
         * @param {user.IAttributeData=} [p] Properties to set
         */
        function AttributeData(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * AttributeData ctx.
         * @member {common.IContext|null|undefined} ctx
         * @memberof user.AttributeData
         * @instance
         */
        AttributeData.prototype.ctx = null;

        /**
         * AttributeData account_id.
         * @member {string} account_id
         * @memberof user.AttributeData
         * @instance
         */
        AttributeData.prototype.account_id = "";

        /**
         * AttributeData user_id.
         * @member {string} user_id
         * @memberof user.AttributeData
         * @instance
         */
        AttributeData.prototype.user_id = "";

        /**
         * AttributeData key.
         * @member {string} key
         * @memberof user.AttributeData
         * @instance
         */
        AttributeData.prototype.key = "";

        /**
         * AttributeData value.
         * @member {string} value
         * @memberof user.AttributeData
         * @instance
         */
        AttributeData.prototype.value = "";

        /**
         * AttributeData state.
         * @member {string} state
         * @memberof user.AttributeData
         * @instance
         */
        AttributeData.prototype.state = "";

        /**
         * AttributeData created.
         * @member {number|Long} created
         * @memberof user.AttributeData
         * @instance
         */
        AttributeData.prototype.created = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

        /**
         * AttributeData modified.
         * @member {number|Long} modified
         * @memberof user.AttributeData
         * @instance
         */
        AttributeData.prototype.modified = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

        return AttributeData;
    })();

    user.User = (function() {

        /**
         * Properties of a User.
         * @memberof user
         * @interface IUser
         * @property {common.IContext|null} [ctx] User ctx
         * @property {string|null} [id] User id
         * @property {string|null} [account_id] User account_id
         * @property {string|null} [fullname] User fullname
         * @property {Array.<string>|null} [phones] User phones
         * @property {Array.<string>|null} [emails] User emails
         * @property {Array.<user.ITrace>|null} [traces] User traces
         * @property {Array.<user.IDevice>|null} [devices] User devices
         * @property {boolean|null} [is_ban] User is_ban
         * @property {string|null} [avatar_url] User avatar_url
         * @property {Array.<user.IAttributeData>|null} [attributes] User attributes
         * @property {Array.<string>|null} [segments] User segments
         * @property {Array.<string>|null} [labels] User labels
         * @property {boolean|null} [unsubscribed] User unsubscribed
         * @property {boolean|null} [marked_spam] User marked_spam
         * @property {boolean|null} [hard_bounced] User hard_bounced
         * @property {number|null} [total_sessions] User total_sessions
         * @property {string|null} [subiz_id] User subiz_id
         * @property {string|null} [timezone] User timezone
         * @property {string|null} [country_code] User country_code
         * @property {string|null} [city] User city
         * @property {string|null} [language] User language
         * @property {Array.<string>|null} [aliases] User aliases
         * @property {number|Long|null} [seen] User seen
         * @property {Array.<user.IField>|null} [fields] User fields
         * @property {number|null} [par] User par
         * @property {number|Long|null} [modified] User modified
         * @property {number|null} [modified_week] User modified_week
         * @property {number|Long|null} [activated] User activated
         */

        /**
         * Constructs a new User.
         * @memberof user
         * @classdesc Represents a User.
         * @implements IUser
         * @constructor
         * @param {user.IUser=} [p] Properties to set
         */
        function User(p) {
            this.phones = [];
            this.emails = [];
            this.traces = [];
            this.devices = [];
            this.attributes = [];
            this.segments = [];
            this.labels = [];
            this.aliases = [];
            this.fields = [];
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * User ctx.
         * @member {common.IContext|null|undefined} ctx
         * @memberof user.User
         * @instance
         */
        User.prototype.ctx = null;

        /**
         * User id.
         * @member {string} id
         * @memberof user.User
         * @instance
         */
        User.prototype.id = "";

        /**
         * User account_id.
         * @member {string} account_id
         * @memberof user.User
         * @instance
         */
        User.prototype.account_id = "";

        /**
         * User fullname.
         * @member {string} fullname
         * @memberof user.User
         * @instance
         */
        User.prototype.fullname = "";

        /**
         * User phones.
         * @member {Array.<string>} phones
         * @memberof user.User
         * @instance
         */
        User.prototype.phones = $util.emptyArray;

        /**
         * User emails.
         * @member {Array.<string>} emails
         * @memberof user.User
         * @instance
         */
        User.prototype.emails = $util.emptyArray;

        /**
         * User traces.
         * @member {Array.<user.ITrace>} traces
         * @memberof user.User
         * @instance
         */
        User.prototype.traces = $util.emptyArray;

        /**
         * User devices.
         * @member {Array.<user.IDevice>} devices
         * @memberof user.User
         * @instance
         */
        User.prototype.devices = $util.emptyArray;

        /**
         * User is_ban.
         * @member {boolean} is_ban
         * @memberof user.User
         * @instance
         */
        User.prototype.is_ban = false;

        /**
         * User avatar_url.
         * @member {string} avatar_url
         * @memberof user.User
         * @instance
         */
        User.prototype.avatar_url = "";

        /**
         * User attributes.
         * @member {Array.<user.IAttributeData>} attributes
         * @memberof user.User
         * @instance
         */
        User.prototype.attributes = $util.emptyArray;

        /**
         * User segments.
         * @member {Array.<string>} segments
         * @memberof user.User
         * @instance
         */
        User.prototype.segments = $util.emptyArray;

        /**
         * User labels.
         * @member {Array.<string>} labels
         * @memberof user.User
         * @instance
         */
        User.prototype.labels = $util.emptyArray;

        /**
         * User unsubscribed.
         * @member {boolean} unsubscribed
         * @memberof user.User
         * @instance
         */
        User.prototype.unsubscribed = false;

        /**
         * User marked_spam.
         * @member {boolean} marked_spam
         * @memberof user.User
         * @instance
         */
        User.prototype.marked_spam = false;

        /**
         * User hard_bounced.
         * @member {boolean} hard_bounced
         * @memberof user.User
         * @instance
         */
        User.prototype.hard_bounced = false;

        /**
         * User total_sessions.
         * @member {number} total_sessions
         * @memberof user.User
         * @instance
         */
        User.prototype.total_sessions = 0;

        /**
         * User subiz_id.
         * @member {string} subiz_id
         * @memberof user.User
         * @instance
         */
        User.prototype.subiz_id = "";

        /**
         * User timezone.
         * @member {string} timezone
         * @memberof user.User
         * @instance
         */
        User.prototype.timezone = "";

        /**
         * User country_code.
         * @member {string} country_code
         * @memberof user.User
         * @instance
         */
        User.prototype.country_code = "";

        /**
         * User city.
         * @member {string} city
         * @memberof user.User
         * @instance
         */
        User.prototype.city = "";

        /**
         * User language.
         * @member {string} language
         * @memberof user.User
         * @instance
         */
        User.prototype.language = "";

        /**
         * User aliases.
         * @member {Array.<string>} aliases
         * @memberof user.User
         * @instance
         */
        User.prototype.aliases = $util.emptyArray;

        /**
         * User seen.
         * @member {number|Long} seen
         * @memberof user.User
         * @instance
         */
        User.prototype.seen = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

        /**
         * User fields.
         * @member {Array.<user.IField>} fields
         * @memberof user.User
         * @instance
         */
        User.prototype.fields = $util.emptyArray;

        /**
         * User par.
         * @member {number} par
         * @memberof user.User
         * @instance
         */
        User.prototype.par = 0;

        /**
         * User modified.
         * @member {number|Long} modified
         * @memberof user.User
         * @instance
         */
        User.prototype.modified = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

        /**
         * User modified_week.
         * @member {number} modified_week
         * @memberof user.User
         * @instance
         */
        User.prototype.modified_week = 0;

        /**
         * User activated.
         * @member {number|Long} activated
         * @memberof user.User
         * @instance
         */
        User.prototype.activated = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

        return User;
    })();

    user.Users = (function() {

        /**
         * Properties of a Users.
         * @memberof user
         * @interface IUsers
         * @property {Array.<user.IUser>|null} [users] Users users
         */

        /**
         * Constructs a new Users.
         * @memberof user
         * @classdesc Represents a Users.
         * @implements IUsers
         * @constructor
         * @param {user.IUsers=} [p] Properties to set
         */
        function Users(p) {
            this.users = [];
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * Users users.
         * @member {Array.<user.IUser>} users
         * @memberof user.Users
         * @instance
         */
        Users.prototype.users = $util.emptyArray;

        return Users;
    })();

    user.Device = (function() {

        /**
         * Properties of a Device.
         * @memberof user
         * @interface IDevice
         * @property {number|null} [id] Device id
         * @property {number|null} [useragent_id] Device useragent_id
         * @property {string|null} [useragent] Device useragent
         * @property {string|null} [screen_resolution] Device screen_resolution
         * @property {number|null} [language_id] Device language_id
         * @property {string|null} [language] Device language
         */

        /**
         * Constructs a new Device.
         * @memberof user
         * @classdesc Represents a Device.
         * @implements IDevice
         * @constructor
         * @param {user.IDevice=} [p] Properties to set
         */
        function Device(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * Device id.
         * @member {number} id
         * @memberof user.Device
         * @instance
         */
        Device.prototype.id = 0;

        /**
         * Device useragent_id.
         * @member {number} useragent_id
         * @memberof user.Device
         * @instance
         */
        Device.prototype.useragent_id = 0;

        /**
         * Device useragent.
         * @member {string} useragent
         * @memberof user.Device
         * @instance
         */
        Device.prototype.useragent = "";

        /**
         * Device screen_resolution.
         * @member {string} screen_resolution
         * @memberof user.Device
         * @instance
         */
        Device.prototype.screen_resolution = "";

        /**
         * Device language_id.
         * @member {number} language_id
         * @memberof user.Device
         * @instance
         */
        Device.prototype.language_id = 0;

        /**
         * Device language.
         * @member {string} language
         * @memberof user.Device
         * @instance
         */
        Device.prototype.language = "";

        return Device;
    })();

    user.Trace = (function() {

        /**
         * Properties of a Trace.
         * @memberof user
         * @interface ITrace
         * @property {string|null} [id] Trace id
         * @property {string|null} [ip] Trace ip
         * @property {number|null} [location_id] Trace location_id
         * @property {string|null} [city_name] Trace city_name
         * @property {string|null} [country_name] Trace country_name
         * @property {string|null} [country_code] Trace country_code
         * @property {string|null} [continent_code] Trace continent_code
         * @property {string|null} [continent_name] Trace continent_name
         * @property {number|null} [latitude] Trace latitude
         * @property {number|null} [longitude] Trace longitude
         * @property {string|null} [postal_code] Trace postal_code
         * @property {string|null} [timezone] Trace timezone
         * @property {string|null} [isp] Trace isp
         */

        /**
         * Constructs a new Trace.
         * @memberof user
         * @classdesc Represents a Trace.
         * @implements ITrace
         * @constructor
         * @param {user.ITrace=} [p] Properties to set
         */
        function Trace(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * Trace id.
         * @member {string} id
         * @memberof user.Trace
         * @instance
         */
        Trace.prototype.id = "";

        /**
         * Trace ip.
         * @member {string} ip
         * @memberof user.Trace
         * @instance
         */
        Trace.prototype.ip = "";

        /**
         * Trace location_id.
         * @member {number} location_id
         * @memberof user.Trace
         * @instance
         */
        Trace.prototype.location_id = 0;

        /**
         * Trace city_name.
         * @member {string} city_name
         * @memberof user.Trace
         * @instance
         */
        Trace.prototype.city_name = "";

        /**
         * Trace country_name.
         * @member {string} country_name
         * @memberof user.Trace
         * @instance
         */
        Trace.prototype.country_name = "";

        /**
         * Trace country_code.
         * @member {string} country_code
         * @memberof user.Trace
         * @instance
         */
        Trace.prototype.country_code = "";

        /**
         * Trace continent_code.
         * @member {string} continent_code
         * @memberof user.Trace
         * @instance
         */
        Trace.prototype.continent_code = "";

        /**
         * Trace continent_name.
         * @member {string} continent_name
         * @memberof user.Trace
         * @instance
         */
        Trace.prototype.continent_name = "";

        /**
         * Trace latitude.
         * @member {number} latitude
         * @memberof user.Trace
         * @instance
         */
        Trace.prototype.latitude = 0;

        /**
         * Trace longitude.
         * @member {number} longitude
         * @memberof user.Trace
         * @instance
         */
        Trace.prototype.longitude = 0;

        /**
         * Trace postal_code.
         * @member {string} postal_code
         * @memberof user.Trace
         * @instance
         */
        Trace.prototype.postal_code = "";

        /**
         * Trace timezone.
         * @member {string} timezone
         * @memberof user.Trace
         * @instance
         */
        Trace.prototype.timezone = "";

        /**
         * Trace isp.
         * @member {string} isp
         * @memberof user.Trace
         * @instance
         */
        Trace.prototype.isp = "";

        return Trace;
    })();

    user.MergeRequest = (function() {

        /**
         * Properties of a MergeRequest.
         * @memberof user
         * @interface IMergeRequest
         * @property {common.IContext|null} [ctx] MergeRequest ctx
         * @property {string|null} [id] MergeRequest id
         * @property {string|null} [recent_id] MergeRequest recent_id
         */

        /**
         * Constructs a new MergeRequest.
         * @memberof user
         * @classdesc Represents a MergeRequest.
         * @implements IMergeRequest
         * @constructor
         * @param {user.IMergeRequest=} [p] Properties to set
         */
        function MergeRequest(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * MergeRequest ctx.
         * @member {common.IContext|null|undefined} ctx
         * @memberof user.MergeRequest
         * @instance
         */
        MergeRequest.prototype.ctx = null;

        /**
         * MergeRequest id.
         * @member {string} id
         * @memberof user.MergeRequest
         * @instance
         */
        MergeRequest.prototype.id = "";

        /**
         * MergeRequest recent_id.
         * @member {string} recent_id
         * @memberof user.MergeRequest
         * @instance
         */
        MergeRequest.prototype.recent_id = "";

        return MergeRequest;
    })();

    user.CreateRequest = (function() {

        /**
         * Properties of a CreateRequest.
         * @memberof user
         * @interface ICreateRequest
         * @property {string|null} [challenge_id] CreateRequest challenge_id
         * @property {string|null} [answer] CreateRequest answer
         */

        /**
         * Constructs a new CreateRequest.
         * @memberof user
         * @classdesc Represents a CreateRequest.
         * @implements ICreateRequest
         * @constructor
         * @param {user.ICreateRequest=} [p] Properties to set
         */
        function CreateRequest(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * CreateRequest challenge_id.
         * @member {string} challenge_id
         * @memberof user.CreateRequest
         * @instance
         */
        CreateRequest.prototype.challenge_id = "";

        /**
         * CreateRequest answer.
         * @member {string} answer
         * @memberof user.CreateRequest
         * @instance
         */
        CreateRequest.prototype.answer = "";

        return CreateRequest;
    })();

    user.Topic = (function() {

        /**
         * Properties of a Topic.
         * @memberof user
         * @interface ITopic
         * @property {common.IContext|null} [ctx] Topic ctx
         * @property {string|null} [account_id] Topic account_id
         * @property {string|null} [topic] Topic topic
         * @property {string|null} [type] Topic type
         * @property {number|Long|null} [updated] Topic updated
         * @property {number|null} [unread] Topic unread
         */

        /**
         * Constructs a new Topic.
         * @memberof user
         * @classdesc Represents a Topic.
         * @implements ITopic
         * @constructor
         * @param {user.ITopic=} [p] Properties to set
         */
        function Topic(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * Topic ctx.
         * @member {common.IContext|null|undefined} ctx
         * @memberof user.Topic
         * @instance
         */
        Topic.prototype.ctx = null;

        /**
         * Topic account_id.
         * @member {string} account_id
         * @memberof user.Topic
         * @instance
         */
        Topic.prototype.account_id = "";

        /**
         * Topic topic.
         * @member {string} topic
         * @memberof user.Topic
         * @instance
         */
        Topic.prototype.topic = "";

        /**
         * Topic type.
         * @member {string} type
         * @memberof user.Topic
         * @instance
         */
        Topic.prototype.type = "";

        /**
         * Topic updated.
         * @member {number|Long} updated
         * @memberof user.Topic
         * @instance
         */
        Topic.prototype.updated = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

        /**
         * Topic unread.
         * @member {number} unread
         * @memberof user.Topic
         * @instance
         */
        Topic.prototype.unread = 0;

        return Topic;
    })();

    user.UnreadTopic = (function() {

        /**
         * Properties of an UnreadTopic.
         * @memberof user
         * @interface IUnreadTopic
         * @property {common.IContext|null} [ctx] UnreadTopic ctx
         * @property {string|null} [topic] UnreadTopic topic
         * @property {string|null} [agent_id] UnreadTopic agent_id
         * @property {string|null} [user_id] UnreadTopic user_id
         * @property {string|null} [type] UnreadTopic type
         * @property {number|Long|null} [updated] UnreadTopic updated
         * @property {number|null} [unread] UnreadTopic unread
         */

        /**
         * Constructs a new UnreadTopic.
         * @memberof user
         * @classdesc Represents an UnreadTopic.
         * @implements IUnreadTopic
         * @constructor
         * @param {user.IUnreadTopic=} [p] Properties to set
         */
        function UnreadTopic(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * UnreadTopic ctx.
         * @member {common.IContext|null|undefined} ctx
         * @memberof user.UnreadTopic
         * @instance
         */
        UnreadTopic.prototype.ctx = null;

        /**
         * UnreadTopic topic.
         * @member {string} topic
         * @memberof user.UnreadTopic
         * @instance
         */
        UnreadTopic.prototype.topic = "";

        /**
         * UnreadTopic agent_id.
         * @member {string} agent_id
         * @memberof user.UnreadTopic
         * @instance
         */
        UnreadTopic.prototype.agent_id = "";

        /**
         * UnreadTopic user_id.
         * @member {string} user_id
         * @memberof user.UnreadTopic
         * @instance
         */
        UnreadTopic.prototype.user_id = "";

        /**
         * UnreadTopic type.
         * @member {string} type
         * @memberof user.UnreadTopic
         * @instance
         */
        UnreadTopic.prototype.type = "";

        /**
         * UnreadTopic updated.
         * @member {number|Long} updated
         * @memberof user.UnreadTopic
         * @instance
         */
        UnreadTopic.prototype.updated = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

        /**
         * UnreadTopic unread.
         * @member {number} unread
         * @memberof user.UnreadTopic
         * @instance
         */
        UnreadTopic.prototype.unread = 0;

        return UnreadTopic;
    })();

    user.ReadTopicRequest = (function() {

        /**
         * Properties of a ReadTopicRequest.
         * @memberof user
         * @interface IReadTopicRequest
         * @property {common.IContext|null} [ctx] ReadTopicRequest ctx
         * @property {string|null} [topic] ReadTopicRequest topic
         * @property {string|null} [user_id] ReadTopicRequest user_id
         * @property {string|null} [agent_id] ReadTopicRequest agent_id
         */

        /**
         * Constructs a new ReadTopicRequest.
         * @memberof user
         * @classdesc Represents a ReadTopicRequest.
         * @implements IReadTopicRequest
         * @constructor
         * @param {user.IReadTopicRequest=} [p] Properties to set
         */
        function ReadTopicRequest(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * ReadTopicRequest ctx.
         * @member {common.IContext|null|undefined} ctx
         * @memberof user.ReadTopicRequest
         * @instance
         */
        ReadTopicRequest.prototype.ctx = null;

        /**
         * ReadTopicRequest topic.
         * @member {string} topic
         * @memberof user.ReadTopicRequest
         * @instance
         */
        ReadTopicRequest.prototype.topic = "";

        /**
         * ReadTopicRequest user_id.
         * @member {string} user_id
         * @memberof user.ReadTopicRequest
         * @instance
         */
        ReadTopicRequest.prototype.user_id = "";

        /**
         * ReadTopicRequest agent_id.
         * @member {string} agent_id
         * @memberof user.ReadTopicRequest
         * @instance
         */
        ReadTopicRequest.prototype.agent_id = "";

        return ReadTopicRequest;
    })();

    user.SubscribeRequest = (function() {

        /**
         * Properties of a SubscribeRequest.
         * @memberof user
         * @interface ISubscribeRequest
         * @property {common.IContext|null} [ctx] SubscribeRequest ctx
         * @property {string|null} [agent_id] SubscribeRequest agent_id
         * @property {Array.<string>|null} [topics] SubscribeRequest topics
         */

        /**
         * Constructs a new SubscribeRequest.
         * @memberof user
         * @classdesc Represents a SubscribeRequest.
         * @implements ISubscribeRequest
         * @constructor
         * @param {user.ISubscribeRequest=} [p] Properties to set
         */
        function SubscribeRequest(p) {
            this.topics = [];
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * SubscribeRequest ctx.
         * @member {common.IContext|null|undefined} ctx
         * @memberof user.SubscribeRequest
         * @instance
         */
        SubscribeRequest.prototype.ctx = null;

        /**
         * SubscribeRequest agent_id.
         * @member {string} agent_id
         * @memberof user.SubscribeRequest
         * @instance
         */
        SubscribeRequest.prototype.agent_id = "";

        /**
         * SubscribeRequest topics.
         * @member {Array.<string>} topics
         * @memberof user.SubscribeRequest
         * @instance
         */
        SubscribeRequest.prototype.topics = $util.emptyArray;

        return SubscribeRequest;
    })();

    /**
     * Event enum.
     * @name user.Event
     * @enum {string}
     * @property {number} UserReadRequested=0 UserReadRequested value
     * @property {number} UserUpdateRequested=2 UserUpdateRequested value
     * @property {number} UserCreateRequested=3 UserCreateRequested value
     * @property {number} UserSearchRequested=4 UserSearchRequested value
     * @property {number} UserEventCreateRequested=5 UserEventCreateRequested value
     * @property {number} UserEventSearchRequested=7 UserEventSearchRequested value
     * @property {number} UserTopicSearchRequested=6 UserTopicSearchRequested value
     * @property {number} UserSegmentationCreateRequested=10 UserSegmentationCreateRequested value
     * @property {number} UserSegmentationUpdateRequested=11 UserSegmentationUpdateRequested value
     * @property {number} UserSegmentationDeleteRequested=12 UserSegmentationDeleteRequested value
     * @property {number} UserSegmentationListRequested=13 UserSegmentationListRequested value
     * @property {number} UserSegmentationReadRequested=14 UserSegmentationReadRequested value
     * @property {number} UserAddToMyListRequested=20 UserAddToMyListRequested value
     * @property {number} UserEventTopicSubscribeRequested=35 UserEventTopicSubscribeRequested value
     * @property {number} UserEventTopicUnsubscribeRequested=36 UserEventTopicUnsubscribeRequested value
     * @property {number} UserReadTopicRequested=41 UserReadTopicRequested value
     * @property {number} UserSubizId=42 UserSubizId value
     * @property {number} UserPresenceReadRequested=44 UserPresenceReadRequested value
     * @property {number} UserPreviewingReadRequested=46 UserPreviewingReadRequested value
     * @property {number} UserListTopRequested=47 UserListTopRequested value
     * @property {number} UserAutomationUpsertRequested=50 UserAutomationUpsertRequested value
     * @property {number} UserAutomationDeleteRequested=51 UserAutomationDeleteRequested value
     * @property {number} UserAutomationListRequested=52 UserAutomationListRequested value
     * @property {number} UserAutomationReadRequested=53 UserAutomationReadRequested value
     * @property {number} AutomationAgentNotificationFired=54 AutomationAgentNotificationFired value
     * @property {number} AutomationConversationMessageFired=55 AutomationConversationMessageFired value
     * @property {number} UserSessionUpdateRequested=65 UserSessionUpdateRequested value
     * @property {number} UserSessionCreateRequested=66 UserSessionCreateRequested value
     * @property {number} UserSessionReadRequested=67 UserSessionReadRequested value
     * @property {number} SegmentationLoop=68 SegmentationLoop value
     * @property {number} AutomationSynced=102 AutomationSynced value
     * @property {number} AutomationFired=103 AutomationFired value
     * @property {number} UserRequested=100 UserRequested value
     * @property {number} UserSynced=101 UserSynced value
     * @property {number} UserUpserted=105 UserUpserted value
     * @property {number} UserV3Synced=106 UserV3Synced value
     */
    user.Event = (function() {
        const valuesById = {}, values = Object.create(valuesById);
        values[valuesById[0] = "UserReadRequested"] = 0;
        values[valuesById[2] = "UserUpdateRequested"] = 2;
        values[valuesById[3] = "UserCreateRequested"] = 3;
        values[valuesById[4] = "UserSearchRequested"] = 4;
        values[valuesById[5] = "UserEventCreateRequested"] = 5;
        values[valuesById[7] = "UserEventSearchRequested"] = 7;
        values[valuesById[6] = "UserTopicSearchRequested"] = 6;
        values[valuesById[10] = "UserSegmentationCreateRequested"] = 10;
        values[valuesById[11] = "UserSegmentationUpdateRequested"] = 11;
        values[valuesById[12] = "UserSegmentationDeleteRequested"] = 12;
        values[valuesById[13] = "UserSegmentationListRequested"] = 13;
        values[valuesById[14] = "UserSegmentationReadRequested"] = 14;
        values[valuesById[20] = "UserAddToMyListRequested"] = 20;
        values[valuesById[35] = "UserEventTopicSubscribeRequested"] = 35;
        values[valuesById[36] = "UserEventTopicUnsubscribeRequested"] = 36;
        values[valuesById[41] = "UserReadTopicRequested"] = 41;
        values[valuesById[42] = "UserSubizId"] = 42;
        values[valuesById[44] = "UserPresenceReadRequested"] = 44;
        values[valuesById[46] = "UserPreviewingReadRequested"] = 46;
        values[valuesById[47] = "UserListTopRequested"] = 47;
        values[valuesById[50] = "UserAutomationUpsertRequested"] = 50;
        values[valuesById[51] = "UserAutomationDeleteRequested"] = 51;
        values[valuesById[52] = "UserAutomationListRequested"] = 52;
        values[valuesById[53] = "UserAutomationReadRequested"] = 53;
        values[valuesById[54] = "AutomationAgentNotificationFired"] = 54;
        values[valuesById[55] = "AutomationConversationMessageFired"] = 55;
        values[valuesById[65] = "UserSessionUpdateRequested"] = 65;
        values[valuesById[66] = "UserSessionCreateRequested"] = 66;
        values[valuesById[67] = "UserSessionReadRequested"] = 67;
        values[valuesById[68] = "SegmentationLoop"] = 68;
        values[valuesById[102] = "AutomationSynced"] = 102;
        values[valuesById[103] = "AutomationFired"] = 103;
        values[valuesById[100] = "UserRequested"] = 100;
        values[valuesById[101] = "UserSynced"] = 101;
        values[valuesById[105] = "UserUpserted"] = 105;
        values[valuesById[106] = "UserV3Synced"] = 106;
        return values;
    })();

    user.SubizIDRequest = (function() {

        /**
         * Properties of a SubizIDRequest.
         * @memberof user
         * @interface ISubizIDRequest
         * @property {common.IContext|null} [ctx] SubizIDRequest ctx
         * @property {string|null} [subiz_id] SubizIDRequest subiz_id
         * @property {string|null} [account_id] SubizIDRequest account_id
         */

        /**
         * Constructs a new SubizIDRequest.
         * @memberof user
         * @classdesc Represents a SubizIDRequest.
         * @implements ISubizIDRequest
         * @constructor
         * @param {user.ISubizIDRequest=} [p] Properties to set
         */
        function SubizIDRequest(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * SubizIDRequest ctx.
         * @member {common.IContext|null|undefined} ctx
         * @memberof user.SubizIDRequest
         * @instance
         */
        SubizIDRequest.prototype.ctx = null;

        /**
         * SubizIDRequest subiz_id.
         * @member {string} subiz_id
         * @memberof user.SubizIDRequest
         * @instance
         */
        SubizIDRequest.prototype.subiz_id = "";

        /**
         * SubizIDRequest account_id.
         * @member {string} account_id
         * @memberof user.SubizIDRequest
         * @instance
         */
        SubizIDRequest.prototype.account_id = "";

        return SubizIDRequest;
    })();

    user.MaskResponse = (function() {

        /**
         * Properties of a MaskResponse.
         * @memberof user
         * @interface IMaskResponse
         * @property {string|null} [subiz_id] MaskResponse subiz_id
         * @property {string|null} [account_id] MaskResponse account_id
         * @property {string|null} [user_id] MaskResponse user_id
         * @property {string|null} [mask] MaskResponse mask
         */

        /**
         * Constructs a new MaskResponse.
         * @memberof user
         * @classdesc Represents a MaskResponse.
         * @implements IMaskResponse
         * @constructor
         * @param {user.IMaskResponse=} [p] Properties to set
         */
        function MaskResponse(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * MaskResponse subiz_id.
         * @member {string} subiz_id
         * @memberof user.MaskResponse
         * @instance
         */
        MaskResponse.prototype.subiz_id = "";

        /**
         * MaskResponse account_id.
         * @member {string} account_id
         * @memberof user.MaskResponse
         * @instance
         */
        MaskResponse.prototype.account_id = "";

        /**
         * MaskResponse user_id.
         * @member {string} user_id
         * @memberof user.MaskResponse
         * @instance
         */
        MaskResponse.prototype.user_id = "";

        /**
         * MaskResponse mask.
         * @member {string} mask
         * @memberof user.MaskResponse
         * @instance
         */
        MaskResponse.prototype.mask = "";

        return MaskResponse;
    })();

    user.SubizIDResponse = (function() {

        /**
         * Properties of a SubizIDResponse.
         * @memberof user
         * @interface ISubizIDResponse
         * @property {common.IContext|null} [ctx] SubizIDResponse ctx
         * @property {string|null} [subiz_id] SubizIDResponse subiz_id
         * @property {string|null} [account_id] SubizIDResponse account_id
         * @property {string|null} [user_id] SubizIDResponse user_id
         */

        /**
         * Constructs a new SubizIDResponse.
         * @memberof user
         * @classdesc Represents a SubizIDResponse.
         * @implements ISubizIDResponse
         * @constructor
         * @param {user.ISubizIDResponse=} [p] Properties to set
         */
        function SubizIDResponse(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * SubizIDResponse ctx.
         * @member {common.IContext|null|undefined} ctx
         * @memberof user.SubizIDResponse
         * @instance
         */
        SubizIDResponse.prototype.ctx = null;

        /**
         * SubizIDResponse subiz_id.
         * @member {string} subiz_id
         * @memberof user.SubizIDResponse
         * @instance
         */
        SubizIDResponse.prototype.subiz_id = "";

        /**
         * SubizIDResponse account_id.
         * @member {string} account_id
         * @memberof user.SubizIDResponse
         * @instance
         */
        SubizIDResponse.prototype.account_id = "";

        /**
         * SubizIDResponse user_id.
         * @member {string} user_id
         * @memberof user.SubizIDResponse
         * @instance
         */
        SubizIDResponse.prototype.user_id = "";

        return SubizIDResponse;
    })();

    user.Segmentations = (function() {

        /**
         * Properties of a Segmentations.
         * @memberof user
         * @interface ISegmentations
         * @property {common.IContext|null} [ctx] Segmentations ctx
         * @property {Array.<user.ISegmentation>|null} [segmentations] Segmentations segmentations
         */

        /**
         * Constructs a new Segmentations.
         * @memberof user
         * @classdesc Represents a Segmentations.
         * @implements ISegmentations
         * @constructor
         * @param {user.ISegmentations=} [p] Properties to set
         */
        function Segmentations(p) {
            this.segmentations = [];
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * Segmentations ctx.
         * @member {common.IContext|null|undefined} ctx
         * @memberof user.Segmentations
         * @instance
         */
        Segmentations.prototype.ctx = null;

        /**
         * Segmentations segmentations.
         * @member {Array.<user.ISegmentation>} segmentations
         * @memberof user.Segmentations
         * @instance
         */
        Segmentations.prototype.segmentations = $util.emptyArray;

        return Segmentations;
    })();

    user.SegmentLoopState = (function() {

        /**
         * Properties of a SegmentLoopState.
         * @memberof user
         * @interface ISegmentLoopState
         * @property {common.IContext|null} [ctx] SegmentLoopState ctx
         * @property {string|null} [account_id] SegmentLoopState account_id
         * @property {number|null} [user_par] SegmentLoopState user_par
         * @property {number|Long|null} [loop_created] SegmentLoopState loop_created
         * @property {number|Long|null} [loop_number] SegmentLoopState loop_number
         */

        /**
         * Constructs a new SegmentLoopState.
         * @memberof user
         * @classdesc Represents a SegmentLoopState.
         * @implements ISegmentLoopState
         * @constructor
         * @param {user.ISegmentLoopState=} [p] Properties to set
         */
        function SegmentLoopState(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * SegmentLoopState ctx.
         * @member {common.IContext|null|undefined} ctx
         * @memberof user.SegmentLoopState
         * @instance
         */
        SegmentLoopState.prototype.ctx = null;

        /**
         * SegmentLoopState account_id.
         * @member {string} account_id
         * @memberof user.SegmentLoopState
         * @instance
         */
        SegmentLoopState.prototype.account_id = "";

        /**
         * SegmentLoopState user_par.
         * @member {number} user_par
         * @memberof user.SegmentLoopState
         * @instance
         */
        SegmentLoopState.prototype.user_par = 0;

        /**
         * SegmentLoopState loop_created.
         * @member {number|Long} loop_created
         * @memberof user.SegmentLoopState
         * @instance
         */
        SegmentLoopState.prototype.loop_created = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

        /**
         * SegmentLoopState loop_number.
         * @member {number|Long} loop_number
         * @memberof user.SegmentLoopState
         * @instance
         */
        SegmentLoopState.prototype.loop_number = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

        return SegmentLoopState;
    })();

    user.UserSegmentCache = (function() {

        /**
         * Properties of a UserSegmentCache.
         * @memberof user
         * @interface IUserSegmentCache
         * @property {common.IContext|null} [ctx] UserSegmentCache ctx
         * @property {string|null} [account_id] UserSegmentCache account_id
         * @property {string|null} [id] UserSegmentCache id
         * @property {string|null} [condition_id] UserSegmentCache condition_id
         */

        /**
         * Constructs a new UserSegmentCache.
         * @memberof user
         * @classdesc Represents a UserSegmentCache.
         * @implements IUserSegmentCache
         * @constructor
         * @param {user.IUserSegmentCache=} [p] Properties to set
         */
        function UserSegmentCache(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * UserSegmentCache ctx.
         * @member {common.IContext|null|undefined} ctx
         * @memberof user.UserSegmentCache
         * @instance
         */
        UserSegmentCache.prototype.ctx = null;

        /**
         * UserSegmentCache account_id.
         * @member {string} account_id
         * @memberof user.UserSegmentCache
         * @instance
         */
        UserSegmentCache.prototype.account_id = "";

        /**
         * UserSegmentCache id.
         * @member {string} id
         * @memberof user.UserSegmentCache
         * @instance
         */
        UserSegmentCache.prototype.id = "";

        /**
         * UserSegmentCache condition_id.
         * @member {string} condition_id
         * @memberof user.UserSegmentCache
         * @instance
         */
        UserSegmentCache.prototype.condition_id = "";

        return UserSegmentCache;
    })();

    user.Segmentation = (function() {

        /**
         * Properties of a Segmentation.
         * @memberof user
         * @interface ISegmentation
         * @property {common.IContext|null} [ctx] Segmentation ctx
         * @property {string|null} [account_id] Segmentation account_id
         * @property {string|null} [id] Segmentation id
         * @property {string|null} [name] Segmentation name
         * @property {string|null} [description] Segmentation description
         * @property {number|Long|null} [user_count] Segmentation user_count
         * @property {number|Long|null} [ran] Segmentation ran
         * @property {number|Long|null} [started_from] Segmentation started_from
         * @property {number|Long|null} [created] Segmentation created
         * @property {number|Long|null} [modified] Segmentation modified
         * @property {string|null} [state] Segmentation state
         * @property {user.ISegmentCondition|null} [condition] Segmentation condition
         * @property {string|null} [current_cursor] Segmentation current_cursor
         */

        /**
         * Constructs a new Segmentation.
         * @memberof user
         * @classdesc Represents a Segmentation.
         * @implements ISegmentation
         * @constructor
         * @param {user.ISegmentation=} [p] Properties to set
         */
        function Segmentation(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * Segmentation ctx.
         * @member {common.IContext|null|undefined} ctx
         * @memberof user.Segmentation
         * @instance
         */
        Segmentation.prototype.ctx = null;

        /**
         * Segmentation account_id.
         * @member {string} account_id
         * @memberof user.Segmentation
         * @instance
         */
        Segmentation.prototype.account_id = "";

        /**
         * Segmentation id.
         * @member {string} id
         * @memberof user.Segmentation
         * @instance
         */
        Segmentation.prototype.id = "";

        /**
         * Segmentation name.
         * @member {string} name
         * @memberof user.Segmentation
         * @instance
         */
        Segmentation.prototype.name = "";

        /**
         * Segmentation description.
         * @member {string} description
         * @memberof user.Segmentation
         * @instance
         */
        Segmentation.prototype.description = "";

        /**
         * Segmentation user_count.
         * @member {number|Long} user_count
         * @memberof user.Segmentation
         * @instance
         */
        Segmentation.prototype.user_count = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

        /**
         * Segmentation ran.
         * @member {number|Long} ran
         * @memberof user.Segmentation
         * @instance
         */
        Segmentation.prototype.ran = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

        /**
         * Segmentation started_from.
         * @member {number|Long} started_from
         * @memberof user.Segmentation
         * @instance
         */
        Segmentation.prototype.started_from = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

        /**
         * Segmentation created.
         * @member {number|Long} created
         * @memberof user.Segmentation
         * @instance
         */
        Segmentation.prototype.created = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

        /**
         * Segmentation modified.
         * @member {number|Long} modified
         * @memberof user.Segmentation
         * @instance
         */
        Segmentation.prototype.modified = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

        /**
         * Segmentation state.
         * @member {string} state
         * @memberof user.Segmentation
         * @instance
         */
        Segmentation.prototype.state = "";

        /**
         * Segmentation condition.
         * @member {user.ISegmentCondition|null|undefined} condition
         * @memberof user.Segmentation
         * @instance
         */
        Segmentation.prototype.condition = null;

        /**
         * Segmentation current_cursor.
         * @member {string} current_cursor
         * @memberof user.Segmentation
         * @instance
         */
        Segmentation.prototype.current_cursor = "";

        /**
         * State enum.
         * @name user.Segmentation.State
         * @enum {string}
         * @property {number} active=0 active value
         * @property {number} inactive=1 inactive value
         */
        Segmentation.State = (function() {
            const valuesById = {}, values = Object.create(valuesById);
            values[valuesById[0] = "active"] = 0;
            values[valuesById[1] = "inactive"] = 1;
            return values;
        })();

        return Segmentation;
    })();

    user.SegmentTracking = (function() {

        /**
         * Properties of a SegmentTracking.
         * @memberof user
         * @interface ISegmentTracking
         * @property {common.IContext|null} [ctx] SegmentTracking ctx
         * @property {number|null} [user_par] SegmentTracking user_par
         * @property {string|null} [account_id] SegmentTracking account_id
         * @property {number|Long|null} [loop_created] SegmentTracking loop_created
         * @property {number|Long|null} [loop_number] SegmentTracking loop_number
         */

        /**
         * Constructs a new SegmentTracking.
         * @memberof user
         * @classdesc Represents a SegmentTracking.
         * @implements ISegmentTracking
         * @constructor
         * @param {user.ISegmentTracking=} [p] Properties to set
         */
        function SegmentTracking(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * SegmentTracking ctx.
         * @member {common.IContext|null|undefined} ctx
         * @memberof user.SegmentTracking
         * @instance
         */
        SegmentTracking.prototype.ctx = null;

        /**
         * SegmentTracking user_par.
         * @member {number} user_par
         * @memberof user.SegmentTracking
         * @instance
         */
        SegmentTracking.prototype.user_par = 0;

        /**
         * SegmentTracking account_id.
         * @member {string} account_id
         * @memberof user.SegmentTracking
         * @instance
         */
        SegmentTracking.prototype.account_id = "";

        /**
         * SegmentTracking loop_created.
         * @member {number|Long} loop_created
         * @memberof user.SegmentTracking
         * @instance
         */
        SegmentTracking.prototype.loop_created = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

        /**
         * SegmentTracking loop_number.
         * @member {number|Long} loop_number
         * @memberof user.SegmentTracking
         * @instance
         */
        SegmentTracking.prototype.loop_number = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

        return SegmentTracking;
    })();

    user.SegmentCondition = (function() {

        /**
         * Properties of a SegmentCondition.
         * @memberof user
         * @interface ISegmentCondition
         * @property {string|null} [join] SegmentCondition join
         * @property {string|null} [id] SegmentCondition id
         * @property {string|null} [event_type] SegmentCondition event_type
         * @property {string|null} [composition] SegmentCondition composition
         * @property {string|null} [key] SegmentCondition key
         * @property {Array.<user.ICondition>|null} [conditions] SegmentCondition conditions
         * @property {user.ISegmentCondition|null} [left_condition] SegmentCondition left_condition
         * @property {user.ISegmentCondition|null} [right_condition] SegmentCondition right_condition
         */

        /**
         * Constructs a new SegmentCondition.
         * @memberof user
         * @classdesc Represents a SegmentCondition.
         * @implements ISegmentCondition
         * @constructor
         * @param {user.ISegmentCondition=} [p] Properties to set
         */
        function SegmentCondition(p) {
            this.conditions = [];
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * SegmentCondition join.
         * @member {string} join
         * @memberof user.SegmentCondition
         * @instance
         */
        SegmentCondition.prototype.join = "";

        /**
         * SegmentCondition id.
         * @member {string} id
         * @memberof user.SegmentCondition
         * @instance
         */
        SegmentCondition.prototype.id = "";

        /**
         * SegmentCondition event_type.
         * @member {string} event_type
         * @memberof user.SegmentCondition
         * @instance
         */
        SegmentCondition.prototype.event_type = "";

        /**
         * SegmentCondition composition.
         * @member {string} composition
         * @memberof user.SegmentCondition
         * @instance
         */
        SegmentCondition.prototype.composition = "";

        /**
         * SegmentCondition key.
         * @member {string} key
         * @memberof user.SegmentCondition
         * @instance
         */
        SegmentCondition.prototype.key = "";

        /**
         * SegmentCondition conditions.
         * @member {Array.<user.ICondition>} conditions
         * @memberof user.SegmentCondition
         * @instance
         */
        SegmentCondition.prototype.conditions = $util.emptyArray;

        /**
         * SegmentCondition left_condition.
         * @member {user.ISegmentCondition|null|undefined} left_condition
         * @memberof user.SegmentCondition
         * @instance
         */
        SegmentCondition.prototype.left_condition = null;

        /**
         * SegmentCondition right_condition.
         * @member {user.ISegmentCondition|null|undefined} right_condition
         * @memberof user.SegmentCondition
         * @instance
         */
        SegmentCondition.prototype.right_condition = null;

        /**
         * JoinOperator enum.
         * @name user.SegmentCondition.JoinOperator
         * @enum {string}
         * @property {number} none=0 none value
         * @property {number} and=1 and value
         * @property {number} or=2 or value
         */
        SegmentCondition.JoinOperator = (function() {
            const valuesById = {}, values = Object.create(valuesById);
            values[valuesById[0] = "none"] = 0;
            values[valuesById[1] = "and"] = 1;
            values[valuesById[2] = "or"] = 2;
            return values;
        })();

        /**
         * Composition enum.
         * @name user.SegmentCondition.Composition
         * @enum {string}
         * @property {number} true=0 true value
         * @property {number} sum=1 sum value
         * @property {number} avg=2 avg value
         * @property {number} count=3 count value
         */
        SegmentCondition.Composition = (function() {
            const valuesById = {}, values = Object.create(valuesById);
            values[valuesById[0] = "true"] = 0;
            values[valuesById[1] = "sum"] = 1;
            values[valuesById[2] = "avg"] = 2;
            values[valuesById[3] = "count"] = 3;
            return values;
        })();

        return SegmentCondition;
    })();

    user.Condition = (function() {

        /**
         * Properties of a Condition.
         * @memberof user
         * @interface ICondition
         * @property {string|null} [join] Condition join
         * @property {string|null} [key] Condition key
         * @property {string|null} [operator] Condition operator
         * @property {string|null} [value] Condition value
         */

        /**
         * Constructs a new Condition.
         * @memberof user
         * @classdesc Represents a Condition.
         * @implements ICondition
         * @constructor
         * @param {user.ICondition=} [p] Properties to set
         */
        function Condition(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * Condition join.
         * @member {string} join
         * @memberof user.Condition
         * @instance
         */
        Condition.prototype.join = "";

        /**
         * Condition key.
         * @member {string} key
         * @memberof user.Condition
         * @instance
         */
        Condition.prototype.key = "";

        /**
         * Condition operator.
         * @member {string} operator
         * @memberof user.Condition
         * @instance
         */
        Condition.prototype.operator = "";

        /**
         * Condition value.
         * @member {string} value
         * @memberof user.Condition
         * @instance
         */
        Condition.prototype.value = "";

        /**
         * JoinOperator enum.
         * @name user.Condition.JoinOperator
         * @enum {string}
         * @property {number} none=0 none value
         * @property {number} and=1 and value
         * @property {number} or=2 or value
         */
        Condition.JoinOperator = (function() {
            const valuesById = {}, values = Object.create(valuesById);
            values[valuesById[0] = "none"] = 0;
            values[valuesById[1] = "and"] = 1;
            values[valuesById[2] = "or"] = 2;
            return values;
        })();

        return Condition;
    })();

    user.UserSearchResult = (function() {

        /**
         * Properties of a UserSearchResult.
         * @memberof user
         * @interface IUserSearchResult
         * @property {common.IContext|null} [ctx] UserSearchResult ctx
         * @property {number|Long|null} [total] UserSearchResult total
         * @property {Array.<user.IUser>|null} [users] UserSearchResult users
         * @property {string|null} [anchor] UserSearchResult anchor
         * @property {Array.<number>|null} [unreads] UserSearchResult unreads
         * @property {Object.<string,number>|null} [unread_counts] UserSearchResult unread_counts
         */

        /**
         * Constructs a new UserSearchResult.
         * @memberof user
         * @classdesc Represents a UserSearchResult.
         * @implements IUserSearchResult
         * @constructor
         * @param {user.IUserSearchResult=} [p] Properties to set
         */
        function UserSearchResult(p) {
            this.users = [];
            this.unreads = [];
            this.unread_counts = {};
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * UserSearchResult ctx.
         * @member {common.IContext|null|undefined} ctx
         * @memberof user.UserSearchResult
         * @instance
         */
        UserSearchResult.prototype.ctx = null;

        /**
         * UserSearchResult total.
         * @member {number|Long} total
         * @memberof user.UserSearchResult
         * @instance
         */
        UserSearchResult.prototype.total = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

        /**
         * UserSearchResult users.
         * @member {Array.<user.IUser>} users
         * @memberof user.UserSearchResult
         * @instance
         */
        UserSearchResult.prototype.users = $util.emptyArray;

        /**
         * UserSearchResult anchor.
         * @member {string} anchor
         * @memberof user.UserSearchResult
         * @instance
         */
        UserSearchResult.prototype.anchor = "";

        /**
         * UserSearchResult unreads.
         * @member {Array.<number>} unreads
         * @memberof user.UserSearchResult
         * @instance
         */
        UserSearchResult.prototype.unreads = $util.emptyArray;

        /**
         * UserSearchResult unread_counts.
         * @member {Object.<string,number>} unread_counts
         * @memberof user.UserSearchResult
         * @instance
         */
        UserSearchResult.prototype.unread_counts = $util.emptyObject;

        return UserSearchResult;
    })();

    user.UserSearchRequest = (function() {

        /**
         * Properties of a UserSearchRequest.
         * @memberof user
         * @interface IUserSearchRequest
         * @property {common.IContext|null} [ctx] UserSearchRequest ctx
         * @property {string|null} [segmentation_id] UserSearchRequest segmentation_id
         * @property {string|null} [query] UserSearchRequest query
         * @property {string|null} [anchor] UserSearchRequest anchor
         * @property {number|null} [limit] UserSearchRequest limit
         * @property {string|null} [agent_id] UserSearchRequest agent_id
         * @property {boolean|null} [unread] UserSearchRequest unread
         */

        /**
         * Constructs a new UserSearchRequest.
         * @memberof user
         * @classdesc Represents a UserSearchRequest.
         * @implements IUserSearchRequest
         * @constructor
         * @param {user.IUserSearchRequest=} [p] Properties to set
         */
        function UserSearchRequest(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * UserSearchRequest ctx.
         * @member {common.IContext|null|undefined} ctx
         * @memberof user.UserSearchRequest
         * @instance
         */
        UserSearchRequest.prototype.ctx = null;

        /**
         * UserSearchRequest segmentation_id.
         * @member {string} segmentation_id
         * @memberof user.UserSearchRequest
         * @instance
         */
        UserSearchRequest.prototype.segmentation_id = "";

        /**
         * UserSearchRequest query.
         * @member {string} query
         * @memberof user.UserSearchRequest
         * @instance
         */
        UserSearchRequest.prototype.query = "";

        /**
         * UserSearchRequest anchor.
         * @member {string} anchor
         * @memberof user.UserSearchRequest
         * @instance
         */
        UserSearchRequest.prototype.anchor = "";

        /**
         * UserSearchRequest limit.
         * @member {number} limit
         * @memberof user.UserSearchRequest
         * @instance
         */
        UserSearchRequest.prototype.limit = 0;

        /**
         * UserSearchRequest agent_id.
         * @member {string} agent_id
         * @memberof user.UserSearchRequest
         * @instance
         */
        UserSearchRequest.prototype.agent_id = "";

        /**
         * UserSearchRequest unread.
         * @member {boolean} unread
         * @memberof user.UserSearchRequest
         * @instance
         */
        UserSearchRequest.prototype.unread = false;

        return UserSearchRequest;
    })();

    user.IndexEvent = (function() {

        /**
         * Properties of an IndexEvent.
         * @memberof user
         * @interface IIndexEvent
         * @property {string|null} [id] IndexEvent id
         * @property {string|null} [account_id] IndexEvent account_id
         * @property {string|null} [user_id] IndexEvent user_id
         * @property {number|Long|null} [created] IndexEvent created
         * @property {string|null} [category] IndexEvent category
         * @property {Array.<string>|null} [topics] IndexEvent topics
         * @property {string|null} [object] IndexEvent object
         * @property {string|null} [text] IndexEvent text
         */

        /**
         * Constructs a new IndexEvent.
         * @memberof user
         * @classdesc Represents an IndexEvent.
         * @implements IIndexEvent
         * @constructor
         * @param {user.IIndexEvent=} [p] Properties to set
         */
        function IndexEvent(p) {
            this.topics = [];
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * IndexEvent id.
         * @member {string} id
         * @memberof user.IndexEvent
         * @instance
         */
        IndexEvent.prototype.id = "";

        /**
         * IndexEvent account_id.
         * @member {string} account_id
         * @memberof user.IndexEvent
         * @instance
         */
        IndexEvent.prototype.account_id = "";

        /**
         * IndexEvent user_id.
         * @member {string} user_id
         * @memberof user.IndexEvent
         * @instance
         */
        IndexEvent.prototype.user_id = "";

        /**
         * IndexEvent created.
         * @member {number|Long} created
         * @memberof user.IndexEvent
         * @instance
         */
        IndexEvent.prototype.created = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

        /**
         * IndexEvent category.
         * @member {string} category
         * @memberof user.IndexEvent
         * @instance
         */
        IndexEvent.prototype.category = "";

        /**
         * IndexEvent topics.
         * @member {Array.<string>} topics
         * @memberof user.IndexEvent
         * @instance
         */
        IndexEvent.prototype.topics = $util.emptyArray;

        /**
         * IndexEvent object.
         * @member {string} object
         * @memberof user.IndexEvent
         * @instance
         */
        IndexEvent.prototype.object = "";

        /**
         * IndexEvent text.
         * @member {string} text
         * @memberof user.IndexEvent
         * @instance
         */
        IndexEvent.prototype.text = "";

        return IndexEvent;
    })();

    user.ListTopicsRequest = (function() {

        /**
         * Properties of a ListTopicsRequest.
         * @memberof user
         * @interface IListTopicsRequest
         * @property {common.IContext|null} [ctx] ListTopicsRequest ctx
         * @property {string|null} [user_id] ListTopicsRequest user_id
         * @property {string|null} [agent_id] ListTopicsRequest agent_id
         * @property {string|null} [anchor] ListTopicsRequest anchor
         * @property {number|null} [limit] ListTopicsRequest limit
         * @property {boolean|null} [unread] ListTopicsRequest unread
         */

        /**
         * Constructs a new ListTopicsRequest.
         * @memberof user
         * @classdesc Represents a ListTopicsRequest.
         * @implements IListTopicsRequest
         * @constructor
         * @param {user.IListTopicsRequest=} [p] Properties to set
         */
        function ListTopicsRequest(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * ListTopicsRequest ctx.
         * @member {common.IContext|null|undefined} ctx
         * @memberof user.ListTopicsRequest
         * @instance
         */
        ListTopicsRequest.prototype.ctx = null;

        /**
         * ListTopicsRequest user_id.
         * @member {string} user_id
         * @memberof user.ListTopicsRequest
         * @instance
         */
        ListTopicsRequest.prototype.user_id = "";

        /**
         * ListTopicsRequest agent_id.
         * @member {string} agent_id
         * @memberof user.ListTopicsRequest
         * @instance
         */
        ListTopicsRequest.prototype.agent_id = "";

        /**
         * ListTopicsRequest anchor.
         * @member {string} anchor
         * @memberof user.ListTopicsRequest
         * @instance
         */
        ListTopicsRequest.prototype.anchor = "";

        /**
         * ListTopicsRequest limit.
         * @member {number} limit
         * @memberof user.ListTopicsRequest
         * @instance
         */
        ListTopicsRequest.prototype.limit = 0;

        /**
         * ListTopicsRequest unread.
         * @member {boolean} unread
         * @memberof user.ListTopicsRequest
         * @instance
         */
        ListTopicsRequest.prototype.unread = false;

        return ListTopicsRequest;
    })();

    user.ListTopicsResult = (function() {

        /**
         * Properties of a ListTopicsResult.
         * @memberof user
         * @interface IListTopicsResult
         * @property {common.IContext|null} [ctx] ListTopicsResult ctx
         * @property {Array.<user.ITopic>|null} [topics] ListTopicsResult topics
         * @property {string|null} [anchor] ListTopicsResult anchor
         */

        /**
         * Constructs a new ListTopicsResult.
         * @memberof user
         * @classdesc Represents a ListTopicsResult.
         * @implements IListTopicsResult
         * @constructor
         * @param {user.IListTopicsResult=} [p] Properties to set
         */
        function ListTopicsResult(p) {
            this.topics = [];
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * ListTopicsResult ctx.
         * @member {common.IContext|null|undefined} ctx
         * @memberof user.ListTopicsResult
         * @instance
         */
        ListTopicsResult.prototype.ctx = null;

        /**
         * ListTopicsResult topics.
         * @member {Array.<user.ITopic>} topics
         * @memberof user.ListTopicsResult
         * @instance
         */
        ListTopicsResult.prototype.topics = $util.emptyArray;

        /**
         * ListTopicsResult anchor.
         * @member {string} anchor
         * @memberof user.ListTopicsResult
         * @instance
         */
        ListTopicsResult.prototype.anchor = "";

        return ListTopicsResult;
    })();

    user.ListNewsRequest = (function() {

        /**
         * Properties of a ListNewsRequest.
         * @memberof user
         * @interface IListNewsRequest
         * @property {common.IContext|null} [ctx] ListNewsRequest ctx
         * @property {string|null} [user_id] ListNewsRequest user_id
         * @property {number|Long|null} [start_time] ListNewsRequest start_time
         * @property {string|null} [limit] ListNewsRequest limit
         */

        /**
         * Constructs a new ListNewsRequest.
         * @memberof user
         * @classdesc Represents a ListNewsRequest.
         * @implements IListNewsRequest
         * @constructor
         * @param {user.IListNewsRequest=} [p] Properties to set
         */
        function ListNewsRequest(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * ListNewsRequest ctx.
         * @member {common.IContext|null|undefined} ctx
         * @memberof user.ListNewsRequest
         * @instance
         */
        ListNewsRequest.prototype.ctx = null;

        /**
         * ListNewsRequest user_id.
         * @member {string} user_id
         * @memberof user.ListNewsRequest
         * @instance
         */
        ListNewsRequest.prototype.user_id = "";

        /**
         * ListNewsRequest start_time.
         * @member {number|Long} start_time
         * @memberof user.ListNewsRequest
         * @instance
         */
        ListNewsRequest.prototype.start_time = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

        /**
         * ListNewsRequest limit.
         * @member {string} limit
         * @memberof user.ListNewsRequest
         * @instance
         */
        ListNewsRequest.prototype.limit = "";

        return ListNewsRequest;
    })();

    user.AddToMyList = (function() {

        /**
         * Properties of an AddToMyList.
         * @memberof user
         * @interface IAddToMyList
         * @property {common.IContext|null} [ctx] AddToMyList ctx
         * @property {string|null} [agent_id] AddToMyList agent_id
         * @property {string|null} [user_id] AddToMyList user_id
         */

        /**
         * Constructs a new AddToMyList.
         * @memberof user
         * @classdesc Represents an AddToMyList.
         * @implements IAddToMyList
         * @constructor
         * @param {user.IAddToMyList=} [p] Properties to set
         */
        function AddToMyList(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * AddToMyList ctx.
         * @member {common.IContext|null|undefined} ctx
         * @memberof user.AddToMyList
         * @instance
         */
        AddToMyList.prototype.ctx = null;

        /**
         * AddToMyList agent_id.
         * @member {string} agent_id
         * @memberof user.AddToMyList
         * @instance
         */
        AddToMyList.prototype.agent_id = "";

        /**
         * AddToMyList user_id.
         * @member {string} user_id
         * @memberof user.AddToMyList
         * @instance
         */
        AddToMyList.prototype.user_id = "";

        return AddToMyList;
    })();

    user.Field = (function() {

        /**
         * Properties of a Field.
         * @memberof user
         * @interface IField
         * @property {string|null} [name] Field name
         * @property {string|null} [account_id] Field account_id
         * @property {string|null} [user_id] Field user_id
         * @property {string|null} [setter] Field setter
         * @property {string|null} [setter_type] Field setter_type
         * @property {number|Long|null} [updated] Field updated
         * @property {string|null} [data] Field data
         * @property {string|null} [id] Field id
         */

        /**
         * Constructs a new Field.
         * @memberof user
         * @classdesc Represents a Field.
         * @implements IField
         * @constructor
         * @param {user.IField=} [p] Properties to set
         */
        function Field(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * Field name.
         * @member {string} name
         * @memberof user.Field
         * @instance
         */
        Field.prototype.name = "";

        /**
         * Field account_id.
         * @member {string} account_id
         * @memberof user.Field
         * @instance
         */
        Field.prototype.account_id = "";

        /**
         * Field user_id.
         * @member {string} user_id
         * @memberof user.Field
         * @instance
         */
        Field.prototype.user_id = "";

        /**
         * Field setter.
         * @member {string} setter
         * @memberof user.Field
         * @instance
         */
        Field.prototype.setter = "";

        /**
         * Field setter_type.
         * @member {string} setter_type
         * @memberof user.Field
         * @instance
         */
        Field.prototype.setter_type = "";

        /**
         * Field updated.
         * @member {number|Long} updated
         * @memberof user.Field
         * @instance
         */
        Field.prototype.updated = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

        /**
         * Field data.
         * @member {string} data
         * @memberof user.Field
         * @instance
         */
        Field.prototype.data = "";

        /**
         * Field id.
         * @member {string} id
         * @memberof user.Field
         * @instance
         */
        Field.prototype.id = "";

        return Field;
    })();

    user.Presence = (function() {

        /**
         * Properties of a Presence.
         * @memberof user
         * @interface IPresence
         * @property {common.IContext|null} [ctx] Presence ctx
         * @property {string|null} [account_id] Presence account_id
         * @property {string|null} [user_id] Presence user_id
         * @property {number|Long|null} [pinged] Presence pinged
         * @property {number|Long|null} [pinged_minute] Presence pinged_minute
         */

        /**
         * Constructs a new Presence.
         * @memberof user
         * @classdesc Represents a Presence.
         * @implements IPresence
         * @constructor
         * @param {user.IPresence=} [p] Properties to set
         */
        function Presence(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * Presence ctx.
         * @member {common.IContext|null|undefined} ctx
         * @memberof user.Presence
         * @instance
         */
        Presence.prototype.ctx = null;

        /**
         * Presence account_id.
         * @member {string} account_id
         * @memberof user.Presence
         * @instance
         */
        Presence.prototype.account_id = "";

        /**
         * Presence user_id.
         * @member {string} user_id
         * @memberof user.Presence
         * @instance
         */
        Presence.prototype.user_id = "";

        /**
         * Presence pinged.
         * @member {number|Long} pinged
         * @memberof user.Presence
         * @instance
         */
        Presence.prototype.pinged = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

        /**
         * Presence pinged_minute.
         * @member {number|Long} pinged_minute
         * @memberof user.Presence
         * @instance
         */
        Presence.prototype.pinged_minute = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

        return Presence;
    })();

    user.Presences = (function() {

        /**
         * Properties of a Presences.
         * @memberof user
         * @interface IPresences
         * @property {common.IContext|null} [ctx] Presences ctx
         * @property {Array.<user.IPresence>|null} [presences] Presences presences
         */

        /**
         * Constructs a new Presences.
         * @memberof user
         * @classdesc Represents a Presences.
         * @implements IPresences
         * @constructor
         * @param {user.IPresences=} [p] Properties to set
         */
        function Presences(p) {
            this.presences = [];
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * Presences ctx.
         * @member {common.IContext|null|undefined} ctx
         * @memberof user.Presences
         * @instance
         */
        Presences.prototype.ctx = null;

        /**
         * Presences presences.
         * @member {Array.<user.IPresence>} presences
         * @memberof user.Presences
         * @instance
         */
        Presences.prototype.presences = $util.emptyArray;

        return Presences;
    })();

    user.Visitor = (function() {

        /**
         * Properties of a Visitor.
         * @memberof user
         * @interface IVisitor
         * @property {common.IContext|null} [ctx] Visitor ctx
         * @property {string|null} [account_id] Visitor account_id
         * @property {user.IUser|null} [user] Visitor user
         * @property {number|Long|null} [pinged] Visitor pinged
         * @property {string|null} [page_url] Visitor page_url
         * @property {number|Long|null} [page_viewed] Visitor page_viewed
         * @property {string|null} [page_title] Visitor page_title
         */

        /**
         * Constructs a new Visitor.
         * @memberof user
         * @classdesc Represents a Visitor.
         * @implements IVisitor
         * @constructor
         * @param {user.IVisitor=} [p] Properties to set
         */
        function Visitor(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * Visitor ctx.
         * @member {common.IContext|null|undefined} ctx
         * @memberof user.Visitor
         * @instance
         */
        Visitor.prototype.ctx = null;

        /**
         * Visitor account_id.
         * @member {string} account_id
         * @memberof user.Visitor
         * @instance
         */
        Visitor.prototype.account_id = "";

        /**
         * Visitor user.
         * @member {user.IUser|null|undefined} user
         * @memberof user.Visitor
         * @instance
         */
        Visitor.prototype.user = null;

        /**
         * Visitor pinged.
         * @member {number|Long} pinged
         * @memberof user.Visitor
         * @instance
         */
        Visitor.prototype.pinged = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

        /**
         * Visitor page_url.
         * @member {string} page_url
         * @memberof user.Visitor
         * @instance
         */
        Visitor.prototype.page_url = "";

        /**
         * Visitor page_viewed.
         * @member {number|Long} page_viewed
         * @memberof user.Visitor
         * @instance
         */
        Visitor.prototype.page_viewed = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

        /**
         * Visitor page_title.
         * @member {string} page_title
         * @memberof user.Visitor
         * @instance
         */
        Visitor.prototype.page_title = "";

        return Visitor;
    })();

    user.Visitors = (function() {

        /**
         * Properties of a Visitors.
         * @memberof user
         * @interface IVisitors
         * @property {common.IContext|null} [ctx] Visitors ctx
         * @property {Array.<user.IVisitor>|null} [visitors] Visitors visitors
         */

        /**
         * Constructs a new Visitors.
         * @memberof user
         * @classdesc Represents a Visitors.
         * @implements IVisitors
         * @constructor
         * @param {user.IVisitors=} [p] Properties to set
         */
        function Visitors(p) {
            this.visitors = [];
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * Visitors ctx.
         * @member {common.IContext|null|undefined} ctx
         * @memberof user.Visitors
         * @instance
         */
        Visitors.prototype.ctx = null;

        /**
         * Visitors visitors.
         * @member {Array.<user.IVisitor>} visitors
         * @memberof user.Visitors
         * @instance
         */
        Visitors.prototype.visitors = $util.emptyArray;

        return Visitors;
    })();

    user.LastView = (function() {

        /**
         * Properties of a LastView.
         * @memberof user
         * @interface ILastView
         * @property {string|null} [account_id] LastView account_id
         * @property {string|null} [user_id] LastView user_id
         * @property {string|null} [url] LastView url
         * @property {string|null} [ua] LastView ua
         * @property {string|null} [ip] LastView ip
         * @property {number|Long|null} [created] LastView created
         * @property {string|null} [event_id] LastView event_id
         * @property {string|null} [title] LastView title
         */

        /**
         * Constructs a new LastView.
         * @memberof user
         * @classdesc Represents a LastView.
         * @implements ILastView
         * @constructor
         * @param {user.ILastView=} [p] Properties to set
         */
        function LastView(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * LastView account_id.
         * @member {string} account_id
         * @memberof user.LastView
         * @instance
         */
        LastView.prototype.account_id = "";

        /**
         * LastView user_id.
         * @member {string} user_id
         * @memberof user.LastView
         * @instance
         */
        LastView.prototype.user_id = "";

        /**
         * LastView url.
         * @member {string} url
         * @memberof user.LastView
         * @instance
         */
        LastView.prototype.url = "";

        /**
         * LastView ua.
         * @member {string} ua
         * @memberof user.LastView
         * @instance
         */
        LastView.prototype.ua = "";

        /**
         * LastView ip.
         * @member {string} ip
         * @memberof user.LastView
         * @instance
         */
        LastView.prototype.ip = "";

        /**
         * LastView created.
         * @member {number|Long} created
         * @memberof user.LastView
         * @instance
         */
        LastView.prototype.created = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

        /**
         * LastView event_id.
         * @member {string} event_id
         * @memberof user.LastView
         * @instance
         */
        LastView.prototype.event_id = "";

        /**
         * LastView title.
         * @member {string} title
         * @memberof user.LastView
         * @instance
         */
        LastView.prototype.title = "";

        return LastView;
    })();

    user.Automation = (function() {

        /**
         * Properties of an Automation.
         * @memberof user
         * @interface IAutomation
         * @property {common.IContext|null} [ctx] Automation ctx
         * @property {string|null} [account_id] Automation account_id
         * @property {string|null} [id] Automation id
         * @property {string|null} [channel] Automation channel
         * @property {string|null} [name] Automation name
         * @property {string|null} [description] Automation description
         * @property {Array.<user.ICondition>|null} [conditions] Automation conditions
         * @property {number|Long|null} [created] Automation created
         * @property {number|Long|null} [modified] Automation modified
         * @property {string|null} [state] Automation state
         * @property {string|null} [action_type] Automation action_type
         * @property {string|null} [action_data] Automation action_data
         * @property {string|null} [scope] Automation scope
         */

        /**
         * Constructs a new Automation.
         * @memberof user
         * @classdesc Represents an Automation.
         * @implements IAutomation
         * @constructor
         * @param {user.IAutomation=} [p] Properties to set
         */
        function Automation(p) {
            this.conditions = [];
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * Automation ctx.
         * @member {common.IContext|null|undefined} ctx
         * @memberof user.Automation
         * @instance
         */
        Automation.prototype.ctx = null;

        /**
         * Automation account_id.
         * @member {string} account_id
         * @memberof user.Automation
         * @instance
         */
        Automation.prototype.account_id = "";

        /**
         * Automation id.
         * @member {string} id
         * @memberof user.Automation
         * @instance
         */
        Automation.prototype.id = "";

        /**
         * Automation channel.
         * @member {string} channel
         * @memberof user.Automation
         * @instance
         */
        Automation.prototype.channel = "";

        /**
         * Automation name.
         * @member {string} name
         * @memberof user.Automation
         * @instance
         */
        Automation.prototype.name = "";

        /**
         * Automation description.
         * @member {string} description
         * @memberof user.Automation
         * @instance
         */
        Automation.prototype.description = "";

        /**
         * Automation conditions.
         * @member {Array.<user.ICondition>} conditions
         * @memberof user.Automation
         * @instance
         */
        Automation.prototype.conditions = $util.emptyArray;

        /**
         * Automation created.
         * @member {number|Long} created
         * @memberof user.Automation
         * @instance
         */
        Automation.prototype.created = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

        /**
         * Automation modified.
         * @member {number|Long} modified
         * @memberof user.Automation
         * @instance
         */
        Automation.prototype.modified = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

        /**
         * Automation state.
         * @member {string} state
         * @memberof user.Automation
         * @instance
         */
        Automation.prototype.state = "";

        /**
         * Automation action_type.
         * @member {string} action_type
         * @memberof user.Automation
         * @instance
         */
        Automation.prototype.action_type = "";

        /**
         * Automation action_data.
         * @member {string} action_data
         * @memberof user.Automation
         * @instance
         */
        Automation.prototype.action_data = "";

        /**
         * Automation scope.
         * @member {string} scope
         * @memberof user.Automation
         * @instance
         */
        Automation.prototype.scope = "";

        /**
         * State enum.
         * @name user.Automation.State
         * @enum {string}
         * @property {number} active=0 active value
         * @property {number} inactive=1 inactive value
         */
        Automation.State = (function() {
            const valuesById = {}, values = Object.create(valuesById);
            values[valuesById[0] = "active"] = 0;
            values[valuesById[1] = "inactive"] = 1;
            return values;
        })();

        /**
         * ActionType enum.
         * @name user.Automation.ActionType
         * @enum {string}
         * @property {number} conversation_message=0 conversation_message value
         * @property {number} agent_notification=1 agent_notification value
         * @property {number} automation_invite_message=4 automation_invite_message value
         */
        Automation.ActionType = (function() {
            const valuesById = {}, values = Object.create(valuesById);
            values[valuesById[0] = "conversation_message"] = 0;
            values[valuesById[1] = "agent_notification"] = 1;
            values[valuesById[4] = "automation_invite_message"] = 4;
            return values;
        })();

        /**
         * Scope enum.
         * @name user.Automation.Scope
         * @enum {string}
         * @property {number} conversation=2 conversation value
         * @property {number} user=3 user value
         */
        Automation.Scope = (function() {
            const valuesById = {}, values = Object.create(valuesById);
            values[valuesById[2] = "conversation"] = 2;
            values[valuesById[3] = "user"] = 3;
            return values;
        })();

        return Automation;
    })();

    user.Automations = (function() {

        /**
         * Properties of an Automations.
         * @memberof user
         * @interface IAutomations
         * @property {common.IContext|null} [ctx] Automations ctx
         * @property {Array.<user.IAutomation>|null} [automations] Automations automations
         */

        /**
         * Constructs a new Automations.
         * @memberof user
         * @classdesc Represents an Automations.
         * @implements IAutomations
         * @constructor
         * @param {user.IAutomations=} [p] Properties to set
         */
        function Automations(p) {
            this.automations = [];
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * Automations ctx.
         * @member {common.IContext|null|undefined} ctx
         * @memberof user.Automations
         * @instance
         */
        Automations.prototype.ctx = null;

        /**
         * Automations automations.
         * @member {Array.<user.IAutomation>} automations
         * @memberof user.Automations
         * @instance
         */
        Automations.prototype.automations = $util.emptyArray;

        return Automations;
    })();

    user.Session = (function() {

        /**
         * Properties of a Session.
         * @memberof user
         * @interface ISession
         * @property {common.IContext|null} [ctx] Session ctx
         * @property {string|null} [account_id] Session account_id
         * @property {string|null} [user_id] Session user_id
         * @property {string|null} [id] Session id
         * @property {string|null} [platform] Session platform
         * @property {string|null} [referrer] Session referrer
         * @property {string|null} [search_engine] Session search_engine
         * @property {number|Long|null} [started] Session started
         * @property {number|Long|null} [tracked] Session tracked
         * @property {string|null} [status] Session status
         * @property {number|null} [events_count] Session events_count
         * @property {number|null} [content_views_count] Session content_views_count
         * @property {string|null} [search_term] Session search_term
         */

        /**
         * Constructs a new Session.
         * @memberof user
         * @classdesc Represents a Session.
         * @implements ISession
         * @constructor
         * @param {user.ISession=} [p] Properties to set
         */
        function Session(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * Session ctx.
         * @member {common.IContext|null|undefined} ctx
         * @memberof user.Session
         * @instance
         */
        Session.prototype.ctx = null;

        /**
         * Session account_id.
         * @member {string} account_id
         * @memberof user.Session
         * @instance
         */
        Session.prototype.account_id = "";

        /**
         * Session user_id.
         * @member {string} user_id
         * @memberof user.Session
         * @instance
         */
        Session.prototype.user_id = "";

        /**
         * Session id.
         * @member {string} id
         * @memberof user.Session
         * @instance
         */
        Session.prototype.id = "";

        /**
         * Session platform.
         * @member {string} platform
         * @memberof user.Session
         * @instance
         */
        Session.prototype.platform = "";

        /**
         * Session referrer.
         * @member {string} referrer
         * @memberof user.Session
         * @instance
         */
        Session.prototype.referrer = "";

        /**
         * Session search_engine.
         * @member {string} search_engine
         * @memberof user.Session
         * @instance
         */
        Session.prototype.search_engine = "";

        /**
         * Session started.
         * @member {number|Long} started
         * @memberof user.Session
         * @instance
         */
        Session.prototype.started = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

        /**
         * Session tracked.
         * @member {number|Long} tracked
         * @memberof user.Session
         * @instance
         */
        Session.prototype.tracked = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

        /**
         * Session status.
         * @member {string} status
         * @memberof user.Session
         * @instance
         */
        Session.prototype.status = "";

        /**
         * Session events_count.
         * @member {number} events_count
         * @memberof user.Session
         * @instance
         */
        Session.prototype.events_count = 0;

        /**
         * Session content_views_count.
         * @member {number} content_views_count
         * @memberof user.Session
         * @instance
         */
        Session.prototype.content_views_count = 0;

        /**
         * Session search_term.
         * @member {string} search_term
         * @memberof user.Session
         * @instance
         */
        Session.prototype.search_term = "";

        /**
         * Platform enum.
         * @name user.Session.Platform
         * @enum {string}
         * @property {number} web=0 web value
         * @property {number} mobile=2 mobile value
         * @property {number} desktop=4 desktop value
         */
        Session.Platform = (function() {
            const valuesById = {}, values = Object.create(valuesById);
            values[valuesById[0] = "web"] = 0;
            values[valuesById[2] = "mobile"] = 2;
            values[valuesById[4] = "desktop"] = 4;
            return values;
        })();

        /**
         * Status enum.
         * @name user.Session.Status
         * @enum {string}
         * @property {number} open=0 open value
         * @property {number} closed=1 closed value
         */
        Session.Status = (function() {
            const valuesById = {}, values = Object.create(valuesById);
            values[valuesById[0] = "open"] = 0;
            values[valuesById[1] = "closed"] = 1;
            return values;
        })();

        return Session;
    })();

    user.DeleteAttrRequest = (function() {

        /**
         * Properties of a DeleteAttrRequest.
         * @memberof user
         * @interface IDeleteAttrRequest
         * @property {common.IContext|null} [ctx] DeleteAttrRequest ctx
         * @property {string|null} [key] DeleteAttrRequest key
         */

        /**
         * Constructs a new DeleteAttrRequest.
         * @memberof user
         * @classdesc Represents a DeleteAttrRequest.
         * @implements IDeleteAttrRequest
         * @constructor
         * @param {user.IDeleteAttrRequest=} [p] Properties to set
         */
        function DeleteAttrRequest(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * DeleteAttrRequest ctx.
         * @member {common.IContext|null|undefined} ctx
         * @memberof user.DeleteAttrRequest
         * @instance
         */
        DeleteAttrRequest.prototype.ctx = null;

        /**
         * DeleteAttrRequest key.
         * @member {string} key
         * @memberof user.DeleteAttrRequest
         * @instance
         */
        DeleteAttrRequest.prototype.key = "";

        return DeleteAttrRequest;
    })();

    user.SegmentationMgr = (function() {

        /**
         * Constructs a new SegmentationMgr service.
         * @memberof user
         * @classdesc Represents a SegmentationMgr
         * @extends $protobuf.rpc.Service
         * @constructor
         * @param {$protobuf.RPCImpl} rpcImpl RPC implementation
         * @param {boolean} [requestDelimited=false] Whether requests are length-delimited
         * @param {boolean} [responseDelimited=false] Whether responses are length-delimited
         */
        function SegmentationMgr(rpcImpl, requestDelimited, responseDelimited) {
            $protobuf.rpc.Service.call(this, rpcImpl, requestDelimited, responseDelimited);
        }

        (SegmentationMgr.prototype = Object.create($protobuf.rpc.Service.prototype)).constructor = SegmentationMgr;

        /**
         * Callback as used by {@link user.SegmentationMgr#createSegment}.
         * @memberof user.SegmentationMgr
         * @typedef CreateSegmentCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {user.Segmentation} [response] Segmentation
         */

        /**
         * Calls CreateSegment.
         * @function createSegment
         * @memberof user.SegmentationMgr
         * @instance
         * @param {user.ISegmentation} request Segmentation message or plain object
         * @param {user.SegmentationMgr.CreateSegmentCallback} callback Node-style callback called with the error, if any, and Segmentation
         * @returns {undefined}
         * @variation 1
         */
        SegmentationMgr.prototype.createSegment = function createSegment(request, callback) {
            return this.rpcCall(createSegment, $root.user.Segmentation, $root.user.Segmentation, request, callback);
        };

        /**
         * Calls CreateSegment.
         * @function createSegment
         * @memberof user.SegmentationMgr
         * @instance
         * @param {user.ISegmentation} request Segmentation message or plain object
         * @returns {Promise<user.Segmentation>} Promise
         * @variation 2
         */

        /**
         * Callback as used by {@link user.SegmentationMgr#updateSegment}.
         * @memberof user.SegmentationMgr
         * @typedef UpdateSegmentCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {user.Segmentation} [response] Segmentation
         */

        /**
         * Calls UpdateSegment.
         * @function updateSegment
         * @memberof user.SegmentationMgr
         * @instance
         * @param {user.ISegmentation} request Segmentation message or plain object
         * @param {user.SegmentationMgr.UpdateSegmentCallback} callback Node-style callback called with the error, if any, and Segmentation
         * @returns {undefined}
         * @variation 1
         */
        SegmentationMgr.prototype.updateSegment = function updateSegment(request, callback) {
            return this.rpcCall(updateSegment, $root.user.Segmentation, $root.user.Segmentation, request, callback);
        };

        /**
         * Calls UpdateSegment.
         * @function updateSegment
         * @memberof user.SegmentationMgr
         * @instance
         * @param {user.ISegmentation} request Segmentation message or plain object
         * @returns {Promise<user.Segmentation>} Promise
         * @variation 2
         */

        /**
         * Callback as used by {@link user.SegmentationMgr#listSegments}.
         * @memberof user.SegmentationMgr
         * @typedef ListSegmentsCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {user.Segmentations} [response] Segmentations
         */

        /**
         * Calls ListSegments.
         * @function listSegments
         * @memberof user.SegmentationMgr
         * @instance
         * @param {common.IId} request Id message or plain object
         * @param {user.SegmentationMgr.ListSegmentsCallback} callback Node-style callback called with the error, if any, and Segmentations
         * @returns {undefined}
         * @variation 1
         */
        SegmentationMgr.prototype.listSegments = function listSegments(request, callback) {
            return this.rpcCall(listSegments, $root.common.Id, $root.user.Segmentations, request, callback);
        };

        /**
         * Calls ListSegments.
         * @function listSegments
         * @memberof user.SegmentationMgr
         * @instance
         * @param {common.IId} request Id message or plain object
         * @returns {Promise<user.Segmentations>} Promise
         * @variation 2
         */

        /**
         * Callback as used by {@link user.SegmentationMgr#deleteSegment}.
         * @memberof user.SegmentationMgr
         * @typedef DeleteSegmentCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {common.Empty} [response] Empty
         */

        /**
         * Calls DeleteSegment.
         * @function deleteSegment
         * @memberof user.SegmentationMgr
         * @instance
         * @param {common.IId} request Id message or plain object
         * @param {user.SegmentationMgr.DeleteSegmentCallback} callback Node-style callback called with the error, if any, and Empty
         * @returns {undefined}
         * @variation 1
         */
        SegmentationMgr.prototype.deleteSegment = function deleteSegment(request, callback) {
            return this.rpcCall(deleteSegment, $root.common.Id, $root.common.Empty, request, callback);
        };

        /**
         * Calls DeleteSegment.
         * @function deleteSegment
         * @memberof user.SegmentationMgr
         * @instance
         * @param {common.IId} request Id message or plain object
         * @returns {Promise<common.Empty>} Promise
         * @variation 2
         */

        /**
         * Callback as used by {@link user.SegmentationMgr#readSegment}.
         * @memberof user.SegmentationMgr
         * @typedef ReadSegmentCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {user.Segmentation} [response] Segmentation
         */

        /**
         * Calls ReadSegment.
         * @function readSegment
         * @memberof user.SegmentationMgr
         * @instance
         * @param {common.IId} request Id message or plain object
         * @param {user.SegmentationMgr.ReadSegmentCallback} callback Node-style callback called with the error, if any, and Segmentation
         * @returns {undefined}
         * @variation 1
         */
        SegmentationMgr.prototype.readSegment = function readSegment(request, callback) {
            return this.rpcCall(readSegment, $root.common.Id, $root.user.Segmentation, request, callback);
        };

        /**
         * Calls ReadSegment.
         * @function readSegment
         * @memberof user.SegmentationMgr
         * @instance
         * @param {common.IId} request Id message or plain object
         * @returns {Promise<user.Segmentation>} Promise
         * @variation 2
         */

        return SegmentationMgr;
    })();

    user.VisitorMgr = (function() {

        /**
         * Constructs a new VisitorMgr service.
         * @memberof user
         * @classdesc Represents a VisitorMgr
         * @extends $protobuf.rpc.Service
         * @constructor
         * @param {$protobuf.RPCImpl} rpcImpl RPC implementation
         * @param {boolean} [requestDelimited=false] Whether requests are length-delimited
         * @param {boolean} [responseDelimited=false] Whether responses are length-delimited
         */
        function VisitorMgr(rpcImpl, requestDelimited, responseDelimited) {
            $protobuf.rpc.Service.call(this, rpcImpl, requestDelimited, responseDelimited);
        }

        (VisitorMgr.prototype = Object.create($protobuf.rpc.Service.prototype)).constructor = VisitorMgr;

        /**
         * Callback as used by {@link user.VisitorMgr#readPresence}.
         * @memberof user.VisitorMgr
         * @typedef ReadPresenceCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {user.Presence} [response] Presence
         */

        /**
         * Calls ReadPresence.
         * @function readPresence
         * @memberof user.VisitorMgr
         * @instance
         * @param {common.IId} request Id message or plain object
         * @param {user.VisitorMgr.ReadPresenceCallback} callback Node-style callback called with the error, if any, and Presence
         * @returns {undefined}
         * @variation 1
         */
        VisitorMgr.prototype.readPresence = function readPresence(request, callback) {
            return this.rpcCall(readPresence, $root.common.Id, $root.user.Presence, request, callback);
        };

        /**
         * Calls ReadPresence.
         * @function readPresence
         * @memberof user.VisitorMgr
         * @instance
         * @param {common.IId} request Id message or plain object
         * @returns {Promise<user.Presence>} Promise
         * @variation 2
         */

        /**
         * Callback as used by {@link user.VisitorMgr#readPresences}.
         * @memberof user.VisitorMgr
         * @typedef ReadPresencesCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {user.Presences} [response] Presences
         */

        /**
         * Calls ReadPresences.
         * @function readPresences
         * @memberof user.VisitorMgr
         * @instance
         * @param {common.IIds} request Ids message or plain object
         * @param {user.VisitorMgr.ReadPresencesCallback} callback Node-style callback called with the error, if any, and Presences
         * @returns {undefined}
         * @variation 1
         */
        VisitorMgr.prototype.readPresences = function readPresences(request, callback) {
            return this.rpcCall(readPresences, $root.common.Ids, $root.user.Presences, request, callback);
        };

        /**
         * Calls ReadPresences.
         * @function readPresences
         * @memberof user.VisitorMgr
         * @instance
         * @param {common.IIds} request Ids message or plain object
         * @returns {Promise<user.Presences>} Promise
         * @variation 2
         */

        /**
         * Callback as used by {@link user.VisitorMgr#readPreview}.
         * @memberof user.VisitorMgr
         * @typedef ReadPreviewCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {user.LastView} [response] LastView
         */

        /**
         * Calls ReadPreview.
         * @function readPreview
         * @memberof user.VisitorMgr
         * @instance
         * @param {common.IId} request Id message or plain object
         * @param {user.VisitorMgr.ReadPreviewCallback} callback Node-style callback called with the error, if any, and LastView
         * @returns {undefined}
         * @variation 1
         */
        VisitorMgr.prototype.readPreview = function readPreview(request, callback) {
            return this.rpcCall(readPreview, $root.common.Id, $root.user.LastView, request, callback);
        };

        /**
         * Calls ReadPreview.
         * @function readPreview
         * @memberof user.VisitorMgr
         * @instance
         * @param {common.IId} request Id message or plain object
         * @returns {Promise<user.LastView>} Promise
         * @variation 2
         */

        /**
         * Callback as used by {@link user.VisitorMgr#listTopVisitors}.
         * @memberof user.VisitorMgr
         * @typedef ListTopVisitorsCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {user.Visitors} [response] Visitors
         */

        /**
         * Calls ListTopVisitors.
         * @function listTopVisitors
         * @memberof user.VisitorMgr
         * @instance
         * @param {common.IId} request Id message or plain object
         * @param {user.VisitorMgr.ListTopVisitorsCallback} callback Node-style callback called with the error, if any, and Visitors
         * @returns {undefined}
         * @variation 1
         */
        VisitorMgr.prototype.listTopVisitors = function listTopVisitors(request, callback) {
            return this.rpcCall(listTopVisitors, $root.common.Id, $root.user.Visitors, request, callback);
        };

        /**
         * Calls ListTopVisitors.
         * @function listTopVisitors
         * @memberof user.VisitorMgr
         * @instance
         * @param {common.IId} request Id message or plain object
         * @returns {Promise<user.Visitors>} Promise
         * @variation 2
         */

        return VisitorMgr;
    })();

    user.UserMgr = (function() {

        /**
         * Constructs a new UserMgr service.
         * @memberof user
         * @classdesc Represents a UserMgr
         * @extends $protobuf.rpc.Service
         * @constructor
         * @param {$protobuf.RPCImpl} rpcImpl RPC implementation
         * @param {boolean} [requestDelimited=false] Whether requests are length-delimited
         * @param {boolean} [responseDelimited=false] Whether responses are length-delimited
         */
        function UserMgr(rpcImpl, requestDelimited, responseDelimited) {
            $protobuf.rpc.Service.call(this, rpcImpl, requestDelimited, responseDelimited);
        }

        (UserMgr.prototype = Object.create($protobuf.rpc.Service.prototype)).constructor = UserMgr;

        /**
         * Callback as used by {@link user.UserMgr#searchUsers}.
         * @memberof user.UserMgr
         * @typedef SearchUsersCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {user.UserSearchResult} [response] UserSearchResult
         */

        /**
         * Calls SearchUsers.
         * @function searchUsers
         * @memberof user.UserMgr
         * @instance
         * @param {user.IUserSearchRequest} request UserSearchRequest message or plain object
         * @param {user.UserMgr.SearchUsersCallback} callback Node-style callback called with the error, if any, and UserSearchResult
         * @returns {undefined}
         * @variation 1
         */
        UserMgr.prototype.searchUsers = function searchUsers(request, callback) {
            return this.rpcCall(searchUsers, $root.user.UserSearchRequest, $root.user.UserSearchResult, request, callback);
        };

        /**
         * Calls SearchUsers.
         * @function searchUsers
         * @memberof user.UserMgr
         * @instance
         * @param {user.IUserSearchRequest} request UserSearchRequest message or plain object
         * @returns {Promise<user.UserSearchResult>} Promise
         * @variation 2
         */

        /**
         * Callback as used by {@link user.UserMgr#subizID}.
         * @memberof user.UserMgr
         * @typedef SubizIDCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {user.SubizIDResponse} [response] SubizIDResponse
         */

        /**
         * Calls SubizID.
         * @function subizID
         * @memberof user.UserMgr
         * @instance
         * @param {user.ISubizIDRequest} request SubizIDRequest message or plain object
         * @param {user.UserMgr.SubizIDCallback} callback Node-style callback called with the error, if any, and SubizIDResponse
         * @returns {undefined}
         * @variation 1
         */
        UserMgr.prototype.subizID = function subizID(request, callback) {
            return this.rpcCall(subizID, $root.user.SubizIDRequest, $root.user.SubizIDResponse, request, callback);
        };

        /**
         * Calls SubizID.
         * @function subizID
         * @memberof user.UserMgr
         * @instance
         * @param {user.ISubizIDRequest} request SubizIDRequest message or plain object
         * @returns {Promise<user.SubizIDResponse>} Promise
         * @variation 2
         */

        /**
         * Callback as used by {@link user.UserMgr#addToMy}.
         * @memberof user.UserMgr
         * @typedef AddToMyCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {common.Empty} [response] Empty
         */

        /**
         * Calls AddToMy.
         * @function addToMy
         * @memberof user.UserMgr
         * @instance
         * @param {user.IAddToMyRequest} request AddToMyRequest message or plain object
         * @param {user.UserMgr.AddToMyCallback} callback Node-style callback called with the error, if any, and Empty
         * @returns {undefined}
         * @variation 1
         */
        UserMgr.prototype.addToMy = function addToMy(request, callback) {
            return this.rpcCall(addToMy, $root.user.AddToMyRequest, $root.common.Empty, request, callback);
        };

        /**
         * Calls AddToMy.
         * @function addToMy
         * @memberof user.UserMgr
         * @instance
         * @param {user.IAddToMyRequest} request AddToMyRequest message or plain object
         * @returns {Promise<common.Empty>} Promise
         * @variation 2
         */

        /**
         * Callback as used by {@link user.UserMgr#createUser}.
         * @memberof user.UserMgr
         * @typedef CreateUserCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {user.User} [response] User
         */

        /**
         * Calls CreateUser.
         * @function createUser
         * @memberof user.UserMgr
         * @instance
         * @param {user.IUser} request User message or plain object
         * @param {user.UserMgr.CreateUserCallback} callback Node-style callback called with the error, if any, and User
         * @returns {undefined}
         * @variation 1
         */
        UserMgr.prototype.createUser = function createUser(request, callback) {
            return this.rpcCall(createUser, $root.user.User, $root.user.User, request, callback);
        };

        /**
         * Calls CreateUser.
         * @function createUser
         * @memberof user.UserMgr
         * @instance
         * @param {user.IUser} request User message or plain object
         * @returns {Promise<user.User>} Promise
         * @variation 2
         */

        /**
         * Callback as used by {@link user.UserMgr#updateUser}.
         * @memberof user.UserMgr
         * @typedef UpdateUserCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {user.User} [response] User
         */

        /**
         * Calls UpdateUser.
         * @function updateUser
         * @memberof user.UserMgr
         * @instance
         * @param {user.IUser} request User message or plain object
         * @param {user.UserMgr.UpdateUserCallback} callback Node-style callback called with the error, if any, and User
         * @returns {undefined}
         * @variation 1
         */
        UserMgr.prototype.updateUser = function updateUser(request, callback) {
            return this.rpcCall(updateUser, $root.user.User, $root.user.User, request, callback);
        };

        /**
         * Calls UpdateUser.
         * @function updateUser
         * @memberof user.UserMgr
         * @instance
         * @param {user.IUser} request User message or plain object
         * @returns {Promise<user.User>} Promise
         * @variation 2
         */

        /**
         * Callback as used by {@link user.UserMgr#readUser}.
         * @memberof user.UserMgr
         * @typedef ReadUserCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {user.User} [response] User
         */

        /**
         * Calls ReadUser.
         * @function readUser
         * @memberof user.UserMgr
         * @instance
         * @param {common.IId} request Id message or plain object
         * @param {user.UserMgr.ReadUserCallback} callback Node-style callback called with the error, if any, and User
         * @returns {undefined}
         * @variation 1
         */
        UserMgr.prototype.readUser = function readUser(request, callback) {
            return this.rpcCall(readUser, $root.common.Id, $root.user.User, request, callback);
        };

        /**
         * Calls ReadUser.
         * @function readUser
         * @memberof user.UserMgr
         * @instance
         * @param {common.IId} request Id message or plain object
         * @returns {Promise<user.User>} Promise
         * @variation 2
         */

        /**
         * Callback as used by {@link user.UserMgr#removeFromMy}.
         * @memberof user.UserMgr
         * @typedef RemoveFromMyCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {common.Empty} [response] Empty
         */

        /**
         * Calls RemoveFromMy.
         * @function removeFromMy
         * @memberof user.UserMgr
         * @instance
         * @param {common.IId} request Id message or plain object
         * @param {user.UserMgr.RemoveFromMyCallback} callback Node-style callback called with the error, if any, and Empty
         * @returns {undefined}
         * @variation 1
         */
        UserMgr.prototype.removeFromMy = function removeFromMy(request, callback) {
            return this.rpcCall(removeFromMy, $root.common.Id, $root.common.Empty, request, callback);
        };

        /**
         * Calls RemoveFromMy.
         * @function removeFromMy
         * @memberof user.UserMgr
         * @instance
         * @param {common.IId} request Id message or plain object
         * @returns {Promise<common.Empty>} Promise
         * @variation 2
         */

        return UserMgr;
    })();

    user.AutomationMgr = (function() {

        /**
         * Constructs a new AutomationMgr service.
         * @memberof user
         * @classdesc Represents an AutomationMgr
         * @extends $protobuf.rpc.Service
         * @constructor
         * @param {$protobuf.RPCImpl} rpcImpl RPC implementation
         * @param {boolean} [requestDelimited=false] Whether requests are length-delimited
         * @param {boolean} [responseDelimited=false] Whether responses are length-delimited
         */
        function AutomationMgr(rpcImpl, requestDelimited, responseDelimited) {
            $protobuf.rpc.Service.call(this, rpcImpl, requestDelimited, responseDelimited);
        }

        (AutomationMgr.prototype = Object.create($protobuf.rpc.Service.prototype)).constructor = AutomationMgr;

        /**
         * Callback as used by {@link user.AutomationMgr#listAutomations}.
         * @memberof user.AutomationMgr
         * @typedef ListAutomationsCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {user.Automations} [response] Automations
         */

        /**
         * Calls ListAutomations.
         * @function listAutomations
         * @memberof user.AutomationMgr
         * @instance
         * @param {common.IId} request Id message or plain object
         * @param {user.AutomationMgr.ListAutomationsCallback} callback Node-style callback called with the error, if any, and Automations
         * @returns {undefined}
         * @variation 1
         */
        AutomationMgr.prototype.listAutomations = function listAutomations(request, callback) {
            return this.rpcCall(listAutomations, $root.common.Id, $root.user.Automations, request, callback);
        };

        /**
         * Calls ListAutomations.
         * @function listAutomations
         * @memberof user.AutomationMgr
         * @instance
         * @param {common.IId} request Id message or plain object
         * @returns {Promise<user.Automations>} Promise
         * @variation 2
         */

        /**
         * Callback as used by {@link user.AutomationMgr#updateAutomation}.
         * @memberof user.AutomationMgr
         * @typedef UpdateAutomationCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {user.Automation} [response] Automation
         */

        /**
         * Calls UpdateAutomation.
         * @function updateAutomation
         * @memberof user.AutomationMgr
         * @instance
         * @param {user.IAutomation} request Automation message or plain object
         * @param {user.AutomationMgr.UpdateAutomationCallback} callback Node-style callback called with the error, if any, and Automation
         * @returns {undefined}
         * @variation 1
         */
        AutomationMgr.prototype.updateAutomation = function updateAutomation(request, callback) {
            return this.rpcCall(updateAutomation, $root.user.Automation, $root.user.Automation, request, callback);
        };

        /**
         * Calls UpdateAutomation.
         * @function updateAutomation
         * @memberof user.AutomationMgr
         * @instance
         * @param {user.IAutomation} request Automation message or plain object
         * @returns {Promise<user.Automation>} Promise
         * @variation 2
         */

        /**
         * Callback as used by {@link user.AutomationMgr#deleteAutomation}.
         * @memberof user.AutomationMgr
         * @typedef DeleteAutomationCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {common.Empty} [response] Empty
         */

        /**
         * Calls DeleteAutomation.
         * @function deleteAutomation
         * @memberof user.AutomationMgr
         * @instance
         * @param {common.IId} request Id message or plain object
         * @param {user.AutomationMgr.DeleteAutomationCallback} callback Node-style callback called with the error, if any, and Empty
         * @returns {undefined}
         * @variation 1
         */
        AutomationMgr.prototype.deleteAutomation = function deleteAutomation(request, callback) {
            return this.rpcCall(deleteAutomation, $root.common.Id, $root.common.Empty, request, callback);
        };

        /**
         * Calls DeleteAutomation.
         * @function deleteAutomation
         * @memberof user.AutomationMgr
         * @instance
         * @param {common.IId} request Id message or plain object
         * @returns {Promise<common.Empty>} Promise
         * @variation 2
         */

        /**
         * Callback as used by {@link user.AutomationMgr#readAutomation}.
         * @memberof user.AutomationMgr
         * @typedef ReadAutomationCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {user.Automation} [response] Automation
         */

        /**
         * Calls ReadAutomation.
         * @function readAutomation
         * @memberof user.AutomationMgr
         * @instance
         * @param {common.IId} request Id message or plain object
         * @param {user.AutomationMgr.ReadAutomationCallback} callback Node-style callback called with the error, if any, and Automation
         * @returns {undefined}
         * @variation 1
         */
        AutomationMgr.prototype.readAutomation = function readAutomation(request, callback) {
            return this.rpcCall(readAutomation, $root.common.Id, $root.user.Automation, request, callback);
        };

        /**
         * Calls ReadAutomation.
         * @function readAutomation
         * @memberof user.AutomationMgr
         * @instance
         * @param {common.IId} request Id message or plain object
         * @returns {Promise<user.Automation>} Promise
         * @variation 2
         */

        /**
         * Callback as used by {@link user.AutomationMgr#createAutomation}.
         * @memberof user.AutomationMgr
         * @typedef CreateAutomationCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {user.Automation} [response] Automation
         */

        /**
         * Calls CreateAutomation.
         * @function createAutomation
         * @memberof user.AutomationMgr
         * @instance
         * @param {user.IAutomation} request Automation message or plain object
         * @param {user.AutomationMgr.CreateAutomationCallback} callback Node-style callback called with the error, if any, and Automation
         * @returns {undefined}
         * @variation 1
         */
        AutomationMgr.prototype.createAutomation = function createAutomation(request, callback) {
            return this.rpcCall(createAutomation, $root.user.Automation, $root.user.Automation, request, callback);
        };

        /**
         * Calls CreateAutomation.
         * @function createAutomation
         * @memberof user.AutomationMgr
         * @instance
         * @param {user.IAutomation} request Automation message or plain object
         * @returns {Promise<user.Automation>} Promise
         * @variation 2
         */

        return AutomationMgr;
    })();

    user.SessionMgr = (function() {

        /**
         * Constructs a new SessionMgr service.
         * @memberof user
         * @classdesc Represents a SessionMgr
         * @extends $protobuf.rpc.Service
         * @constructor
         * @param {$protobuf.RPCImpl} rpcImpl RPC implementation
         * @param {boolean} [requestDelimited=false] Whether requests are length-delimited
         * @param {boolean} [responseDelimited=false] Whether responses are length-delimited
         */
        function SessionMgr(rpcImpl, requestDelimited, responseDelimited) {
            $protobuf.rpc.Service.call(this, rpcImpl, requestDelimited, responseDelimited);
        }

        (SessionMgr.prototype = Object.create($protobuf.rpc.Service.prototype)).constructor = SessionMgr;

        /**
         * Callback as used by {@link user.SessionMgr#createSession}.
         * @memberof user.SessionMgr
         * @typedef CreateSessionCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {user.Session} [response] Session
         */

        /**
         * Calls CreateSession.
         * @function createSession
         * @memberof user.SessionMgr
         * @instance
         * @param {user.ISession} request Session message or plain object
         * @param {user.SessionMgr.CreateSessionCallback} callback Node-style callback called with the error, if any, and Session
         * @returns {undefined}
         * @variation 1
         */
        SessionMgr.prototype.createSession = function createSession(request, callback) {
            return this.rpcCall(createSession, $root.user.Session, $root.user.Session, request, callback);
        };

        /**
         * Calls CreateSession.
         * @function createSession
         * @memberof user.SessionMgr
         * @instance
         * @param {user.ISession} request Session message or plain object
         * @returns {Promise<user.Session>} Promise
         * @variation 2
         */

        /**
         * Callback as used by {@link user.SessionMgr#readSession}.
         * @memberof user.SessionMgr
         * @typedef ReadSessionCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {user.Session} [response] Session
         */

        /**
         * Calls ReadSession.
         * @function readSession
         * @memberof user.SessionMgr
         * @instance
         * @param {user.ISession} request Session message or plain object
         * @param {user.SessionMgr.ReadSessionCallback} callback Node-style callback called with the error, if any, and Session
         * @returns {undefined}
         * @variation 1
         */
        SessionMgr.prototype.readSession = function readSession(request, callback) {
            return this.rpcCall(readSession, $root.user.Session, $root.user.Session, request, callback);
        };

        /**
         * Calls ReadSession.
         * @function readSession
         * @memberof user.SessionMgr
         * @instance
         * @param {user.ISession} request Session message or plain object
         * @returns {Promise<user.Session>} Promise
         * @variation 2
         */

        /**
         * Callback as used by {@link user.SessionMgr#updateSession}.
         * @memberof user.SessionMgr
         * @typedef UpdateSessionCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {user.Session} [response] Session
         */

        /**
         * Calls UpdateSession.
         * @function updateSession
         * @memberof user.SessionMgr
         * @instance
         * @param {user.ISession} request Session message or plain object
         * @param {user.SessionMgr.UpdateSessionCallback} callback Node-style callback called with the error, if any, and Session
         * @returns {undefined}
         * @variation 1
         */
        SessionMgr.prototype.updateSession = function updateSession(request, callback) {
            return this.rpcCall(updateSession, $root.user.Session, $root.user.Session, request, callback);
        };

        /**
         * Calls UpdateSession.
         * @function updateSession
         * @memberof user.SessionMgr
         * @instance
         * @param {user.ISession} request Session message or plain object
         * @returns {Promise<user.Session>} Promise
         * @variation 2
         */

        return SessionMgr;
    })();

    user.AttributeMgr = (function() {

        /**
         * Constructs a new AttributeMgr service.
         * @memberof user
         * @classdesc Represents an AttributeMgr
         * @extends $protobuf.rpc.Service
         * @constructor
         * @param {$protobuf.RPCImpl} rpcImpl RPC implementation
         * @param {boolean} [requestDelimited=false] Whether requests are length-delimited
         * @param {boolean} [responseDelimited=false] Whether responses are length-delimited
         */
        function AttributeMgr(rpcImpl, requestDelimited, responseDelimited) {
            $protobuf.rpc.Service.call(this, rpcImpl, requestDelimited, responseDelimited);
        }

        (AttributeMgr.prototype = Object.create($protobuf.rpc.Service.prototype)).constructor = AttributeMgr;

        /**
         * Callback as used by {@link user.AttributeMgr#listAttributeDefinitions}.
         * @memberof user.AttributeMgr
         * @typedef ListAttributeDefinitionsCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {user.AttributeDefinitions} [response] AttributeDefinitions
         */

        /**
         * Calls ListAttributeDefinitions.
         * @function listAttributeDefinitions
         * @memberof user.AttributeMgr
         * @instance
         * @param {common.IEmpty} request Empty message or plain object
         * @param {user.AttributeMgr.ListAttributeDefinitionsCallback} callback Node-style callback called with the error, if any, and AttributeDefinitions
         * @returns {undefined}
         * @variation 1
         */
        AttributeMgr.prototype.listAttributeDefinitions = function listAttributeDefinitions(request, callback) {
            return this.rpcCall(listAttributeDefinitions, $root.common.Empty, $root.user.AttributeDefinitions, request, callback);
        };

        /**
         * Calls ListAttributeDefinitions.
         * @function listAttributeDefinitions
         * @memberof user.AttributeMgr
         * @instance
         * @param {common.IEmpty} request Empty message or plain object
         * @returns {Promise<user.AttributeDefinitions>} Promise
         * @variation 2
         */

        /**
         * Callback as used by {@link user.AttributeMgr#createAttributeDefinition}.
         * @memberof user.AttributeMgr
         * @typedef CreateAttributeDefinitionCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {user.AttributeDefinition} [response] AttributeDefinition
         */

        /**
         * Calls CreateAttributeDefinition.
         * @function createAttributeDefinition
         * @memberof user.AttributeMgr
         * @instance
         * @param {user.IAttributeDefinition} request AttributeDefinition message or plain object
         * @param {user.AttributeMgr.CreateAttributeDefinitionCallback} callback Node-style callback called with the error, if any, and AttributeDefinition
         * @returns {undefined}
         * @variation 1
         */
        AttributeMgr.prototype.createAttributeDefinition = function createAttributeDefinition(request, callback) {
            return this.rpcCall(createAttributeDefinition, $root.user.AttributeDefinition, $root.user.AttributeDefinition, request, callback);
        };

        /**
         * Calls CreateAttributeDefinition.
         * @function createAttributeDefinition
         * @memberof user.AttributeMgr
         * @instance
         * @param {user.IAttributeDefinition} request AttributeDefinition message or plain object
         * @returns {Promise<user.AttributeDefinition>} Promise
         * @variation 2
         */

        /**
         * Callback as used by {@link user.AttributeMgr#updateAttributeDefinition}.
         * @memberof user.AttributeMgr
         * @typedef UpdateAttributeDefinitionCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {user.AttributeDefinition} [response] AttributeDefinition
         */

        /**
         * Calls UpdateAttributeDefinition.
         * @function updateAttributeDefinition
         * @memberof user.AttributeMgr
         * @instance
         * @param {user.IAttributeDefinition} request AttributeDefinition message or plain object
         * @param {user.AttributeMgr.UpdateAttributeDefinitionCallback} callback Node-style callback called with the error, if any, and AttributeDefinition
         * @returns {undefined}
         * @variation 1
         */
        AttributeMgr.prototype.updateAttributeDefinition = function updateAttributeDefinition(request, callback) {
            return this.rpcCall(updateAttributeDefinition, $root.user.AttributeDefinition, $root.user.AttributeDefinition, request, callback);
        };

        /**
         * Calls UpdateAttributeDefinition.
         * @function updateAttributeDefinition
         * @memberof user.AttributeMgr
         * @instance
         * @param {user.IAttributeDefinition} request AttributeDefinition message or plain object
         * @returns {Promise<user.AttributeDefinition>} Promise
         * @variation 2
         */

        /**
         * Callback as used by {@link user.AttributeMgr#deleteAttributeDefinition}.
         * @memberof user.AttributeMgr
         * @typedef DeleteAttributeDefinitionCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {common.Empty} [response] Empty
         */

        /**
         * Calls DeleteAttributeDefinition.
         * @function deleteAttributeDefinition
         * @memberof user.AttributeMgr
         * @instance
         * @param {user.IDeleteAttrRequest} request DeleteAttrRequest message or plain object
         * @param {user.AttributeMgr.DeleteAttributeDefinitionCallback} callback Node-style callback called with the error, if any, and Empty
         * @returns {undefined}
         * @variation 1
         */
        AttributeMgr.prototype.deleteAttributeDefinition = function deleteAttributeDefinition(request, callback) {
            return this.rpcCall(deleteAttributeDefinition, $root.user.DeleteAttrRequest, $root.common.Empty, request, callback);
        };

        /**
         * Calls DeleteAttributeDefinition.
         * @function deleteAttributeDefinition
         * @memberof user.AttributeMgr
         * @instance
         * @param {user.IDeleteAttrRequest} request DeleteAttrRequest message or plain object
         * @returns {Promise<common.Empty>} Promise
         * @variation 2
         */

        return AttributeMgr;
    })();

    user.AutomationCheck = (function() {

        /**
         * Properties of an AutomationCheck.
         * @memberof user
         * @interface IAutomationCheck
         * @property {string|null} [account_id] AutomationCheck account_id
         * @property {string|null} [automation_id] AutomationCheck automation_id
         * @property {string|null} [user_id] AutomationCheck user_id
         * @property {string|null} [event_id] AutomationCheck event_id
         * @property {string|null} [scope] AutomationCheck scope
         */

        /**
         * Constructs a new AutomationCheck.
         * @memberof user
         * @classdesc Represents an AutomationCheck.
         * @implements IAutomationCheck
         * @constructor
         * @param {user.IAutomationCheck=} [p] Properties to set
         */
        function AutomationCheck(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * AutomationCheck account_id.
         * @member {string} account_id
         * @memberof user.AutomationCheck
         * @instance
         */
        AutomationCheck.prototype.account_id = "";

        /**
         * AutomationCheck automation_id.
         * @member {string} automation_id
         * @memberof user.AutomationCheck
         * @instance
         */
        AutomationCheck.prototype.automation_id = "";

        /**
         * AutomationCheck user_id.
         * @member {string} user_id
         * @memberof user.AutomationCheck
         * @instance
         */
        AutomationCheck.prototype.user_id = "";

        /**
         * AutomationCheck event_id.
         * @member {string} event_id
         * @memberof user.AutomationCheck
         * @instance
         */
        AutomationCheck.prototype.event_id = "";

        /**
         * AutomationCheck scope.
         * @member {string} scope
         * @memberof user.AutomationCheck
         * @instance
         */
        AutomationCheck.prototype.scope = "";

        return AutomationCheck;
    })();

    return user;
})();

export const webhook4 = $root.webhook4 = (() => {

    /**
     * Namespace webhook4.
     * @exports webhook4
     * @namespace
     */
    const webhook4 = {};

    webhook4.Webhook = (function() {

        /**
         * Properties of a Webhook.
         * @memberof webhook4
         * @interface IWebhook
         * @property {common.IContext|null} [ctx] Webhook ctx
         * @property {string|null} [account_id] Webhook account_id
         * @property {string|null} [client_id] Webhook client_id
         * @property {string|null} [url] Webhook url
         * @property {string|null} [secret] Webhook secret
         * @property {Array.<string>|null} [events] Webhook events
         * @property {string|null} [state] Webhook state
         * @property {number|Long|null} [last_hook_at] Webhook last_hook_at
         * @property {number|Long|null} [last_hook_id] Webhook last_hook_id
         * @property {number|Long|null} [events_count] Webhook events_count
         * @property {string|null} [last_hook_response] Webhook last_hook_response
         * @property {number|null} [last_hook_status] Webhook last_hook_status
         * @property {string|null} [last_hook_payload] Webhook last_hook_payload
         * @property {number|Long|null} [backoffs_count] Webhook backoffs_count
         */

        /**
         * Constructs a new Webhook.
         * @memberof webhook4
         * @classdesc Represents a Webhook.
         * @implements IWebhook
         * @constructor
         * @param {webhook4.IWebhook=} [p] Properties to set
         */
        function Webhook(p) {
            this.events = [];
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * Webhook ctx.
         * @member {common.IContext|null|undefined} ctx
         * @memberof webhook4.Webhook
         * @instance
         */
        Webhook.prototype.ctx = null;

        /**
         * Webhook account_id.
         * @member {string} account_id
         * @memberof webhook4.Webhook
         * @instance
         */
        Webhook.prototype.account_id = "";

        /**
         * Webhook client_id.
         * @member {string} client_id
         * @memberof webhook4.Webhook
         * @instance
         */
        Webhook.prototype.client_id = "";

        /**
         * Webhook url.
         * @member {string} url
         * @memberof webhook4.Webhook
         * @instance
         */
        Webhook.prototype.url = "";

        /**
         * Webhook secret.
         * @member {string} secret
         * @memberof webhook4.Webhook
         * @instance
         */
        Webhook.prototype.secret = "";

        /**
         * Webhook events.
         * @member {Array.<string>} events
         * @memberof webhook4.Webhook
         * @instance
         */
        Webhook.prototype.events = $util.emptyArray;

        /**
         * Webhook state.
         * @member {string} state
         * @memberof webhook4.Webhook
         * @instance
         */
        Webhook.prototype.state = "";

        /**
         * Webhook last_hook_at.
         * @member {number|Long} last_hook_at
         * @memberof webhook4.Webhook
         * @instance
         */
        Webhook.prototype.last_hook_at = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

        /**
         * Webhook last_hook_id.
         * @member {number|Long} last_hook_id
         * @memberof webhook4.Webhook
         * @instance
         */
        Webhook.prototype.last_hook_id = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

        /**
         * Webhook events_count.
         * @member {number|Long} events_count
         * @memberof webhook4.Webhook
         * @instance
         */
        Webhook.prototype.events_count = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

        /**
         * Webhook last_hook_response.
         * @member {string} last_hook_response
         * @memberof webhook4.Webhook
         * @instance
         */
        Webhook.prototype.last_hook_response = "";

        /**
         * Webhook last_hook_status.
         * @member {number} last_hook_status
         * @memberof webhook4.Webhook
         * @instance
         */
        Webhook.prototype.last_hook_status = 0;

        /**
         * Webhook last_hook_payload.
         * @member {string} last_hook_payload
         * @memberof webhook4.Webhook
         * @instance
         */
        Webhook.prototype.last_hook_payload = "";

        /**
         * Webhook backoffs_count.
         * @member {number|Long} backoffs_count
         * @memberof webhook4.Webhook
         * @instance
         */
        Webhook.prototype.backoffs_count = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

        /**
         * State enum.
         * @name webhook4.Webhook.State
         * @enum {string}
         * @property {number} normal=0 normal value
         * @property {number} backoff=1 backoff value
         * @property {number} inactive=2 inactive value
         */
        Webhook.State = (function() {
            const valuesById = {}, values = Object.create(valuesById);
            values[valuesById[0] = "normal"] = 0;
            values[valuesById[1] = "backoff"] = 1;
            values[valuesById[2] = "inactive"] = 2;
            return values;
        })();

        return Webhook;
    })();

    /**
     * Event enum.
     * @name webhook4.Event
     * @enum {string}
     * @property {number} Webhook4WebhookReadRequested=3 Webhook4WebhookReadRequested value
     * @property {number} Webhook4WebhookUpserted=4 Webhook4WebhookUpserted value
     * @property {number} Webhook4WebhookVerifyRequested=5 Webhook4WebhookVerifyRequested value
     * @property {number} Webhook4Requested=6 Webhook4Requested value
     * @property {number} Webhook4ClearBackoffRequested=7 Webhook4ClearBackoffRequested value
     * @property {number} Webhook4PurgeQueueRequested=8 Webhook4PurgeQueueRequested value
     * @property {number} Webhook4ShardSend=9 Webhook4ShardSend value
     */
    webhook4.Event = (function() {
        const valuesById = {}, values = Object.create(valuesById);
        values[valuesById[3] = "Webhook4WebhookReadRequested"] = 3;
        values[valuesById[4] = "Webhook4WebhookUpserted"] = 4;
        values[valuesById[5] = "Webhook4WebhookVerifyRequested"] = 5;
        values[valuesById[6] = "Webhook4Requested"] = 6;
        values[valuesById[7] = "Webhook4ClearBackoffRequested"] = 7;
        values[valuesById[8] = "Webhook4PurgeQueueRequested"] = 8;
        values[valuesById[9] = "Webhook4ShardSend"] = 9;
        return values;
    })();

    webhook4.WebhookService = (function() {

        /**
         * Constructs a new WebhookService service.
         * @memberof webhook4
         * @classdesc Represents a WebhookService
         * @extends $protobuf.rpc.Service
         * @constructor
         * @param {$protobuf.RPCImpl} rpcImpl RPC implementation
         * @param {boolean} [requestDelimited=false] Whether requests are length-delimited
         * @param {boolean} [responseDelimited=false] Whether responses are length-delimited
         */
        function WebhookService(rpcImpl, requestDelimited, responseDelimited) {
            $protobuf.rpc.Service.call(this, rpcImpl, requestDelimited, responseDelimited);
        }

        (WebhookService.prototype = Object.create($protobuf.rpc.Service.prototype)).constructor = WebhookService;

        /**
         * Callback as used by {@link webhook4.WebhookService#read}.
         * @memberof webhook4.WebhookService
         * @typedef ReadCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {webhook4.Webhook} [response] Webhook
         */

        /**
         * Calls Read.
         * @function read
         * @memberof webhook4.WebhookService
         * @instance
         * @param {common.IId} request Id message or plain object
         * @param {webhook4.WebhookService.ReadCallback} callback Node-style callback called with the error, if any, and Webhook
         * @returns {undefined}
         * @variation 1
         */
        WebhookService.prototype.read = function read(request, callback) {
            return this.rpcCall(read, $root.common.Id, $root.webhook4.Webhook, request, callback);
        };

        /**
         * Calls Read.
         * @function read
         * @memberof webhook4.WebhookService
         * @instance
         * @param {common.IId} request Id message or plain object
         * @returns {Promise<webhook4.Webhook>} Promise
         * @variation 2
         */

        /**
         * Callback as used by {@link webhook4.WebhookService#clearBackoff}.
         * @memberof webhook4.WebhookService
         * @typedef ClearBackoffCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {common.Empty} [response] Empty
         */

        /**
         * Calls ClearBackoff.
         * @function clearBackoff
         * @memberof webhook4.WebhookService
         * @instance
         * @param {common.IId} request Id message or plain object
         * @param {webhook4.WebhookService.ClearBackoffCallback} callback Node-style callback called with the error, if any, and Empty
         * @returns {undefined}
         * @variation 1
         */
        WebhookService.prototype.clearBackoff = function clearBackoff(request, callback) {
            return this.rpcCall(clearBackoff, $root.common.Id, $root.common.Empty, request, callback);
        };

        /**
         * Calls ClearBackoff.
         * @function clearBackoff
         * @memberof webhook4.WebhookService
         * @instance
         * @param {common.IId} request Id message or plain object
         * @returns {Promise<common.Empty>} Promise
         * @variation 2
         */

        /**
         * Callback as used by {@link webhook4.WebhookService#purgeQueue}.
         * @memberof webhook4.WebhookService
         * @typedef PurgeQueueCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {common.Empty} [response] Empty
         */

        /**
         * Calls PurgeQueue.
         * @function purgeQueue
         * @memberof webhook4.WebhookService
         * @instance
         * @param {common.IId} request Id message or plain object
         * @param {webhook4.WebhookService.PurgeQueueCallback} callback Node-style callback called with the error, if any, and Empty
         * @returns {undefined}
         * @variation 1
         */
        WebhookService.prototype.purgeQueue = function purgeQueue(request, callback) {
            return this.rpcCall(purgeQueue, $root.common.Id, $root.common.Empty, request, callback);
        };

        /**
         * Calls PurgeQueue.
         * @function purgeQueue
         * @memberof webhook4.WebhookService
         * @instance
         * @param {common.IId} request Id message or plain object
         * @returns {Promise<common.Empty>} Promise
         * @variation 2
         */

        return WebhookService;
    })();

    return webhook4;
})();

export const widget = $root.widget = (() => {

    /**
     * Namespace widget.
     * @exports widget
     * @namespace
     */
    const widget = {};

    widget.AllType = (function() {

        /**
         * Properties of an AllType.
         * @memberof widget
         * @interface IAllType
         * @property {widget.ITheme|null} [theme] AllType theme
         * @property {widget.ISound|null} [sound] AllType sound
         * @property {widget.ISetting|null} [setting] AllType setting
         * @property {widget.IUserSetting|null} [us] AllType us
         */

        /**
         * Constructs a new AllType.
         * @memberof widget
         * @classdesc Represents an AllType.
         * @implements IAllType
         * @constructor
         * @param {widget.IAllType=} [p] Properties to set
         */
        function AllType(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * AllType theme.
         * @member {widget.ITheme|null|undefined} theme
         * @memberof widget.AllType
         * @instance
         */
        AllType.prototype.theme = null;

        /**
         * AllType sound.
         * @member {widget.ISound|null|undefined} sound
         * @memberof widget.AllType
         * @instance
         */
        AllType.prototype.sound = null;

        /**
         * AllType setting.
         * @member {widget.ISetting|null|undefined} setting
         * @memberof widget.AllType
         * @instance
         */
        AllType.prototype.setting = null;

        /**
         * AllType us.
         * @member {widget.IUserSetting|null|undefined} us
         * @memberof widget.AllType
         * @instance
         */
        AllType.prototype.us = null;

        return AllType;
    })();

    widget.MyServer = (function() {

        /**
         * Constructs a new MyServer service.
         * @memberof widget
         * @classdesc Represents a MyServer
         * @extends $protobuf.rpc.Service
         * @constructor
         * @param {$protobuf.RPCImpl} rpcImpl RPC implementation
         * @param {boolean} [requestDelimited=false] Whether requests are length-delimited
         * @param {boolean} [responseDelimited=false] Whether responses are length-delimited
         */
        function MyServer(rpcImpl, requestDelimited, responseDelimited) {
            $protobuf.rpc.Service.call(this, rpcImpl, requestDelimited, responseDelimited);
        }

        (MyServer.prototype = Object.create($protobuf.rpc.Service.prototype)).constructor = MyServer;

        /**
         * Callback as used by {@link widget.MyServer#do_}.
         * @memberof widget.MyServer
         * @typedef DoCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {widget.AllType} [response] AllType
         */

        /**
         * Calls Do.
         * @function do
         * @memberof widget.MyServer
         * @instance
         * @param {widget.IAllType} request AllType message or plain object
         * @param {widget.MyServer.DoCallback} callback Node-style callback called with the error, if any, and AllType
         * @returns {undefined}
         * @variation 1
         */
        MyServer.prototype["do"] = function do_(request, callback) {
            return this.rpcCall(do_, $root.widget.AllType, $root.widget.AllType, request, callback);
        };

        /**
         * Calls Do.
         * @function do
         * @memberof widget.MyServer
         * @instance
         * @param {widget.IAllType} request AllType message or plain object
         * @returns {Promise<widget.AllType>} Promise
         * @variation 2
         */

        return MyServer;
    })();

    widget.WidgetService = (function() {

        /**
         * Constructs a new WidgetService service.
         * @memberof widget
         * @classdesc Represents a WidgetService
         * @extends $protobuf.rpc.Service
         * @constructor
         * @param {$protobuf.RPCImpl} rpcImpl RPC implementation
         * @param {boolean} [requestDelimited=false] Whether requests are length-delimited
         * @param {boolean} [responseDelimited=false] Whether responses are length-delimited
         */
        function WidgetService(rpcImpl, requestDelimited, responseDelimited) {
            $protobuf.rpc.Service.call(this, rpcImpl, requestDelimited, responseDelimited);
        }

        (WidgetService.prototype = Object.create($protobuf.rpc.Service.prototype)).constructor = WidgetService;

        /**
         * Callback as used by {@link widget.WidgetService#read}.
         * @memberof widget.WidgetService
         * @typedef ReadCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {widget.Setting} [response] Setting
         */

        /**
         * Calls Read.
         * @function read
         * @memberof widget.WidgetService
         * @instance
         * @param {common.IId} request Id message or plain object
         * @param {widget.WidgetService.ReadCallback} callback Node-style callback called with the error, if any, and Setting
         * @returns {undefined}
         * @variation 1
         */
        WidgetService.prototype.read = function read(request, callback) {
            return this.rpcCall(read, $root.common.Id, $root.widget.Setting, request, callback);
        };

        /**
         * Calls Read.
         * @function read
         * @memberof widget.WidgetService
         * @instance
         * @param {common.IId} request Id message or plain object
         * @returns {Promise<widget.Setting>} Promise
         * @variation 2
         */

        /**
         * Callback as used by {@link widget.WidgetService#update}.
         * @memberof widget.WidgetService
         * @typedef UpdateCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {widget.Setting} [response] Setting
         */

        /**
         * Calls Update.
         * @function update
         * @memberof widget.WidgetService
         * @instance
         * @param {widget.ISetting} request Setting message or plain object
         * @param {widget.WidgetService.UpdateCallback} callback Node-style callback called with the error, if any, and Setting
         * @returns {undefined}
         * @variation 1
         */
        WidgetService.prototype.update = function update(request, callback) {
            return this.rpcCall(update, $root.widget.Setting, $root.widget.Setting, request, callback);
        };

        /**
         * Calls Update.
         * @function update
         * @memberof widget.WidgetService
         * @instance
         * @param {widget.ISetting} request Setting message or plain object
         * @returns {Promise<widget.Setting>} Promise
         * @variation 2
         */

        /**
         * Callback as used by {@link widget.WidgetService#readUserSetting}.
         * @memberof widget.WidgetService
         * @typedef ReadUserSettingCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {widget.UserSetting} [response] UserSetting
         */

        /**
         * Calls ReadUserSetting.
         * @function readUserSetting
         * @memberof widget.WidgetService
         * @instance
         * @param {common.IId} request Id message or plain object
         * @param {widget.WidgetService.ReadUserSettingCallback} callback Node-style callback called with the error, if any, and UserSetting
         * @returns {undefined}
         * @variation 1
         */
        WidgetService.prototype.readUserSetting = function readUserSetting(request, callback) {
            return this.rpcCall(readUserSetting, $root.common.Id, $root.widget.UserSetting, request, callback);
        };

        /**
         * Calls ReadUserSetting.
         * @function readUserSetting
         * @memberof widget.WidgetService
         * @instance
         * @param {common.IId} request Id message or plain object
         * @returns {Promise<widget.UserSetting>} Promise
         * @variation 2
         */

        /**
         * Callback as used by {@link widget.WidgetService#updateUserSetting}.
         * @memberof widget.WidgetService
         * @typedef UpdateUserSettingCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {widget.UserSetting} [response] UserSetting
         */

        /**
         * Calls UpdateUserSetting.
         * @function updateUserSetting
         * @memberof widget.WidgetService
         * @instance
         * @param {widget.IUserSetting} request UserSetting message or plain object
         * @param {widget.WidgetService.UpdateUserSettingCallback} callback Node-style callback called with the error, if any, and UserSetting
         * @returns {undefined}
         * @variation 1
         */
        WidgetService.prototype.updateUserSetting = function updateUserSetting(request, callback) {
            return this.rpcCall(updateUserSetting, $root.widget.UserSetting, $root.widget.UserSetting, request, callback);
        };

        /**
         * Calls UpdateUserSetting.
         * @function updateUserSetting
         * @memberof widget.WidgetService
         * @instance
         * @param {widget.IUserSetting} request UserSetting message or plain object
         * @returns {Promise<widget.UserSetting>} Promise
         * @variation 2
         */

        return WidgetService;
    })();

    widget.Theme = (function() {

        /**
         * Properties of a Theme.
         * @memberof widget
         * @interface ITheme
         * @property {string|null} [account_id] Theme account_id
         * @property {string|null} [custom_css] Theme custom_css
         * @property {string|null} [widget_position] Theme widget_position
         * @property {string|null} [window_mode] Theme window_mode
         */

        /**
         * Constructs a new Theme.
         * @memberof widget
         * @classdesc Represents a Theme.
         * @implements ITheme
         * @constructor
         * @param {widget.ITheme=} [p] Properties to set
         */
        function Theme(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * Theme account_id.
         * @member {string} account_id
         * @memberof widget.Theme
         * @instance
         */
        Theme.prototype.account_id = "";

        /**
         * Theme custom_css.
         * @member {string} custom_css
         * @memberof widget.Theme
         * @instance
         */
        Theme.prototype.custom_css = "";

        /**
         * Theme widget_position.
         * @member {string} widget_position
         * @memberof widget.Theme
         * @instance
         */
        Theme.prototype.widget_position = "";

        /**
         * Theme window_mode.
         * @member {string} window_mode
         * @memberof widget.Theme
         * @instance
         */
        Theme.prototype.window_mode = "";

        /**
         * ButtonPosition enum.
         * @name widget.Theme.ButtonPosition
         * @enum {string}
         * @property {number} left=0 left value
         * @property {number} right=1 right value
         */
        Theme.ButtonPosition = (function() {
            const valuesById = {}, values = Object.create(valuesById);
            values[valuesById[0] = "left"] = 0;
            values[valuesById[1] = "right"] = 1;
            return values;
        })();

        /**
         * WindowMode enum.
         * @name widget.Theme.WindowMode
         * @enum {string}
         * @property {number} mini=0 mini value
         * @property {number} full=1 full value
         */
        Theme.WindowMode = (function() {
            const valuesById = {}, values = Object.create(valuesById);
            values[valuesById[0] = "mini"] = 0;
            values[valuesById[1] = "full"] = 1;
            return values;
        })();

        return Theme;
    })();

    widget.Sound = (function() {

        /**
         * Properties of a Sound.
         * @memberof widget
         * @interface ISound
         * @property {string|null} [account_id] Sound account_id
         * @property {boolean|null} [enabled] Sound enabled
         * @property {string|null} [new_conversation] Sound new_conversation
         * @property {string|null} [file_create] Sound file_create
         * @property {string|null} [new_message] Sound new_message
         * @property {string|null} [message_send_failed] Sound message_send_failed
         */

        /**
         * Constructs a new Sound.
         * @memberof widget
         * @classdesc Represents a Sound.
         * @implements ISound
         * @constructor
         * @param {widget.ISound=} [p] Properties to set
         */
        function Sound(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * Sound account_id.
         * @member {string} account_id
         * @memberof widget.Sound
         * @instance
         */
        Sound.prototype.account_id = "";

        /**
         * Sound enabled.
         * @member {boolean} enabled
         * @memberof widget.Sound
         * @instance
         */
        Sound.prototype.enabled = false;

        /**
         * Sound new_conversation.
         * @member {string} new_conversation
         * @memberof widget.Sound
         * @instance
         */
        Sound.prototype.new_conversation = "";

        /**
         * Sound file_create.
         * @member {string} file_create
         * @memberof widget.Sound
         * @instance
         */
        Sound.prototype.file_create = "";

        /**
         * Sound new_message.
         * @member {string} new_message
         * @memberof widget.Sound
         * @instance
         */
        Sound.prototype.new_message = "";

        /**
         * Sound message_send_failed.
         * @member {string} message_send_failed
         * @memberof widget.Sound
         * @instance
         */
        Sound.prototype.message_send_failed = "";

        return Sound;
    })();

    widget.Setting = (function() {

        /**
         * Properties of a Setting.
         * @memberof widget
         * @interface ISetting
         * @property {common.IContext|null} [ctx] Setting ctx
         * @property {string|null} [account_id] Setting account_id
         * @property {string|null} [widget_version] Setting widget_version
         * @property {widget.ISound|null} [sound] Setting sound
         * @property {string|null} [language] Setting language
         * @property {widget.ITheme|null} [theme] Setting theme
         * @property {number|null} [replytime] Setting replytime
         * @property {Array.<account.IAgent>|null} [agents] Setting agents
         * @property {Array.<string>|null} [agent_ids] Setting agent_ids
         * @property {string|null} [language_url] Setting language_url
         * @property {string|null} [custom_language_url] Setting custom_language_url
         * @property {string|null} [css_url] Setting css_url
         * @property {string|null} [custom_css_url] Setting custom_css_url
         * @property {string|null} [custom_language] Setting custom_language
         */

        /**
         * Constructs a new Setting.
         * @memberof widget
         * @classdesc Represents a Setting.
         * @implements ISetting
         * @constructor
         * @param {widget.ISetting=} [p] Properties to set
         */
        function Setting(p) {
            this.agents = [];
            this.agent_ids = [];
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * Setting ctx.
         * @member {common.IContext|null|undefined} ctx
         * @memberof widget.Setting
         * @instance
         */
        Setting.prototype.ctx = null;

        /**
         * Setting account_id.
         * @member {string} account_id
         * @memberof widget.Setting
         * @instance
         */
        Setting.prototype.account_id = "";

        /**
         * Setting widget_version.
         * @member {string} widget_version
         * @memberof widget.Setting
         * @instance
         */
        Setting.prototype.widget_version = "";

        /**
         * Setting sound.
         * @member {widget.ISound|null|undefined} sound
         * @memberof widget.Setting
         * @instance
         */
        Setting.prototype.sound = null;

        /**
         * Setting language.
         * @member {string} language
         * @memberof widget.Setting
         * @instance
         */
        Setting.prototype.language = "";

        /**
         * Setting theme.
         * @member {widget.ITheme|null|undefined} theme
         * @memberof widget.Setting
         * @instance
         */
        Setting.prototype.theme = null;

        /**
         * Setting replytime.
         * @member {number} replytime
         * @memberof widget.Setting
         * @instance
         */
        Setting.prototype.replytime = 0;

        /**
         * Setting agents.
         * @member {Array.<account.IAgent>} agents
         * @memberof widget.Setting
         * @instance
         */
        Setting.prototype.agents = $util.emptyArray;

        /**
         * Setting agent_ids.
         * @member {Array.<string>} agent_ids
         * @memberof widget.Setting
         * @instance
         */
        Setting.prototype.agent_ids = $util.emptyArray;

        /**
         * Setting language_url.
         * @member {string} language_url
         * @memberof widget.Setting
         * @instance
         */
        Setting.prototype.language_url = "";

        /**
         * Setting custom_language_url.
         * @member {string} custom_language_url
         * @memberof widget.Setting
         * @instance
         */
        Setting.prototype.custom_language_url = "";

        /**
         * Setting css_url.
         * @member {string} css_url
         * @memberof widget.Setting
         * @instance
         */
        Setting.prototype.css_url = "";

        /**
         * Setting custom_css_url.
         * @member {string} custom_css_url
         * @memberof widget.Setting
         * @instance
         */
        Setting.prototype.custom_css_url = "";

        /**
         * Setting custom_language.
         * @member {string} custom_language
         * @memberof widget.Setting
         * @instance
         */
        Setting.prototype.custom_language = "";

        return Setting;
    })();

    widget.UserSetting = (function() {

        /**
         * Properties of a UserSetting.
         * @memberof widget
         * @interface IUserSetting
         * @property {common.IContext|null} [ctx] UserSetting ctx
         * @property {account.IAccount|null} [account] UserSetting account
         * @property {string|null} [account_id] UserSetting account_id
         * @property {user.IUser|null} [user] UserSetting user
         * @property {string|null} [user_id] UserSetting user_id
         * @property {boolean|null} [sound_enabled] UserSetting sound_enabled
         * @property {string|null} [language] UserSetting language
         * @property {boolean|null} [send_transcript] UserSetting send_transcript
         * @property {widget.ISetting|null} [account_setting] UserSetting account_setting
         * @property {boolean|null} [desktop_notification] UserSetting desktop_notification
         */

        /**
         * Constructs a new UserSetting.
         * @memberof widget
         * @classdesc Represents a UserSetting.
         * @implements IUserSetting
         * @constructor
         * @param {widget.IUserSetting=} [p] Properties to set
         */
        function UserSetting(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * UserSetting ctx.
         * @member {common.IContext|null|undefined} ctx
         * @memberof widget.UserSetting
         * @instance
         */
        UserSetting.prototype.ctx = null;

        /**
         * UserSetting account.
         * @member {account.IAccount|null|undefined} account
         * @memberof widget.UserSetting
         * @instance
         */
        UserSetting.prototype.account = null;

        /**
         * UserSetting account_id.
         * @member {string} account_id
         * @memberof widget.UserSetting
         * @instance
         */
        UserSetting.prototype.account_id = "";

        /**
         * UserSetting user.
         * @member {user.IUser|null|undefined} user
         * @memberof widget.UserSetting
         * @instance
         */
        UserSetting.prototype.user = null;

        /**
         * UserSetting user_id.
         * @member {string} user_id
         * @memberof widget.UserSetting
         * @instance
         */
        UserSetting.prototype.user_id = "";

        /**
         * UserSetting sound_enabled.
         * @member {boolean} sound_enabled
         * @memberof widget.UserSetting
         * @instance
         */
        UserSetting.prototype.sound_enabled = false;

        /**
         * UserSetting language.
         * @member {string} language
         * @memberof widget.UserSetting
         * @instance
         */
        UserSetting.prototype.language = "";

        /**
         * UserSetting send_transcript.
         * @member {boolean} send_transcript
         * @memberof widget.UserSetting
         * @instance
         */
        UserSetting.prototype.send_transcript = false;

        /**
         * UserSetting account_setting.
         * @member {widget.ISetting|null|undefined} account_setting
         * @memberof widget.UserSetting
         * @instance
         */
        UserSetting.prototype.account_setting = null;

        /**
         * UserSetting desktop_notification.
         * @member {boolean} desktop_notification
         * @memberof widget.UserSetting
         * @instance
         */
        UserSetting.prototype.desktop_notification = false;

        return UserSetting;
    })();

    widget.Global = (function() {

        /**
         * Properties of a Global.
         * @memberof widget
         * @interface IGlobal
         * @property {common.IContext|null} [ctx] Global ctx
         * @property {string|null} [name] Global name
         * @property {string|null} [data] Global data
         * @property {string|null} [pk] Global pk
         */

        /**
         * Constructs a new Global.
         * @memberof widget
         * @classdesc Represents a Global.
         * @implements IGlobal
         * @constructor
         * @param {widget.IGlobal=} [p] Properties to set
         */
        function Global(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * Global ctx.
         * @member {common.IContext|null|undefined} ctx
         * @memberof widget.Global
         * @instance
         */
        Global.prototype.ctx = null;

        /**
         * Global name.
         * @member {string} name
         * @memberof widget.Global
         * @instance
         */
        Global.prototype.name = "";

        /**
         * Global data.
         * @member {string} data
         * @memberof widget.Global
         * @instance
         */
        Global.prototype.data = "";

        /**
         * Global pk.
         * @member {string} pk
         * @memberof widget.Global
         * @instance
         */
        Global.prototype.pk = "";

        return Global;
    })();

    /**
     * Event enum.
     * @name widget.Event
     * @enum {string}
     * @property {number} WidgetUserSettingReadRequested=0 WidgetUserSettingReadRequested value
     * @property {number} WidgetUserSettingUpdateRequested=1 WidgetUserSettingUpdateRequested value
     * @property {number} WidgetSettingReadRequested=2 WidgetSettingReadRequested value
     * @property {number} WidgetSettingUpdateRequested=3 WidgetSettingUpdateRequested value
     * @property {number} WidgetSettingCssVersionUpdated=4 WidgetSettingCssVersionUpdated value
     * @property {number} WidgetSettingLanguageUpdated=5 WidgetSettingLanguageUpdated value
     * @property {number} WidgetSettingCssVersionUpdateRequested=6 WidgetSettingCssVersionUpdateRequested value
     * @property {number} WidgetSettingLanguageUpdateRequested=7 WidgetSettingLanguageUpdateRequested value
     * @property {number} WidgetRequested=8 WidgetRequested value
     * @property {number} WidgetSettingUpserted=10 WidgetSettingUpserted value
     * @property {number} WidgetSynced=100 WidgetSynced value
     * @property {number} WidgetV3Synced=103 WidgetV3Synced value
     */
    widget.Event = (function() {
        const valuesById = {}, values = Object.create(valuesById);
        values[valuesById[0] = "WidgetUserSettingReadRequested"] = 0;
        values[valuesById[1] = "WidgetUserSettingUpdateRequested"] = 1;
        values[valuesById[2] = "WidgetSettingReadRequested"] = 2;
        values[valuesById[3] = "WidgetSettingUpdateRequested"] = 3;
        values[valuesById[4] = "WidgetSettingCssVersionUpdated"] = 4;
        values[valuesById[5] = "WidgetSettingLanguageUpdated"] = 5;
        values[valuesById[6] = "WidgetSettingCssVersionUpdateRequested"] = 6;
        values[valuesById[7] = "WidgetSettingLanguageUpdateRequested"] = 7;
        values[valuesById[8] = "WidgetRequested"] = 8;
        values[valuesById[10] = "WidgetSettingUpserted"] = 10;
        values[valuesById[100] = "WidgetSynced"] = 100;
        values[valuesById[103] = "WidgetV3Synced"] = 103;
        return values;
    })();

    return widget;
})();

export const ws = $root.ws = (() => {

    /**
     * Namespace ws.
     * @exports ws
     * @namespace
     */
    const ws = {};

    ws.AllType = (function() {

        /**
         * Properties of an AllType.
         * @memberof ws
         * @interface IAllType
         * @property {ws.ISubscribe|null} [sub] AllType sub
         */

        /**
         * Constructs a new AllType.
         * @memberof ws
         * @classdesc Represents an AllType.
         * @implements IAllType
         * @constructor
         * @param {ws.IAllType=} [p] Properties to set
         */
        function AllType(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * AllType sub.
         * @member {ws.ISubscribe|null|undefined} sub
         * @memberof ws.AllType
         * @instance
         */
        AllType.prototype.sub = null;

        return AllType;
    })();

    ws.MyServer = (function() {

        /**
         * Constructs a new MyServer service.
         * @memberof ws
         * @classdesc Represents a MyServer
         * @extends $protobuf.rpc.Service
         * @constructor
         * @param {$protobuf.RPCImpl} rpcImpl RPC implementation
         * @param {boolean} [requestDelimited=false] Whether requests are length-delimited
         * @param {boolean} [responseDelimited=false] Whether responses are length-delimited
         */
        function MyServer(rpcImpl, requestDelimited, responseDelimited) {
            $protobuf.rpc.Service.call(this, rpcImpl, requestDelimited, responseDelimited);
        }

        (MyServer.prototype = Object.create($protobuf.rpc.Service.prototype)).constructor = MyServer;

        /**
         * Callback as used by {@link ws.MyServer#do_}.
         * @memberof ws.MyServer
         * @typedef DoCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {ws.AllType} [response] AllType
         */

        /**
         * Calls Do.
         * @function do
         * @memberof ws.MyServer
         * @instance
         * @param {ws.IAllType} request AllType message or plain object
         * @param {ws.MyServer.DoCallback} callback Node-style callback called with the error, if any, and AllType
         * @returns {undefined}
         * @variation 1
         */
        MyServer.prototype["do"] = function do_(request, callback) {
            return this.rpcCall(do_, $root.ws.AllType, $root.ws.AllType, request, callback);
        };

        /**
         * Calls Do.
         * @function do
         * @memberof ws.MyServer
         * @instance
         * @param {ws.IAllType} request AllType message or plain object
         * @returns {Promise<ws.AllType>} Promise
         * @variation 2
         */

        return MyServer;
    })();

    ws.Subscribe = (function() {

        /**
         * Properties of a Subscribe.
         * @memberof ws
         * @interface ISubscribe
         * @property {Array.<string>|null} [events] Subscribe events
         * @property {string|null} [connection_id] Subscribe connection_id
         */

        /**
         * Constructs a new Subscribe.
         * @memberof ws
         * @classdesc Represents a Subscribe.
         * @implements ISubscribe
         * @constructor
         * @param {ws.ISubscribe=} [p] Properties to set
         */
        function Subscribe(p) {
            this.events = [];
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * Subscribe events.
         * @member {Array.<string>} events
         * @memberof ws.Subscribe
         * @instance
         */
        Subscribe.prototype.events = $util.emptyArray;

        /**
         * Subscribe connection_id.
         * @member {string} connection_id
         * @memberof ws.Subscribe
         * @instance
         */
        Subscribe.prototype.connection_id = "";

        return Subscribe;
    })();

    ws.ListRequest = (function() {

        /**
         * Properties of a ListRequest.
         * @memberof ws
         * @interface IListRequest
         * @property {common.IContext|null} [ctx] ListRequest ctx
         * @property {string|null} [connection_id] ListRequest connection_id
         */

        /**
         * Constructs a new ListRequest.
         * @memberof ws
         * @classdesc Represents a ListRequest.
         * @implements IListRequest
         * @constructor
         * @param {ws.IListRequest=} [p] Properties to set
         */
        function ListRequest(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * ListRequest ctx.
         * @member {common.IContext|null|undefined} ctx
         * @memberof ws.ListRequest
         * @instance
         */
        ListRequest.prototype.ctx = null;

        /**
         * ListRequest connection_id.
         * @member {string} connection_id
         * @memberof ws.ListRequest
         * @instance
         */
        ListRequest.prototype.connection_id = "";

        return ListRequest;
    })();

    /**
     * Event enum.
     * @name ws.Event
     * @enum {string}
     * @property {number} WsSynced=0 WsSynced value
     * @property {number} WsRequested=1 WsRequested value
     * @property {number} WsSend=10 WsSend value
     */
    ws.Event = (function() {
        const valuesById = {}, values = Object.create(valuesById);
        values[valuesById[0] = "WsSynced"] = 0;
        values[valuesById[1] = "WsRequested"] = 1;
        values[valuesById[10] = "WsSend"] = 10;
        return values;
    })();

    ws.Payload = (function() {

        /**
         * Properties of a Payload.
         * @memberof ws
         * @interface IPayload
         * @property {number|Long|null} [id] Payload id
         * @property {string|null} [message] Payload message
         * @property {string|null} [error] Payload error
         */

        /**
         * Constructs a new Payload.
         * @memberof ws
         * @classdesc Represents a Payload.
         * @implements IPayload
         * @constructor
         * @param {ws.IPayload=} [p] Properties to set
         */
        function Payload(p) {
            if (p)
                for (var ks = Object.keys(p), i = 0; i < ks.length; ++i)
                    if (p[ks[i]] != null)
                        this[ks[i]] = p[ks[i]];
        }

        /**
         * Payload id.
         * @member {number|Long} id
         * @memberof ws.Payload
         * @instance
         */
        Payload.prototype.id = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

        /**
         * Payload message.
         * @member {string} message
         * @memberof ws.Payload
         * @instance
         */
        Payload.prototype.message = "";

        /**
         * Payload error.
         * @member {string} error
         * @memberof ws.Payload
         * @instance
         */
        Payload.prototype.error = "";

        return Payload;
    })();

    return ws;
})();

export { $root as default };
