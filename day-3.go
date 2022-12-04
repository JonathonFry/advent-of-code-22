package main

import (
	"fmt"
	"strings"
	"unicode"
)

const day3Input = `
QLFdFCdlLcVqdvFLnFLSSShZwptfHHhfZZZpSwfmHp
rTJRjjbJTgzDJjdsRsfwtfNwtfmZpZNhmmzt
bMdJjsjglnVMFCCc
BZvZMBBBMTtZTgcCPdgtgQCrrV
VspNDDpsGDGnRmRpRplQdrrPcPCjgDCcPQCQQj
RVnsmspnpwFRlHGRwHHlnRSThSSvBTbFTZMqTMZMTZFh
ttffrVJWtWpgtQnZGVnNSLTHSZ
jRsjjmhdRcjcRsFcdwGLMZSnHMTSMNZN
RjczlvjhPCcPjcvRpbglWplJBblqrGgq
NwCWwdNQNDCwGpWNQwmJtZgPZrdJgLZcPhLddr
blqpnFTqrLbcLPtV
MnjqSSpqlFRvSqNDGHvWHQDwfWmm
jfLljlQhDLmlrMJQtJTJrQqQ
NpBbjjsdMCgCCMrb
dwspwGnSHHGsGzDDlFDjVWjfZWnP
wQhTZwvpZFZdqWLnnwSrWC
mfDmMFlDcPLdgDSCLCqg
PmzclsMclMlFsHHsJZFHpT
LfLJWNdJnBLfhndfWdZqcgDZgSqgCCSSSLDF
bQVQmrrjPqQDZSZBCQ
RRtGjVmRsPbPrrnNNpNHHnBnpHns
PfbGNwGBwNcPTbGNQFBVjsjztVztVjHV
hrdCJhmlJhZrLDRmghrmDrzqFsbgtbHqnznzznQtbjtz
WdZdDJCDmrJmLZrLDLDZlClcSccwSPbNPfSNWfGNNWMGNc
QwrnTSgqgFShSdfHPdbS
BGdjmMmZMvfhvCZZPf
BzGmzVGGGzmzGpVBtdnQqqdTQQDDqrpR
PPRPwLQlLtbPmbwgJhrSssNlhhrgsZ
fFTdFvTMNfzVnFqdnjgSSjjsSjhghpJs
dvczcFzNTWVWMFLcLbQQwmbHLCLL
HhLLDfMmTjsjmLmhsmmfZMjGtpGJtdcvnCWtZJcWGddttW
gwrwwgzwgDpJddrJDr
SBwBBbgVsVmRRhDM
SZdmfdZRTQZTQgHVVGGRqZdVCjCcNCNcJRWcCBbJjCPCsNsc
FnhzMMhDDFlzlnvpwMLrMDCsbcjNJbcJbBcBfPhCNbWj
wzwnpDDnLvFnLlttLrzGgTVGQqZtTqSqSdfZTg
FJJWWWrCGWdmzFlTVqqlMhmvVlNh
btpgtfjZjjhgggrNMThl
DpwpfRZDZwwfwjsnjfsZnnnwGcCRCHcCCCLGHLWcrcFCWCHW
ljHHHBtjQthchhZpqqNt
FTmJnPFwzlJPmzTgTgbFwJbMCpbMchhhqbhWCDMDhZcppM
JwlFGwVGnFFGBGjdSsfdsG
QsvpGhjGgswvjjwjjjvPpThJfLZCCLCSSLbFLStCFCSgtH
mDrzdzMqqnMrDMmZnqnzNVRCStlCHFSCtJqlCLFCFLfJfJ
mBDzNRWBDDMBpGsZcGZWGjPp
SlLQhQsvvttFlWsWcfHHMTJfwSHwTfTf
VZmmrRDRfpTHJcRf
jzBnDDgjPchWlsQsBW
LTLVdTSLNTLnTSnrWvdwswsfmJwmwGsffH
FbgCbRRppCpPbgMcZvCcGftGGltwHwGtplQQsfJw
CRBMCvZZRgMgBbDCPcDrjLzWLVrSSzShSSNrBS
hVJJjhjRVRZjQvDfBstsBVNBdwstHsld
pCTCcMqCThTFLFFPWcWSPHtwwdmcBHHmNtHdmwmwBl
rMTCCWPLLPCMFhDnDrjzRrfDJD
pqMpCvMchvFNWSBdVNqQ
zDRJJDGJJtNtmGRRWVdFWWVdSfjb
DDJLmnJmzwGmGLTPhTCNpgcrpv
cpPpbPWVprWcbJrrwpCwwdWrvNNFRqzNnChgqFzFnZvqFlzq
fTtHLfSHSsNDGLSmsLvnhFqzhzlzDhhvRlFz
mSMLHTQTmHMSfBSMTPdBJJNNVddrVrbjbJ
zpCpBTnFgFbncznbblzdhRswdlJsLllJdw
QqqmtWVPWvHDVmqDhjsljwRhlZldhRMQ
tWSHDmVfmrtPHVgGRRbgTRpSgpTc
ssTbzFRtPRwTFZtvbPRMhndBqMMvMBHJnnHMMd
WQVWzlGWVBqqdMQJMq
pVSpSSgLfjDzWrLGWWjDzzfLtbRFFNtPZRssspsNRZcRsNZb
jnPzzGlnnznWnzhvGnnpVFrZmVFcgjrrmZRFtj
fsbgTdwdqBbfwCptVtdZRcrRCp
gsMbgfHsBSwsGhhJWMLnWPPJ
bpmbJpNbbbNGGmRJzJTsfdsvsNdglfhssCvC
hWLwQjZjLhjHFFBLZldvvflrvtfjCrfjlT
QWVQZZFDcDJJhJJc
RmRghgRlNgfGGRmdGqhsgsZFZZpBvHpZppHcgH
tbLCDLnLtSbbbjtPtMLtDPTvHHBpcHcsHvTcHsmZcF
rSLrMJzDznzGmhNlVwdrRr
vWjljMWcnSSpjmzbJVzJrTCmtGJV
NZDDQLRqPJrPzrprTC
gqDqqwpdHWhlgnjH
ccptcpstDvbNvHbLNRZZ
dFjhdnjQFJlFCQSjgngJPJgWWrRNWNRtNCzrVbRzNVzZZL
PhThFSPPSpsmTcqMwt
cLcLlMhGMGcpGzslHFHFvnHlBDvWbT
VHdQwqPJdPwjJPdPQRrmjjjQnTFrbvNWFFffDvbvvNDvbBNN
RwRJCHmmdQJZZLzGphcCtz
hVvFVjvjVWmFRQVZqTpqtwQpwpqZfp
gvDlSBDJSlPLcLdDwzwtptqTTTzwcCCt
JgrJGbLgvnWsmvVr
rwmqqRqrnHQGmnjCCqCzdBzCBJBz
hFLgbWWPWmvtLhPtgpcdjJdcBJdpJjDsgp
lvfhSSPWtTNTTZZmfr
bHDDssRHsjNMbJjJLQJsbTtGvSCzCGQCTzGvSqSBzT
mmVrwhmmpfPStnTSBnhStG
pcwrptZcgFcmpgHRDjjZlDJsjbbD
JJRrmFqJMdFFJMjjJcqGgzSCSHSCscPCHPHGZc
VWpWptnvSmpPGCHC
vQnDLBmbntvLBbnlldTQFlJlFFrNRd
LPDftnHFQfwmBcBGmc
CVqRsdqvdrlsCVsNvqwwSpTNSSDSDDBBTTSN
lqlDRddjbblRbRqrlRRjsbvghHHnPQWjHWQWZHPHWZhFnP
bwQsDcgsJqcsDpcQRQnpqtVSVvgSMMMfMvfVBVfdvM
CGZFrHHPrTZNGGZZHmCZHlVfjfzjSfzBtBBNSBVjvntf
ChrCCLGrTlhJnhDncRbp
nmFnhfTQjSzfjddZWsRRRFRFGl
HDgCwgtQbZlHsrqHHr
cJPCgCCPbpbgDMPvMQnjmnhTfmzLpQQmjz
QFHSQdNMCSgcSgFtttPNFJpCpnTjZlbblpppZplZjz
LqLsWMRRfrrWMmMGpbTnbppbZnTjpnmm
fWBrMqWsGswGfGRMMwrLgtPPdNFBQPHNHNPPcFPP
dngbSppJSSppbVMZQQMjqfQQgwcl
TWmSWtvCRCWjwfjqQqMstq
FhFvRzSTNmhHnVPhGhBJdBpB
gcHPgzGmPPwTsSTsbwbdWD
QjBLLfVhhBqqBFQLrLjVFlNpNDtsSWTDdNptdbqbdS
jCMFLVjFBFjJJLFFMVBFrLnvPzHRmHPGnGWWcCvzHRZm
PDPqWWjhPpPbCsjwjTVbLT
SrtCttGRMddSVwHFSs
JtfvttmrGMRRJzJCqhqqqWQZhCNgJZ
ChrCVFQCVQlwQNwpQcmmcjmWBmddghjjdW
sbDTZStTqqfSBggPmWjWWNsL
THqqSHDTtZTDTHZZbHTzRzFvlFCVprFprQVnCzNppQrz
PdfWCwMWjPSrdgCMnnlGsGQvvpJZvFGnps
DmBhVBLbbVqVBzTRLBRzzTLNNpRQNNZQZppZvlQpZvllvF
zVDtVHBbbTbzDbrjgWjMPtMWPMlj
JLsTTNDsgTMNvDQpLpGpLGNJShrfzCFnSnSrnfzCfTFhWrfw
ZcqrRddHZZVRfzWnVWCzWFnn
tZHZtrHHPdRtdHlcccQggsplpJDNvMGNGMss
cMCLfStfMTCjPMPcGzjftMbgsRNmRgmmGsmnJbNJbghJ
QHVVWrFFWZNShHSgbSJm
qZwwrrpqpZpZFvqrQdFlQVSwLBMBfTTLTjLBCcdTMzMftPPB
SwsdBTvgvJLPNptpCpCmBDtn
wffrzwGFWFNZWpjWZnNm
zrfflbRwJPhbPbsS
HjHHRtwjnjRblQRttHwQGvGWNNBWvqGzfTvfNN
FmScCcrsdVZrpBrVcCVFzffvzzmWGLWgqWqgGWzW
SFVSDDBdsdDSJhnjJltJbPtHRM
FjGFVqWrzQFlQrZzGQzFLTvfwwTgMnvcnbRMLRdnfb
CCttSNsSnRfgncSg
CNspmDBPtPmJJNBJPNpDhQZVzQlhqrGZflfVjQFrQj
dNNdHWcmdmPPptmmWHpPTFFjJPGrQsVsPQGGGJVDrVVGrS
MhZlZhlgflgfnfDtjbjJGbtnVtGS
LtZqlzhzqMZWHHLwdHmFWp
llNRlfwWRwwLlwFNNgRrVCBjdjCVdjpWjtVWCD
HTQqzPqzQPmhhmSPznSsssJtdnMZddtMCjprtMjCnBVnjZ
PzHQmqsGSJPSmQqPbfwNcgNbNgNfBGwR
lPdzlZPzQzMZQGQrTZvvpjHTTpfsTTZb
zRShhtWRnqnqSNRnDTTHvfNJspNsLpTsjL
hBVncVtDSnhDnDBBtGrlzwmmMlGmVrMdrP
HPTZVHVPlHDPlfgnjJFdJdjPjSPqCS
hLRRBhwGhqbtmsRSSSjjdMJjnJGSMj
QrQtqrRrcQDgVglc
ZTwbbZdchZZjmVWHTrHWBVJtBB
glslCDqLLDfGRqlsgLssfrCHBHFHmrHBBppFmCJWWp
fRzvvvgGgNSNvmQbSQ
qPGGPwCTqTzHCvPGqWdLFLssLpstLLspvd
njJchhcbjbDrbcLNlLrpWWrLLHgp
DQhMMMJQMQJnVbbnRHSMPwZmGZPZRTRCwTZmZGwz
zzGNfPbcgdPqLrqvWWVzMq
DGmJtnJTJRhhJMhCQqCLCLrrLM
ZnHDtSZlTBHnBdccGSfGcwjjdb
FpZDpQZDvMwZpCCMdCBPpJGPPLgJGGLffJJL
jlbswNrlPPJJfGlf
bnNwqbHnNwRSrqhbdCcmHddQzddFDdvZ
gbQQQngWPVVtvvPQNVNvWWSHGwDsCCmDtHSlmrssDmHs
fqhMLFFMMZqZMRZqMjRMqLJSCdFlrrldsDrCsDSSHHGCSC
MJRZLZLGMcTqczjNPzNnzPvWBVgNnP
gqdbBffTvlRHbwLl
nMMQJQpGdsFpQsJzNMRLLDlmLLmjFFmLjDRF
pzGMnVcMBfTdtBBV
WSbfmrrrrWdbWmdfDSSStmHjtMtvCLVnqBHCVGtVGnMM
lRcgFRZhJgnMLjvGgv
lcvwTcFTplvwphzcTTJTbsdsPSPDdbmzmDSWPsSm
bbdTjTQTQMsZNqqhJrZslg
jFGVjwfCPVGfwjCVqWhWZFgqWrglllNN
PjfSPzRBjCCfSBCGBLznTndHcdMLbMmmdT
wSVMJSVccdGwGnsgbVTTbRsCRNgN
rHjhHLmrhPJrqjNTRDgBbbRRRs
zqmPPqqpPLzltrMdJcZpfdpGWWJJ
ZhrBBJGrgJhGHttGGVPPcPPF
cnzLqNssfRnpfWqsLfcfWQNMbMVPDtnDtbHFtMbPtVPFFM
jfqzCCLsWQLcjgldjmljmgTd
wghGSSGZPVwgqtwtwCCtFFMM
BvbspnBznvvWHWHHHbCQptQFQlFcqMClqLLq
JWzzsJHWzfWjJrvMBWHBBGDmVDrVhZmmgSPSmZVVrh
ccRMJRsjjgJgcPCSCCVCwsSWVNzp
WQQqnmrBWtqWqdSbVwwBSpbbCSBB
QvDqmqqmgWPWjPvW
msqpjDWspRWwvFvDWWhnbbJfPzFQblJJPlnz
gGGrMTgLVBsBBLdsVTrSCBffHQfdHhnbPPPPffndlbzh
ZCVsCGSScsLZpwNpqmZRqW
PPsGmJPVPQPZmsQCVPJPnPCMDcTcdqDDTqvFhvnTjRDTDchq
BdrtzNBLHStHrdrlwfNThvFhcvbDccThjbFBqq
SSgdHNfSHHgzLHtLNWSPQQPMQVpmmppVCmQZCW
pPssrWWLdndHPJdd
QNQFTLNBFTzzgjfGTjffFNZjCSGnHDnSDJHnDScttDCcDnmd
FVzVLZwZZgswqqrbphbR
VpWCZjCwWnppZpqnhNjjNZjFLtLzQJHdHLQRzWLRzRztHJ
DMGPmPMgTSmsgQzRFbdHRLJgdn
csDMPMGDDvMSSPnDTvrDChhwljlqNNjchNCjNVcf
WpGGmbSGpVWWpjMMTNdfCFNdFfRNwNSF
JsQztzrvrJqsTTRbbvFBhhhv
cLrDqLccsLqbDHGpZWDHgjGlZW
QGMQJMmsJmMCmmqjsRvLvvdgvgVvDVdD
BDcrcNbNppwTpzRdvvchhFvfFv
plBBwWrbpQHDjGmGJl
mzFlTdmSDzrPvCJqqDVVNC
hfRmhgjRhnfwnRHcnhGGvPJQPvvfLfQvNLGv
BhhnjMgRWghpwjRWMRjrZzdbSbsdstTrltdmMs
bLLnbqjpvplnDvNlqpqBWJZSdPJCNdJJThhSPhTd
HFwHHQMMFHGzGwRPPJPTWthTZtJSQr
mfWMHFHWHmgmFcwGwwpbDljqjBDcDnLcVnlb
wBrWBwSWRJMBwdZnPQPgFnwGVF
fLjfbsvDDfvvqqGqZGqmPQgqTGGG
vZLsjzjjZCzJWRNSBR
jTRbRHHqPqTRBHqdjhgvgghhZQdDvvgvhC
WLWWzzFszsmNFGWSFmMrpghCtZvhlQNDgQCDgctC
FJsLsSrDmsFSDLWrzJmmMsGqjRBVbJTBVPVBbBqRjPBjHn
QbwwnDDQDcDfSbDbfhhrvrCtJMvJSCvvJh
FWRjjLjmdZWdWNBFNWNlNQQrMGvvMGgssGvQRvrMJs
BjWdlBpmdmBWFWdpWfPfpVnVwfHpqPQDbq
SqrvlMldqvSWdGPTGzWpWpzpHP
tRwmhtbsRRFsLwGGTVDHppTNdbVp
FRCRQdCFtCLmBhCcmmQdhFdCvnfjffjZlZnjSnvfcSrrgMgn
GQQtNJQWWcqPPhMMtwqD
WpWLlBWZCvhjwMMZqDDP
WgvmLVmHCbpppLgdllHddvCmFGzGnfsJJQJsJncSsccFVffF
HcSsSlTTvvPPWWNMWWgPTPPbGbbrwJQbrrDphrHJJRpRhp
ztfLqqzmRwDGlLDb
fdVtmqjdZBmSvjsPSWlTgv
DPvDhhMRRMhRNDLPMNsbwHwrjgnddqddrWdPtHzr
pcBGSpcVBfJWCcmJGGwHtzgrrtwqzdrtrngG
mllBlBZmMlQWRbQv
SGZBSFMZllJWmzvfpp
NTqbNrhHNHWgNqHrNhNQbbjHJLcnfnzLLnLmfcfccJcfQLcL
HggbTNRRTHqqbVSGMSZVWDMDwVPs
SBsSlvbPlFPvRlbPsMFZLgVLrLsJVgzrCJfVCH
jcNddNdGzZrVgNVJ
tTGwdcmWGdtwQmwmwZdwSlhBPbhPTBFRhlhSMFMR
RzStzTzzvvQvSHVvhVgBqMMFqhPM
ddlLLwNVLWLjbbLrjrbWrwmlhcFmBGgFMMPgBcGBBqPhggMs
dLwdVCVWWdfNwNwLrWrbfbJNptzDDHRnHptHtznHTppnQCtR
RzcfMBHLzpDQFmnDSWNB
dbqjtjVqJZZGjPGJCPGbPndNNDglrmQmNSDgSlSSng
hjCTqhCJbhVCGNvMcfvhfRLhvchz
sDDqDMtqshJhPvhhCpSCCWlZHSWp
bffRcbBGGTwGfGfbNjgSHZSgWwplHCClZZ
RTQBbcnbRNmGbGTQLbmbJVqLllsDVMsPDVVvttMd
nbLBjnqwgfRRBgBwnllbLlwScvPdZPcScZPcdFZJPvZPvcMZ
tChQpphHrrHztssZdcDJcPZcMvWv
hpTHVMQMtQtVpzBfwjfRnfwfnjVl
ljJlvvJQlrlcJcWpPzgthnPnzMgpgSpC
smtmZBmHZTVttHmqFqmzCSZSdndzShPNgPShgP
bVqFHLqLqfHHFwbBLHcwDQrDrtjlQvGjlRQQ
pwhVsPvVVCFtmhPhzqGqqZMZvGTTTMlGWM
drrrrDfDRrNQdQdrRrBdjGWqWqWlGlGtlGbGZGBTLc
tSDfgnHrdDtVSPSshJCSPh
WlWlDqhglLhsdgrcbFdJJpPpdBbB
ZQZvSvzRMSzjZjvZmMMpbFPQFVBrVbPcpbJFLB
SwGZmjvCRSMRjMzZvRnstHftNfswHsflLhNWHf
jsprCvGRQrtjCsQrGsrzvGHhgmHVmHZgggmMGVmhMbHm
FFFdDSdwSffJWqqMzzMmDVbZ
LLcdcfcfPwwBzdTTdtvlsrjCtvPvprnsjR
MvtSqNSWMzjwzFTD
ZRPlcRpQszNgszNwVT
bcZcrcPlcPLLLZllPlbcbLSBfWCvHvWWNSmSqNqfWN
rNdZpMGnddgggwHwzRPCzDDD
vcvhcTLhZLhLPCPHPDPPVvzH
LTmBmthWBchWLttttFJFLlFnGJNsfpdjNsnMnMpnpZdssn
ZHWFCvqBDdqqqCTDHHBWrgppTMhhVpspMPQcSgQVPS
jblbGffntRwltfMQVrrQscphfg
ztJrGtbwGztbmtzzRGnRznWNNCWmHHdFHdFNWWHHqCqZ
WGWSSZvVvqmrmzPm
NgjtwFFlwDsFghNsMtlcjljcPqrQHcZzQznpQQprnqqzHQ
tgMCwNhtgbdLZRbZCT
PQSPQrSGZnGnVFhpVhRRlvLvBDRV
tjctcjTMMpDTvFTlRD
JCftsccFCcmsJJGZGGmPHnQrGwGS
TrjRFFRnpnRCHNFSjSRrffJvJfzqQBsjqQqzzffd
ZtlgMDhZhgmGDLVZLlGtLPqdQQvvfBJJqzzBPdMzdd
VlLDgLLDWtGZwgtRNTNrFTqCwqHTrr
LpcDFDMMPjMLLjpcDGCHgHssGHWnbCBWBHvm
QfZhrhVVdZThlZlfVvVzZrTbgQnBHsCCHgJBsCsJBHmBmn
wwtvfZztlTVlhtrzzlLNpFFRjMPDpRcPFwRj
VzZhhQHQJJWJSSFWDGclbmNPgglPgVGc
ddBTqCjjBCcrqrCRrwGPGmmDGmbpBGNpNNgg
CRMjwsjwsLdLRrQFJSvMFMWZcHFW
JgJJPvtrhRPQQzSRMQFFSF
BLqsjsdLsMBqblnsGbBqVqdwSQSCSWwNFwczQWCNNwNCHn
ljqbpLbbdDlbDbqDDVtMttTTgpJJgThhJrJr
nflndmjbSnlTQGwvWGPHGRGj
NtstcMcDJMvwgHfFvDgR
qqqpLrMsLLqLNNnzbrdlbZSrznfz
ttZCCFjNjnPVCFQPPFbbStrzqzqrrrcwtmJJ
gTTMRMTWsTGGTddHTTbBzBLSmqbbJGzGmqqb
HpgpMTvRhHHTRDhMsHdHDRhjJlVPJjNFJnnFpQQVfPCjnP
VqJVQPpjQqPBbHwldmLfVVmd
tMGvrzzDGCDDddwLbgLvLwcm
TWDWCzTZDGMZtzWWtsFhbRRqRQRjhbNQBBTh
zgLgLHnnzCCvnsHSsZBZBsTRdD
rslllhJjcQNNGjpWJlSRTRdwBVSSNTPVSdPB
jGrGqjJfqccrfqGcGplrJpFvzggqmCtMzmsMnvMvvCgm`

/*
The ASCII value of the lowercase alphabet is from 97 to 122.
And, the ASCII value of the uppercase alphabet is from 65 to 90.
-96
*/
func day3() {
	var score int

	trimmed := strings.TrimSuffix(day3Input, "\n")
	trimmed = strings.TrimPrefix(trimmed, "\n")
	backpacks := strings.Split(trimmed, "\n")

	for i := 0; i < len(backpacks); i += 3 {
		fmt.Println(i)
		fmt.Println(backpacks[i])
		intersection := intersect(strings.Split(backpacks[i], ""), strings.Split(backpacks[i+1], ""), strings.Split(backpacks[i+2], ""))
		fmt.Println(fmt.Sprintf("%v", intersection))
		value := charCodeAt(intersection[0], 0)
		if unicode.IsLower(value) {
			value -= 96
		} else {
			value -= 38
		}
		score += int(value)
	}

	//	for _, line := range strings.Split(strings.TrimSuffix(day3Input, "\n"), "\n") {
	//		if line == "" {
	//			continue
	//		}
	//
	//		chars := strings.Split(line, "")
	//		size := len(line) / 2
	//		first := chars[0:size]
	//		second := chars[size:len(line)]
	//		values := intersect(first, second)
	//
	//		println(line)
	//		println(fmt.Sprintf("%v", first))
	//		println(fmt.Sprintf("%v", second))
	//		println(fmt.Sprintf("%v", values))
	//
	//		value := charCodeAt(values[0], 0)
	//		if unicode.IsLower(value) {
	//			value -= 96
	//		} else {
	//			value -= 38
	//		}
	//
	//		score += int(value)
	//	}
	println(score)
}

func charCodeAt(s string, n int) rune {
	i := 0
	for _, r := range s {
		if i == n {
			return r
		}
		i++
	}
	return 0
}

func intersect[T comparable](a []T, b []T, c []T) []T {
	set := make([]T, 0)

	for _, v := range a {
		if containsGeneric(b, v) && containsGeneric(c, v) {
			set = append(set, v)
		}
	}

	return set
}

func containsGeneric[T comparable](b []T, e T) bool {
	for _, v := range b {
		if v == e {
			return true
		}
	}
	return false
}
