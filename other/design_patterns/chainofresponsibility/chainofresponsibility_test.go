package chainofresponsibility

import (
	"github.com/gookit/goutil/dump"
	"testing"
)

func TestCompleteBoardingProcessor_ProcessFor(t *testing.T) {
	passenger := &Passenger{
		name:                  "邹祝家",
		hasBoardingPass:       false,
		hasLuggage:            false,
		isPassIdentityCheck:   false,
		isPassSecurityCheck:   false,
		isCompleteForBoarding: false,
	}

	// 办理登机牌
	bp := &boardingPassProcessor{}
	bp.ProcessFor(passenger)

	// 托运行李
	lp := &luggageCheckInProcessor{}
	lp.ProcessFor(passenger)

	// 身份校验
	ip := &identityCheckProcessor{}
	ip.ProcessFor(passenger)

	// 安检
	sp := &securityCheckProcessor{}
	sp.ProcessFor(passenger)

	// 完成登机
	cp := &completeBoardingProcessor{}
	cp.ProcessFor(passenger)

}

func TestChainOfResponsibility(t *testing.T) {
	bp := BuildBoardingProcessorChain()

	passenger := &Passenger{
		name:                  "邹祝家",
		hasBoardingPass:       false,
		hasLuggage:            false,
		isPassIdentityCheck:   false,
		isPassSecurityCheck:   false,
		isCompleteForBoarding: false,
	}
	dump.P(bp)
	bp.ProcessFor(passenger)
}

func BuildBoardingProcessorChain() BoardingProcessor {
	cp := &completeBoardingProcessor{}

	securityCheckNode := &securityCheckProcessor{}
	securityCheckNode.SetNextProcessor(cp)

	identityCheckNode := &identityCheckProcessor{}
	identityCheckNode.SetNextProcessor(securityCheckNode)

	luggageCheckInNode := &luggageCheckInProcessor{}
	luggageCheckInNode.SetNextProcessor(identityCheckNode)

	boardingPassNode := &boardingPassProcessor{}
	boardingPassNode.SetNextProcessor(luggageCheckInNode)

	return boardingPassNode
}
