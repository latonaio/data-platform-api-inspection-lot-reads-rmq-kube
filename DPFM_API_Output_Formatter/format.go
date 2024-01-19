package dpfm_api_output_formatter

import (
	"data-platform-api-inspection-lot-reads-rmq-kube/DPFM_API_Caller/requests"
	"database/sql"
	"fmt"
)
func ConvertToHeader(rows *sql.Rows) (*[]Header, error) {
	defer rows.Close()
	header := make([]Header, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.Header{}

		err := rows.Scan(
			&pm.InspectionLot,
			&pm.InspectionLotDate,
			&pm.InspectionPlan,
			&pm.InspectionPlantBusinessPartner,
			&pm.InspectionPlant,
			&pm.Product,
			&pm.ProductSpecification,
			&pm.InspectionSpecification,
			&pm.ProductionOrder,
			&pm.ProductionOrderItem,
			&pm.InspectionLotHeaderText,
			&pm.ExternalReferenceDocument,
			&pm.CertificateAuthorityChain,
			&pm.UsageControlChain,
			&pm.CreationDate,
			&pm.CreationTime,
			&pm.LastChangeDate,
			&pm.LastChangeTime,
			&pm.IsReleased,
			&pm.IsPartiallyConfirmed,
			&pm.IsConfirmed,
			&pm.IsLocked,
			&pm.IsCancelled,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &header, err
		}

		data := pm
		header = append(header, Header{
			InspectionLot:                  data.InspectionLot,
			InspectionLotDate:              data.InspectionLotDate,
			InspectionPlan:                 data.InspectionPlan,
			InspectionPlantBusinessPartner: data.InspectionPlantBusinessPartner,
			InspectionPlant:                data.InspectionPlant,
			Product:                        data.Product,
			ProductSpecification:           data.ProductSpecification,
			InspectionSpecification:        data.InspectionSpecification,
			ProductionOrder:                data.ProductionOrder,
			ProductionOrderItem:            data.ProductionOrderItem,
			InspectionLotHeaderText:        data.InspectionLotHeaderText,
			ExternalReferenceDocument:		data.ExternalReferenceDocument,
			CertificateAuthorityChain:		data.CertificateAuthorityChain,
			UsageControlChain:		  		data.UsageControlChain,
			CreationDate:                   data.CreationDate,
			CreationTime:                   data.CreationTime,
			LastChangeDate:                 data.LastChangeDate,
			LastChangeTime:                 data.LastChangeTime,
			IsReleased:                     data.IsReleased,
			IsPartiallyConfirmed:           data.IsPartiallyConfirmed,
			IsConfirmed:                    data.IsConfirmed,
			IsLocked:                       data.IsLocked,
			IsCancelled:                    data.IsCancelled,
			IsMarkedForDeletion:            data.IsMarkedForDeletion,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &header, nil
	}

	return &header, nil
}

func ConvertToSpecGeneral(rows *sql.Rows) (*[]SpecGeneral, error) {
	defer rows.Close()
	specGeneral := make([]SpecGeneral, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.SpecGeneral{}

		err := rows.Scan(
			&pm.InspectionLot,
			&pm.HeatNumber,
			&pm.CreationDate,
			&pm.CreationTime,
			&pm.LastChangeDate,
			&pm.LastChangeTime,
			&pm.IsReleased,
			&pm.IsLocked,
			&pm.IsCancelled,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &specGeneral, err
		}

		data := pm
		specGeneral = append(specGeneral, SpecGeneral{
			InspectionLot:       	data.InspectionLot,
			HeatNumber:          	data.HeatNumber,
			CreationDate:           data.CreationDate,
			CreationTime:           data.CreationTime,
			LastChangeDate:         data.LastChangeDate,
			LastChangeTime:         data.LastChangeTime,
			IsReleased:             data.IsReleased,
			IsLocked:               data.IsLocked,
			IsCancelled:            data.IsCancelled,
			IsMarkedForDeletion:    data.IsMarkedForDeletion,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &specGeneral, nil
	}

	return &specGeneral, nil
}

func ConvertToSpecDetail(rows *sql.Rows) (*[]SpecDetail, error) {
	defer rows.Close()
	specDetail := make([]SpecDetail, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.SpecDetail{}

		err := rows.Scan(
			&pm.InspectionLot,
			&pm.SpecType,
			&pm.UpperLimitValue,
			&pm.LowerLimitValue,
			&pm.StandardValue,
			&pm.SpecTypeUnit,
			&pm.Formula,
			&pm.CreationDate,
			&pm.CreationTime,
			&pm.LastChangeDate,
			&pm.LastChangeTime,
			&pm.IsReleased,
			&pm.IsLocked,
			&pm.IsCancelled,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &specDetail, err
		}

		data := pm
		specDetail = append(specDetail, SpecDetail{
			InspectionLot:       data.InspectionLot,
			SpecType:            data.SpecType,
			UpperLimitValue:     data.UpperLimitValue,
			LowerLimitValue:     data.LowerLimitValue,
			StandardValue:       data.StandardValue,
			SpecTypeUnit:        data.SpecTypeUnit,
			Formula:	         data.Formula,
			CreationDate:           data.CreationDate,
			CreationTime:           data.CreationTime,
			LastChangeDate:         data.LastChangeDate,
			LastChangeTime:         data.LastChangeTime,
			IsReleased:             data.IsReleased,
			IsLocked:               data.IsLocked,
			IsCancelled:            data.IsCancelled,
			IsMarkedForDeletion:    data.IsMarkedForDeletion,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &specDetail, nil
	}

	return &specDetail, nil
}

func ConvertToComponentComposition(rows *sql.Rows) (*[]ComponentComposition, error) {
	defer rows.Close()
	componentComposition := make([]ComponentComposition, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.ComponentComposition{}

		err := rows.Scan(
			&pm.InspectionLot,
			&pm.ComponentCompositionType,
			&pm.ComponentCompositionUpperLimitInPercent,
			&pm.ComponentCompositionLowerLimitInPercent,
			&pm.ComponentCompositionStandardValueInPercent,
			&pm.CreationDate,
			&pm.CreationTime,
			&pm.LastChangeDate,
			&pm.LastChangeTime,
			&pm.IsReleased,
			&pm.IsLocked,
			&pm.IsCancelled,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &componentComposition, err
		}

		data := pm
		componentComposition = append(componentComposition, ComponentComposition{
			InspectionLot:                              data.InspectionLot,
			ComponentCompositionType:                   data.ComponentCompositionType,
			ComponentCompositionUpperLimitInPercent:     data.ComponentCompositionUpperLimitInPercent,
			ComponentCompositionLowerLimitInPercent:    data.ComponentCompositionLowerLimitInPercent,
			ComponentCompositionStandardValueInPercent: data.ComponentCompositionStandardValueInPercent,
			CreationDate:           					data.CreationDate,
			CreationTime:           					data.CreationTime,
			LastChangeDate:         					data.LastChangeDate,
			LastChangeTime:         					data.LastChangeTime,
			IsReleased:             					data.IsReleased,
			IsLocked:               					data.IsLocked,
			IsCancelled:            					data.IsCancelled,
			IsMarkedForDeletion:    					data.IsMarkedForDeletion,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &componentComposition, nil
	}

	return &componentComposition, nil
}

func ConvertToInspection(rows *sql.Rows) (*[]Inspection, error) {
	defer rows.Close()
	inspection := make([]Inspection, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.Inspection{}

		err := rows.Scan(
			&pm.InspectionLot,
			&pm.Inspection,
			&pm.InspectionDate,
			&pm.InspectionType,
			&pm.InspectionTypeValueUnit,
			&pm.InspectionTypePlannedValue,
			&pm.InspectionTypeCertificateType,
			&pm.InspectionTypeCertificateValueInText,
			&pm.InspectionTypeCertificateValueInQuantity,
			&pm.InspectionLotInspectionText,
			&pm.CreationDate,
			&pm.CreationTime,
			&pm.LastChangeDate,
			&pm.LastChangeTime,
			&pm.IsReleased,
			&pm.IsLocked,
			&pm.IsCancelled,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &inspection, err
		}

		data := pm
		inspection = append(inspection, Inspection{
			InspectionLot:                            data.InspectionLot,
			Inspection:                               data.Inspection,
			InspectionDate:                           data.InspectionDate,
			InspectionType:                           data.InspectionType,
			InspectionTypeValueUnit:                  data.InspectionTypeValueUnit,
			InspectionTypePlannedValue:               data.InspectionTypePlannedValue,
			InspectionTypeCertificateType:            data.InspectionTypeCertificateType,
			InspectionTypeCertificateValueInText:     data.InspectionTypeCertificateValueInText,
			InspectionTypeCertificateValueInQuantity: data.InspectionTypeCertificateValueInQuantity,
			InspectionLotInspectionText:              data.InspectionLotInspectionText,
			CreationDate:           				  data.CreationDate,
			CreationTime:           				  data.CreationTime,
			LastChangeDate:         				  data.LastChangeDate,
			LastChangeTime:         				  data.LastChangeTime,
			IsReleased:             				  data.IsReleased,
			IsLocked:               				  data.IsLocked,
			IsCancelled:            				  data.IsCancelled,
			IsMarkedForDeletion:    				  data.IsMarkedForDeletion,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &inspection, nil
	}

	return &inspection, nil
}

func ConvertToOperation(rows *sql.Rows) (*[]Operation, error) {
	defer rows.Close()
	operation := make([]Operation, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.Operation{}

		err := rows.Scan(
			&pm.InspectionLot,
			&pm.Operations,
			&pm.OperationsItem,
			&pm.OperationID,
			&pm.Inspection,
			&pm.OperationType,
			&pm.SupplyChainRelationshipID,
			&pm.SupplyChainRelationshipDeliveryID,
			&pm.SupplyChainRelationshipDeliveryPlantID,
			&pm.SupplyChainRelationshipProductionPlantID,
			&pm.Product,
			&pm.Buyer,
			&pm.Seller,
			&pm.DeliverToParty,
			&pm.DeliverToPlant,
			&pm.DeliverFromParty,
			&pm.DeliverFromPlant,
			&pm.InspectionPlantBusinessPartner,
			&pm.InspectionPlant,
			&pm.Sequence,
			&pm.SequenceText,
			&pm.OperationText,
			&pm.OperationStatus,
			&pm.ResponsiblePlannerGroup,
			&pm.OperationUnit,
			&pm.StandardLotSizeQuantity,
			&pm.MinimumLotSizeQuantity,
			&pm.MaximumLotSizeQuantity,
			&pm.PlainLongText,
			&pm.WorkCenter,
			&pm.CapacityCategoryCode,
			&pm.OperationCostingRelevancyType,
			&pm.OperationSetupType,
			&pm.OperationSetupGroupCategory,
			&pm.OperationSetupGroup,
			&pm.OperationReferenceQuantity,
			&pm.MaximumWaitDuration,
			&pm.StandardWaitDuration,
			&pm.MinimumWaitDuration,
			&pm.WaitDurationUnit,
			&pm.MaximumQueueDuration,
			&pm.StandardQueueDuration,
			&pm.MinimumQueueDuration,
			&pm.QueueDurationUnit,
			&pm.MaximumMoveDuration,
			&pm.StandardMoveDuration,
			&pm.MinimumMoveDuration,
			&pm.MoveDurationUnit,
			&pm.StandardDeliveryDuration,
			&pm.StandardDeliveryDurationUnit,
			&pm.StandardOperationScrapPercent,
			&pm.PlannedOperationStandardValue,
			&pm.PlannedOperationLowerValue,
			&pm.PlannedOperationUpperValue,
			&pm.PlannedOperationValueUnit,
			&pm.CostElement,
			&pm.OperationErlstSchedldExecStrtDte,
			&pm.OperationErlstSchedldExecStrtTme,
			&pm.OperationErlstSchedldExecEndDte,
			&pm.OperationErlstSchedldExecEndTme,
			&pm.OperationActualExecutionStartDate,
			&pm.OperationActualExecutionStartTime,
			&pm.OperationActualExecutionEndDate,
			&pm.OperationActualExecutionEndTime,
			&pm.OperationConfirmedYieldQuantityInBaseUnit,
			&pm.OperationConfirmedYieldQuantityInProductionUnit,
			&pm.OperationConfirmedYieldQuantityInOperationUnit,
			&pm.ValidityStartDate,
			&pm.ValidityEndDate,
			&pm.CreationDate,
			&pm.CreationTime,
			&pm.LastChangeDate,
			&pm.LastChangeTime,
			&pm.IsReleased,
			&pm.IsPartiallyConfirmed,
			&pm.IsConfirmed,
			&pm.IsLocked,
			&pm.IsCancelled,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &operation, err
		}

		data := pm
		operation = append(operation, Operation{
			InspectionLot:                            				data.InspectionLot,
			Operations:                               				data.Operations,
			OperationsItem:                           				data.OperationsItem,
			OperationID:                              				data.OperationID,
			Inspection:                               				data.Inspection,
			OperationType:                            				data.OperationType,
			SupplyChainRelationshipID:                				data.SupplyChainRelationshipID,
			SupplyChainRelationshipDeliveryID:        				data.SupplyChainRelationshipDeliveryID,
			SupplyChainRelationshipDeliveryPlantID:   				data.SupplyChainRelationshipDeliveryPlantID,
			SupplyChainRelationshipProductionPlantID: 				data.SupplyChainRelationshipProductionPlantID,
			Product:                                  				data.Product,
			Buyer:                                    				data.Buyer,
			Seller:                                   				data.Seller,
			DeliverToParty:                           				data.DeliverToParty,
			DeliverToPlant:                           				data.DeliverToPlant,
			DeliverFromParty:                         				data.DeliverFromParty,
			DeliverFromPlant:                         				data.DeliverFromPlant,
			InspectionPlantBusinessPartner:           				data.InspectionPlantBusinessPartner,
			InspectionPlant:                          				data.InspectionPlant,
			Sequence:                                 				data.Sequence,
			SequenceText:                             				data.SequenceText,
			OperationText:                            				data.OperationText,
			OperationStatus:                          				data.OperationStatus,
			ResponsiblePlannerGroup:                  				data.ResponsiblePlannerGroup,
			OperationUnit:                            				data.OperationUnit,
			StandardLotSizeQuantity:                  				data.StandardLotSizeQuantity,
			MinimumLotSizeQuantity:                   				data.MinimumLotSizeQuantity,
			MaximumLotSizeQuantity:                   				data.MaximumLotSizeQuantity,
			PlainLongText:                            				data.PlainLongText,
			WorkCenter:                               				data.WorkCenter,
			CapacityCategoryCode:                     				data.CapacityCategoryCode,
			OperationCostingRelevancyType:            				data.OperationCostingRelevancyType,
			OperationSetupType:                       				data.OperationSetupType,
			OperationSetupGroupCategory:              				data.OperationSetupGroupCategory,
			OperationSetupGroup:                      				data.OperationSetupGroup,
			OperationReferenceQuantity:               				data.OperationReferenceQuantity,
			MaximumWaitDuration:                      				data.MaximumWaitDuration,
			StandardWaitDuration:                     				data.StandardWaitDuration,
			MinimumWaitDuration:                      				data.MinimumWaitDuration,
			WaitDurationUnit:                         				data.WaitDurationUnit,
			MaximumQueueDuration:                     				data.MaximumQueueDuration,
			StandardQueueDuration:                    				data.StandardQueueDuration,
			MinimumQueueDuration:                     				data.MinimumQueueDuration,
			QueueDurationUnit:                        				data.QueueDurationUnit,
			MaximumMoveDuration:                      				data.MaximumMoveDuration,
			StandardMoveDuration:                     				data.StandardMoveDuration,
			MinimumMoveDuration:                      				data.MinimumMoveDuration,
			MoveDurationUnit:                         				data.MoveDurationUnit,
			StandardDeliveryDuration:                 				data.StandardDeliveryDuration,
			StandardDeliveryDurationUnit:             				data.StandardDeliveryDurationUnit,
			StandardOperationScrapPercent:            				data.StandardOperationScrapPercent,
			PlannedOperationStandardValue:            				data.PlannedOperationStandardValue,
			PlannedOperationLowerValue:            	  				data.PlannedOperationLowerValue,
			PlannedOperationUpperValue:            	  				data.PlannedOperationUpperValue,
			PlannedOperationValueUnit:            	  				data.PlannedOperationValueUnit,
			CostElement:                              				data.CostElement,
			OperationErlstSchedldExecStrtDte:						data.OperationErlstSchedldExecStrtDte,
			OperationErlstSchedldExecStrtTme:						data.OperationErlstSchedldExecStrtTme,
			OperationErlstSchedldExecEndDte:						data.OperationErlstSchedldExecEndDte,
			OperationErlstSchedldExecEndTme:						data.OperationErlstSchedldExecEndTme,
			OperationActualExecutionStartDate:						data.OperationActualExecutionStartDate,
			OperationActualExecutionStartTime:						data.OperationActualExecutionStartTime,
			OperationActualExecutionEndDate:						data.OperationActualExecutionEndDate,
			OperationActualExecutionEndTime:						data.OperationActualExecutionEndTime,
			OperationConfirmedYieldQuantityInBaseUnit:				data.OperationConfirmedYieldQuantityInBaseUnit,
			OperationConfirmedYieldQuantityInProductionUnit:		data.OperationConfirmedYieldQuantityInProductionUnit,
			OperationConfirmedYieldQuantityInOperationUnit:			data.OperationConfirmedYieldQuantityInOperationUnit,
			ValidityStartDate:                        				data.ValidityStartDate,
			ValidityEndDate:                          				data.ValidityEndDate,
			CreationDate:                             				data.CreationDate,
			CreationTime:                             				data.CreationTime,
			LastChangeDate:                           				data.LastChangeDate,
			LastChangeTime:                           				data.LastChangeTime,
			IsReleased:                               				data.IsReleased,
			IsPartiallyConfirmed:                     				data.IsPartiallyConfirmed,
			IsConfirmed:                              				data.IsConfirmed,
			IsLocked:                                 				data.IsLocked,
			IsCancelled:                              				data.IsCancelled,
			IsMarkedForDeletion:                      				data.IsMarkedForDeletion,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &operation, nil
	}

	return &operation, nil
}
