package kel0

import "github.com/danangkonang/ceodeaja-go/config"

func Kel7() {
	db := config.Connect()
	db.Exec(`INSERT INTO kelurahan (id,kelurahan,kecamatan_id,kota_id,provinsi_id) VALUES
		(7001, 'GIRIMUKTI', 684, 33, 3),
		(7002, 'KERTANEGLA', 684, 33, 3),
		(7003, 'MANGKONJAYA', 684, 33, 3),
		(7004, 'PEDANGKAMULYAN', 684, 33, 3),
		(7005, 'PURWARAHARJA', 684, 33, 3),
		(7006, 'WANDASARI', 684, 33, 3),
		(7007, 'BANTARSARI', 685, 33, 3),
		(7008, 'BUNGURSARI', 685, 33, 3),
		(7009, 'BUNGURSARI', 685, 33, 3),
		(7010, 'CIBENING', 685, 33, 3),
		(7011, 'CIBODAS', 685, 33, 3),
		(7012, 'CIBUNGUR', 685, 33, 3),
		(7013, 'CIBUNIGEULIS', 685, 33, 3),
		(7014, 'CIKOPO', 685, 33, 3),
		(7015, 'CINANGKA', 685, 33, 3),
		(7016, 'CIWANGI', 685, 33, 3),
		(7017, 'DANGDEUR', 685, 33, 3),
		(7018, 'KARANGMUKTI', 685, 33, 3),
		(7019, 'SUKAJAYA', 685, 33, 3),
		(7020, 'SUKALAKSANA', 685, 33, 3),
		(7021, 'SUKAMULYA', 685, 33, 3),
		(7022, 'SUKARINDIK', 685, 33, 3),
		(7023, 'WANAKERTA', 685, 33, 3),
		(7024, 'BANJAR SARI', 686, 33, 3),
		(7025, 'BANJAR WANGI', 686, 33, 3),
		(7026, 'BANJAR WARU', 686, 33, 3),
		(7027, 'BENDUNGAN', 686, 33, 3),
		(7028, 'BITUNG SARI', 686, 33, 3),
		(7029, 'BOJONG MURNI', 686, 33, 3),
		(7030, 'BUGEL', 686, 33, 3),
		(7031, 'CIAWI', 686, 33, 3),
		(7032, 'CIAWI', 686, 33, 3),
		(7033, 'CIBEDUG', 686, 33, 3),
		(7034, 'CILEUNGSI', 686, 33, 3),
		(7035, 'CITAMBA', 686, 33, 3),
		(7036, 'CITAPEN', 686, 33, 3),
		(7037, 'GOMBONG', 686, 33, 3),
		(7038, 'JAMBU LUWUK', 686, 33, 3),
		(7039, 'KERTAMUKTI', 686, 33, 3),
		(7040, 'KURNIABAKTI', 686, 33, 3),
		(7041, 'MARGASARI', 686, 33, 3),
		(7042, 'PAKEMITAN', 686, 33, 3),
		(7043, 'PAKEMITANKIDUL', 686, 33, 3),
		(7044, 'PANDANSARI', 686, 33, 3),
		(7045, 'PASIRHUNI', 686, 33, 3),
		(7046, 'SUKAMANTRI', 686, 33, 3),
		(7047, 'TELUK PINANG', 686, 33, 3),
		(7048, 'CIBALONG', 687, 33, 3),
		(7049, 'CIGARONGGONG', 687, 33, 3),
		(7050, 'CISEMPUR', 687, 33, 3),
		(7051, 'EUREUNPALAY', 687, 33, 3),
		(7052, 'KARYA MUKTI', 687, 33, 3),
		(7053, 'KARYASARI', 687, 33, 3),
		(7054, 'MAROKO', 687, 33, 3),
		(7055, 'MEKARMUKTI', 687, 33, 3),
		(7056, 'MEKARSARI', 687, 33, 3),
		(7057, 'MEKARWANGI', 687, 33, 3),
		(7058, 'NAJATEN', 687, 33, 3),
		(7059, 'PARUNG', 687, 33, 3),
		(7060, 'SAGARA', 687, 33, 3),
		(7061, 'SANCANG', 687, 33, 3),
		(7062, 'SETIAWARAS', 687, 33, 3),
		(7063, 'SIMPANG', 687, 33, 3),
		(7064, 'SINGAJAYA', 687, 33, 3),
		(7065, 'AWIPARI', 688, 33, 3),
		(7066, 'BABAKAN', 688, 33, 3),
		(7067, 'CIAKAR', 688, 33, 3),
		(7068, 'CIBEUREUM', 688, 33, 3),
		(7069, 'CIBEUREUMHILIR', 688, 33, 3),
		(7070, 'CIHERANG', 688, 33, 3),
		(7071, 'CIMARA', 688, 33, 3),
		(7072, 'KAWUNGSARI', 688, 33, 3),
		(7073, 'KERSANAGARA', 688, 33, 3),
		(7074, 'KOTABARU', 688, 33, 3),
		(7075, 'LIMUSNUNGGAL', 688, 33, 3),
		(7076, 'MARGABAKTI', 688, 33, 3),
		(7077, 'RANDUSARI', 688, 33, 3),
		(7078, 'SETIAJAYA', 688, 33, 3),
		(7079, 'SETIANAGARA', 688, 33, 3),
		(7080, 'SETIARATU', 688, 33, 3),
		(7081, 'SINDANGPALAY', 688, 33, 3),
		(7082, 'SUKADANA', 688, 33, 3),
		(7083, 'SUKARAPIH', 688, 33, 3),
		(7084, 'SUMURWIRU', 688, 33, 3),
		(7085, 'TARIKOLOT', 688, 33, 3),
		(7086, 'CIDUGALEUN', 689, 33, 3),
		(7087, 'CIGALONTANG', 689, 33, 3),
		(7088, 'JAYAPURA (JAYAPUTRA)', 689, 33, 3),
		(7089, 'KERSAMAJU', 689, 33, 3),
		(7090, 'LENGKONGJAYA', 689, 33, 3),
		(7091, 'NANGGERANG', 689, 33, 3),
		(7092, 'NANGTANG', 689, 33, 3),
		(7093, 'PARENTAS', 689, 33, 3),
		(7094, 'PUSPAMUKTI', 689, 33, 3),
		(7095, 'PUSPARAJA', 689, 33, 3),
		(7096, 'SIRNAGALIH', 689, 33, 3),
		(7097, 'SIRNAPUTRA', 689, 33, 3),
		(7098, 'SIRNARAJA', 689, 33, 3),
		(7099, 'SUKAMANAH', 689, 33, 3),
		(7100, 'TANJUNGKARANG', 689, 33, 3),
		(7101, 'TENJONAGARA', 689, 33, 3),
		(7102, 'ARGASARI', 690, 33, 3),
		(7103, 'CILEMBANG', 690, 33, 3),
		(7104, 'NAGARAWANGI', 690, 33, 3),
		(7105, 'TUGUJAYA', 690, 33, 3),
		(7106, 'TUGURAJA', 690, 33, 3),
		(7107, 'YUDANAGARA', 690, 33, 3),
		(7108, 'CIBEBER', 691, 33, 3),
		(7109, 'CIDADALI', 691, 33, 3),
		(7110, 'CIKADU', 691, 33, 3),
		(7111, 'CIKALONG', 691, 33, 3),
		(7112, 'CIKANCRA', 691, 33, 3),
		(7113, 'CIMANUK', 691, 33, 3),
		(7114, 'KALAPAGENEP', 691, 33, 3),
		(7115, 'KUBANGSARI', 691, 33, 3),
		(7116, 'MANDALAJAYA', 691, 33, 3),
		(7117, 'PANYIARAN', 691, 33, 3),
		(7118, 'SINDANGJAYA', 691, 33, 3),
		(7119, 'SINGKIR', 691, 33, 3),
		(7120, 'TONJONGSARI', 691, 33, 3),
		(7121, 'CAYUR', 692, 33, 3),
		(7122, 'CILUMBA', 692, 33, 3),
		(7123, 'COGREG', 692, 33, 3),
		(7124, 'GUNUNGSARI', 692, 33, 3),
		(7125, 'LENGKONGBARANG', 692, 33, 3),
		(7126, 'LINGGALAKSANA', 692, 33, 3),
		(7127, 'PAKEMITAN', 692, 33, 3),
		(7128, 'SINDANGASIH', 692, 33, 3),
		(7129, 'TANJUNGBARANG', 692, 33, 3),
		(7130, 'ANCOL', 693, 33, 3),
		(7131, 'CIAMPANAN', 693, 33, 3),
		(7132, 'CIJULANG', 693, 33, 3),
		(7133, 'CIKONDANG', 693, 33, 3),
		(7134, 'CINEAM', 693, 33, 3),
		(7135, 'CISARUA', 693, 33, 3),
		(7136, 'MADIASARI', 693, 33, 3),
		(7137, 'NAGARATENGAH', 693, 33, 3),
		(7138, 'PASIRMUKTI', 693, 33, 3),
		(7139, 'RAJADATU', 693, 33, 3),
		(7140, 'BANTARKALONG', 694, 33, 3),
		(7141, 'CIANDUM', 694, 33, 3),
		(7142, 'CIHERAS', 694, 33, 3),
		(7143, 'CIKAWUNGADING', 694, 33, 3),
		(7144, 'CIPANAS', 694, 33, 3),
		(7145, 'CIPATUJAH', 694, 33, 3),
		(7146, 'DARAWATI', 694, 33, 3),
		(7147, 'KERTASARI', 694, 33, 3),
		(7148, 'NAGROG', 694, 33, 3),
		(7149, 'NANGELASARI', 694, 33, 3),
		(7150, 'PADAWARAS', 694, 33, 3),
		(7151, 'PAMEUTINGAN', 694, 33, 3),
		(7152, 'SINDANGKERTA', 694, 33, 3),
		(7153, 'SUKAHURIP', 694, 33, 3),
		(7154, 'TOBONGJAYA', 694, 33, 3),
		(7155, 'CIPEDES', 695, 33, 3),
		(7156, 'NAGARASARI', 695, 33, 3),
		(7157, 'PANGLAYUNGAN', 695, 33, 3),
		(7158, 'SUKAMANAH', 695, 33, 3),
		(7159, 'CIKADU', 696, 33, 3),
		(7160, 'CILEULEUS', 696, 33, 3),
		(7161, 'CISAYONG', 696, 33, 3),
		(7162, 'JATIHURIP', 696, 33, 3),
		(7163, 'MEKARWANGI', 696, 33, 3),
		(7164, 'NUSAWANGI', 696, 33, 3),
		(7165, 'PURWASARI', 696, 33, 3),
		(7166, 'SANTANAMEKAR', 696, 33, 3),
		(7167, 'SUKAJADI', 696, 33, 3),
		(7168, 'SUKAMUKTI', 696, 33, 3),
		(7169, 'SUKARAHARJA', 696, 33, 3),
		(7170, 'SUKASETIA', 696, 33, 3),
		(7171, 'SUKASUKUR', 696, 33, 3),
		(7172, 'BOJONGSARI', 697, 33, 3),
		(7173, 'CIKUYA', 697, 33, 3),
		(7174, 'CINTABODAS', 697, 33, 3),
		(7175, 'CIPICUNG', 697, 33, 3),
		(7176, 'MEKARLAKSANA', 697, 33, 3),
		(7177, 'BOJONGSARI', 698, 33, 3),
		(7178, 'CINUNJANG', 698, 33, 3),
		(7179, 'GIRIWANGI', 698, 33, 3),
		(7180, 'GUNUNG TANJUNG', 698, 33, 3),
		(7181, 'JATIJAYA', 698, 33, 3),
		(7182, 'MALATISUKA', 698, 33, 3),
		(7183, 'TANJUNGSARI', 698, 33, 3),
		(7184, 'INDIHIANG', 699, 33, 3),
		(7185, 'PANYINGKIRAN', 699, 33, 3),
		(7186, 'PARAKANNYASAG', 699, 33, 3),
		(7187, 'SIRNAGALIH', 699, 33, 3),
		(7188, 'SUKAMAJUKALER', 699, 33, 3),
		(7189, 'SUKAMAJUKIDUL', 699, 33, 3),
		(7190, 'BOJONGGAOK', 700, 33, 3),
		(7191, 'CONDONG', 700, 33, 3),
		(7192, 'GERESIK', 700, 33, 3),
		(7193, 'KARANGMULYA', 700, 33, 3),
		(7194, 'KARANGRESIK', 700, 33, 3),
		(7195, 'KARANGSEMBUNG', 700, 33, 3),
		(7196, 'SINDANGRAJA', 700, 33, 3),
		(7197, 'TANJUNGMEKAR', 700, 33, 3),
		(7198, 'CIWARAK', 701, 33, 3),
		(7199, 'JATIWARAS', 701, 33, 3),
		(7200, 'KAPUTIHAN', 701, 33, 3),
		(7201, 'KERSAGALIH', 701, 33, 3),
		(7202, 'KERTARAHAYU', 701, 33, 3),
		(7203, 'MANDALAHURIP', 701, 33, 3),
		(7204, 'MANDALAMEKAR', 701, 33, 3),
		(7205, 'NEGLASARI', 701, 33, 3),
		(7206, 'PAPAYAN', 701, 33, 3),
		(7207, 'SETIAWANGI', 701, 33, 3),
		(7208, 'SUKAKERTA', 701, 33, 3),
		(7209, 'BABAKAN ANYAR', 702, 33, 3),
		(7210, 'BUNIASIH', 702, 33, 3),
		(7211, 'CIBAHAYU', 702, 33, 3),
		(7212, 'CIPAKU', 702, 33, 3),
		(7213, 'DIRGAHAYU', 702, 33, 3),
		(7214, 'HEULEUT', 702, 33, 3),
		(7215, 'KADIPATEN', 702, 33, 3),
		(7216, 'KADIPATEN', 702, 33, 3),
		(7217, 'KARANGSAMBUNG', 702, 33, 3),
		(7218, 'LIANGJULANG', 702, 33, 3),
		(7219, 'MEKARSARI', 702, 33, 3),
		(7220, 'PAGANDON', 702, 33, 3),
		(7221, 'PAMOYANAN', 702, 33, 3),
		(7222, 'BUKIT LANGKAP', 703, 33, 3),
		(7223, 'BUKIT ULU', 703, 33, 3),
		(7224, 'CITALAHAB', 703, 33, 3),
		(7225, 'EMBACANG BARU', 703, 33, 3),
		(7226, 'EMBACANG LAMA', 703, 33, 3),
		(7227, 'KARANG JAYA', 703, 33, 3),
		(7228, 'KARANGJAYA', 703, 33, 3),
		(7229, 'KARANGLAYUNG', 703, 33, 3),
		(7230, 'LUBUK KUMBANG/KUMBUNG', 703, 33, 3),
		(7231, 'MUARA BATANG EMPU', 703, 33, 3),
		(7232, 'MUARA TIKU', 703, 33, 3),
		(7233, 'RANTAU JAYA', 703, 33, 3),
		(7234, 'RANTAU TELANG', 703, 33, 3),
		(7235, 'SIRNAJAYA', 703, 33, 3),
		(7236, 'SUKAMENANG', 703, 33, 3),
		(7237, 'SUKARAJA', 703, 33, 3),
		(7238, 'TANJUNG AGUNG', 703, 33, 3),
		(7239, 'TERUSAN', 703, 33, 3),
		(7240, 'CIAWI', 704, 33, 3),
		(7241, 'CIBATU', 704, 33, 3),
		(7242, 'CIBATUIRENG', 704, 33, 3),
		(7243, 'CIDADAP', 704, 33, 3),
		(7244, 'CIKAPINIS', 704, 33, 3),
		(7245, 'CIKUKULU', 704, 33, 3),
		(7246, 'CIKUPA', 704, 33, 3),
		(7247, 'CINTAWANGI', 704, 33, 3),
		(7248, 'KARANGMEKAR', 704, 33, 3),
		(7249, 'KARANGNUNGGAL', 704, 33, 3),
		(7250, 'KUJANG', 704, 33, 3),
		(7251, 'SARIMANGGU', 704, 33, 3),
		(7252, 'SARIMUKTI', 704, 33, 3),
		(7253, 'SUKAWANGUN', 704, 33, 3),
		(7254, 'CIBEUTI', 705, 33, 3),
		(7255, 'CILAMAJANG', 705, 33, 3),
		(7256, 'GUNUNGGEDE', 705, 33, 3),
		(7257, 'GUNUNGTANDALA', 705, 33, 3),
		(7258, 'KARANGANYAR', 705, 33, 3),
		(7259, 'KARSAMENAK', 705, 33, 3),
		(7260, 'LEUWILIANG', 705, 33, 3),
		(7261, 'TALAGASARI', 705, 33, 3),
		(7262, 'TANJUNG', 705, 33, 3),
		(7263, 'URUG', 705, 33, 3),
		(7264, 'ARJASARI', 706, 33, 3),
		(7265, 'CIAWANG', 706, 33, 3),
		(7266, 'CIGADOG', 706, 33, 3),
		(7267, 'JAYAMUKTI', 706, 33, 3),
		(7268, 'LINGGAMULYA', 706, 33, 3),
		(7269, 'LINGGAWANGI', 706, 33, 3),
		(7270, 'MANDALAGIRI', 706, 33, 3),
		(7271, 'CIGANTANG', 707, 33, 3),
		(7272, 'CIPARI', 707, 33, 3),
		(7273, 'CIPAWITRA', 707, 33, 3),
		(7274, 'KARIKIL', 707, 33, 3),
		(7275, 'LINGGAJAYA', 707, 33, 3),
		(7276, 'MANGKUBUMI', 707, 33, 3),
		(7277, 'SAMBONGJAYA', 707, 33, 3),
		(7278, 'SAMBONGPARI', 707, 33, 3),
		(7279, 'MANGUNREJA', 708, 33, 3),
		(7280, 'MARGAJAYA', 708, 33, 3),
		(7281, 'PASIRSALAM', 708, 33, 3),
		(7282, 'SALEBU', 708, 33, 3),
		(7283, 'SUKALUYU', 708, 33, 3),
		(7284, 'SUKASUKUR', 708, 33, 3),
		(7285, 'BATUSUMUR', 709, 33, 3),
		(7286, 'CIBEBER', 709, 33, 3),
		(7287, 'CIHAUR', 709, 33, 3),
		(7288, 'CILANGKAP', 709, 33, 3),
		(7289, 'GUNAJAYA', 709, 33, 3),
		(7290, 'KALIMANGGIS', 709, 33, 3),
		(7291, 'KAMULYAN', 709, 33, 3),
		(7292, 'MANONJAYA', 709, 33, 3),
		(7293, 'MARGAHAYU', 709, 33, 3),
		(7294, 'MARGALUYU', 709, 33, 3),
		(7295, 'PASIRBATANG', 709, 33, 3),
		(7296, 'PASIRPANJANG', 709, 33, 3),
		(7297, 'CILAMPUNGHILIR', 710, 33, 3),
		(7298, 'CISARUNI', 710, 33, 3),
		(7299, 'MEKARJAYA', 710, 33, 3),
		(7300, 'PADAKEMBANG', 710, 33, 3),
		(7301, 'RANCAPAKU', 710, 33, 3),
		(7302, 'CIPACING', 711, 33, 3),
		(7303, 'GURANTENG', 711, 33, 3),
		(7304, 'NANGGEWER', 711, 33, 3),
		(7305, 'PAGERAGEUNG', 711, 33, 3),
		(7306, 'PAGERSARI', 711, 33, 3),
		(7307, 'PUTERAN', 711, 33, 3),
		(7308, 'SUKADANA', 711, 33, 3),
		(7309, 'SUKAMAJU', 711, 33, 3),
		(7310, 'SUKAPADA', 711, 33, 3),
		(7311, 'TANJUNGKERTA', 711, 33, 3),
		(7312, 'CIBONGAS', 712, 33, 3),
		(7313, 'CIBUNIASIH', 712, 33, 3),
		(7314, 'CIKAWUNG', 712, 33, 3),
		(7315, 'JAYAMUKTI', 712, 33, 3),
		(7316, 'MARGALUYU', 712, 33, 3),
		(7317, 'MEKARSARI', 712, 33, 3),
		(7318, 'NEGLASARI', 712, 33, 3),
		(7319, 'PANCAWANGI', 712, 33, 3),
		(7320, 'PANGLIARAN', 712, 33, 3),
		(7321, 'TAWANG', 712, 33, 3),
		(7322, 'TONJONG', 712, 33, 3),
		(7323, 'BARUMEKAR', 713, 33, 3),
		(7324, 'BURUJULJAYA', 713, 33, 3),
		(7325, 'CIBANTENG', 713, 33, 3),
		(7326, 'CIBUNGUR', 713, 33, 3),
		(7327, 'CIGUNUNG', 713, 33, 3),
		(7328, 'GIRIKENCANA', 713, 33, 3),
		(7329, 'KARYABAKTI', 713, 33, 3),
		(7330, 'PARUNGPONTENG', 713, 33, 3),
		(7331, 'PURBARATU', 714, 33, 3),
		(7332, 'SINGKUP', 714, 33, 3),
		(7333, 'SUKAASIH', 714, 33, 3),
		(7334, 'SUKAJAYA', 714, 33, 3),
		(7335, 'SUKAMENAK', 714, 33, 3),
		(7336, 'SUKANAGARA', 714, 33, 3),
		(7337, 'CIMANGGU', 715, 33, 3),
		(7338, 'LUYUBAKTI', 715, 33, 3),
		(7339, 'MANDALASARI', 715, 33, 3),
		(7340, 'PUSPAHIANG', 715, 33, 3),
		(7341, 'PUSPAJAYA', 715, 33, 3),
		(7342, 'PUSPARAHAYU', 715, 33, 3),
		(7343, 'PUSPASARI', 715, 33, 3),
		(7344, 'SUKASARI', 715, 33, 3),
		(7345, 'DAWAGUNG', 716, 33, 3),
		(7346, 'MANGGUNGJAYA', 716, 33, 3),
		(7347, 'MANGGUNGSARI', 716, 33, 3),
		(7348, 'RAJAMANDALA', 716, 33, 3),
		(7349, 'RAJAPOLAH', 716, 33, 3),
		(7350, 'SUKANAGALIH', 716, 33, 3),
		(7351, 'SUKARAJA', 716, 33, 3),
		(7352, 'TANJUNGPURA', 716, 33, 3),
		(7353, 'JAHIANG', 717, 33, 3),
		(7354, 'KARANGMUKTI', 717, 33, 3),
		(7355, 'KAWUNGSARI', 717, 33, 3),
		(7356, 'KUTAWARINGIN', 717, 33, 3),
		(7357, 'MARGALAKSANA', 717, 33, 3),
		(7358, 'NEGLASARI', 717, 33, 3),
		(7359, 'SALAWU', 717, 33, 3),
		(7360, 'SERANG', 717, 33, 3),
		(7361, 'SUKARASA', 717, 33, 3),
		(7362, 'SUNDAWENANG', 717, 33, 3),
		(7363, 'TANJUNGSARI', 717, 33, 3),
		(7364, 'TENJOWARINGIN', 717, 33, 3),
		(7365, 'BANJARWARINGIN', 718, 33, 3),
		(7366, 'KARYAMANDALA', 718, 33, 3),
		(7367, 'KARYAWANGI', 718, 33, 3),
		(7368, 'KAWITAN', 718, 33, 3),
		(7369, 'MANDALAGUNA', 718, 33, 3),
		(7370, 'MANDALAHAYU', 718, 33, 3),
		(7371, 'MANDALAWANGI', 718, 33, 3),
		(7372, 'MULYASARI', 718, 33, 3),
		(7373, 'TANJUNGSARI', 718, 33, 3),
		(7374, 'JAYAPUTRA', 719, 33, 3),
		(7375, 'JAYARATU', 719, 33, 3),
		(7376, 'LINGGASIRNA', 719, 33, 3),
		(7377, 'SARIWANGI', 719, 33, 3),
		(7378, 'SELAWANGI', 719, 33, 3),
		(7379, 'SIRNASARI', 719, 33, 3),
		(7380, 'SUKAHARJA (SUKARAHARJA)', 719, 33, 3),
		(7381, 'SUKAMULIH', 719, 33, 3),
		(7382, 'CIKADONGDONG', 720, 33, 3),
		(7383, 'CIKUNIR', 720, 33, 3),
		(7384, 'CIKUNTEN', 720, 33, 3),
		(7385, 'CINTARAJA', 720, 33, 3),
		(7386, 'CIPAKAT', 720, 33, 3),
		(7387, 'SINGAPARNA', 720, 33, 3),
		(7388, 'SINGASARI', 720, 33, 3),
		(7389, 'SUKAASIH', 720, 33, 3),
		(7390, 'SUKAHERANG', 720, 33, 3),
		(7391, 'SUKAMULYA', 720, 33, 3),
		(7392, 'CIKALONG', 721, 33, 3),
		(7393, 'CIPAINGEUN', 721, 33, 3),
		(7394, 'CUKANGJAYAGUNA', 721, 33, 3),
		(7395, 'CUKANGKAWUNG', 721, 33, 3),
		(7396, 'LEUWIDULANG', 721, 33, 3),
		(7397, 'MUNCANG', 721, 33, 3),
		(7398, 'PAKALONGAN', 721, 33, 3),
		(7399, 'PARUMASAN', 721, 33, 3),
		(7400, 'RAKSAJAYA', 721, 33, 3),
		(7401, 'SEPATNUNGGAL', 721, 33, 3),
		(7402, 'SODONGHILIR', 721, 33, 3),
		(7403, 'SUKABAKTI', 721, 33, 3),
		(7404, 'BANYURASA', 722, 33, 3),
		(7405, 'BANYURESMI', 722, 33, 3),
		(7406, 'CALINGCING', 722, 33, 3),
		(7407, 'KIARAJANGKUNG', 722, 33, 3),
		(7408, 'KUDADEPA', 722, 33, 3),
		(7409, 'SUKAHENING', 722, 33, 3),
		(7410, 'SUNDAKERTA', 722, 33, 3),
		(7411, 'AIR KEMUNING', 723, 33, 3),
		(7412, 'AIR PETAI', 723, 33, 3),
		(7413, 'BABATAN', 723, 33, 3),
		(7414, 'BUKIT PENINJAUAN I', 723, 33, 3),
		(7415, 'BUKIT PENINJAUAN II', 723, 33, 3),
		(7416, 'CADAS NGAMPAR', 723, 33, 3),
		(7417, 'CAHAYA NEGERI', 723, 33, 3),
		(7418, 'CIBANON', 723, 33, 3),
		(7419, 'CIJUJUNG', 723, 33, 3),
		(7420, 'CIKEAS', 723, 33, 3),
		(7421, 'CILEBUT BARAT', 723, 33, 3),
		(7422, 'CILEBUT TIMUR', 723, 33, 3),
		(7423, 'CIMANDALA', 723, 33, 3),
		(7424, 'CISARUA', 723, 33, 3),
		(7425, 'GUNUNG GEULIS', 723, 33, 3),
		(7426, 'JANGGALA', 723, 33, 3),
		(7427, 'JENGGALU', 723, 33, 3),
		(7428, 'KAYU ARANG', 723, 33, 3),
		(7429, 'KUTI AGUNG', 723, 33, 3),
		(7430, 'LANGENSARI', 723, 33, 3),
		(7431, 'LEUWIBUDAH', 723, 33, 3),
		(7432, 'LIMBANGAN', 723, 33, 3),
		(7433, 'LINGGARAJA', 723, 33, 3),
		(7434, 'LUBUK SAHUNG', 723, 33, 3),
		(7435, 'MARGALAKSANA', 723, 33, 3),
		(7436, 'MARGALUYU', 723, 33, 3),
		(7437, 'MEKARJAYA', 723, 33, 3),
		(7438, 'NAGRAK', 723, 33, 3),
		(7439, 'NIUR', 723, 33, 3),
		(7440, 'PADANG KUAS', 723, 33, 3),
		(7441, 'PADANG PELAWI', 723, 33, 3),
		(7442, 'PASIR JAMBU', 723, 33, 3),
		(7443, 'PASIRHALANG', 723, 33, 3),
		(7444, 'PASIRLAJA', 723, 33, 3),
		(7445, 'RIAK SIABUN', 723, 33, 3),
		(7446, 'RIAK SIABUN I', 723, 33, 3),
		(7447, 'SARI MULYO', 723, 33, 3),
		(7448, 'SELAAWI', 723, 33, 3),
		(7449, 'SELAWANGI', 723, 33, 3),
		(7450, 'SIDO LUHUR', 723, 33, 3),
		(7451, 'SIDO SARI', 723, 33, 3),
		(7452, 'SIRNAJAYA', 723, 33, 3),
		(7453, 'SUKAMEKAR', 723, 33, 3),
		(7454, 'SUKAPURA', 723, 33, 3),
		(7455, 'SUKARAJA', 723, 33, 3),
		(7456, 'SUKARAJA', 723, 33, 3),
		(7457, 'SUKARAJA', 723, 33, 3),
		(7458, 'SUKATANI', 723, 33, 3),
		(7459, 'SUMBER ARUNG/ARUM', 723, 33, 3),
		(7460, 'SUMBER MAKMUR', 723, 33, 3),
		(7461, 'TARUNAJAYA', 723, 33, 3),
		(7462, 'KORPRI JAYA', 724, 33, 3),
		(7463, 'KORPRI RAYA', 724, 33, 3),
		(7464, 'PADASUKA', 724, 33, 3),
		(7465, 'SUKA MENAK', 724, 33, 3),
		(7466, 'SUKAKARSA', 724, 33, 3),
		(7467, 'SUKARAME', 724, 33, 3),
		(7468, 'SUKARAME', 724, 33, 3),
		(7469, 'SUKARAME BARU', 724, 33, 3),
		(7470, 'SUKARAPIH', 724, 33, 3),
		(7471, 'WARGAKERTA', 724, 33, 3),
		(7472, 'WAY DADI', 724, 33, 3),
		(7473, 'WAY DADI BARU', 724, 33, 3),
		(7474, 'GUNUNGSARI', 725, 33, 3),
		(7475, 'INDRAJAYA', 725, 33, 3),
		(7476, 'LINGGAJATI', 725, 33, 3),
		(7477, 'SINAGAR', 725, 33, 3),
		(7478, 'SUKAGALIH', 725, 33, 3),
		(7479, 'SUKAMAHI', 725, 33, 3),
		(7480, 'SUKARATU', 725, 33, 3),
		(7481, 'TAWANGBANTENG', 725, 33, 3),
		(7482, 'BANJARSARI', 726, 33, 3),
		(7483, 'CIPONDOK', 726, 33, 3),
		(7484, 'MARGAMULYA', 726, 33, 3),
		(7485, 'SUKAMENAK', 726, 33, 3),
		(7486, 'SUKAPANCAR', 726, 33, 3),
		(7487, 'SUKARATU', 726, 33, 3),
		(7488, 'SUKARESIK', 726, 33, 3),
		(7489, 'TANJUNGSARI', 726, 33, 3),
		(7490, 'MUGARSARI', 727, 33, 3),
		(7491, 'MULYASARI', 727, 33, 3),
		(7492, 'PASIREURIH', 727, 33, 3),
		(7493, 'SETIAMULYA', 727, 33, 3),
		(7494, 'SETIAWARGI', 727, 33, 3),
		(7495, 'SIRNAGALIH', 727, 33, 3),
		(7496, 'SUKAHURIP', 727, 33, 3),
		(7497, 'SUKAJADI', 727, 33, 3),
		(7498, 'SUKAJAYA', 727, 33, 3),
		(7499, 'SUKALUYU', 727, 33, 3),
		(7500, 'SUKAMANTRI', 727, 33, 3),
		(7501, 'SUKARESMI', 727, 33, 3),
		(7502, 'SUMELAP', 727, 33, 3),
		(7503, 'TAMANJAYA', 727, 33, 3),
		(7504, 'TAMANSARI', 727, 33, 3),
		(7505, 'TAMANSARI', 727, 33, 3),
		(7506, 'CIBALANARIK', 728, 33, 3),
		(7507, 'CIKEUSAL', 728, 33, 3),
		(7508, 'CILOLOHAN', 728, 33, 3),
		(7509, 'CINTAJAYA', 728, 33, 3),
		(7510, 'SUKANAGARA', 728, 33, 3),
		(7511, 'SUKASENANG', 728, 33, 3),
		(7512, 'TANJUNGJAYA', 728, 33, 3),
		(7513, 'BANYUASIH', 729, 33, 3),
		(7514, 'CIKUBANG', 729, 33, 3),
		(7515, 'DEUDEUL', 729, 33, 3),
		(7516, 'KERTARAHARJA', 729, 33, 3),
		(7517, 'PAGERALAM', 729, 33, 3),
		(7518, 'PURWARAHAYU', 729, 33, 3),
		(7519, 'RAKSASARI', 729, 33, 3),
		(7520, 'SINGASARI', 729, 33, 3),
		(7521, 'TARAJU', 729, 33, 3),
		(7522, 'CIKALANG', 730, 33, 3),
		(7523, 'EMPANGSARI', 730, 33, 3),
		(7524, 'KAHURIPAN', 730, 33, 3),
		(7525, 'LENGKONGSARI', 730, 33, 3),
		(7526, 'TAWANGSARI', 730, 33, 3),



		
		(7527, 'BANJARKULON', 731, 34, 4),
		(7528, 'BANJARMANGU', 731, 34, 4),
		(7529, 'BEJI', 731, 34, 4),
		(7530, 'GRIPIT', 731, 34, 4),
		(7531, 'JENGGAWUR', 731, 34, 4),
		(7532, 'KALILUNJAR', 731, 34, 4),
		(7533, 'KENDAGA', 731, 34, 4),
		(7534, 'KESENET', 731, 34, 4),
		(7535, 'MAJATENGAH', 731, 34, 4),
		(7536, 'PASEH', 731, 34, 4),
		(7537, 'PEKANDANGAN', 731, 34, 4),
		(7538, 'PRENDENGAN', 731, 34, 4),
		(7539, 'REJASARI', 731, 34, 4),
		(7540, 'SIGEBLOG', 731, 34, 4),
		(7541, 'SIJENGGUNG', 731, 34, 4),
		(7542, 'SIJERUK', 731, 34, 4),
		(7543, 'SIPEDANG', 731, 34, 4),
		(7544, 'AMPELSARI', 732, 34, 4),
		(7545, 'ARGASOKA', 732, 34, 4),
		(7546, 'CENDANA', 732, 34, 4),
		(7547, 'KARANGTENGAH', 732, 34, 4),
		(7548, 'KRANDEGAN', 732, 34, 4),
		(7549, 'KUTA BANJARNEGARA', 732, 34, 4),
		(7550, 'PARAKANCANGGAH', 732, 34, 4),
		(7551, 'SEMAMPIR', 732, 34, 4),
		(7552, 'SEMARANG', 732, 34, 4),
		(7553, 'SOKANANDI', 732, 34, 4),
		(7554, 'SOKAYASA', 732, 34, 4),
		(7555, 'TLAGAWERA', 732, 34, 4),
		(7556, 'WANGON', 732, 34, 4),
		(7557, 'BAKAL', 733, 34, 4),
		(7558, 'BATUR', 733, 34, 4),
		(7559, 'DIENG KULON', 733, 34, 4),
		(7560, 'KARANGTENGAH', 733, 34, 4),
		(7561, 'KEPAKISAN', 733, 34, 4),
		(7562, 'PASURENAN', 733, 34, 4),
		(7563, 'PEKASIRAN', 733, 34, 4),
		(7564, 'SUMBEREJO', 733, 34, 4),
		(7565, 'BANDINGAN', 734, 34, 4),
		(7566, 'BAWANG', 734, 34, 4),
		(7567, 'BAWANG', 734, 34, 4),
		(7568, 'BINORONG', 734, 34, 4),
		(7569, 'BLAMBANGAN', 734, 34, 4),
		(7570, 'CANDIGUGUR', 734, 34, 4),
		(7571, 'CANDIREJO', 734, 34, 4),
		(7572, 'DELES', 734, 34, 4),
		(7573, 'DEPOK', 734, 34, 4),
		(7574, 'GEMURUH', 734, 34, 4),
		(7575, 'GETAS', 734, 34, 4),
		(7576, 'GUNUNGSARI', 734, 34, 4),
		(7577, 'JAMBANGAN', 734, 34, 4),
		(7578, 'JLAMPRANG', 734, 34, 4),
		(7579, 'JOHO', 734, 34, 4),
		(7580, 'KALIREJO', 734, 34, 4),
		(7581, 'KEBATURAN', 734, 34, 4),
		(7582, 'KEBONDALEM', 734, 34, 4),
		(7583, 'KUTAYASA', 734, 34, 4),
		(7584, 'MAJALENGKA', 734, 34, 4),
		(7585, 'MANTRIANOM', 734, 34, 4),
		(7586, 'MASARAN', 734, 34, 4),
		(7587, 'PANGEMPON', 734, 34, 4),
		(7588, 'PASUSUKAN', 734, 34, 4),
		(7589, 'PRANTEN', 734, 34, 4),
		(7590, 'PUCANG', 734, 34, 4),
		(7591, 'PURBO', 734, 34, 4),
		(7592, 'SANGUBANYU', 734, 34, 4),
		(7593, 'SERANG', 734, 34, 4),
		(7594, 'SIBEBEK', 734, 34, 4),
		(7595, 'SIDOHARJO', 734, 34, 4),
		(7596, 'SOKA', 734, 34, 4),
		(7597, 'SURJO', 734, 34, 4),
		(7598, 'WANADRI', 734, 34, 4),
		(7599, 'WATUURIP', 734, 34, 4),
		(7600, 'WINONG', 734, 34, 4),
		(7601, 'WIRAMASTRA', 734, 34, 4),
		(7602, 'WONO SARI', 734, 34, 4),
		(7603, 'ASINAN', 735, 34, 4),
		(7604, 'BEDANA', 735, 34, 4),
		(7605, 'GUNUNGLANGIT', 735, 34, 4),
		(7606, 'KALIBENING', 735, 34, 4),
		(7607, 'KALIBOMBONG', 735, 34, 4),
		(7608, 'KALISAT KIDUL', 735, 34, 4),
		(7609, 'KARANG ANYAR', 735, 34, 4),
		(7610, 'KASINOMAN', 735, 34, 4),
		(7611, 'KERTASARI', 735, 34, 4),
		(7612, 'MAJATENGAH', 735, 34, 4),
		(7613, 'PLORENGAN', 735, 34, 4),
		(7614, 'SEMBAWA', 735, 34, 4),
		(7615, 'SIDAKANGEN', 735, 34, 4),
		(7616, 'SIKUMPUL', 735, 34, 4),
		(7617, 'SIRUKEM', 735, 34, 4),
		(7618, 'SIRUKUN', 735, 34, 4),
		(7619, 'AMBAL', 736, 34, 4),
		(7620, 'BINANGUN', 736, 34, 4),
		(7621, 'GUMELAR', 736, 34, 4),
		(7622, 'JLEGONG', 736, 34, 4),
		(7623, 'KARANGGONDANG', 736, 34, 4),
		(7624, 'KARANGKOBAR', 736, 34, 4),
		(7625, 'LEKSANA', 736, 34, 4),
		(7626, 'PAGERPELAH', 736, 34, 4),
		(7627, 'PASURUHAN', 736, 34, 4),
		(7628, 'PAWEDEN', 736, 34, 4),
		(7629, 'PURWODADI', 736, 34, 4),
		(7630, 'SAMPANG', 736, 34, 4),
		(7631, 'SLATRI', 736, 34, 4),
		(7632, 'BANTARWARU', 737, 34, 4),
		(7633, 'BLITAR', 737, 34, 4),
		(7634, 'CLAPAR', 737, 34, 4),
		(7635, 'DAWUHAN', 737, 34, 4),
		(7636, 'GUNUNGGIANA', 737, 34, 4),
		(7637, 'KALIURIP', 737, 34, 4),
		(7638, 'KARANGANYAR', 737, 34, 4),
		(7639, 'KENTENG', 737, 34, 4),
		(7640, 'KUTAYASA', 737, 34, 4),
		(7641, 'LIMBANGAN', 737, 34, 4),
		(7642, 'MADUKARA', 737, 34, 4),
		(7643, 'PAGELAK', 737, 34, 4),
		(7644, 'PAKELEN', 737, 34, 4),
		(7645, 'PEKAUMAN', 737, 34, 4),
		(7646, 'PENAWANGAN', 737, 34, 4),
		(7647, 'PETAMBAKAN', 737, 34, 4),
		(7648, 'RAKITAN', 737, 34, 4),
		(7649, 'REJASA', 737, 34, 4),
		(7650, 'SERED', 737, 34, 4),
		(7651, 'TALUNAMBA', 737, 34, 4),
		(7652, 'BANJENGAN', 738, 34, 4),
		(7653, 'BLIMBING', 738, 34, 4),
		(7654, 'CANDIWULAN', 738, 34, 4),
		(7655, 'GLEMPANG', 738, 34, 4),
		(7656, 'JALATUNDA', 738, 34, 4),
		(7657, 'KALIWUNGU', 738, 34, 4),
		(7658, 'KEBAKALAN', 738, 34, 4),
		(7659, 'KEBANARAN', 738, 34, 4),
		(7660, 'KERTAYASA', 738, 34, 4),
		(7661, 'MANDIRAJA KULON', 738, 34, 4),
		(7662, 'MANDIRAJA WETAN', 738, 34, 4),
		(7663, 'PANGGISARI', 738, 34, 4),
		(7664, 'PURWASABA', 738, 34, 4),
		(7665, 'SALAMERTA', 738, 34, 4),
		(7666, 'SIMBANG', 738, 34, 4),
		(7667, 'SOMAWANGI', 738, 34, 4),
		(7668, 'DUREN', 739, 34, 4),
		(7669, 'GENTANSARI', 739, 34, 4),
		(7670, 'GUNUNGJATI', 739, 34, 4),
		(7671, 'KEBUTUHDUWUR', 739, 34, 4),
		(7672, 'KEBUTUHJURANG', 739, 34, 4),
		(7673, 'LEBAKWANGI', 739, 34, 4),
		(7674, 'PAGEDONGAN', 739, 34, 4),
		(7675, 'PESANGKALAN', 739, 34, 4),
		(7676, 'TWELAGIRI', 739, 34, 4),
		(7677, 'ARIBAYA', 740, 34, 4),
		(7678, 'BABADAN', 740, 34, 4),
		(7679, 'GUMINGSIR', 740, 34, 4),
		(7680, 'KALITLAGA', 740, 34, 4),
		(7681, 'KARANGNANGKA', 740, 34, 4),
		(7682, 'KAREKAN', 740, 34, 4),
		(7683, 'KASMARAN', 740, 34, 4),
		(7684, 'KAYUARES', 740, 34, 4),
		(7685, 'LARANGAN', 740, 34, 4),
		(7686, 'MAJASARI', 740, 34, 4),
		(7687, 'METAWANA', 740, 34, 4),
		(7688, 'NAGASARI', 740, 34, 4),
		(7689, 'PAGENTAN', 740, 34, 4),
		(7690, 'PLUMBUNGAN', 740, 34, 4),
		(7691, 'SOKARAJA', 740, 34, 4),
		(7692, 'TEGALJERUK', 740, 34, 4),
		(7693, 'BEJI', 741, 34, 4),
		(7694, 'LAWEN', 741, 34, 4),
		(7695, 'PANDANARUM', 741, 34, 4),
		(7696, 'PASEGERAN', 741, 34, 4),
		(7697, 'PINGIT LOR', 741, 34, 4),
		(7698, 'PRINGAMBA', 741, 34, 4),
		(7699, 'SINDUAJI', 741, 34, 4),
		(7700, 'SIRONGGE', 741, 34, 4),
		(7701, 'BEJI', 742, 34, 4),
		(7702, 'BITING', 742, 34, 4),
		(7703, 'CONDONG CAMPUR', 742, 34, 4),
		(7704, 'DARMAYASA', 742, 34, 4),
		(7705, 'GEMBOL', 742, 34, 4),
		(7706, 'GIRITIRTA', 742, 34, 4),
		(7707, 'GROGOL', 742, 34, 4),
		(7708, 'KALILUNJAR', 742, 34, 4),
		(7709, 'KARANGSARI', 742, 34, 4),
		(7710, 'PANUSUPAN', 742, 34, 4),
		(7711, 'PEGUNDUNGAN', 742, 34, 4),
		(7712, 'PEJAWARAN', 742, 34, 4),
		(7713, 'RATAMBA', 742, 34, 4),
		(7714, 'SARWODADI', 742, 34, 4),
		(7715, 'SEMANGKUNG', 742, 34, 4),
		(7716, 'SIDENGOK', 742, 34, 4),
		(7717, 'TLAHAP', 742, 34, 4),
		(7718, 'BADAKARYA', 743, 34, 4),
		(7719, 'BONDOLHARJO', 743, 34, 4),
		(7720, 'DANAKERTA', 743, 34, 4),
		(7721, 'JEMBANGAN', 743, 34, 4),
		(7722, 'KARANGSARI', 743, 34, 4),
		(7723, 'KECEPIT', 743, 34, 4),
		(7724, 'KLAPA', 743, 34, 4),
		(7725, 'MLAYA', 743, 34, 4),
		(7726, 'PETUGURAN', 743, 34, 4),
		(7727, 'PUNGGELAN', 743, 34, 4),
		(7728, 'PURWASANA', 743, 34, 4),
		(7729, 'SAMBONG', 743, 34, 4),
		(7730, 'SAWANGAN', 743, 34, 4),
		(7731, 'SIDARATA', 743, 34, 4),
		(7732, 'TANJUNGTIRTA', 743, 34, 4),
		(7733, 'TLAGA', 743, 34, 4),
		(7734, 'TRIBUANA', 743, 34, 4),
		(7735, 'DANARAJA', 744, 34, 4),
		(7736, 'GUMIWANG', 744, 34, 4),
		(7737, 'KALIAJIR', 744, 34, 4),
		(7738, 'KALIPELUS', 744, 34, 4),
		(7739, 'KALITENGAH', 744, 34, 4),
		(7740, 'KARANGANYAR', 744, 34, 4),
		(7741, 'KUTAWULUH', 744, 34, 4),
		(7742, 'MERDEN', 744, 34, 4),
		(7743, 'MERTASARI', 744, 34, 4),
		(7744, 'PARAKAN', 744, 34, 4),
		(7745, 'PETIR', 744, 34, 4),
		(7746, 'PUCUNGBEDUG', 744, 34, 4),
		(7747, 'PURWONEGORO', 744, 34, 4),
		(7748, 'KALILANDAK', 745, 34, 4),
		(7749, 'KALIMANDI', 745, 34, 4),
		(7750, 'KALIWINASUH', 745, 34, 4),
		(7751, 'KECITRAN', 745, 34, 4),
		(7752, 'KLAMPOK', 745, 34, 4),
		(7753, 'PAGAK', 745, 34, 4),
		(7754, 'PURWOREJO', 745, 34, 4),
		(7755, 'SIRKANDI', 745, 34, 4),
		(7756, 'ADIPASIR', 746, 34, 4),
		(7757, 'BADAMITA', 746, 34, 4),
		(7758, 'BANDINGAN', 746, 34, 4),
		(7759, 'GELANG', 746, 34, 4),
		(7760, 'KINCANG', 746, 34, 4),
		(7761, 'LENGKONG', 746, 34, 4),
		(7762, 'LUWUNG', 746, 34, 4),
		(7763, 'PINGIT', 746, 34, 4),
		(7764, 'RAKIT', 746, 34, 4),
		(7765, 'SITUWANGI', 746, 34, 4),
		(7766, 'TANJUNGANOM', 746, 34, 4),
		(7767, 'BANDINGAN', 747, 34, 4),
		(7768, 'BOJANEGARA', 747, 34, 4),
		(7769, 'GEMBONGAN', 747, 34, 4),
		(7770, 'KALIBENDA', 747, 34, 4),
		(7771, 'KARANGMANGU', 747, 34, 4),
		(7772, 'KEMIRI', 747, 34, 4),
		(7773, 'PANAWAREN', 747, 34, 4),
		(7774, 'PRIGI', 747, 34, 4),
		(7775, 'PRINGAMBA', 747, 34, 4),
		(7776, 'RANDEGAN', 747, 34, 4),
		(7777, 'SAWAL', 747, 34, 4),
		(7778, 'SIGALUH', 747, 34, 4),
		(7779, 'SINGOMERTO (SINGAMERTA)', 747, 34, 4),
		(7780, 'TUNGGARA', 747, 34, 4),
		(7781, 'WANACIPTA', 747, 34, 4),
		(7782, 'BADRAN', 748, 34, 4),
		(7783, 'BAKALREJO', 748, 34, 4),
		(7784, 'BERTA', 748, 34, 4),
		(7785, 'BOJONG KULON', 748, 34, 4),
		(7786, 'BRENGKOK', 748, 34, 4),
		(7787, 'BUNDER', 748, 34, 4),
		(7788, 'DERIK', 748, 34, 4),
		(7789, 'DERMASARI', 748, 34, 4),
		(7790, 'GENTAN', 748, 34, 4),
		(7791, 'GINTUNG LOR', 748, 34, 4),
		(7792, 'GUMELEM KULON', 748, 34, 4),
		(7793, 'GUMELEM WETAN', 748, 34, 4),
		(7794, 'JATIANOM', 748, 34, 4),
		(7795, 'JATIPURA', 748, 34, 4),
		(7796, 'KARANGJATI', 748, 34, 4),
		(7797, 'KARANGSALAM', 748, 34, 4),
		(7798, 'KEDAWUNG', 748, 34, 4),
		(7799, 'KEDONGDONG', 748, 34, 4),
		(7800, 'KEJIWAN', 748, 34, 4),
		(7801, 'KEMETUL', 748, 34, 4),
		(7802, 'KEMRANGGON', 748, 34, 4),
		(7803, 'KENTENG', 748, 34, 4),
		(7804, 'KETAPANG', 748, 34, 4),
		(7805, 'KORIPAN', 748, 34, 4),
		(7806, 'LUWUNG KENCANA', 748, 34, 4),
		(7807, 'MUNCAR', 748, 34, 4),
		(7808, 'NGASINAN', 748, 34, 4),
		(7809, 'PAKIKIRAN', 748, 34, 4),
		(7810, 'PANERUSAN KULON', 748, 34, 4),
		(7811, 'PANERUSAN WETAN', 748, 34, 4),
		(7812, 'PIASA WETAN', 748, 34, 4),
		(7813, 'SIDOHARJO', 748, 34, 4),
		(7814, 'SUSUKAN', 748, 34, 4),
		(7815, 'SUSUKAN', 748, 34, 4),
		(7816, 'SUSUKAN', 748, 34, 4),
		(7817, 'TANGKIL', 748, 34, 4),
		(7818, 'TAWANG', 748, 34, 4),
		(7819, 'TIMPIK', 748, 34, 4),
		(7820, 'UJUNGGEBANG', 748, 34, 4),
		(7821, 'WIYONG', 748, 34, 4),
		(7822, 'GUMINGSIR', 749, 34, 4),
		(7823, 'KANDANGWANGI', 749, 34, 4),
		(7824, 'KARANGJAMBE', 749, 34, 4),
		(7825, 'KARANGKEMIRI', 749, 34, 4),
		(7826, 'KASALIB', 749, 34, 4),
		(7827, 'LEMAHJAYA', 749, 34, 4),
		(7828, 'LINGGASARI', 749, 34, 4),
		(7829, 'MEDAYU', 749, 34, 4),
		(7830, 'TAPEN', 749, 34, 4),
		(7831, 'WANADADI', 749, 34, 4),
		(7832, 'WANAKARSA', 749, 34, 4),
		(7833, 'BABAKAN', 750, 34, 4),
		(7834, 'BALUN', 750, 34, 4),
		(7835, 'BANTAR', 750, 34, 4),
		(7836, 'CIAWI', 750, 34, 4),
		(7837, 'CIBUNTU', 750, 34, 4),
		(7838, 'DAWUHAN', 750, 34, 4),
		(7839, 'JATILAWANG', 750, 34, 4),
		(7840, 'KARANGTENGAH', 750, 34, 4),
		(7841, 'KASIMPAR', 750, 34, 4),
		(7842, 'KUBANG', 750, 34, 4),
		(7843, 'LEGOKHUNI', 750, 34, 4),
		(7844, 'LEGOKSAYEM', 750, 34, 4),
		(7845, 'NAGROG', 750, 34, 4),
		(7846, 'NANGERANG', 750, 34, 4),
		(7847, 'PAGERGUNUNG', 750, 34, 4),
		(7848, 'PANDANSARI', 750, 34, 4),
		(7849, 'PENANGGUNGAN', 750, 34, 4),
		(7850, 'PESANTREN', 750, 34, 4),
		(7851, 'RAHARJA', 750, 34, 4),
		(7852, 'SAKAMBANG', 750, 34, 4),
		(7853, 'SIMPANG', 750, 34, 4),
		(7854, 'SUKADAMI', 750, 34, 4),
		(7855, 'SUMURUGUL', 750, 34, 4),
		(7856, 'SUSUKAN', 750, 34, 4),
		(7857, 'SUWIDAK', 750, 34, 4),
		(7858, 'TARINGGUL TENGAH', 750, 34, 4),
		(7859, 'TARINGGUL TONGGOH', 750, 34, 4),
		(7860, 'TEMPURAN', 750, 34, 4),
		(7861, 'WANARAJA', 750, 34, 4),
		(7862, 'WANASARI', 750, 34, 4),
		(7863, 'WANAYASA', 750, 34, 4),
		(7864, 'WANAYASA', 750, 34, 4),
		(7865, 'AJIBARANG KULON', 751, 35, 4),
		(7866, 'AJIBARANG WETAN', 751, 35, 4),
		(7867, 'BANJARSARI (BANJARSARI KIDUL)', 751, 35, 4),
		(7868, 'CIBERUNG', 751, 35, 4),
		(7869, 'DARMAKRADENAN', 751, 35, 4),
		(7870, 'JINGKANG', 751, 35, 4),
		(7871, 'KALIBENDA', 751, 35, 4),
		(7872, 'KARANGBAWANG', 751, 35, 4),
		(7873, 'KRACAK', 751, 35, 4),
		(7874, 'LESMANA', 751, 35, 4),
		(7875, 'PANCASAN', 751, 35, 4),
		(7876, 'PANCURENDANG', 751, 35, 4),
		(7877, 'PANDANSARI', 751, 35, 4),
		(7878, 'SAWANGAN', 751, 35, 4),
		(7879, 'TIPAR (TIPAR KIDUL)', 751, 35, 4),
		(7880, 'BANJAREJO', 752, 35, 4),
		(7881, 'BANYU URIP', 752, 35, 4),
		(7882, 'BANYUMAS', 752, 35, 4),
		(7883, 'BANYUWANGI', 752, 35, 4),
		(7884, 'BINANGUN', 752, 35, 4),
		(7885, 'DANARAJA', 752, 35, 4),
		(7886, 'DAWUHAN', 752, 35, 4),
		(7887, 'KALISUBE', 752, 35, 4),
		(7888, 'KARANGRAU', 752, 35, 4),
		(7889, 'KEDUNGGEDE', 752, 35, 4),
		(7890, 'KEDUNGUTER', 752, 35, 4),
		(7891, 'KEJAWAR', 752, 35, 4),
		(7892, 'MULYO REJO', 752, 35, 4),
		(7893, 'NUSA WUNGU', 752, 35, 4),
		(7894, 'PAPRINGAN', 752, 35, 4),
		(7895, 'PASINGGANGAN', 752, 35, 4),
		(7896, 'PEKUNDEN', 752, 35, 4),
		(7897, 'SINAR MULYO/MULYA', 752, 35, 4),
		(7898, 'SRI RAHAYU', 752, 35, 4),
		(7899, 'SRIWUNGU', 752, 35, 4),
		(7900, 'SUDAGARAN', 752, 35, 4),
		(7901, 'SUKAMULYA', 752, 35, 4),
		(7902, 'WAYA KRUI', 752, 35, 4),
		(7903, 'KARANGMANGU', 753, 35, 4),
		(7904, 'KARANGSALAM LOR', 753, 35, 4),
		(7905, 'KARANGTENGAH', 753, 35, 4),
		(7906, 'KEBUMEN', 753, 35, 4),
		(7907, 'KEMUTUG KIDUL', 753, 35, 4),
		(7908, 'KEMUTUG LOR', 753, 35, 4),
		(7909, 'KETENGER', 753, 35, 4),
		(7910, 'KUTASARI', 753, 35, 4),
		(7911, 'PAMIJEN', 753, 35, 4),
		(7912, 'PANDAK', 753, 35, 4),
		(7913, 'PURWOSARI', 753, 35, 4),
		(7914, 'REMPOAH', 753, 35, 4),
		(7915, 'BATUANTEN', 754, 35, 4),
		(7916, 'CIKIDANG', 754, 35, 4),
		(7917, 'CILONGOK', 754, 35, 4),
		(7918, 'CIPETE', 754, 35, 4),
		(7919, 'GUNUNGLURAH', 754, 35, 4),
		(7920, 'JATISABA', 754, 35, 4),
		(7921, 'KALISARI', 754, 35, 4),
		(7922, 'KARANGLO', 754, 35, 4),
		(7923, 'KARANGTENGAH', 754, 35, 4),
		(7924, 'KASEGERAN', 754, 35, 4),
		(7925, 'LANGGONGSARI', 754, 35, 4),
		(7926, 'PAGERAJI', 754, 35, 4),
		(7927, 'PANEMBANGAN', 754, 35, 4),
		(7928, 'PANUSUPAN', 754, 35, 4),
		(7929, 'PEJOGOL', 754, 35, 4),
		(7930, 'PERNASIDI', 754, 35, 4),
		(7931, 'RANCAMAYA', 754, 35, 4),
		(7932, 'SAMBIRATA', 754, 35, 4),
		(7933, 'SOKAWERA', 754, 35, 4),
		(7934, 'SUDIMARA', 754, 35, 4),
		(7935, 'CIHONJE', 755, 35, 4),
		(7936, 'CILANGKAP', 755, 35, 4),
		(7937, 'GANCANG', 755, 35, 4),
		(7938, 'GUMELAR', 755, 35, 4),
		(7939, 'KARANGKEMOJING', 755, 35, 4),
		(7940, 'KEDUNGURANG', 755, 35, 4),
		(7941, 'PANINGKABAN', 755, 35, 4),
		(7942, 'SAMUDRA', 755, 35, 4),
		(7943, 'SAMUDRA KULON', 755, 35, 4),
		(7944, 'TLAGA', 755, 35, 4),
		(7945, 'ADISARA', 756, 35, 4),
		(7946, 'BANTAR', 756, 35, 4),
		(7947, 'GENTAWANGI', 756, 35, 4),
		(7948, 'GUNUNG WETAN', 756, 35, 4),
		(7949, 'KARANGANYAR', 756, 35, 4),
		(7950, 'KARANGLEWAS', 756, 35, 4),
		(7951, 'KEDUNGWRINGIN', 756, 35, 4),
		(7952, 'MARGASANA', 756, 35, 4),
		(7953, 'PEKUNCEN', 756, 35, 4),
		(7954, 'TINGGARJAYA', 756, 35, 4),
		(7955, 'TUNJUNG (TANJUNG)', 756, 35, 4),
		(7956, 'KALIBAGOR', 757, 35, 4),
		(7957, 'KALICUPAK KIDUL (KALIPUCAK KIDUL)', 757, 35, 4),
		(7958, 'KALICUPAK LOR (KALIPUCAK LOR)', 757, 35, 4),
		(7959, 'KALIORI', 757, 35, 4),
		(7960, 'KALISOGRA WETAN', 757, 35, 4),
		(7961, 'KARANGDADAP', 757, 35, 4),
		(7962, 'PAJERUKAN (PEJERUKAN)', 757, 35, 4),
		(7963, 'PEKAJA', 757, 35, 4),
		(7964, 'PETIR', 757, 35, 4),
		(7965, 'SROWOT', 757, 35, 4),
		(7966, 'SURO', 757, 35, 4),
		(7967, 'WLAHAR WETAN', 757, 35, 4),
		(7968, 'BABAKAN', 758, 35, 4),
		(7969, 'JIPANG', 758, 35, 4),
		(7970, 'KARANGGUDE KULON', 758, 35, 4),
		(7971, 'KARANGKEMIRI', 758, 35, 4),
		(7972, 'KARANGLEWAS KIDUL', 758, 35, 4),
		(7973, 'KEDIRI', 758, 35, 4),
		(7974, 'PANGEBATAN (PENGEBATAN)', 758, 35, 4),
		(7975, 'PASIR KULON', 758, 35, 4),
		(7976, 'PASIR LOR', 758, 35, 4),
		(7977, 'PASIR WETAN', 758, 35, 4),
		(7978, 'SINGASARI', 758, 35, 4),
		(7979, 'SUNYALANGU', 758, 35, 4),
		(7980, 'TAMANSARI', 758, 35, 4),
		(7981, 'ADISANA', 759, 35, 4),
		(7982, 'BANGSA', 759, 35, 4),
		(7983, 'CINDAGA', 759, 35, 4),
		(7984, 'GAMBARSARI', 759, 35, 4),
		(7985, 'KALISALAK', 759, 35, 4),
		(7986, 'KALIWEDI', 759, 35, 4),
		(7987, 'KARANGSARI', 759, 35, 4),
		(7988, 'KEBASEN', 759, 35, 4),
		(7989, 'MANDIRANCAN', 759, 35, 4),
		(7990, 'RANDEGAN', 759, 35, 4),
		(7991, 'SAWANGAN', 759, 35, 4),
		(7992, 'TUMIYANG', 759, 35, 4),
		(7993, 'BASEH', 760, 35, 4),
		(7994, 'BEJI', 760, 35, 4),
		(7995, 'DAWUHAN KULON', 760, 35, 4),
		(7996, 'DAWUHAN WETAN', 760, 35, 4),
		(7997, 'DUKUHJATI WETAN', 760, 35, 4),
		(7998, 'KALIKESUR', 760, 35, 4),
		(7999, 'KALISALAK', 760, 35, 4),
		(8000, 'KARANG ANYAR', 760, 35, 4)
	`)
}
