package spec

type ProposerCommitment string

const (
	FULL_BLOCK    ProposerCommitment = "full_block"
	TOB_ROB_SPLIT ProposerCommitment = "tob_rob_split"
)

var VALID_PROPOSER_COMMITMENTS = [2]ProposerCommitment{
	FULL_BLOCK,
	TOB_ROB_SPLIT,
}
