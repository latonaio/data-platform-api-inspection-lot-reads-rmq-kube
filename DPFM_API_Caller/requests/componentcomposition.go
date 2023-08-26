package requests

InspectionLot								int		`json:InspectionLot`
ComponentCompositionType					string	`json:ComponentCompositionType`
ComponentCompositionUperLimitInPercent		float32	`json:ComponentCompositionUperLimitInPercent`
ComponentCompositionLowerLimitInPercent		float32	`json:ComponentCompositionLowerLimitInPercent`
ComponentCompositionStandardValueInPercent	float32	`json:ComponentCompositionStandardValueInPercent`
Creationstring								string	`json:Creationstring`
LastChangestring							string	`json:LastChangestring`
IsCancelled 								*int	`json:IsCancelled `
IsMarkedForDeletion							*int	`json:IsMarkedForDeletion`

}
