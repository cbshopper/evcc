package charger

// Code generated by github.com/evcc-io/evcc/cmd/tools/decorate.go. DO NOT EDIT.

import (
	"github.com/evcc-io/evcc/api"
)

func decorateWallbe(base *Wallbe, meter func() (float64, error), meterEnergy func() (float64, error), phaseCurrents func() (float64, float64, float64, error), chargerEx func(float64) error) api.Charger {
	switch {
	case chargerEx == nil && meter == nil && meterEnergy == nil && phaseCurrents == nil:
		return base

	case chargerEx == nil && meter != nil && meterEnergy == nil && phaseCurrents == nil:
		return &struct {
			*Wallbe
			api.Meter
		}{
			Wallbe: base,
			Meter: &decorateWallbeMeterImpl{
				meter: meter,
			},
		}

	case chargerEx == nil && meter == nil && meterEnergy != nil && phaseCurrents == nil:
		return &struct {
			*Wallbe
			api.MeterEnergy
		}{
			Wallbe: base,
			MeterEnergy: &decorateWallbeMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
		}

	case chargerEx == nil && meter != nil && meterEnergy != nil && phaseCurrents == nil:
		return &struct {
			*Wallbe
			api.Meter
			api.MeterEnergy
		}{
			Wallbe: base,
			Meter: &decorateWallbeMeterImpl{
				meter: meter,
			},
			MeterEnergy: &decorateWallbeMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
		}

	case chargerEx == nil && meter == nil && meterEnergy == nil && phaseCurrents != nil:
		return &struct {
			*Wallbe
			api.PhaseCurrents
		}{
			Wallbe: base,
			PhaseCurrents: &decorateWallbePhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
		}

	case chargerEx == nil && meter != nil && meterEnergy == nil && phaseCurrents != nil:
		return &struct {
			*Wallbe
			api.Meter
			api.PhaseCurrents
		}{
			Wallbe: base,
			Meter: &decorateWallbeMeterImpl{
				meter: meter,
			},
			PhaseCurrents: &decorateWallbePhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
		}

	case chargerEx == nil && meter == nil && meterEnergy != nil && phaseCurrents != nil:
		return &struct {
			*Wallbe
			api.MeterEnergy
			api.PhaseCurrents
		}{
			Wallbe: base,
			MeterEnergy: &decorateWallbeMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhaseCurrents: &decorateWallbePhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
		}

	case chargerEx == nil && meter != nil && meterEnergy != nil && phaseCurrents != nil:
		return &struct {
			*Wallbe
			api.Meter
			api.MeterEnergy
			api.PhaseCurrents
		}{
			Wallbe: base,
			Meter: &decorateWallbeMeterImpl{
				meter: meter,
			},
			MeterEnergy: &decorateWallbeMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhaseCurrents: &decorateWallbePhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
		}

	case chargerEx != nil && meter == nil && meterEnergy == nil && phaseCurrents == nil:
		return &struct {
			*Wallbe
			api.ChargerEx
		}{
			Wallbe: base,
			ChargerEx: &decorateWallbeChargerExImpl{
				chargerEx: chargerEx,
			},
		}

	case chargerEx != nil && meter != nil && meterEnergy == nil && phaseCurrents == nil:
		return &struct {
			*Wallbe
			api.ChargerEx
			api.Meter
		}{
			Wallbe: base,
			ChargerEx: &decorateWallbeChargerExImpl{
				chargerEx: chargerEx,
			},
			Meter: &decorateWallbeMeterImpl{
				meter: meter,
			},
		}

	case chargerEx != nil && meter == nil && meterEnergy != nil && phaseCurrents == nil:
		return &struct {
			*Wallbe
			api.ChargerEx
			api.MeterEnergy
		}{
			Wallbe: base,
			ChargerEx: &decorateWallbeChargerExImpl{
				chargerEx: chargerEx,
			},
			MeterEnergy: &decorateWallbeMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
		}

	case chargerEx != nil && meter != nil && meterEnergy != nil && phaseCurrents == nil:
		return &struct {
			*Wallbe
			api.ChargerEx
			api.Meter
			api.MeterEnergy
		}{
			Wallbe: base,
			ChargerEx: &decorateWallbeChargerExImpl{
				chargerEx: chargerEx,
			},
			Meter: &decorateWallbeMeterImpl{
				meter: meter,
			},
			MeterEnergy: &decorateWallbeMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
		}

	case chargerEx != nil && meter == nil && meterEnergy == nil && phaseCurrents != nil:
		return &struct {
			*Wallbe
			api.ChargerEx
			api.PhaseCurrents
		}{
			Wallbe: base,
			ChargerEx: &decorateWallbeChargerExImpl{
				chargerEx: chargerEx,
			},
			PhaseCurrents: &decorateWallbePhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
		}

	case chargerEx != nil && meter != nil && meterEnergy == nil && phaseCurrents != nil:
		return &struct {
			*Wallbe
			api.ChargerEx
			api.Meter
			api.PhaseCurrents
		}{
			Wallbe: base,
			ChargerEx: &decorateWallbeChargerExImpl{
				chargerEx: chargerEx,
			},
			Meter: &decorateWallbeMeterImpl{
				meter: meter,
			},
			PhaseCurrents: &decorateWallbePhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
		}

	case chargerEx != nil && meter == nil && meterEnergy != nil && phaseCurrents != nil:
		return &struct {
			*Wallbe
			api.ChargerEx
			api.MeterEnergy
			api.PhaseCurrents
		}{
			Wallbe: base,
			ChargerEx: &decorateWallbeChargerExImpl{
				chargerEx: chargerEx,
			},
			MeterEnergy: &decorateWallbeMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhaseCurrents: &decorateWallbePhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
		}

	case chargerEx != nil && meter != nil && meterEnergy != nil && phaseCurrents != nil:
		return &struct {
			*Wallbe
			api.ChargerEx
			api.Meter
			api.MeterEnergy
			api.PhaseCurrents
		}{
			Wallbe: base,
			ChargerEx: &decorateWallbeChargerExImpl{
				chargerEx: chargerEx,
			},
			Meter: &decorateWallbeMeterImpl{
				meter: meter,
			},
			MeterEnergy: &decorateWallbeMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhaseCurrents: &decorateWallbePhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
		}
	}

	return nil
}

type decorateWallbeChargerExImpl struct {
	chargerEx func(float64) error
}

func (impl *decorateWallbeChargerExImpl) MaxCurrentMillis(p0 float64) error {
	return impl.chargerEx(p0)
}

type decorateWallbeMeterImpl struct {
	meter func() (float64, error)
}

func (impl *decorateWallbeMeterImpl) CurrentPower() (float64, error) {
	return impl.meter()
}

type decorateWallbeMeterEnergyImpl struct {
	meterEnergy func() (float64, error)
}

func (impl *decorateWallbeMeterEnergyImpl) TotalEnergy() (float64, error) {
	return impl.meterEnergy()
}

type decorateWallbePhaseCurrentsImpl struct {
	phaseCurrents func() (float64, float64, float64, error)
}

func (impl *decorateWallbePhaseCurrentsImpl) Currents() (float64, float64, float64, error) {
	return impl.phaseCurrents()
}
