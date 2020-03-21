/*

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// PlayerStatsSpec defines the desired state of PlayerStats
type PlayerStatsSpec struct {
	ChanceOfPlayingNextRound int    `json:"chance_of_playing_next_round"`
	ChanceOfPlayingThisRound int    `json:"chance_of_playing_this_round"`
	Code                     int    `json:"code"`
	CostChangeEvent          int    `json:"cost_change_event"`
	CostChangeEventFall      int    `json:"cost_change_event_fall"`
	CostChangeStart          int    `json:"cost_change_start"`
	CostChangeStartFall      int    `json:"cost_change_start_fall"`
	DreamteamCount           int    `json:"dreamteam_count"`
	ElementType              int    `json:"element_type"`
	EpNext                   string `json:"ep_next"`
	EpThis                   string `json:"ep_this"`
	EventPoints              int    `json:"event_points"`
	FirstName                string `json:"first_name"`
	Form                     string `json:"form"`
	Id                       int    `json:"id"`
	InDreamteam              bool   `json:"in_dreamteam"`
	News                     string `json:"news"`
	NewsAdded                string `json:"news_added"`
	NowCost                  int    `json:"now_cost"`
	Photo                    string `json:"photo"`
	PointsPerGame            string `json:"points_per_game"`
	SecondName               string `json:"second_name"`
	SelectedByPercent        string `json:"selected_by_percent"`
	Special                  bool   `json:"special"`
	SquadNumber              string `json:"squad_number"`
	Status                   string `json:"status"`
	Team                     int    `json:"team"`
	TeamCode                 int    `json:"team_code"`
	TotalPoints              int    `json:"total_points"`
	TransfersIn              int    `json:"transfers_in"`
	TransfersInEvent         int    `json:"transfers_in_event"`
	TransfersOut             int    `json:"transfers_out"`
	TransfersOutEvent        int    `json:"transfers_out_event"`
	ValueForm                string `json:"value_form"`
	ValueSeason              string `json:"value_season"`
	WebName                  string `json:"web_name"`
	Minutes                  int    `json:"minutes"`
	GoalsScored              int    `json:"goals_scored"`
	Assists                  int    `json:"assists"`
	CleanSheets              int    `json:"clean_sheets"`
	GoalsConceded            int    `json:"goals_conceded"`
	OwnGoals                 int    `json:"own_goals"`
	PenaltiesSaved           int    `json:"penalties_saved"`
	PenaltiesMissed          int    `json:"penalties_missed"`
	YellowCards              int    `json:"yellow_cards"`
	RedCards                 int    `json:"red_cards"`
	Saves                    int    `json:"saves"`
	Bonus                    int    `json:"bonus"`
	Bps                      int    `json:"bps"`
	Influence                string `json:"influence"`
	Creativity               string `json:"creativity"`
	Threat                   string `json:"threat"`
	IctIndex                 string `json:"ict_index"`
}

// PlayerStatsStatus defines the observed state of PlayerStats
type PlayerStatsStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// // FirstName of player 85
	// FirstName string `json:"firstname,omitempty"`
	// // LastName of player 94
	// LastName string `json:"lastname,omitempty"`
	// // 124
	// IctIndex int `json:"ictindex,omitempty"`
	// // 119
	// Bonus string `json:"bonus,omitempty"`
	// // 86
	// Form string `json:"form,omitempty"`
	// // 105
	// ValueForm string `json:"valueform,omitempty"`
	// // 101
	// TransferIn string `json:"transferin,omitempty"`
	// // 103
	// TransferOut string `json:"transferout,omitempty"`
	// // 91
	// Cost string `json:"cost,omitempty"`
	// // status 97
	// Status string `json:"status,omitempty"`
	// // assits 110
	// Assists string `json:"assists,omitempty"`
	// // goalsScored 110
	// GoalsScored string `json:"goalsscored,omitempty"`
	// // bpsIndex 120
	// BpsIndex string `json:"bpsindex,omitempty"`
	// // cleanSheets
	// CleanSheets string `json:"cleansheets,omitempty"`
	// // penaltiesMissed 114
	// PenaltiesMissed string `json:"penaltiesmissed,omitempty"`
	// // points 100
	// Points string `json:"points,omitempty"`
	// // position 61
	// Position string `json:"position,omitempty"`
	// // inDreamTeam 88
	// InDreamTeam string `json:"indreamteam,omitempty"`
	// // pointsPerGame 93
	// PointsPerGame string `json:"pointspergame,omitempty"`
	// // selectByPercent 95
	// SelectedByPercent string `json:"selectedbypercent,omitempty"`
	// // penaltiesSaved 114
	// PenaltiesSaved string `json:"penaltiessaved,omitempty"`
	// // saves 114
	// Saves string `json:"saves,omitempty"`
}

// +kubebuilder:object:root=true

// PlayerStats is the Schema for the playerstats API
type PlayerStats struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PlayerStatsSpec   `json:"spec,omitempty"`
	Status PlayerStatsStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PlayerStatsList contains a list of PlayerStats
type PlayerStatsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PlayerStats `json:"items"`
}

// // Team Struct to create a recommended team
// type Team struct {
// 	Players []PlayerStatsSpec
// }

func init() {
	SchemeBuilder.Register(&PlayerStats{}, &PlayerStatsList{})
}
