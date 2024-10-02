/*****************************************************************************************************************/

//	@author		Michael Roberts <michael@observerly.com>
//	@package	@observerly/birpc
//	@license	Copyright © 2021-2024 observerly

/*****************************************************************************************************************/

package storage

/*****************************************************************************************************************/

import (
	"context"
	"fmt"
	"time"

	pb "birpc/internal/gen/store/v1"

	"connectrpc.com/connect"
	"github.com/rs/zerolog/log"
)

/*****************************************************************************************************************/

func (s *server) GetFITSAsTIFFHandler(ctx context.Context, req *connect.Request[pb.GetFITSAsGenericHandlerRequest]) (*connect.Response[pb.GetFITSAsGenericHandlerResponse], error) {
	now := time.Now()

	logger := log.With().Str("owner", req.Msg.Owner).Str("bucket", req.Msg.BucketName).Str("location", req.Msg.Location).Str("rfc3339", now.Format(time.RFC3339)).Logger()

	url := fmt.Sprintf("https://%s/%s", req.Msg.BucketName, req.Msg.Location)

	width := int32(0)

	height := int32(0)

	logger.Info().Msg("Returning TIFF Download URL")

	return connect.NewResponse(&pb.GetFITSAsGenericHandlerResponse{
		DownloadUrl: url,
		Height:      height,
		Width:       width,
	}), nil
}

/*****************************************************************************************************************/
