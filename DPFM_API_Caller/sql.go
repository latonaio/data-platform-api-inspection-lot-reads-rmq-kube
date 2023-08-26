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
	for _, fn := range accepter {
		switch fn {
		case "ComponentComposition":
			func() {
				componentComposition = c.ComponentComposition(mtx, input, output, errs, log)	
		case "Header":
			func() {
				header = c.Header(mtx, input, output, errs, log)
			}()
		case "HeaderDoc":
			func() {		
				headerDoc = c.HeaderDoc(mtx, input, output, errs, log)
		case "Inspection":
			func() 
				inspection = c.Inspection(mtx, input, output, errs, log)
		case "Operation":
			func() {
				operation = c.Operation(mtx, input, output, errs, log)
		case "SpecDetail":
			func() {
				specDetail = c.SpecDetail(mtx, input, output, errs, log)
		case "SpecGeneral":
			func() {
				specGeneral = c.SpecGeneral(mtx, input, output, errs, log)
			}()
		default:
		}
		if len(*errs) != 0 {
			break
		}
	}

	data := &dpfm_api_output_formatter.Message{
		ComponentComposition: componentComposition,
		Header: header,
		HeaderDoc:            headerDoc,
		Inspection:           inspection,
		Operation:            operation,
		SpecDetail:           specDetail,
		SpecGeneral:          specGeneral,
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
	where := fmt.Sprintf("WHERE componentComposition.InspectionLot = %d ", input.ComponentComposition.InspectionLot)

	if input.ComponentComposition.IsMarkedForDeletion != nil {
		where = fmt.Sprintf("%s\nAND componentComposition.IsMarkedForDeletion = %v", where, *input.ComponentCompositio.IsMarkedForDeletion)
	}

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_inspection_lot_componentComposition_data AS componentComposition
		` + where + ` ORDER BY componentComposition.IsMarkedForDeletion ASC, componentComposition.IsCancelled ASC, componentComposition.InspectionLot DESC;`,
	)
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
func (c *DPFMAPICaller) HeaderDoc(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.HeaderDoc {
	where := fmt.Sprintf("WHERE headeDocr.InspectionLot = %d ", input.HeaderDoc.InspectionLot)

	if input.Header.IsMarkedForDeletion != nil {
		where = fmt.Sprintf("%s\nAND headerDoc.IsMarkedForDeletion = %v", where, *input.HeaderDoc.IsMarkedForDeletion)
	}

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_inspection_lot_headerDoc_data AS headerDoc
		` + where + ` ORDER BY headerDoc.IsMarkedForDeletion ASC, headerDoc.IsCancelled ASC, headerDoc.InspectionLot DESC;`,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToHeaderDoc(rows)
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
	where := fmt.Sprintf("WHERE inspection.InspectionLot = %d ", input.Inspection.InspectionLot)

	if input.Inspection.IsMarkedForDeletion != nil {
		where = fmt.Sprintf("%s\nAND inspection.IsMarkedForDeletion = %v", where, *input.Inspection.IsMarkedForDeletion)
	}

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_inspection_lot_inspection_data AS inspection
		` + where + ` ORDER BY inspection.IsMarkedForDeletion ASC, inspection.IsCancelled ASC, inspection.InspectionLot DESC;`,
	)
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
func (c *DPFMAPICaller) Operation(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Operation {
	where := fmt.Sprintf("WHERE operation.InspectionLot = %d ", input.Operation.InspectionLot)

	if input.Operation.IsMarkedForDeletion != nil {
		where = fmt.Sprintf("%s\nAND operation.IsMarkedForDeletion = %v", where, *input.Operation.IsMarkedForDeletion)
	}

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_inspection_lot_operation_data AS operation
		` + where + ` ORDER BY operation.IsMarkedForDeletion ASC, operation.IsCancelled ASC, operation.InspectionLot DESC;`,
	)
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
func (c *DPFMAPICaller) SpecDetail(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.SpecDetail {
	where := fmt.Sprintf("WHERE specDetail.InspectionLot = %d ", input.SpecDetail.InspectionLot)

	if input.SpecDetail.IsMarkedForDeletion != nil {
		where = fmt.Sprintf("%s\nAND specDetail.IsMarkedForDeletion = %v", where, *input.SpecDetail.IsMarkedForDeletion)
	}

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_inspection_lot_specDetail_data AS specDetail
		` + where + ` ORDER BY specDetail.IsMarkedForDeletion ASC, specDetail.IsCancelled ASC, specDetail.InspectionLot DESC;`,
	)
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
func (c *DPFMAPICaller) SpecGeneral(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.SpecGeneral {
	where := fmt.Sprintf("WHERE specGeneral.InspectionLot = %d ", input.SpecGeneral.InspectionLot)

	if input.SpecGeneral.IsMarkedForDeletion != nil {
		where = fmt.Sprintf("%s\nAND specGeneral.IsMarkedForDeletion = %v", where, *input.SpecGeneral.IsMarkedForDeletion)
	}

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_inspection_lot_specGeneral_data AS specGeneral
		` + where + ` ORDER BY specGeneral.IsMarkedForDeletion ASC, specGeneral.IsCancelled ASC, specGeneral.InspectionLot DESC;`,
	)
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
