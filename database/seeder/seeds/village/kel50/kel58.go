package kel50

import (
	"fmt"

	"github.com/danangkonang/rest-api/config"
)

func Kel58() {
	db := config.Connect()
	_, err := db.Exec(`INSERT INTO villages (id,village,sub_district_id,regency_id,province_id) VALUES
	(58001, 'MALEI II', 4426, 318, 23),
	(58002, 'MARTASARI', 4426, 318, 23),
	(58003, 'PEDANDA II', 4426, 318, 23),
	(58004, 'LETAWA', 4427, 318, 23),
	(58005, 'MAPONU', 4427, 318, 23),
	(58006, 'SARJO', 4427, 318, 23),
	(58007, 'SARUDE', 4427, 318, 23),
	(58008, 'BULU MARIO', 4428, 318, 23),
	(58009, 'DODA', 4428, 318, 23),
	(58010, 'KUMASARI', 4428, 318, 23),
	(58011, 'PATIKA', 4428, 318, 23),
	(58012, 'SARUDU', 4428, 318, 23),
	(58013, 'JENGENG RAYA', 4429, 318, 23),
	(58014, 'LARIANG', 4429, 318, 23),
	(58015, 'MAKMUR JAYA', 4429, 318, 23),
	(58016, 'PAJALELE (PAKAWA)', 4429, 318, 23),
	(58017, 'TIKKE', 4429, 318, 23),
	(58018, 'ALU (ALLU)', 4430, 319, 23),
	(58019, 'KALUMAMMANG', 4430, 319, 23),
	(58020, 'MOMBI', 4430, 319, 23),
	(58021, 'PAO PAO', 4430, 319, 23),
	(58022, 'PETOOSANG', 4430, 319, 23),
	(58023, 'PUPU URING (PUPPURING)', 4430, 319, 23),
	(58024, 'SARAGIAN', 4430, 319, 23),
	(58025, 'SAYOANG', 4430, 319, 23),
	(58026, 'ANREAPI', 4431, 319, 23),
	(58027, 'DUAMPANUA', 4431, 319, 23),
	(58028, 'KELAPA DUA', 4431, 319, 23),
	(58029, 'KUNYI', 4431, 319, 23),
	(58030, 'PAPPANDANGAN', 4431, 319, 23),
	(58031, 'BALA', 4432, 319, 23),
	(58032, 'BALANIPA', 4432, 319, 23),
	(58033, 'GALUNG TULU', 4432, 319, 23),
	(58034, 'LAMBANAN', 4432, 319, 23),
	(58035, 'LEGO', 4432, 319, 23),
	(58036, 'MOSSO', 4432, 319, 23),
	(58037, 'PALLIS', 4432, 319, 23),
	(58038, 'PAMBUSUANG', 4432, 319, 23),
	(58039, 'SABANG SUBIK', 4432, 319, 23),
	(58040, 'TAMMANGALLE', 4432, 319, 23),
	(58041, 'TAMMEJARRA (TAMMAJARRA)', 4432, 319, 23),
	(58042, 'A YANI PURA', 4433, 319, 23),
	(58043, 'AMASSANGAN', 4433, 319, 23),
	(58044, 'AMOLA', 4433, 319, 23),
	(58045, 'BATETANGNGA', 4433, 319, 23),
	(58046, 'BINUANG', 4433, 319, 23),
	(58047, 'BINUANG', 4433, 319, 23),
	(58048, 'CAKUNG', 4433, 319, 23),
	(58049, 'GEMBOR', 4433, 319, 23),
	(58050, 'GUNUNG BATU', 4433, 319, 23),
	(58051, 'KALEOK', 4433, 319, 23),
	(58052, 'KARANG PUTIH', 4433, 319, 23),
	(58053, 'KUAJANG', 4433, 319, 23),
	(58054, 'LAMARAN', 4433, 319, 23),
	(58055, 'MAMMI', 4433, 319, 23),
	(58056, 'MEKARSARI', 4433, 319, 23),
	(58057, 'MIRRING', 4433, 319, 23),
	(58058, 'PADANG SARI', 4433, 319, 23),
	(58059, 'PAKU', 4433, 319, 23),
	(58060, 'PUALAM SARI', 4433, 319, 23),
	(58061, 'PULAU PINANG', 4433, 319, 23),
	(58062, 'PULAU PINANG UTARA', 4433, 319, 23),
	(58063, 'RAYA BELANTI', 4433, 319, 23),
	(58064, 'REA', 4433, 319, 23),
	(58065, 'RENGED', 4433, 319, 23),
	(58066, 'SUKAMAMPIR', 4433, 319, 23),
	(58067, 'TONYAMAN', 4433, 319, 23),
	(58068, 'TUNGKAP', 4433, 319, 23),
	(58069, 'WARAKAS', 4433, 319, 23),
	(58070, 'BULO', 4434, 319, 23),
	(58071, 'DAALA TIMUR', 4434, 319, 23),
	(58072, 'IHING', 4434, 319, 23),
	(58073, 'KAROMBANG', 4434, 319, 23),
	(58074, 'LENGGO', 4434, 319, 23),
	(58075, 'PULLIWA', 4434, 319, 23),
	(58076, 'SABURA', 4434, 319, 23),
	(58077, 'SEPPORRAKI', 4434, 319, 23),
	(58078, 'TAPAMBANUA (PATAMBANUA)', 4434, 319, 23),
	(58079, 'BONDE', 4435, 319, 23),
	(58080, 'BOTTO', 4435, 319, 23),
	(58081, 'GATTUNGAN', 4435, 319, 23),
	(58082, 'KATUMBANGAN', 4435, 319, 23),
	(58083, 'KATUMBANGAN LEMO', 4435, 319, 23),
	(58084, 'KENJE', 4435, 319, 23),
	(58085, 'LAGI-AGI', 4435, 319, 23),
	(58086, 'LALIKO', 4435, 319, 23),
	(58087, 'LAMPOKO', 4435, 319, 23),
	(58088, 'LAPEO', 4435, 319, 23),
	(58089, 'ONGKO', 4435, 319, 23),
	(58090, 'PADANG', 4435, 319, 23),
	(58091, 'PADANG TIMUR', 4435, 319, 23),
	(58092, 'PANYAMPA', 4435, 319, 23),
	(58093, 'PAPPANG', 4435, 319, 23),
	(58094, 'PARAPPE', 4435, 319, 23),
	(58095, 'SUMARANG (SUMARRANG)', 4435, 319, 23),
	(58096, 'SURUANG', 4435, 319, 23),
	(58097, 'LEMBANG LEMBANG', 4436, 319, 23),
	(58098, 'LIMBORO', 4436, 319, 23),
	(58099, 'NAPO', 4436, 319, 23),
	(58100, 'PALECE', 4436, 319, 23),
	(58101, 'PENDULANGAN (PANDULANGAN)', 4436, 319, 23),
	(58102, 'RENGGEANG', 4436, 319, 23),
	(58103, 'SALARRI', 4436, 319, 23),
	(58104, 'SAMASUNDU', 4436, 319, 23),
	(58105, 'TANDASURA (TANDASSURA)', 4436, 319, 23),
	(58106, 'TANGAN BARU', 4436, 319, 23),
	(58107, 'TODANG TODANG', 4436, 319, 23),
	(58108, 'BARU', 4437, 319, 23),
	(58109, 'BATU PANGA', 4437, 319, 23),
	(58110, 'BATUPANGA DAALA', 4437, 319, 23),
	(58111, 'LUYO (LUJO)', 4437, 319, 23),
	(58112, 'MAMBU', 4437, 319, 23),
	(58113, 'MAPILLI BARAT', 4437, 319, 23),
	(58114, 'PUCCADI', 4437, 319, 23),
	(58115, 'PUSSUI', 4437, 319, 23),
	(58116, 'PUSSUI BARAT', 4437, 319, 23),
	(58117, 'SAMBALI WALI', 4437, 319, 23),
	(58118, 'TENGGELANG (TENGGELAN)', 4437, 319, 23),
	(58119, 'BEROANGIN', 4438, 319, 23),
	(58120, 'BONNE-BONNE', 4438, 319, 23),
	(58121, 'BONRA', 4438, 319, 23),
	(58122, 'BUKU', 4438, 319, 23),
	(58123, 'KURMA', 4438, 319, 23),
	(58124, 'LANDI KANUSUANG', 4438, 319, 23),
	(58125, 'MAPILLI', 4438, 319, 23),
	(58126, 'RAPPANG BARAT', 4438, 319, 23),
	(58127, 'RUMPA', 4438, 319, 23),
	(58128, 'SATTOKO', 4438, 319, 23),
	(58129, 'SEGERANG', 4438, 319, 23),
	(58130, 'UGI BARU', 4438, 319, 23),
	(58131, 'BARUMBUNG', 4439, 319, 23),
	(58132, 'BUNGA-BUNGA', 4439, 319, 23),
	(58133, 'INDO MAKKOMBONG (INDOMAKKONDANG)', 4439, 319, 23),
	(58134, 'MATAKALI', 4439, 319, 23),
	(58135, 'PASIANG', 4439, 319, 23),
	(58136, 'PATAMPANUA', 4439, 319, 23),
	(58137, 'TONRO LIMA', 4439, 319, 23),
	(58138, 'BA\\BA TAPUA', 4440, 319, 23),
	(58139, 'KATIMBANG', 4440, 319, 23),
	(58140, 'LILLI', 4440, 319, 23),
	(58141, 'MAMBU TAPUA', 4440, 319, 23),
	(58142, 'MATANGNGA', 4440, 319, 23),
	(58143, 'RANGOAN', 4440, 319, 23),
	(58144, 'TAPUA', 4440, 319, 23),
	(58145, 'DARMA', 4441, 319, 23),
	(58146, 'LANTORA', 4441, 319, 23),
	(58147, 'MADATTE', 4441, 319, 23),
	(58148, 'MANDING', 4441, 319, 23),
	(58149, 'PEKKABATA', 4441, 319, 23),
	(58150, 'POLEWALI', 4441, 319, 23),
	(58151, 'SULEWATANG', 4441, 319, 23),
	(58152, 'TAKATIDUNG', 4441, 319, 23),
	(58153, 'WATTANG', 4441, 319, 23),
	(58154, 'BANATO REJO', 4442, 319, 23),
	(58155, 'BATU', 4442, 319, 23),
	(58156, 'BUSSU', 4442, 319, 23),
	(58157, 'DAKKA', 4442, 319, 23),
	(58158, 'JAMBU MALEA', 4442, 319, 23),
	(58159, 'KALIMBUA', 4442, 319, 23),
	(58160, 'KURRAK', 4442, 319, 23),
	(58161, 'PALATTA', 4442, 319, 23),
	(58162, 'PELITAKAN', 4442, 319, 23),
	(58163, 'RAPPANG', 4442, 319, 23),
	(58164, 'RISO', 4442, 319, 23),
	(58165, 'TAPANGO', 4442, 319, 23),
	(58166, 'TAPANGO BARAT', 4442, 319, 23),
	(58167, 'TUTTULA', 4442, 319, 23),
	(58168, 'BATULAYA', 4443, 319, 23),
	(58169, 'GALUNG LOMBOK', 4443, 319, 23),
	(58170, 'KARAMA', 4443, 319, 23),
	(58171, 'LEKOPADIS', 4443, 319, 23),
	(58172, 'SEPA BATU', 4443, 319, 23),
	(58173, 'TANDUNG', 4443, 319, 23),
	(58174, 'TANGNGA-TANGNGA', 4443, 319, 23),
	(58175, 'TINAMBUNG', 4443, 319, 23),
	(58176, 'AMBO PADANG', 4444, 319, 23),
	(58177, 'ARABUA', 4444, 319, 23),
	(58178, 'BESOANGIN', 4444, 319, 23),
	(58179, 'BESOANGIN UTARA', 4444, 319, 23),
	(58180, 'PEBURRU', 4444, 319, 23),
	(58181, 'PIRIANG TAPIKO', 4444, 319, 23),
	(58182, 'PODA', 4444, 319, 23),
	(58183, 'POLLEWANI', 4444, 319, 23),
	(58184, 'RATTE', 4444, 319, 23),
	(58185, 'TALOBA', 4444, 319, 23),
	(58186, 'TARAMANU', 4444, 319, 23),
	(58187, 'TARAMANU TUA', 4444, 319, 23),
	(58188, 'TUBBI', 4444, 319, 23),
	(58189, 'ARJOSARI', 4445, 319, 23),
	(58190, 'BAKKA-BAKKA', 4445, 319, 23),
	(58191, 'BANUA BARU', 4445, 319, 23),
	(58192, 'BUMI MULYO', 4445, 319, 23),
	(58193, 'BUMIAYU', 4445, 319, 23),
	(58194, 'CAMPURJO', 4445, 319, 23),
	(58195, 'GELESO', 4445, 319, 23),
	(58196, 'KEBUNSARI', 4445, 319, 23),
	(58197, 'NEPO', 4445, 319, 23),
	(58198, 'SIDODADI', 4445, 319, 23),
	(58199, 'SIDOREJO', 4445, 319, 23),
	(58200, 'SUGIH WARAS', 4445, 319, 23),
	(58201, 'SUMBERJO', 4445, 319, 23),
	(58202, 'TUMPILING', 4445, 319, 23),
	(58203, 'KARATUANG', 4446, 320, 24),
	(58204, 'KAYU LOE', 4446, 320, 24),
	(58205, 'LAMALAKA', 4446, 320, 24),
	(58206, 'LEMBANG', 4446, 320, 24),
	(58207, 'LETTA', 4446, 320, 24),
	(58208, 'MALLILINGI', 4446, 320, 24),
	(58209, 'ONTO', 4446, 320, 24),
	(58210, 'PALLANTIKANG', 4446, 320, 24),
	(58211, 'TAPPANJENG', 4446, 320, 24),
	(58212, 'BONTO ATU', 4447, 320, 24),
	(58213, 'BONTO CINDE', 4447, 320, 24),
	(58214, 'BONTO JAI', 4447, 320, 24),
	(58215, 'BONTO JAYA', 4447, 320, 24),
	(58216, 'BONTO LANGKASA', 4447, 320, 24),
	(58217, 'BONTO LEBANG', 4447, 320, 24),
	(58218, 'BONTO LOE', 4447, 320, 24),
	(58219, 'BONTO MANAI', 4447, 320, 24),
	(58220, 'BONTO RITA', 4447, 320, 24),
	(58221, 'BONTO SALLUANG', 4447, 320, 24),
	(58222, 'BONTO SUNGGU', 4447, 320, 24),
	(58223, 'BARUA', 4448, 320, 24),
	(58224, 'KAMPALA', 4448, 320, 24),
	(58225, 'LONRONG', 4448, 320, 24),
	(58226, 'MAMAMPANG', 4448, 320, 24),
	(58227, 'MAPPILAWING', 4448, 320, 24),
	(58228, 'PA BENTENGAN', 4448, 320, 24),
	(58229, 'PA BUMBUNGAN', 4448, 320, 24),
	(58230, 'PARANGLOE', 4448, 320, 24),
	(58231, 'ULUGALUNG', 4448, 320, 24),
	(58232, 'BAJIMINASA', 4449, 320, 24),
	(58233, 'GANTARANGKEKE', 4449, 320, 24),
	(58234, 'KALOLING', 4449, 320, 24),
	(58235, 'LAYOA', 4449, 320, 24),
	(58236, 'TANAHLOE', 4449, 320, 24),
	(58237, 'TOMBOLO', 4449, 320, 24),
	(58238, 'BARUGA', 4450, 320, 24),
	(58239, 'BATU KARAENG', 4450, 320, 24),
	(58240, 'BIANGKEKE', 4450, 320, 24),
	(58241, 'BIANGLOE', 4450, 320, 24),
	(58242, 'BORONGLOE', 4450, 320, 24),
	(58243, 'LUMPANGAN', 4450, 320, 24),
	(58244, 'NIPA-NIPA', 4450, 320, 24),
	(58245, 'PAJUKUKANG', 4450, 320, 24),
	(58246, 'PAPANLOE', 4450, 320, 24),
	(58247, 'RAPPOA', 4450, 320, 24),
	(58248, 'BONTO BULAENG', 4451, 320, 24),
	(58249, 'BONTO KARAENG', 4451, 320, 24),
	(58250, 'BONTO MACCINI', 4451, 320, 24),
	(58251, 'BONTO MAJANNANG', 4451, 320, 24),
	(58252, 'BONTO MATENE', 4451, 320, 24),
	(58253, 'BONTO TIRO', 4451, 320, 24),
	(58254, 'BALUMBUNG', 4452, 320, 24),
	(58255, 'BANYORANG', 4452, 320, 24),
	(58256, 'BONTO TAPPALANG', 4452, 320, 24),
	(58257, 'BONTO-BONTOA', 4452, 320, 24),
	(58258, 'BONTOBUDDUNG', 4452, 320, 24),
	(58259, 'CAMPAGA', 4452, 320, 24),
	(58260, 'CIKORO', 4452, 320, 24),
	(58261, 'DATARA', 4452, 320, 24),
	(58262, 'ERENG-ERENG', 4452, 320, 24),
	(58263, 'GARING', 4452, 320, 24),
	(58264, 'LABBO', 4452, 320, 24),
	(58265, 'LEMBANG GANTARANGKEKE', 4452, 320, 24),
	(58266, 'MALAKAJI', 4452, 320, 24),
	(58267, 'PATTALLASSANG', 4452, 320, 24),
	(58268, 'PATTANETEANG', 4452, 320, 24),
	(58269, 'RAPPOALA', 4452, 320, 24),
	(58270, 'RAPPOLEMBA', 4452, 320, 24),
	(58271, 'TANETE', 4452, 320, 24),
	(58272, 'BONTO DAENG', 4453, 320, 24),
	(58273, 'BONTO LOJONG', 4453, 320, 24),
	(58274, 'BONTO MARANNU', 4453, 320, 24),
	(58275, 'BONTO RANNU', 4453, 320, 24),
	(58276, 'BONTO TALLASA', 4453, 320, 24),
	(58277, 'BONTO TANGNGA', 4453, 320, 24),
	(58278, 'ANABANUA', 4454, 321, 24),
	(58279, 'COPPO', 4454, 321, 24),
	(58280, 'GALUNG', 4454, 321, 24),
	(58281, 'MANGEMPANG', 4454, 321, 24),
	(58282, 'PALAKKA', 4454, 321, 24),
	(58283, 'SEPEE', 4454, 321, 24),
	(58284, 'SIAWUNG', 4454, 321, 24),
	(58285, 'SUMPANG BINANGAE', 4454, 321, 24),
	(58286, 'TOMPO', 4454, 321, 24),
	(58287, 'TUWUNG', 4454, 321, 24),
	(58288, 'BOJO', 4455, 321, 24),
	(58289, 'BOJO BARU', 4455, 321, 24),
	(58290, 'CILELLANG', 4455, 321, 24),
	(58291, 'KUPA', 4455, 321, 24),
	(58292, 'MALLAWA', 4455, 321, 24),
	(58293, 'MANUBA', 4455, 321, 24),
	(58294, 'NEPO', 4455, 321, 24),
	(58295, 'PALANRO', 4455, 321, 24),
	(58296, 'BACU-BACU', 4456, 321, 24),
	(58297, 'BULO-BULO', 4456, 321, 24),
	(58298, 'GATTARENG', 4456, 321, 24),
	(58299, 'JANGAN-JANGAN', 4456, 321, 24),
	(58300, 'MATTAPPAWALIE', 4456, 321, 24),
	(58301, 'PATTAPPA', 4456, 321, 24),
	(58302, 'PUJANANTING', 4456, 321, 24),
	(58303, 'AJAKKANG', 4457, 321, 24),
	(58304, 'BATUPUTE', 4457, 321, 24),
	(58305, 'KIRU-KIRU', 4457, 321, 24),
	(58306, 'LAWALLU', 4457, 321, 24),
	(58307, 'MANGKOSO', 4457, 321, 24),
	(58308, 'PACCEKKE', 4457, 321, 24),
	(58309, 'SIDDO', 4457, 321, 24),
	(58310, 'HARAPAN', 4458, 321, 24),
	(58311, 'KADING', 4458, 321, 24),
	(58312, 'LEMPANG', 4458, 321, 24),
	(58313, 'LIBURENG', 4458, 321, 24),
	(58314, 'LOMPO RIAJA', 4458, 321, 24),
	(58315, 'LOMPO TENGAH', 4458, 321, 24),
	(58316, 'MATTIROWALIE', 4458, 321, 24),
	(58317, 'CORAWALI', 4459, 321, 24),
	(58318, 'GARESSI', 4459, 321, 24),
	(58319, 'LALABATA', 4459, 321, 24),
	(58320, 'LALOLANG', 4459, 321, 24),
	(58321, 'LASITAE', 4459, 321, 24),
	(58322, 'LIPUKASI', 4459, 321, 24),
	(58323, 'PANCANA', 4459, 321, 24),
	(58324, 'PAO-PAO', 4459, 321, 24),
	(58325, 'TANETE', 4459, 321, 24),
	(58326, 'TELLUMPANUA', 4459, 321, 24),
	(58327, 'ALLAMUNGENG PATUE', 4460, 322, 24),
	(58328, 'AMESSANGENG', 4460, 322, 24),
	(58329, 'LABISSA', 4460, 322, 24),
	(58330, 'LEBBAE', 4460, 322, 24),
	(58331, 'LEPPANGENG', 4460, 322, 24),
	(58332, 'MANCIRI', 4460, 322, 24),
	(58333, 'OPO', 4460, 322, 24),
	(58334, 'PACCIRO', 4460, 322, 24),
	(58335, 'PINCENG PUTE', 4460, 322, 24),
	(58336, 'POMPANUA', 4460, 322, 24),
	(58337, 'POMPANUA RIATTANG (SALEWANGENG)', 4460, 322, 24),
	(58338, 'TELLE', 4460, 322, 24),
	(58339, 'TIMURUNG', 4460, 322, 24),
	(58340, 'WELADO', 4460, 322, 24),
	(58341, 'AJANG LALENG', 4461, 322, 24),
	(58342, 'AMALI RIATTANG', 4461, 322, 24),
	(58343, 'BENTENG TELLUE', 4461, 322, 24),
	(58344, 'BILA', 4461, 322, 24),
	(58345, 'LAPONRONG', 4461, 322, 24),
	(58346, 'LILI RIATTANG', 4461, 322, 24),
	(58347, 'MAMPOTU', 4461, 322, 24),
	(58348, 'MATTARO PURAE', 4461, 322, 24),
	(58349, 'TACIPONG', 4461, 322, 24),
	(58350, 'TASSIPI', 4461, 322, 24),
	(58351, 'TOCINNONG', 4461, 322, 24),
	(58352, 'ULAWENG RIAJA', 4461, 322, 24),
	(58353, 'WAEMPUBBUE', 4461, 322, 24),
	(58354, 'WAEMPUTTANGE', 4461, 322, 24),
	(58355, 'WELLULANG', 4461, 322, 24),
	(58356, 'ABBANUANG', 4462, 322, 24),
	(58357, 'AWO LAGADING', 4462, 322, 24),
	(58358, 'BULUMPAREE', 4462, 322, 24),
	(58359, 'CAKKE BONE', 4462, 322, 24),
	(58360, 'CAREBBU', 4462, 322, 24),
	(58361, 'CARI GADING', 4462, 322, 24),
	(58362, 'CUMPIGA', 4462, 322, 24),
	(58363, 'JALING', 4462, 322, 24),
	(58364, 'KADING', 4462, 322, 24),
	(58365, 'KAJUARA', 4462, 322, 24),
	(58366, 'LAPPO ASE', 4462, 322, 24),
	(58367, 'LETTEKKO', 4462, 322, 24),
	(58368, 'MACCOPE', 4462, 322, 24),
	(58369, 'MALLARI', 4462, 322, 24),
	(58370, 'MAPPALO ULAWENG', 4462, 322, 24),
	(58371, 'MATUJU', 4462, 322, 24),
	(58372, 'PACCING', 4462, 322, 24),
	(58373, 'UNRA', 4462, 322, 24),
	(58374, 'APALA', 4463, 322, 24),
	(58375, 'BACU', 4463, 322, 24),
	(58376, 'BAREBBO', 4463, 322, 24),
	(58377, 'CEMPANIGA', 4463, 322, 24),
	(58378, 'CINGKANG', 4463, 322, 24),
	(58379, 'CINNONG', 4463, 322, 24),
	(58380, 'CONGKO', 4463, 322, 24),
	(58381, 'CORAWALIE', 4463, 322, 24),
	(58382, 'KADING', 4463, 322, 24),
	(58383, 'KAJAOLALIDDONG', 4463, 322, 24),
	(58384, 'KAMPUNO', 4463, 322, 24),
	(58385, 'LAMPOKO', 4463, 322, 24),
	(58386, 'PARIPPUNG', 4463, 322, 24),
	(58387, 'SAMAELO', 4463, 322, 24),
	(58388, 'SUGIALE', 4463, 322, 24),
	(58389, 'TALUNGENG', 4463, 322, 24),
	(58390, 'WATU', 4463, 322, 24),
	(58391, 'WOLLANGI', 4463, 322, 24),
	(58392, 'BENGO', 4464, 322, 24),
	(58393, 'BULU ALLAPPORENGE', 4464, 322, 24),
	(58394, 'LILIRIAWANG', 4464, 322, 24),
	(58395, 'MATTARO PULI', 4464, 322, 24),
	(58396, 'MATTIRO WALIE', 4464, 322, 24),
	(58397, 'SAMAENRE', 4464, 322, 24),
	(58398, 'SELLI', 4464, 322, 24),
	(58399, 'TUNGKE', 4464, 322, 24),
	(58400, 'WALIMPONG', 4464, 322, 24),
	(58401, 'BANA', 4465, 322, 24),
	(58402, 'BONTOJAI', 4465, 322, 24),
	(58403, 'BULUSIRUA', 4465, 322, 24),
	(58404, 'ERE CINNONG', 4465, 322, 24),
	(58405, 'KAHU', 4465, 322, 24),
	(58406, 'LAMONCONG', 4465, 322, 24),
	(58407, 'LANGI', 4465, 322, 24),
	(58408, 'MATTIRO WALIE', 4465, 322, 24),
	(58409, 'PAMMUSURENG', 4465, 322, 24),
	(58410, 'PATTUKU', 4465, 322, 24),
	(58411, 'WATANGCANI', 4465, 322, 24),
	(58412, 'AJALLASSE', 4466, 322, 24),
	(58413, 'AWANG CENGRANA / CENRANA', 4466, 322, 24),
	(58414, 'BAJI PAMAI', 4466, 322, 24),
	(58415, 'CENRANA', 4466, 322, 24),
	(58416, 'CENRANA BARU', 4466, 322, 24),
	(58417, 'LABONGE (LEBBONGNGE)', 4466, 322, 24),
	(58418, 'LABOTTO', 4466, 322, 24),
	(58419, 'LABUAJA', 4466, 322, 24),
	(58420, 'LAIYA', 4466, 322, 24),
	(58421, 'LAONI', 4466, 322, 24),
	(58422, 'LATONRO', 4466, 322, 24),
	(58423, 'LEBBOTENGNGAE', 4466, 322, 24),
	(58424, 'LIMAM POCCOE', 4466, 322, 24),
	(58425, 'NAGA ULENG', 4466, 322, 24),
	(58426, 'PACUBBE', 4466, 322, 24),
	(58427, 'PALLAE', 4466, 322, 24),
	(58428, 'PALLIME', 4466, 322, 24),
	(58429, 'PANYIWI', 4466, 322, 24),
	(58430, 'PUSUNGNGE', 4466, 322, 24),
	(58431, 'ROMPE GADING', 4466, 322, 24),
	(58432, 'UJUNG TANAH', 4466, 322, 24),
	(58433, 'WATANG TA', 4466, 322, 24),
	(58434, 'WATU', 4466, 322, 24),
	(58435, 'ABBUMPUNGENG', 4467, 322, 24),
	(58436, 'AJANGPULU', 4467, 322, 24),
	(58437, 'ARASOE', 4467, 322, 24),
	(58438, 'AWO', 4467, 322, 24),
	(58439, 'CINENNUNG', 4467, 322, 24),
	(58440, 'KANCO', 4467, 322, 24),
	(58441, 'KAWERANG', 4467, 322, 24),
	(58442, 'LOMPU', 4467, 322, 24),
	(58443, 'PADANG LOANG', 4467, 322, 24),
	(58444, 'TANETE', 4467, 322, 24),
	(58445, 'TANETE HARAPAN', 4467, 322, 24),
	(58446, 'WELENRENG', 4467, 322, 24),
	(58447, 'CABBENG', 4468, 322, 24),
	(58448, 'KAMPOTI', 4468, 322, 24),
	(58449, 'LACCORI', 4468, 322, 24),
	(58450, 'LALLATANG', 4468, 322, 24),
	(58451, 'MARIO', 4468, 322, 24),
	(58452, 'MATAJANG', 4468, 322, 24),
	(58453, 'MELLE', 4468, 322, 24),
	(58454, 'PADACENGA (PADA CENGNGA)', 4468, 322, 24),
	(58455, 'PAKKASALO', 4468, 322, 24),
	(58456, 'PANYILI', 4468, 322, 24),
	(58457, 'PATTIRO', 4468, 322, 24),
	(58458, 'PRAJA MAJU', 4468, 322, 24),
	(58459, 'SAILONG', 4468, 322, 24),
	(58460, 'SANRANGENG', 4468, 322, 24),
	(58461, 'SOLO', 4468, 322, 24),
	(58462, 'TAWAROE', 4468, 322, 24),
	(58463, 'TEMPE', 4468, 322, 24),
	(58464, 'TOCINA', 4468, 322, 24),
	(58465, 'UJUNG', 4468, 322, 24),
	(58466, 'ULOE', 4468, 322, 24),
	(58467, 'UNNYI', 4468, 322, 24),
	(58468, 'WATANG PADACENGA (TURUMAME/PADACENGNGA)', 4468, 322, 24),
	(58469, 'ARALLAE', 4469, 322, 24),
	(58470, 'BALLE', 4469, 322, 24),
	(58471, 'BIRU', 4469, 322, 24),
	(58472, 'BONTO PADANG', 4469, 322, 24),
	(58473, 'CAKKELA', 4469, 322, 24),
	(58474, 'CAMMILO', 4469, 322, 24),
	(58475, 'CARIMA', 4469, 322, 24),
	(58476, 'CENRANA', 4469, 322, 24),
	(58477, 'HULO', 4469, 322, 24),
	(58478, 'LABUAJA', 4469, 322, 24),
	(58479, 'LALEPO', 4469, 322, 24),
	(58480, 'MAGGENRANG (MANGGENRANG)', 4469, 322, 24),
	(58481, 'MATAJANG', 4469, 322, 24),
	(58482, 'MATTOANGING', 4469, 322, 24),
	(58483, 'NUSA', 4469, 322, 24),
	(58484, 'PALAKKA', 4469, 322, 24),
	(58485, 'PALATTAE', 4469, 322, 24),
	(58486, 'PASAKA', 4469, 322, 24),
	(58487, 'SANREGO', 4469, 322, 24),
	(58488, 'TOMPONG PATU', 4469, 322, 24),
	(58489, 'ABBUMPUNGENG', 4470, 322, 24),
	(58490, 'ANCU', 4470, 322, 24),
	(58491, 'ANGKUE / ANGKU', 4470, 322, 24),
	(58492, 'AWANG TANGKA', 4470, 322, 24),
	(58493, 'BUARENG', 4470, 322, 24),
	(58494, 'BULU TANAH', 4470, 322, 24),
	(58495, 'GONA', 4470, 322, 24),
	(58496, 'KALERO', 4470, 322, 24),
	(58497, 'LAPPA BOSSE', 4470, 322, 24),
	(58498, 'LEMO', 4470, 322, 24),
	(58499, 'MALLAHAE', 4470, 322, 24),
	(58500, 'MASSANGKAE', 4470, 322, 24),
	(58501, 'PADAELO', 4470, 322, 24),
	(58502, 'POLEWALI', 4470, 322, 24),
	(58503, 'PUDE', 4470, 322, 24),
	(58504, 'RAJA', 4470, 322, 24),
	(58505, 'TARASU', 4470, 322, 24),
	(58506, 'WAETUWO (WAETUO)', 4470, 322, 24),
	(58507, 'BARAKKAE', 4471, 322, 24),
	(58508, 'BARUGAE', 4471, 322, 24),
	(58509, 'LALEBATA', 4471, 322, 24),
	(58510, 'MAMMINASAE', 4471, 322, 24),
	(58511, 'MASSENRENPULU', 4471, 322, 24),
	(58512, 'MATTAMPA BULU', 4471, 322, 24),
	(58513, 'MATTAMPA WALIE', 4471, 322, 24),
	(58514, 'PADAELO', 4471, 322, 24),
	(58515, 'POLEONRO', 4471, 322, 24),
	(58516, 'SEBERANG', 4471, 322, 24),
	(58517, 'SENGENG PALIE', 4471, 322, 24),
	(58518, 'TURUCINNAE', 4471, 322, 24),
	(58519, 'LILI RIATTANG', 4472, 322, 24),
	(58520, 'MATTAMPA WALIE', 4472, 322, 24),
	(58521, 'PATANGKAI', 4472, 322, 24),
	(58522, 'PATTUKU LIMPOE', 4472, 322, 24),
	(58523, 'SENGENG PALIE', 4472, 322, 24),
	(58524, 'TENRI PAKKUA', 4472, 322, 24),
	(58525, 'TONRONGE', 4472, 322, 24),
	(58526, 'UJUNG LAMURU', 4472, 322, 24),
	(58527, 'WAEKECCEE', 4472, 322, 24),
	(58528, 'BARINGENG', 4473, 322, 24),
	(58529, 'BINUANG', 4473, 322, 24),
	(58530, 'BUNE', 4473, 322, 24),
	(58531, 'CEPPAGA', 4473, 322, 24),
	(58532, 'LABURASSENG (LABURASSE)', 4473, 322, 24),
	(58533, 'MALLINRUNG', 4473, 322, 24),
	(58534, 'MARIO', 4473, 322, 24),
	(58535, 'MATTIRO BULU', 4473, 322, 24),
	(58536, 'MATTIRO DECENG', 4473, 322, 24),
	(58537, 'MATTIRO WALIE', 4473, 322, 24),
	(58538, 'PITUMPIDANGE', 4473, 322, 24),
	(58539, 'POLEONRO', 4473, 322, 24),
	(58540, 'POLEWALI', 4473, 322, 24),
	(58541, 'PONRE-PONRE', 4473, 322, 24),
	(58542, 'SUWA', 4473, 322, 24),
	(58543, 'SWADAYA', 4473, 322, 24),
	(58544, 'TANABATUE', 4473, 322, 24),
	(58545, 'TAPPALE', 4473, 322, 24),
	(58546, 'TOMPOBULU', 4473, 322, 24),
	(58547, 'WANUAWARU', 4473, 322, 24),
	(58548, 'BAINANG', 4474, 322, 24),
	(58549, 'CINENNUNG', 4474, 322, 24),
	(58550, 'LEMOAPE', 4474, 322, 24),
	(58551, 'MADURI', 4474, 322, 24),
	(58552, 'MATTANETE BUA', 4474, 322, 24),
	(58553, 'MELLE', 4474, 322, 24),
	(58554, 'MICO', 4474, 322, 24),
	(58555, 'PANYILI', 4474, 322, 24),
	(58556, 'PASEMPE', 4474, 322, 24),
	(58557, 'PASSIPPO', 4474, 322, 24),
	(58558, 'SIAME', 4474, 322, 24),
	(58559, 'TANAH TENGA', 4474, 322, 24),
	(58560, 'TIRONG', 4474, 322, 24),
	(58561, 'URENG', 4474, 322, 24),
	(58562, 'USA', 4474, 322, 24),
	(58563, 'BATULAPPA', 4475, 322, 24),
	(58564, 'BULU ULAWENG', 4475, 322, 24),
	(58565, 'LATELLANG', 4475, 322, 24),
	(58566, 'MADDANRENGPULU', 4475, 322, 24),
	(58567, 'MASAGO', 4475, 322, 24),
	(58568, 'MASSILA', 4475, 322, 24),
	(58569, 'PACCING', 4475, 322, 24),
	(58570, 'PATIMPENG', 4475, 322, 24),
	(58571, 'PATIONGI', 4475, 322, 24),
	(58572, 'TALABANGI', 4475, 322, 24),
	(58573, 'BOLLI', 4476, 322, 24),
	(58574, 'MAPPESANGKA', 4476, 322, 24),
	(58575, 'MATTAMPAE', 4476, 322, 24),
	(58576, 'PATTIMPA', 4476, 322, 24),
	(58577, 'POLEONRO', 4476, 322, 24),
	(58578, 'SALAMPE', 4476, 322, 24),
	(58579, 'SALEBBA', 4476, 322, 24),
	(58580, 'TELLU BOCCOE', 4476, 322, 24),
	(58581, 'TURU ADAE', 4476, 322, 24),
	(58582, 'BELLU', 4477, 322, 24),
	(58583, 'GATTARENG', 4477, 322, 24),
	(58584, 'MALIMONGENG', 4477, 322, 24),
	(58585, 'MANERA', 4477, 322, 24),
	(58586, 'MAPPATOBA', 4477, 322, 24),
	(58587, 'PANCAITANA', 4477, 322, 24),
	(58588, 'TEBBA', 4477, 322, 24),
	(58589, 'ULU BALANG', 4477, 322, 24),
	(58590, 'AJANG PULU', 4478, 322, 24),
	(58591, 'BALIENG TOA', 4478, 322, 24),
	(58592, 'BULIE', 4478, 322, 24),
	(58593, 'CINNONG', 4478, 322, 24),
	(58594, 'KALIBONG', 4478, 322, 24),
	(58595, 'LETTE TANAH', 4478, 322, 24),
	(58596, 'MABBIRING', 4478, 322, 24),
	(58597, 'MALLUSE TASI', 4478, 322, 24),
	(58598, 'MANAJENG', 4478, 322, 24),
	(58599, 'MAROANGING', 4478, 322, 24),
	(58600, 'MASSENRENGPULU', 4478, 322, 24),
	(58601, 'PAKKASALO', 4478, 322, 24),
	(58602, 'PASAKA', 4478, 322, 24),
	(58603, 'PATTIRO BAJO', 4478, 322, 24),
	(58604, 'PATTIRO RIOLO', 4478, 322, 24),
	(58605, 'PATTIRO SOMPE', 4478, 322, 24),
	(58606, 'POLEWALI', 4478, 322, 24),
	(58607, 'SUMPANG MINANGAE', 4478, 322, 24),
	(58608, 'TADANG PALIE', 4478, 322, 24),
	(58609, 'TUNRENG TELLUE', 4478, 322, 24),
	(58610, 'BIRU', 4479, 322, 24),
	(58611, 'BUKAKA', 4479, 322, 24),
	(58612, 'MANURUNGNGE', 4479, 322, 24),
	(58613, 'MASUMPU', 4479, 322, 24),
	(58614, 'PAPPOLO', 4479, 322, 24),
	(58615, 'TA', 4479, 322, 24),
	(58616, 'WALENNAE', 4479, 322, 24),
	(58617, 'WATAMPONE', 4479, 322, 24),
	(58618, 'BULU TEMPE', 4480, 322, 24),
	(58619, 'JEPPEE', 4480, 322, 24),
	(58620, 'MACANANG', 4480, 322, 24),
	(58621, 'MACEGE', 4480, 322, 24),
	(58622, 'MAJANG', 4480, 322, 24),
	(58623, 'MATTIRO WALIE', 4480, 322, 24),
	(58624, 'POLEWALI', 4480, 322, 24),
	(58625, 'WATANG PALAKKA', 4480, 322, 24),
	(58626, 'BAJOE', 4481, 322, 24),
	(58627, 'CELLU', 4481, 322, 24),
	(58628, 'LONRAE', 4481, 322, 24),
	(58629, 'PALLETTE', 4481, 322, 24),
	(58630, 'PANYULA', 4481, 322, 24),
	(58631, 'TIBOJONG', 4481, 322, 24),
	(58632, 'TORO', 4481, 322, 24),
	(58633, 'WAETUO', 4481, 322, 24),
	(58634, 'AMPARITA', 4482, 322, 24),
	(58635, 'ARATENG', 4482, 322, 24),
	(58636, 'BATU PUTIH', 4482, 322, 24),
	(58637, 'BAULA', 4482, 322, 24),
	(58638, 'BONTO MASUNGGU', 4482, 322, 24),
	(58639, 'BUA', 4482, 322, 24),
	(58640, 'ERABARU', 4482, 322, 24),
	(58641, 'GAYA BARU', 4482, 322, 24),
	(58642, 'KALOBBA', 4482, 322, 24),
	(58643, 'LAGORI', 4482, 322, 24),
	(58644, 'LEMBANG LOHE', 4482, 322, 24),
	(58645, 'MANNANTI', 4482, 322, 24),
	(58646, 'MASSAILE', 4482, 322, 24),
	(58647, 'MASSEPE', 4482, 322, 24),
	(58648, 'PAJALELE', 4482, 322, 24),
	(58649, 'PALLAWA', 4482, 322, 24),
	(58650, 'PATTONGKO', 4482, 322, 24),
	(58651, 'POLEWALI', 4482, 322, 24),
	(58652, 'POLEWALI', 4482, 322, 24),
	(58653, 'SADAR', 4482, 322, 24),
	(58654, 'SAMAENRE', 4482, 322, 24),
	(58655, 'SAMATURUE', 4482, 322, 24),
	(58656, 'SAOTENGAH', 4482, 322, 24),
	(58657, 'SUKAMAJU', 4482, 322, 24),
	(58658, 'TAPONG', 4482, 322, 24),
	(58659, 'TELLANGKERE', 4482, 322, 24),
	(58660, 'TELLU LIMPOE', 4482, 322, 24),
	(58661, 'TEPPO', 4482, 322, 24),
	(58662, 'TETEAJI', 4482, 322, 24),
	(58663, 'TODANG PULU', 4482, 322, 24),
	(58664, 'TONDONG', 4482, 322, 24),
	(58665, 'AJJALIRENG', 4483, 322, 24),
	(58666, 'ITTERUNG', 4483, 322, 24),
	(58667, 'LAMURU', 4483, 322, 24),
	(58668, 'LANCA', 4483, 322, 24),
	(58669, 'LAPPAE', 4483, 322, 24),
	(58670, 'LEA', 4483, 322, 24),
	(58671, 'MATTOANGING', 4483, 322, 24),
	(58672, 'OTTING', 4483, 322, 24),
	(58673, 'PADAIDI', 4483, 322, 24),
	(58674, 'PALONGKI', 4483, 322, 24),
	(58675, 'PATANGNGA', 4483, 322, 24),
	(58676, 'PONGKA', 4483, 322, 24),
	(58677, 'SIJELLING', 4483, 322, 24),
	(58678, 'TAJONG', 4483, 322, 24),
	(58679, 'TOKASENG', 4483, 322, 24),
	(58680, 'ULO', 4483, 322, 24),
	(58681, 'WAJI', 4483, 322, 24),
	(58682, 'BACU', 4484, 322, 24),
	(58683, 'BICCOING', 4484, 322, 24),
	(58684, 'BONEPUTE', 4484, 322, 24),
	(58685, 'BULU-BULU', 4484, 322, 24),
	(58686, 'GARECCING', 4484, 322, 24),
	(58687, 'LIBURENG', 4484, 322, 24),
	(58688, 'MUARA', 4484, 322, 24),
	(58689, 'PADATUO', 4484, 322, 24),
	(58690, 'RAPPA', 4484, 322, 24),
	(58691, 'SAMAENRE', 4484, 322, 24),
	(58692, 'UJUNGE', 4484, 322, 24),
	(58693, 'CANI SIRENRENG', 4485, 322, 24),
	(58694, 'CINNONG', 4485, 322, 24),
	(58695, 'GALUNG', 4485, 322, 24),
	(58696, 'JOMPIE', 4485, 322, 24),
	(58697, 'LAMAKKARASENG', 4485, 322, 24),
	(58698, 'LILINA AJANGALE', 4485, 322, 24),
	(58699, 'MANURUNGE', 4485, 322, 24),
	(58700, 'MULA MENREE', 4485, 322, 24),
	(58701, 'PALLAWA RUKKA', 4485, 322, 24),
	(58702, 'SAPPE WALIE', 4485, 322, 24),
	(58703, 'TADANG PALIE', 4485, 322, 24),
	(58704, 'TEA MALALA', 4485, 322, 24),
	(58705, 'TEAMUSU', 4485, 322, 24),
	(58706, 'TIMUSU', 4485, 322, 24),
	(58707, 'ULAWENG CINNONG', 4485, 322, 24),
	(58708, 'ARA', 4486, 323, 24),
	(58709, 'BENJALA', 4486, 323, 24),
	(58710, 'BIRA', 4486, 323, 24),
	(58711, 'DARUBIAH', 4486, 323, 24),
	(58712, 'LEMBANNA', 4486, 323, 24),
	(58713, 'SAPOLOHE', 4486, 323, 24),
	(58714, 'TANAH BERU', 4486, 323, 24),
	(58715, 'TANAH LEMO', 4486, 323, 24),
	(58716, 'BATANG', 4487, 323, 24),
	(58717, 'BONTO BARUA', 4487, 323, 24),
	(58718, 'BONTO BULAENG', 4487, 323, 24),
	(58719, 'BONTO MARANNU', 4487, 323, 24),
	(58720, 'BONTO TANGNGA', 4487, 323, 24),
	(58721, 'BUHUNG BUNDANG', 4487, 323, 24),
	(58722, 'CARAMMING', 4487, 323, 24),
	(58723, 'DWI TIRO', 4487, 323, 24),
	(58724, 'EKATIRO', 4487, 323, 24),
	(58725, 'LAMANDA', 4487, 323, 24),
	(58726, 'PAKUBALAHO', 4487, 323, 24),
	(58727, 'TAMALANREA', 4487, 323, 24),
	(58728, 'TRITIRO', 4487, 323, 24),
	(58729, 'BALANG PESOANG', 4488, 323, 24),
	(58730, 'BALANG TAROANG', 4488, 323, 24),
	(58731, 'BALLASARAJA', 4488, 323, 24),
	(58732, 'BARUGA RIATTANG', 4488, 323, 24),
	(58733, 'BARUGAE', 4488, 323, 24),
	(58734, 'BATULOHE', 4488, 323, 24),
	(58735, 'BONTO BULAENG', 4488, 323, 24),
	(58736, 'BONTO MINASA', 4488, 323, 24),
	(58737, 'BONTOMANGIRING', 4488, 323, 24),
	(58738, 'BULO-BULO', 4488, 323, 24),
	(58739, 'JAWI - JAWI', 4488, 323, 24),
	(58740, 'JOJJOLO', 4488, 323, 24),
	(58741, 'KAMBUNO', 4488, 323, 24),
	(58742, 'SALASSAE', 4488, 323, 24),
	(58743, 'SAPO BONTO', 4488, 323, 24),
	(58744, 'TANETE', 4488, 323, 24),
	(58745, 'TIBONA', 4488, 323, 24),
	(58746, 'BAROMBONG', 4489, 323, 24),
	(58747, 'BENTENG GATARENG', 4489, 323, 24),
	(58748, 'BENTENG MALEWANG', 4489, 323, 24),
	(58749, 'BIALO', 4489, 323, 24),
	(58750, 'BONTO MACINNA', 4489, 323, 24),
	(58751, 'BONTO SUNGGU', 4489, 323, 24),
	(58752, 'BONTOMASILA', 4489, 323, 24),
	(58753, 'BONTONYELENG', 4489, 323, 24),
	(58754, 'BONTORAJA', 4489, 323, 24),
	(58755, 'BUKIT HARAPAN', 4489, 323, 24),
	(58756, 'BUKIT TINGGI', 4489, 323, 24),
	(58757, 'DAMPANG', 4489, 323, 24),
	(58758, 'GATTARENG', 4489, 323, 24),
	(58759, 'JALANJANG', 4489, 323, 24),
	(58760, 'MARIORENNU', 4489, 323, 24),
	(58761, 'MATTEKKO', 4489, 323, 24),
	(58762, 'PADANG', 4489, 323, 24),
	(58763, 'PAENRE LOMPOE', 4489, 323, 24),
	(58764, 'PALAMBARAE', 4489, 323, 24),
	(58765, 'POLEWALI', 4489, 323, 24),
	(58766, 'TACCORONG', 4489, 323, 24),
	(58767, 'BONTO KAMASE', 4490, 323, 24),
	(58768, 'BORONG', 4490, 323, 24),
	(58769, 'GUNTURU', 4490, 323, 24),
	(58770, 'KARASSING', 4490, 323, 24),
	(58771, 'PATARO', 4490, 323, 24),
	(58772, 'SINGA', 4490, 323, 24),
	(58773, 'TANUNTUNG', 4490, 323, 24),
	(58774, 'TUGONDENG', 4490, 323, 24),
	(58775, 'BATUNILAMUNG', 4491, 323, 24),
	(58776, 'BONTO BAJI', 4491, 323, 24),
	(58777, 'BONTO BIRAENG', 4491, 323, 24),
	(58778, 'BONTORANNU', 4491, 323, 24),
	(58779, 'LAIKANG', 4491, 323, 24),
	(58780, 'LEMBANG', 4491, 323, 24),
	(58781, 'LEMBANGLOHE', 4491, 323, 24),
	(58782, 'LEMBANNA', 4491, 323, 24),
	(58783, 'LOLISANG', 4491, 323, 24),
	(58784, 'MALLELENG', 4491, 323, 24),
	(58785, 'MATTOANGING', 4491, 323, 24),
	(58786, 'PANTAMA', 4491, 323, 24),
	(58787, 'PATTIROANG', 4491, 323, 24),
	(58788, 'POSSI TANAH', 4491, 323, 24),
	(58789, 'SANGKALA', 4491, 323, 24),
	(58790, 'SAPANANG', 4491, 323, 24),
	(58791, 'TAMBANGAN', 4491, 323, 24),
	(58792, 'TANAH JAYA', 4491, 323, 24),
	(58793, 'TANAH TOWA', 4491, 323, 24),
	(58794, 'ANRIHUA', 4492, 323, 24),
	(58795, 'BALIBO', 4492, 323, 24),
	(58796, 'BENTENG PALIOI', 4492, 323, 24),
	(58797, 'BORONG RAPPOA', 4492, 323, 24),
	(58798, 'GARUNTUNGAN', 4492, 323, 24),
	(58799, 'KAHAYYA', 4492, 323, 24),
	(58800, 'KINDANG', 4492, 323, 24),
	(58801, 'MATTIROWALIE', 4492, 323, 24),
	(58802, 'ORO GADING', 4492, 323, 24),
	(58803, 'SIPAENRE', 4492, 323, 24),
	(58804, 'SOMBA PALIOLI', 4492, 323, 24),
	(58805, 'SOPA', 4492, 323, 24),
	(58806, 'TAMAONA', 4492, 323, 24),
	(58807, 'ANRANG', 4493, 323, 24),
	(58808, 'BAJIMINASA', 4493, 323, 24),
	(58809, 'BATUKAROPA', 4493, 323, 24),
	(58810, 'BONTO MATENE', 4493, 323, 24),
	(58811, 'BONTOBANGUN', 4493, 323, 24),
	(58812, 'BONTOHARU', 4493, 323, 24),
	(58813, 'BONTOLOHE', 4493, 323, 24),
	(58814, 'BONTOMANAI', 4493, 323, 24),
	(58815, 'BULOLOHE', 4493, 323, 24),
	(58816, 'KARAMA', 4493, 323, 24),
	(58817, 'PALAMPANG', 4493, 323, 24),
	(58818, 'PANGALLOANG', 4493, 323, 24),
	(58819, 'SWATANI', 4493, 323, 24),
	(58820, 'TANAH HARAPAN', 4493, 323, 24),
	(58821, 'TOPANDA', 4493, 323, 24),
	(58822, 'BENTENGNGE', 4494, 323, 24),
	(58823, 'BINTARORE', 4494, 323, 24),
	(58824, 'CAILE', 4494, 323, 24),
	(58825, 'ELA-ELA', 4494, 323, 24),
	(58826, 'KALUMEME', 4494, 323, 24),
	(58827, 'KASIMPURENG', 4494, 323, 24),
	(58828, 'L O K A (LOKA)', 4494, 323, 24),
	(58829, 'TANAH KONGKONG', 4494, 323, 24),
	(58830, 'TERANG-TERANG', 4494, 323, 24),
	(58831, 'BALLEANGING (BALLEANGIN)', 4495, 323, 24),
	(58832, 'BALONG', 4495, 323, 24),
	(58833, 'BIJAWANG', 4495, 323, 24),
	(58834, 'DANNUANG', 4495, 323, 24),
	(58835, 'GARANTA', 4495, 323, 24),
	(58836, 'LONRONG', 4495, 323, 24),
	(58837, 'MANJALLING', 4495, 323, 24),
	(58838, 'MANYAMPA', 4495, 323, 24),
	(58839, 'PACCARAMENGANG', 4495, 323, 24),
	(58840, 'PADANG LOANG', 4495, 323, 24),
	(58841, 'SALEMBA', 4495, 323, 24),
	(58842, 'SEPPANG', 4495, 323, 24),
	(58843, 'TAMATTO', 4495, 323, 24),
	(58844, 'BOLANG', 4496, 324, 24),
	(58845, 'BUNTU SUGI', 4496, 324, 24),
	(58846, 'KALOSI', 4496, 324, 24),
	(58847, 'KAMBIOLANGI', 4496, 324, 24),
	(58848, 'MATA ALLO', 4496, 324, 24),
	(58849, 'PANA', 4496, 324, 24),
	(58850, 'SUMILLAN', 4496, 324, 24),
	(58851, 'TAULO', 4496, 324, 24),
	(58852, 'BAMBA PUANG', 4497, 324, 24),
	(58853, 'BATU NONI', 4497, 324, 24),
	(58854, 'BUBUN LAMBA', 4497, 324, 24),
	(58855, 'LAKAWAN', 4497, 324, 24),
	(58856, 'MAMPU', 4497, 324, 24),
	(58857, 'MANDATTE', 4497, 324, 24),
	(58858, 'MATARAN', 4497, 324, 24),
	(58859, 'PEKALOBEAN', 4497, 324, 24),
	(58860, 'SALU DEWATA', 4497, 324, 24),
	(58861, 'SARURAN', 4497, 324, 24),
	(58862, 'SIAMBO', 4497, 324, 24),
	(58863, 'SINGKI', 4497, 324, 24),
	(58864, 'TAMPO', 4497, 324, 24),
	(58865, 'TANETE', 4497, 324, 24),
	(58866, 'TINDALLUN', 4497, 324, 24),
	(58867, 'BALLA', 4498, 324, 24),
	(58868, 'BANTI', 4498, 324, 24),
	(58869, 'BARAKA', 4498, 324, 24),
	(58870, 'BONE BONE', 4498, 324, 24),
	(58871, 'BONTONGAN', 4498, 324, 24),
	(58872, 'JANGGURARA', 4498, 324, 24),
	(58873, 'KADINGEH', 4498, 324, 24),
	(58874, 'KENDENAN', 4498, 324, 24),
	(58875, 'PANDUNG BATU', 4498, 324, 24),
	(58876, 'PARINDING', 4498, 324, 24),
	(58877, 'PEPANDUNGAN', 4498, 324, 24),
	(58878, 'PERANGIAN (PARANGIAN)', 4498, 324, 24),
	(58879, 'SALUKANAN', 4498, 324, 24),
	(58880, 'TIRO WALI', 4498, 324, 24),
	(58881, 'TOMENAWA', 4498, 324, 24),
	(58882, 'BAROKO', 4499, 324, 24),
	(58883, 'BENTENG ALLA', 4499, 324, 24),
	(58884, 'BENTENG ALLA UTARA', 4499, 324, 24),
	(58885, 'PATONGLOAN', 4499, 324, 24),
	(58886, 'TONGKO', 4499, 324, 24),
	(58887, 'BARUKA', 4500, 324, 24),
	(58888, 'BULO', 4500, 324, 24),
	(58889, 'BUNGIN', 4500, 324, 24),
	(58890, 'SAWITO', 4500, 324, 24),
	(58891, 'TALLANG RILAU', 4500, 324, 24),
	(58892, 'WIRA KARYA', 4500, 324, 24),
	(58893, 'BUNTU MONDONG', 4501, 324, 24),
	(58894, 'ERAN BATU', 4501, 324, 24),
	(58895, 'LANGDA', 4501, 324, 24),
	(58896, 'LATIMOJONG', 4501, 324, 24),
	(58897, 'LEDAN', 4501, 324, 24),
	(58898, 'LUNJEN', 4501, 324, 24),
	(58899, 'PASUI', 4501, 324, 24),
	(58900, 'POTOK ULLIN', 4501, 324, 24),
	(58901, 'CENDANA', 4502, 324, 24),
	(58902, 'KARRANG', 4502, 324, 24),
	(58903, 'LEBANG', 4502, 324, 24),
	(58904, 'MALALIN', 4502, 324, 24),
	(58905, 'PINANG', 4502, 324, 24),
	(58906, 'PUNDILEMO', 4502, 324, 24),
	(58907, 'TAULAN', 4502, 324, 24),
	(58908, 'BUNTU BARANA', 4503, 324, 24),
	(58909, 'BUNTU PEMA', 4503, 324, 24),
	(58910, 'CURIO', 4503, 324, 24),
	(58911, 'MANDALAN', 4503, 324, 24),
	(58912, 'MEKKALE (MEKKALAK)', 4503, 324, 24),
	(58913, 'PABALORAN', 4503, 324, 24),
	(58914, 'PAROMBEAN', 4503, 324, 24),
	(58915, 'SALASSA', 4503, 324, 24),
	(58916, 'SANGLEPONGAN', 4503, 324, 24),
	(58917, 'SUMBANG', 4503, 324, 24),
	(58918, 'TALLUNGURA', 4503, 324, 24),
	(58919, 'BUTTU BATU', 4504, 324, 24),
	(58920, 'CEMBA', 4504, 324, 24),
	(58921, 'GALONTA', 4504, 324, 24),
	(58922, 'JUPPANDANG', 4504, 324, 24),
	(58923, 'KALUPPINI', 4504, 324, 24),
	(58924, 'KARUENG', 4504, 324, 24),
	(58925, 'LEMBANG', 4504, 324, 24),
	(58926, 'LEORAN', 4504, 324, 24),
	(58927, 'LEWAJA', 4504, 324, 24),
	(58928, 'PUSERREN', 4504, 324, 24),
	(58929, 'RANGA', 4504, 324, 24),
	(58930, 'ROSSOAN', 4504, 324, 24),
	(58931, 'TALLU BAMBA', 4504, 324, 24),
	(58932, 'TEMBAN', 4504, 324, 24),
	(58933, 'TOBALU', 4504, 324, 24),
	(58934, 'TOKKONAN', 4504, 324, 24),
	(58935, 'TUARA', 4504, 324, 24),
	(58936, 'TUNGKA', 4504, 324, 24),
	(58937, 'BANGKALA', 4505, 324, 24),
	(58938, 'BARINGIN', 4505, 324, 24),
	(58939, 'BATU MILA', 4505, 324, 24),
	(58940, 'BOIYA', 4505, 324, 24),
	(58941, 'BOTO MALANGGA', 4505, 324, 24),
	(58942, 'KALUPPANG', 4505, 324, 24),
	(58943, 'LABUKU', 4505, 324, 24),
	(58944, 'LEBANI', 4505, 324, 24),
	(58945, 'LIMBUANG', 4505, 324, 24),
	(58946, 'MANGKAWANI (MENGKAWANI)', 4505, 324, 24),
	(58947, 'MATAJANG', 4505, 324, 24),
	(58948, 'ONGKO', 4505, 324, 24),
	(58949, 'PALADANG', 4505, 324, 24),
	(58950, 'PALAKKA', 4505, 324, 24),
	(58951, 'PARIWANG', 4505, 324, 24),
	(58952, 'PASANG', 4505, 324, 24),
	(58953, 'PATONDON SALU', 4505, 324, 24),
	(58954, 'PUNCAK HARAPAN', 4505, 324, 24),
	(58955, 'SALO DUA', 4505, 324, 24),
	(58956, 'TANETE', 4505, 324, 24),
	(58957, 'TAPONG', 4505, 324, 24),
	(58958, 'TUNCUNG', 4505, 324, 24),
	(58959, 'BONTO', 4506, 324, 24),
	(58960, 'BUNTU BATUAN', 4506, 324, 24),
	(58961, 'DULANG', 4506, 324, 24),
	(58962, 'KOLAI', 4506, 324, 24),
	(58963, 'MALUA', 4506, 324, 24),
	(58964, 'RANTE MARIO', 4506, 324, 24),
	(58965, 'TALLUNG TONDOK', 4506, 324, 24),
	(58966, 'TANGRU', 4506, 324, 24),
	(58967, 'BATU KEDE', 4507, 324, 24),
	(58968, 'BUNTU SARONG', 4507, 324, 24),
	(58969, 'MASALLE', 4507, 324, 24),
	(58970, 'MUNDAN', 4507, 324, 24),
	(58971, 'RAMPUNAN', 4507, 324, 24),
	(58972, 'TONGKONAN BASSE', 4507, 324, 24),
	(58973, 'BONE', 4508, 325, 24),
	(58974, 'BONTOSUNGGU', 4508, 325, 24),
	(58975, 'KALEBAJENG', 4508, 325, 24),
	(58976, 'LEMPANGANG', 4508, 325, 24),
	(58977, 'LIMBUNG', 4508, 325, 24),
	(58978, 'MACCINIBAJI', 4508, 325, 24),
	(58979, 'MARADEKAYA', 4508, 325, 24),
	(58980, 'MATA ALLO', 4508, 325, 24),
	(58981, 'PABENTENGAN', 4508, 325, 24),
	(58982, 'PANCIRO', 4508, 325, 24),
	(58983, 'PANYANGKALANG', 4508, 325, 24),
	(58984, 'PARAIKATE (PARAIKATTE)', 4508, 325, 24),
	(58985, 'TANGKEBAJENG', 4508, 325, 24),
	(58986, 'TUBAJENG', 4508, 325, 24),
	(58987, 'BONTOMANAI', 4509, 325, 24),
	(58988, 'BORIMATANGKASA', 4509, 325, 24),
	(58989, 'GENTUNGANG', 4509, 325, 24),
	(58990, 'KALEMANDALLE', 4509, 325, 24),
	(58991, 'MANDALLE (MANDALE)', 4509, 325, 24),
	(58992, 'MANJALLING', 4509, 325, 24),
	(58993, 'TANABANGKA', 4509, 325, 24),
	(58994, 'BENTENG SOMBA OPU', 4510, 325, 24),
	(58995, 'BIRINGNGALLA', 4510, 325, 24),
	(58996, 'KANJILO', 4510, 325, 24),
	(58997, 'LEMBANGPARANG', 4510, 325, 24),
	(58998, 'MONCOBALANG', 4510, 325, 24),
	(58999, 'TAMANNYELENG', 4510, 325, 24),
	(59000, 'TINGGIMAE', 4510, 325, 24)
	`)
	if err != nil {
		panic(err)
	}
	fmt.Println("insert villige 58")
}
