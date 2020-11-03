package kel20

import (
	"fmt"

	"github.com/danangkonang/ceodeaja-go/config"
)

func Kel26() {
	db := config.Connect()
	_, err := db.Exec(`INSERT INTO villages (id,village,sub_district_id,regency_id,province_id) VALUES
	(26001, 'BESAN', 1973, 108, 7),
	(26002, 'DAWAN KALER', 1973, 108, 7),
	(26003, 'DAWAN KLOD', 1973, 108, 7),
	(26004, 'GUNAKSA', 1973, 108, 7),
	(26005, 'KAMPUNG KUSAMBA', 1973, 108, 7),
	(26006, 'KUSAMBA', 1973, 108, 7),
	(26007, 'PAKSEBALI', 1973, 108, 7),
	(26008, 'PESINGGAHAN', 1973, 108, 7),
	(26009, 'PIKAT', 1973, 108, 7),
	(26010, 'SAMPALAN KLOD', 1973, 108, 7),
	(26011, 'SAMPALAN TENGAH', 1973, 108, 7),
	(26012, 'SULANG', 1973, 108, 7),
	(26013, 'AKAH', 1974, 108, 7),
	(26014, 'GELGEL', 1974, 108, 7),
	(26015, 'JUMPAI', 1974, 108, 7),
	(26016, 'KAMASAN', 1974, 108, 7),
	(26017, 'KAMPUNG GELGEL', 1974, 108, 7),
	(26018, 'MANDUANG', 1974, 108, 7),
	(26019, 'SATRA', 1974, 108, 7),
	(26020, 'SELAT', 1974, 108, 7),
	(26021, 'SELISIHAN', 1974, 108, 7),
	(26022, 'SEMARAPURA KAJA', 1974, 108, 7),
	(26023, 'SEMARAPURA KANGIN', 1974, 108, 7),
	(26024, 'SEMARAPURA KAUH', 1974, 108, 7),
	(26025, 'SEMARAPURA KLOD KANGIN', 1974, 108, 7),
	(26026, 'SEMARAPURA KLOD/KELOD', 1974, 108, 7),
	(26027, 'SEMARAPURA TENGAH', 1974, 108, 7),
	(26028, 'TANGKAS', 1974, 108, 7),
	(26029, 'TEGAK', 1974, 108, 7),
	(26030, 'TOJAN', 1974, 108, 7),
	(26031, 'BATUKANDIK', 1975, 108, 7),
	(26032, 'BATUMADEG', 1975, 108, 7),
	(26033, 'BATUNUNGGUL', 1975, 108, 7),
	(26034, 'BUNGA MEKAR', 1975, 108, 7),
	(26035, 'JUNGUTBATU', 1975, 108, 7),
	(26036, 'KAMPUNG TOYAPAKEH', 1975, 108, 7),
	(26037, 'KLUMPU', 1975, 108, 7),
	(26038, 'KUTAMPI', 1975, 108, 7),
	(26039, 'KUTAMPI KALER', 1975, 108, 7),
	(26040, 'LEMBONGAN', 1975, 108, 7),
	(26041, 'PED', 1975, 108, 7),
	(26042, 'PEJUKUTAN', 1975, 108, 7),
	(26043, 'SAKTI', 1975, 108, 7),
	(26044, 'SEKARTAJI', 1975, 108, 7),
	(26045, 'SUANA', 1975, 108, 7),
	(26046, 'TANGLAD', 1975, 108, 7),
	(26047, 'ANGSERI', 1976, 109, 7),
	(26048, 'ANTAPAN', 1976, 109, 7),
	(26049, 'APUAN', 1976, 109, 7),
	(26050, 'BANGLI', 1976, 109, 7),
	(26051, 'BATUNYA', 1976, 109, 7),
	(26052, 'BATURITI', 1976, 109, 7),
	(26053, 'CANDIKUNING', 1976, 109, 7),
	(26054, 'LUWUS', 1976, 109, 7),
	(26055, 'MEKARSARI', 1976, 109, 7),
	(26056, 'PEREAN', 1976, 109, 7),
	(26057, 'PEREAN KANGIN', 1976, 109, 7),
	(26058, 'PEREAN TENGAH', 1976, 109, 7),
	(26059, 'BATUAJI', 1977, 109, 7),
	(26060, 'BATURITI', 1977, 109, 7),
	(26061, 'BELUMBANG', 1977, 109, 7),
	(26062, 'KELATING', 1977, 109, 7),
	(26063, 'KERAMBITAN', 1977, 109, 7),
	(26064, 'KESIUT', 1977, 109, 7),
	(26065, 'KUKUH', 1977, 109, 7),
	(26066, 'MELILING', 1977, 109, 7),
	(26067, 'PANGKUNG KARUNG', 1977, 109, 7),
	(26068, 'PENARUKAN', 1977, 109, 7),
	(26069, 'SAMSAM', 1977, 109, 7),
	(26070, 'SEMBUNG GEDE', 1977, 109, 7),
	(26071, 'TIBU BIU (TIBUBIYU)', 1977, 109, 7),
	(26072, 'TIMPAG', 1977, 109, 7),
	(26073, 'TISTA', 1977, 109, 7),
	(26074, 'BATANNYUH', 1978, 109, 7),
	(26075, 'BERINGKIT', 1978, 109, 7),
	(26076, 'CAUBELAYU', 1978, 109, 7),
	(26077, 'GELUNTUNG', 1978, 109, 7),
	(26078, 'KUKUH', 1978, 109, 7),
	(26079, 'KUWUM', 1978, 109, 7),
	(26080, 'MARGA', 1978, 109, 7),
	(26081, 'MARGA DAJAN PURI', 1978, 109, 7),
	(26082, 'MARGA DAUH PURI', 1978, 109, 7),
	(26083, 'PAYANGAN', 1978, 109, 7),
	(26084, 'PEKEN BELAYU', 1978, 109, 7),
	(26085, 'PETIGA', 1978, 109, 7),
	(26086, 'SELANBAWAK', 1978, 109, 7),
	(26087, 'TEGALJADI', 1978, 109, 7),
	(26088, 'TUA', 1978, 109, 7),
	(26089, 'BABAHAN', 1979, 109, 7),
	(26090, 'BIAUNG', 1979, 109, 7),
	(26091, 'BURUAN', 1979, 109, 7),
	(26092, 'JATILUWIH', 1979, 109, 7),
	(26093, 'JEGU', 1979, 109, 7),
	(26094, 'MENGESTE', 1979, 109, 7),
	(26095, 'PENATAHAN', 1979, 109, 7),
	(26096, 'PENEBEL', 1979, 109, 7),
	(26097, 'PESAGI', 1979, 109, 7),
	(26098, 'PITRA', 1979, 109, 7),
	(26099, 'REJASA', 1979, 109, 7),
	(26100, 'RIANG GEDE', 1979, 109, 7),
	(26101, 'SANGKETAN', 1979, 109, 7),
	(26102, 'SENGANAN', 1979, 109, 7),
	(26103, 'TAJEN', 1979, 109, 7),
	(26104, 'TEGALINGGAH', 1979, 109, 7),
	(26105, 'TENGKUDAK', 1979, 109, 7),
	(26106, 'WONGAYA GEDE', 1979, 109, 7),
	(26107, 'BANTIRAN', 1980, 109, 7),
	(26108, 'BATUNGSEL', 1980, 109, 7),
	(26109, 'BELATUNGAN', 1980, 109, 7),
	(26110, 'BELIMBING', 1980, 109, 7),
	(26111, 'JELIJIH PUNGGANG', 1980, 109, 7),
	(26112, 'KARYA SARI', 1980, 109, 7),
	(26113, 'KEBON PADANGAN', 1980, 109, 7),
	(26114, 'MUNDUK TEMU', 1980, 109, 7),
	(26115, 'PADANGAN', 1980, 109, 7),
	(26116, 'PAJAHAN', 1980, 109, 7),
	(26117, 'PUJUNGAN', 1980, 109, 7),
	(26118, 'PUPUAN', 1980, 109, 7),
	(26119, 'SANDA', 1980, 109, 7),
	(26120, 'ANTAP', 1981, 109, 7),
	(26121, 'BAJERA', 1981, 109, 7),
	(26122, 'BAJERA UTARA', 1981, 109, 7),
	(26123, 'BEREMBENG', 1981, 109, 7),
	(26124, 'MANIKYANG', 1981, 109, 7),
	(26125, 'PUPUAN SAWAH', 1981, 109, 7),
	(26126, 'SELEMADEG', 1981, 109, 7),
	(26127, 'SERAMPINGAN', 1981, 109, 7),
	(26128, 'WANAGIRI', 1981, 109, 7),
	(26129, 'WANAGIRI KAUH', 1981, 109, 7),
	(26130, 'BANTAS', 1982, 109, 7),
	(26131, 'BERABAN', 1982, 109, 7),
	(26132, 'DALANG', 1982, 109, 7),
	(26133, 'GADUNG SARI', 1982, 109, 7),
	(26134, 'GADUNGAN', 1982, 109, 7),
	(26135, 'GUNUNG SALAK', 1982, 109, 7),
	(26136, 'MAMBANG', 1982, 109, 7),
	(26137, 'MEGATI', 1982, 109, 7),
	(26138, 'TANGGUNTITI', 1982, 109, 7),
	(26139, 'TEGAL MENGKEB', 1982, 109, 7),
	(26140, 'ANGKAH', 1983, 109, 7),
	(26141, 'ANTOSARI', 1983, 109, 7),
	(26142, 'BENGKEL SARI', 1983, 109, 7),
	(26143, 'LALANG LINGGAH', 1983, 109, 7),
	(26144, 'LUMBUNG', 1983, 109, 7),
	(26145, 'LUMBUNG KAUH', 1983, 109, 7),
	(26146, 'MUNDEH', 1983, 109, 7),
	(26147, 'MUNDEH KANGIN', 1983, 109, 7),
	(26148, 'MUNDEH KAUH', 1983, 109, 7),
	(26149, 'SELABIH', 1983, 109, 7),
	(26150, 'TIYING GADING', 1983, 109, 7),
	(26151, 'BONGAN (BOONGAN)', 1984, 109, 7),
	(26152, 'BUAHAN', 1984, 109, 7),
	(26153, 'DAUH PEKEN', 1984, 109, 7),
	(26154, 'DEJAN PEKEN (DAJAN PEKEN)', 1984, 109, 7),
	(26155, 'DELOD PEKEN', 1984, 109, 7),
	(26156, 'DENBANTAS', 1984, 109, 7),
	(26157, 'GUBUG', 1984, 109, 7),
	(26158, 'SESANDAN', 1984, 109, 7),
	(26159, 'SUBAMIA', 1984, 109, 7),
	(26160, 'SUDIMARA', 1984, 109, 7),
	(26161, 'TUNJUK', 1984, 109, 7),
	(26162, 'WANASARI', 1984, 109, 7),
	(26163, 'ALUE BAGOK', 1985, 110, 8),
	(26164, 'ALUE BATEE', 1985, 110, 8),
	(26165, 'ALUE SUNDAK', 1985, 110, 8),
	(26166, 'ARONGAN', 1985, 110, 8),
	(26167, 'COT BULOH', 1985, 110, 8),
	(26168, 'COT JURUMUDI', 1985, 110, 8),
	(26169, 'COT KUMBANG', 1985, 110, 8),
	(26170, 'DRIEN RAMPAK', 1985, 110, 8),
	(26171, 'GUNONG PULO', 1985, 110, 8),
	(26172, 'KARANG HAMPA', 1985, 110, 8),
	(26173, 'KEUB', 1985, 110, 8),
	(26174, 'KUBU', 1985, 110, 8),
	(26175, 'PANTE MEUTIA', 1985, 110, 8),
	(26176, 'PANTON BAHAGIA', 1985, 110, 8),
	(26177, 'PANTON MAKMUR', 1985, 110, 8),
	(26178, 'PEULANTEU LB', 1985, 110, 8),
	(26179, 'PEURIBU', 1985, 110, 8),
	(26180, 'RIMBA LANGGEH', 1985, 110, 8),
	(26181, 'SEUNEUBOK LUENG', 1985, 110, 8),
	(26182, 'SEUNEUBOK TEUNGOH', 1985, 110, 8),
	(26183, 'SIMPANG PEUT', 1985, 110, 8),
	(26184, 'SUAK BIDOK', 1985, 110, 8),
	(26185, 'SUAK IE BEUSOU', 1985, 110, 8),
	(26186, 'SUAK KEUMUDEE', 1985, 110, 8),
	(26187, 'TEUPIN PEURAHO', 1985, 110, 8),
	(26188, 'UJONG BEUSA', 1985, 110, 8),
	(26189, 'UJONG SIMPANG', 1985, 110, 8),
	(26190, 'ALUE BAKONG', 1986, 110, 8),
	(26191, 'ALUE LHOK', 1986, 110, 8),
	(26192, 'BEURAWANG', 1986, 110, 8),
	(26193, 'BLANG SIBEUTONG', 1986, 110, 8),
	(26194, 'COT KEUMUNENG', 1986, 110, 8),
	(26195, 'COT LADA', 1986, 110, 8),
	(26196, 'GUNONG PANAH', 1986, 110, 8),
	(26197, 'KUALA PLING', 1986, 110, 8),
	(26198, 'KUTA PADANG', 1986, 110, 8),
	(26199, 'LAYUNG', 1986, 110, 8),
	(26200, 'LICEH', 1986, 110, 8),
	(26201, 'PEULANTEU SP', 1986, 110, 8),
	(26202, 'RAMBONG', 1986, 110, 8),
	(26203, 'SEUMULENG', 1986, 110, 8),
	(26204, 'SEUNEUBOK TRAP', 1986, 110, 8),
	(26205, 'SUAK PANGKAT', 1986, 110, 8),
	(26206, 'ULEE BLANG', 1986, 110, 8),
	(26207, 'BLANG BERANDANG', 1987, 110, 8),
	(26208, 'DRIEN RAMPAK', 1987, 110, 8),
	(26209, 'GAMPA', 1987, 110, 8),
	(26210, 'KAMPUNG BELAKANG', 1987, 110, 8),
	(26211, 'KAMPUNG DARAT', 1987, 110, 8),
	(26212, 'KAMPUNG PASIR', 1987, 110, 8),
	(26213, 'KUTA PADANG', 1987, 110, 8),
	(26214, 'LAPANG', 1987, 110, 8),
	(26215, 'LEUHAN', 1987, 110, 8),
	(26216, 'PADANG SEURAHET', 1987, 110, 8),
	(26217, 'PANGGONG', 1987, 110, 8),
	(26218, 'PASAR ACEH', 1987, 110, 8),
	(26219, 'RUNDENG', 1987, 110, 8),
	(26220, 'SEUNEUBOK', 1987, 110, 8),
	(26221, 'SUAK NIE', 1987, 110, 8),
	(26222, 'SUAK RAYA', 1987, 110, 8),
	(26223, 'SUAK RIBEE', 1987, 110, 8),
	(26224, 'SUAK SIGADENG', 1987, 110, 8),
	(26225, 'SUWAK INDRAPURI', 1987, 110, 8),
	(26226, 'UJONG BAROH', 1987, 110, 8),
	(26227, 'UJUNG KALAK', 1987, 110, 8),
	(26228, 'ALUE LHEE', 1988, 110, 8),
	(26229, 'ALUE LHOK', 1988, 110, 8),
	(26230, 'ALUE ON', 1988, 110, 8),
	(26231, 'ALUE PEUDEUNG', 1988, 110, 8),
	(26232, 'ALUE TAMPAK', 1988, 110, 8),
	(26233, 'BABAH MEULABOH', 1988, 110, 8),
	(26234, 'BATU JAYA', 1988, 110, 8),
	(26235, 'BEUREUGANG', 1988, 110, 8),
	(26236, 'BLANG DALAM', 1988, 110, 8),
	(26237, 'BLANG GEUNANG', 1988, 110, 8),
	(26238, 'DRIEN CALEE', 1988, 110, 8),
	(26239, 'KEUDE ARON', 1988, 110, 8),
	(26240, 'KEUDE TANJONG', 1988, 110, 8),
	(26241, 'KEURAMAT', 1988, 110, 8),
	(26242, 'MAREK', 1988, 110, 8),
	(26243, 'MENUANG TANJONG', 1988, 110, 8),
	(26244, 'MESJID', 1988, 110, 8),
	(26245, 'MEUNASAH ARA', 1988, 110, 8),
	(26246, 'MEUNASAH BULOH', 1988, 110, 8),
	(26247, 'MEUNASAH GANTUNG', 1988, 110, 8),
	(26248, 'MEUNASAH RAMBOT', 1988, 110, 8),
	(26249, 'MEUNASAH RAYEUK', 1988, 110, 8),
	(26250, 'MUKO', 1988, 110, 8),
	(26251, 'PADANG MANCANG', 1988, 110, 8),
	(26252, 'PADANG SIKABU', 1988, 110, 8),
	(26253, 'PALIMBUNGAN', 1988, 110, 8),
	(26254, 'PASI ARA', 1988, 110, 8),
	(26255, 'PASI JAMBU', 1988, 110, 8),
	(26256, 'PASI JEUMPA', 1988, 110, 8),
	(26257, 'PASI KUMBANG', 1988, 110, 8),
	(26258, 'PASI MEUGAT', 1988, 110, 8),
	(26259, 'PASI TEUNGOH', 1988, 110, 8),
	(26260, 'PEUNIA', 1988, 110, 8),
	(26261, 'PUCOK PUNGKIE', 1988, 110, 8),
	(26262, 'PUNGKIE', 1988, 110, 8),
	(26263, 'PUTIM', 1988, 110, 8),
	(26264, 'PUUK', 1988, 110, 8),
	(26265, 'SAWANG TEUBEE', 1988, 110, 8),
	(26266, 'SIMPANG', 1988, 110, 8),
	(26267, 'TANJONG BUNGA', 1988, 110, 8),
	(26268, 'TANJONG MEULABOH', 1988, 110, 8),
	(26269, 'TEULADAN', 1988, 110, 8),
	(26270, 'TEUPIN PANAH', 1988, 110, 8),
	(26271, 'TUMPOK LADANG', 1988, 110, 8),
	(26272, 'BALEE', 1989, 110, 8),
	(26273, 'BUKIT JAYA', 1989, 110, 8),
	(26274, 'BULOH', 1989, 110, 8),
	(26275, 'GUNONG KLENG', 1989, 110, 8),
	(26276, 'LANGUNG', 1989, 110, 8),
	(26277, 'MESJID TUHA', 1989, 110, 8),
	(26278, 'MEUREUBO', 1989, 110, 8),
	(26279, 'PASI ACEH BAROH', 1989, 110, 8),
	(26280, 'PASI ACEH TUNONG', 1989, 110, 8),
	(26281, 'PASI MESJID', 1989, 110, 8),
	(26282, 'PASI PINANG', 1989, 110, 8),
	(26283, 'PAYA BARO RP.', 1989, 110, 8),
	(26284, 'PAYA PEUNAGA', 1989, 110, 8),
	(26285, 'PEUNAGA CUT UJONG', 1989, 110, 8),
	(26286, 'PEUNAGA PASI', 1989, 110, 8),
	(26287, 'PEUNAGA RAYEUK', 1989, 110, 8),
	(26288, 'PUCOK REUDEUP', 1989, 110, 8),
	(26289, 'PULO TEUNGOH RANTO', 1989, 110, 8),
	(26290, 'RANTO PANYANG BARAT', 1989, 110, 8),
	(26291, 'RANTO PANYANG TIMUR', 1989, 110, 8),
	(26292, 'RANUB DONG', 1989, 110, 8),
	(26293, 'REUDEUP', 1989, 110, 8),
	(26294, 'SUMBER BATU', 1989, 110, 8),
	(26295, 'UJONG DRIEN', 1989, 110, 8),
	(26296, 'UJONG TANJONG', 1989, 110, 8),
	(26297, 'UJONG TANOH DARAT', 1989, 110, 8),
	(26298, 'ALUE KEUMANG', 1990, 110, 8),
	(26299, 'BABAH ISEUNG', 1990, 110, 8),
	(26300, 'BABAH KRUENG TEKLEP', 1990, 110, 8),
	(26301, 'BABAH LUENG', 1990, 110, 8),
	(26302, 'BERDIKARI', 1990, 110, 8),
	(26303, 'CANGGAI', 1990, 110, 8),
	(26304, 'GUNONG TAROK', 1990, 110, 8),
	(26305, 'JAMBAK', 1990, 110, 8),
	(26306, 'KEUDE SUWAK AWE', 1990, 110, 8),
	(26307, 'KEUTAMBANG', 1990, 110, 8),
	(26308, 'KRUENG BEUKAH', 1990, 110, 8),
	(26309, 'LANGO', 1990, 110, 8),
	(26310, 'LAWET', 1990, 110, 8),
	(26311, 'LHOK GUCI', 1990, 110, 8),
	(26312, 'LHOK SARI', 1990, 110, 8),
	(26313, 'MANJENG', 1990, 110, 8),
	(26314, 'MENUANG KINCO', 1990, 110, 8),
	(26315, 'PANTE CEUREUMEN (PANTE CERMIN)', 1990, 110, 8),
	(26316, 'PULO TEUNGOH MANYENG', 1990, 110, 8),
	(26317, 'SAWANG RAMBOT', 1990, 110, 8),
	(26318, 'SEUMANTOK', 1990, 110, 8),
	(26319, 'SEUMARA', 1990, 110, 8),
	(26320, 'SIKUNDO', 1990, 110, 8),
	(26321, 'SUWAK AWE', 1990, 110, 8),
	(26322, 'TEGAL SARI', 1990, 110, 8),
	(26323, 'ANTONG', 1991, 110, 8),
	(26324, 'BABAH KRUENG MANGGIE', 1991, 110, 8),
	(26325, 'BARO PAYA', 1991, 110, 8),
	(26326, 'BLANG BALEE', 1991, 110, 8),
	(26327, 'BLANG TEUNGOH', 1991, 110, 8),
	(26328, 'COT MANGGIE', 1991, 110, 8),
	(26329, 'GAMPONG BARO', 1991, 110, 8),
	(26330, 'GUNUNG MATA IE', 1991, 110, 8),
	(26331, 'KUALA MANYE', 1991, 110, 8),
	(26332, 'LEK LEK', 1991, 110, 8),
	(26333, 'MANGGIE', 1991, 110, 8),
	(26334, 'MEUTULANG', 1991, 110, 8),
	(26335, 'MUGO CUT', 1991, 110, 8),
	(26336, 'MUGO RAYEUK', 1991, 110, 8),
	(26337, 'PAYA BARO MEUKO', 1991, 110, 8),
	(26338, 'SIBINTANG', 1991, 110, 8),
	(26339, 'TAMPING', 1991, 110, 8),
	(26340, 'TUWI BUYA', 1991, 110, 8),
	(26341, 'UJONG RAJA', 1991, 110, 8),
	(26342, 'ALUE RAYA', 1992, 110, 8),
	(26343, 'COT AMUN', 1992, 110, 8),
	(26344, 'COT DARAT', 1992, 110, 8),
	(26345, 'COT LAMPISE', 1992, 110, 8),
	(26346, 'COT MESJID', 1992, 110, 8),
	(26347, 'COT PLUH', 1992, 110, 8),
	(26348, 'COT SEULAMAT', 1992, 110, 8),
	(26349, 'COT SEUMEURENG', 1992, 110, 8),
	(26350, 'DEUAH', 1992, 110, 8),
	(26351, 'GAMPONG COT', 1992, 110, 8),
	(26352, 'GAMPONG LADANG', 1992, 110, 8),
	(26353, 'GAMPONG TEUNGOH', 1992, 110, 8),
	(26354, 'KEUREUSENG', 1992, 110, 8),
	(26355, 'KRUENG TINGGAI', 1992, 110, 8),
	(26356, 'KUALA BUBON', 1992, 110, 8),
	(26357, 'LEUKEUN', 1992, 110, 8),
	(26358, 'LHOK BUBON', 1992, 110, 8),
	(26359, 'LUBOK', 1992, 110, 8),
	(26360, 'MESJID BARO', 1992, 110, 8),
	(26361, 'PANGE', 1992, 110, 8),
	(26362, 'PAYA LUMPAT', 1992, 110, 8),
	(26363, 'PINEM', 1992, 110, 8),
	(26364, 'PUCOK LUENG', 1992, 110, 8),
	(26365, 'RANGKILEH', 1992, 110, 8),
	(26366, 'REUSAK', 1992, 110, 8),
	(26367, 'SUAK GEUDUBANG', 1992, 110, 8),
	(26368, 'SUAK PANDAN', 1992, 110, 8),
	(26369, 'SUAK PANTE BREUH', 1992, 110, 8),
	(26370, 'SUAK SEUKEE', 1992, 110, 8),
	(26371, 'SUAK SEUMASEH', 1992, 110, 8),
	(26372, 'SUAK TIMAH', 1992, 110, 8),
	(26373, 'UJONG NGA', 1992, 110, 8),
	(26374, 'DRIEN SIBAK', 1993, 110, 8),
	(26375, 'GASEU', 1993, 110, 8),
	(26376, 'GEUDONG', 1993, 110, 8),
	(26377, 'GLEUNG', 1993, 110, 8),
	(26378, 'GUNONG BULOH', 1993, 110, 8),
	(26379, 'KAJEUNG', 1993, 110, 8),
	(26380, 'LANCONG', 1993, 110, 8),
	(26381, 'LEUBOK BEUTONG', 1993, 110, 8),
	(26382, 'LUENG BARO', 1993, 110, 8),
	(26383, 'PUNGKIE', 1993, 110, 8),
	(26384, 'RAMITIE', 1993, 110, 8),
	(26385, 'SAKUY', 1993, 110, 8),
	(26386, 'SARAH PERLAK', 1993, 110, 8),
	(26387, 'SIPOT', 1993, 110, 8),
	(26388, 'TANOH MIRAH', 1993, 110, 8),
	(26389, 'TUNGKOP', 1993, 110, 8),
	(26390, 'TUTUT', 1993, 110, 8),
	(26391, 'TUWI SAYA', 1993, 110, 8),
	(26392, 'ALUE BLANG', 1994, 110, 8),
	(26393, 'ALUE PANYANG', 1994, 110, 8),
	(26394, 'ALUE SIKAYA', 1994, 110, 8),
	(26395, 'ALUE SUNDAK', 1994, 110, 8),
	(26396, 'ARON BAROH', 1994, 110, 8),
	(26397, 'ARON TUNONG', 1994, 110, 8),
	(26398, 'BAKAT', 1994, 110, 8),
	(26399, 'BLANG MEE', 1994, 110, 8),
	(26400, 'COT KEUMUDEE', 1994, 110, 8),

	(26401, 'COT LAGAN BB/CM', 1994, 110, 8),
	(26402, 'COT MURONG', 1994, 110, 8),
	(26403, 'COT SITUAH', 1994, 110, 8),
	(26404, 'DARUL HUDA', 1994, 110, 8),
	(26405, 'DRIN MANGKO', 1994, 110, 8),
	(26406, 'GEMPA RAYA', 1994, 110, 8),
	(26407, 'GLEE SIBLAH', 1994, 110, 8),
	(26408, 'GUNONG/GUNUNG RAMBONG', 1994, 110, 8),
	(26409, 'GUNUNG HAMPA', 1994, 110, 8),
	(26410, 'IE ITAM BAROH', 1994, 110, 8),
	(26411, 'IE ITAM TUNONG', 1994, 110, 8),
	(26412, 'JAWA', 1994, 110, 8),
	(26413, 'JAWI', 1994, 110, 8),
	(26414, 'KEULEUMBAH', 1994, 110, 8),
	(26415, 'KUALA BHEE', 1994, 110, 8),
	(26416, 'LHUNG TGK YAH', 1994, 110, 8),
	(26417, 'LUENG BULOH', 1994, 110, 8),
	(26418, 'LUENG JAWA', 1994, 110, 8),
	(26419, 'LUENG TANOH THO', 1994, 110, 8),
	(26420, 'PADANG JAWA', 1994, 110, 8),
	(26421, 'PANTON', 1994, 110, 8),
	(26422, 'PASI ACEH', 1994, 110, 8),
	(26423, 'PASI ARA KB', 1994, 110, 8),
	(26424, 'PASI BIRAH', 1994, 110, 8),
	(26425, 'PASI LUNAK', 1994, 110, 8),
	(26426, 'PASI PANDAN', 1994, 110, 8),
	(26427, 'PAYA DUA', 1994, 110, 8),
	(26428, 'PAYA LUAH', 1994, 110, 8),
	(26429, 'PULO IE', 1994, 110, 8),
	(26430, 'RANTO PANYANG', 1994, 110, 8),
	(26431, 'SEUMANTOK', 1994, 110, 8),
	(26432, 'SUAK TRING', 1994, 110, 8),
	(26433, 'TEUMAROM', 1994, 110, 8),
	(26434, 'TINGKEUM PANYANG', 1994, 110, 8),
	(26435, 'ALUE KEUMUNENG', 1995, 110, 8),
	(26436, 'ALUE LEUHOB', 1995, 110, 8),
	(26437, 'ALUE PERMAN', 1995, 110, 8),
	(26438, 'BLANG COT MAMEH', 1995, 110, 8),
	(26439, 'BLANG COT RUBEK', 1995, 110, 8),
	(26440, 'BLANG LUAH', 1995, 110, 8),
	(26441, 'COT LAGAN LM', 1995, 110, 8),
	(26442, 'COT RAMBONG', 1995, 110, 8),
	(26443, 'IE SAYANG', 1995, 110, 8),
	(26444, 'KARAK', 1995, 110, 8),
	(26445, 'KULAM KAJU', 1995, 110, 8),
	(26446, 'LEUBOK PASI ARA', 1995, 110, 8),
	(26447, 'LHOK MALEE', 1995, 110, 8),
	(26448, 'LUENG BARO', 1995, 110, 8),
	(26449, 'MON PASONG', 1995, 110, 8),
	(26450, 'NAPAI', 1995, 110, 8),
	(26451, 'PASI JEUT', 1995, 110, 8),
	(26452, 'PASI MALEE', 1995, 110, 8),
	(26453, 'PASI MALI', 1995, 110, 8),
	(26454, 'PASI PANYANG', 1995, 110, 8),
	(26455, 'PEULUEKUNG', 1995, 110, 8),
	(26456, 'SIMPANG TEUMAROM', 1995, 110, 8),
	(26457, 'ULEE PASI ARA', 1995, 110, 8),
	(26458, 'ULEE PULO', 1995, 110, 8),
	(26459, 'ALUE BILIE', 1996, 110, 8),
	(26460, 'ALUE EMPEUK', 1996, 110, 8),
	(26461, 'ALUE KUYUEN', 1996, 110, 8),
	(26462, 'ALUE MEUGANDA', 1996, 110, 8),
	(26463, 'ALUE SEURALEN', 1996, 110, 8),
	(26464, 'BLANG DALAM', 1996, 110, 8),
	(26465, 'BLANG LUAH KB', 1996, 110, 8),
	(26466, 'BLANG MAKMUR', 1996, 110, 8),
	(26467, 'BUKET MEUGAJAH (BUKIT MEGAJAH)', 1996, 110, 8),
	(26468, 'COT PUNTI', 1996, 110, 8),
	(26469, 'GAMPONG BARO KB', 1996, 110, 8),
	(26470, 'GAMPONG BARO WT', 1996, 110, 8),
	(26471, 'GUNONG PANYANG', 1996, 110, 8),
	(26472, 'KUBU CAPANG (KEUBEU CAPANG)', 1996, 110, 8),
	(26473, 'LEUBOK/LUBOK PANYANG', 1996, 110, 8),
	(26474, 'PASI ARA WT', 1996, 110, 8),
	(26475, 'PASI JANENG', 1996, 110, 8),
	(26476, 'PAYA BARO', 1996, 110, 8),
	(26477, 'PAYA MEUGEUNDRANG', 1996, 110, 8),
	(26478, 'RAMBONG', 1996, 110, 8),
	(26479, 'RAMBONG PINTO', 1996, 110, 8),
	(26480, 'SEUNEUBOK DALAM', 1996, 110, 8),
	(26481, 'SEURADEUK', 1996, 110, 8),
	(26482, 'TANGKEH', 1996, 110, 8),
	(26483, 'TEUMIKET RANOM', 1996, 110, 8),
	(26484, 'TUWI EMPEUK', 1996, 110, 8),
	(26485, 'ALUE JEUREUJAK (JERJAK)', 1997, 111, 8),
	(26486, 'ALUE PEUNAWA', 1997, 111, 8),
	(26487, 'BLANG DALAM', 1997, 111, 8),
	(26488, 'GUNUNG SAMARINDA', 1997, 111, 8),
	(26489, 'IE MERAH', 1997, 111, 8),
	(26490, 'PANTEE CERMIN', 1997, 111, 8),
	(26491, 'PANTEE RAKYAT', 1997, 111, 8),
	(26492, 'ALUE MANGGOTA (ALUM ANGGOTA)', 1998, 111, 8),
	(26493, 'BAHARU', 1998, 111, 8),
	(26494, 'COT JEURAT', 1998, 111, 8),
	(26495, 'GEULUMPANG PAYONG', 1998, 111, 8),
	(26496, 'GUDANG', 1998, 111, 8),
	(26497, 'GUHANG', 1998, 111, 8),
	(26498, 'KEDAI / KEUDE PAYA', 1998, 111, 8),
	(26499, 'KEDAI / KEUDE SIBLAH', 1998, 111, 8),
	(26500, 'KUTA BAHAGIA', 1998, 111, 8),
	(26501, 'KUTA TINGGI', 1998, 111, 8),
	(26502, 'KUTA TUHA', 1998, 111, 8),
	(26503, 'LAMKUTA', 1998, 111, 8),
	(26504, 'LHUNG ASAN', 1998, 111, 8),
	(26505, 'LHUNG TAROK', 1998, 111, 8),
	(26506, 'MATA IE', 1998, 111, 8),
	(26507, 'MEUDANG ARA', 1998, 111, 8),
	(26508, 'PANTON RAYA', 1998, 111, 8),
	(26509, 'PASAR BLANG PIDIE', 1998, 111, 8),
	(26510, 'SEUNALOH', 1998, 111, 8),
	(26511, 'ALUE PADEE', 1999, 111, 8),
	(26512, 'ALUE PISANG', 1999, 111, 8),
	(26513, 'BLANG MAKMUR', 1999, 111, 8),
	(26514, 'BLANG PANYANG', 1999, 111, 8),
	(26515, 'GELANGGANG GAJAH', 1999, 111, 8),
	(26516, 'IE MAMEH', 1999, 111, 8),
	(26517, 'KAMPUNG TENGAH', 1999, 111, 8),
	(26518, 'KEUDEE BARO', 1999, 111, 8),
	(26519, 'KOTA BAHAGIA', 1999, 111, 8),
	(26520, 'KRUENG BATEE', 1999, 111, 8),
	(26521, 'KUALA TERUBUE', 1999, 111, 8),
	(26522, 'LAMA TUHA', 1999, 111, 8),
	(26523, 'LHOK GAJAH', 1999, 111, 8),
	(26524, 'LHUNG GELUMPANG', 1999, 111, 8),
	(26525, 'MUKA BLANG', 1999, 111, 8),
	(26526, 'PADANG SIKABU', 1999, 111, 8),
	(26527, 'PANTO CUT', 1999, 111, 8),
	(26528, 'PASAR KUTA BAHAGIA', 1999, 111, 8),
	(26529, 'COT BAK-U', 2000, 111, 8),
	(26530, 'GEULANGGANG BATEE', 2000, 111, 8),
	(26531, 'KUTA PAYA', 2000, 111, 8),
	(26532, 'MEUNASAH SUKON', 2000, 111, 8),
	(26533, 'MEUNASAH TEUNGOH / TENGAH', 2000, 111, 8),
	(26534, 'MEURANDEH', 2000, 111, 8),
	(26535, 'PADANG KEULELE/KELELE', 2000, 111, 8),
	(26536, 'SUKA DAMAI', 2000, 111, 8),
	(26537, 'UJUNG TANAH', 2000, 111, 8),
	(26538, 'ALUE RAMBOT', 2001, 111, 8),
	(26539, 'BLANG MANGGENG', 2001, 111, 8),
	(26540, 'KEUDEI', 2001, 111, 8),
	(26541, 'LADANG PANAH', 2001, 111, 8),
	(26542, 'LADANG TUHA I', 2001, 111, 8),
	(26543, 'LADANG TUHA II', 2001, 111, 8),
	(26544, 'LHOK PAWOH', 2001, 111, 8),
	(26545, 'LHOK PUNTOL', 2001, 111, 8),
	(26546, 'LHONG BARO', 2001, 111, 8),
	(26547, 'PADANG', 2001, 111, 8),
	(26548, 'PANTE PIRAK', 2001, 111, 8),
	(26549, 'PANTE RAJA', 2001, 111, 8),
	(26550, 'PANTOM MAKMUR', 2001, 111, 8),
	(26551, 'PAYA', 2001, 111, 8),
	(26552, 'PUSU INGIN RAYA', 2001, 111, 8),
	(26553, 'SEJAHTERA', 2001, 111, 8),
	(26554, 'SEUNEULOP', 2001, 111, 8),
	(26555, 'TEUNGOH', 2001, 111, 8),
	(26556, 'TOKOH', 2001, 111, 8),
	(26557, 'UJUNG TANAH', 2001, 111, 8),
	(26558, 'ALUE DAMA', 2002, 111, 8),
	(26559, 'LHANG', 2002, 111, 8),
	(26560, 'PESISIR MON MAMEH', 2002, 111, 8),
	(26561, 'PISANG', 2002, 111, 8),
	(26562, 'RAMBONG', 2002, 111, 8),
	(26563, 'TANGAN TANGAN CUT', 2002, 111, 8),
	(26564, 'UJUNG TANAH', 2002, 111, 8),
	(26565, 'BAHARU', 2003, 111, 8),
	(26566, 'BARAT', 2003, 111, 8),
	(26567, 'BLANG DALAM', 2003, 111, 8),
	(26568, 'COT MANCANG', 2003, 111, 8),
	(26569, 'DURIAN JANGEK', 2003, 111, 8),
	(26570, 'DURIAN RAMPAK', 2003, 111, 8),
	(26571, 'GADANG', 2003, 111, 8),
	(26572, 'KEDAI PALAK KERAMBIL', 2003, 111, 8),
	(26573, 'KEDAI SUSOH', 2003, 111, 8),
	(26574, 'KEPALA BANDAR', 2003, 111, 8),
	(26575, 'LADANG', 2003, 111, 8),
	(26576, 'LAMPOH DRIEN', 2003, 111, 8),
	(26577, 'MEUNASAH', 2003, 111, 8),
	(26578, 'PADANG BARU', 2003, 111, 8),
	(26579, 'PADANG HILIR', 2003, 111, 8),
	(26580, 'PADANG PANJANG', 2003, 111, 8),
	(26581, 'PALAK HILIR', 2003, 111, 8),
	(26582, 'PALAK HULU', 2003, 111, 8),
	(26583, 'PANJANG BARU', 2003, 111, 8),
	(26584, 'PANTAI PERAK', 2003, 111, 8),
	(26585, 'PAWOH', 2003, 111, 8),
	(26586, 'PINANG', 2003, 111, 8),
	(26587, 'PULAU KAYU', 2003, 111, 8),
	(26588, 'RUBEK MEUPAYONG', 2003, 111, 8),
	(26589, 'RUMAH DUA LAPIS', 2003, 111, 8),
	(26590, 'RUMAH PANJANG', 2003, 111, 8),
	(26591, 'TANGAH', 2003, 111, 8),
	(26592, 'UJUNG PADANG', 2003, 111, 8),
	(26593, 'ADAN', 2004, 111, 8),
	(26594, 'BINEH KRUENG', 2004, 111, 8),
	(26595, 'BLANG PADANG', 2004, 111, 8),
	(26596, 'DRIEN JALO', 2004, 111, 8),
	(26597, 'DRIEN KIPAH', 2004, 111, 8),
	(26598, 'GUNUNG CUT', 2004, 111, 8),
	(26599, 'KUTA BAK DRIEN/DRIN', 2004, 111, 8),
	(26600, 'MESJID', 2004, 111, 8),
	(26601, 'PADANG BAK JEUMPA / JUMPA', 2004, 111, 8),
	(26602, 'PADANG JAKLOK (BAKJOK)', 2004, 111, 8),
	(26603, 'PADANG KAWA', 2004, 111, 8),
	(26604, 'PANTE GEULUMPANG', 2004, 111, 8),
	(26605, 'SUAK LABU', 2004, 111, 8),
	(26606, 'SUAK NIBONG', 2004, 111, 8),
	(26607, 'BAET', 2005, 112, 8),
	(26608, 'BLANG KRUENG', 2005, 112, 8),
	(26609, 'CADEK', 2005, 112, 8),
	(26610, 'COT PAYA', 2005, 112, 8),
	(26611, 'KAJHU', 2005, 112, 8),
	(26612, 'KLIENG COT ARON', 2005, 112, 8),
	(26613, 'KLIENG MEURIA', 2005, 112, 8),
	(26614, 'LABUI', 2005, 112, 8),
	(26615, 'LAM ASAN', 2005, 112, 8),
	(26616, 'LAM BADA LHOK', 2005, 112, 8),
	(26617, 'LAM UJONG', 2005, 112, 8),
	(26618, 'LAMPINEUNG', 2005, 112, 8),
	(26619, 'MIRUK LAM REUDEUP', 2005, 112, 8),
	(26620, 'BUENG PAGEU', 2006, 112, 8),
	(26621, 'BUENG SIDOM', 2006, 112, 8),
	(26622, 'COT BAGIE', 2006, 112, 8),
	(26623, 'COT GEUNDREUT', 2006, 112, 8),
	(26624, 'COT HO-HO', 2006, 112, 8),
	(26625, 'COT JAMBO', 2006, 112, 8),
	(26626, 'COT KARIENG', 2006, 112, 8),
	(26627, 'COT LEU-OT', 2006, 112, 8),
	(26628, 'COT MADHI/MADI', 2006, 112, 8),
	(26629, 'COT MALEM', 2006, 112, 8),
	(26630, 'COT MANCANG', 2006, 112, 8),
	(26631, 'COT MEULANGEN', 2006, 112, 8),
	(26632, 'COT MON RAYA', 2006, 112, 8),
	(26633, 'COT NAMBAK', 2006, 112, 8),
	(26634, 'COT PUKLAT', 2006, 112, 8),
	(26635, 'COT RUMPUN', 2006, 112, 8),
	(26636, 'COT SAYUN', 2006, 112, 8),
	(26637, 'DATA MAKMUR', 2006, 112, 8),
	(26638, 'EMPEE BATA', 2006, 112, 8),
	(26639, 'KAMPUNG BLANG', 2006, 112, 8),
	(26640, 'KAYEE KUNYET', 2006, 112, 8),
	(26641, 'LAM SIEM', 2006, 112, 8),
	(26642, 'LAMME', 2006, 112, 8),
	(26643, 'MEULAYO (MELAYO)', 2006, 112, 8),
	(26644, 'PAYA UE (PAYA U)', 2006, 112, 8),
	(26645, 'TEUPIN BATEE', 2006, 112, 8),
	(26646, 'BAYU', 2007, 112, 8),
	(26647, 'DAROY KAMEU', 2007, 112, 8),
	(26648, 'DEUNONG', 2007, 112, 8),
	(26649, 'GAROT GEUCEU', 2007, 112, 8),
	(26650, 'GEUNDRIENG', 2007, 112, 8),
	(26651, 'GUE GAJAH', 2007, 112, 8),
	(26652, 'JEUMPET AJUEN/AJUN', 2007, 112, 8),
	(26653, 'KANDANG', 2007, 112, 8),
	(26654, 'KUTA KARANG', 2007, 112, 8),
	(26655, 'LAGANG', 2007, 112, 8),
	(26656, 'LAM BHEU', 2007, 112, 8),
	(26657, 'LAM BLANG MANYANG', 2007, 112, 8),
	(26658, 'LAM BLANG TRIENG', 2007, 112, 8),
	(26659, 'LAM KAWE', 2007, 112, 8),
	(26660, 'LAMCOT', 2007, 112, 8),
	(26661, 'LAMPASI ENGKING', 2007, 112, 8),
	(26662, 'LAMPEUNEUEN', 2007, 112, 8),
	(26663, 'LAMPEUNEURUT GAMPONG', 2007, 112, 8),
	(26664, 'LAMPEUNEURUT UJONG BLANG', 2007, 112, 8),
	(26665, 'LAMREUNG', 2007, 112, 8),
	(26666, 'LAMSIDAYA', 2007, 112, 8),
	(26667, 'LAMSITEH', 2007, 112, 8),
	(26668, 'LAMTHEUN', 2007, 112, 8),
	(26669, 'LEU UE (MATA IE)', 2007, 112, 8),
	(26670, 'LEUGEU', 2007, 112, 8),
	(26671, 'LHEU BLANG', 2007, 112, 8),
	(26672, 'PASHEU BEUTONG', 2007, 112, 8),
	(26673, 'PAYAROH', 2007, 112, 8),
	(26674, 'PUNIE', 2007, 112, 8),
	(26675, 'TINGKEUM', 2007, 112, 8),
	(26676, 'ULEE LUENG', 2007, 112, 8),
	(26677, 'ULEE TUI/THUI', 2007, 112, 8),
	(26678, 'BILUI/BILUY', 2008, 112, 8),
	(26679, 'BLANG KIREE', 2008, 112, 8),
	(26680, 'EMPETRING (EMPEE TRIENG)', 2008, 112, 8),
	(26681, 'LAM BLEUT', 2008, 112, 8),
	(26682, 'LAMBARO BILUY', 2008, 112, 8),
	(26683, 'LAMBATEE', 2008, 112, 8),
	(26684, 'LAMKUNYET', 2008, 112, 8),
	(26685, 'LAMSOD', 2008, 112, 8),
	(26686, 'LAMTADOK', 2008, 112, 8),
	(26687, 'LHANG', 2008, 112, 8),
	(26688, 'MANEE DAYAH', 2008, 112, 8),
	(26689, 'NEUSOK', 2008, 112, 8),
	(26690, 'TEUBALUI', 2008, 112, 8),
	(26691, 'TURAM', 2008, 112, 8),
	(26692, 'ANGAN', 2009, 112, 8),
	(26693, 'BEURABUNG', 2009, 112, 8),
	(26694, 'GAMPONG BLANG', 2009, 112, 8),
	(26695, 'GAMPONG COT', 2009, 112, 8),
	(26696, 'KRUENG KALEE', 2009, 112, 8),
	(26697, 'LAM ASAN', 2009, 112, 8),
	(26698, 'LAM BIHEU', 2009, 112, 8),
	(26699, 'LAM BITRA', 2009, 112, 8),
	(26700, 'LAM DUROY', 2009, 112, 8),
	(26701, 'LAM GAWEE', 2009, 112, 8),
	(26702, 'LAM KEUNUNG', 2009, 112, 8),
	(26703, 'LAM KLAT', 2009, 112, 8),
	(26704, 'LAM PEUDAYA', 2009, 112, 8),
	(26705, 'LAM PUJA', 2009, 112, 8),
	(26706, 'LAM PUUK', 2009, 112, 8),
	(26707, 'LAM REH', 2009, 112, 8),
	(26708, 'LAM TIMPEUNG', 2009, 112, 8),
	(26709, 'LAM UJONG', 2009, 112, 8),
	(26710, 'LAMBADA PEUKAN', 2009, 112, 8),
	(26711, 'LAMBARO SUKON', 2009, 112, 8),
	(26712, 'LAMBIHEU SIEM', 2009, 112, 8),
	(26713, 'LIE EUE', 2009, 112, 8),
	(26714, 'LIMPOK', 2009, 112, 8),
	(26715, 'MIRUEK TAMAN', 2009, 112, 8),
	(26716, 'SEULEU', 2009, 112, 8),
	(26717, 'SIEM', 2009, 112, 8),
	(26718, 'TANJONG SEULAMAT', 2009, 112, 8),
	(26719, 'TANJUNG DEYAH', 2009, 112, 8),
	(26720, 'TUNGKOP', 2009, 112, 8),
	(26721, 'ANEUK GLEE', 2010, 112, 8),
	(26722, 'COT KAREUNG', 2010, 112, 8),
	(26723, 'CUREH', 2010, 112, 8),
	(26724, 'EMPEE ARA', 2010, 112, 8),
	(26725, 'GROT MEUNASAH BARO', 2010, 112, 8),
	(26726, 'GROT MEUNASAH BLANG', 2010, 112, 8),
	(26727, 'GROT MEUNASAH MANYANG', 2010, 112, 8),
	(26728, 'INDRAPURI', 2010, 112, 8),
	(26729, 'JRUEK BAK KREH', 2010, 112, 8),
	(26730, 'JRUEK BALEE', 2010, 112, 8),
	(26731, 'KRUENG LAMKAREUNG', 2010, 112, 8),
	(26732, 'LAM BEUTONG', 2010, 112, 8),
	(26733, 'LAM ILIE GANTO', 2010, 112, 8),
	(26734, 'LAM ILIE MESJID', 2010, 112, 8),
	(26735, 'LAM ILIE TEUNGOH', 2010, 112, 8),
	(26736, 'LAM LUENG', 2010, 112, 8),
	(26737, 'LAM SIOT', 2010, 112, 8),
	(26738, 'LAMBUNOT', 2010, 112, 8),
	(26739, 'LAMLEUBOK', 2010, 112, 8),
	(26740, 'LAMPANAH BARO', 2010, 112, 8),
	(26741, 'LAMPANAH DAYAH', 2010, 112, 8),
	(26742, 'LAMPANAH RANJO', 2010, 112, 8),
	(26743, 'LAMPANAH TEUNGOH', 2010, 112, 8),
	(26744, 'LAMPANAH TUNONG', 2010, 112, 8),
	(26745, 'LAMPUPOK BARO', 2010, 112, 8),
	(26746, 'LAMPUPOK RAYA', 2010, 112, 8),
	(26747, 'LHEUE JEUMPA', 2010, 112, 8),
	(26748, 'LIMO BLANG', 2010, 112, 8),
	(26749, 'LIMO LAMLUWEUNG (LAMLUENG)', 2010, 112, 8),
	(26750, 'LIMO MESJID', 2010, 112, 8),
	(26751, 'LINGOM', 2010, 112, 8),
	(26752, 'MANGGRA', 2010, 112, 8),
	(26753, 'MEUNARA', 2010, 112, 8),
	(26754, 'MEUSALEE LHOK', 2010, 112, 8),
	(26755, 'MON ALUE', 2010, 112, 8),
	(26756, 'MUREU BARO', 2010, 112, 8),
	(26757, 'MUREU BUENG-UE', 2010, 112, 8),
	(26758, 'MUREU LAM GLUMPANG', 2010, 112, 8),
	(26759, 'MUREU ULEE TITI', 2010, 112, 8),
	(26760, 'PASAR INDRAPURI', 2010, 112, 8),
	(26761, 'REUKIH DAYAH', 2010, 112, 8),
	(26762, 'REUKIH KEUPULA', 2010, 112, 8),
	(26763, 'RITING', 2010, 112, 8),
	(26764, 'SEOT/SEUOT BAROH', 2010, 112, 8),
	(26765, 'SEOT/SEUOT TUNONG', 2010, 112, 8),
	(26766, 'SEULANGAI', 2010, 112, 8),
	(26767, 'SEUREUMO', 2010, 112, 8),
	(26768, 'SIHOM COT', 2010, 112, 8),
	(26769, 'SIHOM LHOK', 2010, 112, 8),
	(26770, 'SINYEU', 2010, 112, 8),
	(26771, 'ULEE KAREUNG', 2010, 112, 8),
	(26772, 'ULEE OE', 2010, 112, 8),
	(26773, 'AJEE CUT', 2011, 112, 8),
	(26774, 'AJEE PAGAR AIR', 2011, 112, 8),
	(26775, 'AJEE RAYEUK', 2011, 112, 8),
	(26776, 'ATEUK ANGGOK', 2011, 112, 8),
	(26777, 'ATEUK LUENG IE', 2011, 112, 8),
	(26778, 'BADA', 2011, 112, 8),
	(26779, 'BAKOY', 2011, 112, 8),
	(26780, 'BINEEH BLANG', 2011, 112, 8),
	(26781, 'BUNG CEUKOK', 2011, 112, 8),
	(26782, 'COT ALUE', 2011, 112, 8),
	(26783, 'COT BADA', 2011, 112, 8),
	(26784, 'COT GOT', 2011, 112, 8),
	(26785, 'COT MONTIWAN', 2011, 112, 8),
	(26786, 'COT SURUY', 2011, 112, 8),
	(26787, 'DHAM CEUKOK', 2011, 112, 8),
	(26788, 'DHAM PULO', 2011, 112, 8),
	(26789, 'GANI', 2011, 112, 8),
	(26790, 'JURONG PEUJEURA', 2011, 112, 8),
	(26791, 'KALUT', 2011, 112, 8),
	(26792, 'KAYEE LEE', 2011, 112, 8),
	(26793, 'LAM COT', 2011, 112, 8),
	(26794, 'LAM SINYEU', 2011, 112, 8),
	(26795, 'LAM TEUNGOH', 2011, 112, 8),
	(26796, 'LAM UE', 2011, 112, 8),
	(26797, 'LAMBADA', 2011, 112, 8),
	(26798, 'LAMBARO', 2011, 112, 8),
	(26799, 'LAMDAYA', 2011, 112, 8),
	(26800, 'LAMPREH L. TEUNGOH', 2011, 112, 8),
	(26801, 'LAMPREH LAM JAMPROK', 2011, 112, 8),
	(26802, 'LHOK SUKON', 2011, 112, 8),
	(26803, 'LUBOK GAPUY', 2011, 112, 8),
	(26804, 'LUBUK BATEE', 2011, 112, 8),
	(26805, 'MEUNASAH BARO', 2011, 112, 8),
	(26806, 'MEUNASAH DAYAH', 2011, 112, 8),
	(26807, 'MEUNASAH KRUENG', 2011, 112, 8),
	(26808, 'MEUNASAH MANYANG LG', 2011, 112, 8),
	(26809, 'MEUNASAH MANYANG PAGAR AIR', 2011, 112, 8),
	(26810, 'MEUNASAH MANYET', 2011, 112, 8),
	(26811, 'MEUNASAH TEUTONG', 2011, 112, 8),
	(26812, 'PALEUH BLANG', 2011, 112, 8),
	(26813, 'PALEUH PULO', 2011, 112, 8),
	(26814, 'PANTEE', 2011, 112, 8),
	(26815, 'PASIE LG', 2011, 112, 8),
	(26816, 'PASIE LUBOK', 2011, 112, 8),
	(26817, 'REULOH', 2011, 112, 8),
	(26818, 'SANTAN', 2011, 112, 8),
	(26819, 'SIRON', 2011, 112, 8),
	(26820, 'TANJONG', 2011, 112, 8),
	(26821, 'TEUBANG PHUI', 2011, 112, 8),
	(26822, 'UJONG XII', 2011, 112, 8),
	(26823, 'BAK SUKON', 2012, 112, 8),
	(26824, 'BANDASAFA', 2012, 112, 8),
	(26825, 'BARIH LHOK', 2012, 112, 8),
	(26826, 'BITHAK', 2012, 112, 8),
	(26827, 'BUENG SIMEK', 2012, 112, 8),
	(26828, 'COT BAYU', 2012, 112, 8),
	(26829, 'GLEE JAI', 2012, 112, 8),
	(26830, 'IE ALANG DAYAH', 2012, 112, 8),
	(26831, 'IE ALANG LAM GHUI', 2012, 112, 8),
	(26832, 'IE ALANG LAM KREUMEUH (KEUREUMEUN)', 2012, 112, 8),
	(26833, 'IE ALANG MESJID', 2012, 112, 8),
	(26834, 'KEUMIREU', 2012, 112, 8),
	(26835, 'KEURUWEUNG BLANG', 2012, 112, 8),
	(26836, 'KEURUWEUNG KRUENG', 2012, 112, 8),
	(26837, 'LAM ALING/ALIENG', 2012, 112, 8),
	(26838, 'LAM BEUGAK', 2012, 112, 8),
	(26839, 'LAM KLENG/KLEENG', 2012, 112, 8),
	(26840, 'LAM LEUOT', 2012, 112, 8),
	(26841, 'LAM LEUPUNG', 2012, 112, 8),
	(26842, 'LAMPAKUK', 2012, 112, 8),
	(26843, 'LAMPOH RAJA', 2012, 112, 8),
	(26844, 'LAMSIE (LAMSIEM)', 2012, 112, 8),
	(26845, 'LAMTUI', 2012, 112, 8),
	(26846, 'LEUPUNG BALEU/BALEE', 2012, 112, 8),
	(26847, 'LEUPUNG BRUK/BRUEK', 2012, 112, 8),
	(26848, 'MAHENG', 2012, 112, 8),
	(26849, 'PAKUK', 2012, 112, 8),
	(26850, 'PASAR LAMPAKUK', 2012, 112, 8),
	(26851, 'SIGAPANG', 2012, 112, 8),
	(26852, 'SIRON BLANG', 2012, 112, 8),
	(26853, 'SIRON KRUENG', 2012, 112, 8),
	(26854, 'TUTUI/TUTUE', 2012, 112, 8),
	(26855, 'AWEE', 2013, 112, 8),
	(26856, 'BARUEH', 2013, 112, 8),
	(26857, 'BEUNG', 2013, 112, 8),
	(26858, 'BUKIT MEUSARA', 2013, 112, 8),
	(26859, 'CUCUM', 2013, 112, 8),
	(26860, 'DATA CUT', 2013, 112, 8),
	(26861, 'JALIN', 2013, 112, 8),
	(26862, 'JANTHO', 2013, 112, 8),
	(26863, 'JANTHO BARU', 2013, 112, 8),
	(26864, 'JANTHO MAKMUR', 2013, 112, 8),
	(26865, 'SUKA TANI', 2013, 112, 8),
	(26866, 'TEUREUBEH', 2013, 112, 8),
	(26867, 'WEU', 2013, 112, 8),
	(26868, 'BUGHU', 2014, 112, 8),
	(26869, 'LAM ARA CUT', 2014, 112, 8),
	(26870, 'LAM ARA ENGKIT', 2014, 112, 8),
	(26871, 'LAM ARA TUNONG', 2014, 112, 8),
	(26872, 'LAMBARO SAMAHANI', 2014, 112, 8),
	(26873, 'LAMSITEH COT', 2014, 112, 8),
	(26874, 'LEUPUNG CUT', 2014, 112, 8),
	(26875, 'LEUPUNG RAYEUK', 2014, 112, 8),
	(26876, 'LEUPUNG RIWAT', 2014, 112, 8),
	(26877, 'LUBOK BATEE', 2014, 112, 8),
	(26878, 'LUBOK BUNI', 2014, 112, 8),
	(26879, 'REULEUNG GEULUMPANG', 2014, 112, 8),
	(26880, 'REULEUNG KARIENG', 2014, 112, 8),
	(26881, 'TEUDAYAH', 2014, 112, 8),
	(26882, 'TUMBO BARO', 2014, 112, 8),
	(26883, 'GLA DEYAH', 2015, 112, 8),
	(26884, 'GLA MEUNASAH BARO', 2015, 112, 8),
	(26885, 'LAM GAPANG', 2015, 112, 8),
	(26886, 'LAMPERMEI (LAMPERMEE)', 2015, 112, 8),
	(26887, 'LUENG IE', 2015, 112, 8),
	(26888, 'MEUNASAH BAET', 2015, 112, 8),
	(26889, 'MEUNASAH BAKTRIENG', 2015, 112, 8),
	(26890, 'MEUNASAH INTAN', 2015, 112, 8),
	(26891, 'MEUNASAH MANYANG', 2015, 112, 8),
	(26892, 'MEUNASAH PAPEUN', 2015, 112, 8),
	(26893, 'MIRUK', 2015, 112, 8),
	(26894, 'RUMPET', 2015, 112, 8),
	(26895, 'ARON', 2016, 112, 8),
	(26896, 'BABAH JURONG', 2016, 112, 8),
	(26897, 'BAK BULOH', 2016, 112, 8),
	(26898, 'BEURANGONG', 2016, 112, 8),
	(26899, 'BUENG BAK JOK', 2016, 112, 8),
	(26900, 'COT BEUT', 2016, 112, 8),
	(26901, 'COT CUT', 2016, 112, 8),
	(26902, 'COT LAMME', 2016, 112, 8),
	(26903, 'COT MANCANG', 2016, 112, 8),
	(26904, 'COT MASAM', 2016, 112, 8),
	(26905, 'COT PEUTANO', 2016, 112, 8),
	(26906, 'COT PREH', 2016, 112, 8),
	(26907, 'COT RAYA', 2016, 112, 8),
	(26908, 'COT YANG', 2016, 112, 8),
	(26909, 'CUCUM', 2016, 112, 8),
	(26910, 'DEYAH', 2016, 112, 8),
	(26911, 'GUE', 2016, 112, 8),
	(26912, 'KRUENG ANOI', 2016, 112, 8),
	(26913, 'LAM ALUE CUT', 2016, 112, 8),
	(26914, 'LAM ALUE RAYA', 2016, 112, 8),
	(26915, 'LAM ASAN', 2016, 112, 8),
	(26916, 'LAM BAET', 2016, 112, 8),
	(26917, 'LAM BUNOT PAYA', 2016, 112, 8),
	(26918, 'LAM BUNOT TANOH', 2016, 112, 8),
	(26919, 'LAM GLUMPANG/GEULUMPANG', 2016, 112, 8),
	(26920, 'LAM NEUHEUN', 2016, 112, 8),
	(26921, 'LAM RAYA', 2016, 112, 8),
	(26922, 'LAM SABANG', 2016, 112, 8),
	(26923, 'LAM SEUNONG', 2016, 112, 8),
	(26924, 'LAM TRIENG', 2016, 112, 8),
	(26925, 'LAMBRO BILEU', 2016, 112, 8),
	(26926, 'LAMBRO DEYAH', 2016, 112, 8),
	(26927, 'LAMCEU', 2016, 112, 8),
	(26928, 'LAMPOH KEUDEE', 2016, 112, 8),
	(26929, 'LAMPOH TAROM', 2016, 112, 8),
	(26930, 'LAMPUUK', 2016, 112, 8),
	(26931, 'LAMROH', 2016, 112, 8),
	(26932, 'LAMTEUBEE GEUPULA', 2016, 112, 8),
	(26933, 'LAMTEUBEE MON ARA', 2016, 112, 8),
	(26934, 'LEUPUNG MESJID', 2016, 112, 8),
	(26935, 'LEUPUNG ULEE ALUE', 2016, 112, 8),
	(26936, 'MEUNASAH BAK TRIENG', 2016, 112, 8),
	(26937, 'PUUK', 2016, 112, 8),
	(26938, 'RABEU', 2016, 112, 8),
	(26939, 'SEUPEU', 2016, 112, 8),
	(26940, 'TUMPOK LAMPOH', 2016, 112, 8),
	(26941, 'UJONG BLANG', 2016, 112, 8),
	(26942, 'LAM KUBU', 2017, 112, 8),
	(26943, 'LAMBARO TUNONG', 2017, 112, 8),
	(26944, 'LAMTAMOT', 2017, 112, 8),
	(26945, 'LON ASAN', 2017, 112, 8),
	(26946, 'LON BAROH', 2017, 112, 8),
	(26947, 'PANCA', 2017, 112, 8),
	(26948, 'PANCA KUBU', 2017, 112, 8),
	(26949, 'PAYA KEUREULEH', 2017, 112, 8),
	(26950, 'SAREE ACEH', 2017, 112, 8),
	(26951, 'SUKA MULYA (SUKAMULIA)', 2017, 112, 8),
	(26952, 'SUKADAMAI', 2017, 112, 8),
	(26953, 'TEULADAN', 2017, 112, 8),
	(26954, 'DEAH MAMPLAM', 2018, 112, 8),
	(26955, 'LAMSENIA (LAMSEUNIA)', 2018, 112, 8),
	(26956, 'LAYEUN', 2018, 112, 8),
	(26957, 'MEUNASAH BAK UE', 2018, 112, 8),
	(26958, 'MEUNASAH MESJID (MESJID LEUPUNG)', 2018, 112, 8),
	(26959, 'PULOT', 2018, 112, 8),
	(26960, 'ATEUK PAYA', 2019, 112, 8),
	(26961, 'KUEH', 2019, 112, 8),
	(26962, 'LAM ATEUK', 2019, 112, 8),
	(26963, 'LAM COK', 2019, 112, 8),
	(26964, 'LAM GABOH', 2019, 112, 8),
	(26965, 'LAM KRUET', 2019, 112, 8),
	(26966, 'LAMBARO KUEH', 2019, 112, 8),
	(26967, 'LAMBARO SEUBUN', 2019, 112, 8),
	(26968, 'LAMPAYA', 2019, 112, 8),
	(26969, 'MEUNASAH BALEE LAMLHOM', 2019, 112, 8),
	(26970, 'MEUNASAH BARO LAM PUUK', 2019, 112, 8),
	(26971, 'MEUNASAH BARO LAMLHOM', 2019, 112, 8),
	(26972, 'MEUNASAH BEUTONG', 2019, 112, 8),
	(26973, 'MEUNASAH BLANG', 2019, 112, 8),
	(26974, 'MEUNASAH CUT', 2019, 112, 8),
	(26975, 'MEUNASAH KARIENG', 2019, 112, 8),
	(26976, 'MEUNASAH LAM GIREK', 2019, 112, 8),
	(26977, 'MEUNASAH MANYANG', 2019, 112, 8),
	(26978, 'MEUNASAH MESJID LAMLHOM', 2019, 112, 8),
	(26979, 'MEUNASAH MON CUT', 2019, 112, 8),
	(26980, 'MNSH MESJID LAM PUUK', 2019, 112, 8),
	(26981, 'MON IKEUN', 2019, 112, 8),
	(26982, 'NAGA UMBANG', 2019, 112, 8),
	(26983, 'NUSA', 2019, 112, 8),
	(26984, 'SEUBUN AYON', 2019, 112, 8),
	(26985, 'SEUBUN KEUTAPANG', 2019, 112, 8),
	(26986, 'TANJONG', 2019, 112, 8),
	(26987, 'WEU RAYA', 2019, 112, 8),
	(26988, 'BAROH BLANG MEE', 2020, 112, 8),
	(26989, 'BAROH GENTEUT', 2020, 112, 8),
	(26990, 'BAROH KRUENG KALA', 2020, 112, 8),
	(26991, 'BIREK', 2020, 112, 8),
	(26992, 'COT', 2020, 112, 8),
	(26993, 'CUNDIN', 2020, 112, 8),
	(26994, 'GAPUY', 2020, 112, 8),
	(26995, 'GLEE BRUEK', 2020, 112, 8),
	(26996, 'JANTANG', 2020, 112, 8),
	(26997, 'KAREUNG', 2020, 112, 8),
	(26998, 'KEUTAPANG', 2020, 112, 8),
	(26999, 'LAM GRIHEU', 2020, 112, 8),
	(27000, 'LAM JUHANG', 2020, 112, 8)
	`)
	if err != nil {
		panic(err)
	}
	fmt.Println("insert villige 26")
}
