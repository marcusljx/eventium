package server
/* AUTOGENERATED USING grpcgen
 *  add your server logic here!
 * For template changes/edits/bugs/new features, please contact marcusljx@gmail.com
 */

import (
	"github.com/marcusljx/eventium/eventium"
	"golang.org/x/net/context"
)


func (s *eventiumLogic) PostEvent(ctx context.Context, req *eventium.PostEventRequest) (*eventium.PostEventResponse, error) {
    //TODO: Implement logic for eventiumLogic.PostEvent
    return nil, nil
}


func (s *eventiumLogic) GetEventByID(ctx context.Context, req *eventium.GetEventByIDRequest) (*eventium.Eventum, error) {
    //TODO: Implement logic for eventiumLogic.GetEventByID
    return nil, nil
}


func (s *eventiumLogic) GetEventByNearbyLocation(ctx context.Context, req *eventium.Location) (*eventium.Eventum, error) {
    //TODO: Implement logic for eventiumLogic.GetEventByNearbyLocation
    return nil, nil
}

