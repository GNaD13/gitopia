package utils

import (
	"fmt"

	"github.com/gitopia/gitopia/x/gitopia/types"
)

func AddAssigneesCommentBody(creator string, assignees []string) string {
	return fmt.Sprintf("@%v assigned to"+JoinAssignees(assignees), creator)
}

func RemoveAssigneesCommentBody(creator string, assignees []string) string {
	return fmt.Sprintf("@%v unassigned"+JoinAssignees(assignees), creator)
}

func AddLabelsCommentBody(creator string, labels []string) string {
	return fmt.Sprintf("@%v added"+JoinLabels(labels)+" label", creator)
}

func RemoveLabelsCommentBody(creator string, labels []string) string {
	return fmt.Sprintf("@%v removed"+JoinLabels(labels)+" label", creator)
}

func UpdateTitleCommentBody(creator string, oldTitle string, newTitle string) string {
	return fmt.Sprintf("@%v changed title from **~~%v~~** to **%v**", creator, oldTitle, newTitle)
}

func UpdateDescriptionCommentBody(creator string) string {
	return fmt.Sprintf("@%v changed the description", creator)
}

func IssueToggleStateCommentBody(creator string, state types.Issue_State) string {
	if state == types.Issue_OPEN {
		return fmt.Sprintf("@%v reopened", creator)
	} else {
		return fmt.Sprintf("@%v closed", creator)
	}
}

func PullRequestToggleStateCommentBody(creator string, state types.PullRequest_State) string {
	if state == types.PullRequest_OPEN {
		return fmt.Sprintf("@%v reopened", creator)
	} else if state == types.PullRequest_CLOSED {
		return fmt.Sprintf("@%v closed", creator)
	} else if state == types.PullRequest_MERGED {
		return fmt.Sprintf("@%v merged", creator)
	}
	return "undefined"
}

func AddReviewersCommentBody(creator string, reviewers []string) string {
	return fmt.Sprintf("@%v requested review from"+JoinReviewers(reviewers), creator)
}

func RemoveReviewersCommentBody(creator string, reviewers []string) string {
	return fmt.Sprintf("@%v removed review request for"+JoinReviewers(reviewers), creator)
}
