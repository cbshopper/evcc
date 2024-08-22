package meter

// Code generated by github.com/evcc-io/evcc/cmd/tools/decorate.go. DO NOT EDIT.

import (
	"github.com/evcc-io/evcc/api"
)

func decorateModbusMbmd(base api.Meter, meterEnergy func() (float64, error), phaseCurrents func() (float64, float64, float64, error), phaseVoltages func() (float64, float64, float64, error), phasePowers func() (float64, float64, float64, error), battery func() (float64, error), batteryCapacity func() float64) api.Meter {
	switch {
	case battery == nil && batteryCapacity == nil && meterEnergy == nil && phaseCurrents == nil && phasePowers == nil && phaseVoltages == nil:
		return base

	case battery == nil && batteryCapacity == nil && meterEnergy != nil && phaseCurrents == nil && phasePowers == nil && phaseVoltages == nil:
		return &struct {
			api.Meter
			api.MeterEnergy
		}{
			Meter: base,
			MeterEnergy: &decorateModbusMbmdMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
		}

	case battery == nil && batteryCapacity == nil && meterEnergy == nil && phaseCurrents != nil && phasePowers == nil && phaseVoltages == nil:
		return &struct {
			api.Meter
			api.PhaseCurrents
		}{
			Meter: base,
			PhaseCurrents: &decorateModbusMbmdPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
		}

	case battery == nil && batteryCapacity == nil && meterEnergy != nil && phaseCurrents != nil && phasePowers == nil && phaseVoltages == nil:
		return &struct {
			api.Meter
			api.MeterEnergy
			api.PhaseCurrents
		}{
			Meter: base,
			MeterEnergy: &decorateModbusMbmdMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhaseCurrents: &decorateModbusMbmdPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
		}

	case battery == nil && batteryCapacity == nil && meterEnergy == nil && phaseCurrents == nil && phasePowers == nil && phaseVoltages != nil:
		return &struct {
			api.Meter
			api.PhaseVoltages
		}{
			Meter: base,
			PhaseVoltages: &decorateModbusMbmdPhaseVoltagesImpl{
				phaseVoltages: phaseVoltages,
			},
		}

	case battery == nil && batteryCapacity == nil && meterEnergy != nil && phaseCurrents == nil && phasePowers == nil && phaseVoltages != nil:
		return &struct {
			api.Meter
			api.MeterEnergy
			api.PhaseVoltages
		}{
			Meter: base,
			MeterEnergy: &decorateModbusMbmdMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhaseVoltages: &decorateModbusMbmdPhaseVoltagesImpl{
				phaseVoltages: phaseVoltages,
			},
		}

	case battery == nil && batteryCapacity == nil && meterEnergy == nil && phaseCurrents != nil && phasePowers == nil && phaseVoltages != nil:
		return &struct {
			api.Meter
			api.PhaseCurrents
			api.PhaseVoltages
		}{
			Meter: base,
			PhaseCurrents: &decorateModbusMbmdPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
			PhaseVoltages: &decorateModbusMbmdPhaseVoltagesImpl{
				phaseVoltages: phaseVoltages,
			},
		}

	case battery == nil && batteryCapacity == nil && meterEnergy != nil && phaseCurrents != nil && phasePowers == nil && phaseVoltages != nil:
		return &struct {
			api.Meter
			api.MeterEnergy
			api.PhaseCurrents
			api.PhaseVoltages
		}{
			Meter: base,
			MeterEnergy: &decorateModbusMbmdMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhaseCurrents: &decorateModbusMbmdPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
			PhaseVoltages: &decorateModbusMbmdPhaseVoltagesImpl{
				phaseVoltages: phaseVoltages,
			},
		}

	case battery == nil && batteryCapacity == nil && meterEnergy == nil && phaseCurrents == nil && phasePowers != nil && phaseVoltages == nil:
		return &struct {
			api.Meter
			api.PhasePowers
		}{
			Meter: base,
			PhasePowers: &decorateModbusMbmdPhasePowersImpl{
				phasePowers: phasePowers,
			},
		}

	case battery == nil && batteryCapacity == nil && meterEnergy != nil && phaseCurrents == nil && phasePowers != nil && phaseVoltages == nil:
		return &struct {
			api.Meter
			api.MeterEnergy
			api.PhasePowers
		}{
			Meter: base,
			MeterEnergy: &decorateModbusMbmdMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhasePowers: &decorateModbusMbmdPhasePowersImpl{
				phasePowers: phasePowers,
			},
		}

	case battery == nil && batteryCapacity == nil && meterEnergy == nil && phaseCurrents != nil && phasePowers != nil && phaseVoltages == nil:
		return &struct {
			api.Meter
			api.PhaseCurrents
			api.PhasePowers
		}{
			Meter: base,
			PhaseCurrents: &decorateModbusMbmdPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
			PhasePowers: &decorateModbusMbmdPhasePowersImpl{
				phasePowers: phasePowers,
			},
		}

	case battery == nil && batteryCapacity == nil && meterEnergy != nil && phaseCurrents != nil && phasePowers != nil && phaseVoltages == nil:
		return &struct {
			api.Meter
			api.MeterEnergy
			api.PhaseCurrents
			api.PhasePowers
		}{
			Meter: base,
			MeterEnergy: &decorateModbusMbmdMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhaseCurrents: &decorateModbusMbmdPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
			PhasePowers: &decorateModbusMbmdPhasePowersImpl{
				phasePowers: phasePowers,
			},
		}

	case battery == nil && batteryCapacity == nil && meterEnergy == nil && phaseCurrents == nil && phasePowers != nil && phaseVoltages != nil:
		return &struct {
			api.Meter
			api.PhasePowers
			api.PhaseVoltages
		}{
			Meter: base,
			PhasePowers: &decorateModbusMbmdPhasePowersImpl{
				phasePowers: phasePowers,
			},
			PhaseVoltages: &decorateModbusMbmdPhaseVoltagesImpl{
				phaseVoltages: phaseVoltages,
			},
		}

	case battery == nil && batteryCapacity == nil && meterEnergy != nil && phaseCurrents == nil && phasePowers != nil && phaseVoltages != nil:
		return &struct {
			api.Meter
			api.MeterEnergy
			api.PhasePowers
			api.PhaseVoltages
		}{
			Meter: base,
			MeterEnergy: &decorateModbusMbmdMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhasePowers: &decorateModbusMbmdPhasePowersImpl{
				phasePowers: phasePowers,
			},
			PhaseVoltages: &decorateModbusMbmdPhaseVoltagesImpl{
				phaseVoltages: phaseVoltages,
			},
		}

	case battery == nil && batteryCapacity == nil && meterEnergy == nil && phaseCurrents != nil && phasePowers != nil && phaseVoltages != nil:
		return &struct {
			api.Meter
			api.PhaseCurrents
			api.PhasePowers
			api.PhaseVoltages
		}{
			Meter: base,
			PhaseCurrents: &decorateModbusMbmdPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
			PhasePowers: &decorateModbusMbmdPhasePowersImpl{
				phasePowers: phasePowers,
			},
			PhaseVoltages: &decorateModbusMbmdPhaseVoltagesImpl{
				phaseVoltages: phaseVoltages,
			},
		}

	case battery == nil && batteryCapacity == nil && meterEnergy != nil && phaseCurrents != nil && phasePowers != nil && phaseVoltages != nil:
		return &struct {
			api.Meter
			api.MeterEnergy
			api.PhaseCurrents
			api.PhasePowers
			api.PhaseVoltages
		}{
			Meter: base,
			MeterEnergy: &decorateModbusMbmdMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhaseCurrents: &decorateModbusMbmdPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
			PhasePowers: &decorateModbusMbmdPhasePowersImpl{
				phasePowers: phasePowers,
			},
			PhaseVoltages: &decorateModbusMbmdPhaseVoltagesImpl{
				phaseVoltages: phaseVoltages,
			},
		}

	case battery != nil && batteryCapacity == nil && meterEnergy == nil && phaseCurrents == nil && phasePowers == nil && phaseVoltages == nil:
		return &struct {
			api.Meter
			api.Battery
		}{
			Meter: base,
			Battery: &decorateModbusMbmdBatteryImpl{
				battery: battery,
			},
		}

	case battery != nil && batteryCapacity == nil && meterEnergy != nil && phaseCurrents == nil && phasePowers == nil && phaseVoltages == nil:
		return &struct {
			api.Meter
			api.Battery
			api.MeterEnergy
		}{
			Meter: base,
			Battery: &decorateModbusMbmdBatteryImpl{
				battery: battery,
			},
			MeterEnergy: &decorateModbusMbmdMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
		}

	case battery != nil && batteryCapacity == nil && meterEnergy == nil && phaseCurrents != nil && phasePowers == nil && phaseVoltages == nil:
		return &struct {
			api.Meter
			api.Battery
			api.PhaseCurrents
		}{
			Meter: base,
			Battery: &decorateModbusMbmdBatteryImpl{
				battery: battery,
			},
			PhaseCurrents: &decorateModbusMbmdPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
		}

	case battery != nil && batteryCapacity == nil && meterEnergy != nil && phaseCurrents != nil && phasePowers == nil && phaseVoltages == nil:
		return &struct {
			api.Meter
			api.Battery
			api.MeterEnergy
			api.PhaseCurrents
		}{
			Meter: base,
			Battery: &decorateModbusMbmdBatteryImpl{
				battery: battery,
			},
			MeterEnergy: &decorateModbusMbmdMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhaseCurrents: &decorateModbusMbmdPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
		}

	case battery != nil && batteryCapacity == nil && meterEnergy == nil && phaseCurrents == nil && phasePowers == nil && phaseVoltages != nil:
		return &struct {
			api.Meter
			api.Battery
			api.PhaseVoltages
		}{
			Meter: base,
			Battery: &decorateModbusMbmdBatteryImpl{
				battery: battery,
			},
			PhaseVoltages: &decorateModbusMbmdPhaseVoltagesImpl{
				phaseVoltages: phaseVoltages,
			},
		}

	case battery != nil && batteryCapacity == nil && meterEnergy != nil && phaseCurrents == nil && phasePowers == nil && phaseVoltages != nil:
		return &struct {
			api.Meter
			api.Battery
			api.MeterEnergy
			api.PhaseVoltages
		}{
			Meter: base,
			Battery: &decorateModbusMbmdBatteryImpl{
				battery: battery,
			},
			MeterEnergy: &decorateModbusMbmdMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhaseVoltages: &decorateModbusMbmdPhaseVoltagesImpl{
				phaseVoltages: phaseVoltages,
			},
		}

	case battery != nil && batteryCapacity == nil && meterEnergy == nil && phaseCurrents != nil && phasePowers == nil && phaseVoltages != nil:
		return &struct {
			api.Meter
			api.Battery
			api.PhaseCurrents
			api.PhaseVoltages
		}{
			Meter: base,
			Battery: &decorateModbusMbmdBatteryImpl{
				battery: battery,
			},
			PhaseCurrents: &decorateModbusMbmdPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
			PhaseVoltages: &decorateModbusMbmdPhaseVoltagesImpl{
				phaseVoltages: phaseVoltages,
			},
		}

	case battery != nil && batteryCapacity == nil && meterEnergy != nil && phaseCurrents != nil && phasePowers == nil && phaseVoltages != nil:
		return &struct {
			api.Meter
			api.Battery
			api.MeterEnergy
			api.PhaseCurrents
			api.PhaseVoltages
		}{
			Meter: base,
			Battery: &decorateModbusMbmdBatteryImpl{
				battery: battery,
			},
			MeterEnergy: &decorateModbusMbmdMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhaseCurrents: &decorateModbusMbmdPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
			PhaseVoltages: &decorateModbusMbmdPhaseVoltagesImpl{
				phaseVoltages: phaseVoltages,
			},
		}

	case battery != nil && batteryCapacity == nil && meterEnergy == nil && phaseCurrents == nil && phasePowers != nil && phaseVoltages == nil:
		return &struct {
			api.Meter
			api.Battery
			api.PhasePowers
		}{
			Meter: base,
			Battery: &decorateModbusMbmdBatteryImpl{
				battery: battery,
			},
			PhasePowers: &decorateModbusMbmdPhasePowersImpl{
				phasePowers: phasePowers,
			},
		}

	case battery != nil && batteryCapacity == nil && meterEnergy != nil && phaseCurrents == nil && phasePowers != nil && phaseVoltages == nil:
		return &struct {
			api.Meter
			api.Battery
			api.MeterEnergy
			api.PhasePowers
		}{
			Meter: base,
			Battery: &decorateModbusMbmdBatteryImpl{
				battery: battery,
			},
			MeterEnergy: &decorateModbusMbmdMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhasePowers: &decorateModbusMbmdPhasePowersImpl{
				phasePowers: phasePowers,
			},
		}

	case battery != nil && batteryCapacity == nil && meterEnergy == nil && phaseCurrents != nil && phasePowers != nil && phaseVoltages == nil:
		return &struct {
			api.Meter
			api.Battery
			api.PhaseCurrents
			api.PhasePowers
		}{
			Meter: base,
			Battery: &decorateModbusMbmdBatteryImpl{
				battery: battery,
			},
			PhaseCurrents: &decorateModbusMbmdPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
			PhasePowers: &decorateModbusMbmdPhasePowersImpl{
				phasePowers: phasePowers,
			},
		}

	case battery != nil && batteryCapacity == nil && meterEnergy != nil && phaseCurrents != nil && phasePowers != nil && phaseVoltages == nil:
		return &struct {
			api.Meter
			api.Battery
			api.MeterEnergy
			api.PhaseCurrents
			api.PhasePowers
		}{
			Meter: base,
			Battery: &decorateModbusMbmdBatteryImpl{
				battery: battery,
			},
			MeterEnergy: &decorateModbusMbmdMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhaseCurrents: &decorateModbusMbmdPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
			PhasePowers: &decorateModbusMbmdPhasePowersImpl{
				phasePowers: phasePowers,
			},
		}

	case battery != nil && batteryCapacity == nil && meterEnergy == nil && phaseCurrents == nil && phasePowers != nil && phaseVoltages != nil:
		return &struct {
			api.Meter
			api.Battery
			api.PhasePowers
			api.PhaseVoltages
		}{
			Meter: base,
			Battery: &decorateModbusMbmdBatteryImpl{
				battery: battery,
			},
			PhasePowers: &decorateModbusMbmdPhasePowersImpl{
				phasePowers: phasePowers,
			},
			PhaseVoltages: &decorateModbusMbmdPhaseVoltagesImpl{
				phaseVoltages: phaseVoltages,
			},
		}

	case battery != nil && batteryCapacity == nil && meterEnergy != nil && phaseCurrents == nil && phasePowers != nil && phaseVoltages != nil:
		return &struct {
			api.Meter
			api.Battery
			api.MeterEnergy
			api.PhasePowers
			api.PhaseVoltages
		}{
			Meter: base,
			Battery: &decorateModbusMbmdBatteryImpl{
				battery: battery,
			},
			MeterEnergy: &decorateModbusMbmdMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhasePowers: &decorateModbusMbmdPhasePowersImpl{
				phasePowers: phasePowers,
			},
			PhaseVoltages: &decorateModbusMbmdPhaseVoltagesImpl{
				phaseVoltages: phaseVoltages,
			},
		}

	case battery != nil && batteryCapacity == nil && meterEnergy == nil && phaseCurrents != nil && phasePowers != nil && phaseVoltages != nil:
		return &struct {
			api.Meter
			api.Battery
			api.PhaseCurrents
			api.PhasePowers
			api.PhaseVoltages
		}{
			Meter: base,
			Battery: &decorateModbusMbmdBatteryImpl{
				battery: battery,
			},
			PhaseCurrents: &decorateModbusMbmdPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
			PhasePowers: &decorateModbusMbmdPhasePowersImpl{
				phasePowers: phasePowers,
			},
			PhaseVoltages: &decorateModbusMbmdPhaseVoltagesImpl{
				phaseVoltages: phaseVoltages,
			},
		}

	case battery != nil && batteryCapacity == nil && meterEnergy != nil && phaseCurrents != nil && phasePowers != nil && phaseVoltages != nil:
		return &struct {
			api.Meter
			api.Battery
			api.MeterEnergy
			api.PhaseCurrents
			api.PhasePowers
			api.PhaseVoltages
		}{
			Meter: base,
			Battery: &decorateModbusMbmdBatteryImpl{
				battery: battery,
			},
			MeterEnergy: &decorateModbusMbmdMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhaseCurrents: &decorateModbusMbmdPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
			PhasePowers: &decorateModbusMbmdPhasePowersImpl{
				phasePowers: phasePowers,
			},
			PhaseVoltages: &decorateModbusMbmdPhaseVoltagesImpl{
				phaseVoltages: phaseVoltages,
			},
		}

	case battery == nil && batteryCapacity != nil && meterEnergy == nil && phaseCurrents == nil && phasePowers == nil && phaseVoltages == nil:
		return &struct {
			api.Meter
			api.BatteryCapacity
		}{
			Meter: base,
			BatteryCapacity: &decorateModbusMbmdBatteryCapacityImpl{
				batteryCapacity: batteryCapacity,
			},
		}

	case battery == nil && batteryCapacity != nil && meterEnergy != nil && phaseCurrents == nil && phasePowers == nil && phaseVoltages == nil:
		return &struct {
			api.Meter
			api.BatteryCapacity
			api.MeterEnergy
		}{
			Meter: base,
			BatteryCapacity: &decorateModbusMbmdBatteryCapacityImpl{
				batteryCapacity: batteryCapacity,
			},
			MeterEnergy: &decorateModbusMbmdMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
		}

	case battery == nil && batteryCapacity != nil && meterEnergy == nil && phaseCurrents != nil && phasePowers == nil && phaseVoltages == nil:
		return &struct {
			api.Meter
			api.BatteryCapacity
			api.PhaseCurrents
		}{
			Meter: base,
			BatteryCapacity: &decorateModbusMbmdBatteryCapacityImpl{
				batteryCapacity: batteryCapacity,
			},
			PhaseCurrents: &decorateModbusMbmdPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
		}

	case battery == nil && batteryCapacity != nil && meterEnergy != nil && phaseCurrents != nil && phasePowers == nil && phaseVoltages == nil:
		return &struct {
			api.Meter
			api.BatteryCapacity
			api.MeterEnergy
			api.PhaseCurrents
		}{
			Meter: base,
			BatteryCapacity: &decorateModbusMbmdBatteryCapacityImpl{
				batteryCapacity: batteryCapacity,
			},
			MeterEnergy: &decorateModbusMbmdMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhaseCurrents: &decorateModbusMbmdPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
		}

	case battery == nil && batteryCapacity != nil && meterEnergy == nil && phaseCurrents == nil && phasePowers == nil && phaseVoltages != nil:
		return &struct {
			api.Meter
			api.BatteryCapacity
			api.PhaseVoltages
		}{
			Meter: base,
			BatteryCapacity: &decorateModbusMbmdBatteryCapacityImpl{
				batteryCapacity: batteryCapacity,
			},
			PhaseVoltages: &decorateModbusMbmdPhaseVoltagesImpl{
				phaseVoltages: phaseVoltages,
			},
		}

	case battery == nil && batteryCapacity != nil && meterEnergy != nil && phaseCurrents == nil && phasePowers == nil && phaseVoltages != nil:
		return &struct {
			api.Meter
			api.BatteryCapacity
			api.MeterEnergy
			api.PhaseVoltages
		}{
			Meter: base,
			BatteryCapacity: &decorateModbusMbmdBatteryCapacityImpl{
				batteryCapacity: batteryCapacity,
			},
			MeterEnergy: &decorateModbusMbmdMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhaseVoltages: &decorateModbusMbmdPhaseVoltagesImpl{
				phaseVoltages: phaseVoltages,
			},
		}

	case battery == nil && batteryCapacity != nil && meterEnergy == nil && phaseCurrents != nil && phasePowers == nil && phaseVoltages != nil:
		return &struct {
			api.Meter
			api.BatteryCapacity
			api.PhaseCurrents
			api.PhaseVoltages
		}{
			Meter: base,
			BatteryCapacity: &decorateModbusMbmdBatteryCapacityImpl{
				batteryCapacity: batteryCapacity,
			},
			PhaseCurrents: &decorateModbusMbmdPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
			PhaseVoltages: &decorateModbusMbmdPhaseVoltagesImpl{
				phaseVoltages: phaseVoltages,
			},
		}

	case battery == nil && batteryCapacity != nil && meterEnergy != nil && phaseCurrents != nil && phasePowers == nil && phaseVoltages != nil:
		return &struct {
			api.Meter
			api.BatteryCapacity
			api.MeterEnergy
			api.PhaseCurrents
			api.PhaseVoltages
		}{
			Meter: base,
			BatteryCapacity: &decorateModbusMbmdBatteryCapacityImpl{
				batteryCapacity: batteryCapacity,
			},
			MeterEnergy: &decorateModbusMbmdMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhaseCurrents: &decorateModbusMbmdPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
			PhaseVoltages: &decorateModbusMbmdPhaseVoltagesImpl{
				phaseVoltages: phaseVoltages,
			},
		}

	case battery == nil && batteryCapacity != nil && meterEnergy == nil && phaseCurrents == nil && phasePowers != nil && phaseVoltages == nil:
		return &struct {
			api.Meter
			api.BatteryCapacity
			api.PhasePowers
		}{
			Meter: base,
			BatteryCapacity: &decorateModbusMbmdBatteryCapacityImpl{
				batteryCapacity: batteryCapacity,
			},
			PhasePowers: &decorateModbusMbmdPhasePowersImpl{
				phasePowers: phasePowers,
			},
		}

	case battery == nil && batteryCapacity != nil && meterEnergy != nil && phaseCurrents == nil && phasePowers != nil && phaseVoltages == nil:
		return &struct {
			api.Meter
			api.BatteryCapacity
			api.MeterEnergy
			api.PhasePowers
		}{
			Meter: base,
			BatteryCapacity: &decorateModbusMbmdBatteryCapacityImpl{
				batteryCapacity: batteryCapacity,
			},
			MeterEnergy: &decorateModbusMbmdMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhasePowers: &decorateModbusMbmdPhasePowersImpl{
				phasePowers: phasePowers,
			},
		}

	case battery == nil && batteryCapacity != nil && meterEnergy == nil && phaseCurrents != nil && phasePowers != nil && phaseVoltages == nil:
		return &struct {
			api.Meter
			api.BatteryCapacity
			api.PhaseCurrents
			api.PhasePowers
		}{
			Meter: base,
			BatteryCapacity: &decorateModbusMbmdBatteryCapacityImpl{
				batteryCapacity: batteryCapacity,
			},
			PhaseCurrents: &decorateModbusMbmdPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
			PhasePowers: &decorateModbusMbmdPhasePowersImpl{
				phasePowers: phasePowers,
			},
		}

	case battery == nil && batteryCapacity != nil && meterEnergy != nil && phaseCurrents != nil && phasePowers != nil && phaseVoltages == nil:
		return &struct {
			api.Meter
			api.BatteryCapacity
			api.MeterEnergy
			api.PhaseCurrents
			api.PhasePowers
		}{
			Meter: base,
			BatteryCapacity: &decorateModbusMbmdBatteryCapacityImpl{
				batteryCapacity: batteryCapacity,
			},
			MeterEnergy: &decorateModbusMbmdMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhaseCurrents: &decorateModbusMbmdPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
			PhasePowers: &decorateModbusMbmdPhasePowersImpl{
				phasePowers: phasePowers,
			},
		}

	case battery == nil && batteryCapacity != nil && meterEnergy == nil && phaseCurrents == nil && phasePowers != nil && phaseVoltages != nil:
		return &struct {
			api.Meter
			api.BatteryCapacity
			api.PhasePowers
			api.PhaseVoltages
		}{
			Meter: base,
			BatteryCapacity: &decorateModbusMbmdBatteryCapacityImpl{
				batteryCapacity: batteryCapacity,
			},
			PhasePowers: &decorateModbusMbmdPhasePowersImpl{
				phasePowers: phasePowers,
			},
			PhaseVoltages: &decorateModbusMbmdPhaseVoltagesImpl{
				phaseVoltages: phaseVoltages,
			},
		}

	case battery == nil && batteryCapacity != nil && meterEnergy != nil && phaseCurrents == nil && phasePowers != nil && phaseVoltages != nil:
		return &struct {
			api.Meter
			api.BatteryCapacity
			api.MeterEnergy
			api.PhasePowers
			api.PhaseVoltages
		}{
			Meter: base,
			BatteryCapacity: &decorateModbusMbmdBatteryCapacityImpl{
				batteryCapacity: batteryCapacity,
			},
			MeterEnergy: &decorateModbusMbmdMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhasePowers: &decorateModbusMbmdPhasePowersImpl{
				phasePowers: phasePowers,
			},
			PhaseVoltages: &decorateModbusMbmdPhaseVoltagesImpl{
				phaseVoltages: phaseVoltages,
			},
		}

	case battery == nil && batteryCapacity != nil && meterEnergy == nil && phaseCurrents != nil && phasePowers != nil && phaseVoltages != nil:
		return &struct {
			api.Meter
			api.BatteryCapacity
			api.PhaseCurrents
			api.PhasePowers
			api.PhaseVoltages
		}{
			Meter: base,
			BatteryCapacity: &decorateModbusMbmdBatteryCapacityImpl{
				batteryCapacity: batteryCapacity,
			},
			PhaseCurrents: &decorateModbusMbmdPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
			PhasePowers: &decorateModbusMbmdPhasePowersImpl{
				phasePowers: phasePowers,
			},
			PhaseVoltages: &decorateModbusMbmdPhaseVoltagesImpl{
				phaseVoltages: phaseVoltages,
			},
		}

	case battery == nil && batteryCapacity != nil && meterEnergy != nil && phaseCurrents != nil && phasePowers != nil && phaseVoltages != nil:
		return &struct {
			api.Meter
			api.BatteryCapacity
			api.MeterEnergy
			api.PhaseCurrents
			api.PhasePowers
			api.PhaseVoltages
		}{
			Meter: base,
			BatteryCapacity: &decorateModbusMbmdBatteryCapacityImpl{
				batteryCapacity: batteryCapacity,
			},
			MeterEnergy: &decorateModbusMbmdMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhaseCurrents: &decorateModbusMbmdPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
			PhasePowers: &decorateModbusMbmdPhasePowersImpl{
				phasePowers: phasePowers,
			},
			PhaseVoltages: &decorateModbusMbmdPhaseVoltagesImpl{
				phaseVoltages: phaseVoltages,
			},
		}

	case battery != nil && batteryCapacity != nil && meterEnergy == nil && phaseCurrents == nil && phasePowers == nil && phaseVoltages == nil:
		return &struct {
			api.Meter
			api.Battery
			api.BatteryCapacity
		}{
			Meter: base,
			Battery: &decorateModbusMbmdBatteryImpl{
				battery: battery,
			},
			BatteryCapacity: &decorateModbusMbmdBatteryCapacityImpl{
				batteryCapacity: batteryCapacity,
			},
		}

	case battery != nil && batteryCapacity != nil && meterEnergy != nil && phaseCurrents == nil && phasePowers == nil && phaseVoltages == nil:
		return &struct {
			api.Meter
			api.Battery
			api.BatteryCapacity
			api.MeterEnergy
		}{
			Meter: base,
			Battery: &decorateModbusMbmdBatteryImpl{
				battery: battery,
			},
			BatteryCapacity: &decorateModbusMbmdBatteryCapacityImpl{
				batteryCapacity: batteryCapacity,
			},
			MeterEnergy: &decorateModbusMbmdMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
		}

	case battery != nil && batteryCapacity != nil && meterEnergy == nil && phaseCurrents != nil && phasePowers == nil && phaseVoltages == nil:
		return &struct {
			api.Meter
			api.Battery
			api.BatteryCapacity
			api.PhaseCurrents
		}{
			Meter: base,
			Battery: &decorateModbusMbmdBatteryImpl{
				battery: battery,
			},
			BatteryCapacity: &decorateModbusMbmdBatteryCapacityImpl{
				batteryCapacity: batteryCapacity,
			},
			PhaseCurrents: &decorateModbusMbmdPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
		}

	case battery != nil && batteryCapacity != nil && meterEnergy != nil && phaseCurrents != nil && phasePowers == nil && phaseVoltages == nil:
		return &struct {
			api.Meter
			api.Battery
			api.BatteryCapacity
			api.MeterEnergy
			api.PhaseCurrents
		}{
			Meter: base,
			Battery: &decorateModbusMbmdBatteryImpl{
				battery: battery,
			},
			BatteryCapacity: &decorateModbusMbmdBatteryCapacityImpl{
				batteryCapacity: batteryCapacity,
			},
			MeterEnergy: &decorateModbusMbmdMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhaseCurrents: &decorateModbusMbmdPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
		}

	case battery != nil && batteryCapacity != nil && meterEnergy == nil && phaseCurrents == nil && phasePowers == nil && phaseVoltages != nil:
		return &struct {
			api.Meter
			api.Battery
			api.BatteryCapacity
			api.PhaseVoltages
		}{
			Meter: base,
			Battery: &decorateModbusMbmdBatteryImpl{
				battery: battery,
			},
			BatteryCapacity: &decorateModbusMbmdBatteryCapacityImpl{
				batteryCapacity: batteryCapacity,
			},
			PhaseVoltages: &decorateModbusMbmdPhaseVoltagesImpl{
				phaseVoltages: phaseVoltages,
			},
		}

	case battery != nil && batteryCapacity != nil && meterEnergy != nil && phaseCurrents == nil && phasePowers == nil && phaseVoltages != nil:
		return &struct {
			api.Meter
			api.Battery
			api.BatteryCapacity
			api.MeterEnergy
			api.PhaseVoltages
		}{
			Meter: base,
			Battery: &decorateModbusMbmdBatteryImpl{
				battery: battery,
			},
			BatteryCapacity: &decorateModbusMbmdBatteryCapacityImpl{
				batteryCapacity: batteryCapacity,
			},
			MeterEnergy: &decorateModbusMbmdMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhaseVoltages: &decorateModbusMbmdPhaseVoltagesImpl{
				phaseVoltages: phaseVoltages,
			},
		}

	case battery != nil && batteryCapacity != nil && meterEnergy == nil && phaseCurrents != nil && phasePowers == nil && phaseVoltages != nil:
		return &struct {
			api.Meter
			api.Battery
			api.BatteryCapacity
			api.PhaseCurrents
			api.PhaseVoltages
		}{
			Meter: base,
			Battery: &decorateModbusMbmdBatteryImpl{
				battery: battery,
			},
			BatteryCapacity: &decorateModbusMbmdBatteryCapacityImpl{
				batteryCapacity: batteryCapacity,
			},
			PhaseCurrents: &decorateModbusMbmdPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
			PhaseVoltages: &decorateModbusMbmdPhaseVoltagesImpl{
				phaseVoltages: phaseVoltages,
			},
		}

	case battery != nil && batteryCapacity != nil && meterEnergy != nil && phaseCurrents != nil && phasePowers == nil && phaseVoltages != nil:
		return &struct {
			api.Meter
			api.Battery
			api.BatteryCapacity
			api.MeterEnergy
			api.PhaseCurrents
			api.PhaseVoltages
		}{
			Meter: base,
			Battery: &decorateModbusMbmdBatteryImpl{
				battery: battery,
			},
			BatteryCapacity: &decorateModbusMbmdBatteryCapacityImpl{
				batteryCapacity: batteryCapacity,
			},
			MeterEnergy: &decorateModbusMbmdMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhaseCurrents: &decorateModbusMbmdPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
			PhaseVoltages: &decorateModbusMbmdPhaseVoltagesImpl{
				phaseVoltages: phaseVoltages,
			},
		}

	case battery != nil && batteryCapacity != nil && meterEnergy == nil && phaseCurrents == nil && phasePowers != nil && phaseVoltages == nil:
		return &struct {
			api.Meter
			api.Battery
			api.BatteryCapacity
			api.PhasePowers
		}{
			Meter: base,
			Battery: &decorateModbusMbmdBatteryImpl{
				battery: battery,
			},
			BatteryCapacity: &decorateModbusMbmdBatteryCapacityImpl{
				batteryCapacity: batteryCapacity,
			},
			PhasePowers: &decorateModbusMbmdPhasePowersImpl{
				phasePowers: phasePowers,
			},
		}

	case battery != nil && batteryCapacity != nil && meterEnergy != nil && phaseCurrents == nil && phasePowers != nil && phaseVoltages == nil:
		return &struct {
			api.Meter
			api.Battery
			api.BatteryCapacity
			api.MeterEnergy
			api.PhasePowers
		}{
			Meter: base,
			Battery: &decorateModbusMbmdBatteryImpl{
				battery: battery,
			},
			BatteryCapacity: &decorateModbusMbmdBatteryCapacityImpl{
				batteryCapacity: batteryCapacity,
			},
			MeterEnergy: &decorateModbusMbmdMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhasePowers: &decorateModbusMbmdPhasePowersImpl{
				phasePowers: phasePowers,
			},
		}

	case battery != nil && batteryCapacity != nil && meterEnergy == nil && phaseCurrents != nil && phasePowers != nil && phaseVoltages == nil:
		return &struct {
			api.Meter
			api.Battery
			api.BatteryCapacity
			api.PhaseCurrents
			api.PhasePowers
		}{
			Meter: base,
			Battery: &decorateModbusMbmdBatteryImpl{
				battery: battery,
			},
			BatteryCapacity: &decorateModbusMbmdBatteryCapacityImpl{
				batteryCapacity: batteryCapacity,
			},
			PhaseCurrents: &decorateModbusMbmdPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
			PhasePowers: &decorateModbusMbmdPhasePowersImpl{
				phasePowers: phasePowers,
			},
		}

	case battery != nil && batteryCapacity != nil && meterEnergy != nil && phaseCurrents != nil && phasePowers != nil && phaseVoltages == nil:
		return &struct {
			api.Meter
			api.Battery
			api.BatteryCapacity
			api.MeterEnergy
			api.PhaseCurrents
			api.PhasePowers
		}{
			Meter: base,
			Battery: &decorateModbusMbmdBatteryImpl{
				battery: battery,
			},
			BatteryCapacity: &decorateModbusMbmdBatteryCapacityImpl{
				batteryCapacity: batteryCapacity,
			},
			MeterEnergy: &decorateModbusMbmdMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhaseCurrents: &decorateModbusMbmdPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
			PhasePowers: &decorateModbusMbmdPhasePowersImpl{
				phasePowers: phasePowers,
			},
		}

	case battery != nil && batteryCapacity != nil && meterEnergy == nil && phaseCurrents == nil && phasePowers != nil && phaseVoltages != nil:
		return &struct {
			api.Meter
			api.Battery
			api.BatteryCapacity
			api.PhasePowers
			api.PhaseVoltages
		}{
			Meter: base,
			Battery: &decorateModbusMbmdBatteryImpl{
				battery: battery,
			},
			BatteryCapacity: &decorateModbusMbmdBatteryCapacityImpl{
				batteryCapacity: batteryCapacity,
			},
			PhasePowers: &decorateModbusMbmdPhasePowersImpl{
				phasePowers: phasePowers,
			},
			PhaseVoltages: &decorateModbusMbmdPhaseVoltagesImpl{
				phaseVoltages: phaseVoltages,
			},
		}

	case battery != nil && batteryCapacity != nil && meterEnergy != nil && phaseCurrents == nil && phasePowers != nil && phaseVoltages != nil:
		return &struct {
			api.Meter
			api.Battery
			api.BatteryCapacity
			api.MeterEnergy
			api.PhasePowers
			api.PhaseVoltages
		}{
			Meter: base,
			Battery: &decorateModbusMbmdBatteryImpl{
				battery: battery,
			},
			BatteryCapacity: &decorateModbusMbmdBatteryCapacityImpl{
				batteryCapacity: batteryCapacity,
			},
			MeterEnergy: &decorateModbusMbmdMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhasePowers: &decorateModbusMbmdPhasePowersImpl{
				phasePowers: phasePowers,
			},
			PhaseVoltages: &decorateModbusMbmdPhaseVoltagesImpl{
				phaseVoltages: phaseVoltages,
			},
		}

	case battery != nil && batteryCapacity != nil && meterEnergy == nil && phaseCurrents != nil && phasePowers != nil && phaseVoltages != nil:
		return &struct {
			api.Meter
			api.Battery
			api.BatteryCapacity
			api.PhaseCurrents
			api.PhasePowers
			api.PhaseVoltages
		}{
			Meter: base,
			Battery: &decorateModbusMbmdBatteryImpl{
				battery: battery,
			},
			BatteryCapacity: &decorateModbusMbmdBatteryCapacityImpl{
				batteryCapacity: batteryCapacity,
			},
			PhaseCurrents: &decorateModbusMbmdPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
			PhasePowers: &decorateModbusMbmdPhasePowersImpl{
				phasePowers: phasePowers,
			},
			PhaseVoltages: &decorateModbusMbmdPhaseVoltagesImpl{
				phaseVoltages: phaseVoltages,
			},
		}

	case battery != nil && batteryCapacity != nil && meterEnergy != nil && phaseCurrents != nil && phasePowers != nil && phaseVoltages != nil:
		return &struct {
			api.Meter
			api.Battery
			api.BatteryCapacity
			api.MeterEnergy
			api.PhaseCurrents
			api.PhasePowers
			api.PhaseVoltages
		}{
			Meter: base,
			Battery: &decorateModbusMbmdBatteryImpl{
				battery: battery,
			},
			BatteryCapacity: &decorateModbusMbmdBatteryCapacityImpl{
				batteryCapacity: batteryCapacity,
			},
			MeterEnergy: &decorateModbusMbmdMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhaseCurrents: &decorateModbusMbmdPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
			PhasePowers: &decorateModbusMbmdPhasePowersImpl{
				phasePowers: phasePowers,
			},
			PhaseVoltages: &decorateModbusMbmdPhaseVoltagesImpl{
				phaseVoltages: phaseVoltages,
			},
		}
	}

	return nil
}

type decorateModbusMbmdBatteryImpl struct {
	battery func() (float64, error)
}

func (impl *decorateModbusMbmdBatteryImpl) Soc() (float64, error) {
	return impl.battery()
}

type decorateModbusMbmdBatteryCapacityImpl struct {
	batteryCapacity func() float64
}

func (impl *decorateModbusMbmdBatteryCapacityImpl) Capacity() float64 {
	return impl.batteryCapacity()
}

type decorateModbusMbmdMeterEnergyImpl struct {
	meterEnergy func() (float64, error)
}

func (impl *decorateModbusMbmdMeterEnergyImpl) TotalEnergy() (float64, error) {
	return impl.meterEnergy()
}

type decorateModbusMbmdPhaseCurrentsImpl struct {
	phaseCurrents func() (float64, float64, float64, error)
}

func (impl *decorateModbusMbmdPhaseCurrentsImpl) Currents() (float64, float64, float64, error) {
	return impl.phaseCurrents()
}

type decorateModbusMbmdPhasePowersImpl struct {
	phasePowers func() (float64, float64, float64, error)
}

func (impl *decorateModbusMbmdPhasePowersImpl) Powers() (float64, float64, float64, error) {
	return impl.phasePowers()
}

type decorateModbusMbmdPhaseVoltagesImpl struct {
	phaseVoltages func() (float64, float64, float64, error)
}

func (impl *decorateModbusMbmdPhaseVoltagesImpl) Voltages() (float64, float64, float64, error) {
	return impl.phaseVoltages()
}
