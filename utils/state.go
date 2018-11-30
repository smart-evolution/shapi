package utils

import (
    "github.com/smart-evolution/smarthome/services/dataflux"
    "github.com/smart-evolution/smarthome/services/email"
    "github.com/smart-evolution/smarthome/services/persistence"
)

var (
    // IsAlerts - are alerts turned on
    IsAlerts    bool
    // SendAlert - should alerts be emailed
    SendAlert   bool
    // DataFlux - data persistance
    DataFlux    dataflux.DataFlux
    // Persistance - data persistance
    Persistance persistence.Persistance
    // Mailer
    Mailer      email.Mailer
)
