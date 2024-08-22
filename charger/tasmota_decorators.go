package charger

// Code generated by github.com/evcc-io/evcc/cmd/tools/decorate.go. DO NOT EDIT.

import (
	"github.com/evcc-io/evcc/api"
)

func decorateTasmota(base *Tasmota, phaseVoltages func() (float64, float64, float64, error), phaseCurrents func() (float64, float64, float64, error)) api.Charger {
	switch {
	case phaseCurrents == nil && phaseVoltages == nil:
		return base

	case phaseCurrents == nil && phaseVoltages != nil:
		return &struct {
			*Tasmota
			api.PhaseVoltages
		}{
			Tasmota: base,
			PhaseVoltages: &decorateTasmotaPhaseVoltagesImpl{
				phaseVoltages: phaseVoltages,
			},
		}

	case phaseCurrents != nil && phaseVoltages == nil:
		return &struct {
			*Tasmota
			api.PhaseCurrents
		}{
			Tasmota: base,
			PhaseCurrents: &decorateTasmotaPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
		}

	case phaseCurrents != nil && phaseVoltages != nil:
		return &struct {
			*Tasmota
			api.PhaseCurrents
			api.PhaseVoltages
		}{
			Tasmota: base,
			PhaseCurrents: &decorateTasmotaPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
			PhaseVoltages: &decorateTasmotaPhaseVoltagesImpl{
				phaseVoltages: phaseVoltages,
			},
		}
	}

	return nil
}

type decorateTasmotaPhaseCurrentsImpl struct {
	phaseCurrents func() (float64, float64, float64, error)
}

func (impl *decorateTasmotaPhaseCurrentsImpl) Currents() (float64, float64, float64, error) {
	return impl.phaseCurrents()
}

type decorateTasmotaPhaseVoltagesImpl struct {
	phaseVoltages func() (float64, float64, float64, error)
}

func (impl *decorateTasmotaPhaseVoltagesImpl) Voltages() (float64, float64, float64, error) {
	return impl.phaseVoltages()
}
