package kel20

import (
	"fmt"

	"github.com/danangkonang/rest-api/config"
)

func Kel22() {
	db := config.Connect()
	_, err := db.Exec(`INSERT INTO villages (id,village,sub_district_id,regency_id,province_id) VALUES
		(22001, 'ANGSANAH', 1701, 89, 6),
		(22002, 'BANYUPELLE', 1701, 89, 6),
		(22003, 'KACOK', 1701, 89, 6),
		(22004, 'LARANGAN BADUNG', 1701, 89, 6),
		(22005, 'PALENGAAN DAJA', 1701, 89, 6),
		(22006, 'PALENGAAN LAOK', 1701, 89, 6),
		(22007, 'PANAAN', 1701, 89, 6),
		(22008, 'POTOAN DAJA', 1701, 89, 6),
		(22009, 'POTOAN LAOK', 1701, 89, 6),
		(22010, 'REK KERREK', 1701, 89, 6),
		(22011, 'ROMBUH', 1701, 89, 6),
		(22012, 'BARURAMBAT KOTA (BARKOT)', 1702, 89, 6),
		(22013, 'BETTET', 1702, 89, 6),
		(22014, 'BUGIH', 1702, 89, 6),
		(22015, 'GLADAK ANYAR', 1702, 89, 6),
		(22016, 'JALMAK', 1702, 89, 6),
		(22017, 'JUNGCANGCANG', 1702, 89, 6),
		(22018, 'KANGENAN', 1702, 89, 6),
		(22019, 'KOLPAJUNG', 1702, 89, 6),
		(22020, 'KOWEL', 1702, 89, 6),
		(22021, 'LADEN', 1702, 89, 6),
		(22022, 'NYLABU DAYA/DAJA', 1702, 89, 6),
		(22023, 'NYLABU LAOK', 1702, 89, 6),
		(22024, 'PANEMPAN', 1702, 89, 6),
		(22025, 'PARTEKER', 1702, 89, 6),
		(22026, 'PATEMON', 1702, 89, 6),
		(22027, 'TEJAH/TEJA BARAT', 1702, 89, 6),
		(22028, 'TEJAH/TEJA TIMUR', 1702, 89, 6),
		(22029, 'TORONAN', 1702, 89, 6),
		(22030, 'BATUKERBUY', 1703, 89, 6),
		(22031, 'BINDANG', 1703, 89, 6),
		(22032, 'DEMPO BARAT', 1703, 89, 6),
		(22033, 'DEMPO TIMUR', 1703, 89, 6),
		(22034, 'SANA DAJA', 1703, 89, 6),
		(22035, 'SANA TENGAH', 1703, 89, 6),
		(22036, 'SOTOBAR (SOTABAR)', 1703, 89, 6),
		(22037, 'TEGANGSER DAYA', 1703, 89, 6),
		(22038, 'TLONTORAJA', 1703, 89, 6),
		(22039, 'AMBENDER', 1704, 89, 6),
		(22040, 'BULANGAN BARAT', 1704, 89, 6),
		(22041, 'BULANGAN BRANTA', 1704, 89, 6),
		(22042, 'BULANGAN HAJI', 1704, 89, 6),
		(22043, 'BULANGAN TIMUR', 1704, 89, 6),
		(22044, 'PALESANGGAR', 1704, 89, 6),
		(22045, 'PASANGGAR', 1704, 89, 6),
		(22046, 'PEGANTENAN (PAGANTENAN)', 1704, 89, 6),
		(22047, 'PLAKPAK', 1704, 89, 6),
		(22048, 'TANJUNG', 1704, 89, 6),
		(22049, 'TEBUL BARAT', 1704, 89, 6),
		(22050, 'TEBUL TIMUR', 1704, 89, 6),
		(22051, 'TLAGAH (TALAGAH)', 1704, 89, 6),
		(22052, 'BADUNG', 1705, 89, 6),
		(22053, 'BANYU BULU', 1705, 89, 6),
		(22054, 'BATU KALANGAN', 1705, 89, 6),
		(22055, 'BILLA\\AN', 1705, 89, 6),
		(22056, 'CAMPOR', 1705, 89, 6),
		(22057, 'CANDI BURUNG', 1705, 89, 6),
		(22058, 'GRO\\OM', 1705, 89, 6),
		(22059, 'JAMBRINGIN', 1705, 89, 6),
		(22060, 'KARANGANYAR', 1705, 89, 6),
		(22061, 'KLAMPAR', 1705, 89, 6),
		(22062, 'KODIK', 1705, 89, 6),
		(22063, 'LENTENG', 1705, 89, 6),
		(22064, 'MAPPER', 1705, 89, 6),
		(22065, 'PANAGUAN', 1705, 89, 6),
		(22066, 'PANGBATOK', 1705, 89, 6),
		(22067, 'PANGLEMAH', 1705, 89, 6),
		(22068, 'PANGTONGGAL', 1705, 89, 6),
		(22069, 'PANGURAYAN', 1705, 89, 6),
		(22070, 'PROPPO', 1705, 89, 6),
		(22071, 'RANG PERANG DAYA', 1705, 89, 6),
		(22072, 'RANG PERANG LAOK', 1705, 89, 6),
		(22073, 'SAMATAN', 1705, 89, 6),
		(22074, 'SAMIRAN', 1705, 89, 6),
		(22075, 'SRAMBAH', 1705, 89, 6),
		(22076, 'TATTANGOH', 1705, 89, 6),
		(22077, 'TLANGOH', 1705, 89, 6),
		(22078, 'TOKET', 1705, 89, 6),
		(22079, 'AMBAT', 1706, 89, 6),
		(22080, 'BANDARAN', 1706, 89, 6),
		(22081, 'BRANTA PASISIR', 1706, 89, 6),
		(22082, 'BRANTA TINGGI', 1706, 89, 6),
		(22083, 'BUKEK', 1706, 89, 6),
		(22084, 'CEGUK', 1706, 89, 6),
		(22085, 'DABUAN', 1706, 89, 6),
		(22086, 'GUGUL', 1706, 89, 6),
		(22087, 'KRAMAT', 1706, 89, 6),
		(22088, 'LARANGAN SLAMPAR', 1706, 89, 6),
		(22089, 'LARANGAN TOKOL', 1706, 89, 6),
		(22090, 'MANGAR', 1706, 89, 6),
		(22091, 'PANGLEGUR', 1706, 89, 6),
		(22092, 'TARO\\AN', 1706, 89, 6),
		(22093, 'TERRAK', 1706, 89, 6),
		(22094, 'TLANAKAN', 1706, 89, 6),
		(22095, 'TLESA', 1706, 89, 6),
		(22096, 'BENDO MUNGAL', 1707, 90, 6),
		(22097, 'DERMO', 1707, 90, 6),
		(22098, 'GEMPENG', 1707, 90, 6),
		(22099, 'KALIANYAR', 1707, 90, 6),
		(22100, 'KALIREJO', 1707, 90, 6),
		(22101, 'KAUMAN', 1707, 90, 6),
		(22102, 'KERSIKAN', 1707, 90, 6),
		(22103, 'KIDUL DALEM', 1707, 90, 6),
		(22104, 'KOLURSARI', 1707, 90, 6),
		(22105, 'LATEK', 1707, 90, 6),
		(22106, 'MANARUWI', 1707, 90, 6),
		(22107, 'MASANGAN', 1707, 90, 6),
		(22108, 'POGAR', 1707, 90, 6),
		(22109, 'RACI', 1707, 90, 6),
		(22110, 'TAMBAKAN', 1707, 90, 6),
		(22111, 'BAUJENG', 1708, 90, 6),
		(22112, 'BEJI', 1708, 90, 6),
		(22113, 'BEJI', 1708, 90, 6),
		(22114, 'BEJI TIMUR', 1708, 90, 6),
		(22115, 'CANGKRING MALANG', 1708, 90, 6),
		(22116, 'GAJAHBENDO', 1708, 90, 6),
		(22117, 'GLANGGANG', 1708, 90, 6),
		(22118, 'GUNUNG GANGSIR', 1708, 90, 6),
		(22119, 'GUNUNG SARI', 1708, 90, 6),
		(22120, 'KEDUNG BOTO', 1708, 90, 6),
		(22121, 'KEDUNG RINGIN', 1708, 90, 6),
		(22122, 'KEMIRI MUKA', 1708, 90, 6),
		(22123, 'KENEP', 1708, 90, 6),
		(22124, 'KUKUSAN', 1708, 90, 6),
		(22125, 'NGEMBE', 1708, 90, 6),
		(22126, 'PAGAK', 1708, 90, 6),
		(22127, 'PONDOK CINA', 1708, 90, 6),
		(22128, 'SIDOWAYAH', 1708, 90, 6),
		(22129, 'TANAH BARU', 1708, 90, 6),
		(22130, 'WONOKOYO', 1708, 90, 6),
		(22131, 'BAKALAN', 1709, 90, 6),
		(22132, 'BLANDONGAN', 1709, 90, 6),
		(22133, 'BUGUL KIDUL', 1709, 90, 6),
		(22134, 'KEPEL', 1709, 90, 6),
		(22135, 'KRAMPYANGAN', 1709, 90, 6),
		(22136, 'TAPAAN', 1709, 90, 6),
		(22137, 'BUKIR', 1710, 90, 6),
		(22138, 'GADINGREJO', 1710, 90, 6),
		(22139, 'GENTONG', 1710, 90, 6),
		(22140, 'KARANGKETUG', 1710, 90, 6),
		(22141, 'KRAPYAKREJO', 1710, 90, 6),
		(22142, 'PETAHUNAN', 1710, 90, 6),
		(22143, 'RANDUSARI', 1710, 90, 6),
		(22144, 'SEBANI', 1710, 90, 6),
		(22145, 'BULUSARI', 1711, 90, 6),
		(22146, 'CARAT', 1711, 90, 6),
		(22147, 'CIKEUSAL', 1711, 90, 6),
		(22148, 'CUPANG', 1711, 90, 6),
		(22149, 'GEMPOL', 1711, 90, 6),
		(22150, 'GEMPOL', 1711, 90, 6),
		(22151, 'JERUK PURUT', 1711, 90, 6),
		(22152, 'KARANGREJO', 1711, 90, 6),
		(22153, 'KEDUNGBUNDER', 1711, 90, 6),
		(22154, 'KEJAPANAN', 1711, 90, 6),
		(22155, 'KEMPEK', 1711, 90, 6),
		(22156, 'KEPULUNGAN', 1711, 90, 6),
		(22157, 'LEGOK', 1711, 90, 6),
		(22158, 'NGERONG', 1711, 90, 6),
		(22159, 'PALIMANAN BARAT', 1711, 90, 6),
		(22160, 'RANDUPITU', 1711, 90, 6),
		(22161, 'SUMBER SUKO', 1711, 90, 6),
		(22162, 'WALAHAR', 1711, 90, 6),
		(22163, 'WATUKOSEK', 1711, 90, 6),
		(22164, 'WINONG', 1711, 90, 6),
		(22165, 'WINONG', 1711, 90, 6),
		(22166, 'WONOSARI', 1711, 90, 6),
		(22167, 'WONOSUNYO', 1711, 90, 6),
		(22168, 'BAJANGAN', 1712, 90, 6),
		(22169, 'BAYEMAN', 1712, 90, 6),
		(22170, 'BRAMBANG', 1712, 90, 6),
		(22171, 'GAYAM', 1712, 90, 6),
		(22172, 'GONDANG REJO', 1712, 90, 6),
		(22173, 'GONDANG WETAN', 1712, 90, 6),
		(22174, 'GROGOL', 1712, 90, 6),
		(22175, 'KALI REJO', 1712, 90, 6),
		(22176, 'KARANG SENTUL', 1712, 90, 6),
		(22177, 'KEBON CANDI', 1712, 90, 6),
		(22178, 'KERSIKAN', 1712, 90, 6),
		(22179, 'LAJUK', 1712, 90, 6),
		(22180, 'PATEGUHAN', 1712, 90, 6),
		(22181, 'PEKANGKUNGAN', 1712, 90, 6),
		(22182, 'RANGGEH', 1712, 90, 6),
		(22183, 'SEKAR PUTIH', 1712, 90, 6),
		(22184, 'TEBAS', 1712, 90, 6),
		(22185, 'TENGGILIS REJO', 1712, 90, 6),
		(22186, 'WONOJATI', 1712, 90, 6),
		(22187, 'WONOSARI', 1712, 90, 6),
		(22188, 'CUKUR GONDANG', 1713, 90, 6),
		(22189, 'GRATI TUNON', 1713, 90, 6),
		(22190, 'KALIPANG', 1713, 90, 6),
		(22191, 'KAMBINGAN REJO', 1713, 90, 6),
		(22192, 'KARANG KLIWON', 1713, 90, 6),
		(22193, 'KARANG LO', 1713, 90, 6),
		(22194, 'KEBON REJO', 1713, 90, 6),
		(22195, 'KEDAWUNG KULON', 1713, 90, 6),
		(22196, 'KEDAWUNG WETAN', 1713, 90, 6),
		(22197, 'PLOSOSARI', 1713, 90, 6),
		(22198, 'RANU KLINDUNGAN', 1713, 90, 6),
		(22199, 'REBALAS', 1713, 90, 6),
		(22200, 'SUMBER AGUNG', 1713, 90, 6),
		(22201, 'SUMBER DAWESARI', 1713, 90, 6),
		(22202, 'TREWUNG', 1713, 90, 6),
		(22203, 'AMBAL AMBIL', 1714, 90, 6),
		(22204, 'BENERWOJO', 1714, 90, 6),
		(22205, 'COBAN JOYO', 1714, 90, 6),
		(22206, 'KEDEMUNGAN', 1714, 90, 6),
		(22207, 'KEDUNG PENGARON', 1714, 90, 6),
		(22208, 'KEJAYAN', 1714, 90, 6),
		(22209, 'KEPUH', 1714, 90, 6),
		(22210, 'KETANGI REJO (KANGIREJO)', 1714, 90, 6),
		(22211, 'KLANGRONG', 1714, 90, 6),
		(22212, 'KLINTER', 1714, 90, 6),
		(22213, 'KURUNG', 1714, 90, 6),
		(22214, 'LINGGO', 1714, 90, 6),
		(22215, 'LOROKAN', 1714, 90, 6),
		(22216, 'LUWUK', 1714, 90, 6),
		(22217, 'ORO ORO PULE', 1714, 90, 6),
		(22218, 'PACAR KELING', 1714, 90, 6),
		(22219, 'PATEBON', 1714, 90, 6),
		(22220, 'RANDU GONG', 1714, 90, 6),
		(22221, 'SLADI', 1714, 90, 6),
		(22222, 'SUMBER BANTENG', 1714, 90, 6),
		(22223, 'SUMBER SUKO', 1714, 90, 6),
		(22224, 'TANGGULANGIN', 1714, 90, 6),
		(22225, 'TUNDO SORO', 1714, 90, 6),
		(22226, 'WANGKAL WETAN', 1714, 90, 6),
		(22227, 'WRATI', 1714, 90, 6),
		(22228, 'ASEMKANDANG', 1715, 90, 6),
		(22229, 'BENDUNGAN', 1715, 90, 6),
		(22230, 'CURAH DUKUH', 1715, 90, 6),
		(22231, 'DHOMPO', 1715, 90, 6),
		(22232, 'GAMBIR KUNING', 1715, 90, 6),
		(22233, 'GERONGAN', 1715, 90, 6),
		(22234, 'JERUK', 1715, 90, 6),
		(22235, 'KADIPATEN', 1715, 90, 6),
		(22236, 'KALIREJO', 1715, 90, 6),
		(22237, 'KARANG ANYAR', 1715, 90, 6),
		(22238, 'KEBOTOHAN', 1715, 90, 6),
		(22239, 'KLAMPIS REJO', 1715, 90, 6),
		(22240, 'KRATON', 1715, 90, 6),
		(22241, 'MULYO REJO', 1715, 90, 6),
		(22242, 'NGABAR', 1715, 90, 6),
		(22243, 'NGEMPIT', 1715, 90, 6),
		(22244, 'PANEMBAHAN', 1715, 90, 6),
		(22245, 'PATEHAN', 1715, 90, 6),
		(22246, 'PLINGGISAN', 1715, 90, 6),
		(22247, 'PUKUL', 1715, 90, 6),
		(22248, 'PULOKERTO', 1715, 90, 6),
		(22249, 'REJOSARI', 1715, 90, 6),
		(22250, 'SELOTAMBAK', 1715, 90, 6),
		(22251, 'SEMARE', 1715, 90, 6),
		(22252, 'SIDOGIRI', 1715, 90, 6),
		(22253, 'SLAMBRIT', 1715, 90, 6),
		(22254, 'TAMBAKREJO', 1715, 90, 6),
		(22255, 'TAMBAKSARI', 1715, 90, 6),
		(22256, 'ALAS TLOGO', 1716, 90, 6),
		(22257, 'BALUNG ANYAR', 1716, 90, 6),
		(22258, 'BRANANG', 1716, 90, 6),
		(22259, 'GEJUG JATI', 1716, 90, 6),
		(22260, 'JATIREJO', 1716, 90, 6),
		(22261, 'PASINAN', 1716, 90, 6),
		(22262, 'ROWO GEMPOL', 1716, 90, 6),
		(22263, 'SEMEDUSARI', 1716, 90, 6),
		(22264, 'TAMBAK LEKOK', 1716, 90, 6),
		(22265, 'TAMPUNG', 1716, 90, 6),
		(22266, 'WATES', 1716, 90, 6),
		(22267, 'DANDANG GENDIS', 1717, 90, 6),
		(22268, 'KAPASAN', 1717, 90, 6),
		(22269, 'KEDAWANG', 1717, 90, 6),
		(22270, 'MLATEN', 1717, 90, 6),
		(22271, 'NGULING', 1717, 90, 6),
		(22272, 'PENUNGGUL', 1717, 90, 6),
		(22273, 'RANDUATI', 1717, 90, 6),
		(22274, 'SANG ANOM', 1717, 90, 6),
		(22275, 'SEBALONG', 1717, 90, 6),
		(22276, 'SEDARUM', 1717, 90, 6),
		(22277, 'SUDIMULYO', 1717, 90, 6),
		(22278, 'SUMBER ANYAR', 1717, 90, 6),
		(22279, 'WATES TANI', 1717, 90, 6),
		(22280, 'WATUPRAPAT', 1717, 90, 6),
		(22281, 'WOT GALIH', 1717, 90, 6),
		(22282, 'BANJAR KEJEN', 1718, 90, 6),
		(22283, 'BANJARSARI', 1718, 90, 6),
		(22284, 'DURENSEWU', 1718, 90, 6),
		(22285, 'JOGOSARI', 1718, 90, 6),
		(22286, 'KARANG JATI', 1718, 90, 6),
		(22287, 'KEBON WARIS', 1718, 90, 6),
		(22288, 'KEMIRI SEWU', 1718, 90, 6),
		(22289, 'KUTOREJO', 1718, 90, 6),
		(22290, 'NOGOSARI', 1718, 90, 6),
		(22291, 'PANDAAN', 1718, 90, 6),
		(22292, 'PETUNGASRI', 1718, 90, 6),
		(22293, 'PLINTAHAN', 1718, 90, 6),
		(22294, 'SEBANI', 1718, 90, 6),
		(22295, 'SUMBER GEDANG', 1718, 90, 6),
		(22296, 'SUMBER REJO', 1718, 90, 6),
		(22297, 'TAWANG REJO', 1718, 90, 6),
		(22298, 'TUNGGUL WULUNG', 1718, 90, 6),
		(22299, 'WEDORO', 1718, 90, 6),
		(22300, 'BALE REJO', 1719, 90, 6),
		(22301, 'BANGILAN', 1719, 90, 6),
		(22302, 'BUGUL LOR', 1719, 90, 6),
		(22303, 'BUMIAYU', 1719, 90, 6),
		(22304, 'KALI GAMBIR', 1719, 90, 6),
		(22305, 'KALITENGAH', 1719, 90, 6),
		(22306, 'KANDANG SAPI', 1719, 90, 6),
		(22307, 'KARANGANYAR', 1719, 90, 6),
		(22308, 'KEBONSARI', 1719, 90, 6),
		(22309, 'MANDARANREJO', 1719, 90, 6),
		(22310, 'MARGO MULYO', 1719, 90, 6),
		(22311, 'MAYANGAN', 1719, 90, 6),
		(22312, 'NGEMPLAKREJO', 1719, 90, 6),
		(22313, 'PANGGUNG ASRI', 1719, 90, 6),
		(22314, 'PANGGUNG REJO', 1719, 90, 6),
		(22315, 'PANGGUNGREJO', 1719, 90, 6),
		(22316, 'PEKUNCEN', 1719, 90, 6),
		(22317, 'PETAMANAN', 1719, 90, 6),
		(22318, 'SERANG', 1719, 90, 6),
		(22319, 'SUMBER AGUNG', 1719, 90, 6),
		(22320, 'SUMBERSIH', 1719, 90, 6),
		(22321, 'TAMBAAN', 1719, 90, 6),
		(22322, 'TRAJENG', 1719, 90, 6),
		(22323, 'AMPELSARI', 1720, 90, 6),
		(22324, 'CENGKRONG', 1720, 90, 6),
		(22325, 'GALIH', 1720, 90, 6),
		(22326, 'JOGOREPUH', 1720, 90, 6),
		(22327, 'KLAKAH', 1720, 90, 6),
		(22328, 'LEMAHBANG', 1720, 90, 6),
		(22329, 'MANGGUAN', 1720, 90, 6),
		(22330, 'NGANTUNGAN', 1720, 90, 6),
		(22331, 'PASREPAN', 1720, 90, 6),
		(22332, 'PETUNG', 1720, 90, 6),
		(22333, 'POH GADING', 1720, 90, 6),
		(22334, 'POH GEDANG', 1720, 90, 6),
		(22335, 'REJOSALAM', 1720, 90, 6),
		(22336, 'SAPULANTE', 1720, 90, 6),
		(22337, 'SIBON', 1720, 90, 6),
		(22338, 'TAMBAKREJO', 1720, 90, 6),
		(22339, 'TEMPURAN', 1720, 90, 6),
		(22340, 'LOGOWOK', 1721, 90, 6),
		(22341, 'PARAS REJO', 1721, 90, 6),
		(22342, 'PLERET', 1721, 90, 6),
		(22343, 'SUKOREJO', 1721, 90, 6),
		(22344, 'SUNGI KULON', 1721, 90, 6),
		(22345, 'SUNGI WETAN', 1721, 90, 6),
		(22346, 'SUSUKAN REJO', 1721, 90, 6),
		(22347, 'TIDU', 1721, 90, 6),
		(22348, 'WARUNG DOWO', 1721, 90, 6),
		(22349, 'BULUKANDANG', 1722, 90, 6),
		(22350, 'CANDI WATES', 1722, 90, 6),
		(22351, 'DAYUREJO', 1722, 90, 6),
		(22352, 'GAMBIRAN', 1722, 90, 6),
		(22353, 'JATIARJO', 1722, 90, 6),
		(22354, 'KETANIRENG', 1722, 90, 6),
		(22355, 'LEDUG', 1722, 90, 6),
		(22356, 'LUMBANG REJO', 1722, 90, 6),
		(22357, 'PECALUKAN', 1722, 90, 6),
		(22358, 'PRIGEN', 1722, 90, 6),
		(22359, 'SEKARJOHO', 1722, 90, 6),
		(22360, 'SUKOLELO', 1722, 90, 6),
		(22361, 'SUKORENO', 1722, 90, 6),
		(22362, 'WATUAGUNG', 1722, 90, 6),
		(22363, 'BANGUN SARI', 1723, 90, 6),
		(22364, 'BANJARSARI', 1723, 90, 6),
		(22365, 'BLENDUNG', 1723, 90, 6),
		(22366, 'BONGKOT', 1723, 90, 6),
		(22367, 'BRAGOLAN', 1723, 90, 6),
		(22368, 'BRONDONGREJO', 1723, 90, 6),
		(22369, 'BUBUTAN', 1723, 90, 6),
		(22370, 'CANDISARI', 1723, 90, 6),
		(22371, 'CAPANG', 1723, 90, 6),
		(22372, 'CINGKRONG', 1723, 90, 6),
		(22373, 'COWEK', 1723, 90, 6),
		(22374, 'DANYANG', 1723, 90, 6),
		(22375, 'DAWUHAN SENGON', 1723, 90, 6),
		(22376, 'GAJAH REJO', 1723, 90, 6),
		(22377, 'GEDANGAN', 1723, 90, 6),
		(22378, 'GENUKSURAN', 1723, 90, 6),
		(22379, 'GEPARANG', 1723, 90, 6),
		(22380, 'GERBO', 1723, 90, 6),
		(22381, 'GESING', 1723, 90, 6),
		(22382, 'GUYANGAN', 1723, 90, 6),
		(22383, 'JATIKONTAL', 1723, 90, 6),
		(22384, 'JATIMALANG', 1723, 90, 6),
		(22385, 'JATISARI', 1723, 90, 6),
		(22386, 'JENAR KIDUL', 1723, 90, 6),
		(22387, 'JENAR LOR', 1723, 90, 6),
		(22388, 'JENAR WETAN', 1723, 90, 6),
		(22389, 'JOGOBOYO', 1723, 90, 6),
		(22390, 'JOGORESAN', 1723, 90, 6),
		(22391, 'KALONGAN', 1723, 90, 6),
		(22392, 'KANDANGAN', 1723, 90, 6),
		(22393, 'KARANGANYAR', 1723, 90, 6),
		(22394, 'KARANGANYAR', 1723, 90, 6),
		(22395, 'KARANGMULYO', 1723, 90, 6),
		(22396, 'KARANGSARI', 1723, 90, 6),
		(22397, 'KARYADADI', 1723, 90, 6),
		(22398, 'KEBONSARI', 1723, 90, 6),
		(22399, 'KEDUNGREJO', 1723, 90, 6),
		(22400, 'KEDUREN', 1723, 90, 6),
		(22401, 'KENTENGREJO', 1723, 90, 6),
		(22402, 'KEPONGGOK', 1723, 90, 6),
		(22403, 'KERTOSARI', 1723, 90, 6),
		(22404, 'KESUGIHAN', 1723, 90, 6),
		(22405, 'KETANGI', 1723, 90, 6),
		(22406, 'KURIPAN', 1723, 90, 6),
		(22407, 'LEBAK REJO', 1723, 90, 6),
		(22408, 'MANGUN HARJO', 1723, 90, 6),
		(22409, 'MARDIHARJO', 1723, 90, 6),
		(22410, 'NAMBUHAN', 1723, 90, 6),
		(22411, 'NAMPU', 1723, 90, 6),
		(22412, 'NAMPUREJO', 1723, 90, 6),
		(22413, 'NGEMBAK', 1723, 90, 6),
		(22414, 'NGLOBAR', 1723, 90, 6),
		(22415, 'NGRAJI', 1723, 90, 6),
		(22416, 'PAGARSARI', 1723, 90, 6),
		(22417, 'PAREREJO', 1723, 90, 6),
		(22418, 'PLANDI', 1723, 90, 6),
		(22419, 'PUCANG SARI', 1723, 90, 6),
		(22420, 'PULOREJO', 1723, 90, 6),
		(22421, 'PUNDENSARI', 1723, 90, 6),
		(22422, 'PURWA KARYA', 1723, 90, 6),
		(22423, 'PURWODADI', 1723, 90, 6),
		(22424, 'PURWODADI', 1723, 90, 6),
		(22425, 'PURWODADI', 1723, 90, 6),
		(22426, 'PURWODADI', 1723, 90, 6),
		(22427, 'PURWOSARI', 1723, 90, 6),
		(22428, 'PUTAT', 1723, 90, 6),
		(22429, 'REJOSARI', 1723, 90, 6),
		(22430, 'SADAR KARYA', 1723, 90, 6),
		(22431, 'SEMUT', 1723, 90, 6),
		(22432, 'SENDANGSARI', 1723, 90, 6),
		(22433, 'SENTUL', 1723, 90, 6),
		(22434, 'SIDOHARJO', 1723, 90, 6),
		(22435, 'SUKOMANAH', 1723, 90, 6),
		(22436, 'SUMBEREJO', 1723, 90, 6),
		(22437, 'SUMBERSARI', 1723, 90, 6),
		(22438, 'TAMBAK SARI', 1723, 90, 6),
		(22439, 'TEGALAREN', 1723, 90, 6),
		(22440, 'TLOGOREJO', 1723, 90, 6),
		(22441, 'TRIKARYA', 1723, 90, 6),
		(22442, 'WARUKARANGANYAR', 1723, 90, 6),
		(22443, 'WATUKURO', 1723, 90, 6),
		(22444, 'BALEDONO', 1724, 90, 6),
		(22445, 'BRENGGONG', 1724, 90, 6),
		(22446, 'CANGKREP KIDUL', 1724, 90, 6),
		(22447, 'CANGKREP LOR', 1724, 90, 6),
		(22448, 'DONORATI', 1724, 90, 6),
		(22449, 'DOPLANG', 1724, 90, 6),
		(22450, 'GANGGENG', 1724, 90, 6),
		(22451, 'KEBONAGUNG', 1724, 90, 6),
		(22452, 'KEDUNG SARI', 1724, 90, 6),
		(22453, 'KESENENG', 1724, 90, 6),
		(22454, 'MRANTI', 1724, 90, 6),
		(22455, 'MUDAL', 1724, 90, 6),
		(22456, 'PACEKELAN', 1724, 90, 6),
		(22457, 'PADUROSO', 1724, 90, 6),
		(22458, 'PANGENJURU TENGAH', 1724, 90, 6),
		(22459, 'PANGENREJO', 1724, 90, 6),
		(22460, 'PLIPIR', 1724, 90, 6),
		(22461, 'POHJENTREK', 1724, 90, 6),
		(22462, 'PURUTREJO', 1724, 90, 6),
		(22463, 'PURWOREJO', 1724, 90, 6),
		(22464, 'PURWOREJO', 1724, 90, 6),
		(22465, 'SEKARGADUNG', 1724, 90, 6),
		(22466, 'SEMAWUNG', 1724, 90, 6),
		(22467, 'SIDO MULYO', 1724, 90, 6),
		(22468, 'SIDOREJO', 1724, 90, 6),
		(22469, 'SINDURJAN', 1724, 90, 6),
		(22470, 'SUDIMORO', 1724, 90, 6),
		(22471, 'TAMBAKREJO', 1724, 90, 6),
		(22472, 'TEMBOKREJO', 1724, 90, 6),
		(22473, 'WIROGUNAN', 1724, 90, 6),
		(22474, 'WONOROTO', 1724, 90, 6),
		(22475, 'WONOTULUS', 1724, 90, 6),
		(22476, 'BAKALAN', 1725, 90, 6),
		(22477, 'CENDONO', 1725, 90, 6),
		(22478, 'DONAN', 1725, 90, 6),
		(22479, 'GAPLUK', 1725, 90, 6),
		(22480, 'GIRIASIH', 1725, 90, 6),
		(22481, 'GIRICAHYO', 1725, 90, 6),
		(22482, 'GIRIJATI', 1725, 90, 6),
		(22483, 'GIRIPURWO', 1725, 90, 6),
		(22484, 'GIRITIRTO', 1725, 90, 6),
		(22485, 'KALIOMBO', 1725, 90, 6),
		(22486, 'KARANG REJO', 1725, 90, 6),
		(22487, 'KAYOMAN', 1725, 90, 6),
		(22488, 'KERTOSARI', 1725, 90, 6),
		(22489, 'KUNIRAN', 1725, 90, 6),
		(22490, 'MARTOPURO', 1725, 90, 6),
		(22491, 'NGREJENG', 1725, 90, 6),
		(22492, 'PAGER', 1725, 90, 6),
		(22493, 'PELEM', 1725, 90, 6),
		(22494, 'POJOK', 1725, 90, 6),
		(22495, 'PUCANGSARI', 1725, 90, 6),
		(22496, 'PUNGGUR', 1725, 90, 6),
		(22497, 'PURWOSARI', 1725, 90, 6),
		(22498, 'PURWOSARI', 1725, 90, 6),
		(22499, 'SEDAHKIDUL', 1725, 90, 6),
		(22500, 'SEKAR MOJO', 1725, 90, 6),
		(22501, 'SENGON AGUNG', 1725, 90, 6),
		(22502, 'SUKODERMO', 1725, 90, 6),
		(22503, 'SUMBER REJO', 1725, 90, 6),
		(22504, 'SUMBER SUKO', 1725, 90, 6),
		(22505, 'TEJOWANGI', 1725, 90, 6),
		(22506, 'TINUMPUK', 1725, 90, 6),
		(22507, 'TLATAH', 1725, 90, 6),
		(22508, 'JANGJANGWULUNG', 1726, 90, 6),
		(22509, 'JIMBARAN', 1726, 90, 6),
		(22510, 'KEDUWUNG', 1726, 90, 6),
		(22511, 'KEMIRI', 1726, 90, 6),
		(22512, 'PALANG SARI', 1726, 90, 6),
		(22513, 'PUSPO', 1726, 90, 6),
		(22514, 'PUSUNG MALANG', 1726, 90, 6),
		(22515, 'ARJOSARI', 1727, 90, 6),
		(22516, 'BANJAREJO', 1727, 90, 6),
		(22517, 'BENDA ASRI', 1727, 90, 6),
		(22518, 'GEMPOL', 1727, 90, 6),
		(22519, 'JARANGAN', 1727, 90, 6),
		(22520, 'JATIREJO', 1727, 90, 6),
		(22521, 'JINTEL', 1727, 90, 6),
		(22522, 'KARANG PANDAN', 1727, 90, 6),
		(22523, 'KAWIS REJO', 1727, 90, 6),
		(22524, 'KEDUNG BAKO', 1727, 90, 6),
		(22525, 'KEDUNG PADANG', 1727, 90, 6),
		(22526, 'KEMANTREN REJO', 1727, 90, 6),
		(22527, 'KETEGAN', 1727, 90, 6),
		(22528, 'KLAGEN', 1727, 90, 6),
		(22529, 'MANIK REJO', 1727, 90, 6),
		(22530, 'MLORAH', 1727, 90, 6),
		(22531, 'MOJOREMBUN', 1727, 90, 6),
		(22532, 'MUNGKUNG', 1727, 90, 6),
		(22533, 'MUSIR KIDUL', 1727, 90, 6),
		(22534, 'MUSIR LOR', 1727, 90, 6),
		(22535, 'NGADIBOYO', 1727, 90, 6),
		(22536, 'NGANGKATAN', 1727, 90, 6),
		(22537, 'PANDAN REJO', 1727, 90, 6),
		(22538, 'PATUGURAN', 1727, 90, 6),
		(22539, 'PUH KEREP', 1727, 90, 6),
		(22540, 'REJOSO', 1727, 90, 6),
		(22541, 'REJOSO KIDUL', 1727, 90, 6),
		(22542, 'REJOSOLOR', 1727, 90, 6),
		(22543, 'SADENG REJO', 1727, 90, 6),
		(22544, 'SAMBI KEREP', 1727, 90, 6),
		(22545, 'SAMBIREJO', 1727, 90, 6),
		(22546, 'SEGORO PURO', 1727, 90, 6),
		(22547, 'SETREN', 1727, 90, 6),
		(22548, 'SIDOKARE', 1727, 90, 6),
		(22549, 'SUKOREJO', 1727, 90, 6),
		(22550, 'TALANG', 1727, 90, 6),
		(22551, 'TALUN', 1727, 90, 6),
		(22552, 'TOYANING', 1727, 90, 6),
		(22553, 'TRITIK', 1727, 90, 6),
		(22554, 'WENGKAL', 1727, 90, 6),
		(22555, 'BANTARBARANG', 1728, 90, 6),
		(22556, 'BODAS KARANGJATI', 1728, 90, 6),
		(22557, 'GEDANGAN', 1728, 90, 6),
		(22558, 'GEGUNUNG KULON', 1728, 90, 6),
		(22559, 'GEGUNUNG WETAN', 1728, 90, 6),
		(22560, 'GENENG WARU', 1728, 90, 6),
		(22561, 'GUNUNGWULED', 1728, 90, 6),
		(22562, 'KABONGAN KIDUL', 1728, 90, 6),
		(22563, 'KABONGAN LOR', 1728, 90, 6),
		(22564, 'KALISAT', 1728, 90, 6),
		(22565, 'KANIGORO', 1728, 90, 6),
		(22566, 'KARANGBAWANG', 1728, 90, 6),
		(22567, 'KASREMAN', 1728, 90, 6),
		(22568, 'KEDUNG BANTENG', 1728, 90, 6),
		(22569, 'KEDUNGREJO', 1728, 90, 6),
		(22570, 'KETANGGI', 1728, 90, 6),
		(22571, 'KRENGIH', 1728, 90, 6),
		(22572, 'KUMENDUNG', 1728, 90, 6),
		(22573, 'KUTOHARJO', 1728, 90, 6),
		(22574, 'LETEH', 1728, 90, 6),
		(22575, 'LOSARI', 1728, 90, 6),
		(22576, 'MAGERSARI', 1728, 90, 6),
		(22577, 'MAKAM', 1728, 90, 6),
		(22578, 'MOJOPARON', 1728, 90, 6),
		(22579, 'MONDOTEKO', 1728, 90, 6),
		(22580, 'NGADEM', 1728, 90, 6),
		(22581, 'NGOTET', 1728, 90, 6),
		(22582, 'OROBULU', 1728, 90, 6),
		(22583, 'OROOMBO KULON', 1728, 90, 6),
		(22584, 'OROOMBO WETAN', 1728, 90, 6),
		(22585, 'PACAR', 1728, 90, 6),
		(22586, 'PADARAN', 1728, 90, 6),
		(22587, 'PAJARAN', 1728, 90, 6),
		(22588, 'PANDEAN', 1728, 90, 6),
		(22589, 'PANDEAN', 1728, 90, 6),
		(22590, 'PANUSUPAN', 1728, 90, 6),
		(22591, 'PASARBANGGI', 1728, 90, 6),
		(22592, 'PEJANGKUNGAN', 1728, 90, 6),
		(22593, 'PEKOREN', 1728, 90, 6),
		(22594, 'PULO', 1728, 90, 6),
		(22595, 'PUNJULHARJO', 1728, 90, 6),
		(22596, 'REMBANG', 1728, 90, 6),
		(22597, 'SAWAHAN', 1728, 90, 6),
		(22598, 'SIDOWAYAH', 1728, 90, 6),
		(22599, 'SIYAR', 1728, 90, 6),
		(22600, 'SRIDADI', 1728, 90, 6),
		(22601, 'SUKOHARJO', 1728, 90, 6),
		(22602, 'SUMAMPIR', 1728, 90, 6),
		(22603, 'SUMBER GLAGAH', 1728, 90, 6),
		(22604, 'SUMBERJO', 1728, 90, 6),
		(22605, 'TAMPUNG', 1728, 90, 6),
		(22606, 'TANALUM', 1728, 90, 6),
		(22607, 'TANJUNGSARI', 1728, 90, 6),
		(22608, 'TASIKAGUNG', 1728, 90, 6),
		(22609, 'TIREMAN', 1728, 90, 6),
		(22610, 'TLOGOMOJO', 1728, 90, 6),
		(22611, 'TRITUNGGAL', 1728, 90, 6),
		(22612, 'TURUSGEDE', 1728, 90, 6),
		(22613, 'WANOGARA KULON', 1728, 90, 6),
		(22614, 'WANOGARA WETAN', 1728, 90, 6),
		(22615, 'WARU', 1728, 90, 6),
		(22616, 'WETON', 1728, 90, 6),
		(22617, 'WLAHAR', 1728, 90, 6),
		(22618, 'BANGUNREJO', 1729, 90, 6),
		(22619, 'BLITAR', 1729, 90, 6),
		(22620, 'BRINGINSARI', 1729, 90, 6),
		(22621, 'CANDI BINANGUN', 1729, 90, 6),
		(22622, 'CURAH REJO', 1729, 90, 6),
		(22623, 'DAMARJATI', 1729, 90, 6),
		(22624, 'DUKUH SARI', 1729, 90, 6),
		(22625, 'GANDUKEPUH', 1729, 90, 6),
		(22626, 'GEGERAN', 1729, 90, 6),
		(22627, 'GELANGLOR', 1729, 90, 6),
		(22628, 'GENTING GUNUNG', 1729, 90, 6),
		(22629, 'GLAGAH SARI', 1729, 90, 6),
		(22630, 'GOLAN', 1729, 90, 6),
		(22631, 'GUNTING', 1729, 90, 6),
		(22632, 'HARJODOWO', 1729, 90, 6),
		(22633, 'KALIBOGOR', 1729, 90, 6),
		(22634, 'KALIMALANG', 1729, 90, 6),
		(22635, 'KALIPAKIS', 1729, 90, 6),
		(22636, 'KALIREJO', 1729, 90, 6),
		(22637, 'KARANG SARI', 1729, 90, 6),
		(22638, 'KARANGLO LOR', 1729, 90, 6),
		(22639, 'KARANGSONO', 1729, 90, 6),
		(22640, 'KEBUMEN', 1729, 90, 6),
		(22641, 'KEDUNG BANTENG', 1729, 90, 6),
		(22642, 'KENDURUAN', 1729, 90, 6),
		(22643, 'KRANGGAN', 1729, 90, 6),
		(22644, 'LECARI', 1729, 90, 6),
		(22645, 'LEMAHBANG', 1729, 90, 6),
		(22646, 'LENGKONG', 1729, 90, 6),
		(22647, 'MOJOTENGAH', 1729, 90, 6),
		(22648, 'MOROSARI', 1729, 90, 6),
		(22649, 'MULYOSARI', 1729, 90, 6),
		(22650, 'NAMBANGREJO', 1729, 90, 6),
		(22651, 'NAMPAN', 1729, 90, 6),
		(22652, 'NGADIMULYO', 1729, 90, 6),
		(22653, 'NGADIWARNO', 1729, 90, 6),
		(22654, 'NGARGOSARI', 1729, 90, 6),
		(22655, 'PAKUKERTO', 1729, 90, 6),
		(22656, 'PAKUNDEN', 1729, 90, 6),
		(22657, 'PERON', 1729, 90, 6),
		(22658, 'PESAREN', 1729, 90, 6),
		(22659, 'PRAJEGAN', 1729, 90, 6),
		(22660, 'PURWO SARI', 1729, 90, 6),
		(22661, 'SEBANDUNG', 1729, 90, 6),
		(22662, 'SELOKATON', 1729, 90, 6),
		(22663, 'SERANGAN', 1729, 90, 6),
		(22664, 'SIDOREJO', 1729, 90, 6),
		(22665, 'SRAGI', 1729, 90, 6),
		(22666, 'SUKORAME', 1729, 90, 6),
		(22667, 'SUKOREJO', 1729, 90, 6),
		(22668, 'SUKOREJO', 1729, 90, 6),
		(22669, 'SUKOREJO', 1729, 90, 6),
		(22670, 'SUKOREJO', 1729, 90, 6),
		(22671, 'SUWAYUWO', 1729, 90, 6),
		(22672, 'TAMAN REJO', 1729, 90, 6),
		(22673, 'TAMPING WINARNO', 1729, 90, 6),
		(22674, 'TANJUNG SARI', 1729, 90, 6),
		(22675, 'TANJUNGARUM', 1729, 90, 6),
		(22676, 'TLUMPU', 1729, 90, 6),
		(22677, 'TRIMULYO', 1729, 90, 6),
		(22678, 'TURI', 1729, 90, 6),
		(22679, 'WONOKERTO', 1729, 90, 6),
		(22680, 'BALEDONO', 1730, 90, 6),
		(22681, 'KANDANGAN', 1730, 90, 6),
		(22682, 'MOROREJO', 1730, 90, 6),
		(22683, 'NGADIWONO', 1730, 90, 6),
		(22684, 'PODOKOYO', 1730, 90, 6),
		(22685, 'SEDAENG', 1730, 90, 6),
		(22686, 'TOSARI', 1730, 90, 6),
		(22687, 'WONOKITRI', 1730, 90, 6),
		(22688, 'ANDONO SARI', 1731, 90, 6),
		(22689, 'BLARANG', 1731, 90, 6),
		(22690, 'GENDRO', 1731, 90, 6),
		(22691, 'KALI PUCANG', 1731, 90, 6),
		(22692, 'KAYU KEBEK', 1731, 90, 6),
		(22693, 'NGADIREJO', 1731, 90, 6),
		(22694, 'NGEMBAL', 1731, 90, 6),
		(22695, 'PUNGGING', 1731, 90, 6),
		(22696, 'SUMBER PITU', 1731, 90, 6),
		(22697, 'TLOGOSARI', 1731, 90, 6),
		(22698, 'TUTUR', 1731, 90, 6),
		(22699, 'WONOSARI', 1731, 90, 6),
		(22700, 'BANDARAN', 1732, 90, 6),
		(22701, 'GADING', 1732, 90, 6),
		(22702, 'JELADRI', 1732, 90, 6),
		(22703, 'KANDUNG', 1732, 90, 6),
		(22704, 'KARANG TENGAH', 1732, 90, 6),
		(22705, 'KEDUNG REJO', 1732, 90, 6),
		(22706, 'LEBAK', 1732, 90, 6),
		(22707, 'MENDALAN', 1732, 90, 6),
		(22708, 'MENYARIK', 1732, 90, 6),
		(22709, 'MINGGIR', 1732, 90, 6),
		(22710, 'PENATAAN', 1732, 90, 6),
		(22711, 'PRODO', 1732, 90, 6),
		(22712, 'SIDEPAN', 1732, 90, 6),
		(22713, 'SRUWI', 1732, 90, 6),
		(22714, 'SUMBER REJO', 1732, 90, 6),
		(22715, 'UMBULAN', 1732, 90, 6),
		(22716, 'WINONGAN KIDUL', 1732, 90, 6),
		(22717, 'WINONGAN LOR', 1732, 90, 6),
		(22718, 'COBAN BLIMBING', 1733, 90, 6),
		(22719, 'JATI GUNTING', 1733, 90, 6),
		(22720, 'KARANG ASEM', 1733, 90, 6),
		(22721, 'KARANG MENGGAH', 1733, 90, 6),
		(22722, 'KARANGJATI ANYAR', 1733, 90, 6),
		(22723, 'KARANGSONO', 1733, 90, 6),
		(22724, 'KENDANG DUKUH', 1733, 90, 6),
		(22725, 'KLUWUT', 1733, 90, 6),
		(22726, 'LEBAKSARI', 1733, 90, 6),
		(22727, 'PAKIJANGAN', 1733, 90, 6),
		(22728, 'REBONO', 1733, 90, 6),
		(22729, 'SAMBISIRAH', 1733, 90, 6),
		(22730, 'TAMANSARI', 1733, 90, 6),
		(22731, 'WONOREJO', 1733, 90, 6),
		(22732, 'WONOSARI', 1733, 90, 6),
		(22733, 'BABADAN', 1734, 91, 6),
		(22734, 'BARENG', 1734, 91, 6),
		(22735, 'CEKOK', 1734, 91, 6),
		(22736, 'GUPOLO', 1734, 91, 6),
		(22737, 'JAPAN', 1734, 91, 6),
		(22738, 'KADIPATEN', 1734, 91, 6),
		(22739, 'KERTOSARI', 1734, 91, 6),
		(22740, 'LEMBAH', 1734, 91, 6),
		(22741, 'NGUNUT', 1734, 91, 6),
		(22742, 'PATIHAN WETAN', 1734, 91, 6),
		(22743, 'POLOREJO', 1734, 91, 6),
		(22744, 'PONDOK', 1734, 91, 6),
		(22745, 'PURWOSARI', 1734, 91, 6),
		(22746, 'SUKOSARI', 1734, 91, 6),
		(22747, 'TRISONO', 1734, 91, 6),
		(22748, 'BADEGAN', 1735, 91, 6),
		(22749, 'BANDARALIM', 1735, 91, 6),
		(22750, 'BITING', 1735, 91, 6),
		(22751, 'DAYAKAN', 1735, 91, 6),
		(22752, 'KAPURAN', 1735, 91, 6),
		(22753, 'KARANGAN', 1735, 91, 6),
		(22754, 'KARANGJOHO', 1735, 91, 6),
		(22755, 'TANJUNG GUNUNG', 1735, 91, 6),
		(22756, 'TANJUNGREJO', 1735, 91, 6),
		(22757, 'WATUBONANG', 1735, 91, 6),
		(22758, 'BAJANG', 1736, 91, 6),
		(22759, 'BALONG', 1736, 91, 6),
		(22760, 'BULAK', 1736, 91, 6),
		(22761, 'BULUKIDUL', 1736, 91, 6),
		(22762, 'DADAPAN', 1736, 91, 6),
		(22763, 'JALEN', 1736, 91, 6),
		(22764, 'KARANGAN', 1736, 91, 6),
		(22765, 'KARANGMOJO', 1736, 91, 6),
		(22766, 'KARANGPATIHAN', 1736, 91, 6),
		(22767, 'MUNENG', 1736, 91, 6),
		(22768, 'NGAMPEL', 1736, 91, 6),
		(22769, 'NGENDUT', 1736, 91, 6),
		(22770, 'NGRAKET', 1736, 91, 6),
		(22771, 'NGUMPUL', 1736, 91, 6),
		(22772, 'PANDAK', 1736, 91, 6),
		(22773, 'PURWOREJO', 1736, 91, 6),
		(22774, 'SEDARAT', 1736, 91, 6),
		(22775, 'SINGKIL', 1736, 91, 6),
		(22776, 'SUMBEREJO', 1736, 91, 6),
		(22777, 'TATUNG', 1736, 91, 6),
		(22778, 'BANCAR', 1737, 91, 6),
		(22779, 'BEDI KULON', 1737, 91, 6),
		(22780, 'BEDI WETAN', 1737, 91, 6),
		(22781, 'BEKARE', 1737, 91, 6),
		(22782, 'BELANG (BLANG)', 1737, 91, 6),
		(22783, 'BUNGKAL', 1737, 91, 6),
		(22784, 'BUNGU', 1737, 91, 6),
		(22785, 'KALISAT', 1737, 91, 6),
		(22786, 'KETONGGO', 1737, 91, 6),
		(22787, 'KORIPAN', 1737, 91, 6),
		(22788, 'KUNTI', 1737, 91, 6),
		(22789, 'KUPUK', 1737, 91, 6),
		(22790, 'KWAJON', 1737, 91, 6),
		(22791, 'MUNGGU', 1737, 91, 6),
		(22792, 'NAMBAK', 1737, 91, 6),
		(22793, 'PADAS', 1737, 91, 6),
		(22794, 'PAGER', 1737, 91, 6),
		(22795, 'PELEM', 1737, 91, 6),
		(22796, 'SAMBILAWANG', 1737, 91, 6),
		(22797, 'BLEMBEM', 1738, 91, 6),
		(22798, 'BRINGINAN', 1738, 91, 6),
		(22799, 'BULU LOR', 1738, 91, 6),
		(22800, 'JAMBON', 1738, 91, 6),
		(22801, 'JONGGOL', 1738, 91, 6),
		(22802, 'KARANGLO KIDUL', 1738, 91, 6),
		(22803, 'KREBET', 1738, 91, 6),
		(22804, 'MENANG', 1738, 91, 6),
		(22805, 'POKO', 1738, 91, 6),
		(22806, 'PULOSARI', 1738, 91, 6),
		(22807, 'SENDANG', 1738, 91, 6),
		(22808, 'SIDOHARJO', 1738, 91, 6),
		(22809, 'SRANDIL', 1738, 91, 6),
		(22810, 'JENANGAN', 1739, 91, 6),
		(22811, 'JIMBE', 1739, 91, 6),
		(22812, 'KEMIRI', 1739, 91, 6),
		(22813, 'MRICAN', 1739, 91, 6),
		(22814, 'NGLAYANG', 1739, 91, 6),
		(22815, 'NGRUPIT', 1739, 91, 6),
		(22816, 'PANJENG', 1739, 91, 6),
		(22817, 'PARINGAN', 1739, 91, 6),
		(22818, 'PINTU', 1739, 91, 6),
		(22819, 'PLALANGAN', 1739, 91, 6),
		(22820, 'SEDAH', 1739, 91, 6),
		(22821, 'SEMANDING', 1739, 91, 6),
		(22822, 'SETONO', 1739, 91, 6),
		(22823, 'SINGOSAREN', 1739, 91, 6),
		(22824, 'SRATEN', 1739, 91, 6),
		(22825, 'TANJUNG SARI', 1739, 91, 6),
		(22826, 'WATES', 1739, 91, 6),
		(22827, 'BANJARSARI', 1740, 91, 6),
		(22828, 'BENDUNG', 1740, 91, 6),
		(22829, 'BUMIJO', 1740, 91, 6),
		(22830, 'CANDEN', 1740, 91, 6),
		(22831, 'CANGGU', 1740, 91, 6),
		(22832, 'COKRODININGRATAN', 1740, 91, 6),
		(22833, 'COPER', 1740, 91, 6),
		(22834, 'GOWONGAN', 1740, 91, 6),
		(22835, 'JETIS', 1740, 91, 6),
		(22836, 'JETIS', 1740, 91, 6),
		(22837, 'JOLOTUNDO', 1740, 91, 6),
		(22838, 'JOSARI', 1740, 91, 6),
		(22839, 'KARANGGEBANG', 1740, 91, 6),
		(22840, 'KRADENAN (KRADINAN)', 1740, 91, 6),
		(22841, 'KUPANG', 1740, 91, 6),
		(22842, 'KUTU KULON', 1740, 91, 6),
		(22843, 'KUTU WETAN', 1740, 91, 6),
		(22844, 'LAKARDOWO', 1740, 91, 6),
		(22845, 'MLIRIP', 1740, 91, 6),
		(22846, 'MOJOLEBAK', 1740, 91, 6),
		(22847, 'MOJOMATI', 1740, 91, 6),
		(22848, 'MOJOREJO', 1740, 91, 6),
		(22849, 'MOJOREJO', 1740, 91, 6),
		(22850, 'NGABAR', 1740, 91, 6),
		(22851, 'NGASINAN', 1740, 91, 6),
		(22852, 'PARINGAN', 1740, 91, 6),
		(22853, 'PATALAN', 1740, 91, 6),
		(22854, 'PENOMPO', 1740, 91, 6),
		(22855, 'PERNING', 1740, 91, 6),
		(22856, 'SAWO', 1740, 91, 6),
		(22857, 'SIDOREJO', 1740, 91, 6),
		(22858, 'SUMBER AGUNG', 1740, 91, 6),
		(22859, 'TEGALSARI', 1740, 91, 6),
		(22860, 'TRIMULYO', 1740, 91, 6),
		(22861, 'TURI', 1740, 91, 6),
		(22862, 'WINONG', 1740, 91, 6),
		(22863, 'WONOKETRO', 1740, 91, 6),
		(22864, 'BAJANG', 1741, 91, 6),
		(22865, 'CANDI', 1741, 91, 6),
		(22866, 'GANDU', 1741, 91, 6),
		(22867, 'GONTOR', 1741, 91, 6),
		(22868, 'JABUNG', 1741, 91, 6),
		(22869, 'JORESAN', 1741, 91, 6),
		(22870, 'KAPONAN', 1741, 91, 6),
		(22871, 'MLARAK', 1741, 91, 6),
		(22872, 'NGLUMPANG', 1741, 91, 6),
		(22873, 'NGRUKEM (NGERUKEM)', 1741, 91, 6),
		(22874, 'SERANGAN', 1741, 91, 6),
		(22875, 'SIWALAN', 1741, 91, 6),
		(22876, 'SUREN', 1741, 91, 6),
		(22877, 'TOTOKAN', 1741, 91, 6),
		(22878, 'TUGU (TUGUREJO)', 1741, 91, 6),
		(22879, 'GONDOWIDO', 1742, 91, 6),
		(22880, 'NGEBEL', 1742, 91, 6),
		(22881, 'NGROGUNG', 1742, 91, 6),
		(22882, 'PUPUS', 1742, 91, 6),
		(22883, 'SAHANG', 1742, 91, 6),
		(22884, 'SEMPU', 1742, 91, 6),
		(22885, 'TALUN', 1742, 91, 6),
		(22886, 'WAGIR LOR', 1742, 91, 6),
		(22887, 'BAOSAN KIDUL', 1743, 91, 6),
		(22888, 'BAOSAN LOR', 1743, 91, 6),
		(22889, 'BINADE', 1743, 91, 6),
		(22890, 'CEPOKO', 1743, 91, 6),
		(22891, 'GEDANGAN', 1743, 91, 6),
		(22892, 'MRAYAN (MERAYAN)', 1743, 91, 6),
		(22893, 'NGRAYUN (NGERAYUN)', 1743, 91, 6),
		(22894, 'SELUR', 1743, 91, 6),
		(22895, 'SENDANG', 1743, 91, 6),
		(22896, 'TEMON', 1743, 91, 6),
		(22897, 'WONODADI', 1743, 91, 6),
		(22898, 'BANGUNSARI', 1744, 91, 6),
		(22899, 'BANYUDONO', 1744, 91, 6),
		(22900, 'BEDURI', 1744, 91, 6),
		(22901, 'BROTONEGARAN', 1744, 91, 6),
		(22902, 'COKROMENGGALAN', 1744, 91, 6),
		(22903, 'JINGGLONG', 1744, 91, 6),
		(22904, 'KAUMAN', 1744, 91, 6),
		(22905, 'KENITEN', 1744, 91, 6),
		(22906, 'KEPATIHAN', 1744, 91, 6),
		(22907, 'MANGKUJAYAN', 1744, 91, 6),
		(22908, 'NOLOGATEN', 1744, 91, 6),
		(22909, 'PAJU', 1744, 91, 6),
		(22910, 'PAKUNDEN', 1744, 91, 6),
		(22911, 'PINGGIR SARI', 1744, 91, 6),
		(22912, 'PURBOSUMAN', 1744, 91, 6),
		(22913, 'SURODIKRAMAN', 1744, 91, 6),
		(22914, 'TAMAN ARUM', 1744, 91, 6),
		(22915, 'TAMBAK BAYAN', 1744, 91, 6),
		(22916, 'TONATAN', 1744, 91, 6),
		(22917, 'BANJARJO', 1745, 91, 6),
		(22918, 'BARENG', 1745, 91, 6),
		(22919, 'KRISIK', 1745, 91, 6),
		(22920, 'PUDAKKULON', 1745, 91, 6),
		(22921, 'PUDAKWETAN', 1745, 91, 6),
		(22922, 'TAMBANG', 1745, 91, 6),
		(22923, 'BANARAN', 1746, 91, 6),
		(22924, 'BEDRUG', 1746, 91, 6),
		(22925, 'BEKIRING', 1746, 91, 6),
		(22926, 'KARANGPATIHAN', 1746, 91, 6),
		(22927, 'KESUGIHAN', 1746, 91, 6),
		(22928, 'MUNGGUNG', 1746, 91, 6),
		(22929, 'PATIK', 1746, 91, 6),
		(22930, 'PLUNTURAN', 1746, 91, 6),
		(22931, 'POMAHAN', 1746, 91, 6),
		(22932, 'PULUNG', 1746, 91, 6),
		(22933, 'PULUNG MERDIKO', 1746, 91, 6),
		(22934, 'SERAG', 1746, 91, 6),
		(22935, 'SIDOHARJO', 1746, 91, 6),
		(22936, 'SINGGAHAN', 1746, 91, 6),
		(22937, 'TEGALREJO', 1746, 91, 6),
		(22938, 'WAGIRKIDUL', 1746, 91, 6),
		(22939, 'WAYANG', 1746, 91, 6),
		(22940, 'WOTAN', 1746, 91, 6),
		(22941, 'BANCANGAN', 1747, 91, 6),
		(22942, 'BANGSALAN', 1747, 91, 6),
		(22943, 'BEDINGIN', 1747, 91, 6),
		(22944, 'BESUKI', 1747, 91, 6),
		(22945, 'BULU', 1747, 91, 6),
		(22946, 'CAMPUREJO', 1747, 91, 6),
		(22947, 'CAMPURSARI', 1747, 91, 6),
		(22948, 'GAJAH', 1747, 91, 6),
		(22949, 'JRAKAH', 1747, 91, 6),
		(22950, 'KEMUNING', 1747, 91, 6),
		(22951, 'MAGUWAN', 1747, 91, 6),
		(22952, 'NGADISANAN', 1747, 91, 6),
		(22953, 'NGLEWAN (NGELEWAN)', 1747, 91, 6),
		(22954, 'SAMBIT', 1747, 91, 6),
		(22955, 'WILANGAN', 1747, 91, 6),
		(22956, 'WRINGINANOM', 1747, 91, 6),
		(22957, 'CARANG REJO', 1748, 91, 6),
		(22958, 'GELANGKULON', 1748, 91, 6),
		(22959, 'GLINGGANG', 1748, 91, 6),
		(22960, 'JENANGAN', 1748, 91, 6),
		(22961, 'KARANG WALUH', 1748, 91, 6),
		(22962, 'KUNTI', 1748, 91, 6),
		(22963, 'NGLURUP', 1748, 91, 6),
		(22964, 'PAGERUKIR', 1748, 91, 6),
		(22965, 'POHIJO', 1748, 91, 6),
		(22966, 'RINGINPUTIH', 1748, 91, 6),
		(22967, 'SAMPUNG', 1748, 91, 6),
		(22968, 'TULUNG', 1748, 91, 6),
		(22969, 'BONDRANG', 1749, 91, 6),
		(22970, 'GROGOL', 1749, 91, 6),
		(22971, 'KETRO', 1749, 91, 6),
		(22972, 'KORI', 1749, 91, 6),
		(22973, 'NGINDENG', 1749, 91, 6),
		(22974, 'PANGKAL', 1749, 91, 6),
		(22975, 'PRAYUNGAN', 1749, 91, 6),
		(22976, 'SAWOO', 1749, 91, 6),
		(22977, 'SRITI (SERINTI)', 1749, 91, 6),
		(22978, 'TEMON', 1749, 91, 6),
		(22979, 'TEMPURAN', 1749, 91, 6),
		(22980, 'TUGUREJO', 1749, 91, 6),
		(22981, 'TUMPAKPELEM', 1749, 91, 6),
		(22982, 'TUMPUK', 1749, 91, 6),
		(22983, 'BETON', 1750, 91, 6),
		(22984, 'BRAHU', 1750, 91, 6),
		(22985, 'DEMANGAN', 1750, 91, 6),
		(22986, 'JARAK', 1750, 91, 6),
		(22987, 'KEPUH RUBUH', 1750, 91, 6),
		(22988, 'MADUSARI', 1750, 91, 6),
		(22989, 'MANGUNSUMAN (NGAMUNSUMAN)', 1750, 91, 6),
		(22990, 'MANUK', 1750, 91, 6),
		(22991, 'NGABAR', 1750, 91, 6),
		(22992, 'PATIHAN KIDUL', 1750, 91, 6),
		(22993, 'PIJERAN', 1750, 91, 6),
		(22994, 'RONOSENTANAN', 1750, 91, 6),
		(22995, 'RONOWIJAYAN', 1750, 91, 6),
		(22996, 'SAWUH', 1750, 91, 6),
		(22997, 'SEKARAN (SAKARAN)', 1750, 91, 6),
		(22998, 'SIMAN', 1750, 91, 6),
		(22999, 'TAJUG', 1750, 91, 6),
		(23000, 'TRANJANG', 1750, 91, 6)
	`)
	if err != nil {
		panic(err)
	}
	fmt.Println("insert villige 22")
}
