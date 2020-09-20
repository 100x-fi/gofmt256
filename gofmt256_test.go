package gofmt256_test

import (
	"github.com/100x-fi/gofmt256"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBuildFrom(t *testing.T) {
	type fields struct {
		header interface{}
		body   interface{}
		footer interface{}
	}
	tests := []struct {
		name      string
		fields    fields
		want      string
		wantError bool
	}{
		{
			name: "when generate format 256 bytes with body successfully",
			fields: fields{
				header: getSubMerchantReportHeader(),
				body:   getSubMerchantReportBody(),
				footer: getSubMerchantReportFooter(),
			},
			want:      "H0000018888888888888100000000X                              03092020                                                                                                                                                                                            \nD000002888888888888803092020100337John Doe                                          7777777             7777777777777       0000000000000000000000000000CETH00000000000000051500                                                                                \nD000003888888888888803092020100739John Doe                                          8888888             8888888888888       0000000000000000000000000000CETH00000000000000746000                                                                                \nD000004888888888888803092020101056John Doe                                          9999999             9999999999999       0000000000000000000000000000CETH00000000000004880700                                                                                \nT000005888888888888800000000000000000000000005678200000003                                                                                                                                                                                                      \n",
			wantError: false,
		},
		{
			name: "when generate format 256 bytes without body successfully",
			fields: fields{
				header: SubMerchantReportHeader{
					RecordType:     "H",
					SequenceNo:     1,
					BankCode:       "888",
					CompanyAccount: "0000000000",
					CompanyName:    "100000000000000X",
					EffectiveDate:  "18082020",
					ServiceCode:    "",
					Spare:          "",
				},
				body: []SubMerchantReportBody{},
				footer: SubMerchantReportFooter{
					RecordType:             "T",
					SequenceNo:             9,
					BankCode:               "888",
					CompanyAccount:         "0000000000",
					TotalDebitAmount:       "0000",
					TotalDebitTransaction:  0,
					TotalCreditAmount:      "0000",
					TotalCreditTransaction: 0,
					Spare:                  "",
				},
			},
			want:      "H0000018880000000000100000000000000X                        18082020                                                                                                                                                                                            \nT000009888000000000000000000000000000000000000000000000000                                                                                                                                                                                                      \n",
			wantError: false,
		},
		{
			name: "when header is not a struct",
			fields: fields{
				header: "not_struct_for_sure",
				body:   getSubMerchantReportBody(),
				footer: getSubMerchantReportFooter(),
			},
			want:      "",
			wantError: true,
		},
		{
			name: "when body is not a slice",
			fields: fields{
				header: getSubMerchantReportHeader(),
				body:   "not_a_slice_for_sure",
				footer: getSubMerchantReportFooter(),
			},
			want:      "",
			wantError: true,
		},
		{
			name: "when footer is not a struct",
			fields: fields{
				header: getSubMerchantReportHeader(),
				body:   getSubMerchantReportBody(),
				footer: "not_a_struct_for_sure",
			},
			want:      "",
			wantError: true,
		},
		{
			name: "when there is a conflict between fields in header",
			fields: fields{
				header: ConflictMock{
					RecordType:     "",
					SequenceNo:     0,
					BankCode:       "",
					CompanyAccount: "",
					CompanyName:    "",
					EffectiveDate:  "",
					ServiceCode:    "",
					Spare:          "",
				},
				body:   getSubMerchantReportBody(),
				footer: getSubMerchantReportFooter(),
			},
			wantError: true,
			want:      "",
		},
		{
			name: "when there is a conflict between fields in body",
			fields: fields{
				header: getSubMerchantReportHeader(),
				body: []ConflictMock{{
					RecordType:     "",
					SequenceNo:     0,
					BankCode:       "",
					CompanyAccount: "",
					CompanyName:    "",
					EffectiveDate:  "",
					ServiceCode:    "",
					Spare:          "",
				}},
				footer: getSubMerchantReportFooter(),
			},
			wantError: true,
			want:      "",
		},
		{
			name: "when there is a conflict between fields in footer",
			fields: fields{
				header: getSubMerchantReportHeader(),
				body:   getSubMerchantReportBody(),
				footer: ConflictMock{
					RecordType:     "",
					SequenceNo:     0,
					BankCode:       "",
					CompanyAccount: "",
					CompanyName:    "",
					EffectiveDate:  "",
					ServiceCode:    "",
					Spare:          "",
				},
			},
			wantError: true,
			want:      "",
		},
		{
			name: "when `from` tag in any field not int",
			fields: fields{
				header: FromTagIsNotInt{
					RecordType:     "",
					SequenceNo:     0,
					BankCode:       "",
					CompanyAccount: "",
					CompanyName:    "",
					EffectiveDate:  "",
					ServiceCode:    "",
					Spare:          "",
				},
				body:   getSubMerchantReportBody(),
				footer: getSubMerchantReportFooter(),
			},
			want:      "",
			wantError: true,
		},
		{
			name: "when `to` tag in any field not int",
			fields: fields{
				header: ToTagIsNotInt{
					RecordType:     "",
					SequenceNo:     0,
					BankCode:       "",
					CompanyAccount: "",
					CompanyName:    "",
					EffectiveDate:  "",
					ServiceCode:    "",
					Spare:          "",
				},
				body:   getSubMerchantReportBody(),
				footer: getSubMerchantReportFooter(),
			},
			want:      "",
			wantError: true,
		},
		{
			name: "when any tag in any field doesn't has right hand value",
			fields: fields{
				header: TagNoRightHandValue{
					RecordType:     "",
					SequenceNo:     0,
					BankCode:       "",
					CompanyAccount: "",
					CompanyName:    "",
					EffectiveDate:  "",
					ServiceCode:    "",
					Spare:          "",
				},
				body:   getSubMerchantReportBody(),
				footer: getSubMerchantReportFooter(),
			},
			want:      "",
			wantError: true,
		},
		{
			name: "when any sub tag has no equal sign",
			fields: fields{
				header: SubTagNoEqual{
					RecordType:     "",
					SequenceNo:     0,
					BankCode:       "",
					CompanyAccount: "",
					CompanyName:    "",
					EffectiveDate:  "",
					ServiceCode:    "",
					Spare:          "",
				},
				body:   getSubMerchantReportBody(),
				footer: getSubMerchantReportFooter(),
			},
			want:      "",
			wantError: true,
		},
		{
			name: "when `from` of any sub tag more than `to`",
			fields: fields{
				header: FromMoreThanTo{
					RecordType:     "",
					SequenceNo:     0,
					BankCode:       "",
					CompanyAccount: "",
					CompanyName:    "",
					EffectiveDate:  "",
					ServiceCode:    "",
					Spare:          "",
				},
				body:   getSubMerchantReportBody(),
				footer: getSubMerchantReportFooter(),
			},
			want:      "",
			wantError: true,
		},
		{
			name: "when `from` of any sub tag is missing",
			fields: fields{
				header: MissingFrom{
					RecordType:     "",
					SequenceNo:     0,
					BankCode:       "",
					CompanyAccount: "",
					CompanyName:    "",
					EffectiveDate:  "",
					ServiceCode:    "",
					Spare:          "",
				},
				body:   getSubMerchantReportBody(),
				footer: getSubMerchantReportFooter(),
			},
			want:      "",
			wantError: true,
		},
		{
			name: "when `to` of any sub tag is missing",
			fields: fields{
				header: MissingTo{
					RecordType:     "",
					SequenceNo:     0,
					BankCode:       "",
					CompanyAccount: "",
					CompanyName:    "",
					EffectiveDate:  "",
					ServiceCode:    "",
					Spare:          "",
				},
				body:   getSubMerchantReportBody(),
				footer: getSubMerchantReportFooter(),
			},
			want:      "",
			wantError: true,
		},
		{
			name: "when `from` is minus",
			fields: fields{
				header: MinusFrom{
					RecordType:     "",
					SequenceNo:     0,
					BankCode:       "",
					CompanyAccount: "",
					CompanyName:    "",
					EffectiveDate:  "",
					ServiceCode:    "",
					Spare:          "",
				},
				body:   getSubMerchantReportBody(),
				footer: getSubMerchantReportFooter(),
			},
			want:      "",
			wantError: true,
		},
		{
			name: "when `to` is minus",
			fields: fields{
				header: MinusTo{
					RecordType:     "",
					SequenceNo:     0,
					BankCode:       "",
					CompanyAccount: "",
					CompanyName:    "",
					EffectiveDate:  "",
					ServiceCode:    "",
					Spare:          "",
				},
				body:   getSubMerchantReportBody(),
				footer: getSubMerchantReportFooter(),
			},
			want:      "",
			wantError: true,
		},
		{
			name: "when `from` more than 256",
			fields: fields{
				header: FromMoreThan256{
					RecordType:     "",
					SequenceNo:     0,
					BankCode:       "",
					CompanyAccount: "",
					CompanyName:    "",
					EffectiveDate:  "",
					ServiceCode:    "",
					Spare:          "",
				},
				body:   getSubMerchantReportBody(),
				footer: getSubMerchantReportFooter(),
			},
			want:      "",
			wantError: true,
		},
		{
			name: "when data is longer than allocated length",
			fields: fields{
				header: SubMerchantReportHeader{
					RecordType:     "AAAAAA",
					SequenceNo:     0,
					BankCode:       "",
					CompanyAccount: "",
					CompanyName:    "",
					EffectiveDate:  "",
					ServiceCode:    "",
					Spare:          "",
				},
				body:   getSubMerchantReportBody(),
				footer: getSubMerchantReportFooter(),
			},
			want:      "",
			wantError: true,
		},
		{
			name: "when not fill 256",
			fields: fields{
				header: NotFill256{
					RecordType:     "",
					SequenceNo:     0,
					BankCode:       "",
					CompanyAccount: "",
					CompanyName:    "",
					EffectiveDate:  "",
					ServiceCode:    "",
					Spare:          "",
				},
				body:   getSubMerchantReportBody(),
				footer: getSubMerchantReportFooter(),
			},
			want:      "",
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			builder := gofmt256.New(tt.fields.header, tt.fields.body, tt.fields.footer)
			got, err := builder.Build()
			if (err != nil) != tt.wantError {
				t.Errorf("gofmt256.Build() err %v, wantErr %v", err, tt.wantError)
			}
			assert.Equal(t, tt.want, got)
		})
	}
}
