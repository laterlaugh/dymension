package types

import (
	fmt "fmt"
	"time"

	"cosmossdk.io/math"
)

// Default parameter values

var (
	DefaultTakerFee                                     = "0.02"                      // 2%
	DefaultCreationFee                                  = math.NewInt(1).MulRaw(1e18) /* 1 Rollapp token */
	DefaultMinPlanDuration                              = 0 * time.Hour               // no enforced minimum by default
	DefaultIncentivePlanMinimumNumEpochsPaidOver        = uint64(364)                 // default: min 364 days (based on 1 day distribution epoch)
	DefaultIncentivePlanMinimumStartTimeAfterSettlement = 60 * time.Minute            // default: min 1 hour after settlement
	DefaultMinLiquidityPart                             = "0.4"                       // default: at least 40% goes to the liquidity pool
)

// NewParams creates a new Params object
func NewParams(takerFee, liquidityPart math.LegacyDec, creationFee math.Int, minPlanDuration time.Duration, minIncentivePlanParams IncentivePlanParams) Params {
	return Params{
		TakerFee:                              takerFee,
		CreationFee:                           creationFee,
		MinPlanDuration:                       minPlanDuration,
		IncentivesMinStartTimeAfterSettlement: minIncentivePlanParams.StartTimeAfterSettlement,
		IncentivesMinNumEpochsPaidOver:        minIncentivePlanParams.NumEpochsPaidOver,
		MinLiquidityPart:                      liquidityPart,
	}
}

// DefaultParams returns a default set of parameters.
func DefaultParams() Params {
	return Params{
		TakerFee:                              math.LegacyMustNewDecFromStr(DefaultTakerFee),
		CreationFee:                           DefaultCreationFee,
		MinPlanDuration:                       DefaultMinPlanDuration,
		IncentivesMinStartTimeAfterSettlement: DefaultIncentivePlanMinimumStartTimeAfterSettlement,
		IncentivesMinNumEpochsPaidOver:        DefaultIncentivePlanMinimumNumEpochsPaidOver,
		MinLiquidityPart:                      math.LegacyMustNewDecFromStr(DefaultMinLiquidityPart),
	}
}

// Validate checks that the parameters have valid values.
func (p Params) Validate() error {
	if err := validateTakerFee(p.TakerFee); err != nil {
		return err
	}

	if err := validateCreationFee(p.CreationFee); err != nil {
		return err
	}

	if p.MinPlanDuration < 0 {
		return fmt.Errorf("minimum plan duration must be non-negative: %v", p.MinPlanDuration)
	}

	if p.IncentivesMinNumEpochsPaidOver < 1 {
		return fmt.Errorf("incentive plan num epochs paid over must be greater than 0: %d", p.IncentivesMinNumEpochsPaidOver)
	}

	if p.IncentivesMinStartTimeAfterSettlement <= 0 {
		return fmt.Errorf("incentive plan start time after settlement must be greater than 0: %v", p.IncentivesMinStartTimeAfterSettlement)
	}

	if !p.MinLiquidityPart.IsPositive() {
		return fmt.Errorf("min liquidity part must be positive: %s", p.MinLiquidityPart)
	}

	return nil
}

func validateTakerFee(i interface{}) error {
	v, ok := i.(math.LegacyDec)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.IsNil() || v.IsNegative() {
		return fmt.Errorf("taker fee must be a non-negative decimal: %s", v)
	}

	if v.GTE(math.LegacyOneDec()) {
		return fmt.Errorf("taker fee must be less than 1: %s", v)
	}

	return nil
}

func validateCreationFee(i interface{}) error {
	v, ok := i.(math.Int)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	// creation fee must be a positive integer greater than 1^18 (1 Rollapp token)
	if v.LT(math.NewIntWithDecimal(1, 18)) {
		return fmt.Errorf("creation fee must be a positive integer: %s", v)
	}

	return nil
}
