// Subiz Block step contract v1. See BlockStep/BlockStepBatch in header.proto.
// Reference fixtures only; no editor adapter or server apply implementation.
// ProseMirror clients translate their transactions into this format.
// Neither ProseMirror positions, marks, slices nor stepType cross this boundary.
// Embedded blocks use header.Block protobuf JSON (not ProseMirror Node JSON).

/**
 * @typedef {Object} Block Native header.Block; fields are defined in header.proto.
 * @typedef {{type: "splice_text", path: number[], splice_text: {offset: number, delete_count: number, text?: string}}} TextStep
 * @typedef {{type: "splice_content", path: number[], splice_content: {index: number, delete_count: number, blocks?: Block[]}}} ContentStep
 * @typedef {{type: "set_properties", path: number[], set_properties: {fields: string[], values: Block}}} PropertiesStep
 * @typedef {TextStep | ContentStep | PropertiesStep} Step
 *
 * type selects the operation; set only its matching payload.
 * [] is root; [0,1] is root.content[0].content[1].
 * Resolve paths after all prior steps, without implicit text-node normalization.
 * Text offsets count Unicode scalar values: Array.from(text).length / Go runes.
 * Thus "A😀B" has length 3 here, unlike JavaScript string.length (4).
 * Do not split user-visible grapheme clusters when translating UI selections.
 * All ranges must fit; invalid input rejects the batch, never clamps offsets.
 * Properties use exact top-level proto names; values use protobuf JSON names.
 * Use protojson for transport. An omitted empty path/blocks/text has its default
 * meaning; an explicit zero offset/index/delete_count must retain presence.
 */

// Examples are independent unless explicitly grouped in a sequence below.
var documentBefore = {
  type: "div",
  content: [{ type: "paragraph", content: [{ type: "text", text: "Xin chao" }] }],
};

// Change one character: "Xin chao" -> "Xin chào". No paragraph retransmission.
/** @type {Step} */
var replaceCharacter = {
  type: "splice_text",
  path: [0, 0],
  splice_text: { offset: 6, delete_count: 1, text: "à" },
};

/** @type {Step} */
var insertText = {
  type: "splice_text",
  path: [0, 0],
  splice_text: { offset: 8, delete_count: 0, text: "!" },
};

/** @type {Step} */
var deleteText = {
  type: "splice_text",
  path: [0, 0],
  splice_text: { offset: 7, delete_count: 1 },
};

/** @type {Step} */
var boldWholeText = {
  type: "set_properties",
  path: [0, 0],
  set_properties: { fields: ["bold"], values: { bold: true } },
};

// A listed property omitted from values resets to its protobuf default.
/** @type {Step} */
var removeBold = {
  type: "set_properties",
  path: [0, 0],
  set_properties: { fields: ["bold"], values: {} },
};

// Bold only "Xin": replace one text run with two native Block text runs.
/** @type {Step} */
var boldRange = {
  type: "splice_content",
  path: [0],
  splice_content: {
    index: 0,
    delete_count: 1,
    blocks: [
      { type: "text", text: "Xin", bold: true },
      { type: "text", text: " chao" },
    ],
  },
};

/** @type {Step} */
var insertParagraph = {
  type: "splice_content",
  path: [],
  splice_content: {
    index: 1,
    delete_count: 0,
    blocks: [{ type: "paragraph", content: [{ type: "text", text: "Subiz" }] }],
  },
};

/** @type {Step} */
var deleteParagraph = {
  type: "splice_content",
  path: [],
  splice_content: { index: 0, delete_count: 1 },
};

/** @type {Step} */
var changeToHeading = {
  type: "set_properties",
  path: [0],
  set_properties: { fields: ["type", "level"], values: { type: "heading", level: "2" } },
};
// level is int64 in header.Block: canonical protojson encodes it as a string.

// Enter after "Xin ": replace one paragraph with two, in one operation.
/** @type {Step} */
var splitParagraph = {
  type: "splice_content",
  path: [],
  splice_content: {
    index: 0,
    delete_count: 1,
    blocks: [
      { type: "paragraph", content: [{ type: "text", text: "Xin " }] },
      { type: "paragraph", content: [{ type: "text", text: "chao" }] },
    ],
  },
};

// Sequential example: boldRange changes paths; second run is now [0,1].
var formatThenEdit = [
  boldRange,
  { type: "splice_text", path: [0, 1], splice_text: { offset: 3, delete_count: 1, text: "à" } },
];

/**
 * @typedef {{document_id: string, schema_version: 1, base_revision: string,
 * client_id: string, batch_id: string, steps: Step[]}} StepBatch
 * Revision belongs to the server, not to ProseMirror or a step counter.
 * Apply/persist the entire batch atomically only on matching base_revision.
 * Retry with the same batch_id + identical payload; changed payload needs a new ID.
 * Deduplicate before revision check. Confirm only the acknowledged batch;
 * preserve edits made while it was in flight. Send dependent batches in order.
 * A conflict requires reconciliation; these steps do not implement OT/CRDT.
 * Stock prosemirror-collab cannot consume this wire format directly.
 */

/** @type {StepBatch} */
var saveChanges = {
  document_id: "doc_abc",
  schema_version: 1,
  base_revision: "42",
  client_id: "editor_A",
  batch_id: "batch_105",
  steps: [replaceCharacter],
};

var saveResult = {
  document_id: "doc_abc",
  client_id: "editor_A",
  batch_id: "batch_105",
  revision: "43",
};
