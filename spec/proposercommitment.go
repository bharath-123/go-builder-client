package spec

type ProposerCommitment uint64

const (
	FULL_BLOCK    ProposerCommitment = 0
	TOB_ROB_SPLIT ProposerCommitment = 1
)

var PROPOSER_COMMITMENT_NAMES = map[ProposerCommitment]string{
	FULL_BLOCK:    "full_block",
	TOB_ROB_SPLIT: "tob_rob_split",
}

var VALID_PROPOSER_COMMITMENTS = [2]ProposerCommitment{
	FULL_BLOCK,
	TOB_ROB_SPLIT,
}
