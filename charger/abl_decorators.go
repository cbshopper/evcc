package charger

// Code generated by github.com/evcc-io/evcc/cmd/tools/decorate.go. DO NOT EDIT.

import (
	"github.com/evcc-io/evcc/api"
)

func decorateABLeMH(base *ABLeMH, meter func() (float64, error), phaseCurrents func() (float64, float64, float64, error)) api.Charger {
	switch {
	case meter == nil && phaseCurrents == nil:
		return base

	case meter != nil && phaseCurrents == nil:
		return &struct {
			*ABLeMH
			api.Meter
		}{
			ABLeMH: base,
			Meter: &decorateABLeMHMeterImpl{
				meter: meter,
			},
		}

	case meter == nil && phaseCurrents != nil:
		return &struct {
			*ABLeMH
			api.PhaseCurrents
		}{
			ABLeMH: base,
			PhaseCurrents: &decorateABLeMHPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
		}

	case meter != nil && phaseCurrents != nil:
		return &struct {
			*ABLeMH
			api.Meter
			api.PhaseCurrents
		}{
			ABLeMH: base,
			Meter: &decorateABLeMHMeterImpl{
				meter: meter,
			},
			PhaseCurrents: &decorateABLeMHPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
		}
	}

	return nil
}

type decorateABLeMHMeterImpl struct {
	meter func() (float64, error)
}

func (impl *decorateABLeMHMeterImpl) CurrentPower() (float64, error) {
	return impl.meter()
}

type decorateABLeMHPhaseCurrentsImpl struct {
	phaseCurrents func() (float64, float64, float64, error)
}

func (impl *decorateABLeMHPhaseCurrentsImpl) Currents() (float64, float64, float64, error) {
	return impl.phaseCurrents()
}
