package requests

type SpecGeneral struct {
	InspectionLot	    	int	    `json:"InspectionLot"`
    HeatNumber	        	string	`json:"HeatNumber"`
	CreationDate            string   `json:"CreationDate"`
	CreationTime            string   `json:"CreationTime"`
	LastChangeDate          string   `json:"LastChangeDate"`
	LastChangeTime          string   `json:"LastChangeTime"`
	IsReleased              *bool    `json:"IsReleased"`
	IsLocked                *bool    `json:"IsLocked"`
	IsCancelled             *bool    `json:"IsCancelled"`
	IsMarkedForDeletion     *bool    `json:"IsMarkedForDeletion"`
}
