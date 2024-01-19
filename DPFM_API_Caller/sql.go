package dpfm_api_caller

import (
	"context"
	dpfm_api_input_reader "data-platform-api-inspection-lot-reads-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-inspection-lot-reads-rmq-kube/DPFM_API_Output_Formatter"
	"fmt"
	"sync"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

func (c *DPFMAPICaller) readSqlProcess(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	accepter []string,
	errs *[]error,
	log *logger.Logger,
) interface{} {
	var header *[]dpfm_api_output_formatter.Header
	var specGeneral *[]dpfm_api_output_formatter.SpecGeneral
	var specDetail *[]dpfm_api_output_formatter.SpecDetail
	var componentComposition *[]dpfm_api_output_formatter.ComponentComposition
	var inspection *[]dpfm_api_output_formatter.Inspection
	var operation *[]dpfm_api_output_formatter.Operation
	for _, fn := range accepter {
		switch fn {
		case "Header":
			func() {
				header = c.Header(mtx, input, output, errs, log)
			}()
		case "SpecGeneral":
			func() {
				specGeneral = c.SpecGeneral(mtx, input, output, errs, log)
			}()
		case "SpecGenerals":
			func() {
				specGeneral = c.SpecGenerals(mtx, input, output, errs, log)
			}()
		case "SpecDetail":
			func() {
				specDetail = c.SpecDetail(mtx, input, output, errs, log)
			}()
		case "SpecDetails":
			func() {
				specDetail = c.SpecDetails(mtx, input, output, errs, log)
			}()
		case "ComponentComposition":
			func() {
				componentComposition = c.ComponentComposition(mtx, input, output, errs, log)
			}()
		case "ComponentCompositions":
			func() {
				componentComposition = c.ComponentCompositions(mtx, input, output, errs, log)
			}()
		case "Inspection":
			func() {
				inspection = c.Inspection(mtx, input, output, errs, log)
			}()
		case "Inspections":
			func() {
				inspection = c.Inspections(mtx, input, output, errs, log)
			}()
		case "Operation":
			func() {
				operation = c.Operation(mtx, input, output, errs, log)
			}()
		case "Operations":
			func() {
				operation = c.Operations(mtx, input, output, errs, log)
			}()
		default:
		}
		if len(*errs) != 0 {
			break
		}
	}

	data := &dpfm_api_output_formatter.Message{
		Header:               header,
		SpecGeneral:          specGeneral,
		SpecDetail:           specDetail,
		ComponentComposition: componentComposition,
		Inspection:           inspection,
		Operation:            operation,
	}

	return data
}

func (c *DPFMAPICaller) Header(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Header {
	where := fmt.Sprintf("WHERE header.InspectionLot = %d ", input.Header.InspectionLot)

	if input.Header.IsMarkedForDeletion != nil {
		where = fmt.Sprintf("%s\nAND header.IsMarkedForDeletion = %v", where, *input.Header.IsMarkedForDeletion)
	}

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_inspection_lot_header_data AS header
		` + where + ` ORDER BY header.IsMarkedForDeletion ASC, header.IsCancelled ASC, header.InspectionLot DESC;`,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToHeader(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) SpecGeneral(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.SpecGeneral {
	where := "WHERE 1 = 1"

	where = fmt.Sprintf("%s\nAND InspectionLot = %d", where, input.Header.InspectionLot)
	where = fmt.Sprintf("%s\nAND IsMarkedForDeletion = %v", where, *input.Header.SpecGeneral[0].IsMarkedForDeletion)

	groupBy := "\nGROUP BY InspectionLot, IsMarkedForDeletion"

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_inspection_lot_spec_general_data
		` + where + groupBy + `;`)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToSpecGeneral(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) SpecGenerals(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.SpecGeneral {
	specGeneral := &dpfm_api_input_reader.SpecGeneral{}
	if len(input.Header.SpecGeneral) > 0 {
		specGeneral = &input.Header.SpecGeneral[0]
	}
	where := "WHERE 1 = 1"

	where = fmt.Sprintf("%s\nAND specGeneral.InspectionLot = '%v'", where, specGeneral.InspectionLot)
	if specGeneral != nil {
		if specGeneral.IsCancelled != nil {
			where = fmt.Sprintf("%s\nAND specGeneral.IsCancelled = %v", where, *specGeneral.IsCancelled)
		}
		if specGeneral.IsMarkedForDeletion != nil {
			where = fmt.Sprintf("%s\nAND specGeneral.IsMarkedForDeletion = %v", where, *specGeneral.IsMarkedForDeletion)
		}
	}

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_inspection_lot_spec_general_data as specGeneral
		` + where + ` ORDER BY IsMarkedForDeletion ASC, InspectionLot ASC;`,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	data, err := dpfm_api_output_formatter.ConvertToSpecGeneral(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) SpecDetail(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.SpecDetail {
	where := "WHERE 1 = 1"

	where = fmt.Sprintf("%s\nAND InspectionLot = %d", where, input.Header.InspectionLot)
	where = fmt.Sprintf("%s\nAND SpecType = \"%s\"", where, input.Header.SpecDetail[0].SpecType)

	if input.Header.SpecDetail[0].IsMarkedForDeletion != nil {
		where = fmt.Sprintf("%s\nAND IsMarkedForDeletion = %v", where, *input.Header.SpecDetail[0].IsMarkedForDeletion)
	}
	groupBy := "\nGROUP BY InspectionLot, SpecType, IsMarkedForDeletion"

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_inspection_lot_spec_detail_data
		` + where + groupBy + `;`)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToSpecDetail(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) SpecDetails(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.SpecDetail {
	specDetail := &dpfm_api_input_reader.SpecDetail{}
	if len(input.Header.SpecDetail) > 0 {
		specDetail = &input.Header.SpecDetail[0]
	}
	where := "WHERE 1 = 1"

	where = fmt.Sprintf("%s\nAND specDetail.InspectionLot = '%v'", where, specDetail.InspectionLot)
	if specDetail != nil {
		if specDetail.IsCancelled != nil {
			where = fmt.Sprintf("%s\nAND specDetail.IsCancelled = %v", where, *specDetail.IsCancelled)
		}
		if specDetail.IsMarkedForDeletion != nil {
			where = fmt.Sprintf("%s\nAND specDetail.IsMarkedForDeletion = %v", where, *specDetail.IsMarkedForDeletion)
		}
	}

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_inspection_lot_spec_detail_data as specDetail
		` + where + ` ORDER BY IsMarkedForDeletion ASC, InspectionLot ASC;`,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	data, err := dpfm_api_output_formatter.ConvertToSpecDetail(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) ComponentComposition(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.ComponentComposition {
	where := "WHERE 1 = 1"

	where = fmt.Sprintf("%s\nAND InspectionLot = %d", where, input.Header.InspectionLot)
	where = fmt.Sprintf("%s\nAND ComponentCompositionType = \"%s\"", where, input.Header.ComponentComposition[0].ComponentCompositionType)

	if input.Header.ComponentComposition[0].IsMarkedForDeletion != nil {
		where = fmt.Sprintf("%s\nAND IsMarkedForDeletion = %v", where, input.Header.ComponentComposition[0].IsMarkedForDeletion)
	}
	groupBy := "\nGROUP BY InspectionLot, ComponentCompositionType, IsMarkedForDeletion"

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_inspection_lot_component_composition_data
		` + where + groupBy + `;`)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToComponentComposition(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) ComponentCompositions(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.ComponentComposition {
	componentComposition := &dpfm_api_input_reader.ComponentComposition{}
	if len(input.Header.ComponentComposition) > 0 {
		componentComposition = &input.Header.ComponentComposition[0]
	}
	where := "WHERE 1 = 1"

	where = fmt.Sprintf("%s\nAND componentComposition.InspectionLot = '%v'", where, componentComposition.InspectionLot)
	if componentComposition != nil {
		if componentComposition.IsCancelled != nil {
			where = fmt.Sprintf("%s\nAND componentComposition.IsCancelled = %v", where, *componentComposition.IsCancelled)
		}
		if componentComposition.IsMarkedForDeletion != nil {
			where = fmt.Sprintf("%s\nAND componentComposition.IsMarkedForDeletion = %v", where, *componentComposition.IsMarkedForDeletion)
		}
	}

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_inspection_lot_component_composition_data as componentComposition
		` + where + ` ORDER BY IsMarkedForDeletion ASC, InspectionLot ASC;`,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	data, err := dpfm_api_output_formatter.ConvertToComponentComposition(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) Inspection(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Inspection {
	where := "WHERE 1 = 1"

	where = fmt.Sprintf("%s\nAND InspectionLot = %d", where, input.Header.InspectionLot)
	where = fmt.Sprintf("%s\nAND Inspection = %d", where, input.Header.Inspection[0].Inspection)

	if input.Header.Inspection[0].IsMarkedForDeletion != nil {
		where = fmt.Sprintf("%s\nAND IsMarkedForDeletion = %v", where, *input.Header.Inspection[0].IsMarkedForDeletion)
	}

	groupBy := "\nGROUP BY InspectionLot, Inspection, IsMarkedForDeletion"

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_inspection_lot_inspection_data
		` + where + groupBy + `;`)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToInspection(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) Inspections(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Inspection {
	inspection := &dpfm_api_input_reader.Inspection{}
	if len(input.Header.Inspection) > 0 {
		inspection = &input.Header.Inspection[0]
	}
	where := "WHERE 1 = 1"

	where = fmt.Sprintf("%s\nAND inspection.InspectionLot = '%v'", where, inspection.InspectionLot)
	if inspection != nil {
		if inspection.IsCancelled != nil {
			where = fmt.Sprintf("%s\nAND inspection.IsCancelled = %v", where, *inspection.IsCancelled)
		}
		if inspection.IsMarkedForDeletion != nil {
			where = fmt.Sprintf("%s\nAND inspection.IsMarkedForDeletion = %v", where, *inspection.IsMarkedForDeletion)
		}
	}

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_inspection_lot_inspection_data as inspection
		` + where + ` ORDER BY IsMarkedForDeletion ASC, InspectionLot ASC;`,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	data, err := dpfm_api_output_formatter.ConvertToInspection(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) Operation(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Operation {
	where := "WHERE 1 = 1"

	where = fmt.Sprintf("%s\nAND InspectionLot = %d", where, input.Header.InspectionLot)
	where = fmt.Sprintf("%s\nAND Operations = %d", where, input.Header.Operation[0].Operations)
	where = fmt.Sprintf("%s\nAND OperationsItem = %d", where, input.Header.Operation[0].OperationsItem)
	where = fmt.Sprintf("%s\nAND OperationID = %d", where, input.Header.Operation[0].OperationID)

	if input.Header.Operation[0].IsMarkedForDeletion != nil {
		where = fmt.Sprintf("%s\nAND IsMarkedForDeletion = %v", where, *input.Header.Operation[0].IsMarkedForDeletion)
	}
	groupBy := "\nGROUP BY InspectionLot, Operations, OperationsItem, OperationID, IsMarkedForDeletion"

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_inspection_lot_operation_data
		` + where + groupBy + `;`)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToOperation(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) Operations(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Operation {
	operation := &dpfm_api_input_reader.Operation{}
	if len(input.Header.SpecGeneral) > 0 {
		operation = &input.Header.Operation[0]
	}
	where := "WHERE 1 = 1"

	where = fmt.Sprintf("%s\nAND operation.InspectionLot = '%v'", where, operation.InspectionLot)
	if operation != nil {
		if operation.IsCancelled != nil {
			where = fmt.Sprintf("%s\nAND operation.IsCancelled = %v", where, *operation.IsCancelled)
		}
		if operation.IsMarkedForDeletion != nil {
			where = fmt.Sprintf("%s\nAND operation.IsMarkedForDeletion = %v", where, *operation.IsMarkedForDeletion)
		}
	}

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_inspection_lot_operation_data as operation
		` + where + ` ORDER BY IsMarkedForDeletion ASC, InspectionLot ASC;`,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	data, err := dpfm_api_output_formatter.ConvertToOperation(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}
