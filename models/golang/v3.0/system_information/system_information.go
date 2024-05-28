// Copyright 2024 MobilityData
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    systemInformation, err := UnmarshalSystemInformation(bytes)
//    bytes, err = systemInformation.Marshal()

package system_information



import "encoding/json"

func UnmarshalSystemInformation(data []byte) (SystemInformation, error) {
	var r SystemInformation
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SystemInformation) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

// Details including system operator, system location, year implemented, URL, contact info,
// time zone.
type SystemInformation struct {
	// Response data in the form of name:value pairs.                                                     
	Data                                                                                        Data      `json:"data"`
	// Last time the data in the feed was updated in RFC3339 format.                                      
	LastUpdated                                                                                 string `json:"last_updated"`
	// Number of seconds before the data in the feed will be updated again (0 if the data should          
	// always be refreshed).                                                                              
	TTL                                                                                         int64     `json:"ttl"`
	// GBFS version number to which the feed conforms, according to the versioning framework              
	// (added in v1.1).                                                                                   
	Version                                                                                     Version   `json:"version"`
}

// Response data in the form of name:value pairs.
type Data struct {
	// If the feed license requires attribution, name of the organization to which attribution                                
	// should be provided. An array with one object per supported language with the following                                 
	// keys:                                                                                                                  
	AttributionOrganizationName                                                                 []AttributionOrganizationName `json:"attribution_organization_name,omitempty"`
	// URL of the organization to which attribution should be provided.                                                       
	AttributionURL                                                                              *string                       `json:"attribution_url,omitempty"`
	// An object where each key defines one of the items listed below (added in v2.3-RC).                                     
	BrandAssets                                                                                 *BrandAssets                  `json:"brand_assets,omitempty"`
	// Email address actively monitored by the operator's customer service department.                                        
	Email                                                                                       *string                       `json:"email,omitempty"`
	// A single contact email address for consumers of this feed to report technical issues                                   
	// (added in v1.1).                                                                                                       
	FeedContactEmail                                                                            string                        `json:"feed_contact_email"`
	// List of languages used in translated strings. Each element in the list must be of type                                 
	// Language.                                                                                                              
	Languages                                                                                   []string                      `json:"languages"`
	// REQUIRED if the dataset is provided under a standard license. An identifier for a                                      
	// standard license from the SPDX License List. Provide license_id rather than license_url                                
	// if the license is included in the SPDX License List. See the GBFS wiki for a comparison                                
	// of a subset of standard licenses. If the license_id and license_url fields are blank or                                
	// omitted, this indicates that the feed is provided under the Creative Commons Universal                                 
	// Public Domain Dedication.                                                                                              
	LicenseID                                                                                   *LicenseID                    `json:"license_id,omitempty"`
	// A fully qualified URL of a page that defines the license terms for the GBFS data for this                              
	// system.                                                                                                                
	LicenseURL                                                                                  *string                       `json:"license_url,omitempty"`
	// REQUIRED if the producer publishes datasets for more than one system geography, for                                    
	// example Berlin and Paris. A fully qualified URL pointing to the manifest.json file for                                 
	// the publisher.                                                                                                         
	ManifestURL                                                                                 *string                       `json:"manifest_url,omitempty"`
	// Name of the system to be displayed to customers. An array with one object per supported                                
	// language with the following keys:                                                                                      
	Name                                                                                        []Name                        `json:"name"`
	// Hours and dates of operation for the system in OSM opening_hours format. (added in v3.0)                               
	OpeningHours                                                                                string                        `json:"opening_hours"`
	// Name of the system operator. An array with one object per supported language with the                                  
	// following keys:                                                                                                        
	Operator                                                                                    []Operator                    `json:"operator,omitempty"`
	// This OPTIONAL field SHOULD contain a single voice telephone number for the specified                                   
	// system’s customer service department. MUST be in E.164 format as defined in Field Types.                               
	PhoneNumber                                                                                 *string                       `json:"phone_number,omitempty"`
	// The date that the privacy policy provided at privacy_url was last updated (added in                                    
	// v2.3-RC).                                                                                                              
	PrivacyLastUpdated                                                                          *string                       `json:"privacy_last_updated,omitempty"`
	// A fully qualified URL pointing to the privacy policy for the service. An array with one                                
	// object per supported language with the following keys:                                                                 
	PrivacyURL                                                                                  []PrivacyURL                  `json:"privacy_url,omitempty"`
	// URL where a customer can purchase a membership.                                                                        
	PurchaseURL                                                                                 *string                       `json:"purchase_url,omitempty"`
	// Contains rental app information in the android and ios JSON objects (added in v1.1).                                   
	RentalApps                                                                                  *RentalApps                   `json:"rental_apps,omitempty"`
	// Abbreviation for a system. An array with one object per supported language with the                                    
	// following keys:                                                                                                        
	ShortName                                                                                   []ShortName                   `json:"short_name,omitempty"`
	// Date that the system began operations.                                                                                 
	StartDate                                                                                   *string                       `json:"start_date,omitempty"`
	// Identifier for this vehicle share system. This should be globally unique (even between                                 
	// different systems).                                                                                                    
	SystemID                                                                                    string                        `json:"system_id"`
	// Date after which this data source will no longer be available to consuming applications.                               
	TerminationDate                                                                             *string                       `json:"termination_date,omitempty"`
	// The date that the terms of service provided at terms_url were last updated (added in                                   
	// v2.3-RC)                                                                                                               
	TermsLastUpdated                                                                            *string                       `json:"terms_last_updated,omitempty"`
	// A fully qualified URL pointing to the terms of service (also often called "terms of use"                               
	// or "terms and conditions") for the service. An array with one object per supported                                     
	// language with the following keys:                                                                                      
	TermsURL                                                                                    []TermsURL                    `json:"terms_url,omitempty"`
	// The time zone where the system is located.                                                                             
	Timezone                                                                                    Timezone                      `json:"timezone"`
	// The URL of the vehicle share system.                                                                                   
	URL                                                                                         *string                       `json:"url,omitempty"`
}

type AttributionOrganizationName struct {
	// IETF BCP 47 language code.       
	Language                     string `json:"language"`
	// The translated text.             
	Text                         string `json:"text"`
}

// An object where each key defines one of the items listed below (added in v2.3-RC).
type BrandAssets struct {
	// A fully qualified URL pointing to the location of a graphic file representing the brand        
	// for the service (added in v2.3-RC).                                                            
	BrandImageURL                                                                             string  `json:"brand_image_url"`
	// A fully qualified URL pointing to the location of a graphic file representing the brand        
	// for the service for use in dark mode (added in v2.3-RC).                                       
	BrandImageURLDark                                                                         *string `json:"brand_image_url_dark,omitempty"`
	// Date that indicates the last time any included brand assets were updated (added in             
	// v2.3-RC).                                                                                      
	BrandLastModified                                                                         string  `json:"brand_last_modified"`
	// A fully qualified URL pointing to the location of a page that defines the license terms        
	// of brand icons, colors or other trademark information (added in v2.3-RC).                      
	BrandTermsURL                                                                             *string `json:"brand_terms_url,omitempty"`
	// Color used to represent the brand for the service (added in v2.3-RC)                           
	Color                                                                                     *string `json:"color,omitempty"`
}

type Name struct {
	// IETF BCP 47 language code.       
	Language                     string `json:"language"`
	// The translated text.             
	Text                         string `json:"text"`
}

type Operator struct {
	// IETF BCP 47 language code.       
	Language                     string `json:"language"`
	// The translated text.             
	Text                         string `json:"text"`
}

type PrivacyURL struct {
	// IETF BCP 47 language code.       
	Language                     string `json:"language"`
	// The translated text.             
	Text                         string `json:"text"`
}

// Contains rental app information in the android and ios JSON objects (added in v1.1).
type RentalApps struct {
	// Contains rental app download and app discovery information for the Android platform.         
	// (added in v1.1)                                                                              
	Android                                                                                *Android `json:"android,omitempty"`
	// Contains rental information for the iOS platform (added in v1.1).                            
	Ios                                                                                    *Ios     `json:"ios,omitempty"`
}

// Contains rental app download and app discovery information for the Android platform.
// (added in v1.1)
type Android struct {
	// URI that can be used to discover if the rental Android app is installed on the device       
	// (added in v1.1).                                                                            
	DiscoveryURI                                                                            string `json:"discovery_uri"`
	// URI where the rental Android app can be downloaded from (added in v1.1).                    
	StoreURI                                                                                string `json:"store_uri"`
}

// Contains rental information for the iOS platform (added in v1.1).
type Ios struct {
	// URI that can be used to discover if the rental iOS app is installed on the device (added       
	// in v1.1).                                                                                      
	DiscoveryURI                                                                               string `json:"discovery_uri"`
	// URI where the rental iOS app can be downloaded from (added in v1.1).                           
	StoreURI                                                                                   string `json:"store_uri"`
}

type ShortName struct {
	// IETF BCP 47 language code.       
	Language                     string `json:"language"`
	// The translated text.             
	Text                         string `json:"text"`
}

type TermsURL struct {
	// IETF BCP 47 language code.       
	Language                     string `json:"language"`
	// The translated text.             
	Text                         string `json:"text"`
}

// REQUIRED if the dataset is provided under a standard license. An identifier for a
// standard license from the SPDX License List. Provide license_id rather than license_url
// if the license is included in the SPDX License List. See the GBFS wiki for a comparison
// of a subset of standard licenses. If the license_id and license_url fields are blank or
// omitted, this indicates that the feed is provided under the Creative Commons Universal
// Public Domain Dedication.
type LicenseID string

const (
	AAL                               LicenseID = "AAL"
	ADSL                              LicenseID = "ADSL"
	AGPL10Only                        LicenseID = "AGPL-1.0-only"
	AGPL10OrLater                     LicenseID = "AGPL-1.0-or-later"
	AGPL30Only                        LicenseID = "AGPL-3.0-only"
	AGPL30OrLater                     LicenseID = "AGPL-3.0-or-later"
	ANTLRPDFallback                   LicenseID = "ANTLR-PD-fallback"
	APL10                             LicenseID = "APL-1.0"
	Abstyles                          LicenseID = "Abstyles"
	AdaCoreDoc                        LicenseID = "AdaCore-doc"
	Adobe2006                         LicenseID = "Adobe-2006"
	AdobeGlyph                        LicenseID = "Adobe-Glyph"
	Afl11                             LicenseID = "AFL-1.1"
	Afl12                             LicenseID = "AFL-1.2"
	Afl20                             LicenseID = "AFL-2.0"
	Afl21                             LicenseID = "AFL-2.1"
	Afl30                             LicenseID = "AFL-3.0"
	Afmparse                          LicenseID = "Afmparse"
	Aladdin                           LicenseID = "Aladdin"
	Amdplpa                           LicenseID = "AMDPLPA"
	Aml                               LicenseID = "AML"
	Ampas                             LicenseID = "AMPAS"
	AntlrPD                           LicenseID = "ANTLR-PD"
	Apache10                          LicenseID = "Apache-1.0"
	Apache11                          LicenseID = "Apache-1.1"
	Apache20                          LicenseID = "Apache-2.0"
	Apafml                            LicenseID = "APAFML"
	AppS2P                            LicenseID = "App-s2p"
	Apsl10                            LicenseID = "APSL-1.0"
	Apsl11                            LicenseID = "APSL-1.1"
	Apsl12                            LicenseID = "APSL-1.2"
	Apsl20                            LicenseID = "APSL-2.0"
	Arphic1999                        LicenseID = "Arphic-1999"
	Artistic10                        LicenseID = "Artistic-1.0"
	Artistic10Cl8                     LicenseID = "Artistic-1.0-cl8"
	Artistic10PERL                    LicenseID = "Artistic-1.0-Perl"
	Artistic20                        LicenseID = "Artistic-2.0"
	BSD1Clause                        LicenseID = "BSD-1-Clause"
	BSD2Clause                        LicenseID = "BSD-2-Clause"
	BSD2ClausePatent                  LicenseID = "BSD-2-Clause-Patent"
	BSD2ClauseViews                   LicenseID = "BSD-2-Clause-Views"
	BSD3Clause                        LicenseID = "BSD-3-Clause"
	BSD3ClauseAttribution             LicenseID = "BSD-3-Clause-Attribution"
	BSD3ClauseClear                   LicenseID = "BSD-3-Clause-Clear"
	BSD3ClauseLBNL                    LicenseID = "BSD-3-Clause-LBNL"
	BSD3ClauseModification            LicenseID = "BSD-3-Clause-Modification"
	BSD3ClauseNoMilitaryLicense       LicenseID = "BSD-3-Clause-No-Military-License"
	BSD3ClauseNoNuclearLicense        LicenseID = "BSD-3-Clause-No-Nuclear-License"
	BSD3ClauseNoNuclearLicense2014    LicenseID = "BSD-3-Clause-No-Nuclear-License-2014"
	BSD3ClauseNoNuclearWarranty       LicenseID = "BSD-3-Clause-No-Nuclear-Warranty"
	BSD3ClauseOpenMPI                 LicenseID = "BSD-3-Clause-Open-MPI"
	BSD43Reno                         LicenseID = "BSD-4.3RENO"
	BSD43Tahoe                        LicenseID = "BSD-4.3TAHOE"
	BSD4Clause                        LicenseID = "BSD-4-Clause"
	BSD4ClauseShortened               LicenseID = "BSD-4-Clause-Shortened"
	BSD4ClauseUC                      LicenseID = "BSD-4-Clause-UC"
	BSDAdvertisingAcknowledgement     LicenseID = "BSD-Advertising-Acknowledgement"
	BSDAttributionHPNDDisclaimer      LicenseID = "BSD-Attribution-HPND-disclaimer"
	BSDProtection                     LicenseID = "BSD-Protection"
	BSDSourceCode                     LicenseID = "BSD-Source-Code"
	Baekmuk                           LicenseID = "Baekmuk"
	Bahyph                            LicenseID = "Bahyph"
	Barr                              LicenseID = "Barr"
	Beerware                          LicenseID = "Beerware"
	BitTorrent10                      LicenseID = "BitTorrent-1.0"
	BitTorrent11                      LicenseID = "BitTorrent-1.1"
	BitstreamCharter                  LicenseID = "Bitstream-Charter"
	BitstreamVera                     LicenseID = "Bitstream-Vera"
	Blessing                          LicenseID = "blessing"
	BlueOak100                        LicenseID = "BlueOak-1.0.0"
	Borceux                           LicenseID = "Borceux"
	BrianGladman3Clause               LicenseID = "Brian-Gladman-3-Clause"
	Bsl10                             LicenseID = "BSL-1.0"
	Busl11                            LicenseID = "BUSL-1.1"
	Bzip2106                          LicenseID = "bzip2-1.0.6"
	CAL10CombinedWorkException        LicenseID = "CAL-1.0-Combined-Work-Exception"
	CDLAPermissive10                  LicenseID = "CDLA-Permissive-1.0"
	CDLAPermissive20                  LicenseID = "CDLA-Permissive-2.0"
	CDLASharing10                     LicenseID = "CDLA-Sharing-1.0"
	CMUMach                           LicenseID = "CMU-Mach"
	CNRIJython                        LicenseID = "CNRI-Jython"
	CNRIPython                        LicenseID = "CNRI-Python"
	CNRIPythonGPLCompatible           LicenseID = "CNRI-Python-GPL-Compatible"
	CUAOpl10                          LicenseID = "CUA-OPL-1.0"
	CUda10                            LicenseID = "C-UDA-1.0"
	Cal10                             LicenseID = "CAL-1.0"
	Caldera                           LicenseID = "Caldera"
	Catosl11                          LicenseID = "CATOSL-1.1"
	Cc010                             LicenseID = "CC0-1.0"
	CcBy10                            LicenseID = "CC-BY-1.0"
	CcBy20                            LicenseID = "CC-BY-2.0"
	CcBy25                            LicenseID = "CC-BY-2.5"
	CcBy25Au                          LicenseID = "CC-BY-2.5-AU"
	CcBy30                            LicenseID = "CC-BY-3.0"
	CcBy30At                          LicenseID = "CC-BY-3.0-AT"
	CcBy30De                          LicenseID = "CC-BY-3.0-DE"
	CcBy30Igo                         LicenseID = "CC-BY-3.0-IGO"
	CcBy30Nl                          LicenseID = "CC-BY-3.0-NL"
	CcBy30Us                          LicenseID = "CC-BY-3.0-US"
	CcBy40                            LicenseID = "CC-BY-4.0"
	CcByNc10                          LicenseID = "CC-BY-NC-1.0"
	CcByNc20                          LicenseID = "CC-BY-NC-2.0"
	CcByNc25                          LicenseID = "CC-BY-NC-2.5"
	CcByNc30                          LicenseID = "CC-BY-NC-3.0"
	CcByNc30De                        LicenseID = "CC-BY-NC-3.0-DE"
	CcByNc40                          LicenseID = "CC-BY-NC-4.0"
	CcByNcNd10                        LicenseID = "CC-BY-NC-ND-1.0"
	CcByNcNd20                        LicenseID = "CC-BY-NC-ND-2.0"
	CcByNcNd25                        LicenseID = "CC-BY-NC-ND-2.5"
	CcByNcNd30                        LicenseID = "CC-BY-NC-ND-3.0"
	CcByNcNd30De                      LicenseID = "CC-BY-NC-ND-3.0-DE"
	CcByNcNd30Igo                     LicenseID = "CC-BY-NC-ND-3.0-IGO"
	CcByNcNd40                        LicenseID = "CC-BY-NC-ND-4.0"
	CcByNcSa10                        LicenseID = "CC-BY-NC-SA-1.0"
	CcByNcSa20                        LicenseID = "CC-BY-NC-SA-2.0"
	CcByNcSa20De                      LicenseID = "CC-BY-NC-SA-2.0-DE"
	CcByNcSa20Fr                      LicenseID = "CC-BY-NC-SA-2.0-FR"
	CcByNcSa20Uk                      LicenseID = "CC-BY-NC-SA-2.0-UK"
	CcByNcSa25                        LicenseID = "CC-BY-NC-SA-2.5"
	CcByNcSa30                        LicenseID = "CC-BY-NC-SA-3.0"
	CcByNcSa30De                      LicenseID = "CC-BY-NC-SA-3.0-DE"
	CcByNcSa30Igo                     LicenseID = "CC-BY-NC-SA-3.0-IGO"
	CcByNcSa40                        LicenseID = "CC-BY-NC-SA-4.0"
	CcByNd10                          LicenseID = "CC-BY-ND-1.0"
	CcByNd20                          LicenseID = "CC-BY-ND-2.0"
	CcByNd25                          LicenseID = "CC-BY-ND-2.5"
	CcByNd30                          LicenseID = "CC-BY-ND-3.0"
	CcByNd30De                        LicenseID = "CC-BY-ND-3.0-DE"
	CcByNd40                          LicenseID = "CC-BY-ND-4.0"
	CcBySa10                          LicenseID = "CC-BY-SA-1.0"
	CcBySa20                          LicenseID = "CC-BY-SA-2.0"
	CcBySa20Uk                        LicenseID = "CC-BY-SA-2.0-UK"
	CcBySa21Jp                        LicenseID = "CC-BY-SA-2.1-JP"
	CcBySa25                          LicenseID = "CC-BY-SA-2.5"
	CcBySa30                          LicenseID = "CC-BY-SA-3.0"
	CcBySa30At                        LicenseID = "CC-BY-SA-3.0-AT"
	CcBySa30De                        LicenseID = "CC-BY-SA-3.0-DE"
	CcBySa40                          LicenseID = "CC-BY-SA-4.0"
	CcPddc                            LicenseID = "CC-PDDC"
	Cddl10                            LicenseID = "CDDL-1.0"
	Cddl11                            LicenseID = "CDDL-1.1"
	Cdl10                             LicenseID = "CDL-1.0"
	Cecill10                          LicenseID = "CECILL-1.0"
	Cecill11                          LicenseID = "CECILL-1.1"
	Cecill20                          LicenseID = "CECILL-2.0"
	Cecill21                          LicenseID = "CECILL-2.1"
	CecillB                           LicenseID = "CECILL-B"
	CecillC                           LicenseID = "CECILL-C"
	CernOhl11                         LicenseID = "CERN-OHL-1.1"
	CernOhl12                         LicenseID = "CERN-OHL-1.2"
	CernOhlP20                        LicenseID = "CERN-OHL-P-2.0"
	CernOhlS20                        LicenseID = "CERN-OHL-S-2.0"
	CernOhlW20                        LicenseID = "CERN-OHL-W-2.0"
	Cfitsio                           LicenseID = "CFITSIO"
	Checkmk                           LicenseID = "checkmk"
	ClArtistic                        LicenseID = "ClArtistic"
	Clips                             LicenseID = "Clips"
	Coil10                            LicenseID = "COIL-1.0"
	CommunitySpec10                   LicenseID = "Community-Spec-1.0"
	Condor11                          LicenseID = "Condor-1.1"
	CopyleftNext030                   LicenseID = "copyleft-next-0.3.0"
	CopyleftNext031                   LicenseID = "copyleft-next-0.3.1"
	CornellLosslessJPEG               LicenseID = "Cornell-Lossless-JPEG"
	Cpal10                            LicenseID = "CPAL-1.0"
	Cpl10                             LicenseID = "CPL-1.0"
	Cpol102                           LicenseID = "CPOL-1.02"
	Crossword                         LicenseID = "Crossword"
	CrystalStacker                    LicenseID = "CrystalStacker"
	Cube                              LicenseID = "Cube"
	Curl                              LicenseID = "curl"
	DFsl10                            LicenseID = "D-FSL-1.0"
	DLDeBy20                          LicenseID = "DL-DE-BY-2.0"
	Diffmark                          LicenseID = "diffmark"
	Doc                               LicenseID = "DOC"
	Dotseqn                           LicenseID = "Dotseqn"
	Drl10                             LicenseID = "DRL-1.0"
	Dsdp                              LicenseID = "DSDP"
	Dvipdfm                           LicenseID = "dvipdfm"
	EGenix                            LicenseID = "eGenix"
	EUDatagrid                        LicenseID = "EUDatagrid"
	Ecl10                             LicenseID = "ECL-1.0"
	Ecl20                             LicenseID = "ECL-2.0"
	Efl10                             LicenseID = "EFL-1.0"
	Efl20                             LicenseID = "EFL-2.0"
	Elastic20                         LicenseID = "Elastic-2.0"
	Entessa                           LicenseID = "Entessa"
	Epics                             LicenseID = "EPICS"
	Epl10                             LicenseID = "EPL-1.0"
	Epl20                             LicenseID = "EPL-2.0"
	ErlPL11                           LicenseID = "ErlPL-1.1"
	Etalab20                          LicenseID = "etalab-2.0"
	Eupl10                            LicenseID = "EUPL-1.0"
	Eupl11                            LicenseID = "EUPL-1.1"
	Eupl12                            LicenseID = "EUPL-1.2"
	Eurosym                           LicenseID = "Eurosym"
	Fair                              LicenseID = "Fair"
	FdkAAC                            LicenseID = "FDK-AAC"
	Frameworx10                       LicenseID = "Frameworx-1.0"
	FreeBSDDOC                        LicenseID = "FreeBSD-DOC"
	FreeImage                         LicenseID = "FreeImage"
	Fsfap                             LicenseID = "FSFAP"
	Fsful                             LicenseID = "FSFUL"
	Fsfullr                           LicenseID = "FSFULLR"
	Fsfullrwd                         LicenseID = "FSFULLRWD"
	Ftl                               LicenseID = "FTL"
	GFDL11InvariantsOnly              LicenseID = "GFDL-1.1-invariants-only"
	GFDL11InvariantsOrLater           LicenseID = "GFDL-1.1-invariants-or-later"
	GFDL11NoInvariantsOnly            LicenseID = "GFDL-1.1-no-invariants-only"
	GFDL11NoInvariantsOrLater         LicenseID = "GFDL-1.1-no-invariants-or-later"
	GFDL11Only                        LicenseID = "GFDL-1.1-only"
	GFDL11OrLater                     LicenseID = "GFDL-1.1-or-later"
	GFDL12InvariantsOnly              LicenseID = "GFDL-1.2-invariants-only"
	GFDL12InvariantsOrLater           LicenseID = "GFDL-1.2-invariants-or-later"
	GFDL12NoInvariantsOnly            LicenseID = "GFDL-1.2-no-invariants-only"
	GFDL12NoInvariantsOrLater         LicenseID = "GFDL-1.2-no-invariants-or-later"
	GFDL12Only                        LicenseID = "GFDL-1.2-only"
	GFDL12OrLater                     LicenseID = "GFDL-1.2-or-later"
	GFDL13InvariantsOnly              LicenseID = "GFDL-1.3-invariants-only"
	GFDL13InvariantsOrLater           LicenseID = "GFDL-1.3-invariants-or-later"
	GFDL13NoInvariantsOnly            LicenseID = "GFDL-1.3-no-invariants-only"
	GFDL13NoInvariantsOrLater         LicenseID = "GFDL-1.3-no-invariants-or-later"
	GFDL13Only                        LicenseID = "GFDL-1.3-only"
	GFDL13OrLater                     LicenseID = "GFDL-1.3-or-later"
	GPL10Only                         LicenseID = "GPL-1.0-only"
	GPL10OrLater                      LicenseID = "GPL-1.0-or-later"
	GPL20Only                         LicenseID = "GPL-2.0-only"
	GPL20OrLater                      LicenseID = "GPL-2.0-or-later"
	GPL30Only                         LicenseID = "GPL-3.0-only"
	GPL30OrLater                      LicenseID = "GPL-3.0-or-later"
	GSOAP13B                          LicenseID = "gSOAP-1.3b"
	Gd                                LicenseID = "GD"
	Giftware                          LicenseID = "Giftware"
	Gl2PS                             LicenseID = "GL2PS"
	Glide                             LicenseID = "Glide"
	Glulxe                            LicenseID = "Glulxe"
	Glwtpl                            LicenseID = "GLWTPL"
	Gnuplot                           LicenseID = "gnuplot"
	GraphicsGems                      LicenseID = "Graphics-Gems"
	HP1986                            LicenseID = "HP-1986"
	HPNDExportUS                      LicenseID = "HPND-export-US"
	HPNDMarkusKuhn                    LicenseID = "HPND-Markus-Kuhn"
	HPNDSellVariant                   LicenseID = "HPND-sell-variant"
	HPNDSellVariantMITDisclaimer      LicenseID = "HPND-sell-variant-MIT-disclaimer"
	HaskellReport                     LicenseID = "HaskellReport"
	Hippocratic21                     LicenseID = "Hippocratic-2.1"
	Hpnd                              LicenseID = "HPND"
	Htmltidy                          LicenseID = "HTMLTIDY"
	IBMPibs                           LicenseID = "IBM-pibs"
	IECCodeComponentsEULA             LicenseID = "IEC-Code-Components-EULA"
	IJGShort                          LicenseID = "IJG-short"
	IMatix                            LicenseID = "iMatix"
	IPL10                             LicenseID = "IPL-1.0"
	ISC                               LicenseID = "ISC"
	Icu                               LicenseID = "ICU"
	Ijg                               LicenseID = "IJG"
	ImageMagick                       LicenseID = "ImageMagick"
	Imlib2                            LicenseID = "Imlib2"
	InfoZIP                           LicenseID = "Info-ZIP"
	Intel                             LicenseID = "Intel"
	IntelACPI                         LicenseID = "Intel-ACPI"
	Interbase10                       LicenseID = "Interbase-1.0"
	Ipa                               LicenseID = "IPA"
	JPLImage                          LicenseID = "JPL-image"
	JSON                              LicenseID = "JSON"
	Jam                               LicenseID = "Jam"
	JasPer20                          LicenseID = "JasPer-2.0"
	Jpnic                             LicenseID = "JPNIC"
	Kazlib                            LicenseID = "Kazlib"
	KnuthCTAN                         LicenseID = "Knuth-CTAN"
	LGPL20Only                        LicenseID = "LGPL-2.0-only"
	LGPL20OrLater                     LicenseID = "LGPL-2.0-or-later"
	LGPL21Only                        LicenseID = "LGPL-2.1-only"
	LGPL21OrLater                     LicenseID = "LGPL-2.1-or-later"
	LGPL30Only                        LicenseID = "LGPL-3.0-only"
	LGPL30OrLater                     LicenseID = "LGPL-3.0-or-later"
	LPPL13A                           LicenseID = "LPPL-1.3a"
	LPPL13C                           LicenseID = "LPPL-1.3c"
	LZMASDK911To920                   LicenseID = "LZMA-SDK-9.11-to-9.20"
	Lal12                             LicenseID = "LAL-1.2"
	Lal13                             LicenseID = "LAL-1.3"
	Latex2E                           LicenseID = "Latex2e"
	Leptonica                         LicenseID = "Leptonica"
	Lgpllr                            LicenseID = "LGPLLR"
	LiLiQP11                          LicenseID = "LiLiQ-P-1.1"
	LiLiQR11                          LicenseID = "LiLiQ-R-1.1"
	LiLiQRplus11                      LicenseID = "LiLiQ-Rplus-1.1"
	Libpng                            LicenseID = "Libpng"
	Libpng20                          LicenseID = "libpng-2.0"
	Libselinux10                      LicenseID = "libselinux-1.0"
	Libtiff                           LicenseID = "libtiff"
	LibutilDavidNugent                LicenseID = "libutil-David-Nugent"
	LinuxManPagesCopyleft             LicenseID = "Linux-man-pages-copyleft"
	LinuxOpenIB                       LicenseID = "Linux-OpenIB"
	Loop                              LicenseID = "LOOP"
	Lpl10                             LicenseID = "LPL-1.0"
	Lpl102                            LicenseID = "LPL-1.02"
	Lppl10                            LicenseID = "LPPL-1.0"
	Lppl11                            LicenseID = "LPPL-1.1"
	Lppl12                            LicenseID = "LPPL-1.2"
	LzmaSDK922                        LicenseID = "LZMA-SDK-9.22"
	MIT                               LicenseID = "MIT"
	MIT0                              LicenseID = "MIT-0"
	MITAdvertising                    LicenseID = "MIT-advertising"
	MITCmu                            LicenseID = "MIT-CMU"
	MITEnna                           LicenseID = "MIT-enna"
	MITFeh                            LicenseID = "MIT-feh"
	MITModernVariant                  LicenseID = "MIT-Modern-Variant"
	MITOpenGroup                      LicenseID = "MIT-open-group"
	MITWu                             LicenseID = "MIT-Wu"
	MPL10                             LicenseID = "MPL-1.0"
	MPL11                             LicenseID = "MPL-1.1"
	MPL20                             LicenseID = "MPL-2.0"
	MPL20NoCopyleftException          LicenseID = "MPL-2.0-no-copyleft-exception"
	MSLpl                             LicenseID = "MS-LPL"
	MSPl                              LicenseID = "MS-PL"
	MSRl                              LicenseID = "MS-RL"
	MakeIndex                         LicenseID = "MakeIndex"
	MartinBirgmeier                   LicenseID = "Martin-Birgmeier"
	Minpack                           LicenseID = "Minpack"
	MirOS                             LicenseID = "MirOS"
	Mitnfa                            LicenseID = "MITNFA"
	Motosoto                          LicenseID = "Motosoto"
	MpiPermissive                     LicenseID = "mpi-permissive"
	Mpich2                            LicenseID = "mpich2"
	Mplus                             LicenseID = "mplus"
	Mtll                              LicenseID = "MTLL"
	MulanPSL10                        LicenseID = "MulanPSL-1.0"
	MulanPSL20                        LicenseID = "MulanPSL-2.0"
	Multics                           LicenseID = "Multics"
	Mup                               LicenseID = "Mup"
	NCSA                              LicenseID = "NCSA"
	NISTPD                            LicenseID = "NIST-PD"
	NISTPDFallback                    LicenseID = "NIST-PD-fallback"
	NPL10                             LicenseID = "NPL-1.0"
	NPL11                             LicenseID = "NPL-1.1"
	NTP                               LicenseID = "NTP"
	NTP0                              LicenseID = "NTP-0"
	Naist2003                         LicenseID = "NAIST-2003"
	Nasa13                            LicenseID = "NASA-1.3"
	Naumen                            LicenseID = "Naumen"
	Nbpl10                            LicenseID = "NBPL-1.0"
	NcglUk20                          LicenseID = "NCGL-UK-2.0"
	NetCDF                            LicenseID = "NetCDF"
	NetSNMP                           LicenseID = "Net-SNMP"
	Newsletr                          LicenseID = "Newsletr"
	Ngpl                              LicenseID = "NGPL"
	Nicta10                           LicenseID = "NICTA-1.0"
	Nlod10                            LicenseID = "NLOD-1.0"
	Nlod20                            LicenseID = "NLOD-2.0"
	Nlpl                              LicenseID = "NLPL"
	Nokia                             LicenseID = "Nokia"
	Nosl                              LicenseID = "NOSL"
	Noweb                             LicenseID = "Noweb"
	Nposl30                           LicenseID = "NPOSL-3.0"
	Nrl                               LicenseID = "NRL"
	ODBL10                            LicenseID = "ODbL-1.0"
	ODCBy10                           LicenseID = "ODC-By-1.0"
	OFL10NoRFN                        LicenseID = "OFL-1.0-no-RFN"
	OFL11NoRFN                        LicenseID = "OFL-1.1-no-RFN"
	OGDLTaiwan10                      LicenseID = "OGDL-Taiwan-1.0"
	OGLCanada20                       LicenseID = "OGL-Canada-2.0"
	OUda10                            LicenseID = "O-UDA-1.0"
	OcctPl                            LicenseID = "OCCT-PL"
	Oclc20                            LicenseID = "OCLC-2.0"
	Offis                             LicenseID = "OFFIS"
	Ofl10                             LicenseID = "OFL-1.0"
	Ofl10Rfn                          LicenseID = "OFL-1.0-RFN"
	Ofl11                             LicenseID = "OFL-1.1"
	Ofl11Rfn                          LicenseID = "OFL-1.1-RFN"
	Ogc10                             LicenseID = "OGC-1.0"
	OglUk10                           LicenseID = "OGL-UK-1.0"
	OglUk20                           LicenseID = "OGL-UK-2.0"
	OglUk30                           LicenseID = "OGL-UK-3.0"
	Ogtsl                             LicenseID = "OGTSL"
	Oldap11                           LicenseID = "OLDAP-1.1"
	Oldap12                           LicenseID = "OLDAP-1.2"
	Oldap13                           LicenseID = "OLDAP-1.3"
	Oldap14                           LicenseID = "OLDAP-1.4"
	Oldap20                           LicenseID = "OLDAP-2.0"
	Oldap201                          LicenseID = "OLDAP-2.0.1"
	Oldap21                           LicenseID = "OLDAP-2.1"
	Oldap22                           LicenseID = "OLDAP-2.2"
	Oldap221                          LicenseID = "OLDAP-2.2.1"
	Oldap222                          LicenseID = "OLDAP-2.2.2"
	Oldap23                           LicenseID = "OLDAP-2.3"
	Oldap24                           LicenseID = "OLDAP-2.4"
	Oldap25                           LicenseID = "OLDAP-2.5"
	Oldap26                           LicenseID = "OLDAP-2.6"
	Oldap27                           LicenseID = "OLDAP-2.7"
	Oldap28                           LicenseID = "OLDAP-2.8"
	Oml                               LicenseID = "OML"
	OpenPBS23                         LicenseID = "OpenPBS-2.3"
	OpenSSL                           LicenseID = "OpenSSL"
	Opl10                             LicenseID = "OPL-1.0"
	Opubl10                           LicenseID = "OPUBL-1.0"
	OsetPl21                          LicenseID = "OSET-PL-2.1"
	Osl10                             LicenseID = "OSL-1.0"
	Osl11                             LicenseID = "OSL-1.1"
	Osl20                             LicenseID = "OSL-2.0"
	Osl21                             LicenseID = "OSL-2.1"
	Osl30                             LicenseID = "OSL-3.0"
	PHP30                             LicenseID = "PHP-3.0"
	PHP301                            LicenseID = "PHP-3.01"
	Parity600                         LicenseID = "Parity-6.0.0"
	Parity700                         LicenseID = "Parity-7.0.0"
	Pddl10                            LicenseID = "PDDL-1.0"
	Plexus                            LicenseID = "Plexus"
	PolyFormNoncommercial100          LicenseID = "PolyForm-Noncommercial-1.0.0"
	PolyFormSmallBusiness100          LicenseID = "PolyForm-Small-Business-1.0.0"
	PostgreSQL                        LicenseID = "PostgreSQL"
	Psf20                             LicenseID = "PSF-2.0"
	Psfrag                            LicenseID = "psfrag"
	Psutils                           LicenseID = "psutils"
	Python20                          LicenseID = "Python-2.0"
	Python201                         LicenseID = "Python-2.0.1"
	Qhull                             LicenseID = "Qhull"
	Qpl10                             LicenseID = "QPL-1.0"
	Qpl10Inria2004                    LicenseID = "QPL-1.0-INRIA-2004"
	RHeCos11                          LicenseID = "RHeCos-1.1"
	RSAMd                             LicenseID = "RSA-MD"
	Rdisc                             LicenseID = "Rdisc"
	Rpl11                             LicenseID = "RPL-1.1"
	Rpl15                             LicenseID = "RPL-1.5"
	Rpsl10                            LicenseID = "RPSL-1.0"
	Rscpl                             LicenseID = "RSCPL"
	Ruby                              LicenseID = "Ruby"
	SAXPD                             LicenseID = "SAX-PD"
	SGIB10                            LicenseID = "SGI-B-1.0"
	SGIB11                            LicenseID = "SGI-B-1.1"
	SGIB20                            LicenseID = "SGI-B-2.0"
	SSHOpenSSH                        LicenseID = "SSH-OpenSSH"
	SSHShort                          LicenseID = "SSH-short"
	Saxpath                           LicenseID = "Saxpath"
	Scea                              LicenseID = "SCEA"
	SchemeReport                      LicenseID = "SchemeReport"
	Sendmail                          LicenseID = "Sendmail"
	Sendmail823                       LicenseID = "Sendmail-8.23"
	Shl05                             LicenseID = "SHL-0.5"
	Shl051                            LicenseID = "SHL-0.51"
	SimPL20                           LicenseID = "SimPL-2.0"
	Sissl                             LicenseID = "SISSL"
	Sissl12                           LicenseID = "SISSL-1.2"
	Sleepycat                         LicenseID = "Sleepycat"
	Smlnj                             LicenseID = "SMLNJ"
	Smppl                             LicenseID = "SMPPL"
	Snia                              LicenseID = "SNIA"
	Snprintf                          LicenseID = "snprintf"
	Spencer86                         LicenseID = "Spencer-86"
	Spencer94                         LicenseID = "Spencer-94"
	Spencer99                         LicenseID = "Spencer-99"
	Spl10                             LicenseID = "SPL-1.0"
	Sspl10                            LicenseID = "SSPL-1.0"
	SugarCRM113                       LicenseID = "SugarCRM-1.1.3"
	SunPro                            LicenseID = "SunPro"
	Swl                               LicenseID = "SWL"
	Symlinks                          LicenseID = "Symlinks"
	TCPWrappers                       LicenseID = "TCP-wrappers"
	TMate                             LicenseID = "TMate"
	TUBerlin10                        LicenseID = "TU-Berlin-1.0"
	TUBerlin20                        LicenseID = "TU-Berlin-2.0"
	TaprOhl10                         LicenseID = "TAPR-OHL-1.0"
	Tcl                               LicenseID = "TCL"
	The0BSD                           LicenseID = "0BSD"
	Torque11                          LicenseID = "TORQUE-1.1"
	Tosl                              LicenseID = "TOSL"
	Tpdl                              LicenseID = "TPDL"
	Tpl10                             LicenseID = "TPL-1.0"
	Ttwl                              LicenseID = "TTWL"
	Ucar                              LicenseID = "UCAR"
	Ucl10                             LicenseID = "UCL-1.0"
	UnicodeDFS2015                    LicenseID = "Unicode-DFS-2015"
	UnicodeDFS2016                    LicenseID = "Unicode-DFS-2016"
	UnicodeTOU                        LicenseID = "Unicode-TOU"
	Unlicense                         LicenseID = "Unlicense"
	Upl10                             LicenseID = "UPL-1.0"
	Vim                               LicenseID = "Vim"
	Vostrom                           LicenseID = "VOSTROM"
	Vsl10                             LicenseID = "VSL-1.0"
	W3C                               LicenseID = "W3C"
	W3C19980720                       LicenseID = "W3C-19980720"
	W3C20150513                       LicenseID = "W3C-20150513"
	W3M                               LicenseID = "w3m"
	Watcom10                          LicenseID = "Watcom-1.0"
	Wsuipa                            LicenseID = "Wsuipa"
	Wtfpl                             LicenseID = "WTFPL"
	X11                               LicenseID = "X11"
	X11DistributeModificationsVariant LicenseID = "X11-distribute-modifications-variant"
	XFree8611                         LicenseID = "XFree86-1.1"
	XSkat                             LicenseID = "XSkat"
	Xerox                             LicenseID = "Xerox"
	Xinetd                            LicenseID = "xinetd"
	Xlock                             LicenseID = "xlock"
	Xnet                              LicenseID = "Xnet"
	Xpp                               LicenseID = "xpp"
	Ypl10                             LicenseID = "YPL-1.0"
	Ypl11                             LicenseID = "YPL-1.1"
	ZPL11                             LicenseID = "ZPL-1.1"
	ZPL20                             LicenseID = "ZPL-2.0"
	ZPL21                             LicenseID = "ZPL-2.1"
	Zed                               LicenseID = "Zed"
	Zend20                            LicenseID = "Zend-2.0"
	Zimbra13                          LicenseID = "Zimbra-1.3"
	Zimbra14                          LicenseID = "Zimbra-1.4"
	Zlib                              LicenseID = "Zlib"
	ZlibAcknowledgement               LicenseID = "zlib-acknowledgement"
)

// The time zone where the system is located.
type Timezone string

const (
	AfricaAbidjan               Timezone = "Africa/Abidjan"
	AfricaAlgiers               Timezone = "Africa/Algiers"
	AfricaBissau                Timezone = "Africa/Bissau"
	AfricaCairo                 Timezone = "Africa/Cairo"
	AfricaCasablanca            Timezone = "Africa/Casablanca"
	AfricaCeuta                 Timezone = "Africa/Ceuta"
	AfricaElAaiun               Timezone = "Africa/El_Aaiun"
	AfricaJohannesburg          Timezone = "Africa/Johannesburg"
	AfricaJuba                  Timezone = "Africa/Juba"
	AfricaKhartoum              Timezone = "Africa/Khartoum"
	AfricaLagos                 Timezone = "Africa/Lagos"
	AfricaMaputo                Timezone = "Africa/Maputo"
	AfricaMonrovia              Timezone = "Africa/Monrovia"
	AfricaNairobi               Timezone = "Africa/Nairobi"
	AfricaNdjamena              Timezone = "Africa/Ndjamena"
	AfricaSaoTome               Timezone = "Africa/Sao_Tome"
	AfricaTripoli               Timezone = "Africa/Tripoli"
	AfricaTunis                 Timezone = "Africa/Tunis"
	AfricaWindhoek              Timezone = "Africa/Windhoek"
	AmericaAdak                 Timezone = "America/Adak"
	AmericaAnchorage            Timezone = "America/Anchorage"
	AmericaAraguaina            Timezone = "America/Araguaina"
	AmericaArgentinaBuenosAires Timezone = "America/Argentina/Buenos_Aires"
	AmericaArgentinaCatamarca   Timezone = "America/Argentina/Catamarca"
	AmericaArgentinaCordoba     Timezone = "America/Argentina/Cordoba"
	AmericaArgentinaJujuy       Timezone = "America/Argentina/Jujuy"
	AmericaArgentinaLaRioja     Timezone = "America/Argentina/La_Rioja"
	AmericaArgentinaMendoza     Timezone = "America/Argentina/Mendoza"
	AmericaArgentinaRioGallegos Timezone = "America/Argentina/Rio_Gallegos"
	AmericaArgentinaSANJuan     Timezone = "America/Argentina/San_Juan"
	AmericaArgentinaSANLuis     Timezone = "America/Argentina/San_Luis"
	AmericaArgentinaSalta       Timezone = "America/Argentina/Salta"
	AmericaArgentinaTucuman     Timezone = "America/Argentina/Tucuman"
	AmericaArgentinaUshuaia     Timezone = "America/Argentina/Ushuaia"
	AmericaAsuncion             Timezone = "America/Asuncion"
	AmericaBahia                Timezone = "America/Bahia"
	AmericaBahiaBanderas        Timezone = "America/Bahia_Banderas"
	AmericaBarbados             Timezone = "America/Barbados"
	AmericaBelem                Timezone = "America/Belem"
	AmericaBelize               Timezone = "America/Belize"
	AmericaBoaVista             Timezone = "America/Boa_Vista"
	AmericaBogota               Timezone = "America/Bogota"
	AmericaBoise                Timezone = "America/Boise"
	AmericaCambridgeBay         Timezone = "America/Cambridge_Bay"
	AmericaCampoGrande          Timezone = "America/Campo_Grande"
	AmericaCancun               Timezone = "America/Cancun"
	AmericaCaracas              Timezone = "America/Caracas"
	AmericaCayenne              Timezone = "America/Cayenne"
	AmericaChicago              Timezone = "America/Chicago"
	AmericaChihuahua            Timezone = "America/Chihuahua"
	AmericaCostaRica            Timezone = "America/Costa_Rica"
	AmericaCuiaba               Timezone = "America/Cuiaba"
	AmericaDanmarkshavn         Timezone = "America/Danmarkshavn"
	AmericaDawson               Timezone = "America/Dawson"
	AmericaDawsonCreek          Timezone = "America/Dawson_Creek"
	AmericaDenver               Timezone = "America/Denver"
	AmericaDetroit              Timezone = "America/Detroit"
	AmericaEdmonton             Timezone = "America/Edmonton"
	AmericaEirunepe             Timezone = "America/Eirunepe"
	AmericaElSalvador           Timezone = "America/El_Salvador"
	AmericaFortNelson           Timezone = "America/Fort_Nelson"
	AmericaFortaleza            Timezone = "America/Fortaleza"
	AmericaGlaceBay             Timezone = "America/Glace_Bay"
	AmericaGooseBay             Timezone = "America/Goose_Bay"
	AmericaGrandTurk            Timezone = "America/Grand_Turk"
	AmericaGuatemala            Timezone = "America/Guatemala"
	AmericaGuayaquil            Timezone = "America/Guayaquil"
	AmericaGuyana               Timezone = "America/Guyana"
	AmericaHalifax              Timezone = "America/Halifax"
	AmericaHavana               Timezone = "America/Havana"
	AmericaHermosillo           Timezone = "America/Hermosillo"
	AmericaIndianaIndianapolis  Timezone = "America/Indiana/Indianapolis"
	AmericaIndianaKnox          Timezone = "America/Indiana/Knox"
	AmericaIndianaMarengo       Timezone = "America/Indiana/Marengo"
	AmericaIndianaPetersburg    Timezone = "America/Indiana/Petersburg"
	AmericaIndianaTellCity      Timezone = "America/Indiana/Tell_City"
	AmericaIndianaVevay         Timezone = "America/Indiana/Vevay"
	AmericaIndianaVincennes     Timezone = "America/Indiana/Vincennes"
	AmericaIndianaWinamac       Timezone = "America/Indiana/Winamac"
	AmericaInuvik               Timezone = "America/Inuvik"
	AmericaIqaluit              Timezone = "America/Iqaluit"
	AmericaJamaica              Timezone = "America/Jamaica"
	AmericaJuneau               Timezone = "America/Juneau"
	AmericaKentuckyLouisville   Timezone = "America/Kentucky/Louisville"
	AmericaKentuckyMonticello   Timezone = "America/Kentucky/Monticello"
	AmericaLaPaz                Timezone = "America/La_Paz"
	AmericaLima                 Timezone = "America/Lima"
	AmericaLosAngeles           Timezone = "America/Los_Angeles"
	AmericaMaceio               Timezone = "America/Maceio"
	AmericaManagua              Timezone = "America/Managua"
	AmericaManaus               Timezone = "America/Manaus"
	AmericaMartinique           Timezone = "America/Martinique"
	AmericaMatamoros            Timezone = "America/Matamoros"
	AmericaMazatlan             Timezone = "America/Mazatlan"
	AmericaMenominee            Timezone = "America/Menominee"
	AmericaMerida               Timezone = "America/Merida"
	AmericaMetlakatla           Timezone = "America/Metlakatla"
	AmericaMexicoCity           Timezone = "America/Mexico_City"
	AmericaMiquelon             Timezone = "America/Miquelon"
	AmericaMoncton              Timezone = "America/Moncton"
	AmericaMonterrey            Timezone = "America/Monterrey"
	AmericaMontevideo           Timezone = "America/Montevideo"
	AmericaNewYork              Timezone = "America/New_York"
	AmericaNipigon              Timezone = "America/Nipigon"
	AmericaNome                 Timezone = "America/Nome"
	AmericaNoronha              Timezone = "America/Noronha"
	AmericaNorthDakotaBeulah    Timezone = "America/North_Dakota/Beulah"
	AmericaNorthDakotaCenter    Timezone = "America/North_Dakota/Center"
	AmericaNorthDakotaNewSalem  Timezone = "America/North_Dakota/New_Salem"
	AmericaNuuk                 Timezone = "America/Nuuk"
	AmericaOjinaga              Timezone = "America/Ojinaga"
	AmericaPanama               Timezone = "America/Panama"
	AmericaPangnirtung          Timezone = "America/Pangnirtung"
	AmericaParamaribo           Timezone = "America/Paramaribo"
	AmericaPhoenix              Timezone = "America/Phoenix"
	AmericaPortAuPrince         Timezone = "America/Port-au-Prince"
	AmericaPortoVelho           Timezone = "America/Porto_Velho"
	AmericaPuertoRico           Timezone = "America/Puerto_Rico"
	AmericaPuntaArenas          Timezone = "America/Punta_Arenas"
	AmericaRainyRiver           Timezone = "America/Rainy_River"
	AmericaRankinInlet          Timezone = "America/Rankin_Inlet"
	AmericaRecife               Timezone = "America/Recife"
	AmericaRegina               Timezone = "America/Regina"
	AmericaResolute             Timezone = "America/Resolute"
	AmericaRioBranco            Timezone = "America/Rio_Branco"
	AmericaSantarem             Timezone = "America/Santarem"
	AmericaSantiago             Timezone = "America/Santiago"
	AmericaSantoDomingo         Timezone = "America/Santo_Domingo"
	AmericaSaoPaulo             Timezone = "America/Sao_Paulo"
	AmericaScoresbysund         Timezone = "America/Scoresbysund"
	AmericaSitka                Timezone = "America/Sitka"
	AmericaStJohns              Timezone = "America/St_Johns"
	AmericaSwiftCurrent         Timezone = "America/Swift_Current"
	AmericaTegucigalpa          Timezone = "America/Tegucigalpa"
	AmericaThule                Timezone = "America/Thule"
	AmericaThunderBay           Timezone = "America/Thunder_Bay"
	AmericaTijuana              Timezone = "America/Tijuana"
	AmericaToronto              Timezone = "America/Toronto"
	AmericaVancouver            Timezone = "America/Vancouver"
	AmericaWhitehorse           Timezone = "America/Whitehorse"
	AmericaWinnipeg             Timezone = "America/Winnipeg"
	AmericaYakutat              Timezone = "America/Yakutat"
	AmericaYellowknife          Timezone = "America/Yellowknife"
	AntarcticaCasey             Timezone = "Antarctica/Casey"
	AntarcticaDavis             Timezone = "Antarctica/Davis"
	AntarcticaMacquarie         Timezone = "Antarctica/Macquarie"
	AntarcticaMawson            Timezone = "Antarctica/Mawson"
	AntarcticaPalmer            Timezone = "Antarctica/Palmer"
	AntarcticaRothera           Timezone = "Antarctica/Rothera"
	AntarcticaTroll             Timezone = "Antarctica/Troll"
	AntarcticaVostok            Timezone = "Antarctica/Vostok"
	AsiaAlmaty                  Timezone = "Asia/Almaty"
	AsiaAmman                   Timezone = "Asia/Amman"
	AsiaAnadyr                  Timezone = "Asia/Anadyr"
	AsiaAqtau                   Timezone = "Asia/Aqtau"
	AsiaAqtobe                  Timezone = "Asia/Aqtobe"
	AsiaAshgabat                Timezone = "Asia/Ashgabat"
	AsiaAtyrau                  Timezone = "Asia/Atyrau"
	AsiaBaghdad                 Timezone = "Asia/Baghdad"
	AsiaBaku                    Timezone = "Asia/Baku"
	AsiaBangkok                 Timezone = "Asia/Bangkok"
	AsiaBarnaul                 Timezone = "Asia/Barnaul"
	AsiaBeirut                  Timezone = "Asia/Beirut"
	AsiaBishkek                 Timezone = "Asia/Bishkek"
	AsiaBrunei                  Timezone = "Asia/Brunei"
	AsiaChita                   Timezone = "Asia/Chita"
	AsiaChoibalsan              Timezone = "Asia/Choibalsan"
	AsiaColombo                 Timezone = "Asia/Colombo"
	AsiaDamascus                Timezone = "Asia/Damascus"
	AsiaDhaka                   Timezone = "Asia/Dhaka"
	AsiaDili                    Timezone = "Asia/Dili"
	AsiaDubai                   Timezone = "Asia/Dubai"
	AsiaDushanbe                Timezone = "Asia/Dushanbe"
	AsiaFamagusta               Timezone = "Asia/Famagusta"
	AsiaGaza                    Timezone = "Asia/Gaza"
	AsiaHebron                  Timezone = "Asia/Hebron"
	AsiaHoChiMinh               Timezone = "Asia/Ho_Chi_Minh"
	AsiaHongKong                Timezone = "Asia/Hong_Kong"
	AsiaHovd                    Timezone = "Asia/Hovd"
	AsiaIrkutsk                 Timezone = "Asia/Irkutsk"
	AsiaJakarta                 Timezone = "Asia/Jakarta"
	AsiaJayapura                Timezone = "Asia/Jayapura"
	AsiaJerusalem               Timezone = "Asia/Jerusalem"
	AsiaKabul                   Timezone = "Asia/Kabul"
	AsiaKamchatka               Timezone = "Asia/Kamchatka"
	AsiaKarachi                 Timezone = "Asia/Karachi"
	AsiaKathmandu               Timezone = "Asia/Kathmandu"
	AsiaKhandyga                Timezone = "Asia/Khandyga"
	AsiaKolkata                 Timezone = "Asia/Kolkata"
	AsiaKrasnoyarsk             Timezone = "Asia/Krasnoyarsk"
	AsiaKualaLumpur             Timezone = "Asia/Kuala_Lumpur"
	AsiaKuching                 Timezone = "Asia/Kuching"
	AsiaMacau                   Timezone = "Asia/Macau"
	AsiaMagadan                 Timezone = "Asia/Magadan"
	AsiaMakassar                Timezone = "Asia/Makassar"
	AsiaManila                  Timezone = "Asia/Manila"
	AsiaNicosia                 Timezone = "Asia/Nicosia"
	AsiaNovokuznetsk            Timezone = "Asia/Novokuznetsk"
	AsiaNovosibirsk             Timezone = "Asia/Novosibirsk"
	AsiaOmsk                    Timezone = "Asia/Omsk"
	AsiaOral                    Timezone = "Asia/Oral"
	AsiaPontianak               Timezone = "Asia/Pontianak"
	AsiaPyongyang               Timezone = "Asia/Pyongyang"
	AsiaQatar                   Timezone = "Asia/Qatar"
	AsiaQostanay                Timezone = "Asia/Qostanay"
	AsiaQyzylorda               Timezone = "Asia/Qyzylorda"
	AsiaRiyadh                  Timezone = "Asia/Riyadh"
	AsiaSakhalin                Timezone = "Asia/Sakhalin"
	AsiaSamarkand               Timezone = "Asia/Samarkand"
	AsiaSeoul                   Timezone = "Asia/Seoul"
	AsiaShanghai                Timezone = "Asia/Shanghai"
	AsiaSingapore               Timezone = "Asia/Singapore"
	AsiaSrednekolymsk           Timezone = "Asia/Srednekolymsk"
	AsiaTaipei                  Timezone = "Asia/Taipei"
	AsiaTashkent                Timezone = "Asia/Tashkent"
	AsiaTbilisi                 Timezone = "Asia/Tbilisi"
	AsiaTehran                  Timezone = "Asia/Tehran"
	AsiaThimphu                 Timezone = "Asia/Thimphu"
	AsiaTokyo                   Timezone = "Asia/Tokyo"
	AsiaTomsk                   Timezone = "Asia/Tomsk"
	AsiaUlaanbaatar             Timezone = "Asia/Ulaanbaatar"
	AsiaUrumqi                  Timezone = "Asia/Urumqi"
	AsiaUstNera                 Timezone = "Asia/Ust-Nera"
	AsiaVladivostok             Timezone = "Asia/Vladivostok"
	AsiaYakutsk                 Timezone = "Asia/Yakutsk"
	AsiaYangon                  Timezone = "Asia/Yangon"
	AsiaYekaterinburg           Timezone = "Asia/Yekaterinburg"
	AsiaYerevan                 Timezone = "Asia/Yerevan"
	AtlanticAzores              Timezone = "Atlantic/Azores"
	AtlanticBermuda             Timezone = "Atlantic/Bermuda"
	AtlanticCanary              Timezone = "Atlantic/Canary"
	AtlanticCapeVerde           Timezone = "Atlantic/Cape_Verde"
	AtlanticFaroe               Timezone = "Atlantic/Faroe"
	AtlanticMadeira             Timezone = "Atlantic/Madeira"
	AtlanticReykjavik           Timezone = "Atlantic/Reykjavik"
	AtlanticSouthGeorgia        Timezone = "Atlantic/South_Georgia"
	AtlanticStanley             Timezone = "Atlantic/Stanley"
	AustraliaAdelaide           Timezone = "Australia/Adelaide"
	AustraliaBrisbane           Timezone = "Australia/Brisbane"
	AustraliaBrokenHill         Timezone = "Australia/Broken_Hill"
	AustraliaDarwin             Timezone = "Australia/Darwin"
	AustraliaEucla              Timezone = "Australia/Eucla"
	AustraliaHobart             Timezone = "Australia/Hobart"
	AustraliaLindeman           Timezone = "Australia/Lindeman"
	AustraliaLordHowe           Timezone = "Australia/Lord_Howe"
	AustraliaMelbourne          Timezone = "Australia/Melbourne"
	AustraliaPerth              Timezone = "Australia/Perth"
	AustraliaSydney             Timezone = "Australia/Sydney"
	Cet                         Timezone = "CET"
	Cst6Cdt                     Timezone = "CST6CDT"
	Eet                         Timezone = "EET"
	Est                         Timezone = "EST"
	Est5Edt                     Timezone = "EST5EDT"
	EtcGMT                      Timezone = "Etc/GMT"
	EtcGMT1                     Timezone = "Etc/GMT-1"
	EtcGMT10                    Timezone = "Etc/GMT-10"
	EtcGMT11                    Timezone = "Etc/GMT-11"
	EtcGMT12                    Timezone = "Etc/GMT-12"
	EtcGMT13                    Timezone = "Etc/GMT-13"
	EtcGMT14                    Timezone = "Etc/GMT-14"
	EtcGMT2                     Timezone = "Etc/GMT-2"
	EtcGMT3                     Timezone = "Etc/GMT-3"
	EtcGMT4                     Timezone = "Etc/GMT-4"
	EtcGMT5                     Timezone = "Etc/GMT-5"
	EtcGMT6                     Timezone = "Etc/GMT-6"
	EtcGMT7                     Timezone = "Etc/GMT-7"
	EtcGMT8                     Timezone = "Etc/GMT-8"
	EtcGMT9                     Timezone = "Etc/GMT-9"
	EtcUTC                      Timezone = "Etc/UTC"
	EuropeAmsterdam             Timezone = "Europe/Amsterdam"
	EuropeAndorra               Timezone = "Europe/Andorra"
	EuropeAstrakhan             Timezone = "Europe/Astrakhan"
	EuropeAthens                Timezone = "Europe/Athens"
	EuropeBelgrade              Timezone = "Europe/Belgrade"
	EuropeBerlin                Timezone = "Europe/Berlin"
	EuropeBrussels              Timezone = "Europe/Brussels"
	EuropeBucharest             Timezone = "Europe/Bucharest"
	EuropeBudapest              Timezone = "Europe/Budapest"
	EuropeChisinau              Timezone = "Europe/Chisinau"
	EuropeCopenhagen            Timezone = "Europe/Copenhagen"
	EuropeDublin                Timezone = "Europe/Dublin"
	EuropeGibraltar             Timezone = "Europe/Gibraltar"
	EuropeHelsinki              Timezone = "Europe/Helsinki"
	EuropeIstanbul              Timezone = "Europe/Istanbul"
	EuropeKaliningrad           Timezone = "Europe/Kaliningrad"
	EuropeKiev                  Timezone = "Europe/Kiev"
	EuropeKirov                 Timezone = "Europe/Kirov"
	EuropeLisbon                Timezone = "Europe/Lisbon"
	EuropeLondon                Timezone = "Europe/London"
	EuropeLuxembourg            Timezone = "Europe/Luxembourg"
	EuropeMadrid                Timezone = "Europe/Madrid"
	EuropeMalta                 Timezone = "Europe/Malta"
	EuropeMinsk                 Timezone = "Europe/Minsk"
	EuropeMonaco                Timezone = "Europe/Monaco"
	EuropeMoscow                Timezone = "Europe/Moscow"
	EuropeOslo                  Timezone = "Europe/Oslo"
	EuropeParis                 Timezone = "Europe/Paris"
	EuropePrague                Timezone = "Europe/Prague"
	EuropeRiga                  Timezone = "Europe/Riga"
	EuropeRome                  Timezone = "Europe/Rome"
	EuropeSamara                Timezone = "Europe/Samara"
	EuropeSaratov               Timezone = "Europe/Saratov"
	EuropeSimferopol            Timezone = "Europe/Simferopol"
	EuropeSofia                 Timezone = "Europe/Sofia"
	EuropeStockholm             Timezone = "Europe/Stockholm"
	EuropeTallinn               Timezone = "Europe/Tallinn"
	EuropeTirane                Timezone = "Europe/Tirane"
	EuropeUlyanovsk             Timezone = "Europe/Ulyanovsk"
	EuropeUzhgorod              Timezone = "Europe/Uzhgorod"
	EuropeVienna                Timezone = "Europe/Vienna"
	EuropeVilnius               Timezone = "Europe/Vilnius"
	EuropeVolgograd             Timezone = "Europe/Volgograd"
	EuropeWarsaw                Timezone = "Europe/Warsaw"
	EuropeZaporozhye            Timezone = "Europe/Zaporozhye"
	EuropeZurich                Timezone = "Europe/Zurich"
	Hst                         Timezone = "HST"
	IndianChagos                Timezone = "Indian/Chagos"
	IndianChristmas             Timezone = "Indian/Christmas"
	IndianCocos                 Timezone = "Indian/Cocos"
	IndianKerguelen             Timezone = "Indian/Kerguelen"
	IndianMahe                  Timezone = "Indian/Mahe"
	IndianMaldives              Timezone = "Indian/Maldives"
	IndianMauritius             Timezone = "Indian/Mauritius"
	IndianReunion               Timezone = "Indian/Reunion"
	Met                         Timezone = "MET"
	Mst                         Timezone = "MST"
	Mst7Mdt                     Timezone = "MST7MDT"
	PacificApia                 Timezone = "Pacific/Apia"
	PacificAuckland             Timezone = "Pacific/Auckland"
	PacificBougainville         Timezone = "Pacific/Bougainville"
	PacificChatham              Timezone = "Pacific/Chatham"
	PacificChuuk                Timezone = "Pacific/Chuuk"
	PacificEaster               Timezone = "Pacific/Easter"
	PacificEfate                Timezone = "Pacific/Efate"
	PacificFakaofo              Timezone = "Pacific/Fakaofo"
	PacificFiji                 Timezone = "Pacific/Fiji"
	PacificFunafuti             Timezone = "Pacific/Funafuti"
	PacificGalapagos            Timezone = "Pacific/Galapagos"
	PacificGambier              Timezone = "Pacific/Gambier"
	PacificGuadalcanal          Timezone = "Pacific/Guadalcanal"
	PacificGuam                 Timezone = "Pacific/Guam"
	PacificHonolulu             Timezone = "Pacific/Honolulu"
	PacificKanton               Timezone = "Pacific/Kanton"
	PacificKiritimati           Timezone = "Pacific/Kiritimati"
	PacificKosrae               Timezone = "Pacific/Kosrae"
	PacificKwajalein            Timezone = "Pacific/Kwajalein"
	PacificMajuro               Timezone = "Pacific/Majuro"
	PacificMarquesas            Timezone = "Pacific/Marquesas"
	PacificNauru                Timezone = "Pacific/Nauru"
	PacificNiue                 Timezone = "Pacific/Niue"
	PacificNorfolk              Timezone = "Pacific/Norfolk"
	PacificNoumea               Timezone = "Pacific/Noumea"
	PacificPagoPago             Timezone = "Pacific/Pago_Pago"
	PacificPalau                Timezone = "Pacific/Palau"
	PacificPitcairn             Timezone = "Pacific/Pitcairn"
	PacificPohnpei              Timezone = "Pacific/Pohnpei"
	PacificPortMoresby          Timezone = "Pacific/Port_Moresby"
	PacificRarotonga            Timezone = "Pacific/Rarotonga"
	PacificTahiti               Timezone = "Pacific/Tahiti"
	PacificTarawa               Timezone = "Pacific/Tarawa"
	PacificTongatapu            Timezone = "Pacific/Tongatapu"
	PacificWake                 Timezone = "Pacific/Wake"
	PacificWallis               Timezone = "Pacific/Wallis"
	Pst8Pdt                     Timezone = "PST8PDT"
	TimezoneEtcGMT1             Timezone = "Etc/GMT+1"
	TimezoneEtcGMT10            Timezone = "Etc/GMT+10"
	TimezoneEtcGMT11            Timezone = "Etc/GMT+11"
	TimezoneEtcGMT12            Timezone = "Etc/GMT+12"
	TimezoneEtcGMT2             Timezone = "Etc/GMT+2"
	TimezoneEtcGMT3             Timezone = "Etc/GMT+3"
	TimezoneEtcGMT4             Timezone = "Etc/GMT+4"
	TimezoneEtcGMT5             Timezone = "Etc/GMT+5"
	TimezoneEtcGMT6             Timezone = "Etc/GMT+6"
	TimezoneEtcGMT7             Timezone = "Etc/GMT+7"
	TimezoneEtcGMT8             Timezone = "Etc/GMT+8"
	TimezoneEtcGMT9             Timezone = "Etc/GMT+9"
	Wet                         Timezone = "WET"
)

type Version string

const (
	The30 Version = "3.0"
)
