import * as permfilef from "./perm.json" with { type: "json" };
let permfile = permfilef.default;

let ScopeM = {}; // map[string]map[string]bool{}

function _setPrimaryScope(k) {
  if (!ScopeM[k]) ScopeM[k] = {}; // map[string]bool{}
  ScopeM[k][k] = true;
}

function expandScope(permfile, scope) {
  let permStruct = permfile[scope];

  if (!ScopeM[scope]) ScopeM[scope] = {}; // map[string]bool{}
  let ss = (permStruct.scope || "").split(" ");
  lo_map(ss, (subScope) => {
    if (subScope == "") return;
    expandScope(permfile, subScope);
    lo_map(ScopeM[subScope], (v, k) => {
      ScopeM[scope][k] = v;
    });
  });
}

function initscope() {
  let objM = {}; // map[string]bool{} // list of object
  lo_map(permfile, (p, scope) => {
    if (!ScopeM[scope]) ScopeM[scope] = {}; // map[string]bool{}
    lo_map(p.perm, (str, obj) => {
      lo_map(str.split(" "), (k) => {
        if (k == "") return;
        if (!ScopeM[scope]) ScopeM[scope] = {}; // map[string]bool{}
        ScopeM[scope][obj + ":" + k] = true;
        objM[obj] = true; // ticket, ticket_type, account...
      });
    });
  });

  lo_map(permfile, (p, scope) => {
    expandScope(permfile, scope);
  });

  lo_map(objM, (_, obj) => {
    _setPrimaryScope(obj + ":read");
    _setPrimaryScope(obj + ":read:all");
    _setPrimaryScope(obj + ":read:own");
    _setPrimaryScope(obj + ":read:unassigned");
    _setPrimaryScope(obj + ":read:none");

    _setPrimaryScope(obj + ":import");
    _setPrimaryScope(obj + ":import:all");
    _setPrimaryScope(obj + ":import:own");
    _setPrimaryScope(obj + ":import:unassigned");
    _setPrimaryScope(obj + ":import:none");

    _setPrimaryScope(obj + ":invite");
    _setPrimaryScope(obj + ":invite:all");
    _setPrimaryScope(obj + ":invite:own");
    _setPrimaryScope(obj + ":invite:unassigned");
    _setPrimaryScope(obj + ":invite:none");

    _setPrimaryScope(obj + ":update");
    _setPrimaryScope(obj + ":update:all");
    _setPrimaryScope(obj + ":update:own");
    _setPrimaryScope(obj + ":update:unassigned");
    _setPrimaryScope(obj + ":update:none");

    _setPrimaryScope(obj + ":create");
    _setPrimaryScope(obj + ":create:all");
    _setPrimaryScope(obj + ":create:own");
    _setPrimaryScope(obj + ":create:unassigned");
    _setPrimaryScope(obj + ":create:none");

    _setPrimaryScope(obj + ":delete");
    _setPrimaryScope(obj + ":delete:all");
    _setPrimaryScope(obj + ":delete:own");
    _setPrimaryScope(obj + ":delete:unassigned");
    _setPrimaryScope(obj + ":delete:none");
  });
}

function lo_map(collection, iteratee) {
  const result = [];

  if (Array.isArray(collection)) {
    for (let i = 0; i < collection.length; i++) {
      result.push(iteratee(collection[i], i, collection));
    }
  } else if (collection && typeof collection === "object") {
    for (const key in collection) {
      if (Object.prototype.hasOwnProperty.call(collection, key)) {
        result.push(iteratee(collection[key], key, collection));
      }
    }
  }

  return result;
}

function joinMap(a, b) {
  lo_map(b, (v, k) => {
    if (v) a[k] = v;
  });
}

initscope();

// objectType: 'user', 'ticket'
// action: 'update' 'export'
// return true/false: allow access or not
// examples:
//   accessFeature(["agent"], "user", "export") => true
//   accessFeature(["account_manage"], "user", "export") => true
//   accessFeature(["agent", "view_other_convos"], "conversation", "view") => true
//   accessFeature(["agent"], "bot", "update") => false
export var accessFeature = (agentscopes, objectType, action) => {
  if (action == "") return true;

  let permM = {}; // map[string]bool{}
  lo_map(agentscopes, (scope) => {
    // agent's account-wide scope
    joinMap(permM, ScopeM[scope]);
  });

  // ticket:read ticket:write
  if (permM[objectType + ":" + action + ":none"]) return false;

  if (
    permM[objectType + ":" + action] ||
    permM[objectType + ":" + action + ":all"]
  )
    return true;
  return false;
};
