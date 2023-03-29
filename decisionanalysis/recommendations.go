package decisionanalysis

import (
	"fmt"
)

type decisionNode struct {
	question  string
	yesBranch *decisionNode
	noBranch  *decisionNode
	result    string
}

func askQuestion(question string) bool {
	var answer string
	fmt.Print(question + " (y/n) ")
	fmt.Scanln(&answer)
	return answer == "y"
}

func makeDecisionTree(cost float64, leaseTerm int, interestRate float64, residualValue float64, maintenanceCost *float64, repairCost *float64, marketDemand *string, competitiveLandscape *string, taxImplications *string) *decisionNode {
	rootNode := &decisionNode{
		question: fmt.Sprintf("Is the cost of the asset greater than $%.2f?", cost),
		yesBranch: &decisionNode{
			question: fmt.Sprintf("Will the asset be leased for more than %d years?", leaseTerm),
			yesBranch: &decisionNode{
				question: fmt.Sprintf("Is the interest rate higher than %.2f%%?", interestRate*100),
				yesBranch: &decisionNode{
					question: "Are the maintenance costs low?",
					yesBranch: &decisionNode{
						question: fmt.Sprintf("Will the residual value of the asset be more than %.2f%% of the initial cost?", residualValue*100),
						yesBranch: &decisionNode{
							result: "Leasing is recommended.",
						},
						noBranch: &decisionNode{
							result: "Buying is recommended.",
						},
					},
					noBranch: &decisionNode{
						result: "Buying is recommended.",
					},
				},
				noBranch: &decisionNode{
					result: "Buying is recommended.",
				},
			},
			noBranch: &decisionNode{
				question: fmt.Sprintf("Will the residual value of the asset be more than %.2f%% of the initial cost?", residualValue*100),
				yesBranch: &decisionNode{
					result: "Leasing is recommended.",
				},
				noBranch: &decisionNode{
					result: "Buying is recommended.",
				},
			},
		},
		noBranch: &decisionNode{
			result: "Buying is recommended.",
		},
	}

	if maintenanceCost != nil {
		rootNode.yesBranch.yesBranch.yesBranch.yesBranch = &decisionNode{
			question: fmt.Sprintf("Are the maintenance costs high (%.2f)?", *maintenanceCost),
			yesBranch: &decisionNode{
				result: "Leasing is recommended.",
			},
			noBranch: &decisionNode{
				result: "Buying is recommended.",
			},
		}
	}

	if marketDemand != nil {
		rootNode.yesBranch.yesBranch.noBranch.yesBranch.noBranch = &decisionNode{
			question: fmt.Sprintf("Is the market demand for the asset high (%s)?", *marketDemand),
			yesBranch: &decisionNode{
				result: "Leasing is recommended.",
			},
			noBranch: &decisionNode{
				result: "Buying is recommended.",
			},
		}
	}

	if competitiveLandscape != nil {
		rootNode.yesBranch.yesBranch.noBranch.noBranch = &decisionNode{
			question: fmt.Sprintf("Is the competitive landscape for the asset favorable to leasing (%s)?", *competitiveLandscape),
			yesBranch: &decisionNode{
				result: "Leasing is recommended.",
			},
			noBranch: &decisionNode{
				result: "Buying is recommended.",
			},
		}
	}

	if repairCost != nil {
		rootNode.yesBranch.yesBranch.yesBranch.noBranch = &decisionNode{
			question: fmt.Sprintf("Are the repair costs high (%.2f)?", *repairCost),
			yesBranch: &decisionNode{
				result: "Leasing is recommended.",
			},
			noBranch: &decisionNode{
				result: "Buying is recommended.",
			},
		}
	}

	if taxImplications != nil {
		rootNode.yesBranch.yesBranch.yesBranch.yesBranch.noBranch = &decisionNode{
			question: fmt.Sprintf("Are there tax implications for leasing (%s)?", *taxImplications),
			yesBranch: &decisionNode{
				result: "Leasing is recommended.",
			},
			noBranch: &decisionNode{
				result: "Buying is recommended.",
			},
		}
	}

	return rootNode
}

func getRecommendation(node *decisionNode) string {
	if node.result != "" {
		return node.result
	}

	if askQuestion(node.question) {
		return getRecommendation(node.yesBranch)
	} else {
		return getRecommendation(node.noBranch)
	}
}

func GetRecommendation(cost float64, leaseTerm int, interestRate float64, residualValue float64, maintenanceCost *float64, repairCost *float64, marketDemand *string, competitiveLandscape *string, taxImplications *string) string {
	rootNode := makeDecisionTree(cost, leaseTerm, interestRate, residualValue, maintenanceCost, repairCost, marketDemand, competitiveLandscape, taxImplications)
	return getRecommendation(rootNode)
}
