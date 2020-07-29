// WARNING: This file has automatically been generated
// Code generated by https://git.io/c-for-go. DO NOT EDIT.

package generated

/*
#cgo LDFLAGS: -L${SRCDIR}/.. -lfilcrypto
#cgo pkg-config: ${SRCDIR}/../filcrypto.pc
#include "../filcrypto.h"
#include <stdlib.h>
#include "cgo_helpers.h"
*/
import "C"
import (
	"runtime"
	"unsafe"
)

// FilAggregate function as declared in filecoin-ffi/filcrypto.h:286
func FilAggregate(flattenedSignaturesPtr string, flattenedSignaturesLen uint) *FilAggregateResponse {
	flattenedSignaturesPtr = safeString(flattenedSignaturesPtr)
	cflattenedSignaturesPtr, _ := unpackPUint8TString(flattenedSignaturesPtr)
	cflattenedSignaturesLen, _ := (C.size_t)(flattenedSignaturesLen), cgoAllocsUnknown
	__ret := C.fil_aggregate(cflattenedSignaturesPtr, cflattenedSignaturesLen)
	runtime.KeepAlive(flattenedSignaturesPtr)
	__v := NewFilAggregateResponseRef(unsafe.Pointer(__ret))
	return __v
}

// FilClearCache function as declared in filecoin-ffi/filcrypto.h:289
func FilClearCache(sectorSize uint64, cacheDirPath string) *FilClearCacheResponse {
	csectorSize, _ := (C.uint64_t)(sectorSize), cgoAllocsUnknown
	cacheDirPath = safeString(cacheDirPath)
	ccacheDirPath, _ := unpackPCharString(cacheDirPath)
	__ret := C.fil_clear_cache(csectorSize, ccacheDirPath)
	runtime.KeepAlive(cacheDirPath)
	__v := NewFilClearCacheResponseRef(unsafe.Pointer(__ret))
	return __v
}

// FilDestroyAggregateResponse function as declared in filecoin-ffi/filcrypto.h:291
func FilDestroyAggregateResponse(ptr *FilAggregateResponse) {
	cptr, _ := ptr.PassRef()
	C.fil_destroy_aggregate_response(cptr)
}

// FilDestroyClearCacheResponse function as declared in filecoin-ffi/filcrypto.h:293
func FilDestroyClearCacheResponse(ptr *FilClearCacheResponse) {
	cptr, _ := ptr.PassRef()
	C.fil_destroy_clear_cache_response(cptr)
}

// FilDestroyFauxrepResponse function as declared in filecoin-ffi/filcrypto.h:295
func FilDestroyFauxrepResponse(ptr *FilFauxRepResponse) {
	cptr, _ := ptr.PassRef()
	C.fil_destroy_fauxrep_response(cptr)
}

// FilDestroyFinalizeTicketResponse function as declared in filecoin-ffi/filcrypto.h:297
func FilDestroyFinalizeTicketResponse(ptr *FilFinalizeTicketResponse) {
	cptr, _ := ptr.PassRef()
	C.fil_destroy_finalize_ticket_response(cptr)
}

// FilDestroyGenerateDataCommitmentResponse function as declared in filecoin-ffi/filcrypto.h:299
func FilDestroyGenerateDataCommitmentResponse(ptr *FilGenerateDataCommitmentResponse) {
	cptr, _ := ptr.PassRef()
	C.fil_destroy_generate_data_commitment_response(cptr)
}

// FilDestroyGeneratePieceCommitmentResponse function as declared in filecoin-ffi/filcrypto.h:301
func FilDestroyGeneratePieceCommitmentResponse(ptr *FilGeneratePieceCommitmentResponse) {
	cptr, _ := ptr.PassRef()
	C.fil_destroy_generate_piece_commitment_response(cptr)
}

// FilDestroyGenerateWindowPostResponse function as declared in filecoin-ffi/filcrypto.h:303
func FilDestroyGenerateWindowPostResponse(ptr *FilGenerateWindowPoStResponse) {
	cptr, _ := ptr.PassRef()
	C.fil_destroy_generate_window_post_response(cptr)
}

// FilDestroyGenerateWinningPostResponse function as declared in filecoin-ffi/filcrypto.h:305
func FilDestroyGenerateWinningPostResponse(ptr *FilGenerateWinningPoStResponse) {
	cptr, _ := ptr.PassRef()
	C.fil_destroy_generate_winning_post_response(cptr)
}

// FilDestroyGenerateWinningPostSectorChallenge function as declared in filecoin-ffi/filcrypto.h:307
func FilDestroyGenerateWinningPostSectorChallenge(ptr *FilGenerateWinningPoStSectorChallenge) {
	cptr, _ := ptr.PassRef()
	C.fil_destroy_generate_winning_post_sector_challenge(cptr)
}

// FilDestroyGpuDeviceResponse function as declared in filecoin-ffi/filcrypto.h:309
func FilDestroyGpuDeviceResponse(ptr *FilGpuDeviceResponse) {
	cptr, _ := ptr.PassRef()
	C.fil_destroy_gpu_device_response(cptr)
}

// FilDestroyHashResponse function as declared in filecoin-ffi/filcrypto.h:311
func FilDestroyHashResponse(ptr *FilHashResponse) {
	cptr, _ := ptr.PassRef()
	C.fil_destroy_hash_response(cptr)
}

// FilDestroyInitLogFdResponse function as declared in filecoin-ffi/filcrypto.h:313
func FilDestroyInitLogFdResponse(ptr *FilInitLogFdResponse) {
	cptr, _ := ptr.PassRef()
	C.fil_destroy_init_log_fd_response(cptr)
}

// FilDestroyPrivateKeyGenerateResponse function as declared in filecoin-ffi/filcrypto.h:315
func FilDestroyPrivateKeyGenerateResponse(ptr *FilPrivateKeyGenerateResponse) {
	cptr, _ := ptr.PassRef()
	C.fil_destroy_private_key_generate_response(cptr)
}

// FilDestroyPrivateKeyPublicKeyResponse function as declared in filecoin-ffi/filcrypto.h:317
func FilDestroyPrivateKeyPublicKeyResponse(ptr *FilPrivateKeyPublicKeyResponse) {
	cptr, _ := ptr.PassRef()
	C.fil_destroy_private_key_public_key_response(cptr)
}

// FilDestroyPrivateKeySignResponse function as declared in filecoin-ffi/filcrypto.h:319
func FilDestroyPrivateKeySignResponse(ptr *FilPrivateKeySignResponse) {
	cptr, _ := ptr.PassRef()
	C.fil_destroy_private_key_sign_response(cptr)
}

// FilDestroySealCommitPhase1Response function as declared in filecoin-ffi/filcrypto.h:321
func FilDestroySealCommitPhase1Response(ptr *FilSealCommitPhase1Response) {
	cptr, _ := ptr.PassRef()
	C.fil_destroy_seal_commit_phase1_response(cptr)
}

// FilDestroySealCommitPhase2Response function as declared in filecoin-ffi/filcrypto.h:323
func FilDestroySealCommitPhase2Response(ptr *FilSealCommitPhase2Response) {
	cptr, _ := ptr.PassRef()
	C.fil_destroy_seal_commit_phase2_response(cptr)
}

// FilDestroySealPreCommitPhase1Response function as declared in filecoin-ffi/filcrypto.h:325
func FilDestroySealPreCommitPhase1Response(ptr *FilSealPreCommitPhase1Response) {
	cptr, _ := ptr.PassRef()
	C.fil_destroy_seal_pre_commit_phase1_response(cptr)
}

// FilDestroySealPreCommitPhase2Response function as declared in filecoin-ffi/filcrypto.h:327
func FilDestroySealPreCommitPhase2Response(ptr *FilSealPreCommitPhase2Response) {
	cptr, _ := ptr.PassRef()
	C.fil_destroy_seal_pre_commit_phase2_response(cptr)
}

// FilDestroyStringResponse function as declared in filecoin-ffi/filcrypto.h:329
func FilDestroyStringResponse(ptr *FilStringResponse) {
	cptr, _ := ptr.PassRef()
	C.fil_destroy_string_response(cptr)
}

// FilDestroyUnsealRangeResponse function as declared in filecoin-ffi/filcrypto.h:331
func FilDestroyUnsealRangeResponse(ptr *FilUnsealRangeResponse) {
	cptr, _ := ptr.PassRef()
	C.fil_destroy_unseal_range_response(cptr)
}

// FilDestroyVerifySealResponse function as declared in filecoin-ffi/filcrypto.h:337
func FilDestroyVerifySealResponse(ptr *FilVerifySealResponse) {
	cptr, _ := ptr.PassRef()
	C.fil_destroy_verify_seal_response(cptr)
}

// FilDestroyVerifyWindowPostResponse function as declared in filecoin-ffi/filcrypto.h:339
func FilDestroyVerifyWindowPostResponse(ptr *FilVerifyWindowPoStResponse) {
	cptr, _ := ptr.PassRef()
	C.fil_destroy_verify_window_post_response(cptr)
}

// FilDestroyVerifyWinningPostResponse function as declared in filecoin-ffi/filcrypto.h:345
func FilDestroyVerifyWinningPostResponse(ptr *FilVerifyWinningPoStResponse) {
	cptr, _ := ptr.PassRef()
	C.fil_destroy_verify_winning_post_response(cptr)
}

// FilDestroyWriteWithAlignmentResponse function as declared in filecoin-ffi/filcrypto.h:347
func FilDestroyWriteWithAlignmentResponse(ptr *FilWriteWithAlignmentResponse) {
	cptr, _ := ptr.PassRef()
	C.fil_destroy_write_with_alignment_response(cptr)
}

// FilDestroyWriteWithoutAlignmentResponse function as declared in filecoin-ffi/filcrypto.h:349
func FilDestroyWriteWithoutAlignmentResponse(ptr *FilWriteWithoutAlignmentResponse) {
	cptr, _ := ptr.PassRef()
	C.fil_destroy_write_without_alignment_response(cptr)
}

// FilFauxrep function as declared in filecoin-ffi/filcrypto.h:351
func FilFauxrep(registeredProof FilRegisteredSealProof, cacheDirPath string, sealedSectorPath string) *FilFauxRepResponse {
	cregisteredProof, _ := (C.fil_RegisteredSealProof)(registeredProof), cgoAllocsUnknown
	cacheDirPath = safeString(cacheDirPath)
	ccacheDirPath, _ := unpackPCharString(cacheDirPath)
	sealedSectorPath = safeString(sealedSectorPath)
	csealedSectorPath, _ := unpackPCharString(sealedSectorPath)
	__ret := C.fil_fauxrep(cregisteredProof, ccacheDirPath, csealedSectorPath)
	runtime.KeepAlive(sealedSectorPath)
	runtime.KeepAlive(cacheDirPath)
	__v := NewFilFauxRepResponseRef(unsafe.Pointer(__ret))
	return __v
}

// FilFauxrep2 function as declared in filecoin-ffi/filcrypto.h:355
func FilFauxrep2(registeredProof FilRegisteredSealProof, cacheDirPath string, existingPAuxPath string) *FilFauxRepResponse {
	cregisteredProof, _ := (C.fil_RegisteredSealProof)(registeredProof), cgoAllocsUnknown
	cacheDirPath = safeString(cacheDirPath)
	ccacheDirPath, _ := unpackPCharString(cacheDirPath)
	existingPAuxPath = safeString(existingPAuxPath)
	cexistingPAuxPath, _ := unpackPCharString(existingPAuxPath)
	__ret := C.fil_fauxrep2(cregisteredProof, ccacheDirPath, cexistingPAuxPath)
	runtime.KeepAlive(existingPAuxPath)
	runtime.KeepAlive(cacheDirPath)
	__v := NewFilFauxRepResponseRef(unsafe.Pointer(__ret))
	return __v
}

// FilGenerateDataCommitment function as declared in filecoin-ffi/filcrypto.h:362
func FilGenerateDataCommitment(registeredProof FilRegisteredSealProof, piecesPtr []FilPublicPieceInfo, piecesLen uint) *FilGenerateDataCommitmentResponse {
	cregisteredProof, _ := (C.fil_RegisteredSealProof)(registeredProof), cgoAllocsUnknown
	cpiecesPtr, _ := unpackArgSFilPublicPieceInfo(piecesPtr)
	cpiecesLen, _ := (C.size_t)(piecesLen), cgoAllocsUnknown
	__ret := C.fil_generate_data_commitment(cregisteredProof, cpiecesPtr, cpiecesLen)
	packSFilPublicPieceInfo(piecesPtr, cpiecesPtr)
	__v := NewFilGenerateDataCommitmentResponseRef(unsafe.Pointer(__ret))
	return __v
}

// FilGeneratePieceCommitment function as declared in filecoin-ffi/filcrypto.h:370
func FilGeneratePieceCommitment(registeredProof FilRegisteredSealProof, pieceFdRaw int32, unpaddedPieceSize uint64) *FilGeneratePieceCommitmentResponse {
	cregisteredProof, _ := (C.fil_RegisteredSealProof)(registeredProof), cgoAllocsUnknown
	cpieceFdRaw, _ := (C.int)(pieceFdRaw), cgoAllocsUnknown
	cunpaddedPieceSize, _ := (C.uint64_t)(unpaddedPieceSize), cgoAllocsUnknown
	__ret := C.fil_generate_piece_commitment(cregisteredProof, cpieceFdRaw, cunpaddedPieceSize)
	__v := NewFilGeneratePieceCommitmentResponseRef(unsafe.Pointer(__ret))
	return __v
}

// FilGenerateWindowPost function as declared in filecoin-ffi/filcrypto.h:378
func FilGenerateWindowPost(randomness Fil32ByteArray, replicasPtr []FilPrivateReplicaInfo, replicasLen uint, proverId Fil32ByteArray) *FilGenerateWindowPoStResponse {
	crandomness, _ := randomness.PassValue()
	creplicasPtr, _ := unpackArgSFilPrivateReplicaInfo(replicasPtr)
	creplicasLen, _ := (C.size_t)(replicasLen), cgoAllocsUnknown
	cproverId, _ := proverId.PassValue()
	__ret := C.fil_generate_window_post(crandomness, creplicasPtr, creplicasLen, cproverId)
	packSFilPrivateReplicaInfo(replicasPtr, creplicasPtr)
	__v := NewFilGenerateWindowPoStResponseRef(unsafe.Pointer(__ret))
	return __v
}

// FilGenerateWinningPost function as declared in filecoin-ffi/filcrypto.h:387
func FilGenerateWinningPost(randomness Fil32ByteArray, replicasPtr []FilPrivateReplicaInfo, replicasLen uint, proverId Fil32ByteArray) *FilGenerateWinningPoStResponse {
	crandomness, _ := randomness.PassValue()
	creplicasPtr, _ := unpackArgSFilPrivateReplicaInfo(replicasPtr)
	creplicasLen, _ := (C.size_t)(replicasLen), cgoAllocsUnknown
	cproverId, _ := proverId.PassValue()
	__ret := C.fil_generate_winning_post(crandomness, creplicasPtr, creplicasLen, cproverId)
	packSFilPrivateReplicaInfo(replicasPtr, creplicasPtr)
	__v := NewFilGenerateWinningPoStResponseRef(unsafe.Pointer(__ret))
	return __v
}

// FilGenerateWinningPostSectorChallenge function as declared in filecoin-ffi/filcrypto.h:396
func FilGenerateWinningPostSectorChallenge(registeredProof FilRegisteredPoStProof, randomness Fil32ByteArray, sectorSetLen uint64, proverId Fil32ByteArray) *FilGenerateWinningPoStSectorChallenge {
	cregisteredProof, _ := (C.fil_RegisteredPoStProof)(registeredProof), cgoAllocsUnknown
	crandomness, _ := randomness.PassValue()
	csectorSetLen, _ := (C.uint64_t)(sectorSetLen), cgoAllocsUnknown
	cproverId, _ := proverId.PassValue()
	__ret := C.fil_generate_winning_post_sector_challenge(cregisteredProof, crandomness, csectorSetLen, cproverId)
	__v := NewFilGenerateWinningPoStSectorChallengeRef(unsafe.Pointer(__ret))
	return __v
}

// FilGetGpuDevices function as declared in filecoin-ffi/filcrypto.h:404
func FilGetGpuDevices() *FilGpuDeviceResponse {
	__ret := C.fil_get_gpu_devices()
	__v := NewFilGpuDeviceResponseRef(unsafe.Pointer(__ret))
	return __v
}

// FilGetMaxUserBytesPerStagedSector function as declared in filecoin-ffi/filcrypto.h:410
func FilGetMaxUserBytesPerStagedSector(registeredProof FilRegisteredSealProof) uint64 {
	cregisteredProof, _ := (C.fil_RegisteredSealProof)(registeredProof), cgoAllocsUnknown
	__ret := C.fil_get_max_user_bytes_per_staged_sector(cregisteredProof)
	__v := (uint64)(__ret)
	return __v
}

// FilGetPostCircuitIdentifier function as declared in filecoin-ffi/filcrypto.h:416
func FilGetPostCircuitIdentifier(registeredProof FilRegisteredPoStProof) *FilStringResponse {
	cregisteredProof, _ := (C.fil_RegisteredPoStProof)(registeredProof), cgoAllocsUnknown
	__ret := C.fil_get_post_circuit_identifier(cregisteredProof)
	__v := NewFilStringResponseRef(unsafe.Pointer(__ret))
	return __v
}

// FilGetPostParamsCid function as declared in filecoin-ffi/filcrypto.h:422
func FilGetPostParamsCid(registeredProof FilRegisteredPoStProof) *FilStringResponse {
	cregisteredProof, _ := (C.fil_RegisteredPoStProof)(registeredProof), cgoAllocsUnknown
	__ret := C.fil_get_post_params_cid(cregisteredProof)
	__v := NewFilStringResponseRef(unsafe.Pointer(__ret))
	return __v
}

// FilGetPostParamsPath function as declared in filecoin-ffi/filcrypto.h:429
func FilGetPostParamsPath(registeredProof FilRegisteredPoStProof) *FilStringResponse {
	cregisteredProof, _ := (C.fil_RegisteredPoStProof)(registeredProof), cgoAllocsUnknown
	__ret := C.fil_get_post_params_path(cregisteredProof)
	__v := NewFilStringResponseRef(unsafe.Pointer(__ret))
	return __v
}

// FilGetPostVerifyingKeyCid function as declared in filecoin-ffi/filcrypto.h:435
func FilGetPostVerifyingKeyCid(registeredProof FilRegisteredPoStProof) *FilStringResponse {
	cregisteredProof, _ := (C.fil_RegisteredPoStProof)(registeredProof), cgoAllocsUnknown
	__ret := C.fil_get_post_verifying_key_cid(cregisteredProof)
	__v := NewFilStringResponseRef(unsafe.Pointer(__ret))
	return __v
}

// FilGetPostVerifyingKeyPath function as declared in filecoin-ffi/filcrypto.h:442
func FilGetPostVerifyingKeyPath(registeredProof FilRegisteredPoStProof) *FilStringResponse {
	cregisteredProof, _ := (C.fil_RegisteredPoStProof)(registeredProof), cgoAllocsUnknown
	__ret := C.fil_get_post_verifying_key_path(cregisteredProof)
	__v := NewFilStringResponseRef(unsafe.Pointer(__ret))
	return __v
}

// FilGetPostVersion function as declared in filecoin-ffi/filcrypto.h:448
func FilGetPostVersion(registeredProof FilRegisteredPoStProof) *FilStringResponse {
	cregisteredProof, _ := (C.fil_RegisteredPoStProof)(registeredProof), cgoAllocsUnknown
	__ret := C.fil_get_post_version(cregisteredProof)
	__v := NewFilStringResponseRef(unsafe.Pointer(__ret))
	return __v
}

// FilGetSealCircuitIdentifier function as declared in filecoin-ffi/filcrypto.h:454
func FilGetSealCircuitIdentifier(registeredProof FilRegisteredSealProof) *FilStringResponse {
	cregisteredProof, _ := (C.fil_RegisteredSealProof)(registeredProof), cgoAllocsUnknown
	__ret := C.fil_get_seal_circuit_identifier(cregisteredProof)
	__v := NewFilStringResponseRef(unsafe.Pointer(__ret))
	return __v
}

// FilGetSealParamsCid function as declared in filecoin-ffi/filcrypto.h:460
func FilGetSealParamsCid(registeredProof FilRegisteredSealProof) *FilStringResponse {
	cregisteredProof, _ := (C.fil_RegisteredSealProof)(registeredProof), cgoAllocsUnknown
	__ret := C.fil_get_seal_params_cid(cregisteredProof)
	__v := NewFilStringResponseRef(unsafe.Pointer(__ret))
	return __v
}

// FilGetSealParamsPath function as declared in filecoin-ffi/filcrypto.h:467
func FilGetSealParamsPath(registeredProof FilRegisteredSealProof) *FilStringResponse {
	cregisteredProof, _ := (C.fil_RegisteredSealProof)(registeredProof), cgoAllocsUnknown
	__ret := C.fil_get_seal_params_path(cregisteredProof)
	__v := NewFilStringResponseRef(unsafe.Pointer(__ret))
	return __v
}

// FilGetSealVerifyingKeyCid function as declared in filecoin-ffi/filcrypto.h:473
func FilGetSealVerifyingKeyCid(registeredProof FilRegisteredSealProof) *FilStringResponse {
	cregisteredProof, _ := (C.fil_RegisteredSealProof)(registeredProof), cgoAllocsUnknown
	__ret := C.fil_get_seal_verifying_key_cid(cregisteredProof)
	__v := NewFilStringResponseRef(unsafe.Pointer(__ret))
	return __v
}

// FilGetSealVerifyingKeyPath function as declared in filecoin-ffi/filcrypto.h:480
func FilGetSealVerifyingKeyPath(registeredProof FilRegisteredSealProof) *FilStringResponse {
	cregisteredProof, _ := (C.fil_RegisteredSealProof)(registeredProof), cgoAllocsUnknown
	__ret := C.fil_get_seal_verifying_key_path(cregisteredProof)
	__v := NewFilStringResponseRef(unsafe.Pointer(__ret))
	return __v
}

// FilGetSealVersion function as declared in filecoin-ffi/filcrypto.h:486
func FilGetSealVersion(registeredProof FilRegisteredSealProof) *FilStringResponse {
	cregisteredProof, _ := (C.fil_RegisteredSealProof)(registeredProof), cgoAllocsUnknown
	__ret := C.fil_get_seal_version(cregisteredProof)
	__v := NewFilStringResponseRef(unsafe.Pointer(__ret))
	return __v
}

// FilHash function as declared in filecoin-ffi/filcrypto.h:496
func FilHash(messagePtr string, messageLen uint) *FilHashResponse {
	messagePtr = safeString(messagePtr)
	cmessagePtr, _ := unpackPUint8TString(messagePtr)
	cmessageLen, _ := (C.size_t)(messageLen), cgoAllocsUnknown
	__ret := C.fil_hash(cmessagePtr, cmessageLen)
	runtime.KeepAlive(messagePtr)
	__v := NewFilHashResponseRef(unsafe.Pointer(__ret))
	return __v
}

// FilHashVerify function as declared in filecoin-ffi/filcrypto.h:510
func FilHashVerify(signaturePtr string, flattenedMessagesPtr string, flattenedMessagesLen uint, messageSizesPtr []uint, messageSizesLen uint, flattenedPublicKeysPtr string, flattenedPublicKeysLen uint) int32 {
	signaturePtr = safeString(signaturePtr)
	csignaturePtr, _ := unpackPUint8TString(signaturePtr)
	flattenedMessagesPtr = safeString(flattenedMessagesPtr)
	cflattenedMessagesPtr, _ := unpackPUint8TString(flattenedMessagesPtr)
	cflattenedMessagesLen, _ := (C.size_t)(flattenedMessagesLen), cgoAllocsUnknown
	cmessageSizesPtr, _ := (*C.size_t)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&messageSizesPtr)).Data)), cgoAllocsUnknown
	cmessageSizesLen, _ := (C.size_t)(messageSizesLen), cgoAllocsUnknown
	flattenedPublicKeysPtr = safeString(flattenedPublicKeysPtr)
	cflattenedPublicKeysPtr, _ := unpackPUint8TString(flattenedPublicKeysPtr)
	cflattenedPublicKeysLen, _ := (C.size_t)(flattenedPublicKeysLen), cgoAllocsUnknown
	__ret := C.fil_hash_verify(csignaturePtr, cflattenedMessagesPtr, cflattenedMessagesLen, cmessageSizesPtr, cmessageSizesLen, cflattenedPublicKeysPtr, cflattenedPublicKeysLen)
	runtime.KeepAlive(flattenedPublicKeysPtr)
	runtime.KeepAlive(flattenedMessagesPtr)
	runtime.KeepAlive(signaturePtr)
	__v := (int32)(__ret)
	return __v
}

// FilInitLogFd function as declared in filecoin-ffi/filcrypto.h:527
func FilInitLogFd(logFd int32) *FilInitLogFdResponse {
	clogFd, _ := (C.int)(logFd), cgoAllocsUnknown
	__ret := C.fil_init_log_fd(clogFd)
	__v := NewFilInitLogFdResponseRef(unsafe.Pointer(__ret))
	return __v
}

// FilPrivateKeyGenerate function as declared in filecoin-ffi/filcrypto.h:532
func FilPrivateKeyGenerate() *FilPrivateKeyGenerateResponse {
	__ret := C.fil_private_key_generate()
	__v := NewFilPrivateKeyGenerateResponseRef(unsafe.Pointer(__ret))
	return __v
}

// FilPrivateKeyGenerateWithSeed function as declared in filecoin-ffi/filcrypto.h:545
func FilPrivateKeyGenerateWithSeed(rawSeed Fil32ByteArray) *FilPrivateKeyGenerateResponse {
	crawSeed, _ := rawSeed.PassValue()
	__ret := C.fil_private_key_generate_with_seed(crawSeed)
	__v := NewFilPrivateKeyGenerateResponseRef(unsafe.Pointer(__ret))
	return __v
}

// FilPrivateKeyPublicKey function as declared in filecoin-ffi/filcrypto.h:556
func FilPrivateKeyPublicKey(rawPrivateKeyPtr string) *FilPrivateKeyPublicKeyResponse {
	rawPrivateKeyPtr = safeString(rawPrivateKeyPtr)
	crawPrivateKeyPtr, _ := unpackPUint8TString(rawPrivateKeyPtr)
	__ret := C.fil_private_key_public_key(crawPrivateKeyPtr)
	runtime.KeepAlive(rawPrivateKeyPtr)
	__v := NewFilPrivateKeyPublicKeyResponseRef(unsafe.Pointer(__ret))
	return __v
}

// FilPrivateKeySign function as declared in filecoin-ffi/filcrypto.h:569
func FilPrivateKeySign(rawPrivateKeyPtr string, messagePtr string, messageLen uint) *FilPrivateKeySignResponse {
	rawPrivateKeyPtr = safeString(rawPrivateKeyPtr)
	crawPrivateKeyPtr, _ := unpackPUint8TString(rawPrivateKeyPtr)
	messagePtr = safeString(messagePtr)
	cmessagePtr, _ := unpackPUint8TString(messagePtr)
	cmessageLen, _ := (C.size_t)(messageLen), cgoAllocsUnknown
	__ret := C.fil_private_key_sign(crawPrivateKeyPtr, cmessagePtr, cmessageLen)
	runtime.KeepAlive(messagePtr)
	runtime.KeepAlive(rawPrivateKeyPtr)
	__v := NewFilPrivateKeySignResponseRef(unsafe.Pointer(__ret))
	return __v
}

// FilSealCommitPhase1 function as declared in filecoin-ffi/filcrypto.h:577
func FilSealCommitPhase1(registeredProof FilRegisteredSealProof, commR Fil32ByteArray, commD Fil32ByteArray, cacheDirPath string, replicaPath string, sectorId uint64, proverId Fil32ByteArray, ticket Fil32ByteArray, seed Fil32ByteArray, piecesPtr []FilPublicPieceInfo, piecesLen uint) *FilSealCommitPhase1Response {
	cregisteredProof, _ := (C.fil_RegisteredSealProof)(registeredProof), cgoAllocsUnknown
	ccommR, _ := commR.PassValue()
	ccommD, _ := commD.PassValue()
	cacheDirPath = safeString(cacheDirPath)
	ccacheDirPath, _ := unpackPCharString(cacheDirPath)
	replicaPath = safeString(replicaPath)
	creplicaPath, _ := unpackPCharString(replicaPath)
	csectorId, _ := (C.uint64_t)(sectorId), cgoAllocsUnknown
	cproverId, _ := proverId.PassValue()
	cticket, _ := ticket.PassValue()
	cseed, _ := seed.PassValue()
	cpiecesPtr, _ := unpackArgSFilPublicPieceInfo(piecesPtr)
	cpiecesLen, _ := (C.size_t)(piecesLen), cgoAllocsUnknown
	__ret := C.fil_seal_commit_phase1(cregisteredProof, ccommR, ccommD, ccacheDirPath, creplicaPath, csectorId, cproverId, cticket, cseed, cpiecesPtr, cpiecesLen)
	packSFilPublicPieceInfo(piecesPtr, cpiecesPtr)
	runtime.KeepAlive(replicaPath)
	runtime.KeepAlive(cacheDirPath)
	__v := NewFilSealCommitPhase1ResponseRef(unsafe.Pointer(__ret))
	return __v
}

// FilSealCommitPhase2 function as declared in filecoin-ffi/filcrypto.h:589
func FilSealCommitPhase2(sealCommitPhase1OutputPtr string, sealCommitPhase1OutputLen uint, sectorId uint64, proverId Fil32ByteArray) *FilSealCommitPhase2Response {
	sealCommitPhase1OutputPtr = safeString(sealCommitPhase1OutputPtr)
	csealCommitPhase1OutputPtr, _ := unpackPUint8TString(sealCommitPhase1OutputPtr)
	csealCommitPhase1OutputLen, _ := (C.size_t)(sealCommitPhase1OutputLen), cgoAllocsUnknown
	csectorId, _ := (C.uint64_t)(sectorId), cgoAllocsUnknown
	cproverId, _ := proverId.PassValue()
	__ret := C.fil_seal_commit_phase2(csealCommitPhase1OutputPtr, csealCommitPhase1OutputLen, csectorId, cproverId)
	runtime.KeepAlive(sealCommitPhase1OutputPtr)
	__v := NewFilSealCommitPhase2ResponseRef(unsafe.Pointer(__ret))
	return __v
}

// FilSealPreCommitPhase1 function as declared in filecoin-ffi/filcrypto.h:598
func FilSealPreCommitPhase1(registeredProof FilRegisteredSealProof, cacheDirPath string, stagedSectorPath string, sealedSectorPath string, sectorId uint64, proverId Fil32ByteArray, ticket Fil32ByteArray, piecesPtr []FilPublicPieceInfo, piecesLen uint) *FilSealPreCommitPhase1Response {
	cregisteredProof, _ := (C.fil_RegisteredSealProof)(registeredProof), cgoAllocsUnknown
	cacheDirPath = safeString(cacheDirPath)
	ccacheDirPath, _ := unpackPCharString(cacheDirPath)
	stagedSectorPath = safeString(stagedSectorPath)
	cstagedSectorPath, _ := unpackPCharString(stagedSectorPath)
	sealedSectorPath = safeString(sealedSectorPath)
	csealedSectorPath, _ := unpackPCharString(sealedSectorPath)
	csectorId, _ := (C.uint64_t)(sectorId), cgoAllocsUnknown
	cproverId, _ := proverId.PassValue()
	cticket, _ := ticket.PassValue()
	cpiecesPtr, _ := unpackArgSFilPublicPieceInfo(piecesPtr)
	cpiecesLen, _ := (C.size_t)(piecesLen), cgoAllocsUnknown
	__ret := C.fil_seal_pre_commit_phase1(cregisteredProof, ccacheDirPath, cstagedSectorPath, csealedSectorPath, csectorId, cproverId, cticket, cpiecesPtr, cpiecesLen)
	packSFilPublicPieceInfo(piecesPtr, cpiecesPtr)
	runtime.KeepAlive(sealedSectorPath)
	runtime.KeepAlive(stagedSectorPath)
	runtime.KeepAlive(cacheDirPath)
	__v := NewFilSealPreCommitPhase1ResponseRef(unsafe.Pointer(__ret))
	return __v
}

// FilSealPreCommitPhase2 function as declared in filecoin-ffi/filcrypto.h:612
func FilSealPreCommitPhase2(sealPreCommitPhase1OutputPtr string, sealPreCommitPhase1OutputLen uint, cacheDirPath string, sealedSectorPath string) *FilSealPreCommitPhase2Response {
	sealPreCommitPhase1OutputPtr = safeString(sealPreCommitPhase1OutputPtr)
	csealPreCommitPhase1OutputPtr, _ := unpackPUint8TString(sealPreCommitPhase1OutputPtr)
	csealPreCommitPhase1OutputLen, _ := (C.size_t)(sealPreCommitPhase1OutputLen), cgoAllocsUnknown
	cacheDirPath = safeString(cacheDirPath)
	ccacheDirPath, _ := unpackPCharString(cacheDirPath)
	sealedSectorPath = safeString(sealedSectorPath)
	csealedSectorPath, _ := unpackPCharString(sealedSectorPath)
	__ret := C.fil_seal_pre_commit_phase2(csealPreCommitPhase1OutputPtr, csealPreCommitPhase1OutputLen, ccacheDirPath, csealedSectorPath)
	runtime.KeepAlive(sealedSectorPath)
	runtime.KeepAlive(cacheDirPath)
	runtime.KeepAlive(sealPreCommitPhase1OutputPtr)
	__v := NewFilSealPreCommitPhase2ResponseRef(unsafe.Pointer(__ret))
	return __v
}

// FilUnsealRange function as declared in filecoin-ffi/filcrypto.h:620
func FilUnsealRange(registeredProof FilRegisteredSealProof, cacheDirPath string, sealedSectorFdRaw int32, unsealOutputFdRaw int32, sectorId uint64, proverId Fil32ByteArray, ticket Fil32ByteArray, commD Fil32ByteArray, unpaddedByteIndex uint64, unpaddedBytesAmount uint64) *FilUnsealRangeResponse {
	cregisteredProof, _ := (C.fil_RegisteredSealProof)(registeredProof), cgoAllocsUnknown
	cacheDirPath = safeString(cacheDirPath)
	ccacheDirPath, _ := unpackPCharString(cacheDirPath)
	csealedSectorFdRaw, _ := (C.int)(sealedSectorFdRaw), cgoAllocsUnknown
	cunsealOutputFdRaw, _ := (C.int)(unsealOutputFdRaw), cgoAllocsUnknown
	csectorId, _ := (C.uint64_t)(sectorId), cgoAllocsUnknown
	cproverId, _ := proverId.PassValue()
	cticket, _ := ticket.PassValue()
	ccommD, _ := commD.PassValue()
	cunpaddedByteIndex, _ := (C.uint64_t)(unpaddedByteIndex), cgoAllocsUnknown
	cunpaddedBytesAmount, _ := (C.uint64_t)(unpaddedBytesAmount), cgoAllocsUnknown
	__ret := C.fil_unseal_range(cregisteredProof, ccacheDirPath, csealedSectorFdRaw, cunsealOutputFdRaw, csectorId, cproverId, cticket, ccommD, cunpaddedByteIndex, cunpaddedBytesAmount)
	runtime.KeepAlive(cacheDirPath)
	__v := NewFilUnsealRangeResponseRef(unsafe.Pointer(__ret))
	return __v
}

// FilVerify function as declared in filecoin-ffi/filcrypto.h:642
func FilVerify(signaturePtr string, flattenedDigestsPtr string, flattenedDigestsLen uint, flattenedPublicKeysPtr string, flattenedPublicKeysLen uint) int32 {
	signaturePtr = safeString(signaturePtr)
	csignaturePtr, _ := unpackPUint8TString(signaturePtr)
	flattenedDigestsPtr = safeString(flattenedDigestsPtr)
	cflattenedDigestsPtr, _ := unpackPUint8TString(flattenedDigestsPtr)
	cflattenedDigestsLen, _ := (C.size_t)(flattenedDigestsLen), cgoAllocsUnknown
	flattenedPublicKeysPtr = safeString(flattenedPublicKeysPtr)
	cflattenedPublicKeysPtr, _ := unpackPUint8TString(flattenedPublicKeysPtr)
	cflattenedPublicKeysLen, _ := (C.size_t)(flattenedPublicKeysLen), cgoAllocsUnknown
	__ret := C.fil_verify(csignaturePtr, cflattenedDigestsPtr, cflattenedDigestsLen, cflattenedPublicKeysPtr, cflattenedPublicKeysLen)
	runtime.KeepAlive(flattenedPublicKeysPtr)
	runtime.KeepAlive(flattenedDigestsPtr)
	runtime.KeepAlive(signaturePtr)
	__v := (int32)(__ret)
	return __v
}

// FilVerifySeal function as declared in filecoin-ffi/filcrypto.h:652
func FilVerifySeal(registeredProof FilRegisteredSealProof, commR Fil32ByteArray, commD Fil32ByteArray, proverId Fil32ByteArray, ticket Fil32ByteArray, seed Fil32ByteArray, sectorId uint64, proofPtr string, proofLen uint) *FilVerifySealResponse {
	cregisteredProof, _ := (C.fil_RegisteredSealProof)(registeredProof), cgoAllocsUnknown
	ccommR, _ := commR.PassValue()
	ccommD, _ := commD.PassValue()
	cproverId, _ := proverId.PassValue()
	cticket, _ := ticket.PassValue()
	cseed, _ := seed.PassValue()
	csectorId, _ := (C.uint64_t)(sectorId), cgoAllocsUnknown
	proofPtr = safeString(proofPtr)
	cproofPtr, _ := unpackPUint8TString(proofPtr)
	cproofLen, _ := (C.size_t)(proofLen), cgoAllocsUnknown
	__ret := C.fil_verify_seal(cregisteredProof, ccommR, ccommD, cproverId, cticket, cseed, csectorId, cproofPtr, cproofLen)
	runtime.KeepAlive(proofPtr)
	__v := NewFilVerifySealResponseRef(unsafe.Pointer(__ret))
	return __v
}

// FilVerifyWindowPost function as declared in filecoin-ffi/filcrypto.h:665
func FilVerifyWindowPost(randomness Fil32ByteArray, replicasPtr []FilPublicReplicaInfo, replicasLen uint, proofsPtr []FilPoStProof, proofsLen uint, proverId Fil32ByteArray) *FilVerifyWindowPoStResponse {
	crandomness, _ := randomness.PassValue()
	creplicasPtr, _ := unpackArgSFilPublicReplicaInfo(replicasPtr)
	creplicasLen, _ := (C.size_t)(replicasLen), cgoAllocsUnknown
	cproofsPtr, _ := unpackArgSFilPoStProof(proofsPtr)
	cproofsLen, _ := (C.size_t)(proofsLen), cgoAllocsUnknown
	cproverId, _ := proverId.PassValue()
	__ret := C.fil_verify_window_post(crandomness, creplicasPtr, creplicasLen, cproofsPtr, cproofsLen, cproverId)
	packSFilPoStProof(proofsPtr, cproofsPtr)
	packSFilPublicReplicaInfo(replicasPtr, creplicasPtr)
	__v := NewFilVerifyWindowPoStResponseRef(unsafe.Pointer(__ret))
	return __v
}

// FilVerifyWinningPost function as declared in filecoin-ffi/filcrypto.h:675
func FilVerifyWinningPost(randomness Fil32ByteArray, replicasPtr []FilPublicReplicaInfo, replicasLen uint, proofsPtr []FilPoStProof, proofsLen uint, proverId Fil32ByteArray) *FilVerifyWinningPoStResponse {
	crandomness, _ := randomness.PassValue()
	creplicasPtr, _ := unpackArgSFilPublicReplicaInfo(replicasPtr)
	creplicasLen, _ := (C.size_t)(replicasLen), cgoAllocsUnknown
	cproofsPtr, _ := unpackArgSFilPoStProof(proofsPtr)
	cproofsLen, _ := (C.size_t)(proofsLen), cgoAllocsUnknown
	cproverId, _ := proverId.PassValue()
	__ret := C.fil_verify_winning_post(crandomness, creplicasPtr, creplicasLen, cproofsPtr, cproofsLen, cproverId)
	packSFilPoStProof(proofsPtr, cproofsPtr)
	packSFilPublicReplicaInfo(replicasPtr, creplicasPtr)
	__v := NewFilVerifyWinningPoStResponseRef(unsafe.Pointer(__ret))
	return __v
}

// FilWriteWithAlignment function as declared in filecoin-ffi/filcrypto.h:686
func FilWriteWithAlignment(registeredProof FilRegisteredSealProof, srcFd int32, srcSize uint64, dstFd int32, existingPieceSizesPtr []uint64, existingPieceSizesLen uint) *FilWriteWithAlignmentResponse {
	cregisteredProof, _ := (C.fil_RegisteredSealProof)(registeredProof), cgoAllocsUnknown
	csrcFd, _ := (C.int)(srcFd), cgoAllocsUnknown
	csrcSize, _ := (C.uint64_t)(srcSize), cgoAllocsUnknown
	cdstFd, _ := (C.int)(dstFd), cgoAllocsUnknown
	cexistingPieceSizesPtr, _ := (*C.uint64_t)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&existingPieceSizesPtr)).Data)), cgoAllocsUnknown
	cexistingPieceSizesLen, _ := (C.size_t)(existingPieceSizesLen), cgoAllocsUnknown
	__ret := C.fil_write_with_alignment(cregisteredProof, csrcFd, csrcSize, cdstFd, cexistingPieceSizesPtr, cexistingPieceSizesLen)
	__v := NewFilWriteWithAlignmentResponseRef(unsafe.Pointer(__ret))
	return __v
}

// FilWriteWithoutAlignment function as declared in filecoin-ffi/filcrypto.h:697
func FilWriteWithoutAlignment(registeredProof FilRegisteredSealProof, srcFd int32, srcSize uint64, dstFd int32) *FilWriteWithoutAlignmentResponse {
	cregisteredProof, _ := (C.fil_RegisteredSealProof)(registeredProof), cgoAllocsUnknown
	csrcFd, _ := (C.int)(srcFd), cgoAllocsUnknown
	csrcSize, _ := (C.uint64_t)(srcSize), cgoAllocsUnknown
	cdstFd, _ := (C.int)(dstFd), cgoAllocsUnknown
	__ret := C.fil_write_without_alignment(cregisteredProof, csrcFd, csrcSize, cdstFd)
	__v := NewFilWriteWithoutAlignmentResponseRef(unsafe.Pointer(__ret))
	return __v
}
