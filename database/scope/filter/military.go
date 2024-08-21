package filter

import (
	"gorm.io/gorm"
)

func Military(db *gorm.DB) *gorm.DB {
	return db.Where(
		"model IN ?", []string{
			"E3CF", /* BOEING E-3 Sentry */
			"E3TF", /* BOEING E-3 Sentry */
			"E6",   /* BOEING E-6 Mercury */
			"C5",   /* LOCKHEED C-5 Galaxy */
			"C17",  /* BOEING C-17 Globemaster 3 */
			"C130", /* LOCKHEED C130 Hercules */
			"C135", /* BOEING C-135 Stratolifter */
			"C30J", /* LOCKHEED MARTIN C-130J Super Hercules */
			"A400", /* AIRBUS A-400M Atlas */
			"K35R", /* BOEING KC-135 Stratotanker */
			"K35E", /* BOEING KC-135E Stratotanker */
			"R135", /* BOEING RC-135 */
			"C27J", /* ALENIA C-27J Spartan */
			"C160", /* TRANSALL C-160 */
			"RFAL", /* DASSAULT	Rafale */
		},
	).Or(
		"callsign LIKE ?", "ASF%", /* Österreichische Luftstreitkräfte (Austrian Air Force) */
	).Or(
		"callsign LIKE ?", "FAG%", /* Fuerza Aérea Argentina (Argentine Air Force) */
	).Or(
		"callsign LIKE ?", "BFW%", /* قوة دفاع البحرين (Bahrain Defence Force) */
	).Or(
		"callsign LIKE ?", "BAF%", /* Force aérienne belge (Belgian Air Force) */
	).Or(
		"callsign LIKE ?", "AYB%", /* Force terrestre belge (Belgian Land Component) */
	).Or(
		"callsign LIKE ?", "NYB%", /* Force marine belge (Belgian Navy) */
	).Or(
		"callsign LIKE ?", "BRS%", /* Força Aérea Brasileira (Brazilian Air Force) */
	).Or(
		"callsign LIKE ?", "HRZ%", /* Hrvatsko ratno zrakoplovstvo (Croatian Air Force) */
	).Or(
		"callsign LIKE ?", "CEF%", /* Vzdušné síly (Czech Air Force) */
	).Or(
		"callsign LIKE ?", "CFC%", /* Royal Canadian Air Force */
	).Or(
		"callsign LIKE ?", "DAF%", /* Flyvevåbnet (Royal Danish Air Force) */
	).Or(
		"callsign LIKE ?", "EGY%", /* القوات الجوية المصرية (Egyptian Air Force) */
	).Or(
		"callsign LIKE ?", "EEF%", /* Õhuvägi (Estonian Air Force) */
	).Or(
		"callsign LIKE ?", "FNF%", /* Ilmavoimat (Finnish Air Force) */
	).Or(
		"callsign LIKE ?", "FRU%", /* Sécurité Civil (French Rescue) */
	).Or(
		"callsign LIKE ?", "FAF%", /* Armée de l"air et de l"espace (French Air and Space Force) */
	).Or(
		"callsign LIKE ?", "FMY%", /* Aviation Légère De L"Armée De Terre (French Army Light Aviation) */
	).Or(
		"callsign LIKE ?", "FNY%", /* Marine Nationale (France Navy) */
	).Or(
		"callsign LIKE ?", "CTM%", /* Armée de l"Air (French Air Force) */
	).Or(
		"callsign LIKE ?", "DCN%", /* Bundeswehr (Federal Armed Forces) */
	).Or(
		"callsign LIKE ?", "GAF%", /* Luftwaffe (German Air Force) */
	).Or(
		"callsign LIKE ?", "GAM%", /* Heer (German Army) */
	).Or(
		"callsign LIKE ?", "GNY%", /* Deutsche Marine (German Navy) */
	).Or(
		"callsign LIKE ?", "HAF%", /* Πολεμική Αεροπορία (Hellenic Air Force) */
	).Or(
		"callsign LIKE ?", "HNA%", /* Ελληνικό Πολεμικό Ναυτικό (Hellenic Navy) */
	).Or(
		"callsign LIKE ?", "HUF%", /* Magyar Légierő (Hungarian Air Force) */
	).Or(
		"callsign LIKE ?", "IFC%", /* Indian Air Force */
	).Or(
		"callsign LIKE ?", "IAF%", /* זְרוֹעַ הָאֲוִיר וְהֶחָלָל (Israeli Air Force) */
	).Or(
		"callsign LIKE ?", "LAF%", /* Latvijas Gaisa spēki (Latvian Air Force) */
	).Or(
		"callsign LIKE ?", "LYF%", /* Lietuvos karinės oro pajėgos (Lithuanian Air Force) */
	).Or(
		"callsign LIKE ?", "NDF%", /* Namibian Defence Force */
	).Or(
		"callsign LIKE ?", "KRC%", /* Royal New Zealand Air Force */
	).Or(
		"callsign LIKE ?", "KIW%", /* Royal New Zealand Air Force */
	).Or(
		"callsign LIKE ?", "NGR%", /* Nigerian Air Force */
	).Or(
		"callsign LIKE ?", "FPR%", /* Fuerza Aérea del Perú (Peruvian Air Force) */
	).Or(
		"callsign LIKE ?", "PLF%", /* Siły Powietrzne (Polish Air Force) */
	).Or(
		"callsign LIKE ?", "PNY%", /* Marynarka Wojenna (Polish Navy) */
	).Or(
		"callsign LIKE ?", "AFP%", /* Força Aérea Portuguesa (Portuguese Air Force) */
	).Or(
		"callsign LIKE ?", "POA%", /* Exército Português (Portuguese Army) */
	).Or(
		"callsign LIKE ?", "PON%", /* Marinha Portuguesa (Portuguese Navy) */
	).Or(
		"callsign LIKE ?", "ROF%", /* Forțele Aeriene Române (Romanian Air Force) */
	).Or(
		"callsign LIKE ?", "RFR%", /* Royal Air Force */
	).Or(
		"callsign LIKE ?", "RRR%", /* Royal Air Force */
	).Or(
		"callsign LIKE ?", "RRF%", /* Royal Air Force */
	).Or(
		"callsign LIKE ?", "SHF%", /* Royal Air Force */
	).Or(
		"callsign LIKE ?", "NVY%", /* Royal Navy */
	).Or(
		"callsign LIKE ?", "MJN%", /* سلاح الجو السلطاني عمان (Royal Air Force of Oman) */
	).Or(
		"callsign LIKE ?", "ASY%", /* Royal Australian Air Force */
	).Or(
		"callsign LIKE ?", "RJZ%", /* سلاح الجو الملكي الأردني (Royal Jordanian Air Force) */
	).Or(
		"callsign LIKE ?", "RMF%", /* Tentera Udara Diraja Malaysia (Royal Malaysian Air Force) */
	).Or(
		"callsign LIKE ?", "NAF%", /* Koninklijke Luchtmacht (Royal Netherlands Air Force) */
	).Or(
		"callsign LIKE ?", "NRN%", /* Koninklijke Marine (Royal Netherlands Navy) */
	).Or(
		"callsign LIKE ?", "NOW%", /* Luftforsvaret (Royal Norwegian Air Force) */
	).Or(
		"callsign LIKE ?", "RSF%", /* الْقُوَّاتُ الْجَوِّيَّةُ الْمَلَكِيَّةْ ٱلسُّعُوْدِيَّة (Royal Saudi Air Force) */
	).Or(
		"callsign LIKE ?", "RFF%", /* Военно-воздушные силы России (Russian Air Force) */
	).Or(
		"callsign LIKE ?", "AME%", /* Fuerza Aerea Española (Spanish Air Force) */
	).Or(
		"callsign LIKE ?", "SQF%", /* Vzdušné sily Ozbrojených síl Slovenskej republiky (Slovak Air Force) */
	).Or(
		"callsign LIKE ?", "SIV%", /* Slovenska vojska (Slovenian Armed Forces) */
	).Or(
		"callsign LIKE ?", "SUI%", /* Schweizer Luftwaffe (Swiss Air Force) */
	).Or(
		"callsign LIKE ?", "LMG%", /* South African Air Force */
	).Or(
		"callsign LIKE ?", "SAF%", /* Angkatan Udara Republik Singapura (Republic of Singapore Air Force) */
	).Or(
		"callsign LIKE ?", "HVK%", /* Türk Hava Kuvvetleri (Turkish Air Force) */
	).Or(
		"callsign LIKE ?", "UAF%", /* القوات الجوية والدفاع الجوي الاماراتي (United Arab Emirates Air Force) */
	).Or(
		"callsign LIKE ?", "AIO%", /* United States Air Force */
	).Or(
		"callsign LIKE ?", "RCH%", /* Air Mobility Command (United States Air Force) */
	).Or(
		"callsign LIKE ?", "IAM%", /* Aeronautica Militare (Italian Air Force) */
	).Or(
		"callsign LIKE ?", "IEI%", /* Esercito Italiano (Italian Army) */
	).Or(
		"callsign LIKE ?", "MMI%", /* Marina Militare (Italian Navy) */
	).Or(
		"callsign LIKE ?", "IRL%", /* Irish Air Corps */
	).Or(
		"callsign LIKE ?", "SVF%", /* Svenska flygvapnet (Sweden Air Force) */
	).Or(
		"callsign LIKE ?", "MMF%", /* Multinational Multi-Role Tanker Transport Fleet (NATO) */
	).Or(
		"callsign LIKE ?", "KJD%", /* القُوَّاتُ الجَوِّيَّةُ الجَزَائِرِيَّةُ (Algerian Air Force) */
	).Or(
		"callsign LIKE ?", "AFB%", /* Военновъздушни сили (Bulgarian Air Force) */
	).Or(
		"callsign LIKE ?", "SRB%", /* Ратно ваздухопловство и противваздухопловна (Serbian Air Force and Air Defence) */
	).Or(
		"callsign LIKE ?", "TUN%", /* جيش الطيران (Tunisian Air Force) */
	).Or(
		"callsign LIKE ?", "LBF%", /* القوات الجوية الليبية (Libyan Air Force) */
	)
}
