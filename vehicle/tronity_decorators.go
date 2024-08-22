package vehicle

// Code generated by github.com/evcc-io/evcc/cmd/tools/decorate.go. DO NOT EDIT.

import (
	"github.com/evcc-io/evcc/api"
)

func decorateTronity(base *Tronity, chargeState func() (api.ChargeStatus, error), vehicleOdometer func() (float64, error), chargeController func(bool) error) api.Vehicle {
	switch {
	case chargeController == nil && chargeState == nil && vehicleOdometer == nil:
		return base

	case chargeController == nil && chargeState != nil && vehicleOdometer == nil:
		return &struct {
			*Tronity
			api.ChargeState
		}{
			Tronity: base,
			ChargeState: &decorateTronityChargeStateImpl{
				chargeState: chargeState,
			},
		}

	case chargeController == nil && chargeState == nil && vehicleOdometer != nil:
		return &struct {
			*Tronity
			api.VehicleOdometer
		}{
			Tronity: base,
			VehicleOdometer: &decorateTronityVehicleOdometerImpl{
				vehicleOdometer: vehicleOdometer,
			},
		}

	case chargeController == nil && chargeState != nil && vehicleOdometer != nil:
		return &struct {
			*Tronity
			api.ChargeState
			api.VehicleOdometer
		}{
			Tronity: base,
			ChargeState: &decorateTronityChargeStateImpl{
				chargeState: chargeState,
			},
			VehicleOdometer: &decorateTronityVehicleOdometerImpl{
				vehicleOdometer: vehicleOdometer,
			},
		}

	case chargeController != nil && chargeState == nil && vehicleOdometer == nil:
		return &struct {
			*Tronity
			api.ChargeController
		}{
			Tronity: base,
			ChargeController: &decorateTronityChargeControllerImpl{
				chargeController: chargeController,
			},
		}

	case chargeController != nil && chargeState != nil && vehicleOdometer == nil:
		return &struct {
			*Tronity
			api.ChargeController
			api.ChargeState
		}{
			Tronity: base,
			ChargeController: &decorateTronityChargeControllerImpl{
				chargeController: chargeController,
			},
			ChargeState: &decorateTronityChargeStateImpl{
				chargeState: chargeState,
			},
		}

	case chargeController != nil && chargeState == nil && vehicleOdometer != nil:
		return &struct {
			*Tronity
			api.ChargeController
			api.VehicleOdometer
		}{
			Tronity: base,
			ChargeController: &decorateTronityChargeControllerImpl{
				chargeController: chargeController,
			},
			VehicleOdometer: &decorateTronityVehicleOdometerImpl{
				vehicleOdometer: vehicleOdometer,
			},
		}

	case chargeController != nil && chargeState != nil && vehicleOdometer != nil:
		return &struct {
			*Tronity
			api.ChargeController
			api.ChargeState
			api.VehicleOdometer
		}{
			Tronity: base,
			ChargeController: &decorateTronityChargeControllerImpl{
				chargeController: chargeController,
			},
			ChargeState: &decorateTronityChargeStateImpl{
				chargeState: chargeState,
			},
			VehicleOdometer: &decorateTronityVehicleOdometerImpl{
				vehicleOdometer: vehicleOdometer,
			},
		}
	}

	return nil
}

type decorateTronityChargeControllerImpl struct {
	chargeController func(bool) error
}

func (impl *decorateTronityChargeControllerImpl) ChargeEnable(p0 bool) error {
	return impl.chargeController(p0)
}

type decorateTronityChargeStateImpl struct {
	chargeState func() (api.ChargeStatus, error)
}

func (impl *decorateTronityChargeStateImpl) Status() (api.ChargeStatus, error) {
	return impl.chargeState()
}

type decorateTronityVehicleOdometerImpl struct {
	vehicleOdometer func() (float64, error)
}

func (impl *decorateTronityVehicleOdometerImpl) Odometer() (float64, error) {
	return impl.vehicleOdometer()
}
