
//line machine.rl:1
package influx

import (
	"fmt"
)

var (
	ErrFieldParse = fmt.Errorf("expected field")
	ErrTagParse = fmt.Errorf("expected tag")
	ErrTimestampParse = fmt.Errorf("expected timestamp")
	ErrParse = fmt.Errorf("parse error")
)


//line machine.go:18
const LineProtocol_start int = 490
const LineProtocol_first_final int = 490
const LineProtocol_error int = 0

const LineProtocol_en_nextline int = 489
const LineProtocol_en_main int = 490


//line machine.rl:179


type machine struct {
	data       []byte
	cs         int
	p, pe, eof int
	pb         int
	handler    Handler
	err        error

	count int
}

func NewMachine(handler Handler) *machine {
	m := &machine{
		handler: handler,
	}

	
//line machine.rl:198
	
//line machine.rl:199
	
//line machine.rl:200
	
//line machine.rl:201
	
//line machine.rl:202
	
//line machine.go:57
	{
	 m.cs = LineProtocol_start
	}

//line machine.rl:203

	return m
}

func (m *machine) SetData(data []byte) {
	m.data = data
	m.p = 0
	m.pb = 0
	m.pe = len(data)
	m.eof = len(data)
	m.cs = 0
	m.err = nil

	
//line machine.go:77
	{
	 m.cs = LineProtocol_start
	}

//line machine.rl:217
}

// ParseLine parses a line of input and returns true if more data can be
// parsed.
func (m *machine) ParseLine() bool {
	if m.data == nil || m.p >= m.pe {
		return false
	}

	m.err = nil
	var key []byte
	var yield bool

	
//line machine.go:97
	{
	if ( m.p) == ( m.pe) {
		goto _test_eof
	}
	goto _resume

_again:
	switch  m.cs {
	case 490:
		goto st490
	case 1:
		goto st1
	case 2:
		goto st2
	case 3:
		goto st3
	case 0:
		goto st0
	case 4:
		goto st4
	case 5:
		goto st5
	case 6:
		goto st6
	case 491:
		goto st491
	case 492:
		goto st492
	case 493:
		goto st493
	case 7:
		goto st7
	case 8:
		goto st8
	case 9:
		goto st9
	case 10:
		goto st10
	case 11:
		goto st11
	case 12:
		goto st12
	case 13:
		goto st13
	case 14:
		goto st14
	case 15:
		goto st15
	case 16:
		goto st16
	case 17:
		goto st17
	case 18:
		goto st18
	case 19:
		goto st19
	case 20:
		goto st20
	case 21:
		goto st21
	case 22:
		goto st22
	case 23:
		goto st23
	case 24:
		goto st24
	case 25:
		goto st25
	case 494:
		goto st494
	case 495:
		goto st495
	case 26:
		goto st26
	case 496:
		goto st496
	case 497:
		goto st497
	case 498:
		goto st498
	case 27:
		goto st27
	case 499:
		goto st499
	case 500:
		goto st500
	case 501:
		goto st501
	case 502:
		goto st502
	case 503:
		goto st503
	case 504:
		goto st504
	case 505:
		goto st505
	case 506:
		goto st506
	case 507:
		goto st507
	case 508:
		goto st508
	case 509:
		goto st509
	case 510:
		goto st510
	case 511:
		goto st511
	case 512:
		goto st512
	case 513:
		goto st513
	case 514:
		goto st514
	case 515:
		goto st515
	case 516:
		goto st516
	case 28:
		goto st28
	case 29:
		goto st29
	case 517:
		goto st517
	case 518:
		goto st518
	case 519:
		goto st519
	case 30:
		goto st30
	case 31:
		goto st31
	case 32:
		goto st32
	case 33:
		goto st33
	case 34:
		goto st34
	case 520:
		goto st520
	case 521:
		goto st521
	case 522:
		goto st522
	case 523:
		goto st523
	case 35:
		goto st35
	case 524:
		goto st524
	case 525:
		goto st525
	case 526:
		goto st526
	case 527:
		goto st527
	case 36:
		goto st36
	case 528:
		goto st528
	case 529:
		goto st529
	case 530:
		goto st530
	case 531:
		goto st531
	case 532:
		goto st532
	case 533:
		goto st533
	case 534:
		goto st534
	case 535:
		goto st535
	case 536:
		goto st536
	case 537:
		goto st537
	case 538:
		goto st538
	case 539:
		goto st539
	case 540:
		goto st540
	case 541:
		goto st541
	case 542:
		goto st542
	case 543:
		goto st543
	case 544:
		goto st544
	case 545:
		goto st545
	case 37:
		goto st37
	case 38:
		goto st38
	case 39:
		goto st39
	case 40:
		goto st40
	case 41:
		goto st41
	case 42:
		goto st42
	case 43:
		goto st43
	case 44:
		goto st44
	case 45:
		goto st45
	case 546:
		goto st546
	case 547:
		goto st547
	case 548:
		goto st548
	case 46:
		goto st46
	case 47:
		goto st47
	case 48:
		goto st48
	case 49:
		goto st49
	case 50:
		goto st50
	case 51:
		goto st51
	case 549:
		goto st549
	case 550:
		goto st550
	case 52:
		goto st52
	case 551:
		goto st551
	case 552:
		goto st552
	case 553:
		goto st553
	case 554:
		goto st554
	case 555:
		goto st555
	case 556:
		goto st556
	case 557:
		goto st557
	case 558:
		goto st558
	case 559:
		goto st559
	case 560:
		goto st560
	case 561:
		goto st561
	case 562:
		goto st562
	case 563:
		goto st563
	case 564:
		goto st564
	case 565:
		goto st565
	case 566:
		goto st566
	case 567:
		goto st567
	case 568:
		goto st568
	case 569:
		goto st569
	case 570:
		goto st570
	case 53:
		goto st53
	case 571:
		goto st571
	case 572:
		goto st572
	case 573:
		goto st573
	case 54:
		goto st54
	case 574:
		goto st574
	case 575:
		goto st575
	case 576:
		goto st576
	case 577:
		goto st577
	case 578:
		goto st578
	case 579:
		goto st579
	case 580:
		goto st580
	case 581:
		goto st581
	case 582:
		goto st582
	case 583:
		goto st583
	case 584:
		goto st584
	case 585:
		goto st585
	case 586:
		goto st586
	case 587:
		goto st587
	case 588:
		goto st588
	case 589:
		goto st589
	case 590:
		goto st590
	case 591:
		goto st591
	case 592:
		goto st592
	case 593:
		goto st593
	case 55:
		goto st55
	case 56:
		goto st56
	case 57:
		goto st57
	case 594:
		goto st594
	case 595:
		goto st595
	case 596:
		goto st596
	case 58:
		goto st58
	case 597:
		goto st597
	case 598:
		goto st598
	case 59:
		goto st59
	case 599:
		goto st599
	case 60:
		goto st60
	case 600:
		goto st600
	case 601:
		goto st601
	case 602:
		goto st602
	case 603:
		goto st603
	case 604:
		goto st604
	case 605:
		goto st605
	case 606:
		goto st606
	case 607:
		goto st607
	case 608:
		goto st608
	case 609:
		goto st609
	case 610:
		goto st610
	case 611:
		goto st611
	case 612:
		goto st612
	case 613:
		goto st613
	case 614:
		goto st614
	case 615:
		goto st615
	case 616:
		goto st616
	case 617:
		goto st617
	case 618:
		goto st618
	case 619:
		goto st619
	case 61:
		goto st61
	case 62:
		goto st62
	case 63:
		goto st63
	case 64:
		goto st64
	case 65:
		goto st65
	case 66:
		goto st66
	case 67:
		goto st67
	case 620:
		goto st620
	case 68:
		goto st68
	case 69:
		goto st69
	case 70:
		goto st70
	case 71:
		goto st71
	case 621:
		goto st621
	case 622:
		goto st622
	case 623:
		goto st623
	case 624:
		goto st624
	case 72:
		goto st72
	case 73:
		goto st73
	case 74:
		goto st74
	case 75:
		goto st75
	case 625:
		goto st625
	case 626:
		goto st626
	case 76:
		goto st76
	case 627:
		goto st627
	case 77:
		goto st77
	case 78:
		goto st78
	case 628:
		goto st628
	case 629:
		goto st629
	case 79:
		goto st79
	case 630:
		goto st630
	case 631:
		goto st631
	case 80:
		goto st80
	case 632:
		goto st632
	case 633:
		goto st633
	case 634:
		goto st634
	case 635:
		goto st635
	case 636:
		goto st636
	case 637:
		goto st637
	case 638:
		goto st638
	case 639:
		goto st639
	case 640:
		goto st640
	case 641:
		goto st641
	case 642:
		goto st642
	case 643:
		goto st643
	case 644:
		goto st644
	case 645:
		goto st645
	case 646:
		goto st646
	case 647:
		goto st647
	case 648:
		goto st648
	case 649:
		goto st649
	case 81:
		goto st81
	case 650:
		goto st650
	case 651:
		goto st651
	case 652:
		goto st652
	case 82:
		goto st82
	case 653:
		goto st653
	case 654:
		goto st654
	case 655:
		goto st655
	case 83:
		goto st83
	case 656:
		goto st656
	case 657:
		goto st657
	case 658:
		goto st658
	case 659:
		goto st659
	case 660:
		goto st660
	case 661:
		goto st661
	case 662:
		goto st662
	case 663:
		goto st663
	case 664:
		goto st664
	case 665:
		goto st665
	case 666:
		goto st666
	case 667:
		goto st667
	case 668:
		goto st668
	case 669:
		goto st669
	case 670:
		goto st670
	case 671:
		goto st671
	case 672:
		goto st672
	case 673:
		goto st673
	case 674:
		goto st674
	case 84:
		goto st84
	case 85:
		goto st85
	case 86:
		goto st86
	case 675:
		goto st675
	case 87:
		goto st87
	case 676:
		goto st676
	case 677:
		goto st677
	case 678:
		goto st678
	case 679:
		goto st679
	case 680:
		goto st680
	case 681:
		goto st681
	case 682:
		goto st682
	case 683:
		goto st683
	case 684:
		goto st684
	case 685:
		goto st685
	case 686:
		goto st686
	case 687:
		goto st687
	case 688:
		goto st688
	case 689:
		goto st689
	case 690:
		goto st690
	case 691:
		goto st691
	case 692:
		goto st692
	case 693:
		goto st693
	case 694:
		goto st694
	case 695:
		goto st695
	case 696:
		goto st696
	case 697:
		goto st697
	case 88:
		goto st88
	case 89:
		goto st89
	case 90:
		goto st90
	case 91:
		goto st91
	case 92:
		goto st92
	case 698:
		goto st698
	case 699:
		goto st699
	case 700:
		goto st700
	case 701:
		goto st701
	case 702:
		goto st702
	case 93:
		goto st93
	case 94:
		goto st94
	case 95:
		goto st95
	case 703:
		goto st703
	case 96:
		goto st96
	case 97:
		goto st97
	case 98:
		goto st98
	case 704:
		goto st704
	case 99:
		goto st99
	case 100:
		goto st100
	case 705:
		goto st705
	case 706:
		goto st706
	case 101:
		goto st101
	case 102:
		goto st102
	case 707:
		goto st707
	case 708:
		goto st708
	case 709:
		goto st709
	case 103:
		goto st103
	case 710:
		goto st710
	case 711:
		goto st711
	case 712:
		goto st712
	case 713:
		goto st713
	case 714:
		goto st714
	case 715:
		goto st715
	case 716:
		goto st716
	case 717:
		goto st717
	case 718:
		goto st718
	case 719:
		goto st719
	case 720:
		goto st720
	case 721:
		goto st721
	case 722:
		goto st722
	case 723:
		goto st723
	case 724:
		goto st724
	case 725:
		goto st725
	case 726:
		goto st726
	case 727:
		goto st727
	case 728:
		goto st728
	case 729:
		goto st729
	case 730:
		goto st730
	case 731:
		goto st731
	case 732:
		goto st732
	case 733:
		goto st733
	case 104:
		goto st104
	case 105:
		goto st105
	case 106:
		goto st106
	case 734:
		goto st734
	case 107:
		goto st107
	case 108:
		goto st108
	case 109:
		goto st109
	case 735:
		goto st735
	case 110:
		goto st110
	case 111:
		goto st111
	case 736:
		goto st736
	case 737:
		goto st737
	case 112:
		goto st112
	case 738:
		goto st738
	case 113:
		goto st113
	case 739:
		goto st739
	case 740:
		goto st740
	case 741:
		goto st741
	case 114:
		goto st114
	case 115:
		goto st115
	case 116:
		goto st116
	case 742:
		goto st742
	case 117:
		goto st117
	case 118:
		goto st118
	case 119:
		goto st119
	case 743:
		goto st743
	case 120:
		goto st120
	case 121:
		goto st121
	case 744:
		goto st744
	case 745:
		goto st745
	case 746:
		goto st746
	case 747:
		goto st747
	case 748:
		goto st748
	case 749:
		goto st749
	case 750:
		goto st750
	case 751:
		goto st751
	case 752:
		goto st752
	case 753:
		goto st753
	case 754:
		goto st754
	case 755:
		goto st755
	case 756:
		goto st756
	case 757:
		goto st757
	case 758:
		goto st758
	case 759:
		goto st759
	case 760:
		goto st760
	case 761:
		goto st761
	case 762:
		goto st762
	case 763:
		goto st763
	case 122:
		goto st122
	case 764:
		goto st764
	case 765:
		goto st765
	case 766:
		goto st766
	case 123:
		goto st123
	case 767:
		goto st767
	case 768:
		goto st768
	case 124:
		goto st124
	case 769:
		goto st769
	case 770:
		goto st770
	case 771:
		goto st771
	case 772:
		goto st772
	case 773:
		goto st773
	case 774:
		goto st774
	case 775:
		goto st775
	case 776:
		goto st776
	case 777:
		goto st777
	case 778:
		goto st778
	case 779:
		goto st779
	case 780:
		goto st780
	case 781:
		goto st781
	case 782:
		goto st782
	case 783:
		goto st783
	case 784:
		goto st784
	case 785:
		goto st785
	case 786:
		goto st786
	case 787:
		goto st787
	case 125:
		goto st125
	case 788:
		goto st788
	case 789:
		goto st789
	case 790:
		goto st790
	case 126:
		goto st126
	case 127:
		goto st127
	case 128:
		goto st128
	case 791:
		goto st791
	case 129:
		goto st129
	case 130:
		goto st130
	case 131:
		goto st131
	case 792:
		goto st792
	case 132:
		goto st132
	case 133:
		goto st133
	case 793:
		goto st793
	case 794:
		goto st794
	case 134:
		goto st134
	case 135:
		goto st135
	case 136:
		goto st136
	case 137:
		goto st137
	case 138:
		goto st138
	case 139:
		goto st139
	case 795:
		goto st795
	case 796:
		goto st796
	case 797:
		goto st797
	case 798:
		goto st798
	case 799:
		goto st799
	case 800:
		goto st800
	case 801:
		goto st801
	case 802:
		goto st802
	case 803:
		goto st803
	case 804:
		goto st804
	case 805:
		goto st805
	case 806:
		goto st806
	case 807:
		goto st807
	case 808:
		goto st808
	case 809:
		goto st809
	case 810:
		goto st810
	case 811:
		goto st811
	case 812:
		goto st812
	case 813:
		goto st813
	case 140:
		goto st140
	case 141:
		goto st141
	case 142:
		goto st142
	case 814:
		goto st814
	case 815:
		goto st815
	case 816:
		goto st816
	case 817:
		goto st817
	case 818:
		goto st818
	case 143:
		goto st143
	case 144:
		goto st144
	case 145:
		goto st145
	case 819:
		goto st819
	case 146:
		goto st146
	case 147:
		goto st147
	case 148:
		goto st148
	case 820:
		goto st820
	case 149:
		goto st149
	case 150:
		goto st150
	case 821:
		goto st821
	case 822:
		goto st822
	case 151:
		goto st151
	case 152:
		goto st152
	case 153:
		goto st153
	case 823:
		goto st823
	case 824:
		goto st824
	case 825:
		goto st825
	case 826:
		goto st826
	case 827:
		goto st827
	case 154:
		goto st154
	case 155:
		goto st155
	case 156:
		goto st156
	case 828:
		goto st828
	case 157:
		goto st157
	case 158:
		goto st158
	case 159:
		goto st159
	case 829:
		goto st829
	case 160:
		goto st160
	case 161:
		goto st161
	case 830:
		goto st830
	case 831:
		goto st831
	case 162:
		goto st162
	case 163:
		goto st163
	case 164:
		goto st164
	case 832:
		goto st832
	case 833:
		goto st833
	case 834:
		goto st834
	case 835:
		goto st835
	case 836:
		goto st836
	case 837:
		goto st837
	case 838:
		goto st838
	case 839:
		goto st839
	case 840:
		goto st840
	case 841:
		goto st841
	case 842:
		goto st842
	case 843:
		goto st843
	case 844:
		goto st844
	case 845:
		goto st845
	case 846:
		goto st846
	case 847:
		goto st847
	case 848:
		goto st848
	case 849:
		goto st849
	case 850:
		goto st850
	case 165:
		goto st165
	case 166:
		goto st166
	case 167:
		goto st167
	case 168:
		goto st168
	case 851:
		goto st851
	case 852:
		goto st852
	case 853:
		goto st853
	case 854:
		goto st854
	case 855:
		goto st855
	case 169:
		goto st169
	case 170:
		goto st170
	case 171:
		goto st171
	case 856:
		goto st856
	case 172:
		goto st172
	case 173:
		goto st173
	case 174:
		goto st174
	case 857:
		goto st857
	case 175:
		goto st175
	case 176:
		goto st176
	case 858:
		goto st858
	case 859:
		goto st859
	case 177:
		goto st177
	case 178:
		goto st178
	case 179:
		goto st179
	case 860:
		goto st860
	case 861:
		goto st861
	case 862:
		goto st862
	case 863:
		goto st863
	case 864:
		goto st864
	case 180:
		goto st180
	case 181:
		goto st181
	case 182:
		goto st182
	case 865:
		goto st865
	case 183:
		goto st183
	case 184:
		goto st184
	case 185:
		goto st185
	case 866:
		goto st866
	case 186:
		goto st186
	case 187:
		goto st187
	case 867:
		goto st867
	case 868:
		goto st868
	case 188:
		goto st188
	case 189:
		goto st189
	case 190:
		goto st190
	case 191:
		goto st191
	case 192:
		goto st192
	case 193:
		goto st193
	case 869:
		goto st869
	case 870:
		goto st870
	case 871:
		goto st871
	case 872:
		goto st872
	case 873:
		goto st873
	case 874:
		goto st874
	case 875:
		goto st875
	case 876:
		goto st876
	case 877:
		goto st877
	case 878:
		goto st878
	case 879:
		goto st879
	case 880:
		goto st880
	case 881:
		goto st881
	case 882:
		goto st882
	case 883:
		goto st883
	case 884:
		goto st884
	case 885:
		goto st885
	case 886:
		goto st886
	case 887:
		goto st887
	case 194:
		goto st194
	case 888:
		goto st888
	case 889:
		goto st889
	case 890:
		goto st890
	case 891:
		goto st891
	case 892:
		goto st892
	case 195:
		goto st195
	case 196:
		goto st196
	case 197:
		goto st197
	case 893:
		goto st893
	case 198:
		goto st198
	case 199:
		goto st199
	case 200:
		goto st200
	case 894:
		goto st894
	case 201:
		goto st201
	case 202:
		goto st202
	case 895:
		goto st895
	case 896:
		goto st896
	case 203:
		goto st203
	case 204:
		goto st204
	case 205:
		goto st205
	case 206:
		goto st206
	case 207:
		goto st207
	case 208:
		goto st208
	case 209:
		goto st209
	case 210:
		goto st210
	case 211:
		goto st211
	case 212:
		goto st212
	case 897:
		goto st897
	case 898:
		goto st898
	case 899:
		goto st899
	case 213:
		goto st213
	case 214:
		goto st214
	case 215:
		goto st215
	case 900:
		goto st900
	case 901:
		goto st901
	case 216:
		goto st216
	case 902:
		goto st902
	case 903:
		goto st903
	case 904:
		goto st904
	case 905:
		goto st905
	case 906:
		goto st906
	case 907:
		goto st907
	case 908:
		goto st908
	case 909:
		goto st909
	case 910:
		goto st910
	case 911:
		goto st911
	case 912:
		goto st912
	case 913:
		goto st913
	case 914:
		goto st914
	case 915:
		goto st915
	case 916:
		goto st916
	case 917:
		goto st917
	case 918:
		goto st918
	case 919:
		goto st919
	case 920:
		goto st920
	case 217:
		goto st217
	case 921:
		goto st921
	case 922:
		goto st922
	case 218:
		goto st218
	case 219:
		goto st219
	case 220:
		goto st220
	case 221:
		goto st221
	case 923:
		goto st923
	case 222:
		goto st222
	case 223:
		goto st223
	case 224:
		goto st224
	case 924:
		goto st924
	case 925:
		goto st925
	case 225:
		goto st225
	case 226:
		goto st226
	case 227:
		goto st227
	case 228:
		goto st228
	case 229:
		goto st229
	case 926:
		goto st926
	case 927:
		goto st927
	case 928:
		goto st928
	case 929:
		goto st929
	case 930:
		goto st930
	case 230:
		goto st230
	case 231:
		goto st231
	case 232:
		goto st232
	case 931:
		goto st931
	case 233:
		goto st233
	case 234:
		goto st234
	case 235:
		goto st235
	case 932:
		goto st932
	case 236:
		goto st236
	case 237:
		goto st237
	case 933:
		goto st933
	case 934:
		goto st934
	case 238:
		goto st238
	case 239:
		goto st239
	case 935:
		goto st935
	case 936:
		goto st936
	case 937:
		goto st937
	case 938:
		goto st938
	case 939:
		goto st939
	case 240:
		goto st240
	case 241:
		goto st241
	case 242:
		goto st242
	case 940:
		goto st940
	case 243:
		goto st243
	case 244:
		goto st244
	case 245:
		goto st245
	case 941:
		goto st941
	case 246:
		goto st246
	case 247:
		goto st247
	case 942:
		goto st942
	case 943:
		goto st943
	case 248:
		goto st248
	case 249:
		goto st249
	case 944:
		goto st944
	case 945:
		goto st945
	case 946:
		goto st946
	case 947:
		goto st947
	case 948:
		goto st948
	case 250:
		goto st250
	case 251:
		goto st251
	case 252:
		goto st252
	case 949:
		goto st949
	case 253:
		goto st253
	case 254:
		goto st254
	case 255:
		goto st255
	case 950:
		goto st950
	case 256:
		goto st256
	case 257:
		goto st257
	case 951:
		goto st951
	case 952:
		goto st952
	case 258:
		goto st258
	case 259:
		goto st259
	case 260:
		goto st260
	case 261:
		goto st261
	case 262:
		goto st262
	case 263:
		goto st263
	case 953:
		goto st953
	case 954:
		goto st954
	case 955:
		goto st955
	case 956:
		goto st956
	case 264:
		goto st264
	case 957:
		goto st957
	case 958:
		goto st958
	case 959:
		goto st959
	case 960:
		goto st960
	case 961:
		goto st961
	case 962:
		goto st962
	case 963:
		goto st963
	case 964:
		goto st964
	case 965:
		goto st965
	case 966:
		goto st966
	case 967:
		goto st967
	case 968:
		goto st968
	case 969:
		goto st969
	case 970:
		goto st970
	case 971:
		goto st971
	case 972:
		goto st972
	case 973:
		goto st973
	case 974:
		goto st974
	case 975:
		goto st975
	case 976:
		goto st976
	case 265:
		goto st265
	case 266:
		goto st266
	case 267:
		goto st267
	case 268:
		goto st268
	case 269:
		goto st269
	case 270:
		goto st270
	case 271:
		goto st271
	case 272:
		goto st272
	case 977:
		goto st977
	case 978:
		goto st978
	case 273:
		goto st273
	case 979:
		goto st979
	case 980:
		goto st980
	case 981:
		goto st981
	case 982:
		goto st982
	case 983:
		goto st983
	case 984:
		goto st984
	case 985:
		goto st985
	case 986:
		goto st986
	case 987:
		goto st987
	case 988:
		goto st988
	case 989:
		goto st989
	case 990:
		goto st990
	case 991:
		goto st991
	case 992:
		goto st992
	case 993:
		goto st993
	case 994:
		goto st994
	case 995:
		goto st995
	case 996:
		goto st996
	case 997:
		goto st997
	case 274:
		goto st274
	case 275:
		goto st275
	case 998:
		goto st998
	case 999:
		goto st999
	case 276:
		goto st276
	case 1000:
		goto st1000
	case 1001:
		goto st1001
	case 1002:
		goto st1002
	case 1003:
		goto st1003
	case 1004:
		goto st1004
	case 1005:
		goto st1005
	case 1006:
		goto st1006
	case 1007:
		goto st1007
	case 1008:
		goto st1008
	case 1009:
		goto st1009
	case 1010:
		goto st1010
	case 1011:
		goto st1011
	case 1012:
		goto st1012
	case 1013:
		goto st1013
	case 1014:
		goto st1014
	case 1015:
		goto st1015
	case 1016:
		goto st1016
	case 1017:
		goto st1017
	case 1018:
		goto st1018
	case 277:
		goto st277
	case 1019:
		goto st1019
	case 1020:
		goto st1020
	case 278:
		goto st278
	case 279:
		goto st279
	case 280:
		goto st280
	case 281:
		goto st281
	case 1021:
		goto st1021
	case 1022:
		goto st1022
	case 1023:
		goto st1023
	case 1024:
		goto st1024
	case 1025:
		goto st1025
	case 282:
		goto st282
	case 283:
		goto st283
	case 284:
		goto st284
	case 1026:
		goto st1026
	case 285:
		goto st285
	case 286:
		goto st286
	case 287:
		goto st287
	case 1027:
		goto st1027
	case 288:
		goto st288
	case 289:
		goto st289
	case 1028:
		goto st1028
	case 1029:
		goto st1029
	case 290:
		goto st290
	case 291:
		goto st291
	case 1030:
		goto st1030
	case 1031:
		goto st1031
	case 1032:
		goto st1032
	case 292:
		goto st292
	case 1033:
		goto st1033
	case 1034:
		goto st1034
	case 293:
		goto st293
	case 1035:
		goto st1035
	case 1036:
		goto st1036
	case 1037:
		goto st1037
	case 1038:
		goto st1038
	case 1039:
		goto st1039
	case 1040:
		goto st1040
	case 1041:
		goto st1041
	case 1042:
		goto st1042
	case 1043:
		goto st1043
	case 1044:
		goto st1044
	case 1045:
		goto st1045
	case 1046:
		goto st1046
	case 1047:
		goto st1047
	case 1048:
		goto st1048
	case 1049:
		goto st1049
	case 1050:
		goto st1050
	case 1051:
		goto st1051
	case 1052:
		goto st1052
	case 1053:
		goto st1053
	case 1054:
		goto st1054
	case 294:
		goto st294
	case 295:
		goto st295
	case 1055:
		goto st1055
	case 1056:
		goto st1056
	case 296:
		goto st296
	case 1057:
		goto st1057
	case 1058:
		goto st1058
	case 1059:
		goto st1059
	case 1060:
		goto st1060
	case 1061:
		goto st1061
	case 1062:
		goto st1062
	case 1063:
		goto st1063
	case 1064:
		goto st1064
	case 1065:
		goto st1065
	case 1066:
		goto st1066
	case 1067:
		goto st1067
	case 1068:
		goto st1068
	case 1069:
		goto st1069
	case 1070:
		goto st1070
	case 1071:
		goto st1071
	case 1072:
		goto st1072
	case 1073:
		goto st1073
	case 1074:
		goto st1074
	case 1075:
		goto st1075
	case 1076:
		goto st1076
	case 297:
		goto st297
	case 1077:
		goto st1077
	case 1078:
		goto st1078
	case 1079:
		goto st1079
	case 298:
		goto st298
	case 1080:
		goto st1080
	case 1081:
		goto st1081
	case 1082:
		goto st1082
	case 1083:
		goto st1083
	case 1084:
		goto st1084
	case 1085:
		goto st1085
	case 1086:
		goto st1086
	case 1087:
		goto st1087
	case 1088:
		goto st1088
	case 1089:
		goto st1089
	case 1090:
		goto st1090
	case 1091:
		goto st1091
	case 1092:
		goto st1092
	case 1093:
		goto st1093
	case 1094:
		goto st1094
	case 1095:
		goto st1095
	case 1096:
		goto st1096
	case 1097:
		goto st1097
	case 1098:
		goto st1098
	case 1099:
		goto st1099
	case 1100:
		goto st1100
	case 1101:
		goto st1101
	case 1102:
		goto st1102
	case 299:
		goto st299
	case 1103:
		goto st1103
	case 1104:
		goto st1104
	case 1105:
		goto st1105
	case 1106:
		goto st1106
	case 1107:
		goto st1107
	case 1108:
		goto st1108
	case 1109:
		goto st1109
	case 1110:
		goto st1110
	case 1111:
		goto st1111
	case 1112:
		goto st1112
	case 1113:
		goto st1113
	case 1114:
		goto st1114
	case 1115:
		goto st1115
	case 1116:
		goto st1116
	case 1117:
		goto st1117
	case 1118:
		goto st1118
	case 1119:
		goto st1119
	case 1120:
		goto st1120
	case 1121:
		goto st1121
	case 1122:
		goto st1122
	case 1123:
		goto st1123
	case 1124:
		goto st1124
	case 300:
		goto st300
	case 301:
		goto st301
	case 1125:
		goto st1125
	case 1126:
		goto st1126
	case 302:
		goto st302
	case 1127:
		goto st1127
	case 1128:
		goto st1128
	case 1129:
		goto st1129
	case 1130:
		goto st1130
	case 1131:
		goto st1131
	case 1132:
		goto st1132
	case 1133:
		goto st1133
	case 1134:
		goto st1134
	case 1135:
		goto st1135
	case 1136:
		goto st1136
	case 1137:
		goto st1137
	case 1138:
		goto st1138
	case 1139:
		goto st1139
	case 1140:
		goto st1140
	case 1141:
		goto st1141
	case 1142:
		goto st1142
	case 1143:
		goto st1143
	case 1144:
		goto st1144
	case 1145:
		goto st1145
	case 303:
		goto st303
	case 1146:
		goto st1146
	case 1147:
		goto st1147
	case 1148:
		goto st1148
	case 304:
		goto st304
	case 1149:
		goto st1149
	case 1150:
		goto st1150
	case 305:
		goto st305
	case 1151:
		goto st1151
	case 1152:
		goto st1152
	case 1153:
		goto st1153
	case 1154:
		goto st1154
	case 1155:
		goto st1155
	case 1156:
		goto st1156
	case 1157:
		goto st1157
	case 1158:
		goto st1158
	case 1159:
		goto st1159
	case 1160:
		goto st1160
	case 1161:
		goto st1161
	case 1162:
		goto st1162
	case 1163:
		goto st1163
	case 1164:
		goto st1164
	case 1165:
		goto st1165
	case 1166:
		goto st1166
	case 1167:
		goto st1167
	case 1168:
		goto st1168
	case 1169:
		goto st1169
	case 306:
		goto st306
	case 307:
		goto st307
	case 308:
		goto st308
	case 309:
		goto st309
	case 1170:
		goto st1170
	case 1171:
		goto st1171
	case 1172:
		goto st1172
	case 1173:
		goto st1173
	case 1174:
		goto st1174
	case 310:
		goto st310
	case 311:
		goto st311
	case 312:
		goto st312
	case 1175:
		goto st1175
	case 313:
		goto st313
	case 314:
		goto st314
	case 315:
		goto st315
	case 1176:
		goto st1176
	case 316:
		goto st316
	case 317:
		goto st317
	case 1177:
		goto st1177
	case 1178:
		goto st1178
	case 318:
		goto st318
	case 319:
		goto st319
	case 1179:
		goto st1179
	case 1180:
		goto st1180
	case 1181:
		goto st1181
	case 1182:
		goto st1182
	case 1183:
		goto st1183
	case 320:
		goto st320
	case 321:
		goto st321
	case 322:
		goto st322
	case 1184:
		goto st1184
	case 323:
		goto st323
	case 324:
		goto st324
	case 325:
		goto st325
	case 1185:
		goto st1185
	case 326:
		goto st326
	case 327:
		goto st327
	case 1186:
		goto st1186
	case 1187:
		goto st1187
	case 328:
		goto st328
	case 329:
		goto st329
	case 1188:
		goto st1188
	case 1189:
		goto st1189
	case 1190:
		goto st1190
	case 330:
		goto st330
	case 1191:
		goto st1191
	case 1192:
		goto st1192
	case 1193:
		goto st1193
	case 1194:
		goto st1194
	case 1195:
		goto st1195
	case 1196:
		goto st1196
	case 1197:
		goto st1197
	case 1198:
		goto st1198
	case 1199:
		goto st1199
	case 1200:
		goto st1200
	case 1201:
		goto st1201
	case 1202:
		goto st1202
	case 1203:
		goto st1203
	case 1204:
		goto st1204
	case 1205:
		goto st1205
	case 1206:
		goto st1206
	case 1207:
		goto st1207
	case 1208:
		goto st1208
	case 1209:
		goto st1209
	case 1210:
		goto st1210
	case 1211:
		goto st1211
	case 1212:
		goto st1212
	case 1213:
		goto st1213
	case 331:
		goto st331
	case 332:
		goto st332
	case 333:
		goto st333
	case 1214:
		goto st1214
	case 334:
		goto st334
	case 335:
		goto st335
	case 336:
		goto st336
	case 1215:
		goto st1215
	case 337:
		goto st337
	case 338:
		goto st338
	case 1216:
		goto st1216
	case 1217:
		goto st1217
	case 339:
		goto st339
	case 340:
		goto st340
	case 1218:
		goto st1218
	case 1219:
		goto st1219
	case 1220:
		goto st1220
	case 1221:
		goto st1221
	case 1222:
		goto st1222
	case 341:
		goto st341
	case 1223:
		goto st1223
	case 1224:
		goto st1224
	case 342:
		goto st342
	case 1225:
		goto st1225
	case 1226:
		goto st1226
	case 1227:
		goto st1227
	case 1228:
		goto st1228
	case 1229:
		goto st1229
	case 1230:
		goto st1230
	case 1231:
		goto st1231
	case 1232:
		goto st1232
	case 1233:
		goto st1233
	case 1234:
		goto st1234
	case 1235:
		goto st1235
	case 1236:
		goto st1236
	case 1237:
		goto st1237
	case 1238:
		goto st1238
	case 1239:
		goto st1239
	case 1240:
		goto st1240
	case 1241:
		goto st1241
	case 1242:
		goto st1242
	case 1243:
		goto st1243
	case 343:
		goto st343
	case 1244:
		goto st1244
	case 1245:
		goto st1245
	case 1246:
		goto st1246
	case 344:
		goto st344
	case 345:
		goto st345
	case 346:
		goto st346
	case 1247:
		goto st1247
	case 347:
		goto st347
	case 348:
		goto st348
	case 349:
		goto st349
	case 1248:
		goto st1248
	case 350:
		goto st350
	case 351:
		goto st351
	case 1249:
		goto st1249
	case 1250:
		goto st1250
	case 352:
		goto st352
	case 353:
		goto st353
	case 354:
		goto st354
	case 355:
		goto st355
	case 356:
		goto st356
	case 357:
		goto st357
	case 1251:
		goto st1251
	case 1252:
		goto st1252
	case 1253:
		goto st1253
	case 1254:
		goto st1254
	case 1255:
		goto st1255
	case 1256:
		goto st1256
	case 1257:
		goto st1257
	case 1258:
		goto st1258
	case 1259:
		goto st1259
	case 1260:
		goto st1260
	case 1261:
		goto st1261
	case 1262:
		goto st1262
	case 1263:
		goto st1263
	case 1264:
		goto st1264
	case 1265:
		goto st1265
	case 1266:
		goto st1266
	case 1267:
		goto st1267
	case 1268:
		goto st1268
	case 1269:
		goto st1269
	case 358:
		goto st358
	case 359:
		goto st359
	case 360:
		goto st360
	case 1270:
		goto st1270
	case 361:
		goto st361
	case 362:
		goto st362
	case 363:
		goto st363
	case 364:
		goto st364
	case 1271:
		goto st1271
	case 1272:
		goto st1272
	case 1273:
		goto st1273
	case 1274:
		goto st1274
	case 1275:
		goto st1275
	case 365:
		goto st365
	case 366:
		goto st366
	case 367:
		goto st367
	case 1276:
		goto st1276
	case 368:
		goto st368
	case 369:
		goto st369
	case 370:
		goto st370
	case 1277:
		goto st1277
	case 371:
		goto st371
	case 372:
		goto st372
	case 1278:
		goto st1278
	case 1279:
		goto st1279
	case 373:
		goto st373
	case 374:
		goto st374
	case 1280:
		goto st1280
	case 1281:
		goto st1281
	case 1282:
		goto st1282
	case 1283:
		goto st1283
	case 1284:
		goto st1284
	case 375:
		goto st375
	case 376:
		goto st376
	case 377:
		goto st377
	case 1285:
		goto st1285
	case 378:
		goto st378
	case 379:
		goto st379
	case 380:
		goto st380
	case 1286:
		goto st1286
	case 381:
		goto st381
	case 382:
		goto st382
	case 1287:
		goto st1287
	case 1288:
		goto st1288
	case 383:
		goto st383
	case 384:
		goto st384
	case 385:
		goto st385
	case 386:
		goto st386
	case 387:
		goto st387
	case 388:
		goto st388
	case 1289:
		goto st1289
	case 1290:
		goto st1290
	case 389:
		goto st389
	case 1291:
		goto st1291
	case 1292:
		goto st1292
	case 390:
		goto st390
	case 1293:
		goto st1293
	case 1294:
		goto st1294
	case 1295:
		goto st1295
	case 1296:
		goto st1296
	case 1297:
		goto st1297
	case 1298:
		goto st1298
	case 1299:
		goto st1299
	case 1300:
		goto st1300
	case 1301:
		goto st1301
	case 1302:
		goto st1302
	case 1303:
		goto st1303
	case 1304:
		goto st1304
	case 1305:
		goto st1305
	case 1306:
		goto st1306
	case 1307:
		goto st1307
	case 1308:
		goto st1308
	case 1309:
		goto st1309
	case 1310:
		goto st1310
	case 391:
		goto st391
	case 392:
		goto st392
	case 393:
		goto st393
	case 394:
		goto st394
	case 395:
		goto st395
	case 396:
		goto st396
	case 397:
		goto st397
	case 398:
		goto st398
	case 399:
		goto st399
	case 400:
		goto st400
	case 401:
		goto st401
	case 402:
		goto st402
	case 403:
		goto st403
	case 1311:
		goto st1311
	case 404:
		goto st404
	case 405:
		goto st405
	case 406:
		goto st406
	case 1312:
		goto st1312
	case 407:
		goto st407
	case 408:
		goto st408
	case 409:
		goto st409
	case 410:
		goto st410
	case 1313:
		goto st1313
	case 1314:
		goto st1314
	case 1315:
		goto st1315
	case 1316:
		goto st1316
	case 1317:
		goto st1317
	case 411:
		goto st411
	case 412:
		goto st412
	case 413:
		goto st413
	case 1318:
		goto st1318
	case 414:
		goto st414
	case 415:
		goto st415
	case 416:
		goto st416
	case 1319:
		goto st1319
	case 417:
		goto st417
	case 418:
		goto st418
	case 1320:
		goto st1320
	case 1321:
		goto st1321
	case 419:
		goto st419
	case 420:
		goto st420
	case 421:
		goto st421
	case 422:
		goto st422
	case 423:
		goto st423
	case 424:
		goto st424
	case 425:
		goto st425
	case 426:
		goto st426
	case 427:
		goto st427
	case 428:
		goto st428
	case 429:
		goto st429
	case 430:
		goto st430
	case 1322:
		goto st1322
	case 1323:
		goto st1323
	case 1324:
		goto st1324
	case 431:
		goto st431
	case 1325:
		goto st1325
	case 1326:
		goto st1326
	case 1327:
		goto st1327
	case 1328:
		goto st1328
	case 1329:
		goto st1329
	case 1330:
		goto st1330
	case 1331:
		goto st1331
	case 1332:
		goto st1332
	case 1333:
		goto st1333
	case 1334:
		goto st1334
	case 1335:
		goto st1335
	case 1336:
		goto st1336
	case 1337:
		goto st1337
	case 1338:
		goto st1338
	case 1339:
		goto st1339
	case 1340:
		goto st1340
	case 1341:
		goto st1341
	case 1342:
		goto st1342
	case 1343:
		goto st1343
	case 1344:
		goto st1344
	case 1345:
		goto st1345
	case 1346:
		goto st1346
	case 432:
		goto st432
	case 1347:
		goto st1347
	case 1348:
		goto st1348
	case 1349:
		goto st1349
	case 1350:
		goto st1350
	case 433:
		goto st433
	case 1351:
		goto st1351
	case 1352:
		goto st1352
	case 1353:
		goto st1353
	case 1354:
		goto st1354
	case 1355:
		goto st1355
	case 1356:
		goto st1356
	case 1357:
		goto st1357
	case 1358:
		goto st1358
	case 1359:
		goto st1359
	case 1360:
		goto st1360
	case 1361:
		goto st1361
	case 1362:
		goto st1362
	case 1363:
		goto st1363
	case 1364:
		goto st1364
	case 1365:
		goto st1365
	case 1366:
		goto st1366
	case 1367:
		goto st1367
	case 1368:
		goto st1368
	case 1369:
		goto st1369
	case 1370:
		goto st1370
	case 434:
		goto st434
	case 435:
		goto st435
	case 436:
		goto st436
	case 1371:
		goto st1371
	case 1372:
		goto st1372
	case 437:
		goto st437
	case 1373:
		goto st1373
	case 1374:
		goto st1374
	case 1375:
		goto st1375
	case 1376:
		goto st1376
	case 1377:
		goto st1377
	case 1378:
		goto st1378
	case 1379:
		goto st1379
	case 1380:
		goto st1380
	case 1381:
		goto st1381
	case 1382:
		goto st1382
	case 1383:
		goto st1383
	case 1384:
		goto st1384
	case 1385:
		goto st1385
	case 1386:
		goto st1386
	case 1387:
		goto st1387
	case 1388:
		goto st1388
	case 1389:
		goto st1389
	case 1390:
		goto st1390
	case 1391:
		goto st1391
	case 1392:
		goto st1392
	case 438:
		goto st438
	case 1393:
		goto st1393
	case 1394:
		goto st1394
	case 1395:
		goto st1395
	case 439:
		goto st439
	case 1396:
		goto st1396
	case 1397:
		goto st1397
	case 1398:
		goto st1398
	case 1399:
		goto st1399
	case 1400:
		goto st1400
	case 1401:
		goto st1401
	case 1402:
		goto st1402
	case 1403:
		goto st1403
	case 1404:
		goto st1404
	case 1405:
		goto st1405
	case 1406:
		goto st1406
	case 1407:
		goto st1407
	case 1408:
		goto st1408
	case 1409:
		goto st1409
	case 1410:
		goto st1410
	case 1411:
		goto st1411
	case 1412:
		goto st1412
	case 1413:
		goto st1413
	case 1414:
		goto st1414
	case 1415:
		goto st1415
	case 1416:
		goto st1416
	case 1417:
		goto st1417
	case 440:
		goto st440
	case 1418:
		goto st1418
	case 1419:
		goto st1419
	case 1420:
		goto st1420
	case 1421:
		goto st1421
	case 1422:
		goto st1422
	case 1423:
		goto st1423
	case 1424:
		goto st1424
	case 1425:
		goto st1425
	case 1426:
		goto st1426
	case 1427:
		goto st1427
	case 1428:
		goto st1428
	case 1429:
		goto st1429
	case 1430:
		goto st1430
	case 1431:
		goto st1431
	case 1432:
		goto st1432
	case 1433:
		goto st1433
	case 1434:
		goto st1434
	case 1435:
		goto st1435
	case 1436:
		goto st1436
	case 1437:
		goto st1437
	case 1438:
		goto st1438
	case 1439:
		goto st1439
	case 441:
		goto st441
	case 442:
		goto st442
	case 443:
		goto st443
	case 444:
		goto st444
	case 1440:
		goto st1440
	case 1441:
		goto st1441
	case 1442:
		goto st1442
	case 1443:
		goto st1443
	case 1444:
		goto st1444
	case 445:
		goto st445
	case 446:
		goto st446
	case 447:
		goto st447
	case 1445:
		goto st1445
	case 448:
		goto st448
	case 449:
		goto st449
	case 450:
		goto st450
	case 1446:
		goto st1446
	case 451:
		goto st451
	case 452:
		goto st452
	case 1447:
		goto st1447
	case 1448:
		goto st1448
	case 453:
		goto st453
	case 454:
		goto st454
	case 1449:
		goto st1449
	case 1450:
		goto st1450
	case 1451:
		goto st1451
	case 455:
		goto st455
	case 1452:
		goto st1452
	case 1453:
		goto st1453
	case 1454:
		goto st1454
	case 1455:
		goto st1455
	case 1456:
		goto st1456
	case 1457:
		goto st1457
	case 1458:
		goto st1458
	case 1459:
		goto st1459
	case 1460:
		goto st1460
	case 1461:
		goto st1461
	case 1462:
		goto st1462
	case 1463:
		goto st1463
	case 1464:
		goto st1464
	case 1465:
		goto st1465
	case 1466:
		goto st1466
	case 1467:
		goto st1467
	case 1468:
		goto st1468
	case 1469:
		goto st1469
	case 1470:
		goto st1470
	case 1471:
		goto st1471
	case 1472:
		goto st1472
	case 1473:
		goto st1473
	case 1474:
		goto st1474
	case 1475:
		goto st1475
	case 456:
		goto st456
	case 457:
		goto st457
	case 458:
		goto st458
	case 1476:
		goto st1476
	case 459:
		goto st459
	case 460:
		goto st460
	case 461:
		goto st461
	case 1477:
		goto st1477
	case 462:
		goto st462
	case 463:
		goto st463
	case 1478:
		goto st1478
	case 1479:
		goto st1479
	case 464:
		goto st464
	case 465:
		goto st465
	case 466:
		goto st466
	case 1480:
		goto st1480
	case 1481:
		goto st1481
	case 1482:
		goto st1482
	case 1483:
		goto st1483
	case 1484:
		goto st1484
	case 467:
		goto st467
	case 468:
		goto st468
	case 469:
		goto st469
	case 1485:
		goto st1485
	case 470:
		goto st470
	case 471:
		goto st471
	case 472:
		goto st472
	case 1486:
		goto st1486
	case 473:
		goto st473
	case 474:
		goto st474
	case 1487:
		goto st1487
	case 1488:
		goto st1488
	case 475:
		goto st475
	case 1489:
		goto st1489
	case 1490:
		goto st1490
	case 1491:
		goto st1491
	case 1492:
		goto st1492
	case 1493:
		goto st1493
	case 476:
		goto st476
	case 477:
		goto st477
	case 478:
		goto st478
	case 1494:
		goto st1494
	case 479:
		goto st479
	case 480:
		goto st480
	case 481:
		goto st481
	case 1495:
		goto st1495
	case 482:
		goto st482
	case 483:
		goto st483
	case 1496:
		goto st1496
	case 1497:
		goto st1497
	case 484:
		goto st484
	case 485:
		goto st485
	case 486:
		goto st486
	case 487:
		goto st487
	case 488:
		goto st488
	case 489:
		goto st489
	case 1498:
		goto st1498
	}

	if ( m.p)++; ( m.p) == ( m.pe) {
		goto _test_eof
	}
_resume:
	switch  m.cs {
	case 490:
		goto st_case_490
	case 1:
		goto st_case_1
	case 2:
		goto st_case_2
	case 3:
		goto st_case_3
	case 0:
		goto st_case_0
	case 4:
		goto st_case_4
	case 5:
		goto st_case_5
	case 6:
		goto st_case_6
	case 491:
		goto st_case_491
	case 492:
		goto st_case_492
	case 493:
		goto st_case_493
	case 7:
		goto st_case_7
	case 8:
		goto st_case_8
	case 9:
		goto st_case_9
	case 10:
		goto st_case_10
	case 11:
		goto st_case_11
	case 12:
		goto st_case_12
	case 13:
		goto st_case_13
	case 14:
		goto st_case_14
	case 15:
		goto st_case_15
	case 16:
		goto st_case_16
	case 17:
		goto st_case_17
	case 18:
		goto st_case_18
	case 19:
		goto st_case_19
	case 20:
		goto st_case_20
	case 21:
		goto st_case_21
	case 22:
		goto st_case_22
	case 23:
		goto st_case_23
	case 24:
		goto st_case_24
	case 25:
		goto st_case_25
	case 494:
		goto st_case_494
	case 495:
		goto st_case_495
	case 26:
		goto st_case_26
	case 496:
		goto st_case_496
	case 497:
		goto st_case_497
	case 498:
		goto st_case_498
	case 27:
		goto st_case_27
	case 499:
		goto st_case_499
	case 500:
		goto st_case_500
	case 501:
		goto st_case_501
	case 502:
		goto st_case_502
	case 503:
		goto st_case_503
	case 504:
		goto st_case_504
	case 505:
		goto st_case_505
	case 506:
		goto st_case_506
	case 507:
		goto st_case_507
	case 508:
		goto st_case_508
	case 509:
		goto st_case_509
	case 510:
		goto st_case_510
	case 511:
		goto st_case_511
	case 512:
		goto st_case_512
	case 513:
		goto st_case_513
	case 514:
		goto st_case_514
	case 515:
		goto st_case_515
	case 516:
		goto st_case_516
	case 28:
		goto st_case_28
	case 29:
		goto st_case_29
	case 517:
		goto st_case_517
	case 518:
		goto st_case_518
	case 519:
		goto st_case_519
	case 30:
		goto st_case_30
	case 31:
		goto st_case_31
	case 32:
		goto st_case_32
	case 33:
		goto st_case_33
	case 34:
		goto st_case_34
	case 520:
		goto st_case_520
	case 521:
		goto st_case_521
	case 522:
		goto st_case_522
	case 523:
		goto st_case_523
	case 35:
		goto st_case_35
	case 524:
		goto st_case_524
	case 525:
		goto st_case_525
	case 526:
		goto st_case_526
	case 527:
		goto st_case_527
	case 36:
		goto st_case_36
	case 528:
		goto st_case_528
	case 529:
		goto st_case_529
	case 530:
		goto st_case_530
	case 531:
		goto st_case_531
	case 532:
		goto st_case_532
	case 533:
		goto st_case_533
	case 534:
		goto st_case_534
	case 535:
		goto st_case_535
	case 536:
		goto st_case_536
	case 537:
		goto st_case_537
	case 538:
		goto st_case_538
	case 539:
		goto st_case_539
	case 540:
		goto st_case_540
	case 541:
		goto st_case_541
	case 542:
		goto st_case_542
	case 543:
		goto st_case_543
	case 544:
		goto st_case_544
	case 545:
		goto st_case_545
	case 37:
		goto st_case_37
	case 38:
		goto st_case_38
	case 39:
		goto st_case_39
	case 40:
		goto st_case_40
	case 41:
		goto st_case_41
	case 42:
		goto st_case_42
	case 43:
		goto st_case_43
	case 44:
		goto st_case_44
	case 45:
		goto st_case_45
	case 546:
		goto st_case_546
	case 547:
		goto st_case_547
	case 548:
		goto st_case_548
	case 46:
		goto st_case_46
	case 47:
		goto st_case_47
	case 48:
		goto st_case_48
	case 49:
		goto st_case_49
	case 50:
		goto st_case_50
	case 51:
		goto st_case_51
	case 549:
		goto st_case_549
	case 550:
		goto st_case_550
	case 52:
		goto st_case_52
	case 551:
		goto st_case_551
	case 552:
		goto st_case_552
	case 553:
		goto st_case_553
	case 554:
		goto st_case_554
	case 555:
		goto st_case_555
	case 556:
		goto st_case_556
	case 557:
		goto st_case_557
	case 558:
		goto st_case_558
	case 559:
		goto st_case_559
	case 560:
		goto st_case_560
	case 561:
		goto st_case_561
	case 562:
		goto st_case_562
	case 563:
		goto st_case_563
	case 564:
		goto st_case_564
	case 565:
		goto st_case_565
	case 566:
		goto st_case_566
	case 567:
		goto st_case_567
	case 568:
		goto st_case_568
	case 569:
		goto st_case_569
	case 570:
		goto st_case_570
	case 53:
		goto st_case_53
	case 571:
		goto st_case_571
	case 572:
		goto st_case_572
	case 573:
		goto st_case_573
	case 54:
		goto st_case_54
	case 574:
		goto st_case_574
	case 575:
		goto st_case_575
	case 576:
		goto st_case_576
	case 577:
		goto st_case_577
	case 578:
		goto st_case_578
	case 579:
		goto st_case_579
	case 580:
		goto st_case_580
	case 581:
		goto st_case_581
	case 582:
		goto st_case_582
	case 583:
		goto st_case_583
	case 584:
		goto st_case_584
	case 585:
		goto st_case_585
	case 586:
		goto st_case_586
	case 587:
		goto st_case_587
	case 588:
		goto st_case_588
	case 589:
		goto st_case_589
	case 590:
		goto st_case_590
	case 591:
		goto st_case_591
	case 592:
		goto st_case_592
	case 593:
		goto st_case_593
	case 55:
		goto st_case_55
	case 56:
		goto st_case_56
	case 57:
		goto st_case_57
	case 594:
		goto st_case_594
	case 595:
		goto st_case_595
	case 596:
		goto st_case_596
	case 58:
		goto st_case_58
	case 597:
		goto st_case_597
	case 598:
		goto st_case_598
	case 59:
		goto st_case_59
	case 599:
		goto st_case_599
	case 60:
		goto st_case_60
	case 600:
		goto st_case_600
	case 601:
		goto st_case_601
	case 602:
		goto st_case_602
	case 603:
		goto st_case_603
	case 604:
		goto st_case_604
	case 605:
		goto st_case_605
	case 606:
		goto st_case_606
	case 607:
		goto st_case_607
	case 608:
		goto st_case_608
	case 609:
		goto st_case_609
	case 610:
		goto st_case_610
	case 611:
		goto st_case_611
	case 612:
		goto st_case_612
	case 613:
		goto st_case_613
	case 614:
		goto st_case_614
	case 615:
		goto st_case_615
	case 616:
		goto st_case_616
	case 617:
		goto st_case_617
	case 618:
		goto st_case_618
	case 619:
		goto st_case_619
	case 61:
		goto st_case_61
	case 62:
		goto st_case_62
	case 63:
		goto st_case_63
	case 64:
		goto st_case_64
	case 65:
		goto st_case_65
	case 66:
		goto st_case_66
	case 67:
		goto st_case_67
	case 620:
		goto st_case_620
	case 68:
		goto st_case_68
	case 69:
		goto st_case_69
	case 70:
		goto st_case_70
	case 71:
		goto st_case_71
	case 621:
		goto st_case_621
	case 622:
		goto st_case_622
	case 623:
		goto st_case_623
	case 624:
		goto st_case_624
	case 72:
		goto st_case_72
	case 73:
		goto st_case_73
	case 74:
		goto st_case_74
	case 75:
		goto st_case_75
	case 625:
		goto st_case_625
	case 626:
		goto st_case_626
	case 76:
		goto st_case_76
	case 627:
		goto st_case_627
	case 77:
		goto st_case_77
	case 78:
		goto st_case_78
	case 628:
		goto st_case_628
	case 629:
		goto st_case_629
	case 79:
		goto st_case_79
	case 630:
		goto st_case_630
	case 631:
		goto st_case_631
	case 80:
		goto st_case_80
	case 632:
		goto st_case_632
	case 633:
		goto st_case_633
	case 634:
		goto st_case_634
	case 635:
		goto st_case_635
	case 636:
		goto st_case_636
	case 637:
		goto st_case_637
	case 638:
		goto st_case_638
	case 639:
		goto st_case_639
	case 640:
		goto st_case_640
	case 641:
		goto st_case_641
	case 642:
		goto st_case_642
	case 643:
		goto st_case_643
	case 644:
		goto st_case_644
	case 645:
		goto st_case_645
	case 646:
		goto st_case_646
	case 647:
		goto st_case_647
	case 648:
		goto st_case_648
	case 649:
		goto st_case_649
	case 81:
		goto st_case_81
	case 650:
		goto st_case_650
	case 651:
		goto st_case_651
	case 652:
		goto st_case_652
	case 82:
		goto st_case_82
	case 653:
		goto st_case_653
	case 654:
		goto st_case_654
	case 655:
		goto st_case_655
	case 83:
		goto st_case_83
	case 656:
		goto st_case_656
	case 657:
		goto st_case_657
	case 658:
		goto st_case_658
	case 659:
		goto st_case_659
	case 660:
		goto st_case_660
	case 661:
		goto st_case_661
	case 662:
		goto st_case_662
	case 663:
		goto st_case_663
	case 664:
		goto st_case_664
	case 665:
		goto st_case_665
	case 666:
		goto st_case_666
	case 667:
		goto st_case_667
	case 668:
		goto st_case_668
	case 669:
		goto st_case_669
	case 670:
		goto st_case_670
	case 671:
		goto st_case_671
	case 672:
		goto st_case_672
	case 673:
		goto st_case_673
	case 674:
		goto st_case_674
	case 84:
		goto st_case_84
	case 85:
		goto st_case_85
	case 86:
		goto st_case_86
	case 675:
		goto st_case_675
	case 87:
		goto st_case_87
	case 676:
		goto st_case_676
	case 677:
		goto st_case_677
	case 678:
		goto st_case_678
	case 679:
		goto st_case_679
	case 680:
		goto st_case_680
	case 681:
		goto st_case_681
	case 682:
		goto st_case_682
	case 683:
		goto st_case_683
	case 684:
		goto st_case_684
	case 685:
		goto st_case_685
	case 686:
		goto st_case_686
	case 687:
		goto st_case_687
	case 688:
		goto st_case_688
	case 689:
		goto st_case_689
	case 690:
		goto st_case_690
	case 691:
		goto st_case_691
	case 692:
		goto st_case_692
	case 693:
		goto st_case_693
	case 694:
		goto st_case_694
	case 695:
		goto st_case_695
	case 696:
		goto st_case_696
	case 697:
		goto st_case_697
	case 88:
		goto st_case_88
	case 89:
		goto st_case_89
	case 90:
		goto st_case_90
	case 91:
		goto st_case_91
	case 92:
		goto st_case_92
	case 698:
		goto st_case_698
	case 699:
		goto st_case_699
	case 700:
		goto st_case_700
	case 701:
		goto st_case_701
	case 702:
		goto st_case_702
	case 93:
		goto st_case_93
	case 94:
		goto st_case_94
	case 95:
		goto st_case_95
	case 703:
		goto st_case_703
	case 96:
		goto st_case_96
	case 97:
		goto st_case_97
	case 98:
		goto st_case_98
	case 704:
		goto st_case_704
	case 99:
		goto st_case_99
	case 100:
		goto st_case_100
	case 705:
		goto st_case_705
	case 706:
		goto st_case_706
	case 101:
		goto st_case_101
	case 102:
		goto st_case_102
	case 707:
		goto st_case_707
	case 708:
		goto st_case_708
	case 709:
		goto st_case_709
	case 103:
		goto st_case_103
	case 710:
		goto st_case_710
	case 711:
		goto st_case_711
	case 712:
		goto st_case_712
	case 713:
		goto st_case_713
	case 714:
		goto st_case_714
	case 715:
		goto st_case_715
	case 716:
		goto st_case_716
	case 717:
		goto st_case_717
	case 718:
		goto st_case_718
	case 719:
		goto st_case_719
	case 720:
		goto st_case_720
	case 721:
		goto st_case_721
	case 722:
		goto st_case_722
	case 723:
		goto st_case_723
	case 724:
		goto st_case_724
	case 725:
		goto st_case_725
	case 726:
		goto st_case_726
	case 727:
		goto st_case_727
	case 728:
		goto st_case_728
	case 729:
		goto st_case_729
	case 730:
		goto st_case_730
	case 731:
		goto st_case_731
	case 732:
		goto st_case_732
	case 733:
		goto st_case_733
	case 104:
		goto st_case_104
	case 105:
		goto st_case_105
	case 106:
		goto st_case_106
	case 734:
		goto st_case_734
	case 107:
		goto st_case_107
	case 108:
		goto st_case_108
	case 109:
		goto st_case_109
	case 735:
		goto st_case_735
	case 110:
		goto st_case_110
	case 111:
		goto st_case_111
	case 736:
		goto st_case_736
	case 737:
		goto st_case_737
	case 112:
		goto st_case_112
	case 738:
		goto st_case_738
	case 113:
		goto st_case_113
	case 739:
		goto st_case_739
	case 740:
		goto st_case_740
	case 741:
		goto st_case_741
	case 114:
		goto st_case_114
	case 115:
		goto st_case_115
	case 116:
		goto st_case_116
	case 742:
		goto st_case_742
	case 117:
		goto st_case_117
	case 118:
		goto st_case_118
	case 119:
		goto st_case_119
	case 743:
		goto st_case_743
	case 120:
		goto st_case_120
	case 121:
		goto st_case_121
	case 744:
		goto st_case_744
	case 745:
		goto st_case_745
	case 746:
		goto st_case_746
	case 747:
		goto st_case_747
	case 748:
		goto st_case_748
	case 749:
		goto st_case_749
	case 750:
		goto st_case_750
	case 751:
		goto st_case_751
	case 752:
		goto st_case_752
	case 753:
		goto st_case_753
	case 754:
		goto st_case_754
	case 755:
		goto st_case_755
	case 756:
		goto st_case_756
	case 757:
		goto st_case_757
	case 758:
		goto st_case_758
	case 759:
		goto st_case_759
	case 760:
		goto st_case_760
	case 761:
		goto st_case_761
	case 762:
		goto st_case_762
	case 763:
		goto st_case_763
	case 122:
		goto st_case_122
	case 764:
		goto st_case_764
	case 765:
		goto st_case_765
	case 766:
		goto st_case_766
	case 123:
		goto st_case_123
	case 767:
		goto st_case_767
	case 768:
		goto st_case_768
	case 124:
		goto st_case_124
	case 769:
		goto st_case_769
	case 770:
		goto st_case_770
	case 771:
		goto st_case_771
	case 772:
		goto st_case_772
	case 773:
		goto st_case_773
	case 774:
		goto st_case_774
	case 775:
		goto st_case_775
	case 776:
		goto st_case_776
	case 777:
		goto st_case_777
	case 778:
		goto st_case_778
	case 779:
		goto st_case_779
	case 780:
		goto st_case_780
	case 781:
		goto st_case_781
	case 782:
		goto st_case_782
	case 783:
		goto st_case_783
	case 784:
		goto st_case_784
	case 785:
		goto st_case_785
	case 786:
		goto st_case_786
	case 787:
		goto st_case_787
	case 125:
		goto st_case_125
	case 788:
		goto st_case_788
	case 789:
		goto st_case_789
	case 790:
		goto st_case_790
	case 126:
		goto st_case_126
	case 127:
		goto st_case_127
	case 128:
		goto st_case_128
	case 791:
		goto st_case_791
	case 129:
		goto st_case_129
	case 130:
		goto st_case_130
	case 131:
		goto st_case_131
	case 792:
		goto st_case_792
	case 132:
		goto st_case_132
	case 133:
		goto st_case_133
	case 793:
		goto st_case_793
	case 794:
		goto st_case_794
	case 134:
		goto st_case_134
	case 135:
		goto st_case_135
	case 136:
		goto st_case_136
	case 137:
		goto st_case_137
	case 138:
		goto st_case_138
	case 139:
		goto st_case_139
	case 795:
		goto st_case_795
	case 796:
		goto st_case_796
	case 797:
		goto st_case_797
	case 798:
		goto st_case_798
	case 799:
		goto st_case_799
	case 800:
		goto st_case_800
	case 801:
		goto st_case_801
	case 802:
		goto st_case_802
	case 803:
		goto st_case_803
	case 804:
		goto st_case_804
	case 805:
		goto st_case_805
	case 806:
		goto st_case_806
	case 807:
		goto st_case_807
	case 808:
		goto st_case_808
	case 809:
		goto st_case_809
	case 810:
		goto st_case_810
	case 811:
		goto st_case_811
	case 812:
		goto st_case_812
	case 813:
		goto st_case_813
	case 140:
		goto st_case_140
	case 141:
		goto st_case_141
	case 142:
		goto st_case_142
	case 814:
		goto st_case_814
	case 815:
		goto st_case_815
	case 816:
		goto st_case_816
	case 817:
		goto st_case_817
	case 818:
		goto st_case_818
	case 143:
		goto st_case_143
	case 144:
		goto st_case_144
	case 145:
		goto st_case_145
	case 819:
		goto st_case_819
	case 146:
		goto st_case_146
	case 147:
		goto st_case_147
	case 148:
		goto st_case_148
	case 820:
		goto st_case_820
	case 149:
		goto st_case_149
	case 150:
		goto st_case_150
	case 821:
		goto st_case_821
	case 822:
		goto st_case_822
	case 151:
		goto st_case_151
	case 152:
		goto st_case_152
	case 153:
		goto st_case_153
	case 823:
		goto st_case_823
	case 824:
		goto st_case_824
	case 825:
		goto st_case_825
	case 826:
		goto st_case_826
	case 827:
		goto st_case_827
	case 154:
		goto st_case_154
	case 155:
		goto st_case_155
	case 156:
		goto st_case_156
	case 828:
		goto st_case_828
	case 157:
		goto st_case_157
	case 158:
		goto st_case_158
	case 159:
		goto st_case_159
	case 829:
		goto st_case_829
	case 160:
		goto st_case_160
	case 161:
		goto st_case_161
	case 830:
		goto st_case_830
	case 831:
		goto st_case_831
	case 162:
		goto st_case_162
	case 163:
		goto st_case_163
	case 164:
		goto st_case_164
	case 832:
		goto st_case_832
	case 833:
		goto st_case_833
	case 834:
		goto st_case_834
	case 835:
		goto st_case_835
	case 836:
		goto st_case_836
	case 837:
		goto st_case_837
	case 838:
		goto st_case_838
	case 839:
		goto st_case_839
	case 840:
		goto st_case_840
	case 841:
		goto st_case_841
	case 842:
		goto st_case_842
	case 843:
		goto st_case_843
	case 844:
		goto st_case_844
	case 845:
		goto st_case_845
	case 846:
		goto st_case_846
	case 847:
		goto st_case_847
	case 848:
		goto st_case_848
	case 849:
		goto st_case_849
	case 850:
		goto st_case_850
	case 165:
		goto st_case_165
	case 166:
		goto st_case_166
	case 167:
		goto st_case_167
	case 168:
		goto st_case_168
	case 851:
		goto st_case_851
	case 852:
		goto st_case_852
	case 853:
		goto st_case_853
	case 854:
		goto st_case_854
	case 855:
		goto st_case_855
	case 169:
		goto st_case_169
	case 170:
		goto st_case_170
	case 171:
		goto st_case_171
	case 856:
		goto st_case_856
	case 172:
		goto st_case_172
	case 173:
		goto st_case_173
	case 174:
		goto st_case_174
	case 857:
		goto st_case_857
	case 175:
		goto st_case_175
	case 176:
		goto st_case_176
	case 858:
		goto st_case_858
	case 859:
		goto st_case_859
	case 177:
		goto st_case_177
	case 178:
		goto st_case_178
	case 179:
		goto st_case_179
	case 860:
		goto st_case_860
	case 861:
		goto st_case_861
	case 862:
		goto st_case_862
	case 863:
		goto st_case_863
	case 864:
		goto st_case_864
	case 180:
		goto st_case_180
	case 181:
		goto st_case_181
	case 182:
		goto st_case_182
	case 865:
		goto st_case_865
	case 183:
		goto st_case_183
	case 184:
		goto st_case_184
	case 185:
		goto st_case_185
	case 866:
		goto st_case_866
	case 186:
		goto st_case_186
	case 187:
		goto st_case_187
	case 867:
		goto st_case_867
	case 868:
		goto st_case_868
	case 188:
		goto st_case_188
	case 189:
		goto st_case_189
	case 190:
		goto st_case_190
	case 191:
		goto st_case_191
	case 192:
		goto st_case_192
	case 193:
		goto st_case_193
	case 869:
		goto st_case_869
	case 870:
		goto st_case_870
	case 871:
		goto st_case_871
	case 872:
		goto st_case_872
	case 873:
		goto st_case_873
	case 874:
		goto st_case_874
	case 875:
		goto st_case_875
	case 876:
		goto st_case_876
	case 877:
		goto st_case_877
	case 878:
		goto st_case_878
	case 879:
		goto st_case_879
	case 880:
		goto st_case_880
	case 881:
		goto st_case_881
	case 882:
		goto st_case_882
	case 883:
		goto st_case_883
	case 884:
		goto st_case_884
	case 885:
		goto st_case_885
	case 886:
		goto st_case_886
	case 887:
		goto st_case_887
	case 194:
		goto st_case_194
	case 888:
		goto st_case_888
	case 889:
		goto st_case_889
	case 890:
		goto st_case_890
	case 891:
		goto st_case_891
	case 892:
		goto st_case_892
	case 195:
		goto st_case_195
	case 196:
		goto st_case_196
	case 197:
		goto st_case_197
	case 893:
		goto st_case_893
	case 198:
		goto st_case_198
	case 199:
		goto st_case_199
	case 200:
		goto st_case_200
	case 894:
		goto st_case_894
	case 201:
		goto st_case_201
	case 202:
		goto st_case_202
	case 895:
		goto st_case_895
	case 896:
		goto st_case_896
	case 203:
		goto st_case_203
	case 204:
		goto st_case_204
	case 205:
		goto st_case_205
	case 206:
		goto st_case_206
	case 207:
		goto st_case_207
	case 208:
		goto st_case_208
	case 209:
		goto st_case_209
	case 210:
		goto st_case_210
	case 211:
		goto st_case_211
	case 212:
		goto st_case_212
	case 897:
		goto st_case_897
	case 898:
		goto st_case_898
	case 899:
		goto st_case_899
	case 213:
		goto st_case_213
	case 214:
		goto st_case_214
	case 215:
		goto st_case_215
	case 900:
		goto st_case_900
	case 901:
		goto st_case_901
	case 216:
		goto st_case_216
	case 902:
		goto st_case_902
	case 903:
		goto st_case_903
	case 904:
		goto st_case_904
	case 905:
		goto st_case_905
	case 906:
		goto st_case_906
	case 907:
		goto st_case_907
	case 908:
		goto st_case_908
	case 909:
		goto st_case_909
	case 910:
		goto st_case_910
	case 911:
		goto st_case_911
	case 912:
		goto st_case_912
	case 913:
		goto st_case_913
	case 914:
		goto st_case_914
	case 915:
		goto st_case_915
	case 916:
		goto st_case_916
	case 917:
		goto st_case_917
	case 918:
		goto st_case_918
	case 919:
		goto st_case_919
	case 920:
		goto st_case_920
	case 217:
		goto st_case_217
	case 921:
		goto st_case_921
	case 922:
		goto st_case_922
	case 218:
		goto st_case_218
	case 219:
		goto st_case_219
	case 220:
		goto st_case_220
	case 221:
		goto st_case_221
	case 923:
		goto st_case_923
	case 222:
		goto st_case_222
	case 223:
		goto st_case_223
	case 224:
		goto st_case_224
	case 924:
		goto st_case_924
	case 925:
		goto st_case_925
	case 225:
		goto st_case_225
	case 226:
		goto st_case_226
	case 227:
		goto st_case_227
	case 228:
		goto st_case_228
	case 229:
		goto st_case_229
	case 926:
		goto st_case_926
	case 927:
		goto st_case_927
	case 928:
		goto st_case_928
	case 929:
		goto st_case_929
	case 930:
		goto st_case_930
	case 230:
		goto st_case_230
	case 231:
		goto st_case_231
	case 232:
		goto st_case_232
	case 931:
		goto st_case_931
	case 233:
		goto st_case_233
	case 234:
		goto st_case_234
	case 235:
		goto st_case_235
	case 932:
		goto st_case_932
	case 236:
		goto st_case_236
	case 237:
		goto st_case_237
	case 933:
		goto st_case_933
	case 934:
		goto st_case_934
	case 238:
		goto st_case_238
	case 239:
		goto st_case_239
	case 935:
		goto st_case_935
	case 936:
		goto st_case_936
	case 937:
		goto st_case_937
	case 938:
		goto st_case_938
	case 939:
		goto st_case_939
	case 240:
		goto st_case_240
	case 241:
		goto st_case_241
	case 242:
		goto st_case_242
	case 940:
		goto st_case_940
	case 243:
		goto st_case_243
	case 244:
		goto st_case_244
	case 245:
		goto st_case_245
	case 941:
		goto st_case_941
	case 246:
		goto st_case_246
	case 247:
		goto st_case_247
	case 942:
		goto st_case_942
	case 943:
		goto st_case_943
	case 248:
		goto st_case_248
	case 249:
		goto st_case_249
	case 944:
		goto st_case_944
	case 945:
		goto st_case_945
	case 946:
		goto st_case_946
	case 947:
		goto st_case_947
	case 948:
		goto st_case_948
	case 250:
		goto st_case_250
	case 251:
		goto st_case_251
	case 252:
		goto st_case_252
	case 949:
		goto st_case_949
	case 253:
		goto st_case_253
	case 254:
		goto st_case_254
	case 255:
		goto st_case_255
	case 950:
		goto st_case_950
	case 256:
		goto st_case_256
	case 257:
		goto st_case_257
	case 951:
		goto st_case_951
	case 952:
		goto st_case_952
	case 258:
		goto st_case_258
	case 259:
		goto st_case_259
	case 260:
		goto st_case_260
	case 261:
		goto st_case_261
	case 262:
		goto st_case_262
	case 263:
		goto st_case_263
	case 953:
		goto st_case_953
	case 954:
		goto st_case_954
	case 955:
		goto st_case_955
	case 956:
		goto st_case_956
	case 264:
		goto st_case_264
	case 957:
		goto st_case_957
	case 958:
		goto st_case_958
	case 959:
		goto st_case_959
	case 960:
		goto st_case_960
	case 961:
		goto st_case_961
	case 962:
		goto st_case_962
	case 963:
		goto st_case_963
	case 964:
		goto st_case_964
	case 965:
		goto st_case_965
	case 966:
		goto st_case_966
	case 967:
		goto st_case_967
	case 968:
		goto st_case_968
	case 969:
		goto st_case_969
	case 970:
		goto st_case_970
	case 971:
		goto st_case_971
	case 972:
		goto st_case_972
	case 973:
		goto st_case_973
	case 974:
		goto st_case_974
	case 975:
		goto st_case_975
	case 976:
		goto st_case_976
	case 265:
		goto st_case_265
	case 266:
		goto st_case_266
	case 267:
		goto st_case_267
	case 268:
		goto st_case_268
	case 269:
		goto st_case_269
	case 270:
		goto st_case_270
	case 271:
		goto st_case_271
	case 272:
		goto st_case_272
	case 977:
		goto st_case_977
	case 978:
		goto st_case_978
	case 273:
		goto st_case_273
	case 979:
		goto st_case_979
	case 980:
		goto st_case_980
	case 981:
		goto st_case_981
	case 982:
		goto st_case_982
	case 983:
		goto st_case_983
	case 984:
		goto st_case_984
	case 985:
		goto st_case_985
	case 986:
		goto st_case_986
	case 987:
		goto st_case_987
	case 988:
		goto st_case_988
	case 989:
		goto st_case_989
	case 990:
		goto st_case_990
	case 991:
		goto st_case_991
	case 992:
		goto st_case_992
	case 993:
		goto st_case_993
	case 994:
		goto st_case_994
	case 995:
		goto st_case_995
	case 996:
		goto st_case_996
	case 997:
		goto st_case_997
	case 274:
		goto st_case_274
	case 275:
		goto st_case_275
	case 998:
		goto st_case_998
	case 999:
		goto st_case_999
	case 276:
		goto st_case_276
	case 1000:
		goto st_case_1000
	case 1001:
		goto st_case_1001
	case 1002:
		goto st_case_1002
	case 1003:
		goto st_case_1003
	case 1004:
		goto st_case_1004
	case 1005:
		goto st_case_1005
	case 1006:
		goto st_case_1006
	case 1007:
		goto st_case_1007
	case 1008:
		goto st_case_1008
	case 1009:
		goto st_case_1009
	case 1010:
		goto st_case_1010
	case 1011:
		goto st_case_1011
	case 1012:
		goto st_case_1012
	case 1013:
		goto st_case_1013
	case 1014:
		goto st_case_1014
	case 1015:
		goto st_case_1015
	case 1016:
		goto st_case_1016
	case 1017:
		goto st_case_1017
	case 1018:
		goto st_case_1018
	case 277:
		goto st_case_277
	case 1019:
		goto st_case_1019
	case 1020:
		goto st_case_1020
	case 278:
		goto st_case_278
	case 279:
		goto st_case_279
	case 280:
		goto st_case_280
	case 281:
		goto st_case_281
	case 1021:
		goto st_case_1021
	case 1022:
		goto st_case_1022
	case 1023:
		goto st_case_1023
	case 1024:
		goto st_case_1024
	case 1025:
		goto st_case_1025
	case 282:
		goto st_case_282
	case 283:
		goto st_case_283
	case 284:
		goto st_case_284
	case 1026:
		goto st_case_1026
	case 285:
		goto st_case_285
	case 286:
		goto st_case_286
	case 287:
		goto st_case_287
	case 1027:
		goto st_case_1027
	case 288:
		goto st_case_288
	case 289:
		goto st_case_289
	case 1028:
		goto st_case_1028
	case 1029:
		goto st_case_1029
	case 290:
		goto st_case_290
	case 291:
		goto st_case_291
	case 1030:
		goto st_case_1030
	case 1031:
		goto st_case_1031
	case 1032:
		goto st_case_1032
	case 292:
		goto st_case_292
	case 1033:
		goto st_case_1033
	case 1034:
		goto st_case_1034
	case 293:
		goto st_case_293
	case 1035:
		goto st_case_1035
	case 1036:
		goto st_case_1036
	case 1037:
		goto st_case_1037
	case 1038:
		goto st_case_1038
	case 1039:
		goto st_case_1039
	case 1040:
		goto st_case_1040
	case 1041:
		goto st_case_1041
	case 1042:
		goto st_case_1042
	case 1043:
		goto st_case_1043
	case 1044:
		goto st_case_1044
	case 1045:
		goto st_case_1045
	case 1046:
		goto st_case_1046
	case 1047:
		goto st_case_1047
	case 1048:
		goto st_case_1048
	case 1049:
		goto st_case_1049
	case 1050:
		goto st_case_1050
	case 1051:
		goto st_case_1051
	case 1052:
		goto st_case_1052
	case 1053:
		goto st_case_1053
	case 1054:
		goto st_case_1054
	case 294:
		goto st_case_294
	case 295:
		goto st_case_295
	case 1055:
		goto st_case_1055
	case 1056:
		goto st_case_1056
	case 296:
		goto st_case_296
	case 1057:
		goto st_case_1057
	case 1058:
		goto st_case_1058
	case 1059:
		goto st_case_1059
	case 1060:
		goto st_case_1060
	case 1061:
		goto st_case_1061
	case 1062:
		goto st_case_1062
	case 1063:
		goto st_case_1063
	case 1064:
		goto st_case_1064
	case 1065:
		goto st_case_1065
	case 1066:
		goto st_case_1066
	case 1067:
		goto st_case_1067
	case 1068:
		goto st_case_1068
	case 1069:
		goto st_case_1069
	case 1070:
		goto st_case_1070
	case 1071:
		goto st_case_1071
	case 1072:
		goto st_case_1072
	case 1073:
		goto st_case_1073
	case 1074:
		goto st_case_1074
	case 1075:
		goto st_case_1075
	case 1076:
		goto st_case_1076
	case 297:
		goto st_case_297
	case 1077:
		goto st_case_1077
	case 1078:
		goto st_case_1078
	case 1079:
		goto st_case_1079
	case 298:
		goto st_case_298
	case 1080:
		goto st_case_1080
	case 1081:
		goto st_case_1081
	case 1082:
		goto st_case_1082
	case 1083:
		goto st_case_1083
	case 1084:
		goto st_case_1084
	case 1085:
		goto st_case_1085
	case 1086:
		goto st_case_1086
	case 1087:
		goto st_case_1087
	case 1088:
		goto st_case_1088
	case 1089:
		goto st_case_1089
	case 1090:
		goto st_case_1090
	case 1091:
		goto st_case_1091
	case 1092:
		goto st_case_1092
	case 1093:
		goto st_case_1093
	case 1094:
		goto st_case_1094
	case 1095:
		goto st_case_1095
	case 1096:
		goto st_case_1096
	case 1097:
		goto st_case_1097
	case 1098:
		goto st_case_1098
	case 1099:
		goto st_case_1099
	case 1100:
		goto st_case_1100
	case 1101:
		goto st_case_1101
	case 1102:
		goto st_case_1102
	case 299:
		goto st_case_299
	case 1103:
		goto st_case_1103
	case 1104:
		goto st_case_1104
	case 1105:
		goto st_case_1105
	case 1106:
		goto st_case_1106
	case 1107:
		goto st_case_1107
	case 1108:
		goto st_case_1108
	case 1109:
		goto st_case_1109
	case 1110:
		goto st_case_1110
	case 1111:
		goto st_case_1111
	case 1112:
		goto st_case_1112
	case 1113:
		goto st_case_1113
	case 1114:
		goto st_case_1114
	case 1115:
		goto st_case_1115
	case 1116:
		goto st_case_1116
	case 1117:
		goto st_case_1117
	case 1118:
		goto st_case_1118
	case 1119:
		goto st_case_1119
	case 1120:
		goto st_case_1120
	case 1121:
		goto st_case_1121
	case 1122:
		goto st_case_1122
	case 1123:
		goto st_case_1123
	case 1124:
		goto st_case_1124
	case 300:
		goto st_case_300
	case 301:
		goto st_case_301
	case 1125:
		goto st_case_1125
	case 1126:
		goto st_case_1126
	case 302:
		goto st_case_302
	case 1127:
		goto st_case_1127
	case 1128:
		goto st_case_1128
	case 1129:
		goto st_case_1129
	case 1130:
		goto st_case_1130
	case 1131:
		goto st_case_1131
	case 1132:
		goto st_case_1132
	case 1133:
		goto st_case_1133
	case 1134:
		goto st_case_1134
	case 1135:
		goto st_case_1135
	case 1136:
		goto st_case_1136
	case 1137:
		goto st_case_1137
	case 1138:
		goto st_case_1138
	case 1139:
		goto st_case_1139
	case 1140:
		goto st_case_1140
	case 1141:
		goto st_case_1141
	case 1142:
		goto st_case_1142
	case 1143:
		goto st_case_1143
	case 1144:
		goto st_case_1144
	case 1145:
		goto st_case_1145
	case 303:
		goto st_case_303
	case 1146:
		goto st_case_1146
	case 1147:
		goto st_case_1147
	case 1148:
		goto st_case_1148
	case 304:
		goto st_case_304
	case 1149:
		goto st_case_1149
	case 1150:
		goto st_case_1150
	case 305:
		goto st_case_305
	case 1151:
		goto st_case_1151
	case 1152:
		goto st_case_1152
	case 1153:
		goto st_case_1153
	case 1154:
		goto st_case_1154
	case 1155:
		goto st_case_1155
	case 1156:
		goto st_case_1156
	case 1157:
		goto st_case_1157
	case 1158:
		goto st_case_1158
	case 1159:
		goto st_case_1159
	case 1160:
		goto st_case_1160
	case 1161:
		goto st_case_1161
	case 1162:
		goto st_case_1162
	case 1163:
		goto st_case_1163
	case 1164:
		goto st_case_1164
	case 1165:
		goto st_case_1165
	case 1166:
		goto st_case_1166
	case 1167:
		goto st_case_1167
	case 1168:
		goto st_case_1168
	case 1169:
		goto st_case_1169
	case 306:
		goto st_case_306
	case 307:
		goto st_case_307
	case 308:
		goto st_case_308
	case 309:
		goto st_case_309
	case 1170:
		goto st_case_1170
	case 1171:
		goto st_case_1171
	case 1172:
		goto st_case_1172
	case 1173:
		goto st_case_1173
	case 1174:
		goto st_case_1174
	case 310:
		goto st_case_310
	case 311:
		goto st_case_311
	case 312:
		goto st_case_312
	case 1175:
		goto st_case_1175
	case 313:
		goto st_case_313
	case 314:
		goto st_case_314
	case 315:
		goto st_case_315
	case 1176:
		goto st_case_1176
	case 316:
		goto st_case_316
	case 317:
		goto st_case_317
	case 1177:
		goto st_case_1177
	case 1178:
		goto st_case_1178
	case 318:
		goto st_case_318
	case 319:
		goto st_case_319
	case 1179:
		goto st_case_1179
	case 1180:
		goto st_case_1180
	case 1181:
		goto st_case_1181
	case 1182:
		goto st_case_1182
	case 1183:
		goto st_case_1183
	case 320:
		goto st_case_320
	case 321:
		goto st_case_321
	case 322:
		goto st_case_322
	case 1184:
		goto st_case_1184
	case 323:
		goto st_case_323
	case 324:
		goto st_case_324
	case 325:
		goto st_case_325
	case 1185:
		goto st_case_1185
	case 326:
		goto st_case_326
	case 327:
		goto st_case_327
	case 1186:
		goto st_case_1186
	case 1187:
		goto st_case_1187
	case 328:
		goto st_case_328
	case 329:
		goto st_case_329
	case 1188:
		goto st_case_1188
	case 1189:
		goto st_case_1189
	case 1190:
		goto st_case_1190
	case 330:
		goto st_case_330
	case 1191:
		goto st_case_1191
	case 1192:
		goto st_case_1192
	case 1193:
		goto st_case_1193
	case 1194:
		goto st_case_1194
	case 1195:
		goto st_case_1195
	case 1196:
		goto st_case_1196
	case 1197:
		goto st_case_1197
	case 1198:
		goto st_case_1198
	case 1199:
		goto st_case_1199
	case 1200:
		goto st_case_1200
	case 1201:
		goto st_case_1201
	case 1202:
		goto st_case_1202
	case 1203:
		goto st_case_1203
	case 1204:
		goto st_case_1204
	case 1205:
		goto st_case_1205
	case 1206:
		goto st_case_1206
	case 1207:
		goto st_case_1207
	case 1208:
		goto st_case_1208
	case 1209:
		goto st_case_1209
	case 1210:
		goto st_case_1210
	case 1211:
		goto st_case_1211
	case 1212:
		goto st_case_1212
	case 1213:
		goto st_case_1213
	case 331:
		goto st_case_331
	case 332:
		goto st_case_332
	case 333:
		goto st_case_333
	case 1214:
		goto st_case_1214
	case 334:
		goto st_case_334
	case 335:
		goto st_case_335
	case 336:
		goto st_case_336
	case 1215:
		goto st_case_1215
	case 337:
		goto st_case_337
	case 338:
		goto st_case_338
	case 1216:
		goto st_case_1216
	case 1217:
		goto st_case_1217
	case 339:
		goto st_case_339
	case 340:
		goto st_case_340
	case 1218:
		goto st_case_1218
	case 1219:
		goto st_case_1219
	case 1220:
		goto st_case_1220
	case 1221:
		goto st_case_1221
	case 1222:
		goto st_case_1222
	case 341:
		goto st_case_341
	case 1223:
		goto st_case_1223
	case 1224:
		goto st_case_1224
	case 342:
		goto st_case_342
	case 1225:
		goto st_case_1225
	case 1226:
		goto st_case_1226
	case 1227:
		goto st_case_1227
	case 1228:
		goto st_case_1228
	case 1229:
		goto st_case_1229
	case 1230:
		goto st_case_1230
	case 1231:
		goto st_case_1231
	case 1232:
		goto st_case_1232
	case 1233:
		goto st_case_1233
	case 1234:
		goto st_case_1234
	case 1235:
		goto st_case_1235
	case 1236:
		goto st_case_1236
	case 1237:
		goto st_case_1237
	case 1238:
		goto st_case_1238
	case 1239:
		goto st_case_1239
	case 1240:
		goto st_case_1240
	case 1241:
		goto st_case_1241
	case 1242:
		goto st_case_1242
	case 1243:
		goto st_case_1243
	case 343:
		goto st_case_343
	case 1244:
		goto st_case_1244
	case 1245:
		goto st_case_1245
	case 1246:
		goto st_case_1246
	case 344:
		goto st_case_344
	case 345:
		goto st_case_345
	case 346:
		goto st_case_346
	case 1247:
		goto st_case_1247
	case 347:
		goto st_case_347
	case 348:
		goto st_case_348
	case 349:
		goto st_case_349
	case 1248:
		goto st_case_1248
	case 350:
		goto st_case_350
	case 351:
		goto st_case_351
	case 1249:
		goto st_case_1249
	case 1250:
		goto st_case_1250
	case 352:
		goto st_case_352
	case 353:
		goto st_case_353
	case 354:
		goto st_case_354
	case 355:
		goto st_case_355
	case 356:
		goto st_case_356
	case 357:
		goto st_case_357
	case 1251:
		goto st_case_1251
	case 1252:
		goto st_case_1252
	case 1253:
		goto st_case_1253
	case 1254:
		goto st_case_1254
	case 1255:
		goto st_case_1255
	case 1256:
		goto st_case_1256
	case 1257:
		goto st_case_1257
	case 1258:
		goto st_case_1258
	case 1259:
		goto st_case_1259
	case 1260:
		goto st_case_1260
	case 1261:
		goto st_case_1261
	case 1262:
		goto st_case_1262
	case 1263:
		goto st_case_1263
	case 1264:
		goto st_case_1264
	case 1265:
		goto st_case_1265
	case 1266:
		goto st_case_1266
	case 1267:
		goto st_case_1267
	case 1268:
		goto st_case_1268
	case 1269:
		goto st_case_1269
	case 358:
		goto st_case_358
	case 359:
		goto st_case_359
	case 360:
		goto st_case_360
	case 1270:
		goto st_case_1270
	case 361:
		goto st_case_361
	case 362:
		goto st_case_362
	case 363:
		goto st_case_363
	case 364:
		goto st_case_364
	case 1271:
		goto st_case_1271
	case 1272:
		goto st_case_1272
	case 1273:
		goto st_case_1273
	case 1274:
		goto st_case_1274
	case 1275:
		goto st_case_1275
	case 365:
		goto st_case_365
	case 366:
		goto st_case_366
	case 367:
		goto st_case_367
	case 1276:
		goto st_case_1276
	case 368:
		goto st_case_368
	case 369:
		goto st_case_369
	case 370:
		goto st_case_370
	case 1277:
		goto st_case_1277
	case 371:
		goto st_case_371
	case 372:
		goto st_case_372
	case 1278:
		goto st_case_1278
	case 1279:
		goto st_case_1279
	case 373:
		goto st_case_373
	case 374:
		goto st_case_374
	case 1280:
		goto st_case_1280
	case 1281:
		goto st_case_1281
	case 1282:
		goto st_case_1282
	case 1283:
		goto st_case_1283
	case 1284:
		goto st_case_1284
	case 375:
		goto st_case_375
	case 376:
		goto st_case_376
	case 377:
		goto st_case_377
	case 1285:
		goto st_case_1285
	case 378:
		goto st_case_378
	case 379:
		goto st_case_379
	case 380:
		goto st_case_380
	case 1286:
		goto st_case_1286
	case 381:
		goto st_case_381
	case 382:
		goto st_case_382
	case 1287:
		goto st_case_1287
	case 1288:
		goto st_case_1288
	case 383:
		goto st_case_383
	case 384:
		goto st_case_384
	case 385:
		goto st_case_385
	case 386:
		goto st_case_386
	case 387:
		goto st_case_387
	case 388:
		goto st_case_388
	case 1289:
		goto st_case_1289
	case 1290:
		goto st_case_1290
	case 389:
		goto st_case_389
	case 1291:
		goto st_case_1291
	case 1292:
		goto st_case_1292
	case 390:
		goto st_case_390
	case 1293:
		goto st_case_1293
	case 1294:
		goto st_case_1294
	case 1295:
		goto st_case_1295
	case 1296:
		goto st_case_1296
	case 1297:
		goto st_case_1297
	case 1298:
		goto st_case_1298
	case 1299:
		goto st_case_1299
	case 1300:
		goto st_case_1300
	case 1301:
		goto st_case_1301
	case 1302:
		goto st_case_1302
	case 1303:
		goto st_case_1303
	case 1304:
		goto st_case_1304
	case 1305:
		goto st_case_1305
	case 1306:
		goto st_case_1306
	case 1307:
		goto st_case_1307
	case 1308:
		goto st_case_1308
	case 1309:
		goto st_case_1309
	case 1310:
		goto st_case_1310
	case 391:
		goto st_case_391
	case 392:
		goto st_case_392
	case 393:
		goto st_case_393
	case 394:
		goto st_case_394
	case 395:
		goto st_case_395
	case 396:
		goto st_case_396
	case 397:
		goto st_case_397
	case 398:
		goto st_case_398
	case 399:
		goto st_case_399
	case 400:
		goto st_case_400
	case 401:
		goto st_case_401
	case 402:
		goto st_case_402
	case 403:
		goto st_case_403
	case 1311:
		goto st_case_1311
	case 404:
		goto st_case_404
	case 405:
		goto st_case_405
	case 406:
		goto st_case_406
	case 1312:
		goto st_case_1312
	case 407:
		goto st_case_407
	case 408:
		goto st_case_408
	case 409:
		goto st_case_409
	case 410:
		goto st_case_410
	case 1313:
		goto st_case_1313
	case 1314:
		goto st_case_1314
	case 1315:
		goto st_case_1315
	case 1316:
		goto st_case_1316
	case 1317:
		goto st_case_1317
	case 411:
		goto st_case_411
	case 412:
		goto st_case_412
	case 413:
		goto st_case_413
	case 1318:
		goto st_case_1318
	case 414:
		goto st_case_414
	case 415:
		goto st_case_415
	case 416:
		goto st_case_416
	case 1319:
		goto st_case_1319
	case 417:
		goto st_case_417
	case 418:
		goto st_case_418
	case 1320:
		goto st_case_1320
	case 1321:
		goto st_case_1321
	case 419:
		goto st_case_419
	case 420:
		goto st_case_420
	case 421:
		goto st_case_421
	case 422:
		goto st_case_422
	case 423:
		goto st_case_423
	case 424:
		goto st_case_424
	case 425:
		goto st_case_425
	case 426:
		goto st_case_426
	case 427:
		goto st_case_427
	case 428:
		goto st_case_428
	case 429:
		goto st_case_429
	case 430:
		goto st_case_430
	case 1322:
		goto st_case_1322
	case 1323:
		goto st_case_1323
	case 1324:
		goto st_case_1324
	case 431:
		goto st_case_431
	case 1325:
		goto st_case_1325
	case 1326:
		goto st_case_1326
	case 1327:
		goto st_case_1327
	case 1328:
		goto st_case_1328
	case 1329:
		goto st_case_1329
	case 1330:
		goto st_case_1330
	case 1331:
		goto st_case_1331
	case 1332:
		goto st_case_1332
	case 1333:
		goto st_case_1333
	case 1334:
		goto st_case_1334
	case 1335:
		goto st_case_1335
	case 1336:
		goto st_case_1336
	case 1337:
		goto st_case_1337
	case 1338:
		goto st_case_1338
	case 1339:
		goto st_case_1339
	case 1340:
		goto st_case_1340
	case 1341:
		goto st_case_1341
	case 1342:
		goto st_case_1342
	case 1343:
		goto st_case_1343
	case 1344:
		goto st_case_1344
	case 1345:
		goto st_case_1345
	case 1346:
		goto st_case_1346
	case 432:
		goto st_case_432
	case 1347:
		goto st_case_1347
	case 1348:
		goto st_case_1348
	case 1349:
		goto st_case_1349
	case 1350:
		goto st_case_1350
	case 433:
		goto st_case_433
	case 1351:
		goto st_case_1351
	case 1352:
		goto st_case_1352
	case 1353:
		goto st_case_1353
	case 1354:
		goto st_case_1354
	case 1355:
		goto st_case_1355
	case 1356:
		goto st_case_1356
	case 1357:
		goto st_case_1357
	case 1358:
		goto st_case_1358
	case 1359:
		goto st_case_1359
	case 1360:
		goto st_case_1360
	case 1361:
		goto st_case_1361
	case 1362:
		goto st_case_1362
	case 1363:
		goto st_case_1363
	case 1364:
		goto st_case_1364
	case 1365:
		goto st_case_1365
	case 1366:
		goto st_case_1366
	case 1367:
		goto st_case_1367
	case 1368:
		goto st_case_1368
	case 1369:
		goto st_case_1369
	case 1370:
		goto st_case_1370
	case 434:
		goto st_case_434
	case 435:
		goto st_case_435
	case 436:
		goto st_case_436
	case 1371:
		goto st_case_1371
	case 1372:
		goto st_case_1372
	case 437:
		goto st_case_437
	case 1373:
		goto st_case_1373
	case 1374:
		goto st_case_1374
	case 1375:
		goto st_case_1375
	case 1376:
		goto st_case_1376
	case 1377:
		goto st_case_1377
	case 1378:
		goto st_case_1378
	case 1379:
		goto st_case_1379
	case 1380:
		goto st_case_1380
	case 1381:
		goto st_case_1381
	case 1382:
		goto st_case_1382
	case 1383:
		goto st_case_1383
	case 1384:
		goto st_case_1384
	case 1385:
		goto st_case_1385
	case 1386:
		goto st_case_1386
	case 1387:
		goto st_case_1387
	case 1388:
		goto st_case_1388
	case 1389:
		goto st_case_1389
	case 1390:
		goto st_case_1390
	case 1391:
		goto st_case_1391
	case 1392:
		goto st_case_1392
	case 438:
		goto st_case_438
	case 1393:
		goto st_case_1393
	case 1394:
		goto st_case_1394
	case 1395:
		goto st_case_1395
	case 439:
		goto st_case_439
	case 1396:
		goto st_case_1396
	case 1397:
		goto st_case_1397
	case 1398:
		goto st_case_1398
	case 1399:
		goto st_case_1399
	case 1400:
		goto st_case_1400
	case 1401:
		goto st_case_1401
	case 1402:
		goto st_case_1402
	case 1403:
		goto st_case_1403
	case 1404:
		goto st_case_1404
	case 1405:
		goto st_case_1405
	case 1406:
		goto st_case_1406
	case 1407:
		goto st_case_1407
	case 1408:
		goto st_case_1408
	case 1409:
		goto st_case_1409
	case 1410:
		goto st_case_1410
	case 1411:
		goto st_case_1411
	case 1412:
		goto st_case_1412
	case 1413:
		goto st_case_1413
	case 1414:
		goto st_case_1414
	case 1415:
		goto st_case_1415
	case 1416:
		goto st_case_1416
	case 1417:
		goto st_case_1417
	case 440:
		goto st_case_440
	case 1418:
		goto st_case_1418
	case 1419:
		goto st_case_1419
	case 1420:
		goto st_case_1420
	case 1421:
		goto st_case_1421
	case 1422:
		goto st_case_1422
	case 1423:
		goto st_case_1423
	case 1424:
		goto st_case_1424
	case 1425:
		goto st_case_1425
	case 1426:
		goto st_case_1426
	case 1427:
		goto st_case_1427
	case 1428:
		goto st_case_1428
	case 1429:
		goto st_case_1429
	case 1430:
		goto st_case_1430
	case 1431:
		goto st_case_1431
	case 1432:
		goto st_case_1432
	case 1433:
		goto st_case_1433
	case 1434:
		goto st_case_1434
	case 1435:
		goto st_case_1435
	case 1436:
		goto st_case_1436
	case 1437:
		goto st_case_1437
	case 1438:
		goto st_case_1438
	case 1439:
		goto st_case_1439
	case 441:
		goto st_case_441
	case 442:
		goto st_case_442
	case 443:
		goto st_case_443
	case 444:
		goto st_case_444
	case 1440:
		goto st_case_1440
	case 1441:
		goto st_case_1441
	case 1442:
		goto st_case_1442
	case 1443:
		goto st_case_1443
	case 1444:
		goto st_case_1444
	case 445:
		goto st_case_445
	case 446:
		goto st_case_446
	case 447:
		goto st_case_447
	case 1445:
		goto st_case_1445
	case 448:
		goto st_case_448
	case 449:
		goto st_case_449
	case 450:
		goto st_case_450
	case 1446:
		goto st_case_1446
	case 451:
		goto st_case_451
	case 452:
		goto st_case_452
	case 1447:
		goto st_case_1447
	case 1448:
		goto st_case_1448
	case 453:
		goto st_case_453
	case 454:
		goto st_case_454
	case 1449:
		goto st_case_1449
	case 1450:
		goto st_case_1450
	case 1451:
		goto st_case_1451
	case 455:
		goto st_case_455
	case 1452:
		goto st_case_1452
	case 1453:
		goto st_case_1453
	case 1454:
		goto st_case_1454
	case 1455:
		goto st_case_1455
	case 1456:
		goto st_case_1456
	case 1457:
		goto st_case_1457
	case 1458:
		goto st_case_1458
	case 1459:
		goto st_case_1459
	case 1460:
		goto st_case_1460
	case 1461:
		goto st_case_1461
	case 1462:
		goto st_case_1462
	case 1463:
		goto st_case_1463
	case 1464:
		goto st_case_1464
	case 1465:
		goto st_case_1465
	case 1466:
		goto st_case_1466
	case 1467:
		goto st_case_1467
	case 1468:
		goto st_case_1468
	case 1469:
		goto st_case_1469
	case 1470:
		goto st_case_1470
	case 1471:
		goto st_case_1471
	case 1472:
		goto st_case_1472
	case 1473:
		goto st_case_1473
	case 1474:
		goto st_case_1474
	case 1475:
		goto st_case_1475
	case 456:
		goto st_case_456
	case 457:
		goto st_case_457
	case 458:
		goto st_case_458
	case 1476:
		goto st_case_1476
	case 459:
		goto st_case_459
	case 460:
		goto st_case_460
	case 461:
		goto st_case_461
	case 1477:
		goto st_case_1477
	case 462:
		goto st_case_462
	case 463:
		goto st_case_463
	case 1478:
		goto st_case_1478
	case 1479:
		goto st_case_1479
	case 464:
		goto st_case_464
	case 465:
		goto st_case_465
	case 466:
		goto st_case_466
	case 1480:
		goto st_case_1480
	case 1481:
		goto st_case_1481
	case 1482:
		goto st_case_1482
	case 1483:
		goto st_case_1483
	case 1484:
		goto st_case_1484
	case 467:
		goto st_case_467
	case 468:
		goto st_case_468
	case 469:
		goto st_case_469
	case 1485:
		goto st_case_1485
	case 470:
		goto st_case_470
	case 471:
		goto st_case_471
	case 472:
		goto st_case_472
	case 1486:
		goto st_case_1486
	case 473:
		goto st_case_473
	case 474:
		goto st_case_474
	case 1487:
		goto st_case_1487
	case 1488:
		goto st_case_1488
	case 475:
		goto st_case_475
	case 1489:
		goto st_case_1489
	case 1490:
		goto st_case_1490
	case 1491:
		goto st_case_1491
	case 1492:
		goto st_case_1492
	case 1493:
		goto st_case_1493
	case 476:
		goto st_case_476
	case 477:
		goto st_case_477
	case 478:
		goto st_case_478
	case 1494:
		goto st_case_1494
	case 479:
		goto st_case_479
	case 480:
		goto st_case_480
	case 481:
		goto st_case_481
	case 1495:
		goto st_case_1495
	case 482:
		goto st_case_482
	case 483:
		goto st_case_483
	case 1496:
		goto st_case_1496
	case 1497:
		goto st_case_1497
	case 484:
		goto st_case_484
	case 485:
		goto st_case_485
	case 486:
		goto st_case_486
	case 487:
		goto st_case_487
	case 488:
		goto st_case_488
	case 489:
		goto st_case_489
	case 1498:
		goto st_case_1498
	}
	goto st_out
	st490:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof490
		}
	st_case_490:
		switch ( m.data)[( m.p)] {
		case 10:
			goto st0
		case 11:
			goto tr839
		case 13:
			goto st0
		case 32:
			goto st484
		case 44:
			goto st0
		case 92:
			goto tr840
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto st484
		}
		goto tr837
tr837:
//line machine.rl:17

	m.pb = m.p

	goto st1
	st1:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1
		}
	st_case_1:
//line machine.go:6145
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr2
		case 11:
			goto tr3
		case 13:
			goto tr2
		case 32:
			goto tr1
		case 44:
			goto tr4
		case 92:
			goto st434
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1
		}
		goto st1
tr1:
//line machine.rl:58

	m.handler.SetMeasurement(m.text())

	goto st2
tr47:
//line machine.rl:66

	m.handler.AddTag(key, m.text())

	goto st2
	st2:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof2
		}
	st_case_2:
//line machine.go:6181
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr2
		case 11:
			goto tr8
		case 13:
			goto tr2
		case 32:
			goto st2
		case 44:
			goto tr2
		case 61:
			goto tr2
		case 92:
			goto tr9
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto st2
		}
		goto tr6
tr6:
//line machine.rl:17

	m.pb = m.p

	goto st3
	st3:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof3
		}
	st_case_3:
//line machine.go:6213
		switch ( m.data)[( m.p)] {
		case 32:
			goto tr2
		case 44:
			goto tr2
		case 61:
			goto tr11
		case 92:
			goto st27
		}
		switch {
		case ( m.data)[( m.p)] > 10:
			if 12 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto tr2
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr2
		}
		goto st3
tr2:
	 m.cs = 0
//line machine.rl:26

	m.err = ErrFieldParse
	( m.p)--

	 m.cs = 489;
	{( m.p)++; goto _out }

	goto _again
tr39:
	 m.cs = 0
//line machine.rl:33

	m.err = ErrTagParse
	( m.p)--

	 m.cs = 489;
	{( m.p)++; goto _out }

	goto _again
tr48:
	 m.cs = 0
//line machine.rl:33

	m.err = ErrTagParse
	( m.p)--

	 m.cs = 489;
	{( m.p)++; goto _out }

//line machine.rl:26

	m.err = ErrFieldParse
	( m.p)--

	 m.cs = 489;
	{( m.p)++; goto _out }

	goto _again
tr88:
	 m.cs = 0
//line machine.rl:26

	m.err = ErrFieldParse
	( m.p)--

	 m.cs = 489;
	{( m.p)++; goto _out }

//line machine.rl:40

	m.err = ErrTimestampParse
	( m.p)--

	 m.cs = 489;
	{( m.p)++; goto _out }

	goto _again
tr177:
	 m.cs = 0
//line machine.rl:33

	m.err = ErrTagParse
	( m.p)--

	 m.cs = 489;
	{( m.p)++; goto _out }

//line machine.rl:40

	m.err = ErrTimestampParse
	( m.p)--

	 m.cs = 489;
	{( m.p)++; goto _out }

	goto _again
tr179:
	 m.cs = 0
//line machine.rl:33

	m.err = ErrTagParse
	( m.p)--

	 m.cs = 489;
	{( m.p)++; goto _out }

//line machine.rl:26

	m.err = ErrFieldParse
	( m.p)--

	 m.cs = 489;
	{( m.p)++; goto _out }

//line machine.rl:40

	m.err = ErrTimestampParse
	( m.p)--

	 m.cs = 489;
	{( m.p)++; goto _out }

	goto _again
tr351:
	 m.cs = 0
//line machine.rl:40

	m.err = ErrTimestampParse
	( m.p)--

	 m.cs = 489;
	{( m.p)++; goto _out }

	goto _again
tr419:
	 m.cs = 0
//line machine.rl:40

	m.err = ErrTimestampParse
	( m.p)--

	 m.cs = 489;
	{( m.p)++; goto _out }

//line machine.rl:26

	m.err = ErrFieldParse
	( m.p)--

	 m.cs = 489;
	{( m.p)++; goto _out }

	goto _again
tr495:
	 m.cs = 0
//line machine.rl:26

	m.err = ErrFieldParse
	( m.p)--

	 m.cs = 489;
	{( m.p)++; goto _out }

//line machine.rl:33

	m.err = ErrTagParse
	( m.p)--

	 m.cs = 489;
	{( m.p)++; goto _out }

	goto _again
tr529:
	 m.cs = 0
//line machine.rl:40

	m.err = ErrTimestampParse
	( m.p)--

	 m.cs = 489;
	{( m.p)++; goto _out }

//line machine.rl:33

	m.err = ErrTagParse
	( m.p)--

	 m.cs = 489;
	{( m.p)++; goto _out }

//line machine.rl:26

	m.err = ErrFieldParse
	( m.p)--

	 m.cs = 489;
	{( m.p)++; goto _out }

	goto _again
tr538:
	 m.cs = 0
//line machine.rl:40

	m.err = ErrTimestampParse
	( m.p)--

	 m.cs = 489;
	{( m.p)++; goto _out }

//line machine.rl:33

	m.err = ErrTagParse
	( m.p)--

	 m.cs = 489;
	{( m.p)++; goto _out }

	goto _again
tr565:
	 m.cs = 0
//line machine.rl:26

	m.err = ErrFieldParse
	( m.p)--

	 m.cs = 489;
	{( m.p)++; goto _out }

//line machine.rl:33

	m.err = ErrTagParse
	( m.p)--

	 m.cs = 489;
	{( m.p)++; goto _out }

//line machine.rl:40

	m.err = ErrTimestampParse
	( m.p)--

	 m.cs = 489;
	{( m.p)++; goto _out }

	goto _again
//line machine.go:6461
st_case_0:
	st0:
		 m.cs = 0
		goto _out
tr11:
//line machine.rl:70

	key = m.text()

	goto st4
	st4:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof4
		}
	st_case_4:
//line machine.go:6477
		switch ( m.data)[( m.p)] {
		case 34:
			goto st5
		case 45:
			goto tr14
		case 46:
			goto tr15
		case 48:
			goto tr16
		case 70:
			goto tr18
		case 84:
			goto tr19
		case 102:
			goto tr20
		case 116:
			goto tr21
		}
		if 49 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
			goto tr17
		}
		goto tr2
	st5:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof5
		}
	st_case_5:
		switch ( m.data)[( m.p)] {
		case 34:
			goto tr23
		case 92:
			goto tr24
		}
		goto tr22
tr22:
//line machine.rl:17

	m.pb = m.p

	goto st6
	st6:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof6
		}
	st_case_6:
//line machine.go:6523
		switch ( m.data)[( m.p)] {
		case 34:
			goto tr26
		case 92:
			goto st59
		}
		goto st6
tr23:
//line machine.rl:17

	m.pb = m.p

//line machine.rl:86

	m.handler.AddString(key, m.text())

	goto st491
tr26:
//line machine.rl:86

	m.handler.AddString(key, m.text())

	goto st491
	st491:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof491
		}
	st_case_491:
//line machine.go:6552
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr850
		case 13:
			goto tr850
		case 32:
			goto st492
		case 44:
			goto st28
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto st492
		}
		goto tr2
tr1350:
//line machine.rl:78

	m.handler.AddFloat(key, m.text())

	goto st492
tr1354:
//line machine.rl:74

	m.handler.AddInt(key, m.text())

	goto st492
tr1356:
//line machine.rl:82

	m.handler.AddBool(key, m.text())

	goto st492
	st492:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof492
		}
	st_case_492:
//line machine.go:6590
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr850
		case 13:
			goto tr850
		case 32:
			goto st492
		case 45:
			goto tr852
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto tr853
			}
		case ( m.data)[( m.p)] >= 9:
			goto st492
		}
		goto tr351
tr850:
//line machine.rl:21

	yield = true
	{( m.p)++;  m.cs = 493; goto _out }

	goto st493
tr859:
//line machine.rl:90

	m.handler.SetTimestamp(m.text())

//line machine.rl:21

	yield = true
	{( m.p)++;  m.cs = 493; goto _out }

	goto st493
tr1270:
//line machine.rl:78

	m.handler.AddFloat(key, m.text())

//line machine.rl:21

	yield = true
	{( m.p)++;  m.cs = 493; goto _out }

	goto st493
tr1276:
//line machine.rl:74

	m.handler.AddInt(key, m.text())

//line machine.rl:21

	yield = true
	{( m.p)++;  m.cs = 493; goto _out }

	goto st493
tr1280:
//line machine.rl:82

	m.handler.AddBool(key, m.text())

//line machine.rl:21

	yield = true
	{( m.p)++;  m.cs = 493; goto _out }

	goto st493
	st493:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof493
		}
	st_case_493:
//line machine.go:6666
		switch ( m.data)[( m.p)] {
		case 10:
			goto st493
		case 11:
			goto tr343
		case 13:
			goto st493
		case 32:
			goto st188
		case 44:
			goto st0
		case 92:
			goto tr344
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto st188
		}
		goto tr341
tr341:
//line machine.rl:17

	m.pb = m.p

	goto st7
	st7:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof7
		}
	st_case_7:
//line machine.go:6696
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr2
		case 11:
			goto tr29
		case 13:
			goto tr2
		case 32:
			goto tr1
		case 44:
			goto tr4
		case 92:
			goto st76
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1
		}
		goto st7
tr29:
//line machine.rl:58

	m.handler.SetMeasurement(m.text())

	goto st8
	st8:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof8
		}
	st_case_8:
//line machine.go:6726
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr2
		case 11:
			goto tr32
		case 13:
			goto tr2
		case 32:
			goto tr1
		case 44:
			goto tr4
		case 61:
			goto st7
		case 92:
			goto tr33
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1
		}
		goto tr31
tr31:
//line machine.rl:17

	m.pb = m.p

	goto st9
	st9:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof9
		}
	st_case_9:
//line machine.go:6758
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr2
		case 11:
			goto tr35
		case 13:
			goto tr2
		case 32:
			goto tr1
		case 44:
			goto tr4
		case 61:
			goto tr36
		case 92:
			goto st36
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1
		}
		goto st9
tr35:
//line machine.rl:58

	m.handler.SetMeasurement(m.text())

	goto st10
tr32:
//line machine.rl:58

	m.handler.SetMeasurement(m.text())

//line machine.rl:17

	m.pb = m.p

	goto st10
	st10:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof10
		}
	st_case_10:
//line machine.go:6800
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr2
		case 11:
			goto tr32
		case 13:
			goto tr2
		case 32:
			goto tr1
		case 44:
			goto tr4
		case 61:
			goto tr36
		case 92:
			goto tr33
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1
		}
		goto tr31
tr4:
//line machine.rl:58

	m.handler.SetMeasurement(m.text())

	goto st11
tr50:
//line machine.rl:66

	m.handler.AddTag(key, m.text())

	goto st11
	st11:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof11
		}
	st_case_11:
//line machine.go:6838
		switch ( m.data)[( m.p)] {
		case 32:
			goto tr39
		case 44:
			goto tr39
		case 61:
			goto tr39
		case 92:
			goto tr40
		}
		switch {
		case ( m.data)[( m.p)] > 10:
			if 12 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto tr39
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr39
		}
		goto tr38
tr38:
//line machine.rl:17

	m.pb = m.p

	goto st12
	st12:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof12
		}
	st_case_12:
//line machine.go:6869
		switch ( m.data)[( m.p)] {
		case 32:
			goto tr39
		case 44:
			goto tr39
		case 61:
			goto tr42
		case 92:
			goto st20
		}
		switch {
		case ( m.data)[( m.p)] > 10:
			if 12 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto tr39
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr39
		}
		goto st12
tr42:
//line machine.rl:62

	key = m.text()

	goto st13
	st13:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof13
		}
	st_case_13:
//line machine.go:6900
		switch ( m.data)[( m.p)] {
		case 32:
			goto tr39
		case 44:
			goto tr39
		case 61:
			goto tr39
		case 92:
			goto tr45
		}
		switch {
		case ( m.data)[( m.p)] > 10:
			if 12 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto tr39
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr39
		}
		goto tr44
tr44:
//line machine.rl:17

	m.pb = m.p

	goto st14
	st14:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof14
		}
	st_case_14:
//line machine.go:6931
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr48
		case 11:
			goto tr49
		case 13:
			goto tr48
		case 32:
			goto tr47
		case 44:
			goto tr50
		case 61:
			goto tr48
		case 92:
			goto st19
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr47
		}
		goto st14
tr49:
//line machine.rl:66

	m.handler.AddTag(key, m.text())

	goto st15
	st15:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof15
		}
	st_case_15:
//line machine.go:6963
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr48
		case 11:
			goto tr53
		case 13:
			goto tr48
		case 32:
			goto tr47
		case 44:
			goto tr50
		case 61:
			goto tr48
		case 92:
			goto tr54
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr47
		}
		goto tr52
tr52:
//line machine.rl:17

	m.pb = m.p

	goto st16
	st16:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof16
		}
	st_case_16:
//line machine.go:6995
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr48
		case 11:
			goto tr56
		case 13:
			goto tr48
		case 32:
			goto tr47
		case 44:
			goto tr50
		case 61:
			goto tr11
		case 92:
			goto st18
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr47
		}
		goto st16
tr56:
//line machine.rl:66

	m.handler.AddTag(key, m.text())

	goto st17
tr53:
//line machine.rl:66

	m.handler.AddTag(key, m.text())

//line machine.rl:17

	m.pb = m.p

	goto st17
	st17:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof17
		}
	st_case_17:
//line machine.go:7037
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr48
		case 11:
			goto tr53
		case 13:
			goto tr48
		case 32:
			goto tr47
		case 44:
			goto tr50
		case 61:
			goto tr11
		case 92:
			goto tr54
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr47
		}
		goto tr52
tr54:
//line machine.rl:17

	m.pb = m.p

	goto st18
	st18:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof18
		}
	st_case_18:
//line machine.go:7069
		switch {
		case ( m.data)[( m.p)] > 10:
			if 12 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto tr48
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr48
		}
		goto st16
tr45:
//line machine.rl:17

	m.pb = m.p

	goto st19
	st19:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof19
		}
	st_case_19:
//line machine.go:7090
		switch {
		case ( m.data)[( m.p)] > 10:
			if 12 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto tr39
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr39
		}
		goto st14
tr40:
//line machine.rl:17

	m.pb = m.p

	goto st20
	st20:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof20
		}
	st_case_20:
//line machine.go:7111
		switch {
		case ( m.data)[( m.p)] > 10:
			if 12 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto tr39
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr39
		}
		goto st12
tr36:
//line machine.rl:70

	key = m.text()

	goto st21
tr350:
//line machine.rl:17

	m.pb = m.p

//line machine.rl:70

	key = m.text()

	goto st21
	st21:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof21
		}
	st_case_21:
//line machine.go:7142
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr2
		case 11:
			goto tr29
		case 13:
			goto tr2
		case 32:
			goto tr1
		case 34:
			goto st22
		case 44:
			goto tr4
		case 45:
			goto tr59
		case 46:
			goto tr60
		case 48:
			goto tr61
		case 70:
			goto tr63
		case 84:
			goto tr64
		case 92:
			goto st76
		case 102:
			goto tr65
		case 116:
			goto tr66
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 49 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto tr62
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1
		}
		goto st7
	st22:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof22
		}
	st_case_22:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr22
		case 11:
			goto tr69
		case 13:
			goto tr22
		case 32:
			goto tr68
		case 34:
			goto tr70
		case 44:
			goto tr71
		case 92:
			goto tr72
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr68
		}
		goto tr67
tr67:
//line machine.rl:17

	m.pb = m.p

	goto st23
	st23:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof23
		}
	st_case_23:
//line machine.go:7218
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr75
		case 13:
			goto st6
		case 32:
			goto tr74
		case 34:
			goto tr76
		case 44:
			goto tr77
		case 92:
			goto st113
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr74
		}
		goto st23
tr74:
//line machine.rl:58

	m.handler.SetMeasurement(m.text())

	goto st24
tr68:
//line machine.rl:58

	m.handler.SetMeasurement(m.text())

//line machine.rl:17

	m.pb = m.p

	goto st24
tr199:
//line machine.rl:66

	m.handler.AddTag(key, m.text())

	goto st24
	st24:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof24
		}
	st_case_24:
//line machine.go:7266
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr81
		case 13:
			goto st6
		case 32:
			goto st24
		case 34:
			goto tr82
		case 44:
			goto st6
		case 61:
			goto st6
		case 92:
			goto tr83
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto st24
		}
		goto tr79
tr79:
//line machine.rl:17

	m.pb = m.p

	goto st25
	st25:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof25
		}
	st_case_25:
//line machine.go:7300
		switch ( m.data)[( m.p)] {
		case 32:
			goto st6
		case 34:
			goto tr85
		case 44:
			goto st6
		case 61:
			goto tr86
		case 92:
			goto st60
		}
		switch {
		case ( m.data)[( m.p)] > 10:
			if 12 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto st6
			}
		case ( m.data)[( m.p)] >= 9:
			goto st6
		}
		goto st25
tr82:
//line machine.rl:17

	m.pb = m.p

//line machine.rl:86

	m.handler.AddString(key, m.text())

	goto st494
tr85:
//line machine.rl:86

	m.handler.AddString(key, m.text())

	goto st494
tr322:
//line machine.rl:86

	m.handler.AddString(key, m.text())

//line machine.rl:17

	m.pb = m.p

	goto st494
	st494:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof494
		}
	st_case_494:
//line machine.go:7353
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr850
		case 11:
			goto st495
		case 13:
			goto tr850
		case 32:
			goto st492
		case 44:
			goto st28
		case 61:
			goto tr11
		case 92:
			goto st27
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto st492
		}
		goto st3
	st495:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof495
		}
	st_case_495:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr850
		case 11:
			goto st495
		case 13:
			goto tr850
		case 32:
			goto st492
		case 44:
			goto tr88
		case 45:
			goto tr856
		case 61:
			goto tr11
		case 92:
			goto st27
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto tr857
			}
		case ( m.data)[( m.p)] >= 9:
			goto st492
		}
		goto st3
tr856:
//line machine.rl:17

	m.pb = m.p

	goto st26
	st26:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof26
		}
	st_case_26:
//line machine.go:7417
		switch ( m.data)[( m.p)] {
		case 32:
			goto tr88
		case 44:
			goto tr88
		case 61:
			goto tr11
		case 92:
			goto st27
		}
		switch {
		case ( m.data)[( m.p)] < 12:
			if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 10 {
				goto tr88
			}
		case ( m.data)[( m.p)] > 13:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st496
			}
		default:
			goto tr88
		}
		goto st3
tr857:
//line machine.rl:17

	m.pb = m.p

	goto st496
	st496:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof496
		}
	st_case_496:
//line machine.go:7452
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr860
		case 13:
			goto tr859
		case 32:
			goto tr858
		case 44:
			goto tr88
		case 61:
			goto tr11
		case 92:
			goto st27
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st499
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr858
		}
		goto st3
tr858:
//line machine.rl:90

	m.handler.SetTimestamp(m.text())

	goto st497
	st497:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof497
		}
	st_case_497:
//line machine.go:7489
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr850
		case 13:
			goto tr850
		case 32:
			goto st497
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto st497
		}
		goto st0
tr860:
//line machine.rl:90

	m.handler.SetTimestamp(m.text())

	goto st498
	st498:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof498
		}
	st_case_498:
//line machine.go:7513
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr850
		case 11:
			goto st498
		case 13:
			goto tr850
		case 32:
			goto st497
		case 44:
			goto tr2
		case 61:
			goto tr11
		case 92:
			goto st27
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto st497
		}
		goto st3
tr9:
//line machine.rl:17

	m.pb = m.p

	goto st27
	st27:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof27
		}
	st_case_27:
//line machine.go:7545
		switch {
		case ( m.data)[( m.p)] > 10:
			if 12 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto tr2
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr2
		}
		goto st3
	st499:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof499
		}
	st_case_499:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr860
		case 13:
			goto tr859
		case 32:
			goto tr858
		case 44:
			goto tr88
		case 61:
			goto tr11
		case 92:
			goto st27
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st500
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr858
		}
		goto st3
	st500:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof500
		}
	st_case_500:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr860
		case 13:
			goto tr859
		case 32:
			goto tr858
		case 44:
			goto tr88
		case 61:
			goto tr11
		case 92:
			goto st27
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st501
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr858
		}
		goto st3
	st501:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof501
		}
	st_case_501:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr860
		case 13:
			goto tr859
		case 32:
			goto tr858
		case 44:
			goto tr88
		case 61:
			goto tr11
		case 92:
			goto st27
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st502
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr858
		}
		goto st3
	st502:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof502
		}
	st_case_502:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr860
		case 13:
			goto tr859
		case 32:
			goto tr858
		case 44:
			goto tr88
		case 61:
			goto tr11
		case 92:
			goto st27
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st503
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr858
		}
		goto st3
	st503:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof503
		}
	st_case_503:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr860
		case 13:
			goto tr859
		case 32:
			goto tr858
		case 44:
			goto tr88
		case 61:
			goto tr11
		case 92:
			goto st27
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st504
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr858
		}
		goto st3
	st504:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof504
		}
	st_case_504:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr860
		case 13:
			goto tr859
		case 32:
			goto tr858
		case 44:
			goto tr88
		case 61:
			goto tr11
		case 92:
			goto st27
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st505
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr858
		}
		goto st3
	st505:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof505
		}
	st_case_505:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr860
		case 13:
			goto tr859
		case 32:
			goto tr858
		case 44:
			goto tr88
		case 61:
			goto tr11
		case 92:
			goto st27
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st506
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr858
		}
		goto st3
	st506:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof506
		}
	st_case_506:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr860
		case 13:
			goto tr859
		case 32:
			goto tr858
		case 44:
			goto tr88
		case 61:
			goto tr11
		case 92:
			goto st27
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st507
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr858
		}
		goto st3
	st507:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof507
		}
	st_case_507:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr860
		case 13:
			goto tr859
		case 32:
			goto tr858
		case 44:
			goto tr88
		case 61:
			goto tr11
		case 92:
			goto st27
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st508
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr858
		}
		goto st3
	st508:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof508
		}
	st_case_508:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr860
		case 13:
			goto tr859
		case 32:
			goto tr858
		case 44:
			goto tr88
		case 61:
			goto tr11
		case 92:
			goto st27
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st509
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr858
		}
		goto st3
	st509:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof509
		}
	st_case_509:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr860
		case 13:
			goto tr859
		case 32:
			goto tr858
		case 44:
			goto tr88
		case 61:
			goto tr11
		case 92:
			goto st27
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st510
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr858
		}
		goto st3
	st510:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof510
		}
	st_case_510:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr860
		case 13:
			goto tr859
		case 32:
			goto tr858
		case 44:
			goto tr88
		case 61:
			goto tr11
		case 92:
			goto st27
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st511
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr858
		}
		goto st3
	st511:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof511
		}
	st_case_511:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr860
		case 13:
			goto tr859
		case 32:
			goto tr858
		case 44:
			goto tr88
		case 61:
			goto tr11
		case 92:
			goto st27
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st512
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr858
		}
		goto st3
	st512:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof512
		}
	st_case_512:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr860
		case 13:
			goto tr859
		case 32:
			goto tr858
		case 44:
			goto tr88
		case 61:
			goto tr11
		case 92:
			goto st27
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st513
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr858
		}
		goto st3
	st513:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof513
		}
	st_case_513:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr860
		case 13:
			goto tr859
		case 32:
			goto tr858
		case 44:
			goto tr88
		case 61:
			goto tr11
		case 92:
			goto st27
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st514
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr858
		}
		goto st3
	st514:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof514
		}
	st_case_514:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr860
		case 13:
			goto tr859
		case 32:
			goto tr858
		case 44:
			goto tr88
		case 61:
			goto tr11
		case 92:
			goto st27
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st515
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr858
		}
		goto st3
	st515:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof515
		}
	st_case_515:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr860
		case 13:
			goto tr859
		case 32:
			goto tr858
		case 44:
			goto tr88
		case 61:
			goto tr11
		case 92:
			goto st27
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st516
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr858
		}
		goto st3
	st516:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof516
		}
	st_case_516:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr860
		case 13:
			goto tr859
		case 32:
			goto tr858
		case 44:
			goto tr88
		case 61:
			goto tr11
		case 92:
			goto st27
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr858
		}
		goto st3
tr1351:
//line machine.rl:78

	m.handler.AddFloat(key, m.text())

	goto st28
tr1355:
//line machine.rl:74

	m.handler.AddInt(key, m.text())

	goto st28
tr1357:
//line machine.rl:82

	m.handler.AddBool(key, m.text())

	goto st28
	st28:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof28
		}
	st_case_28:
//line machine.go:8113
		switch ( m.data)[( m.p)] {
		case 32:
			goto tr2
		case 44:
			goto tr2
		case 61:
			goto tr2
		case 92:
			goto tr9
		}
		switch {
		case ( m.data)[( m.p)] > 10:
			if 12 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto tr2
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr2
		}
		goto tr6
tr86:
//line machine.rl:70

	key = m.text()

	goto st29
	st29:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof29
		}
	st_case_29:
//line machine.go:8144
		switch ( m.data)[( m.p)] {
		case 34:
			goto tr90
		case 45:
			goto tr91
		case 46:
			goto tr92
		case 48:
			goto tr93
		case 70:
			goto tr95
		case 84:
			goto tr96
		case 92:
			goto st59
		case 102:
			goto tr97
		case 116:
			goto tr98
		}
		if 49 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
			goto tr94
		}
		goto st6
tr90:
//line machine.rl:86

	m.handler.AddString(key, m.text())

	goto st517
	st517:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof517
		}
	st_case_517:
//line machine.go:8180
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr882
		case 13:
			goto tr882
		case 32:
			goto tr881
		case 34:
			goto tr23
		case 44:
			goto tr883
		case 92:
			goto tr24
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr881
		}
		goto tr22
tr881:
//line machine.rl:17

	m.pb = m.p

	goto st518
tr1305:
//line machine.rl:78

	m.handler.AddFloat(key, m.text())

	goto st518
tr1309:
//line machine.rl:74

	m.handler.AddInt(key, m.text())

	goto st518
tr1311:
//line machine.rl:82

	m.handler.AddBool(key, m.text())

	goto st518
	st518:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof518
		}
	st_case_518:
//line machine.go:8228
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr885
		case 13:
			goto tr885
		case 32:
			goto st518
		case 34:
			goto tr26
		case 45:
			goto tr886
		case 92:
			goto st59
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto tr887
			}
		case ( m.data)[( m.p)] >= 9:
			goto st518
		}
		goto st6
tr885:
//line machine.rl:21

	yield = true
	{( m.p)++;  m.cs = 519; goto _out }

	goto st519
tr990:
//line machine.rl:90

	m.handler.SetTimestamp(m.text())

//line machine.rl:21

	yield = true
	{( m.p)++;  m.cs = 519; goto _out }

	goto st519
tr882:
//line machine.rl:17

	m.pb = m.p

//line machine.rl:21

	yield = true
	{( m.p)++;  m.cs = 519; goto _out }

	goto st519
tr1026:
//line machine.rl:78

	m.handler.AddFloat(key, m.text())

//line machine.rl:21

	yield = true
	{( m.p)++;  m.cs = 519; goto _out }

	goto st519
tr1162:
//line machine.rl:74

	m.handler.AddInt(key, m.text())

//line machine.rl:21

	yield = true
	{( m.p)++;  m.cs = 519; goto _out }

	goto st519
tr1166:
//line machine.rl:82

	m.handler.AddBool(key, m.text())

//line machine.rl:21

	yield = true
	{( m.p)++;  m.cs = 519; goto _out }

	goto st519
tr1392:
//line machine.rl:21

	yield = true
	{( m.p)++;  m.cs = 519; goto _out }

//line machine.rl:17

	m.pb = m.p

	goto st519
	st519:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof519
		}
	st_case_519:
//line machine.go:8330
		switch ( m.data)[( m.p)] {
		case 10:
			goto st519
		case 11:
			goto tr100
		case 13:
			goto st519
		case 32:
			goto st30
		case 34:
			goto tr70
		case 44:
			goto st6
		case 92:
			goto tr72
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto st30
		}
		goto tr67
	st30:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof30
		}
	st_case_30:
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr100
		case 13:
			goto st6
		case 32:
			goto st30
		case 34:
			goto tr70
		case 44:
			goto st6
		case 92:
			goto tr72
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto st30
		}
		goto tr67
tr100:
//line machine.rl:17

	m.pb = m.p

	goto st31
	st31:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof31
		}
	st_case_31:
//line machine.go:8387
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr102
		case 13:
			goto st6
		case 32:
			goto tr101
		case 34:
			goto tr70
		case 44:
			goto tr77
		case 92:
			goto tr72
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr101
		}
		goto tr67
tr101:
//line machine.rl:58

	m.handler.SetMeasurement(m.text())

	goto st32
	st32:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof32
		}
	st_case_32:
//line machine.go:8419
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr105
		case 13:
			goto st6
		case 32:
			goto st32
		case 34:
			goto tr106
		case 44:
			goto st6
		case 61:
			goto tr67
		case 92:
			goto tr107
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto st32
		}
		goto tr103
tr103:
//line machine.rl:17

	m.pb = m.p

	goto st33
	st33:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof33
		}
	st_case_33:
//line machine.go:8453
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr109
		case 13:
			goto st6
		case 32:
			goto tr74
		case 34:
			goto tr110
		case 44:
			goto tr77
		case 61:
			goto tr111
		case 92:
			goto st74
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr74
		}
		goto st33
tr109:
//line machine.rl:58

	m.handler.SetMeasurement(m.text())

	goto st34
tr113:
//line machine.rl:58

	m.handler.SetMeasurement(m.text())

//line machine.rl:17

	m.pb = m.p

	goto st34
	st34:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof34
		}
	st_case_34:
//line machine.go:8497
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr113
		case 13:
			goto st6
		case 32:
			goto tr74
		case 34:
			goto tr106
		case 44:
			goto tr77
		case 61:
			goto tr111
		case 92:
			goto tr107
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr74
		}
		goto tr103
tr106:
//line machine.rl:17

	m.pb = m.p

//line machine.rl:86

	m.handler.AddString(key, m.text())

	goto st520
tr110:
//line machine.rl:86

	m.handler.AddString(key, m.text())

	goto st520
	st520:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof520
		}
	st_case_520:
//line machine.go:8541
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr850
		case 11:
			goto tr890
		case 13:
			goto tr850
		case 32:
			goto tr889
		case 44:
			goto tr891
		case 61:
			goto tr36
		case 92:
			goto st36
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr889
		}
		goto st9
tr889:
//line machine.rl:58

	m.handler.SetMeasurement(m.text())

	goto st521
tr954:
//line machine.rl:66

	m.handler.AddTag(key, m.text())

	goto st521
tr1317:
//line machine.rl:58

	m.handler.SetMeasurement(m.text())

//line machine.rl:78

	m.handler.AddFloat(key, m.text())

	goto st521
tr1269:
//line machine.rl:66

	m.handler.AddTag(key, m.text())

//line machine.rl:78

	m.handler.AddFloat(key, m.text())

	goto st521
tr1275:
//line machine.rl:66

	m.handler.AddTag(key, m.text())

//line machine.rl:74

	m.handler.AddInt(key, m.text())

	goto st521
tr1279:
//line machine.rl:66

	m.handler.AddTag(key, m.text())

//line machine.rl:82

	m.handler.AddBool(key, m.text())

	goto st521
tr1322:
//line machine.rl:58

	m.handler.SetMeasurement(m.text())

//line machine.rl:74

	m.handler.AddInt(key, m.text())

	goto st521
tr1325:
//line machine.rl:58

	m.handler.SetMeasurement(m.text())

//line machine.rl:82

	m.handler.AddBool(key, m.text())

	goto st521
	st521:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof521
		}
	st_case_521:
//line machine.go:8639
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr850
		case 11:
			goto tr893
		case 13:
			goto tr850
		case 32:
			goto st521
		case 44:
			goto tr88
		case 45:
			goto tr856
		case 61:
			goto tr88
		case 92:
			goto tr9
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto tr857
			}
		case ( m.data)[( m.p)] >= 9:
			goto st521
		}
		goto tr6
tr893:
//line machine.rl:17

	m.pb = m.p

	goto st522
	st522:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof522
		}
	st_case_522:
//line machine.go:8678
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr850
		case 11:
			goto tr893
		case 13:
			goto tr850
		case 32:
			goto st521
		case 44:
			goto tr88
		case 45:
			goto tr856
		case 61:
			goto tr11
		case 92:
			goto tr9
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto tr857
			}
		case ( m.data)[( m.p)] >= 9:
			goto st521
		}
		goto tr6
tr890:
//line machine.rl:58

	m.handler.SetMeasurement(m.text())

	goto st523
tr894:
//line machine.rl:58

	m.handler.SetMeasurement(m.text())

//line machine.rl:17

	m.pb = m.p

	goto st523
	st523:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof523
		}
	st_case_523:
//line machine.go:8727
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr850
		case 11:
			goto tr894
		case 13:
			goto tr850
		case 32:
			goto tr889
		case 44:
			goto tr4
		case 45:
			goto tr895
		case 61:
			goto tr36
		case 92:
			goto tr33
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto tr896
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr889
		}
		goto tr31
tr895:
//line machine.rl:17

	m.pb = m.p

	goto st35
	st35:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof35
		}
	st_case_35:
//line machine.go:8766
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr88
		case 11:
			goto tr35
		case 13:
			goto tr88
		case 32:
			goto tr1
		case 44:
			goto tr4
		case 61:
			goto tr36
		case 92:
			goto st36
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st524
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1
		}
		goto st9
tr896:
//line machine.rl:17

	m.pb = m.p

	goto st524
	st524:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof524
		}
	st_case_524:
//line machine.go:8803
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr898
		case 13:
			goto tr859
		case 32:
			goto tr897
		case 44:
			goto tr4
		case 61:
			goto tr36
		case 92:
			goto st36
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st528
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr897
		}
		goto st9
tr902:
//line machine.rl:58

	m.handler.SetMeasurement(m.text())

	goto st525
tr963:
//line machine.rl:66

	m.handler.AddTag(key, m.text())

	goto st525
tr897:
//line machine.rl:58

	m.handler.SetMeasurement(m.text())

//line machine.rl:90

	m.handler.SetTimestamp(m.text())

	goto st525
tr960:
//line machine.rl:66

	m.handler.AddTag(key, m.text())

//line machine.rl:90

	m.handler.SetTimestamp(m.text())

	goto st525
tr1449:
//line machine.rl:90

	m.handler.SetTimestamp(m.text())

//line machine.rl:58

	m.handler.SetMeasurement(m.text())

	goto st525
tr1476:
//line machine.rl:90

	m.handler.SetTimestamp(m.text())

//line machine.rl:66

	m.handler.AddTag(key, m.text())

	goto st525
	st525:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof525
		}
	st_case_525:
//line machine.go:8886
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr850
		case 11:
			goto tr901
		case 13:
			goto tr850
		case 32:
			goto st525
		case 44:
			goto tr2
		case 61:
			goto tr2
		case 92:
			goto tr9
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto st525
		}
		goto tr6
tr901:
//line machine.rl:17

	m.pb = m.p

	goto st526
	st526:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof526
		}
	st_case_526:
//line machine.go:8918
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr850
		case 11:
			goto tr901
		case 13:
			goto tr850
		case 32:
			goto st525
		case 44:
			goto tr2
		case 61:
			goto tr11
		case 92:
			goto tr9
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto st525
		}
		goto tr6
tr903:
//line machine.rl:58

	m.handler.SetMeasurement(m.text())

//line machine.rl:17

	m.pb = m.p

	goto st527
tr898:
//line machine.rl:58

	m.handler.SetMeasurement(m.text())

//line machine.rl:90

	m.handler.SetTimestamp(m.text())

	goto st527
	st527:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof527
		}
	st_case_527:
//line machine.go:8964
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr850
		case 11:
			goto tr903
		case 13:
			goto tr850
		case 32:
			goto tr902
		case 44:
			goto tr4
		case 61:
			goto tr36
		case 92:
			goto tr33
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr902
		}
		goto tr31
tr33:
//line machine.rl:17

	m.pb = m.p

	goto st36
	st36:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof36
		}
	st_case_36:
//line machine.go:8996
		if ( m.data)[( m.p)] == 92 {
			goto st3
		}
		switch {
		case ( m.data)[( m.p)] > 10:
			if 12 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto tr2
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr2
		}
		goto st9
	st528:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof528
		}
	st_case_528:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr898
		case 13:
			goto tr859
		case 32:
			goto tr897
		case 44:
			goto tr4
		case 61:
			goto tr36
		case 92:
			goto st36
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st529
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr897
		}
		goto st9
	st529:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof529
		}
	st_case_529:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr898
		case 13:
			goto tr859
		case 32:
			goto tr897
		case 44:
			goto tr4
		case 61:
			goto tr36
		case 92:
			goto st36
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st530
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr897
		}
		goto st9
	st530:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof530
		}
	st_case_530:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr898
		case 13:
			goto tr859
		case 32:
			goto tr897
		case 44:
			goto tr4
		case 61:
			goto tr36
		case 92:
			goto st36
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st531
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr897
		}
		goto st9
	st531:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof531
		}
	st_case_531:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr898
		case 13:
			goto tr859
		case 32:
			goto tr897
		case 44:
			goto tr4
		case 61:
			goto tr36
		case 92:
			goto st36
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st532
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr897
		}
		goto st9
	st532:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof532
		}
	st_case_532:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr898
		case 13:
			goto tr859
		case 32:
			goto tr897
		case 44:
			goto tr4
		case 61:
			goto tr36
		case 92:
			goto st36
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st533
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr897
		}
		goto st9
	st533:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof533
		}
	st_case_533:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr898
		case 13:
			goto tr859
		case 32:
			goto tr897
		case 44:
			goto tr4
		case 61:
			goto tr36
		case 92:
			goto st36
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st534
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr897
		}
		goto st9
	st534:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof534
		}
	st_case_534:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr898
		case 13:
			goto tr859
		case 32:
			goto tr897
		case 44:
			goto tr4
		case 61:
			goto tr36
		case 92:
			goto st36
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st535
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr897
		}
		goto st9
	st535:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof535
		}
	st_case_535:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr898
		case 13:
			goto tr859
		case 32:
			goto tr897
		case 44:
			goto tr4
		case 61:
			goto tr36
		case 92:
			goto st36
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st536
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr897
		}
		goto st9
	st536:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof536
		}
	st_case_536:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr898
		case 13:
			goto tr859
		case 32:
			goto tr897
		case 44:
			goto tr4
		case 61:
			goto tr36
		case 92:
			goto st36
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st537
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr897
		}
		goto st9
	st537:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof537
		}
	st_case_537:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr898
		case 13:
			goto tr859
		case 32:
			goto tr897
		case 44:
			goto tr4
		case 61:
			goto tr36
		case 92:
			goto st36
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st538
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr897
		}
		goto st9
	st538:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof538
		}
	st_case_538:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr898
		case 13:
			goto tr859
		case 32:
			goto tr897
		case 44:
			goto tr4
		case 61:
			goto tr36
		case 92:
			goto st36
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st539
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr897
		}
		goto st9
	st539:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof539
		}
	st_case_539:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr898
		case 13:
			goto tr859
		case 32:
			goto tr897
		case 44:
			goto tr4
		case 61:
			goto tr36
		case 92:
			goto st36
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st540
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr897
		}
		goto st9
	st540:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof540
		}
	st_case_540:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr898
		case 13:
			goto tr859
		case 32:
			goto tr897
		case 44:
			goto tr4
		case 61:
			goto tr36
		case 92:
			goto st36
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st541
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr897
		}
		goto st9
	st541:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof541
		}
	st_case_541:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr898
		case 13:
			goto tr859
		case 32:
			goto tr897
		case 44:
			goto tr4
		case 61:
			goto tr36
		case 92:
			goto st36
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st542
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr897
		}
		goto st9
	st542:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof542
		}
	st_case_542:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr898
		case 13:
			goto tr859
		case 32:
			goto tr897
		case 44:
			goto tr4
		case 61:
			goto tr36
		case 92:
			goto st36
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st543
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr897
		}
		goto st9
	st543:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof543
		}
	st_case_543:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr898
		case 13:
			goto tr859
		case 32:
			goto tr897
		case 44:
			goto tr4
		case 61:
			goto tr36
		case 92:
			goto st36
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st544
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr897
		}
		goto st9
	st544:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof544
		}
	st_case_544:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr898
		case 13:
			goto tr859
		case 32:
			goto tr897
		case 44:
			goto tr4
		case 61:
			goto tr36
		case 92:
			goto st36
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st545
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr897
		}
		goto st9
	st545:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof545
		}
	st_case_545:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr898
		case 13:
			goto tr859
		case 32:
			goto tr897
		case 44:
			goto tr4
		case 61:
			goto tr36
		case 92:
			goto st36
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr897
		}
		goto st9
tr891:
//line machine.rl:58

	m.handler.SetMeasurement(m.text())

	goto st37
tr956:
//line machine.rl:66

	m.handler.AddTag(key, m.text())

	goto st37
tr1319:
//line machine.rl:58

	m.handler.SetMeasurement(m.text())

//line machine.rl:78

	m.handler.AddFloat(key, m.text())

	goto st37
tr1272:
//line machine.rl:66

	m.handler.AddTag(key, m.text())

//line machine.rl:78

	m.handler.AddFloat(key, m.text())

	goto st37
tr1278:
//line machine.rl:66

	m.handler.AddTag(key, m.text())

//line machine.rl:74

	m.handler.AddInt(key, m.text())

	goto st37
tr1282:
//line machine.rl:66

	m.handler.AddTag(key, m.text())

//line machine.rl:82

	m.handler.AddBool(key, m.text())

	goto st37
tr1324:
//line machine.rl:58

	m.handler.SetMeasurement(m.text())

//line machine.rl:74

	m.handler.AddInt(key, m.text())

	goto st37
tr1327:
//line machine.rl:58

	m.handler.SetMeasurement(m.text())

//line machine.rl:82

	m.handler.AddBool(key, m.text())

	goto st37
	st37:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof37
		}
	st_case_37:
//line machine.go:9621
		switch ( m.data)[( m.p)] {
		case 32:
			goto tr48
		case 44:
			goto tr48
		case 61:
			goto tr48
		case 92:
			goto tr116
		}
		switch {
		case ( m.data)[( m.p)] > 10:
			if 12 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto tr48
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr48
		}
		goto tr115
tr115:
//line machine.rl:17

	m.pb = m.p

	goto st38
	st38:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof38
		}
	st_case_38:
//line machine.go:9652
		switch ( m.data)[( m.p)] {
		case 32:
			goto tr48
		case 44:
			goto tr48
		case 61:
			goto tr118
		case 92:
			goto st80
		}
		switch {
		case ( m.data)[( m.p)] > 10:
			if 12 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto tr48
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr48
		}
		goto st38
tr118:
//line machine.rl:62

	key = m.text()

//line machine.rl:70

	key = m.text()

	goto st39
	st39:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof39
		}
	st_case_39:
//line machine.go:9687
		switch ( m.data)[( m.p)] {
		case 32:
			goto tr48
		case 34:
			goto tr120
		case 44:
			goto tr48
		case 45:
			goto tr121
		case 46:
			goto tr122
		case 48:
			goto tr123
		case 61:
			goto tr48
		case 70:
			goto tr125
		case 84:
			goto tr126
		case 92:
			goto tr45
		case 102:
			goto tr127
		case 116:
			goto tr128
		}
		switch {
		case ( m.data)[( m.p)] < 12:
			if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 10 {
				goto tr48
			}
		case ( m.data)[( m.p)] > 13:
			if 49 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto tr124
			}
		default:
			goto tr48
		}
		goto tr44
tr120:
//line machine.rl:17

	m.pb = m.p

	goto st40
	st40:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof40
		}
	st_case_40:
//line machine.go:9738
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr22
		case 11:
			goto tr131
		case 13:
			goto tr22
		case 32:
			goto tr130
		case 34:
			goto tr132
		case 44:
			goto tr133
		case 61:
			goto tr22
		case 92:
			goto tr134
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr130
		}
		goto tr129
tr129:
//line machine.rl:17

	m.pb = m.p

	goto st41
	st41:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof41
		}
	st_case_41:
//line machine.go:9772
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr137
		case 13:
			goto st6
		case 32:
			goto tr136
		case 34:
			goto tr138
		case 44:
			goto tr139
		case 61:
			goto st6
		case 92:
			goto st55
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr136
		}
		goto st41
tr159:
//line machine.rl:58

	m.handler.SetMeasurement(m.text())

	goto st42
tr136:
//line machine.rl:66

	m.handler.AddTag(key, m.text())

	goto st42
tr130:
//line machine.rl:66

	m.handler.AddTag(key, m.text())

//line machine.rl:17

	m.pb = m.p

	goto st42
	st42:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof42
		}
	st_case_42:
//line machine.go:9822
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr143
		case 13:
			goto st6
		case 32:
			goto st42
		case 34:
			goto tr82
		case 44:
			goto st6
		case 61:
			goto st6
		case 92:
			goto tr144
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto st42
		}
		goto tr141
tr141:
//line machine.rl:17

	m.pb = m.p

	goto st43
	st43:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof43
		}
	st_case_43:
//line machine.go:9856
		switch ( m.data)[( m.p)] {
		case 32:
			goto st6
		case 34:
			goto tr85
		case 44:
			goto st6
		case 61:
			goto tr146
		case 92:
			goto st83
		}
		switch {
		case ( m.data)[( m.p)] > 10:
			if 12 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto st6
			}
		case ( m.data)[( m.p)] >= 9:
			goto st6
		}
		goto st43
tr146:
//line machine.rl:70

	key = m.text()

	goto st44
	st44:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof44
		}
	st_case_44:
//line machine.go:9889
		switch ( m.data)[( m.p)] {
		case 34:
			goto tr90
		case 45:
			goto tr148
		case 46:
			goto tr149
		case 48:
			goto tr150
		case 70:
			goto tr152
		case 84:
			goto tr153
		case 92:
			goto st59
		case 102:
			goto tr154
		case 116:
			goto tr155
		}
		if 49 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
			goto tr151
		}
		goto st6
tr148:
//line machine.rl:17

	m.pb = m.p

	goto st45
	st45:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof45
		}
	st_case_45:
//line machine.go:9925
		switch ( m.data)[( m.p)] {
		case 34:
			goto tr26
		case 48:
			goto st546
		case 92:
			goto st59
		}
		if 49 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
			goto st816
		}
		goto st6
tr150:
//line machine.rl:17

	m.pb = m.p

	goto st546
	st546:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof546
		}
	st_case_546:
//line machine.go:9949
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr922
		case 13:
			goto tr922
		case 32:
			goto tr921
		case 34:
			goto tr26
		case 44:
			goto tr923
		case 46:
			goto st815
		case 92:
			goto st59
		case 105:
			goto st817
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr921
		}
		goto st6
tr1259:
//line machine.rl:17

	m.pb = m.p

	goto st547
tr921:
//line machine.rl:78

	m.handler.AddFloat(key, m.text())

	goto st547
tr1261:
//line machine.rl:74

	m.handler.AddInt(key, m.text())

	goto st547
tr1263:
//line machine.rl:82

	m.handler.AddBool(key, m.text())

	goto st547
	st547:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof547
		}
	st_case_547:
//line machine.go:10001
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr927
		case 13:
			goto tr927
		case 32:
			goto st547
		case 34:
			goto tr26
		case 45:
			goto tr928
		case 92:
			goto st59
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto tr929
			}
		case ( m.data)[( m.p)] >= 9:
			goto st547
		}
		goto st6
tr927:
//line machine.rl:21

	yield = true
	{( m.p)++;  m.cs = 548; goto _out }

	goto st548
tr1063:
//line machine.rl:90

	m.handler.SetTimestamp(m.text())

//line machine.rl:21

	yield = true
	{( m.p)++;  m.cs = 548; goto _out }

	goto st548
tr1055:
//line machine.rl:17

	m.pb = m.p

//line machine.rl:21

	yield = true
	{( m.p)++;  m.cs = 548; goto _out }

	goto st548
tr922:
//line machine.rl:78

	m.handler.AddFloat(key, m.text())

//line machine.rl:21

	yield = true
	{( m.p)++;  m.cs = 548; goto _out }

	goto st548
tr1119:
//line machine.rl:74

	m.handler.AddInt(key, m.text())

//line machine.rl:21

	yield = true
	{( m.p)++;  m.cs = 548; goto _out }

	goto st548
tr1123:
//line machine.rl:82

	m.handler.AddBool(key, m.text())

//line machine.rl:21

	yield = true
	{( m.p)++;  m.cs = 548; goto _out }

	goto st548
tr1398:
//line machine.rl:21

	yield = true
	{( m.p)++;  m.cs = 548; goto _out }

//line machine.rl:17

	m.pb = m.p

	goto st548
	st548:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof548
		}
	st_case_548:
//line machine.go:10103
		switch ( m.data)[( m.p)] {
		case 10:
			goto st548
		case 11:
			goto tr291
		case 13:
			goto st548
		case 32:
			goto st134
		case 34:
			goto tr70
		case 44:
			goto st6
		case 92:
			goto tr292
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto st134
		}
		goto tr289
tr289:
//line machine.rl:17

	m.pb = m.p

	goto st46
	st46:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof46
		}
	st_case_46:
//line machine.go:10135
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr160
		case 13:
			goto st6
		case 32:
			goto tr159
		case 34:
			goto tr76
		case 44:
			goto tr161
		case 92:
			goto st125
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr159
		}
		goto st46
tr160:
//line machine.rl:58

	m.handler.SetMeasurement(m.text())

	goto st47
	st47:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof47
		}
	st_case_47:
//line machine.go:10167
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr164
		case 13:
			goto st6
		case 32:
			goto tr159
		case 34:
			goto tr106
		case 44:
			goto tr161
		case 61:
			goto st46
		case 92:
			goto tr165
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr159
		}
		goto tr163
tr163:
//line machine.rl:17

	m.pb = m.p

	goto st48
	st48:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof48
		}
	st_case_48:
//line machine.go:10201
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr167
		case 13:
			goto st6
		case 32:
			goto tr159
		case 34:
			goto tr110
		case 44:
			goto tr161
		case 61:
			goto tr168
		case 92:
			goto st124
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr159
		}
		goto st48
tr167:
//line machine.rl:58

	m.handler.SetMeasurement(m.text())

	goto st49
tr164:
//line machine.rl:58

	m.handler.SetMeasurement(m.text())

//line machine.rl:17

	m.pb = m.p

	goto st49
	st49:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof49
		}
	st_case_49:
//line machine.go:10245
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr164
		case 13:
			goto st6
		case 32:
			goto tr159
		case 34:
			goto tr106
		case 44:
			goto tr161
		case 61:
			goto tr168
		case 92:
			goto tr165
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr159
		}
		goto tr163
tr161:
//line machine.rl:58

	m.handler.SetMeasurement(m.text())

	goto st50
tr139:
//line machine.rl:66

	m.handler.AddTag(key, m.text())

	goto st50
tr133:
//line machine.rl:66

	m.handler.AddTag(key, m.text())

//line machine.rl:17

	m.pb = m.p

	goto st50
	st50:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof50
		}
	st_case_50:
//line machine.go:10295
		switch ( m.data)[( m.p)] {
		case 32:
			goto st6
		case 34:
			goto tr171
		case 44:
			goto st6
		case 61:
			goto st6
		case 92:
			goto tr172
		}
		switch {
		case ( m.data)[( m.p)] > 10:
			if 12 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto st6
			}
		case ( m.data)[( m.p)] >= 9:
			goto st6
		}
		goto tr170
tr170:
//line machine.rl:17

	m.pb = m.p

	goto st51
	st51:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof51
		}
	st_case_51:
//line machine.go:10328
		switch ( m.data)[( m.p)] {
		case 32:
			goto st6
		case 34:
			goto tr174
		case 44:
			goto st6
		case 61:
			goto tr175
		case 92:
			goto st56
		}
		switch {
		case ( m.data)[( m.p)] > 10:
			if 12 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto st6
			}
		case ( m.data)[( m.p)] >= 9:
			goto st6
		}
		goto st51
tr171:
//line machine.rl:17

	m.pb = m.p

//line machine.rl:86

	m.handler.AddString(key, m.text())

	goto st549
tr174:
//line machine.rl:86

	m.handler.AddString(key, m.text())

	goto st549
	st549:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof549
		}
	st_case_549:
//line machine.go:10371
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr850
		case 11:
			goto st550
		case 13:
			goto tr850
		case 32:
			goto st492
		case 44:
			goto st28
		case 61:
			goto tr42
		case 92:
			goto st20
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto st492
		}
		goto st12
	st550:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof550
		}
	st_case_550:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr850
		case 11:
			goto st550
		case 13:
			goto tr850
		case 32:
			goto st492
		case 44:
			goto tr177
		case 45:
			goto tr932
		case 61:
			goto tr42
		case 92:
			goto st20
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto tr933
			}
		case ( m.data)[( m.p)] >= 9:
			goto st492
		}
		goto st12
tr932:
//line machine.rl:17

	m.pb = m.p

	goto st52
	st52:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof52
		}
	st_case_52:
//line machine.go:10435
		switch ( m.data)[( m.p)] {
		case 32:
			goto tr177
		case 44:
			goto tr177
		case 61:
			goto tr42
		case 92:
			goto st20
		}
		switch {
		case ( m.data)[( m.p)] < 12:
			if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 10 {
				goto tr177
			}
		case ( m.data)[( m.p)] > 13:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st551
			}
		default:
			goto tr177
		}
		goto st12
tr933:
//line machine.rl:17

	m.pb = m.p

	goto st551
	st551:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof551
		}
	st_case_551:
//line machine.go:10470
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr934
		case 13:
			goto tr859
		case 32:
			goto tr858
		case 44:
			goto tr177
		case 61:
			goto tr42
		case 92:
			goto st20
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st553
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr858
		}
		goto st12
tr934:
//line machine.rl:90

	m.handler.SetTimestamp(m.text())

	goto st552
	st552:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof552
		}
	st_case_552:
//line machine.go:10507
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr850
		case 11:
			goto st552
		case 13:
			goto tr850
		case 32:
			goto st497
		case 44:
			goto tr39
		case 61:
			goto tr42
		case 92:
			goto st20
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto st497
		}
		goto st12
	st553:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof553
		}
	st_case_553:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr934
		case 13:
			goto tr859
		case 32:
			goto tr858
		case 44:
			goto tr177
		case 61:
			goto tr42
		case 92:
			goto st20
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st554
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr858
		}
		goto st12
	st554:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof554
		}
	st_case_554:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr934
		case 13:
			goto tr859
		case 32:
			goto tr858
		case 44:
			goto tr177
		case 61:
			goto tr42
		case 92:
			goto st20
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st555
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr858
		}
		goto st12
	st555:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof555
		}
	st_case_555:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr934
		case 13:
			goto tr859
		case 32:
			goto tr858
		case 44:
			goto tr177
		case 61:
			goto tr42
		case 92:
			goto st20
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st556
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr858
		}
		goto st12
	st556:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof556
		}
	st_case_556:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr934
		case 13:
			goto tr859
		case 32:
			goto tr858
		case 44:
			goto tr177
		case 61:
			goto tr42
		case 92:
			goto st20
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st557
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr858
		}
		goto st12
	st557:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof557
		}
	st_case_557:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr934
		case 13:
			goto tr859
		case 32:
			goto tr858
		case 44:
			goto tr177
		case 61:
			goto tr42
		case 92:
			goto st20
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st558
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr858
		}
		goto st12
	st558:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof558
		}
	st_case_558:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr934
		case 13:
			goto tr859
		case 32:
			goto tr858
		case 44:
			goto tr177
		case 61:
			goto tr42
		case 92:
			goto st20
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st559
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr858
		}
		goto st12
	st559:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof559
		}
	st_case_559:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr934
		case 13:
			goto tr859
		case 32:
			goto tr858
		case 44:
			goto tr177
		case 61:
			goto tr42
		case 92:
			goto st20
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st560
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr858
		}
		goto st12
	st560:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof560
		}
	st_case_560:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr934
		case 13:
			goto tr859
		case 32:
			goto tr858
		case 44:
			goto tr177
		case 61:
			goto tr42
		case 92:
			goto st20
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st561
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr858
		}
		goto st12
	st561:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof561
		}
	st_case_561:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr934
		case 13:
			goto tr859
		case 32:
			goto tr858
		case 44:
			goto tr177
		case 61:
			goto tr42
		case 92:
			goto st20
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st562
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr858
		}
		goto st12
	st562:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof562
		}
	st_case_562:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr934
		case 13:
			goto tr859
		case 32:
			goto tr858
		case 44:
			goto tr177
		case 61:
			goto tr42
		case 92:
			goto st20
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st563
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr858
		}
		goto st12
	st563:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof563
		}
	st_case_563:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr934
		case 13:
			goto tr859
		case 32:
			goto tr858
		case 44:
			goto tr177
		case 61:
			goto tr42
		case 92:
			goto st20
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st564
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr858
		}
		goto st12
	st564:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof564
		}
	st_case_564:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr934
		case 13:
			goto tr859
		case 32:
			goto tr858
		case 44:
			goto tr177
		case 61:
			goto tr42
		case 92:
			goto st20
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st565
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr858
		}
		goto st12
	st565:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof565
		}
	st_case_565:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr934
		case 13:
			goto tr859
		case 32:
			goto tr858
		case 44:
			goto tr177
		case 61:
			goto tr42
		case 92:
			goto st20
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st566
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr858
		}
		goto st12
	st566:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof566
		}
	st_case_566:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr934
		case 13:
			goto tr859
		case 32:
			goto tr858
		case 44:
			goto tr177
		case 61:
			goto tr42
		case 92:
			goto st20
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st567
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr858
		}
		goto st12
	st567:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof567
		}
	st_case_567:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr934
		case 13:
			goto tr859
		case 32:
			goto tr858
		case 44:
			goto tr177
		case 61:
			goto tr42
		case 92:
			goto st20
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st568
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr858
		}
		goto st12
	st568:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof568
		}
	st_case_568:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr934
		case 13:
			goto tr859
		case 32:
			goto tr858
		case 44:
			goto tr177
		case 61:
			goto tr42
		case 92:
			goto st20
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st569
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr858
		}
		goto st12
	st569:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof569
		}
	st_case_569:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr934
		case 13:
			goto tr859
		case 32:
			goto tr858
		case 44:
			goto tr177
		case 61:
			goto tr42
		case 92:
			goto st20
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st570
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr858
		}
		goto st12
	st570:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof570
		}
	st_case_570:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr934
		case 13:
			goto tr859
		case 32:
			goto tr858
		case 44:
			goto tr177
		case 61:
			goto tr42
		case 92:
			goto st20
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr858
		}
		goto st12
tr175:
//line machine.rl:62

	key = m.text()

	goto st53
	st53:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof53
		}
	st_case_53:
//line machine.go:11074
		switch ( m.data)[( m.p)] {
		case 32:
			goto st6
		case 34:
			goto tr132
		case 44:
			goto st6
		case 61:
			goto st6
		case 92:
			goto tr134
		}
		switch {
		case ( m.data)[( m.p)] > 10:
			if 12 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto st6
			}
		case ( m.data)[( m.p)] >= 9:
			goto st6
		}
		goto tr129
tr132:
//line machine.rl:17

	m.pb = m.p

//line machine.rl:86

	m.handler.AddString(key, m.text())

	goto st571
tr138:
//line machine.rl:86

	m.handler.AddString(key, m.text())

	goto st571
	st571:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof571
		}
	st_case_571:
//line machine.go:11117
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr850
		case 11:
			goto tr955
		case 13:
			goto tr850
		case 32:
			goto tr954
		case 44:
			goto tr956
		case 61:
			goto tr48
		case 92:
			goto st19
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr954
		}
		goto st14
tr955:
//line machine.rl:66

	m.handler.AddTag(key, m.text())

	goto st572
tr1271:
//line machine.rl:66

	m.handler.AddTag(key, m.text())

//line machine.rl:78

	m.handler.AddFloat(key, m.text())

	goto st572
tr1277:
//line machine.rl:66

	m.handler.AddTag(key, m.text())

//line machine.rl:74

	m.handler.AddInt(key, m.text())

	goto st572
tr1281:
//line machine.rl:66

	m.handler.AddTag(key, m.text())

//line machine.rl:82

	m.handler.AddBool(key, m.text())

	goto st572
	st572:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof572
		}
	st_case_572:
//line machine.go:11179
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr850
		case 11:
			goto tr957
		case 13:
			goto tr850
		case 32:
			goto tr954
		case 44:
			goto tr50
		case 45:
			goto tr958
		case 61:
			goto tr179
		case 92:
			goto tr54
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto tr959
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr954
		}
		goto tr52
tr1016:
//line machine.rl:66

	m.handler.AddTag(key, m.text())

	goto st573
tr957:
//line machine.rl:66

	m.handler.AddTag(key, m.text())

//line machine.rl:17

	m.pb = m.p

	goto st573
	st573:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof573
		}
	st_case_573:
//line machine.go:11228
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr850
		case 11:
			goto tr957
		case 13:
			goto tr850
		case 32:
			goto tr954
		case 44:
			goto tr50
		case 45:
			goto tr958
		case 61:
			goto tr11
		case 92:
			goto tr54
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto tr959
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr954
		}
		goto tr52
tr958:
//line machine.rl:17

	m.pb = m.p

	goto st54
	st54:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof54
		}
	st_case_54:
//line machine.go:11267
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr179
		case 11:
			goto tr56
		case 13:
			goto tr179
		case 32:
			goto tr47
		case 44:
			goto tr50
		case 61:
			goto tr11
		case 92:
			goto st18
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st574
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr47
		}
		goto st16
tr959:
//line machine.rl:17

	m.pb = m.p

	goto st574
	st574:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof574
		}
	st_case_574:
//line machine.go:11304
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr961
		case 13:
			goto tr859
		case 32:
			goto tr960
		case 44:
			goto tr50
		case 61:
			goto tr11
		case 92:
			goto st18
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st576
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr960
		}
		goto st16
tr964:
//line machine.rl:66

	m.handler.AddTag(key, m.text())

//line machine.rl:17

	m.pb = m.p

	goto st575
tr961:
//line machine.rl:66

	m.handler.AddTag(key, m.text())

//line machine.rl:90

	m.handler.SetTimestamp(m.text())

	goto st575
tr1477:
//line machine.rl:90

	m.handler.SetTimestamp(m.text())

//line machine.rl:66

	m.handler.AddTag(key, m.text())

	goto st575
	st575:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof575
		}
	st_case_575:
//line machine.go:11365
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr850
		case 11:
			goto tr964
		case 13:
			goto tr850
		case 32:
			goto tr963
		case 44:
			goto tr50
		case 61:
			goto tr11
		case 92:
			goto tr54
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr963
		}
		goto tr52
	st576:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof576
		}
	st_case_576:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr961
		case 13:
			goto tr859
		case 32:
			goto tr960
		case 44:
			goto tr50
		case 61:
			goto tr11
		case 92:
			goto st18
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st577
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr960
		}
		goto st16
	st577:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof577
		}
	st_case_577:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr961
		case 13:
			goto tr859
		case 32:
			goto tr960
		case 44:
			goto tr50
		case 61:
			goto tr11
		case 92:
			goto st18
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st578
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr960
		}
		goto st16
	st578:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof578
		}
	st_case_578:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr961
		case 13:
			goto tr859
		case 32:
			goto tr960
		case 44:
			goto tr50
		case 61:
			goto tr11
		case 92:
			goto st18
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st579
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr960
		}
		goto st16
	st579:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof579
		}
	st_case_579:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr961
		case 13:
			goto tr859
		case 32:
			goto tr960
		case 44:
			goto tr50
		case 61:
			goto tr11
		case 92:
			goto st18
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st580
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr960
		}
		goto st16
	st580:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof580
		}
	st_case_580:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr961
		case 13:
			goto tr859
		case 32:
			goto tr960
		case 44:
			goto tr50
		case 61:
			goto tr11
		case 92:
			goto st18
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st581
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr960
		}
		goto st16
	st581:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof581
		}
	st_case_581:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr961
		case 13:
			goto tr859
		case 32:
			goto tr960
		case 44:
			goto tr50
		case 61:
			goto tr11
		case 92:
			goto st18
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st582
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr960
		}
		goto st16
	st582:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof582
		}
	st_case_582:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr961
		case 13:
			goto tr859
		case 32:
			goto tr960
		case 44:
			goto tr50
		case 61:
			goto tr11
		case 92:
			goto st18
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st583
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr960
		}
		goto st16
	st583:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof583
		}
	st_case_583:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr961
		case 13:
			goto tr859
		case 32:
			goto tr960
		case 44:
			goto tr50
		case 61:
			goto tr11
		case 92:
			goto st18
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st584
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr960
		}
		goto st16
	st584:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof584
		}
	st_case_584:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr961
		case 13:
			goto tr859
		case 32:
			goto tr960
		case 44:
			goto tr50
		case 61:
			goto tr11
		case 92:
			goto st18
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st585
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr960
		}
		goto st16
	st585:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof585
		}
	st_case_585:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr961
		case 13:
			goto tr859
		case 32:
			goto tr960
		case 44:
			goto tr50
		case 61:
			goto tr11
		case 92:
			goto st18
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st586
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr960
		}
		goto st16
	st586:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof586
		}
	st_case_586:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr961
		case 13:
			goto tr859
		case 32:
			goto tr960
		case 44:
			goto tr50
		case 61:
			goto tr11
		case 92:
			goto st18
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st587
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr960
		}
		goto st16
	st587:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof587
		}
	st_case_587:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr961
		case 13:
			goto tr859
		case 32:
			goto tr960
		case 44:
			goto tr50
		case 61:
			goto tr11
		case 92:
			goto st18
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st588
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr960
		}
		goto st16
	st588:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof588
		}
	st_case_588:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr961
		case 13:
			goto tr859
		case 32:
			goto tr960
		case 44:
			goto tr50
		case 61:
			goto tr11
		case 92:
			goto st18
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st589
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr960
		}
		goto st16
	st589:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof589
		}
	st_case_589:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr961
		case 13:
			goto tr859
		case 32:
			goto tr960
		case 44:
			goto tr50
		case 61:
			goto tr11
		case 92:
			goto st18
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st590
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr960
		}
		goto st16
	st590:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof590
		}
	st_case_590:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr961
		case 13:
			goto tr859
		case 32:
			goto tr960
		case 44:
			goto tr50
		case 61:
			goto tr11
		case 92:
			goto st18
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st591
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr960
		}
		goto st16
	st591:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof591
		}
	st_case_591:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr961
		case 13:
			goto tr859
		case 32:
			goto tr960
		case 44:
			goto tr50
		case 61:
			goto tr11
		case 92:
			goto st18
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st592
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr960
		}
		goto st16
	st592:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof592
		}
	st_case_592:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr961
		case 13:
			goto tr859
		case 32:
			goto tr960
		case 44:
			goto tr50
		case 61:
			goto tr11
		case 92:
			goto st18
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st593
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr960
		}
		goto st16
	st593:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof593
		}
	st_case_593:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr961
		case 13:
			goto tr859
		case 32:
			goto tr960
		case 44:
			goto tr50
		case 61:
			goto tr11
		case 92:
			goto st18
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr960
		}
		goto st16
tr134:
//line machine.rl:17

	m.pb = m.p

	goto st55
	st55:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof55
		}
	st_case_55:
//line machine.go:11932
		switch ( m.data)[( m.p)] {
		case 34:
			goto st41
		case 92:
			goto st41
		}
		switch {
		case ( m.data)[( m.p)] > 10:
			if 12 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto tr48
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr48
		}
		goto st14
tr172:
//line machine.rl:17

	m.pb = m.p

	goto st56
	st56:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof56
		}
	st_case_56:
//line machine.go:11959
		switch ( m.data)[( m.p)] {
		case 34:
			goto st51
		case 92:
			goto st51
		}
		switch {
		case ( m.data)[( m.p)] > 10:
			if 12 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto tr48
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr48
		}
		goto st12
tr168:
//line machine.rl:70

	key = m.text()

	goto st57
tr298:
//line machine.rl:17

	m.pb = m.p

//line machine.rl:70

	key = m.text()

	goto st57
	st57:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof57
		}
	st_case_57:
//line machine.go:11996
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr160
		case 13:
			goto st6
		case 32:
			goto tr159
		case 34:
			goto tr181
		case 44:
			goto tr161
		case 45:
			goto tr182
		case 46:
			goto tr183
		case 48:
			goto tr184
		case 70:
			goto tr186
		case 84:
			goto tr187
		case 92:
			goto st125
		case 102:
			goto tr188
		case 116:
			goto tr189
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 49 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto tr185
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr159
		}
		goto st46
tr181:
//line machine.rl:86

	m.handler.AddString(key, m.text())

	goto st594
	st594:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof594
		}
	st_case_594:
//line machine.go:12047
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr882
		case 11:
			goto tr983
		case 13:
			goto tr882
		case 32:
			goto tr982
		case 34:
			goto tr70
		case 44:
			goto tr984
		case 92:
			goto tr72
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr982
		}
		goto tr67
tr1012:
//line machine.rl:58

	m.handler.SetMeasurement(m.text())

	goto st595
tr982:
//line machine.rl:58

	m.handler.SetMeasurement(m.text())

//line machine.rl:17

	m.pb = m.p

	goto st595
tr1135:
//line machine.rl:66

	m.handler.AddTag(key, m.text())

	goto st595
tr1025:
//line machine.rl:58

	m.handler.SetMeasurement(m.text())

//line machine.rl:78

	m.handler.AddFloat(key, m.text())

	goto st595
tr1130:
//line machine.rl:66

	m.handler.AddTag(key, m.text())

//line machine.rl:78

	m.handler.AddFloat(key, m.text())

	goto st595
tr1161:
//line machine.rl:66

	m.handler.AddTag(key, m.text())

//line machine.rl:74

	m.handler.AddInt(key, m.text())

	goto st595
tr1165:
//line machine.rl:66

	m.handler.AddTag(key, m.text())

//line machine.rl:82

	m.handler.AddBool(key, m.text())

	goto st595
tr1173:
//line machine.rl:58

	m.handler.SetMeasurement(m.text())

//line machine.rl:74

	m.handler.AddInt(key, m.text())

	goto st595
tr1176:
//line machine.rl:58

	m.handler.SetMeasurement(m.text())

//line machine.rl:82

	m.handler.AddBool(key, m.text())

	goto st595
	st595:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof595
		}
	st_case_595:
//line machine.go:12155
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr885
		case 11:
			goto tr986
		case 13:
			goto tr885
		case 32:
			goto st595
		case 34:
			goto tr82
		case 44:
			goto st6
		case 45:
			goto tr987
		case 61:
			goto st6
		case 92:
			goto tr83
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto tr988
			}
		case ( m.data)[( m.p)] >= 9:
			goto st595
		}
		goto tr79
tr986:
//line machine.rl:17

	m.pb = m.p

	goto st596
	st596:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof596
		}
	st_case_596:
//line machine.go:12196
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr885
		case 11:
			goto tr986
		case 13:
			goto tr885
		case 32:
			goto st595
		case 34:
			goto tr82
		case 44:
			goto st6
		case 45:
			goto tr987
		case 61:
			goto tr86
		case 92:
			goto tr83
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto tr988
			}
		case ( m.data)[( m.p)] >= 9:
			goto st595
		}
		goto tr79
tr987:
//line machine.rl:17

	m.pb = m.p

	goto st58
	st58:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof58
		}
	st_case_58:
//line machine.go:12237
		switch ( m.data)[( m.p)] {
		case 32:
			goto st6
		case 34:
			goto tr85
		case 44:
			goto st6
		case 61:
			goto tr86
		case 92:
			goto st60
		}
		switch {
		case ( m.data)[( m.p)] < 12:
			if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 10 {
				goto st6
			}
		case ( m.data)[( m.p)] > 13:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st597
			}
		default:
			goto st6
		}
		goto st25
tr988:
//line machine.rl:17

	m.pb = m.p

	goto st597
	st597:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof597
		}
	st_case_597:
//line machine.go:12274
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr990
		case 11:
			goto tr991
		case 13:
			goto tr990
		case 32:
			goto tr989
		case 34:
			goto tr85
		case 44:
			goto st6
		case 61:
			goto tr86
		case 92:
			goto st60
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st600
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr989
		}
		goto st25
tr989:
//line machine.rl:90

	m.handler.SetTimestamp(m.text())

	goto st598
	st598:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof598
		}
	st_case_598:
//line machine.go:12313
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr885
		case 13:
			goto tr885
		case 32:
			goto st598
		case 34:
			goto tr26
		case 92:
			goto st59
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto st598
		}
		goto st6
tr24:
//line machine.rl:17

	m.pb = m.p

	goto st59
	st59:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof59
		}
	st_case_59:
//line machine.go:12341
		switch ( m.data)[( m.p)] {
		case 34:
			goto st6
		case 92:
			goto st6
		}
		goto tr2
tr991:
//line machine.rl:90

	m.handler.SetTimestamp(m.text())

	goto st599
	st599:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof599
		}
	st_case_599:
//line machine.go:12360
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr885
		case 11:
			goto st599
		case 13:
			goto tr885
		case 32:
			goto st598
		case 34:
			goto tr85
		case 44:
			goto st6
		case 61:
			goto tr86
		case 92:
			goto st60
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto st598
		}
		goto st25
tr83:
//line machine.rl:17

	m.pb = m.p

	goto st60
	st60:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof60
		}
	st_case_60:
//line machine.go:12394
		switch ( m.data)[( m.p)] {
		case 34:
			goto st25
		case 92:
			goto st25
		}
		switch {
		case ( m.data)[( m.p)] > 10:
			if 12 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto tr2
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr2
		}
		goto st3
	st600:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof600
		}
	st_case_600:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr990
		case 11:
			goto tr991
		case 13:
			goto tr990
		case 32:
			goto tr989
		case 34:
			goto tr85
		case 44:
			goto st6
		case 61:
			goto tr86
		case 92:
			goto st60
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st601
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr989
		}
		goto st25
	st601:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof601
		}
	st_case_601:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr990
		case 11:
			goto tr991
		case 13:
			goto tr990
		case 32:
			goto tr989
		case 34:
			goto tr85
		case 44:
			goto st6
		case 61:
			goto tr86
		case 92:
			goto st60
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st602
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr989
		}
		goto st25
	st602:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof602
		}
	st_case_602:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr990
		case 11:
			goto tr991
		case 13:
			goto tr990
		case 32:
			goto tr989
		case 34:
			goto tr85
		case 44:
			goto st6
		case 61:
			goto tr86
		case 92:
			goto st60
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st603
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr989
		}
		goto st25
	st603:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof603
		}
	st_case_603:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr990
		case 11:
			goto tr991
		case 13:
			goto tr990
		case 32:
			goto tr989
		case 34:
			goto tr85
		case 44:
			goto st6
		case 61:
			goto tr86
		case 92:
			goto st60
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st604
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr989
		}
		goto st25
	st604:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof604
		}
	st_case_604:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr990
		case 11:
			goto tr991
		case 13:
			goto tr990
		case 32:
			goto tr989
		case 34:
			goto tr85
		case 44:
			goto st6
		case 61:
			goto tr86
		case 92:
			goto st60
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st605
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr989
		}
		goto st25
	st605:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof605
		}
	st_case_605:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr990
		case 11:
			goto tr991
		case 13:
			goto tr990
		case 32:
			goto tr989
		case 34:
			goto tr85
		case 44:
			goto st6
		case 61:
			goto tr86
		case 92:
			goto st60
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st606
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr989
		}
		goto st25
	st606:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof606
		}
	st_case_606:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr990
		case 11:
			goto tr991
		case 13:
			goto tr990
		case 32:
			goto tr989
		case 34:
			goto tr85
		case 44:
			goto st6
		case 61:
			goto tr86
		case 92:
			goto st60
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st607
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr989
		}
		goto st25
	st607:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof607
		}
	st_case_607:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr990
		case 11:
			goto tr991
		case 13:
			goto tr990
		case 32:
			goto tr989
		case 34:
			goto tr85
		case 44:
			goto st6
		case 61:
			goto tr86
		case 92:
			goto st60
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st608
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr989
		}
		goto st25
	st608:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof608
		}
	st_case_608:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr990
		case 11:
			goto tr991
		case 13:
			goto tr990
		case 32:
			goto tr989
		case 34:
			goto tr85
		case 44:
			goto st6
		case 61:
			goto tr86
		case 92:
			goto st60
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st609
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr989
		}
		goto st25
	st609:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof609
		}
	st_case_609:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr990
		case 11:
			goto tr991
		case 13:
			goto tr990
		case 32:
			goto tr989
		case 34:
			goto tr85
		case 44:
			goto st6
		case 61:
			goto tr86
		case 92:
			goto st60
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st610
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr989
		}
		goto st25
	st610:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof610
		}
	st_case_610:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr990
		case 11:
			goto tr991
		case 13:
			goto tr990
		case 32:
			goto tr989
		case 34:
			goto tr85
		case 44:
			goto st6
		case 61:
			goto tr86
		case 92:
			goto st60
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st611
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr989
		}
		goto st25
	st611:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof611
		}
	st_case_611:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr990
		case 11:
			goto tr991
		case 13:
			goto tr990
		case 32:
			goto tr989
		case 34:
			goto tr85
		case 44:
			goto st6
		case 61:
			goto tr86
		case 92:
			goto st60
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st612
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr989
		}
		goto st25
	st612:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof612
		}
	st_case_612:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr990
		case 11:
			goto tr991
		case 13:
			goto tr990
		case 32:
			goto tr989
		case 34:
			goto tr85
		case 44:
			goto st6
		case 61:
			goto tr86
		case 92:
			goto st60
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st613
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr989
		}
		goto st25
	st613:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof613
		}
	st_case_613:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr990
		case 11:
			goto tr991
		case 13:
			goto tr990
		case 32:
			goto tr989
		case 34:
			goto tr85
		case 44:
			goto st6
		case 61:
			goto tr86
		case 92:
			goto st60
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st614
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr989
		}
		goto st25
	st614:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof614
		}
	st_case_614:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr990
		case 11:
			goto tr991
		case 13:
			goto tr990
		case 32:
			goto tr989
		case 34:
			goto tr85
		case 44:
			goto st6
		case 61:
			goto tr86
		case 92:
			goto st60
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st615
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr989
		}
		goto st25
	st615:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof615
		}
	st_case_615:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr990
		case 11:
			goto tr991
		case 13:
			goto tr990
		case 32:
			goto tr989
		case 34:
			goto tr85
		case 44:
			goto st6
		case 61:
			goto tr86
		case 92:
			goto st60
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st616
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr989
		}
		goto st25
	st616:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof616
		}
	st_case_616:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr990
		case 11:
			goto tr991
		case 13:
			goto tr990
		case 32:
			goto tr989
		case 34:
			goto tr85
		case 44:
			goto st6
		case 61:
			goto tr86
		case 92:
			goto st60
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st617
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr989
		}
		goto st25
	st617:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof617
		}
	st_case_617:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr990
		case 11:
			goto tr991
		case 13:
			goto tr990
		case 32:
			goto tr989
		case 34:
			goto tr85
		case 44:
			goto st6
		case 61:
			goto tr86
		case 92:
			goto st60
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr989
		}
		goto st25
tr983:
//line machine.rl:58

	m.handler.SetMeasurement(m.text())

//line machine.rl:17

	m.pb = m.p

	goto st618
tr1027:
//line machine.rl:58

	m.handler.SetMeasurement(m.text())

//line machine.rl:78

	m.handler.AddFloat(key, m.text())

	goto st618
tr1174:
//line machine.rl:58

	m.handler.SetMeasurement(m.text())

//line machine.rl:74

	m.handler.AddInt(key, m.text())

	goto st618
tr1177:
//line machine.rl:58

	m.handler.SetMeasurement(m.text())

//line machine.rl:82

	m.handler.AddBool(key, m.text())

	goto st618
	st618:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof618
		}
	st_case_618:
//line machine.go:13026
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr885
		case 11:
			goto tr1013
		case 13:
			goto tr885
		case 32:
			goto tr1012
		case 34:
			goto tr106
		case 44:
			goto tr77
		case 45:
			goto tr1014
		case 61:
			goto st23
		case 92:
			goto tr107
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto tr1015
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1012
		}
		goto tr103
tr1013:
//line machine.rl:58

	m.handler.SetMeasurement(m.text())

//line machine.rl:17

	m.pb = m.p

	goto st619
	st619:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof619
		}
	st_case_619:
//line machine.go:13071
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr885
		case 11:
			goto tr1013
		case 13:
			goto tr885
		case 32:
			goto tr1012
		case 34:
			goto tr106
		case 44:
			goto tr77
		case 45:
			goto tr1014
		case 61:
			goto tr111
		case 92:
			goto tr107
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto tr1015
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1012
		}
		goto tr103
tr77:
//line machine.rl:58

	m.handler.SetMeasurement(m.text())

	goto st61
tr71:
//line machine.rl:58

	m.handler.SetMeasurement(m.text())

//line machine.rl:17

	m.pb = m.p

	goto st61
tr201:
//line machine.rl:66

	m.handler.AddTag(key, m.text())

	goto st61
	st61:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof61
		}
	st_case_61:
//line machine.go:13128
		switch ( m.data)[( m.p)] {
		case 32:
			goto st6
		case 34:
			goto tr171
		case 44:
			goto st6
		case 61:
			goto st6
		case 92:
			goto tr192
		}
		switch {
		case ( m.data)[( m.p)] > 10:
			if 12 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto st6
			}
		case ( m.data)[( m.p)] >= 9:
			goto st6
		}
		goto tr191
tr191:
//line machine.rl:17

	m.pb = m.p

	goto st62
	st62:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof62
		}
	st_case_62:
//line machine.go:13161
		switch ( m.data)[( m.p)] {
		case 32:
			goto st6
		case 34:
			goto tr174
		case 44:
			goto st6
		case 61:
			goto tr194
		case 92:
			goto st70
		}
		switch {
		case ( m.data)[( m.p)] > 10:
			if 12 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto st6
			}
		case ( m.data)[( m.p)] >= 9:
			goto st6
		}
		goto st62
tr194:
//line machine.rl:62

	key = m.text()

	goto st63
	st63:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof63
		}
	st_case_63:
//line machine.go:13194
		switch ( m.data)[( m.p)] {
		case 32:
			goto st6
		case 34:
			goto tr132
		case 44:
			goto st6
		case 61:
			goto st6
		case 92:
			goto tr197
		}
		switch {
		case ( m.data)[( m.p)] > 10:
			if 12 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto st6
			}
		case ( m.data)[( m.p)] >= 9:
			goto st6
		}
		goto tr196
tr196:
//line machine.rl:17

	m.pb = m.p

	goto st64
	st64:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof64
		}
	st_case_64:
//line machine.go:13227
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr200
		case 13:
			goto st6
		case 32:
			goto tr199
		case 34:
			goto tr138
		case 44:
			goto tr201
		case 61:
			goto st6
		case 92:
			goto st69
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr199
		}
		goto st64
tr200:
//line machine.rl:66

	m.handler.AddTag(key, m.text())

	goto st65
	st65:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof65
		}
	st_case_65:
//line machine.go:13261
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr204
		case 13:
			goto st6
		case 32:
			goto tr199
		case 34:
			goto tr205
		case 44:
			goto tr201
		case 61:
			goto st6
		case 92:
			goto tr206
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr199
		}
		goto tr203
tr203:
//line machine.rl:17

	m.pb = m.p

	goto st66
	st66:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof66
		}
	st_case_66:
//line machine.go:13295
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr208
		case 13:
			goto st6
		case 32:
			goto tr199
		case 34:
			goto tr209
		case 44:
			goto tr201
		case 61:
			goto tr86
		case 92:
			goto st68
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr199
		}
		goto st66
tr208:
//line machine.rl:66

	m.handler.AddTag(key, m.text())

	goto st67
tr204:
//line machine.rl:66

	m.handler.AddTag(key, m.text())

//line machine.rl:17

	m.pb = m.p

	goto st67
	st67:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof67
		}
	st_case_67:
//line machine.go:13339
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr204
		case 13:
			goto st6
		case 32:
			goto tr199
		case 34:
			goto tr205
		case 44:
			goto tr201
		case 61:
			goto tr86
		case 92:
			goto tr206
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr199
		}
		goto tr203
tr205:
//line machine.rl:17

	m.pb = m.p

//line machine.rl:86

	m.handler.AddString(key, m.text())

	goto st620
tr209:
//line machine.rl:86

	m.handler.AddString(key, m.text())

	goto st620
	st620:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof620
		}
	st_case_620:
//line machine.go:13383
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr850
		case 11:
			goto tr1016
		case 13:
			goto tr850
		case 32:
			goto tr954
		case 44:
			goto tr956
		case 61:
			goto tr11
		case 92:
			goto st18
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr954
		}
		goto st16
tr206:
//line machine.rl:17

	m.pb = m.p

	goto st68
	st68:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof68
		}
	st_case_68:
//line machine.go:13415
		switch ( m.data)[( m.p)] {
		case 34:
			goto st66
		case 92:
			goto st66
		}
		switch {
		case ( m.data)[( m.p)] > 10:
			if 12 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto tr48
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr48
		}
		goto st16
tr197:
//line machine.rl:17

	m.pb = m.p

	goto st69
	st69:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof69
		}
	st_case_69:
//line machine.go:13442
		switch ( m.data)[( m.p)] {
		case 34:
			goto st64
		case 92:
			goto st64
		}
		switch {
		case ( m.data)[( m.p)] > 10:
			if 12 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto tr48
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr48
		}
		goto st14
tr192:
//line machine.rl:17

	m.pb = m.p

	goto st70
	st70:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof70
		}
	st_case_70:
//line machine.go:13469
		switch ( m.data)[( m.p)] {
		case 34:
			goto st62
		case 92:
			goto st62
		}
		switch {
		case ( m.data)[( m.p)] > 10:
			if 12 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto tr48
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr48
		}
		goto st12
tr1014:
//line machine.rl:17

	m.pb = m.p

	goto st71
	st71:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof71
		}
	st_case_71:
//line machine.go:13496
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr109
		case 13:
			goto st6
		case 32:
			goto tr74
		case 34:
			goto tr110
		case 44:
			goto tr77
		case 61:
			goto tr111
		case 92:
			goto st74
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st621
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr74
		}
		goto st33
tr1015:
//line machine.rl:17

	m.pb = m.p

	goto st621
	st621:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof621
		}
	st_case_621:
//line machine.go:13535
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr990
		case 11:
			goto tr1018
		case 13:
			goto tr990
		case 32:
			goto tr1017
		case 34:
			goto tr110
		case 44:
			goto tr77
		case 61:
			goto tr111
		case 92:
			goto st74
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st746
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1017
		}
		goto st33
tr1022:
//line machine.rl:58

	m.handler.SetMeasurement(m.text())

	goto st622
tr1142:
//line machine.rl:66

	m.handler.AddTag(key, m.text())

	goto st622
tr1017:
//line machine.rl:58

	m.handler.SetMeasurement(m.text())

//line machine.rl:90

	m.handler.SetTimestamp(m.text())

	goto st622
tr1139:
//line machine.rl:66

	m.handler.AddTag(key, m.text())

//line machine.rl:90

	m.handler.SetTimestamp(m.text())

	goto st622
tr1543:
//line machine.rl:90

	m.handler.SetTimestamp(m.text())

//line machine.rl:58

	m.handler.SetMeasurement(m.text())

	goto st622
tr1735:
//line machine.rl:90

	m.handler.SetTimestamp(m.text())

//line machine.rl:66

	m.handler.AddTag(key, m.text())

	goto st622
	st622:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof622
		}
	st_case_622:
//line machine.go:13620
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr885
		case 11:
			goto tr1021
		case 13:
			goto tr885
		case 32:
			goto st622
		case 34:
			goto tr82
		case 44:
			goto st6
		case 61:
			goto st6
		case 92:
			goto tr83
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto st622
		}
		goto tr79
tr1021:
//line machine.rl:17

	m.pb = m.p

	goto st623
	st623:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof623
		}
	st_case_623:
//line machine.go:13654
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr885
		case 11:
			goto tr1021
		case 13:
			goto tr885
		case 32:
			goto st622
		case 34:
			goto tr82
		case 44:
			goto st6
		case 61:
			goto tr86
		case 92:
			goto tr83
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto st622
		}
		goto tr79
tr1023:
//line machine.rl:58

	m.handler.SetMeasurement(m.text())

//line machine.rl:17

	m.pb = m.p

	goto st624
tr1018:
//line machine.rl:58

	m.handler.SetMeasurement(m.text())

//line machine.rl:90

	m.handler.SetTimestamp(m.text())

	goto st624
	st624:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof624
		}
	st_case_624:
//line machine.go:13702
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr885
		case 11:
			goto tr1023
		case 13:
			goto tr885
		case 32:
			goto tr1022
		case 34:
			goto tr106
		case 44:
			goto tr77
		case 61:
			goto tr111
		case 92:
			goto tr107
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1022
		}
		goto tr103
tr111:
//line machine.rl:70

	key = m.text()

	goto st72
tr319:
//line machine.rl:17

	m.pb = m.p

//line machine.rl:70

	key = m.text()

	goto st72
	st72:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof72
		}
	st_case_72:
//line machine.go:13746
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr75
		case 13:
			goto st6
		case 32:
			goto tr74
		case 34:
			goto tr181
		case 44:
			goto tr77
		case 45:
			goto tr212
		case 46:
			goto tr213
		case 48:
			goto tr214
		case 70:
			goto tr216
		case 84:
			goto tr217
		case 92:
			goto st113
		case 102:
			goto tr218
		case 116:
			goto tr219
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 49 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto tr215
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr74
		}
		goto st23
tr75:
//line machine.rl:58

	m.handler.SetMeasurement(m.text())

	goto st73
tr69:
//line machine.rl:58

	m.handler.SetMeasurement(m.text())

//line machine.rl:17

	m.pb = m.p

	goto st73
	st73:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof73
		}
	st_case_73:
//line machine.go:13807
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr113
		case 13:
			goto st6
		case 32:
			goto tr74
		case 34:
			goto tr106
		case 44:
			goto tr77
		case 61:
			goto st23
		case 92:
			goto tr107
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr74
		}
		goto tr103
tr107:
//line machine.rl:17

	m.pb = m.p

	goto st74
	st74:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof74
		}
	st_case_74:
//line machine.go:13841
		switch ( m.data)[( m.p)] {
		case 34:
			goto st33
		case 92:
			goto st25
		}
		switch {
		case ( m.data)[( m.p)] > 10:
			if 12 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto tr2
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr2
		}
		goto st9
tr212:
//line machine.rl:17

	m.pb = m.p

	goto st75
	st75:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof75
		}
	st_case_75:
//line machine.go:13868
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr75
		case 13:
			goto st6
		case 32:
			goto tr74
		case 34:
			goto tr76
		case 44:
			goto tr77
		case 48:
			goto st627
		case 92:
			goto st113
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 49 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st740
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr74
		}
		goto st23
tr70:
//line machine.rl:17

	m.pb = m.p

//line machine.rl:86

	m.handler.AddString(key, m.text())

	goto st625
tr76:
//line machine.rl:86

	m.handler.AddString(key, m.text())

	goto st625
	st625:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof625
		}
	st_case_625:
//line machine.go:13917
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr850
		case 11:
			goto tr1024
		case 13:
			goto tr850
		case 32:
			goto tr889
		case 44:
			goto tr891
		case 92:
			goto st76
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr889
		}
		goto st7
tr1024:
//line machine.rl:58

	m.handler.SetMeasurement(m.text())

	goto st626
tr1318:
//line machine.rl:58

	m.handler.SetMeasurement(m.text())

//line machine.rl:78

	m.handler.AddFloat(key, m.text())

	goto st626
tr1323:
//line machine.rl:58

	m.handler.SetMeasurement(m.text())

//line machine.rl:74

	m.handler.AddInt(key, m.text())

	goto st626
tr1326:
//line machine.rl:58

	m.handler.SetMeasurement(m.text())

//line machine.rl:82

	m.handler.AddBool(key, m.text())

	goto st626
	st626:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof626
		}
	st_case_626:
//line machine.go:13977
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr850
		case 11:
			goto tr894
		case 13:
			goto tr850
		case 32:
			goto tr889
		case 44:
			goto tr4
		case 45:
			goto tr895
		case 61:
			goto st7
		case 92:
			goto tr33
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto tr896
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr889
		}
		goto tr31
tr344:
//line machine.rl:17

	m.pb = m.p

	goto st76
	st76:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof76
		}
	st_case_76:
//line machine.go:14016
		if ( m.data)[( m.p)] == 92 {
			goto st0
		}
		switch {
		case ( m.data)[( m.p)] > 10:
			if 12 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto st0
			}
		case ( m.data)[( m.p)] >= 9:
			goto st0
		}
		goto st7
tr214:
//line machine.rl:17

	m.pb = m.p

	goto st627
	st627:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof627
		}
	st_case_627:
//line machine.go:14040
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1026
		case 11:
			goto tr1027
		case 13:
			goto tr1026
		case 32:
			goto tr1025
		case 34:
			goto tr76
		case 44:
			goto tr1028
		case 46:
			goto st738
		case 92:
			goto st113
		case 105:
			goto st739
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1025
		}
		goto st23
tr984:
//line machine.rl:58

	m.handler.SetMeasurement(m.text())

//line machine.rl:17

	m.pb = m.p

	goto st77
tr1028:
//line machine.rl:58

	m.handler.SetMeasurement(m.text())

//line machine.rl:78

	m.handler.AddFloat(key, m.text())

	goto st77
tr1132:
//line machine.rl:66

	m.handler.AddTag(key, m.text())

//line machine.rl:78

	m.handler.AddFloat(key, m.text())

	goto st77
tr1164:
//line machine.rl:66

	m.handler.AddTag(key, m.text())

//line machine.rl:74

	m.handler.AddInt(key, m.text())

	goto st77
tr1168:
//line machine.rl:66

	m.handler.AddTag(key, m.text())

//line machine.rl:82

	m.handler.AddBool(key, m.text())

	goto st77
tr1175:
//line machine.rl:58

	m.handler.SetMeasurement(m.text())

//line machine.rl:74

	m.handler.AddInt(key, m.text())

	goto st77
tr1178:
//line machine.rl:58

	m.handler.SetMeasurement(m.text())

//line machine.rl:82

	m.handler.AddBool(key, m.text())

	goto st77
	st77:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof77
		}
	st_case_77:
//line machine.go:14140
		switch ( m.data)[( m.p)] {
		case 32:
			goto st6
		case 34:
			goto tr224
		case 44:
			goto st6
		case 61:
			goto st6
		case 92:
			goto tr225
		}
		switch {
		case ( m.data)[( m.p)] > 10:
			if 12 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto st6
			}
		case ( m.data)[( m.p)] >= 9:
			goto st6
		}
		goto tr223
tr223:
//line machine.rl:17

	m.pb = m.p

	goto st78
	st78:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof78
		}
	st_case_78:
//line machine.go:14173
		switch ( m.data)[( m.p)] {
		case 32:
			goto st6
		case 34:
			goto tr227
		case 44:
			goto st6
		case 61:
			goto tr228
		case 92:
			goto st112
		}
		switch {
		case ( m.data)[( m.p)] > 10:
			if 12 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto st6
			}
		case ( m.data)[( m.p)] >= 9:
			goto st6
		}
		goto st78
tr224:
//line machine.rl:17

	m.pb = m.p

//line machine.rl:86

	m.handler.AddString(key, m.text())

	goto st628
tr227:
//line machine.rl:86

	m.handler.AddString(key, m.text())

	goto st628
	st628:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof628
		}
	st_case_628:
//line machine.go:14216
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr850
		case 11:
			goto st629
		case 13:
			goto tr850
		case 32:
			goto st492
		case 44:
			goto st28
		case 61:
			goto tr118
		case 92:
			goto st80
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto st492
		}
		goto st38
	st629:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof629
		}
	st_case_629:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr850
		case 11:
			goto st629
		case 13:
			goto tr850
		case 32:
			goto st492
		case 44:
			goto tr179
		case 45:
			goto tr1032
		case 61:
			goto tr118
		case 92:
			goto st80
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto tr1033
			}
		case ( m.data)[( m.p)] >= 9:
			goto st492
		}
		goto st38
tr1032:
//line machine.rl:17

	m.pb = m.p

	goto st79
	st79:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof79
		}
	st_case_79:
//line machine.go:14280
		switch ( m.data)[( m.p)] {
		case 32:
			goto tr179
		case 44:
			goto tr179
		case 61:
			goto tr118
		case 92:
			goto st80
		}
		switch {
		case ( m.data)[( m.p)] < 12:
			if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 10 {
				goto tr179
			}
		case ( m.data)[( m.p)] > 13:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st630
			}
		default:
			goto tr179
		}
		goto st38
tr1033:
//line machine.rl:17

	m.pb = m.p

	goto st630
	st630:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof630
		}
	st_case_630:
//line machine.go:14315
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr1034
		case 13:
			goto tr859
		case 32:
			goto tr858
		case 44:
			goto tr179
		case 61:
			goto tr118
		case 92:
			goto st80
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st632
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr858
		}
		goto st38
tr1034:
//line machine.rl:90

	m.handler.SetTimestamp(m.text())

	goto st631
	st631:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof631
		}
	st_case_631:
//line machine.go:14352
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr850
		case 11:
			goto st631
		case 13:
			goto tr850
		case 32:
			goto st497
		case 44:
			goto tr48
		case 61:
			goto tr118
		case 92:
			goto st80
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto st497
		}
		goto st38
tr116:
//line machine.rl:17

	m.pb = m.p

	goto st80
	st80:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof80
		}
	st_case_80:
//line machine.go:14384
		switch {
		case ( m.data)[( m.p)] > 10:
			if 12 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto tr48
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr48
		}
		goto st38
	st632:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof632
		}
	st_case_632:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr1034
		case 13:
			goto tr859
		case 32:
			goto tr858
		case 44:
			goto tr179
		case 61:
			goto tr118
		case 92:
			goto st80
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st633
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr858
		}
		goto st38
	st633:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof633
		}
	st_case_633:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr1034
		case 13:
			goto tr859
		case 32:
			goto tr858
		case 44:
			goto tr179
		case 61:
			goto tr118
		case 92:
			goto st80
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st634
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr858
		}
		goto st38
	st634:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof634
		}
	st_case_634:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr1034
		case 13:
			goto tr859
		case 32:
			goto tr858
		case 44:
			goto tr179
		case 61:
			goto tr118
		case 92:
			goto st80
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st635
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr858
		}
		goto st38
	st635:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof635
		}
	st_case_635:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr1034
		case 13:
			goto tr859
		case 32:
			goto tr858
		case 44:
			goto tr179
		case 61:
			goto tr118
		case 92:
			goto st80
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st636
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr858
		}
		goto st38
	st636:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof636
		}
	st_case_636:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr1034
		case 13:
			goto tr859
		case 32:
			goto tr858
		case 44:
			goto tr179
		case 61:
			goto tr118
		case 92:
			goto st80
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st637
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr858
		}
		goto st38
	st637:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof637
		}
	st_case_637:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr1034
		case 13:
			goto tr859
		case 32:
			goto tr858
		case 44:
			goto tr179
		case 61:
			goto tr118
		case 92:
			goto st80
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st638
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr858
		}
		goto st38
	st638:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof638
		}
	st_case_638:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr1034
		case 13:
			goto tr859
		case 32:
			goto tr858
		case 44:
			goto tr179
		case 61:
			goto tr118
		case 92:
			goto st80
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st639
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr858
		}
		goto st38
	st639:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof639
		}
	st_case_639:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr1034
		case 13:
			goto tr859
		case 32:
			goto tr858
		case 44:
			goto tr179
		case 61:
			goto tr118
		case 92:
			goto st80
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st640
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr858
		}
		goto st38
	st640:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof640
		}
	st_case_640:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr1034
		case 13:
			goto tr859
		case 32:
			goto tr858
		case 44:
			goto tr179
		case 61:
			goto tr118
		case 92:
			goto st80
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st641
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr858
		}
		goto st38
	st641:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof641
		}
	st_case_641:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr1034
		case 13:
			goto tr859
		case 32:
			goto tr858
		case 44:
			goto tr179
		case 61:
			goto tr118
		case 92:
			goto st80
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st642
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr858
		}
		goto st38
	st642:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof642
		}
	st_case_642:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr1034
		case 13:
			goto tr859
		case 32:
			goto tr858
		case 44:
			goto tr179
		case 61:
			goto tr118
		case 92:
			goto st80
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st643
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr858
		}
		goto st38
	st643:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof643
		}
	st_case_643:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr1034
		case 13:
			goto tr859
		case 32:
			goto tr858
		case 44:
			goto tr179
		case 61:
			goto tr118
		case 92:
			goto st80
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st644
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr858
		}
		goto st38
	st644:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof644
		}
	st_case_644:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr1034
		case 13:
			goto tr859
		case 32:
			goto tr858
		case 44:
			goto tr179
		case 61:
			goto tr118
		case 92:
			goto st80
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st645
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr858
		}
		goto st38
	st645:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof645
		}
	st_case_645:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr1034
		case 13:
			goto tr859
		case 32:
			goto tr858
		case 44:
			goto tr179
		case 61:
			goto tr118
		case 92:
			goto st80
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st646
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr858
		}
		goto st38
	st646:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof646
		}
	st_case_646:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr1034
		case 13:
			goto tr859
		case 32:
			goto tr858
		case 44:
			goto tr179
		case 61:
			goto tr118
		case 92:
			goto st80
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st647
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr858
		}
		goto st38
	st647:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof647
		}
	st_case_647:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr1034
		case 13:
			goto tr859
		case 32:
			goto tr858
		case 44:
			goto tr179
		case 61:
			goto tr118
		case 92:
			goto st80
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st648
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr858
		}
		goto st38
	st648:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof648
		}
	st_case_648:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr1034
		case 13:
			goto tr859
		case 32:
			goto tr858
		case 44:
			goto tr179
		case 61:
			goto tr118
		case 92:
			goto st80
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st649
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr858
		}
		goto st38
	st649:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof649
		}
	st_case_649:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr1034
		case 13:
			goto tr859
		case 32:
			goto tr858
		case 44:
			goto tr179
		case 61:
			goto tr118
		case 92:
			goto st80
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr858
		}
		goto st38
tr228:
//line machine.rl:62

	key = m.text()

//line machine.rl:70

	key = m.text()

	goto st81
	st81:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof81
		}
	st_case_81:
//line machine.go:14944
		switch ( m.data)[( m.p)] {
		case 32:
			goto st6
		case 34:
			goto tr231
		case 44:
			goto st6
		case 45:
			goto tr232
		case 46:
			goto tr233
		case 48:
			goto tr234
		case 61:
			goto st6
		case 70:
			goto tr236
		case 84:
			goto tr237
		case 92:
			goto tr197
		case 102:
			goto tr238
		case 116:
			goto tr239
		}
		switch {
		case ( m.data)[( m.p)] < 12:
			if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 10 {
				goto st6
			}
		case ( m.data)[( m.p)] > 13:
			if 49 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto tr235
			}
		default:
			goto st6
		}
		goto tr196
tr231:
//line machine.rl:17

	m.pb = m.p

//line machine.rl:86

	m.handler.AddString(key, m.text())

	goto st650
	st650:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof650
		}
	st_case_650:
//line machine.go:14999
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1055
		case 11:
			goto tr1056
		case 13:
			goto tr1055
		case 32:
			goto tr1054
		case 34:
			goto tr132
		case 44:
			goto tr1057
		case 61:
			goto tr22
		case 92:
			goto tr134
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1054
		}
		goto tr129
tr1205:
//line machine.rl:58

	m.handler.SetMeasurement(m.text())

	goto st651
tr1085:
//line machine.rl:66

	m.handler.AddTag(key, m.text())

	goto st651
tr1054:
//line machine.rl:66

	m.handler.AddTag(key, m.text())

//line machine.rl:17

	m.pb = m.p

	goto st651
tr1200:
//line machine.rl:58

	m.handler.SetMeasurement(m.text())

//line machine.rl:78

	m.handler.AddFloat(key, m.text())

	goto st651
tr1113:
//line machine.rl:66

	m.handler.AddTag(key, m.text())

//line machine.rl:78

	m.handler.AddFloat(key, m.text())

	goto st651
tr1118:
//line machine.rl:66

	m.handler.AddTag(key, m.text())

//line machine.rl:74

	m.handler.AddInt(key, m.text())

	goto st651
tr1122:
//line machine.rl:66

	m.handler.AddTag(key, m.text())

//line machine.rl:82

	m.handler.AddBool(key, m.text())

	goto st651
tr1231:
//line machine.rl:58

	m.handler.SetMeasurement(m.text())

//line machine.rl:74

	m.handler.AddInt(key, m.text())

	goto st651
tr1234:
//line machine.rl:58

	m.handler.SetMeasurement(m.text())

//line machine.rl:82

	m.handler.AddBool(key, m.text())

	goto st651
	st651:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof651
		}
	st_case_651:
//line machine.go:15109
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr927
		case 11:
			goto tr1059
		case 13:
			goto tr927
		case 32:
			goto st651
		case 34:
			goto tr82
		case 44:
			goto st6
		case 45:
			goto tr1060
		case 61:
			goto st6
		case 92:
			goto tr144
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto tr1061
			}
		case ( m.data)[( m.p)] >= 9:
			goto st651
		}
		goto tr141
tr1059:
//line machine.rl:17

	m.pb = m.p

	goto st652
	st652:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof652
		}
	st_case_652:
//line machine.go:15150
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr927
		case 11:
			goto tr1059
		case 13:
			goto tr927
		case 32:
			goto st651
		case 34:
			goto tr82
		case 44:
			goto st6
		case 45:
			goto tr1060
		case 61:
			goto tr146
		case 92:
			goto tr144
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto tr1061
			}
		case ( m.data)[( m.p)] >= 9:
			goto st651
		}
		goto tr141
tr1060:
//line machine.rl:17

	m.pb = m.p

	goto st82
	st82:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof82
		}
	st_case_82:
//line machine.go:15191
		switch ( m.data)[( m.p)] {
		case 32:
			goto st6
		case 34:
			goto tr85
		case 44:
			goto st6
		case 61:
			goto tr146
		case 92:
			goto st83
		}
		switch {
		case ( m.data)[( m.p)] < 12:
			if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 10 {
				goto st6
			}
		case ( m.data)[( m.p)] > 13:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st653
			}
		default:
			goto st6
		}
		goto st43
tr1061:
//line machine.rl:17

	m.pb = m.p

	goto st653
	st653:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof653
		}
	st_case_653:
//line machine.go:15228
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1063
		case 11:
			goto tr1064
		case 13:
			goto tr1063
		case 32:
			goto tr1062
		case 34:
			goto tr85
		case 44:
			goto st6
		case 61:
			goto tr146
		case 92:
			goto st83
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st656
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1062
		}
		goto st43
tr1062:
//line machine.rl:90

	m.handler.SetTimestamp(m.text())

	goto st654
	st654:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof654
		}
	st_case_654:
//line machine.go:15267
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr927
		case 13:
			goto tr927
		case 32:
			goto st654
		case 34:
			goto tr26
		case 92:
			goto st59
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto st654
		}
		goto st6
tr1064:
//line machine.rl:90

	m.handler.SetTimestamp(m.text())

	goto st655
	st655:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof655
		}
	st_case_655:
//line machine.go:15295
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr927
		case 11:
			goto st655
		case 13:
			goto tr927
		case 32:
			goto st654
		case 34:
			goto tr85
		case 44:
			goto st6
		case 61:
			goto tr146
		case 92:
			goto st83
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto st654
		}
		goto st43
tr144:
//line machine.rl:17

	m.pb = m.p

	goto st83
	st83:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof83
		}
	st_case_83:
//line machine.go:15329
		switch ( m.data)[( m.p)] {
		case 34:
			goto st43
		case 92:
			goto st43
		}
		switch {
		case ( m.data)[( m.p)] > 10:
			if 12 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto tr2
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr2
		}
		goto st3
	st656:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof656
		}
	st_case_656:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1063
		case 11:
			goto tr1064
		case 13:
			goto tr1063
		case 32:
			goto tr1062
		case 34:
			goto tr85
		case 44:
			goto st6
		case 61:
			goto tr146
		case 92:
			goto st83
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st657
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1062
		}
		goto st43
	st657:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof657
		}
	st_case_657:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1063
		case 11:
			goto tr1064
		case 13:
			goto tr1063
		case 32:
			goto tr1062
		case 34:
			goto tr85
		case 44:
			goto st6
		case 61:
			goto tr146
		case 92:
			goto st83
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st658
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1062
		}
		goto st43
	st658:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof658
		}
	st_case_658:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1063
		case 11:
			goto tr1064
		case 13:
			goto tr1063
		case 32:
			goto tr1062
		case 34:
			goto tr85
		case 44:
			goto st6
		case 61:
			goto tr146
		case 92:
			goto st83
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st659
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1062
		}
		goto st43
	st659:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof659
		}
	st_case_659:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1063
		case 11:
			goto tr1064
		case 13:
			goto tr1063
		case 32:
			goto tr1062
		case 34:
			goto tr85
		case 44:
			goto st6
		case 61:
			goto tr146
		case 92:
			goto st83
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st660
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1062
		}
		goto st43
	st660:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof660
		}
	st_case_660:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1063
		case 11:
			goto tr1064
		case 13:
			goto tr1063
		case 32:
			goto tr1062
		case 34:
			goto tr85
		case 44:
			goto st6
		case 61:
			goto tr146
		case 92:
			goto st83
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st661
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1062
		}
		goto st43
	st661:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof661
		}
	st_case_661:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1063
		case 11:
			goto tr1064
		case 13:
			goto tr1063
		case 32:
			goto tr1062
		case 34:
			goto tr85
		case 44:
			goto st6
		case 61:
			goto tr146
		case 92:
			goto st83
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st662
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1062
		}
		goto st43
	st662:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof662
		}
	st_case_662:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1063
		case 11:
			goto tr1064
		case 13:
			goto tr1063
		case 32:
			goto tr1062
		case 34:
			goto tr85
		case 44:
			goto st6
		case 61:
			goto tr146
		case 92:
			goto st83
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st663
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1062
		}
		goto st43
	st663:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof663
		}
	st_case_663:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1063
		case 11:
			goto tr1064
		case 13:
			goto tr1063
		case 32:
			goto tr1062
		case 34:
			goto tr85
		case 44:
			goto st6
		case 61:
			goto tr146
		case 92:
			goto st83
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st664
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1062
		}
		goto st43
	st664:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof664
		}
	st_case_664:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1063
		case 11:
			goto tr1064
		case 13:
			goto tr1063
		case 32:
			goto tr1062
		case 34:
			goto tr85
		case 44:
			goto st6
		case 61:
			goto tr146
		case 92:
			goto st83
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st665
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1062
		}
		goto st43
	st665:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof665
		}
	st_case_665:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1063
		case 11:
			goto tr1064
		case 13:
			goto tr1063
		case 32:
			goto tr1062
		case 34:
			goto tr85
		case 44:
			goto st6
		case 61:
			goto tr146
		case 92:
			goto st83
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st666
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1062
		}
		goto st43
	st666:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof666
		}
	st_case_666:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1063
		case 11:
			goto tr1064
		case 13:
			goto tr1063
		case 32:
			goto tr1062
		case 34:
			goto tr85
		case 44:
			goto st6
		case 61:
			goto tr146
		case 92:
			goto st83
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st667
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1062
		}
		goto st43
	st667:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof667
		}
	st_case_667:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1063
		case 11:
			goto tr1064
		case 13:
			goto tr1063
		case 32:
			goto tr1062
		case 34:
			goto tr85
		case 44:
			goto st6
		case 61:
			goto tr146
		case 92:
			goto st83
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st668
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1062
		}
		goto st43
	st668:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof668
		}
	st_case_668:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1063
		case 11:
			goto tr1064
		case 13:
			goto tr1063
		case 32:
			goto tr1062
		case 34:
			goto tr85
		case 44:
			goto st6
		case 61:
			goto tr146
		case 92:
			goto st83
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st669
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1062
		}
		goto st43
	st669:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof669
		}
	st_case_669:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1063
		case 11:
			goto tr1064
		case 13:
			goto tr1063
		case 32:
			goto tr1062
		case 34:
			goto tr85
		case 44:
			goto st6
		case 61:
			goto tr146
		case 92:
			goto st83
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st670
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1062
		}
		goto st43
	st670:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof670
		}
	st_case_670:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1063
		case 11:
			goto tr1064
		case 13:
			goto tr1063
		case 32:
			goto tr1062
		case 34:
			goto tr85
		case 44:
			goto st6
		case 61:
			goto tr146
		case 92:
			goto st83
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st671
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1062
		}
		goto st43
	st671:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof671
		}
	st_case_671:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1063
		case 11:
			goto tr1064
		case 13:
			goto tr1063
		case 32:
			goto tr1062
		case 34:
			goto tr85
		case 44:
			goto st6
		case 61:
			goto tr146
		case 92:
			goto st83
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st672
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1062
		}
		goto st43
	st672:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof672
		}
	st_case_672:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1063
		case 11:
			goto tr1064
		case 13:
			goto tr1063
		case 32:
			goto tr1062
		case 34:
			goto tr85
		case 44:
			goto st6
		case 61:
			goto tr146
		case 92:
			goto st83
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st673
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1062
		}
		goto st43
	st673:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof673
		}
	st_case_673:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1063
		case 11:
			goto tr1064
		case 13:
			goto tr1063
		case 32:
			goto tr1062
		case 34:
			goto tr85
		case 44:
			goto st6
		case 61:
			goto tr146
		case 92:
			goto st83
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1062
		}
		goto st43
tr1056:
//line machine.rl:66

	m.handler.AddTag(key, m.text())

//line machine.rl:17

	m.pb = m.p

	goto st674
tr1114:
//line machine.rl:66

	m.handler.AddTag(key, m.text())

//line machine.rl:78

	m.handler.AddFloat(key, m.text())

	goto st674
tr1120:
//line machine.rl:66

	m.handler.AddTag(key, m.text())

//line machine.rl:74

	m.handler.AddInt(key, m.text())

	goto st674
tr1124:
//line machine.rl:66

	m.handler.AddTag(key, m.text())

//line machine.rl:82

	m.handler.AddBool(key, m.text())

	goto st674
	st674:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof674
		}
	st_case_674:
//line machine.go:15961
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr927
		case 11:
			goto tr1086
		case 13:
			goto tr927
		case 32:
			goto tr1085
		case 34:
			goto tr205
		case 44:
			goto tr139
		case 45:
			goto tr1087
		case 61:
			goto st6
		case 92:
			goto tr246
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto tr1088
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1085
		}
		goto tr244
tr244:
//line machine.rl:17

	m.pb = m.p

	goto st84
	st84:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof84
		}
	st_case_84:
//line machine.go:16002
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr242
		case 13:
			goto st6
		case 32:
			goto tr136
		case 34:
			goto tr209
		case 44:
			goto tr139
		case 61:
			goto tr146
		case 92:
			goto st86
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr136
		}
		goto st84
tr242:
//line machine.rl:66

	m.handler.AddTag(key, m.text())

	goto st85
tr245:
//line machine.rl:66

	m.handler.AddTag(key, m.text())

//line machine.rl:17

	m.pb = m.p

	goto st85
	st85:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof85
		}
	st_case_85:
//line machine.go:16046
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr245
		case 13:
			goto st6
		case 32:
			goto tr136
		case 34:
			goto tr205
		case 44:
			goto tr139
		case 61:
			goto tr146
		case 92:
			goto tr246
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr136
		}
		goto tr244
tr246:
//line machine.rl:17

	m.pb = m.p

	goto st86
	st86:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof86
		}
	st_case_86:
//line machine.go:16080
		switch ( m.data)[( m.p)] {
		case 34:
			goto st84
		case 92:
			goto st84
		}
		switch {
		case ( m.data)[( m.p)] > 10:
			if 12 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto tr48
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr48
		}
		goto st16
tr1086:
//line machine.rl:66

	m.handler.AddTag(key, m.text())

//line machine.rl:17

	m.pb = m.p

	goto st675
	st675:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof675
		}
	st_case_675:
//line machine.go:16111
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr927
		case 11:
			goto tr1086
		case 13:
			goto tr927
		case 32:
			goto tr1085
		case 34:
			goto tr205
		case 44:
			goto tr139
		case 45:
			goto tr1087
		case 61:
			goto tr146
		case 92:
			goto tr246
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto tr1088
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1085
		}
		goto tr244
tr1087:
//line machine.rl:17

	m.pb = m.p

	goto st87
	st87:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof87
		}
	st_case_87:
//line machine.go:16152
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr242
		case 13:
			goto st6
		case 32:
			goto tr136
		case 34:
			goto tr209
		case 44:
			goto tr139
		case 61:
			goto tr146
		case 92:
			goto st86
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st676
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr136
		}
		goto st84
tr1088:
//line machine.rl:17

	m.pb = m.p

	goto st676
	st676:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof676
		}
	st_case_676:
//line machine.go:16191
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1063
		case 11:
			goto tr1090
		case 13:
			goto tr1063
		case 32:
			goto tr1089
		case 34:
			goto tr209
		case 44:
			goto tr139
		case 61:
			goto tr146
		case 92:
			goto st86
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st680
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1089
		}
		goto st84
tr1212:
//line machine.rl:58

	m.handler.SetMeasurement(m.text())

	goto st677
tr1094:
//line machine.rl:66

	m.handler.AddTag(key, m.text())

	goto st677
tr1209:
//line machine.rl:58

	m.handler.SetMeasurement(m.text())

//line machine.rl:90

	m.handler.SetTimestamp(m.text())

	goto st677
tr1089:
//line machine.rl:66

	m.handler.AddTag(key, m.text())

//line machine.rl:90

	m.handler.SetTimestamp(m.text())

	goto st677
tr1676:
//line machine.rl:90

	m.handler.SetTimestamp(m.text())

//line machine.rl:66

	m.handler.AddTag(key, m.text())

	goto st677
	st677:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof677
		}
	st_case_677:
//line machine.go:16266
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr927
		case 11:
			goto tr1093
		case 13:
			goto tr927
		case 32:
			goto st677
		case 34:
			goto tr82
		case 44:
			goto st6
		case 61:
			goto st6
		case 92:
			goto tr144
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto st677
		}
		goto tr141
tr1093:
//line machine.rl:17

	m.pb = m.p

	goto st678
	st678:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof678
		}
	st_case_678:
//line machine.go:16300
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr927
		case 11:
			goto tr1093
		case 13:
			goto tr927
		case 32:
			goto st677
		case 34:
			goto tr82
		case 44:
			goto st6
		case 61:
			goto tr146
		case 92:
			goto tr144
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto st677
		}
		goto tr141
tr1095:
//line machine.rl:66

	m.handler.AddTag(key, m.text())

//line machine.rl:17

	m.pb = m.p

	goto st679
tr1090:
//line machine.rl:66

	m.handler.AddTag(key, m.text())

//line machine.rl:90

	m.handler.SetTimestamp(m.text())

	goto st679
tr1677:
//line machine.rl:90

	m.handler.SetTimestamp(m.text())

//line machine.rl:66

	m.handler.AddTag(key, m.text())

	goto st679
	st679:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof679
		}
	st_case_679:
//line machine.go:16358
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr927
		case 11:
			goto tr1095
		case 13:
			goto tr927
		case 32:
			goto tr1094
		case 34:
			goto tr205
		case 44:
			goto tr139
		case 61:
			goto tr146
		case 92:
			goto tr246
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1094
		}
		goto tr244
	st680:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof680
		}
	st_case_680:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1063
		case 11:
			goto tr1090
		case 13:
			goto tr1063
		case 32:
			goto tr1089
		case 34:
			goto tr209
		case 44:
			goto tr139
		case 61:
			goto tr146
		case 92:
			goto st86
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st681
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1089
		}
		goto st84
	st681:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof681
		}
	st_case_681:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1063
		case 11:
			goto tr1090
		case 13:
			goto tr1063
		case 32:
			goto tr1089
		case 34:
			goto tr209
		case 44:
			goto tr139
		case 61:
			goto tr146
		case 92:
			goto st86
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st682
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1089
		}
		goto st84
	st682:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof682
		}
	st_case_682:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1063
		case 11:
			goto tr1090
		case 13:
			goto tr1063
		case 32:
			goto tr1089
		case 34:
			goto tr209
		case 44:
			goto tr139
		case 61:
			goto tr146
		case 92:
			goto st86
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st683
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1089
		}
		goto st84
	st683:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof683
		}
	st_case_683:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1063
		case 11:
			goto tr1090
		case 13:
			goto tr1063
		case 32:
			goto tr1089
		case 34:
			goto tr209
		case 44:
			goto tr139
		case 61:
			goto tr146
		case 92:
			goto st86
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st684
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1089
		}
		goto st84
	st684:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof684
		}
	st_case_684:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1063
		case 11:
			goto tr1090
		case 13:
			goto tr1063
		case 32:
			goto tr1089
		case 34:
			goto tr209
		case 44:
			goto tr139
		case 61:
			goto tr146
		case 92:
			goto st86
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st685
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1089
		}
		goto st84
	st685:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof685
		}
	st_case_685:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1063
		case 11:
			goto tr1090
		case 13:
			goto tr1063
		case 32:
			goto tr1089
		case 34:
			goto tr209
		case 44:
			goto tr139
		case 61:
			goto tr146
		case 92:
			goto st86
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st686
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1089
		}
		goto st84
	st686:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof686
		}
	st_case_686:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1063
		case 11:
			goto tr1090
		case 13:
			goto tr1063
		case 32:
			goto tr1089
		case 34:
			goto tr209
		case 44:
			goto tr139
		case 61:
			goto tr146
		case 92:
			goto st86
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st687
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1089
		}
		goto st84
	st687:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof687
		}
	st_case_687:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1063
		case 11:
			goto tr1090
		case 13:
			goto tr1063
		case 32:
			goto tr1089
		case 34:
			goto tr209
		case 44:
			goto tr139
		case 61:
			goto tr146
		case 92:
			goto st86
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st688
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1089
		}
		goto st84
	st688:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof688
		}
	st_case_688:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1063
		case 11:
			goto tr1090
		case 13:
			goto tr1063
		case 32:
			goto tr1089
		case 34:
			goto tr209
		case 44:
			goto tr139
		case 61:
			goto tr146
		case 92:
			goto st86
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st689
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1089
		}
		goto st84
	st689:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof689
		}
	st_case_689:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1063
		case 11:
			goto tr1090
		case 13:
			goto tr1063
		case 32:
			goto tr1089
		case 34:
			goto tr209
		case 44:
			goto tr139
		case 61:
			goto tr146
		case 92:
			goto st86
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st690
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1089
		}
		goto st84
	st690:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof690
		}
	st_case_690:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1063
		case 11:
			goto tr1090
		case 13:
			goto tr1063
		case 32:
			goto tr1089
		case 34:
			goto tr209
		case 44:
			goto tr139
		case 61:
			goto tr146
		case 92:
			goto st86
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st691
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1089
		}
		goto st84
	st691:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof691
		}
	st_case_691:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1063
		case 11:
			goto tr1090
		case 13:
			goto tr1063
		case 32:
			goto tr1089
		case 34:
			goto tr209
		case 44:
			goto tr139
		case 61:
			goto tr146
		case 92:
			goto st86
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st692
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1089
		}
		goto st84
	st692:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof692
		}
	st_case_692:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1063
		case 11:
			goto tr1090
		case 13:
			goto tr1063
		case 32:
			goto tr1089
		case 34:
			goto tr209
		case 44:
			goto tr139
		case 61:
			goto tr146
		case 92:
			goto st86
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st693
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1089
		}
		goto st84
	st693:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof693
		}
	st_case_693:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1063
		case 11:
			goto tr1090
		case 13:
			goto tr1063
		case 32:
			goto tr1089
		case 34:
			goto tr209
		case 44:
			goto tr139
		case 61:
			goto tr146
		case 92:
			goto st86
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st694
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1089
		}
		goto st84
	st694:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof694
		}
	st_case_694:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1063
		case 11:
			goto tr1090
		case 13:
			goto tr1063
		case 32:
			goto tr1089
		case 34:
			goto tr209
		case 44:
			goto tr139
		case 61:
			goto tr146
		case 92:
			goto st86
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st695
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1089
		}
		goto st84
	st695:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof695
		}
	st_case_695:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1063
		case 11:
			goto tr1090
		case 13:
			goto tr1063
		case 32:
			goto tr1089
		case 34:
			goto tr209
		case 44:
			goto tr139
		case 61:
			goto tr146
		case 92:
			goto st86
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st696
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1089
		}
		goto st84
	st696:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof696
		}
	st_case_696:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1063
		case 11:
			goto tr1090
		case 13:
			goto tr1063
		case 32:
			goto tr1089
		case 34:
			goto tr209
		case 44:
			goto tr139
		case 61:
			goto tr146
		case 92:
			goto st86
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st697
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1089
		}
		goto st84
	st697:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof697
		}
	st_case_697:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1063
		case 11:
			goto tr1090
		case 13:
			goto tr1063
		case 32:
			goto tr1089
		case 34:
			goto tr209
		case 44:
			goto tr139
		case 61:
			goto tr146
		case 92:
			goto st86
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1089
		}
		goto st84
tr1057:
//line machine.rl:66

	m.handler.AddTag(key, m.text())

//line machine.rl:17

	m.pb = m.p

	goto st88
tr1202:
//line machine.rl:58

	m.handler.SetMeasurement(m.text())

//line machine.rl:78

	m.handler.AddFloat(key, m.text())

	goto st88
tr1115:
//line machine.rl:66

	m.handler.AddTag(key, m.text())

//line machine.rl:78

	m.handler.AddFloat(key, m.text())

	goto st88
tr1121:
//line machine.rl:66

	m.handler.AddTag(key, m.text())

//line machine.rl:74

	m.handler.AddInt(key, m.text())

	goto st88
tr1125:
//line machine.rl:66

	m.handler.AddTag(key, m.text())

//line machine.rl:82

	m.handler.AddBool(key, m.text())

	goto st88
tr1233:
//line machine.rl:58

	m.handler.SetMeasurement(m.text())

//line machine.rl:74

	m.handler.AddInt(key, m.text())

	goto st88
tr1236:
//line machine.rl:58

	m.handler.SetMeasurement(m.text())

//line machine.rl:82

	m.handler.AddBool(key, m.text())

	goto st88
	st88:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof88
		}
	st_case_88:
//line machine.go:17027
		switch ( m.data)[( m.p)] {
		case 32:
			goto st6
		case 34:
			goto tr224
		case 44:
			goto st6
		case 61:
			goto st6
		case 92:
			goto tr249
		}
		switch {
		case ( m.data)[( m.p)] > 10:
			if 12 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto st6
			}
		case ( m.data)[( m.p)] >= 9:
			goto st6
		}
		goto tr248
tr248:
//line machine.rl:17

	m.pb = m.p

	goto st89
	st89:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof89
		}
	st_case_89:
//line machine.go:17060
		switch ( m.data)[( m.p)] {
		case 32:
			goto st6
		case 34:
			goto tr227
		case 44:
			goto st6
		case 61:
			goto tr251
		case 92:
			goto st101
		}
		switch {
		case ( m.data)[( m.p)] > 10:
			if 12 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto st6
			}
		case ( m.data)[( m.p)] >= 9:
			goto st6
		}
		goto st89
tr251:
//line machine.rl:62

	key = m.text()

//line machine.rl:70

	key = m.text()

	goto st90
	st90:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof90
		}
	st_case_90:
//line machine.go:17097
		switch ( m.data)[( m.p)] {
		case 32:
			goto st6
		case 34:
			goto tr231
		case 44:
			goto st6
		case 45:
			goto tr253
		case 46:
			goto tr254
		case 48:
			goto tr255
		case 61:
			goto st6
		case 70:
			goto tr257
		case 84:
			goto tr258
		case 92:
			goto tr134
		case 102:
			goto tr259
		case 116:
			goto tr260
		}
		switch {
		case ( m.data)[( m.p)] < 12:
			if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 10 {
				goto st6
			}
		case ( m.data)[( m.p)] > 13:
			if 49 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto tr256
			}
		default:
			goto st6
		}
		goto tr129
tr253:
//line machine.rl:17

	m.pb = m.p

	goto st91
	st91:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof91
		}
	st_case_91:
//line machine.go:17148
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr137
		case 13:
			goto st6
		case 32:
			goto tr136
		case 34:
			goto tr138
		case 44:
			goto tr139
		case 48:
			goto st698
		case 61:
			goto st6
		case 92:
			goto st55
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 49 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st701
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr136
		}
		goto st41
tr137:
//line machine.rl:66

	m.handler.AddTag(key, m.text())

	goto st92
tr131:
//line machine.rl:66

	m.handler.AddTag(key, m.text())

//line machine.rl:17

	m.pb = m.p

	goto st92
	st92:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof92
		}
	st_case_92:
//line machine.go:17199
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr245
		case 13:
			goto st6
		case 32:
			goto tr136
		case 34:
			goto tr205
		case 44:
			goto tr139
		case 61:
			goto st6
		case 92:
			goto tr246
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr136
		}
		goto tr244
tr255:
//line machine.rl:17

	m.pb = m.p

	goto st698
	st698:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof698
		}
	st_case_698:
//line machine.go:17233
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr922
		case 11:
			goto tr1114
		case 13:
			goto tr922
		case 32:
			goto tr1113
		case 34:
			goto tr138
		case 44:
			goto tr1115
		case 46:
			goto st699
		case 61:
			goto st6
		case 92:
			goto st55
		case 105:
			goto st700
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1113
		}
		goto st41
tr254:
//line machine.rl:17

	m.pb = m.p

	goto st699
	st699:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof699
		}
	st_case_699:
//line machine.go:17271
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr922
		case 11:
			goto tr1114
		case 13:
			goto tr922
		case 32:
			goto tr1113
		case 34:
			goto tr138
		case 44:
			goto tr1115
		case 61:
			goto st6
		case 92:
			goto st55
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st699
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1113
		}
		goto st41
	st700:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof700
		}
	st_case_700:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1119
		case 11:
			goto tr1120
		case 13:
			goto tr1119
		case 32:
			goto tr1118
		case 34:
			goto tr138
		case 44:
			goto tr1121
		case 61:
			goto st6
		case 92:
			goto st55
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1118
		}
		goto st41
tr256:
//line machine.rl:17

	m.pb = m.p

	goto st701
	st701:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof701
		}
	st_case_701:
//line machine.go:17337
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr922
		case 11:
			goto tr1114
		case 13:
			goto tr922
		case 32:
			goto tr1113
		case 34:
			goto tr138
		case 44:
			goto tr1115
		case 46:
			goto st699
		case 61:
			goto st6
		case 92:
			goto st55
		case 105:
			goto st700
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st701
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1113
		}
		goto st41
tr257:
//line machine.rl:17

	m.pb = m.p

	goto st702
	st702:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof702
		}
	st_case_702:
//line machine.go:17380
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1123
		case 11:
			goto tr1124
		case 13:
			goto tr1123
		case 32:
			goto tr1122
		case 34:
			goto tr138
		case 44:
			goto tr1125
		case 61:
			goto st6
		case 65:
			goto st93
		case 92:
			goto st55
		case 97:
			goto st96
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1122
		}
		goto st41
	st93:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof93
		}
	st_case_93:
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr137
		case 13:
			goto st6
		case 32:
			goto tr136
		case 34:
			goto tr138
		case 44:
			goto tr139
		case 61:
			goto st6
		case 76:
			goto st94
		case 92:
			goto st55
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr136
		}
		goto st41
	st94:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof94
		}
	st_case_94:
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr137
		case 13:
			goto st6
		case 32:
			goto tr136
		case 34:
			goto tr138
		case 44:
			goto tr139
		case 61:
			goto st6
		case 83:
			goto st95
		case 92:
			goto st55
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr136
		}
		goto st41
	st95:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof95
		}
	st_case_95:
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr137
		case 13:
			goto st6
		case 32:
			goto tr136
		case 34:
			goto tr138
		case 44:
			goto tr139
		case 61:
			goto st6
		case 69:
			goto st703
		case 92:
			goto st55
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr136
		}
		goto st41
	st703:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof703
		}
	st_case_703:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1123
		case 11:
			goto tr1124
		case 13:
			goto tr1123
		case 32:
			goto tr1122
		case 34:
			goto tr138
		case 44:
			goto tr1125
		case 61:
			goto st6
		case 92:
			goto st55
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1122
		}
		goto st41
	st96:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof96
		}
	st_case_96:
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr137
		case 13:
			goto st6
		case 32:
			goto tr136
		case 34:
			goto tr138
		case 44:
			goto tr139
		case 61:
			goto st6
		case 92:
			goto st55
		case 108:
			goto st97
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr136
		}
		goto st41
	st97:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof97
		}
	st_case_97:
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr137
		case 13:
			goto st6
		case 32:
			goto tr136
		case 34:
			goto tr138
		case 44:
			goto tr139
		case 61:
			goto st6
		case 92:
			goto st55
		case 115:
			goto st98
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr136
		}
		goto st41
	st98:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof98
		}
	st_case_98:
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr137
		case 13:
			goto st6
		case 32:
			goto tr136
		case 34:
			goto tr138
		case 44:
			goto tr139
		case 61:
			goto st6
		case 92:
			goto st55
		case 101:
			goto st703
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr136
		}
		goto st41
tr258:
//line machine.rl:17

	m.pb = m.p

	goto st704
	st704:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof704
		}
	st_case_704:
//line machine.go:17619
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1123
		case 11:
			goto tr1124
		case 13:
			goto tr1123
		case 32:
			goto tr1122
		case 34:
			goto tr138
		case 44:
			goto tr1125
		case 61:
			goto st6
		case 82:
			goto st99
		case 92:
			goto st55
		case 114:
			goto st100
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1122
		}
		goto st41
	st99:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof99
		}
	st_case_99:
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr137
		case 13:
			goto st6
		case 32:
			goto tr136
		case 34:
			goto tr138
		case 44:
			goto tr139
		case 61:
			goto st6
		case 85:
			goto st95
		case 92:
			goto st55
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr136
		}
		goto st41
	st100:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof100
		}
	st_case_100:
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr137
		case 13:
			goto st6
		case 32:
			goto tr136
		case 34:
			goto tr138
		case 44:
			goto tr139
		case 61:
			goto st6
		case 92:
			goto st55
		case 117:
			goto st98
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr136
		}
		goto st41
tr259:
//line machine.rl:17

	m.pb = m.p

	goto st705
	st705:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof705
		}
	st_case_705:
//line machine.go:17715
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1123
		case 11:
			goto tr1124
		case 13:
			goto tr1123
		case 32:
			goto tr1122
		case 34:
			goto tr138
		case 44:
			goto tr1125
		case 61:
			goto st6
		case 92:
			goto st55
		case 97:
			goto st96
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1122
		}
		goto st41
tr260:
//line machine.rl:17

	m.pb = m.p

	goto st706
	st706:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof706
		}
	st_case_706:
//line machine.go:17751
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1123
		case 11:
			goto tr1124
		case 13:
			goto tr1123
		case 32:
			goto tr1122
		case 34:
			goto tr138
		case 44:
			goto tr1125
		case 61:
			goto st6
		case 92:
			goto st55
		case 114:
			goto st100
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1122
		}
		goto st41
tr249:
//line machine.rl:17

	m.pb = m.p

	goto st101
	st101:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof101
		}
	st_case_101:
//line machine.go:17787
		switch ( m.data)[( m.p)] {
		case 34:
			goto st89
		case 92:
			goto st89
		}
		switch {
		case ( m.data)[( m.p)] > 10:
			if 12 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto tr48
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr48
		}
		goto st38
tr232:
//line machine.rl:17

	m.pb = m.p

	goto st102
	st102:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof102
		}
	st_case_102:
//line machine.go:17814
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr200
		case 13:
			goto st6
		case 32:
			goto tr199
		case 34:
			goto tr138
		case 44:
			goto tr201
		case 48:
			goto st707
		case 61:
			goto st6
		case 92:
			goto st69
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 49 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st732
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr199
		}
		goto st64
tr234:
//line machine.rl:17

	m.pb = m.p

	goto st707
	st707:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof707
		}
	st_case_707:
//line machine.go:17855
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1026
		case 11:
			goto tr1131
		case 13:
			goto tr1026
		case 32:
			goto tr1130
		case 34:
			goto tr138
		case 44:
			goto tr1132
		case 46:
			goto st730
		case 61:
			goto st6
		case 92:
			goto st69
		case 105:
			goto st731
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1130
		}
		goto st64
tr1131:
//line machine.rl:66

	m.handler.AddTag(key, m.text())

//line machine.rl:78

	m.handler.AddFloat(key, m.text())

	goto st708
tr1163:
//line machine.rl:66

	m.handler.AddTag(key, m.text())

//line machine.rl:74

	m.handler.AddInt(key, m.text())

	goto st708
tr1167:
//line machine.rl:66

	m.handler.AddTag(key, m.text())

//line machine.rl:82

	m.handler.AddBool(key, m.text())

	goto st708
	st708:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof708
		}
	st_case_708:
//line machine.go:17917
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr885
		case 11:
			goto tr1136
		case 13:
			goto tr885
		case 32:
			goto tr1135
		case 34:
			goto tr205
		case 44:
			goto tr201
		case 45:
			goto tr1137
		case 61:
			goto st6
		case 92:
			goto tr206
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto tr1138
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1135
		}
		goto tr203
tr1136:
//line machine.rl:66

	m.handler.AddTag(key, m.text())

//line machine.rl:17

	m.pb = m.p

	goto st709
	st709:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof709
		}
	st_case_709:
//line machine.go:17962
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr885
		case 11:
			goto tr1136
		case 13:
			goto tr885
		case 32:
			goto tr1135
		case 34:
			goto tr205
		case 44:
			goto tr201
		case 45:
			goto tr1137
		case 61:
			goto tr86
		case 92:
			goto tr206
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto tr1138
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1135
		}
		goto tr203
tr1137:
//line machine.rl:17

	m.pb = m.p

	goto st103
	st103:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof103
		}
	st_case_103:
//line machine.go:18003
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr208
		case 13:
			goto st6
		case 32:
			goto tr199
		case 34:
			goto tr209
		case 44:
			goto tr201
		case 61:
			goto tr86
		case 92:
			goto st68
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st710
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr199
		}
		goto st66
tr1138:
//line machine.rl:17

	m.pb = m.p

	goto st710
	st710:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof710
		}
	st_case_710:
//line machine.go:18042
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr990
		case 11:
			goto tr1140
		case 13:
			goto tr990
		case 32:
			goto tr1139
		case 34:
			goto tr209
		case 44:
			goto tr201
		case 61:
			goto tr86
		case 92:
			goto st68
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st712
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1139
		}
		goto st66
tr1143:
//line machine.rl:66

	m.handler.AddTag(key, m.text())

//line machine.rl:17

	m.pb = m.p

	goto st711
tr1140:
//line machine.rl:66

	m.handler.AddTag(key, m.text())

//line machine.rl:90

	m.handler.SetTimestamp(m.text())

	goto st711
tr1736:
//line machine.rl:90

	m.handler.SetTimestamp(m.text())

//line machine.rl:66

	m.handler.AddTag(key, m.text())

	goto st711
	st711:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof711
		}
	st_case_711:
//line machine.go:18105
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr885
		case 11:
			goto tr1143
		case 13:
			goto tr885
		case 32:
			goto tr1142
		case 34:
			goto tr205
		case 44:
			goto tr201
		case 61:
			goto tr86
		case 92:
			goto tr206
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1142
		}
		goto tr203
	st712:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof712
		}
	st_case_712:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr990
		case 11:
			goto tr1140
		case 13:
			goto tr990
		case 32:
			goto tr1139
		case 34:
			goto tr209
		case 44:
			goto tr201
		case 61:
			goto tr86
		case 92:
			goto st68
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st713
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1139
		}
		goto st66
	st713:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof713
		}
	st_case_713:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr990
		case 11:
			goto tr1140
		case 13:
			goto tr990
		case 32:
			goto tr1139
		case 34:
			goto tr209
		case 44:
			goto tr201
		case 61:
			goto tr86
		case 92:
			goto st68
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st714
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1139
		}
		goto st66
	st714:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof714
		}
	st_case_714:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr990
		case 11:
			goto tr1140
		case 13:
			goto tr990
		case 32:
			goto tr1139
		case 34:
			goto tr209
		case 44:
			goto tr201
		case 61:
			goto tr86
		case 92:
			goto st68
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st715
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1139
		}
		goto st66
	st715:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof715
		}
	st_case_715:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr990
		case 11:
			goto tr1140
		case 13:
			goto tr990
		case 32:
			goto tr1139
		case 34:
			goto tr209
		case 44:
			goto tr201
		case 61:
			goto tr86
		case 92:
			goto st68
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st716
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1139
		}
		goto st66
	st716:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof716
		}
	st_case_716:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr990
		case 11:
			goto tr1140
		case 13:
			goto tr990
		case 32:
			goto tr1139
		case 34:
			goto tr209
		case 44:
			goto tr201
		case 61:
			goto tr86
		case 92:
			goto st68
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st717
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1139
		}
		goto st66
	st717:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof717
		}
	st_case_717:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr990
		case 11:
			goto tr1140
		case 13:
			goto tr990
		case 32:
			goto tr1139
		case 34:
			goto tr209
		case 44:
			goto tr201
		case 61:
			goto tr86
		case 92:
			goto st68
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st718
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1139
		}
		goto st66
	st718:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof718
		}
	st_case_718:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr990
		case 11:
			goto tr1140
		case 13:
			goto tr990
		case 32:
			goto tr1139
		case 34:
			goto tr209
		case 44:
			goto tr201
		case 61:
			goto tr86
		case 92:
			goto st68
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st719
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1139
		}
		goto st66
	st719:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof719
		}
	st_case_719:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr990
		case 11:
			goto tr1140
		case 13:
			goto tr990
		case 32:
			goto tr1139
		case 34:
			goto tr209
		case 44:
			goto tr201
		case 61:
			goto tr86
		case 92:
			goto st68
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st720
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1139
		}
		goto st66
	st720:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof720
		}
	st_case_720:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr990
		case 11:
			goto tr1140
		case 13:
			goto tr990
		case 32:
			goto tr1139
		case 34:
			goto tr209
		case 44:
			goto tr201
		case 61:
			goto tr86
		case 92:
			goto st68
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st721
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1139
		}
		goto st66
	st721:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof721
		}
	st_case_721:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr990
		case 11:
			goto tr1140
		case 13:
			goto tr990
		case 32:
			goto tr1139
		case 34:
			goto tr209
		case 44:
			goto tr201
		case 61:
			goto tr86
		case 92:
			goto st68
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st722
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1139
		}
		goto st66
	st722:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof722
		}
	st_case_722:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr990
		case 11:
			goto tr1140
		case 13:
			goto tr990
		case 32:
			goto tr1139
		case 34:
			goto tr209
		case 44:
			goto tr201
		case 61:
			goto tr86
		case 92:
			goto st68
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st723
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1139
		}
		goto st66
	st723:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof723
		}
	st_case_723:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr990
		case 11:
			goto tr1140
		case 13:
			goto tr990
		case 32:
			goto tr1139
		case 34:
			goto tr209
		case 44:
			goto tr201
		case 61:
			goto tr86
		case 92:
			goto st68
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st724
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1139
		}
		goto st66
	st724:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof724
		}
	st_case_724:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr990
		case 11:
			goto tr1140
		case 13:
			goto tr990
		case 32:
			goto tr1139
		case 34:
			goto tr209
		case 44:
			goto tr201
		case 61:
			goto tr86
		case 92:
			goto st68
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st725
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1139
		}
		goto st66
	st725:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof725
		}
	st_case_725:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr990
		case 11:
			goto tr1140
		case 13:
			goto tr990
		case 32:
			goto tr1139
		case 34:
			goto tr209
		case 44:
			goto tr201
		case 61:
			goto tr86
		case 92:
			goto st68
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st726
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1139
		}
		goto st66
	st726:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof726
		}
	st_case_726:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr990
		case 11:
			goto tr1140
		case 13:
			goto tr990
		case 32:
			goto tr1139
		case 34:
			goto tr209
		case 44:
			goto tr201
		case 61:
			goto tr86
		case 92:
			goto st68
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st727
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1139
		}
		goto st66
	st727:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof727
		}
	st_case_727:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr990
		case 11:
			goto tr1140
		case 13:
			goto tr990
		case 32:
			goto tr1139
		case 34:
			goto tr209
		case 44:
			goto tr201
		case 61:
			goto tr86
		case 92:
			goto st68
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st728
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1139
		}
		goto st66
	st728:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof728
		}
	st_case_728:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr990
		case 11:
			goto tr1140
		case 13:
			goto tr990
		case 32:
			goto tr1139
		case 34:
			goto tr209
		case 44:
			goto tr201
		case 61:
			goto tr86
		case 92:
			goto st68
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st729
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1139
		}
		goto st66
	st729:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof729
		}
	st_case_729:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr990
		case 11:
			goto tr1140
		case 13:
			goto tr990
		case 32:
			goto tr1139
		case 34:
			goto tr209
		case 44:
			goto tr201
		case 61:
			goto tr86
		case 92:
			goto st68
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1139
		}
		goto st66
tr233:
//line machine.rl:17

	m.pb = m.p

	goto st730
	st730:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof730
		}
	st_case_730:
//line machine.go:18710
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1026
		case 11:
			goto tr1131
		case 13:
			goto tr1026
		case 32:
			goto tr1130
		case 34:
			goto tr138
		case 44:
			goto tr1132
		case 61:
			goto st6
		case 92:
			goto st69
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st730
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1130
		}
		goto st64
	st731:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof731
		}
	st_case_731:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1162
		case 11:
			goto tr1163
		case 13:
			goto tr1162
		case 32:
			goto tr1161
		case 34:
			goto tr138
		case 44:
			goto tr1164
		case 61:
			goto st6
		case 92:
			goto st69
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1161
		}
		goto st64
tr235:
//line machine.rl:17

	m.pb = m.p

	goto st732
	st732:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof732
		}
	st_case_732:
//line machine.go:18776
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1026
		case 11:
			goto tr1131
		case 13:
			goto tr1026
		case 32:
			goto tr1130
		case 34:
			goto tr138
		case 44:
			goto tr1132
		case 46:
			goto st730
		case 61:
			goto st6
		case 92:
			goto st69
		case 105:
			goto st731
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st732
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1130
		}
		goto st64
tr236:
//line machine.rl:17

	m.pb = m.p

	goto st733
	st733:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof733
		}
	st_case_733:
//line machine.go:18819
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1166
		case 11:
			goto tr1167
		case 13:
			goto tr1166
		case 32:
			goto tr1165
		case 34:
			goto tr138
		case 44:
			goto tr1168
		case 61:
			goto st6
		case 65:
			goto st104
		case 92:
			goto st69
		case 97:
			goto st107
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1165
		}
		goto st64
	st104:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof104
		}
	st_case_104:
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr200
		case 13:
			goto st6
		case 32:
			goto tr199
		case 34:
			goto tr138
		case 44:
			goto tr201
		case 61:
			goto st6
		case 76:
			goto st105
		case 92:
			goto st69
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr199
		}
		goto st64
	st105:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof105
		}
	st_case_105:
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr200
		case 13:
			goto st6
		case 32:
			goto tr199
		case 34:
			goto tr138
		case 44:
			goto tr201
		case 61:
			goto st6
		case 83:
			goto st106
		case 92:
			goto st69
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr199
		}
		goto st64
	st106:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof106
		}
	st_case_106:
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr200
		case 13:
			goto st6
		case 32:
			goto tr199
		case 34:
			goto tr138
		case 44:
			goto tr201
		case 61:
			goto st6
		case 69:
			goto st734
		case 92:
			goto st69
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr199
		}
		goto st64
	st734:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof734
		}
	st_case_734:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1166
		case 11:
			goto tr1167
		case 13:
			goto tr1166
		case 32:
			goto tr1165
		case 34:
			goto tr138
		case 44:
			goto tr1168
		case 61:
			goto st6
		case 92:
			goto st69
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1165
		}
		goto st64
	st107:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof107
		}
	st_case_107:
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr200
		case 13:
			goto st6
		case 32:
			goto tr199
		case 34:
			goto tr138
		case 44:
			goto tr201
		case 61:
			goto st6
		case 92:
			goto st69
		case 108:
			goto st108
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr199
		}
		goto st64
	st108:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof108
		}
	st_case_108:
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr200
		case 13:
			goto st6
		case 32:
			goto tr199
		case 34:
			goto tr138
		case 44:
			goto tr201
		case 61:
			goto st6
		case 92:
			goto st69
		case 115:
			goto st109
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr199
		}
		goto st64
	st109:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof109
		}
	st_case_109:
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr200
		case 13:
			goto st6
		case 32:
			goto tr199
		case 34:
			goto tr138
		case 44:
			goto tr201
		case 61:
			goto st6
		case 92:
			goto st69
		case 101:
			goto st734
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr199
		}
		goto st64
tr237:
//line machine.rl:17

	m.pb = m.p

	goto st735
	st735:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof735
		}
	st_case_735:
//line machine.go:19058
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1166
		case 11:
			goto tr1167
		case 13:
			goto tr1166
		case 32:
			goto tr1165
		case 34:
			goto tr138
		case 44:
			goto tr1168
		case 61:
			goto st6
		case 82:
			goto st110
		case 92:
			goto st69
		case 114:
			goto st111
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1165
		}
		goto st64
	st110:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof110
		}
	st_case_110:
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr200
		case 13:
			goto st6
		case 32:
			goto tr199
		case 34:
			goto tr138
		case 44:
			goto tr201
		case 61:
			goto st6
		case 85:
			goto st106
		case 92:
			goto st69
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr199
		}
		goto st64
	st111:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof111
		}
	st_case_111:
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr200
		case 13:
			goto st6
		case 32:
			goto tr199
		case 34:
			goto tr138
		case 44:
			goto tr201
		case 61:
			goto st6
		case 92:
			goto st69
		case 117:
			goto st109
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr199
		}
		goto st64
tr238:
//line machine.rl:17

	m.pb = m.p

	goto st736
	st736:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof736
		}
	st_case_736:
//line machine.go:19154
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1166
		case 11:
			goto tr1167
		case 13:
			goto tr1166
		case 32:
			goto tr1165
		case 34:
			goto tr138
		case 44:
			goto tr1168
		case 61:
			goto st6
		case 92:
			goto st69
		case 97:
			goto st107
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1165
		}
		goto st64
tr239:
//line machine.rl:17

	m.pb = m.p

	goto st737
	st737:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof737
		}
	st_case_737:
//line machine.go:19190
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1166
		case 11:
			goto tr1167
		case 13:
			goto tr1166
		case 32:
			goto tr1165
		case 34:
			goto tr138
		case 44:
			goto tr1168
		case 61:
			goto st6
		case 92:
			goto st69
		case 114:
			goto st111
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1165
		}
		goto st64
tr225:
//line machine.rl:17

	m.pb = m.p

	goto st112
	st112:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof112
		}
	st_case_112:
//line machine.go:19226
		switch ( m.data)[( m.p)] {
		case 34:
			goto st78
		case 92:
			goto st78
		}
		switch {
		case ( m.data)[( m.p)] > 10:
			if 12 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto tr48
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr48
		}
		goto st38
tr213:
//line machine.rl:17

	m.pb = m.p

	goto st738
	st738:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof738
		}
	st_case_738:
//line machine.go:19253
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1026
		case 11:
			goto tr1027
		case 13:
			goto tr1026
		case 32:
			goto tr1025
		case 34:
			goto tr76
		case 44:
			goto tr1028
		case 92:
			goto st113
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st738
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1025
		}
		goto st23
tr72:
//line machine.rl:17

	m.pb = m.p

	goto st113
	st113:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof113
		}
	st_case_113:
//line machine.go:19290
		switch ( m.data)[( m.p)] {
		case 34:
			goto st23
		case 92:
			goto st6
		}
		switch {
		case ( m.data)[( m.p)] > 10:
			if 12 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto tr2
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr2
		}
		goto st7
	st739:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof739
		}
	st_case_739:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1162
		case 11:
			goto tr1174
		case 13:
			goto tr1162
		case 32:
			goto tr1173
		case 34:
			goto tr76
		case 44:
			goto tr1175
		case 92:
			goto st113
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1173
		}
		goto st23
tr215:
//line machine.rl:17

	m.pb = m.p

	goto st740
	st740:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof740
		}
	st_case_740:
//line machine.go:19342
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1026
		case 11:
			goto tr1027
		case 13:
			goto tr1026
		case 32:
			goto tr1025
		case 34:
			goto tr76
		case 44:
			goto tr1028
		case 46:
			goto st738
		case 92:
			goto st113
		case 105:
			goto st739
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st740
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1025
		}
		goto st23
tr216:
//line machine.rl:17

	m.pb = m.p

	goto st741
	st741:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof741
		}
	st_case_741:
//line machine.go:19383
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1166
		case 11:
			goto tr1177
		case 13:
			goto tr1166
		case 32:
			goto tr1176
		case 34:
			goto tr76
		case 44:
			goto tr1178
		case 65:
			goto st114
		case 92:
			goto st113
		case 97:
			goto st117
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1176
		}
		goto st23
	st114:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof114
		}
	st_case_114:
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr75
		case 13:
			goto st6
		case 32:
			goto tr74
		case 34:
			goto tr76
		case 44:
			goto tr77
		case 76:
			goto st115
		case 92:
			goto st113
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr74
		}
		goto st23
	st115:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof115
		}
	st_case_115:
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr75
		case 13:
			goto st6
		case 32:
			goto tr74
		case 34:
			goto tr76
		case 44:
			goto tr77
		case 83:
			goto st116
		case 92:
			goto st113
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr74
		}
		goto st23
	st116:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof116
		}
	st_case_116:
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr75
		case 13:
			goto st6
		case 32:
			goto tr74
		case 34:
			goto tr76
		case 44:
			goto tr77
		case 69:
			goto st742
		case 92:
			goto st113
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr74
		}
		goto st23
	st742:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof742
		}
	st_case_742:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1166
		case 11:
			goto tr1177
		case 13:
			goto tr1166
		case 32:
			goto tr1176
		case 34:
			goto tr76
		case 44:
			goto tr1178
		case 92:
			goto st113
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1176
		}
		goto st23
	st117:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof117
		}
	st_case_117:
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr75
		case 13:
			goto st6
		case 32:
			goto tr74
		case 34:
			goto tr76
		case 44:
			goto tr77
		case 92:
			goto st113
		case 108:
			goto st118
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr74
		}
		goto st23
	st118:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof118
		}
	st_case_118:
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr75
		case 13:
			goto st6
		case 32:
			goto tr74
		case 34:
			goto tr76
		case 44:
			goto tr77
		case 92:
			goto st113
		case 115:
			goto st119
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr74
		}
		goto st23
	st119:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof119
		}
	st_case_119:
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr75
		case 13:
			goto st6
		case 32:
			goto tr74
		case 34:
			goto tr76
		case 44:
			goto tr77
		case 92:
			goto st113
		case 101:
			goto st742
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr74
		}
		goto st23
tr217:
//line machine.rl:17

	m.pb = m.p

	goto st743
	st743:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof743
		}
	st_case_743:
//line machine.go:19606
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1166
		case 11:
			goto tr1177
		case 13:
			goto tr1166
		case 32:
			goto tr1176
		case 34:
			goto tr76
		case 44:
			goto tr1178
		case 82:
			goto st120
		case 92:
			goto st113
		case 114:
			goto st121
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1176
		}
		goto st23
	st120:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof120
		}
	st_case_120:
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr75
		case 13:
			goto st6
		case 32:
			goto tr74
		case 34:
			goto tr76
		case 44:
			goto tr77
		case 85:
			goto st116
		case 92:
			goto st113
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr74
		}
		goto st23
	st121:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof121
		}
	st_case_121:
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr75
		case 13:
			goto st6
		case 32:
			goto tr74
		case 34:
			goto tr76
		case 44:
			goto tr77
		case 92:
			goto st113
		case 117:
			goto st119
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr74
		}
		goto st23
tr218:
//line machine.rl:17

	m.pb = m.p

	goto st744
	st744:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof744
		}
	st_case_744:
//line machine.go:19696
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1166
		case 11:
			goto tr1177
		case 13:
			goto tr1166
		case 32:
			goto tr1176
		case 34:
			goto tr76
		case 44:
			goto tr1178
		case 92:
			goto st113
		case 97:
			goto st117
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1176
		}
		goto st23
tr219:
//line machine.rl:17

	m.pb = m.p

	goto st745
	st745:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof745
		}
	st_case_745:
//line machine.go:19730
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1166
		case 11:
			goto tr1177
		case 13:
			goto tr1166
		case 32:
			goto tr1176
		case 34:
			goto tr76
		case 44:
			goto tr1178
		case 92:
			goto st113
		case 114:
			goto st121
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1176
		}
		goto st23
	st746:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof746
		}
	st_case_746:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr990
		case 11:
			goto tr1018
		case 13:
			goto tr990
		case 32:
			goto tr1017
		case 34:
			goto tr110
		case 44:
			goto tr77
		case 61:
			goto tr111
		case 92:
			goto st74
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st747
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1017
		}
		goto st33
	st747:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof747
		}
	st_case_747:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr990
		case 11:
			goto tr1018
		case 13:
			goto tr990
		case 32:
			goto tr1017
		case 34:
			goto tr110
		case 44:
			goto tr77
		case 61:
			goto tr111
		case 92:
			goto st74
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st748
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1017
		}
		goto st33
	st748:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof748
		}
	st_case_748:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr990
		case 11:
			goto tr1018
		case 13:
			goto tr990
		case 32:
			goto tr1017
		case 34:
			goto tr110
		case 44:
			goto tr77
		case 61:
			goto tr111
		case 92:
			goto st74
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st749
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1017
		}
		goto st33
	st749:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof749
		}
	st_case_749:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr990
		case 11:
			goto tr1018
		case 13:
			goto tr990
		case 32:
			goto tr1017
		case 34:
			goto tr110
		case 44:
			goto tr77
		case 61:
			goto tr111
		case 92:
			goto st74
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st750
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1017
		}
		goto st33
	st750:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof750
		}
	st_case_750:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr990
		case 11:
			goto tr1018
		case 13:
			goto tr990
		case 32:
			goto tr1017
		case 34:
			goto tr110
		case 44:
			goto tr77
		case 61:
			goto tr111
		case 92:
			goto st74
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st751
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1017
		}
		goto st33
	st751:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof751
		}
	st_case_751:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr990
		case 11:
			goto tr1018
		case 13:
			goto tr990
		case 32:
			goto tr1017
		case 34:
			goto tr110
		case 44:
			goto tr77
		case 61:
			goto tr111
		case 92:
			goto st74
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st752
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1017
		}
		goto st33
	st752:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof752
		}
	st_case_752:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr990
		case 11:
			goto tr1018
		case 13:
			goto tr990
		case 32:
			goto tr1017
		case 34:
			goto tr110
		case 44:
			goto tr77
		case 61:
			goto tr111
		case 92:
			goto st74
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st753
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1017
		}
		goto st33
	st753:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof753
		}
	st_case_753:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr990
		case 11:
			goto tr1018
		case 13:
			goto tr990
		case 32:
			goto tr1017
		case 34:
			goto tr110
		case 44:
			goto tr77
		case 61:
			goto tr111
		case 92:
			goto st74
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st754
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1017
		}
		goto st33
	st754:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof754
		}
	st_case_754:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr990
		case 11:
			goto tr1018
		case 13:
			goto tr990
		case 32:
			goto tr1017
		case 34:
			goto tr110
		case 44:
			goto tr77
		case 61:
			goto tr111
		case 92:
			goto st74
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st755
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1017
		}
		goto st33
	st755:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof755
		}
	st_case_755:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr990
		case 11:
			goto tr1018
		case 13:
			goto tr990
		case 32:
			goto tr1017
		case 34:
			goto tr110
		case 44:
			goto tr77
		case 61:
			goto tr111
		case 92:
			goto st74
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st756
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1017
		}
		goto st33
	st756:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof756
		}
	st_case_756:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr990
		case 11:
			goto tr1018
		case 13:
			goto tr990
		case 32:
			goto tr1017
		case 34:
			goto tr110
		case 44:
			goto tr77
		case 61:
			goto tr111
		case 92:
			goto st74
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st757
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1017
		}
		goto st33
	st757:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof757
		}
	st_case_757:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr990
		case 11:
			goto tr1018
		case 13:
			goto tr990
		case 32:
			goto tr1017
		case 34:
			goto tr110
		case 44:
			goto tr77
		case 61:
			goto tr111
		case 92:
			goto st74
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st758
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1017
		}
		goto st33
	st758:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof758
		}
	st_case_758:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr990
		case 11:
			goto tr1018
		case 13:
			goto tr990
		case 32:
			goto tr1017
		case 34:
			goto tr110
		case 44:
			goto tr77
		case 61:
			goto tr111
		case 92:
			goto st74
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st759
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1017
		}
		goto st33
	st759:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof759
		}
	st_case_759:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr990
		case 11:
			goto tr1018
		case 13:
			goto tr990
		case 32:
			goto tr1017
		case 34:
			goto tr110
		case 44:
			goto tr77
		case 61:
			goto tr111
		case 92:
			goto st74
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st760
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1017
		}
		goto st33
	st760:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof760
		}
	st_case_760:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr990
		case 11:
			goto tr1018
		case 13:
			goto tr990
		case 32:
			goto tr1017
		case 34:
			goto tr110
		case 44:
			goto tr77
		case 61:
			goto tr111
		case 92:
			goto st74
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st761
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1017
		}
		goto st33
	st761:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof761
		}
	st_case_761:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr990
		case 11:
			goto tr1018
		case 13:
			goto tr990
		case 32:
			goto tr1017
		case 34:
			goto tr110
		case 44:
			goto tr77
		case 61:
			goto tr111
		case 92:
			goto st74
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st762
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1017
		}
		goto st33
	st762:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof762
		}
	st_case_762:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr990
		case 11:
			goto tr1018
		case 13:
			goto tr990
		case 32:
			goto tr1017
		case 34:
			goto tr110
		case 44:
			goto tr77
		case 61:
			goto tr111
		case 92:
			goto st74
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st763
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1017
		}
		goto st33
	st763:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof763
		}
	st_case_763:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr990
		case 11:
			goto tr1018
		case 13:
			goto tr990
		case 32:
			goto tr1017
		case 34:
			goto tr110
		case 44:
			goto tr77
		case 61:
			goto tr111
		case 92:
			goto st74
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1017
		}
		goto st33
tr182:
//line machine.rl:17

	m.pb = m.p

	goto st122
	st122:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof122
		}
	st_case_122:
//line machine.go:20335
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr160
		case 13:
			goto st6
		case 32:
			goto tr159
		case 34:
			goto tr76
		case 44:
			goto tr161
		case 48:
			goto st764
		case 92:
			goto st125
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 49 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st789
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr159
		}
		goto st46
tr184:
//line machine.rl:17

	m.pb = m.p

	goto st764
	st764:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof764
		}
	st_case_764:
//line machine.go:20374
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr922
		case 11:
			goto tr1201
		case 13:
			goto tr922
		case 32:
			goto tr1200
		case 34:
			goto tr76
		case 44:
			goto tr1202
		case 46:
			goto st787
		case 92:
			goto st125
		case 105:
			goto st788
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1200
		}
		goto st46
tr1201:
//line machine.rl:58

	m.handler.SetMeasurement(m.text())

//line machine.rl:78

	m.handler.AddFloat(key, m.text())

	goto st765
tr1232:
//line machine.rl:58

	m.handler.SetMeasurement(m.text())

//line machine.rl:74

	m.handler.AddInt(key, m.text())

	goto st765
tr1235:
//line machine.rl:58

	m.handler.SetMeasurement(m.text())

//line machine.rl:82

	m.handler.AddBool(key, m.text())

	goto st765
	st765:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof765
		}
	st_case_765:
//line machine.go:20434
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr927
		case 11:
			goto tr1206
		case 13:
			goto tr927
		case 32:
			goto tr1205
		case 34:
			goto tr106
		case 44:
			goto tr161
		case 45:
			goto tr1207
		case 61:
			goto st46
		case 92:
			goto tr165
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto tr1208
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1205
		}
		goto tr163
tr1206:
//line machine.rl:58

	m.handler.SetMeasurement(m.text())

//line machine.rl:17

	m.pb = m.p

	goto st766
	st766:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof766
		}
	st_case_766:
//line machine.go:20479
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr927
		case 11:
			goto tr1206
		case 13:
			goto tr927
		case 32:
			goto tr1205
		case 34:
			goto tr106
		case 44:
			goto tr161
		case 45:
			goto tr1207
		case 61:
			goto tr168
		case 92:
			goto tr165
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto tr1208
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1205
		}
		goto tr163
tr1207:
//line machine.rl:17

	m.pb = m.p

	goto st123
	st123:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof123
		}
	st_case_123:
//line machine.go:20520
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr167
		case 13:
			goto st6
		case 32:
			goto tr159
		case 34:
			goto tr110
		case 44:
			goto tr161
		case 61:
			goto tr168
		case 92:
			goto st124
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st767
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr159
		}
		goto st48
tr1208:
//line machine.rl:17

	m.pb = m.p

	goto st767
	st767:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof767
		}
	st_case_767:
//line machine.go:20559
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1063
		case 11:
			goto tr1210
		case 13:
			goto tr1063
		case 32:
			goto tr1209
		case 34:
			goto tr110
		case 44:
			goto tr161
		case 61:
			goto tr168
		case 92:
			goto st124
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st769
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1209
		}
		goto st48
tr1213:
//line machine.rl:58

	m.handler.SetMeasurement(m.text())

//line machine.rl:17

	m.pb = m.p

	goto st768
tr1210:
//line machine.rl:58

	m.handler.SetMeasurement(m.text())

//line machine.rl:90

	m.handler.SetTimestamp(m.text())

	goto st768
	st768:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof768
		}
	st_case_768:
//line machine.go:20612
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr927
		case 11:
			goto tr1213
		case 13:
			goto tr927
		case 32:
			goto tr1212
		case 34:
			goto tr106
		case 44:
			goto tr161
		case 61:
			goto tr168
		case 92:
			goto tr165
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1212
		}
		goto tr163
tr165:
//line machine.rl:17

	m.pb = m.p

	goto st124
	st124:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof124
		}
	st_case_124:
//line machine.go:20646
		switch ( m.data)[( m.p)] {
		case 34:
			goto st48
		case 92:
			goto st43
		}
		switch {
		case ( m.data)[( m.p)] > 10:
			if 12 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto tr2
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr2
		}
		goto st9
	st769:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof769
		}
	st_case_769:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1063
		case 11:
			goto tr1210
		case 13:
			goto tr1063
		case 32:
			goto tr1209
		case 34:
			goto tr110
		case 44:
			goto tr161
		case 61:
			goto tr168
		case 92:
			goto st124
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st770
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1209
		}
		goto st48
	st770:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof770
		}
	st_case_770:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1063
		case 11:
			goto tr1210
		case 13:
			goto tr1063
		case 32:
			goto tr1209
		case 34:
			goto tr110
		case 44:
			goto tr161
		case 61:
			goto tr168
		case 92:
			goto st124
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st771
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1209
		}
		goto st48
	st771:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof771
		}
	st_case_771:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1063
		case 11:
			goto tr1210
		case 13:
			goto tr1063
		case 32:
			goto tr1209
		case 34:
			goto tr110
		case 44:
			goto tr161
		case 61:
			goto tr168
		case 92:
			goto st124
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st772
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1209
		}
		goto st48
	st772:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof772
		}
	st_case_772:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1063
		case 11:
			goto tr1210
		case 13:
			goto tr1063
		case 32:
			goto tr1209
		case 34:
			goto tr110
		case 44:
			goto tr161
		case 61:
			goto tr168
		case 92:
			goto st124
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st773
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1209
		}
		goto st48
	st773:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof773
		}
	st_case_773:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1063
		case 11:
			goto tr1210
		case 13:
			goto tr1063
		case 32:
			goto tr1209
		case 34:
			goto tr110
		case 44:
			goto tr161
		case 61:
			goto tr168
		case 92:
			goto st124
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st774
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1209
		}
		goto st48
	st774:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof774
		}
	st_case_774:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1063
		case 11:
			goto tr1210
		case 13:
			goto tr1063
		case 32:
			goto tr1209
		case 34:
			goto tr110
		case 44:
			goto tr161
		case 61:
			goto tr168
		case 92:
			goto st124
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st775
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1209
		}
		goto st48
	st775:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof775
		}
	st_case_775:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1063
		case 11:
			goto tr1210
		case 13:
			goto tr1063
		case 32:
			goto tr1209
		case 34:
			goto tr110
		case 44:
			goto tr161
		case 61:
			goto tr168
		case 92:
			goto st124
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st776
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1209
		}
		goto st48
	st776:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof776
		}
	st_case_776:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1063
		case 11:
			goto tr1210
		case 13:
			goto tr1063
		case 32:
			goto tr1209
		case 34:
			goto tr110
		case 44:
			goto tr161
		case 61:
			goto tr168
		case 92:
			goto st124
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st777
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1209
		}
		goto st48
	st777:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof777
		}
	st_case_777:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1063
		case 11:
			goto tr1210
		case 13:
			goto tr1063
		case 32:
			goto tr1209
		case 34:
			goto tr110
		case 44:
			goto tr161
		case 61:
			goto tr168
		case 92:
			goto st124
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st778
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1209
		}
		goto st48
	st778:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof778
		}
	st_case_778:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1063
		case 11:
			goto tr1210
		case 13:
			goto tr1063
		case 32:
			goto tr1209
		case 34:
			goto tr110
		case 44:
			goto tr161
		case 61:
			goto tr168
		case 92:
			goto st124
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st779
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1209
		}
		goto st48
	st779:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof779
		}
	st_case_779:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1063
		case 11:
			goto tr1210
		case 13:
			goto tr1063
		case 32:
			goto tr1209
		case 34:
			goto tr110
		case 44:
			goto tr161
		case 61:
			goto tr168
		case 92:
			goto st124
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st780
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1209
		}
		goto st48
	st780:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof780
		}
	st_case_780:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1063
		case 11:
			goto tr1210
		case 13:
			goto tr1063
		case 32:
			goto tr1209
		case 34:
			goto tr110
		case 44:
			goto tr161
		case 61:
			goto tr168
		case 92:
			goto st124
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st781
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1209
		}
		goto st48
	st781:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof781
		}
	st_case_781:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1063
		case 11:
			goto tr1210
		case 13:
			goto tr1063
		case 32:
			goto tr1209
		case 34:
			goto tr110
		case 44:
			goto tr161
		case 61:
			goto tr168
		case 92:
			goto st124
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st782
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1209
		}
		goto st48
	st782:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof782
		}
	st_case_782:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1063
		case 11:
			goto tr1210
		case 13:
			goto tr1063
		case 32:
			goto tr1209
		case 34:
			goto tr110
		case 44:
			goto tr161
		case 61:
			goto tr168
		case 92:
			goto st124
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st783
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1209
		}
		goto st48
	st783:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof783
		}
	st_case_783:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1063
		case 11:
			goto tr1210
		case 13:
			goto tr1063
		case 32:
			goto tr1209
		case 34:
			goto tr110
		case 44:
			goto tr161
		case 61:
			goto tr168
		case 92:
			goto st124
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st784
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1209
		}
		goto st48
	st784:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof784
		}
	st_case_784:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1063
		case 11:
			goto tr1210
		case 13:
			goto tr1063
		case 32:
			goto tr1209
		case 34:
			goto tr110
		case 44:
			goto tr161
		case 61:
			goto tr168
		case 92:
			goto st124
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st785
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1209
		}
		goto st48
	st785:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof785
		}
	st_case_785:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1063
		case 11:
			goto tr1210
		case 13:
			goto tr1063
		case 32:
			goto tr1209
		case 34:
			goto tr110
		case 44:
			goto tr161
		case 61:
			goto tr168
		case 92:
			goto st124
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st786
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1209
		}
		goto st48
	st786:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof786
		}
	st_case_786:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1063
		case 11:
			goto tr1210
		case 13:
			goto tr1063
		case 32:
			goto tr1209
		case 34:
			goto tr110
		case 44:
			goto tr161
		case 61:
			goto tr168
		case 92:
			goto st124
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1209
		}
		goto st48
tr183:
//line machine.rl:17

	m.pb = m.p

	goto st787
	st787:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof787
		}
	st_case_787:
//line machine.go:21244
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr922
		case 11:
			goto tr1201
		case 13:
			goto tr922
		case 32:
			goto tr1200
		case 34:
			goto tr76
		case 44:
			goto tr1202
		case 92:
			goto st125
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st787
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1200
		}
		goto st46
tr292:
//line machine.rl:17

	m.pb = m.p

	goto st125
	st125:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof125
		}
	st_case_125:
//line machine.go:21281
		switch ( m.data)[( m.p)] {
		case 34:
			goto st46
		case 92:
			goto st6
		}
		switch {
		case ( m.data)[( m.p)] > 10:
			if 12 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto tr2
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr2
		}
		goto st7
	st788:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof788
		}
	st_case_788:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1119
		case 11:
			goto tr1232
		case 13:
			goto tr1119
		case 32:
			goto tr1231
		case 34:
			goto tr76
		case 44:
			goto tr1233
		case 92:
			goto st125
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1231
		}
		goto st46
tr185:
//line machine.rl:17

	m.pb = m.p

	goto st789
	st789:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof789
		}
	st_case_789:
//line machine.go:21333
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr922
		case 11:
			goto tr1201
		case 13:
			goto tr922
		case 32:
			goto tr1200
		case 34:
			goto tr76
		case 44:
			goto tr1202
		case 46:
			goto st787
		case 92:
			goto st125
		case 105:
			goto st788
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st789
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1200
		}
		goto st46
tr186:
//line machine.rl:17

	m.pb = m.p

	goto st790
	st790:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof790
		}
	st_case_790:
//line machine.go:21374
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1123
		case 11:
			goto tr1235
		case 13:
			goto tr1123
		case 32:
			goto tr1234
		case 34:
			goto tr76
		case 44:
			goto tr1236
		case 65:
			goto st126
		case 92:
			goto st125
		case 97:
			goto st129
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1234
		}
		goto st46
	st126:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof126
		}
	st_case_126:
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr160
		case 13:
			goto st6
		case 32:
			goto tr159
		case 34:
			goto tr76
		case 44:
			goto tr161
		case 76:
			goto st127
		case 92:
			goto st125
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr159
		}
		goto st46
	st127:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof127
		}
	st_case_127:
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr160
		case 13:
			goto st6
		case 32:
			goto tr159
		case 34:
			goto tr76
		case 44:
			goto tr161
		case 83:
			goto st128
		case 92:
			goto st125
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr159
		}
		goto st46
	st128:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof128
		}
	st_case_128:
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr160
		case 13:
			goto st6
		case 32:
			goto tr159
		case 34:
			goto tr76
		case 44:
			goto tr161
		case 69:
			goto st791
		case 92:
			goto st125
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr159
		}
		goto st46
	st791:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof791
		}
	st_case_791:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1123
		case 11:
			goto tr1235
		case 13:
			goto tr1123
		case 32:
			goto tr1234
		case 34:
			goto tr76
		case 44:
			goto tr1236
		case 92:
			goto st125
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1234
		}
		goto st46
	st129:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof129
		}
	st_case_129:
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr160
		case 13:
			goto st6
		case 32:
			goto tr159
		case 34:
			goto tr76
		case 44:
			goto tr161
		case 92:
			goto st125
		case 108:
			goto st130
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr159
		}
		goto st46
	st130:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof130
		}
	st_case_130:
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr160
		case 13:
			goto st6
		case 32:
			goto tr159
		case 34:
			goto tr76
		case 44:
			goto tr161
		case 92:
			goto st125
		case 115:
			goto st131
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr159
		}
		goto st46
	st131:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof131
		}
	st_case_131:
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr160
		case 13:
			goto st6
		case 32:
			goto tr159
		case 34:
			goto tr76
		case 44:
			goto tr161
		case 92:
			goto st125
		case 101:
			goto st791
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr159
		}
		goto st46
tr187:
//line machine.rl:17

	m.pb = m.p

	goto st792
	st792:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof792
		}
	st_case_792:
//line machine.go:21597
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1123
		case 11:
			goto tr1235
		case 13:
			goto tr1123
		case 32:
			goto tr1234
		case 34:
			goto tr76
		case 44:
			goto tr1236
		case 82:
			goto st132
		case 92:
			goto st125
		case 114:
			goto st133
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1234
		}
		goto st46
	st132:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof132
		}
	st_case_132:
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr160
		case 13:
			goto st6
		case 32:
			goto tr159
		case 34:
			goto tr76
		case 44:
			goto tr161
		case 85:
			goto st128
		case 92:
			goto st125
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr159
		}
		goto st46
	st133:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof133
		}
	st_case_133:
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr160
		case 13:
			goto st6
		case 32:
			goto tr159
		case 34:
			goto tr76
		case 44:
			goto tr161
		case 92:
			goto st125
		case 117:
			goto st131
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr159
		}
		goto st46
tr188:
//line machine.rl:17

	m.pb = m.p

	goto st793
	st793:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof793
		}
	st_case_793:
//line machine.go:21687
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1123
		case 11:
			goto tr1235
		case 13:
			goto tr1123
		case 32:
			goto tr1234
		case 34:
			goto tr76
		case 44:
			goto tr1236
		case 92:
			goto st125
		case 97:
			goto st129
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1234
		}
		goto st46
tr189:
//line machine.rl:17

	m.pb = m.p

	goto st794
	st794:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof794
		}
	st_case_794:
//line machine.go:21721
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1123
		case 11:
			goto tr1235
		case 13:
			goto tr1123
		case 32:
			goto tr1234
		case 34:
			goto tr76
		case 44:
			goto tr1236
		case 92:
			goto st125
		case 114:
			goto st133
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1234
		}
		goto st46
	st134:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof134
		}
	st_case_134:
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr291
		case 13:
			goto st6
		case 32:
			goto st134
		case 34:
			goto tr70
		case 44:
			goto st6
		case 92:
			goto tr292
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto st134
		}
		goto tr289
tr291:
//line machine.rl:17

	m.pb = m.p

	goto st135
	st135:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof135
		}
	st_case_135:
//line machine.go:21780
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr294
		case 13:
			goto st6
		case 32:
			goto tr293
		case 34:
			goto tr70
		case 44:
			goto tr161
		case 92:
			goto tr292
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr293
		}
		goto tr289
tr293:
//line machine.rl:58

	m.handler.SetMeasurement(m.text())

	goto st136
	st136:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof136
		}
	st_case_136:
//line machine.go:21812
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr296
		case 13:
			goto st6
		case 32:
			goto st136
		case 34:
			goto tr106
		case 44:
			goto st6
		case 61:
			goto tr289
		case 92:
			goto tr165
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto st136
		}
		goto tr163
tr296:
//line machine.rl:17

	m.pb = m.p

	goto st137
tr297:
//line machine.rl:17

	m.pb = m.p

//line machine.rl:58

	m.handler.SetMeasurement(m.text())

	goto st137
	st137:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof137
		}
	st_case_137:
//line machine.go:21856
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr297
		case 13:
			goto st6
		case 32:
			goto tr293
		case 34:
			goto tr106
		case 44:
			goto tr161
		case 61:
			goto tr298
		case 92:
			goto tr165
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr293
		}
		goto tr163
tr294:
//line machine.rl:17

	m.pb = m.p

//line machine.rl:58

	m.handler.SetMeasurement(m.text())

	goto st138
	st138:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof138
		}
	st_case_138:
//line machine.go:21894
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr297
		case 13:
			goto st6
		case 32:
			goto tr293
		case 34:
			goto tr106
		case 44:
			goto tr161
		case 61:
			goto tr289
		case 92:
			goto tr165
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr293
		}
		goto tr163
tr928:
//line machine.rl:17

	m.pb = m.p

	goto st139
	st139:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof139
		}
	st_case_139:
//line machine.go:21928
		switch ( m.data)[( m.p)] {
		case 34:
			goto tr26
		case 92:
			goto st59
		}
		if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
			goto st795
		}
		goto st6
tr929:
//line machine.rl:17

	m.pb = m.p

	goto st795
	st795:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof795
		}
	st_case_795:
//line machine.go:21950
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1063
		case 13:
			goto tr1063
		case 32:
			goto tr1062
		case 34:
			goto tr26
		case 92:
			goto st59
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st796
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1062
		}
		goto st6
	st796:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof796
		}
	st_case_796:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1063
		case 13:
			goto tr1063
		case 32:
			goto tr1062
		case 34:
			goto tr26
		case 92:
			goto st59
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st797
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1062
		}
		goto st6
	st797:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof797
		}
	st_case_797:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1063
		case 13:
			goto tr1063
		case 32:
			goto tr1062
		case 34:
			goto tr26
		case 92:
			goto st59
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st798
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1062
		}
		goto st6
	st798:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof798
		}
	st_case_798:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1063
		case 13:
			goto tr1063
		case 32:
			goto tr1062
		case 34:
			goto tr26
		case 92:
			goto st59
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st799
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1062
		}
		goto st6
	st799:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof799
		}
	st_case_799:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1063
		case 13:
			goto tr1063
		case 32:
			goto tr1062
		case 34:
			goto tr26
		case 92:
			goto st59
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st800
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1062
		}
		goto st6
	st800:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof800
		}
	st_case_800:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1063
		case 13:
			goto tr1063
		case 32:
			goto tr1062
		case 34:
			goto tr26
		case 92:
			goto st59
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st801
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1062
		}
		goto st6
	st801:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof801
		}
	st_case_801:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1063
		case 13:
			goto tr1063
		case 32:
			goto tr1062
		case 34:
			goto tr26
		case 92:
			goto st59
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st802
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1062
		}
		goto st6
	st802:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof802
		}
	st_case_802:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1063
		case 13:
			goto tr1063
		case 32:
			goto tr1062
		case 34:
			goto tr26
		case 92:
			goto st59
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st803
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1062
		}
		goto st6
	st803:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof803
		}
	st_case_803:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1063
		case 13:
			goto tr1063
		case 32:
			goto tr1062
		case 34:
			goto tr26
		case 92:
			goto st59
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st804
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1062
		}
		goto st6
	st804:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof804
		}
	st_case_804:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1063
		case 13:
			goto tr1063
		case 32:
			goto tr1062
		case 34:
			goto tr26
		case 92:
			goto st59
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st805
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1062
		}
		goto st6
	st805:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof805
		}
	st_case_805:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1063
		case 13:
			goto tr1063
		case 32:
			goto tr1062
		case 34:
			goto tr26
		case 92:
			goto st59
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st806
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1062
		}
		goto st6
	st806:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof806
		}
	st_case_806:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1063
		case 13:
			goto tr1063
		case 32:
			goto tr1062
		case 34:
			goto tr26
		case 92:
			goto st59
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st807
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1062
		}
		goto st6
	st807:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof807
		}
	st_case_807:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1063
		case 13:
			goto tr1063
		case 32:
			goto tr1062
		case 34:
			goto tr26
		case 92:
			goto st59
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st808
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1062
		}
		goto st6
	st808:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof808
		}
	st_case_808:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1063
		case 13:
			goto tr1063
		case 32:
			goto tr1062
		case 34:
			goto tr26
		case 92:
			goto st59
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st809
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1062
		}
		goto st6
	st809:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof809
		}
	st_case_809:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1063
		case 13:
			goto tr1063
		case 32:
			goto tr1062
		case 34:
			goto tr26
		case 92:
			goto st59
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st810
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1062
		}
		goto st6
	st810:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof810
		}
	st_case_810:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1063
		case 13:
			goto tr1063
		case 32:
			goto tr1062
		case 34:
			goto tr26
		case 92:
			goto st59
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st811
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1062
		}
		goto st6
	st811:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof811
		}
	st_case_811:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1063
		case 13:
			goto tr1063
		case 32:
			goto tr1062
		case 34:
			goto tr26
		case 92:
			goto st59
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st812
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1062
		}
		goto st6
	st812:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof812
		}
	st_case_812:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1063
		case 13:
			goto tr1063
		case 32:
			goto tr1062
		case 34:
			goto tr26
		case 92:
			goto st59
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st813
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1062
		}
		goto st6
	st813:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof813
		}
	st_case_813:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1063
		case 13:
			goto tr1063
		case 32:
			goto tr1062
		case 34:
			goto tr26
		case 92:
			goto st59
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1062
		}
		goto st6
tr1260:
//line machine.rl:17

	m.pb = m.p

	goto st140
tr923:
//line machine.rl:78

	m.handler.AddFloat(key, m.text())

	goto st140
tr1262:
//line machine.rl:74

	m.handler.AddInt(key, m.text())

	goto st140
tr1264:
//line machine.rl:82

	m.handler.AddBool(key, m.text())

	goto st140
	st140:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof140
		}
	st_case_140:
//line machine.go:22464
		switch ( m.data)[( m.p)] {
		case 32:
			goto st6
		case 34:
			goto tr82
		case 44:
			goto st6
		case 61:
			goto st6
		case 92:
			goto tr301
		}
		switch {
		case ( m.data)[( m.p)] > 10:
			if 12 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto st6
			}
		case ( m.data)[( m.p)] >= 9:
			goto st6
		}
		goto tr300
tr300:
//line machine.rl:17

	m.pb = m.p

	goto st141
	st141:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof141
		}
	st_case_141:
//line machine.go:22497
		switch ( m.data)[( m.p)] {
		case 32:
			goto st6
		case 34:
			goto tr85
		case 44:
			goto st6
		case 61:
			goto tr303
		case 92:
			goto st151
		}
		switch {
		case ( m.data)[( m.p)] > 10:
			if 12 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto st6
			}
		case ( m.data)[( m.p)] >= 9:
			goto st6
		}
		goto st141
tr303:
//line machine.rl:70

	key = m.text()

	goto st142
	st142:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof142
		}
	st_case_142:
//line machine.go:22530
		switch ( m.data)[( m.p)] {
		case 34:
			goto tr305
		case 45:
			goto tr148
		case 46:
			goto tr149
		case 48:
			goto tr150
		case 70:
			goto tr152
		case 84:
			goto tr153
		case 92:
			goto st59
		case 102:
			goto tr154
		case 116:
			goto tr155
		}
		if 49 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
			goto tr151
		}
		goto st6
tr305:
//line machine.rl:86

	m.handler.AddString(key, m.text())

	goto st814
	st814:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof814
		}
	st_case_814:
//line machine.go:22566
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1055
		case 13:
			goto tr1055
		case 32:
			goto tr1259
		case 34:
			goto tr23
		case 44:
			goto tr1260
		case 92:
			goto tr24
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1259
		}
		goto tr22
tr149:
//line machine.rl:17

	m.pb = m.p

	goto st815
	st815:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof815
		}
	st_case_815:
//line machine.go:22596
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr922
		case 13:
			goto tr922
		case 32:
			goto tr921
		case 34:
			goto tr26
		case 44:
			goto tr923
		case 92:
			goto st59
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st815
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr921
		}
		goto st6
tr151:
//line machine.rl:17

	m.pb = m.p

	goto st816
	st816:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof816
		}
	st_case_816:
//line machine.go:22631
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr922
		case 13:
			goto tr922
		case 32:
			goto tr921
		case 34:
			goto tr26
		case 44:
			goto tr923
		case 46:
			goto st815
		case 92:
			goto st59
		case 105:
			goto st817
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st816
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr921
		}
		goto st6
	st817:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof817
		}
	st_case_817:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1119
		case 13:
			goto tr1119
		case 32:
			goto tr1261
		case 34:
			goto tr26
		case 44:
			goto tr1262
		case 92:
			goto st59
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1261
		}
		goto st6
tr152:
//line machine.rl:17

	m.pb = m.p

	goto st818
	st818:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof818
		}
	st_case_818:
//line machine.go:22693
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1123
		case 13:
			goto tr1123
		case 32:
			goto tr1263
		case 34:
			goto tr26
		case 44:
			goto tr1264
		case 65:
			goto st143
		case 92:
			goto st59
		case 97:
			goto st146
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1263
		}
		goto st6
	st143:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof143
		}
	st_case_143:
		switch ( m.data)[( m.p)] {
		case 34:
			goto tr26
		case 76:
			goto st144
		case 92:
			goto st59
		}
		goto st6
	st144:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof144
		}
	st_case_144:
		switch ( m.data)[( m.p)] {
		case 34:
			goto tr26
		case 83:
			goto st145
		case 92:
			goto st59
		}
		goto st6
	st145:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof145
		}
	st_case_145:
		switch ( m.data)[( m.p)] {
		case 34:
			goto tr26
		case 69:
			goto st819
		case 92:
			goto st59
		}
		goto st6
	st819:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof819
		}
	st_case_819:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1123
		case 13:
			goto tr1123
		case 32:
			goto tr1263
		case 34:
			goto tr26
		case 44:
			goto tr1264
		case 92:
			goto st59
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1263
		}
		goto st6
	st146:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof146
		}
	st_case_146:
		switch ( m.data)[( m.p)] {
		case 34:
			goto tr26
		case 92:
			goto st59
		case 108:
			goto st147
		}
		goto st6
	st147:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof147
		}
	st_case_147:
		switch ( m.data)[( m.p)] {
		case 34:
			goto tr26
		case 92:
			goto st59
		case 115:
			goto st148
		}
		goto st6
	st148:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof148
		}
	st_case_148:
		switch ( m.data)[( m.p)] {
		case 34:
			goto tr26
		case 92:
			goto st59
		case 101:
			goto st819
		}
		goto st6
tr153:
//line machine.rl:17

	m.pb = m.p

	goto st820
	st820:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof820
		}
	st_case_820:
//line machine.go:22834
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1123
		case 13:
			goto tr1123
		case 32:
			goto tr1263
		case 34:
			goto tr26
		case 44:
			goto tr1264
		case 82:
			goto st149
		case 92:
			goto st59
		case 114:
			goto st150
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1263
		}
		goto st6
	st149:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof149
		}
	st_case_149:
		switch ( m.data)[( m.p)] {
		case 34:
			goto tr26
		case 85:
			goto st145
		case 92:
			goto st59
		}
		goto st6
	st150:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof150
		}
	st_case_150:
		switch ( m.data)[( m.p)] {
		case 34:
			goto tr26
		case 92:
			goto st59
		case 117:
			goto st148
		}
		goto st6
tr154:
//line machine.rl:17

	m.pb = m.p

	goto st821
	st821:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof821
		}
	st_case_821:
//line machine.go:22896
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1123
		case 13:
			goto tr1123
		case 32:
			goto tr1263
		case 34:
			goto tr26
		case 44:
			goto tr1264
		case 92:
			goto st59
		case 97:
			goto st146
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1263
		}
		goto st6
tr155:
//line machine.rl:17

	m.pb = m.p

	goto st822
	st822:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof822
		}
	st_case_822:
//line machine.go:22928
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1123
		case 13:
			goto tr1123
		case 32:
			goto tr1263
		case 34:
			goto tr26
		case 44:
			goto tr1264
		case 92:
			goto st59
		case 114:
			goto st150
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1263
		}
		goto st6
tr301:
//line machine.rl:17

	m.pb = m.p

	goto st151
	st151:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof151
		}
	st_case_151:
//line machine.go:22960
		switch ( m.data)[( m.p)] {
		case 34:
			goto st141
		case 92:
			goto st141
		}
		switch {
		case ( m.data)[( m.p)] > 10:
			if 12 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto tr2
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr2
		}
		goto st3
tr143:
//line machine.rl:17

	m.pb = m.p

	goto st152
	st152:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof152
		}
	st_case_152:
//line machine.go:22987
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr143
		case 13:
			goto st6
		case 32:
			goto st42
		case 34:
			goto tr82
		case 44:
			goto st6
		case 61:
			goto tr146
		case 92:
			goto tr144
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto st42
		}
		goto tr141
tr121:
//line machine.rl:17

	m.pb = m.p

	goto st153
	st153:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof153
		}
	st_case_153:
//line machine.go:23021
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr48
		case 11:
			goto tr49
		case 13:
			goto tr48
		case 32:
			goto tr47
		case 44:
			goto tr50
		case 48:
			goto st823
		case 61:
			goto tr48
		case 92:
			goto st19
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 49 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st826
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr47
		}
		goto st14
tr123:
//line machine.rl:17

	m.pb = m.p

	goto st823
	st823:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof823
		}
	st_case_823:
//line machine.go:23060
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1270
		case 11:
			goto tr1271
		case 13:
			goto tr1270
		case 32:
			goto tr1269
		case 44:
			goto tr1272
		case 46:
			goto st824
		case 61:
			goto tr48
		case 92:
			goto st19
		case 105:
			goto st825
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1269
		}
		goto st14
tr122:
//line machine.rl:17

	m.pb = m.p

	goto st824
	st824:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof824
		}
	st_case_824:
//line machine.go:23096
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1270
		case 11:
			goto tr1271
		case 13:
			goto tr1270
		case 32:
			goto tr1269
		case 44:
			goto tr1272
		case 61:
			goto tr48
		case 92:
			goto st19
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st824
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1269
		}
		goto st14
	st825:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof825
		}
	st_case_825:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1276
		case 11:
			goto tr1277
		case 13:
			goto tr1276
		case 32:
			goto tr1275
		case 44:
			goto tr1278
		case 61:
			goto tr48
		case 92:
			goto st19
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1275
		}
		goto st14
tr124:
//line machine.rl:17

	m.pb = m.p

	goto st826
	st826:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof826
		}
	st_case_826:
//line machine.go:23158
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1270
		case 11:
			goto tr1271
		case 13:
			goto tr1270
		case 32:
			goto tr1269
		case 44:
			goto tr1272
		case 46:
			goto st824
		case 61:
			goto tr48
		case 92:
			goto st19
		case 105:
			goto st825
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st826
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1269
		}
		goto st14
tr125:
//line machine.rl:17

	m.pb = m.p

	goto st827
	st827:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof827
		}
	st_case_827:
//line machine.go:23199
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1280
		case 11:
			goto tr1281
		case 13:
			goto tr1280
		case 32:
			goto tr1279
		case 44:
			goto tr1282
		case 61:
			goto tr48
		case 65:
			goto st154
		case 92:
			goto st19
		case 97:
			goto st157
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1279
		}
		goto st14
	st154:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof154
		}
	st_case_154:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr48
		case 11:
			goto tr49
		case 13:
			goto tr48
		case 32:
			goto tr47
		case 44:
			goto tr50
		case 61:
			goto tr48
		case 76:
			goto st155
		case 92:
			goto st19
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr47
		}
		goto st14
	st155:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof155
		}
	st_case_155:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr48
		case 11:
			goto tr49
		case 13:
			goto tr48
		case 32:
			goto tr47
		case 44:
			goto tr50
		case 61:
			goto tr48
		case 83:
			goto st156
		case 92:
			goto st19
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr47
		}
		goto st14
	st156:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof156
		}
	st_case_156:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr48
		case 11:
			goto tr49
		case 13:
			goto tr48
		case 32:
			goto tr47
		case 44:
			goto tr50
		case 61:
			goto tr48
		case 69:
			goto st828
		case 92:
			goto st19
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr47
		}
		goto st14
	st828:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof828
		}
	st_case_828:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1280
		case 11:
			goto tr1281
		case 13:
			goto tr1280
		case 32:
			goto tr1279
		case 44:
			goto tr1282
		case 61:
			goto tr48
		case 92:
			goto st19
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1279
		}
		goto st14
	st157:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof157
		}
	st_case_157:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr48
		case 11:
			goto tr49
		case 13:
			goto tr48
		case 32:
			goto tr47
		case 44:
			goto tr50
		case 61:
			goto tr48
		case 92:
			goto st19
		case 108:
			goto st158
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr47
		}
		goto st14
	st158:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof158
		}
	st_case_158:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr48
		case 11:
			goto tr49
		case 13:
			goto tr48
		case 32:
			goto tr47
		case 44:
			goto tr50
		case 61:
			goto tr48
		case 92:
			goto st19
		case 115:
			goto st159
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr47
		}
		goto st14
	st159:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof159
		}
	st_case_159:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr48
		case 11:
			goto tr49
		case 13:
			goto tr48
		case 32:
			goto tr47
		case 44:
			goto tr50
		case 61:
			goto tr48
		case 92:
			goto st19
		case 101:
			goto st828
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr47
		}
		goto st14
tr126:
//line machine.rl:17

	m.pb = m.p

	goto st829
	st829:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof829
		}
	st_case_829:
//line machine.go:23422
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1280
		case 11:
			goto tr1281
		case 13:
			goto tr1280
		case 32:
			goto tr1279
		case 44:
			goto tr1282
		case 61:
			goto tr48
		case 82:
			goto st160
		case 92:
			goto st19
		case 114:
			goto st161
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1279
		}
		goto st14
	st160:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof160
		}
	st_case_160:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr48
		case 11:
			goto tr49
		case 13:
			goto tr48
		case 32:
			goto tr47
		case 44:
			goto tr50
		case 61:
			goto tr48
		case 85:
			goto st156
		case 92:
			goto st19
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr47
		}
		goto st14
	st161:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof161
		}
	st_case_161:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr48
		case 11:
			goto tr49
		case 13:
			goto tr48
		case 32:
			goto tr47
		case 44:
			goto tr50
		case 61:
			goto tr48
		case 92:
			goto st19
		case 117:
			goto st159
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr47
		}
		goto st14
tr127:
//line machine.rl:17

	m.pb = m.p

	goto st830
	st830:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof830
		}
	st_case_830:
//line machine.go:23512
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1280
		case 11:
			goto tr1281
		case 13:
			goto tr1280
		case 32:
			goto tr1279
		case 44:
			goto tr1282
		case 61:
			goto tr48
		case 92:
			goto st19
		case 97:
			goto st157
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1279
		}
		goto st14
tr128:
//line machine.rl:17

	m.pb = m.p

	goto st831
	st831:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof831
		}
	st_case_831:
//line machine.go:23546
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1280
		case 11:
			goto tr1281
		case 13:
			goto tr1280
		case 32:
			goto tr1279
		case 44:
			goto tr1282
		case 61:
			goto tr48
		case 92:
			goto st19
		case 114:
			goto st161
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1279
		}
		goto st14
tr105:
//line machine.rl:17

	m.pb = m.p

	goto st162
tr318:
//line machine.rl:17

	m.pb = m.p

//line machine.rl:58

	m.handler.SetMeasurement(m.text())

	goto st162
	st162:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof162
		}
	st_case_162:
//line machine.go:23590
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr318
		case 13:
			goto st6
		case 32:
			goto tr101
		case 34:
			goto tr106
		case 44:
			goto tr77
		case 61:
			goto tr319
		case 92:
			goto tr107
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr101
		}
		goto tr103
tr102:
//line machine.rl:17

	m.pb = m.p

//line machine.rl:58

	m.handler.SetMeasurement(m.text())

	goto st163
	st163:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof163
		}
	st_case_163:
//line machine.go:23628
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr318
		case 13:
			goto st6
		case 32:
			goto tr101
		case 34:
			goto tr106
		case 44:
			goto tr77
		case 61:
			goto tr67
		case 92:
			goto tr107
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr101
		}
		goto tr103
tr886:
//line machine.rl:17

	m.pb = m.p

	goto st164
	st164:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof164
		}
	st_case_164:
//line machine.go:23662
		switch ( m.data)[( m.p)] {
		case 34:
			goto tr26
		case 92:
			goto st59
		}
		if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
			goto st832
		}
		goto st6
tr887:
//line machine.rl:17

	m.pb = m.p

	goto st832
	st832:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof832
		}
	st_case_832:
//line machine.go:23684
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr990
		case 13:
			goto tr990
		case 32:
			goto tr989
		case 34:
			goto tr26
		case 92:
			goto st59
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st833
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr989
		}
		goto st6
	st833:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof833
		}
	st_case_833:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr990
		case 13:
			goto tr990
		case 32:
			goto tr989
		case 34:
			goto tr26
		case 92:
			goto st59
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st834
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr989
		}
		goto st6
	st834:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof834
		}
	st_case_834:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr990
		case 13:
			goto tr990
		case 32:
			goto tr989
		case 34:
			goto tr26
		case 92:
			goto st59
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st835
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr989
		}
		goto st6
	st835:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof835
		}
	st_case_835:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr990
		case 13:
			goto tr990
		case 32:
			goto tr989
		case 34:
			goto tr26
		case 92:
			goto st59
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st836
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr989
		}
		goto st6
	st836:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof836
		}
	st_case_836:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr990
		case 13:
			goto tr990
		case 32:
			goto tr989
		case 34:
			goto tr26
		case 92:
			goto st59
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st837
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr989
		}
		goto st6
	st837:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof837
		}
	st_case_837:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr990
		case 13:
			goto tr990
		case 32:
			goto tr989
		case 34:
			goto tr26
		case 92:
			goto st59
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st838
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr989
		}
		goto st6
	st838:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof838
		}
	st_case_838:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr990
		case 13:
			goto tr990
		case 32:
			goto tr989
		case 34:
			goto tr26
		case 92:
			goto st59
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st839
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr989
		}
		goto st6
	st839:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof839
		}
	st_case_839:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr990
		case 13:
			goto tr990
		case 32:
			goto tr989
		case 34:
			goto tr26
		case 92:
			goto st59
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st840
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr989
		}
		goto st6
	st840:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof840
		}
	st_case_840:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr990
		case 13:
			goto tr990
		case 32:
			goto tr989
		case 34:
			goto tr26
		case 92:
			goto st59
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st841
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr989
		}
		goto st6
	st841:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof841
		}
	st_case_841:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr990
		case 13:
			goto tr990
		case 32:
			goto tr989
		case 34:
			goto tr26
		case 92:
			goto st59
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st842
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr989
		}
		goto st6
	st842:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof842
		}
	st_case_842:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr990
		case 13:
			goto tr990
		case 32:
			goto tr989
		case 34:
			goto tr26
		case 92:
			goto st59
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st843
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr989
		}
		goto st6
	st843:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof843
		}
	st_case_843:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr990
		case 13:
			goto tr990
		case 32:
			goto tr989
		case 34:
			goto tr26
		case 92:
			goto st59
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st844
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr989
		}
		goto st6
	st844:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof844
		}
	st_case_844:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr990
		case 13:
			goto tr990
		case 32:
			goto tr989
		case 34:
			goto tr26
		case 92:
			goto st59
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st845
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr989
		}
		goto st6
	st845:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof845
		}
	st_case_845:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr990
		case 13:
			goto tr990
		case 32:
			goto tr989
		case 34:
			goto tr26
		case 92:
			goto st59
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st846
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr989
		}
		goto st6
	st846:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof846
		}
	st_case_846:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr990
		case 13:
			goto tr990
		case 32:
			goto tr989
		case 34:
			goto tr26
		case 92:
			goto st59
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st847
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr989
		}
		goto st6
	st847:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof847
		}
	st_case_847:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr990
		case 13:
			goto tr990
		case 32:
			goto tr989
		case 34:
			goto tr26
		case 92:
			goto st59
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st848
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr989
		}
		goto st6
	st848:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof848
		}
	st_case_848:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr990
		case 13:
			goto tr990
		case 32:
			goto tr989
		case 34:
			goto tr26
		case 92:
			goto st59
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st849
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr989
		}
		goto st6
	st849:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof849
		}
	st_case_849:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr990
		case 13:
			goto tr990
		case 32:
			goto tr989
		case 34:
			goto tr26
		case 92:
			goto st59
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st850
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr989
		}
		goto st6
	st850:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof850
		}
	st_case_850:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr990
		case 13:
			goto tr990
		case 32:
			goto tr989
		case 34:
			goto tr26
		case 92:
			goto st59
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr989
		}
		goto st6
tr883:
//line machine.rl:17

	m.pb = m.p

	goto st165
tr1306:
//line machine.rl:78

	m.handler.AddFloat(key, m.text())

	goto st165
tr1310:
//line machine.rl:74

	m.handler.AddInt(key, m.text())

	goto st165
tr1312:
//line machine.rl:82

	m.handler.AddBool(key, m.text())

	goto st165
	st165:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof165
		}
	st_case_165:
//line machine.go:24198
		switch ( m.data)[( m.p)] {
		case 32:
			goto st6
		case 34:
			goto tr322
		case 44:
			goto st6
		case 61:
			goto st6
		case 92:
			goto tr323
		}
		switch {
		case ( m.data)[( m.p)] > 10:
			if 12 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto st6
			}
		case ( m.data)[( m.p)] >= 9:
			goto st6
		}
		goto tr321
tr321:
//line machine.rl:17

	m.pb = m.p

	goto st166
	st166:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof166
		}
	st_case_166:
//line machine.go:24231
		switch ( m.data)[( m.p)] {
		case 32:
			goto st6
		case 34:
			goto tr85
		case 44:
			goto st6
		case 61:
			goto tr325
		case 92:
			goto st177
		}
		switch {
		case ( m.data)[( m.p)] > 10:
			if 12 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto st6
			}
		case ( m.data)[( m.p)] >= 9:
			goto st6
		}
		goto st166
tr325:
//line machine.rl:70

	key = m.text()

	goto st167
	st167:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof167
		}
	st_case_167:
//line machine.go:24264
		switch ( m.data)[( m.p)] {
		case 34:
			goto tr305
		case 45:
			goto tr91
		case 46:
			goto tr92
		case 48:
			goto tr93
		case 70:
			goto tr95
		case 84:
			goto tr96
		case 92:
			goto st59
		case 102:
			goto tr97
		case 116:
			goto tr98
		}
		if 49 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
			goto tr94
		}
		goto st6
tr91:
//line machine.rl:17

	m.pb = m.p

	goto st168
	st168:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof168
		}
	st_case_168:
//line machine.go:24300
		switch ( m.data)[( m.p)] {
		case 34:
			goto tr26
		case 48:
			goto st851
		case 92:
			goto st59
		}
		if 49 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
			goto st854
		}
		goto st6
tr93:
//line machine.rl:17

	m.pb = m.p

	goto st851
	st851:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof851
		}
	st_case_851:
//line machine.go:24324
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1026
		case 13:
			goto tr1026
		case 32:
			goto tr1305
		case 34:
			goto tr26
		case 44:
			goto tr1306
		case 46:
			goto st852
		case 92:
			goto st59
		case 105:
			goto st853
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1305
		}
		goto st6
tr92:
//line machine.rl:17

	m.pb = m.p

	goto st852
	st852:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof852
		}
	st_case_852:
//line machine.go:24358
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1026
		case 13:
			goto tr1026
		case 32:
			goto tr1305
		case 34:
			goto tr26
		case 44:
			goto tr1306
		case 92:
			goto st59
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st852
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1305
		}
		goto st6
	st853:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof853
		}
	st_case_853:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1162
		case 13:
			goto tr1162
		case 32:
			goto tr1309
		case 34:
			goto tr26
		case 44:
			goto tr1310
		case 92:
			goto st59
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1309
		}
		goto st6
tr94:
//line machine.rl:17

	m.pb = m.p

	goto st854
	st854:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof854
		}
	st_case_854:
//line machine.go:24416
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1026
		case 13:
			goto tr1026
		case 32:
			goto tr1305
		case 34:
			goto tr26
		case 44:
			goto tr1306
		case 46:
			goto st852
		case 92:
			goto st59
		case 105:
			goto st853
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st854
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1305
		}
		goto st6
tr95:
//line machine.rl:17

	m.pb = m.p

	goto st855
	st855:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof855
		}
	st_case_855:
//line machine.go:24455
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1166
		case 13:
			goto tr1166
		case 32:
			goto tr1311
		case 34:
			goto tr26
		case 44:
			goto tr1312
		case 65:
			goto st169
		case 92:
			goto st59
		case 97:
			goto st172
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1311
		}
		goto st6
	st169:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof169
		}
	st_case_169:
		switch ( m.data)[( m.p)] {
		case 34:
			goto tr26
		case 76:
			goto st170
		case 92:
			goto st59
		}
		goto st6
	st170:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof170
		}
	st_case_170:
		switch ( m.data)[( m.p)] {
		case 34:
			goto tr26
		case 83:
			goto st171
		case 92:
			goto st59
		}
		goto st6
	st171:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof171
		}
	st_case_171:
		switch ( m.data)[( m.p)] {
		case 34:
			goto tr26
		case 69:
			goto st856
		case 92:
			goto st59
		}
		goto st6
	st856:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof856
		}
	st_case_856:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1166
		case 13:
			goto tr1166
		case 32:
			goto tr1311
		case 34:
			goto tr26
		case 44:
			goto tr1312
		case 92:
			goto st59
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1311
		}
		goto st6
	st172:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof172
		}
	st_case_172:
		switch ( m.data)[( m.p)] {
		case 34:
			goto tr26
		case 92:
			goto st59
		case 108:
			goto st173
		}
		goto st6
	st173:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof173
		}
	st_case_173:
		switch ( m.data)[( m.p)] {
		case 34:
			goto tr26
		case 92:
			goto st59
		case 115:
			goto st174
		}
		goto st6
	st174:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof174
		}
	st_case_174:
		switch ( m.data)[( m.p)] {
		case 34:
			goto tr26
		case 92:
			goto st59
		case 101:
			goto st856
		}
		goto st6
tr96:
//line machine.rl:17

	m.pb = m.p

	goto st857
	st857:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof857
		}
	st_case_857:
//line machine.go:24596
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1166
		case 13:
			goto tr1166
		case 32:
			goto tr1311
		case 34:
			goto tr26
		case 44:
			goto tr1312
		case 82:
			goto st175
		case 92:
			goto st59
		case 114:
			goto st176
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1311
		}
		goto st6
	st175:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof175
		}
	st_case_175:
		switch ( m.data)[( m.p)] {
		case 34:
			goto tr26
		case 85:
			goto st171
		case 92:
			goto st59
		}
		goto st6
	st176:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof176
		}
	st_case_176:
		switch ( m.data)[( m.p)] {
		case 34:
			goto tr26
		case 92:
			goto st59
		case 117:
			goto st174
		}
		goto st6
tr97:
//line machine.rl:17

	m.pb = m.p

	goto st858
	st858:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof858
		}
	st_case_858:
//line machine.go:24658
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1166
		case 13:
			goto tr1166
		case 32:
			goto tr1311
		case 34:
			goto tr26
		case 44:
			goto tr1312
		case 92:
			goto st59
		case 97:
			goto st172
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1311
		}
		goto st6
tr98:
//line machine.rl:17

	m.pb = m.p

	goto st859
	st859:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof859
		}
	st_case_859:
//line machine.go:24690
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1166
		case 13:
			goto tr1166
		case 32:
			goto tr1311
		case 34:
			goto tr26
		case 44:
			goto tr1312
		case 92:
			goto st59
		case 114:
			goto st176
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1311
		}
		goto st6
tr323:
//line machine.rl:17

	m.pb = m.p

	goto st177
	st177:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof177
		}
	st_case_177:
//line machine.go:24722
		switch ( m.data)[( m.p)] {
		case 34:
			goto st166
		case 92:
			goto st166
		}
		switch {
		case ( m.data)[( m.p)] > 10:
			if 12 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto tr2
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr2
		}
		goto st3
tr81:
//line machine.rl:17

	m.pb = m.p

	goto st178
	st178:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof178
		}
	st_case_178:
//line machine.go:24749
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr81
		case 13:
			goto st6
		case 32:
			goto st24
		case 34:
			goto tr82
		case 44:
			goto st6
		case 61:
			goto tr86
		case 92:
			goto tr83
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto st24
		}
		goto tr79
tr59:
//line machine.rl:17

	m.pb = m.p

	goto st179
	st179:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof179
		}
	st_case_179:
//line machine.go:24783
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr2
		case 11:
			goto tr29
		case 13:
			goto tr2
		case 32:
			goto tr1
		case 44:
			goto tr4
		case 48:
			goto st860
		case 92:
			goto st76
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 49 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st863
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1
		}
		goto st7
tr61:
//line machine.rl:17

	m.pb = m.p

	goto st860
	st860:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof860
		}
	st_case_860:
//line machine.go:24820
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1270
		case 11:
			goto tr1318
		case 13:
			goto tr1270
		case 32:
			goto tr1317
		case 44:
			goto tr1319
		case 46:
			goto st861
		case 92:
			goto st76
		case 105:
			goto st862
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1317
		}
		goto st7
tr60:
//line machine.rl:17

	m.pb = m.p

	goto st861
	st861:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof861
		}
	st_case_861:
//line machine.go:24854
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1270
		case 11:
			goto tr1318
		case 13:
			goto tr1270
		case 32:
			goto tr1317
		case 44:
			goto tr1319
		case 92:
			goto st76
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st861
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1317
		}
		goto st7
	st862:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof862
		}
	st_case_862:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1276
		case 11:
			goto tr1323
		case 13:
			goto tr1276
		case 32:
			goto tr1322
		case 44:
			goto tr1324
		case 92:
			goto st76
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1322
		}
		goto st7
tr62:
//line machine.rl:17

	m.pb = m.p

	goto st863
	st863:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof863
		}
	st_case_863:
//line machine.go:24912
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1270
		case 11:
			goto tr1318
		case 13:
			goto tr1270
		case 32:
			goto tr1317
		case 44:
			goto tr1319
		case 46:
			goto st861
		case 92:
			goto st76
		case 105:
			goto st862
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st863
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1317
		}
		goto st7
tr63:
//line machine.rl:17

	m.pb = m.p

	goto st864
	st864:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof864
		}
	st_case_864:
//line machine.go:24951
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1280
		case 11:
			goto tr1326
		case 13:
			goto tr1280
		case 32:
			goto tr1325
		case 44:
			goto tr1327
		case 65:
			goto st180
		case 92:
			goto st76
		case 97:
			goto st183
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1325
		}
		goto st7
	st180:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof180
		}
	st_case_180:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr2
		case 11:
			goto tr29
		case 13:
			goto tr2
		case 32:
			goto tr1
		case 44:
			goto tr4
		case 76:
			goto st181
		case 92:
			goto st76
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1
		}
		goto st7
	st181:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof181
		}
	st_case_181:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr2
		case 11:
			goto tr29
		case 13:
			goto tr2
		case 32:
			goto tr1
		case 44:
			goto tr4
		case 83:
			goto st182
		case 92:
			goto st76
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1
		}
		goto st7
	st182:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof182
		}
	st_case_182:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr2
		case 11:
			goto tr29
		case 13:
			goto tr2
		case 32:
			goto tr1
		case 44:
			goto tr4
		case 69:
			goto st865
		case 92:
			goto st76
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1
		}
		goto st7
	st865:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof865
		}
	st_case_865:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1280
		case 11:
			goto tr1326
		case 13:
			goto tr1280
		case 32:
			goto tr1325
		case 44:
			goto tr1327
		case 92:
			goto st76
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1325
		}
		goto st7
	st183:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof183
		}
	st_case_183:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr2
		case 11:
			goto tr29
		case 13:
			goto tr2
		case 32:
			goto tr1
		case 44:
			goto tr4
		case 92:
			goto st76
		case 108:
			goto st184
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1
		}
		goto st7
	st184:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof184
		}
	st_case_184:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr2
		case 11:
			goto tr29
		case 13:
			goto tr2
		case 32:
			goto tr1
		case 44:
			goto tr4
		case 92:
			goto st76
		case 115:
			goto st185
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1
		}
		goto st7
	st185:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof185
		}
	st_case_185:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr2
		case 11:
			goto tr29
		case 13:
			goto tr2
		case 32:
			goto tr1
		case 44:
			goto tr4
		case 92:
			goto st76
		case 101:
			goto st865
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1
		}
		goto st7
tr64:
//line machine.rl:17

	m.pb = m.p

	goto st866
	st866:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof866
		}
	st_case_866:
//line machine.go:25158
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1280
		case 11:
			goto tr1326
		case 13:
			goto tr1280
		case 32:
			goto tr1325
		case 44:
			goto tr1327
		case 82:
			goto st186
		case 92:
			goto st76
		case 114:
			goto st187
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1325
		}
		goto st7
	st186:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof186
		}
	st_case_186:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr2
		case 11:
			goto tr29
		case 13:
			goto tr2
		case 32:
			goto tr1
		case 44:
			goto tr4
		case 85:
			goto st182
		case 92:
			goto st76
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1
		}
		goto st7
	st187:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof187
		}
	st_case_187:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr2
		case 11:
			goto tr29
		case 13:
			goto tr2
		case 32:
			goto tr1
		case 44:
			goto tr4
		case 92:
			goto st76
		case 117:
			goto st185
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1
		}
		goto st7
tr65:
//line machine.rl:17

	m.pb = m.p

	goto st867
	st867:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof867
		}
	st_case_867:
//line machine.go:25242
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1280
		case 11:
			goto tr1326
		case 13:
			goto tr1280
		case 32:
			goto tr1325
		case 44:
			goto tr1327
		case 92:
			goto st76
		case 97:
			goto st183
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1325
		}
		goto st7
tr66:
//line machine.rl:17

	m.pb = m.p

	goto st868
	st868:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof868
		}
	st_case_868:
//line machine.go:25274
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1280
		case 11:
			goto tr1326
		case 13:
			goto tr1280
		case 32:
			goto tr1325
		case 44:
			goto tr1327
		case 92:
			goto st76
		case 114:
			goto st187
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1325
		}
		goto st7
	st188:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof188
		}
	st_case_188:
		switch ( m.data)[( m.p)] {
		case 10:
			goto st0
		case 11:
			goto tr343
		case 13:
			goto st0
		case 32:
			goto st188
		case 44:
			goto st0
		case 92:
			goto tr344
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto st188
		}
		goto tr341
tr343:
//line machine.rl:17

	m.pb = m.p

	goto st189
	st189:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof189
		}
	st_case_189:
//line machine.go:25329
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr2
		case 11:
			goto tr346
		case 13:
			goto tr2
		case 32:
			goto tr345
		case 44:
			goto tr4
		case 92:
			goto tr344
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr345
		}
		goto tr341
tr345:
//line machine.rl:58

	m.handler.SetMeasurement(m.text())

	goto st190
	st190:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof190
		}
	st_case_190:
//line machine.go:25359
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr2
		case 11:
			goto tr348
		case 13:
			goto tr2
		case 32:
			goto st190
		case 44:
			goto tr2
		case 61:
			goto tr341
		case 92:
			goto tr33
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto st190
		}
		goto tr31
tr348:
//line machine.rl:17

	m.pb = m.p

	goto st191
tr349:
//line machine.rl:17

	m.pb = m.p

//line machine.rl:58

	m.handler.SetMeasurement(m.text())

	goto st191
	st191:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof191
		}
	st_case_191:
//line machine.go:25401
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr2
		case 11:
			goto tr349
		case 13:
			goto tr2
		case 32:
			goto tr345
		case 44:
			goto tr4
		case 61:
			goto tr350
		case 92:
			goto tr33
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr345
		}
		goto tr31
tr346:
//line machine.rl:17

	m.pb = m.p

//line machine.rl:58

	m.handler.SetMeasurement(m.text())

	goto st192
	st192:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof192
		}
	st_case_192:
//line machine.go:25437
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr2
		case 11:
			goto tr349
		case 13:
			goto tr2
		case 32:
			goto tr345
		case 44:
			goto tr4
		case 61:
			goto tr341
		case 92:
			goto tr33
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr345
		}
		goto tr31
tr852:
//line machine.rl:17

	m.pb = m.p

	goto st193
	st193:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof193
		}
	st_case_193:
//line machine.go:25469
		if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
			goto st869
		}
		goto tr351
tr853:
//line machine.rl:17

	m.pb = m.p

	goto st869
	st869:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof869
		}
	st_case_869:
//line machine.go:25485
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 13:
			goto tr859
		case 32:
			goto tr858
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st870
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr858
		}
		goto tr351
	st870:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof870
		}
	st_case_870:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 13:
			goto tr859
		case 32:
			goto tr858
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st871
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr858
		}
		goto tr351
	st871:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof871
		}
	st_case_871:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 13:
			goto tr859
		case 32:
			goto tr858
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st872
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr858
		}
		goto tr351
	st872:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof872
		}
	st_case_872:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 13:
			goto tr859
		case 32:
			goto tr858
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st873
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr858
		}
		goto tr351
	st873:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof873
		}
	st_case_873:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 13:
			goto tr859
		case 32:
			goto tr858
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st874
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr858
		}
		goto tr351
	st874:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof874
		}
	st_case_874:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 13:
			goto tr859
		case 32:
			goto tr858
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st875
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr858
		}
		goto tr351
	st875:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof875
		}
	st_case_875:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 13:
			goto tr859
		case 32:
			goto tr858
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st876
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr858
		}
		goto tr351
	st876:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof876
		}
	st_case_876:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 13:
			goto tr859
		case 32:
			goto tr858
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st877
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr858
		}
		goto tr351
	st877:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof877
		}
	st_case_877:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 13:
			goto tr859
		case 32:
			goto tr858
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st878
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr858
		}
		goto tr351
	st878:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof878
		}
	st_case_878:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 13:
			goto tr859
		case 32:
			goto tr858
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st879
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr858
		}
		goto tr351
	st879:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof879
		}
	st_case_879:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 13:
			goto tr859
		case 32:
			goto tr858
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st880
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr858
		}
		goto tr351
	st880:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof880
		}
	st_case_880:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 13:
			goto tr859
		case 32:
			goto tr858
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st881
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr858
		}
		goto tr351
	st881:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof881
		}
	st_case_881:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 13:
			goto tr859
		case 32:
			goto tr858
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st882
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr858
		}
		goto tr351
	st882:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof882
		}
	st_case_882:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 13:
			goto tr859
		case 32:
			goto tr858
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st883
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr858
		}
		goto tr351
	st883:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof883
		}
	st_case_883:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 13:
			goto tr859
		case 32:
			goto tr858
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st884
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr858
		}
		goto tr351
	st884:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof884
		}
	st_case_884:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 13:
			goto tr859
		case 32:
			goto tr858
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st885
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr858
		}
		goto tr351
	st885:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof885
		}
	st_case_885:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 13:
			goto tr859
		case 32:
			goto tr858
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st886
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr858
		}
		goto tr351
	st886:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof886
		}
	st_case_886:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 13:
			goto tr859
		case 32:
			goto tr858
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st887
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr858
		}
		goto tr351
	st887:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof887
		}
	st_case_887:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 13:
			goto tr859
		case 32:
			goto tr858
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr858
		}
		goto tr351
tr14:
//line machine.rl:17

	m.pb = m.p

	goto st194
	st194:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof194
		}
	st_case_194:
//line machine.go:25905
		if ( m.data)[( m.p)] == 48 {
			goto st888
		}
		if 49 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
			goto st891
		}
		goto tr2
tr16:
//line machine.rl:17

	m.pb = m.p

	goto st888
	st888:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof888
		}
	st_case_888:
//line machine.go:25924
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1270
		case 13:
			goto tr1270
		case 32:
			goto tr1350
		case 44:
			goto tr1351
		case 46:
			goto st889
		case 105:
			goto st890
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1350
		}
		goto tr2
tr15:
//line machine.rl:17

	m.pb = m.p

	goto st889
	st889:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof889
		}
	st_case_889:
//line machine.go:25954
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1270
		case 13:
			goto tr1270
		case 32:
			goto tr1350
		case 44:
			goto tr1351
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st889
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1350
		}
		goto tr2
	st890:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof890
		}
	st_case_890:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1276
		case 13:
			goto tr1276
		case 32:
			goto tr1354
		case 44:
			goto tr1355
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1354
		}
		goto tr2
tr17:
//line machine.rl:17

	m.pb = m.p

	goto st891
	st891:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof891
		}
	st_case_891:
//line machine.go:26004
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1270
		case 13:
			goto tr1270
		case 32:
			goto tr1350
		case 44:
			goto tr1351
		case 46:
			goto st889
		case 105:
			goto st890
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st891
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1350
		}
		goto tr2
tr18:
//line machine.rl:17

	m.pb = m.p

	goto st892
	st892:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof892
		}
	st_case_892:
//line machine.go:26039
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1280
		case 13:
			goto tr1280
		case 32:
			goto tr1356
		case 44:
			goto tr1357
		case 65:
			goto st195
		case 97:
			goto st198
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1356
		}
		goto tr2
	st195:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof195
		}
	st_case_195:
		if ( m.data)[( m.p)] == 76 {
			goto st196
		}
		goto tr2
	st196:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof196
		}
	st_case_196:
		if ( m.data)[( m.p)] == 83 {
			goto st197
		}
		goto tr2
	st197:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof197
		}
	st_case_197:
		if ( m.data)[( m.p)] == 69 {
			goto st893
		}
		goto tr2
	st893:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof893
		}
	st_case_893:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1280
		case 13:
			goto tr1280
		case 32:
			goto tr1356
		case 44:
			goto tr1357
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1356
		}
		goto tr2
	st198:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof198
		}
	st_case_198:
		if ( m.data)[( m.p)] == 108 {
			goto st199
		}
		goto tr2
	st199:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof199
		}
	st_case_199:
		if ( m.data)[( m.p)] == 115 {
			goto st200
		}
		goto tr2
	st200:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof200
		}
	st_case_200:
		if ( m.data)[( m.p)] == 101 {
			goto st893
		}
		goto tr2
tr19:
//line machine.rl:17

	m.pb = m.p

	goto st894
	st894:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof894
		}
	st_case_894:
//line machine.go:26142
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1280
		case 13:
			goto tr1280
		case 32:
			goto tr1356
		case 44:
			goto tr1357
		case 82:
			goto st201
		case 114:
			goto st202
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1356
		}
		goto tr2
	st201:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof201
		}
	st_case_201:
		if ( m.data)[( m.p)] == 85 {
			goto st197
		}
		goto tr2
	st202:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof202
		}
	st_case_202:
		if ( m.data)[( m.p)] == 117 {
			goto st200
		}
		goto tr2
tr20:
//line machine.rl:17

	m.pb = m.p

	goto st895
	st895:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof895
		}
	st_case_895:
//line machine.go:26190
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1280
		case 13:
			goto tr1280
		case 32:
			goto tr1356
		case 44:
			goto tr1357
		case 97:
			goto st198
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1356
		}
		goto tr2
tr21:
//line machine.rl:17

	m.pb = m.p

	goto st896
	st896:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof896
		}
	st_case_896:
//line machine.go:26218
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1280
		case 13:
			goto tr1280
		case 32:
			goto tr1356
		case 44:
			goto tr1357
		case 114:
			goto st202
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1356
		}
		goto tr2
tr8:
//line machine.rl:17

	m.pb = m.p

	goto st203
	st203:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof203
		}
	st_case_203:
//line machine.go:26246
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr2
		case 11:
			goto tr8
		case 13:
			goto tr2
		case 32:
			goto st2
		case 44:
			goto tr2
		case 61:
			goto tr11
		case 92:
			goto tr9
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto st2
		}
		goto tr6
tr3:
//line machine.rl:58

	m.handler.SetMeasurement(m.text())

	goto st204
	st204:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof204
		}
	st_case_204:
//line machine.go:26278
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr2
		case 11:
			goto tr361
		case 13:
			goto tr2
		case 32:
			goto tr1
		case 44:
			goto tr4
		case 61:
			goto st1
		case 92:
			goto tr362
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1
		}
		goto tr360
tr360:
//line machine.rl:17

	m.pb = m.p

	goto st205
	st205:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof205
		}
	st_case_205:
//line machine.go:26310
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr2
		case 11:
			goto tr364
		case 13:
			goto tr2
		case 32:
			goto tr1
		case 44:
			goto tr4
		case 61:
			goto tr365
		case 92:
			goto st390
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1
		}
		goto st205
tr364:
//line machine.rl:58

	m.handler.SetMeasurement(m.text())

	goto st206
tr361:
//line machine.rl:58

	m.handler.SetMeasurement(m.text())

//line machine.rl:17

	m.pb = m.p

	goto st206
	st206:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof206
		}
	st_case_206:
//line machine.go:26352
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr2
		case 11:
			goto tr361
		case 13:
			goto tr2
		case 32:
			goto tr1
		case 44:
			goto tr4
		case 61:
			goto tr365
		case 92:
			goto tr362
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1
		}
		goto tr360
tr365:
//line machine.rl:70

	key = m.text()

	goto st207
tr846:
//line machine.rl:17

	m.pb = m.p

//line machine.rl:70

	key = m.text()

	goto st207
	st207:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof207
		}
	st_case_207:
//line machine.go:26394
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr2
		case 11:
			goto tr3
		case 13:
			goto tr2
		case 32:
			goto tr1
		case 34:
			goto st208
		case 44:
			goto tr4
		case 45:
			goto tr368
		case 46:
			goto tr369
		case 48:
			goto tr370
		case 70:
			goto tr372
		case 84:
			goto tr373
		case 92:
			goto st434
		case 102:
			goto tr374
		case 116:
			goto tr375
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 49 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto tr371
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1
		}
		goto st1
	st208:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof208
		}
	st_case_208:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr22
		case 11:
			goto tr378
		case 13:
			goto tr22
		case 32:
			goto tr377
		case 34:
			goto tr379
		case 44:
			goto tr380
		case 92:
			goto tr381
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr377
		}
		goto tr376
tr376:
//line machine.rl:17

	m.pb = m.p

	goto st209
	st209:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof209
		}
	st_case_209:
//line machine.go:26470
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr384
		case 13:
			goto st6
		case 32:
			goto tr383
		case 34:
			goto tr385
		case 44:
			goto tr386
		case 92:
			goto st465
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr383
		}
		goto st209
tr383:
//line machine.rl:58

	m.handler.SetMeasurement(m.text())

	goto st210
tr377:
//line machine.rl:58

	m.handler.SetMeasurement(m.text())

//line machine.rl:17

	m.pb = m.p

	goto st210
tr755:
//line machine.rl:66

	m.handler.AddTag(key, m.text())

	goto st210
	st210:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof210
		}
	st_case_210:
//line machine.go:26518
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr390
		case 13:
			goto st6
		case 32:
			goto st210
		case 34:
			goto tr82
		case 44:
			goto st6
		case 61:
			goto st6
		case 92:
			goto tr391
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto st210
		}
		goto tr388
tr388:
//line machine.rl:17

	m.pb = m.p

	goto st211
	st211:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof211
		}
	st_case_211:
//line machine.go:26552
		switch ( m.data)[( m.p)] {
		case 32:
			goto st6
		case 34:
			goto tr85
		case 44:
			goto st6
		case 61:
			goto tr393
		case 92:
			goto st384
		}
		switch {
		case ( m.data)[( m.p)] > 10:
			if 12 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto st6
			}
		case ( m.data)[( m.p)] >= 9:
			goto st6
		}
		goto st211
tr393:
//line machine.rl:70

	key = m.text()

	goto st212
	st212:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof212
		}
	st_case_212:
//line machine.go:26585
		switch ( m.data)[( m.p)] {
		case 34:
			goto tr395
		case 45:
			goto tr396
		case 46:
			goto tr397
		case 48:
			goto tr398
		case 70:
			goto tr400
		case 84:
			goto tr401
		case 92:
			goto st59
		case 102:
			goto tr402
		case 116:
			goto tr403
		}
		if 49 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
			goto tr399
		}
		goto st6
tr395:
//line machine.rl:86

	m.handler.AddString(key, m.text())

	goto st897
	st897:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof897
		}
	st_case_897:
//line machine.go:26621
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1363
		case 13:
			goto tr1363
		case 32:
			goto tr1362
		case 34:
			goto tr23
		case 44:
			goto tr1364
		case 92:
			goto tr24
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1362
		}
		goto tr22
tr1362:
//line machine.rl:17

	m.pb = m.p

	goto st898
tr1414:
//line machine.rl:78

	m.handler.AddFloat(key, m.text())

	goto st898
tr1419:
//line machine.rl:74

	m.handler.AddInt(key, m.text())

	goto st898
tr1422:
//line machine.rl:82

	m.handler.AddBool(key, m.text())

	goto st898
	st898:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof898
		}
	st_case_898:
//line machine.go:26669
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1366
		case 13:
			goto tr1366
		case 32:
			goto st898
		case 34:
			goto tr26
		case 45:
			goto tr1367
		case 92:
			goto st59
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto tr1368
			}
		case ( m.data)[( m.p)] >= 9:
			goto st898
		}
		goto st6
tr1366:
//line machine.rl:21

	yield = true
	{( m.p)++;  m.cs = 899; goto _out }

	goto st899
tr1595:
//line machine.rl:90

	m.handler.SetTimestamp(m.text())

//line machine.rl:21

	yield = true
	{( m.p)++;  m.cs = 899; goto _out }

	goto st899
tr1363:
//line machine.rl:17

	m.pb = m.p

//line machine.rl:21

	yield = true
	{( m.p)++;  m.cs = 899; goto _out }

	goto st899
tr1415:
//line machine.rl:78

	m.handler.AddFloat(key, m.text())

//line machine.rl:21

	yield = true
	{( m.p)++;  m.cs = 899; goto _out }

	goto st899
tr1420:
//line machine.rl:74

	m.handler.AddInt(key, m.text())

//line machine.rl:21

	yield = true
	{( m.p)++;  m.cs = 899; goto _out }

	goto st899
tr1423:
//line machine.rl:82

	m.handler.AddBool(key, m.text())

//line machine.rl:21

	yield = true
	{( m.p)++;  m.cs = 899; goto _out }

	goto st899
	st899:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof899
		}
	st_case_899:
//line machine.go:26760
		switch ( m.data)[( m.p)] {
		case 10:
			goto st899
		case 11:
			goto tr641
		case 13:
			goto st899
		case 32:
			goto st352
		case 34:
			goto tr642
		case 44:
			goto st6
		case 92:
			goto tr643
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto st352
		}
		goto tr639
tr639:
//line machine.rl:17

	m.pb = m.p

	goto st213
	st213:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof213
		}
	st_case_213:
//line machine.go:26792
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr406
		case 13:
			goto st6
		case 32:
			goto tr405
		case 34:
			goto tr407
		case 44:
			goto tr408
		case 92:
			goto st343
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr405
		}
		goto st213
tr405:
//line machine.rl:58

	m.handler.SetMeasurement(m.text())

	goto st214
tr516:
//line machine.rl:66

	m.handler.AddTag(key, m.text())

	goto st214
tr510:
//line machine.rl:17

	m.pb = m.p

//line machine.rl:66

	m.handler.AddTag(key, m.text())

	goto st214
	st214:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof214
		}
	st_case_214:
//line machine.go:26840
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr412
		case 13:
			goto st6
		case 32:
			goto st214
		case 34:
			goto tr413
		case 44:
			goto st6
		case 61:
			goto st6
		case 92:
			goto tr414
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto st214
		}
		goto tr410
tr410:
//line machine.rl:17

	m.pb = m.p

	goto st215
	st215:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof215
		}
	st_case_215:
//line machine.go:26874
		switch ( m.data)[( m.p)] {
		case 32:
			goto st6
		case 34:
			goto tr416
		case 44:
			goto st6
		case 61:
			goto tr417
		case 92:
			goto st259
		}
		switch {
		case ( m.data)[( m.p)] > 10:
			if 12 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto st6
			}
		case ( m.data)[( m.p)] >= 9:
			goto st6
		}
		goto st215
tr416:
//line machine.rl:86

	m.handler.AddString(key, m.text())

	goto st900
tr413:
//line machine.rl:86

	m.handler.AddString(key, m.text())

//line machine.rl:17

	m.pb = m.p

	goto st900
	st900:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof900
		}
	st_case_900:
//line machine.go:26917
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr850
		case 11:
			goto st901
		case 13:
			goto tr850
		case 32:
			goto st492
		case 44:
			goto st28
		case 61:
			goto tr11
		case 92:
			goto st27
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto st492
		}
		goto st3
	st901:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof901
		}
	st_case_901:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr850
		case 11:
			goto st901
		case 13:
			goto tr850
		case 32:
			goto st492
		case 44:
			goto tr419
		case 45:
			goto tr1371
		case 61:
			goto tr11
		case 92:
			goto st27
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto tr1372
			}
		case ( m.data)[( m.p)] >= 9:
			goto st492
		}
		goto st3
tr1371:
//line machine.rl:17

	m.pb = m.p

	goto st216
	st216:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof216
		}
	st_case_216:
//line machine.go:26981
		switch ( m.data)[( m.p)] {
		case 32:
			goto tr419
		case 44:
			goto tr419
		case 61:
			goto tr11
		case 92:
			goto st27
		}
		switch {
		case ( m.data)[( m.p)] < 12:
			if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 10 {
				goto tr419
			}
		case ( m.data)[( m.p)] > 13:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st902
			}
		default:
			goto tr419
		}
		goto st3
tr1372:
//line machine.rl:17

	m.pb = m.p

	goto st902
	st902:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof902
		}
	st_case_902:
//line machine.go:27016
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr860
		case 13:
			goto tr859
		case 32:
			goto tr858
		case 44:
			goto tr419
		case 61:
			goto tr11
		case 92:
			goto st27
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st903
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr858
		}
		goto st3
	st903:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof903
		}
	st_case_903:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr860
		case 13:
			goto tr859
		case 32:
			goto tr858
		case 44:
			goto tr419
		case 61:
			goto tr11
		case 92:
			goto st27
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st904
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr858
		}
		goto st3
	st904:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof904
		}
	st_case_904:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr860
		case 13:
			goto tr859
		case 32:
			goto tr858
		case 44:
			goto tr419
		case 61:
			goto tr11
		case 92:
			goto st27
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st905
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr858
		}
		goto st3
	st905:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof905
		}
	st_case_905:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr860
		case 13:
			goto tr859
		case 32:
			goto tr858
		case 44:
			goto tr419
		case 61:
			goto tr11
		case 92:
			goto st27
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st906
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr858
		}
		goto st3
	st906:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof906
		}
	st_case_906:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr860
		case 13:
			goto tr859
		case 32:
			goto tr858
		case 44:
			goto tr419
		case 61:
			goto tr11
		case 92:
			goto st27
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st907
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr858
		}
		goto st3
	st907:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof907
		}
	st_case_907:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr860
		case 13:
			goto tr859
		case 32:
			goto tr858
		case 44:
			goto tr419
		case 61:
			goto tr11
		case 92:
			goto st27
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st908
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr858
		}
		goto st3
	st908:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof908
		}
	st_case_908:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr860
		case 13:
			goto tr859
		case 32:
			goto tr858
		case 44:
			goto tr419
		case 61:
			goto tr11
		case 92:
			goto st27
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st909
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr858
		}
		goto st3
	st909:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof909
		}
	st_case_909:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr860
		case 13:
			goto tr859
		case 32:
			goto tr858
		case 44:
			goto tr419
		case 61:
			goto tr11
		case 92:
			goto st27
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st910
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr858
		}
		goto st3
	st910:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof910
		}
	st_case_910:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr860
		case 13:
			goto tr859
		case 32:
			goto tr858
		case 44:
			goto tr419
		case 61:
			goto tr11
		case 92:
			goto st27
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st911
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr858
		}
		goto st3
	st911:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof911
		}
	st_case_911:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr860
		case 13:
			goto tr859
		case 32:
			goto tr858
		case 44:
			goto tr419
		case 61:
			goto tr11
		case 92:
			goto st27
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st912
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr858
		}
		goto st3
	st912:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof912
		}
	st_case_912:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr860
		case 13:
			goto tr859
		case 32:
			goto tr858
		case 44:
			goto tr419
		case 61:
			goto tr11
		case 92:
			goto st27
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st913
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr858
		}
		goto st3
	st913:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof913
		}
	st_case_913:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr860
		case 13:
			goto tr859
		case 32:
			goto tr858
		case 44:
			goto tr419
		case 61:
			goto tr11
		case 92:
			goto st27
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st914
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr858
		}
		goto st3
	st914:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof914
		}
	st_case_914:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr860
		case 13:
			goto tr859
		case 32:
			goto tr858
		case 44:
			goto tr419
		case 61:
			goto tr11
		case 92:
			goto st27
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st915
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr858
		}
		goto st3
	st915:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof915
		}
	st_case_915:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr860
		case 13:
			goto tr859
		case 32:
			goto tr858
		case 44:
			goto tr419
		case 61:
			goto tr11
		case 92:
			goto st27
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st916
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr858
		}
		goto st3
	st916:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof916
		}
	st_case_916:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr860
		case 13:
			goto tr859
		case 32:
			goto tr858
		case 44:
			goto tr419
		case 61:
			goto tr11
		case 92:
			goto st27
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st917
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr858
		}
		goto st3
	st917:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof917
		}
	st_case_917:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr860
		case 13:
			goto tr859
		case 32:
			goto tr858
		case 44:
			goto tr419
		case 61:
			goto tr11
		case 92:
			goto st27
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st918
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr858
		}
		goto st3
	st918:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof918
		}
	st_case_918:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr860
		case 13:
			goto tr859
		case 32:
			goto tr858
		case 44:
			goto tr419
		case 61:
			goto tr11
		case 92:
			goto st27
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st919
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr858
		}
		goto st3
	st919:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof919
		}
	st_case_919:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr860
		case 13:
			goto tr859
		case 32:
			goto tr858
		case 44:
			goto tr419
		case 61:
			goto tr11
		case 92:
			goto st27
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st920
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr858
		}
		goto st3
	st920:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof920
		}
	st_case_920:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr860
		case 13:
			goto tr859
		case 32:
			goto tr858
		case 44:
			goto tr419
		case 61:
			goto tr11
		case 92:
			goto st27
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr858
		}
		goto st3
tr417:
//line machine.rl:70

	key = m.text()

	goto st217
	st217:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof217
		}
	st_case_217:
//line machine.go:27588
		switch ( m.data)[( m.p)] {
		case 34:
			goto tr421
		case 45:
			goto tr422
		case 46:
			goto tr423
		case 48:
			goto tr424
		case 70:
			goto tr426
		case 84:
			goto tr427
		case 92:
			goto st59
		case 102:
			goto tr428
		case 116:
			goto tr429
		}
		if 49 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
			goto tr425
		}
		goto st6
tr421:
//line machine.rl:86

	m.handler.AddString(key, m.text())

	goto st921
	st921:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof921
		}
	st_case_921:
//line machine.go:27624
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1392
		case 13:
			goto tr1392
		case 32:
			goto tr1391
		case 34:
			goto tr23
		case 44:
			goto tr1393
		case 92:
			goto tr24
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1391
		}
		goto tr22
tr1391:
//line machine.rl:17

	m.pb = m.p

	goto st922
tr1429:
//line machine.rl:78

	m.handler.AddFloat(key, m.text())

	goto st922
tr1433:
//line machine.rl:74

	m.handler.AddInt(key, m.text())

	goto st922
tr1435:
//line machine.rl:82

	m.handler.AddBool(key, m.text())

	goto st922
	st922:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof922
		}
	st_case_922:
//line machine.go:27672
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr885
		case 13:
			goto tr885
		case 32:
			goto st922
		case 34:
			goto tr26
		case 45:
			goto tr1395
		case 92:
			goto st59
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto tr887
			}
		case ( m.data)[( m.p)] >= 9:
			goto st922
		}
		goto st6
tr1395:
//line machine.rl:17

	m.pb = m.p

	goto st218
	st218:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof218
		}
	st_case_218:
//line machine.go:27707
		switch ( m.data)[( m.p)] {
		case 34:
			goto tr26
		case 92:
			goto st59
		}
		if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
			goto st832
		}
		goto st6
tr1393:
//line machine.rl:17

	m.pb = m.p

	goto st219
tr1430:
//line machine.rl:78

	m.handler.AddFloat(key, m.text())

	goto st219
tr1434:
//line machine.rl:74

	m.handler.AddInt(key, m.text())

	goto st219
tr1436:
//line machine.rl:82

	m.handler.AddBool(key, m.text())

	goto st219
	st219:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof219
		}
	st_case_219:
//line machine.go:27747
		switch ( m.data)[( m.p)] {
		case 32:
			goto st6
		case 34:
			goto tr82
		case 44:
			goto st6
		case 61:
			goto st6
		case 92:
			goto tr431
		}
		switch {
		case ( m.data)[( m.p)] > 10:
			if 12 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto st6
			}
		case ( m.data)[( m.p)] >= 9:
			goto st6
		}
		goto tr430
tr430:
//line machine.rl:17

	m.pb = m.p

	goto st220
	st220:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof220
		}
	st_case_220:
//line machine.go:27780
		switch ( m.data)[( m.p)] {
		case 32:
			goto st6
		case 34:
			goto tr85
		case 44:
			goto st6
		case 61:
			goto tr433
		case 92:
			goto st258
		}
		switch {
		case ( m.data)[( m.p)] > 10:
			if 12 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto st6
			}
		case ( m.data)[( m.p)] >= 9:
			goto st6
		}
		goto st220
tr433:
//line machine.rl:70

	key = m.text()

	goto st221
	st221:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof221
		}
	st_case_221:
//line machine.go:27813
		switch ( m.data)[( m.p)] {
		case 34:
			goto tr435
		case 45:
			goto tr436
		case 46:
			goto tr437
		case 48:
			goto tr438
		case 70:
			goto tr440
		case 84:
			goto tr441
		case 92:
			goto st59
		case 102:
			goto tr442
		case 116:
			goto tr443
		}
		if 49 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
			goto tr439
		}
		goto st6
tr435:
//line machine.rl:86

	m.handler.AddString(key, m.text())

	goto st923
	st923:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof923
		}
	st_case_923:
//line machine.go:27849
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1363
		case 13:
			goto tr1363
		case 32:
			goto tr1362
		case 34:
			goto tr23
		case 44:
			goto tr1396
		case 92:
			goto tr24
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1362
		}
		goto tr22
tr1396:
//line machine.rl:17

	m.pb = m.p

	goto st222
tr1416:
//line machine.rl:78

	m.handler.AddFloat(key, m.text())

	goto st222
tr1421:
//line machine.rl:74

	m.handler.AddInt(key, m.text())

	goto st222
tr1424:
//line machine.rl:82

	m.handler.AddBool(key, m.text())

	goto st222
	st222:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof222
		}
	st_case_222:
//line machine.go:27897
		switch ( m.data)[( m.p)] {
		case 32:
			goto st6
		case 34:
			goto tr413
		case 44:
			goto st6
		case 61:
			goto st6
		case 92:
			goto tr445
		}
		switch {
		case ( m.data)[( m.p)] > 10:
			if 12 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto st6
			}
		case ( m.data)[( m.p)] >= 9:
			goto st6
		}
		goto tr444
tr444:
//line machine.rl:17

	m.pb = m.p

	goto st223
	st223:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof223
		}
	st_case_223:
//line machine.go:27930
		switch ( m.data)[( m.p)] {
		case 32:
			goto st6
		case 34:
			goto tr416
		case 44:
			goto st6
		case 61:
			goto tr447
		case 92:
			goto st248
		}
		switch {
		case ( m.data)[( m.p)] > 10:
			if 12 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto st6
			}
		case ( m.data)[( m.p)] >= 9:
			goto st6
		}
		goto st223
tr447:
//line machine.rl:70

	key = m.text()

	goto st224
	st224:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof224
		}
	st_case_224:
//line machine.go:27963
		switch ( m.data)[( m.p)] {
		case 34:
			goto tr449
		case 45:
			goto tr422
		case 46:
			goto tr423
		case 48:
			goto tr424
		case 70:
			goto tr426
		case 84:
			goto tr427
		case 92:
			goto st59
		case 102:
			goto tr428
		case 116:
			goto tr429
		}
		if 49 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
			goto tr425
		}
		goto st6
tr449:
//line machine.rl:86

	m.handler.AddString(key, m.text())

	goto st924
	st924:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof924
		}
	st_case_924:
//line machine.go:27999
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1398
		case 13:
			goto tr1398
		case 32:
			goto tr1397
		case 34:
			goto tr23
		case 44:
			goto tr1399
		case 92:
			goto tr24
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1397
		}
		goto tr22
tr1397:
//line machine.rl:17

	m.pb = m.p

	goto st925
tr1402:
//line machine.rl:78

	m.handler.AddFloat(key, m.text())

	goto st925
tr1406:
//line machine.rl:74

	m.handler.AddInt(key, m.text())

	goto st925
tr1408:
//line machine.rl:82

	m.handler.AddBool(key, m.text())

	goto st925
	st925:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof925
		}
	st_case_925:
//line machine.go:28047
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr927
		case 13:
			goto tr927
		case 32:
			goto st925
		case 34:
			goto tr26
		case 45:
			goto tr1401
		case 92:
			goto st59
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto tr929
			}
		case ( m.data)[( m.p)] >= 9:
			goto st925
		}
		goto st6
tr1401:
//line machine.rl:17

	m.pb = m.p

	goto st225
	st225:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof225
		}
	st_case_225:
//line machine.go:28082
		switch ( m.data)[( m.p)] {
		case 34:
			goto tr26
		case 92:
			goto st59
		}
		if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
			goto st795
		}
		goto st6
tr1399:
//line machine.rl:17

	m.pb = m.p

	goto st226
tr1403:
//line machine.rl:78

	m.handler.AddFloat(key, m.text())

	goto st226
tr1407:
//line machine.rl:74

	m.handler.AddInt(key, m.text())

	goto st226
tr1409:
//line machine.rl:82

	m.handler.AddBool(key, m.text())

	goto st226
	st226:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof226
		}
	st_case_226:
//line machine.go:28122
		switch ( m.data)[( m.p)] {
		case 32:
			goto st6
		case 34:
			goto tr82
		case 44:
			goto st6
		case 61:
			goto st6
		case 92:
			goto tr451
		}
		switch {
		case ( m.data)[( m.p)] > 10:
			if 12 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto st6
			}
		case ( m.data)[( m.p)] >= 9:
			goto st6
		}
		goto tr450
tr450:
//line machine.rl:17

	m.pb = m.p

	goto st227
	st227:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof227
		}
	st_case_227:
//line machine.go:28155
		switch ( m.data)[( m.p)] {
		case 32:
			goto st6
		case 34:
			goto tr85
		case 44:
			goto st6
		case 61:
			goto tr453
		case 92:
			goto st238
		}
		switch {
		case ( m.data)[( m.p)] > 10:
			if 12 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto st6
			}
		case ( m.data)[( m.p)] >= 9:
			goto st6
		}
		goto st227
tr453:
//line machine.rl:70

	key = m.text()

	goto st228
	st228:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof228
		}
	st_case_228:
//line machine.go:28188
		switch ( m.data)[( m.p)] {
		case 34:
			goto tr435
		case 45:
			goto tr455
		case 46:
			goto tr456
		case 48:
			goto tr457
		case 70:
			goto tr459
		case 84:
			goto tr460
		case 92:
			goto st59
		case 102:
			goto tr461
		case 116:
			goto tr462
		}
		if 49 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
			goto tr458
		}
		goto st6
tr455:
//line machine.rl:17

	m.pb = m.p

	goto st229
	st229:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof229
		}
	st_case_229:
//line machine.go:28224
		switch ( m.data)[( m.p)] {
		case 34:
			goto tr26
		case 48:
			goto st926
		case 92:
			goto st59
		}
		if 49 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
			goto st929
		}
		goto st6
tr457:
//line machine.rl:17

	m.pb = m.p

	goto st926
	st926:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof926
		}
	st_case_926:
//line machine.go:28248
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr922
		case 13:
			goto tr922
		case 32:
			goto tr1402
		case 34:
			goto tr26
		case 44:
			goto tr1403
		case 46:
			goto st927
		case 92:
			goto st59
		case 105:
			goto st928
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1402
		}
		goto st6
tr456:
//line machine.rl:17

	m.pb = m.p

	goto st927
	st927:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof927
		}
	st_case_927:
//line machine.go:28282
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr922
		case 13:
			goto tr922
		case 32:
			goto tr1402
		case 34:
			goto tr26
		case 44:
			goto tr1403
		case 92:
			goto st59
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st927
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1402
		}
		goto st6
	st928:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof928
		}
	st_case_928:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1119
		case 13:
			goto tr1119
		case 32:
			goto tr1406
		case 34:
			goto tr26
		case 44:
			goto tr1407
		case 92:
			goto st59
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1406
		}
		goto st6
tr458:
//line machine.rl:17

	m.pb = m.p

	goto st929
	st929:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof929
		}
	st_case_929:
//line machine.go:28340
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr922
		case 13:
			goto tr922
		case 32:
			goto tr1402
		case 34:
			goto tr26
		case 44:
			goto tr1403
		case 46:
			goto st927
		case 92:
			goto st59
		case 105:
			goto st928
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st929
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1402
		}
		goto st6
tr459:
//line machine.rl:17

	m.pb = m.p

	goto st930
	st930:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof930
		}
	st_case_930:
//line machine.go:28379
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1123
		case 13:
			goto tr1123
		case 32:
			goto tr1408
		case 34:
			goto tr26
		case 44:
			goto tr1409
		case 65:
			goto st230
		case 92:
			goto st59
		case 97:
			goto st233
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1408
		}
		goto st6
	st230:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof230
		}
	st_case_230:
		switch ( m.data)[( m.p)] {
		case 34:
			goto tr26
		case 76:
			goto st231
		case 92:
			goto st59
		}
		goto st6
	st231:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof231
		}
	st_case_231:
		switch ( m.data)[( m.p)] {
		case 34:
			goto tr26
		case 83:
			goto st232
		case 92:
			goto st59
		}
		goto st6
	st232:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof232
		}
	st_case_232:
		switch ( m.data)[( m.p)] {
		case 34:
			goto tr26
		case 69:
			goto st931
		case 92:
			goto st59
		}
		goto st6
	st931:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof931
		}
	st_case_931:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1123
		case 13:
			goto tr1123
		case 32:
			goto tr1408
		case 34:
			goto tr26
		case 44:
			goto tr1409
		case 92:
			goto st59
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1408
		}
		goto st6
	st233:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof233
		}
	st_case_233:
		switch ( m.data)[( m.p)] {
		case 34:
			goto tr26
		case 92:
			goto st59
		case 108:
			goto st234
		}
		goto st6
	st234:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof234
		}
	st_case_234:
		switch ( m.data)[( m.p)] {
		case 34:
			goto tr26
		case 92:
			goto st59
		case 115:
			goto st235
		}
		goto st6
	st235:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof235
		}
	st_case_235:
		switch ( m.data)[( m.p)] {
		case 34:
			goto tr26
		case 92:
			goto st59
		case 101:
			goto st931
		}
		goto st6
tr460:
//line machine.rl:17

	m.pb = m.p

	goto st932
	st932:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof932
		}
	st_case_932:
//line machine.go:28520
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1123
		case 13:
			goto tr1123
		case 32:
			goto tr1408
		case 34:
			goto tr26
		case 44:
			goto tr1409
		case 82:
			goto st236
		case 92:
			goto st59
		case 114:
			goto st237
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1408
		}
		goto st6
	st236:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof236
		}
	st_case_236:
		switch ( m.data)[( m.p)] {
		case 34:
			goto tr26
		case 85:
			goto st232
		case 92:
			goto st59
		}
		goto st6
	st237:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof237
		}
	st_case_237:
		switch ( m.data)[( m.p)] {
		case 34:
			goto tr26
		case 92:
			goto st59
		case 117:
			goto st235
		}
		goto st6
tr461:
//line machine.rl:17

	m.pb = m.p

	goto st933
	st933:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof933
		}
	st_case_933:
//line machine.go:28582
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1123
		case 13:
			goto tr1123
		case 32:
			goto tr1408
		case 34:
			goto tr26
		case 44:
			goto tr1409
		case 92:
			goto st59
		case 97:
			goto st233
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1408
		}
		goto st6
tr462:
//line machine.rl:17

	m.pb = m.p

	goto st934
	st934:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof934
		}
	st_case_934:
//line machine.go:28614
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1123
		case 13:
			goto tr1123
		case 32:
			goto tr1408
		case 34:
			goto tr26
		case 44:
			goto tr1409
		case 92:
			goto st59
		case 114:
			goto st237
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1408
		}
		goto st6
tr451:
//line machine.rl:17

	m.pb = m.p

	goto st238
	st238:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof238
		}
	st_case_238:
//line machine.go:28646
		switch ( m.data)[( m.p)] {
		case 34:
			goto st227
		case 92:
			goto st227
		}
		switch {
		case ( m.data)[( m.p)] > 10:
			if 12 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto tr2
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr2
		}
		goto st3
tr422:
//line machine.rl:17

	m.pb = m.p

	goto st239
	st239:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof239
		}
	st_case_239:
//line machine.go:28673
		switch ( m.data)[( m.p)] {
		case 34:
			goto tr26
		case 48:
			goto st935
		case 92:
			goto st59
		}
		if 49 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
			goto st938
		}
		goto st6
tr424:
//line machine.rl:17

	m.pb = m.p

	goto st935
	st935:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof935
		}
	st_case_935:
//line machine.go:28697
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1415
		case 13:
			goto tr1415
		case 32:
			goto tr1414
		case 34:
			goto tr26
		case 44:
			goto tr1416
		case 46:
			goto st936
		case 92:
			goto st59
		case 105:
			goto st937
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1414
		}
		goto st6
tr423:
//line machine.rl:17

	m.pb = m.p

	goto st936
	st936:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof936
		}
	st_case_936:
//line machine.go:28731
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1415
		case 13:
			goto tr1415
		case 32:
			goto tr1414
		case 34:
			goto tr26
		case 44:
			goto tr1416
		case 92:
			goto st59
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st936
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1414
		}
		goto st6
	st937:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof937
		}
	st_case_937:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1420
		case 13:
			goto tr1420
		case 32:
			goto tr1419
		case 34:
			goto tr26
		case 44:
			goto tr1421
		case 92:
			goto st59
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1419
		}
		goto st6
tr425:
//line machine.rl:17

	m.pb = m.p

	goto st938
	st938:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof938
		}
	st_case_938:
//line machine.go:28789
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1415
		case 13:
			goto tr1415
		case 32:
			goto tr1414
		case 34:
			goto tr26
		case 44:
			goto tr1416
		case 46:
			goto st936
		case 92:
			goto st59
		case 105:
			goto st937
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st938
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1414
		}
		goto st6
tr426:
//line machine.rl:17

	m.pb = m.p

	goto st939
	st939:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof939
		}
	st_case_939:
//line machine.go:28828
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1423
		case 13:
			goto tr1423
		case 32:
			goto tr1422
		case 34:
			goto tr26
		case 44:
			goto tr1424
		case 65:
			goto st240
		case 92:
			goto st59
		case 97:
			goto st243
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1422
		}
		goto st6
	st240:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof240
		}
	st_case_240:
		switch ( m.data)[( m.p)] {
		case 34:
			goto tr26
		case 76:
			goto st241
		case 92:
			goto st59
		}
		goto st6
	st241:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof241
		}
	st_case_241:
		switch ( m.data)[( m.p)] {
		case 34:
			goto tr26
		case 83:
			goto st242
		case 92:
			goto st59
		}
		goto st6
	st242:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof242
		}
	st_case_242:
		switch ( m.data)[( m.p)] {
		case 34:
			goto tr26
		case 69:
			goto st940
		case 92:
			goto st59
		}
		goto st6
	st940:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof940
		}
	st_case_940:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1423
		case 13:
			goto tr1423
		case 32:
			goto tr1422
		case 34:
			goto tr26
		case 44:
			goto tr1424
		case 92:
			goto st59
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1422
		}
		goto st6
	st243:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof243
		}
	st_case_243:
		switch ( m.data)[( m.p)] {
		case 34:
			goto tr26
		case 92:
			goto st59
		case 108:
			goto st244
		}
		goto st6
	st244:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof244
		}
	st_case_244:
		switch ( m.data)[( m.p)] {
		case 34:
			goto tr26
		case 92:
			goto st59
		case 115:
			goto st245
		}
		goto st6
	st245:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof245
		}
	st_case_245:
		switch ( m.data)[( m.p)] {
		case 34:
			goto tr26
		case 92:
			goto st59
		case 101:
			goto st940
		}
		goto st6
tr427:
//line machine.rl:17

	m.pb = m.p

	goto st941
	st941:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof941
		}
	st_case_941:
//line machine.go:28969
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1423
		case 13:
			goto tr1423
		case 32:
			goto tr1422
		case 34:
			goto tr26
		case 44:
			goto tr1424
		case 82:
			goto st246
		case 92:
			goto st59
		case 114:
			goto st247
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1422
		}
		goto st6
	st246:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof246
		}
	st_case_246:
		switch ( m.data)[( m.p)] {
		case 34:
			goto tr26
		case 85:
			goto st242
		case 92:
			goto st59
		}
		goto st6
	st247:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof247
		}
	st_case_247:
		switch ( m.data)[( m.p)] {
		case 34:
			goto tr26
		case 92:
			goto st59
		case 117:
			goto st245
		}
		goto st6
tr428:
//line machine.rl:17

	m.pb = m.p

	goto st942
	st942:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof942
		}
	st_case_942:
//line machine.go:29031
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1423
		case 13:
			goto tr1423
		case 32:
			goto tr1422
		case 34:
			goto tr26
		case 44:
			goto tr1424
		case 92:
			goto st59
		case 97:
			goto st243
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1422
		}
		goto st6
tr429:
//line machine.rl:17

	m.pb = m.p

	goto st943
	st943:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof943
		}
	st_case_943:
//line machine.go:29063
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1423
		case 13:
			goto tr1423
		case 32:
			goto tr1422
		case 34:
			goto tr26
		case 44:
			goto tr1424
		case 92:
			goto st59
		case 114:
			goto st247
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1422
		}
		goto st6
tr445:
//line machine.rl:17

	m.pb = m.p

	goto st248
	st248:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof248
		}
	st_case_248:
//line machine.go:29095
		switch ( m.data)[( m.p)] {
		case 34:
			goto st223
		case 92:
			goto st223
		}
		switch {
		case ( m.data)[( m.p)] > 10:
			if 12 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto tr2
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr2
		}
		goto st3
tr436:
//line machine.rl:17

	m.pb = m.p

	goto st249
	st249:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof249
		}
	st_case_249:
//line machine.go:29122
		switch ( m.data)[( m.p)] {
		case 34:
			goto tr26
		case 48:
			goto st944
		case 92:
			goto st59
		}
		if 49 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
			goto st947
		}
		goto st6
tr438:
//line machine.rl:17

	m.pb = m.p

	goto st944
	st944:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof944
		}
	st_case_944:
//line machine.go:29146
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1026
		case 13:
			goto tr1026
		case 32:
			goto tr1429
		case 34:
			goto tr26
		case 44:
			goto tr1430
		case 46:
			goto st945
		case 92:
			goto st59
		case 105:
			goto st946
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1429
		}
		goto st6
tr437:
//line machine.rl:17

	m.pb = m.p

	goto st945
	st945:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof945
		}
	st_case_945:
//line machine.go:29180
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1026
		case 13:
			goto tr1026
		case 32:
			goto tr1429
		case 34:
			goto tr26
		case 44:
			goto tr1430
		case 92:
			goto st59
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st945
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1429
		}
		goto st6
	st946:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof946
		}
	st_case_946:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1162
		case 13:
			goto tr1162
		case 32:
			goto tr1433
		case 34:
			goto tr26
		case 44:
			goto tr1434
		case 92:
			goto st59
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1433
		}
		goto st6
tr439:
//line machine.rl:17

	m.pb = m.p

	goto st947
	st947:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof947
		}
	st_case_947:
//line machine.go:29238
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1026
		case 13:
			goto tr1026
		case 32:
			goto tr1429
		case 34:
			goto tr26
		case 44:
			goto tr1430
		case 46:
			goto st945
		case 92:
			goto st59
		case 105:
			goto st946
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st947
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1429
		}
		goto st6
tr440:
//line machine.rl:17

	m.pb = m.p

	goto st948
	st948:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof948
		}
	st_case_948:
//line machine.go:29277
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1166
		case 13:
			goto tr1166
		case 32:
			goto tr1435
		case 34:
			goto tr26
		case 44:
			goto tr1436
		case 65:
			goto st250
		case 92:
			goto st59
		case 97:
			goto st253
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1435
		}
		goto st6
	st250:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof250
		}
	st_case_250:
		switch ( m.data)[( m.p)] {
		case 34:
			goto tr26
		case 76:
			goto st251
		case 92:
			goto st59
		}
		goto st6
	st251:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof251
		}
	st_case_251:
		switch ( m.data)[( m.p)] {
		case 34:
			goto tr26
		case 83:
			goto st252
		case 92:
			goto st59
		}
		goto st6
	st252:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof252
		}
	st_case_252:
		switch ( m.data)[( m.p)] {
		case 34:
			goto tr26
		case 69:
			goto st949
		case 92:
			goto st59
		}
		goto st6
	st949:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof949
		}
	st_case_949:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1166
		case 13:
			goto tr1166
		case 32:
			goto tr1435
		case 34:
			goto tr26
		case 44:
			goto tr1436
		case 92:
			goto st59
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1435
		}
		goto st6
	st253:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof253
		}
	st_case_253:
		switch ( m.data)[( m.p)] {
		case 34:
			goto tr26
		case 92:
			goto st59
		case 108:
			goto st254
		}
		goto st6
	st254:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof254
		}
	st_case_254:
		switch ( m.data)[( m.p)] {
		case 34:
			goto tr26
		case 92:
			goto st59
		case 115:
			goto st255
		}
		goto st6
	st255:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof255
		}
	st_case_255:
		switch ( m.data)[( m.p)] {
		case 34:
			goto tr26
		case 92:
			goto st59
		case 101:
			goto st949
		}
		goto st6
tr441:
//line machine.rl:17

	m.pb = m.p

	goto st950
	st950:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof950
		}
	st_case_950:
//line machine.go:29418
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1166
		case 13:
			goto tr1166
		case 32:
			goto tr1435
		case 34:
			goto tr26
		case 44:
			goto tr1436
		case 82:
			goto st256
		case 92:
			goto st59
		case 114:
			goto st257
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1435
		}
		goto st6
	st256:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof256
		}
	st_case_256:
		switch ( m.data)[( m.p)] {
		case 34:
			goto tr26
		case 85:
			goto st252
		case 92:
			goto st59
		}
		goto st6
	st257:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof257
		}
	st_case_257:
		switch ( m.data)[( m.p)] {
		case 34:
			goto tr26
		case 92:
			goto st59
		case 117:
			goto st255
		}
		goto st6
tr442:
//line machine.rl:17

	m.pb = m.p

	goto st951
	st951:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof951
		}
	st_case_951:
//line machine.go:29480
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1166
		case 13:
			goto tr1166
		case 32:
			goto tr1435
		case 34:
			goto tr26
		case 44:
			goto tr1436
		case 92:
			goto st59
		case 97:
			goto st253
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1435
		}
		goto st6
tr443:
//line machine.rl:17

	m.pb = m.p

	goto st952
	st952:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof952
		}
	st_case_952:
//line machine.go:29512
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1166
		case 13:
			goto tr1166
		case 32:
			goto tr1435
		case 34:
			goto tr26
		case 44:
			goto tr1436
		case 92:
			goto st59
		case 114:
			goto st257
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1435
		}
		goto st6
tr431:
//line machine.rl:17

	m.pb = m.p

	goto st258
	st258:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof258
		}
	st_case_258:
//line machine.go:29544
		switch ( m.data)[( m.p)] {
		case 34:
			goto st220
		case 92:
			goto st220
		}
		switch {
		case ( m.data)[( m.p)] > 10:
			if 12 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto tr2
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr2
		}
		goto st3
tr414:
//line machine.rl:17

	m.pb = m.p

	goto st259
	st259:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof259
		}
	st_case_259:
//line machine.go:29571
		switch ( m.data)[( m.p)] {
		case 34:
			goto st215
		case 92:
			goto st215
		}
		switch {
		case ( m.data)[( m.p)] > 10:
			if 12 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto tr2
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr2
		}
		goto st3
tr412:
//line machine.rl:17

	m.pb = m.p

	goto st260
	st260:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof260
		}
	st_case_260:
//line machine.go:29598
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr412
		case 13:
			goto st6
		case 32:
			goto st214
		case 34:
			goto tr413
		case 44:
			goto st6
		case 61:
			goto tr417
		case 92:
			goto tr414
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto st214
		}
		goto tr410
tr406:
//line machine.rl:58

	m.handler.SetMeasurement(m.text())

	goto st261
	st261:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof261
		}
	st_case_261:
//line machine.go:29632
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr485
		case 13:
			goto st6
		case 32:
			goto tr405
		case 34:
			goto tr486
		case 44:
			goto tr408
		case 61:
			goto st213
		case 92:
			goto tr487
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr405
		}
		goto tr484
tr484:
//line machine.rl:17

	m.pb = m.p

	goto st262
	st262:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof262
		}
	st_case_262:
//line machine.go:29666
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr489
		case 13:
			goto st6
		case 32:
			goto tr405
		case 34:
			goto tr490
		case 44:
			goto tr408
		case 61:
			goto tr491
		case 92:
			goto st342
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr405
		}
		goto st262
tr489:
//line machine.rl:58

	m.handler.SetMeasurement(m.text())

	goto st263
tr485:
//line machine.rl:58

	m.handler.SetMeasurement(m.text())

//line machine.rl:17

	m.pb = m.p

	goto st263
	st263:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof263
		}
	st_case_263:
//line machine.go:29710
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr485
		case 13:
			goto st6
		case 32:
			goto tr405
		case 34:
			goto tr486
		case 44:
			goto tr408
		case 61:
			goto tr491
		case 92:
			goto tr487
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr405
		}
		goto tr484
tr490:
//line machine.rl:86

	m.handler.AddString(key, m.text())

	goto st953
tr486:
//line machine.rl:86

	m.handler.AddString(key, m.text())

//line machine.rl:17

	m.pb = m.p

	goto st953
	st953:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof953
		}
	st_case_953:
//line machine.go:29754
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr850
		case 11:
			goto tr1442
		case 13:
			goto tr850
		case 32:
			goto tr1441
		case 44:
			goto tr1443
		case 61:
			goto tr36
		case 92:
			goto st36
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1441
		}
		goto st9
tr1441:
//line machine.rl:58

	m.handler.SetMeasurement(m.text())

	goto st954
tr1470:
//line machine.rl:66

	m.handler.AddTag(key, m.text())

	goto st954
tr1518:
//line machine.rl:78

	m.handler.AddFloat(key, m.text())

//line machine.rl:66

	m.handler.AddTag(key, m.text())

	goto st954
tr1523:
//line machine.rl:74

	m.handler.AddInt(key, m.text())

//line machine.rl:66

	m.handler.AddTag(key, m.text())

	goto st954
tr1526:
//line machine.rl:82

	m.handler.AddBool(key, m.text())

//line machine.rl:66

	m.handler.AddTag(key, m.text())

	goto st954
	st954:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof954
		}
	st_case_954:
//line machine.go:29822
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr850
		case 11:
			goto tr1445
		case 13:
			goto tr850
		case 32:
			goto st954
		case 44:
			goto tr419
		case 45:
			goto tr1371
		case 61:
			goto tr419
		case 92:
			goto tr9
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto tr1372
			}
		case ( m.data)[( m.p)] >= 9:
			goto st954
		}
		goto tr6
tr1445:
//line machine.rl:17

	m.pb = m.p

	goto st955
	st955:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof955
		}
	st_case_955:
//line machine.go:29861
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr850
		case 11:
			goto tr1445
		case 13:
			goto tr850
		case 32:
			goto st954
		case 44:
			goto tr419
		case 45:
			goto tr1371
		case 61:
			goto tr11
		case 92:
			goto tr9
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto tr1372
			}
		case ( m.data)[( m.p)] >= 9:
			goto st954
		}
		goto tr6
tr1442:
//line machine.rl:58

	m.handler.SetMeasurement(m.text())

	goto st956
tr1446:
//line machine.rl:58

	m.handler.SetMeasurement(m.text())

//line machine.rl:17

	m.pb = m.p

	goto st956
	st956:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof956
		}
	st_case_956:
//line machine.go:29910
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr850
		case 11:
			goto tr1446
		case 13:
			goto tr850
		case 32:
			goto tr1441
		case 44:
			goto tr4
		case 45:
			goto tr1447
		case 61:
			goto tr36
		case 92:
			goto tr33
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto tr1448
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1441
		}
		goto tr31
tr1447:
//line machine.rl:17

	m.pb = m.p

	goto st264
	st264:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof264
		}
	st_case_264:
//line machine.go:29949
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr419
		case 11:
			goto tr35
		case 13:
			goto tr419
		case 32:
			goto tr1
		case 44:
			goto tr4
		case 61:
			goto tr36
		case 92:
			goto st36
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st957
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1
		}
		goto st9
tr1448:
//line machine.rl:17

	m.pb = m.p

	goto st957
	st957:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof957
		}
	st_case_957:
//line machine.go:29986
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr1450
		case 13:
			goto tr859
		case 32:
			goto tr1449
		case 44:
			goto tr4
		case 61:
			goto tr36
		case 92:
			goto st36
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st959
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1449
		}
		goto st9
tr1452:
//line machine.rl:58

	m.handler.SetMeasurement(m.text())

//line machine.rl:17

	m.pb = m.p

	goto st958
tr1450:
//line machine.rl:90

	m.handler.SetTimestamp(m.text())

//line machine.rl:58

	m.handler.SetMeasurement(m.text())

	goto st958
	st958:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof958
		}
	st_case_958:
//line machine.go:30037
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr850
		case 11:
			goto tr1452
		case 13:
			goto tr850
		case 32:
			goto tr902
		case 44:
			goto tr4
		case 61:
			goto tr36
		case 92:
			goto tr33
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr902
		}
		goto tr31
	st959:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof959
		}
	st_case_959:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr1450
		case 13:
			goto tr859
		case 32:
			goto tr1449
		case 44:
			goto tr4
		case 61:
			goto tr36
		case 92:
			goto st36
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st960
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1449
		}
		goto st9
	st960:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof960
		}
	st_case_960:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr1450
		case 13:
			goto tr859
		case 32:
			goto tr1449
		case 44:
			goto tr4
		case 61:
			goto tr36
		case 92:
			goto st36
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st961
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1449
		}
		goto st9
	st961:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof961
		}
	st_case_961:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr1450
		case 13:
			goto tr859
		case 32:
			goto tr1449
		case 44:
			goto tr4
		case 61:
			goto tr36
		case 92:
			goto st36
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st962
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1449
		}
		goto st9
	st962:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof962
		}
	st_case_962:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr1450
		case 13:
			goto tr859
		case 32:
			goto tr1449
		case 44:
			goto tr4
		case 61:
			goto tr36
		case 92:
			goto st36
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st963
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1449
		}
		goto st9
	st963:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof963
		}
	st_case_963:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr1450
		case 13:
			goto tr859
		case 32:
			goto tr1449
		case 44:
			goto tr4
		case 61:
			goto tr36
		case 92:
			goto st36
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st964
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1449
		}
		goto st9
	st964:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof964
		}
	st_case_964:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr1450
		case 13:
			goto tr859
		case 32:
			goto tr1449
		case 44:
			goto tr4
		case 61:
			goto tr36
		case 92:
			goto st36
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st965
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1449
		}
		goto st9
	st965:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof965
		}
	st_case_965:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr1450
		case 13:
			goto tr859
		case 32:
			goto tr1449
		case 44:
			goto tr4
		case 61:
			goto tr36
		case 92:
			goto st36
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st966
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1449
		}
		goto st9
	st966:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof966
		}
	st_case_966:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr1450
		case 13:
			goto tr859
		case 32:
			goto tr1449
		case 44:
			goto tr4
		case 61:
			goto tr36
		case 92:
			goto st36
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st967
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1449
		}
		goto st9
	st967:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof967
		}
	st_case_967:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr1450
		case 13:
			goto tr859
		case 32:
			goto tr1449
		case 44:
			goto tr4
		case 61:
			goto tr36
		case 92:
			goto st36
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st968
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1449
		}
		goto st9
	st968:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof968
		}
	st_case_968:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr1450
		case 13:
			goto tr859
		case 32:
			goto tr1449
		case 44:
			goto tr4
		case 61:
			goto tr36
		case 92:
			goto st36
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st969
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1449
		}
		goto st9
	st969:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof969
		}
	st_case_969:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr1450
		case 13:
			goto tr859
		case 32:
			goto tr1449
		case 44:
			goto tr4
		case 61:
			goto tr36
		case 92:
			goto st36
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st970
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1449
		}
		goto st9
	st970:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof970
		}
	st_case_970:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr1450
		case 13:
			goto tr859
		case 32:
			goto tr1449
		case 44:
			goto tr4
		case 61:
			goto tr36
		case 92:
			goto st36
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st971
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1449
		}
		goto st9
	st971:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof971
		}
	st_case_971:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr1450
		case 13:
			goto tr859
		case 32:
			goto tr1449
		case 44:
			goto tr4
		case 61:
			goto tr36
		case 92:
			goto st36
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st972
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1449
		}
		goto st9
	st972:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof972
		}
	st_case_972:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr1450
		case 13:
			goto tr859
		case 32:
			goto tr1449
		case 44:
			goto tr4
		case 61:
			goto tr36
		case 92:
			goto st36
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st973
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1449
		}
		goto st9
	st973:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof973
		}
	st_case_973:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr1450
		case 13:
			goto tr859
		case 32:
			goto tr1449
		case 44:
			goto tr4
		case 61:
			goto tr36
		case 92:
			goto st36
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st974
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1449
		}
		goto st9
	st974:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof974
		}
	st_case_974:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr1450
		case 13:
			goto tr859
		case 32:
			goto tr1449
		case 44:
			goto tr4
		case 61:
			goto tr36
		case 92:
			goto st36
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st975
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1449
		}
		goto st9
	st975:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof975
		}
	st_case_975:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr1450
		case 13:
			goto tr859
		case 32:
			goto tr1449
		case 44:
			goto tr4
		case 61:
			goto tr36
		case 92:
			goto st36
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st976
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1449
		}
		goto st9
	st976:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof976
		}
	st_case_976:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr1450
		case 13:
			goto tr859
		case 32:
			goto tr1449
		case 44:
			goto tr4
		case 61:
			goto tr36
		case 92:
			goto st36
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1449
		}
		goto st9
tr1443:
//line machine.rl:58

	m.handler.SetMeasurement(m.text())

	goto st265
tr1472:
//line machine.rl:66

	m.handler.AddTag(key, m.text())

	goto st265
tr1520:
//line machine.rl:78

	m.handler.AddFloat(key, m.text())

//line machine.rl:66

	m.handler.AddTag(key, m.text())

	goto st265
tr1525:
//line machine.rl:74

	m.handler.AddInt(key, m.text())

//line machine.rl:66

	m.handler.AddTag(key, m.text())

	goto st265
tr1528:
//line machine.rl:82

	m.handler.AddBool(key, m.text())

//line machine.rl:66

	m.handler.AddTag(key, m.text())

	goto st265
	st265:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof265
		}
	st_case_265:
//line machine.go:30640
		switch ( m.data)[( m.p)] {
		case 32:
			goto tr495
		case 44:
			goto tr495
		case 61:
			goto tr495
		case 92:
			goto tr496
		}
		switch {
		case ( m.data)[( m.p)] > 10:
			if 12 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto tr495
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr495
		}
		goto tr494
tr494:
//line machine.rl:17

	m.pb = m.p

	goto st266
	st266:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof266
		}
	st_case_266:
//line machine.go:30671
		switch ( m.data)[( m.p)] {
		case 32:
			goto tr495
		case 44:
			goto tr495
		case 61:
			goto tr498
		case 92:
			goto st290
		}
		switch {
		case ( m.data)[( m.p)] > 10:
			if 12 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto tr495
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr495
		}
		goto st266
tr498:
//line machine.rl:70

	key = m.text()

//line machine.rl:62

	key = m.text()

	goto st267
	st267:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof267
		}
	st_case_267:
//line machine.go:30706
		switch ( m.data)[( m.p)] {
		case 32:
			goto tr495
		case 34:
			goto tr500
		case 44:
			goto tr495
		case 45:
			goto tr501
		case 46:
			goto tr502
		case 48:
			goto tr503
		case 61:
			goto tr495
		case 70:
			goto tr505
		case 84:
			goto tr506
		case 92:
			goto tr45
		case 102:
			goto tr507
		case 116:
			goto tr508
		}
		switch {
		case ( m.data)[( m.p)] < 12:
			if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 10 {
				goto tr495
			}
		case ( m.data)[( m.p)] > 13:
			if 49 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto tr504
			}
		default:
			goto tr495
		}
		goto tr44
tr500:
//line machine.rl:17

	m.pb = m.p

	goto st268
	st268:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof268
		}
	st_case_268:
//line machine.go:30757
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr22
		case 11:
			goto tr511
		case 13:
			goto tr22
		case 32:
			goto tr510
		case 34:
			goto tr512
		case 44:
			goto tr513
		case 61:
			goto tr22
		case 92:
			goto tr514
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr510
		}
		goto tr509
tr509:
//line machine.rl:17

	m.pb = m.p

	goto st269
	st269:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof269
		}
	st_case_269:
//line machine.go:30791
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr517
		case 13:
			goto st6
		case 32:
			goto tr516
		case 34:
			goto tr518
		case 44:
			goto tr519
		case 61:
			goto st6
		case 92:
			goto st278
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr516
		}
		goto st269
tr517:
//line machine.rl:66

	m.handler.AddTag(key, m.text())

	goto st270
tr511:
//line machine.rl:17

	m.pb = m.p

//line machine.rl:66

	m.handler.AddTag(key, m.text())

	goto st270
	st270:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof270
		}
	st_case_270:
//line machine.go:30835
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr522
		case 13:
			goto st6
		case 32:
			goto tr516
		case 34:
			goto tr523
		case 44:
			goto tr519
		case 61:
			goto st6
		case 92:
			goto tr524
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr516
		}
		goto tr521
tr521:
//line machine.rl:17

	m.pb = m.p

	goto st271
	st271:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof271
		}
	st_case_271:
//line machine.go:30869
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr526
		case 13:
			goto st6
		case 32:
			goto tr516
		case 34:
			goto tr527
		case 44:
			goto tr519
		case 61:
			goto tr417
		case 92:
			goto st280
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr516
		}
		goto st271
tr526:
//line machine.rl:66

	m.handler.AddTag(key, m.text())

	goto st272
tr522:
//line machine.rl:66

	m.handler.AddTag(key, m.text())

//line machine.rl:17

	m.pb = m.p

	goto st272
	st272:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof272
		}
	st_case_272:
//line machine.go:30913
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr522
		case 13:
			goto st6
		case 32:
			goto tr516
		case 34:
			goto tr523
		case 44:
			goto tr519
		case 61:
			goto tr417
		case 92:
			goto tr524
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr516
		}
		goto tr521
tr527:
//line machine.rl:86

	m.handler.AddString(key, m.text())

	goto st977
tr523:
//line machine.rl:86

	m.handler.AddString(key, m.text())

//line machine.rl:17

	m.pb = m.p

	goto st977
	st977:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof977
		}
	st_case_977:
//line machine.go:30957
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr850
		case 11:
			goto tr1471
		case 13:
			goto tr850
		case 32:
			goto tr1470
		case 44:
			goto tr1472
		case 61:
			goto tr11
		case 92:
			goto st18
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1470
		}
		goto st16
tr1471:
//line machine.rl:66

	m.handler.AddTag(key, m.text())

	goto st978
tr1473:
//line machine.rl:66

	m.handler.AddTag(key, m.text())

//line machine.rl:17

	m.pb = m.p

	goto st978
	st978:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof978
		}
	st_case_978:
//line machine.go:30999
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr850
		case 11:
			goto tr1473
		case 13:
			goto tr850
		case 32:
			goto tr1470
		case 44:
			goto tr50
		case 45:
			goto tr1474
		case 61:
			goto tr11
		case 92:
			goto tr54
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto tr1475
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1470
		}
		goto tr52
tr1474:
//line machine.rl:17

	m.pb = m.p

	goto st273
	st273:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof273
		}
	st_case_273:
//line machine.go:31038
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr529
		case 11:
			goto tr56
		case 13:
			goto tr529
		case 32:
			goto tr47
		case 44:
			goto tr50
		case 61:
			goto tr11
		case 92:
			goto st18
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st979
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr47
		}
		goto st16
tr1475:
//line machine.rl:17

	m.pb = m.p

	goto st979
	st979:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof979
		}
	st_case_979:
//line machine.go:31075
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr1477
		case 13:
			goto tr859
		case 32:
			goto tr1476
		case 44:
			goto tr50
		case 61:
			goto tr11
		case 92:
			goto st18
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st980
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1476
		}
		goto st16
	st980:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof980
		}
	st_case_980:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr1477
		case 13:
			goto tr859
		case 32:
			goto tr1476
		case 44:
			goto tr50
		case 61:
			goto tr11
		case 92:
			goto st18
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st981
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1476
		}
		goto st16
	st981:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof981
		}
	st_case_981:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr1477
		case 13:
			goto tr859
		case 32:
			goto tr1476
		case 44:
			goto tr50
		case 61:
			goto tr11
		case 92:
			goto st18
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st982
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1476
		}
		goto st16
	st982:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof982
		}
	st_case_982:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr1477
		case 13:
			goto tr859
		case 32:
			goto tr1476
		case 44:
			goto tr50
		case 61:
			goto tr11
		case 92:
			goto st18
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st983
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1476
		}
		goto st16
	st983:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof983
		}
	st_case_983:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr1477
		case 13:
			goto tr859
		case 32:
			goto tr1476
		case 44:
			goto tr50
		case 61:
			goto tr11
		case 92:
			goto st18
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st984
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1476
		}
		goto st16
	st984:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof984
		}
	st_case_984:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr1477
		case 13:
			goto tr859
		case 32:
			goto tr1476
		case 44:
			goto tr50
		case 61:
			goto tr11
		case 92:
			goto st18
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st985
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1476
		}
		goto st16
	st985:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof985
		}
	st_case_985:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr1477
		case 13:
			goto tr859
		case 32:
			goto tr1476
		case 44:
			goto tr50
		case 61:
			goto tr11
		case 92:
			goto st18
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st986
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1476
		}
		goto st16
	st986:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof986
		}
	st_case_986:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr1477
		case 13:
			goto tr859
		case 32:
			goto tr1476
		case 44:
			goto tr50
		case 61:
			goto tr11
		case 92:
			goto st18
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st987
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1476
		}
		goto st16
	st987:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof987
		}
	st_case_987:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr1477
		case 13:
			goto tr859
		case 32:
			goto tr1476
		case 44:
			goto tr50
		case 61:
			goto tr11
		case 92:
			goto st18
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st988
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1476
		}
		goto st16
	st988:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof988
		}
	st_case_988:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr1477
		case 13:
			goto tr859
		case 32:
			goto tr1476
		case 44:
			goto tr50
		case 61:
			goto tr11
		case 92:
			goto st18
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st989
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1476
		}
		goto st16
	st989:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof989
		}
	st_case_989:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr1477
		case 13:
			goto tr859
		case 32:
			goto tr1476
		case 44:
			goto tr50
		case 61:
			goto tr11
		case 92:
			goto st18
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st990
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1476
		}
		goto st16
	st990:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof990
		}
	st_case_990:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr1477
		case 13:
			goto tr859
		case 32:
			goto tr1476
		case 44:
			goto tr50
		case 61:
			goto tr11
		case 92:
			goto st18
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st991
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1476
		}
		goto st16
	st991:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof991
		}
	st_case_991:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr1477
		case 13:
			goto tr859
		case 32:
			goto tr1476
		case 44:
			goto tr50
		case 61:
			goto tr11
		case 92:
			goto st18
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st992
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1476
		}
		goto st16
	st992:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof992
		}
	st_case_992:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr1477
		case 13:
			goto tr859
		case 32:
			goto tr1476
		case 44:
			goto tr50
		case 61:
			goto tr11
		case 92:
			goto st18
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st993
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1476
		}
		goto st16
	st993:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof993
		}
	st_case_993:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr1477
		case 13:
			goto tr859
		case 32:
			goto tr1476
		case 44:
			goto tr50
		case 61:
			goto tr11
		case 92:
			goto st18
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st994
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1476
		}
		goto st16
	st994:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof994
		}
	st_case_994:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr1477
		case 13:
			goto tr859
		case 32:
			goto tr1476
		case 44:
			goto tr50
		case 61:
			goto tr11
		case 92:
			goto st18
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st995
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1476
		}
		goto st16
	st995:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof995
		}
	st_case_995:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr1477
		case 13:
			goto tr859
		case 32:
			goto tr1476
		case 44:
			goto tr50
		case 61:
			goto tr11
		case 92:
			goto st18
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st996
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1476
		}
		goto st16
	st996:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof996
		}
	st_case_996:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr1477
		case 13:
			goto tr859
		case 32:
			goto tr1476
		case 44:
			goto tr50
		case 61:
			goto tr11
		case 92:
			goto st18
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st997
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1476
		}
		goto st16
	st997:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof997
		}
	st_case_997:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr1477
		case 13:
			goto tr859
		case 32:
			goto tr1476
		case 44:
			goto tr50
		case 61:
			goto tr11
		case 92:
			goto st18
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1476
		}
		goto st16
tr408:
//line machine.rl:58

	m.handler.SetMeasurement(m.text())

	goto st274
tr519:
//line machine.rl:66

	m.handler.AddTag(key, m.text())

	goto st274
tr513:
//line machine.rl:17

	m.pb = m.p

//line machine.rl:66

	m.handler.AddTag(key, m.text())

	goto st274
	st274:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof274
		}
	st_case_274:
//line machine.go:31663
		switch ( m.data)[( m.p)] {
		case 32:
			goto st6
		case 34:
			goto tr532
		case 44:
			goto st6
		case 61:
			goto st6
		case 92:
			goto tr533
		}
		switch {
		case ( m.data)[( m.p)] > 10:
			if 12 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto st6
			}
		case ( m.data)[( m.p)] >= 9:
			goto st6
		}
		goto tr531
tr531:
//line machine.rl:17

	m.pb = m.p

	goto st275
	st275:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof275
		}
	st_case_275:
//line machine.go:31696
		switch ( m.data)[( m.p)] {
		case 32:
			goto st6
		case 34:
			goto tr535
		case 44:
			goto st6
		case 61:
			goto tr536
		case 92:
			goto st279
		}
		switch {
		case ( m.data)[( m.p)] > 10:
			if 12 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto st6
			}
		case ( m.data)[( m.p)] >= 9:
			goto st6
		}
		goto st275
tr535:
//line machine.rl:86

	m.handler.AddString(key, m.text())

	goto st998
tr532:
//line machine.rl:86

	m.handler.AddString(key, m.text())

//line machine.rl:17

	m.pb = m.p

	goto st998
	st998:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof998
		}
	st_case_998:
//line machine.go:31739
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr850
		case 11:
			goto st999
		case 13:
			goto tr850
		case 32:
			goto st492
		case 44:
			goto st28
		case 61:
			goto tr42
		case 92:
			goto st20
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto st492
		}
		goto st12
	st999:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof999
		}
	st_case_999:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr850
		case 11:
			goto st999
		case 13:
			goto tr850
		case 32:
			goto st492
		case 44:
			goto tr538
		case 45:
			goto tr1497
		case 61:
			goto tr42
		case 92:
			goto st20
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto tr1498
			}
		case ( m.data)[( m.p)] >= 9:
			goto st492
		}
		goto st12
tr1497:
//line machine.rl:17

	m.pb = m.p

	goto st276
	st276:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof276
		}
	st_case_276:
//line machine.go:31803
		switch ( m.data)[( m.p)] {
		case 32:
			goto tr538
		case 44:
			goto tr538
		case 61:
			goto tr42
		case 92:
			goto st20
		}
		switch {
		case ( m.data)[( m.p)] < 12:
			if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 10 {
				goto tr538
			}
		case ( m.data)[( m.p)] > 13:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1000
			}
		default:
			goto tr538
		}
		goto st12
tr1498:
//line machine.rl:17

	m.pb = m.p

	goto st1000
	st1000:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1000
		}
	st_case_1000:
//line machine.go:31838
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr934
		case 13:
			goto tr859
		case 32:
			goto tr858
		case 44:
			goto tr538
		case 61:
			goto tr42
		case 92:
			goto st20
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1001
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr858
		}
		goto st12
	st1001:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1001
		}
	st_case_1001:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr934
		case 13:
			goto tr859
		case 32:
			goto tr858
		case 44:
			goto tr538
		case 61:
			goto tr42
		case 92:
			goto st20
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1002
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr858
		}
		goto st12
	st1002:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1002
		}
	st_case_1002:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr934
		case 13:
			goto tr859
		case 32:
			goto tr858
		case 44:
			goto tr538
		case 61:
			goto tr42
		case 92:
			goto st20
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1003
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr858
		}
		goto st12
	st1003:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1003
		}
	st_case_1003:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr934
		case 13:
			goto tr859
		case 32:
			goto tr858
		case 44:
			goto tr538
		case 61:
			goto tr42
		case 92:
			goto st20
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1004
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr858
		}
		goto st12
	st1004:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1004
		}
	st_case_1004:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr934
		case 13:
			goto tr859
		case 32:
			goto tr858
		case 44:
			goto tr538
		case 61:
			goto tr42
		case 92:
			goto st20
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1005
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr858
		}
		goto st12
	st1005:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1005
		}
	st_case_1005:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr934
		case 13:
			goto tr859
		case 32:
			goto tr858
		case 44:
			goto tr538
		case 61:
			goto tr42
		case 92:
			goto st20
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1006
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr858
		}
		goto st12
	st1006:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1006
		}
	st_case_1006:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr934
		case 13:
			goto tr859
		case 32:
			goto tr858
		case 44:
			goto tr538
		case 61:
			goto tr42
		case 92:
			goto st20
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1007
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr858
		}
		goto st12
	st1007:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1007
		}
	st_case_1007:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr934
		case 13:
			goto tr859
		case 32:
			goto tr858
		case 44:
			goto tr538
		case 61:
			goto tr42
		case 92:
			goto st20
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1008
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr858
		}
		goto st12
	st1008:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1008
		}
	st_case_1008:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr934
		case 13:
			goto tr859
		case 32:
			goto tr858
		case 44:
			goto tr538
		case 61:
			goto tr42
		case 92:
			goto st20
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1009
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr858
		}
		goto st12
	st1009:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1009
		}
	st_case_1009:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr934
		case 13:
			goto tr859
		case 32:
			goto tr858
		case 44:
			goto tr538
		case 61:
			goto tr42
		case 92:
			goto st20
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1010
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr858
		}
		goto st12
	st1010:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1010
		}
	st_case_1010:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr934
		case 13:
			goto tr859
		case 32:
			goto tr858
		case 44:
			goto tr538
		case 61:
			goto tr42
		case 92:
			goto st20
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1011
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr858
		}
		goto st12
	st1011:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1011
		}
	st_case_1011:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr934
		case 13:
			goto tr859
		case 32:
			goto tr858
		case 44:
			goto tr538
		case 61:
			goto tr42
		case 92:
			goto st20
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1012
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr858
		}
		goto st12
	st1012:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1012
		}
	st_case_1012:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr934
		case 13:
			goto tr859
		case 32:
			goto tr858
		case 44:
			goto tr538
		case 61:
			goto tr42
		case 92:
			goto st20
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1013
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr858
		}
		goto st12
	st1013:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1013
		}
	st_case_1013:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr934
		case 13:
			goto tr859
		case 32:
			goto tr858
		case 44:
			goto tr538
		case 61:
			goto tr42
		case 92:
			goto st20
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1014
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr858
		}
		goto st12
	st1014:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1014
		}
	st_case_1014:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr934
		case 13:
			goto tr859
		case 32:
			goto tr858
		case 44:
			goto tr538
		case 61:
			goto tr42
		case 92:
			goto st20
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1015
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr858
		}
		goto st12
	st1015:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1015
		}
	st_case_1015:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr934
		case 13:
			goto tr859
		case 32:
			goto tr858
		case 44:
			goto tr538
		case 61:
			goto tr42
		case 92:
			goto st20
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1016
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr858
		}
		goto st12
	st1016:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1016
		}
	st_case_1016:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr934
		case 13:
			goto tr859
		case 32:
			goto tr858
		case 44:
			goto tr538
		case 61:
			goto tr42
		case 92:
			goto st20
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1017
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr858
		}
		goto st12
	st1017:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1017
		}
	st_case_1017:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr934
		case 13:
			goto tr859
		case 32:
			goto tr858
		case 44:
			goto tr538
		case 61:
			goto tr42
		case 92:
			goto st20
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1018
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr858
		}
		goto st12
	st1018:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1018
		}
	st_case_1018:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr934
		case 13:
			goto tr859
		case 32:
			goto tr858
		case 44:
			goto tr538
		case 61:
			goto tr42
		case 92:
			goto st20
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr858
		}
		goto st12
tr536:
//line machine.rl:62

	key = m.text()

	goto st277
	st277:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof277
		}
	st_case_277:
//line machine.go:32410
		switch ( m.data)[( m.p)] {
		case 32:
			goto st6
		case 34:
			goto tr540
		case 44:
			goto st6
		case 61:
			goto st6
		case 92:
			goto tr514
		}
		switch {
		case ( m.data)[( m.p)] > 10:
			if 12 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto st6
			}
		case ( m.data)[( m.p)] >= 9:
			goto st6
		}
		goto tr509
tr512:
//line machine.rl:17

	m.pb = m.p

//line machine.rl:86

	m.handler.AddString(key, m.text())

	goto st1019
tr518:
//line machine.rl:86

	m.handler.AddString(key, m.text())

	goto st1019
tr540:
//line machine.rl:86

	m.handler.AddString(key, m.text())

//line machine.rl:17

	m.pb = m.p

	goto st1019
	st1019:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1019
		}
	st_case_1019:
//line machine.go:32463
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr850
		case 11:
			goto tr1517
		case 13:
			goto tr850
		case 32:
			goto tr1470
		case 44:
			goto tr1472
		case 61:
			goto tr495
		case 92:
			goto st19
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1470
		}
		goto st14
tr1517:
//line machine.rl:66

	m.handler.AddTag(key, m.text())

	goto st1020
tr1519:
//line machine.rl:78

	m.handler.AddFloat(key, m.text())

//line machine.rl:66

	m.handler.AddTag(key, m.text())

	goto st1020
tr1524:
//line machine.rl:74

	m.handler.AddInt(key, m.text())

//line machine.rl:66

	m.handler.AddTag(key, m.text())

	goto st1020
tr1527:
//line machine.rl:82

	m.handler.AddBool(key, m.text())

//line machine.rl:66

	m.handler.AddTag(key, m.text())

	goto st1020
	st1020:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1020
		}
	st_case_1020:
//line machine.go:32525
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr850
		case 11:
			goto tr1473
		case 13:
			goto tr850
		case 32:
			goto tr1470
		case 44:
			goto tr50
		case 45:
			goto tr1474
		case 61:
			goto tr529
		case 92:
			goto tr54
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto tr1475
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1470
		}
		goto tr52
tr514:
//line machine.rl:17

	m.pb = m.p

	goto st278
	st278:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof278
		}
	st_case_278:
//line machine.go:32564
		switch ( m.data)[( m.p)] {
		case 34:
			goto st269
		case 92:
			goto st269
		}
		switch {
		case ( m.data)[( m.p)] > 10:
			if 12 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto tr495
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr495
		}
		goto st14
tr533:
//line machine.rl:17

	m.pb = m.p

	goto st279
	st279:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof279
		}
	st_case_279:
//line machine.go:32591
		switch ( m.data)[( m.p)] {
		case 34:
			goto st275
		case 92:
			goto st275
		}
		switch {
		case ( m.data)[( m.p)] > 10:
			if 12 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto tr495
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr495
		}
		goto st12
tr524:
//line machine.rl:17

	m.pb = m.p

	goto st280
	st280:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof280
		}
	st_case_280:
//line machine.go:32618
		switch ( m.data)[( m.p)] {
		case 34:
			goto st271
		case 92:
			goto st271
		}
		switch {
		case ( m.data)[( m.p)] > 10:
			if 12 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto tr495
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr495
		}
		goto st16
tr501:
//line machine.rl:17

	m.pb = m.p

	goto st281
	st281:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof281
		}
	st_case_281:
//line machine.go:32645
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr495
		case 11:
			goto tr49
		case 13:
			goto tr495
		case 32:
			goto tr47
		case 44:
			goto tr50
		case 48:
			goto st1021
		case 61:
			goto tr495
		case 92:
			goto st19
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 49 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1024
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr47
		}
		goto st14
tr503:
//line machine.rl:17

	m.pb = m.p

	goto st1021
	st1021:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1021
		}
	st_case_1021:
//line machine.go:32684
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1270
		case 11:
			goto tr1519
		case 13:
			goto tr1270
		case 32:
			goto tr1518
		case 44:
			goto tr1520
		case 46:
			goto st1022
		case 61:
			goto tr495
		case 92:
			goto st19
		case 105:
			goto st1023
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1518
		}
		goto st14
tr502:
//line machine.rl:17

	m.pb = m.p

	goto st1022
	st1022:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1022
		}
	st_case_1022:
//line machine.go:32720
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1270
		case 11:
			goto tr1519
		case 13:
			goto tr1270
		case 32:
			goto tr1518
		case 44:
			goto tr1520
		case 61:
			goto tr495
		case 92:
			goto st19
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1022
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1518
		}
		goto st14
	st1023:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1023
		}
	st_case_1023:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1276
		case 11:
			goto tr1524
		case 13:
			goto tr1276
		case 32:
			goto tr1523
		case 44:
			goto tr1525
		case 61:
			goto tr495
		case 92:
			goto st19
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1523
		}
		goto st14
tr504:
//line machine.rl:17

	m.pb = m.p

	goto st1024
	st1024:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1024
		}
	st_case_1024:
//line machine.go:32782
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1270
		case 11:
			goto tr1519
		case 13:
			goto tr1270
		case 32:
			goto tr1518
		case 44:
			goto tr1520
		case 46:
			goto st1022
		case 61:
			goto tr495
		case 92:
			goto st19
		case 105:
			goto st1023
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1024
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1518
		}
		goto st14
tr505:
//line machine.rl:17

	m.pb = m.p

	goto st1025
	st1025:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1025
		}
	st_case_1025:
//line machine.go:32823
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1280
		case 11:
			goto tr1527
		case 13:
			goto tr1280
		case 32:
			goto tr1526
		case 44:
			goto tr1528
		case 61:
			goto tr495
		case 65:
			goto st282
		case 92:
			goto st19
		case 97:
			goto st285
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1526
		}
		goto st14
	st282:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof282
		}
	st_case_282:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr495
		case 11:
			goto tr49
		case 13:
			goto tr495
		case 32:
			goto tr47
		case 44:
			goto tr50
		case 61:
			goto tr495
		case 76:
			goto st283
		case 92:
			goto st19
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr47
		}
		goto st14
	st283:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof283
		}
	st_case_283:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr495
		case 11:
			goto tr49
		case 13:
			goto tr495
		case 32:
			goto tr47
		case 44:
			goto tr50
		case 61:
			goto tr495
		case 83:
			goto st284
		case 92:
			goto st19
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr47
		}
		goto st14
	st284:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof284
		}
	st_case_284:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr495
		case 11:
			goto tr49
		case 13:
			goto tr495
		case 32:
			goto tr47
		case 44:
			goto tr50
		case 61:
			goto tr495
		case 69:
			goto st1026
		case 92:
			goto st19
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr47
		}
		goto st14
	st1026:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1026
		}
	st_case_1026:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1280
		case 11:
			goto tr1527
		case 13:
			goto tr1280
		case 32:
			goto tr1526
		case 44:
			goto tr1528
		case 61:
			goto tr495
		case 92:
			goto st19
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1526
		}
		goto st14
	st285:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof285
		}
	st_case_285:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr495
		case 11:
			goto tr49
		case 13:
			goto tr495
		case 32:
			goto tr47
		case 44:
			goto tr50
		case 61:
			goto tr495
		case 92:
			goto st19
		case 108:
			goto st286
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr47
		}
		goto st14
	st286:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof286
		}
	st_case_286:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr495
		case 11:
			goto tr49
		case 13:
			goto tr495
		case 32:
			goto tr47
		case 44:
			goto tr50
		case 61:
			goto tr495
		case 92:
			goto st19
		case 115:
			goto st287
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr47
		}
		goto st14
	st287:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof287
		}
	st_case_287:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr495
		case 11:
			goto tr49
		case 13:
			goto tr495
		case 32:
			goto tr47
		case 44:
			goto tr50
		case 61:
			goto tr495
		case 92:
			goto st19
		case 101:
			goto st1026
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr47
		}
		goto st14
tr506:
//line machine.rl:17

	m.pb = m.p

	goto st1027
	st1027:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1027
		}
	st_case_1027:
//line machine.go:33046
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1280
		case 11:
			goto tr1527
		case 13:
			goto tr1280
		case 32:
			goto tr1526
		case 44:
			goto tr1528
		case 61:
			goto tr495
		case 82:
			goto st288
		case 92:
			goto st19
		case 114:
			goto st289
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1526
		}
		goto st14
	st288:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof288
		}
	st_case_288:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr495
		case 11:
			goto tr49
		case 13:
			goto tr495
		case 32:
			goto tr47
		case 44:
			goto tr50
		case 61:
			goto tr495
		case 85:
			goto st284
		case 92:
			goto st19
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr47
		}
		goto st14
	st289:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof289
		}
	st_case_289:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr495
		case 11:
			goto tr49
		case 13:
			goto tr495
		case 32:
			goto tr47
		case 44:
			goto tr50
		case 61:
			goto tr495
		case 92:
			goto st19
		case 117:
			goto st287
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr47
		}
		goto st14
tr507:
//line machine.rl:17

	m.pb = m.p

	goto st1028
	st1028:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1028
		}
	st_case_1028:
//line machine.go:33136
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1280
		case 11:
			goto tr1527
		case 13:
			goto tr1280
		case 32:
			goto tr1526
		case 44:
			goto tr1528
		case 61:
			goto tr495
		case 92:
			goto st19
		case 97:
			goto st285
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1526
		}
		goto st14
tr508:
//line machine.rl:17

	m.pb = m.p

	goto st1029
	st1029:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1029
		}
	st_case_1029:
//line machine.go:33170
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1280
		case 11:
			goto tr1527
		case 13:
			goto tr1280
		case 32:
			goto tr1526
		case 44:
			goto tr1528
		case 61:
			goto tr495
		case 92:
			goto st19
		case 114:
			goto st289
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1526
		}
		goto st14
tr496:
//line machine.rl:17

	m.pb = m.p

	goto st290
	st290:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof290
		}
	st_case_290:
//line machine.go:33204
		switch {
		case ( m.data)[( m.p)] > 10:
			if 12 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto tr495
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr495
		}
		goto st266
tr491:
//line machine.rl:70

	key = m.text()

	goto st291
tr649:
//line machine.rl:17

	m.pb = m.p

//line machine.rl:70

	key = m.text()

	goto st291
	st291:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof291
		}
	st_case_291:
//line machine.go:33235
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr406
		case 13:
			goto st6
		case 32:
			goto tr405
		case 34:
			goto tr548
		case 44:
			goto tr408
		case 45:
			goto tr549
		case 46:
			goto tr550
		case 48:
			goto tr551
		case 70:
			goto tr553
		case 84:
			goto tr554
		case 92:
			goto st343
		case 102:
			goto tr555
		case 116:
			goto tr556
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 49 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto tr552
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr405
		}
		goto st213
tr548:
//line machine.rl:86

	m.handler.AddString(key, m.text())

	goto st1030
	st1030:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1030
		}
	st_case_1030:
//line machine.go:33286
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1392
		case 11:
			goto tr1534
		case 13:
			goto tr1392
		case 32:
			goto tr1533
		case 34:
			goto tr70
		case 44:
			goto tr1535
		case 92:
			goto tr72
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1533
		}
		goto tr67
tr1539:
//line machine.rl:58

	m.handler.SetMeasurement(m.text())

	goto st1031
tr1533:
//line machine.rl:58

	m.handler.SetMeasurement(m.text())

//line machine.rl:17

	m.pb = m.p

	goto st1031
tr1731:
//line machine.rl:66

	m.handler.AddTag(key, m.text())

	goto st1031
tr1726:
//line machine.rl:78

	m.handler.AddFloat(key, m.text())

//line machine.rl:66

	m.handler.AddTag(key, m.text())

	goto st1031
tr1755:
//line machine.rl:74

	m.handler.AddInt(key, m.text())

//line machine.rl:66

	m.handler.AddTag(key, m.text())

	goto st1031
tr1758:
//line machine.rl:82

	m.handler.AddBool(key, m.text())

//line machine.rl:66

	m.handler.AddTag(key, m.text())

	goto st1031
	st1031:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1031
		}
	st_case_1031:
//line machine.go:33364
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr885
		case 11:
			goto tr1537
		case 13:
			goto tr885
		case 32:
			goto st1031
		case 34:
			goto tr82
		case 44:
			goto st6
		case 45:
			goto tr1538
		case 61:
			goto st6
		case 92:
			goto tr83
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto tr988
			}
		case ( m.data)[( m.p)] >= 9:
			goto st1031
		}
		goto tr79
tr1537:
//line machine.rl:17

	m.pb = m.p

	goto st1032
	st1032:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1032
		}
	st_case_1032:
//line machine.go:33405
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr885
		case 11:
			goto tr1537
		case 13:
			goto tr885
		case 32:
			goto st1031
		case 34:
			goto tr82
		case 44:
			goto st6
		case 45:
			goto tr1538
		case 61:
			goto tr86
		case 92:
			goto tr83
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto tr988
			}
		case ( m.data)[( m.p)] >= 9:
			goto st1031
		}
		goto tr79
tr1538:
//line machine.rl:17

	m.pb = m.p

	goto st292
	st292:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof292
		}
	st_case_292:
//line machine.go:33446
		switch ( m.data)[( m.p)] {
		case 32:
			goto st6
		case 34:
			goto tr85
		case 44:
			goto st6
		case 61:
			goto tr86
		case 92:
			goto st60
		}
		switch {
		case ( m.data)[( m.p)] < 12:
			if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 10 {
				goto st6
			}
		case ( m.data)[( m.p)] > 13:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st597
			}
		default:
			goto st6
		}
		goto st25
tr1534:
//line machine.rl:58

	m.handler.SetMeasurement(m.text())

//line machine.rl:17

	m.pb = m.p

	goto st1033
	st1033:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1033
		}
	st_case_1033:
//line machine.go:33487
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr885
		case 11:
			goto tr1540
		case 13:
			goto tr885
		case 32:
			goto tr1539
		case 34:
			goto tr106
		case 44:
			goto tr77
		case 45:
			goto tr1541
		case 61:
			goto st23
		case 92:
			goto tr107
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto tr1542
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1539
		}
		goto tr103
tr1540:
//line machine.rl:58

	m.handler.SetMeasurement(m.text())

//line machine.rl:17

	m.pb = m.p

	goto st1034
	st1034:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1034
		}
	st_case_1034:
//line machine.go:33532
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr885
		case 11:
			goto tr1540
		case 13:
			goto tr885
		case 32:
			goto tr1539
		case 34:
			goto tr106
		case 44:
			goto tr77
		case 45:
			goto tr1541
		case 61:
			goto tr111
		case 92:
			goto tr107
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto tr1542
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1539
		}
		goto tr103
tr1541:
//line machine.rl:17

	m.pb = m.p

	goto st293
	st293:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof293
		}
	st_case_293:
//line machine.go:33573
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr109
		case 13:
			goto st6
		case 32:
			goto tr74
		case 34:
			goto tr110
		case 44:
			goto tr77
		case 61:
			goto tr111
		case 92:
			goto st74
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1035
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr74
		}
		goto st33
tr1542:
//line machine.rl:17

	m.pb = m.p

	goto st1035
	st1035:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1035
		}
	st_case_1035:
//line machine.go:33612
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr990
		case 11:
			goto tr1544
		case 13:
			goto tr990
		case 32:
			goto tr1543
		case 34:
			goto tr110
		case 44:
			goto tr77
		case 61:
			goto tr111
		case 92:
			goto st74
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1037
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1543
		}
		goto st33
tr1546:
//line machine.rl:58

	m.handler.SetMeasurement(m.text())

//line machine.rl:17

	m.pb = m.p

	goto st1036
tr1544:
//line machine.rl:90

	m.handler.SetTimestamp(m.text())

//line machine.rl:58

	m.handler.SetMeasurement(m.text())

	goto st1036
	st1036:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1036
		}
	st_case_1036:
//line machine.go:33665
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr885
		case 11:
			goto tr1546
		case 13:
			goto tr885
		case 32:
			goto tr1022
		case 34:
			goto tr106
		case 44:
			goto tr77
		case 61:
			goto tr111
		case 92:
			goto tr107
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1022
		}
		goto tr103
	st1037:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1037
		}
	st_case_1037:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr990
		case 11:
			goto tr1544
		case 13:
			goto tr990
		case 32:
			goto tr1543
		case 34:
			goto tr110
		case 44:
			goto tr77
		case 61:
			goto tr111
		case 92:
			goto st74
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1038
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1543
		}
		goto st33
	st1038:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1038
		}
	st_case_1038:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr990
		case 11:
			goto tr1544
		case 13:
			goto tr990
		case 32:
			goto tr1543
		case 34:
			goto tr110
		case 44:
			goto tr77
		case 61:
			goto tr111
		case 92:
			goto st74
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1039
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1543
		}
		goto st33
	st1039:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1039
		}
	st_case_1039:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr990
		case 11:
			goto tr1544
		case 13:
			goto tr990
		case 32:
			goto tr1543
		case 34:
			goto tr110
		case 44:
			goto tr77
		case 61:
			goto tr111
		case 92:
			goto st74
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1040
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1543
		}
		goto st33
	st1040:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1040
		}
	st_case_1040:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr990
		case 11:
			goto tr1544
		case 13:
			goto tr990
		case 32:
			goto tr1543
		case 34:
			goto tr110
		case 44:
			goto tr77
		case 61:
			goto tr111
		case 92:
			goto st74
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1041
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1543
		}
		goto st33
	st1041:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1041
		}
	st_case_1041:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr990
		case 11:
			goto tr1544
		case 13:
			goto tr990
		case 32:
			goto tr1543
		case 34:
			goto tr110
		case 44:
			goto tr77
		case 61:
			goto tr111
		case 92:
			goto st74
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1042
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1543
		}
		goto st33
	st1042:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1042
		}
	st_case_1042:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr990
		case 11:
			goto tr1544
		case 13:
			goto tr990
		case 32:
			goto tr1543
		case 34:
			goto tr110
		case 44:
			goto tr77
		case 61:
			goto tr111
		case 92:
			goto st74
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1043
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1543
		}
		goto st33
	st1043:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1043
		}
	st_case_1043:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr990
		case 11:
			goto tr1544
		case 13:
			goto tr990
		case 32:
			goto tr1543
		case 34:
			goto tr110
		case 44:
			goto tr77
		case 61:
			goto tr111
		case 92:
			goto st74
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1044
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1543
		}
		goto st33
	st1044:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1044
		}
	st_case_1044:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr990
		case 11:
			goto tr1544
		case 13:
			goto tr990
		case 32:
			goto tr1543
		case 34:
			goto tr110
		case 44:
			goto tr77
		case 61:
			goto tr111
		case 92:
			goto st74
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1045
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1543
		}
		goto st33
	st1045:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1045
		}
	st_case_1045:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr990
		case 11:
			goto tr1544
		case 13:
			goto tr990
		case 32:
			goto tr1543
		case 34:
			goto tr110
		case 44:
			goto tr77
		case 61:
			goto tr111
		case 92:
			goto st74
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1046
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1543
		}
		goto st33
	st1046:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1046
		}
	st_case_1046:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr990
		case 11:
			goto tr1544
		case 13:
			goto tr990
		case 32:
			goto tr1543
		case 34:
			goto tr110
		case 44:
			goto tr77
		case 61:
			goto tr111
		case 92:
			goto st74
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1047
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1543
		}
		goto st33
	st1047:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1047
		}
	st_case_1047:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr990
		case 11:
			goto tr1544
		case 13:
			goto tr990
		case 32:
			goto tr1543
		case 34:
			goto tr110
		case 44:
			goto tr77
		case 61:
			goto tr111
		case 92:
			goto st74
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1048
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1543
		}
		goto st33
	st1048:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1048
		}
	st_case_1048:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr990
		case 11:
			goto tr1544
		case 13:
			goto tr990
		case 32:
			goto tr1543
		case 34:
			goto tr110
		case 44:
			goto tr77
		case 61:
			goto tr111
		case 92:
			goto st74
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1049
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1543
		}
		goto st33
	st1049:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1049
		}
	st_case_1049:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr990
		case 11:
			goto tr1544
		case 13:
			goto tr990
		case 32:
			goto tr1543
		case 34:
			goto tr110
		case 44:
			goto tr77
		case 61:
			goto tr111
		case 92:
			goto st74
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1050
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1543
		}
		goto st33
	st1050:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1050
		}
	st_case_1050:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr990
		case 11:
			goto tr1544
		case 13:
			goto tr990
		case 32:
			goto tr1543
		case 34:
			goto tr110
		case 44:
			goto tr77
		case 61:
			goto tr111
		case 92:
			goto st74
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1051
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1543
		}
		goto st33
	st1051:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1051
		}
	st_case_1051:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr990
		case 11:
			goto tr1544
		case 13:
			goto tr990
		case 32:
			goto tr1543
		case 34:
			goto tr110
		case 44:
			goto tr77
		case 61:
			goto tr111
		case 92:
			goto st74
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1052
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1543
		}
		goto st33
	st1052:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1052
		}
	st_case_1052:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr990
		case 11:
			goto tr1544
		case 13:
			goto tr990
		case 32:
			goto tr1543
		case 34:
			goto tr110
		case 44:
			goto tr77
		case 61:
			goto tr111
		case 92:
			goto st74
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1053
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1543
		}
		goto st33
	st1053:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1053
		}
	st_case_1053:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr990
		case 11:
			goto tr1544
		case 13:
			goto tr990
		case 32:
			goto tr1543
		case 34:
			goto tr110
		case 44:
			goto tr77
		case 61:
			goto tr111
		case 92:
			goto st74
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1054
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1543
		}
		goto st33
	st1054:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1054
		}
	st_case_1054:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr990
		case 11:
			goto tr1544
		case 13:
			goto tr990
		case 32:
			goto tr1543
		case 34:
			goto tr110
		case 44:
			goto tr77
		case 61:
			goto tr111
		case 92:
			goto st74
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1543
		}
		goto st33
tr1535:
//line machine.rl:58

	m.handler.SetMeasurement(m.text())

//line machine.rl:17

	m.pb = m.p

	goto st294
tr1728:
//line machine.rl:78

	m.handler.AddFloat(key, m.text())

//line machine.rl:66

	m.handler.AddTag(key, m.text())

	goto st294
tr1757:
//line machine.rl:74

	m.handler.AddInt(key, m.text())

//line machine.rl:66

	m.handler.AddTag(key, m.text())

	goto st294
tr1760:
//line machine.rl:82

	m.handler.AddBool(key, m.text())

//line machine.rl:66

	m.handler.AddTag(key, m.text())

	goto st294
	st294:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof294
		}
	st_case_294:
//line machine.go:34304
		switch ( m.data)[( m.p)] {
		case 32:
			goto st6
		case 34:
			goto tr559
		case 44:
			goto st6
		case 61:
			goto st6
		case 92:
			goto tr560
		}
		switch {
		case ( m.data)[( m.p)] > 10:
			if 12 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto st6
			}
		case ( m.data)[( m.p)] >= 9:
			goto st6
		}
		goto tr558
tr558:
//line machine.rl:17

	m.pb = m.p

	goto st295
	st295:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof295
		}
	st_case_295:
//line machine.go:34337
		switch ( m.data)[( m.p)] {
		case 32:
			goto st6
		case 34:
			goto tr562
		case 44:
			goto st6
		case 61:
			goto tr563
		case 92:
			goto st339
		}
		switch {
		case ( m.data)[( m.p)] > 10:
			if 12 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto st6
			}
		case ( m.data)[( m.p)] >= 9:
			goto st6
		}
		goto st295
tr559:
//line machine.rl:17

	m.pb = m.p

//line machine.rl:86

	m.handler.AddString(key, m.text())

	goto st1055
tr562:
//line machine.rl:86

	m.handler.AddString(key, m.text())

	goto st1055
	st1055:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1055
		}
	st_case_1055:
//line machine.go:34380
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr850
		case 11:
			goto st1056
		case 13:
			goto tr850
		case 32:
			goto st492
		case 44:
			goto st28
		case 61:
			goto tr498
		case 92:
			goto st290
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto st492
		}
		goto st266
	st1056:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1056
		}
	st_case_1056:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr850
		case 11:
			goto st1056
		case 13:
			goto tr850
		case 32:
			goto st492
		case 44:
			goto tr565
		case 45:
			goto tr1565
		case 61:
			goto tr498
		case 92:
			goto st290
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto tr1566
			}
		case ( m.data)[( m.p)] >= 9:
			goto st492
		}
		goto st266
tr1565:
//line machine.rl:17

	m.pb = m.p

	goto st296
	st296:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof296
		}
	st_case_296:
//line machine.go:34444
		switch ( m.data)[( m.p)] {
		case 32:
			goto tr565
		case 44:
			goto tr565
		case 61:
			goto tr498
		case 92:
			goto st290
		}
		switch {
		case ( m.data)[( m.p)] < 12:
			if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 10 {
				goto tr565
			}
		case ( m.data)[( m.p)] > 13:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1057
			}
		default:
			goto tr565
		}
		goto st266
tr1566:
//line machine.rl:17

	m.pb = m.p

	goto st1057
	st1057:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1057
		}
	st_case_1057:
//line machine.go:34479
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr1567
		case 13:
			goto tr859
		case 32:
			goto tr858
		case 44:
			goto tr565
		case 61:
			goto tr498
		case 92:
			goto st290
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1059
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr858
		}
		goto st266
tr1567:
//line machine.rl:90

	m.handler.SetTimestamp(m.text())

	goto st1058
	st1058:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1058
		}
	st_case_1058:
//line machine.go:34516
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr850
		case 11:
			goto st1058
		case 13:
			goto tr850
		case 32:
			goto st497
		case 44:
			goto tr495
		case 61:
			goto tr498
		case 92:
			goto st290
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto st497
		}
		goto st266
	st1059:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1059
		}
	st_case_1059:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr1567
		case 13:
			goto tr859
		case 32:
			goto tr858
		case 44:
			goto tr565
		case 61:
			goto tr498
		case 92:
			goto st290
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1060
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr858
		}
		goto st266
	st1060:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1060
		}
	st_case_1060:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr1567
		case 13:
			goto tr859
		case 32:
			goto tr858
		case 44:
			goto tr565
		case 61:
			goto tr498
		case 92:
			goto st290
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1061
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr858
		}
		goto st266
	st1061:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1061
		}
	st_case_1061:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr1567
		case 13:
			goto tr859
		case 32:
			goto tr858
		case 44:
			goto tr565
		case 61:
			goto tr498
		case 92:
			goto st290
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1062
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr858
		}
		goto st266
	st1062:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1062
		}
	st_case_1062:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr1567
		case 13:
			goto tr859
		case 32:
			goto tr858
		case 44:
			goto tr565
		case 61:
			goto tr498
		case 92:
			goto st290
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1063
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr858
		}
		goto st266
	st1063:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1063
		}
	st_case_1063:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr1567
		case 13:
			goto tr859
		case 32:
			goto tr858
		case 44:
			goto tr565
		case 61:
			goto tr498
		case 92:
			goto st290
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1064
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr858
		}
		goto st266
	st1064:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1064
		}
	st_case_1064:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr1567
		case 13:
			goto tr859
		case 32:
			goto tr858
		case 44:
			goto tr565
		case 61:
			goto tr498
		case 92:
			goto st290
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1065
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr858
		}
		goto st266
	st1065:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1065
		}
	st_case_1065:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr1567
		case 13:
			goto tr859
		case 32:
			goto tr858
		case 44:
			goto tr565
		case 61:
			goto tr498
		case 92:
			goto st290
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1066
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr858
		}
		goto st266
	st1066:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1066
		}
	st_case_1066:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr1567
		case 13:
			goto tr859
		case 32:
			goto tr858
		case 44:
			goto tr565
		case 61:
			goto tr498
		case 92:
			goto st290
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1067
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr858
		}
		goto st266
	st1067:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1067
		}
	st_case_1067:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr1567
		case 13:
			goto tr859
		case 32:
			goto tr858
		case 44:
			goto tr565
		case 61:
			goto tr498
		case 92:
			goto st290
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1068
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr858
		}
		goto st266
	st1068:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1068
		}
	st_case_1068:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr1567
		case 13:
			goto tr859
		case 32:
			goto tr858
		case 44:
			goto tr565
		case 61:
			goto tr498
		case 92:
			goto st290
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1069
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr858
		}
		goto st266
	st1069:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1069
		}
	st_case_1069:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr1567
		case 13:
			goto tr859
		case 32:
			goto tr858
		case 44:
			goto tr565
		case 61:
			goto tr498
		case 92:
			goto st290
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1070
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr858
		}
		goto st266
	st1070:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1070
		}
	st_case_1070:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr1567
		case 13:
			goto tr859
		case 32:
			goto tr858
		case 44:
			goto tr565
		case 61:
			goto tr498
		case 92:
			goto st290
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1071
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr858
		}
		goto st266
	st1071:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1071
		}
	st_case_1071:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr1567
		case 13:
			goto tr859
		case 32:
			goto tr858
		case 44:
			goto tr565
		case 61:
			goto tr498
		case 92:
			goto st290
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1072
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr858
		}
		goto st266
	st1072:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1072
		}
	st_case_1072:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr1567
		case 13:
			goto tr859
		case 32:
			goto tr858
		case 44:
			goto tr565
		case 61:
			goto tr498
		case 92:
			goto st290
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1073
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr858
		}
		goto st266
	st1073:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1073
		}
	st_case_1073:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr1567
		case 13:
			goto tr859
		case 32:
			goto tr858
		case 44:
			goto tr565
		case 61:
			goto tr498
		case 92:
			goto st290
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1074
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr858
		}
		goto st266
	st1074:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1074
		}
	st_case_1074:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr1567
		case 13:
			goto tr859
		case 32:
			goto tr858
		case 44:
			goto tr565
		case 61:
			goto tr498
		case 92:
			goto st290
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1075
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr858
		}
		goto st266
	st1075:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1075
		}
	st_case_1075:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr1567
		case 13:
			goto tr859
		case 32:
			goto tr858
		case 44:
			goto tr565
		case 61:
			goto tr498
		case 92:
			goto st290
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1076
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr858
		}
		goto st266
	st1076:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1076
		}
	st_case_1076:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr1567
		case 13:
			goto tr859
		case 32:
			goto tr858
		case 44:
			goto tr565
		case 61:
			goto tr498
		case 92:
			goto st290
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr858
		}
		goto st266
tr563:
//line machine.rl:70

	key = m.text()

//line machine.rl:62

	key = m.text()

	goto st297
	st297:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof297
		}
	st_case_297:
//line machine.go:35087
		switch ( m.data)[( m.p)] {
		case 32:
			goto st6
		case 34:
			goto tr567
		case 44:
			goto st6
		case 45:
			goto tr568
		case 46:
			goto tr569
		case 48:
			goto tr570
		case 61:
			goto st6
		case 70:
			goto tr572
		case 84:
			goto tr573
		case 92:
			goto tr197
		case 102:
			goto tr574
		case 116:
			goto tr575
		}
		switch {
		case ( m.data)[( m.p)] < 12:
			if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 10 {
				goto st6
			}
		case ( m.data)[( m.p)] > 13:
			if 49 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto tr571
			}
		default:
			goto st6
		}
		goto tr196
tr567:
//line machine.rl:17

	m.pb = m.p

//line machine.rl:86

	m.handler.AddString(key, m.text())

	goto st1077
	st1077:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1077
		}
	st_case_1077:
//line machine.go:35142
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1363
		case 11:
			goto tr1588
		case 13:
			goto tr1363
		case 32:
			goto tr1587
		case 34:
			goto tr512
		case 44:
			goto tr1589
		case 61:
			goto tr22
		case 92:
			goto tr514
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1587
		}
		goto tr509
tr1771:
//line machine.rl:58

	m.handler.SetMeasurement(m.text())

	goto st1078
tr1617:
//line machine.rl:66

	m.handler.AddTag(key, m.text())

	goto st1078
tr1587:
//line machine.rl:17

	m.pb = m.p

//line machine.rl:66

	m.handler.AddTag(key, m.text())

	goto st1078
tr1766:
//line machine.rl:58

	m.handler.SetMeasurement(m.text())

//line machine.rl:78

	m.handler.AddFloat(key, m.text())

	goto st1078
tr1711:
//line machine.rl:66

	m.handler.AddTag(key, m.text())

//line machine.rl:78

	m.handler.AddFloat(key, m.text())

	goto st1078
tr1716:
//line machine.rl:66

	m.handler.AddTag(key, m.text())

//line machine.rl:74

	m.handler.AddInt(key, m.text())

	goto st1078
tr1719:
//line machine.rl:66

	m.handler.AddTag(key, m.text())

//line machine.rl:82

	m.handler.AddBool(key, m.text())

	goto st1078
tr1797:
//line machine.rl:58

	m.handler.SetMeasurement(m.text())

//line machine.rl:74

	m.handler.AddInt(key, m.text())

	goto st1078
tr1800:
//line machine.rl:58

	m.handler.SetMeasurement(m.text())

//line machine.rl:82

	m.handler.AddBool(key, m.text())

	goto st1078
	st1078:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1078
		}
	st_case_1078:
//line machine.go:35252
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1366
		case 11:
			goto tr1591
		case 13:
			goto tr1366
		case 32:
			goto st1078
		case 34:
			goto tr413
		case 44:
			goto st6
		case 45:
			goto tr1592
		case 61:
			goto st6
		case 92:
			goto tr414
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto tr1593
			}
		case ( m.data)[( m.p)] >= 9:
			goto st1078
		}
		goto tr410
tr1591:
//line machine.rl:17

	m.pb = m.p

	goto st1079
	st1079:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1079
		}
	st_case_1079:
//line machine.go:35293
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1366
		case 11:
			goto tr1591
		case 13:
			goto tr1366
		case 32:
			goto st1078
		case 34:
			goto tr413
		case 44:
			goto st6
		case 45:
			goto tr1592
		case 61:
			goto tr417
		case 92:
			goto tr414
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto tr1593
			}
		case ( m.data)[( m.p)] >= 9:
			goto st1078
		}
		goto tr410
tr1592:
//line machine.rl:17

	m.pb = m.p

	goto st298
	st298:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof298
		}
	st_case_298:
//line machine.go:35334
		switch ( m.data)[( m.p)] {
		case 32:
			goto st6
		case 34:
			goto tr416
		case 44:
			goto st6
		case 61:
			goto tr417
		case 92:
			goto st259
		}
		switch {
		case ( m.data)[( m.p)] < 12:
			if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 10 {
				goto st6
			}
		case ( m.data)[( m.p)] > 13:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1080
			}
		default:
			goto st6
		}
		goto st215
tr1593:
//line machine.rl:17

	m.pb = m.p

	goto st1080
	st1080:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1080
		}
	st_case_1080:
//line machine.go:35371
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 11:
			goto tr1596
		case 13:
			goto tr1595
		case 32:
			goto tr1594
		case 34:
			goto tr416
		case 44:
			goto st6
		case 61:
			goto tr417
		case 92:
			goto st259
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1083
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1594
		}
		goto st215
tr1594:
//line machine.rl:90

	m.handler.SetTimestamp(m.text())

	goto st1081
	st1081:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1081
		}
	st_case_1081:
//line machine.go:35410
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1366
		case 13:
			goto tr1366
		case 32:
			goto st1081
		case 34:
			goto tr26
		case 92:
			goto st59
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto st1081
		}
		goto st6
tr1596:
//line machine.rl:90

	m.handler.SetTimestamp(m.text())

	goto st1082
	st1082:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1082
		}
	st_case_1082:
//line machine.go:35438
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1366
		case 11:
			goto st1082
		case 13:
			goto tr1366
		case 32:
			goto st1081
		case 34:
			goto tr416
		case 44:
			goto st6
		case 61:
			goto tr417
		case 92:
			goto st259
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto st1081
		}
		goto st215
	st1083:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1083
		}
	st_case_1083:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 11:
			goto tr1596
		case 13:
			goto tr1595
		case 32:
			goto tr1594
		case 34:
			goto tr416
		case 44:
			goto st6
		case 61:
			goto tr417
		case 92:
			goto st259
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1084
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1594
		}
		goto st215
	st1084:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1084
		}
	st_case_1084:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 11:
			goto tr1596
		case 13:
			goto tr1595
		case 32:
			goto tr1594
		case 34:
			goto tr416
		case 44:
			goto st6
		case 61:
			goto tr417
		case 92:
			goto st259
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1085
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1594
		}
		goto st215
	st1085:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1085
		}
	st_case_1085:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 11:
			goto tr1596
		case 13:
			goto tr1595
		case 32:
			goto tr1594
		case 34:
			goto tr416
		case 44:
			goto st6
		case 61:
			goto tr417
		case 92:
			goto st259
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1086
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1594
		}
		goto st215
	st1086:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1086
		}
	st_case_1086:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 11:
			goto tr1596
		case 13:
			goto tr1595
		case 32:
			goto tr1594
		case 34:
			goto tr416
		case 44:
			goto st6
		case 61:
			goto tr417
		case 92:
			goto st259
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1087
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1594
		}
		goto st215
	st1087:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1087
		}
	st_case_1087:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 11:
			goto tr1596
		case 13:
			goto tr1595
		case 32:
			goto tr1594
		case 34:
			goto tr416
		case 44:
			goto st6
		case 61:
			goto tr417
		case 92:
			goto st259
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1088
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1594
		}
		goto st215
	st1088:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1088
		}
	st_case_1088:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 11:
			goto tr1596
		case 13:
			goto tr1595
		case 32:
			goto tr1594
		case 34:
			goto tr416
		case 44:
			goto st6
		case 61:
			goto tr417
		case 92:
			goto st259
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1089
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1594
		}
		goto st215
	st1089:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1089
		}
	st_case_1089:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 11:
			goto tr1596
		case 13:
			goto tr1595
		case 32:
			goto tr1594
		case 34:
			goto tr416
		case 44:
			goto st6
		case 61:
			goto tr417
		case 92:
			goto st259
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1090
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1594
		}
		goto st215
	st1090:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1090
		}
	st_case_1090:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 11:
			goto tr1596
		case 13:
			goto tr1595
		case 32:
			goto tr1594
		case 34:
			goto tr416
		case 44:
			goto st6
		case 61:
			goto tr417
		case 92:
			goto st259
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1091
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1594
		}
		goto st215
	st1091:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1091
		}
	st_case_1091:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 11:
			goto tr1596
		case 13:
			goto tr1595
		case 32:
			goto tr1594
		case 34:
			goto tr416
		case 44:
			goto st6
		case 61:
			goto tr417
		case 92:
			goto st259
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1092
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1594
		}
		goto st215
	st1092:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1092
		}
	st_case_1092:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 11:
			goto tr1596
		case 13:
			goto tr1595
		case 32:
			goto tr1594
		case 34:
			goto tr416
		case 44:
			goto st6
		case 61:
			goto tr417
		case 92:
			goto st259
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1093
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1594
		}
		goto st215
	st1093:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1093
		}
	st_case_1093:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 11:
			goto tr1596
		case 13:
			goto tr1595
		case 32:
			goto tr1594
		case 34:
			goto tr416
		case 44:
			goto st6
		case 61:
			goto tr417
		case 92:
			goto st259
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1094
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1594
		}
		goto st215
	st1094:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1094
		}
	st_case_1094:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 11:
			goto tr1596
		case 13:
			goto tr1595
		case 32:
			goto tr1594
		case 34:
			goto tr416
		case 44:
			goto st6
		case 61:
			goto tr417
		case 92:
			goto st259
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1095
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1594
		}
		goto st215
	st1095:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1095
		}
	st_case_1095:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 11:
			goto tr1596
		case 13:
			goto tr1595
		case 32:
			goto tr1594
		case 34:
			goto tr416
		case 44:
			goto st6
		case 61:
			goto tr417
		case 92:
			goto st259
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1096
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1594
		}
		goto st215
	st1096:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1096
		}
	st_case_1096:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 11:
			goto tr1596
		case 13:
			goto tr1595
		case 32:
			goto tr1594
		case 34:
			goto tr416
		case 44:
			goto st6
		case 61:
			goto tr417
		case 92:
			goto st259
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1097
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1594
		}
		goto st215
	st1097:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1097
		}
	st_case_1097:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 11:
			goto tr1596
		case 13:
			goto tr1595
		case 32:
			goto tr1594
		case 34:
			goto tr416
		case 44:
			goto st6
		case 61:
			goto tr417
		case 92:
			goto st259
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1098
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1594
		}
		goto st215
	st1098:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1098
		}
	st_case_1098:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 11:
			goto tr1596
		case 13:
			goto tr1595
		case 32:
			goto tr1594
		case 34:
			goto tr416
		case 44:
			goto st6
		case 61:
			goto tr417
		case 92:
			goto st259
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1099
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1594
		}
		goto st215
	st1099:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1099
		}
	st_case_1099:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 11:
			goto tr1596
		case 13:
			goto tr1595
		case 32:
			goto tr1594
		case 34:
			goto tr416
		case 44:
			goto st6
		case 61:
			goto tr417
		case 92:
			goto st259
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1100
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1594
		}
		goto st215
	st1100:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1100
		}
	st_case_1100:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 11:
			goto tr1596
		case 13:
			goto tr1595
		case 32:
			goto tr1594
		case 34:
			goto tr416
		case 44:
			goto st6
		case 61:
			goto tr417
		case 92:
			goto st259
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1594
		}
		goto st215
tr1588:
//line machine.rl:17

	m.pb = m.p

//line machine.rl:66

	m.handler.AddTag(key, m.text())

	goto st1101
tr1712:
//line machine.rl:66

	m.handler.AddTag(key, m.text())

//line machine.rl:78

	m.handler.AddFloat(key, m.text())

	goto st1101
tr1717:
//line machine.rl:66

	m.handler.AddTag(key, m.text())

//line machine.rl:74

	m.handler.AddInt(key, m.text())

	goto st1101
tr1720:
//line machine.rl:66

	m.handler.AddTag(key, m.text())

//line machine.rl:82

	m.handler.AddBool(key, m.text())

	goto st1101
	st1101:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1101
		}
	st_case_1101:
//line machine.go:36077
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1366
		case 11:
			goto tr1618
		case 13:
			goto tr1366
		case 32:
			goto tr1617
		case 34:
			goto tr523
		case 44:
			goto tr519
		case 45:
			goto tr1619
		case 61:
			goto st6
		case 92:
			goto tr524
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto tr1620
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1617
		}
		goto tr521
tr1618:
//line machine.rl:66

	m.handler.AddTag(key, m.text())

//line machine.rl:17

	m.pb = m.p

	goto st1102
	st1102:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1102
		}
	st_case_1102:
//line machine.go:36122
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1366
		case 11:
			goto tr1618
		case 13:
			goto tr1366
		case 32:
			goto tr1617
		case 34:
			goto tr523
		case 44:
			goto tr519
		case 45:
			goto tr1619
		case 61:
			goto tr417
		case 92:
			goto tr524
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto tr1620
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1617
		}
		goto tr521
tr1619:
//line machine.rl:17

	m.pb = m.p

	goto st299
	st299:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof299
		}
	st_case_299:
//line machine.go:36163
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr526
		case 13:
			goto st6
		case 32:
			goto tr516
		case 34:
			goto tr527
		case 44:
			goto tr519
		case 61:
			goto tr417
		case 92:
			goto st280
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1103
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr516
		}
		goto st271
tr1620:
//line machine.rl:17

	m.pb = m.p

	goto st1103
	st1103:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1103
		}
	st_case_1103:
//line machine.go:36202
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 11:
			goto tr1622
		case 13:
			goto tr1595
		case 32:
			goto tr1621
		case 34:
			goto tr527
		case 44:
			goto tr519
		case 61:
			goto tr417
		case 92:
			goto st280
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1107
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1621
		}
		goto st271
tr1778:
//line machine.rl:58

	m.handler.SetMeasurement(m.text())

	goto st1104
tr1626:
//line machine.rl:66

	m.handler.AddTag(key, m.text())

	goto st1104
tr1775:
//line machine.rl:58

	m.handler.SetMeasurement(m.text())

//line machine.rl:90

	m.handler.SetTimestamp(m.text())

	goto st1104
tr1621:
//line machine.rl:66

	m.handler.AddTag(key, m.text())

//line machine.rl:90

	m.handler.SetTimestamp(m.text())

	goto st1104
	st1104:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1104
		}
	st_case_1104:
//line machine.go:36267
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1366
		case 11:
			goto tr1625
		case 13:
			goto tr1366
		case 32:
			goto st1104
		case 34:
			goto tr413
		case 44:
			goto st6
		case 61:
			goto st6
		case 92:
			goto tr414
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto st1104
		}
		goto tr410
tr1625:
//line machine.rl:17

	m.pb = m.p

	goto st1105
	st1105:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1105
		}
	st_case_1105:
//line machine.go:36301
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1366
		case 11:
			goto tr1625
		case 13:
			goto tr1366
		case 32:
			goto st1104
		case 34:
			goto tr413
		case 44:
			goto st6
		case 61:
			goto tr417
		case 92:
			goto tr414
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto st1104
		}
		goto tr410
tr1627:
//line machine.rl:66

	m.handler.AddTag(key, m.text())

//line machine.rl:17

	m.pb = m.p

	goto st1106
tr1622:
//line machine.rl:66

	m.handler.AddTag(key, m.text())

//line machine.rl:90

	m.handler.SetTimestamp(m.text())

	goto st1106
	st1106:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1106
		}
	st_case_1106:
//line machine.go:36349
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1366
		case 11:
			goto tr1627
		case 13:
			goto tr1366
		case 32:
			goto tr1626
		case 34:
			goto tr523
		case 44:
			goto tr519
		case 61:
			goto tr417
		case 92:
			goto tr524
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1626
		}
		goto tr521
	st1107:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1107
		}
	st_case_1107:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 11:
			goto tr1622
		case 13:
			goto tr1595
		case 32:
			goto tr1621
		case 34:
			goto tr527
		case 44:
			goto tr519
		case 61:
			goto tr417
		case 92:
			goto st280
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1108
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1621
		}
		goto st271
	st1108:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1108
		}
	st_case_1108:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 11:
			goto tr1622
		case 13:
			goto tr1595
		case 32:
			goto tr1621
		case 34:
			goto tr527
		case 44:
			goto tr519
		case 61:
			goto tr417
		case 92:
			goto st280
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1109
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1621
		}
		goto st271
	st1109:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1109
		}
	st_case_1109:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 11:
			goto tr1622
		case 13:
			goto tr1595
		case 32:
			goto tr1621
		case 34:
			goto tr527
		case 44:
			goto tr519
		case 61:
			goto tr417
		case 92:
			goto st280
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1110
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1621
		}
		goto st271
	st1110:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1110
		}
	st_case_1110:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 11:
			goto tr1622
		case 13:
			goto tr1595
		case 32:
			goto tr1621
		case 34:
			goto tr527
		case 44:
			goto tr519
		case 61:
			goto tr417
		case 92:
			goto st280
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1111
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1621
		}
		goto st271
	st1111:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1111
		}
	st_case_1111:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 11:
			goto tr1622
		case 13:
			goto tr1595
		case 32:
			goto tr1621
		case 34:
			goto tr527
		case 44:
			goto tr519
		case 61:
			goto tr417
		case 92:
			goto st280
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1112
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1621
		}
		goto st271
	st1112:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1112
		}
	st_case_1112:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 11:
			goto tr1622
		case 13:
			goto tr1595
		case 32:
			goto tr1621
		case 34:
			goto tr527
		case 44:
			goto tr519
		case 61:
			goto tr417
		case 92:
			goto st280
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1113
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1621
		}
		goto st271
	st1113:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1113
		}
	st_case_1113:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 11:
			goto tr1622
		case 13:
			goto tr1595
		case 32:
			goto tr1621
		case 34:
			goto tr527
		case 44:
			goto tr519
		case 61:
			goto tr417
		case 92:
			goto st280
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1114
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1621
		}
		goto st271
	st1114:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1114
		}
	st_case_1114:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 11:
			goto tr1622
		case 13:
			goto tr1595
		case 32:
			goto tr1621
		case 34:
			goto tr527
		case 44:
			goto tr519
		case 61:
			goto tr417
		case 92:
			goto st280
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1115
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1621
		}
		goto st271
	st1115:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1115
		}
	st_case_1115:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 11:
			goto tr1622
		case 13:
			goto tr1595
		case 32:
			goto tr1621
		case 34:
			goto tr527
		case 44:
			goto tr519
		case 61:
			goto tr417
		case 92:
			goto st280
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1116
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1621
		}
		goto st271
	st1116:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1116
		}
	st_case_1116:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 11:
			goto tr1622
		case 13:
			goto tr1595
		case 32:
			goto tr1621
		case 34:
			goto tr527
		case 44:
			goto tr519
		case 61:
			goto tr417
		case 92:
			goto st280
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1117
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1621
		}
		goto st271
	st1117:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1117
		}
	st_case_1117:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 11:
			goto tr1622
		case 13:
			goto tr1595
		case 32:
			goto tr1621
		case 34:
			goto tr527
		case 44:
			goto tr519
		case 61:
			goto tr417
		case 92:
			goto st280
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1118
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1621
		}
		goto st271
	st1118:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1118
		}
	st_case_1118:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 11:
			goto tr1622
		case 13:
			goto tr1595
		case 32:
			goto tr1621
		case 34:
			goto tr527
		case 44:
			goto tr519
		case 61:
			goto tr417
		case 92:
			goto st280
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1119
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1621
		}
		goto st271
	st1119:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1119
		}
	st_case_1119:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 11:
			goto tr1622
		case 13:
			goto tr1595
		case 32:
			goto tr1621
		case 34:
			goto tr527
		case 44:
			goto tr519
		case 61:
			goto tr417
		case 92:
			goto st280
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1120
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1621
		}
		goto st271
	st1120:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1120
		}
	st_case_1120:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 11:
			goto tr1622
		case 13:
			goto tr1595
		case 32:
			goto tr1621
		case 34:
			goto tr527
		case 44:
			goto tr519
		case 61:
			goto tr417
		case 92:
			goto st280
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1121
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1621
		}
		goto st271
	st1121:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1121
		}
	st_case_1121:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 11:
			goto tr1622
		case 13:
			goto tr1595
		case 32:
			goto tr1621
		case 34:
			goto tr527
		case 44:
			goto tr519
		case 61:
			goto tr417
		case 92:
			goto st280
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1122
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1621
		}
		goto st271
	st1122:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1122
		}
	st_case_1122:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 11:
			goto tr1622
		case 13:
			goto tr1595
		case 32:
			goto tr1621
		case 34:
			goto tr527
		case 44:
			goto tr519
		case 61:
			goto tr417
		case 92:
			goto st280
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1123
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1621
		}
		goto st271
	st1123:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1123
		}
	st_case_1123:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 11:
			goto tr1622
		case 13:
			goto tr1595
		case 32:
			goto tr1621
		case 34:
			goto tr527
		case 44:
			goto tr519
		case 61:
			goto tr417
		case 92:
			goto st280
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1124
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1621
		}
		goto st271
	st1124:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1124
		}
	st_case_1124:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 11:
			goto tr1622
		case 13:
			goto tr1595
		case 32:
			goto tr1621
		case 34:
			goto tr527
		case 44:
			goto tr519
		case 61:
			goto tr417
		case 92:
			goto st280
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1621
		}
		goto st271
tr1589:
//line machine.rl:17

	m.pb = m.p

//line machine.rl:66

	m.handler.AddTag(key, m.text())

	goto st300
tr1768:
//line machine.rl:58

	m.handler.SetMeasurement(m.text())

//line machine.rl:78

	m.handler.AddFloat(key, m.text())

	goto st300
tr1713:
//line machine.rl:66

	m.handler.AddTag(key, m.text())

//line machine.rl:78

	m.handler.AddFloat(key, m.text())

	goto st300
tr1718:
//line machine.rl:66

	m.handler.AddTag(key, m.text())

//line machine.rl:74

	m.handler.AddInt(key, m.text())

	goto st300
tr1721:
//line machine.rl:66

	m.handler.AddTag(key, m.text())

//line machine.rl:82

	m.handler.AddBool(key, m.text())

	goto st300
tr1799:
//line machine.rl:58

	m.handler.SetMeasurement(m.text())

//line machine.rl:74

	m.handler.AddInt(key, m.text())

	goto st300
tr1802:
//line machine.rl:58

	m.handler.SetMeasurement(m.text())

//line machine.rl:82

	m.handler.AddBool(key, m.text())

	goto st300
	st300:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof300
		}
	st_case_300:
//line machine.go:37018
		switch ( m.data)[( m.p)] {
		case 32:
			goto st6
		case 34:
			goto tr579
		case 44:
			goto st6
		case 61:
			goto st6
		case 92:
			goto tr580
		}
		switch {
		case ( m.data)[( m.p)] > 10:
			if 12 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto st6
			}
		case ( m.data)[( m.p)] >= 9:
			goto st6
		}
		goto tr578
tr578:
//line machine.rl:17

	m.pb = m.p

	goto st301
	st301:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof301
		}
	st_case_301:
//line machine.go:37051
		switch ( m.data)[( m.p)] {
		case 32:
			goto st6
		case 34:
			goto tr582
		case 44:
			goto st6
		case 61:
			goto tr583
		case 92:
			goto st328
		}
		switch {
		case ( m.data)[( m.p)] > 10:
			if 12 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto st6
			}
		case ( m.data)[( m.p)] >= 9:
			goto st6
		}
		goto st301
tr582:
//line machine.rl:86

	m.handler.AddString(key, m.text())

	goto st1125
tr579:
//line machine.rl:86

	m.handler.AddString(key, m.text())

//line machine.rl:17

	m.pb = m.p

	goto st1125
	st1125:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1125
		}
	st_case_1125:
//line machine.go:37094
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr850
		case 11:
			goto st1126
		case 13:
			goto tr850
		case 32:
			goto st492
		case 44:
			goto st28
		case 61:
			goto tr118
		case 92:
			goto st80
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto st492
		}
		goto st38
	st1126:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1126
		}
	st_case_1126:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr850
		case 11:
			goto st1126
		case 13:
			goto tr850
		case 32:
			goto st492
		case 44:
			goto tr529
		case 45:
			goto tr1646
		case 61:
			goto tr118
		case 92:
			goto st80
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto tr1647
			}
		case ( m.data)[( m.p)] >= 9:
			goto st492
		}
		goto st38
tr1646:
//line machine.rl:17

	m.pb = m.p

	goto st302
	st302:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof302
		}
	st_case_302:
//line machine.go:37158
		switch ( m.data)[( m.p)] {
		case 32:
			goto tr529
		case 44:
			goto tr529
		case 61:
			goto tr118
		case 92:
			goto st80
		}
		switch {
		case ( m.data)[( m.p)] < 12:
			if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 10 {
				goto tr529
			}
		case ( m.data)[( m.p)] > 13:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1127
			}
		default:
			goto tr529
		}
		goto st38
tr1647:
//line machine.rl:17

	m.pb = m.p

	goto st1127
	st1127:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1127
		}
	st_case_1127:
//line machine.go:37193
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr1034
		case 13:
			goto tr859
		case 32:
			goto tr858
		case 44:
			goto tr529
		case 61:
			goto tr118
		case 92:
			goto st80
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1128
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr858
		}
		goto st38
	st1128:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1128
		}
	st_case_1128:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr1034
		case 13:
			goto tr859
		case 32:
			goto tr858
		case 44:
			goto tr529
		case 61:
			goto tr118
		case 92:
			goto st80
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1129
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr858
		}
		goto st38
	st1129:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1129
		}
	st_case_1129:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr1034
		case 13:
			goto tr859
		case 32:
			goto tr858
		case 44:
			goto tr529
		case 61:
			goto tr118
		case 92:
			goto st80
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1130
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr858
		}
		goto st38
	st1130:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1130
		}
	st_case_1130:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr1034
		case 13:
			goto tr859
		case 32:
			goto tr858
		case 44:
			goto tr529
		case 61:
			goto tr118
		case 92:
			goto st80
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1131
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr858
		}
		goto st38
	st1131:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1131
		}
	st_case_1131:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr1034
		case 13:
			goto tr859
		case 32:
			goto tr858
		case 44:
			goto tr529
		case 61:
			goto tr118
		case 92:
			goto st80
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1132
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr858
		}
		goto st38
	st1132:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1132
		}
	st_case_1132:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr1034
		case 13:
			goto tr859
		case 32:
			goto tr858
		case 44:
			goto tr529
		case 61:
			goto tr118
		case 92:
			goto st80
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1133
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr858
		}
		goto st38
	st1133:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1133
		}
	st_case_1133:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr1034
		case 13:
			goto tr859
		case 32:
			goto tr858
		case 44:
			goto tr529
		case 61:
			goto tr118
		case 92:
			goto st80
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1134
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr858
		}
		goto st38
	st1134:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1134
		}
	st_case_1134:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr1034
		case 13:
			goto tr859
		case 32:
			goto tr858
		case 44:
			goto tr529
		case 61:
			goto tr118
		case 92:
			goto st80
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1135
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr858
		}
		goto st38
	st1135:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1135
		}
	st_case_1135:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr1034
		case 13:
			goto tr859
		case 32:
			goto tr858
		case 44:
			goto tr529
		case 61:
			goto tr118
		case 92:
			goto st80
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1136
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr858
		}
		goto st38
	st1136:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1136
		}
	st_case_1136:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr1034
		case 13:
			goto tr859
		case 32:
			goto tr858
		case 44:
			goto tr529
		case 61:
			goto tr118
		case 92:
			goto st80
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1137
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr858
		}
		goto st38
	st1137:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1137
		}
	st_case_1137:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr1034
		case 13:
			goto tr859
		case 32:
			goto tr858
		case 44:
			goto tr529
		case 61:
			goto tr118
		case 92:
			goto st80
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1138
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr858
		}
		goto st38
	st1138:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1138
		}
	st_case_1138:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr1034
		case 13:
			goto tr859
		case 32:
			goto tr858
		case 44:
			goto tr529
		case 61:
			goto tr118
		case 92:
			goto st80
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1139
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr858
		}
		goto st38
	st1139:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1139
		}
	st_case_1139:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr1034
		case 13:
			goto tr859
		case 32:
			goto tr858
		case 44:
			goto tr529
		case 61:
			goto tr118
		case 92:
			goto st80
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1140
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr858
		}
		goto st38
	st1140:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1140
		}
	st_case_1140:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr1034
		case 13:
			goto tr859
		case 32:
			goto tr858
		case 44:
			goto tr529
		case 61:
			goto tr118
		case 92:
			goto st80
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1141
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr858
		}
		goto st38
	st1141:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1141
		}
	st_case_1141:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr1034
		case 13:
			goto tr859
		case 32:
			goto tr858
		case 44:
			goto tr529
		case 61:
			goto tr118
		case 92:
			goto st80
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1142
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr858
		}
		goto st38
	st1142:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1142
		}
	st_case_1142:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr1034
		case 13:
			goto tr859
		case 32:
			goto tr858
		case 44:
			goto tr529
		case 61:
			goto tr118
		case 92:
			goto st80
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1143
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr858
		}
		goto st38
	st1143:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1143
		}
	st_case_1143:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr1034
		case 13:
			goto tr859
		case 32:
			goto tr858
		case 44:
			goto tr529
		case 61:
			goto tr118
		case 92:
			goto st80
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1144
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr858
		}
		goto st38
	st1144:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1144
		}
	st_case_1144:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr1034
		case 13:
			goto tr859
		case 32:
			goto tr858
		case 44:
			goto tr529
		case 61:
			goto tr118
		case 92:
			goto st80
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1145
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr858
		}
		goto st38
	st1145:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1145
		}
	st_case_1145:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr1034
		case 13:
			goto tr859
		case 32:
			goto tr858
		case 44:
			goto tr529
		case 61:
			goto tr118
		case 92:
			goto st80
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr858
		}
		goto st38
tr583:
//line machine.rl:62

	key = m.text()

//line machine.rl:70

	key = m.text()

	goto st303
	st303:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof303
		}
	st_case_303:
//line machine.go:37769
		switch ( m.data)[( m.p)] {
		case 32:
			goto st6
		case 34:
			goto tr586
		case 44:
			goto st6
		case 45:
			goto tr587
		case 46:
			goto tr588
		case 48:
			goto tr589
		case 61:
			goto st6
		case 70:
			goto tr591
		case 84:
			goto tr592
		case 92:
			goto tr514
		case 102:
			goto tr593
		case 116:
			goto tr594
		}
		switch {
		case ( m.data)[( m.p)] < 12:
			if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 10 {
				goto st6
			}
		case ( m.data)[( m.p)] > 13:
			if 49 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto tr590
			}
		default:
			goto st6
		}
		goto tr509
tr586:
//line machine.rl:86

	m.handler.AddString(key, m.text())

//line machine.rl:17

	m.pb = m.p

	goto st1146
	st1146:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1146
		}
	st_case_1146:
//line machine.go:37824
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1398
		case 11:
			goto tr1667
		case 13:
			goto tr1398
		case 32:
			goto tr1666
		case 34:
			goto tr132
		case 44:
			goto tr1668
		case 61:
			goto tr22
		case 92:
			goto tr134
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1666
		}
		goto tr129
tr1672:
//line machine.rl:66

	m.handler.AddTag(key, m.text())

	goto st1147
tr1666:
//line machine.rl:66

	m.handler.AddTag(key, m.text())

//line machine.rl:17

	m.pb = m.p

	goto st1147
tr1696:
//line machine.rl:78

	m.handler.AddFloat(key, m.text())

//line machine.rl:66

	m.handler.AddTag(key, m.text())

	goto st1147
tr1701:
//line machine.rl:74

	m.handler.AddInt(key, m.text())

//line machine.rl:66

	m.handler.AddTag(key, m.text())

	goto st1147
tr1704:
//line machine.rl:82

	m.handler.AddBool(key, m.text())

//line machine.rl:66

	m.handler.AddTag(key, m.text())

	goto st1147
	st1147:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1147
		}
	st_case_1147:
//line machine.go:37898
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr927
		case 11:
			goto tr1670
		case 13:
			goto tr927
		case 32:
			goto st1147
		case 34:
			goto tr82
		case 44:
			goto st6
		case 45:
			goto tr1671
		case 61:
			goto st6
		case 92:
			goto tr144
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto tr1061
			}
		case ( m.data)[( m.p)] >= 9:
			goto st1147
		}
		goto tr141
tr1670:
//line machine.rl:17

	m.pb = m.p

	goto st1148
	st1148:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1148
		}
	st_case_1148:
//line machine.go:37939
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr927
		case 11:
			goto tr1670
		case 13:
			goto tr927
		case 32:
			goto st1147
		case 34:
			goto tr82
		case 44:
			goto st6
		case 45:
			goto tr1671
		case 61:
			goto tr146
		case 92:
			goto tr144
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto tr1061
			}
		case ( m.data)[( m.p)] >= 9:
			goto st1147
		}
		goto tr141
tr1671:
//line machine.rl:17

	m.pb = m.p

	goto st304
	st304:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof304
		}
	st_case_304:
//line machine.go:37980
		switch ( m.data)[( m.p)] {
		case 32:
			goto st6
		case 34:
			goto tr85
		case 44:
			goto st6
		case 61:
			goto tr146
		case 92:
			goto st83
		}
		switch {
		case ( m.data)[( m.p)] < 12:
			if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 10 {
				goto st6
			}
		case ( m.data)[( m.p)] > 13:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st653
			}
		default:
			goto st6
		}
		goto st43
tr1667:
//line machine.rl:66

	m.handler.AddTag(key, m.text())

//line machine.rl:17

	m.pb = m.p

	goto st1149
tr1697:
//line machine.rl:78

	m.handler.AddFloat(key, m.text())

//line machine.rl:66

	m.handler.AddTag(key, m.text())

	goto st1149
tr1702:
//line machine.rl:74

	m.handler.AddInt(key, m.text())

//line machine.rl:66

	m.handler.AddTag(key, m.text())

	goto st1149
tr1705:
//line machine.rl:82

	m.handler.AddBool(key, m.text())

//line machine.rl:66

	m.handler.AddTag(key, m.text())

	goto st1149
	st1149:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1149
		}
	st_case_1149:
//line machine.go:38051
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr927
		case 11:
			goto tr1673
		case 13:
			goto tr927
		case 32:
			goto tr1672
		case 34:
			goto tr205
		case 44:
			goto tr139
		case 45:
			goto tr1674
		case 61:
			goto st6
		case 92:
			goto tr246
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto tr1675
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1672
		}
		goto tr244
tr1673:
//line machine.rl:66

	m.handler.AddTag(key, m.text())

//line machine.rl:17

	m.pb = m.p

	goto st1150
	st1150:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1150
		}
	st_case_1150:
//line machine.go:38096
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr927
		case 11:
			goto tr1673
		case 13:
			goto tr927
		case 32:
			goto tr1672
		case 34:
			goto tr205
		case 44:
			goto tr139
		case 45:
			goto tr1674
		case 61:
			goto tr146
		case 92:
			goto tr246
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto tr1675
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1672
		}
		goto tr244
tr1674:
//line machine.rl:17

	m.pb = m.p

	goto st305
	st305:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof305
		}
	st_case_305:
//line machine.go:38137
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr242
		case 13:
			goto st6
		case 32:
			goto tr136
		case 34:
			goto tr209
		case 44:
			goto tr139
		case 61:
			goto tr146
		case 92:
			goto st86
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1151
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr136
		}
		goto st84
tr1675:
//line machine.rl:17

	m.pb = m.p

	goto st1151
	st1151:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1151
		}
	st_case_1151:
//line machine.go:38176
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1063
		case 11:
			goto tr1677
		case 13:
			goto tr1063
		case 32:
			goto tr1676
		case 34:
			goto tr209
		case 44:
			goto tr139
		case 61:
			goto tr146
		case 92:
			goto st86
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1152
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1676
		}
		goto st84
	st1152:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1152
		}
	st_case_1152:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1063
		case 11:
			goto tr1677
		case 13:
			goto tr1063
		case 32:
			goto tr1676
		case 34:
			goto tr209
		case 44:
			goto tr139
		case 61:
			goto tr146
		case 92:
			goto st86
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1153
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1676
		}
		goto st84
	st1153:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1153
		}
	st_case_1153:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1063
		case 11:
			goto tr1677
		case 13:
			goto tr1063
		case 32:
			goto tr1676
		case 34:
			goto tr209
		case 44:
			goto tr139
		case 61:
			goto tr146
		case 92:
			goto st86
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1154
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1676
		}
		goto st84
	st1154:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1154
		}
	st_case_1154:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1063
		case 11:
			goto tr1677
		case 13:
			goto tr1063
		case 32:
			goto tr1676
		case 34:
			goto tr209
		case 44:
			goto tr139
		case 61:
			goto tr146
		case 92:
			goto st86
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1155
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1676
		}
		goto st84
	st1155:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1155
		}
	st_case_1155:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1063
		case 11:
			goto tr1677
		case 13:
			goto tr1063
		case 32:
			goto tr1676
		case 34:
			goto tr209
		case 44:
			goto tr139
		case 61:
			goto tr146
		case 92:
			goto st86
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1156
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1676
		}
		goto st84
	st1156:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1156
		}
	st_case_1156:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1063
		case 11:
			goto tr1677
		case 13:
			goto tr1063
		case 32:
			goto tr1676
		case 34:
			goto tr209
		case 44:
			goto tr139
		case 61:
			goto tr146
		case 92:
			goto st86
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1157
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1676
		}
		goto st84
	st1157:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1157
		}
	st_case_1157:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1063
		case 11:
			goto tr1677
		case 13:
			goto tr1063
		case 32:
			goto tr1676
		case 34:
			goto tr209
		case 44:
			goto tr139
		case 61:
			goto tr146
		case 92:
			goto st86
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1158
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1676
		}
		goto st84
	st1158:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1158
		}
	st_case_1158:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1063
		case 11:
			goto tr1677
		case 13:
			goto tr1063
		case 32:
			goto tr1676
		case 34:
			goto tr209
		case 44:
			goto tr139
		case 61:
			goto tr146
		case 92:
			goto st86
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1159
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1676
		}
		goto st84
	st1159:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1159
		}
	st_case_1159:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1063
		case 11:
			goto tr1677
		case 13:
			goto tr1063
		case 32:
			goto tr1676
		case 34:
			goto tr209
		case 44:
			goto tr139
		case 61:
			goto tr146
		case 92:
			goto st86
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1160
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1676
		}
		goto st84
	st1160:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1160
		}
	st_case_1160:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1063
		case 11:
			goto tr1677
		case 13:
			goto tr1063
		case 32:
			goto tr1676
		case 34:
			goto tr209
		case 44:
			goto tr139
		case 61:
			goto tr146
		case 92:
			goto st86
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1161
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1676
		}
		goto st84
	st1161:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1161
		}
	st_case_1161:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1063
		case 11:
			goto tr1677
		case 13:
			goto tr1063
		case 32:
			goto tr1676
		case 34:
			goto tr209
		case 44:
			goto tr139
		case 61:
			goto tr146
		case 92:
			goto st86
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1162
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1676
		}
		goto st84
	st1162:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1162
		}
	st_case_1162:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1063
		case 11:
			goto tr1677
		case 13:
			goto tr1063
		case 32:
			goto tr1676
		case 34:
			goto tr209
		case 44:
			goto tr139
		case 61:
			goto tr146
		case 92:
			goto st86
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1163
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1676
		}
		goto st84
	st1163:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1163
		}
	st_case_1163:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1063
		case 11:
			goto tr1677
		case 13:
			goto tr1063
		case 32:
			goto tr1676
		case 34:
			goto tr209
		case 44:
			goto tr139
		case 61:
			goto tr146
		case 92:
			goto st86
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1164
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1676
		}
		goto st84
	st1164:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1164
		}
	st_case_1164:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1063
		case 11:
			goto tr1677
		case 13:
			goto tr1063
		case 32:
			goto tr1676
		case 34:
			goto tr209
		case 44:
			goto tr139
		case 61:
			goto tr146
		case 92:
			goto st86
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1165
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1676
		}
		goto st84
	st1165:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1165
		}
	st_case_1165:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1063
		case 11:
			goto tr1677
		case 13:
			goto tr1063
		case 32:
			goto tr1676
		case 34:
			goto tr209
		case 44:
			goto tr139
		case 61:
			goto tr146
		case 92:
			goto st86
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1166
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1676
		}
		goto st84
	st1166:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1166
		}
	st_case_1166:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1063
		case 11:
			goto tr1677
		case 13:
			goto tr1063
		case 32:
			goto tr1676
		case 34:
			goto tr209
		case 44:
			goto tr139
		case 61:
			goto tr146
		case 92:
			goto st86
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1167
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1676
		}
		goto st84
	st1167:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1167
		}
	st_case_1167:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1063
		case 11:
			goto tr1677
		case 13:
			goto tr1063
		case 32:
			goto tr1676
		case 34:
			goto tr209
		case 44:
			goto tr139
		case 61:
			goto tr146
		case 92:
			goto st86
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1168
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1676
		}
		goto st84
	st1168:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1168
		}
	st_case_1168:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1063
		case 11:
			goto tr1677
		case 13:
			goto tr1063
		case 32:
			goto tr1676
		case 34:
			goto tr209
		case 44:
			goto tr139
		case 61:
			goto tr146
		case 92:
			goto st86
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1169
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1676
		}
		goto st84
	st1169:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1169
		}
	st_case_1169:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1063
		case 11:
			goto tr1677
		case 13:
			goto tr1063
		case 32:
			goto tr1676
		case 34:
			goto tr209
		case 44:
			goto tr139
		case 61:
			goto tr146
		case 92:
			goto st86
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1676
		}
		goto st84
tr1668:
//line machine.rl:66

	m.handler.AddTag(key, m.text())

//line machine.rl:17

	m.pb = m.p

	goto st306
tr1698:
//line machine.rl:78

	m.handler.AddFloat(key, m.text())

//line machine.rl:66

	m.handler.AddTag(key, m.text())

	goto st306
tr1703:
//line machine.rl:74

	m.handler.AddInt(key, m.text())

//line machine.rl:66

	m.handler.AddTag(key, m.text())

	goto st306
tr1706:
//line machine.rl:82

	m.handler.AddBool(key, m.text())

//line machine.rl:66

	m.handler.AddTag(key, m.text())

	goto st306
	st306:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof306
		}
	st_case_306:
//line machine.go:38820
		switch ( m.data)[( m.p)] {
		case 32:
			goto st6
		case 34:
			goto tr559
		case 44:
			goto st6
		case 61:
			goto st6
		case 92:
			goto tr597
		}
		switch {
		case ( m.data)[( m.p)] > 10:
			if 12 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto st6
			}
		case ( m.data)[( m.p)] >= 9:
			goto st6
		}
		goto tr596
tr596:
//line machine.rl:17

	m.pb = m.p

	goto st307
	st307:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof307
		}
	st_case_307:
//line machine.go:38853
		switch ( m.data)[( m.p)] {
		case 32:
			goto st6
		case 34:
			goto tr562
		case 44:
			goto st6
		case 61:
			goto tr599
		case 92:
			goto st318
		}
		switch {
		case ( m.data)[( m.p)] > 10:
			if 12 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto st6
			}
		case ( m.data)[( m.p)] >= 9:
			goto st6
		}
		goto st307
tr599:
//line machine.rl:70

	key = m.text()

//line machine.rl:62

	key = m.text()

	goto st308
	st308:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof308
		}
	st_case_308:
//line machine.go:38890
		switch ( m.data)[( m.p)] {
		case 32:
			goto st6
		case 34:
			goto tr567
		case 44:
			goto st6
		case 45:
			goto tr601
		case 46:
			goto tr602
		case 48:
			goto tr603
		case 61:
			goto st6
		case 70:
			goto tr605
		case 84:
			goto tr606
		case 92:
			goto tr134
		case 102:
			goto tr607
		case 116:
			goto tr608
		}
		switch {
		case ( m.data)[( m.p)] < 12:
			if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 10 {
				goto st6
			}
		case ( m.data)[( m.p)] > 13:
			if 49 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto tr604
			}
		default:
			goto st6
		}
		goto tr129
tr601:
//line machine.rl:17

	m.pb = m.p

	goto st309
	st309:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof309
		}
	st_case_309:
//line machine.go:38941
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr137
		case 13:
			goto st6
		case 32:
			goto tr136
		case 34:
			goto tr138
		case 44:
			goto tr139
		case 48:
			goto st1170
		case 61:
			goto st6
		case 92:
			goto st55
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 49 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1173
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr136
		}
		goto st41
tr603:
//line machine.rl:17

	m.pb = m.p

	goto st1170
	st1170:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1170
		}
	st_case_1170:
//line machine.go:38982
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr922
		case 11:
			goto tr1697
		case 13:
			goto tr922
		case 32:
			goto tr1696
		case 34:
			goto tr138
		case 44:
			goto tr1698
		case 46:
			goto st1171
		case 61:
			goto st6
		case 92:
			goto st55
		case 105:
			goto st1172
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1696
		}
		goto st41
tr602:
//line machine.rl:17

	m.pb = m.p

	goto st1171
	st1171:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1171
		}
	st_case_1171:
//line machine.go:39020
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr922
		case 11:
			goto tr1697
		case 13:
			goto tr922
		case 32:
			goto tr1696
		case 34:
			goto tr138
		case 44:
			goto tr1698
		case 61:
			goto st6
		case 92:
			goto st55
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1171
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1696
		}
		goto st41
	st1172:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1172
		}
	st_case_1172:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1119
		case 11:
			goto tr1702
		case 13:
			goto tr1119
		case 32:
			goto tr1701
		case 34:
			goto tr138
		case 44:
			goto tr1703
		case 61:
			goto st6
		case 92:
			goto st55
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1701
		}
		goto st41
tr604:
//line machine.rl:17

	m.pb = m.p

	goto st1173
	st1173:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1173
		}
	st_case_1173:
//line machine.go:39086
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr922
		case 11:
			goto tr1697
		case 13:
			goto tr922
		case 32:
			goto tr1696
		case 34:
			goto tr138
		case 44:
			goto tr1698
		case 46:
			goto st1171
		case 61:
			goto st6
		case 92:
			goto st55
		case 105:
			goto st1172
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1173
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1696
		}
		goto st41
tr605:
//line machine.rl:17

	m.pb = m.p

	goto st1174
	st1174:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1174
		}
	st_case_1174:
//line machine.go:39129
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1123
		case 11:
			goto tr1705
		case 13:
			goto tr1123
		case 32:
			goto tr1704
		case 34:
			goto tr138
		case 44:
			goto tr1706
		case 61:
			goto st6
		case 65:
			goto st310
		case 92:
			goto st55
		case 97:
			goto st313
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1704
		}
		goto st41
	st310:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof310
		}
	st_case_310:
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr137
		case 13:
			goto st6
		case 32:
			goto tr136
		case 34:
			goto tr138
		case 44:
			goto tr139
		case 61:
			goto st6
		case 76:
			goto st311
		case 92:
			goto st55
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr136
		}
		goto st41
	st311:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof311
		}
	st_case_311:
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr137
		case 13:
			goto st6
		case 32:
			goto tr136
		case 34:
			goto tr138
		case 44:
			goto tr139
		case 61:
			goto st6
		case 83:
			goto st312
		case 92:
			goto st55
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr136
		}
		goto st41
	st312:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof312
		}
	st_case_312:
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr137
		case 13:
			goto st6
		case 32:
			goto tr136
		case 34:
			goto tr138
		case 44:
			goto tr139
		case 61:
			goto st6
		case 69:
			goto st1175
		case 92:
			goto st55
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr136
		}
		goto st41
	st1175:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1175
		}
	st_case_1175:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1123
		case 11:
			goto tr1705
		case 13:
			goto tr1123
		case 32:
			goto tr1704
		case 34:
			goto tr138
		case 44:
			goto tr1706
		case 61:
			goto st6
		case 92:
			goto st55
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1704
		}
		goto st41
	st313:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof313
		}
	st_case_313:
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr137
		case 13:
			goto st6
		case 32:
			goto tr136
		case 34:
			goto tr138
		case 44:
			goto tr139
		case 61:
			goto st6
		case 92:
			goto st55
		case 108:
			goto st314
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr136
		}
		goto st41
	st314:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof314
		}
	st_case_314:
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr137
		case 13:
			goto st6
		case 32:
			goto tr136
		case 34:
			goto tr138
		case 44:
			goto tr139
		case 61:
			goto st6
		case 92:
			goto st55
		case 115:
			goto st315
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr136
		}
		goto st41
	st315:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof315
		}
	st_case_315:
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr137
		case 13:
			goto st6
		case 32:
			goto tr136
		case 34:
			goto tr138
		case 44:
			goto tr139
		case 61:
			goto st6
		case 92:
			goto st55
		case 101:
			goto st1175
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr136
		}
		goto st41
tr606:
//line machine.rl:17

	m.pb = m.p

	goto st1176
	st1176:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1176
		}
	st_case_1176:
//line machine.go:39368
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1123
		case 11:
			goto tr1705
		case 13:
			goto tr1123
		case 32:
			goto tr1704
		case 34:
			goto tr138
		case 44:
			goto tr1706
		case 61:
			goto st6
		case 82:
			goto st316
		case 92:
			goto st55
		case 114:
			goto st317
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1704
		}
		goto st41
	st316:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof316
		}
	st_case_316:
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr137
		case 13:
			goto st6
		case 32:
			goto tr136
		case 34:
			goto tr138
		case 44:
			goto tr139
		case 61:
			goto st6
		case 85:
			goto st312
		case 92:
			goto st55
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr136
		}
		goto st41
	st317:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof317
		}
	st_case_317:
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr137
		case 13:
			goto st6
		case 32:
			goto tr136
		case 34:
			goto tr138
		case 44:
			goto tr139
		case 61:
			goto st6
		case 92:
			goto st55
		case 117:
			goto st315
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr136
		}
		goto st41
tr607:
//line machine.rl:17

	m.pb = m.p

	goto st1177
	st1177:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1177
		}
	st_case_1177:
//line machine.go:39464
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1123
		case 11:
			goto tr1705
		case 13:
			goto tr1123
		case 32:
			goto tr1704
		case 34:
			goto tr138
		case 44:
			goto tr1706
		case 61:
			goto st6
		case 92:
			goto st55
		case 97:
			goto st313
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1704
		}
		goto st41
tr608:
//line machine.rl:17

	m.pb = m.p

	goto st1178
	st1178:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1178
		}
	st_case_1178:
//line machine.go:39500
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1123
		case 11:
			goto tr1705
		case 13:
			goto tr1123
		case 32:
			goto tr1704
		case 34:
			goto tr138
		case 44:
			goto tr1706
		case 61:
			goto st6
		case 92:
			goto st55
		case 114:
			goto st317
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1704
		}
		goto st41
tr597:
//line machine.rl:17

	m.pb = m.p

	goto st318
	st318:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof318
		}
	st_case_318:
//line machine.go:39536
		switch ( m.data)[( m.p)] {
		case 34:
			goto st307
		case 92:
			goto st307
		}
		switch {
		case ( m.data)[( m.p)] > 10:
			if 12 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto tr495
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr495
		}
		goto st266
tr587:
//line machine.rl:17

	m.pb = m.p

	goto st319
	st319:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof319
		}
	st_case_319:
//line machine.go:39563
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr517
		case 13:
			goto st6
		case 32:
			goto tr516
		case 34:
			goto tr518
		case 44:
			goto tr519
		case 48:
			goto st1179
		case 61:
			goto st6
		case 92:
			goto st278
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 49 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1182
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr516
		}
		goto st269
tr589:
//line machine.rl:17

	m.pb = m.p

	goto st1179
	st1179:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1179
		}
	st_case_1179:
//line machine.go:39604
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1415
		case 11:
			goto tr1712
		case 13:
			goto tr1415
		case 32:
			goto tr1711
		case 34:
			goto tr518
		case 44:
			goto tr1713
		case 46:
			goto st1180
		case 61:
			goto st6
		case 92:
			goto st278
		case 105:
			goto st1181
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1711
		}
		goto st269
tr588:
//line machine.rl:17

	m.pb = m.p

	goto st1180
	st1180:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1180
		}
	st_case_1180:
//line machine.go:39642
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1415
		case 11:
			goto tr1712
		case 13:
			goto tr1415
		case 32:
			goto tr1711
		case 34:
			goto tr518
		case 44:
			goto tr1713
		case 61:
			goto st6
		case 92:
			goto st278
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1180
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1711
		}
		goto st269
	st1181:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1181
		}
	st_case_1181:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1420
		case 11:
			goto tr1717
		case 13:
			goto tr1420
		case 32:
			goto tr1716
		case 34:
			goto tr518
		case 44:
			goto tr1718
		case 61:
			goto st6
		case 92:
			goto st278
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1716
		}
		goto st269
tr590:
//line machine.rl:17

	m.pb = m.p

	goto st1182
	st1182:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1182
		}
	st_case_1182:
//line machine.go:39708
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1415
		case 11:
			goto tr1712
		case 13:
			goto tr1415
		case 32:
			goto tr1711
		case 34:
			goto tr518
		case 44:
			goto tr1713
		case 46:
			goto st1180
		case 61:
			goto st6
		case 92:
			goto st278
		case 105:
			goto st1181
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1182
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1711
		}
		goto st269
tr591:
//line machine.rl:17

	m.pb = m.p

	goto st1183
	st1183:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1183
		}
	st_case_1183:
//line machine.go:39751
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1423
		case 11:
			goto tr1720
		case 13:
			goto tr1423
		case 32:
			goto tr1719
		case 34:
			goto tr518
		case 44:
			goto tr1721
		case 61:
			goto st6
		case 65:
			goto st320
		case 92:
			goto st278
		case 97:
			goto st323
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1719
		}
		goto st269
	st320:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof320
		}
	st_case_320:
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr517
		case 13:
			goto st6
		case 32:
			goto tr516
		case 34:
			goto tr518
		case 44:
			goto tr519
		case 61:
			goto st6
		case 76:
			goto st321
		case 92:
			goto st278
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr516
		}
		goto st269
	st321:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof321
		}
	st_case_321:
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr517
		case 13:
			goto st6
		case 32:
			goto tr516
		case 34:
			goto tr518
		case 44:
			goto tr519
		case 61:
			goto st6
		case 83:
			goto st322
		case 92:
			goto st278
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr516
		}
		goto st269
	st322:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof322
		}
	st_case_322:
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr517
		case 13:
			goto st6
		case 32:
			goto tr516
		case 34:
			goto tr518
		case 44:
			goto tr519
		case 61:
			goto st6
		case 69:
			goto st1184
		case 92:
			goto st278
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr516
		}
		goto st269
	st1184:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1184
		}
	st_case_1184:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1423
		case 11:
			goto tr1720
		case 13:
			goto tr1423
		case 32:
			goto tr1719
		case 34:
			goto tr518
		case 44:
			goto tr1721
		case 61:
			goto st6
		case 92:
			goto st278
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1719
		}
		goto st269
	st323:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof323
		}
	st_case_323:
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr517
		case 13:
			goto st6
		case 32:
			goto tr516
		case 34:
			goto tr518
		case 44:
			goto tr519
		case 61:
			goto st6
		case 92:
			goto st278
		case 108:
			goto st324
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr516
		}
		goto st269
	st324:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof324
		}
	st_case_324:
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr517
		case 13:
			goto st6
		case 32:
			goto tr516
		case 34:
			goto tr518
		case 44:
			goto tr519
		case 61:
			goto st6
		case 92:
			goto st278
		case 115:
			goto st325
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr516
		}
		goto st269
	st325:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof325
		}
	st_case_325:
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr517
		case 13:
			goto st6
		case 32:
			goto tr516
		case 34:
			goto tr518
		case 44:
			goto tr519
		case 61:
			goto st6
		case 92:
			goto st278
		case 101:
			goto st1184
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr516
		}
		goto st269
tr592:
//line machine.rl:17

	m.pb = m.p

	goto st1185
	st1185:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1185
		}
	st_case_1185:
//line machine.go:39990
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1423
		case 11:
			goto tr1720
		case 13:
			goto tr1423
		case 32:
			goto tr1719
		case 34:
			goto tr518
		case 44:
			goto tr1721
		case 61:
			goto st6
		case 82:
			goto st326
		case 92:
			goto st278
		case 114:
			goto st327
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1719
		}
		goto st269
	st326:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof326
		}
	st_case_326:
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr517
		case 13:
			goto st6
		case 32:
			goto tr516
		case 34:
			goto tr518
		case 44:
			goto tr519
		case 61:
			goto st6
		case 85:
			goto st322
		case 92:
			goto st278
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr516
		}
		goto st269
	st327:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof327
		}
	st_case_327:
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr517
		case 13:
			goto st6
		case 32:
			goto tr516
		case 34:
			goto tr518
		case 44:
			goto tr519
		case 61:
			goto st6
		case 92:
			goto st278
		case 117:
			goto st325
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr516
		}
		goto st269
tr593:
//line machine.rl:17

	m.pb = m.p

	goto st1186
	st1186:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1186
		}
	st_case_1186:
//line machine.go:40086
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1423
		case 11:
			goto tr1720
		case 13:
			goto tr1423
		case 32:
			goto tr1719
		case 34:
			goto tr518
		case 44:
			goto tr1721
		case 61:
			goto st6
		case 92:
			goto st278
		case 97:
			goto st323
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1719
		}
		goto st269
tr594:
//line machine.rl:17

	m.pb = m.p

	goto st1187
	st1187:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1187
		}
	st_case_1187:
//line machine.go:40122
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1423
		case 11:
			goto tr1720
		case 13:
			goto tr1423
		case 32:
			goto tr1719
		case 34:
			goto tr518
		case 44:
			goto tr1721
		case 61:
			goto st6
		case 92:
			goto st278
		case 114:
			goto st327
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1719
		}
		goto st269
tr580:
//line machine.rl:17

	m.pb = m.p

	goto st328
	st328:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof328
		}
	st_case_328:
//line machine.go:40158
		switch ( m.data)[( m.p)] {
		case 34:
			goto st301
		case 92:
			goto st301
		}
		switch {
		case ( m.data)[( m.p)] > 10:
			if 12 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto tr495
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr495
		}
		goto st38
tr568:
//line machine.rl:17

	m.pb = m.p

	goto st329
	st329:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof329
		}
	st_case_329:
//line machine.go:40185
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr200
		case 13:
			goto st6
		case 32:
			goto tr199
		case 34:
			goto tr138
		case 44:
			goto tr201
		case 48:
			goto st1188
		case 61:
			goto st6
		case 92:
			goto st69
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 49 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1212
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr199
		}
		goto st64
tr570:
//line machine.rl:17

	m.pb = m.p

	goto st1188
	st1188:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1188
		}
	st_case_1188:
//line machine.go:40226
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1026
		case 11:
			goto tr1727
		case 13:
			goto tr1026
		case 32:
			goto tr1726
		case 34:
			goto tr138
		case 44:
			goto tr1728
		case 46:
			goto st1210
		case 61:
			goto st6
		case 92:
			goto st69
		case 105:
			goto st1211
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1726
		}
		goto st64
tr1727:
//line machine.rl:78

	m.handler.AddFloat(key, m.text())

//line machine.rl:66

	m.handler.AddTag(key, m.text())

	goto st1189
tr1756:
//line machine.rl:74

	m.handler.AddInt(key, m.text())

//line machine.rl:66

	m.handler.AddTag(key, m.text())

	goto st1189
tr1759:
//line machine.rl:82

	m.handler.AddBool(key, m.text())

//line machine.rl:66

	m.handler.AddTag(key, m.text())

	goto st1189
	st1189:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1189
		}
	st_case_1189:
//line machine.go:40288
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr885
		case 11:
			goto tr1732
		case 13:
			goto tr885
		case 32:
			goto tr1731
		case 34:
			goto tr205
		case 44:
			goto tr201
		case 45:
			goto tr1733
		case 61:
			goto st6
		case 92:
			goto tr206
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto tr1734
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1731
		}
		goto tr203
tr1732:
//line machine.rl:66

	m.handler.AddTag(key, m.text())

//line machine.rl:17

	m.pb = m.p

	goto st1190
	st1190:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1190
		}
	st_case_1190:
//line machine.go:40333
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr885
		case 11:
			goto tr1732
		case 13:
			goto tr885
		case 32:
			goto tr1731
		case 34:
			goto tr205
		case 44:
			goto tr201
		case 45:
			goto tr1733
		case 61:
			goto tr86
		case 92:
			goto tr206
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto tr1734
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1731
		}
		goto tr203
tr1733:
//line machine.rl:17

	m.pb = m.p

	goto st330
	st330:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof330
		}
	st_case_330:
//line machine.go:40374
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr208
		case 13:
			goto st6
		case 32:
			goto tr199
		case 34:
			goto tr209
		case 44:
			goto tr201
		case 61:
			goto tr86
		case 92:
			goto st68
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1191
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr199
		}
		goto st66
tr1734:
//line machine.rl:17

	m.pb = m.p

	goto st1191
	st1191:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1191
		}
	st_case_1191:
//line machine.go:40413
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr990
		case 11:
			goto tr1736
		case 13:
			goto tr990
		case 32:
			goto tr1735
		case 34:
			goto tr209
		case 44:
			goto tr201
		case 61:
			goto tr86
		case 92:
			goto st68
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1192
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1735
		}
		goto st66
	st1192:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1192
		}
	st_case_1192:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr990
		case 11:
			goto tr1736
		case 13:
			goto tr990
		case 32:
			goto tr1735
		case 34:
			goto tr209
		case 44:
			goto tr201
		case 61:
			goto tr86
		case 92:
			goto st68
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1193
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1735
		}
		goto st66
	st1193:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1193
		}
	st_case_1193:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr990
		case 11:
			goto tr1736
		case 13:
			goto tr990
		case 32:
			goto tr1735
		case 34:
			goto tr209
		case 44:
			goto tr201
		case 61:
			goto tr86
		case 92:
			goto st68
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1194
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1735
		}
		goto st66
	st1194:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1194
		}
	st_case_1194:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr990
		case 11:
			goto tr1736
		case 13:
			goto tr990
		case 32:
			goto tr1735
		case 34:
			goto tr209
		case 44:
			goto tr201
		case 61:
			goto tr86
		case 92:
			goto st68
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1195
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1735
		}
		goto st66
	st1195:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1195
		}
	st_case_1195:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr990
		case 11:
			goto tr1736
		case 13:
			goto tr990
		case 32:
			goto tr1735
		case 34:
			goto tr209
		case 44:
			goto tr201
		case 61:
			goto tr86
		case 92:
			goto st68
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1196
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1735
		}
		goto st66
	st1196:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1196
		}
	st_case_1196:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr990
		case 11:
			goto tr1736
		case 13:
			goto tr990
		case 32:
			goto tr1735
		case 34:
			goto tr209
		case 44:
			goto tr201
		case 61:
			goto tr86
		case 92:
			goto st68
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1197
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1735
		}
		goto st66
	st1197:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1197
		}
	st_case_1197:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr990
		case 11:
			goto tr1736
		case 13:
			goto tr990
		case 32:
			goto tr1735
		case 34:
			goto tr209
		case 44:
			goto tr201
		case 61:
			goto tr86
		case 92:
			goto st68
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1198
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1735
		}
		goto st66
	st1198:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1198
		}
	st_case_1198:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr990
		case 11:
			goto tr1736
		case 13:
			goto tr990
		case 32:
			goto tr1735
		case 34:
			goto tr209
		case 44:
			goto tr201
		case 61:
			goto tr86
		case 92:
			goto st68
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1199
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1735
		}
		goto st66
	st1199:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1199
		}
	st_case_1199:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr990
		case 11:
			goto tr1736
		case 13:
			goto tr990
		case 32:
			goto tr1735
		case 34:
			goto tr209
		case 44:
			goto tr201
		case 61:
			goto tr86
		case 92:
			goto st68
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1200
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1735
		}
		goto st66
	st1200:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1200
		}
	st_case_1200:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr990
		case 11:
			goto tr1736
		case 13:
			goto tr990
		case 32:
			goto tr1735
		case 34:
			goto tr209
		case 44:
			goto tr201
		case 61:
			goto tr86
		case 92:
			goto st68
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1201
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1735
		}
		goto st66
	st1201:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1201
		}
	st_case_1201:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr990
		case 11:
			goto tr1736
		case 13:
			goto tr990
		case 32:
			goto tr1735
		case 34:
			goto tr209
		case 44:
			goto tr201
		case 61:
			goto tr86
		case 92:
			goto st68
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1202
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1735
		}
		goto st66
	st1202:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1202
		}
	st_case_1202:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr990
		case 11:
			goto tr1736
		case 13:
			goto tr990
		case 32:
			goto tr1735
		case 34:
			goto tr209
		case 44:
			goto tr201
		case 61:
			goto tr86
		case 92:
			goto st68
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1203
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1735
		}
		goto st66
	st1203:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1203
		}
	st_case_1203:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr990
		case 11:
			goto tr1736
		case 13:
			goto tr990
		case 32:
			goto tr1735
		case 34:
			goto tr209
		case 44:
			goto tr201
		case 61:
			goto tr86
		case 92:
			goto st68
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1204
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1735
		}
		goto st66
	st1204:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1204
		}
	st_case_1204:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr990
		case 11:
			goto tr1736
		case 13:
			goto tr990
		case 32:
			goto tr1735
		case 34:
			goto tr209
		case 44:
			goto tr201
		case 61:
			goto tr86
		case 92:
			goto st68
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1205
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1735
		}
		goto st66
	st1205:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1205
		}
	st_case_1205:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr990
		case 11:
			goto tr1736
		case 13:
			goto tr990
		case 32:
			goto tr1735
		case 34:
			goto tr209
		case 44:
			goto tr201
		case 61:
			goto tr86
		case 92:
			goto st68
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1206
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1735
		}
		goto st66
	st1206:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1206
		}
	st_case_1206:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr990
		case 11:
			goto tr1736
		case 13:
			goto tr990
		case 32:
			goto tr1735
		case 34:
			goto tr209
		case 44:
			goto tr201
		case 61:
			goto tr86
		case 92:
			goto st68
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1207
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1735
		}
		goto st66
	st1207:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1207
		}
	st_case_1207:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr990
		case 11:
			goto tr1736
		case 13:
			goto tr990
		case 32:
			goto tr1735
		case 34:
			goto tr209
		case 44:
			goto tr201
		case 61:
			goto tr86
		case 92:
			goto st68
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1208
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1735
		}
		goto st66
	st1208:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1208
		}
	st_case_1208:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr990
		case 11:
			goto tr1736
		case 13:
			goto tr990
		case 32:
			goto tr1735
		case 34:
			goto tr209
		case 44:
			goto tr201
		case 61:
			goto tr86
		case 92:
			goto st68
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1209
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1735
		}
		goto st66
	st1209:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1209
		}
	st_case_1209:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr990
		case 11:
			goto tr1736
		case 13:
			goto tr990
		case 32:
			goto tr1735
		case 34:
			goto tr209
		case 44:
			goto tr201
		case 61:
			goto tr86
		case 92:
			goto st68
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1735
		}
		goto st66
tr569:
//line machine.rl:17

	m.pb = m.p

	goto st1210
	st1210:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1210
		}
	st_case_1210:
//line machine.go:41023
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1026
		case 11:
			goto tr1727
		case 13:
			goto tr1026
		case 32:
			goto tr1726
		case 34:
			goto tr138
		case 44:
			goto tr1728
		case 61:
			goto st6
		case 92:
			goto st69
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1210
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1726
		}
		goto st64
	st1211:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1211
		}
	st_case_1211:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1162
		case 11:
			goto tr1756
		case 13:
			goto tr1162
		case 32:
			goto tr1755
		case 34:
			goto tr138
		case 44:
			goto tr1757
		case 61:
			goto st6
		case 92:
			goto st69
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1755
		}
		goto st64
tr571:
//line machine.rl:17

	m.pb = m.p

	goto st1212
	st1212:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1212
		}
	st_case_1212:
//line machine.go:41089
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1026
		case 11:
			goto tr1727
		case 13:
			goto tr1026
		case 32:
			goto tr1726
		case 34:
			goto tr138
		case 44:
			goto tr1728
		case 46:
			goto st1210
		case 61:
			goto st6
		case 92:
			goto st69
		case 105:
			goto st1211
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1212
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1726
		}
		goto st64
tr572:
//line machine.rl:17

	m.pb = m.p

	goto st1213
	st1213:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1213
		}
	st_case_1213:
//line machine.go:41132
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1166
		case 11:
			goto tr1759
		case 13:
			goto tr1166
		case 32:
			goto tr1758
		case 34:
			goto tr138
		case 44:
			goto tr1760
		case 61:
			goto st6
		case 65:
			goto st331
		case 92:
			goto st69
		case 97:
			goto st334
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1758
		}
		goto st64
	st331:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof331
		}
	st_case_331:
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr200
		case 13:
			goto st6
		case 32:
			goto tr199
		case 34:
			goto tr138
		case 44:
			goto tr201
		case 61:
			goto st6
		case 76:
			goto st332
		case 92:
			goto st69
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr199
		}
		goto st64
	st332:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof332
		}
	st_case_332:
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr200
		case 13:
			goto st6
		case 32:
			goto tr199
		case 34:
			goto tr138
		case 44:
			goto tr201
		case 61:
			goto st6
		case 83:
			goto st333
		case 92:
			goto st69
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr199
		}
		goto st64
	st333:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof333
		}
	st_case_333:
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr200
		case 13:
			goto st6
		case 32:
			goto tr199
		case 34:
			goto tr138
		case 44:
			goto tr201
		case 61:
			goto st6
		case 69:
			goto st1214
		case 92:
			goto st69
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr199
		}
		goto st64
	st1214:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1214
		}
	st_case_1214:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1166
		case 11:
			goto tr1759
		case 13:
			goto tr1166
		case 32:
			goto tr1758
		case 34:
			goto tr138
		case 44:
			goto tr1760
		case 61:
			goto st6
		case 92:
			goto st69
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1758
		}
		goto st64
	st334:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof334
		}
	st_case_334:
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr200
		case 13:
			goto st6
		case 32:
			goto tr199
		case 34:
			goto tr138
		case 44:
			goto tr201
		case 61:
			goto st6
		case 92:
			goto st69
		case 108:
			goto st335
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr199
		}
		goto st64
	st335:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof335
		}
	st_case_335:
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr200
		case 13:
			goto st6
		case 32:
			goto tr199
		case 34:
			goto tr138
		case 44:
			goto tr201
		case 61:
			goto st6
		case 92:
			goto st69
		case 115:
			goto st336
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr199
		}
		goto st64
	st336:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof336
		}
	st_case_336:
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr200
		case 13:
			goto st6
		case 32:
			goto tr199
		case 34:
			goto tr138
		case 44:
			goto tr201
		case 61:
			goto st6
		case 92:
			goto st69
		case 101:
			goto st1214
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr199
		}
		goto st64
tr573:
//line machine.rl:17

	m.pb = m.p

	goto st1215
	st1215:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1215
		}
	st_case_1215:
//line machine.go:41371
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1166
		case 11:
			goto tr1759
		case 13:
			goto tr1166
		case 32:
			goto tr1758
		case 34:
			goto tr138
		case 44:
			goto tr1760
		case 61:
			goto st6
		case 82:
			goto st337
		case 92:
			goto st69
		case 114:
			goto st338
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1758
		}
		goto st64
	st337:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof337
		}
	st_case_337:
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr200
		case 13:
			goto st6
		case 32:
			goto tr199
		case 34:
			goto tr138
		case 44:
			goto tr201
		case 61:
			goto st6
		case 85:
			goto st333
		case 92:
			goto st69
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr199
		}
		goto st64
	st338:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof338
		}
	st_case_338:
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr200
		case 13:
			goto st6
		case 32:
			goto tr199
		case 34:
			goto tr138
		case 44:
			goto tr201
		case 61:
			goto st6
		case 92:
			goto st69
		case 117:
			goto st336
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr199
		}
		goto st64
tr574:
//line machine.rl:17

	m.pb = m.p

	goto st1216
	st1216:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1216
		}
	st_case_1216:
//line machine.go:41467
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1166
		case 11:
			goto tr1759
		case 13:
			goto tr1166
		case 32:
			goto tr1758
		case 34:
			goto tr138
		case 44:
			goto tr1760
		case 61:
			goto st6
		case 92:
			goto st69
		case 97:
			goto st334
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1758
		}
		goto st64
tr575:
//line machine.rl:17

	m.pb = m.p

	goto st1217
	st1217:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1217
		}
	st_case_1217:
//line machine.go:41503
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1166
		case 11:
			goto tr1759
		case 13:
			goto tr1166
		case 32:
			goto tr1758
		case 34:
			goto tr138
		case 44:
			goto tr1760
		case 61:
			goto st6
		case 92:
			goto st69
		case 114:
			goto st338
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1758
		}
		goto st64
tr560:
//line machine.rl:17

	m.pb = m.p

	goto st339
	st339:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof339
		}
	st_case_339:
//line machine.go:41539
		switch ( m.data)[( m.p)] {
		case 34:
			goto st295
		case 92:
			goto st295
		}
		switch {
		case ( m.data)[( m.p)] > 10:
			if 12 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto tr495
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr495
		}
		goto st266
tr549:
//line machine.rl:17

	m.pb = m.p

	goto st340
	st340:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof340
		}
	st_case_340:
//line machine.go:41566
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr406
		case 13:
			goto st6
		case 32:
			goto tr405
		case 34:
			goto tr407
		case 44:
			goto tr408
		case 48:
			goto st1220
		case 92:
			goto st343
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 49 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1245
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr405
		}
		goto st213
tr407:
//line machine.rl:86

	m.handler.AddString(key, m.text())

	goto st1218
tr642:
//line machine.rl:86

	m.handler.AddString(key, m.text())

//line machine.rl:17

	m.pb = m.p

	goto st1218
	st1218:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1218
		}
	st_case_1218:
//line machine.go:41615
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr850
		case 11:
			goto tr1765
		case 13:
			goto tr850
		case 32:
			goto tr1441
		case 44:
			goto tr1443
		case 92:
			goto st76
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1441
		}
		goto st7
tr1765:
//line machine.rl:58

	m.handler.SetMeasurement(m.text())

	goto st1219
	st1219:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1219
		}
	st_case_1219:
//line machine.go:41645
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr850
		case 11:
			goto tr1446
		case 13:
			goto tr850
		case 32:
			goto tr1441
		case 44:
			goto tr4
		case 45:
			goto tr1447
		case 61:
			goto st7
		case 92:
			goto tr33
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto tr1448
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1441
		}
		goto tr31
tr551:
//line machine.rl:17

	m.pb = m.p

	goto st1220
	st1220:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1220
		}
	st_case_1220:
//line machine.go:41684
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1415
		case 11:
			goto tr1767
		case 13:
			goto tr1415
		case 32:
			goto tr1766
		case 34:
			goto tr407
		case 44:
			goto tr1768
		case 46:
			goto st1243
		case 92:
			goto st343
		case 105:
			goto st1244
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1766
		}
		goto st213
tr1767:
//line machine.rl:58

	m.handler.SetMeasurement(m.text())

//line machine.rl:78

	m.handler.AddFloat(key, m.text())

	goto st1221
tr1798:
//line machine.rl:58

	m.handler.SetMeasurement(m.text())

//line machine.rl:74

	m.handler.AddInt(key, m.text())

	goto st1221
tr1801:
//line machine.rl:58

	m.handler.SetMeasurement(m.text())

//line machine.rl:82

	m.handler.AddBool(key, m.text())

	goto st1221
	st1221:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1221
		}
	st_case_1221:
//line machine.go:41744
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1366
		case 11:
			goto tr1772
		case 13:
			goto tr1366
		case 32:
			goto tr1771
		case 34:
			goto tr486
		case 44:
			goto tr408
		case 45:
			goto tr1773
		case 61:
			goto st213
		case 92:
			goto tr487
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto tr1774
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1771
		}
		goto tr484
tr1772:
//line machine.rl:58

	m.handler.SetMeasurement(m.text())

//line machine.rl:17

	m.pb = m.p

	goto st1222
	st1222:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1222
		}
	st_case_1222:
//line machine.go:41789
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1366
		case 11:
			goto tr1772
		case 13:
			goto tr1366
		case 32:
			goto tr1771
		case 34:
			goto tr486
		case 44:
			goto tr408
		case 45:
			goto tr1773
		case 61:
			goto tr491
		case 92:
			goto tr487
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto tr1774
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1771
		}
		goto tr484
tr1773:
//line machine.rl:17

	m.pb = m.p

	goto st341
	st341:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof341
		}
	st_case_341:
//line machine.go:41830
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr489
		case 13:
			goto st6
		case 32:
			goto tr405
		case 34:
			goto tr490
		case 44:
			goto tr408
		case 61:
			goto tr491
		case 92:
			goto st342
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1223
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr405
		}
		goto st262
tr1774:
//line machine.rl:17

	m.pb = m.p

	goto st1223
	st1223:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1223
		}
	st_case_1223:
//line machine.go:41869
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 11:
			goto tr1776
		case 13:
			goto tr1595
		case 32:
			goto tr1775
		case 34:
			goto tr490
		case 44:
			goto tr408
		case 61:
			goto tr491
		case 92:
			goto st342
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1225
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1775
		}
		goto st262
tr1779:
//line machine.rl:58

	m.handler.SetMeasurement(m.text())

//line machine.rl:17

	m.pb = m.p

	goto st1224
tr1776:
//line machine.rl:58

	m.handler.SetMeasurement(m.text())

//line machine.rl:90

	m.handler.SetTimestamp(m.text())

	goto st1224
	st1224:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1224
		}
	st_case_1224:
//line machine.go:41922
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1366
		case 11:
			goto tr1779
		case 13:
			goto tr1366
		case 32:
			goto tr1778
		case 34:
			goto tr486
		case 44:
			goto tr408
		case 61:
			goto tr491
		case 92:
			goto tr487
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1778
		}
		goto tr484
tr487:
//line machine.rl:17

	m.pb = m.p

	goto st342
	st342:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof342
		}
	st_case_342:
//line machine.go:41956
		switch ( m.data)[( m.p)] {
		case 34:
			goto st262
		case 92:
			goto st215
		}
		switch {
		case ( m.data)[( m.p)] > 10:
			if 12 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto tr2
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr2
		}
		goto st9
	st1225:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1225
		}
	st_case_1225:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 11:
			goto tr1776
		case 13:
			goto tr1595
		case 32:
			goto tr1775
		case 34:
			goto tr490
		case 44:
			goto tr408
		case 61:
			goto tr491
		case 92:
			goto st342
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1226
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1775
		}
		goto st262
	st1226:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1226
		}
	st_case_1226:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 11:
			goto tr1776
		case 13:
			goto tr1595
		case 32:
			goto tr1775
		case 34:
			goto tr490
		case 44:
			goto tr408
		case 61:
			goto tr491
		case 92:
			goto st342
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1227
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1775
		}
		goto st262
	st1227:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1227
		}
	st_case_1227:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 11:
			goto tr1776
		case 13:
			goto tr1595
		case 32:
			goto tr1775
		case 34:
			goto tr490
		case 44:
			goto tr408
		case 61:
			goto tr491
		case 92:
			goto st342
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1228
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1775
		}
		goto st262
	st1228:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1228
		}
	st_case_1228:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 11:
			goto tr1776
		case 13:
			goto tr1595
		case 32:
			goto tr1775
		case 34:
			goto tr490
		case 44:
			goto tr408
		case 61:
			goto tr491
		case 92:
			goto st342
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1229
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1775
		}
		goto st262
	st1229:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1229
		}
	st_case_1229:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 11:
			goto tr1776
		case 13:
			goto tr1595
		case 32:
			goto tr1775
		case 34:
			goto tr490
		case 44:
			goto tr408
		case 61:
			goto tr491
		case 92:
			goto st342
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1230
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1775
		}
		goto st262
	st1230:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1230
		}
	st_case_1230:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 11:
			goto tr1776
		case 13:
			goto tr1595
		case 32:
			goto tr1775
		case 34:
			goto tr490
		case 44:
			goto tr408
		case 61:
			goto tr491
		case 92:
			goto st342
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1231
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1775
		}
		goto st262
	st1231:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1231
		}
	st_case_1231:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 11:
			goto tr1776
		case 13:
			goto tr1595
		case 32:
			goto tr1775
		case 34:
			goto tr490
		case 44:
			goto tr408
		case 61:
			goto tr491
		case 92:
			goto st342
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1232
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1775
		}
		goto st262
	st1232:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1232
		}
	st_case_1232:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 11:
			goto tr1776
		case 13:
			goto tr1595
		case 32:
			goto tr1775
		case 34:
			goto tr490
		case 44:
			goto tr408
		case 61:
			goto tr491
		case 92:
			goto st342
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1233
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1775
		}
		goto st262
	st1233:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1233
		}
	st_case_1233:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 11:
			goto tr1776
		case 13:
			goto tr1595
		case 32:
			goto tr1775
		case 34:
			goto tr490
		case 44:
			goto tr408
		case 61:
			goto tr491
		case 92:
			goto st342
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1234
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1775
		}
		goto st262
	st1234:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1234
		}
	st_case_1234:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 11:
			goto tr1776
		case 13:
			goto tr1595
		case 32:
			goto tr1775
		case 34:
			goto tr490
		case 44:
			goto tr408
		case 61:
			goto tr491
		case 92:
			goto st342
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1235
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1775
		}
		goto st262
	st1235:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1235
		}
	st_case_1235:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 11:
			goto tr1776
		case 13:
			goto tr1595
		case 32:
			goto tr1775
		case 34:
			goto tr490
		case 44:
			goto tr408
		case 61:
			goto tr491
		case 92:
			goto st342
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1236
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1775
		}
		goto st262
	st1236:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1236
		}
	st_case_1236:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 11:
			goto tr1776
		case 13:
			goto tr1595
		case 32:
			goto tr1775
		case 34:
			goto tr490
		case 44:
			goto tr408
		case 61:
			goto tr491
		case 92:
			goto st342
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1237
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1775
		}
		goto st262
	st1237:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1237
		}
	st_case_1237:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 11:
			goto tr1776
		case 13:
			goto tr1595
		case 32:
			goto tr1775
		case 34:
			goto tr490
		case 44:
			goto tr408
		case 61:
			goto tr491
		case 92:
			goto st342
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1238
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1775
		}
		goto st262
	st1238:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1238
		}
	st_case_1238:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 11:
			goto tr1776
		case 13:
			goto tr1595
		case 32:
			goto tr1775
		case 34:
			goto tr490
		case 44:
			goto tr408
		case 61:
			goto tr491
		case 92:
			goto st342
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1239
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1775
		}
		goto st262
	st1239:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1239
		}
	st_case_1239:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 11:
			goto tr1776
		case 13:
			goto tr1595
		case 32:
			goto tr1775
		case 34:
			goto tr490
		case 44:
			goto tr408
		case 61:
			goto tr491
		case 92:
			goto st342
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1240
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1775
		}
		goto st262
	st1240:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1240
		}
	st_case_1240:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 11:
			goto tr1776
		case 13:
			goto tr1595
		case 32:
			goto tr1775
		case 34:
			goto tr490
		case 44:
			goto tr408
		case 61:
			goto tr491
		case 92:
			goto st342
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1241
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1775
		}
		goto st262
	st1241:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1241
		}
	st_case_1241:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 11:
			goto tr1776
		case 13:
			goto tr1595
		case 32:
			goto tr1775
		case 34:
			goto tr490
		case 44:
			goto tr408
		case 61:
			goto tr491
		case 92:
			goto st342
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1242
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1775
		}
		goto st262
	st1242:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1242
		}
	st_case_1242:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 11:
			goto tr1776
		case 13:
			goto tr1595
		case 32:
			goto tr1775
		case 34:
			goto tr490
		case 44:
			goto tr408
		case 61:
			goto tr491
		case 92:
			goto st342
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1775
		}
		goto st262
tr550:
//line machine.rl:17

	m.pb = m.p

	goto st1243
	st1243:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1243
		}
	st_case_1243:
//line machine.go:42554
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1415
		case 11:
			goto tr1767
		case 13:
			goto tr1415
		case 32:
			goto tr1766
		case 34:
			goto tr407
		case 44:
			goto tr1768
		case 92:
			goto st343
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1243
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1766
		}
		goto st213
tr643:
//line machine.rl:17

	m.pb = m.p

	goto st343
	st343:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof343
		}
	st_case_343:
//line machine.go:42591
		switch ( m.data)[( m.p)] {
		case 34:
			goto st213
		case 92:
			goto st6
		}
		switch {
		case ( m.data)[( m.p)] > 10:
			if 12 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto tr2
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr2
		}
		goto st7
	st1244:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1244
		}
	st_case_1244:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1420
		case 11:
			goto tr1798
		case 13:
			goto tr1420
		case 32:
			goto tr1797
		case 34:
			goto tr407
		case 44:
			goto tr1799
		case 92:
			goto st343
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1797
		}
		goto st213
tr552:
//line machine.rl:17

	m.pb = m.p

	goto st1245
	st1245:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1245
		}
	st_case_1245:
//line machine.go:42643
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1415
		case 11:
			goto tr1767
		case 13:
			goto tr1415
		case 32:
			goto tr1766
		case 34:
			goto tr407
		case 44:
			goto tr1768
		case 46:
			goto st1243
		case 92:
			goto st343
		case 105:
			goto st1244
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1245
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1766
		}
		goto st213
tr553:
//line machine.rl:17

	m.pb = m.p

	goto st1246
	st1246:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1246
		}
	st_case_1246:
//line machine.go:42684
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1423
		case 11:
			goto tr1801
		case 13:
			goto tr1423
		case 32:
			goto tr1800
		case 34:
			goto tr407
		case 44:
			goto tr1802
		case 65:
			goto st344
		case 92:
			goto st343
		case 97:
			goto st347
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1800
		}
		goto st213
	st344:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof344
		}
	st_case_344:
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr406
		case 13:
			goto st6
		case 32:
			goto tr405
		case 34:
			goto tr407
		case 44:
			goto tr408
		case 76:
			goto st345
		case 92:
			goto st343
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr405
		}
		goto st213
	st345:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof345
		}
	st_case_345:
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr406
		case 13:
			goto st6
		case 32:
			goto tr405
		case 34:
			goto tr407
		case 44:
			goto tr408
		case 83:
			goto st346
		case 92:
			goto st343
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr405
		}
		goto st213
	st346:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof346
		}
	st_case_346:
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr406
		case 13:
			goto st6
		case 32:
			goto tr405
		case 34:
			goto tr407
		case 44:
			goto tr408
		case 69:
			goto st1247
		case 92:
			goto st343
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr405
		}
		goto st213
	st1247:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1247
		}
	st_case_1247:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1423
		case 11:
			goto tr1801
		case 13:
			goto tr1423
		case 32:
			goto tr1800
		case 34:
			goto tr407
		case 44:
			goto tr1802
		case 92:
			goto st343
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1800
		}
		goto st213
	st347:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof347
		}
	st_case_347:
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr406
		case 13:
			goto st6
		case 32:
			goto tr405
		case 34:
			goto tr407
		case 44:
			goto tr408
		case 92:
			goto st343
		case 108:
			goto st348
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr405
		}
		goto st213
	st348:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof348
		}
	st_case_348:
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr406
		case 13:
			goto st6
		case 32:
			goto tr405
		case 34:
			goto tr407
		case 44:
			goto tr408
		case 92:
			goto st343
		case 115:
			goto st349
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr405
		}
		goto st213
	st349:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof349
		}
	st_case_349:
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr406
		case 13:
			goto st6
		case 32:
			goto tr405
		case 34:
			goto tr407
		case 44:
			goto tr408
		case 92:
			goto st343
		case 101:
			goto st1247
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr405
		}
		goto st213
tr554:
//line machine.rl:17

	m.pb = m.p

	goto st1248
	st1248:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1248
		}
	st_case_1248:
//line machine.go:42907
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1423
		case 11:
			goto tr1801
		case 13:
			goto tr1423
		case 32:
			goto tr1800
		case 34:
			goto tr407
		case 44:
			goto tr1802
		case 82:
			goto st350
		case 92:
			goto st343
		case 114:
			goto st351
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1800
		}
		goto st213
	st350:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof350
		}
	st_case_350:
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr406
		case 13:
			goto st6
		case 32:
			goto tr405
		case 34:
			goto tr407
		case 44:
			goto tr408
		case 85:
			goto st346
		case 92:
			goto st343
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr405
		}
		goto st213
	st351:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof351
		}
	st_case_351:
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr406
		case 13:
			goto st6
		case 32:
			goto tr405
		case 34:
			goto tr407
		case 44:
			goto tr408
		case 92:
			goto st343
		case 117:
			goto st349
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr405
		}
		goto st213
tr555:
//line machine.rl:17

	m.pb = m.p

	goto st1249
	st1249:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1249
		}
	st_case_1249:
//line machine.go:42997
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1423
		case 11:
			goto tr1801
		case 13:
			goto tr1423
		case 32:
			goto tr1800
		case 34:
			goto tr407
		case 44:
			goto tr1802
		case 92:
			goto st343
		case 97:
			goto st347
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1800
		}
		goto st213
tr556:
//line machine.rl:17

	m.pb = m.p

	goto st1250
	st1250:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1250
		}
	st_case_1250:
//line machine.go:43031
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1423
		case 11:
			goto tr1801
		case 13:
			goto tr1423
		case 32:
			goto tr1800
		case 34:
			goto tr407
		case 44:
			goto tr1802
		case 92:
			goto st343
		case 114:
			goto st351
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1800
		}
		goto st213
	st352:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof352
		}
	st_case_352:
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr641
		case 13:
			goto st6
		case 32:
			goto st352
		case 34:
			goto tr642
		case 44:
			goto st6
		case 92:
			goto tr643
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto st352
		}
		goto tr639
tr641:
//line machine.rl:17

	m.pb = m.p

	goto st353
	st353:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof353
		}
	st_case_353:
//line machine.go:43090
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr645
		case 13:
			goto st6
		case 32:
			goto tr644
		case 34:
			goto tr642
		case 44:
			goto tr408
		case 92:
			goto tr643
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr644
		}
		goto tr639
tr644:
//line machine.rl:58

	m.handler.SetMeasurement(m.text())

	goto st354
	st354:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof354
		}
	st_case_354:
//line machine.go:43122
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr647
		case 13:
			goto st6
		case 32:
			goto st354
		case 34:
			goto tr486
		case 44:
			goto st6
		case 61:
			goto tr639
		case 92:
			goto tr487
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto st354
		}
		goto tr484
tr647:
//line machine.rl:17

	m.pb = m.p

	goto st355
tr648:
//line machine.rl:17

	m.pb = m.p

//line machine.rl:58

	m.handler.SetMeasurement(m.text())

	goto st355
	st355:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof355
		}
	st_case_355:
//line machine.go:43166
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr648
		case 13:
			goto st6
		case 32:
			goto tr644
		case 34:
			goto tr486
		case 44:
			goto tr408
		case 61:
			goto tr649
		case 92:
			goto tr487
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr644
		}
		goto tr484
tr645:
//line machine.rl:17

	m.pb = m.p

//line machine.rl:58

	m.handler.SetMeasurement(m.text())

	goto st356
	st356:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof356
		}
	st_case_356:
//line machine.go:43204
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr648
		case 13:
			goto st6
		case 32:
			goto tr644
		case 34:
			goto tr486
		case 44:
			goto tr408
		case 61:
			goto tr639
		case 92:
			goto tr487
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr644
		}
		goto tr484
tr1367:
//line machine.rl:17

	m.pb = m.p

	goto st357
	st357:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof357
		}
	st_case_357:
//line machine.go:43238
		switch ( m.data)[( m.p)] {
		case 34:
			goto tr26
		case 92:
			goto st59
		}
		if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
			goto st1251
		}
		goto st6
tr1368:
//line machine.rl:17

	m.pb = m.p

	goto st1251
	st1251:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1251
		}
	st_case_1251:
//line machine.go:43260
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 13:
			goto tr1595
		case 32:
			goto tr1594
		case 34:
			goto tr26
		case 92:
			goto st59
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1252
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1594
		}
		goto st6
	st1252:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1252
		}
	st_case_1252:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 13:
			goto tr1595
		case 32:
			goto tr1594
		case 34:
			goto tr26
		case 92:
			goto st59
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1253
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1594
		}
		goto st6
	st1253:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1253
		}
	st_case_1253:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 13:
			goto tr1595
		case 32:
			goto tr1594
		case 34:
			goto tr26
		case 92:
			goto st59
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1254
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1594
		}
		goto st6
	st1254:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1254
		}
	st_case_1254:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 13:
			goto tr1595
		case 32:
			goto tr1594
		case 34:
			goto tr26
		case 92:
			goto st59
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1255
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1594
		}
		goto st6
	st1255:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1255
		}
	st_case_1255:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 13:
			goto tr1595
		case 32:
			goto tr1594
		case 34:
			goto tr26
		case 92:
			goto st59
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1256
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1594
		}
		goto st6
	st1256:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1256
		}
	st_case_1256:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 13:
			goto tr1595
		case 32:
			goto tr1594
		case 34:
			goto tr26
		case 92:
			goto st59
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1257
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1594
		}
		goto st6
	st1257:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1257
		}
	st_case_1257:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 13:
			goto tr1595
		case 32:
			goto tr1594
		case 34:
			goto tr26
		case 92:
			goto st59
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1258
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1594
		}
		goto st6
	st1258:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1258
		}
	st_case_1258:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 13:
			goto tr1595
		case 32:
			goto tr1594
		case 34:
			goto tr26
		case 92:
			goto st59
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1259
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1594
		}
		goto st6
	st1259:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1259
		}
	st_case_1259:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 13:
			goto tr1595
		case 32:
			goto tr1594
		case 34:
			goto tr26
		case 92:
			goto st59
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1260
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1594
		}
		goto st6
	st1260:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1260
		}
	st_case_1260:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 13:
			goto tr1595
		case 32:
			goto tr1594
		case 34:
			goto tr26
		case 92:
			goto st59
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1261
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1594
		}
		goto st6
	st1261:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1261
		}
	st_case_1261:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 13:
			goto tr1595
		case 32:
			goto tr1594
		case 34:
			goto tr26
		case 92:
			goto st59
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1262
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1594
		}
		goto st6
	st1262:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1262
		}
	st_case_1262:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 13:
			goto tr1595
		case 32:
			goto tr1594
		case 34:
			goto tr26
		case 92:
			goto st59
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1263
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1594
		}
		goto st6
	st1263:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1263
		}
	st_case_1263:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 13:
			goto tr1595
		case 32:
			goto tr1594
		case 34:
			goto tr26
		case 92:
			goto st59
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1264
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1594
		}
		goto st6
	st1264:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1264
		}
	st_case_1264:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 13:
			goto tr1595
		case 32:
			goto tr1594
		case 34:
			goto tr26
		case 92:
			goto st59
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1265
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1594
		}
		goto st6
	st1265:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1265
		}
	st_case_1265:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 13:
			goto tr1595
		case 32:
			goto tr1594
		case 34:
			goto tr26
		case 92:
			goto st59
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1266
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1594
		}
		goto st6
	st1266:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1266
		}
	st_case_1266:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 13:
			goto tr1595
		case 32:
			goto tr1594
		case 34:
			goto tr26
		case 92:
			goto st59
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1267
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1594
		}
		goto st6
	st1267:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1267
		}
	st_case_1267:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 13:
			goto tr1595
		case 32:
			goto tr1594
		case 34:
			goto tr26
		case 92:
			goto st59
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1268
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1594
		}
		goto st6
	st1268:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1268
		}
	st_case_1268:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 13:
			goto tr1595
		case 32:
			goto tr1594
		case 34:
			goto tr26
		case 92:
			goto st59
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1269
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1594
		}
		goto st6
	st1269:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1269
		}
	st_case_1269:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 13:
			goto tr1595
		case 32:
			goto tr1594
		case 34:
			goto tr26
		case 92:
			goto st59
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1594
		}
		goto st6
tr1364:
//line machine.rl:17

	m.pb = m.p

	goto st358
tr1835:
//line machine.rl:78

	m.handler.AddFloat(key, m.text())

	goto st358
tr1838:
//line machine.rl:74

	m.handler.AddInt(key, m.text())

	goto st358
tr1839:
//line machine.rl:82

	m.handler.AddBool(key, m.text())

	goto st358
	st358:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof358
		}
	st_case_358:
//line machine.go:43774
		switch ( m.data)[( m.p)] {
		case 32:
			goto st6
		case 34:
			goto tr322
		case 44:
			goto st6
		case 61:
			goto st6
		case 92:
			goto tr652
		}
		switch {
		case ( m.data)[( m.p)] > 10:
			if 12 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto st6
			}
		case ( m.data)[( m.p)] >= 9:
			goto st6
		}
		goto tr651
tr651:
//line machine.rl:17

	m.pb = m.p

	goto st359
	st359:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof359
		}
	st_case_359:
//line machine.go:43807
		switch ( m.data)[( m.p)] {
		case 32:
			goto st6
		case 34:
			goto tr85
		case 44:
			goto st6
		case 61:
			goto tr654
		case 92:
			goto st383
		}
		switch {
		case ( m.data)[( m.p)] > 10:
			if 12 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto st6
			}
		case ( m.data)[( m.p)] >= 9:
			goto st6
		}
		goto st359
tr654:
//line machine.rl:70

	key = m.text()

	goto st360
	st360:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof360
		}
	st_case_360:
//line machine.go:43840
		switch ( m.data)[( m.p)] {
		case 34:
			goto tr656
		case 45:
			goto tr396
		case 46:
			goto tr397
		case 48:
			goto tr398
		case 70:
			goto tr400
		case 84:
			goto tr401
		case 92:
			goto st59
		case 102:
			goto tr402
		case 116:
			goto tr403
		}
		if 49 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
			goto tr399
		}
		goto st6
tr656:
//line machine.rl:86

	m.handler.AddString(key, m.text())

	goto st1270
	st1270:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1270
		}
	st_case_1270:
//line machine.go:43876
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1363
		case 13:
			goto tr1363
		case 32:
			goto tr1362
		case 34:
			goto tr23
		case 44:
			goto tr1825
		case 92:
			goto tr24
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1362
		}
		goto tr22
tr1825:
//line machine.rl:17

	m.pb = m.p

	goto st361
tr1826:
//line machine.rl:78

	m.handler.AddFloat(key, m.text())

	goto st361
tr1829:
//line machine.rl:74

	m.handler.AddInt(key, m.text())

	goto st361
tr1830:
//line machine.rl:82

	m.handler.AddBool(key, m.text())

	goto st361
	st361:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof361
		}
	st_case_361:
//line machine.go:43924
		switch ( m.data)[( m.p)] {
		case 32:
			goto st6
		case 34:
			goto tr82
		case 44:
			goto st6
		case 61:
			goto st6
		case 92:
			goto tr658
		}
		switch {
		case ( m.data)[( m.p)] > 10:
			if 12 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto st6
			}
		case ( m.data)[( m.p)] >= 9:
			goto st6
		}
		goto tr657
tr657:
//line machine.rl:17

	m.pb = m.p

	goto st362
	st362:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof362
		}
	st_case_362:
//line machine.go:43957
		switch ( m.data)[( m.p)] {
		case 32:
			goto st6
		case 34:
			goto tr85
		case 44:
			goto st6
		case 61:
			goto tr660
		case 92:
			goto st373
		}
		switch {
		case ( m.data)[( m.p)] > 10:
			if 12 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto st6
			}
		case ( m.data)[( m.p)] >= 9:
			goto st6
		}
		goto st362
tr660:
//line machine.rl:70

	key = m.text()

	goto st363
	st363:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof363
		}
	st_case_363:
//line machine.go:43990
		switch ( m.data)[( m.p)] {
		case 34:
			goto tr656
		case 45:
			goto tr662
		case 46:
			goto tr663
		case 48:
			goto tr664
		case 70:
			goto tr666
		case 84:
			goto tr667
		case 92:
			goto st59
		case 102:
			goto tr668
		case 116:
			goto tr669
		}
		if 49 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
			goto tr665
		}
		goto st6
tr662:
//line machine.rl:17

	m.pb = m.p

	goto st364
	st364:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof364
		}
	st_case_364:
//line machine.go:44026
		switch ( m.data)[( m.p)] {
		case 34:
			goto tr26
		case 48:
			goto st1271
		case 92:
			goto st59
		}
		if 49 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
			goto st1274
		}
		goto st6
tr664:
//line machine.rl:17

	m.pb = m.p

	goto st1271
	st1271:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1271
		}
	st_case_1271:
//line machine.go:44050
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1415
		case 13:
			goto tr1415
		case 32:
			goto tr1414
		case 34:
			goto tr26
		case 44:
			goto tr1826
		case 46:
			goto st1272
		case 92:
			goto st59
		case 105:
			goto st1273
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1414
		}
		goto st6
tr663:
//line machine.rl:17

	m.pb = m.p

	goto st1272
	st1272:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1272
		}
	st_case_1272:
//line machine.go:44084
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1415
		case 13:
			goto tr1415
		case 32:
			goto tr1414
		case 34:
			goto tr26
		case 44:
			goto tr1826
		case 92:
			goto st59
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1272
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1414
		}
		goto st6
	st1273:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1273
		}
	st_case_1273:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1420
		case 13:
			goto tr1420
		case 32:
			goto tr1419
		case 34:
			goto tr26
		case 44:
			goto tr1829
		case 92:
			goto st59
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1419
		}
		goto st6
tr665:
//line machine.rl:17

	m.pb = m.p

	goto st1274
	st1274:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1274
		}
	st_case_1274:
//line machine.go:44142
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1415
		case 13:
			goto tr1415
		case 32:
			goto tr1414
		case 34:
			goto tr26
		case 44:
			goto tr1826
		case 46:
			goto st1272
		case 92:
			goto st59
		case 105:
			goto st1273
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1274
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1414
		}
		goto st6
tr666:
//line machine.rl:17

	m.pb = m.p

	goto st1275
	st1275:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1275
		}
	st_case_1275:
//line machine.go:44181
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1423
		case 13:
			goto tr1423
		case 32:
			goto tr1422
		case 34:
			goto tr26
		case 44:
			goto tr1830
		case 65:
			goto st365
		case 92:
			goto st59
		case 97:
			goto st368
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1422
		}
		goto st6
	st365:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof365
		}
	st_case_365:
		switch ( m.data)[( m.p)] {
		case 34:
			goto tr26
		case 76:
			goto st366
		case 92:
			goto st59
		}
		goto st6
	st366:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof366
		}
	st_case_366:
		switch ( m.data)[( m.p)] {
		case 34:
			goto tr26
		case 83:
			goto st367
		case 92:
			goto st59
		}
		goto st6
	st367:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof367
		}
	st_case_367:
		switch ( m.data)[( m.p)] {
		case 34:
			goto tr26
		case 69:
			goto st1276
		case 92:
			goto st59
		}
		goto st6
	st1276:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1276
		}
	st_case_1276:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1423
		case 13:
			goto tr1423
		case 32:
			goto tr1422
		case 34:
			goto tr26
		case 44:
			goto tr1830
		case 92:
			goto st59
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1422
		}
		goto st6
	st368:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof368
		}
	st_case_368:
		switch ( m.data)[( m.p)] {
		case 34:
			goto tr26
		case 92:
			goto st59
		case 108:
			goto st369
		}
		goto st6
	st369:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof369
		}
	st_case_369:
		switch ( m.data)[( m.p)] {
		case 34:
			goto tr26
		case 92:
			goto st59
		case 115:
			goto st370
		}
		goto st6
	st370:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof370
		}
	st_case_370:
		switch ( m.data)[( m.p)] {
		case 34:
			goto tr26
		case 92:
			goto st59
		case 101:
			goto st1276
		}
		goto st6
tr667:
//line machine.rl:17

	m.pb = m.p

	goto st1277
	st1277:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1277
		}
	st_case_1277:
//line machine.go:44322
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1423
		case 13:
			goto tr1423
		case 32:
			goto tr1422
		case 34:
			goto tr26
		case 44:
			goto tr1830
		case 82:
			goto st371
		case 92:
			goto st59
		case 114:
			goto st372
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1422
		}
		goto st6
	st371:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof371
		}
	st_case_371:
		switch ( m.data)[( m.p)] {
		case 34:
			goto tr26
		case 85:
			goto st367
		case 92:
			goto st59
		}
		goto st6
	st372:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof372
		}
	st_case_372:
		switch ( m.data)[( m.p)] {
		case 34:
			goto tr26
		case 92:
			goto st59
		case 117:
			goto st370
		}
		goto st6
tr668:
//line machine.rl:17

	m.pb = m.p

	goto st1278
	st1278:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1278
		}
	st_case_1278:
//line machine.go:44384
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1423
		case 13:
			goto tr1423
		case 32:
			goto tr1422
		case 34:
			goto tr26
		case 44:
			goto tr1830
		case 92:
			goto st59
		case 97:
			goto st368
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1422
		}
		goto st6
tr669:
//line machine.rl:17

	m.pb = m.p

	goto st1279
	st1279:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1279
		}
	st_case_1279:
//line machine.go:44416
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1423
		case 13:
			goto tr1423
		case 32:
			goto tr1422
		case 34:
			goto tr26
		case 44:
			goto tr1830
		case 92:
			goto st59
		case 114:
			goto st372
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1422
		}
		goto st6
tr658:
//line machine.rl:17

	m.pb = m.p

	goto st373
	st373:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof373
		}
	st_case_373:
//line machine.go:44448
		switch ( m.data)[( m.p)] {
		case 34:
			goto st362
		case 92:
			goto st362
		}
		switch {
		case ( m.data)[( m.p)] > 10:
			if 12 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto tr2
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr2
		}
		goto st3
tr396:
//line machine.rl:17

	m.pb = m.p

	goto st374
	st374:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof374
		}
	st_case_374:
//line machine.go:44475
		switch ( m.data)[( m.p)] {
		case 34:
			goto tr26
		case 48:
			goto st1280
		case 92:
			goto st59
		}
		if 49 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
			goto st1283
		}
		goto st6
tr398:
//line machine.rl:17

	m.pb = m.p

	goto st1280
	st1280:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1280
		}
	st_case_1280:
//line machine.go:44499
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1415
		case 13:
			goto tr1415
		case 32:
			goto tr1414
		case 34:
			goto tr26
		case 44:
			goto tr1835
		case 46:
			goto st1281
		case 92:
			goto st59
		case 105:
			goto st1282
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1414
		}
		goto st6
tr397:
//line machine.rl:17

	m.pb = m.p

	goto st1281
	st1281:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1281
		}
	st_case_1281:
//line machine.go:44533
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1415
		case 13:
			goto tr1415
		case 32:
			goto tr1414
		case 34:
			goto tr26
		case 44:
			goto tr1835
		case 92:
			goto st59
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1281
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1414
		}
		goto st6
	st1282:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1282
		}
	st_case_1282:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1420
		case 13:
			goto tr1420
		case 32:
			goto tr1419
		case 34:
			goto tr26
		case 44:
			goto tr1838
		case 92:
			goto st59
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1419
		}
		goto st6
tr399:
//line machine.rl:17

	m.pb = m.p

	goto st1283
	st1283:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1283
		}
	st_case_1283:
//line machine.go:44591
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1415
		case 13:
			goto tr1415
		case 32:
			goto tr1414
		case 34:
			goto tr26
		case 44:
			goto tr1835
		case 46:
			goto st1281
		case 92:
			goto st59
		case 105:
			goto st1282
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1283
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1414
		}
		goto st6
tr400:
//line machine.rl:17

	m.pb = m.p

	goto st1284
	st1284:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1284
		}
	st_case_1284:
//line machine.go:44630
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1423
		case 13:
			goto tr1423
		case 32:
			goto tr1422
		case 34:
			goto tr26
		case 44:
			goto tr1839
		case 65:
			goto st375
		case 92:
			goto st59
		case 97:
			goto st378
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1422
		}
		goto st6
	st375:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof375
		}
	st_case_375:
		switch ( m.data)[( m.p)] {
		case 34:
			goto tr26
		case 76:
			goto st376
		case 92:
			goto st59
		}
		goto st6
	st376:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof376
		}
	st_case_376:
		switch ( m.data)[( m.p)] {
		case 34:
			goto tr26
		case 83:
			goto st377
		case 92:
			goto st59
		}
		goto st6
	st377:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof377
		}
	st_case_377:
		switch ( m.data)[( m.p)] {
		case 34:
			goto tr26
		case 69:
			goto st1285
		case 92:
			goto st59
		}
		goto st6
	st1285:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1285
		}
	st_case_1285:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1423
		case 13:
			goto tr1423
		case 32:
			goto tr1422
		case 34:
			goto tr26
		case 44:
			goto tr1839
		case 92:
			goto st59
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1422
		}
		goto st6
	st378:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof378
		}
	st_case_378:
		switch ( m.data)[( m.p)] {
		case 34:
			goto tr26
		case 92:
			goto st59
		case 108:
			goto st379
		}
		goto st6
	st379:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof379
		}
	st_case_379:
		switch ( m.data)[( m.p)] {
		case 34:
			goto tr26
		case 92:
			goto st59
		case 115:
			goto st380
		}
		goto st6
	st380:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof380
		}
	st_case_380:
		switch ( m.data)[( m.p)] {
		case 34:
			goto tr26
		case 92:
			goto st59
		case 101:
			goto st1285
		}
		goto st6
tr401:
//line machine.rl:17

	m.pb = m.p

	goto st1286
	st1286:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1286
		}
	st_case_1286:
//line machine.go:44771
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1423
		case 13:
			goto tr1423
		case 32:
			goto tr1422
		case 34:
			goto tr26
		case 44:
			goto tr1839
		case 82:
			goto st381
		case 92:
			goto st59
		case 114:
			goto st382
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1422
		}
		goto st6
	st381:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof381
		}
	st_case_381:
		switch ( m.data)[( m.p)] {
		case 34:
			goto tr26
		case 85:
			goto st377
		case 92:
			goto st59
		}
		goto st6
	st382:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof382
		}
	st_case_382:
		switch ( m.data)[( m.p)] {
		case 34:
			goto tr26
		case 92:
			goto st59
		case 117:
			goto st380
		}
		goto st6
tr402:
//line machine.rl:17

	m.pb = m.p

	goto st1287
	st1287:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1287
		}
	st_case_1287:
//line machine.go:44833
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1423
		case 13:
			goto tr1423
		case 32:
			goto tr1422
		case 34:
			goto tr26
		case 44:
			goto tr1839
		case 92:
			goto st59
		case 97:
			goto st378
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1422
		}
		goto st6
tr403:
//line machine.rl:17

	m.pb = m.p

	goto st1288
	st1288:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1288
		}
	st_case_1288:
//line machine.go:44865
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1423
		case 13:
			goto tr1423
		case 32:
			goto tr1422
		case 34:
			goto tr26
		case 44:
			goto tr1839
		case 92:
			goto st59
		case 114:
			goto st382
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1422
		}
		goto st6
tr652:
//line machine.rl:17

	m.pb = m.p

	goto st383
	st383:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof383
		}
	st_case_383:
//line machine.go:44897
		switch ( m.data)[( m.p)] {
		case 34:
			goto st359
		case 92:
			goto st359
		}
		switch {
		case ( m.data)[( m.p)] > 10:
			if 12 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto tr2
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr2
		}
		goto st3
tr391:
//line machine.rl:17

	m.pb = m.p

	goto st384
	st384:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof384
		}
	st_case_384:
//line machine.go:44924
		switch ( m.data)[( m.p)] {
		case 34:
			goto st211
		case 92:
			goto st211
		}
		switch {
		case ( m.data)[( m.p)] > 10:
			if 12 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto tr2
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr2
		}
		goto st3
tr390:
//line machine.rl:17

	m.pb = m.p

	goto st385
	st385:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof385
		}
	st_case_385:
//line machine.go:44951
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr390
		case 13:
			goto st6
		case 32:
			goto st210
		case 34:
			goto tr82
		case 44:
			goto st6
		case 61:
			goto tr393
		case 92:
			goto tr391
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto st210
		}
		goto tr388
tr384:
//line machine.rl:58

	m.handler.SetMeasurement(m.text())

	goto st386
tr378:
//line machine.rl:58

	m.handler.SetMeasurement(m.text())

//line machine.rl:17

	m.pb = m.p

	goto st386
	st386:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof386
		}
	st_case_386:
//line machine.go:44995
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr685
		case 13:
			goto st6
		case 32:
			goto tr383
		case 34:
			goto tr686
		case 44:
			goto tr386
		case 61:
			goto st209
		case 92:
			goto tr687
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr383
		}
		goto tr684
tr684:
//line machine.rl:17

	m.pb = m.p

	goto st387
	st387:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof387
		}
	st_case_387:
//line machine.go:45029
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr689
		case 13:
			goto st6
		case 32:
			goto tr383
		case 34:
			goto tr690
		case 44:
			goto tr386
		case 61:
			goto tr691
		case 92:
			goto st433
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr383
		}
		goto st387
tr689:
//line machine.rl:58

	m.handler.SetMeasurement(m.text())

	goto st388
tr685:
//line machine.rl:58

	m.handler.SetMeasurement(m.text())

//line machine.rl:17

	m.pb = m.p

	goto st388
	st388:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof388
		}
	st_case_388:
//line machine.go:45073
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr685
		case 13:
			goto st6
		case 32:
			goto tr383
		case 34:
			goto tr686
		case 44:
			goto tr386
		case 61:
			goto tr691
		case 92:
			goto tr687
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr383
		}
		goto tr684
tr686:
//line machine.rl:17

	m.pb = m.p

//line machine.rl:86

	m.handler.AddString(key, m.text())

	goto st1289
tr690:
//line machine.rl:86

	m.handler.AddString(key, m.text())

	goto st1289
	st1289:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1289
		}
	st_case_1289:
//line machine.go:45117
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr850
		case 11:
			goto tr1844
		case 13:
			goto tr850
		case 32:
			goto tr889
		case 44:
			goto tr1845
		case 61:
			goto tr365
		case 92:
			goto st390
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr889
		}
		goto st205
tr1844:
//line machine.rl:58

	m.handler.SetMeasurement(m.text())

	goto st1290
tr1846:
//line machine.rl:58

	m.handler.SetMeasurement(m.text())

//line machine.rl:17

	m.pb = m.p

	goto st1290
	st1290:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1290
		}
	st_case_1290:
//line machine.go:45159
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr850
		case 11:
			goto tr1846
		case 13:
			goto tr850
		case 32:
			goto tr889
		case 44:
			goto tr4
		case 45:
			goto tr1847
		case 61:
			goto tr365
		case 92:
			goto tr362
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto tr1848
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr889
		}
		goto tr360
tr1847:
//line machine.rl:17

	m.pb = m.p

	goto st389
	st389:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof389
		}
	st_case_389:
//line machine.go:45198
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr88
		case 11:
			goto tr364
		case 13:
			goto tr88
		case 32:
			goto tr1
		case 44:
			goto tr4
		case 61:
			goto tr365
		case 92:
			goto st390
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1291
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1
		}
		goto st205
tr1848:
//line machine.rl:17

	m.pb = m.p

	goto st1291
	st1291:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1291
		}
	st_case_1291:
//line machine.go:45235
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr1849
		case 13:
			goto tr859
		case 32:
			goto tr897
		case 44:
			goto tr4
		case 61:
			goto tr365
		case 92:
			goto st390
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1293
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr897
		}
		goto st205
tr1851:
//line machine.rl:58

	m.handler.SetMeasurement(m.text())

//line machine.rl:17

	m.pb = m.p

	goto st1292
tr1849:
//line machine.rl:58

	m.handler.SetMeasurement(m.text())

//line machine.rl:90

	m.handler.SetTimestamp(m.text())

	goto st1292
	st1292:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1292
		}
	st_case_1292:
//line machine.go:45286
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr850
		case 11:
			goto tr1851
		case 13:
			goto tr850
		case 32:
			goto tr902
		case 44:
			goto tr4
		case 61:
			goto tr365
		case 92:
			goto tr362
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr902
		}
		goto tr360
tr362:
//line machine.rl:17

	m.pb = m.p

	goto st390
	st390:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof390
		}
	st_case_390:
//line machine.go:45318
		if ( m.data)[( m.p)] == 92 {
			goto st3
		}
		switch {
		case ( m.data)[( m.p)] > 10:
			if 12 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto tr2
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr2
		}
		goto st205
	st1293:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1293
		}
	st_case_1293:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr1849
		case 13:
			goto tr859
		case 32:
			goto tr897
		case 44:
			goto tr4
		case 61:
			goto tr365
		case 92:
			goto st390
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1294
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr897
		}
		goto st205
	st1294:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1294
		}
	st_case_1294:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr1849
		case 13:
			goto tr859
		case 32:
			goto tr897
		case 44:
			goto tr4
		case 61:
			goto tr365
		case 92:
			goto st390
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1295
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr897
		}
		goto st205
	st1295:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1295
		}
	st_case_1295:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr1849
		case 13:
			goto tr859
		case 32:
			goto tr897
		case 44:
			goto tr4
		case 61:
			goto tr365
		case 92:
			goto st390
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1296
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr897
		}
		goto st205
	st1296:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1296
		}
	st_case_1296:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr1849
		case 13:
			goto tr859
		case 32:
			goto tr897
		case 44:
			goto tr4
		case 61:
			goto tr365
		case 92:
			goto st390
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1297
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr897
		}
		goto st205
	st1297:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1297
		}
	st_case_1297:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr1849
		case 13:
			goto tr859
		case 32:
			goto tr897
		case 44:
			goto tr4
		case 61:
			goto tr365
		case 92:
			goto st390
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1298
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr897
		}
		goto st205
	st1298:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1298
		}
	st_case_1298:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr1849
		case 13:
			goto tr859
		case 32:
			goto tr897
		case 44:
			goto tr4
		case 61:
			goto tr365
		case 92:
			goto st390
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1299
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr897
		}
		goto st205
	st1299:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1299
		}
	st_case_1299:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr1849
		case 13:
			goto tr859
		case 32:
			goto tr897
		case 44:
			goto tr4
		case 61:
			goto tr365
		case 92:
			goto st390
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1300
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr897
		}
		goto st205
	st1300:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1300
		}
	st_case_1300:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr1849
		case 13:
			goto tr859
		case 32:
			goto tr897
		case 44:
			goto tr4
		case 61:
			goto tr365
		case 92:
			goto st390
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1301
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr897
		}
		goto st205
	st1301:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1301
		}
	st_case_1301:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr1849
		case 13:
			goto tr859
		case 32:
			goto tr897
		case 44:
			goto tr4
		case 61:
			goto tr365
		case 92:
			goto st390
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1302
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr897
		}
		goto st205
	st1302:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1302
		}
	st_case_1302:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr1849
		case 13:
			goto tr859
		case 32:
			goto tr897
		case 44:
			goto tr4
		case 61:
			goto tr365
		case 92:
			goto st390
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1303
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr897
		}
		goto st205
	st1303:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1303
		}
	st_case_1303:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr1849
		case 13:
			goto tr859
		case 32:
			goto tr897
		case 44:
			goto tr4
		case 61:
			goto tr365
		case 92:
			goto st390
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1304
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr897
		}
		goto st205
	st1304:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1304
		}
	st_case_1304:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr1849
		case 13:
			goto tr859
		case 32:
			goto tr897
		case 44:
			goto tr4
		case 61:
			goto tr365
		case 92:
			goto st390
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1305
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr897
		}
		goto st205
	st1305:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1305
		}
	st_case_1305:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr1849
		case 13:
			goto tr859
		case 32:
			goto tr897
		case 44:
			goto tr4
		case 61:
			goto tr365
		case 92:
			goto st390
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1306
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr897
		}
		goto st205
	st1306:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1306
		}
	st_case_1306:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr1849
		case 13:
			goto tr859
		case 32:
			goto tr897
		case 44:
			goto tr4
		case 61:
			goto tr365
		case 92:
			goto st390
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1307
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr897
		}
		goto st205
	st1307:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1307
		}
	st_case_1307:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr1849
		case 13:
			goto tr859
		case 32:
			goto tr897
		case 44:
			goto tr4
		case 61:
			goto tr365
		case 92:
			goto st390
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1308
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr897
		}
		goto st205
	st1308:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1308
		}
	st_case_1308:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr1849
		case 13:
			goto tr859
		case 32:
			goto tr897
		case 44:
			goto tr4
		case 61:
			goto tr365
		case 92:
			goto st390
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1309
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr897
		}
		goto st205
	st1309:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1309
		}
	st_case_1309:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr1849
		case 13:
			goto tr859
		case 32:
			goto tr897
		case 44:
			goto tr4
		case 61:
			goto tr365
		case 92:
			goto st390
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1310
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr897
		}
		goto st205
	st1310:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1310
		}
	st_case_1310:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr1849
		case 13:
			goto tr859
		case 32:
			goto tr897
		case 44:
			goto tr4
		case 61:
			goto tr365
		case 92:
			goto st390
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr897
		}
		goto st205
tr1845:
//line machine.rl:58

	m.handler.SetMeasurement(m.text())

	goto st391
tr1869:
//line machine.rl:66

	m.handler.AddTag(key, m.text())

	goto st391
tr2085:
//line machine.rl:58

	m.handler.SetMeasurement(m.text())

//line machine.rl:78

	m.handler.AddFloat(key, m.text())

	goto st391
tr1870:
//line machine.rl:66

	m.handler.AddTag(key, m.text())

//line machine.rl:78

	m.handler.AddFloat(key, m.text())

	goto st391
tr1873:
//line machine.rl:66

	m.handler.AddTag(key, m.text())

//line machine.rl:74

	m.handler.AddInt(key, m.text())

	goto st391
tr1874:
//line machine.rl:66

	m.handler.AddTag(key, m.text())

//line machine.rl:82

	m.handler.AddBool(key, m.text())

	goto st391
tr2089:
//line machine.rl:58

	m.handler.SetMeasurement(m.text())

//line machine.rl:74

	m.handler.AddInt(key, m.text())

	goto st391
tr2091:
//line machine.rl:58

	m.handler.SetMeasurement(m.text())

//line machine.rl:82

	m.handler.AddBool(key, m.text())

	goto st391
	st391:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof391
		}
	st_case_391:
//line machine.go:45943
		switch ( m.data)[( m.p)] {
		case 32:
			goto tr48
		case 44:
			goto tr48
		case 61:
			goto tr48
		case 92:
			goto tr695
		}
		switch {
		case ( m.data)[( m.p)] > 10:
			if 12 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto tr48
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr48
		}
		goto tr694
tr694:
//line machine.rl:17

	m.pb = m.p

	goto st392
	st392:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof392
		}
	st_case_392:
//line machine.go:45974
		switch ( m.data)[( m.p)] {
		case 32:
			goto tr48
		case 44:
			goto tr48
		case 61:
			goto tr697
		case 92:
			goto st419
		}
		switch {
		case ( m.data)[( m.p)] > 10:
			if 12 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto tr48
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr48
		}
		goto st392
tr697:
//line machine.rl:62

	key = m.text()

//line machine.rl:70

	key = m.text()

	goto st393
	st393:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof393
		}
	st_case_393:
//line machine.go:46009
		switch ( m.data)[( m.p)] {
		case 32:
			goto tr48
		case 34:
			goto tr699
		case 44:
			goto tr48
		case 45:
			goto tr700
		case 46:
			goto tr701
		case 48:
			goto tr702
		case 61:
			goto tr48
		case 70:
			goto tr704
		case 84:
			goto tr705
		case 92:
			goto tr45
		case 102:
			goto tr706
		case 116:
			goto tr707
		}
		switch {
		case ( m.data)[( m.p)] < 12:
			if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 10 {
				goto tr48
			}
		case ( m.data)[( m.p)] > 13:
			if 49 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto tr703
			}
		default:
			goto tr48
		}
		goto tr44
tr699:
//line machine.rl:17

	m.pb = m.p

	goto st394
	st394:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof394
		}
	st_case_394:
//line machine.go:46060
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr22
		case 11:
			goto tr710
		case 13:
			goto tr22
		case 32:
			goto tr709
		case 34:
			goto tr711
		case 44:
			goto tr712
		case 61:
			goto tr22
		case 92:
			goto tr713
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr709
		}
		goto tr708
tr708:
//line machine.rl:17

	m.pb = m.p

	goto st395
	st395:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof395
		}
	st_case_395:
//line machine.go:46094
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr716
		case 13:
			goto st6
		case 32:
			goto tr715
		case 34:
			goto tr717
		case 44:
			goto tr718
		case 61:
			goto st6
		case 92:
			goto st407
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr715
		}
		goto st395
tr715:
//line machine.rl:66

	m.handler.AddTag(key, m.text())

	goto st396
tr709:
//line machine.rl:66

	m.handler.AddTag(key, m.text())

//line machine.rl:17

	m.pb = m.p

	goto st396
	st396:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof396
		}
	st_case_396:
//line machine.go:46138
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr722
		case 13:
			goto st6
		case 32:
			goto st396
		case 34:
			goto tr82
		case 44:
			goto st6
		case 61:
			goto st6
		case 92:
			goto tr723
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto st396
		}
		goto tr720
tr720:
//line machine.rl:17

	m.pb = m.p

	goto st397
	st397:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof397
		}
	st_case_397:
//line machine.go:46172
		switch ( m.data)[( m.p)] {
		case 32:
			goto st6
		case 34:
			goto tr85
		case 44:
			goto st6
		case 61:
			goto tr725
		case 92:
			goto st399
		}
		switch {
		case ( m.data)[( m.p)] > 10:
			if 12 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto st6
			}
		case ( m.data)[( m.p)] >= 9:
			goto st6
		}
		goto st397
tr725:
//line machine.rl:70

	key = m.text()

	goto st398
	st398:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof398
		}
	st_case_398:
//line machine.go:46205
		switch ( m.data)[( m.p)] {
		case 34:
			goto tr395
		case 45:
			goto tr662
		case 46:
			goto tr663
		case 48:
			goto tr664
		case 70:
			goto tr666
		case 84:
			goto tr667
		case 92:
			goto st59
		case 102:
			goto tr668
		case 116:
			goto tr669
		}
		if 49 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
			goto tr665
		}
		goto st6
tr723:
//line machine.rl:17

	m.pb = m.p

	goto st399
	st399:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof399
		}
	st_case_399:
//line machine.go:46241
		switch ( m.data)[( m.p)] {
		case 34:
			goto st397
		case 92:
			goto st397
		}
		switch {
		case ( m.data)[( m.p)] > 10:
			if 12 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto tr2
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr2
		}
		goto st3
tr722:
//line machine.rl:17

	m.pb = m.p

	goto st400
	st400:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof400
		}
	st_case_400:
//line machine.go:46268
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr722
		case 13:
			goto st6
		case 32:
			goto st396
		case 34:
			goto tr82
		case 44:
			goto st6
		case 61:
			goto tr725
		case 92:
			goto tr723
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto st396
		}
		goto tr720
tr716:
//line machine.rl:66

	m.handler.AddTag(key, m.text())

	goto st401
tr710:
//line machine.rl:66

	m.handler.AddTag(key, m.text())

//line machine.rl:17

	m.pb = m.p

	goto st401
	st401:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof401
		}
	st_case_401:
//line machine.go:46312
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr728
		case 13:
			goto st6
		case 32:
			goto tr715
		case 34:
			goto tr729
		case 44:
			goto tr718
		case 61:
			goto st6
		case 92:
			goto tr730
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr715
		}
		goto tr727
tr727:
//line machine.rl:17

	m.pb = m.p

	goto st402
	st402:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof402
		}
	st_case_402:
//line machine.go:46346
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr732
		case 13:
			goto st6
		case 32:
			goto tr715
		case 34:
			goto tr733
		case 44:
			goto tr718
		case 61:
			goto tr725
		case 92:
			goto st409
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr715
		}
		goto st402
tr732:
//line machine.rl:66

	m.handler.AddTag(key, m.text())

	goto st403
tr728:
//line machine.rl:66

	m.handler.AddTag(key, m.text())

//line machine.rl:17

	m.pb = m.p

	goto st403
	st403:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof403
		}
	st_case_403:
//line machine.go:46390
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr728
		case 13:
			goto st6
		case 32:
			goto tr715
		case 34:
			goto tr729
		case 44:
			goto tr718
		case 61:
			goto tr725
		case 92:
			goto tr730
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr715
		}
		goto tr727
tr729:
//line machine.rl:17

	m.pb = m.p

//line machine.rl:86

	m.handler.AddString(key, m.text())

	goto st1311
tr733:
//line machine.rl:86

	m.handler.AddString(key, m.text())

	goto st1311
	st1311:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1311
		}
	st_case_1311:
//line machine.go:46434
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr850
		case 11:
			goto tr1016
		case 13:
			goto tr850
		case 32:
			goto tr954
		case 44:
			goto tr1869
		case 61:
			goto tr11
		case 92:
			goto st18
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr954
		}
		goto st16
tr718:
//line machine.rl:66

	m.handler.AddTag(key, m.text())

	goto st404
tr712:
//line machine.rl:66

	m.handler.AddTag(key, m.text())

//line machine.rl:17

	m.pb = m.p

	goto st404
	st404:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof404
		}
	st_case_404:
//line machine.go:46476
		switch ( m.data)[( m.p)] {
		case 32:
			goto st6
		case 34:
			goto tr171
		case 44:
			goto st6
		case 61:
			goto st6
		case 92:
			goto tr736
		}
		switch {
		case ( m.data)[( m.p)] > 10:
			if 12 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto st6
			}
		case ( m.data)[( m.p)] >= 9:
			goto st6
		}
		goto tr735
tr735:
//line machine.rl:17

	m.pb = m.p

	goto st405
	st405:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof405
		}
	st_case_405:
//line machine.go:46509
		switch ( m.data)[( m.p)] {
		case 32:
			goto st6
		case 34:
			goto tr174
		case 44:
			goto st6
		case 61:
			goto tr738
		case 92:
			goto st408
		}
		switch {
		case ( m.data)[( m.p)] > 10:
			if 12 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto st6
			}
		case ( m.data)[( m.p)] >= 9:
			goto st6
		}
		goto st405
tr738:
//line machine.rl:62

	key = m.text()

	goto st406
	st406:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof406
		}
	st_case_406:
//line machine.go:46542
		switch ( m.data)[( m.p)] {
		case 32:
			goto st6
		case 34:
			goto tr711
		case 44:
			goto st6
		case 61:
			goto st6
		case 92:
			goto tr713
		}
		switch {
		case ( m.data)[( m.p)] > 10:
			if 12 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto st6
			}
		case ( m.data)[( m.p)] >= 9:
			goto st6
		}
		goto tr708
tr711:
//line machine.rl:17

	m.pb = m.p

//line machine.rl:86

	m.handler.AddString(key, m.text())

	goto st1312
tr717:
//line machine.rl:86

	m.handler.AddString(key, m.text())

	goto st1312
	st1312:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1312
		}
	st_case_1312:
//line machine.go:46585
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr850
		case 11:
			goto tr955
		case 13:
			goto tr850
		case 32:
			goto tr954
		case 44:
			goto tr1869
		case 61:
			goto tr48
		case 92:
			goto st19
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr954
		}
		goto st14
tr713:
//line machine.rl:17

	m.pb = m.p

	goto st407
	st407:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof407
		}
	st_case_407:
//line machine.go:46617
		switch ( m.data)[( m.p)] {
		case 34:
			goto st395
		case 92:
			goto st395
		}
		switch {
		case ( m.data)[( m.p)] > 10:
			if 12 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto tr48
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr48
		}
		goto st14
tr736:
//line machine.rl:17

	m.pb = m.p

	goto st408
	st408:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof408
		}
	st_case_408:
//line machine.go:46644
		switch ( m.data)[( m.p)] {
		case 34:
			goto st405
		case 92:
			goto st405
		}
		switch {
		case ( m.data)[( m.p)] > 10:
			if 12 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto tr48
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr48
		}
		goto st12
tr730:
//line machine.rl:17

	m.pb = m.p

	goto st409
	st409:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof409
		}
	st_case_409:
//line machine.go:46671
		switch ( m.data)[( m.p)] {
		case 34:
			goto st402
		case 92:
			goto st402
		}
		switch {
		case ( m.data)[( m.p)] > 10:
			if 12 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto tr48
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr48
		}
		goto st16
tr700:
//line machine.rl:17

	m.pb = m.p

	goto st410
	st410:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof410
		}
	st_case_410:
//line machine.go:46698
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr48
		case 11:
			goto tr49
		case 13:
			goto tr48
		case 32:
			goto tr47
		case 44:
			goto tr50
		case 48:
			goto st1313
		case 61:
			goto tr48
		case 92:
			goto st19
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 49 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1316
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr47
		}
		goto st14
tr702:
//line machine.rl:17

	m.pb = m.p

	goto st1313
	st1313:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1313
		}
	st_case_1313:
//line machine.go:46737
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1270
		case 11:
			goto tr1271
		case 13:
			goto tr1270
		case 32:
			goto tr1269
		case 44:
			goto tr1870
		case 46:
			goto st1314
		case 61:
			goto tr48
		case 92:
			goto st19
		case 105:
			goto st1315
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1269
		}
		goto st14
tr701:
//line machine.rl:17

	m.pb = m.p

	goto st1314
	st1314:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1314
		}
	st_case_1314:
//line machine.go:46773
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1270
		case 11:
			goto tr1271
		case 13:
			goto tr1270
		case 32:
			goto tr1269
		case 44:
			goto tr1870
		case 61:
			goto tr48
		case 92:
			goto st19
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1314
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1269
		}
		goto st14
	st1315:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1315
		}
	st_case_1315:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1276
		case 11:
			goto tr1277
		case 13:
			goto tr1276
		case 32:
			goto tr1275
		case 44:
			goto tr1873
		case 61:
			goto tr48
		case 92:
			goto st19
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1275
		}
		goto st14
tr703:
//line machine.rl:17

	m.pb = m.p

	goto st1316
	st1316:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1316
		}
	st_case_1316:
//line machine.go:46835
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1270
		case 11:
			goto tr1271
		case 13:
			goto tr1270
		case 32:
			goto tr1269
		case 44:
			goto tr1870
		case 46:
			goto st1314
		case 61:
			goto tr48
		case 92:
			goto st19
		case 105:
			goto st1315
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1316
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1269
		}
		goto st14
tr704:
//line machine.rl:17

	m.pb = m.p

	goto st1317
	st1317:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1317
		}
	st_case_1317:
//line machine.go:46876
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1280
		case 11:
			goto tr1281
		case 13:
			goto tr1280
		case 32:
			goto tr1279
		case 44:
			goto tr1874
		case 61:
			goto tr48
		case 65:
			goto st411
		case 92:
			goto st19
		case 97:
			goto st414
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1279
		}
		goto st14
	st411:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof411
		}
	st_case_411:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr48
		case 11:
			goto tr49
		case 13:
			goto tr48
		case 32:
			goto tr47
		case 44:
			goto tr50
		case 61:
			goto tr48
		case 76:
			goto st412
		case 92:
			goto st19
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr47
		}
		goto st14
	st412:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof412
		}
	st_case_412:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr48
		case 11:
			goto tr49
		case 13:
			goto tr48
		case 32:
			goto tr47
		case 44:
			goto tr50
		case 61:
			goto tr48
		case 83:
			goto st413
		case 92:
			goto st19
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr47
		}
		goto st14
	st413:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof413
		}
	st_case_413:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr48
		case 11:
			goto tr49
		case 13:
			goto tr48
		case 32:
			goto tr47
		case 44:
			goto tr50
		case 61:
			goto tr48
		case 69:
			goto st1318
		case 92:
			goto st19
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr47
		}
		goto st14
	st1318:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1318
		}
	st_case_1318:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1280
		case 11:
			goto tr1281
		case 13:
			goto tr1280
		case 32:
			goto tr1279
		case 44:
			goto tr1874
		case 61:
			goto tr48
		case 92:
			goto st19
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1279
		}
		goto st14
	st414:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof414
		}
	st_case_414:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr48
		case 11:
			goto tr49
		case 13:
			goto tr48
		case 32:
			goto tr47
		case 44:
			goto tr50
		case 61:
			goto tr48
		case 92:
			goto st19
		case 108:
			goto st415
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr47
		}
		goto st14
	st415:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof415
		}
	st_case_415:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr48
		case 11:
			goto tr49
		case 13:
			goto tr48
		case 32:
			goto tr47
		case 44:
			goto tr50
		case 61:
			goto tr48
		case 92:
			goto st19
		case 115:
			goto st416
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr47
		}
		goto st14
	st416:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof416
		}
	st_case_416:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr48
		case 11:
			goto tr49
		case 13:
			goto tr48
		case 32:
			goto tr47
		case 44:
			goto tr50
		case 61:
			goto tr48
		case 92:
			goto st19
		case 101:
			goto st1318
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr47
		}
		goto st14
tr705:
//line machine.rl:17

	m.pb = m.p

	goto st1319
	st1319:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1319
		}
	st_case_1319:
//line machine.go:47099
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1280
		case 11:
			goto tr1281
		case 13:
			goto tr1280
		case 32:
			goto tr1279
		case 44:
			goto tr1874
		case 61:
			goto tr48
		case 82:
			goto st417
		case 92:
			goto st19
		case 114:
			goto st418
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1279
		}
		goto st14
	st417:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof417
		}
	st_case_417:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr48
		case 11:
			goto tr49
		case 13:
			goto tr48
		case 32:
			goto tr47
		case 44:
			goto tr50
		case 61:
			goto tr48
		case 85:
			goto st413
		case 92:
			goto st19
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr47
		}
		goto st14
	st418:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof418
		}
	st_case_418:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr48
		case 11:
			goto tr49
		case 13:
			goto tr48
		case 32:
			goto tr47
		case 44:
			goto tr50
		case 61:
			goto tr48
		case 92:
			goto st19
		case 117:
			goto st416
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr47
		}
		goto st14
tr706:
//line machine.rl:17

	m.pb = m.p

	goto st1320
	st1320:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1320
		}
	st_case_1320:
//line machine.go:47189
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1280
		case 11:
			goto tr1281
		case 13:
			goto tr1280
		case 32:
			goto tr1279
		case 44:
			goto tr1874
		case 61:
			goto tr48
		case 92:
			goto st19
		case 97:
			goto st414
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1279
		}
		goto st14
tr707:
//line machine.rl:17

	m.pb = m.p

	goto st1321
	st1321:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1321
		}
	st_case_1321:
//line machine.go:47223
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1280
		case 11:
			goto tr1281
		case 13:
			goto tr1280
		case 32:
			goto tr1279
		case 44:
			goto tr1874
		case 61:
			goto tr48
		case 92:
			goto st19
		case 114:
			goto st418
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1279
		}
		goto st14
tr695:
//line machine.rl:17

	m.pb = m.p

	goto st419
	st419:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof419
		}
	st_case_419:
//line machine.go:47257
		switch {
		case ( m.data)[( m.p)] > 10:
			if 12 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto tr48
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr48
		}
		goto st392
tr386:
//line machine.rl:58

	m.handler.SetMeasurement(m.text())

	goto st420
tr380:
//line machine.rl:58

	m.handler.SetMeasurement(m.text())

//line machine.rl:17

	m.pb = m.p

	goto st420
tr757:
//line machine.rl:66

	m.handler.AddTag(key, m.text())

	goto st420
	st420:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof420
		}
	st_case_420:
//line machine.go:47294
		switch ( m.data)[( m.p)] {
		case 32:
			goto st6
		case 34:
			goto tr171
		case 44:
			goto st6
		case 61:
			goto st6
		case 92:
			goto tr748
		}
		switch {
		case ( m.data)[( m.p)] > 10:
			if 12 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto st6
			}
		case ( m.data)[( m.p)] >= 9:
			goto st6
		}
		goto tr747
tr747:
//line machine.rl:17

	m.pb = m.p

	goto st421
	st421:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof421
		}
	st_case_421:
//line machine.go:47327
		switch ( m.data)[( m.p)] {
		case 32:
			goto st6
		case 34:
			goto tr174
		case 44:
			goto st6
		case 61:
			goto tr750
		case 92:
			goto st429
		}
		switch {
		case ( m.data)[( m.p)] > 10:
			if 12 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto st6
			}
		case ( m.data)[( m.p)] >= 9:
			goto st6
		}
		goto st421
tr750:
//line machine.rl:62

	key = m.text()

	goto st422
	st422:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof422
		}
	st_case_422:
//line machine.go:47360
		switch ( m.data)[( m.p)] {
		case 32:
			goto st6
		case 34:
			goto tr711
		case 44:
			goto st6
		case 61:
			goto st6
		case 92:
			goto tr753
		}
		switch {
		case ( m.data)[( m.p)] > 10:
			if 12 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto st6
			}
		case ( m.data)[( m.p)] >= 9:
			goto st6
		}
		goto tr752
tr752:
//line machine.rl:17

	m.pb = m.p

	goto st423
	st423:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof423
		}
	st_case_423:
//line machine.go:47393
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr756
		case 13:
			goto st6
		case 32:
			goto tr755
		case 34:
			goto tr717
		case 44:
			goto tr757
		case 61:
			goto st6
		case 92:
			goto st428
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr755
		}
		goto st423
tr756:
//line machine.rl:66

	m.handler.AddTag(key, m.text())

	goto st424
	st424:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof424
		}
	st_case_424:
//line machine.go:47427
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr760
		case 13:
			goto st6
		case 32:
			goto tr755
		case 34:
			goto tr729
		case 44:
			goto tr757
		case 61:
			goto st6
		case 92:
			goto tr761
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr755
		}
		goto tr759
tr759:
//line machine.rl:17

	m.pb = m.p

	goto st425
	st425:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof425
		}
	st_case_425:
//line machine.go:47461
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr763
		case 13:
			goto st6
		case 32:
			goto tr755
		case 34:
			goto tr733
		case 44:
			goto tr757
		case 61:
			goto tr393
		case 92:
			goto st427
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr755
		}
		goto st425
tr763:
//line machine.rl:66

	m.handler.AddTag(key, m.text())

	goto st426
tr760:
//line machine.rl:66

	m.handler.AddTag(key, m.text())

//line machine.rl:17

	m.pb = m.p

	goto st426
	st426:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof426
		}
	st_case_426:
//line machine.go:47505
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr760
		case 13:
			goto st6
		case 32:
			goto tr755
		case 34:
			goto tr729
		case 44:
			goto tr757
		case 61:
			goto tr393
		case 92:
			goto tr761
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr755
		}
		goto tr759
tr761:
//line machine.rl:17

	m.pb = m.p

	goto st427
	st427:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof427
		}
	st_case_427:
//line machine.go:47539
		switch ( m.data)[( m.p)] {
		case 34:
			goto st425
		case 92:
			goto st425
		}
		switch {
		case ( m.data)[( m.p)] > 10:
			if 12 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto tr48
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr48
		}
		goto st16
tr753:
//line machine.rl:17

	m.pb = m.p

	goto st428
	st428:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof428
		}
	st_case_428:
//line machine.go:47566
		switch ( m.data)[( m.p)] {
		case 34:
			goto st423
		case 92:
			goto st423
		}
		switch {
		case ( m.data)[( m.p)] > 10:
			if 12 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto tr48
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr48
		}
		goto st14
tr748:
//line machine.rl:17

	m.pb = m.p

	goto st429
	st429:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof429
		}
	st_case_429:
//line machine.go:47593
		switch ( m.data)[( m.p)] {
		case 34:
			goto st421
		case 92:
			goto st421
		}
		switch {
		case ( m.data)[( m.p)] > 10:
			if 12 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto tr48
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr48
		}
		goto st12
tr691:
//line machine.rl:70

	key = m.text()

	goto st430
	st430:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof430
		}
	st_case_430:
//line machine.go:47620
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr384
		case 13:
			goto st6
		case 32:
			goto tr383
		case 34:
			goto tr765
		case 44:
			goto tr386
		case 45:
			goto tr766
		case 46:
			goto tr767
		case 48:
			goto tr768
		case 70:
			goto tr770
		case 84:
			goto tr771
		case 92:
			goto st465
		case 102:
			goto tr772
		case 116:
			goto tr773
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 49 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto tr769
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr383
		}
		goto st209
tr765:
//line machine.rl:86

	m.handler.AddString(key, m.text())

	goto st1322
	st1322:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1322
		}
	st_case_1322:
//line machine.go:47671
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1363
		case 11:
			goto tr1880
		case 13:
			goto tr1363
		case 32:
			goto tr1879
		case 34:
			goto tr379
		case 44:
			goto tr1881
		case 92:
			goto tr381
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1879
		}
		goto tr376
tr1906:
//line machine.rl:58

	m.handler.SetMeasurement(m.text())

	goto st1323
tr1879:
//line machine.rl:58

	m.handler.SetMeasurement(m.text())

//line machine.rl:17

	m.pb = m.p

	goto st1323
tr2033:
//line machine.rl:66

	m.handler.AddTag(key, m.text())

	goto st1323
tr2069:
//line machine.rl:58

	m.handler.SetMeasurement(m.text())

//line machine.rl:78

	m.handler.AddFloat(key, m.text())

	goto st1323
tr2028:
//line machine.rl:66

	m.handler.AddTag(key, m.text())

//line machine.rl:78

	m.handler.AddFloat(key, m.text())

	goto st1323
tr2059:
//line machine.rl:66

	m.handler.AddTag(key, m.text())

//line machine.rl:74

	m.handler.AddInt(key, m.text())

	goto st1323
tr2062:
//line machine.rl:66

	m.handler.AddTag(key, m.text())

//line machine.rl:82

	m.handler.AddBool(key, m.text())

	goto st1323
tr2074:
//line machine.rl:58

	m.handler.SetMeasurement(m.text())

//line machine.rl:74

	m.handler.AddInt(key, m.text())

	goto st1323
tr2077:
//line machine.rl:58

	m.handler.SetMeasurement(m.text())

//line machine.rl:82

	m.handler.AddBool(key, m.text())

	goto st1323
	st1323:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1323
		}
	st_case_1323:
//line machine.go:47779
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1366
		case 11:
			goto tr1883
		case 13:
			goto tr1366
		case 32:
			goto st1323
		case 34:
			goto tr82
		case 44:
			goto st6
		case 45:
			goto tr1884
		case 61:
			goto st6
		case 92:
			goto tr391
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto tr1885
			}
		case ( m.data)[( m.p)] >= 9:
			goto st1323
		}
		goto tr388
tr1883:
//line machine.rl:17

	m.pb = m.p

	goto st1324
	st1324:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1324
		}
	st_case_1324:
//line machine.go:47820
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1366
		case 11:
			goto tr1883
		case 13:
			goto tr1366
		case 32:
			goto st1323
		case 34:
			goto tr82
		case 44:
			goto st6
		case 45:
			goto tr1884
		case 61:
			goto tr393
		case 92:
			goto tr391
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto tr1885
			}
		case ( m.data)[( m.p)] >= 9:
			goto st1323
		}
		goto tr388
tr1884:
//line machine.rl:17

	m.pb = m.p

	goto st431
	st431:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof431
		}
	st_case_431:
//line machine.go:47861
		switch ( m.data)[( m.p)] {
		case 32:
			goto st6
		case 34:
			goto tr85
		case 44:
			goto st6
		case 61:
			goto tr393
		case 92:
			goto st384
		}
		switch {
		case ( m.data)[( m.p)] < 12:
			if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 10 {
				goto st6
			}
		case ( m.data)[( m.p)] > 13:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1325
			}
		default:
			goto st6
		}
		goto st211
tr1885:
//line machine.rl:17

	m.pb = m.p

	goto st1325
	st1325:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1325
		}
	st_case_1325:
//line machine.go:47898
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 11:
			goto tr1886
		case 13:
			goto tr1595
		case 32:
			goto tr1594
		case 34:
			goto tr85
		case 44:
			goto st6
		case 61:
			goto tr393
		case 92:
			goto st384
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1327
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1594
		}
		goto st211
tr1886:
//line machine.rl:90

	m.handler.SetTimestamp(m.text())

	goto st1326
	st1326:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1326
		}
	st_case_1326:
//line machine.go:47937
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1366
		case 11:
			goto st1326
		case 13:
			goto tr1366
		case 32:
			goto st1081
		case 34:
			goto tr85
		case 44:
			goto st6
		case 61:
			goto tr393
		case 92:
			goto st384
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto st1081
		}
		goto st211
	st1327:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1327
		}
	st_case_1327:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 11:
			goto tr1886
		case 13:
			goto tr1595
		case 32:
			goto tr1594
		case 34:
			goto tr85
		case 44:
			goto st6
		case 61:
			goto tr393
		case 92:
			goto st384
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1328
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1594
		}
		goto st211
	st1328:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1328
		}
	st_case_1328:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 11:
			goto tr1886
		case 13:
			goto tr1595
		case 32:
			goto tr1594
		case 34:
			goto tr85
		case 44:
			goto st6
		case 61:
			goto tr393
		case 92:
			goto st384
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1329
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1594
		}
		goto st211
	st1329:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1329
		}
	st_case_1329:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 11:
			goto tr1886
		case 13:
			goto tr1595
		case 32:
			goto tr1594
		case 34:
			goto tr85
		case 44:
			goto st6
		case 61:
			goto tr393
		case 92:
			goto st384
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1330
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1594
		}
		goto st211
	st1330:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1330
		}
	st_case_1330:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 11:
			goto tr1886
		case 13:
			goto tr1595
		case 32:
			goto tr1594
		case 34:
			goto tr85
		case 44:
			goto st6
		case 61:
			goto tr393
		case 92:
			goto st384
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1331
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1594
		}
		goto st211
	st1331:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1331
		}
	st_case_1331:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 11:
			goto tr1886
		case 13:
			goto tr1595
		case 32:
			goto tr1594
		case 34:
			goto tr85
		case 44:
			goto st6
		case 61:
			goto tr393
		case 92:
			goto st384
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1332
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1594
		}
		goto st211
	st1332:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1332
		}
	st_case_1332:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 11:
			goto tr1886
		case 13:
			goto tr1595
		case 32:
			goto tr1594
		case 34:
			goto tr85
		case 44:
			goto st6
		case 61:
			goto tr393
		case 92:
			goto st384
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1333
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1594
		}
		goto st211
	st1333:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1333
		}
	st_case_1333:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 11:
			goto tr1886
		case 13:
			goto tr1595
		case 32:
			goto tr1594
		case 34:
			goto tr85
		case 44:
			goto st6
		case 61:
			goto tr393
		case 92:
			goto st384
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1334
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1594
		}
		goto st211
	st1334:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1334
		}
	st_case_1334:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 11:
			goto tr1886
		case 13:
			goto tr1595
		case 32:
			goto tr1594
		case 34:
			goto tr85
		case 44:
			goto st6
		case 61:
			goto tr393
		case 92:
			goto st384
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1335
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1594
		}
		goto st211
	st1335:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1335
		}
	st_case_1335:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 11:
			goto tr1886
		case 13:
			goto tr1595
		case 32:
			goto tr1594
		case 34:
			goto tr85
		case 44:
			goto st6
		case 61:
			goto tr393
		case 92:
			goto st384
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1336
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1594
		}
		goto st211
	st1336:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1336
		}
	st_case_1336:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 11:
			goto tr1886
		case 13:
			goto tr1595
		case 32:
			goto tr1594
		case 34:
			goto tr85
		case 44:
			goto st6
		case 61:
			goto tr393
		case 92:
			goto st384
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1337
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1594
		}
		goto st211
	st1337:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1337
		}
	st_case_1337:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 11:
			goto tr1886
		case 13:
			goto tr1595
		case 32:
			goto tr1594
		case 34:
			goto tr85
		case 44:
			goto st6
		case 61:
			goto tr393
		case 92:
			goto st384
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1338
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1594
		}
		goto st211
	st1338:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1338
		}
	st_case_1338:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 11:
			goto tr1886
		case 13:
			goto tr1595
		case 32:
			goto tr1594
		case 34:
			goto tr85
		case 44:
			goto st6
		case 61:
			goto tr393
		case 92:
			goto st384
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1339
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1594
		}
		goto st211
	st1339:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1339
		}
	st_case_1339:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 11:
			goto tr1886
		case 13:
			goto tr1595
		case 32:
			goto tr1594
		case 34:
			goto tr85
		case 44:
			goto st6
		case 61:
			goto tr393
		case 92:
			goto st384
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1340
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1594
		}
		goto st211
	st1340:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1340
		}
	st_case_1340:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 11:
			goto tr1886
		case 13:
			goto tr1595
		case 32:
			goto tr1594
		case 34:
			goto tr85
		case 44:
			goto st6
		case 61:
			goto tr393
		case 92:
			goto st384
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1341
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1594
		}
		goto st211
	st1341:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1341
		}
	st_case_1341:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 11:
			goto tr1886
		case 13:
			goto tr1595
		case 32:
			goto tr1594
		case 34:
			goto tr85
		case 44:
			goto st6
		case 61:
			goto tr393
		case 92:
			goto st384
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1342
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1594
		}
		goto st211
	st1342:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1342
		}
	st_case_1342:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 11:
			goto tr1886
		case 13:
			goto tr1595
		case 32:
			goto tr1594
		case 34:
			goto tr85
		case 44:
			goto st6
		case 61:
			goto tr393
		case 92:
			goto st384
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1343
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1594
		}
		goto st211
	st1343:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1343
		}
	st_case_1343:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 11:
			goto tr1886
		case 13:
			goto tr1595
		case 32:
			goto tr1594
		case 34:
			goto tr85
		case 44:
			goto st6
		case 61:
			goto tr393
		case 92:
			goto st384
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1344
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1594
		}
		goto st211
	st1344:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1344
		}
	st_case_1344:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 11:
			goto tr1886
		case 13:
			goto tr1595
		case 32:
			goto tr1594
		case 34:
			goto tr85
		case 44:
			goto st6
		case 61:
			goto tr393
		case 92:
			goto st384
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1594
		}
		goto st211
tr1880:
//line machine.rl:58

	m.handler.SetMeasurement(m.text())

//line machine.rl:17

	m.pb = m.p

	goto st1345
tr2070:
//line machine.rl:58

	m.handler.SetMeasurement(m.text())

//line machine.rl:78

	m.handler.AddFloat(key, m.text())

	goto st1345
tr2075:
//line machine.rl:58

	m.handler.SetMeasurement(m.text())

//line machine.rl:74

	m.handler.AddInt(key, m.text())

	goto st1345
tr2078:
//line machine.rl:58

	m.handler.SetMeasurement(m.text())

//line machine.rl:82

	m.handler.AddBool(key, m.text())

	goto st1345
	st1345:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1345
		}
	st_case_1345:
//line machine.go:48576
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1366
		case 11:
			goto tr1907
		case 13:
			goto tr1366
		case 32:
			goto tr1906
		case 34:
			goto tr686
		case 44:
			goto tr386
		case 45:
			goto tr1908
		case 61:
			goto st209
		case 92:
			goto tr687
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto tr1909
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1906
		}
		goto tr684
tr1907:
//line machine.rl:58

	m.handler.SetMeasurement(m.text())

//line machine.rl:17

	m.pb = m.p

	goto st1346
	st1346:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1346
		}
	st_case_1346:
//line machine.go:48621
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1366
		case 11:
			goto tr1907
		case 13:
			goto tr1366
		case 32:
			goto tr1906
		case 34:
			goto tr686
		case 44:
			goto tr386
		case 45:
			goto tr1908
		case 61:
			goto tr691
		case 92:
			goto tr687
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto tr1909
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1906
		}
		goto tr684
tr1908:
//line machine.rl:17

	m.pb = m.p

	goto st432
	st432:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof432
		}
	st_case_432:
//line machine.go:48662
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr689
		case 13:
			goto st6
		case 32:
			goto tr383
		case 34:
			goto tr690
		case 44:
			goto tr386
		case 61:
			goto tr691
		case 92:
			goto st433
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1347
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr383
		}
		goto st387
tr1909:
//line machine.rl:17

	m.pb = m.p

	goto st1347
	st1347:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1347
		}
	st_case_1347:
//line machine.go:48701
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 11:
			goto tr1911
		case 13:
			goto tr1595
		case 32:
			goto tr1910
		case 34:
			goto tr690
		case 44:
			goto tr386
		case 61:
			goto tr691
		case 92:
			goto st433
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1351
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1910
		}
		goto st387
tr1915:
//line machine.rl:58

	m.handler.SetMeasurement(m.text())

	goto st1348
tr2040:
//line machine.rl:66

	m.handler.AddTag(key, m.text())

	goto st1348
tr1910:
//line machine.rl:58

	m.handler.SetMeasurement(m.text())

//line machine.rl:90

	m.handler.SetTimestamp(m.text())

	goto st1348
tr2037:
//line machine.rl:66

	m.handler.AddTag(key, m.text())

//line machine.rl:90

	m.handler.SetTimestamp(m.text())

	goto st1348
	st1348:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1348
		}
	st_case_1348:
//line machine.go:48766
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1366
		case 11:
			goto tr1914
		case 13:
			goto tr1366
		case 32:
			goto st1348
		case 34:
			goto tr82
		case 44:
			goto st6
		case 61:
			goto st6
		case 92:
			goto tr391
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto st1348
		}
		goto tr388
tr1914:
//line machine.rl:17

	m.pb = m.p

	goto st1349
	st1349:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1349
		}
	st_case_1349:
//line machine.go:48800
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1366
		case 11:
			goto tr1914
		case 13:
			goto tr1366
		case 32:
			goto st1348
		case 34:
			goto tr82
		case 44:
			goto st6
		case 61:
			goto tr393
		case 92:
			goto tr391
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto st1348
		}
		goto tr388
tr1916:
//line machine.rl:58

	m.handler.SetMeasurement(m.text())

//line machine.rl:17

	m.pb = m.p

	goto st1350
tr1911:
//line machine.rl:58

	m.handler.SetMeasurement(m.text())

//line machine.rl:90

	m.handler.SetTimestamp(m.text())

	goto st1350
	st1350:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1350
		}
	st_case_1350:
//line machine.go:48848
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1366
		case 11:
			goto tr1916
		case 13:
			goto tr1366
		case 32:
			goto tr1915
		case 34:
			goto tr686
		case 44:
			goto tr386
		case 61:
			goto tr691
		case 92:
			goto tr687
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1915
		}
		goto tr684
tr687:
//line machine.rl:17

	m.pb = m.p

	goto st433
	st433:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof433
		}
	st_case_433:
//line machine.go:48882
		switch ( m.data)[( m.p)] {
		case 34:
			goto st387
		case 92:
			goto st211
		}
		switch {
		case ( m.data)[( m.p)] > 10:
			if 12 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto tr2
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr2
		}
		goto st205
	st1351:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1351
		}
	st_case_1351:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 11:
			goto tr1911
		case 13:
			goto tr1595
		case 32:
			goto tr1910
		case 34:
			goto tr690
		case 44:
			goto tr386
		case 61:
			goto tr691
		case 92:
			goto st433
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1352
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1910
		}
		goto st387
	st1352:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1352
		}
	st_case_1352:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 11:
			goto tr1911
		case 13:
			goto tr1595
		case 32:
			goto tr1910
		case 34:
			goto tr690
		case 44:
			goto tr386
		case 61:
			goto tr691
		case 92:
			goto st433
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1353
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1910
		}
		goto st387
	st1353:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1353
		}
	st_case_1353:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 11:
			goto tr1911
		case 13:
			goto tr1595
		case 32:
			goto tr1910
		case 34:
			goto tr690
		case 44:
			goto tr386
		case 61:
			goto tr691
		case 92:
			goto st433
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1354
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1910
		}
		goto st387
	st1354:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1354
		}
	st_case_1354:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 11:
			goto tr1911
		case 13:
			goto tr1595
		case 32:
			goto tr1910
		case 34:
			goto tr690
		case 44:
			goto tr386
		case 61:
			goto tr691
		case 92:
			goto st433
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1355
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1910
		}
		goto st387
	st1355:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1355
		}
	st_case_1355:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 11:
			goto tr1911
		case 13:
			goto tr1595
		case 32:
			goto tr1910
		case 34:
			goto tr690
		case 44:
			goto tr386
		case 61:
			goto tr691
		case 92:
			goto st433
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1356
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1910
		}
		goto st387
	st1356:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1356
		}
	st_case_1356:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 11:
			goto tr1911
		case 13:
			goto tr1595
		case 32:
			goto tr1910
		case 34:
			goto tr690
		case 44:
			goto tr386
		case 61:
			goto tr691
		case 92:
			goto st433
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1357
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1910
		}
		goto st387
	st1357:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1357
		}
	st_case_1357:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 11:
			goto tr1911
		case 13:
			goto tr1595
		case 32:
			goto tr1910
		case 34:
			goto tr690
		case 44:
			goto tr386
		case 61:
			goto tr691
		case 92:
			goto st433
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1358
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1910
		}
		goto st387
	st1358:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1358
		}
	st_case_1358:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 11:
			goto tr1911
		case 13:
			goto tr1595
		case 32:
			goto tr1910
		case 34:
			goto tr690
		case 44:
			goto tr386
		case 61:
			goto tr691
		case 92:
			goto st433
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1359
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1910
		}
		goto st387
	st1359:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1359
		}
	st_case_1359:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 11:
			goto tr1911
		case 13:
			goto tr1595
		case 32:
			goto tr1910
		case 34:
			goto tr690
		case 44:
			goto tr386
		case 61:
			goto tr691
		case 92:
			goto st433
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1360
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1910
		}
		goto st387
	st1360:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1360
		}
	st_case_1360:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 11:
			goto tr1911
		case 13:
			goto tr1595
		case 32:
			goto tr1910
		case 34:
			goto tr690
		case 44:
			goto tr386
		case 61:
			goto tr691
		case 92:
			goto st433
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1361
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1910
		}
		goto st387
	st1361:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1361
		}
	st_case_1361:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 11:
			goto tr1911
		case 13:
			goto tr1595
		case 32:
			goto tr1910
		case 34:
			goto tr690
		case 44:
			goto tr386
		case 61:
			goto tr691
		case 92:
			goto st433
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1362
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1910
		}
		goto st387
	st1362:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1362
		}
	st_case_1362:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 11:
			goto tr1911
		case 13:
			goto tr1595
		case 32:
			goto tr1910
		case 34:
			goto tr690
		case 44:
			goto tr386
		case 61:
			goto tr691
		case 92:
			goto st433
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1363
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1910
		}
		goto st387
	st1363:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1363
		}
	st_case_1363:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 11:
			goto tr1911
		case 13:
			goto tr1595
		case 32:
			goto tr1910
		case 34:
			goto tr690
		case 44:
			goto tr386
		case 61:
			goto tr691
		case 92:
			goto st433
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1364
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1910
		}
		goto st387
	st1364:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1364
		}
	st_case_1364:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 11:
			goto tr1911
		case 13:
			goto tr1595
		case 32:
			goto tr1910
		case 34:
			goto tr690
		case 44:
			goto tr386
		case 61:
			goto tr691
		case 92:
			goto st433
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1365
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1910
		}
		goto st387
	st1365:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1365
		}
	st_case_1365:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 11:
			goto tr1911
		case 13:
			goto tr1595
		case 32:
			goto tr1910
		case 34:
			goto tr690
		case 44:
			goto tr386
		case 61:
			goto tr691
		case 92:
			goto st433
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1366
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1910
		}
		goto st387
	st1366:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1366
		}
	st_case_1366:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 11:
			goto tr1911
		case 13:
			goto tr1595
		case 32:
			goto tr1910
		case 34:
			goto tr690
		case 44:
			goto tr386
		case 61:
			goto tr691
		case 92:
			goto st433
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1367
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1910
		}
		goto st387
	st1367:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1367
		}
	st_case_1367:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 11:
			goto tr1911
		case 13:
			goto tr1595
		case 32:
			goto tr1910
		case 34:
			goto tr690
		case 44:
			goto tr386
		case 61:
			goto tr691
		case 92:
			goto st433
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1368
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1910
		}
		goto st387
	st1368:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1368
		}
	st_case_1368:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 11:
			goto tr1911
		case 13:
			goto tr1595
		case 32:
			goto tr1910
		case 34:
			goto tr690
		case 44:
			goto tr386
		case 61:
			goto tr691
		case 92:
			goto st433
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1910
		}
		goto st387
tr379:
//line machine.rl:17

	m.pb = m.p

//line machine.rl:86

	m.handler.AddString(key, m.text())

	goto st1369
tr385:
//line machine.rl:86

	m.handler.AddString(key, m.text())

	goto st1369
	st1369:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1369
		}
	st_case_1369:
//line machine.go:49490
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr850
		case 11:
			goto tr1934
		case 13:
			goto tr850
		case 32:
			goto tr889
		case 44:
			goto tr1845
		case 92:
			goto st434
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr889
		}
		goto st1
tr1934:
//line machine.rl:58

	m.handler.SetMeasurement(m.text())

	goto st1370
tr2084:
//line machine.rl:58

	m.handler.SetMeasurement(m.text())

//line machine.rl:78

	m.handler.AddFloat(key, m.text())

	goto st1370
tr2088:
//line machine.rl:58

	m.handler.SetMeasurement(m.text())

//line machine.rl:74

	m.handler.AddInt(key, m.text())

	goto st1370
tr2090:
//line machine.rl:58

	m.handler.SetMeasurement(m.text())

//line machine.rl:82

	m.handler.AddBool(key, m.text())

	goto st1370
	st1370:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1370
		}
	st_case_1370:
//line machine.go:49550
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr850
		case 11:
			goto tr1846
		case 13:
			goto tr850
		case 32:
			goto tr889
		case 44:
			goto tr4
		case 45:
			goto tr1847
		case 61:
			goto st1
		case 92:
			goto tr362
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto tr1848
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr889
		}
		goto tr360
tr840:
//line machine.rl:17

	m.pb = m.p

	goto st434
	st434:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof434
		}
	st_case_434:
//line machine.go:49589
		if ( m.data)[( m.p)] == 92 {
			goto st0
		}
		switch {
		case ( m.data)[( m.p)] > 10:
			if 12 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto st0
			}
		case ( m.data)[( m.p)] >= 9:
			goto st0
		}
		goto st1
tr1881:
//line machine.rl:58

	m.handler.SetMeasurement(m.text())

//line machine.rl:17

	m.pb = m.p

	goto st435
tr2071:
//line machine.rl:58

	m.handler.SetMeasurement(m.text())

//line machine.rl:78

	m.handler.AddFloat(key, m.text())

	goto st435
tr2030:
//line machine.rl:66

	m.handler.AddTag(key, m.text())

//line machine.rl:78

	m.handler.AddFloat(key, m.text())

	goto st435
tr2061:
//line machine.rl:66

	m.handler.AddTag(key, m.text())

//line machine.rl:74

	m.handler.AddInt(key, m.text())

	goto st435
tr2064:
//line machine.rl:66

	m.handler.AddTag(key, m.text())

//line machine.rl:82

	m.handler.AddBool(key, m.text())

	goto st435
tr2076:
//line machine.rl:58

	m.handler.SetMeasurement(m.text())

//line machine.rl:74

	m.handler.AddInt(key, m.text())

	goto st435
tr2079:
//line machine.rl:58

	m.handler.SetMeasurement(m.text())

//line machine.rl:82

	m.handler.AddBool(key, m.text())

	goto st435
	st435:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof435
		}
	st_case_435:
//line machine.go:49677
		switch ( m.data)[( m.p)] {
		case 32:
			goto st6
		case 34:
			goto tr777
		case 44:
			goto st6
		case 61:
			goto st6
		case 92:
			goto tr778
		}
		switch {
		case ( m.data)[( m.p)] > 10:
			if 12 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto st6
			}
		case ( m.data)[( m.p)] >= 9:
			goto st6
		}
		goto tr776
tr776:
//line machine.rl:17

	m.pb = m.p

	goto st436
	st436:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof436
		}
	st_case_436:
//line machine.go:49710
		switch ( m.data)[( m.p)] {
		case 32:
			goto st6
		case 34:
			goto tr780
		case 44:
			goto st6
		case 61:
			goto tr781
		case 92:
			goto st464
		}
		switch {
		case ( m.data)[( m.p)] > 10:
			if 12 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto st6
			}
		case ( m.data)[( m.p)] >= 9:
			goto st6
		}
		goto st436
tr777:
//line machine.rl:17

	m.pb = m.p

//line machine.rl:86

	m.handler.AddString(key, m.text())

	goto st1371
tr780:
//line machine.rl:86

	m.handler.AddString(key, m.text())

	goto st1371
	st1371:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1371
		}
	st_case_1371:
//line machine.go:49753
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr850
		case 11:
			goto st1372
		case 13:
			goto tr850
		case 32:
			goto st492
		case 44:
			goto st28
		case 61:
			goto tr697
		case 92:
			goto st419
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto st492
		}
		goto st392
	st1372:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1372
		}
	st_case_1372:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr850
		case 11:
			goto st1372
		case 13:
			goto tr850
		case 32:
			goto st492
		case 44:
			goto tr179
		case 45:
			goto tr1936
		case 61:
			goto tr697
		case 92:
			goto st419
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto tr1937
			}
		case ( m.data)[( m.p)] >= 9:
			goto st492
		}
		goto st392
tr1936:
//line machine.rl:17

	m.pb = m.p

	goto st437
	st437:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof437
		}
	st_case_437:
//line machine.go:49817
		switch ( m.data)[( m.p)] {
		case 32:
			goto tr179
		case 44:
			goto tr179
		case 61:
			goto tr697
		case 92:
			goto st419
		}
		switch {
		case ( m.data)[( m.p)] < 12:
			if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 10 {
				goto tr179
			}
		case ( m.data)[( m.p)] > 13:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1373
			}
		default:
			goto tr179
		}
		goto st392
tr1937:
//line machine.rl:17

	m.pb = m.p

	goto st1373
	st1373:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1373
		}
	st_case_1373:
//line machine.go:49852
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr1938
		case 13:
			goto tr859
		case 32:
			goto tr858
		case 44:
			goto tr179
		case 61:
			goto tr697
		case 92:
			goto st419
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1375
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr858
		}
		goto st392
tr1938:
//line machine.rl:90

	m.handler.SetTimestamp(m.text())

	goto st1374
	st1374:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1374
		}
	st_case_1374:
//line machine.go:49889
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr850
		case 11:
			goto st1374
		case 13:
			goto tr850
		case 32:
			goto st497
		case 44:
			goto tr48
		case 61:
			goto tr697
		case 92:
			goto st419
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto st497
		}
		goto st392
	st1375:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1375
		}
	st_case_1375:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr1938
		case 13:
			goto tr859
		case 32:
			goto tr858
		case 44:
			goto tr179
		case 61:
			goto tr697
		case 92:
			goto st419
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1376
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr858
		}
		goto st392
	st1376:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1376
		}
	st_case_1376:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr1938
		case 13:
			goto tr859
		case 32:
			goto tr858
		case 44:
			goto tr179
		case 61:
			goto tr697
		case 92:
			goto st419
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1377
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr858
		}
		goto st392
	st1377:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1377
		}
	st_case_1377:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr1938
		case 13:
			goto tr859
		case 32:
			goto tr858
		case 44:
			goto tr179
		case 61:
			goto tr697
		case 92:
			goto st419
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1378
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr858
		}
		goto st392
	st1378:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1378
		}
	st_case_1378:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr1938
		case 13:
			goto tr859
		case 32:
			goto tr858
		case 44:
			goto tr179
		case 61:
			goto tr697
		case 92:
			goto st419
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1379
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr858
		}
		goto st392
	st1379:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1379
		}
	st_case_1379:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr1938
		case 13:
			goto tr859
		case 32:
			goto tr858
		case 44:
			goto tr179
		case 61:
			goto tr697
		case 92:
			goto st419
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1380
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr858
		}
		goto st392
	st1380:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1380
		}
	st_case_1380:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr1938
		case 13:
			goto tr859
		case 32:
			goto tr858
		case 44:
			goto tr179
		case 61:
			goto tr697
		case 92:
			goto st419
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1381
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr858
		}
		goto st392
	st1381:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1381
		}
	st_case_1381:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr1938
		case 13:
			goto tr859
		case 32:
			goto tr858
		case 44:
			goto tr179
		case 61:
			goto tr697
		case 92:
			goto st419
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1382
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr858
		}
		goto st392
	st1382:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1382
		}
	st_case_1382:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr1938
		case 13:
			goto tr859
		case 32:
			goto tr858
		case 44:
			goto tr179
		case 61:
			goto tr697
		case 92:
			goto st419
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1383
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr858
		}
		goto st392
	st1383:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1383
		}
	st_case_1383:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr1938
		case 13:
			goto tr859
		case 32:
			goto tr858
		case 44:
			goto tr179
		case 61:
			goto tr697
		case 92:
			goto st419
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1384
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr858
		}
		goto st392
	st1384:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1384
		}
	st_case_1384:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr1938
		case 13:
			goto tr859
		case 32:
			goto tr858
		case 44:
			goto tr179
		case 61:
			goto tr697
		case 92:
			goto st419
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1385
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr858
		}
		goto st392
	st1385:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1385
		}
	st_case_1385:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr1938
		case 13:
			goto tr859
		case 32:
			goto tr858
		case 44:
			goto tr179
		case 61:
			goto tr697
		case 92:
			goto st419
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1386
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr858
		}
		goto st392
	st1386:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1386
		}
	st_case_1386:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr1938
		case 13:
			goto tr859
		case 32:
			goto tr858
		case 44:
			goto tr179
		case 61:
			goto tr697
		case 92:
			goto st419
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1387
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr858
		}
		goto st392
	st1387:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1387
		}
	st_case_1387:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr1938
		case 13:
			goto tr859
		case 32:
			goto tr858
		case 44:
			goto tr179
		case 61:
			goto tr697
		case 92:
			goto st419
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1388
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr858
		}
		goto st392
	st1388:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1388
		}
	st_case_1388:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr1938
		case 13:
			goto tr859
		case 32:
			goto tr858
		case 44:
			goto tr179
		case 61:
			goto tr697
		case 92:
			goto st419
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1389
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr858
		}
		goto st392
	st1389:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1389
		}
	st_case_1389:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr1938
		case 13:
			goto tr859
		case 32:
			goto tr858
		case 44:
			goto tr179
		case 61:
			goto tr697
		case 92:
			goto st419
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1390
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr858
		}
		goto st392
	st1390:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1390
		}
	st_case_1390:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr1938
		case 13:
			goto tr859
		case 32:
			goto tr858
		case 44:
			goto tr179
		case 61:
			goto tr697
		case 92:
			goto st419
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1391
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr858
		}
		goto st392
	st1391:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1391
		}
	st_case_1391:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr1938
		case 13:
			goto tr859
		case 32:
			goto tr858
		case 44:
			goto tr179
		case 61:
			goto tr697
		case 92:
			goto st419
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1392
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr858
		}
		goto st392
	st1392:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1392
		}
	st_case_1392:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr859
		case 11:
			goto tr1938
		case 13:
			goto tr859
		case 32:
			goto tr858
		case 44:
			goto tr179
		case 61:
			goto tr697
		case 92:
			goto st419
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr858
		}
		goto st392
tr781:
//line machine.rl:62

	key = m.text()

//line machine.rl:70

	key = m.text()

	goto st438
	st438:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof438
		}
	st_case_438:
//line machine.go:50460
		switch ( m.data)[( m.p)] {
		case 32:
			goto st6
		case 34:
			goto tr784
		case 44:
			goto st6
		case 45:
			goto tr785
		case 46:
			goto tr786
		case 48:
			goto tr787
		case 61:
			goto st6
		case 70:
			goto tr789
		case 84:
			goto tr790
		case 92:
			goto tr753
		case 102:
			goto tr791
		case 116:
			goto tr792
		}
		switch {
		case ( m.data)[( m.p)] < 12:
			if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 10 {
				goto st6
			}
		case ( m.data)[( m.p)] > 13:
			if 49 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto tr788
			}
		default:
			goto st6
		}
		goto tr752
tr784:
//line machine.rl:17

	m.pb = m.p

//line machine.rl:86

	m.handler.AddString(key, m.text())

	goto st1393
	st1393:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1393
		}
	st_case_1393:
//line machine.go:50515
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1363
		case 11:
			goto tr1959
		case 13:
			goto tr1363
		case 32:
			goto tr1958
		case 34:
			goto tr711
		case 44:
			goto tr1960
		case 61:
			goto tr22
		case 92:
			goto tr713
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1958
		}
		goto tr708
tr1985:
//line machine.rl:66

	m.handler.AddTag(key, m.text())

	goto st1394
tr1958:
//line machine.rl:66

	m.handler.AddTag(key, m.text())

//line machine.rl:17

	m.pb = m.p

	goto st1394
tr2013:
//line machine.rl:66

	m.handler.AddTag(key, m.text())

//line machine.rl:78

	m.handler.AddFloat(key, m.text())

	goto st1394
tr2018:
//line machine.rl:66

	m.handler.AddTag(key, m.text())

//line machine.rl:74

	m.handler.AddInt(key, m.text())

	goto st1394
tr2021:
//line machine.rl:66

	m.handler.AddTag(key, m.text())

//line machine.rl:82

	m.handler.AddBool(key, m.text())

	goto st1394
	st1394:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1394
		}
	st_case_1394:
//line machine.go:50589
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1366
		case 11:
			goto tr1962
		case 13:
			goto tr1366
		case 32:
			goto st1394
		case 34:
			goto tr82
		case 44:
			goto st6
		case 45:
			goto tr1963
		case 61:
			goto st6
		case 92:
			goto tr723
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto tr1964
			}
		case ( m.data)[( m.p)] >= 9:
			goto st1394
		}
		goto tr720
tr1962:
//line machine.rl:17

	m.pb = m.p

	goto st1395
	st1395:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1395
		}
	st_case_1395:
//line machine.go:50630
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1366
		case 11:
			goto tr1962
		case 13:
			goto tr1366
		case 32:
			goto st1394
		case 34:
			goto tr82
		case 44:
			goto st6
		case 45:
			goto tr1963
		case 61:
			goto tr725
		case 92:
			goto tr723
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto tr1964
			}
		case ( m.data)[( m.p)] >= 9:
			goto st1394
		}
		goto tr720
tr1963:
//line machine.rl:17

	m.pb = m.p

	goto st439
	st439:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof439
		}
	st_case_439:
//line machine.go:50671
		switch ( m.data)[( m.p)] {
		case 32:
			goto st6
		case 34:
			goto tr85
		case 44:
			goto st6
		case 61:
			goto tr725
		case 92:
			goto st399
		}
		switch {
		case ( m.data)[( m.p)] < 12:
			if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 10 {
				goto st6
			}
		case ( m.data)[( m.p)] > 13:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1396
			}
		default:
			goto st6
		}
		goto st397
tr1964:
//line machine.rl:17

	m.pb = m.p

	goto st1396
	st1396:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1396
		}
	st_case_1396:
//line machine.go:50708
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 11:
			goto tr1965
		case 13:
			goto tr1595
		case 32:
			goto tr1594
		case 34:
			goto tr85
		case 44:
			goto st6
		case 61:
			goto tr725
		case 92:
			goto st399
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1398
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1594
		}
		goto st397
tr1965:
//line machine.rl:90

	m.handler.SetTimestamp(m.text())

	goto st1397
	st1397:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1397
		}
	st_case_1397:
//line machine.go:50747
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1366
		case 11:
			goto st1397
		case 13:
			goto tr1366
		case 32:
			goto st1081
		case 34:
			goto tr85
		case 44:
			goto st6
		case 61:
			goto tr725
		case 92:
			goto st399
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto st1081
		}
		goto st397
	st1398:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1398
		}
	st_case_1398:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 11:
			goto tr1965
		case 13:
			goto tr1595
		case 32:
			goto tr1594
		case 34:
			goto tr85
		case 44:
			goto st6
		case 61:
			goto tr725
		case 92:
			goto st399
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1399
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1594
		}
		goto st397
	st1399:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1399
		}
	st_case_1399:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 11:
			goto tr1965
		case 13:
			goto tr1595
		case 32:
			goto tr1594
		case 34:
			goto tr85
		case 44:
			goto st6
		case 61:
			goto tr725
		case 92:
			goto st399
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1400
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1594
		}
		goto st397
	st1400:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1400
		}
	st_case_1400:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 11:
			goto tr1965
		case 13:
			goto tr1595
		case 32:
			goto tr1594
		case 34:
			goto tr85
		case 44:
			goto st6
		case 61:
			goto tr725
		case 92:
			goto st399
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1401
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1594
		}
		goto st397
	st1401:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1401
		}
	st_case_1401:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 11:
			goto tr1965
		case 13:
			goto tr1595
		case 32:
			goto tr1594
		case 34:
			goto tr85
		case 44:
			goto st6
		case 61:
			goto tr725
		case 92:
			goto st399
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1402
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1594
		}
		goto st397
	st1402:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1402
		}
	st_case_1402:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 11:
			goto tr1965
		case 13:
			goto tr1595
		case 32:
			goto tr1594
		case 34:
			goto tr85
		case 44:
			goto st6
		case 61:
			goto tr725
		case 92:
			goto st399
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1403
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1594
		}
		goto st397
	st1403:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1403
		}
	st_case_1403:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 11:
			goto tr1965
		case 13:
			goto tr1595
		case 32:
			goto tr1594
		case 34:
			goto tr85
		case 44:
			goto st6
		case 61:
			goto tr725
		case 92:
			goto st399
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1404
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1594
		}
		goto st397
	st1404:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1404
		}
	st_case_1404:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 11:
			goto tr1965
		case 13:
			goto tr1595
		case 32:
			goto tr1594
		case 34:
			goto tr85
		case 44:
			goto st6
		case 61:
			goto tr725
		case 92:
			goto st399
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1405
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1594
		}
		goto st397
	st1405:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1405
		}
	st_case_1405:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 11:
			goto tr1965
		case 13:
			goto tr1595
		case 32:
			goto tr1594
		case 34:
			goto tr85
		case 44:
			goto st6
		case 61:
			goto tr725
		case 92:
			goto st399
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1406
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1594
		}
		goto st397
	st1406:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1406
		}
	st_case_1406:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 11:
			goto tr1965
		case 13:
			goto tr1595
		case 32:
			goto tr1594
		case 34:
			goto tr85
		case 44:
			goto st6
		case 61:
			goto tr725
		case 92:
			goto st399
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1407
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1594
		}
		goto st397
	st1407:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1407
		}
	st_case_1407:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 11:
			goto tr1965
		case 13:
			goto tr1595
		case 32:
			goto tr1594
		case 34:
			goto tr85
		case 44:
			goto st6
		case 61:
			goto tr725
		case 92:
			goto st399
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1408
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1594
		}
		goto st397
	st1408:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1408
		}
	st_case_1408:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 11:
			goto tr1965
		case 13:
			goto tr1595
		case 32:
			goto tr1594
		case 34:
			goto tr85
		case 44:
			goto st6
		case 61:
			goto tr725
		case 92:
			goto st399
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1409
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1594
		}
		goto st397
	st1409:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1409
		}
	st_case_1409:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 11:
			goto tr1965
		case 13:
			goto tr1595
		case 32:
			goto tr1594
		case 34:
			goto tr85
		case 44:
			goto st6
		case 61:
			goto tr725
		case 92:
			goto st399
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1410
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1594
		}
		goto st397
	st1410:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1410
		}
	st_case_1410:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 11:
			goto tr1965
		case 13:
			goto tr1595
		case 32:
			goto tr1594
		case 34:
			goto tr85
		case 44:
			goto st6
		case 61:
			goto tr725
		case 92:
			goto st399
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1411
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1594
		}
		goto st397
	st1411:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1411
		}
	st_case_1411:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 11:
			goto tr1965
		case 13:
			goto tr1595
		case 32:
			goto tr1594
		case 34:
			goto tr85
		case 44:
			goto st6
		case 61:
			goto tr725
		case 92:
			goto st399
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1412
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1594
		}
		goto st397
	st1412:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1412
		}
	st_case_1412:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 11:
			goto tr1965
		case 13:
			goto tr1595
		case 32:
			goto tr1594
		case 34:
			goto tr85
		case 44:
			goto st6
		case 61:
			goto tr725
		case 92:
			goto st399
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1413
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1594
		}
		goto st397
	st1413:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1413
		}
	st_case_1413:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 11:
			goto tr1965
		case 13:
			goto tr1595
		case 32:
			goto tr1594
		case 34:
			goto tr85
		case 44:
			goto st6
		case 61:
			goto tr725
		case 92:
			goto st399
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1414
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1594
		}
		goto st397
	st1414:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1414
		}
	st_case_1414:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 11:
			goto tr1965
		case 13:
			goto tr1595
		case 32:
			goto tr1594
		case 34:
			goto tr85
		case 44:
			goto st6
		case 61:
			goto tr725
		case 92:
			goto st399
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1415
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1594
		}
		goto st397
	st1415:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1415
		}
	st_case_1415:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 11:
			goto tr1965
		case 13:
			goto tr1595
		case 32:
			goto tr1594
		case 34:
			goto tr85
		case 44:
			goto st6
		case 61:
			goto tr725
		case 92:
			goto st399
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1594
		}
		goto st397
tr1959:
//line machine.rl:66

	m.handler.AddTag(key, m.text())

//line machine.rl:17

	m.pb = m.p

	goto st1416
tr2014:
//line machine.rl:66

	m.handler.AddTag(key, m.text())

//line machine.rl:78

	m.handler.AddFloat(key, m.text())

	goto st1416
tr2019:
//line machine.rl:66

	m.handler.AddTag(key, m.text())

//line machine.rl:74

	m.handler.AddInt(key, m.text())

	goto st1416
tr2022:
//line machine.rl:66

	m.handler.AddTag(key, m.text())

//line machine.rl:82

	m.handler.AddBool(key, m.text())

	goto st1416
	st1416:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1416
		}
	st_case_1416:
//line machine.go:51386
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1366
		case 11:
			goto tr1986
		case 13:
			goto tr1366
		case 32:
			goto tr1985
		case 34:
			goto tr729
		case 44:
			goto tr718
		case 45:
			goto tr1987
		case 61:
			goto st6
		case 92:
			goto tr730
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto tr1988
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1985
		}
		goto tr727
tr1986:
//line machine.rl:66

	m.handler.AddTag(key, m.text())

//line machine.rl:17

	m.pb = m.p

	goto st1417
	st1417:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1417
		}
	st_case_1417:
//line machine.go:51431
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1366
		case 11:
			goto tr1986
		case 13:
			goto tr1366
		case 32:
			goto tr1985
		case 34:
			goto tr729
		case 44:
			goto tr718
		case 45:
			goto tr1987
		case 61:
			goto tr725
		case 92:
			goto tr730
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto tr1988
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1985
		}
		goto tr727
tr1987:
//line machine.rl:17

	m.pb = m.p

	goto st440
	st440:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof440
		}
	st_case_440:
//line machine.go:51472
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr732
		case 13:
			goto st6
		case 32:
			goto tr715
		case 34:
			goto tr733
		case 44:
			goto tr718
		case 61:
			goto tr725
		case 92:
			goto st409
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1418
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr715
		}
		goto st402
tr1988:
//line machine.rl:17

	m.pb = m.p

	goto st1418
	st1418:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1418
		}
	st_case_1418:
//line machine.go:51511
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 11:
			goto tr1990
		case 13:
			goto tr1595
		case 32:
			goto tr1989
		case 34:
			goto tr733
		case 44:
			goto tr718
		case 61:
			goto tr725
		case 92:
			goto st409
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1422
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1989
		}
		goto st402
tr1994:
//line machine.rl:66

	m.handler.AddTag(key, m.text())

	goto st1419
tr1989:
//line machine.rl:66

	m.handler.AddTag(key, m.text())

//line machine.rl:90

	m.handler.SetTimestamp(m.text())

	goto st1419
	st1419:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1419
		}
	st_case_1419:
//line machine.go:51560
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1366
		case 11:
			goto tr1993
		case 13:
			goto tr1366
		case 32:
			goto st1419
		case 34:
			goto tr82
		case 44:
			goto st6
		case 61:
			goto st6
		case 92:
			goto tr723
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto st1419
		}
		goto tr720
tr1993:
//line machine.rl:17

	m.pb = m.p

	goto st1420
	st1420:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1420
		}
	st_case_1420:
//line machine.go:51594
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1366
		case 11:
			goto tr1993
		case 13:
			goto tr1366
		case 32:
			goto st1419
		case 34:
			goto tr82
		case 44:
			goto st6
		case 61:
			goto tr725
		case 92:
			goto tr723
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto st1419
		}
		goto tr720
tr1995:
//line machine.rl:66

	m.handler.AddTag(key, m.text())

//line machine.rl:17

	m.pb = m.p

	goto st1421
tr1990:
//line machine.rl:66

	m.handler.AddTag(key, m.text())

//line machine.rl:90

	m.handler.SetTimestamp(m.text())

	goto st1421
	st1421:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1421
		}
	st_case_1421:
//line machine.go:51642
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1366
		case 11:
			goto tr1995
		case 13:
			goto tr1366
		case 32:
			goto tr1994
		case 34:
			goto tr729
		case 44:
			goto tr718
		case 61:
			goto tr725
		case 92:
			goto tr730
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1994
		}
		goto tr727
	st1422:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1422
		}
	st_case_1422:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 11:
			goto tr1990
		case 13:
			goto tr1595
		case 32:
			goto tr1989
		case 34:
			goto tr733
		case 44:
			goto tr718
		case 61:
			goto tr725
		case 92:
			goto st409
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1423
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1989
		}
		goto st402
	st1423:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1423
		}
	st_case_1423:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 11:
			goto tr1990
		case 13:
			goto tr1595
		case 32:
			goto tr1989
		case 34:
			goto tr733
		case 44:
			goto tr718
		case 61:
			goto tr725
		case 92:
			goto st409
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1424
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1989
		}
		goto st402
	st1424:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1424
		}
	st_case_1424:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 11:
			goto tr1990
		case 13:
			goto tr1595
		case 32:
			goto tr1989
		case 34:
			goto tr733
		case 44:
			goto tr718
		case 61:
			goto tr725
		case 92:
			goto st409
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1425
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1989
		}
		goto st402
	st1425:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1425
		}
	st_case_1425:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 11:
			goto tr1990
		case 13:
			goto tr1595
		case 32:
			goto tr1989
		case 34:
			goto tr733
		case 44:
			goto tr718
		case 61:
			goto tr725
		case 92:
			goto st409
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1426
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1989
		}
		goto st402
	st1426:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1426
		}
	st_case_1426:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 11:
			goto tr1990
		case 13:
			goto tr1595
		case 32:
			goto tr1989
		case 34:
			goto tr733
		case 44:
			goto tr718
		case 61:
			goto tr725
		case 92:
			goto st409
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1427
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1989
		}
		goto st402
	st1427:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1427
		}
	st_case_1427:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 11:
			goto tr1990
		case 13:
			goto tr1595
		case 32:
			goto tr1989
		case 34:
			goto tr733
		case 44:
			goto tr718
		case 61:
			goto tr725
		case 92:
			goto st409
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1428
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1989
		}
		goto st402
	st1428:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1428
		}
	st_case_1428:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 11:
			goto tr1990
		case 13:
			goto tr1595
		case 32:
			goto tr1989
		case 34:
			goto tr733
		case 44:
			goto tr718
		case 61:
			goto tr725
		case 92:
			goto st409
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1429
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1989
		}
		goto st402
	st1429:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1429
		}
	st_case_1429:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 11:
			goto tr1990
		case 13:
			goto tr1595
		case 32:
			goto tr1989
		case 34:
			goto tr733
		case 44:
			goto tr718
		case 61:
			goto tr725
		case 92:
			goto st409
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1430
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1989
		}
		goto st402
	st1430:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1430
		}
	st_case_1430:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 11:
			goto tr1990
		case 13:
			goto tr1595
		case 32:
			goto tr1989
		case 34:
			goto tr733
		case 44:
			goto tr718
		case 61:
			goto tr725
		case 92:
			goto st409
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1431
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1989
		}
		goto st402
	st1431:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1431
		}
	st_case_1431:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 11:
			goto tr1990
		case 13:
			goto tr1595
		case 32:
			goto tr1989
		case 34:
			goto tr733
		case 44:
			goto tr718
		case 61:
			goto tr725
		case 92:
			goto st409
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1432
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1989
		}
		goto st402
	st1432:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1432
		}
	st_case_1432:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 11:
			goto tr1990
		case 13:
			goto tr1595
		case 32:
			goto tr1989
		case 34:
			goto tr733
		case 44:
			goto tr718
		case 61:
			goto tr725
		case 92:
			goto st409
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1433
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1989
		}
		goto st402
	st1433:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1433
		}
	st_case_1433:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 11:
			goto tr1990
		case 13:
			goto tr1595
		case 32:
			goto tr1989
		case 34:
			goto tr733
		case 44:
			goto tr718
		case 61:
			goto tr725
		case 92:
			goto st409
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1434
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1989
		}
		goto st402
	st1434:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1434
		}
	st_case_1434:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 11:
			goto tr1990
		case 13:
			goto tr1595
		case 32:
			goto tr1989
		case 34:
			goto tr733
		case 44:
			goto tr718
		case 61:
			goto tr725
		case 92:
			goto st409
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1435
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1989
		}
		goto st402
	st1435:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1435
		}
	st_case_1435:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 11:
			goto tr1990
		case 13:
			goto tr1595
		case 32:
			goto tr1989
		case 34:
			goto tr733
		case 44:
			goto tr718
		case 61:
			goto tr725
		case 92:
			goto st409
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1436
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1989
		}
		goto st402
	st1436:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1436
		}
	st_case_1436:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 11:
			goto tr1990
		case 13:
			goto tr1595
		case 32:
			goto tr1989
		case 34:
			goto tr733
		case 44:
			goto tr718
		case 61:
			goto tr725
		case 92:
			goto st409
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1437
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1989
		}
		goto st402
	st1437:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1437
		}
	st_case_1437:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 11:
			goto tr1990
		case 13:
			goto tr1595
		case 32:
			goto tr1989
		case 34:
			goto tr733
		case 44:
			goto tr718
		case 61:
			goto tr725
		case 92:
			goto st409
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1438
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1989
		}
		goto st402
	st1438:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1438
		}
	st_case_1438:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 11:
			goto tr1990
		case 13:
			goto tr1595
		case 32:
			goto tr1989
		case 34:
			goto tr733
		case 44:
			goto tr718
		case 61:
			goto tr725
		case 92:
			goto st409
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1439
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1989
		}
		goto st402
	st1439:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1439
		}
	st_case_1439:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 11:
			goto tr1990
		case 13:
			goto tr1595
		case 32:
			goto tr1989
		case 34:
			goto tr733
		case 44:
			goto tr718
		case 61:
			goto tr725
		case 92:
			goto st409
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1989
		}
		goto st402
tr1960:
//line machine.rl:66

	m.handler.AddTag(key, m.text())

//line machine.rl:17

	m.pb = m.p

	goto st441
tr2015:
//line machine.rl:66

	m.handler.AddTag(key, m.text())

//line machine.rl:78

	m.handler.AddFloat(key, m.text())

	goto st441
tr2020:
//line machine.rl:66

	m.handler.AddTag(key, m.text())

//line machine.rl:74

	m.handler.AddInt(key, m.text())

	goto st441
tr2023:
//line machine.rl:66

	m.handler.AddTag(key, m.text())

//line machine.rl:82

	m.handler.AddBool(key, m.text())

	goto st441
	st441:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof441
		}
	st_case_441:
//line machine.go:52281
		switch ( m.data)[( m.p)] {
		case 32:
			goto st6
		case 34:
			goto tr777
		case 44:
			goto st6
		case 61:
			goto st6
		case 92:
			goto tr796
		}
		switch {
		case ( m.data)[( m.p)] > 10:
			if 12 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto st6
			}
		case ( m.data)[( m.p)] >= 9:
			goto st6
		}
		goto tr795
tr795:
//line machine.rl:17

	m.pb = m.p

	goto st442
	st442:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof442
		}
	st_case_442:
//line machine.go:52314
		switch ( m.data)[( m.p)] {
		case 32:
			goto st6
		case 34:
			goto tr780
		case 44:
			goto st6
		case 61:
			goto tr798
		case 92:
			goto st453
		}
		switch {
		case ( m.data)[( m.p)] > 10:
			if 12 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto st6
			}
		case ( m.data)[( m.p)] >= 9:
			goto st6
		}
		goto st442
tr798:
//line machine.rl:62

	key = m.text()

//line machine.rl:70

	key = m.text()

	goto st443
	st443:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof443
		}
	st_case_443:
//line machine.go:52351
		switch ( m.data)[( m.p)] {
		case 32:
			goto st6
		case 34:
			goto tr784
		case 44:
			goto st6
		case 45:
			goto tr800
		case 46:
			goto tr801
		case 48:
			goto tr802
		case 61:
			goto st6
		case 70:
			goto tr804
		case 84:
			goto tr805
		case 92:
			goto tr713
		case 102:
			goto tr806
		case 116:
			goto tr807
		}
		switch {
		case ( m.data)[( m.p)] < 12:
			if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 10 {
				goto st6
			}
		case ( m.data)[( m.p)] > 13:
			if 49 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto tr803
			}
		default:
			goto st6
		}
		goto tr708
tr800:
//line machine.rl:17

	m.pb = m.p

	goto st444
	st444:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof444
		}
	st_case_444:
//line machine.go:52402
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr716
		case 13:
			goto st6
		case 32:
			goto tr715
		case 34:
			goto tr717
		case 44:
			goto tr718
		case 48:
			goto st1440
		case 61:
			goto st6
		case 92:
			goto st407
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 49 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1443
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr715
		}
		goto st395
tr802:
//line machine.rl:17

	m.pb = m.p

	goto st1440
	st1440:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1440
		}
	st_case_1440:
//line machine.go:52443
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1415
		case 11:
			goto tr2014
		case 13:
			goto tr1415
		case 32:
			goto tr2013
		case 34:
			goto tr717
		case 44:
			goto tr2015
		case 46:
			goto st1441
		case 61:
			goto st6
		case 92:
			goto st407
		case 105:
			goto st1442
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr2013
		}
		goto st395
tr801:
//line machine.rl:17

	m.pb = m.p

	goto st1441
	st1441:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1441
		}
	st_case_1441:
//line machine.go:52481
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1415
		case 11:
			goto tr2014
		case 13:
			goto tr1415
		case 32:
			goto tr2013
		case 34:
			goto tr717
		case 44:
			goto tr2015
		case 61:
			goto st6
		case 92:
			goto st407
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1441
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr2013
		}
		goto st395
	st1442:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1442
		}
	st_case_1442:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1420
		case 11:
			goto tr2019
		case 13:
			goto tr1420
		case 32:
			goto tr2018
		case 34:
			goto tr717
		case 44:
			goto tr2020
		case 61:
			goto st6
		case 92:
			goto st407
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr2018
		}
		goto st395
tr803:
//line machine.rl:17

	m.pb = m.p

	goto st1443
	st1443:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1443
		}
	st_case_1443:
//line machine.go:52547
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1415
		case 11:
			goto tr2014
		case 13:
			goto tr1415
		case 32:
			goto tr2013
		case 34:
			goto tr717
		case 44:
			goto tr2015
		case 46:
			goto st1441
		case 61:
			goto st6
		case 92:
			goto st407
		case 105:
			goto st1442
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1443
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr2013
		}
		goto st395
tr804:
//line machine.rl:17

	m.pb = m.p

	goto st1444
	st1444:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1444
		}
	st_case_1444:
//line machine.go:52590
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1423
		case 11:
			goto tr2022
		case 13:
			goto tr1423
		case 32:
			goto tr2021
		case 34:
			goto tr717
		case 44:
			goto tr2023
		case 61:
			goto st6
		case 65:
			goto st445
		case 92:
			goto st407
		case 97:
			goto st448
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr2021
		}
		goto st395
	st445:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof445
		}
	st_case_445:
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr716
		case 13:
			goto st6
		case 32:
			goto tr715
		case 34:
			goto tr717
		case 44:
			goto tr718
		case 61:
			goto st6
		case 76:
			goto st446
		case 92:
			goto st407
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr715
		}
		goto st395
	st446:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof446
		}
	st_case_446:
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr716
		case 13:
			goto st6
		case 32:
			goto tr715
		case 34:
			goto tr717
		case 44:
			goto tr718
		case 61:
			goto st6
		case 83:
			goto st447
		case 92:
			goto st407
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr715
		}
		goto st395
	st447:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof447
		}
	st_case_447:
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr716
		case 13:
			goto st6
		case 32:
			goto tr715
		case 34:
			goto tr717
		case 44:
			goto tr718
		case 61:
			goto st6
		case 69:
			goto st1445
		case 92:
			goto st407
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr715
		}
		goto st395
	st1445:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1445
		}
	st_case_1445:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1423
		case 11:
			goto tr2022
		case 13:
			goto tr1423
		case 32:
			goto tr2021
		case 34:
			goto tr717
		case 44:
			goto tr2023
		case 61:
			goto st6
		case 92:
			goto st407
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr2021
		}
		goto st395
	st448:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof448
		}
	st_case_448:
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr716
		case 13:
			goto st6
		case 32:
			goto tr715
		case 34:
			goto tr717
		case 44:
			goto tr718
		case 61:
			goto st6
		case 92:
			goto st407
		case 108:
			goto st449
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr715
		}
		goto st395
	st449:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof449
		}
	st_case_449:
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr716
		case 13:
			goto st6
		case 32:
			goto tr715
		case 34:
			goto tr717
		case 44:
			goto tr718
		case 61:
			goto st6
		case 92:
			goto st407
		case 115:
			goto st450
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr715
		}
		goto st395
	st450:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof450
		}
	st_case_450:
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr716
		case 13:
			goto st6
		case 32:
			goto tr715
		case 34:
			goto tr717
		case 44:
			goto tr718
		case 61:
			goto st6
		case 92:
			goto st407
		case 101:
			goto st1445
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr715
		}
		goto st395
tr805:
//line machine.rl:17

	m.pb = m.p

	goto st1446
	st1446:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1446
		}
	st_case_1446:
//line machine.go:52829
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1423
		case 11:
			goto tr2022
		case 13:
			goto tr1423
		case 32:
			goto tr2021
		case 34:
			goto tr717
		case 44:
			goto tr2023
		case 61:
			goto st6
		case 82:
			goto st451
		case 92:
			goto st407
		case 114:
			goto st452
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr2021
		}
		goto st395
	st451:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof451
		}
	st_case_451:
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr716
		case 13:
			goto st6
		case 32:
			goto tr715
		case 34:
			goto tr717
		case 44:
			goto tr718
		case 61:
			goto st6
		case 85:
			goto st447
		case 92:
			goto st407
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr715
		}
		goto st395
	st452:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof452
		}
	st_case_452:
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr716
		case 13:
			goto st6
		case 32:
			goto tr715
		case 34:
			goto tr717
		case 44:
			goto tr718
		case 61:
			goto st6
		case 92:
			goto st407
		case 117:
			goto st450
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr715
		}
		goto st395
tr806:
//line machine.rl:17

	m.pb = m.p

	goto st1447
	st1447:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1447
		}
	st_case_1447:
//line machine.go:52925
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1423
		case 11:
			goto tr2022
		case 13:
			goto tr1423
		case 32:
			goto tr2021
		case 34:
			goto tr717
		case 44:
			goto tr2023
		case 61:
			goto st6
		case 92:
			goto st407
		case 97:
			goto st448
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr2021
		}
		goto st395
tr807:
//line machine.rl:17

	m.pb = m.p

	goto st1448
	st1448:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1448
		}
	st_case_1448:
//line machine.go:52961
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1423
		case 11:
			goto tr2022
		case 13:
			goto tr1423
		case 32:
			goto tr2021
		case 34:
			goto tr717
		case 44:
			goto tr2023
		case 61:
			goto st6
		case 92:
			goto st407
		case 114:
			goto st452
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr2021
		}
		goto st395
tr796:
//line machine.rl:17

	m.pb = m.p

	goto st453
	st453:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof453
		}
	st_case_453:
//line machine.go:52997
		switch ( m.data)[( m.p)] {
		case 34:
			goto st442
		case 92:
			goto st442
		}
		switch {
		case ( m.data)[( m.p)] > 10:
			if 12 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto tr48
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr48
		}
		goto st392
tr785:
//line machine.rl:17

	m.pb = m.p

	goto st454
	st454:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof454
		}
	st_case_454:
//line machine.go:53024
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr756
		case 13:
			goto st6
		case 32:
			goto tr755
		case 34:
			goto tr717
		case 44:
			goto tr757
		case 48:
			goto st1449
		case 61:
			goto st6
		case 92:
			goto st428
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 49 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1474
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr755
		}
		goto st423
tr787:
//line machine.rl:17

	m.pb = m.p

	goto st1449
	st1449:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1449
		}
	st_case_1449:
//line machine.go:53065
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1415
		case 11:
			goto tr2029
		case 13:
			goto tr1415
		case 32:
			goto tr2028
		case 34:
			goto tr717
		case 44:
			goto tr2030
		case 46:
			goto st1472
		case 61:
			goto st6
		case 92:
			goto st428
		case 105:
			goto st1473
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr2028
		}
		goto st423
tr2029:
//line machine.rl:66

	m.handler.AddTag(key, m.text())

//line machine.rl:78

	m.handler.AddFloat(key, m.text())

	goto st1450
tr2060:
//line machine.rl:66

	m.handler.AddTag(key, m.text())

//line machine.rl:74

	m.handler.AddInt(key, m.text())

	goto st1450
tr2063:
//line machine.rl:66

	m.handler.AddTag(key, m.text())

//line machine.rl:82

	m.handler.AddBool(key, m.text())

	goto st1450
	st1450:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1450
		}
	st_case_1450:
//line machine.go:53127
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1366
		case 11:
			goto tr2034
		case 13:
			goto tr1366
		case 32:
			goto tr2033
		case 34:
			goto tr729
		case 44:
			goto tr757
		case 45:
			goto tr2035
		case 61:
			goto st6
		case 92:
			goto tr761
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto tr2036
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr2033
		}
		goto tr759
tr2034:
//line machine.rl:66

	m.handler.AddTag(key, m.text())

//line machine.rl:17

	m.pb = m.p

	goto st1451
	st1451:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1451
		}
	st_case_1451:
//line machine.go:53172
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1366
		case 11:
			goto tr2034
		case 13:
			goto tr1366
		case 32:
			goto tr2033
		case 34:
			goto tr729
		case 44:
			goto tr757
		case 45:
			goto tr2035
		case 61:
			goto tr393
		case 92:
			goto tr761
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto tr2036
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr2033
		}
		goto tr759
tr2035:
//line machine.rl:17

	m.pb = m.p

	goto st455
	st455:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof455
		}
	st_case_455:
//line machine.go:53213
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr763
		case 13:
			goto st6
		case 32:
			goto tr755
		case 34:
			goto tr733
		case 44:
			goto tr757
		case 61:
			goto tr393
		case 92:
			goto st427
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1452
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr755
		}
		goto st425
tr2036:
//line machine.rl:17

	m.pb = m.p

	goto st1452
	st1452:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1452
		}
	st_case_1452:
//line machine.go:53252
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 11:
			goto tr2038
		case 13:
			goto tr1595
		case 32:
			goto tr2037
		case 34:
			goto tr733
		case 44:
			goto tr757
		case 61:
			goto tr393
		case 92:
			goto st427
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1454
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr2037
		}
		goto st425
tr2041:
//line machine.rl:66

	m.handler.AddTag(key, m.text())

//line machine.rl:17

	m.pb = m.p

	goto st1453
tr2038:
//line machine.rl:66

	m.handler.AddTag(key, m.text())

//line machine.rl:90

	m.handler.SetTimestamp(m.text())

	goto st1453
	st1453:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1453
		}
	st_case_1453:
//line machine.go:53305
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1366
		case 11:
			goto tr2041
		case 13:
			goto tr1366
		case 32:
			goto tr2040
		case 34:
			goto tr729
		case 44:
			goto tr757
		case 61:
			goto tr393
		case 92:
			goto tr761
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr2040
		}
		goto tr759
	st1454:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1454
		}
	st_case_1454:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 11:
			goto tr2038
		case 13:
			goto tr1595
		case 32:
			goto tr2037
		case 34:
			goto tr733
		case 44:
			goto tr757
		case 61:
			goto tr393
		case 92:
			goto st427
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1455
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr2037
		}
		goto st425
	st1455:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1455
		}
	st_case_1455:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 11:
			goto tr2038
		case 13:
			goto tr1595
		case 32:
			goto tr2037
		case 34:
			goto tr733
		case 44:
			goto tr757
		case 61:
			goto tr393
		case 92:
			goto st427
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1456
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr2037
		}
		goto st425
	st1456:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1456
		}
	st_case_1456:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 11:
			goto tr2038
		case 13:
			goto tr1595
		case 32:
			goto tr2037
		case 34:
			goto tr733
		case 44:
			goto tr757
		case 61:
			goto tr393
		case 92:
			goto st427
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1457
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr2037
		}
		goto st425
	st1457:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1457
		}
	st_case_1457:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 11:
			goto tr2038
		case 13:
			goto tr1595
		case 32:
			goto tr2037
		case 34:
			goto tr733
		case 44:
			goto tr757
		case 61:
			goto tr393
		case 92:
			goto st427
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1458
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr2037
		}
		goto st425
	st1458:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1458
		}
	st_case_1458:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 11:
			goto tr2038
		case 13:
			goto tr1595
		case 32:
			goto tr2037
		case 34:
			goto tr733
		case 44:
			goto tr757
		case 61:
			goto tr393
		case 92:
			goto st427
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1459
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr2037
		}
		goto st425
	st1459:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1459
		}
	st_case_1459:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 11:
			goto tr2038
		case 13:
			goto tr1595
		case 32:
			goto tr2037
		case 34:
			goto tr733
		case 44:
			goto tr757
		case 61:
			goto tr393
		case 92:
			goto st427
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1460
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr2037
		}
		goto st425
	st1460:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1460
		}
	st_case_1460:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 11:
			goto tr2038
		case 13:
			goto tr1595
		case 32:
			goto tr2037
		case 34:
			goto tr733
		case 44:
			goto tr757
		case 61:
			goto tr393
		case 92:
			goto st427
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1461
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr2037
		}
		goto st425
	st1461:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1461
		}
	st_case_1461:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 11:
			goto tr2038
		case 13:
			goto tr1595
		case 32:
			goto tr2037
		case 34:
			goto tr733
		case 44:
			goto tr757
		case 61:
			goto tr393
		case 92:
			goto st427
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1462
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr2037
		}
		goto st425
	st1462:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1462
		}
	st_case_1462:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 11:
			goto tr2038
		case 13:
			goto tr1595
		case 32:
			goto tr2037
		case 34:
			goto tr733
		case 44:
			goto tr757
		case 61:
			goto tr393
		case 92:
			goto st427
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1463
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr2037
		}
		goto st425
	st1463:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1463
		}
	st_case_1463:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 11:
			goto tr2038
		case 13:
			goto tr1595
		case 32:
			goto tr2037
		case 34:
			goto tr733
		case 44:
			goto tr757
		case 61:
			goto tr393
		case 92:
			goto st427
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1464
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr2037
		}
		goto st425
	st1464:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1464
		}
	st_case_1464:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 11:
			goto tr2038
		case 13:
			goto tr1595
		case 32:
			goto tr2037
		case 34:
			goto tr733
		case 44:
			goto tr757
		case 61:
			goto tr393
		case 92:
			goto st427
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1465
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr2037
		}
		goto st425
	st1465:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1465
		}
	st_case_1465:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 11:
			goto tr2038
		case 13:
			goto tr1595
		case 32:
			goto tr2037
		case 34:
			goto tr733
		case 44:
			goto tr757
		case 61:
			goto tr393
		case 92:
			goto st427
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1466
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr2037
		}
		goto st425
	st1466:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1466
		}
	st_case_1466:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 11:
			goto tr2038
		case 13:
			goto tr1595
		case 32:
			goto tr2037
		case 34:
			goto tr733
		case 44:
			goto tr757
		case 61:
			goto tr393
		case 92:
			goto st427
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1467
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr2037
		}
		goto st425
	st1467:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1467
		}
	st_case_1467:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 11:
			goto tr2038
		case 13:
			goto tr1595
		case 32:
			goto tr2037
		case 34:
			goto tr733
		case 44:
			goto tr757
		case 61:
			goto tr393
		case 92:
			goto st427
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1468
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr2037
		}
		goto st425
	st1468:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1468
		}
	st_case_1468:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 11:
			goto tr2038
		case 13:
			goto tr1595
		case 32:
			goto tr2037
		case 34:
			goto tr733
		case 44:
			goto tr757
		case 61:
			goto tr393
		case 92:
			goto st427
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1469
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr2037
		}
		goto st425
	st1469:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1469
		}
	st_case_1469:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 11:
			goto tr2038
		case 13:
			goto tr1595
		case 32:
			goto tr2037
		case 34:
			goto tr733
		case 44:
			goto tr757
		case 61:
			goto tr393
		case 92:
			goto st427
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1470
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr2037
		}
		goto st425
	st1470:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1470
		}
	st_case_1470:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 11:
			goto tr2038
		case 13:
			goto tr1595
		case 32:
			goto tr2037
		case 34:
			goto tr733
		case 44:
			goto tr757
		case 61:
			goto tr393
		case 92:
			goto st427
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1471
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr2037
		}
		goto st425
	st1471:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1471
		}
	st_case_1471:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1595
		case 11:
			goto tr2038
		case 13:
			goto tr1595
		case 32:
			goto tr2037
		case 34:
			goto tr733
		case 44:
			goto tr757
		case 61:
			goto tr393
		case 92:
			goto st427
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr2037
		}
		goto st425
tr786:
//line machine.rl:17

	m.pb = m.p

	goto st1472
	st1472:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1472
		}
	st_case_1472:
//line machine.go:53910
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1415
		case 11:
			goto tr2029
		case 13:
			goto tr1415
		case 32:
			goto tr2028
		case 34:
			goto tr717
		case 44:
			goto tr2030
		case 61:
			goto st6
		case 92:
			goto st428
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1472
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr2028
		}
		goto st423
	st1473:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1473
		}
	st_case_1473:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1420
		case 11:
			goto tr2060
		case 13:
			goto tr1420
		case 32:
			goto tr2059
		case 34:
			goto tr717
		case 44:
			goto tr2061
		case 61:
			goto st6
		case 92:
			goto st428
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr2059
		}
		goto st423
tr788:
//line machine.rl:17

	m.pb = m.p

	goto st1474
	st1474:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1474
		}
	st_case_1474:
//line machine.go:53976
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1415
		case 11:
			goto tr2029
		case 13:
			goto tr1415
		case 32:
			goto tr2028
		case 34:
			goto tr717
		case 44:
			goto tr2030
		case 46:
			goto st1472
		case 61:
			goto st6
		case 92:
			goto st428
		case 105:
			goto st1473
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1474
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr2028
		}
		goto st423
tr789:
//line machine.rl:17

	m.pb = m.p

	goto st1475
	st1475:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1475
		}
	st_case_1475:
//line machine.go:54019
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1423
		case 11:
			goto tr2063
		case 13:
			goto tr1423
		case 32:
			goto tr2062
		case 34:
			goto tr717
		case 44:
			goto tr2064
		case 61:
			goto st6
		case 65:
			goto st456
		case 92:
			goto st428
		case 97:
			goto st459
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr2062
		}
		goto st423
	st456:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof456
		}
	st_case_456:
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr756
		case 13:
			goto st6
		case 32:
			goto tr755
		case 34:
			goto tr717
		case 44:
			goto tr757
		case 61:
			goto st6
		case 76:
			goto st457
		case 92:
			goto st428
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr755
		}
		goto st423
	st457:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof457
		}
	st_case_457:
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr756
		case 13:
			goto st6
		case 32:
			goto tr755
		case 34:
			goto tr717
		case 44:
			goto tr757
		case 61:
			goto st6
		case 83:
			goto st458
		case 92:
			goto st428
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr755
		}
		goto st423
	st458:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof458
		}
	st_case_458:
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr756
		case 13:
			goto st6
		case 32:
			goto tr755
		case 34:
			goto tr717
		case 44:
			goto tr757
		case 61:
			goto st6
		case 69:
			goto st1476
		case 92:
			goto st428
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr755
		}
		goto st423
	st1476:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1476
		}
	st_case_1476:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1423
		case 11:
			goto tr2063
		case 13:
			goto tr1423
		case 32:
			goto tr2062
		case 34:
			goto tr717
		case 44:
			goto tr2064
		case 61:
			goto st6
		case 92:
			goto st428
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr2062
		}
		goto st423
	st459:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof459
		}
	st_case_459:
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr756
		case 13:
			goto st6
		case 32:
			goto tr755
		case 34:
			goto tr717
		case 44:
			goto tr757
		case 61:
			goto st6
		case 92:
			goto st428
		case 108:
			goto st460
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr755
		}
		goto st423
	st460:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof460
		}
	st_case_460:
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr756
		case 13:
			goto st6
		case 32:
			goto tr755
		case 34:
			goto tr717
		case 44:
			goto tr757
		case 61:
			goto st6
		case 92:
			goto st428
		case 115:
			goto st461
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr755
		}
		goto st423
	st461:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof461
		}
	st_case_461:
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr756
		case 13:
			goto st6
		case 32:
			goto tr755
		case 34:
			goto tr717
		case 44:
			goto tr757
		case 61:
			goto st6
		case 92:
			goto st428
		case 101:
			goto st1476
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr755
		}
		goto st423
tr790:
//line machine.rl:17

	m.pb = m.p

	goto st1477
	st1477:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1477
		}
	st_case_1477:
//line machine.go:54258
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1423
		case 11:
			goto tr2063
		case 13:
			goto tr1423
		case 32:
			goto tr2062
		case 34:
			goto tr717
		case 44:
			goto tr2064
		case 61:
			goto st6
		case 82:
			goto st462
		case 92:
			goto st428
		case 114:
			goto st463
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr2062
		}
		goto st423
	st462:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof462
		}
	st_case_462:
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr756
		case 13:
			goto st6
		case 32:
			goto tr755
		case 34:
			goto tr717
		case 44:
			goto tr757
		case 61:
			goto st6
		case 85:
			goto st458
		case 92:
			goto st428
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr755
		}
		goto st423
	st463:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof463
		}
	st_case_463:
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr756
		case 13:
			goto st6
		case 32:
			goto tr755
		case 34:
			goto tr717
		case 44:
			goto tr757
		case 61:
			goto st6
		case 92:
			goto st428
		case 117:
			goto st461
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr755
		}
		goto st423
tr791:
//line machine.rl:17

	m.pb = m.p

	goto st1478
	st1478:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1478
		}
	st_case_1478:
//line machine.go:54354
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1423
		case 11:
			goto tr2063
		case 13:
			goto tr1423
		case 32:
			goto tr2062
		case 34:
			goto tr717
		case 44:
			goto tr2064
		case 61:
			goto st6
		case 92:
			goto st428
		case 97:
			goto st459
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr2062
		}
		goto st423
tr792:
//line machine.rl:17

	m.pb = m.p

	goto st1479
	st1479:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1479
		}
	st_case_1479:
//line machine.go:54390
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1423
		case 11:
			goto tr2063
		case 13:
			goto tr1423
		case 32:
			goto tr2062
		case 34:
			goto tr717
		case 44:
			goto tr2064
		case 61:
			goto st6
		case 92:
			goto st428
		case 114:
			goto st463
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr2062
		}
		goto st423
tr778:
//line machine.rl:17

	m.pb = m.p

	goto st464
	st464:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof464
		}
	st_case_464:
//line machine.go:54426
		switch ( m.data)[( m.p)] {
		case 34:
			goto st436
		case 92:
			goto st436
		}
		switch {
		case ( m.data)[( m.p)] > 10:
			if 12 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto tr48
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr48
		}
		goto st392
tr381:
//line machine.rl:17

	m.pb = m.p

	goto st465
	st465:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof465
		}
	st_case_465:
//line machine.go:54453
		switch ( m.data)[( m.p)] {
		case 34:
			goto st209
		case 92:
			goto st6
		}
		switch {
		case ( m.data)[( m.p)] > 10:
			if 12 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 13 {
				goto tr2
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr2
		}
		goto st1
tr766:
//line machine.rl:17

	m.pb = m.p

	goto st466
	st466:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof466
		}
	st_case_466:
//line machine.go:54480
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr384
		case 13:
			goto st6
		case 32:
			goto tr383
		case 34:
			goto tr385
		case 44:
			goto tr386
		case 48:
			goto st1480
		case 92:
			goto st465
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 49 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1483
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr383
		}
		goto st209
tr768:
//line machine.rl:17

	m.pb = m.p

	goto st1480
	st1480:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1480
		}
	st_case_1480:
//line machine.go:54519
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1415
		case 11:
			goto tr2070
		case 13:
			goto tr1415
		case 32:
			goto tr2069
		case 34:
			goto tr385
		case 44:
			goto tr2071
		case 46:
			goto st1481
		case 92:
			goto st465
		case 105:
			goto st1482
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr2069
		}
		goto st209
tr767:
//line machine.rl:17

	m.pb = m.p

	goto st1481
	st1481:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1481
		}
	st_case_1481:
//line machine.go:54555
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1415
		case 11:
			goto tr2070
		case 13:
			goto tr1415
		case 32:
			goto tr2069
		case 34:
			goto tr385
		case 44:
			goto tr2071
		case 92:
			goto st465
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1481
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr2069
		}
		goto st209
	st1482:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1482
		}
	st_case_1482:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1420
		case 11:
			goto tr2075
		case 13:
			goto tr1420
		case 32:
			goto tr2074
		case 34:
			goto tr385
		case 44:
			goto tr2076
		case 92:
			goto st465
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr2074
		}
		goto st209
tr769:
//line machine.rl:17

	m.pb = m.p

	goto st1483
	st1483:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1483
		}
	st_case_1483:
//line machine.go:54617
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1415
		case 11:
			goto tr2070
		case 13:
			goto tr1415
		case 32:
			goto tr2069
		case 34:
			goto tr385
		case 44:
			goto tr2071
		case 46:
			goto st1481
		case 92:
			goto st465
		case 105:
			goto st1482
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1483
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr2069
		}
		goto st209
tr770:
//line machine.rl:17

	m.pb = m.p

	goto st1484
	st1484:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1484
		}
	st_case_1484:
//line machine.go:54658
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1423
		case 11:
			goto tr2078
		case 13:
			goto tr1423
		case 32:
			goto tr2077
		case 34:
			goto tr385
		case 44:
			goto tr2079
		case 65:
			goto st467
		case 92:
			goto st465
		case 97:
			goto st470
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr2077
		}
		goto st209
	st467:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof467
		}
	st_case_467:
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr384
		case 13:
			goto st6
		case 32:
			goto tr383
		case 34:
			goto tr385
		case 44:
			goto tr386
		case 76:
			goto st468
		case 92:
			goto st465
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr383
		}
		goto st209
	st468:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof468
		}
	st_case_468:
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr384
		case 13:
			goto st6
		case 32:
			goto tr383
		case 34:
			goto tr385
		case 44:
			goto tr386
		case 83:
			goto st469
		case 92:
			goto st465
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr383
		}
		goto st209
	st469:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof469
		}
	st_case_469:
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr384
		case 13:
			goto st6
		case 32:
			goto tr383
		case 34:
			goto tr385
		case 44:
			goto tr386
		case 69:
			goto st1485
		case 92:
			goto st465
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr383
		}
		goto st209
	st1485:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1485
		}
	st_case_1485:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1423
		case 11:
			goto tr2078
		case 13:
			goto tr1423
		case 32:
			goto tr2077
		case 34:
			goto tr385
		case 44:
			goto tr2079
		case 92:
			goto st465
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr2077
		}
		goto st209
	st470:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof470
		}
	st_case_470:
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr384
		case 13:
			goto st6
		case 32:
			goto tr383
		case 34:
			goto tr385
		case 44:
			goto tr386
		case 92:
			goto st465
		case 108:
			goto st471
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr383
		}
		goto st209
	st471:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof471
		}
	st_case_471:
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr384
		case 13:
			goto st6
		case 32:
			goto tr383
		case 34:
			goto tr385
		case 44:
			goto tr386
		case 92:
			goto st465
		case 115:
			goto st472
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr383
		}
		goto st209
	st472:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof472
		}
	st_case_472:
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr384
		case 13:
			goto st6
		case 32:
			goto tr383
		case 34:
			goto tr385
		case 44:
			goto tr386
		case 92:
			goto st465
		case 101:
			goto st1485
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr383
		}
		goto st209
tr771:
//line machine.rl:17

	m.pb = m.p

	goto st1486
	st1486:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1486
		}
	st_case_1486:
//line machine.go:54881
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1423
		case 11:
			goto tr2078
		case 13:
			goto tr1423
		case 32:
			goto tr2077
		case 34:
			goto tr385
		case 44:
			goto tr2079
		case 82:
			goto st473
		case 92:
			goto st465
		case 114:
			goto st474
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr2077
		}
		goto st209
	st473:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof473
		}
	st_case_473:
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr384
		case 13:
			goto st6
		case 32:
			goto tr383
		case 34:
			goto tr385
		case 44:
			goto tr386
		case 85:
			goto st469
		case 92:
			goto st465
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr383
		}
		goto st209
	st474:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof474
		}
	st_case_474:
		switch ( m.data)[( m.p)] {
		case 10:
			goto st6
		case 11:
			goto tr384
		case 13:
			goto st6
		case 32:
			goto tr383
		case 34:
			goto tr385
		case 44:
			goto tr386
		case 92:
			goto st465
		case 117:
			goto st472
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr383
		}
		goto st209
tr772:
//line machine.rl:17

	m.pb = m.p

	goto st1487
	st1487:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1487
		}
	st_case_1487:
//line machine.go:54971
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1423
		case 11:
			goto tr2078
		case 13:
			goto tr1423
		case 32:
			goto tr2077
		case 34:
			goto tr385
		case 44:
			goto tr2079
		case 92:
			goto st465
		case 97:
			goto st470
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr2077
		}
		goto st209
tr773:
//line machine.rl:17

	m.pb = m.p

	goto st1488
	st1488:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1488
		}
	st_case_1488:
//line machine.go:55005
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1423
		case 11:
			goto tr2078
		case 13:
			goto tr1423
		case 32:
			goto tr2077
		case 34:
			goto tr385
		case 44:
			goto tr2079
		case 92:
			goto st465
		case 114:
			goto st474
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr2077
		}
		goto st209
tr368:
//line machine.rl:17

	m.pb = m.p

	goto st475
	st475:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof475
		}
	st_case_475:
//line machine.go:55039
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr2
		case 11:
			goto tr3
		case 13:
			goto tr2
		case 32:
			goto tr1
		case 44:
			goto tr4
		case 48:
			goto st1489
		case 92:
			goto st434
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 49 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1492
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1
		}
		goto st1
tr370:
//line machine.rl:17

	m.pb = m.p

	goto st1489
	st1489:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1489
		}
	st_case_1489:
//line machine.go:55076
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1270
		case 11:
			goto tr2084
		case 13:
			goto tr1270
		case 32:
			goto tr1317
		case 44:
			goto tr2085
		case 46:
			goto st1490
		case 92:
			goto st434
		case 105:
			goto st1491
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1317
		}
		goto st1
tr369:
//line machine.rl:17

	m.pb = m.p

	goto st1490
	st1490:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1490
		}
	st_case_1490:
//line machine.go:55110
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1270
		case 11:
			goto tr2084
		case 13:
			goto tr1270
		case 32:
			goto tr1317
		case 44:
			goto tr2085
		case 92:
			goto st434
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1490
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1317
		}
		goto st1
	st1491:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1491
		}
	st_case_1491:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1276
		case 11:
			goto tr2088
		case 13:
			goto tr1276
		case 32:
			goto tr1322
		case 44:
			goto tr2089
		case 92:
			goto st434
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1322
		}
		goto st1
tr371:
//line machine.rl:17

	m.pb = m.p

	goto st1492
	st1492:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1492
		}
	st_case_1492:
//line machine.go:55168
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1270
		case 11:
			goto tr2084
		case 13:
			goto tr1270
		case 32:
			goto tr1317
		case 44:
			goto tr2085
		case 46:
			goto st1490
		case 92:
			goto st434
		case 105:
			goto st1491
		}
		switch {
		case ( m.data)[( m.p)] > 12:
			if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
				goto st1492
			}
		case ( m.data)[( m.p)] >= 9:
			goto tr1317
		}
		goto st1
tr372:
//line machine.rl:17

	m.pb = m.p

	goto st1493
	st1493:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1493
		}
	st_case_1493:
//line machine.go:55207
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1280
		case 11:
			goto tr2090
		case 13:
			goto tr1280
		case 32:
			goto tr1325
		case 44:
			goto tr2091
		case 65:
			goto st476
		case 92:
			goto st434
		case 97:
			goto st479
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1325
		}
		goto st1
	st476:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof476
		}
	st_case_476:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr2
		case 11:
			goto tr3
		case 13:
			goto tr2
		case 32:
			goto tr1
		case 44:
			goto tr4
		case 76:
			goto st477
		case 92:
			goto st434
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1
		}
		goto st1
	st477:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof477
		}
	st_case_477:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr2
		case 11:
			goto tr3
		case 13:
			goto tr2
		case 32:
			goto tr1
		case 44:
			goto tr4
		case 83:
			goto st478
		case 92:
			goto st434
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1
		}
		goto st1
	st478:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof478
		}
	st_case_478:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr2
		case 11:
			goto tr3
		case 13:
			goto tr2
		case 32:
			goto tr1
		case 44:
			goto tr4
		case 69:
			goto st1494
		case 92:
			goto st434
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1
		}
		goto st1
	st1494:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1494
		}
	st_case_1494:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1280
		case 11:
			goto tr2090
		case 13:
			goto tr1280
		case 32:
			goto tr1325
		case 44:
			goto tr2091
		case 92:
			goto st434
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1325
		}
		goto st1
	st479:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof479
		}
	st_case_479:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr2
		case 11:
			goto tr3
		case 13:
			goto tr2
		case 32:
			goto tr1
		case 44:
			goto tr4
		case 92:
			goto st434
		case 108:
			goto st480
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1
		}
		goto st1
	st480:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof480
		}
	st_case_480:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr2
		case 11:
			goto tr3
		case 13:
			goto tr2
		case 32:
			goto tr1
		case 44:
			goto tr4
		case 92:
			goto st434
		case 115:
			goto st481
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1
		}
		goto st1
	st481:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof481
		}
	st_case_481:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr2
		case 11:
			goto tr3
		case 13:
			goto tr2
		case 32:
			goto tr1
		case 44:
			goto tr4
		case 92:
			goto st434
		case 101:
			goto st1494
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1
		}
		goto st1
tr373:
//line machine.rl:17

	m.pb = m.p

	goto st1495
	st1495:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1495
		}
	st_case_1495:
//line machine.go:55414
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1280
		case 11:
			goto tr2090
		case 13:
			goto tr1280
		case 32:
			goto tr1325
		case 44:
			goto tr2091
		case 82:
			goto st482
		case 92:
			goto st434
		case 114:
			goto st483
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1325
		}
		goto st1
	st482:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof482
		}
	st_case_482:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr2
		case 11:
			goto tr3
		case 13:
			goto tr2
		case 32:
			goto tr1
		case 44:
			goto tr4
		case 85:
			goto st478
		case 92:
			goto st434
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1
		}
		goto st1
	st483:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof483
		}
	st_case_483:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr2
		case 11:
			goto tr3
		case 13:
			goto tr2
		case 32:
			goto tr1
		case 44:
			goto tr4
		case 92:
			goto st434
		case 117:
			goto st481
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1
		}
		goto st1
tr374:
//line machine.rl:17

	m.pb = m.p

	goto st1496
	st1496:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1496
		}
	st_case_1496:
//line machine.go:55498
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1280
		case 11:
			goto tr2090
		case 13:
			goto tr1280
		case 32:
			goto tr1325
		case 44:
			goto tr2091
		case 92:
			goto st434
		case 97:
			goto st479
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1325
		}
		goto st1
tr375:
//line machine.rl:17

	m.pb = m.p

	goto st1497
	st1497:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1497
		}
	st_case_1497:
//line machine.go:55530
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr1280
		case 11:
			goto tr2090
		case 13:
			goto tr1280
		case 32:
			goto tr1325
		case 44:
			goto tr2091
		case 92:
			goto st434
		case 114:
			goto st483
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr1325
		}
		goto st1
	st484:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof484
		}
	st_case_484:
		switch ( m.data)[( m.p)] {
		case 10:
			goto st0
		case 11:
			goto tr839
		case 13:
			goto st0
		case 32:
			goto st484
		case 44:
			goto st0
		case 92:
			goto tr840
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto st484
		}
		goto tr837
tr839:
//line machine.rl:17

	m.pb = m.p

	goto st485
	st485:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof485
		}
	st_case_485:
//line machine.go:55585
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr2
		case 11:
			goto tr842
		case 13:
			goto tr2
		case 32:
			goto tr841
		case 44:
			goto tr4
		case 92:
			goto tr840
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr841
		}
		goto tr837
tr841:
//line machine.rl:58

	m.handler.SetMeasurement(m.text())

	goto st486
	st486:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof486
		}
	st_case_486:
//line machine.go:55615
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr2
		case 11:
			goto tr844
		case 13:
			goto tr2
		case 32:
			goto st486
		case 44:
			goto tr2
		case 61:
			goto tr837
		case 92:
			goto tr362
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto st486
		}
		goto tr360
tr844:
//line machine.rl:17

	m.pb = m.p

	goto st487
tr845:
//line machine.rl:17

	m.pb = m.p

//line machine.rl:58

	m.handler.SetMeasurement(m.text())

	goto st487
	st487:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof487
		}
	st_case_487:
//line machine.go:55657
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr2
		case 11:
			goto tr845
		case 13:
			goto tr2
		case 32:
			goto tr841
		case 44:
			goto tr4
		case 61:
			goto tr846
		case 92:
			goto tr362
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr841
		}
		goto tr360
tr842:
//line machine.rl:17

	m.pb = m.p

//line machine.rl:58

	m.handler.SetMeasurement(m.text())

	goto st488
	st488:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof488
		}
	st_case_488:
//line machine.go:55693
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr2
		case 11:
			goto tr845
		case 13:
			goto tr2
		case 32:
			goto tr841
		case 44:
			goto tr4
		case 61:
			goto tr837
		case 92:
			goto tr362
		}
		if 9 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 12 {
			goto tr841
		}
		goto tr360
	st489:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof489
		}
	st_case_489:
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr848
		case 13:
			goto tr848
		}
		goto st489
tr848:
//line machine.rl:54

	{goto st490 }

	goto st1498
	st1498:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof1498
		}
	st_case_1498:
//line machine.go:55737
		switch ( m.data)[( m.p)] {
		case 10:
			goto tr848
		case 13:
			goto tr848
		}
		goto st0
	st_out:
	_test_eof490:  m.cs = 490; goto _test_eof
	_test_eof1:  m.cs = 1; goto _test_eof
	_test_eof2:  m.cs = 2; goto _test_eof
	_test_eof3:  m.cs = 3; goto _test_eof
	_test_eof4:  m.cs = 4; goto _test_eof
	_test_eof5:  m.cs = 5; goto _test_eof
	_test_eof6:  m.cs = 6; goto _test_eof
	_test_eof491:  m.cs = 491; goto _test_eof
	_test_eof492:  m.cs = 492; goto _test_eof
	_test_eof493:  m.cs = 493; goto _test_eof
	_test_eof7:  m.cs = 7; goto _test_eof
	_test_eof8:  m.cs = 8; goto _test_eof
	_test_eof9:  m.cs = 9; goto _test_eof
	_test_eof10:  m.cs = 10; goto _test_eof
	_test_eof11:  m.cs = 11; goto _test_eof
	_test_eof12:  m.cs = 12; goto _test_eof
	_test_eof13:  m.cs = 13; goto _test_eof
	_test_eof14:  m.cs = 14; goto _test_eof
	_test_eof15:  m.cs = 15; goto _test_eof
	_test_eof16:  m.cs = 16; goto _test_eof
	_test_eof17:  m.cs = 17; goto _test_eof
	_test_eof18:  m.cs = 18; goto _test_eof
	_test_eof19:  m.cs = 19; goto _test_eof
	_test_eof20:  m.cs = 20; goto _test_eof
	_test_eof21:  m.cs = 21; goto _test_eof
	_test_eof22:  m.cs = 22; goto _test_eof
	_test_eof23:  m.cs = 23; goto _test_eof
	_test_eof24:  m.cs = 24; goto _test_eof
	_test_eof25:  m.cs = 25; goto _test_eof
	_test_eof494:  m.cs = 494; goto _test_eof
	_test_eof495:  m.cs = 495; goto _test_eof
	_test_eof26:  m.cs = 26; goto _test_eof
	_test_eof496:  m.cs = 496; goto _test_eof
	_test_eof497:  m.cs = 497; goto _test_eof
	_test_eof498:  m.cs = 498; goto _test_eof
	_test_eof27:  m.cs = 27; goto _test_eof
	_test_eof499:  m.cs = 499; goto _test_eof
	_test_eof500:  m.cs = 500; goto _test_eof
	_test_eof501:  m.cs = 501; goto _test_eof
	_test_eof502:  m.cs = 502; goto _test_eof
	_test_eof503:  m.cs = 503; goto _test_eof
	_test_eof504:  m.cs = 504; goto _test_eof
	_test_eof505:  m.cs = 505; goto _test_eof
	_test_eof506:  m.cs = 506; goto _test_eof
	_test_eof507:  m.cs = 507; goto _test_eof
	_test_eof508:  m.cs = 508; goto _test_eof
	_test_eof509:  m.cs = 509; goto _test_eof
	_test_eof510:  m.cs = 510; goto _test_eof
	_test_eof511:  m.cs = 511; goto _test_eof
	_test_eof512:  m.cs = 512; goto _test_eof
	_test_eof513:  m.cs = 513; goto _test_eof
	_test_eof514:  m.cs = 514; goto _test_eof
	_test_eof515:  m.cs = 515; goto _test_eof
	_test_eof516:  m.cs = 516; goto _test_eof
	_test_eof28:  m.cs = 28; goto _test_eof
	_test_eof29:  m.cs = 29; goto _test_eof
	_test_eof517:  m.cs = 517; goto _test_eof
	_test_eof518:  m.cs = 518; goto _test_eof
	_test_eof519:  m.cs = 519; goto _test_eof
	_test_eof30:  m.cs = 30; goto _test_eof
	_test_eof31:  m.cs = 31; goto _test_eof
	_test_eof32:  m.cs = 32; goto _test_eof
	_test_eof33:  m.cs = 33; goto _test_eof
	_test_eof34:  m.cs = 34; goto _test_eof
	_test_eof520:  m.cs = 520; goto _test_eof
	_test_eof521:  m.cs = 521; goto _test_eof
	_test_eof522:  m.cs = 522; goto _test_eof
	_test_eof523:  m.cs = 523; goto _test_eof
	_test_eof35:  m.cs = 35; goto _test_eof
	_test_eof524:  m.cs = 524; goto _test_eof
	_test_eof525:  m.cs = 525; goto _test_eof
	_test_eof526:  m.cs = 526; goto _test_eof
	_test_eof527:  m.cs = 527; goto _test_eof
	_test_eof36:  m.cs = 36; goto _test_eof
	_test_eof528:  m.cs = 528; goto _test_eof
	_test_eof529:  m.cs = 529; goto _test_eof
	_test_eof530:  m.cs = 530; goto _test_eof
	_test_eof531:  m.cs = 531; goto _test_eof
	_test_eof532:  m.cs = 532; goto _test_eof
	_test_eof533:  m.cs = 533; goto _test_eof
	_test_eof534:  m.cs = 534; goto _test_eof
	_test_eof535:  m.cs = 535; goto _test_eof
	_test_eof536:  m.cs = 536; goto _test_eof
	_test_eof537:  m.cs = 537; goto _test_eof
	_test_eof538:  m.cs = 538; goto _test_eof
	_test_eof539:  m.cs = 539; goto _test_eof
	_test_eof540:  m.cs = 540; goto _test_eof
	_test_eof541:  m.cs = 541; goto _test_eof
	_test_eof542:  m.cs = 542; goto _test_eof
	_test_eof543:  m.cs = 543; goto _test_eof
	_test_eof544:  m.cs = 544; goto _test_eof
	_test_eof545:  m.cs = 545; goto _test_eof
	_test_eof37:  m.cs = 37; goto _test_eof
	_test_eof38:  m.cs = 38; goto _test_eof
	_test_eof39:  m.cs = 39; goto _test_eof
	_test_eof40:  m.cs = 40; goto _test_eof
	_test_eof41:  m.cs = 41; goto _test_eof
	_test_eof42:  m.cs = 42; goto _test_eof
	_test_eof43:  m.cs = 43; goto _test_eof
	_test_eof44:  m.cs = 44; goto _test_eof
	_test_eof45:  m.cs = 45; goto _test_eof
	_test_eof546:  m.cs = 546; goto _test_eof
	_test_eof547:  m.cs = 547; goto _test_eof
	_test_eof548:  m.cs = 548; goto _test_eof
	_test_eof46:  m.cs = 46; goto _test_eof
	_test_eof47:  m.cs = 47; goto _test_eof
	_test_eof48:  m.cs = 48; goto _test_eof
	_test_eof49:  m.cs = 49; goto _test_eof
	_test_eof50:  m.cs = 50; goto _test_eof
	_test_eof51:  m.cs = 51; goto _test_eof
	_test_eof549:  m.cs = 549; goto _test_eof
	_test_eof550:  m.cs = 550; goto _test_eof
	_test_eof52:  m.cs = 52; goto _test_eof
	_test_eof551:  m.cs = 551; goto _test_eof
	_test_eof552:  m.cs = 552; goto _test_eof
	_test_eof553:  m.cs = 553; goto _test_eof
	_test_eof554:  m.cs = 554; goto _test_eof
	_test_eof555:  m.cs = 555; goto _test_eof
	_test_eof556:  m.cs = 556; goto _test_eof
	_test_eof557:  m.cs = 557; goto _test_eof
	_test_eof558:  m.cs = 558; goto _test_eof
	_test_eof559:  m.cs = 559; goto _test_eof
	_test_eof560:  m.cs = 560; goto _test_eof
	_test_eof561:  m.cs = 561; goto _test_eof
	_test_eof562:  m.cs = 562; goto _test_eof
	_test_eof563:  m.cs = 563; goto _test_eof
	_test_eof564:  m.cs = 564; goto _test_eof
	_test_eof565:  m.cs = 565; goto _test_eof
	_test_eof566:  m.cs = 566; goto _test_eof
	_test_eof567:  m.cs = 567; goto _test_eof
	_test_eof568:  m.cs = 568; goto _test_eof
	_test_eof569:  m.cs = 569; goto _test_eof
	_test_eof570:  m.cs = 570; goto _test_eof
	_test_eof53:  m.cs = 53; goto _test_eof
	_test_eof571:  m.cs = 571; goto _test_eof
	_test_eof572:  m.cs = 572; goto _test_eof
	_test_eof573:  m.cs = 573; goto _test_eof
	_test_eof54:  m.cs = 54; goto _test_eof
	_test_eof574:  m.cs = 574; goto _test_eof
	_test_eof575:  m.cs = 575; goto _test_eof
	_test_eof576:  m.cs = 576; goto _test_eof
	_test_eof577:  m.cs = 577; goto _test_eof
	_test_eof578:  m.cs = 578; goto _test_eof
	_test_eof579:  m.cs = 579; goto _test_eof
	_test_eof580:  m.cs = 580; goto _test_eof
	_test_eof581:  m.cs = 581; goto _test_eof
	_test_eof582:  m.cs = 582; goto _test_eof
	_test_eof583:  m.cs = 583; goto _test_eof
	_test_eof584:  m.cs = 584; goto _test_eof
	_test_eof585:  m.cs = 585; goto _test_eof
	_test_eof586:  m.cs = 586; goto _test_eof
	_test_eof587:  m.cs = 587; goto _test_eof
	_test_eof588:  m.cs = 588; goto _test_eof
	_test_eof589:  m.cs = 589; goto _test_eof
	_test_eof590:  m.cs = 590; goto _test_eof
	_test_eof591:  m.cs = 591; goto _test_eof
	_test_eof592:  m.cs = 592; goto _test_eof
	_test_eof593:  m.cs = 593; goto _test_eof
	_test_eof55:  m.cs = 55; goto _test_eof
	_test_eof56:  m.cs = 56; goto _test_eof
	_test_eof57:  m.cs = 57; goto _test_eof
	_test_eof594:  m.cs = 594; goto _test_eof
	_test_eof595:  m.cs = 595; goto _test_eof
	_test_eof596:  m.cs = 596; goto _test_eof
	_test_eof58:  m.cs = 58; goto _test_eof
	_test_eof597:  m.cs = 597; goto _test_eof
	_test_eof598:  m.cs = 598; goto _test_eof
	_test_eof59:  m.cs = 59; goto _test_eof
	_test_eof599:  m.cs = 599; goto _test_eof
	_test_eof60:  m.cs = 60; goto _test_eof
	_test_eof600:  m.cs = 600; goto _test_eof
	_test_eof601:  m.cs = 601; goto _test_eof
	_test_eof602:  m.cs = 602; goto _test_eof
	_test_eof603:  m.cs = 603; goto _test_eof
	_test_eof604:  m.cs = 604; goto _test_eof
	_test_eof605:  m.cs = 605; goto _test_eof
	_test_eof606:  m.cs = 606; goto _test_eof
	_test_eof607:  m.cs = 607; goto _test_eof
	_test_eof608:  m.cs = 608; goto _test_eof
	_test_eof609:  m.cs = 609; goto _test_eof
	_test_eof610:  m.cs = 610; goto _test_eof
	_test_eof611:  m.cs = 611; goto _test_eof
	_test_eof612:  m.cs = 612; goto _test_eof
	_test_eof613:  m.cs = 613; goto _test_eof
	_test_eof614:  m.cs = 614; goto _test_eof
	_test_eof615:  m.cs = 615; goto _test_eof
	_test_eof616:  m.cs = 616; goto _test_eof
	_test_eof617:  m.cs = 617; goto _test_eof
	_test_eof618:  m.cs = 618; goto _test_eof
	_test_eof619:  m.cs = 619; goto _test_eof
	_test_eof61:  m.cs = 61; goto _test_eof
	_test_eof62:  m.cs = 62; goto _test_eof
	_test_eof63:  m.cs = 63; goto _test_eof
	_test_eof64:  m.cs = 64; goto _test_eof
	_test_eof65:  m.cs = 65; goto _test_eof
	_test_eof66:  m.cs = 66; goto _test_eof
	_test_eof67:  m.cs = 67; goto _test_eof
	_test_eof620:  m.cs = 620; goto _test_eof
	_test_eof68:  m.cs = 68; goto _test_eof
	_test_eof69:  m.cs = 69; goto _test_eof
	_test_eof70:  m.cs = 70; goto _test_eof
	_test_eof71:  m.cs = 71; goto _test_eof
	_test_eof621:  m.cs = 621; goto _test_eof
	_test_eof622:  m.cs = 622; goto _test_eof
	_test_eof623:  m.cs = 623; goto _test_eof
	_test_eof624:  m.cs = 624; goto _test_eof
	_test_eof72:  m.cs = 72; goto _test_eof
	_test_eof73:  m.cs = 73; goto _test_eof
	_test_eof74:  m.cs = 74; goto _test_eof
	_test_eof75:  m.cs = 75; goto _test_eof
	_test_eof625:  m.cs = 625; goto _test_eof
	_test_eof626:  m.cs = 626; goto _test_eof
	_test_eof76:  m.cs = 76; goto _test_eof
	_test_eof627:  m.cs = 627; goto _test_eof
	_test_eof77:  m.cs = 77; goto _test_eof
	_test_eof78:  m.cs = 78; goto _test_eof
	_test_eof628:  m.cs = 628; goto _test_eof
	_test_eof629:  m.cs = 629; goto _test_eof
	_test_eof79:  m.cs = 79; goto _test_eof
	_test_eof630:  m.cs = 630; goto _test_eof
	_test_eof631:  m.cs = 631; goto _test_eof
	_test_eof80:  m.cs = 80; goto _test_eof
	_test_eof632:  m.cs = 632; goto _test_eof
	_test_eof633:  m.cs = 633; goto _test_eof
	_test_eof634:  m.cs = 634; goto _test_eof
	_test_eof635:  m.cs = 635; goto _test_eof
	_test_eof636:  m.cs = 636; goto _test_eof
	_test_eof637:  m.cs = 637; goto _test_eof
	_test_eof638:  m.cs = 638; goto _test_eof
	_test_eof639:  m.cs = 639; goto _test_eof
	_test_eof640:  m.cs = 640; goto _test_eof
	_test_eof641:  m.cs = 641; goto _test_eof
	_test_eof642:  m.cs = 642; goto _test_eof
	_test_eof643:  m.cs = 643; goto _test_eof
	_test_eof644:  m.cs = 644; goto _test_eof
	_test_eof645:  m.cs = 645; goto _test_eof
	_test_eof646:  m.cs = 646; goto _test_eof
	_test_eof647:  m.cs = 647; goto _test_eof
	_test_eof648:  m.cs = 648; goto _test_eof
	_test_eof649:  m.cs = 649; goto _test_eof
	_test_eof81:  m.cs = 81; goto _test_eof
	_test_eof650:  m.cs = 650; goto _test_eof
	_test_eof651:  m.cs = 651; goto _test_eof
	_test_eof652:  m.cs = 652; goto _test_eof
	_test_eof82:  m.cs = 82; goto _test_eof
	_test_eof653:  m.cs = 653; goto _test_eof
	_test_eof654:  m.cs = 654; goto _test_eof
	_test_eof655:  m.cs = 655; goto _test_eof
	_test_eof83:  m.cs = 83; goto _test_eof
	_test_eof656:  m.cs = 656; goto _test_eof
	_test_eof657:  m.cs = 657; goto _test_eof
	_test_eof658:  m.cs = 658; goto _test_eof
	_test_eof659:  m.cs = 659; goto _test_eof
	_test_eof660:  m.cs = 660; goto _test_eof
	_test_eof661:  m.cs = 661; goto _test_eof
	_test_eof662:  m.cs = 662; goto _test_eof
	_test_eof663:  m.cs = 663; goto _test_eof
	_test_eof664:  m.cs = 664; goto _test_eof
	_test_eof665:  m.cs = 665; goto _test_eof
	_test_eof666:  m.cs = 666; goto _test_eof
	_test_eof667:  m.cs = 667; goto _test_eof
	_test_eof668:  m.cs = 668; goto _test_eof
	_test_eof669:  m.cs = 669; goto _test_eof
	_test_eof670:  m.cs = 670; goto _test_eof
	_test_eof671:  m.cs = 671; goto _test_eof
	_test_eof672:  m.cs = 672; goto _test_eof
	_test_eof673:  m.cs = 673; goto _test_eof
	_test_eof674:  m.cs = 674; goto _test_eof
	_test_eof84:  m.cs = 84; goto _test_eof
	_test_eof85:  m.cs = 85; goto _test_eof
	_test_eof86:  m.cs = 86; goto _test_eof
	_test_eof675:  m.cs = 675; goto _test_eof
	_test_eof87:  m.cs = 87; goto _test_eof
	_test_eof676:  m.cs = 676; goto _test_eof
	_test_eof677:  m.cs = 677; goto _test_eof
	_test_eof678:  m.cs = 678; goto _test_eof
	_test_eof679:  m.cs = 679; goto _test_eof
	_test_eof680:  m.cs = 680; goto _test_eof
	_test_eof681:  m.cs = 681; goto _test_eof
	_test_eof682:  m.cs = 682; goto _test_eof
	_test_eof683:  m.cs = 683; goto _test_eof
	_test_eof684:  m.cs = 684; goto _test_eof
	_test_eof685:  m.cs = 685; goto _test_eof
	_test_eof686:  m.cs = 686; goto _test_eof
	_test_eof687:  m.cs = 687; goto _test_eof
	_test_eof688:  m.cs = 688; goto _test_eof
	_test_eof689:  m.cs = 689; goto _test_eof
	_test_eof690:  m.cs = 690; goto _test_eof
	_test_eof691:  m.cs = 691; goto _test_eof
	_test_eof692:  m.cs = 692; goto _test_eof
	_test_eof693:  m.cs = 693; goto _test_eof
	_test_eof694:  m.cs = 694; goto _test_eof
	_test_eof695:  m.cs = 695; goto _test_eof
	_test_eof696:  m.cs = 696; goto _test_eof
	_test_eof697:  m.cs = 697; goto _test_eof
	_test_eof88:  m.cs = 88; goto _test_eof
	_test_eof89:  m.cs = 89; goto _test_eof
	_test_eof90:  m.cs = 90; goto _test_eof
	_test_eof91:  m.cs = 91; goto _test_eof
	_test_eof92:  m.cs = 92; goto _test_eof
	_test_eof698:  m.cs = 698; goto _test_eof
	_test_eof699:  m.cs = 699; goto _test_eof
	_test_eof700:  m.cs = 700; goto _test_eof
	_test_eof701:  m.cs = 701; goto _test_eof
	_test_eof702:  m.cs = 702; goto _test_eof
	_test_eof93:  m.cs = 93; goto _test_eof
	_test_eof94:  m.cs = 94; goto _test_eof
	_test_eof95:  m.cs = 95; goto _test_eof
	_test_eof703:  m.cs = 703; goto _test_eof
	_test_eof96:  m.cs = 96; goto _test_eof
	_test_eof97:  m.cs = 97; goto _test_eof
	_test_eof98:  m.cs = 98; goto _test_eof
	_test_eof704:  m.cs = 704; goto _test_eof
	_test_eof99:  m.cs = 99; goto _test_eof
	_test_eof100:  m.cs = 100; goto _test_eof
	_test_eof705:  m.cs = 705; goto _test_eof
	_test_eof706:  m.cs = 706; goto _test_eof
	_test_eof101:  m.cs = 101; goto _test_eof
	_test_eof102:  m.cs = 102; goto _test_eof
	_test_eof707:  m.cs = 707; goto _test_eof
	_test_eof708:  m.cs = 708; goto _test_eof
	_test_eof709:  m.cs = 709; goto _test_eof
	_test_eof103:  m.cs = 103; goto _test_eof
	_test_eof710:  m.cs = 710; goto _test_eof
	_test_eof711:  m.cs = 711; goto _test_eof
	_test_eof712:  m.cs = 712; goto _test_eof
	_test_eof713:  m.cs = 713; goto _test_eof
	_test_eof714:  m.cs = 714; goto _test_eof
	_test_eof715:  m.cs = 715; goto _test_eof
	_test_eof716:  m.cs = 716; goto _test_eof
	_test_eof717:  m.cs = 717; goto _test_eof
	_test_eof718:  m.cs = 718; goto _test_eof
	_test_eof719:  m.cs = 719; goto _test_eof
	_test_eof720:  m.cs = 720; goto _test_eof
	_test_eof721:  m.cs = 721; goto _test_eof
	_test_eof722:  m.cs = 722; goto _test_eof
	_test_eof723:  m.cs = 723; goto _test_eof
	_test_eof724:  m.cs = 724; goto _test_eof
	_test_eof725:  m.cs = 725; goto _test_eof
	_test_eof726:  m.cs = 726; goto _test_eof
	_test_eof727:  m.cs = 727; goto _test_eof
	_test_eof728:  m.cs = 728; goto _test_eof
	_test_eof729:  m.cs = 729; goto _test_eof
	_test_eof730:  m.cs = 730; goto _test_eof
	_test_eof731:  m.cs = 731; goto _test_eof
	_test_eof732:  m.cs = 732; goto _test_eof
	_test_eof733:  m.cs = 733; goto _test_eof
	_test_eof104:  m.cs = 104; goto _test_eof
	_test_eof105:  m.cs = 105; goto _test_eof
	_test_eof106:  m.cs = 106; goto _test_eof
	_test_eof734:  m.cs = 734; goto _test_eof
	_test_eof107:  m.cs = 107; goto _test_eof
	_test_eof108:  m.cs = 108; goto _test_eof
	_test_eof109:  m.cs = 109; goto _test_eof
	_test_eof735:  m.cs = 735; goto _test_eof
	_test_eof110:  m.cs = 110; goto _test_eof
	_test_eof111:  m.cs = 111; goto _test_eof
	_test_eof736:  m.cs = 736; goto _test_eof
	_test_eof737:  m.cs = 737; goto _test_eof
	_test_eof112:  m.cs = 112; goto _test_eof
	_test_eof738:  m.cs = 738; goto _test_eof
	_test_eof113:  m.cs = 113; goto _test_eof
	_test_eof739:  m.cs = 739; goto _test_eof
	_test_eof740:  m.cs = 740; goto _test_eof
	_test_eof741:  m.cs = 741; goto _test_eof
	_test_eof114:  m.cs = 114; goto _test_eof
	_test_eof115:  m.cs = 115; goto _test_eof
	_test_eof116:  m.cs = 116; goto _test_eof
	_test_eof742:  m.cs = 742; goto _test_eof
	_test_eof117:  m.cs = 117; goto _test_eof
	_test_eof118:  m.cs = 118; goto _test_eof
	_test_eof119:  m.cs = 119; goto _test_eof
	_test_eof743:  m.cs = 743; goto _test_eof
	_test_eof120:  m.cs = 120; goto _test_eof
	_test_eof121:  m.cs = 121; goto _test_eof
	_test_eof744:  m.cs = 744; goto _test_eof
	_test_eof745:  m.cs = 745; goto _test_eof
	_test_eof746:  m.cs = 746; goto _test_eof
	_test_eof747:  m.cs = 747; goto _test_eof
	_test_eof748:  m.cs = 748; goto _test_eof
	_test_eof749:  m.cs = 749; goto _test_eof
	_test_eof750:  m.cs = 750; goto _test_eof
	_test_eof751:  m.cs = 751; goto _test_eof
	_test_eof752:  m.cs = 752; goto _test_eof
	_test_eof753:  m.cs = 753; goto _test_eof
	_test_eof754:  m.cs = 754; goto _test_eof
	_test_eof755:  m.cs = 755; goto _test_eof
	_test_eof756:  m.cs = 756; goto _test_eof
	_test_eof757:  m.cs = 757; goto _test_eof
	_test_eof758:  m.cs = 758; goto _test_eof
	_test_eof759:  m.cs = 759; goto _test_eof
	_test_eof760:  m.cs = 760; goto _test_eof
	_test_eof761:  m.cs = 761; goto _test_eof
	_test_eof762:  m.cs = 762; goto _test_eof
	_test_eof763:  m.cs = 763; goto _test_eof
	_test_eof122:  m.cs = 122; goto _test_eof
	_test_eof764:  m.cs = 764; goto _test_eof
	_test_eof765:  m.cs = 765; goto _test_eof
	_test_eof766:  m.cs = 766; goto _test_eof
	_test_eof123:  m.cs = 123; goto _test_eof
	_test_eof767:  m.cs = 767; goto _test_eof
	_test_eof768:  m.cs = 768; goto _test_eof
	_test_eof124:  m.cs = 124; goto _test_eof
	_test_eof769:  m.cs = 769; goto _test_eof
	_test_eof770:  m.cs = 770; goto _test_eof
	_test_eof771:  m.cs = 771; goto _test_eof
	_test_eof772:  m.cs = 772; goto _test_eof
	_test_eof773:  m.cs = 773; goto _test_eof
	_test_eof774:  m.cs = 774; goto _test_eof
	_test_eof775:  m.cs = 775; goto _test_eof
	_test_eof776:  m.cs = 776; goto _test_eof
	_test_eof777:  m.cs = 777; goto _test_eof
	_test_eof778:  m.cs = 778; goto _test_eof
	_test_eof779:  m.cs = 779; goto _test_eof
	_test_eof780:  m.cs = 780; goto _test_eof
	_test_eof781:  m.cs = 781; goto _test_eof
	_test_eof782:  m.cs = 782; goto _test_eof
	_test_eof783:  m.cs = 783; goto _test_eof
	_test_eof784:  m.cs = 784; goto _test_eof
	_test_eof785:  m.cs = 785; goto _test_eof
	_test_eof786:  m.cs = 786; goto _test_eof
	_test_eof787:  m.cs = 787; goto _test_eof
	_test_eof125:  m.cs = 125; goto _test_eof
	_test_eof788:  m.cs = 788; goto _test_eof
	_test_eof789:  m.cs = 789; goto _test_eof
	_test_eof790:  m.cs = 790; goto _test_eof
	_test_eof126:  m.cs = 126; goto _test_eof
	_test_eof127:  m.cs = 127; goto _test_eof
	_test_eof128:  m.cs = 128; goto _test_eof
	_test_eof791:  m.cs = 791; goto _test_eof
	_test_eof129:  m.cs = 129; goto _test_eof
	_test_eof130:  m.cs = 130; goto _test_eof
	_test_eof131:  m.cs = 131; goto _test_eof
	_test_eof792:  m.cs = 792; goto _test_eof
	_test_eof132:  m.cs = 132; goto _test_eof
	_test_eof133:  m.cs = 133; goto _test_eof
	_test_eof793:  m.cs = 793; goto _test_eof
	_test_eof794:  m.cs = 794; goto _test_eof
	_test_eof134:  m.cs = 134; goto _test_eof
	_test_eof135:  m.cs = 135; goto _test_eof
	_test_eof136:  m.cs = 136; goto _test_eof
	_test_eof137:  m.cs = 137; goto _test_eof
	_test_eof138:  m.cs = 138; goto _test_eof
	_test_eof139:  m.cs = 139; goto _test_eof
	_test_eof795:  m.cs = 795; goto _test_eof
	_test_eof796:  m.cs = 796; goto _test_eof
	_test_eof797:  m.cs = 797; goto _test_eof
	_test_eof798:  m.cs = 798; goto _test_eof
	_test_eof799:  m.cs = 799; goto _test_eof
	_test_eof800:  m.cs = 800; goto _test_eof
	_test_eof801:  m.cs = 801; goto _test_eof
	_test_eof802:  m.cs = 802; goto _test_eof
	_test_eof803:  m.cs = 803; goto _test_eof
	_test_eof804:  m.cs = 804; goto _test_eof
	_test_eof805:  m.cs = 805; goto _test_eof
	_test_eof806:  m.cs = 806; goto _test_eof
	_test_eof807:  m.cs = 807; goto _test_eof
	_test_eof808:  m.cs = 808; goto _test_eof
	_test_eof809:  m.cs = 809; goto _test_eof
	_test_eof810:  m.cs = 810; goto _test_eof
	_test_eof811:  m.cs = 811; goto _test_eof
	_test_eof812:  m.cs = 812; goto _test_eof
	_test_eof813:  m.cs = 813; goto _test_eof
	_test_eof140:  m.cs = 140; goto _test_eof
	_test_eof141:  m.cs = 141; goto _test_eof
	_test_eof142:  m.cs = 142; goto _test_eof
	_test_eof814:  m.cs = 814; goto _test_eof
	_test_eof815:  m.cs = 815; goto _test_eof
	_test_eof816:  m.cs = 816; goto _test_eof
	_test_eof817:  m.cs = 817; goto _test_eof
	_test_eof818:  m.cs = 818; goto _test_eof
	_test_eof143:  m.cs = 143; goto _test_eof
	_test_eof144:  m.cs = 144; goto _test_eof
	_test_eof145:  m.cs = 145; goto _test_eof
	_test_eof819:  m.cs = 819; goto _test_eof
	_test_eof146:  m.cs = 146; goto _test_eof
	_test_eof147:  m.cs = 147; goto _test_eof
	_test_eof148:  m.cs = 148; goto _test_eof
	_test_eof820:  m.cs = 820; goto _test_eof
	_test_eof149:  m.cs = 149; goto _test_eof
	_test_eof150:  m.cs = 150; goto _test_eof
	_test_eof821:  m.cs = 821; goto _test_eof
	_test_eof822:  m.cs = 822; goto _test_eof
	_test_eof151:  m.cs = 151; goto _test_eof
	_test_eof152:  m.cs = 152; goto _test_eof
	_test_eof153:  m.cs = 153; goto _test_eof
	_test_eof823:  m.cs = 823; goto _test_eof
	_test_eof824:  m.cs = 824; goto _test_eof
	_test_eof825:  m.cs = 825; goto _test_eof
	_test_eof826:  m.cs = 826; goto _test_eof
	_test_eof827:  m.cs = 827; goto _test_eof
	_test_eof154:  m.cs = 154; goto _test_eof
	_test_eof155:  m.cs = 155; goto _test_eof
	_test_eof156:  m.cs = 156; goto _test_eof
	_test_eof828:  m.cs = 828; goto _test_eof
	_test_eof157:  m.cs = 157; goto _test_eof
	_test_eof158:  m.cs = 158; goto _test_eof
	_test_eof159:  m.cs = 159; goto _test_eof
	_test_eof829:  m.cs = 829; goto _test_eof
	_test_eof160:  m.cs = 160; goto _test_eof
	_test_eof161:  m.cs = 161; goto _test_eof
	_test_eof830:  m.cs = 830; goto _test_eof
	_test_eof831:  m.cs = 831; goto _test_eof
	_test_eof162:  m.cs = 162; goto _test_eof
	_test_eof163:  m.cs = 163; goto _test_eof
	_test_eof164:  m.cs = 164; goto _test_eof
	_test_eof832:  m.cs = 832; goto _test_eof
	_test_eof833:  m.cs = 833; goto _test_eof
	_test_eof834:  m.cs = 834; goto _test_eof
	_test_eof835:  m.cs = 835; goto _test_eof
	_test_eof836:  m.cs = 836; goto _test_eof
	_test_eof837:  m.cs = 837; goto _test_eof
	_test_eof838:  m.cs = 838; goto _test_eof
	_test_eof839:  m.cs = 839; goto _test_eof
	_test_eof840:  m.cs = 840; goto _test_eof
	_test_eof841:  m.cs = 841; goto _test_eof
	_test_eof842:  m.cs = 842; goto _test_eof
	_test_eof843:  m.cs = 843; goto _test_eof
	_test_eof844:  m.cs = 844; goto _test_eof
	_test_eof845:  m.cs = 845; goto _test_eof
	_test_eof846:  m.cs = 846; goto _test_eof
	_test_eof847:  m.cs = 847; goto _test_eof
	_test_eof848:  m.cs = 848; goto _test_eof
	_test_eof849:  m.cs = 849; goto _test_eof
	_test_eof850:  m.cs = 850; goto _test_eof
	_test_eof165:  m.cs = 165; goto _test_eof
	_test_eof166:  m.cs = 166; goto _test_eof
	_test_eof167:  m.cs = 167; goto _test_eof
	_test_eof168:  m.cs = 168; goto _test_eof
	_test_eof851:  m.cs = 851; goto _test_eof
	_test_eof852:  m.cs = 852; goto _test_eof
	_test_eof853:  m.cs = 853; goto _test_eof
	_test_eof854:  m.cs = 854; goto _test_eof
	_test_eof855:  m.cs = 855; goto _test_eof
	_test_eof169:  m.cs = 169; goto _test_eof
	_test_eof170:  m.cs = 170; goto _test_eof
	_test_eof171:  m.cs = 171; goto _test_eof
	_test_eof856:  m.cs = 856; goto _test_eof
	_test_eof172:  m.cs = 172; goto _test_eof
	_test_eof173:  m.cs = 173; goto _test_eof
	_test_eof174:  m.cs = 174; goto _test_eof
	_test_eof857:  m.cs = 857; goto _test_eof
	_test_eof175:  m.cs = 175; goto _test_eof
	_test_eof176:  m.cs = 176; goto _test_eof
	_test_eof858:  m.cs = 858; goto _test_eof
	_test_eof859:  m.cs = 859; goto _test_eof
	_test_eof177:  m.cs = 177; goto _test_eof
	_test_eof178:  m.cs = 178; goto _test_eof
	_test_eof179:  m.cs = 179; goto _test_eof
	_test_eof860:  m.cs = 860; goto _test_eof
	_test_eof861:  m.cs = 861; goto _test_eof
	_test_eof862:  m.cs = 862; goto _test_eof
	_test_eof863:  m.cs = 863; goto _test_eof
	_test_eof864:  m.cs = 864; goto _test_eof
	_test_eof180:  m.cs = 180; goto _test_eof
	_test_eof181:  m.cs = 181; goto _test_eof
	_test_eof182:  m.cs = 182; goto _test_eof
	_test_eof865:  m.cs = 865; goto _test_eof
	_test_eof183:  m.cs = 183; goto _test_eof
	_test_eof184:  m.cs = 184; goto _test_eof
	_test_eof185:  m.cs = 185; goto _test_eof
	_test_eof866:  m.cs = 866; goto _test_eof
	_test_eof186:  m.cs = 186; goto _test_eof
	_test_eof187:  m.cs = 187; goto _test_eof
	_test_eof867:  m.cs = 867; goto _test_eof
	_test_eof868:  m.cs = 868; goto _test_eof
	_test_eof188:  m.cs = 188; goto _test_eof
	_test_eof189:  m.cs = 189; goto _test_eof
	_test_eof190:  m.cs = 190; goto _test_eof
	_test_eof191:  m.cs = 191; goto _test_eof
	_test_eof192:  m.cs = 192; goto _test_eof
	_test_eof193:  m.cs = 193; goto _test_eof
	_test_eof869:  m.cs = 869; goto _test_eof
	_test_eof870:  m.cs = 870; goto _test_eof
	_test_eof871:  m.cs = 871; goto _test_eof
	_test_eof872:  m.cs = 872; goto _test_eof
	_test_eof873:  m.cs = 873; goto _test_eof
	_test_eof874:  m.cs = 874; goto _test_eof
	_test_eof875:  m.cs = 875; goto _test_eof
	_test_eof876:  m.cs = 876; goto _test_eof
	_test_eof877:  m.cs = 877; goto _test_eof
	_test_eof878:  m.cs = 878; goto _test_eof
	_test_eof879:  m.cs = 879; goto _test_eof
	_test_eof880:  m.cs = 880; goto _test_eof
	_test_eof881:  m.cs = 881; goto _test_eof
	_test_eof882:  m.cs = 882; goto _test_eof
	_test_eof883:  m.cs = 883; goto _test_eof
	_test_eof884:  m.cs = 884; goto _test_eof
	_test_eof885:  m.cs = 885; goto _test_eof
	_test_eof886:  m.cs = 886; goto _test_eof
	_test_eof887:  m.cs = 887; goto _test_eof
	_test_eof194:  m.cs = 194; goto _test_eof
	_test_eof888:  m.cs = 888; goto _test_eof
	_test_eof889:  m.cs = 889; goto _test_eof
	_test_eof890:  m.cs = 890; goto _test_eof
	_test_eof891:  m.cs = 891; goto _test_eof
	_test_eof892:  m.cs = 892; goto _test_eof
	_test_eof195:  m.cs = 195; goto _test_eof
	_test_eof196:  m.cs = 196; goto _test_eof
	_test_eof197:  m.cs = 197; goto _test_eof
	_test_eof893:  m.cs = 893; goto _test_eof
	_test_eof198:  m.cs = 198; goto _test_eof
	_test_eof199:  m.cs = 199; goto _test_eof
	_test_eof200:  m.cs = 200; goto _test_eof
	_test_eof894:  m.cs = 894; goto _test_eof
	_test_eof201:  m.cs = 201; goto _test_eof
	_test_eof202:  m.cs = 202; goto _test_eof
	_test_eof895:  m.cs = 895; goto _test_eof
	_test_eof896:  m.cs = 896; goto _test_eof
	_test_eof203:  m.cs = 203; goto _test_eof
	_test_eof204:  m.cs = 204; goto _test_eof
	_test_eof205:  m.cs = 205; goto _test_eof
	_test_eof206:  m.cs = 206; goto _test_eof
	_test_eof207:  m.cs = 207; goto _test_eof
	_test_eof208:  m.cs = 208; goto _test_eof
	_test_eof209:  m.cs = 209; goto _test_eof
	_test_eof210:  m.cs = 210; goto _test_eof
	_test_eof211:  m.cs = 211; goto _test_eof
	_test_eof212:  m.cs = 212; goto _test_eof
	_test_eof897:  m.cs = 897; goto _test_eof
	_test_eof898:  m.cs = 898; goto _test_eof
	_test_eof899:  m.cs = 899; goto _test_eof
	_test_eof213:  m.cs = 213; goto _test_eof
	_test_eof214:  m.cs = 214; goto _test_eof
	_test_eof215:  m.cs = 215; goto _test_eof
	_test_eof900:  m.cs = 900; goto _test_eof
	_test_eof901:  m.cs = 901; goto _test_eof
	_test_eof216:  m.cs = 216; goto _test_eof
	_test_eof902:  m.cs = 902; goto _test_eof
	_test_eof903:  m.cs = 903; goto _test_eof
	_test_eof904:  m.cs = 904; goto _test_eof
	_test_eof905:  m.cs = 905; goto _test_eof
	_test_eof906:  m.cs = 906; goto _test_eof
	_test_eof907:  m.cs = 907; goto _test_eof
	_test_eof908:  m.cs = 908; goto _test_eof
	_test_eof909:  m.cs = 909; goto _test_eof
	_test_eof910:  m.cs = 910; goto _test_eof
	_test_eof911:  m.cs = 911; goto _test_eof
	_test_eof912:  m.cs = 912; goto _test_eof
	_test_eof913:  m.cs = 913; goto _test_eof
	_test_eof914:  m.cs = 914; goto _test_eof
	_test_eof915:  m.cs = 915; goto _test_eof
	_test_eof916:  m.cs = 916; goto _test_eof
	_test_eof917:  m.cs = 917; goto _test_eof
	_test_eof918:  m.cs = 918; goto _test_eof
	_test_eof919:  m.cs = 919; goto _test_eof
	_test_eof920:  m.cs = 920; goto _test_eof
	_test_eof217:  m.cs = 217; goto _test_eof
	_test_eof921:  m.cs = 921; goto _test_eof
	_test_eof922:  m.cs = 922; goto _test_eof
	_test_eof218:  m.cs = 218; goto _test_eof
	_test_eof219:  m.cs = 219; goto _test_eof
	_test_eof220:  m.cs = 220; goto _test_eof
	_test_eof221:  m.cs = 221; goto _test_eof
	_test_eof923:  m.cs = 923; goto _test_eof
	_test_eof222:  m.cs = 222; goto _test_eof
	_test_eof223:  m.cs = 223; goto _test_eof
	_test_eof224:  m.cs = 224; goto _test_eof
	_test_eof924:  m.cs = 924; goto _test_eof
	_test_eof925:  m.cs = 925; goto _test_eof
	_test_eof225:  m.cs = 225; goto _test_eof
	_test_eof226:  m.cs = 226; goto _test_eof
	_test_eof227:  m.cs = 227; goto _test_eof
	_test_eof228:  m.cs = 228; goto _test_eof
	_test_eof229:  m.cs = 229; goto _test_eof
	_test_eof926:  m.cs = 926; goto _test_eof
	_test_eof927:  m.cs = 927; goto _test_eof
	_test_eof928:  m.cs = 928; goto _test_eof
	_test_eof929:  m.cs = 929; goto _test_eof
	_test_eof930:  m.cs = 930; goto _test_eof
	_test_eof230:  m.cs = 230; goto _test_eof
	_test_eof231:  m.cs = 231; goto _test_eof
	_test_eof232:  m.cs = 232; goto _test_eof
	_test_eof931:  m.cs = 931; goto _test_eof
	_test_eof233:  m.cs = 233; goto _test_eof
	_test_eof234:  m.cs = 234; goto _test_eof
	_test_eof235:  m.cs = 235; goto _test_eof
	_test_eof932:  m.cs = 932; goto _test_eof
	_test_eof236:  m.cs = 236; goto _test_eof
	_test_eof237:  m.cs = 237; goto _test_eof
	_test_eof933:  m.cs = 933; goto _test_eof
	_test_eof934:  m.cs = 934; goto _test_eof
	_test_eof238:  m.cs = 238; goto _test_eof
	_test_eof239:  m.cs = 239; goto _test_eof
	_test_eof935:  m.cs = 935; goto _test_eof
	_test_eof936:  m.cs = 936; goto _test_eof
	_test_eof937:  m.cs = 937; goto _test_eof
	_test_eof938:  m.cs = 938; goto _test_eof
	_test_eof939:  m.cs = 939; goto _test_eof
	_test_eof240:  m.cs = 240; goto _test_eof
	_test_eof241:  m.cs = 241; goto _test_eof
	_test_eof242:  m.cs = 242; goto _test_eof
	_test_eof940:  m.cs = 940; goto _test_eof
	_test_eof243:  m.cs = 243; goto _test_eof
	_test_eof244:  m.cs = 244; goto _test_eof
	_test_eof245:  m.cs = 245; goto _test_eof
	_test_eof941:  m.cs = 941; goto _test_eof
	_test_eof246:  m.cs = 246; goto _test_eof
	_test_eof247:  m.cs = 247; goto _test_eof
	_test_eof942:  m.cs = 942; goto _test_eof
	_test_eof943:  m.cs = 943; goto _test_eof
	_test_eof248:  m.cs = 248; goto _test_eof
	_test_eof249:  m.cs = 249; goto _test_eof
	_test_eof944:  m.cs = 944; goto _test_eof
	_test_eof945:  m.cs = 945; goto _test_eof
	_test_eof946:  m.cs = 946; goto _test_eof
	_test_eof947:  m.cs = 947; goto _test_eof
	_test_eof948:  m.cs = 948; goto _test_eof
	_test_eof250:  m.cs = 250; goto _test_eof
	_test_eof251:  m.cs = 251; goto _test_eof
	_test_eof252:  m.cs = 252; goto _test_eof
	_test_eof949:  m.cs = 949; goto _test_eof
	_test_eof253:  m.cs = 253; goto _test_eof
	_test_eof254:  m.cs = 254; goto _test_eof
	_test_eof255:  m.cs = 255; goto _test_eof
	_test_eof950:  m.cs = 950; goto _test_eof
	_test_eof256:  m.cs = 256; goto _test_eof
	_test_eof257:  m.cs = 257; goto _test_eof
	_test_eof951:  m.cs = 951; goto _test_eof
	_test_eof952:  m.cs = 952; goto _test_eof
	_test_eof258:  m.cs = 258; goto _test_eof
	_test_eof259:  m.cs = 259; goto _test_eof
	_test_eof260:  m.cs = 260; goto _test_eof
	_test_eof261:  m.cs = 261; goto _test_eof
	_test_eof262:  m.cs = 262; goto _test_eof
	_test_eof263:  m.cs = 263; goto _test_eof
	_test_eof953:  m.cs = 953; goto _test_eof
	_test_eof954:  m.cs = 954; goto _test_eof
	_test_eof955:  m.cs = 955; goto _test_eof
	_test_eof956:  m.cs = 956; goto _test_eof
	_test_eof264:  m.cs = 264; goto _test_eof
	_test_eof957:  m.cs = 957; goto _test_eof
	_test_eof958:  m.cs = 958; goto _test_eof
	_test_eof959:  m.cs = 959; goto _test_eof
	_test_eof960:  m.cs = 960; goto _test_eof
	_test_eof961:  m.cs = 961; goto _test_eof
	_test_eof962:  m.cs = 962; goto _test_eof
	_test_eof963:  m.cs = 963; goto _test_eof
	_test_eof964:  m.cs = 964; goto _test_eof
	_test_eof965:  m.cs = 965; goto _test_eof
	_test_eof966:  m.cs = 966; goto _test_eof
	_test_eof967:  m.cs = 967; goto _test_eof
	_test_eof968:  m.cs = 968; goto _test_eof
	_test_eof969:  m.cs = 969; goto _test_eof
	_test_eof970:  m.cs = 970; goto _test_eof
	_test_eof971:  m.cs = 971; goto _test_eof
	_test_eof972:  m.cs = 972; goto _test_eof
	_test_eof973:  m.cs = 973; goto _test_eof
	_test_eof974:  m.cs = 974; goto _test_eof
	_test_eof975:  m.cs = 975; goto _test_eof
	_test_eof976:  m.cs = 976; goto _test_eof
	_test_eof265:  m.cs = 265; goto _test_eof
	_test_eof266:  m.cs = 266; goto _test_eof
	_test_eof267:  m.cs = 267; goto _test_eof
	_test_eof268:  m.cs = 268; goto _test_eof
	_test_eof269:  m.cs = 269; goto _test_eof
	_test_eof270:  m.cs = 270; goto _test_eof
	_test_eof271:  m.cs = 271; goto _test_eof
	_test_eof272:  m.cs = 272; goto _test_eof
	_test_eof977:  m.cs = 977; goto _test_eof
	_test_eof978:  m.cs = 978; goto _test_eof
	_test_eof273:  m.cs = 273; goto _test_eof
	_test_eof979:  m.cs = 979; goto _test_eof
	_test_eof980:  m.cs = 980; goto _test_eof
	_test_eof981:  m.cs = 981; goto _test_eof
	_test_eof982:  m.cs = 982; goto _test_eof
	_test_eof983:  m.cs = 983; goto _test_eof
	_test_eof984:  m.cs = 984; goto _test_eof
	_test_eof985:  m.cs = 985; goto _test_eof
	_test_eof986:  m.cs = 986; goto _test_eof
	_test_eof987:  m.cs = 987; goto _test_eof
	_test_eof988:  m.cs = 988; goto _test_eof
	_test_eof989:  m.cs = 989; goto _test_eof
	_test_eof990:  m.cs = 990; goto _test_eof
	_test_eof991:  m.cs = 991; goto _test_eof
	_test_eof992:  m.cs = 992; goto _test_eof
	_test_eof993:  m.cs = 993; goto _test_eof
	_test_eof994:  m.cs = 994; goto _test_eof
	_test_eof995:  m.cs = 995; goto _test_eof
	_test_eof996:  m.cs = 996; goto _test_eof
	_test_eof997:  m.cs = 997; goto _test_eof
	_test_eof274:  m.cs = 274; goto _test_eof
	_test_eof275:  m.cs = 275; goto _test_eof
	_test_eof998:  m.cs = 998; goto _test_eof
	_test_eof999:  m.cs = 999; goto _test_eof
	_test_eof276:  m.cs = 276; goto _test_eof
	_test_eof1000:  m.cs = 1000; goto _test_eof
	_test_eof1001:  m.cs = 1001; goto _test_eof
	_test_eof1002:  m.cs = 1002; goto _test_eof
	_test_eof1003:  m.cs = 1003; goto _test_eof
	_test_eof1004:  m.cs = 1004; goto _test_eof
	_test_eof1005:  m.cs = 1005; goto _test_eof
	_test_eof1006:  m.cs = 1006; goto _test_eof
	_test_eof1007:  m.cs = 1007; goto _test_eof
	_test_eof1008:  m.cs = 1008; goto _test_eof
	_test_eof1009:  m.cs = 1009; goto _test_eof
	_test_eof1010:  m.cs = 1010; goto _test_eof
	_test_eof1011:  m.cs = 1011; goto _test_eof
	_test_eof1012:  m.cs = 1012; goto _test_eof
	_test_eof1013:  m.cs = 1013; goto _test_eof
	_test_eof1014:  m.cs = 1014; goto _test_eof
	_test_eof1015:  m.cs = 1015; goto _test_eof
	_test_eof1016:  m.cs = 1016; goto _test_eof
	_test_eof1017:  m.cs = 1017; goto _test_eof
	_test_eof1018:  m.cs = 1018; goto _test_eof
	_test_eof277:  m.cs = 277; goto _test_eof
	_test_eof1019:  m.cs = 1019; goto _test_eof
	_test_eof1020:  m.cs = 1020; goto _test_eof
	_test_eof278:  m.cs = 278; goto _test_eof
	_test_eof279:  m.cs = 279; goto _test_eof
	_test_eof280:  m.cs = 280; goto _test_eof
	_test_eof281:  m.cs = 281; goto _test_eof
	_test_eof1021:  m.cs = 1021; goto _test_eof
	_test_eof1022:  m.cs = 1022; goto _test_eof
	_test_eof1023:  m.cs = 1023; goto _test_eof
	_test_eof1024:  m.cs = 1024; goto _test_eof
	_test_eof1025:  m.cs = 1025; goto _test_eof
	_test_eof282:  m.cs = 282; goto _test_eof
	_test_eof283:  m.cs = 283; goto _test_eof
	_test_eof284:  m.cs = 284; goto _test_eof
	_test_eof1026:  m.cs = 1026; goto _test_eof
	_test_eof285:  m.cs = 285; goto _test_eof
	_test_eof286:  m.cs = 286; goto _test_eof
	_test_eof287:  m.cs = 287; goto _test_eof
	_test_eof1027:  m.cs = 1027; goto _test_eof
	_test_eof288:  m.cs = 288; goto _test_eof
	_test_eof289:  m.cs = 289; goto _test_eof
	_test_eof1028:  m.cs = 1028; goto _test_eof
	_test_eof1029:  m.cs = 1029; goto _test_eof
	_test_eof290:  m.cs = 290; goto _test_eof
	_test_eof291:  m.cs = 291; goto _test_eof
	_test_eof1030:  m.cs = 1030; goto _test_eof
	_test_eof1031:  m.cs = 1031; goto _test_eof
	_test_eof1032:  m.cs = 1032; goto _test_eof
	_test_eof292:  m.cs = 292; goto _test_eof
	_test_eof1033:  m.cs = 1033; goto _test_eof
	_test_eof1034:  m.cs = 1034; goto _test_eof
	_test_eof293:  m.cs = 293; goto _test_eof
	_test_eof1035:  m.cs = 1035; goto _test_eof
	_test_eof1036:  m.cs = 1036; goto _test_eof
	_test_eof1037:  m.cs = 1037; goto _test_eof
	_test_eof1038:  m.cs = 1038; goto _test_eof
	_test_eof1039:  m.cs = 1039; goto _test_eof
	_test_eof1040:  m.cs = 1040; goto _test_eof
	_test_eof1041:  m.cs = 1041; goto _test_eof
	_test_eof1042:  m.cs = 1042; goto _test_eof
	_test_eof1043:  m.cs = 1043; goto _test_eof
	_test_eof1044:  m.cs = 1044; goto _test_eof
	_test_eof1045:  m.cs = 1045; goto _test_eof
	_test_eof1046:  m.cs = 1046; goto _test_eof
	_test_eof1047:  m.cs = 1047; goto _test_eof
	_test_eof1048:  m.cs = 1048; goto _test_eof
	_test_eof1049:  m.cs = 1049; goto _test_eof
	_test_eof1050:  m.cs = 1050; goto _test_eof
	_test_eof1051:  m.cs = 1051; goto _test_eof
	_test_eof1052:  m.cs = 1052; goto _test_eof
	_test_eof1053:  m.cs = 1053; goto _test_eof
	_test_eof1054:  m.cs = 1054; goto _test_eof
	_test_eof294:  m.cs = 294; goto _test_eof
	_test_eof295:  m.cs = 295; goto _test_eof
	_test_eof1055:  m.cs = 1055; goto _test_eof
	_test_eof1056:  m.cs = 1056; goto _test_eof
	_test_eof296:  m.cs = 296; goto _test_eof
	_test_eof1057:  m.cs = 1057; goto _test_eof
	_test_eof1058:  m.cs = 1058; goto _test_eof
	_test_eof1059:  m.cs = 1059; goto _test_eof
	_test_eof1060:  m.cs = 1060; goto _test_eof
	_test_eof1061:  m.cs = 1061; goto _test_eof
	_test_eof1062:  m.cs = 1062; goto _test_eof
	_test_eof1063:  m.cs = 1063; goto _test_eof
	_test_eof1064:  m.cs = 1064; goto _test_eof
	_test_eof1065:  m.cs = 1065; goto _test_eof
	_test_eof1066:  m.cs = 1066; goto _test_eof
	_test_eof1067:  m.cs = 1067; goto _test_eof
	_test_eof1068:  m.cs = 1068; goto _test_eof
	_test_eof1069:  m.cs = 1069; goto _test_eof
	_test_eof1070:  m.cs = 1070; goto _test_eof
	_test_eof1071:  m.cs = 1071; goto _test_eof
	_test_eof1072:  m.cs = 1072; goto _test_eof
	_test_eof1073:  m.cs = 1073; goto _test_eof
	_test_eof1074:  m.cs = 1074; goto _test_eof
	_test_eof1075:  m.cs = 1075; goto _test_eof
	_test_eof1076:  m.cs = 1076; goto _test_eof
	_test_eof297:  m.cs = 297; goto _test_eof
	_test_eof1077:  m.cs = 1077; goto _test_eof
	_test_eof1078:  m.cs = 1078; goto _test_eof
	_test_eof1079:  m.cs = 1079; goto _test_eof
	_test_eof298:  m.cs = 298; goto _test_eof
	_test_eof1080:  m.cs = 1080; goto _test_eof
	_test_eof1081:  m.cs = 1081; goto _test_eof
	_test_eof1082:  m.cs = 1082; goto _test_eof
	_test_eof1083:  m.cs = 1083; goto _test_eof
	_test_eof1084:  m.cs = 1084; goto _test_eof
	_test_eof1085:  m.cs = 1085; goto _test_eof
	_test_eof1086:  m.cs = 1086; goto _test_eof
	_test_eof1087:  m.cs = 1087; goto _test_eof
	_test_eof1088:  m.cs = 1088; goto _test_eof
	_test_eof1089:  m.cs = 1089; goto _test_eof
	_test_eof1090:  m.cs = 1090; goto _test_eof
	_test_eof1091:  m.cs = 1091; goto _test_eof
	_test_eof1092:  m.cs = 1092; goto _test_eof
	_test_eof1093:  m.cs = 1093; goto _test_eof
	_test_eof1094:  m.cs = 1094; goto _test_eof
	_test_eof1095:  m.cs = 1095; goto _test_eof
	_test_eof1096:  m.cs = 1096; goto _test_eof
	_test_eof1097:  m.cs = 1097; goto _test_eof
	_test_eof1098:  m.cs = 1098; goto _test_eof
	_test_eof1099:  m.cs = 1099; goto _test_eof
	_test_eof1100:  m.cs = 1100; goto _test_eof
	_test_eof1101:  m.cs = 1101; goto _test_eof
	_test_eof1102:  m.cs = 1102; goto _test_eof
	_test_eof299:  m.cs = 299; goto _test_eof
	_test_eof1103:  m.cs = 1103; goto _test_eof
	_test_eof1104:  m.cs = 1104; goto _test_eof
	_test_eof1105:  m.cs = 1105; goto _test_eof
	_test_eof1106:  m.cs = 1106; goto _test_eof
	_test_eof1107:  m.cs = 1107; goto _test_eof
	_test_eof1108:  m.cs = 1108; goto _test_eof
	_test_eof1109:  m.cs = 1109; goto _test_eof
	_test_eof1110:  m.cs = 1110; goto _test_eof
	_test_eof1111:  m.cs = 1111; goto _test_eof
	_test_eof1112:  m.cs = 1112; goto _test_eof
	_test_eof1113:  m.cs = 1113; goto _test_eof
	_test_eof1114:  m.cs = 1114; goto _test_eof
	_test_eof1115:  m.cs = 1115; goto _test_eof
	_test_eof1116:  m.cs = 1116; goto _test_eof
	_test_eof1117:  m.cs = 1117; goto _test_eof
	_test_eof1118:  m.cs = 1118; goto _test_eof
	_test_eof1119:  m.cs = 1119; goto _test_eof
	_test_eof1120:  m.cs = 1120; goto _test_eof
	_test_eof1121:  m.cs = 1121; goto _test_eof
	_test_eof1122:  m.cs = 1122; goto _test_eof
	_test_eof1123:  m.cs = 1123; goto _test_eof
	_test_eof1124:  m.cs = 1124; goto _test_eof
	_test_eof300:  m.cs = 300; goto _test_eof
	_test_eof301:  m.cs = 301; goto _test_eof
	_test_eof1125:  m.cs = 1125; goto _test_eof
	_test_eof1126:  m.cs = 1126; goto _test_eof
	_test_eof302:  m.cs = 302; goto _test_eof
	_test_eof1127:  m.cs = 1127; goto _test_eof
	_test_eof1128:  m.cs = 1128; goto _test_eof
	_test_eof1129:  m.cs = 1129; goto _test_eof
	_test_eof1130:  m.cs = 1130; goto _test_eof
	_test_eof1131:  m.cs = 1131; goto _test_eof
	_test_eof1132:  m.cs = 1132; goto _test_eof
	_test_eof1133:  m.cs = 1133; goto _test_eof
	_test_eof1134:  m.cs = 1134; goto _test_eof
	_test_eof1135:  m.cs = 1135; goto _test_eof
	_test_eof1136:  m.cs = 1136; goto _test_eof
	_test_eof1137:  m.cs = 1137; goto _test_eof
	_test_eof1138:  m.cs = 1138; goto _test_eof
	_test_eof1139:  m.cs = 1139; goto _test_eof
	_test_eof1140:  m.cs = 1140; goto _test_eof
	_test_eof1141:  m.cs = 1141; goto _test_eof
	_test_eof1142:  m.cs = 1142; goto _test_eof
	_test_eof1143:  m.cs = 1143; goto _test_eof
	_test_eof1144:  m.cs = 1144; goto _test_eof
	_test_eof1145:  m.cs = 1145; goto _test_eof
	_test_eof303:  m.cs = 303; goto _test_eof
	_test_eof1146:  m.cs = 1146; goto _test_eof
	_test_eof1147:  m.cs = 1147; goto _test_eof
	_test_eof1148:  m.cs = 1148; goto _test_eof
	_test_eof304:  m.cs = 304; goto _test_eof
	_test_eof1149:  m.cs = 1149; goto _test_eof
	_test_eof1150:  m.cs = 1150; goto _test_eof
	_test_eof305:  m.cs = 305; goto _test_eof
	_test_eof1151:  m.cs = 1151; goto _test_eof
	_test_eof1152:  m.cs = 1152; goto _test_eof
	_test_eof1153:  m.cs = 1153; goto _test_eof
	_test_eof1154:  m.cs = 1154; goto _test_eof
	_test_eof1155:  m.cs = 1155; goto _test_eof
	_test_eof1156:  m.cs = 1156; goto _test_eof
	_test_eof1157:  m.cs = 1157; goto _test_eof
	_test_eof1158:  m.cs = 1158; goto _test_eof
	_test_eof1159:  m.cs = 1159; goto _test_eof
	_test_eof1160:  m.cs = 1160; goto _test_eof
	_test_eof1161:  m.cs = 1161; goto _test_eof
	_test_eof1162:  m.cs = 1162; goto _test_eof
	_test_eof1163:  m.cs = 1163; goto _test_eof
	_test_eof1164:  m.cs = 1164; goto _test_eof
	_test_eof1165:  m.cs = 1165; goto _test_eof
	_test_eof1166:  m.cs = 1166; goto _test_eof
	_test_eof1167:  m.cs = 1167; goto _test_eof
	_test_eof1168:  m.cs = 1168; goto _test_eof
	_test_eof1169:  m.cs = 1169; goto _test_eof
	_test_eof306:  m.cs = 306; goto _test_eof
	_test_eof307:  m.cs = 307; goto _test_eof
	_test_eof308:  m.cs = 308; goto _test_eof
	_test_eof309:  m.cs = 309; goto _test_eof
	_test_eof1170:  m.cs = 1170; goto _test_eof
	_test_eof1171:  m.cs = 1171; goto _test_eof
	_test_eof1172:  m.cs = 1172; goto _test_eof
	_test_eof1173:  m.cs = 1173; goto _test_eof
	_test_eof1174:  m.cs = 1174; goto _test_eof
	_test_eof310:  m.cs = 310; goto _test_eof
	_test_eof311:  m.cs = 311; goto _test_eof
	_test_eof312:  m.cs = 312; goto _test_eof
	_test_eof1175:  m.cs = 1175; goto _test_eof
	_test_eof313:  m.cs = 313; goto _test_eof
	_test_eof314:  m.cs = 314; goto _test_eof
	_test_eof315:  m.cs = 315; goto _test_eof
	_test_eof1176:  m.cs = 1176; goto _test_eof
	_test_eof316:  m.cs = 316; goto _test_eof
	_test_eof317:  m.cs = 317; goto _test_eof
	_test_eof1177:  m.cs = 1177; goto _test_eof
	_test_eof1178:  m.cs = 1178; goto _test_eof
	_test_eof318:  m.cs = 318; goto _test_eof
	_test_eof319:  m.cs = 319; goto _test_eof
	_test_eof1179:  m.cs = 1179; goto _test_eof
	_test_eof1180:  m.cs = 1180; goto _test_eof
	_test_eof1181:  m.cs = 1181; goto _test_eof
	_test_eof1182:  m.cs = 1182; goto _test_eof
	_test_eof1183:  m.cs = 1183; goto _test_eof
	_test_eof320:  m.cs = 320; goto _test_eof
	_test_eof321:  m.cs = 321; goto _test_eof
	_test_eof322:  m.cs = 322; goto _test_eof
	_test_eof1184:  m.cs = 1184; goto _test_eof
	_test_eof323:  m.cs = 323; goto _test_eof
	_test_eof324:  m.cs = 324; goto _test_eof
	_test_eof325:  m.cs = 325; goto _test_eof
	_test_eof1185:  m.cs = 1185; goto _test_eof
	_test_eof326:  m.cs = 326; goto _test_eof
	_test_eof327:  m.cs = 327; goto _test_eof
	_test_eof1186:  m.cs = 1186; goto _test_eof
	_test_eof1187:  m.cs = 1187; goto _test_eof
	_test_eof328:  m.cs = 328; goto _test_eof
	_test_eof329:  m.cs = 329; goto _test_eof
	_test_eof1188:  m.cs = 1188; goto _test_eof
	_test_eof1189:  m.cs = 1189; goto _test_eof
	_test_eof1190:  m.cs = 1190; goto _test_eof
	_test_eof330:  m.cs = 330; goto _test_eof
	_test_eof1191:  m.cs = 1191; goto _test_eof
	_test_eof1192:  m.cs = 1192; goto _test_eof
	_test_eof1193:  m.cs = 1193; goto _test_eof
	_test_eof1194:  m.cs = 1194; goto _test_eof
	_test_eof1195:  m.cs = 1195; goto _test_eof
	_test_eof1196:  m.cs = 1196; goto _test_eof
	_test_eof1197:  m.cs = 1197; goto _test_eof
	_test_eof1198:  m.cs = 1198; goto _test_eof
	_test_eof1199:  m.cs = 1199; goto _test_eof
	_test_eof1200:  m.cs = 1200; goto _test_eof
	_test_eof1201:  m.cs = 1201; goto _test_eof
	_test_eof1202:  m.cs = 1202; goto _test_eof
	_test_eof1203:  m.cs = 1203; goto _test_eof
	_test_eof1204:  m.cs = 1204; goto _test_eof
	_test_eof1205:  m.cs = 1205; goto _test_eof
	_test_eof1206:  m.cs = 1206; goto _test_eof
	_test_eof1207:  m.cs = 1207; goto _test_eof
	_test_eof1208:  m.cs = 1208; goto _test_eof
	_test_eof1209:  m.cs = 1209; goto _test_eof
	_test_eof1210:  m.cs = 1210; goto _test_eof
	_test_eof1211:  m.cs = 1211; goto _test_eof
	_test_eof1212:  m.cs = 1212; goto _test_eof
	_test_eof1213:  m.cs = 1213; goto _test_eof
	_test_eof331:  m.cs = 331; goto _test_eof
	_test_eof332:  m.cs = 332; goto _test_eof
	_test_eof333:  m.cs = 333; goto _test_eof
	_test_eof1214:  m.cs = 1214; goto _test_eof
	_test_eof334:  m.cs = 334; goto _test_eof
	_test_eof335:  m.cs = 335; goto _test_eof
	_test_eof336:  m.cs = 336; goto _test_eof
	_test_eof1215:  m.cs = 1215; goto _test_eof
	_test_eof337:  m.cs = 337; goto _test_eof
	_test_eof338:  m.cs = 338; goto _test_eof
	_test_eof1216:  m.cs = 1216; goto _test_eof
	_test_eof1217:  m.cs = 1217; goto _test_eof
	_test_eof339:  m.cs = 339; goto _test_eof
	_test_eof340:  m.cs = 340; goto _test_eof
	_test_eof1218:  m.cs = 1218; goto _test_eof
	_test_eof1219:  m.cs = 1219; goto _test_eof
	_test_eof1220:  m.cs = 1220; goto _test_eof
	_test_eof1221:  m.cs = 1221; goto _test_eof
	_test_eof1222:  m.cs = 1222; goto _test_eof
	_test_eof341:  m.cs = 341; goto _test_eof
	_test_eof1223:  m.cs = 1223; goto _test_eof
	_test_eof1224:  m.cs = 1224; goto _test_eof
	_test_eof342:  m.cs = 342; goto _test_eof
	_test_eof1225:  m.cs = 1225; goto _test_eof
	_test_eof1226:  m.cs = 1226; goto _test_eof
	_test_eof1227:  m.cs = 1227; goto _test_eof
	_test_eof1228:  m.cs = 1228; goto _test_eof
	_test_eof1229:  m.cs = 1229; goto _test_eof
	_test_eof1230:  m.cs = 1230; goto _test_eof
	_test_eof1231:  m.cs = 1231; goto _test_eof
	_test_eof1232:  m.cs = 1232; goto _test_eof
	_test_eof1233:  m.cs = 1233; goto _test_eof
	_test_eof1234:  m.cs = 1234; goto _test_eof
	_test_eof1235:  m.cs = 1235; goto _test_eof
	_test_eof1236:  m.cs = 1236; goto _test_eof
	_test_eof1237:  m.cs = 1237; goto _test_eof
	_test_eof1238:  m.cs = 1238; goto _test_eof
	_test_eof1239:  m.cs = 1239; goto _test_eof
	_test_eof1240:  m.cs = 1240; goto _test_eof
	_test_eof1241:  m.cs = 1241; goto _test_eof
	_test_eof1242:  m.cs = 1242; goto _test_eof
	_test_eof1243:  m.cs = 1243; goto _test_eof
	_test_eof343:  m.cs = 343; goto _test_eof
	_test_eof1244:  m.cs = 1244; goto _test_eof
	_test_eof1245:  m.cs = 1245; goto _test_eof
	_test_eof1246:  m.cs = 1246; goto _test_eof
	_test_eof344:  m.cs = 344; goto _test_eof
	_test_eof345:  m.cs = 345; goto _test_eof
	_test_eof346:  m.cs = 346; goto _test_eof
	_test_eof1247:  m.cs = 1247; goto _test_eof
	_test_eof347:  m.cs = 347; goto _test_eof
	_test_eof348:  m.cs = 348; goto _test_eof
	_test_eof349:  m.cs = 349; goto _test_eof
	_test_eof1248:  m.cs = 1248; goto _test_eof
	_test_eof350:  m.cs = 350; goto _test_eof
	_test_eof351:  m.cs = 351; goto _test_eof
	_test_eof1249:  m.cs = 1249; goto _test_eof
	_test_eof1250:  m.cs = 1250; goto _test_eof
	_test_eof352:  m.cs = 352; goto _test_eof
	_test_eof353:  m.cs = 353; goto _test_eof
	_test_eof354:  m.cs = 354; goto _test_eof
	_test_eof355:  m.cs = 355; goto _test_eof
	_test_eof356:  m.cs = 356; goto _test_eof
	_test_eof357:  m.cs = 357; goto _test_eof
	_test_eof1251:  m.cs = 1251; goto _test_eof
	_test_eof1252:  m.cs = 1252; goto _test_eof
	_test_eof1253:  m.cs = 1253; goto _test_eof
	_test_eof1254:  m.cs = 1254; goto _test_eof
	_test_eof1255:  m.cs = 1255; goto _test_eof
	_test_eof1256:  m.cs = 1256; goto _test_eof
	_test_eof1257:  m.cs = 1257; goto _test_eof
	_test_eof1258:  m.cs = 1258; goto _test_eof
	_test_eof1259:  m.cs = 1259; goto _test_eof
	_test_eof1260:  m.cs = 1260; goto _test_eof
	_test_eof1261:  m.cs = 1261; goto _test_eof
	_test_eof1262:  m.cs = 1262; goto _test_eof
	_test_eof1263:  m.cs = 1263; goto _test_eof
	_test_eof1264:  m.cs = 1264; goto _test_eof
	_test_eof1265:  m.cs = 1265; goto _test_eof
	_test_eof1266:  m.cs = 1266; goto _test_eof
	_test_eof1267:  m.cs = 1267; goto _test_eof
	_test_eof1268:  m.cs = 1268; goto _test_eof
	_test_eof1269:  m.cs = 1269; goto _test_eof
	_test_eof358:  m.cs = 358; goto _test_eof
	_test_eof359:  m.cs = 359; goto _test_eof
	_test_eof360:  m.cs = 360; goto _test_eof
	_test_eof1270:  m.cs = 1270; goto _test_eof
	_test_eof361:  m.cs = 361; goto _test_eof
	_test_eof362:  m.cs = 362; goto _test_eof
	_test_eof363:  m.cs = 363; goto _test_eof
	_test_eof364:  m.cs = 364; goto _test_eof
	_test_eof1271:  m.cs = 1271; goto _test_eof
	_test_eof1272:  m.cs = 1272; goto _test_eof
	_test_eof1273:  m.cs = 1273; goto _test_eof
	_test_eof1274:  m.cs = 1274; goto _test_eof
	_test_eof1275:  m.cs = 1275; goto _test_eof
	_test_eof365:  m.cs = 365; goto _test_eof
	_test_eof366:  m.cs = 366; goto _test_eof
	_test_eof367:  m.cs = 367; goto _test_eof
	_test_eof1276:  m.cs = 1276; goto _test_eof
	_test_eof368:  m.cs = 368; goto _test_eof
	_test_eof369:  m.cs = 369; goto _test_eof
	_test_eof370:  m.cs = 370; goto _test_eof
	_test_eof1277:  m.cs = 1277; goto _test_eof
	_test_eof371:  m.cs = 371; goto _test_eof
	_test_eof372:  m.cs = 372; goto _test_eof
	_test_eof1278:  m.cs = 1278; goto _test_eof
	_test_eof1279:  m.cs = 1279; goto _test_eof
	_test_eof373:  m.cs = 373; goto _test_eof
	_test_eof374:  m.cs = 374; goto _test_eof
	_test_eof1280:  m.cs = 1280; goto _test_eof
	_test_eof1281:  m.cs = 1281; goto _test_eof
	_test_eof1282:  m.cs = 1282; goto _test_eof
	_test_eof1283:  m.cs = 1283; goto _test_eof
	_test_eof1284:  m.cs = 1284; goto _test_eof
	_test_eof375:  m.cs = 375; goto _test_eof
	_test_eof376:  m.cs = 376; goto _test_eof
	_test_eof377:  m.cs = 377; goto _test_eof
	_test_eof1285:  m.cs = 1285; goto _test_eof
	_test_eof378:  m.cs = 378; goto _test_eof
	_test_eof379:  m.cs = 379; goto _test_eof
	_test_eof380:  m.cs = 380; goto _test_eof
	_test_eof1286:  m.cs = 1286; goto _test_eof
	_test_eof381:  m.cs = 381; goto _test_eof
	_test_eof382:  m.cs = 382; goto _test_eof
	_test_eof1287:  m.cs = 1287; goto _test_eof
	_test_eof1288:  m.cs = 1288; goto _test_eof
	_test_eof383:  m.cs = 383; goto _test_eof
	_test_eof384:  m.cs = 384; goto _test_eof
	_test_eof385:  m.cs = 385; goto _test_eof
	_test_eof386:  m.cs = 386; goto _test_eof
	_test_eof387:  m.cs = 387; goto _test_eof
	_test_eof388:  m.cs = 388; goto _test_eof
	_test_eof1289:  m.cs = 1289; goto _test_eof
	_test_eof1290:  m.cs = 1290; goto _test_eof
	_test_eof389:  m.cs = 389; goto _test_eof
	_test_eof1291:  m.cs = 1291; goto _test_eof
	_test_eof1292:  m.cs = 1292; goto _test_eof
	_test_eof390:  m.cs = 390; goto _test_eof
	_test_eof1293:  m.cs = 1293; goto _test_eof
	_test_eof1294:  m.cs = 1294; goto _test_eof
	_test_eof1295:  m.cs = 1295; goto _test_eof
	_test_eof1296:  m.cs = 1296; goto _test_eof
	_test_eof1297:  m.cs = 1297; goto _test_eof
	_test_eof1298:  m.cs = 1298; goto _test_eof
	_test_eof1299:  m.cs = 1299; goto _test_eof
	_test_eof1300:  m.cs = 1300; goto _test_eof
	_test_eof1301:  m.cs = 1301; goto _test_eof
	_test_eof1302:  m.cs = 1302; goto _test_eof
	_test_eof1303:  m.cs = 1303; goto _test_eof
	_test_eof1304:  m.cs = 1304; goto _test_eof
	_test_eof1305:  m.cs = 1305; goto _test_eof
	_test_eof1306:  m.cs = 1306; goto _test_eof
	_test_eof1307:  m.cs = 1307; goto _test_eof
	_test_eof1308:  m.cs = 1308; goto _test_eof
	_test_eof1309:  m.cs = 1309; goto _test_eof
	_test_eof1310:  m.cs = 1310; goto _test_eof
	_test_eof391:  m.cs = 391; goto _test_eof
	_test_eof392:  m.cs = 392; goto _test_eof
	_test_eof393:  m.cs = 393; goto _test_eof
	_test_eof394:  m.cs = 394; goto _test_eof
	_test_eof395:  m.cs = 395; goto _test_eof
	_test_eof396:  m.cs = 396; goto _test_eof
	_test_eof397:  m.cs = 397; goto _test_eof
	_test_eof398:  m.cs = 398; goto _test_eof
	_test_eof399:  m.cs = 399; goto _test_eof
	_test_eof400:  m.cs = 400; goto _test_eof
	_test_eof401:  m.cs = 401; goto _test_eof
	_test_eof402:  m.cs = 402; goto _test_eof
	_test_eof403:  m.cs = 403; goto _test_eof
	_test_eof1311:  m.cs = 1311; goto _test_eof
	_test_eof404:  m.cs = 404; goto _test_eof
	_test_eof405:  m.cs = 405; goto _test_eof
	_test_eof406:  m.cs = 406; goto _test_eof
	_test_eof1312:  m.cs = 1312; goto _test_eof
	_test_eof407:  m.cs = 407; goto _test_eof
	_test_eof408:  m.cs = 408; goto _test_eof
	_test_eof409:  m.cs = 409; goto _test_eof
	_test_eof410:  m.cs = 410; goto _test_eof
	_test_eof1313:  m.cs = 1313; goto _test_eof
	_test_eof1314:  m.cs = 1314; goto _test_eof
	_test_eof1315:  m.cs = 1315; goto _test_eof
	_test_eof1316:  m.cs = 1316; goto _test_eof
	_test_eof1317:  m.cs = 1317; goto _test_eof
	_test_eof411:  m.cs = 411; goto _test_eof
	_test_eof412:  m.cs = 412; goto _test_eof
	_test_eof413:  m.cs = 413; goto _test_eof
	_test_eof1318:  m.cs = 1318; goto _test_eof
	_test_eof414:  m.cs = 414; goto _test_eof
	_test_eof415:  m.cs = 415; goto _test_eof
	_test_eof416:  m.cs = 416; goto _test_eof
	_test_eof1319:  m.cs = 1319; goto _test_eof
	_test_eof417:  m.cs = 417; goto _test_eof
	_test_eof418:  m.cs = 418; goto _test_eof
	_test_eof1320:  m.cs = 1320; goto _test_eof
	_test_eof1321:  m.cs = 1321; goto _test_eof
	_test_eof419:  m.cs = 419; goto _test_eof
	_test_eof420:  m.cs = 420; goto _test_eof
	_test_eof421:  m.cs = 421; goto _test_eof
	_test_eof422:  m.cs = 422; goto _test_eof
	_test_eof423:  m.cs = 423; goto _test_eof
	_test_eof424:  m.cs = 424; goto _test_eof
	_test_eof425:  m.cs = 425; goto _test_eof
	_test_eof426:  m.cs = 426; goto _test_eof
	_test_eof427:  m.cs = 427; goto _test_eof
	_test_eof428:  m.cs = 428; goto _test_eof
	_test_eof429:  m.cs = 429; goto _test_eof
	_test_eof430:  m.cs = 430; goto _test_eof
	_test_eof1322:  m.cs = 1322; goto _test_eof
	_test_eof1323:  m.cs = 1323; goto _test_eof
	_test_eof1324:  m.cs = 1324; goto _test_eof
	_test_eof431:  m.cs = 431; goto _test_eof
	_test_eof1325:  m.cs = 1325; goto _test_eof
	_test_eof1326:  m.cs = 1326; goto _test_eof
	_test_eof1327:  m.cs = 1327; goto _test_eof
	_test_eof1328:  m.cs = 1328; goto _test_eof
	_test_eof1329:  m.cs = 1329; goto _test_eof
	_test_eof1330:  m.cs = 1330; goto _test_eof
	_test_eof1331:  m.cs = 1331; goto _test_eof
	_test_eof1332:  m.cs = 1332; goto _test_eof
	_test_eof1333:  m.cs = 1333; goto _test_eof
	_test_eof1334:  m.cs = 1334; goto _test_eof
	_test_eof1335:  m.cs = 1335; goto _test_eof
	_test_eof1336:  m.cs = 1336; goto _test_eof
	_test_eof1337:  m.cs = 1337; goto _test_eof
	_test_eof1338:  m.cs = 1338; goto _test_eof
	_test_eof1339:  m.cs = 1339; goto _test_eof
	_test_eof1340:  m.cs = 1340; goto _test_eof
	_test_eof1341:  m.cs = 1341; goto _test_eof
	_test_eof1342:  m.cs = 1342; goto _test_eof
	_test_eof1343:  m.cs = 1343; goto _test_eof
	_test_eof1344:  m.cs = 1344; goto _test_eof
	_test_eof1345:  m.cs = 1345; goto _test_eof
	_test_eof1346:  m.cs = 1346; goto _test_eof
	_test_eof432:  m.cs = 432; goto _test_eof
	_test_eof1347:  m.cs = 1347; goto _test_eof
	_test_eof1348:  m.cs = 1348; goto _test_eof
	_test_eof1349:  m.cs = 1349; goto _test_eof
	_test_eof1350:  m.cs = 1350; goto _test_eof
	_test_eof433:  m.cs = 433; goto _test_eof
	_test_eof1351:  m.cs = 1351; goto _test_eof
	_test_eof1352:  m.cs = 1352; goto _test_eof
	_test_eof1353:  m.cs = 1353; goto _test_eof
	_test_eof1354:  m.cs = 1354; goto _test_eof
	_test_eof1355:  m.cs = 1355; goto _test_eof
	_test_eof1356:  m.cs = 1356; goto _test_eof
	_test_eof1357:  m.cs = 1357; goto _test_eof
	_test_eof1358:  m.cs = 1358; goto _test_eof
	_test_eof1359:  m.cs = 1359; goto _test_eof
	_test_eof1360:  m.cs = 1360; goto _test_eof
	_test_eof1361:  m.cs = 1361; goto _test_eof
	_test_eof1362:  m.cs = 1362; goto _test_eof
	_test_eof1363:  m.cs = 1363; goto _test_eof
	_test_eof1364:  m.cs = 1364; goto _test_eof
	_test_eof1365:  m.cs = 1365; goto _test_eof
	_test_eof1366:  m.cs = 1366; goto _test_eof
	_test_eof1367:  m.cs = 1367; goto _test_eof
	_test_eof1368:  m.cs = 1368; goto _test_eof
	_test_eof1369:  m.cs = 1369; goto _test_eof
	_test_eof1370:  m.cs = 1370; goto _test_eof
	_test_eof434:  m.cs = 434; goto _test_eof
	_test_eof435:  m.cs = 435; goto _test_eof
	_test_eof436:  m.cs = 436; goto _test_eof
	_test_eof1371:  m.cs = 1371; goto _test_eof
	_test_eof1372:  m.cs = 1372; goto _test_eof
	_test_eof437:  m.cs = 437; goto _test_eof
	_test_eof1373:  m.cs = 1373; goto _test_eof
	_test_eof1374:  m.cs = 1374; goto _test_eof
	_test_eof1375:  m.cs = 1375; goto _test_eof
	_test_eof1376:  m.cs = 1376; goto _test_eof
	_test_eof1377:  m.cs = 1377; goto _test_eof
	_test_eof1378:  m.cs = 1378; goto _test_eof
	_test_eof1379:  m.cs = 1379; goto _test_eof
	_test_eof1380:  m.cs = 1380; goto _test_eof
	_test_eof1381:  m.cs = 1381; goto _test_eof
	_test_eof1382:  m.cs = 1382; goto _test_eof
	_test_eof1383:  m.cs = 1383; goto _test_eof
	_test_eof1384:  m.cs = 1384; goto _test_eof
	_test_eof1385:  m.cs = 1385; goto _test_eof
	_test_eof1386:  m.cs = 1386; goto _test_eof
	_test_eof1387:  m.cs = 1387; goto _test_eof
	_test_eof1388:  m.cs = 1388; goto _test_eof
	_test_eof1389:  m.cs = 1389; goto _test_eof
	_test_eof1390:  m.cs = 1390; goto _test_eof
	_test_eof1391:  m.cs = 1391; goto _test_eof
	_test_eof1392:  m.cs = 1392; goto _test_eof
	_test_eof438:  m.cs = 438; goto _test_eof
	_test_eof1393:  m.cs = 1393; goto _test_eof
	_test_eof1394:  m.cs = 1394; goto _test_eof
	_test_eof1395:  m.cs = 1395; goto _test_eof
	_test_eof439:  m.cs = 439; goto _test_eof
	_test_eof1396:  m.cs = 1396; goto _test_eof
	_test_eof1397:  m.cs = 1397; goto _test_eof
	_test_eof1398:  m.cs = 1398; goto _test_eof
	_test_eof1399:  m.cs = 1399; goto _test_eof
	_test_eof1400:  m.cs = 1400; goto _test_eof
	_test_eof1401:  m.cs = 1401; goto _test_eof
	_test_eof1402:  m.cs = 1402; goto _test_eof
	_test_eof1403:  m.cs = 1403; goto _test_eof
	_test_eof1404:  m.cs = 1404; goto _test_eof
	_test_eof1405:  m.cs = 1405; goto _test_eof
	_test_eof1406:  m.cs = 1406; goto _test_eof
	_test_eof1407:  m.cs = 1407; goto _test_eof
	_test_eof1408:  m.cs = 1408; goto _test_eof
	_test_eof1409:  m.cs = 1409; goto _test_eof
	_test_eof1410:  m.cs = 1410; goto _test_eof
	_test_eof1411:  m.cs = 1411; goto _test_eof
	_test_eof1412:  m.cs = 1412; goto _test_eof
	_test_eof1413:  m.cs = 1413; goto _test_eof
	_test_eof1414:  m.cs = 1414; goto _test_eof
	_test_eof1415:  m.cs = 1415; goto _test_eof
	_test_eof1416:  m.cs = 1416; goto _test_eof
	_test_eof1417:  m.cs = 1417; goto _test_eof
	_test_eof440:  m.cs = 440; goto _test_eof
	_test_eof1418:  m.cs = 1418; goto _test_eof
	_test_eof1419:  m.cs = 1419; goto _test_eof
	_test_eof1420:  m.cs = 1420; goto _test_eof
	_test_eof1421:  m.cs = 1421; goto _test_eof
	_test_eof1422:  m.cs = 1422; goto _test_eof
	_test_eof1423:  m.cs = 1423; goto _test_eof
	_test_eof1424:  m.cs = 1424; goto _test_eof
	_test_eof1425:  m.cs = 1425; goto _test_eof
	_test_eof1426:  m.cs = 1426; goto _test_eof
	_test_eof1427:  m.cs = 1427; goto _test_eof
	_test_eof1428:  m.cs = 1428; goto _test_eof
	_test_eof1429:  m.cs = 1429; goto _test_eof
	_test_eof1430:  m.cs = 1430; goto _test_eof
	_test_eof1431:  m.cs = 1431; goto _test_eof
	_test_eof1432:  m.cs = 1432; goto _test_eof
	_test_eof1433:  m.cs = 1433; goto _test_eof
	_test_eof1434:  m.cs = 1434; goto _test_eof
	_test_eof1435:  m.cs = 1435; goto _test_eof
	_test_eof1436:  m.cs = 1436; goto _test_eof
	_test_eof1437:  m.cs = 1437; goto _test_eof
	_test_eof1438:  m.cs = 1438; goto _test_eof
	_test_eof1439:  m.cs = 1439; goto _test_eof
	_test_eof441:  m.cs = 441; goto _test_eof
	_test_eof442:  m.cs = 442; goto _test_eof
	_test_eof443:  m.cs = 443; goto _test_eof
	_test_eof444:  m.cs = 444; goto _test_eof
	_test_eof1440:  m.cs = 1440; goto _test_eof
	_test_eof1441:  m.cs = 1441; goto _test_eof
	_test_eof1442:  m.cs = 1442; goto _test_eof
	_test_eof1443:  m.cs = 1443; goto _test_eof
	_test_eof1444:  m.cs = 1444; goto _test_eof
	_test_eof445:  m.cs = 445; goto _test_eof
	_test_eof446:  m.cs = 446; goto _test_eof
	_test_eof447:  m.cs = 447; goto _test_eof
	_test_eof1445:  m.cs = 1445; goto _test_eof
	_test_eof448:  m.cs = 448; goto _test_eof
	_test_eof449:  m.cs = 449; goto _test_eof
	_test_eof450:  m.cs = 450; goto _test_eof
	_test_eof1446:  m.cs = 1446; goto _test_eof
	_test_eof451:  m.cs = 451; goto _test_eof
	_test_eof452:  m.cs = 452; goto _test_eof
	_test_eof1447:  m.cs = 1447; goto _test_eof
	_test_eof1448:  m.cs = 1448; goto _test_eof
	_test_eof453:  m.cs = 453; goto _test_eof
	_test_eof454:  m.cs = 454; goto _test_eof
	_test_eof1449:  m.cs = 1449; goto _test_eof
	_test_eof1450:  m.cs = 1450; goto _test_eof
	_test_eof1451:  m.cs = 1451; goto _test_eof
	_test_eof455:  m.cs = 455; goto _test_eof
	_test_eof1452:  m.cs = 1452; goto _test_eof
	_test_eof1453:  m.cs = 1453; goto _test_eof
	_test_eof1454:  m.cs = 1454; goto _test_eof
	_test_eof1455:  m.cs = 1455; goto _test_eof
	_test_eof1456:  m.cs = 1456; goto _test_eof
	_test_eof1457:  m.cs = 1457; goto _test_eof
	_test_eof1458:  m.cs = 1458; goto _test_eof
	_test_eof1459:  m.cs = 1459; goto _test_eof
	_test_eof1460:  m.cs = 1460; goto _test_eof
	_test_eof1461:  m.cs = 1461; goto _test_eof
	_test_eof1462:  m.cs = 1462; goto _test_eof
	_test_eof1463:  m.cs = 1463; goto _test_eof
	_test_eof1464:  m.cs = 1464; goto _test_eof
	_test_eof1465:  m.cs = 1465; goto _test_eof
	_test_eof1466:  m.cs = 1466; goto _test_eof
	_test_eof1467:  m.cs = 1467; goto _test_eof
	_test_eof1468:  m.cs = 1468; goto _test_eof
	_test_eof1469:  m.cs = 1469; goto _test_eof
	_test_eof1470:  m.cs = 1470; goto _test_eof
	_test_eof1471:  m.cs = 1471; goto _test_eof
	_test_eof1472:  m.cs = 1472; goto _test_eof
	_test_eof1473:  m.cs = 1473; goto _test_eof
	_test_eof1474:  m.cs = 1474; goto _test_eof
	_test_eof1475:  m.cs = 1475; goto _test_eof
	_test_eof456:  m.cs = 456; goto _test_eof
	_test_eof457:  m.cs = 457; goto _test_eof
	_test_eof458:  m.cs = 458; goto _test_eof
	_test_eof1476:  m.cs = 1476; goto _test_eof
	_test_eof459:  m.cs = 459; goto _test_eof
	_test_eof460:  m.cs = 460; goto _test_eof
	_test_eof461:  m.cs = 461; goto _test_eof
	_test_eof1477:  m.cs = 1477; goto _test_eof
	_test_eof462:  m.cs = 462; goto _test_eof
	_test_eof463:  m.cs = 463; goto _test_eof
	_test_eof1478:  m.cs = 1478; goto _test_eof
	_test_eof1479:  m.cs = 1479; goto _test_eof
	_test_eof464:  m.cs = 464; goto _test_eof
	_test_eof465:  m.cs = 465; goto _test_eof
	_test_eof466:  m.cs = 466; goto _test_eof
	_test_eof1480:  m.cs = 1480; goto _test_eof
	_test_eof1481:  m.cs = 1481; goto _test_eof
	_test_eof1482:  m.cs = 1482; goto _test_eof
	_test_eof1483:  m.cs = 1483; goto _test_eof
	_test_eof1484:  m.cs = 1484; goto _test_eof
	_test_eof467:  m.cs = 467; goto _test_eof
	_test_eof468:  m.cs = 468; goto _test_eof
	_test_eof469:  m.cs = 469; goto _test_eof
	_test_eof1485:  m.cs = 1485; goto _test_eof
	_test_eof470:  m.cs = 470; goto _test_eof
	_test_eof471:  m.cs = 471; goto _test_eof
	_test_eof472:  m.cs = 472; goto _test_eof
	_test_eof1486:  m.cs = 1486; goto _test_eof
	_test_eof473:  m.cs = 473; goto _test_eof
	_test_eof474:  m.cs = 474; goto _test_eof
	_test_eof1487:  m.cs = 1487; goto _test_eof
	_test_eof1488:  m.cs = 1488; goto _test_eof
	_test_eof475:  m.cs = 475; goto _test_eof
	_test_eof1489:  m.cs = 1489; goto _test_eof
	_test_eof1490:  m.cs = 1490; goto _test_eof
	_test_eof1491:  m.cs = 1491; goto _test_eof
	_test_eof1492:  m.cs = 1492; goto _test_eof
	_test_eof1493:  m.cs = 1493; goto _test_eof
	_test_eof476:  m.cs = 476; goto _test_eof
	_test_eof477:  m.cs = 477; goto _test_eof
	_test_eof478:  m.cs = 478; goto _test_eof
	_test_eof1494:  m.cs = 1494; goto _test_eof
	_test_eof479:  m.cs = 479; goto _test_eof
	_test_eof480:  m.cs = 480; goto _test_eof
	_test_eof481:  m.cs = 481; goto _test_eof
	_test_eof1495:  m.cs = 1495; goto _test_eof
	_test_eof482:  m.cs = 482; goto _test_eof
	_test_eof483:  m.cs = 483; goto _test_eof
	_test_eof1496:  m.cs = 1496; goto _test_eof
	_test_eof1497:  m.cs = 1497; goto _test_eof
	_test_eof484:  m.cs = 484; goto _test_eof
	_test_eof485:  m.cs = 485; goto _test_eof
	_test_eof486:  m.cs = 486; goto _test_eof
	_test_eof487:  m.cs = 487; goto _test_eof
	_test_eof488:  m.cs = 488; goto _test_eof
	_test_eof489:  m.cs = 489; goto _test_eof
	_test_eof1498:  m.cs = 1498; goto _test_eof

	_test_eof: {}
	if ( m.p) == ( m.eof) {
		switch  m.cs {
		case 491, 492, 494, 495, 497, 498, 517, 518, 521, 522, 525, 526, 547, 549, 550, 552, 571, 572, 573, 575, 595, 596, 598, 599, 620, 622, 623, 628, 629, 631, 650, 651, 652, 654, 655, 674, 675, 677, 678, 679, 708, 709, 711, 814, 897, 898, 900, 901, 921, 922, 923, 924, 925, 954, 955, 977, 978, 998, 999, 1019, 1020, 1031, 1032, 1055, 1056, 1058, 1077, 1078, 1079, 1081, 1082, 1101, 1102, 1104, 1105, 1106, 1125, 1126, 1146, 1147, 1148, 1149, 1150, 1189, 1190, 1270, 1311, 1312, 1323, 1324, 1326, 1348, 1349, 1371, 1372, 1374, 1393, 1394, 1395, 1397, 1416, 1417, 1419, 1420, 1421, 1450, 1451, 1453:
//line machine.rl:21

	yield = true
	{( m.p)++;  m.cs = 0; goto _out }

		case 2, 3, 4, 5, 6, 24, 25, 27, 28, 29, 30, 32, 36, 42, 43, 44, 45, 59, 60, 74, 83, 113, 124, 125, 134, 136, 140, 141, 142, 143, 144, 145, 146, 147, 148, 149, 150, 151, 152, 165, 166, 167, 168, 169, 170, 171, 172, 173, 174, 175, 176, 177, 178, 190, 194, 195, 196, 197, 198, 199, 200, 201, 202, 203, 210, 211, 212, 214, 215, 217, 219, 220, 221, 222, 223, 224, 226, 227, 228, 229, 230, 231, 232, 233, 234, 235, 236, 237, 238, 239, 240, 241, 242, 243, 244, 245, 246, 247, 248, 249, 250, 251, 252, 253, 254, 255, 256, 257, 258, 259, 260, 342, 343, 352, 354, 358, 359, 360, 361, 362, 363, 364, 365, 366, 367, 368, 369, 370, 371, 372, 373, 374, 375, 376, 377, 378, 379, 380, 381, 382, 383, 384, 385, 390, 396, 397, 398, 399, 400, 433, 465, 486:
//line machine.rl:26

	m.err = ErrFieldParse
	( m.p)--

	 m.cs = 489;
	{( m.p)++;  m.cs = 0; goto _out }

		case 11, 12, 13, 19, 20:
//line machine.rl:33

	m.err = ErrTagParse
	( m.p)--

	 m.cs = 489;
	{( m.p)++;  m.cs = 0; goto _out }

		case 193:
//line machine.rl:40

	m.err = ErrTimestampParse
	( m.p)--

	 m.cs = 489;
	{( m.p)++;  m.cs = 0; goto _out }

		case 953, 956, 958, 1030, 1033, 1034, 1036, 1218, 1219:
//line machine.rl:21

	yield = true
	{( m.p)++;  m.cs = 0; goto _out }

//line machine.rl:58

	m.handler.SetMeasurement(m.text())

		case 265, 266, 267, 268, 269, 270, 271, 272, 274, 275, 277, 278, 279, 280, 281, 282, 283, 284, 285, 286, 287, 288, 289, 290, 294, 295, 297, 300, 301, 303, 306, 307, 308, 309, 310, 311, 312, 313, 314, 315, 316, 317, 318, 319, 320, 321, 322, 323, 324, 325, 326, 327, 328, 329, 331, 332, 333, 334, 335, 336, 337, 338, 339:
//line machine.rl:26

	m.err = ErrFieldParse
	( m.p)--

	 m.cs = 489;
	{( m.p)++;  m.cs = 0; goto _out }

//line machine.rl:33

	m.err = ErrTagParse
	( m.p)--

	 m.cs = 489;
	{( m.p)++;  m.cs = 0; goto _out }

		case 26, 58, 82, 139, 164, 298, 357, 431, 439:
//line machine.rl:26

	m.err = ErrFieldParse
	( m.p)--

	 m.cs = 489;
	{( m.p)++;  m.cs = 0; goto _out }

//line machine.rl:40

	m.err = ErrTimestampParse
	( m.p)--

	 m.cs = 489;
	{( m.p)++;  m.cs = 0; goto _out }

		case 213, 261, 262, 263, 291, 340, 344, 345, 346, 347, 348, 349, 350, 351, 353, 355, 356:
//line machine.rl:26

	m.err = ErrFieldParse
	( m.p)--

	 m.cs = 489;
	{( m.p)++;  m.cs = 0; goto _out }

//line machine.rl:58

	m.handler.SetMeasurement(m.text())

		case 14, 15, 16, 17, 18, 37, 38, 39, 40, 41, 50, 51, 53, 55, 56, 61, 62, 63, 64, 65, 66, 67, 68, 69, 70, 77, 78, 80, 81, 84, 85, 86, 88, 89, 90, 91, 92, 93, 94, 95, 96, 97, 98, 99, 100, 101, 102, 104, 105, 106, 107, 108, 109, 110, 111, 112, 153, 154, 155, 156, 157, 158, 159, 160, 161, 391, 392, 393, 394, 395, 401, 402, 403, 404, 405, 406, 407, 408, 409, 410, 411, 412, 413, 414, 415, 416, 417, 418, 419, 420, 421, 422, 423, 424, 425, 426, 427, 428, 429, 435, 436, 438, 441, 442, 443, 444, 445, 446, 447, 448, 449, 450, 451, 452, 453, 454, 456, 457, 458, 459, 460, 461, 462, 463, 464:
//line machine.rl:33

	m.err = ErrTagParse
	( m.p)--

	 m.cs = 489;
	{( m.p)++;  m.cs = 0; goto _out }

//line machine.rl:26

	m.err = ErrFieldParse
	( m.p)--

	 m.cs = 489;
	{( m.p)++;  m.cs = 0; goto _out }

		case 52:
//line machine.rl:33

	m.err = ErrTagParse
	( m.p)--

	 m.cs = 489;
	{( m.p)++;  m.cs = 0; goto _out }

//line machine.rl:40

	m.err = ErrTimestampParse
	( m.p)--

	 m.cs = 489;
	{( m.p)++;  m.cs = 0; goto _out }

		case 216, 218, 225, 292, 304:
//line machine.rl:40

	m.err = ErrTimestampParse
	( m.p)--

	 m.cs = 489;
	{( m.p)++;  m.cs = 0; goto _out }

//line machine.rl:26

	m.err = ErrFieldParse
	( m.p)--

	 m.cs = 489;
	{( m.p)++;  m.cs = 0; goto _out }

		case 276:
//line machine.rl:40

	m.err = ErrTimestampParse
	( m.p)--

	 m.cs = 489;
	{( m.p)++;  m.cs = 0; goto _out }

//line machine.rl:33

	m.err = ErrTagParse
	( m.p)--

	 m.cs = 489;
	{( m.p)++;  m.cs = 0; goto _out }

		case 520, 523, 527, 594, 618, 619, 624, 625, 626, 765, 766, 768, 1221, 1222, 1224, 1289, 1290, 1292, 1322, 1345, 1346, 1350, 1369, 1370:
//line machine.rl:58

	m.handler.SetMeasurement(m.text())

//line machine.rl:21

	yield = true
	{( m.p)++;  m.cs = 0; goto _out }

		case 1, 7, 8, 9, 10, 21, 22, 23, 31, 33, 34, 46, 47, 48, 49, 57, 72, 73, 75, 114, 115, 116, 117, 118, 119, 120, 121, 122, 126, 127, 128, 129, 130, 131, 132, 133, 135, 137, 138, 162, 163, 179, 180, 181, 182, 183, 184, 185, 186, 187, 189, 191, 192, 204, 205, 206, 207, 208, 209, 386, 387, 388, 430, 466, 467, 468, 469, 470, 471, 472, 473, 474, 475, 476, 477, 478, 479, 480, 481, 482, 483, 485, 487, 488:
//line machine.rl:58

	m.handler.SetMeasurement(m.text())

//line machine.rl:26

	m.err = ErrFieldParse
	( m.p)--

	 m.cs = 489;
	{( m.p)++;  m.cs = 0; goto _out }

		case 700, 731, 817, 825, 853, 890, 928, 937, 946, 1023, 1172, 1181, 1211, 1273, 1282, 1315, 1442, 1473:
//line machine.rl:74

	m.handler.AddInt(key, m.text())

//line machine.rl:21

	yield = true
	{( m.p)++;  m.cs = 0; goto _out }

		case 546, 698, 699, 701, 707, 730, 732, 815, 816, 823, 824, 826, 851, 852, 854, 888, 889, 891, 926, 927, 929, 935, 936, 938, 944, 945, 947, 1021, 1022, 1024, 1170, 1171, 1173, 1179, 1180, 1182, 1188, 1210, 1212, 1271, 1272, 1274, 1280, 1281, 1283, 1313, 1314, 1316, 1440, 1441, 1443, 1449, 1472, 1474:
//line machine.rl:78

	m.handler.AddFloat(key, m.text())

//line machine.rl:21

	yield = true
	{( m.p)++;  m.cs = 0; goto _out }

		case 702, 703, 704, 705, 706, 733, 734, 735, 736, 737, 818, 819, 820, 821, 822, 827, 828, 829, 830, 831, 855, 856, 857, 858, 859, 892, 893, 894, 895, 896, 930, 931, 932, 933, 934, 939, 940, 941, 942, 943, 948, 949, 950, 951, 952, 1025, 1026, 1027, 1028, 1029, 1174, 1175, 1176, 1177, 1178, 1183, 1184, 1185, 1186, 1187, 1213, 1214, 1215, 1216, 1217, 1275, 1276, 1277, 1278, 1279, 1284, 1285, 1286, 1287, 1288, 1317, 1318, 1319, 1320, 1321, 1444, 1445, 1446, 1447, 1448, 1475, 1476, 1477, 1478, 1479:
//line machine.rl:82

	m.handler.AddBool(key, m.text())

//line machine.rl:21

	yield = true
	{( m.p)++;  m.cs = 0; goto _out }

		case 496, 499, 500, 501, 502, 503, 504, 505, 506, 507, 508, 509, 510, 511, 512, 513, 514, 515, 516, 551, 553, 554, 555, 556, 557, 558, 559, 560, 561, 562, 563, 564, 565, 566, 567, 568, 569, 570, 574, 576, 577, 578, 579, 580, 581, 582, 583, 584, 585, 586, 587, 588, 589, 590, 591, 592, 593, 597, 600, 601, 602, 603, 604, 605, 606, 607, 608, 609, 610, 611, 612, 613, 614, 615, 616, 617, 630, 632, 633, 634, 635, 636, 637, 638, 639, 640, 641, 642, 643, 644, 645, 646, 647, 648, 649, 653, 656, 657, 658, 659, 660, 661, 662, 663, 664, 665, 666, 667, 668, 669, 670, 671, 672, 673, 676, 680, 681, 682, 683, 684, 685, 686, 687, 688, 689, 690, 691, 692, 693, 694, 695, 696, 697, 710, 712, 713, 714, 715, 716, 717, 718, 719, 720, 721, 722, 723, 724, 725, 726, 727, 728, 729, 795, 796, 797, 798, 799, 800, 801, 802, 803, 804, 805, 806, 807, 808, 809, 810, 811, 812, 813, 832, 833, 834, 835, 836, 837, 838, 839, 840, 841, 842, 843, 844, 845, 846, 847, 848, 849, 850, 869, 870, 871, 872, 873, 874, 875, 876, 877, 878, 879, 880, 881, 882, 883, 884, 885, 886, 887, 902, 903, 904, 905, 906, 907, 908, 909, 910, 911, 912, 913, 914, 915, 916, 917, 918, 919, 920, 979, 980, 981, 982, 983, 984, 985, 986, 987, 988, 989, 990, 991, 992, 993, 994, 995, 996, 997, 1000, 1001, 1002, 1003, 1004, 1005, 1006, 1007, 1008, 1009, 1010, 1011, 1012, 1013, 1014, 1015, 1016, 1017, 1018, 1057, 1059, 1060, 1061, 1062, 1063, 1064, 1065, 1066, 1067, 1068, 1069, 1070, 1071, 1072, 1073, 1074, 1075, 1076, 1080, 1083, 1084, 1085, 1086, 1087, 1088, 1089, 1090, 1091, 1092, 1093, 1094, 1095, 1096, 1097, 1098, 1099, 1100, 1103, 1107, 1108, 1109, 1110, 1111, 1112, 1113, 1114, 1115, 1116, 1117, 1118, 1119, 1120, 1121, 1122, 1123, 1124, 1127, 1128, 1129, 1130, 1131, 1132, 1133, 1134, 1135, 1136, 1137, 1138, 1139, 1140, 1141, 1142, 1143, 1144, 1145, 1151, 1152, 1153, 1154, 1155, 1156, 1157, 1158, 1159, 1160, 1161, 1162, 1163, 1164, 1165, 1166, 1167, 1168, 1169, 1191, 1192, 1193, 1194, 1195, 1196, 1197, 1198, 1199, 1200, 1201, 1202, 1203, 1204, 1205, 1206, 1207, 1208, 1209, 1251, 1252, 1253, 1254, 1255, 1256, 1257, 1258, 1259, 1260, 1261, 1262, 1263, 1264, 1265, 1266, 1267, 1268, 1269, 1325, 1327, 1328, 1329, 1330, 1331, 1332, 1333, 1334, 1335, 1336, 1337, 1338, 1339, 1340, 1341, 1342, 1343, 1344, 1373, 1375, 1376, 1377, 1378, 1379, 1380, 1381, 1382, 1383, 1384, 1385, 1386, 1387, 1388, 1389, 1390, 1391, 1392, 1396, 1398, 1399, 1400, 1401, 1402, 1403, 1404, 1405, 1406, 1407, 1408, 1409, 1410, 1411, 1412, 1413, 1414, 1415, 1418, 1422, 1423, 1424, 1425, 1426, 1427, 1428, 1429, 1430, 1431, 1432, 1433, 1434, 1435, 1436, 1437, 1438, 1439, 1452, 1454, 1455, 1456, 1457, 1458, 1459, 1460, 1461, 1462, 1463, 1464, 1465, 1466, 1467, 1468, 1469, 1470, 1471:
//line machine.rl:90

	m.handler.SetTimestamp(m.text())

//line machine.rl:21

	yield = true
	{( m.p)++;  m.cs = 0; goto _out }

		case 296, 299:
//line machine.rl:26

	m.err = ErrFieldParse
	( m.p)--

	 m.cs = 489;
	{( m.p)++;  m.cs = 0; goto _out }

//line machine.rl:33

	m.err = ErrTagParse
	( m.p)--

	 m.cs = 489;
	{( m.p)++;  m.cs = 0; goto _out }

//line machine.rl:40

	m.err = ErrTimestampParse
	( m.p)--

	 m.cs = 489;
	{( m.p)++;  m.cs = 0; goto _out }

		case 341:
//line machine.rl:26

	m.err = ErrFieldParse
	( m.p)--

	 m.cs = 489;
	{( m.p)++;  m.cs = 0; goto _out }

//line machine.rl:58

	m.handler.SetMeasurement(m.text())

//line machine.rl:40

	m.err = ErrTimestampParse
	( m.p)--

	 m.cs = 489;
	{( m.p)++;  m.cs = 0; goto _out }

		case 54, 79, 87, 103, 437, 440, 455:
//line machine.rl:33

	m.err = ErrTagParse
	( m.p)--

	 m.cs = 489;
	{( m.p)++;  m.cs = 0; goto _out }

//line machine.rl:26

	m.err = ErrFieldParse
	( m.p)--

	 m.cs = 489;
	{( m.p)++;  m.cs = 0; goto _out }

//line machine.rl:40

	m.err = ErrTimestampParse
	( m.p)--

	 m.cs = 489;
	{( m.p)++;  m.cs = 0; goto _out }

		case 273, 302, 305, 330:
//line machine.rl:40

	m.err = ErrTimestampParse
	( m.p)--

	 m.cs = 489;
	{( m.p)++;  m.cs = 0; goto _out }

//line machine.rl:33

	m.err = ErrTagParse
	( m.p)--

	 m.cs = 489;
	{( m.p)++;  m.cs = 0; goto _out }

//line machine.rl:26

	m.err = ErrFieldParse
	( m.p)--

	 m.cs = 489;
	{( m.p)++;  m.cs = 0; goto _out }

		case 264, 293:
//line machine.rl:40

	m.err = ErrTimestampParse
	( m.p)--

	 m.cs = 489;
	{( m.p)++;  m.cs = 0; goto _out }

//line machine.rl:58

	m.handler.SetMeasurement(m.text())

//line machine.rl:26

	m.err = ErrFieldParse
	( m.p)--

	 m.cs = 489;
	{( m.p)++;  m.cs = 0; goto _out }

		case 35, 71, 123, 389, 432:
//line machine.rl:58

	m.handler.SetMeasurement(m.text())

//line machine.rl:26

	m.err = ErrFieldParse
	( m.p)--

	 m.cs = 489;
	{( m.p)++;  m.cs = 0; goto _out }

//line machine.rl:40

	m.err = ErrTimestampParse
	( m.p)--

	 m.cs = 489;
	{( m.p)++;  m.cs = 0; goto _out }

		case 739, 788, 862, 1244, 1482, 1491:
//line machine.rl:58

	m.handler.SetMeasurement(m.text())

//line machine.rl:74

	m.handler.AddInt(key, m.text())

//line machine.rl:21

	yield = true
	{( m.p)++;  m.cs = 0; goto _out }

		case 627, 738, 740, 764, 787, 789, 860, 861, 863, 1220, 1243, 1245, 1480, 1481, 1483, 1489, 1490, 1492:
//line machine.rl:58

	m.handler.SetMeasurement(m.text())

//line machine.rl:78

	m.handler.AddFloat(key, m.text())

//line machine.rl:21

	yield = true
	{( m.p)++;  m.cs = 0; goto _out }

		case 741, 742, 743, 744, 745, 790, 791, 792, 793, 794, 864, 865, 866, 867, 868, 1246, 1247, 1248, 1249, 1250, 1484, 1485, 1486, 1487, 1488, 1493, 1494, 1495, 1496, 1497:
//line machine.rl:58

	m.handler.SetMeasurement(m.text())

//line machine.rl:82

	m.handler.AddBool(key, m.text())

//line machine.rl:21

	yield = true
	{( m.p)++;  m.cs = 0; goto _out }

		case 524, 528, 529, 530, 531, 532, 533, 534, 535, 536, 537, 538, 539, 540, 541, 542, 543, 544, 545, 621, 746, 747, 748, 749, 750, 751, 752, 753, 754, 755, 756, 757, 758, 759, 760, 761, 762, 763, 767, 769, 770, 771, 772, 773, 774, 775, 776, 777, 778, 779, 780, 781, 782, 783, 784, 785, 786, 1223, 1225, 1226, 1227, 1228, 1229, 1230, 1231, 1232, 1233, 1234, 1235, 1236, 1237, 1238, 1239, 1240, 1241, 1242, 1291, 1293, 1294, 1295, 1296, 1297, 1298, 1299, 1300, 1301, 1302, 1303, 1304, 1305, 1306, 1307, 1308, 1309, 1310, 1347, 1351, 1352, 1353, 1354, 1355, 1356, 1357, 1358, 1359, 1360, 1361, 1362, 1363, 1364, 1365, 1366, 1367, 1368:
//line machine.rl:58

	m.handler.SetMeasurement(m.text())

//line machine.rl:90

	m.handler.SetTimestamp(m.text())

//line machine.rl:21

	yield = true
	{( m.p)++;  m.cs = 0; goto _out }

		case 957, 959, 960, 961, 962, 963, 964, 965, 966, 967, 968, 969, 970, 971, 972, 973, 974, 975, 976, 1035, 1037, 1038, 1039, 1040, 1041, 1042, 1043, 1044, 1045, 1046, 1047, 1048, 1049, 1050, 1051, 1052, 1053, 1054:
//line machine.rl:90

	m.handler.SetTimestamp(m.text())

//line machine.rl:21

	yield = true
	{( m.p)++;  m.cs = 0; goto _out }

//line machine.rl:58

	m.handler.SetMeasurement(m.text())

//line machine.go:57677
		}
	}

	_out: {}
	}

//line machine.rl:231

	// Even though there was an error, return true. On the next call to this
	// function we will attempt to scan to the next line of input and recover.
	if m.err != nil {
		return true
	}

	// Don't check the error state in the case that we just yielded, because
	// the yield indicates we just completed parsing a line.
	if !yield && m.cs == LineProtocol_error {
		m.err = ErrParse
		return true
	}

	return true
}

// Err returns the error that occurred on the last call to ParseLine.  If the
// result is nil, then the line was parsed successfully.
func (m *machine) Err() error {
	return m.err
}

// Position returns the current position into the input.
func (m *machine) Position() int {
	return m.p
}

func (m *machine) text() []byte {
	return m.data[m.pb:m.p]
}
