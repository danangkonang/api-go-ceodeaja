package subDistrict

import (
	"fmt"

	"github.com/danangkonang/rest-api/config"
)

func Kec5() {
	db := config.Connect()
	_, err := db.Exec(`INSERT INTO sub_districts (id,subDistrict,regency_id,province_id) VALUES
		(4001, 'MIHING RAYA', 277, 19),
		(4002, 'MIRI MANASA', 277, 19),
		(4003, 'RUNGAN', 277, 19),
		(4004, 'RUNGAN BARAT', 277, 19),
		(4005, 'RUNGAN HULU', 277, 19),
		(4006, 'SEPANG (SEPANG SIMIN)', 277, 19),
		(4007, 'TEWAH', 277, 19),
		(4008, 'BASARANG', 278, 19),
		(4009, 'BATAGUH', 278, 19),
		(4010, 'DADAHUP', 278, 19),
		(4011, 'KAPUAS BARAT', 278, 19),
		(4012, 'KAPUAS HILIR', 278, 19),
		(4013, 'KAPUAS HULU', 278, 19),
		(4014, 'KAPUAS KUALA', 278, 19),
		(4015, 'KAPUAS MURUNG', 278, 19),
		(4016, 'KAPUAS TENGAH', 278, 19),
		(4017, 'KAPUAS TIMUR', 278, 19),
		(4018, 'MANDAU TALAWANG', 278, 19),
		(4019, 'MANTANGAI', 278, 19),
		(4020, 'PASAK TALAWANG', 278, 19),
		(4021, 'PULAU PETAK', 278, 19),
		(4022, 'TAMBAN CATUR', 278, 19),
		(4023, 'TIMPAH', 278, 19),
		(4024, 'BUKIT RAYA', 279, 19),
		(4025, 'KAMIPANG', 279, 19),
		(4026, 'KATINGAN HILIR', 279, 19),
		(4027, 'KATINGAN HULU', 279, 19),
		(4028, 'KATINGAN KUALA', 279, 19),
		(4029, 'KATINGAN TENGAH', 279, 19),
		(4030, 'MARIKIT', 279, 19),
		(4031, 'MENDAWAI', 279, 19),
		(4032, 'PETAK MALAI', 279, 19),
		(4033, 'PULAU MALAN', 279, 19),
		(4034, 'SANAMAN MANTIKEI (SENAMANG MANTIKEI)', 279, 19),
		(4035, 'TASIK PAYAWAN', 279, 19),
		(4036, 'TEWANG SANGGALANG GARING (SANGALANG)', 279, 19),
		(4037, 'ARUT SELATAN', 280, 19),
		(4038, 'ARUT UTARA', 280, 19),
		(4039, 'KOTAWARINGIN LAMA', 280, 19),
		(4040, 'KUMAI', 280, 19),
		(4041, 'PANGKALAN BANTENG', 280, 19),
		(4042, 'PANGKALAN LADA', 280, 19),
		(4043, 'ANTANG KALANG', 281, 19),
		(4044, 'BAAMANG', 281, 19),
		(4045, 'BUKIT SANTUEI', 281, 19),
		(4046, 'CEMPAGA', 281, 19),
		(4047, 'CEMPAGA HULU', 281, 19),
		(4048, 'KOTA BESI', 281, 19),
		(4049, 'MENTAWA BARU (KETAPANG)', 281, 19),
		(4050, 'MENTAYA HILIR SELATAN', 281, 19),
		(4051, 'MENTAYA HILIR UTARA', 281, 19),
		(4052, 'MENTAYA HULU', 281, 19),
		(4053, 'PARENGGEAN', 281, 19),
		(4054, 'PULAU HANAUT', 281, 19),
		(4055, 'SERANAU', 281, 19),
		(4056, 'TELAGA ANTANG', 281, 19),
		(4057, 'TELAWANG', 281, 19),
		(4058, 'TELUK SAMPIT', 281, 19),
		(4059, 'TUALAN HULU', 281, 19),
		(4060, 'BATANGKAWA', 282, 19),
		(4061, 'BELANTIKAN RAYA', 282, 19),
		(4062, 'BULIK', 282, 19),
		(4063, 'BULIK TIMUR', 282, 19),
		(4064, 'DELANG', 282, 19),
		(4065, 'LAMANDAU', 282, 19),
		(4066, 'MENTHOBI RAYA', 282, 19),
		(4067, 'SEMATU JAYA', 282, 19),
		(4068, 'BARITO TUHUP RAYA', 283, 19),
		(4069, 'LAUNG TUHUP', 283, 19),
		(4070, 'MURUNG', 283, 19),
		(4071, 'PERMATA INTAN', 283, 19),
		(4072, 'SERIBU RIAM', 283, 19),
		(4073, 'SUMBER BARITO', 283, 19),
		(4074, 'SUNGAI BABUAT', 283, 19),
		(4075, 'TANAH SIANG', 283, 19),
		(4076, 'TANAH SIANG SELATAN', 283, 19),
		(4077, 'UUT MURUNG', 283, 19),
		(4078, 'BUKIT BATU', 284, 19),
		(4079, 'JEKAN RAYA', 284, 19),
		(4080, 'PAHANDUT', 284, 19),
		(4081, 'RAKUMPIT', 284, 19),
		(4082, 'SEBANGAU', 284, 19),
		(4083, 'BANAMA TINGANG', 285, 19),
		(4084, 'JABIREN RAYA', 285, 19),
		(4085, 'KAHAYAN HILIR', 285, 19),
		(4086, 'KAHAYAN KUALA', 285, 19),
		(4087, 'KAHAYAN TENGAH', 285, 19),
		(4088, 'MALIKU', 285, 19),
		(4089, 'PANDIH BATU', 285, 19),
		(4090, 'SEBANGAU KUALA', 285, 19),
		(4091, 'DANAU SELULUK', 286, 19),
		(4092, 'DANAU SEMBULUH', 286, 19),
		(4093, 'HANAU', 286, 19),
		(4094, 'SERUYAN HILIR', 286, 19),
		(4095, 'SERUYAN HILIR TIMUR', 286, 19),
		(4096, 'SERUYAN HULU', 286, 19),
		(4097, 'SERUYAN RAYA', 286, 19),
		(4098, 'SERUYAN TENGAH', 286, 19),
		(4099, 'SULING TAMBUN', 286, 19),
		(4100, 'BALAI RIAM', 287, 19),
		(4101, 'JELAI', 287, 19),
		(4102, 'PANTAI LUNCI', 287, 19),
		(4103, 'PERMATA KECUBUNG', 287, 19),
		(4104, 'SUKAMARA', 287, 19),
		(4105, 'AWAYAN', 288, 20),
		(4106, 'BATU MANDI', 288, 20),
		(4107, 'HALONG', 288, 20),
		(4108, 'JUAI', 288, 20),
		(4109, 'LAMPIHONG', 288, 20),
		(4110, 'PARINGIN', 288, 20),
		(4111, 'PARINGIN SELATAN', 288, 20),
		(4112, 'TEBING TINGGI', 288, 20),
		(4113, 'BANJAR BARU SELATAN', 290, 20),
		(4114, 'BANJAR BARU UTARA', 290, 20),
		(4115, 'CEMPAKA', 290, 20),
		(4116, 'LANDASAN ULIN', 290, 20),
		(4117, 'LIANG ANGGANG', 290, 20),
		(4118, 'BANJARMASIN BARAT', 291, 20),
		(4119, 'BANJARMASIN SELATAN', 291, 20),
		(4120, 'BANJARMASIN TENGAH', 291, 20),
		(4121, 'BANJARMASIN TIMUR', 291, 20),
		(4122, 'BANJARMASIN UTARA', 291, 20),
		(4123, 'ALALAK', 292, 20),
		(4124, 'ANJIR MUARA', 292, 20),
		(4125, 'ANJIR PASAR', 292, 20),
		(4126, 'BAKUMPAI', 292, 20),
		(4127, 'BARAMBAI', 292, 20),
		(4128, 'BELAWANG', 292, 20),
		(4129, 'CERBON', 292, 20),
		(4130, 'JEJANGKIT', 292, 20),
		(4131, 'MANDASTANA', 292, 20),
		(4132, 'MARABAHAN', 292, 20),
		(4133, 'MEKAR SARI', 292, 20),
		(4134, 'RANTAU BADAUH', 292, 20),
		(4135, 'TABUKAN', 292, 20),
		(4136, 'TABUNGANEN', 292, 20),
		(4137, 'TAMBAN', 292, 20),
		(4138, 'WANARAYA', 292, 20),
		(4139, 'ANGKINANG', 293, 20),
		(4140, 'DAHA BARAT', 293, 20),
		(4141, 'DAHA SELATAN', 293, 20),
		(4142, 'DAHA UTARA', 293, 20),
		(4143, 'KALUMPANG (KELUMPANG)', 293, 20),
		(4144, 'KANDANGAN', 293, 20),
		(4145, 'LOKSADO', 293, 20),
		(4146, 'PADANG BATUNG', 293, 20),
		(4147, 'SIMPUR', 293, 20),
		(4148, 'TELAGA LANGSAT', 293, 20),
		(4149, 'BARABAI', 294, 20),
		(4150, 'BATANG ALAI SELATAN', 294, 20),
		(4151, 'BATANG ALAI TIMUR', 294, 20),
		(4152, 'BATANG ALAI UTARA', 294, 20),
		(4153, 'BATU BENAWA', 294, 20),
		(4154, 'HANTAKAN', 294, 20),
		(4155, 'HARUYAN', 294, 20),
		(4156, 'LABUAN AMAS SELATAN', 294, 20),
		(4157, 'LABUAN AMAS UTARA', 294, 20),
		(4158, 'LIMPASU', 294, 20),
		(4159, 'PANDAWAN', 294, 20),
		(4160, 'AMUNTAI SELATAN', 295, 20),
		(4161, 'AMUNTAI TENGAH', 295, 20),
		(4162, 'AMUNTAI UTARA', 295, 20),
		(4163, 'BABIRIK', 295, 20),
		(4164, 'BANJANG', 295, 20),
		(4165, 'DANAU PANGGANG', 295, 20),
		(4166, 'HAUR GADING', 295, 20),
		(4167, 'PAMINGGIR', 295, 20),
		(4168, 'SUNGAI PANDAN', 295, 20),
		(4169, 'SUNGAI TABUKAN', 295, 20),
		(4170, 'HAMPANG', 296, 20),
		(4171, 'KELUMPANG BARAT', 296, 20),
		(4172, 'KELUMPANG HILIR', 296, 20),
		(4173, 'KELUMPANG HULU', 296, 20),
		(4174, 'KELUMPANG SELATAN', 296, 20),
		(4175, 'KELUMPANG TENGAH', 296, 20),
		(4176, 'KELUMPANG UTARA', 296, 20),
		(4177, 'PAMUKAN BARAT', 296, 20),
		(4178, 'PAMUKAN SELATAN', 296, 20),
		(4179, 'PAMUKAN UTARA', 296, 20),
		(4180, 'PULAU LAUT BARAT', 296, 20),
		(4181, 'PULAU LAUT KEPULAUAN', 296, 20),
		(4182, 'PULAU LAUT SELATAN', 296, 20),
		(4183, 'PULAU LAUT TANJUNG SELAYAR', 296, 20),
		(4184, 'PULAU LAUT TENGAH', 296, 20),
		(4185, 'PULAU LAUT TIMUR', 296, 20),
		(4186, 'PULAU LAUT UTARA', 296, 20),
		(4187, 'PULAU SEBUKU', 296, 20),
		(4188, 'SAMPANAHAN', 296, 20),
		(4189, 'SUNGAI DURIAN', 296, 20),
		(4190, 'BANUA LAWAS', 297, 20),
		(4191, 'BINTANG ARA', 297, 20),
		(4192, 'HARUAI', 297, 20),
		(4193, 'JARO', 297, 20),
		(4194, 'KELUA (KLUA)', 297, 20),
		(4195, 'MUARA HARUS', 297, 20),
		(4196, 'MUARA UYA', 297, 20),
		(4197, 'MURUNG PUDAK', 297, 20),
		(4198, 'PUGAAN', 297, 20),
		(4199, 'TANTA', 297, 20),
		(4200, 'UPAU', 297, 20),
		(4201, 'ANGSANA', 298, 20),
		(4202, 'BATULICIN', 298, 20),
		(4203, 'KARANG BINTANG', 298, 20),
		(4204, 'KURANJI', 298, 20),
		(4205, 'KUSAN HILIR', 298, 20),
		(4206, 'KUSAN HULU', 298, 20),
		(4207, 'MANTEWE', 298, 20),
		(4208, 'SATUI', 298, 20),
		(4209, 'SIMPANG EMPAT', 298, 20),
		(4210, 'SUNGAI LOBAN', 298, 20),
		(4211, 'BAJUIN', 299, 20),
		(4212, 'BATI-BATI', 299, 20),
		(4213, 'BUMI MAKMUR', 299, 20),
		(4214, 'JORONG', 299, 20),
		(4215, 'KINTAP', 299, 20),
		(4216, 'KURAU', 299, 20),
		(4217, 'PANYIPATAN', 299, 20),
		(4218, 'PELAIHARI', 299, 20),
		(4219, 'TAKISUNG', 299, 20),
		(4220, 'TAMBANG ULANG', 299, 20),
		(4221, 'BAKARANGAN', 300, 20),
		(4222, 'BUNGUR', 300, 20),
		(4223, 'CANDI LARAS SELATAN', 300, 20),
		(4224, 'CANDI LARAS UTARA', 300, 20),
		(4225, 'HATUNGUN', 300, 20),
		(4226, 'LOKPAIKAT', 300, 20),
		(4227, 'PIANI', 300, 20),
		(4228, 'SALAM BABARIS', 300, 20),
		(4229, 'TAPIN SELATAN', 300, 20),
		(4230, 'TAPIN TENGAH', 300, 20),
		(4231, 'TAPIN UTARA', 300, 20),
		(4232, 'BALIKPAPAN BARAT', 301, 21),
		(4233, 'BALIKPAPAN KOTA', 301, 21),
		(4234, 'BALIKPAPAN SELATAN', 301, 21),
		(4235, 'BALIKPAPAN TENGAH', 301, 21),
		(4236, 'BALIKPAPAN TIMUR', 301, 21),
		(4237, 'BALIKPAPAN UTARA', 301, 21),
		(4238, 'BIATAN', 302, 21),
		(4239, 'BIDUK-BIDUK', 302, 21),
		(4240, 'DERAWAN (PULAU DERAWAN)', 302, 21),
		(4241, 'GUNUNG TABUR', 302, 21),
		(4242, 'KELAY', 302, 21),
		(4243, 'MARATUA', 302, 21),
		(4244, 'SAMBALIUNG', 302, 21),
		(4245, 'SEGAH', 302, 21),
		(4246, 'TABALAR', 302, 21),
		(4247, 'TALISAYAN', 302, 21),
		(4248, 'TANJUNG REDEB', 302, 21),
		(4249, 'TELUK BAYUR', 302, 21),
		(4250, 'BONTANG BARAT', 303, 21),
		(4251, 'BONTANG SELATAN', 303, 21),
		(4252, 'BONTANG UTARA', 303, 21),
		(4253, 'BARONG TONGKOK', 304, 21),
		(4254, 'BENTIAN BESAR', 304, 21),
		(4255, 'BONGAN', 304, 21),
		(4256, 'DAMAI', 304, 21),
		(4257, 'JEMPANG', 304, 21),
		(4258, 'LAHAM', 304, 21),
		(4259, 'LINGGANG BIGUNG', 304, 21),
		(4260, 'LONG APARI', 304, 21),
		(4261, 'LONG BAGUN', 304, 21),
		(4262, 'LONG HUBUNG', 304, 21),
		(4263, 'LONG IRAM', 304, 21),
		(4264, 'LONG PAHANGAI', 304, 21),
		(4265, 'MANOR BULATIN (MOOK MANAAR BULATN)', 304, 21),
		(4266, 'MELAK', 304, 21),
		(4267, 'MUARA LAWA', 304, 21),
		(4268, 'MUARA PAHU', 304, 21),
		(4269, 'NYUATAN', 304, 21),
		(4270, 'PENYINGGAHAN', 304, 21),
		(4271, 'SEKOLAQ DARAT', 304, 21),
		(4272, 'SILUQ NGURAI', 304, 21),
		(4273, 'TERING', 304, 21),
		(4274, 'ANGGANA', 305, 21),
		(4275, 'KEMBANG JANGGUT', 305, 21),
		(4276, 'KENOHAN', 305, 21),
		(4277, 'KOTA BANGUN', 305, 21),
		(4278, 'LOA JANAN', 305, 21),
		(4279, 'LOA KULU', 305, 21),
		(4280, 'MARANG KAYU', 305, 21),
		(4281, 'MUARA BADAK', 305, 21),
		(4282, 'MUARA JAWA', 305, 21),
		(4283, 'MUARA KAMAN', 305, 21),
		(4284, 'MUARA MUNTAI', 305, 21),
		(4285, 'MUARA WIS', 305, 21),
		(4286, 'SAMBOJA (SEMBOJA)', 305, 21),
		(4287, 'SANGA-SANGA', 305, 21),
		(4288, 'SEBULU', 305, 21),
		(4289, 'TENGGARONG', 305, 21),
		(4290, 'TENGGARONG SEBERANG', 305, 21),
		(4291, 'BENGALON', 306, 21),
		(4292, 'BUSANG', 306, 21),
		(4293, 'KALIORANG', 306, 21),
		(4294, 'KARANGAN', 306, 21),
		(4295, 'KAUBUN', 306, 21),
		(4296, 'KONGBENG', 306, 21),
		(4297, 'LONG MESANGAT (MESENGAT)', 306, 21),
		(4298, 'MUARA ANCALONG', 306, 21),
		(4299, 'MUARA BENGKAL', 306, 21),
		(4300, 'MUARA WAHAU', 306, 21),
		(4301, 'RANTAU PULUNG', 306, 21),
		(4302, 'SANDARAN', 306, 21),
		(4303, 'SANGATTA SELATAN', 306, 21),
		(4304, 'SANGATTA UTARA', 306, 21),
		(4305, 'SANGKULIRANG', 306, 21),
		(4306, 'TELEN', 306, 21),
		(4307, 'TELUK PANDAN', 306, 21),
		(4308, 'BATU ENGAU', 307, 21),
		(4309, 'BATU SOPANG', 307, 21),
		(4310, 'KUARO', 307, 21),
		(4311, 'LONG IKIS', 307, 21),
		(4312, 'LONG KALI', 307, 21),
		(4313, 'MUARA KOMAM', 307, 21),
		(4314, 'MUARA SAMU', 307, 21),
		(4315, 'PASIR BELENGKONG', 307, 21),
		(4316, 'TANAH GROGOT', 307, 21),
		(4317, 'TANJUNG HARAPAN', 307, 21),
		(4318, 'BABULU', 308, 21),
		(4319, 'PENAJAM', 308, 21),
		(4320, 'SEPAKU', 308, 21),
		(4321, 'WARU', 308, 21),
		(4322, 'LOA JANAN ILIR', 309, 21),
		(4323, 'PALARAN', 309, 21),
		(4324, 'SAMARINDA ILIR', 309, 21),
		(4325, 'SAMARINDA KOTA', 309, 21),
		(4326, 'SAMARINDA SEBERANG', 309, 21),
		(4327, 'SAMARINDA ULU', 309, 21),
		(4328, 'SAMARINDA UTARA', 309, 21),
		(4329, 'SAMBUTAN', 309, 21),
		(4330, 'SUNGAI KUNJANG', 309, 21),
		(4331, 'SUNGAI PINANG', 309, 21),
		(4332, 'PESO', 310, 22),
		(4333, 'PESO HILIR/ILIR', 310, 22),
		(4334, 'PULAU BUNYU', 310, 22),
		(4335, 'SEKATAK', 310, 22),
		(4336, 'TANJUNG PALAS', 310, 22),
		(4337, 'TANJUNG PALAS BARAT', 310, 22),
		(4338, 'TANJUNG PALAS TENGAH', 310, 22),
		(4339, 'TANJUNG PALAS TIMUR', 310, 22),
		(4340, 'TANJUNG PALAS UTARA', 310, 22),
		(4341, 'TANJUNG SELOR', 310, 22),
		(4342, 'BAHAU HULU', 311, 22),
		(4343, 'KAYAN SELATAN', 311, 22),
		(4344, 'MALINAU BARAT', 311, 22),
		(4345, 'MALINAU KOTA', 311, 22),
		(4346, 'MALINAU SELATAN', 311, 22),
		(4347, 'MALINAU SELATAN HILIR', 311, 22),
		(4348, 'MALINAU SELATAN HULU', 311, 22),
		(4349, 'MALINAU UTARA', 311, 22),
		(4350, 'MENTARANG', 311, 22),
		(4351, 'MENTARANG HULU', 311, 22),
		(4352, 'PUJUNGAN', 311, 22),
		(4353, 'SUNGAI BOH', 311, 22),
		(4354, 'SUNGAI TUBU', 311, 22),
		(4355, 'KRAYAN', 312, 22),
		(4356, 'KRAYAN SELATAN', 312, 22),
		(4357, 'LUMBIS', 312, 22),
		(4358, 'LUMBIS OGONG', 312, 22),
		(4359, 'NUNUKAN', 312, 22),
		(4360, 'NUNUKAN SELATAN', 312, 22),
		(4361, 'SEBATIK', 312, 22),
		(4362, 'SEBATIK BARAT', 312, 22),
		(4363, 'SEBATIK TENGAH', 312, 22),
		(4364, 'SEBATIK TIMUR', 312, 22),
		(4365, 'SEBATIK UTARA', 312, 22),
		(4366, 'SEBUKU', 312, 22),
		(4367, 'SEI MENGGARIS', 312, 22),
		(4368, 'SEMBAKUNG', 312, 22),
		(4369, 'TULIN ONSOI', 312, 22),
		(4370, 'BETAYAU', 313, 22),
		(4371, 'SESAYAP', 313, 22),
		(4372, 'SESAYAP HILIR', 313, 22),
		(4373, 'TANA LIA (TANAH LIA)', 313, 22),
		(4374, 'TARAKAN BARAT', 314, 22),
		(4375, 'TARAKAN TENGAH', 314, 22),
		(4376, 'TARAKAN TIMUR', 314, 22),
		(4377, 'TARAKAN UTARA', 314, 22),
		(4378, 'BANGGAE', 315, 23),
		(4379, 'BANGGAE TIMUR', 315, 23),
		(4380, 'MALUNDA', 315, 23),
		(4381, 'PAMBOANG', 315, 23),
		(4382, 'TAMMEREDO SENDANA', 315, 23),
		(4383, 'TUBO (TUBO SENDANA)', 315, 23),
		(4384, 'ULUMUNDA', 315, 23),
		(4385, 'ARALLE (ARRALE)', 316, 23),
		(4386, 'BALLA', 316, 23),
		(4387, 'BAMBANG', 316, 23),
		(4388, 'BUNTUMALANGKA', 316, 23),
		(4389, 'MAMASA', 316, 23),
		(4390, 'MAMBI', 316, 23),
		(4391, 'MEHALAAN', 316, 23),
		(4392, 'MESSAWA', 316, 23),
		(4393, 'NOSU', 316, 23),
		(4394, 'PANA', 316, 23),
		(4395, 'RANTEBULAHAN TIMUR', 316, 23),
		(4396, 'SESENA PADANG', 316, 23),
		(4397, 'SUMARORONG', 316, 23),
		(4398, 'TABANG', 316, 23),
		(4399, 'TABULAHAN', 316, 23),
		(4400, 'TANDUK KALUA', 316, 23),
		(4401, 'TAWALIAN', 316, 23),
		(4402, 'BONEHAU', 317, 23),
		(4403, 'BUDONG-BUDONG', 317, 23),
		(4404, 'KALUKKU', 317, 23),
		(4405, 'KALUMPANG', 317, 23),
		(4406, 'KAROSSA', 317, 23),
		(4407, 'KEP. BALA BALAKANG', 317, 23),
		(4408, 'MAMUJU', 317, 23),
		(4409, 'PANGALE', 317, 23),
		(4410, 'PAPALANG', 317, 23),
		(4411, 'SAMPAGA', 317, 23),
		(4412, 'SIMBORO DAN KEPULAUAN', 317, 23),
		(4413, 'TAPALANG', 317, 23),
		(4414, 'TAPALANG BARAT', 317, 23),
		(4415, 'TOBADAK', 317, 23),
		(4416, 'TOMMO', 317, 23),
		(4417, 'TOPOYO', 317, 23),
		(4418, 'BAMBAIRA', 318, 23),
		(4419, 'BAMBALAMOTU', 318, 23),
		(4420, 'BARAS', 318, 23),
		(4421, 'BULU TABA', 318, 23),
		(4422, 'DAPURANG', 318, 23),
		(4423, 'DURIPOKU', 318, 23),
		(4424, 'LARIANG', 318, 23),
		(4425, 'PASANGKAYU', 318, 23),
		(4426, 'PEDONGGA', 318, 23),
		(4427, 'SARJO', 318, 23),
		(4428, 'SARUDU', 318, 23),
		(4429, 'TIKKE RAYA', 318, 23),
		(4430, 'ALU (ALLU)', 319, 23),
		(4431, 'ANREAPI', 319, 23),
		(4432, 'BALANIPA', 319, 23),
		(4433, 'BINUANG', 319, 23),
		(4434, 'BULO', 319, 23),
		(4435, 'CAMPALAGIAN', 319, 23),
		(4436, 'LIMBORO', 319, 23),
		(4437, 'LUYO', 319, 23),
		(4438, 'MAPILLI', 319, 23),
		(4439, 'MATAKALI', 319, 23),
		(4440, 'MATANGNGA', 319, 23),
		(4441, 'POLEWALI', 319, 23),
		(4442, 'TAPANGO', 319, 23),
		(4443, 'TINAMBUNG', 319, 23),
		(4444, 'TUBBI TARAMANU (TUTAR/TUTALLU)', 319, 23),
		(4445, 'WONOMULYO', 319, 23),
		(4446, 'BANTAENG', 320, 24),
		(4447, 'BISSAPPU', 320, 24),
		(4448, 'EREMERASA', 320, 24),
		(4449, 'GANTARANG KEKE (GANTARENG KEKE)', 320, 24),
		(4450, 'PAJUKUKANG', 320, 24),
		(4451, 'SINOA', 320, 24),
		(4452, 'TOMPOBULU', 320, 24),
		(4453, 'ULUERE', 320, 24),
		(4454, 'BARRU', 321, 24),
		(4455, 'MALLUSETASI', 321, 24),
		(4456, 'PUJANANTING', 321, 24),
		(4457, 'SOPPENG RIAJA', 321, 24),
		(4458, 'TANETE RIAJA', 321, 24),
		(4459, 'TANETE RILAU', 321, 24),
		(4460, 'AJANGALE', 322, 24),
		(4461, 'AMALI', 322, 24),
		(4462, 'AWANGPONE', 322, 24),
		(4463, 'BAREBBO', 322, 24),
		(4464, 'BENGO', 322, 24),
		(4465, 'BONTOCANI', 322, 24),
		(4466, 'CENRANA', 322, 24),
		(4467, 'CINA', 322, 24),
		(4468, 'DUA BOCCOE', 322, 24),
		(4469, 'KAHU', 322, 24),
		(4470, 'KAJUARA', 322, 24),
		(4471, 'LAMURU', 322, 24),
		(4472, 'LAPPARIAJA', 322, 24),
		(4473, 'LIBURENG', 322, 24),
		(4474, 'PALAKKA', 322, 24),
		(4475, 'PATIMPENG', 322, 24),
		(4476, 'PONRE', 322, 24),
		(4477, 'SALOMEKKO', 322, 24),
		(4478, 'SIBULUE', 322, 24),
		(4479, 'TANETE RIATTANG', 322, 24),
		(4480, 'TANETE RIATTANG BARAT', 322, 24),
		(4481, 'TANETE RIATTANG TIMUR', 322, 24),
		(4482, 'TELLU LIMPOE', 322, 24),
		(4483, 'TELLU SIATTINGE', 322, 24),
		(4484, 'TONRA', 322, 24),
		(4485, 'ULAWENG', 322, 24),
		(4486, 'BONTO BAHARI', 323, 24),
		(4487, 'BONTOTIRO', 323, 24),
		(4488, 'BULUKUMBA (BULUKUMPA)', 323, 24),
		(4489, 'GANTORANG/GANTARANG (GANGKING)', 323, 24),
		(4490, 'HERO LANGE-LANGE (HERLANG)', 323, 24),
		(4491, 'KAJANG', 323, 24),
		(4492, 'KINDANG', 323, 24),
		(4493, 'RILAU ALE', 323, 24),
		(4494, 'UJUNG BULU', 323, 24),
		(4495, 'UJUNG LOE', 323, 24),
		(4496, 'ALLA', 324, 24),
		(4497, 'ANGGERAJA', 324, 24),
		(4498, 'BARAKA', 324, 24),
		(4499, 'BAROKO', 324, 24),
		(4500, 'BUNGIN', 324, 24),
		(4501, 'BUNTU BATU', 324, 24),
		(4502, 'CENDANA', 324, 24),
		(4503, 'CURIO', 324, 24),
		(4504, 'ENREKANG', 324, 24),
		(4505, 'MAIWA', 324, 24),
		(4506, 'MALUA', 324, 24),
		(4507, 'MASALLE', 324, 24),
		(4508, 'BAJENG', 325, 24),
		(4509, 'BAJENG BARAT', 325, 24),
		(4510, 'BAROMBONG', 325, 24),
		(4511, 'BIRINGBULU', 325, 24),
		(4512, 'BONTOLEMPANGANG', 325, 24),
		(4513, 'BONTOMARANNU', 325, 24),
		(4514, 'BONTONOMPO', 325, 24),
		(4515, 'BONTONOMPO SELATAN', 325, 24),
		(4516, 'BUNGAYA', 325, 24),
		(4517, 'MANUJU', 325, 24),
		(4518, 'PALLANGGA', 325, 24),
		(4519, 'PARANGLOE', 325, 24),
		(4520, 'PATTALLASSANG', 325, 24),
		(4521, 'SOMBA OPU (UPU)', 325, 24),
		(4522, 'TINGGIMONCONG', 325, 24),
		(4523, 'TOMBOLO PAO', 325, 24),
		(4524, 'ARUNGKEKE', 326, 24),
		(4525, 'BANGKALA', 326, 24),
		(4526, 'BANGKALA BARAT', 326, 24),
		(4527, 'BATANG', 326, 24),
		(4528, 'BINAMU', 326, 24),
		(4529, 'BONTORAMBA', 326, 24),
		(4530, 'KELARA', 326, 24),
		(4531, 'TAMALATEA', 326, 24),
		(4532, 'TAROWANG', 326, 24),
		(4533, 'TURATEA', 326, 24),
		(4534, 'BAJO', 327, 24),
		(4535, 'BAJO BARAT', 327, 24),
		(4536, 'BASSE SANGTEMPE UTARA', 327, 24),
		(4537, 'BASSESANG TEMPE (BASTEM)', 327, 24),
		(4538, 'BELOPA', 327, 24),
		(4539, 'BELOPA UTARA', 327, 24),
		(4540, 'BUA', 327, 24),
		(4541, 'BUA PONRANG (BUPON)', 327, 24),
		(4542, 'KAMANRE', 327, 24),
		(4543, 'LAMASI', 327, 24),
		(4544, 'LAMASI TIMUR', 327, 24),
		(4545, 'LAROMPONG', 327, 24),
		(4546, 'LAROMPONG SELATAN', 327, 24),
		(4547, 'LATIMOJONG', 327, 24),
		(4548, 'PONRANG', 327, 24),
		(4549, 'PONRANG SELATAN', 327, 24),
		(4550, 'SULI', 327, 24),
		(4551, 'SULI BARAT', 327, 24),
		(4552, 'WALENRANG', 327, 24),
		(4553, 'WALENRANG BARAT', 327, 24),
		(4554, 'WALENRANG TIMUR', 327, 24),
		(4555, 'WALENRANG UTARA', 327, 24),
		(4556, 'ANGKONA', 328, 24),
		(4557, 'BURAU', 328, 24),
		(4558, 'KALAENA', 328, 24),
		(4559, 'MALILI', 328, 24),
		(4560, 'MANGKUTANA', 328, 24),
		(4561, 'NUHA', 328, 24),
		(4562, 'TOMONI', 328, 24),
		(4563, 'TOMONI TIMUR', 328, 24),
		(4564, 'TOWUTI', 328, 24),
		(4565, 'WASUPONDA', 328, 24),
		(4566, 'WOTU', 328, 24),
		(4567, 'BAEBUNTA', 329, 24),
		(4568, 'BONE-BONE', 329, 24),
		(4569, 'LIMBONG', 329, 24),
		(4570, 'MALANGKE', 329, 24),
		(4571, 'MALANGKE BARAT', 329, 24),
		(4572, 'MAPPEDECENG', 329, 24),
		(4573, 'MASAMBA', 329, 24),
		(4574, 'RAMPI', 329, 24),
		(4575, 'SABBANG', 329, 24),
		(4576, 'SEKO', 329, 24),
		(4577, 'SUKAMAJU', 329, 24),
		(4578, 'TANA LILI', 329, 24),
		(4579, 'BIRING KANAYA', 330, 24),
		(4580, 'BONTOALA', 330, 24),
		(4581, 'MAKASSAR', 330, 24),
		(4582, 'MAMAJANG', 330, 24),
		(4583, 'MANGGALA', 330, 24),
		(4584, 'MARISO', 330, 24),
		(4585, 'PANAKKUKANG', 330, 24),
		(4586, 'RAPPOCINI', 330, 24),
		(4587, 'TALLO', 330, 24),
		(4588, 'TAMALANREA', 330, 24),
		(4589, 'TAMALATE', 330, 24),
		(4590, 'UJUNG PANDANG', 330, 24),
		(4591, 'UJUNG TANAH', 330, 24),
		(4592, 'WAJO', 330, 24),
		(4593, 'BANTIMURUNG', 331, 24),
		(4594, 'BONTOA (MAROS UTARA)', 331, 24),
		(4595, 'CAMBA', 331, 24),
		(4596, 'LAU', 331, 24),
		(4597, 'MALLAWA', 331, 24),
		(4598, 'MANDAI', 331, 24),
		(4599, 'MAROS BARU', 331, 24),
		(4600, 'MARUSU', 331, 24),
		(4601, 'MONCONGLOE', 331, 24),
		(4602, 'SIMBANG', 331, 24),
		(4603, 'TANRALILI', 331, 24),
		(4604, 'TOMPU BULU', 331, 24),
		(4605, 'TURIKALE', 331, 24),
		(4606, 'BARA', 332, 24),
		(4607, 'MUNGKAJANG', 332, 24),
		(4608, 'SENDANA', 332, 24),
		(4609, 'TELLUWANUA', 332, 24),
		(4610, 'WARA', 332, 24),
		(4611, 'WARA BARAT', 332, 24),
		(4612, 'WARA SELATAN', 332, 24),
		(4613, 'WARA TIMUR', 332, 24),
		(4614, 'WARA UTARA', 332, 24),
		(4615, 'BALOCCI', 333, 24),
		(4616, 'BUNGORO', 333, 24),
		(4617, 'LABAKKANG', 333, 24),
		(4618, 'LIUKANG KALMAS (KALUKUANG MASALIMA)', 333, 24),
		(4619, 'LIUKANG TANGAYA', 333, 24),
		(4620, 'LIUKANG TUPABBIRING', 333, 24),
		(4621, 'LIUKANG TUPABBIRING UTARA', 333, 24),
		(4622, 'MANDALLE', 333, 24),
		(4623, 'MARANG (MA RANG)', 333, 24),
		(4624, 'MINASA TENE', 333, 24),
		(4625, 'PANGKAJENE', 333, 24),
		(4626, 'SEGERI', 333, 24),
		(4627, 'TONDONG TALLASA', 333, 24),
		(4628, 'BACUKIKI', 334, 24),
		(4629, 'BACUKIKI BARAT', 334, 24),
		(4630, 'SOREANG', 334, 24),
		(4631, 'UJUNG', 334, 24),
		(4632, 'BATULAPPA', 335, 24),
		(4633, 'CEMPA', 335, 24),
		(4634, 'DUAMPANUA', 335, 24),
		(4635, 'LANRISANG', 335, 24),
		(4636, 'LEMBANG', 335, 24),
		(4637, 'MATTIRO BULU', 335, 24),
		(4638, 'MATTIRO SOMPE', 335, 24),
		(4639, 'PALETEANG', 335, 24),
		(4640, 'PATAMPANUA', 335, 24),
		(4641, 'SUPPA', 335, 24),
		(4642, 'TIROANG', 335, 24),
		(4643, 'WATANG SAWITTO', 335, 24),
		(4644, 'BENTENG', 336, 24),
		(4645, 'BONTOHARU', 336, 24),
		(4646, 'BONTOMANAI', 336, 24),
		(4647, 'BONTOMATENE', 336, 24),
		(4648, 'BONTOSIKUYU', 336, 24),
		(4649, 'BUKI', 336, 24),
		(4650, 'PASILAMBENA', 336, 24),
		(4651, 'PASIMARANNU', 336, 24),
		(4652, 'PASIMASSUNGGU', 336, 24),
		(4653, 'PASIMASUNGGU TIMUR', 336, 24),
		(4654, 'TAKABONERATE', 336, 24),
		(4655, 'BARANTI', 337, 24),
		(4656, 'DUA PITUE', 337, 24),
		(4657, 'KULO', 337, 24),
		(4658, 'MARITENGNGAE', 337, 24),
		(4659, 'PANCA LAUTAN (LAUTANG)', 337, 24),
		(4660, 'PANCA RIJANG', 337, 24),
		(4661, 'PITU RAISE/RIASE', 337, 24),
		(4662, 'PITU RIAWA', 337, 24),
		(4663, 'WATANG PULU', 337, 24),
		(4664, 'WATTANG SIDENRENG (WATANG SIDENRENG)', 337, 24),
		(4665, 'BULUPODDO', 338, 24),
		(4666, 'PULAU SEMBILAN', 338, 24),
		(4667, 'SINJAI BARAT', 338, 24),
		(4668, 'SINJAI BORONG', 338, 24),
		(4669, 'SINJAI SELATAN', 338, 24),
		(4670, 'SINJAI TENGAH', 338, 24),
		(4671, 'SINJAI TIMUR', 338, 24),
		(4672, 'SINJAI UTARA', 338, 24),
		(4673, 'CITTA', 339, 24),
		(4674, 'DONRI-DONRI', 339, 24),
		(4675, 'GANRA', 339, 24),
		(4676, 'LALABATA', 339, 24),
		(4677, 'LILI RILAU', 339, 24),
		(4678, 'LILIRAJA (LILI RIAJA)', 339, 24),
		(4679, 'MARIO RIAWA', 339, 24),
		(4680, 'MARIO RIWAWO', 339, 24),
		(4681, 'GALESONG', 340, 24),
		(4682, 'GALESONG SELATAN', 340, 24),
		(4683, 'GALESONG UTARA', 340, 24),
		(4684, 'MANGARA BOMBANG', 340, 24),
		(4685, 'MAPPAKASUNGGU', 340, 24),
		(4686, 'PATALLASSANG', 340, 24),
		(4687, 'POLOMBANGKENG SELATAN (POLOBANGKENG)', 340, 24),
		(4688, 'POLOMBANGKENG UTARA (POLOBANGKENG)', 340, 24),
		(4689, 'SANROBONE', 340, 24),
		(4690, 'BITTUANG', 341, 24),
		(4691, 'BONGGAKARADENG', 341, 24),
		(4692, 'GANDANG BATU SILLANAN', 341, 24),
		(4693, 'KURRA', 341, 24),
		(4694, 'MAKALE', 341, 24),
		(4695, 'MAKALE SELATAN', 341, 24),
		(4696, 'MAKALE UTARA', 341, 24),
		(4697, 'MALIMBONG BALEPE', 341, 24),
		(4698, 'MAPPAK', 341, 24),
		(4699, 'MASANDA', 341, 24),
		(4700, 'MENGKENDEK', 341, 24),
		(4701, 'RANO', 341, 24),
		(4702, 'RANTETAYO', 341, 24),
		(4703, 'REMBON', 341, 24),
		(4704, 'SALUPUTTI', 341, 24),
		(4705, 'SANGALLA (SANGGALA)', 341, 24),
		(4706, 'SANGALLA SELATAN', 341, 24),
		(4707, 'SANGALLA UTARA', 341, 24),
		(4708, 'SIMBUANG', 341, 24),
		(4709, 'AWAN RANTE KARUA', 342, 24),
		(4710, 'BALUSU', 342, 24),
		(4711, 'BANGKELEKILA', 342, 24),
		(4712, 'BARUPPU', 342, 24),
		(4713, 'BUNTAO', 342, 24),
		(4714, 'BUNTU PEPASAN', 342, 24),
		(4715, 'DENDE\\ PIONGAN NAPO', 342, 24),
		(4716, 'KAPALLA PITU (KAPALA PITU)', 342, 24),
		(4717, 'KESU', 342, 24),
		(4718, 'NANGGALA', 342, 24),
		(4719, 'RANTEBUA', 342, 24),
		(4720, 'RANTEPAO', 342, 24),
		(4721, 'RINDINGALLO', 342, 24),
		(4722, 'SA\\DAN', 342, 24),
		(4723, 'SANGGALANGI', 342, 24),
		(4724, 'SESEAN', 342, 24),
		(4725, 'SESEAN SULOARA', 342, 24),
		(4726, 'SOPAI', 342, 24),
		(4727, 'TALLUNGLIPU', 342, 24),
		(4728, 'TONDON', 342, 24),
		(4729, 'BELAWA', 343, 24),
		(4730, 'BOLA', 343, 24),
		(4731, 'GILIRENG', 343, 24),
		(4732, 'KEERA', 343, 24),
		(4733, 'MAJAULENG', 343, 24),
		(4734, 'MANIANG PAJO', 343, 24),
		(4735, 'PAMMANA', 343, 24),
		(4736, 'PENRANG', 343, 24),
		(4737, 'PITUMPANUA', 343, 24),
		(4738, 'SABBANG PARU', 343, 24),
		(4739, 'SAJOANGING', 343, 24),
		(4740, 'TAKKALALLA', 343, 24),
		(4741, 'TANA SITOLO', 343, 24),
		(4742, 'TEMPE', 343, 24),
		(4743, 'BATUPOARO', 344, 25),
		(4744, 'BETOAMBARI', 344, 25),
		(4745, 'BUNGI', 344, 25),
		(4746, 'KOKALUKUNA', 344, 25),
		(4747, 'LEA-LEA', 344, 25),
		(4748, 'MURHUM', 344, 25),
		(4749, 'SORA WALIO (SOROWALIO)', 344, 25),
		(4750, 'WOLIO', 344, 25),
		(4751, 'KABAENA', 345, 25),
		(4752, 'KABAENA BARAT', 345, 25),
		(4753, 'KABAENA SELATAN', 345, 25),
		(4754, 'KABAENA TENGAH', 345, 25),
		(4755, 'KABAENA TIMUR', 345, 25),
		(4756, 'KABAENA UTARA', 345, 25),
		(4757, 'KEPULAUAN MASALOKA RAYA', 345, 25),
		(4758, 'LENTARAI JAYA S. (LANTARI JAYA)', 345, 25),
		(4759, 'MATA OLEO', 345, 25),
		(4760, 'MATA USU', 345, 25),
		(4761, 'POLEANG', 345, 25),
		(4762, 'POLEANG BARAT', 345, 25),
		(4763, 'POLEANG SELATAN', 345, 25),
		(4764, 'POLEANG TENGAH', 345, 25),
		(4765, 'POLEANG TENGGARA', 345, 25),
		(4766, 'POLEANG TIMUR', 345, 25),
		(4767, 'POLEANG UTARA', 345, 25),
		(4768, 'RAROWATU', 345, 25),
		(4769, 'RAROWATU UTARA', 345, 25),
		(4770, 'RUMBIA', 345, 25),
		(4771, 'RUMBIA TENGAH', 345, 25),
		(4772, 'TONTONUNU (TONTONUWU)', 345, 25),
		(4773, 'BATAUGA', 346, 25),
		(4774, 'BATU ATAS', 346, 25),
		(4775, 'GU', 346, 25),
		(4776, 'KADATUA', 346, 25),
		(4777, 'KAPONTORI', 346, 25),
		(4778, 'LAKUDO', 346, 25),
		(4779, 'LAPANDEWA', 346, 25),
		(4780, 'LASALIMU', 346, 25),
		(4781, 'LASALIMU SELATAN', 346, 25),
		(4782, 'MAWASANGKA', 346, 25),
		(4783, 'MAWASANGKA TENGAH', 346, 25),
		(4784, 'MAWASANGKA TIMUR', 346, 25),
		(4785, 'PASAR WAJO', 346, 25),
		(4786, 'SAMPOLAWA', 346, 25),
		(4787, 'SANGIA MAMBULU', 346, 25),
		(4788, 'SIOMPU', 346, 25),
		(4789, 'SIOMPU BARAT', 346, 25),
		(4790, 'SIONTAPIA (SIONTAPINA)', 346, 25),
		(4791, 'TALAGA RAYA (TELAGA RAYA)', 346, 25),
		(4792, 'WABULA', 346, 25),
		(4793, 'WOLOWA', 346, 25),
		(4794, 'BONEGUNU', 347, 25),
		(4795, 'KAMBOWA', 347, 25),
		(4796, 'KULISUSU (KALINGSUSU/KALISUSU)', 347, 25),
		(4797, 'KULISUSU BARAT', 347, 25),
		(4798, 'KULISUSU UTARA', 347, 25),
		(4799, 'WAKORUMBA UTARA', 347, 25),
		(4800, 'ABELI', 348, 25),
		(4801, 'BARUGA', 348, 25),
		(4802, 'KADIA', 348, 25),
		(4803, 'KAMBU', 348, 25),
		(4804, 'KENDARI', 348, 25),
		(4805, 'KENDARI BARAT', 348, 25),
		(4806, 'MANDONGA', 348, 25),
		(4807, 'POASIA', 348, 25),
		(4808, 'PUUWATU', 348, 25),
		(4809, 'WUA-WUA', 348, 25),
		(4810, 'BAULA', 349, 25),
		(4811, 'KOLAKA', 349, 25),
		(4812, 'LADONGI', 349, 25),
		(4813, 'LALOLAE', 349, 25),
		(4814, 'LAMBANDIA (LAMBADIA)', 349, 25),
		(4815, 'LATAMBAGA', 349, 25),
		(4816, 'LOEA', 349, 25),
		(4817, 'MOWEWE', 349, 25),
		(4818, 'POLI POLIA', 349, 25),
		(4819, 'POLINGGONA', 349, 25),
		(4820, 'POMALAA', 349, 25),
		(4821, 'SAMATURU', 349, 25),
		(4822, 'TANGGETADA', 349, 25),
		(4823, 'TINONDO', 349, 25),
		(4824, 'TIRAWUTA', 349, 25),
		(4825, 'TOARI', 349, 25),
		(4826, 'ULUIWOI', 349, 25),
		(4827, 'WATUMBANGGA (WATUBANGGO)', 349, 25),
		(4828, 'WOLO', 349, 25),
		(4829, 'WUNDULAKO', 349, 25),
		(4830, 'BATU PUTIH', 350, 25),
		(4831, 'KATOI', 350, 25),
		(4832, 'KODEOHA', 350, 25),
		(4833, 'LASUSUA', 350, 25),
		(4834, 'LOMBAI (LAMBAI)', 350, 25),
		(4835, 'NGAPA', 350, 25),
		(4836, 'PAKUE', 350, 25),
		(4837, 'PAKUE TENGAH', 350, 25),
		(4838, 'PAKUE UTARA', 350, 25),
		(4839, 'POREHU', 350, 25),
		(4840, 'RANTEANGIN', 350, 25),
		(4841, 'TIWU', 350, 25),
		(4842, 'TOLALA', 350, 25),
		(4843, 'WATUNOHU', 350, 25),
		(4844, 'WAWO', 350, 25),
		(4845, 'ABUKI', 351, 25),
		(4846, 'AMONGGEDO', 351, 25),
		(4847, 'ANGGABERI', 351, 25),
		(4848, 'ASINUA', 351, 25),
		(4849, 'BESULUTU', 351, 25),
		(4850, 'BONDOALA', 351, 25),
		(4851, 'KAPOIALA (KAPOYALA)', 351, 25),
		(4852, 'KONAWE', 351, 25),
		(4853, 'LALONGGASUMEETO', 351, 25),
		(4854, 'LAMBUYA', 351, 25),
		(4855, 'LATOMA', 351, 25),
		(4856, 'MELUHU', 351, 25),
		(4857, 'ONEMBUTE', 351, 25),
		(4858, 'PONDIDAHA', 351, 25),
		(4859, 'PURIALA', 351, 25),
		(4860, 'ROUTA', 351, 25),
		(4861, 'SAMPARA', 351, 25),
		(4862, 'SOROPIA', 351, 25),
		(4863, 'TONGAUNA', 351, 25),
		(4864, 'UEPAI (UWEPAI)', 351, 25),
		(4865, 'UNAAHA', 351, 25),
		(4866, 'WAWONII BARAT', 351, 25),
		(4867, 'WAWONII SELATAN', 351, 25),
		(4868, 'WAWONII TENGAH', 351, 25),
		(4869, 'WAWONII TENGGARA', 351, 25),
		(4870, 'WAWONII TIMUR', 351, 25),
		(4871, 'WAWONII TIMUR LAUT', 351, 25),
		(4872, 'WAWONII UTARA', 351, 25),
		(4873, 'WAWOTOBI', 351, 25),
		(4874, 'WONGGEDUKU', 351, 25),
		(4875, 'ANDOOLO', 352, 25),
		(4876, 'ANGATA', 352, 25),
		(4877, 'BAITO', 352, 25),
		(4878, 'BASALA', 352, 25),
		(4879, 'BENUA', 352, 25),
		(4880, 'BUKE', 352, 25),
		(4881, 'KOLONO', 352, 25),
		(4882, 'LAEYA', 352, 25),
		(4883, 'LAINEA', 352, 25),
		(4884, 'LALEMBUU / LALUMBUU', 352, 25),
		(4885, 'LANDONO', 352, 25),
		(4886, 'LAONTI', 352, 25),
		(4887, 'MORAMO', 352, 25),
		(4888, 'MORAMO UTARA', 352, 25),
		(4889, 'MOWILA', 352, 25),
		(4890, 'PALANGGA', 352, 25),
		(4891, 'PALANGGA SELATAN', 352, 25),
		(4892, 'RANOMEETO', 352, 25),
		(4893, 'RANOMEETO BARAT', 352, 25),
		(4894, 'TINANGGEA', 352, 25),
		(4895, 'WOLASI', 352, 25),
		(4896, 'ANDOWIA', 353, 25),
		(4897, 'ASERA', 353, 25),
		(4898, 'LANGGIKIMA', 353, 25),
		(4899, 'LASOLO', 353, 25),
		(4900, 'MOLAWE', 353, 25),
		(4901, 'MOTUI', 353, 25),
		(4902, 'OHEO', 353, 25),
		(4903, 'SAWA', 353, 25),
		(4904, 'WIWIRANO', 353, 25),
		(4905, 'BARANGKA', 354, 25),
		(4906, 'BATALAIWARU (BATALAIWORU)', 354, 25),
		(4907, 'BATUKARA', 354, 25),
		(4908, 'BONE (BONE TONDO)', 354, 25),
		(4909, 'DURUKA', 354, 25),
		(4910, 'KABANGKA', 354, 25),
		(4911, 'KABAWO', 354, 25),
		(4912, 'KATOBU', 354, 25),
		(4913, 'KONTU KOWUNA', 354, 25),
		(4914, 'KONTUNAGA', 354, 25),
		(4915, 'KUSAMBI', 354, 25),
		(4916, 'LASALEPA', 354, 25),
		(4917, 'LAWA', 354, 25),
		(4918, 'LOHIA', 354, 25),
		(4919, 'MAGINTI', 354, 25),
		(4920, 'MALIGANO', 354, 25),
		(4921, 'MAROBO', 354, 25),
		(4922, 'NAPABALANO', 354, 25),
		(4923, 'NAPANO KUSAMBI', 354, 25),
		(4924, 'PASI KOLAGA', 354, 25),
		(4925, 'SAWERIGADI (SAWERIGADING/SEWERGADI)', 354, 25),
		(4926, 'TIWORO KEPULAUAN', 354, 25),
		(4927, 'TIWORO SELATAN', 354, 25),
		(4928, 'TIWORO TENGAH', 354, 25),
		(4929, 'TIWORO UTARA', 354, 25),
		(4930, 'TONGKUNO', 354, 25),
		(4931, 'TONGKUNO SELATAN', 354, 25),
		(4932, 'TOWEA', 354, 25),
		(4933, 'WA DAGA', 354, 25),
		(4934, 'WAKORUMBA SELATAN', 354, 25),
		(4935, 'WATOPUTE', 354, 25),
		(4936, 'BINONGKO', 355, 25),
		(4937, 'KALEDUPA', 355, 25),
		(4938, 'KALEDUPA SELATAN', 355, 25),
		(4939, 'TOGO BINONGKO', 355, 25),
		(4940, 'TOMIA', 355, 25),
		(4941, 'TOMIA TIMUR', 355, 25),
		(4942, 'WANGI-WANGI', 355, 25),
		(4943, 'WANGI-WANGI SELATAN', 355, 25),
		(4944, 'BALANTAK', 356, 26),
		(4945, 'BALANTAK SELATAN', 356, 26),
		(4946, 'BALANTAK UTARA', 356, 26),
		(4947, 'BATUI', 356, 26),
		(4948, 'BATUI SELATAN', 356, 26),
		(4949, 'BUALEMO (BOALEMO)', 356, 26),
		(4950, 'BUNTA', 356, 26),
		(4951, 'KINTOM', 356, 26),
		(4952, 'LAMALA', 356, 26),
		(4953, 'LOBU', 356, 26),
		(4954, 'LUWUK', 356, 26),
		(4955, 'LUWUK SELATAN', 356, 26),
		(4956, 'LUWUK TIMUR', 356, 26),
		(4957, 'LUWUK UTARA', 356, 26),
		(4958, 'MANTOH', 356, 26),
		(4959, 'MASAMA', 356, 26),
		(4960, 'MOILONG', 356, 26),
		(4961, 'NAMBO', 356, 26),
		(4962, 'NUHON', 356, 26),
		(4963, 'PAGIMANA', 356, 26),
		(4964, 'SIMPANG RAYA', 356, 26),
		(4965, 'TOILI', 356, 26),
		(4966, 'TOILI BARAT', 356, 26),
		(4967, 'BANGGAI', 357, 26),
		(4968, 'BANGGAI SELATAN', 357, 26),
		(4969, 'BANGGAI TENGAH', 357, 26),
		(4970, 'BANGGAI UTARA', 357, 26),
		(4971, 'BANGKURUNG', 357, 26),
		(4972, 'BOKAN KEPULAUAN', 357, 26),
		(4973, 'BUKO', 357, 26),
		(4974, 'BUKO SELATAN', 357, 26),
		(4975, 'BULAGI', 357, 26),
		(4976, 'BULAGI SELATAN', 357, 26),
		(4977, 'BULAGI UTARA', 357, 26),
		(4978, 'LABOBO (LOBANGKURUNG)', 357, 26),
		(4979, 'LIANG', 357, 26),
		(4980, 'PELING TENGAH', 357, 26),
		(4981, 'TINANGKUNG', 357, 26),
		(4982, 'TINANGKUNG SELATAN', 357, 26),
		(4983, 'TINANGKUNG UTARA', 357, 26),
		(4984, 'TOTIKUM (TOTIKUNG)', 357, 26),
		(4985, 'TOTIKUM SELATAN', 357, 26),
		(4986, 'BOKAT', 358, 26),
		(4987, 'BUKAL', 358, 26),
		(4988, 'BUNOBOGU', 358, 26),
		(4989, 'GADUNG', 358, 26),
		(4990, 'KARAMAT', 358, 26),
		(4991, 'LAKEA (LIPUNOTO)', 358, 26),
		(4992, 'MOMUNU', 358, 26),
		(4993, 'PALELEH', 358, 26),
		(4994, 'PALELEH BARAT', 358, 26),
		(4995, 'TILOAN', 358, 26),
		(4996, 'BALAESANG', 359, 26),
		(4997, 'BALAESANG TANJUNG', 359, 26),
		(4998, 'BANAWA', 359, 26),
		(4999, 'BANAWA SELATAN', 359, 26),
		(5000, 'BANAWA TENGAH', 359, 26)
	`)
	if err != nil {
		panic(err)
	}
	fmt.Println("insert kecamatan 5")
}