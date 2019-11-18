package route

import (
    "github.com/gorilla/mux"
    messaging "github.com/oms-services/intercom/messaging"
    "log"
    "net/http"
)

type Route struct {
    Name        string
    Method      string
    Pattern     string
    HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
    Route{
        "CreateUser",
        "POST",
        "/createuser",
        messaging.CreateUser,
    },
    Route{
        "SendInAppMessage",
        "POST",
        "/inappmessage",
        messaging.SendInAppMessage,
    },
    Route{
        "SendEmailMessage",
        "POST",
        "/emailmessage",
        messaging.SendEmailMessage,
    },
    Route{
        "SendUserMessage",
        "POST",
        "/usermessage",
        messaging.SendUserMessage,
    },
}

func NewRouter() *mux.Router {
    router := mux.NewRouter().StrictSlash(true)
    for _, route := range routes {
        var handler http.Handler
        log.Println(route.Name)
        handler = route.HandlerFunc

        router.
            Methods(route.Method).
            Path(route.Pattern).
            Name(route.Name).
            Handler(handler)
    }
    return router
}
