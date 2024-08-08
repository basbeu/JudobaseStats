package judobase

type Contest struct {
	IDCompetition         *string `json:"id_competition"`
	IDFight               *string `json:"id_fight"`
	IDPersonBlue          *string `json:"id_person_blue"`
	IDPersonWhite         *string `json:"id_person_white"`
	IDWinner              *string `json:"id_winner"`
	IsFinished            *string `json:"is_finished"`
	Round                 *string `json:"round"`
	Duration              *string `json:"duration"`
	Ippon                 *string `json:"ippon"`
	Waza                  *string `json:"waza"`
	Yuko                  *string `json:"yuko"`
	Penalty               *string `json:"penalty"`
	IpponBlue             *string `json:"ippon_b"`
	WazaBlue              *string `json:"waza_b"`
	YukoBlue              *string `json:"yuko_b"`
	PenaltyBlue           *string `json:"penalty_b"`
	IpponWhite            *string `json:"ippon_w"`
	WazaWhite             *string `json:"waza_w"`
	YukoWhite             *string `json:"yuko_w"`
	PenaltyWhite          *string `json:"penalty_w"`
	Type                  *string `json:"type"`
	RoundCode             *string `json:"round_code"`
	RoundName             *string `json:"round_name"`
	DateStart             *int64  `json:"date_start_ts,string"`
	FirstHajime           *int64  `json:"first_hajime_at_ts,string"`
	GoldenScore           *string `json:"gs"`
	Bye                   *string `json:"bye"`
	UpdatedAt             *string `json:"updated_at"`
	HSKWhite              *string `json:"hsk_w"`
	HSKBlue               *string `json:"hsk_b"`
	Tagged                *string `json:"tagged"`
	KodokanTagged         *string `json:"kodokan_tagged"`
	Published             *string `json:"published"`
	ScCountdownOffset     *string `json:"sc_countdown_offset"`
	FightNo               *string `json:"fight_no"`
	ContestCodeLong       *string `json:"contest_code_long"`
	Mat                   *string `json:"mat"`
	IDCompetitionTeams    *string `json:"id_competition_teams"`
	IDFightTeam           *string `json:"id_fight_team"`
	CompetitionName       *string `json:"competition_name"`
	CompetitionDate       *string `json:"competition_date"`
	ExternalID            *string `json:"external_id"`
	City                  *string `json:"city"`
	CompYear              *string `json:"comp_year"`
	Age                   *string `json:"age"`
	FightDuration         *string `json:"fight_duration"`
	Weight                *string `json:"weight"`
	IDWeight              *string `json:"id_weight"`
	DateRaw               *string `json:"date_raw"`
	PersonWhite           *string `json:"person_white"`
	IDIJFWhite            *string `json:"id_ijf_white"`
	FamilyNameWhite       *string `json:"family_name_white"`
	GivenNameWhite        *string `json:"given_name_white"`
	TimestampVersionWhite *string `json:"timestamp_version_white"`
	PersonBlue            *string `json:"person_blue"`
	IDIJFBlue             *string `json:"id_ijf_blue"`
	FamilyNameBlue        *string `json:"family_name_blue"`
	GivenNameBlue         *string `json:"given_name_blue"`
	TimestampVersionBlue  *string `json:"timestamp_version_blue"`
	CountryWhite          *string `json:"country_white"`
	CountryShortWhite     *string `json:"country_short_white"`
	IDCountryWhite        *string `json:"id_country_white"`
	CountryBlue           *string `json:"country_blue"`
	CountryShortBlue      *string `json:"country_short_blue"`
	IDCountryBlue         *string `json:"id_country_blue"`
	PictureFolder1        *string `json:"picture_folder_1"`
	PictureFilename1      *string `json:"picture_filename_1"`
	PictureFolder2        *string `json:"picture_folder_2"`
	PictureFilename2      *string `json:"picture_filename_2"`
	Media                 *string `json:"media"`
	RankName              *string `json:"rank_name"`
	PersonalPictureWhite  *string `json:"personal_picture_white"`
	PersonalPictureBlue   *string `json:"personal_picture_blue"`
}

type Category struct {
	Name     string
	Contests []Contest `json:"contests"`
}

type Competition struct {
	Name       string
	Categories []Category
}
