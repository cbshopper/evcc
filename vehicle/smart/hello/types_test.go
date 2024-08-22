package hello

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

const data = `{
    "code": "1000",
    "data": {
        "result": {
            "serviceResult": {
                "error": null,
                "operationResult": 1
            },
            "sessionId": "PS0485700000000290629****"
        },
        "vehicleStatus": {
            "basicVehicleStatus": {
                "usageMode": "1",
                "engineStatus": "engine_off",
                "position": {
                    "altitude": "",
                    "posCanBeTrusted": "",
                    "latitude": "",
                    "carLocatorStatUploadEn": "false",
                    "marsCoordinates": "",
                    "longitude": ""
                },
                "carMode": "0",
                "speed": "0.0",
                "speedValidity": "false",
                "direction": "0"
            },
            "notification": {
                "notifForEmgyCallStatus": "0"
            },
            "eg": {
                "enableRunning": "false",
                "blocked": {
                    "status": "0"
                },
                "panicStatus": "false"
            },
            "parkTime": {
                "status": "1723920451549"
            },
            "theftNotification": {
                "time": "1716550899",
                "activated": "2"
            },
            "configuration": {
                "propulsionType": "4",
                "fuelType": "4",
                "vin": "HESX****"
            },
            "updateTime": "1723990612062",
            "additionalVehicleStatus": {
                "maintenanceStatus": {
                    "tyreTempWarningPassengerRear": "0",
                    "daysToService": "191",
                    "engineHrsToService": "500",
                    "odometer": "4973.000",
                    "brakeFluidLevelStatus": "3",
                    "tyreTempDriverRear": "22.000",
                    "tyreTempWarningPassenger": "0",
                    "tyreTempWarningDriverRear": "0",
                    "mainBatteryStatus": {
                        "stateOfCharge": "1",
                        "chargeLevel": "96.2",
                        "energyLevel": "0",
                        "stateOfHealth": "0",
                        "powerLevel": "0",
                        "voltage": "14.275"
                    },
                    "tyreTempDriver": "22.000",
                    "tyreTempPassengerRear": "22.000",
                    "tyrePreWarningDriver": "0",
                    "distanceToService": "25027",
                    "tyrePreWarningPassengerRear": "0",
                    "tyreTempWarningDriver": "0",
                    "tyreStatusPassengerRear": "244.394",
                    "tyreStatusPassenger": "241.648",
                    "tyreStatusDriverRear": "241.648",
                    "serviceWarningStatus": "0",
                    "tyreStatusDriver": "248.513",
                    "tyreTempPassenger": "23.000",
                    "tyrePreWarningDriverRear": "0",
                    "tyrePreWarningPassenger": "0",
                    "washerFluidLevelStatus": "1"
                },
                "electricVehicleStatus": {
                    "disChargeUAct": "0.0",
                    "disChargeSts": "0",
                    "wptFineAlignt": "0",
                    "chargeLidAcStatus": "2",
                    "distanceToEmptyOnBatteryOnly": "201",
                    "distanceToEmptyOnBattery100Soc": "402",
                    "chargeSts": "0",
                    "averPowerConsumption": "-102.3",
                    "chargerState": "4",
                    "timeToTargetDisCharged": "2047",
                    "distanceToEmptyOnBattery20Soc": "80",
                    "disChargeConnectStatus": "2",
                    "chargeLidDcAcStatus": "1",
                    "dcChargeSts": "0",
                    "ptReady": "0",
                    "chargeLevel": "50",
                    "statusOfChargerConnection": "2",
                    "dcDcActvd": "1",
                    "indPowerConsumption": "0.0",
                    "dcDcConnectStatus": "0",
                    "disChargeIAct": "0.0",
                    "dcChargeIAct": "0.5",
                    "chargeUAct": "0.0",
                    "bookChargeSts": "0",
                    "chargeIAct": "0.000",
                    "timeToFullyCharged": "2047"
                },
                "chargeHvSts": "1",
                "drivingBehaviourStatus": {
                    "gearAutoStatus": "0",
                    "gearManualStatus": "0",
                    "engineSpeed": "0.000"
                },
                "runningStatus": {
                    "ahbc": "0",
                    "goodbye": "0",
                    "homeSafe": "0",
                    "cornrgLi": "0",
                    "frntFog": "0",
                    "stopLi": "0",
                    "tripMeter1": "4118.8",
                    "approach": "0",
                    "tripMeter2": "0.0",
                    "indFuelConsumption": "0",
                    "hiBeam": "0",
                    "engineCoolantLevelStatus": "3",
                    "fuelEnLevel": "0",
                    "loBeam": "0",
                    "posLiRe": "0",
                    "ltgShow": "0",
                    "welcome": "0",
                    "drl": "0",
                    "fuelLevelPct": "0",
                    "ahl": "0",
                    "fuelEnCns": "0",
                    "trunIndrLe": "0",
                    "trunIndrRi": "0",
                    "afs": "0",
                    "dbl": "0",
                    "avgSpeed": "24",
                    "posLiFrnt": "0",
                    "reverseLi": "0",
                    "hwl": "0",
                    "reFog": "0",
                    "flash": "0",
                    "allwl": "0",
                    "fuelEnCnsFild": "0"
                },
                "trailerStatus": {
                    "trailerTurningLampSts": "0",
                    "trailerFogLampSts": "0",
                    "trailerBreakLampSts": "0",
                    "trailerReversingLampSts": "0",
                    "trailerPosLampSts": "0"
                },
                "climateStatus": {
                    "drvHeatSts": "0",
                    "winPosDriver": "0",
                    "rrVentDetail": "0",
                    "rlVentSts": "0",
                    "passVentSts": "0",
                    "interiorTemp": "24.900",
                    "passVentDetail": "0",
                    "sunroofPos": "101",
                    "cdsClimateActive": "false",
                    "sunroofOpenStatus": "1",
                    "rrHeatingDetail": "0",
                    "winStatusPassenger": "2",
                    "fragActive": false,
                    "winStatusDriver": "2",
                    "drvVentSts": "0",
                    "winStatusPassengerRear": "2",
                    "sunCurtainRearOpenStatus": "1",
                    "preClimateActive": false,
                    "rlHeatingDetail": "0",
                    "winPosPassengerRear": "0",
                    "curtainPos": "0",
                    "rlVentDetail": "0",
                    "curtainOpenStatus": "1",
                    "climateOverHeatProActive": "true",
                    "rrVentSts": "0",
                    "rrHeatingSts": "0",
                    "winPosPassenger": "0",
                    "steerWhlHeatingSts": "2",
                    "drvVentDetail": "0",
                    "winPosDriverRear": "0",
                    "exteriorTemp": "23.500",
                    "rlHeatingSts": "0",
                    "winStatusDriverRear": "2",
                    "defrost": "false",
                    "drvHeatDetail": "2",
                    "passHeatingDetail": "2",
                    "airBlowerActive": "false",
                    "sunCurtainRearPos": "101",
                    "passHeatingSts": "0"
                },
                "drivingSafetyStatus": {
                    "doorLockStatusDriverRear": "1",
                    "srsCrashStatus": "0",
                    "doorOpenStatusPassengerRear": "0",
                    "doorPosPassengerRear": "0",
                    "doorOpenStatusDriver": "0",
                    "seatBeltStatusPassenger": "false",
                    "doorPosDriver": "0",
                    "seatBeltStatusThPassengerRear": "false",
                    "electricParkBrakeStatus": "1",
                    "doorLockStatusDriver": "1",
                    "seatBeltStatusThDriverRear": "false",
                    "tankFlapStatus": "2",
                    "seatBeltStatusPassengerRear": "false",
                    "doorOpenStatusPassenger": "0",
                    "doorPosPassenger": "0",
                    "vehicleAlarm": {
                        "alrmSt": "1",
                        "alrmTrgSrc": "7"
                    },
                    "doorPosDriverRear": "0",
                    "centralLockingStatus": "2",
                    "seatBeltStatusDriver": "false",
                    "doorLockStatusPassenger": "1",
                    "seatBeltStatusMidRear": "false",
                    "trunkLockStatus": "1",
                    "seatBeltStatusDriverRear": "false",
                    "engineHoodOpenStatus": "0",
                    "doorOpenStatusDriverRear": "0",
                    "doorLockStatusPassengerRear": "1",
                    "trunkOpenStatus": "0"
                },
                "pollutionStatus": {
                    "interiorPM25": "1",
                    "interiorSecondPM25Level": "0",
                    "interiorPM25Level": "0",
                    "relHumSts": "80",
                    "exteriorPM25Level": "0"
                }
            },
            "temStatus": {
                "swVersion": null,
                "serialNumber": null,
                "powerSource": null,
                "networkAccessStatus": {
                    "mobileNetwork": null,
                    "simInfo": {
                        "iccId": null,
                        "imsi": null,
                        "msisdn": null
                    }
                },
                "mcuVersion": null,
                "mpuVersion": null,
                "backupBattery": {
                    "stateOfCharge": null,
                    "stateOfHealth": null,
                    "voltage": null
                },
                "hwVersion": null,
                "powerMode": null,
                "healthStatus": null,
                "sleepCycleNextWakeupTime": null,
                "rvsEnable": "true",
                "imei": null,
                "state": null,
                "connectivityStatus": null
            }
        }
    },
    "success": true,
    "hint": null,
    "httpStatus": "OK",
    "sessionId": "6fab3f23b494976b2101eaff1f2c87cf",
    "message": "operation succeed"
}`

func TestUnmarshal(t *testing.T) {
	var res VehicleStatus
	require.NoError(t, json.Unmarshal([]byte(data), &res))
}
