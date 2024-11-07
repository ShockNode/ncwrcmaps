package fishingareas

import (
	"context"
	"net/http"
	"strconv"

	"github.com/ShockNode/ncwrcmaps/pkg/httphelper"
)

type Location struct {
	ID int `json:"locationID"`
	/*
		Provided in Partnership with WRC -- NCWRC fishing regulations apply at these areas.
		Non-WRC Affiliated -- These public access areas are provided as a convenience to anglers, but are not affiliated with the Commission.

		const locationTypes = {
		   BAA: 1,
		   PFA: 3,
		   Non_WRC: 7
		};

		if (wrcSite) {
		    // WRC_SPONSORED (1)
		} else if (locationTypeID == Non_WRC) {
		    // NON_AFFILIATED (3)
		} else {
		    // WRC_PARTNER (2)
		}
	*/
	TypeID        int     `json:"locationTypeID"`
	Name          string  `json:"locationName"`
	Latitude      float64 `json:"latitude"`
	Longitude     float64 `json:"longitude"`
	WaterBodyName string  `json:"waterBodyName"`
	WRCSite       bool    `json:"wrcSite"`
	OperatedBy    string  `json:"operatedBy"`
}

type LocationDetails struct {
	Location
	WheelchairAccessible bool          `json:"wheelchairAccessible"`
	CanoeAccess          bool          `json:"canoeAccess"`
	FishingPierAccess    bool          `json:"fishingPierAccess"`
	ShorelineAccess      bool          `json:"shorelineAccess"`
	BoatRamp             bool          `json:"boatRamp"`
	TroutStocked         bool          `json:"troutStocked"`
	WarmWaterFishStocked bool          `json:"warmWaterFishStocked"`
	Description          string        `json:"description"`
	County               string        `json:"county"`
	Directions           string        `json:"directions"`
	LatitudeDMSString    string        `json:"latitudeDMSString"`
	LongitudeDMSString   string        `json:"longitudeDMSString"`
	Management           string        `json:"management"`
	OwnerName            string        `json:"ownerName"`
	LocationPhotoID      int           `json:"locationPhotoID"`
	SpeciesInfo          []SpeciesInfo `json:"speciesInfo"`
	WaterbodyInfo        WaterbodyInfo `json:"waterbodyInfo"`
}

type SpeciesInfo struct {
	CommonName string `json:"commonName"`
	Stocked    bool   `json:"stocked"`
	Wild       bool   `json:"wild"`
}

type WaterbodyInfo struct {
	Name string `json:"waterbodyName"`
	/*
	    1, 2 - "Stream/River";
	    3, 4 - "Lake/Pond";
	   	5, 6 - "Ocean/Sound";
	*/
	TypeID      int     `json:"waterbodyTypeID"`
	SizeInAcres float64 `json:"sizeInAcres"`
}

var UrlPaws = "https://ncpaws.org/NCWRCMaps/FishingAreas/Home"
var EndpointGetFishingAreaInfo = "/GetFishingAreaInfo"

type Requestor struct {
	Client http.Client
}

// GetFishingAreaInfo
func (requestor *Requestor) GetFishingAreaInfo(ctx context.Context, locationID int) (*LocationDetails, error) {
	url := UrlPaws + EndpointGetFishingAreaInfo
	details, err := httphelper.Get[LocationDetails](ctx, requestor.Client, url, nil, map[string]string{"locationID": strconv.Itoa(locationID)})
	if err != nil {
		return nil, err
	}
	return details, nil
}
