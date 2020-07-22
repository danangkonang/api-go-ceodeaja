package kel50

import (
	"fmt"

	"github.com/danangkonang/rest-api/config"
)

func Kel57() {
	db := config.Connect()
	_, err := db.Exec(`INSERT INTO villages (id,village,sub_district_id,regency_id,province_id) VALUES
	(57001, 'BUKIT RAYA (SEPAKU I)', 4320, 308, 21),
	(57002, 'BUMI HARAPAN', 4320, 308, 21),
	(57003, 'KARANG JINAWI', 4320, 308, 21),
	(57004, 'MARIDAN', 4320, 308, 21),
	(57005, 'MENTAWIR', 4320, 308, 21),
	(57006, 'PEMALUAN', 4320, 308, 21),
	(57007, 'SEMOI II', 4320, 308, 21),
	(57008, 'SEPAKU', 4320, 308, 21),
	(57009, 'SUKARAJA (SEPAKU II)', 4320, 308, 21),
	(57010, 'SUKO MULYO', 4320, 308, 21),
	(57011, 'TELEMOW', 4320, 308, 21),
	(57012, 'TENGIN BARU (SEPAKU III)', 4320, 308, 21),
	(57013, 'WONOSARI', 4320, 308, 21),
	(57014, 'API API', 4321, 308, 21),
	(57015, 'BAJUR', 4321, 308, 21),
	(57016, 'BANGUN MULYA', 4321, 308, 21),
	(57017, 'BERBEK', 4321, 308, 21),
	(57018, 'BUNGURASIH', 4321, 308, 21),
	(57019, 'JANTI', 4321, 308, 21),
	(57020, 'KEDUNGREJO', 4321, 308, 21),
	(57021, 'KEPUH KIRIMAN', 4321, 308, 21),
	(57022, 'KUREKSARI', 4321, 308, 21),
	(57023, 'MEDAENG', 4321, 308, 21),
	(57024, 'NGINGAS', 4321, 308, 21),
	(57025, 'PEPELEGI', 4321, 308, 21),
	(57026, 'RAGANG', 4321, 308, 21),
	(57027, 'SANA LAOK', 4321, 308, 21),
	(57028, 'SESULU', 4321, 308, 21),
	(57029, 'SUMBER WARU', 4321, 308, 21),
	(57030, 'TAMBAK OSO', 4321, 308, 21),
	(57031, 'TAMBAK REJO', 4321, 308, 21),
	(57032, 'TAMBAK SAWAH', 4321, 308, 21),
	(57033, 'TAMBAK SUMUR', 4321, 308, 21),
	(57034, 'TAMPOJUNG GUWA', 4321, 308, 21),
	(57035, 'TAMPOJUNG PREGI', 4321, 308, 21),
	(57036, 'TAMPOJUNG TENGAH', 4321, 308, 21),
	(57037, 'TAMPOJUNG TENGGINA', 4321, 308, 21),
	(57038, 'TEGANGSER LAOK (TAGANGSER LAOK)', 4321, 308, 21),
	(57039, 'TLONTO ARES', 4321, 308, 21),
	(57040, 'TROPODO', 4321, 308, 21),
	(57041, 'WADUNGASRI', 4321, 308, 21),
	(57042, 'WARU', 4321, 308, 21),
	(57043, 'WARU', 4321, 308, 21),
	(57044, 'WARU BARAT', 4321, 308, 21),
	(57045, 'WARU TIMUR', 4321, 308, 21),
	(57046, 'WEDORO', 4321, 308, 21),
	(57047, 'HARAPAN BARU', 4322, 309, 21),
	(57048, 'RAPAK DALAM', 4322, 309, 21),
	(57049, 'SENGKOTEK', 4322, 309, 21),
	(57050, 'SIMPANG TIGA (LOA JANAN ILIR)', 4322, 309, 21),
	(57051, 'TANI AMAN', 4322, 309, 21),
	(57052, 'BANTUAS', 4323, 309, 21),
	(57053, 'BUKUAN', 4323, 309, 21),
	(57054, 'HANDIL BAKTI', 4323, 309, 21),
	(57055, 'RAWA MAKMUR', 4323, 309, 21),
	(57056, 'SIMPANG PASIR', 4323, 309, 21),
	(57057, 'PELITA', 4324, 309, 21),
	(57058, 'SEI/SUNGAI DAMA', 4324, 309, 21),
	(57059, 'SELILI', 4324, 309, 21),
	(57060, 'SIDODAMAI', 4324, 309, 21),
	(57061, 'SIDOMULYO', 4324, 309, 21),
	(57062, 'BUGIS', 4325, 309, 21),
	(57063, 'KARANG MUMUS', 4325, 309, 21),
	(57064, 'PASAR PAGI', 4325, 309, 21),
	(57065, 'PELABUHAN', 4325, 309, 21),
	(57066, 'SEI/SUNGAI PINANG LUAR', 4325, 309, 21),
	(57067, 'BAQA/BAKA/RAPAK DALAM', 4326, 309, 21),
	(57068, 'MESJID', 4326, 309, 21),
	(57069, 'SEI/SUNGAI KELEDANG', 4326, 309, 21),
	(57070, 'AIR HITAM', 4327, 309, 21),
	(57071, 'AIR PUTIH', 4327, 309, 21),
	(57072, 'BUKIT PINANG', 4327, 309, 21),
	(57073, 'DADI MULYA', 4327, 309, 21),
	(57074, 'GUNUNG KELUA', 4327, 309, 21),
	(57075, 'JAWA', 4327, 309, 21),
	(57076, 'SIDODADI', 4327, 309, 21),
	(57077, 'TELUK LERONG ILIR', 4327, 309, 21),
	(57078, 'LEMPAKE', 4328, 309, 21),
	(57079, 'SEI/SUNGAI SIRING', 4328, 309, 21),
	(57080, 'SEMPAJA SELATAN', 4328, 309, 21),
	(57081, 'SEMPAJA UTARA', 4328, 309, 21),
	(57082, 'TANAH MERAH', 4328, 309, 21),
	(57083, 'MAKROMAN', 4329, 309, 21),
	(57084, 'PULAU ATAS', 4329, 309, 21),
	(57085, 'SAMBUTAN', 4329, 309, 21),
	(57086, 'SEI/SUNGAI KAPIH', 4329, 309, 21),
	(57087, 'SINDANG SARI', 4329, 309, 21),
	(57088, 'KARANG ANYAR', 4330, 309, 21),
	(57089, 'KARANG ASAM ILIR', 4330, 309, 21),
	(57090, 'KARANG ASAM ULU', 4330, 309, 21),
	(57091, 'LOA BAKUNG', 4330, 309, 21),
	(57092, 'LOA BUAH', 4330, 309, 21),
	(57093, 'LOK BAHU', 4330, 309, 21),
	(57094, 'TELUK LERONG ULU', 4330, 309, 21),
	(57095, 'BANDARA', 4331, 309, 21),
	(57096, 'GUNUNG LINGAI', 4331, 309, 21),
	(57097, 'MUGIREJO', 4331, 309, 21),
	(57098, 'PENYANDINGAN', 4331, 309, 21),
	(57099, 'PINANG JAYA', 4331, 309, 21),
	(57100, 'PINANG MAS', 4331, 309, 21),
	(57101, 'SEI/SUNGAI PINANG DALAM', 4331, 309, 21),
	(57102, 'SERIJABO', 4331, 309, 21),
	(57103, 'SERIJABO BARU', 4331, 309, 21),
	(57104, 'SUNGAI PINANG', 4331, 309, 21),
	(57105, 'SUNGAI PINANG I', 4331, 309, 21),
	(57106, 'SUNGAI PINANG II', 4331, 309, 21),
	(57107, 'SUNGAI PINANG III', 4331, 309, 21),
	(57108, 'SUNGAI PINANG LAGATI', 4331, 309, 21),
	(57109, 'SUNGAI PINANG NIBUNG', 4331, 309, 21),
	(57110, 'TAJUNG SERIAN', 4331, 309, 21),
	(57111, 'TALANG DUKUN', 4331, 309, 21),
	(57112, 'TEMINDUNG PERMAI', 4331, 309, 21),
	(57113, 'LEPAK ARU', 4332, 310, 22),
	(57114, 'LONG BIA', 4332, 310, 22),
	(57115, 'LONG BUANG', 4332, 310, 22),
	(57116, 'LONG LASAN', 4332, 310, 22),
	(57117, 'LONG LEJU', 4332, 310, 22),
	(57118, 'LONG LIAN', 4332, 310, 22),
	(57119, 'LONG PELEBAN', 4332, 310, 22),
	(57120, 'LONG PESO', 4332, 310, 22),
	(57121, 'LONG YIIN/PELAAH', 4332, 310, 22),
	(57122, 'MUARA PENGEAN', 4332, 310, 22),
	(57123, 'LONG BANG', 4333, 310, 22),
	(57124, 'LONG BANG HULU', 4333, 310, 22),
	(57125, 'LONG LEMBU', 4333, 310, 22),
	(57126, 'LONG TELENJAU', 4333, 310, 22),
	(57127, 'LONG TUNGU', 4333, 310, 22),
	(57128, 'NAHA AYA', 4333, 310, 22),
	(57129, 'PULAU BUNYU BARAT', 4334, 310, 22),
	(57130, 'PULAU BUNYU SELATAN', 4334, 310, 22),
	(57131, 'PULAU BUNYU TIMUR', 4334, 310, 22),
	(57132, 'AMBALAT', 4335, 310, 22),
	(57133, 'ANJAR ARIP (ANJARAREF)', 4335, 310, 22),
	(57134, 'BAMBANG', 4335, 310, 22),
	(57135, 'BEKILIU', 4335, 310, 22),
	(57136, 'BUNAU', 4335, 310, 22),
	(57137, 'KELINCAWAN', 4335, 310, 22),
	(57138, 'KELISING (KELIISING)', 4335, 310, 22),
	(57139, 'KELUMBUNAN (KELEMBUNAN)', 4335, 310, 22),
	(57140, 'KENDARI', 4335, 310, 22),
	(57141, 'KERITING', 4335, 310, 22),
	(57142, 'LIU AGU (LIAGU)', 4335, 310, 22),
	(57143, 'MARITAM', 4335, 310, 22),
	(57144, 'PARU ABANG', 4335, 310, 22),
	(57145, 'PENTIAN', 4335, 310, 22),
	(57146, 'PUNAN DULAU', 4335, 310, 22),
	(57147, 'PUNGIT', 4335, 310, 22),
	(57148, 'SEKATAK BENGARA', 4335, 310, 22),
	(57149, 'SEKATAK BUJI', 4335, 310, 22),
	(57150, 'TENGGILING', 4335, 310, 22),
	(57151, 'TERINDAK', 4335, 310, 22),
	(57152, 'TURUNG', 4335, 310, 22),
	(57153, 'UJANG', 4335, 310, 22),
	(57154, 'ANTUTAN', 4336, 310, 22),
	(57155, 'GUNUNG PUTIH', 4336, 310, 22),
	(57156, 'KARANG ANYAR', 4336, 310, 22),
	(57157, 'PEJALIN', 4336, 310, 22),
	(57158, 'TANJUNG PALAS ILIR/HILIR', 4336, 310, 22),
	(57159, 'TANJUNG PALAS TENGAH', 4336, 310, 22),
	(57160, 'TANJUNG PALAS ULU/HULU', 4336, 310, 22),
	(57161, 'TERAS BARU', 4336, 310, 22),
	(57162, 'TERAS NAWANG', 4336, 310, 22),
	(57163, 'LONG BELUAH', 4337, 310, 22),
	(57164, 'LONG PARI', 4337, 310, 22),
	(57165, 'LONG SAM', 4337, 310, 22),
	(57166, 'MARA HILIR', 4337, 310, 22),
	(57167, 'MARA SATU', 4337, 310, 22),
	(57168, 'SALIM BATU', 4338, 310, 22),
	(57169, 'SILVA RAHAYU', 4338, 310, 22),
	(57170, 'TANJUNG BUKA', 4338, 310, 22),
	(57171, 'BINAI', 4339, 310, 22),
	(57172, 'MANGKUPADI', 4339, 310, 22),
	(57173, 'PURA SAJAU', 4339, 310, 22),
	(57174, 'SAJAU', 4339, 310, 22),
	(57175, 'SAJAU TIMUR', 4339, 310, 22),
	(57176, 'TANAH KUNING', 4339, 310, 22),
	(57177, 'TANJUNG AGUNG', 4339, 310, 22),
	(57178, 'WONO MULYO', 4339, 310, 22),
	(57179, 'ARDI MULYA', 4340, 310, 22),
	(57180, 'KARANG AGUNG', 4340, 310, 22),
	(57181, 'KELUBIR', 4340, 310, 22),
	(57182, 'PANCA AGUNG', 4340, 310, 22),
	(57183, 'PIMPING', 4340, 310, 22),
	(57184, 'RUHUI RAHAYU', 4340, 310, 22),
	(57185, 'APUNG', 4341, 310, 22),
	(57186, 'BUMI RAHAYU', 4341, 310, 22),
	(57187, 'GUNUNG SARI', 4341, 310, 22),
	(57188, 'GUNUNG SERIANG', 4341, 310, 22),
	(57189, 'JELARAI SELOR', 4341, 310, 22),
	(57190, 'TANJUNG SELOR ILIR/HILIR', 4341, 310, 22),

	(57191, 'TANJUNG SELOR TIMUR', 4341, 310, 22),
	(57192, 'TANJUNG SELOR ULU/HULU', 4341, 310, 22),
	(57193, 'TENGKAPAK', 4341, 310, 22),
	(57194, 'APAU PING MADING', 4342, 311, 22),
	(57195, 'LONG ALANGO', 4342, 311, 22),
	(57196, 'LONG ATUA/LONG KEMUAT', 4342, 311, 22),
	(57197, 'LONG BERINI MARARIAN', 4342, 311, 22),
	(57198, 'LONG TABULO/TEBULO', 4342, 311, 22),
	(57199, 'LONG ULI', 4342, 311, 22),
	(57200, 'LIDUNG PAYAU', 4343, 311, 22),
	(57201, 'LONG AMPUNG', 4343, 311, 22),
	(57202, 'LONG URO', 4343, 311, 22),
	(57203, 'METULANG', 4343, 311, 22),
	(57204, 'SUNGAI BARANG', 4343, 311, 22),
	(57205, 'KUALA LAPANG', 4344, 311, 22),
	(57206, 'LONG BILA', 4344, 311, 22),
	(57207, 'LONG KENIPE', 4344, 311, 22),
	(57208, 'PUNAN BENGALUN', 4344, 311, 22),
	(57209, 'SEMPAYANG', 4344, 311, 22),
	(57210, 'SENTABAN', 4344, 311, 22),
	(57211, 'SESUA', 4344, 311, 22),
	(57212, 'TANJUNG LAPANG', 4344, 311, 22),
	(57213, 'TARAS', 4344, 311, 22),
	(57214, 'BATU LIDUNG', 4345, 311, 22),
	(57215, 'MALINAU HILIR', 4345, 311, 22),
	(57216, 'MALINAU HULU', 4345, 311, 22),
	(57217, 'MALINAU KOTA', 4345, 311, 22),
	(57218, 'PELITA KANAAN', 4345, 311, 22),
	(57219, 'TANJUNG KERANJANG', 4345, 311, 22),
	(57220, 'BILA BEKAYUK', 4346, 311, 22),
	(57221, 'LABAN NYARIT', 4346, 311, 22),
	(57222, 'LANGAP', 4346, 311, 22),
	(57223, 'LONG LOREH', 4346, 311, 22),
	(57224, 'NUNUK TANAH KIBANG', 4346, 311, 22),
	(57225, 'PAYA SETURAN', 4346, 311, 22),
	(57226, 'PELENCAU', 4346, 311, 22),
	(57227, 'PUNAN RIAN', 4346, 311, 22),
	(57228, 'SENGAYAN', 4346, 311, 22),
	(57229, 'BATU KAJANG', 4347, 311, 22),
	(57230, 'GONG SOLOK', 4347, 311, 22),
	(57231, 'LONG ADIU', 4347, 311, 22),
	(57232, 'PUNAN GONG SOLOK', 4347, 311, 22),
	(57233, 'PUNAN LONG ADIU', 4347, 311, 22),
	(57234, 'PUNAN SETARAP', 4347, 311, 22),
	(57235, 'SETARAP', 4347, 311, 22),
	(57236, 'SETULANG', 4347, 311, 22),
	(57237, 'HALANGA', 4348, 311, 22),
	(57238, 'LONG JALAN', 4348, 311, 22),
	(57239, 'LONG LAKE', 4348, 311, 22),
	(57240, 'LONG RAT', 4348, 311, 22),
	(57241, 'METUT', 4348, 311, 22),
	(57242, 'NAHAKRAMO BARU', 4348, 311, 22),
	(57243, 'PUNAN MIRAU', 4348, 311, 22),
	(57244, 'TANJUNG NANGA', 4348, 311, 22),
	(57245, 'BELAYAN', 4349, 311, 22),
	(57246, 'KALIAMOK', 4349, 311, 22),
	(57247, 'KELAPIS', 4349, 311, 22),
	(57248, 'LUBOK MANIS', 4349, 311, 22),
	(57249, 'LUSO', 4349, 311, 22),
	(57250, 'MALINAU SEBERANG', 4349, 311, 22),
	(57251, 'PUTAT', 4349, 311, 22),
	(57252, 'RESPEN TUBU', 4349, 311, 22),
	(57253, 'SALAP', 4349, 311, 22),
	(57254, 'SEMANGGARIS', 4349, 311, 22),
	(57255, 'SEMBUAK WAROT/WAROD', 4349, 311, 22),
	(57256, 'SERUYUNG', 4349, 311, 22),
	(57257, 'HARAPAN MAJU', 4350, 311, 22),
	(57258, 'LIDUNG KEMENCI', 4350, 311, 22),
	(57259, 'LONG BISAI', 4350, 311, 22),
	(57260, 'LONG GAFID', 4350, 311, 22),
	(57261, 'LONG LIKU', 4350, 311, 22),
	(57262, 'MENTARANG BARU', 4350, 311, 22),
	(57263, 'PAKING', 4350, 311, 22),
	(57264, 'PULAU SAPI', 4350, 311, 22),
	(57265, 'TEMALANG', 4350, 311, 22),
	(57266, 'LONG BERANG', 4351, 311, 22),
	(57267, 'LONG KEBINU', 4351, 311, 22),
	(57268, 'LONG MEKATIF', 4351, 311, 22),
	(57269, 'LONG PALA', 4351, 311, 22),
	(57270, 'LONG SEMAMU', 4351, 311, 22),
	(57271, 'LONG SIMAU', 4351, 311, 22),
	(57272, 'LONG SULIT', 4351, 311, 22),
	(57273, 'LAME', 4352, 311, 22),
	(57274, 'LONG ARAN', 4352, 311, 22),
	(57275, 'LONG BELAKA PITAU', 4352, 311, 22),
	(57276, 'LONG BENA', 4352, 311, 22),
	(57277, 'LONG JELET', 4352, 311, 22),
	(57278, 'LONG KETAMAN', 4352, 311, 22),
	(57279, 'LONG PALIRAN', 4352, 311, 22),
	(57280, 'LONG PUA', 4352, 311, 22),
	(57281, 'LONG PUJUNGAN', 4352, 311, 22),
	(57282, 'AGUNG BARU', 4353, 311, 22),
	(57283, 'DATA BARU', 4353, 311, 22),
	(57284, 'DUMU MAHAK', 4353, 311, 22),
	(57285, 'LONG LEBUSAN', 4353, 311, 22),
	(57286, 'LONG TOP', 4353, 311, 22),
	(57287, 'MAHAK BARU', 4353, 311, 22),
	(57288, 'LONG NYAU', 4354, 311, 22),
	(57289, 'LONG PADA', 4354, 311, 22),
	(57290, 'LONG RANAU', 4354, 311, 22),
	(57291, 'LONG TITI', 4354, 311, 22),
	(57292, 'RIAN TUBU', 4354, 311, 22),
	(57293, 'BA SIKOR', 4355, 312, 22),
	(57294, 'BUDUK KINANGAN', 4355, 312, 22),
	(57295, 'BUDUK KUBUL', 4355, 312, 22),
	(57296, 'BUDUK TUMU', 4355, 312, 22),
	(57297, 'BUNGAYAN', 4355, 312, 22),
	(57298, 'CINGLAT', 4355, 312, 22),
	(57299, 'KAMPUNG BARU', 4355, 312, 22),
	(57300, 'LEMBADA', 4355, 312, 22),
	(57301, 'LEMBUDUD', 4355, 312, 22),
	(57302, 'LEPATAR', 4355, 312, 22),
	(57303, 'LIANG ALIG', 4355, 312, 22),
	(57304, 'LIANG BUA', 4355, 312, 22),
	(57305, 'LIANG BUTAN', 4355, 312, 22),
	(57306, 'LIANG TUER', 4355, 312, 22),
	(57307, 'LIANG TURAN', 4355, 312, 22),
	(57308, 'LONG BAWAN', 4355, 312, 22),
	(57309, 'LONG BERAYANG', 4355, 312, 22),
	(57310, 'LONG BIA DUNG', 4355, 312, 22),
	(57311, 'LONG KABID', 4355, 312, 22),
	(57312, 'LONG KATUNG', 4355, 312, 22),
	(57313, 'LONG MANGAN', 4355, 312, 22),
	(57314, 'LONG MATUNG', 4355, 312, 22),
	(57315, 'LONG NAWANG', 4355, 312, 22),
	(57316, 'LONG NUAT', 4355, 312, 22),
	(57317, 'LONG PUAK', 4355, 312, 22),
	(57318, 'LONG RUPAN', 4355, 312, 22),
	(57319, 'LONG SEPAYANG', 4355, 312, 22),
	(57320, 'LONG TENEM', 4355, 312, 22),
	(57321, 'LONG TUGUL', 4355, 312, 22),
	(57322, 'LONG UMUNG', 4355, 312, 22),
	(57323, 'MA LIBU', 4355, 312, 22),
	(57324, 'PA API', 4355, 312, 22),
	(57325, 'PA BETUNG', 4355, 312, 22),
	(57326, 'PA BUTAL', 4355, 312, 22),
	(57327, 'PA DELUNG', 4355, 312, 22),
	(57328, 'PA INAN', 4355, 312, 22),
	(57329, 'PA KEBUAN', 4355, 312, 22),
	(57330, 'PA KEMUT', 4355, 312, 22),
	(57331, 'PA KIDANG', 4355, 312, 22),
	(57332, 'PA LIDUNG', 4355, 312, 22),
	(57333, 'PA LUTUT', 4355, 312, 22),
	(57334, 'PA MALADE (PA MELADE)', 4355, 312, 22),
	(57335, 'PA MATUNG', 4355, 312, 22),
	(57336, 'PA MERING', 4355, 312, 22),
	(57337, 'PA MULAK', 4355, 312, 22),
	(57338, 'PA NADO', 4355, 312, 22),
	(57339, 'PA PADI', 4355, 312, 22),
	(57340, 'PA PALA', 4355, 312, 22),
	(57341, 'PA PANI', 4355, 312, 22),
	(57342, 'PA PAWAN', 4355, 312, 22),
	(57343, 'PA PAYAK', 4355, 312, 22),
	(57344, 'PA PIRIT', 4355, 312, 22),
	(57345, 'PA PUTUK', 4355, 312, 22),
	(57346, 'PA RANGEB (PARANGEB)', 4355, 312, 22),
	(57347, 'PA RAYE', 4355, 312, 22),
	(57348, 'PA RUPAI', 4355, 312, 22),
	(57349, 'PA SIRE (PASIRE)', 4355, 312, 22),
	(57350, 'PA TERUTUN', 4355, 312, 22),
	(57351, 'PA UMUNG', 4355, 312, 22),
	(57352, 'PA URUD', 4355, 312, 22),
	(57353, 'SEMBUDUD', 4355, 312, 22),
	(57354, 'SINAR BARU', 4355, 312, 22),
	(57355, 'WA LAYA', 4355, 312, 22),
	(57356, 'WA YAGUNG', 4355, 312, 22),
	(57357, 'WA YANUD', 4355, 312, 22),
	(57358, 'BA LIKU', 4356, 312, 22),
	(57359, 'BINUANG (BA BINUANG)', 4356, 312, 22),
	(57360, 'LIANG LUNUK', 4356, 312, 22),
	(57361, 'LONG BIRAR', 4356, 312, 22),
	(57362, 'LONG BUDUNG', 4356, 312, 22),
	(57363, 'LONG KELUPAN / KLUPAN', 4356, 312, 22),
	(57364, 'LONG MUTAN', 4356, 312, 22),
	(57365, 'LONG PADI', 4356, 312, 22),
	(57366, 'LONG PASIA', 4356, 312, 22),
	(57367, 'LONG PUPUNG', 4356, 312, 22),
	(57368, 'LONG RIAN', 4356, 312, 22),
	(57369, 'LONG RUNGAN', 4356, 312, 22),
	(57370, 'PA AMAI', 4356, 312, 22),
	(57371, 'PA DALAN', 4356, 312, 22),
	(57372, 'PA IBANG', 4356, 312, 22),
	(57373, 'PA KABER', 4356, 312, 22),
	(57374, 'PA MILAU', 4356, 312, 22),
	(57375, 'PA SING', 4356, 312, 22),
	(57376, 'PA TERA', 4356, 312, 22),
	(57377, 'PA UPAN', 4356, 312, 22),
	(57378, 'PA URANG', 4356, 312, 22),
	(57379, 'PA YALAU', 4356, 312, 22),
	(57380, 'TANG BADUI', 4356, 312, 22),
	(57381, 'TANG PAYE', 4356, 312, 22),
	(57382, 'BULAN BULAN', 4357, 312, 22),
	(57383, 'DABULON (DUBULON)', 4357, 312, 22),
	(57384, 'DERALON', 4357, 312, 22),
	(57385, 'KALAMPISING', 4357, 312, 22),
	(57386, 'LIANG', 4357, 312, 22),
	(57387, 'LIBANG', 4357, 312, 22),
	(57388, 'LIKOS', 4357, 312, 22),
	(57389, 'LINTONG', 4357, 312, 22),
	(57390, 'MANSALONG', 4357, 312, 22),
	(57391, 'NAINSID', 4357, 312, 22),
	(57392, 'PA LEMUMUT', 4357, 312, 22),
	(57393, 'PA LOO', 4357, 312, 22),
	(57394, 'PATAL I (SATU)', 4357, 312, 22),
	(57395, 'PATAL II (DUA)', 4357, 312, 22),
	(57396, 'PODONG', 4357, 312, 22),
	(57397, 'PULU BULAWAN', 4357, 312, 22),
	(57398, 'SALUDAN', 4357, 312, 22),
	(57399, 'SANGKUP', 4357, 312, 22),
	(57400, 'SAPUYAN', 4357, 312, 22),
	(57401, 'SASIBU', 4357, 312, 22),
	(57402, 'SEDONGON', 4357, 312, 22),
	(57403, 'SEMALAT', 4357, 312, 22),
	(57404, 'SIAWANG', 4357, 312, 22),
	(57405, 'SUMALUMUNG', 4357, 312, 22),
	(57406, 'TALUAN', 4357, 312, 22),
	(57407, 'TANJUNG HILIR', 4357, 312, 22),
	(57408, 'TANJUNG HULU', 4357, 312, 22),
	(57409, 'TUBUS', 4357, 312, 22),
	(57410, 'BATUNG', 4358, 312, 22),
	(57411, 'BOKOK', 4358, 312, 22),
	(57412, 'BULU LAUN HILIR', 4358, 312, 22),
	(57413, 'BULU LAUN HULU', 4358, 312, 22),
	(57414, 'BULU LUAN HULU', 4358, 312, 22),
	(57415, 'BULU MENGELOM', 4358, 312, 22),
	(57416, 'DUYAN', 4358, 312, 22),
	(57417, 'JUKUP', 4358, 312, 22),
	(57418, 'KABUNGALOR', 4358, 312, 22),
	(57419, 'KALAM BUKU', 4358, 312, 22),
	(57420, 'KALISUN', 4358, 312, 22),
	(57421, 'KUYO', 4358, 312, 22),
	(57422, 'LABANG', 4358, 312, 22),
	(57423, 'LABUK', 4358, 312, 22),
	(57424, 'LAGAS', 4358, 312, 22),
	(57425, 'LANGGASON', 4358, 312, 22),
	(57426, 'LEPAGA', 4358, 312, 22),
	(57427, 'LIMPAKON', 4358, 312, 22),
	(57428, 'LINSAYUNG', 4358, 312, 22),
	(57429, 'LONG BULU', 4358, 312, 22),
	(57430, 'MAMANIS', 4358, 312, 22),
	(57431, 'NAN SAPAN', 4358, 312, 22),
	(57432, 'NANTUKIDAN', 4358, 312, 22),
	(57433, 'NGAWOL', 4358, 312, 22),
	(57434, 'PALUAN', 4358, 312, 22),
	(57435, 'PANAS', 4358, 312, 22),
	(57436, 'PAYANG', 4358, 312, 22),
	(57437, 'SALAN', 4358, 312, 22),
	(57438, 'SAMUNTI', 4358, 312, 22),
	(57439, 'SANAL', 4358, 312, 22),
	(57440, 'SEDALIT', 4358, 312, 22),
	(57441, 'SEMATA', 4358, 312, 22),
	(57442, 'SIBALU', 4358, 312, 22),
	(57443, 'SINAMPALA DUA', 4358, 312, 22),
	(57444, 'SINAMPALA SATU', 4358, 312, 22),
	(57445, 'SUMANTIPAL', 4358, 312, 22),
	(57446, 'SUMENTOBOL', 4358, 312, 22),
	(57447, 'SUNGOI', 4358, 312, 22),
	(57448, 'SUYADON', 4358, 312, 22),
	(57449, 'TAMBALANG HILIR', 4358, 312, 22),
	(57450, 'TAMBALANG ULU', 4358, 312, 22),
	(57451, 'TANTALUJUK', 4358, 312, 22),
	(57452, 'TANTU LIBING', 4358, 312, 22),
	(57453, 'TAU LUMBIS', 4358, 312, 22),
	(57454, 'TETAGAS', 4358, 312, 22),
	(57455, 'TUBEL ALUNG (UBEL ALUNG)', 4358, 312, 22),
	(57456, 'TUKULEN (TUKULON)', 4358, 312, 22),
	(57457, 'TUMATALAS', 4358, 312, 22),
	(57458, 'UBEL SULOK', 4358, 312, 22),
	(57459, 'BINUSAN', 4359, 312, 22),
	(57460, 'NUNUKAN BARAT', 4359, 312, 22),
	(57461, 'NUNUKAN TENGAH', 4359, 312, 22),
	(57462, 'NUNUKAN TIMUR', 4359, 312, 22),
	(57463, 'NUNUKAN UTARA', 4359, 312, 22),
	(57464, 'MANSAPA', 4360, 312, 22),
	(57465, 'NUNUKAN SELATAN', 4360, 312, 22),
	(57466, 'SELISUN', 4360, 312, 22),
	(57467, 'TANJUNG HARAPAN', 4360, 312, 22),
	(57468, 'BALANSIKU', 4361, 312, 22),
	(57469, 'PADAIDI', 4361, 312, 22),
	(57470, 'SUNGAI MANURUNG', 4361, 312, 22),
	(57471, 'TANJUNG KARANG', 4361, 312, 22),
	(57472, 'BAMBANGAN', 4362, 312, 22),
	(57473, 'BINALAWAN', 4362, 312, 22),
	(57474, 'LIANG BUNYU', 4362, 312, 22),
	(57475, 'SETABU', 4362, 312, 22),
	(57476, 'AJI KUNING', 4363, 312, 22),
	(57477, 'BUKIT HARAPAN', 4363, 312, 22),
	(57478, 'MASPUL', 4363, 312, 22),
	(57479, 'SEI LIMAU', 4363, 312, 22),
	(57480, 'BUKIT ARU INDAH', 4364, 312, 22),
	(57481, 'SEI/SUNGAI NYAMUK', 4364, 312, 22),
	(57482, 'TANJUNG ARU', 4364, 312, 22),
	(57483, 'TANJUNG HARAPAN', 4364, 312, 22),
	(57484, 'LAPRI', 4365, 312, 22),
	(57485, 'SEBERANG', 4365, 312, 22),
	(57486, 'SEI/SUNGAI PANCANG', 4365, 312, 22),
	(57487, 'APAS', 4366, 312, 22),
	(57488, 'BEBANAS', 4366, 312, 22),
	(57489, 'HARAPAN', 4366, 312, 22),
	(57490, 'KEKAYAP', 4366, 312, 22),
	(57491, 'KUNYIT', 4366, 312, 22),
	(57492, 'LULU', 4366, 312, 22),
	(57493, 'MELASU BARU', 4366, 312, 22),
	(57494, 'PEMBELIANGAN', 4366, 312, 22),
	(57495, 'SUJAU', 4366, 312, 22),
	(57496, 'TETABAN', 4366, 312, 22),
	(57497, 'SAMAENRE SEMAJA', 4367, 312, 22),
	(57498, 'SEKADUYAN TAKA', 4367, 312, 22),
	(57499, 'SRINANTI', 4367, 312, 22),
	(57500, 'TABUR LESTARI', 4367, 312, 22),
	(57501, 'ATAP', 4368, 312, 22),
	(57502, 'BINANUN', 4368, 312, 22),
	(57503, 'BUTAS BAGU', 4368, 312, 22),
	(57504, 'KATUL', 4368, 312, 22),
	(57505, 'LABUK', 4368, 312, 22),
	(57506, 'LIUK BULU', 4368, 312, 22),
	(57507, 'LUBAKAN', 4368, 312, 22),
	(57508, 'LUBOK BUAT', 4368, 312, 22),
	(57509, 'MAMBULU', 4368, 312, 22),
	(57510, 'MANUK BUNGKUL', 4368, 312, 22),
	(57511, 'PAGALUYON', 4368, 312, 22),
	(57512, 'PAGAR', 4368, 312, 22),
	(57513, 'PELAJU (PLAJU)', 4368, 312, 22),
	(57514, 'PULAU KERAS', 4368, 312, 22),
	(57515, 'SABULUAN', 4368, 312, 22),
	(57516, 'SADUMAN', 4368, 312, 22),
	(57517, 'TAGUL', 4368, 312, 22),
	(57518, 'TEPIAN', 4368, 312, 22),
	(57519, 'TUJUNG', 4368, 312, 22),
	(57520, 'TULANG', 4368, 312, 22),
	(57521, 'BLATIKON (BELATIKAN)', 4369, 312, 22),
	(57522, 'KULUNSAYAN (KALUN SAYAN)', 4369, 312, 22),
	(57523, 'NAPUTI', 4369, 312, 22),
	(57524, 'SALANG', 4369, 312, 22),
	(57525, 'SEBUKU I (SANUR)', 4369, 312, 22),
	(57526, 'SEBUKU II (MAKMUR)', 4369, 312, 22),
	(57527, 'SEKIKILAN', 4369, 312, 22),
	(57528, 'SEMUNAD', 4369, 312, 22),
	(57529, 'TAU BARU', 4369, 312, 22),
	(57530, 'TEMBALANG', 4369, 312, 22),
	(57531, 'TINAMPAK I', 4369, 312, 22),
	(57532, 'TINAMPAK II', 4369, 312, 22),
	(57533, 'BEBAKUNG', 4370, 313, 22),
	(57534, 'BUANG BARU (BUONG BARU)', 4370, 313, 22),
	(57535, 'KUJAU', 4370, 313, 22),
	(57536, 'MANING', 4370, 313, 22),
	(57537, 'MENDUPE/MENDOPU', 4370, 313, 22),
	(57538, 'PERIUK', 4370, 313, 22),
	(57539, 'GUNAWAN', 4371, 313, 22),
	(57540, 'LIMBU SEDULUN/SIDULUN', 4371, 313, 22),
	(57541, 'RIAN', 4371, 313, 22),
	(57542, 'SEBAWANG', 4371, 313, 22),
	(57543, 'SEBIDAI/SEBIDAY', 4371, 313, 22),
	(57544, 'SEDULUN', 4371, 313, 22),
	(57545, 'SEPUTUK', 4371, 313, 22),
	(57546, 'TIDENG PALE', 4371, 313, 22),
	(57547, 'TIDENG PALE TIMUR', 4371, 313, 22),
	(57548, 'BADAN BIKIS', 4372, 313, 22),
	(57549, 'BEBATU', 4372, 313, 22),
	(57550, 'MENDELUTUNG (MENJELUTUNG)', 4372, 313, 22),
	(57551, 'SELUDAU', 4372, 313, 22),
	(57552, 'SENGKONG', 4372, 313, 22),
	(57553, 'SEPALA DALUNG (SESAYAP DALUNG)', 4372, 313, 22),
	(57554, 'SESAYAP/SEPALA', 4372, 313, 22),
	(57555, 'SAMBUNGAN', 4373, 313, 22),
	(57556, 'TANAH MERAH', 4373, 313, 22),
	(57557, 'TENGKU DACING', 4373, 313, 22),
	(57558, 'KARANG ANYAR', 4374, 314, 22),
	(57559, 'KARANG ANYAR PANTAI', 4374, 314, 22),
	(57560, 'KARANG BALIK', 4374, 314, 22),
	(57561, 'KARANG HARAPAN', 4374, 314, 22),
	(57562, 'KARANG REJO', 4374, 314, 22),
	(57563, 'KAMPUNG SATU (SKIP)', 4375, 314, 22),
	(57564, 'PAMUSIAN', 4375, 314, 22),
	(57565, 'SEBENGKOK', 4375, 314, 22),
	(57566, 'SELUMIT', 4375, 314, 22),
	(57567, 'SELUMIT PANTAI', 4375, 314, 22),
	(57568, 'GUNUNG LINGKAS', 4376, 314, 22),
	(57569, 'KAMPUNG EMPAT', 4376, 314, 22),
	(57570, 'KAMPUNG ENAM', 4376, 314, 22),
	(57571, 'LINGKAS UJUNG', 4376, 314, 22),
	(57572, 'MAMBURUNGAN', 4376, 314, 22),
	(57573, 'MAMBURUNGAN TIMUR', 4376, 314, 22),
	(57574, 'PANTAI AMAL', 4376, 314, 22),
	(57575, 'JUATA KRIKIL/KERIKIL', 4377, 314, 22),
	(57576, 'JUATA LAUT', 4377, 314, 22),
	(57577, 'JUATA PERMAI', 4377, 314, 22),
	(57578, 'BANGGAE', 4378, 315, 23),
	(57579, 'BARU', 4378, 315, 23),
	(57580, 'GALUNG', 4378, 315, 23),
	(57581, 'PALIPI SOREANG', 4378, 315, 23),
	(57582, 'PAMBOBORANG', 4378, 315, 23),
	(57583, 'PANGALI-ALI', 4378, 315, 23),
	(57584, 'RANGAS', 4378, 315, 23),
	(57585, 'TOTOLI', 4378, 315, 23),
	(57586, 'BARUGA', 4379, 315, 23),
	(57587, 'BARUGA DUA', 4379, 315, 23),
	(57588, 'BAURUNG', 4379, 315, 23),
	(57589, 'BUTTU BARUGA', 4379, 315, 23),
	(57590, 'LABUANG', 4379, 315, 23),
	(57591, 'LABUANG UTARA', 4379, 315, 23),
	(57592, 'LEMBANG', 4379, 315, 23),
	(57593, 'TANDE', 4379, 315, 23),
	(57594, 'TANDE TIMUR', 4379, 315, 23),
	(57595, 'BAMBANGAN', 4380, 315, 23),
	(57596, 'KAYUANGIN', 4380, 315, 23),
	(57597, 'LAMUNGANG BATU', 4380, 315, 23),
	(57598, 'LOMBANG', 4380, 315, 23),
	(57599, 'LOMBANG TIMUR', 4380, 315, 23),
	(57600, 'LOMBONG', 4380, 315, 23),
	(57601, 'LOMBONG TIMUR', 4380, 315, 23),
	(57602, 'MALIAYA', 4380, 315, 23),
	(57603, 'MALUNDA', 4380, 315, 23),
	(57604, 'MEKKATTA', 4380, 315, 23),
	(57605, 'MEKKATTA SELATAN', 4380, 315, 23),
	(57606, 'SALUTAHONGAN', 4380, 315, 23),
	(57607, 'ADOLANG', 4381, 315, 23),
	(57608, 'ADOLANG DHUA', 4381, 315, 23),
	(57609, 'BABABULO', 4381, 315, 23),
	(57610, 'BABABULO UTARA', 4381, 315, 23),
	(57611, 'BALOMBONG', 4381, 315, 23),
	(57612, 'BANUA ADOLANG', 4381, 315, 23),
	(57613, 'BETTENG', 4381, 315, 23),
	(57614, 'BONDE', 4381, 315, 23),
	(57615, 'BONDE UTARA', 4381, 315, 23),
	(57616, 'BUTTU PAMBOANG', 4381, 315, 23),
	(57617, 'LALAMPANUA', 4381, 315, 23),
	(57618, 'PESULOANG', 4381, 315, 23),
	(57619, 'SIMBANG', 4381, 315, 23),
	(57620, 'SIRINDU', 4381, 315, 23),
	(57621, 'TINAMBUNG', 4381, 315, 23),
	(57622, 'AWO', 4382, 315, 23),
	(57623, 'MANYAMBA', 4382, 315, 23),
	(57624, 'SEPPONG', 4382, 315, 23),
	(57625, 'TALLAMBALAO', 4382, 315, 23),
	(57626, 'TAMMERODO', 4382, 315, 23),
	(57627, 'TAMMERODO UTARA', 4382, 315, 23),
	(57628, 'ULIDANG', 4382, 315, 23),
	(57629, 'BONDE BONDE', 4383, 315, 23),
	(57630, 'ONANG SELATAN', 4383, 315, 23),
	(57631, 'ONANG UTARA', 4383, 315, 23),
	(57632, 'TUBO', 4383, 315, 23),
	(57633, 'TUBO POANG', 4383, 315, 23),
	(57634, 'TUBO SELATAN', 4383, 315, 23),
	(57635, 'TUBO TENGAH', 4383, 315, 23),
	(57636, 'KABIRAAN', 4384, 315, 23),
	(57637, 'PANGGALO', 4384, 315, 23),
	(57638, 'POPENGA', 4384, 315, 23),
	(57639, 'SALUTAMBUNG', 4384, 315, 23),
	(57640, 'SAMBABO', 4384, 315, 23),
	(57641, 'SULAI', 4384, 315, 23),
	(57642, 'TANDEALLO', 4384, 315, 23),
	(57643, 'ULUMANDA', 4384, 315, 23),
	(57644, 'ARALLE', 4385, 316, 23),
	(57645, 'ARALLE UTARA', 4385, 316, 23),
	(57646, 'BARURU', 4385, 316, 23),
	(57647, 'PAMOSEANG PANGGA', 4385, 316, 23),
	(57648, 'PANETEAN', 4385, 316, 23),
	(57649, 'RALLE ANAK', 4385, 316, 23),
	(57650, 'RALLEANAK UTARA', 4385, 316, 23),
	(57651, 'UHAIDAHO', 4385, 316, 23),
	(57652, 'UHAILANU', 4385, 316, 23),
	(57653, 'BALLA', 4386, 316, 23),
	(57654, 'BALLA BARAT', 4386, 316, 23),
	(57655, 'BALLA SATANETEAN', 4386, 316, 23),
	(57656, 'BALLA TUMUKA', 4386, 316, 23),
	(57657, 'PIDARA', 4386, 316, 23),
	(57658, 'SIPAKUAN', 4386, 316, 23),
	(57659, 'BAMBANG', 4387, 316, 23),
	(57660, 'BAMBANG TIMUR', 4387, 316, 23),
	(57661, 'LEMBANG MOKALLANG', 4387, 316, 23),
	(57662, 'LIMBADEBATA', 4387, 316, 23),
	(57663, 'MASOSO', 4387, 316, 23),
	(57664, 'MINANGA', 4387, 316, 23),
	(57665, 'RANTELEMO', 4387, 316, 23),
	(57666, 'SALUASSING', 4387, 316, 23),
	(57667, 'SALUBULO', 4387, 316, 23),
	(57668, 'SALUKADI', 4387, 316, 23),
	(57669, 'SALUKEPOPOK', 4387, 316, 23),
	(57670, 'SALUNDENGEN', 4387, 316, 23),
	(57671, 'SIKAMASE', 4387, 316, 23),
	(57672, 'TANETE TOMBA', 4387, 316, 23),
	(57673, 'ULUMAMBI', 4387, 316, 23),
	(57674, 'ULUMAMBI BARAT', 4387, 316, 23),
	(57675, 'ARALLE TIMUR', 4388, 316, 23),
	(57676, 'BUNTU MALANGKA', 4388, 316, 23),
	(57677, 'KABAE', 4388, 316, 23),
	(57678, 'KEBANGA', 4388, 316, 23),
	(57679, 'RANTEBERAN (RANTE BERANG)', 4388, 316, 23),
	(57680, 'SALURINDU', 4388, 316, 23),
	(57681, 'SALUTAMBUN', 4388, 316, 23),
	(57682, 'TAORA', 4388, 316, 23),
	(57683, 'BOMBONG LAMBE', 4389, 316, 23),
	(57684, 'BUNTU BUDA', 4389, 316, 23),
	(57685, 'LAMBANAN', 4389, 316, 23),
	(57686, 'MAMASA', 4389, 316, 23),
	(57687, 'MAMBULILING', 4389, 316, 23),
	(57688, 'MBANA SALULO', 4389, 316, 23),
	(57689, 'OSANGO', 4389, 316, 23),
	(57690, 'PEBASSIAN', 4389, 316, 23),
	(57691, 'RAMBUSARATU', 4389, 316, 23),
	(57692, 'TAUPE', 4389, 316, 23),
	(57693, 'TONDOK BAKARU', 4389, 316, 23),
	(57694, 'BUJUNG MANURUNG', 4390, 316, 23),
	(57695, 'MAMBI', 4390, 316, 23),
	(57696, 'PAMOSEANG', 4390, 316, 23),
	(57697, 'RANTE BULAHAN', 4390, 316, 23),
	(57698, 'SALUBANUA', 4390, 316, 23),
	(57699, 'SELUMAKA', 4390, 316, 23),
	(57700, 'SENDANA', 4390, 316, 23),
	(57701, 'SONDONGLAYU', 4390, 316, 23),
	(57702, 'TALIPUKKI', 4390, 316, 23),
	(57703, 'TAPALINA', 4390, 316, 23),
	(57704, 'BOTTENG', 4391, 316, 23),
	(57705, 'MEHALAAN', 4391, 316, 23),
	(57706, 'MESAKADA', 4391, 316, 23),
	(57707, 'SALUKONTA', 4391, 316, 23),
	(57708, 'MAKUANG', 4392, 316, 23),
	(57709, 'MALIMBONG', 4392, 316, 23),
	(57710, 'MATANDE', 4392, 316, 23),
	(57711, 'MESSAWA', 4392, 316, 23),
	(57712, 'PASMA MAMBU', 4392, 316, 23),
	(57713, 'RIPPUNG', 4392, 316, 23),
	(57714, 'SEPANG', 4392, 316, 23),
	(57715, 'SIPAI', 4392, 316, 23),
	(57716, 'BATUPAPAN', 4393, 316, 23),
	(57717, 'MASEWWE', 4393, 316, 23),
	(57718, 'MINANGA', 4393, 316, 23),
	(57719, 'MINANGA TIMUR', 4393, 316, 23),
	(57720, 'NOSU', 4393, 316, 23),
	(57721, 'SIWI', 4393, 316, 23),
	(57722, 'DATUBARINGAN', 4394, 316, 23),
	(57723, 'KARAKA', 4394, 316, 23),
	(57724, 'MAMULLU', 4394, 316, 23),
	(57725, 'MANIPI', 4394, 316, 23),
	(57726, 'PANA', 4394, 316, 23),
	(57727, 'PANURA', 4394, 316, 23),
	(57728, 'SAPAN', 4394, 316, 23),
	(57729, 'TALLANG BULAWAN', 4394, 316, 23),
	(57730, 'ULUSALU', 4394, 316, 23),
	(57731, 'BAMBANG BUDA', 4395, 316, 23),
	(57732, 'BUANGIN', 4395, 316, 23),
	(57733, 'LEKO', 4395, 316, 23),
	(57734, 'SALUMOKANAN', 4395, 316, 23),
	(57735, 'SALUMOKANAN BARAT', 4395, 316, 23),
	(57736, 'SALUMOKANAN UTARA', 4395, 316, 23),
	(57737, 'LISUAN ADA', 4396, 316, 23),
	(57738, 'MALANGKENA PADANG', 4396, 316, 23),
	(57739, 'OROBUA', 4396, 316, 23),
	(57740, 'OROBUA SELATAN', 4396, 316, 23),
	(57741, 'OROBUA TIMUR', 4396, 316, 23),
	(57742, 'PALADAN', 4396, 316, 23),
	(57743, 'RANTE PUANG', 4396, 316, 23),
	(57744, 'SATENETEAN', 4396, 316, 23),
	(57745, 'BANEA', 4397, 316, 23),
	(57746, 'BATANG URU', 4397, 316, 23),
	(57747, 'BATANGNGURU TIMUR', 4397, 316, 23),
	(57748, 'SALUBALO', 4397, 316, 23),
	(57749, 'SASAKAN', 4397, 316, 23),
	(57750, 'SIBANAWA', 4397, 316, 23),
	(57751, 'SUMARORONG', 4397, 316, 23),
	(57752, 'TABONE', 4397, 316, 23),
	(57753, 'TADISI', 4397, 316, 23),
	(57754, 'BAKADISURA', 4398, 316, 23),
	(57755, 'BILA TALANG', 4398, 316, 23),
	(57756, 'BULUK SEN', 4398, 316, 23),
	(57757, 'GUNUNG SARI', 4398, 316, 23),
	(57758, 'KALAMA', 4398, 316, 23),
	(57759, 'KAMPUNG BARU', 4398, 316, 23),
	(57760, 'LONG LALANG', 4398, 316, 23),
	(57761, 'MASUPU', 4398, 316, 23),
	(57762, 'MUARA BELINAU', 4398, 316, 23),
	(57763, 'MUARA KEBAQ', 4398, 316, 23),
	(57764, 'MUARA PEDOHON', 4398, 316, 23),
	(57765, 'MUARA RITAN', 4398, 316, 23),
	(57766, 'MUARA SALUNG', 4398, 316, 23),
	(57767, 'MUARA TIQ', 4398, 316, 23),
	(57768, 'MUARA TUBOQ', 4398, 316, 23),
	(57769, 'RITAN BARU', 4398, 316, 23),
	(57770, 'SIDOMULYO', 4398, 316, 23),
	(57771, 'TABANG', 4398, 316, 23),
	(57772, 'TABANG BARAT', 4398, 316, 23),
	(57773, 'TABANG LAMA', 4398, 316, 23),
	(57774, 'TADOKALUA', 4398, 316, 23),
	(57775, 'TUKUNG RITAN', 4398, 316, 23),
	(57776, 'UMAQ BEKUAY/BEKUAI', 4398, 316, 23),
	(57777, 'UMAQ DIAN', 4398, 316, 23),
	(57778, 'UMAQ TUKUNG', 4398, 316, 23),
	(57779, 'BURANA', 4399, 316, 23),
	(57780, 'GANDANG DEWATA', 4399, 316, 23),
	(57781, 'LAKAHANG', 4399, 316, 23),
	(57782, 'MALATIRO', 4399, 316, 23),
	(57783, 'PANGANDARAN', 4399, 316, 23),
	(57784, 'PERIANGAN', 4399, 316, 23),
	(57785, 'SALUBAKKA', 4399, 316, 23),
	(57786, 'SALULEANG', 4399, 316, 23),
	(57787, 'TABULAHAN', 4399, 316, 23),
	(57788, 'TAMPAK KURRA', 4399, 316, 23),
	(57789, 'BALLABATU', 4400, 316, 23),
	(57790, 'MALABO', 4400, 316, 23),
	(57791, 'MANNABABA', 4400, 316, 23),
	(57792, 'MESAKADA', 4400, 316, 23),
	(57793, 'MINAKE', 4400, 316, 23),
	(57794, 'PARONDOBULAWAN', 4400, 316, 23),
	(57795, 'SINDAGAMANIK', 4400, 316, 23),
	(57796, 'TAMALANTIK', 4400, 316, 23),
	(57797, 'KARIANGO', 4401, 316, 23),
	(57798, 'RANTE TANGNGA', 4401, 316, 23),
	(57799, 'TAWALIAN', 4401, 316, 23),
	(57800, 'TAWALIAN TIMUR', 4401, 316, 23),
	(57801, 'BANUADA', 4402, 317, 23),
	(57802, 'BONEHAU', 4402, 317, 23),
	(57803, 'BUTTUADA', 4402, 317, 23),
	(57804, 'HINUA', 4402, 317, 23),
	(57805, 'KINATANG', 4402, 317, 23),
	(57806, 'LUMIKA', 4402, 317, 23),
	(57807, 'LUMIKA I', 4402, 317, 23),
	(57808, 'MAPPU', 4402, 317, 23),
	(57809, 'SALUTIWO', 4402, 317, 23),
	(57810, 'TAMALEA (TALONDO I)', 4402, 317, 23),
	(57811, 'BABANA', 4403, 317, 23),
	(57812, 'BARAKKANG', 4403, 317, 23),
	(57813, 'BOJO', 4403, 317, 23),
	(57814, 'KIRE', 4403, 317, 23),
	(57815, 'LEMBAH HADA', 4403, 317, 23),
	(57816, 'LUMU', 4403, 317, 23),
	(57817, 'PASSAPA (PASAPA)', 4403, 317, 23),
	(57818, 'PONTANAKAYANG', 4403, 317, 23),
	(57819, 'SALUGATTA (SALOGATTA)', 4403, 317, 23),
	(57820, 'SALUMANURUNG', 4403, 317, 23),
	(57821, 'TINALI', 4403, 317, 23),
	(57822, 'BEBANGA', 4404, 317, 23),
	(57823, 'BELANG-BELANG', 4404, 317, 23),
	(57824, 'BERU-BERU', 4404, 317, 23),
	(57825, 'GULILING', 4404, 317, 23),
	(57826, 'KABULOANG', 4404, 317, 23),
	(57827, 'KALUKKU', 4404, 317, 23),
	(57828, 'KALUKKU BARAT', 4404, 317, 23),
	(57829, 'KEANG', 4404, 317, 23),
	(57830, 'PAMMULUKANG', 4404, 317, 23),
	(57831, 'POKKANG', 4404, 317, 23),
	(57832, 'SINYONYOI', 4404, 317, 23),
	(57833, 'SONDOANG', 4404, 317, 23),
	(57834, 'UHAIMATE', 4404, 317, 23),
	(57835, 'BATU MAKKADA', 4405, 317, 23),
	(57836, 'KALUMPANG', 4405, 317, 23),
	(57837, 'KARAMA', 4405, 317, 23),
	(57838, 'KARATAUN', 4405, 317, 23),
	(57839, 'KONDO BULO', 4405, 317, 23),
	(57840, 'LASA', 4405, 317, 23),
	(57841, 'LIMBONG', 4405, 317, 23),
	(57842, 'MAKKALIKI', 4405, 317, 23),
	(57843, 'POLIO (BA\\SAN)', 4405, 317, 23),
	(57844, 'SALUMAKKI', 4405, 317, 23),
	(57845, 'SANDAPANG', 4405, 317, 23),
	(57846, 'SIRAUN', 4405, 317, 23),
	(57847, 'TUMONGA', 4405, 317, 23),
	(57848, 'KADAILA', 4406, 317, 23),
	(57849, 'KAMBUNONG', 4406, 317, 23),
	(57850, 'KAROSSA', 4406, 317, 23),
	(57851, 'KAYUCALLA', 4406, 317, 23),
	(57852, 'LARA', 4406, 317, 23),
	(57853, 'LARA III', 4406, 317, 23),
	(57854, 'LEMBAH HOPO', 4406, 317, 23),
	(57855, 'MORA BENGGAULU', 4406, 317, 23),
	(57856, 'MORA IV', 4406, 317, 23),
	(57857, 'SUKAMAJU', 4406, 317, 23),
	(57858, 'TASOKKO', 4406, 317, 23),
	(57859, 'BALA BALAKANG', 4407, 317, 23),
	(57860, 'BALA BALAKANG TIMUR', 4407, 317, 23),
	(57861, 'BAMBU', 4408, 317, 23),
	(57862, 'BATU PANNU', 4408, 317, 23),
	(57863, 'BINANGA', 4408, 317, 23),
	(57864, 'KARAMPUANG', 4408, 317, 23),
	(57865, 'KAREMA', 4408, 317, 23),
	(57866, 'MAMUNYU', 4408, 317, 23),
	(57867, 'RIMUKU', 4408, 317, 23),
	(57868, 'TADUI', 4408, 317, 23),
	(57869, 'KOMBILING', 4409, 317, 23),
	(57870, 'KUO', 4409, 317, 23),
	(57871, 'LAMBA-LAMBA', 4409, 317, 23),
	(57872, 'LEMO-LEMO', 4409, 317, 23),
	(57873, 'PANGALE', 4409, 317, 23),
	(57874, 'POLOCAMBA', 4409, 317, 23),
	(57875, 'POLOLERENG', 4409, 317, 23),
	(57876, 'POLOPANGALE', 4409, 317, 23),
	(57877, 'SARTANAMA', 4409, 317, 23),
	(57878, 'BATU AMPA', 4410, 317, 23),
	(57879, 'BODA-BODA', 4410, 317, 23),
	(57880, 'BONDA', 4410, 317, 23),
	(57881, 'PAPALANG', 4410, 317, 23),
	(57882, 'SALUKAYU (SALOKAYU)', 4410, 317, 23),
	(57883, 'SISANGO', 4410, 317, 23),
	(57884, 'SUKADAMAI', 4410, 317, 23),
	(57885, 'TOABO', 4410, 317, 23),
	(57886, 'TOPARE (TAPORE)', 4410, 317, 23),
	(57887, 'BUNDE', 4411, 317, 23),
	(57888, 'KALONDING', 4411, 317, 23),
	(57889, 'LOSSO', 4411, 317, 23),
	(57890, 'SALUBARANA', 4411, 317, 23),
	(57891, 'SAMPAGA', 4411, 317, 23),
	(57892, 'TANAM BUAH (TANAMBUA)', 4411, 317, 23),
	(57893, 'TARAILU', 4411, 317, 23),
	(57894, 'BOTTENG', 4412, 317, 23),
	(57895, 'BOTTENG UTARA', 4412, 317, 23),
	(57896, 'KUPTD BOTTENG I', 4412, 317, 23),
	(57897, 'KUPTD BOTTENG II', 4412, 317, 23),
	(57898, 'RANGAS', 4412, 317, 23),
	(57899, 'SALLETTO', 4412, 317, 23),
	(57900, 'SIMBORO', 4412, 317, 23),
	(57901, 'SUMARE', 4412, 317, 23),
	(57902, 'BELA', 4413, 317, 23),
	(57903, 'GALUNG', 4413, 317, 23),
	(57904, 'KASAMBANG', 4413, 317, 23),
	(57905, 'KOPEANG', 4413, 317, 23),
	(57906, 'OROBATU', 4413, 317, 23),
	(57907, 'RANTEDODA', 4413, 317, 23),
	(57908, 'TAANG (TA\\AN)', 4413, 317, 23),
	(57909, 'TAKANDEANG', 4413, 317, 23),
	(57910, 'TAMPALANG', 4413, 317, 23),
	(57911, 'AHU', 4414, 317, 23),
	(57912, 'DUNGKAIT', 4414, 317, 23),
	(57913, 'LABUANG RANO', 4414, 317, 23),
	(57914, 'LEBANI', 4414, 317, 23),
	(57915, 'PANGASAAN', 4414, 317, 23),
	(57916, 'PASABU (PASA BU)', 4414, 317, 23),
	(57917, 'TANETE PAO', 4414, 317, 23),
	(57918, 'BAMBADARU', 4415, 317, 23),
	(57919, 'BATU PARIGI', 4415, 317, 23),
	(57920, 'MAHAHE', 4415, 317, 23),
	(57921, 'PALONGAAN', 4415, 317, 23),
	(57922, 'SALUADAK (SALO ADAK)', 4415, 317, 23),
	(57923, 'SALUBAJA (SALO BAJA)', 4415, 317, 23),
	(57924, 'SEJATI', 4415, 317, 23),
	(57925, 'TOBADAK', 4415, 317, 23),
	(57926, 'BUANA SAKTI', 4416, 317, 23),
	(57927, 'CAMPALOGA', 4416, 317, 23),
	(57928, 'KAKULASANG (KAKULLASAN)', 4416, 317, 23),
	(57929, 'KALEPU', 4416, 317, 23),
	(57930, 'LELING', 4416, 317, 23),
	(57931, 'LELING BARAT', 4416, 317, 23),
	(57932, 'LELING UTARA', 4416, 317, 23),
	(57933, 'MALINO', 4416, 317, 23),
	(57934, 'RANTE MARIO', 4416, 317, 23),
	(57935, 'SALUDENGEN', 4416, 317, 23),
	(57936, 'SANDANA', 4416, 317, 23),
	(57937, 'TAMEJARRA (TAMMEJARRA)', 4416, 317, 23),
	(57938, 'TAMEMONGGA', 4416, 317, 23),
	(57939, 'TOMMO', 4416, 317, 23),
	(57940, 'BAMBAMANURUNG', 4417, 317, 23),
	(57941, 'BUDONG-BUDONG', 4417, 317, 23),
	(57942, 'KABUBU', 4417, 317, 23),
	(57943, 'PANGALLOANG', 4417, 317, 23),
	(57944, 'PARAILI', 4417, 317, 23),
	(57945, 'SALULEKBO', 4417, 317, 23),
	(57946, 'SALUPANGKANG (SALOPANGKANG)', 4417, 317, 23),
	(57947, 'SALUPANGKANG IV (SALUPANGKANG IV)', 4417, 317, 23),
	(57948, 'SINABATTA', 4417, 317, 23),
	(57949, 'TABOLANG', 4417, 317, 23),
	(57950, 'TANGKAU (TANGKOU)', 4417, 317, 23),
	(57951, 'TAPILINA (TAPPILINA)', 4417, 317, 23),
	(57952, 'TOPOYO', 4417, 317, 23),
	(57953, 'TUMBU', 4417, 317, 23),
	(57954, 'WAEPUTEH', 4417, 317, 23),
	(57955, 'BAMBAIRA', 4418, 318, 23),
	(57956, 'KALUKUNANGKA', 4418, 318, 23),
	(57957, 'KASOLOANG', 4418, 318, 23),
	(57958, 'TAMPAURE', 4418, 318, 23),
	(57959, 'BAMBALAMOTU', 4419, 318, 23),
	(57960, 'KALOLA', 4419, 318, 23),
	(57961, 'PANGIANG', 4419, 318, 23),
	(57962, 'POLEWALI', 4419, 318, 23),
	(57963, 'RANDOMAYANG', 4419, 318, 23),
	(57964, 'WULAI', 4419, 318, 23),
	(57965, 'BALANTI', 4420, 318, 23),
	(57966, 'BARAS', 4420, 318, 23),
	(57967, 'BULU PARIGI', 4420, 318, 23),
	(57968, 'KASANO', 4420, 318, 23),
	(57969, 'MOTU', 4420, 318, 23),
	(57970, 'TOWONI (TORONI)', 4420, 318, 23),
	(57971, 'BUKIT HARAPAN', 4421, 318, 23),
	(57972, 'KARAVE (KARABE)', 4421, 318, 23),
	(57973, 'KASTA BUANA', 4421, 318, 23),
	(57974, 'LELEJAE', 4421, 318, 23),
	(57975, 'LILIMORI (LELEMORI)', 4421, 318, 23),
	(57976, 'OMPI', 4421, 318, 23),
	(57977, 'SUMBERSARI', 4421, 318, 23),
	(57978, 'BENGGAULU', 4422, 318, 23),
	(57979, 'BULU BONGGU', 4422, 318, 23),
	(57980, 'DAPURANG', 4422, 318, 23),
	(57981, 'SERASA', 4422, 318, 23),
	(57982, 'TIRTA BUANA', 4422, 318, 23),
	(57983, 'SAPTANAJAYA', 4423, 318, 23),
	(57984, 'SIPAKAINGA', 4423, 318, 23),
	(57985, 'TAMMARUNNANG', 4423, 318, 23),
	(57986, 'TARANGGI', 4423, 318, 23),
	(57987, 'BAJAWALI', 4424, 318, 23),
	(57988, 'BAMBAKORO', 4424, 318, 23),
	(57989, 'BATU MATORU', 4424, 318, 23),
	(57990, 'KENANGAN', 4424, 318, 23),
	(57991, 'KULU', 4424, 318, 23),
	(57992, 'PARABU', 4424, 318, 23),
	(57993, 'SINGGANI', 4424, 318, 23),
	(57994, 'AKO', 4425, 318, 23),
	(57995, 'KARYA BERSAMA', 4425, 318, 23),
	(57996, 'MALEI', 4425, 318, 23),
	(57997, 'MARTAJAYA', 4425, 318, 23),
	(57998, 'PASANGKAYU', 4425, 318, 23),
	(57999, 'PEDANDA', 4425, 318, 23),
	(58000, 'BATU OGE', 4426, 318, 23)
	`)
	if err != nil {
		panic(err)
	}
	fmt.Println("insert villige 57")
}