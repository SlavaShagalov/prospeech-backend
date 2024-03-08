package http

import (
	"bytes"
	pAudios "github.com/SlavaShagalov/prospeech-backend/internal/audios"
	pFiles "github.com/SlavaShagalov/prospeech-backend/internal/files"
	mw "github.com/SlavaShagalov/prospeech-backend/internal/middleware"
	"github.com/SlavaShagalov/prospeech-backend/internal/pkg/constants"
	pErrors "github.com/SlavaShagalov/prospeech-backend/internal/pkg/errors"
	pHTTP "github.com/SlavaShagalov/prospeech-backend/internal/pkg/http"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"io"
	"net/http"
)

const (
	fileFormKey = "file"
)

type delivery struct {
	uc  pAudios.Usecase
	log *zap.Logger
}

func RegisterHandlers(mux *mux.Router, uc pAudios.Usecase, log *zap.Logger, checkAuth mw.Middleware) {
	del := delivery{
		uc:  uc,
		log: log,
	}

	const (
		audiosPrefix = "/audios"
		audiosPath   = constants.ApiPrefix + audiosPrefix
		audioPath    = audiosPath + "/{id}"
	)

	mux.HandleFunc(audiosPath, checkAuth(del.create)).Methods(http.MethodPost)
	mux.HandleFunc(audiosPath, checkAuth(del.list)).Methods(http.MethodGet)

	mux.HandleFunc(audioPath, checkAuth(del.get)).Methods(http.MethodGet)
	mux.HandleFunc(audioPath, checkAuth(del.partialUpdate)).Methods(http.MethodPatch)
	mux.HandleFunc(audioPath, checkAuth(del.delete)).Methods(http.MethodDelete)
}

func (del *delivery) create(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(mw.ContextUserID).(int64)
	if !ok {
		pHTTP.HandleError(w, r, pErrors.ErrReadBody)
		return
	}

	file, header, err := r.FormFile(fileFormKey)
	if err != nil {
		pHTTP.HandleError(w, r, err)
		return
	}
	buf := bytes.NewBuffer(nil)
	_, err = io.Copy(buf, file)
	if err != nil {
		pHTTP.HandleError(w, r, err)
		return
	}

	params := pAudios.CreateParams{
		UserID: userID,
		File: pFiles.File{
			Name: header.Filename,
			Data: buf.Bytes(),
		},
	}

	audio, err := del.uc.Create(r.Context(), &params)
	if err != nil {
		pHTTP.HandleError(w, r, err)
		return
	}

	response := newCreateResponse(audio)
	pHTTP.SendJSON(w, r, http.StatusOK, response)
}

func (del *delivery) list(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(mw.ContextUserID).(int64)
	if !ok {
		pHTTP.HandleError(w, r, pErrors.ErrReadBody)
		return
	}

	audios, err := del.uc.List(r.Context(), userID)
	if err != nil {
		pHTTP.HandleError(w, r, err)
		return
	}

	response := newListResponse(audios)
	pHTTP.SendJSON(w, r, http.StatusOK, response)
}

func (del *delivery) get(w http.ResponseWriter, r *http.Request) {
	//vars := mux.Vars(r)
	//audioID, err := strconv.Atoi(vars["id"])
	//if err != nil {
	//	pHTTP.HandleError(w, r, err)
	//	return
	//}
	//
	//audio, err := del.uc.Get(ctx, audioID)
	//if err != nil {
	//	pHTTP.HandleError(w, r, err)
	//	return
	//}
	//
	//response := newGetResponse(&audio)
	//pHTTP.SendJSON(w, r, http.StatusOK, response)
}

func (del *delivery) partialUpdate(w http.ResponseWriter, r *http.Request) {
	//
	//vars := mux.Vars(r)
	//audioID, err := strconv.Atoi(vars["id"])
	//if err != nil {
	//	pHTTP.HandleError(w, r, err)
	//	return
	//}
	//
	//body, err := pHTTP.ReadBody(r, del.log)
	//if err != nil {
	//	pHTTP.HandleError(w, r, err)
	//	return
	//}
	//
	//var request partialUpdateRequest
	//err = request.UnmarshalJSON(body)
	//if err != nil {
	//	pHTTP.HandleError(w, r, pErrors.ErrReadBody)
	//	return
	//}
	//
	//params := pAudios.PartialUpdateParams{ID: audioID}
	//params.UpdateTitle = request.Title != nil
	//if params.UpdateTitle {
	//	params.Title = *request.Title
	//}
	//params.UpdateDescription = request.Description != nil
	//if params.UpdateDescription {
	//	params.Description = *request.Description
	//}
	//
	//audio, err := del.uc.PartialUpdate(ctx, &params)
	//if err != nil {
	//	pHTTP.HandleError(w, r, err)
	//	return
	//}
	//
	//response := newGetResponse(&audio)
	//pHTTP.SendJSON(w, r, http.StatusOK, response)
}

func (del *delivery) delete(w http.ResponseWriter, r *http.Request) {
	//
	//vars := mux.Vars(r)
	//audioID, err := strconv.Atoi(vars["id"])
	//if err != nil {
	//	pHTTP.HandleError(w, r, err)
	//	return
	//}
	//
	//err = del.uc.Delete(ctx, audioID)
	//if err != nil {
	//	pHTTP.HandleError(w, r, err)
	//	return
	//}
	//
	//w.WriteHeader(http.StatusNoContent)
}
