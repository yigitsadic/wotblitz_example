package database

import (
	"context"
	"github.com/yigitsadic/wotblitz_example/ent"
	"github.com/yigitsadic/wotblitz_example/graph/model"
	"log"
)

func SeedDB(ctx context.Context, client *ent.Client) {
	count, _ := client.Tank.Query().Where().Count(ctx)
	if count > 0 {
		log.Println("Tanks found on db. Skipped seed.")
		return
	}

	tx, err := client.Tx(ctx)
	if err != nil {
		panic("unable to initialize transaction")
	}

	r35, _ := tx.Tank.Create().
		SetName("R35").
		SetTier(1).
		SetTankClass(model.TankClassMediumTank.String()).
		SetIsPremium(false).
		SetCountry(model.CountryFrance.String()).
		Save(ctx)

	amx_38, _ := tx.Tank.Create().
		SetName("AMX 38").
		SetTier(2).
		SetTankClass(model.TankClassMediumTank.String()).
		SetIsPremium(false).
		SetCountry(model.CountryFrance.String()).
		AddPreviousTanks(r35).
		Save(ctx)

	r35.Update().AddNextTanks(amx_38).Save(ctx)

	d2, _ := tx.Tank.Create().
		SetName("D2").
		SetTier(3).
		SetTankClass(model.TankClassHeavyTank.String()).
		SetIsPremium(false).
		SetCountry(model.CountryFrance.String()).
		AddPreviousTanks(amx_38).
		Save(ctx)

	amx_38.Update().AddNextTanks(d2).Save(ctx)

	sau_40, _ := tx.Tank.Create().
		SetName("SAu 40").
		SetTier(4).
		SetTankClass(model.TankClassTankDestroyer.String()).
		SetCountry(model.CountryFrance.String()).
		AddPreviousTanks(d2).
		Save(ctx)

	b1, _ := tx.Tank.Create().
		SetName("B1").
		SetTier(4).
		SetTankClass(model.TankClassHeavyTank.String()).
		SetCountry(model.CountryFrance.String()).
		AddPreviousTanks(d2).
		Save(ctx)

	d2.Update().AddNextTanks(sau_40, b1).Save(ctx)

	s35_ca, _ := tx.Tank.Create().
		SetName("S35 CA").
		SetCountry(model.CountryFrance.String()).
		SetTier(5).
		SetTankClass(model.TankClassTankDestroyer.String()).
		AddPreviousTanks(sau_40).
		Save(ctx)

	s35_ca.Update().AddPreviousTanks(sau_40).Save(ctx)

	bdr_g1b, _ := tx.Tank.Create().
		SetName("BDR G1 B").
		SetTier(5).
		SetTankClass(model.TankClassHeavyTank.String()).
		SetCountry(model.CountryFrance.String()).
		AddPreviousTanks(b1).
		Save(ctx)

	amx_elc_bis, _ := tx.Tank.Create().
		SetName("AMX ELC bis").
		SetTier(5).
		SetTankClass(model.TankClassLightTank.String()).
		SetCountry(model.CountryFrance.String()).
		AddPreviousTanks(b1).
		Save(ctx)

	b1.Update().AddNextTanks(bdr_g1b, amx_elc_bis).Save(ctx)

	amx_v39, _ := tx.Tank.Create().
		SetName("AMX V39").
		SetTier(6).
		SetTankClass(model.TankClassTankDestroyer.String()).
		SetCountry(model.CountryFrance.String()).
		AddPreviousTanks(s35_ca).
		Save(ctx)
	s35_ca.Update().AddNextTanks(amx_v39).Save(ctx)

	arl_44, _ := tx.Tank.Create().
		SetName("ARL 44").
		SetTier(6).
		SetTankClass(model.TankClassHeavyTank.String()).
		SetCountry(model.CountryFrance.String()).
		AddPreviousTanks(bdr_g1b).
		Save(ctx)

	bdr_g1b.Update().AddNextTanks(arl_44).Save(ctx)

	amx_12_t, _ := tx.Tank.Create().
		SetName("AMX 12 t").
		SetTier(6).
		SetTankClass(model.TankClassLightTank.String()).
		SetCountry(model.CountryFrance.String()).
		AddPreviousTanks(amx_elc_bis).
		Save(ctx)

	amx_elc_bis.Update().AddNextTanks(amx_12_t).Save(ctx)

	amx_ac_46, _ := tx.Tank.Create().
		SetName("AMX AC 46").
		SetTier(7).
		SetTankClass(model.TankClassTankDestroyer.String()).
		SetCountry(model.CountryFrance.String()).
		AddPreviousTanks(arl_44).
		Save(ctx)
	arl_44.Update().AddNextTanks(amx_ac_46).Save(ctx)

	amx_m4_45, _ := tx.Tank.Create().
		SetName("AMX M4 45").
		SetTier(7).
		SetTankClass(model.TankClassHeavyTank.String()).
		SetCountry(model.CountryFrance.String()).
		AddPreviousTanks(arl_44).
		Save(ctx)
	arl_44.Update().AddNextTanks(amx_m4_45).Save(ctx)

	amx_13_75, _ := tx.Tank.Create().
		SetName("AMX 13 75").
		SetTier(7).
		SetTankClass(model.TankClassLightTank.String()).
		SetCountry(model.CountryFrance.String()).
		AddPreviousTanks(amx_12_t).
		Save(ctx)
	amx_12_t.Update().AddNextTanks(amx_13_75).Save(ctx)

	amx_ac_48, _ := tx.Tank.Create().
		SetName("AMX AC 48").
		SetTier(8).
		SetTankClass(model.TankClassTankDestroyer.String()).
		SetCountry(model.CountryFrance.String()).
		AddPreviousTanks(amx_ac_46).
		Save(ctx)
	amx_ac_46.Update().AddNextTanks(amx_ac_48).Save(ctx)

	amx_50_100, _ := tx.Tank.Create().
		SetName("AMX 50 100").
		SetTier(8).
		SetTankClass(model.TankClassHeavyTank.String()).
		SetCountry(model.CountryFrance.String()).
		AddPreviousTanks(amx_m4_45).
		Save(ctx)
	amx_m4_45.Update().AddNextTanks(amx_50_100).Save(ctx)

	amx_13_90, _ := tx.Tank.Create().
		SetName("AMX 13 90").
		SetTier(8).
		SetTankClass(model.TankClassLightTank.String()).
		SetCountry(model.CountryFrance.String()).
		AddPreviousTanks(amx_13_75).
		Save(ctx)
	amx_13_75.Update().AddNextTanks(amx_13_90).Save(ctx)

	tx.Tank.Create().
		SetName("FCM 50 t").
		SetTier(8).
		SetTankClass(model.TankClassMediumTank.String()).
		SetCountry(model.CountryFrance.String()).
		SetIsPremium(true).
		Save(ctx)

	foch, _ := tx.Tank.Create().
		SetName("Foch").
		SetTier(9).
		SetTankClass(model.TankClassTankDestroyer.String()).
		SetCountry(model.CountryFrance.String()).
		AddPreviousTanks(amx_ac_48).
		Save(ctx)
	amx_ac_48.Update().AddNextTanks(foch).Save(ctx)

	amx_50_120, _ := tx.Tank.Create().
		SetName("AMX 50 120").
		SetTier(9).
		SetTankClass(model.TankClassHeavyTank.String()).
		SetCountry(model.CountryFrance.String()).
		AddPreviousTanks(amx_50_100).
		Save(ctx)
	amx_50_100.Update().AddNextTanks(amx_50_120).Save(ctx)

	bc_25_t_ap, _ := tx.Tank.Create().
		SetName("B-C 25 t AP").
		SetTier(9).
		SetTankClass(model.TankClassLightTank.String()).
		SetCountry(model.CountryFrance.String()).
		AddPreviousTanks(amx_13_90).
		Save(ctx)
	amx_13_90.Update().AddNextTanks(bc_25_t_ap).Save(ctx)

	foch_155, _ := tx.Tank.Create().
		SetName("Foch 155").
		SetTier(10).
		SetTankClass(model.TankClassTankDestroyer.String()).
		SetCountry(model.CountryFrance.String()).
		AddPreviousTanks(foch).
		Save(ctx)
	foch.Update().AddNextTanks(foch_155).Save(ctx)

	amx_50_b, _ := tx.Tank.Create().
		SetName("AMX 50 B").
		SetTier(10).
		SetTankClass(model.TankClassHeavyTank.String()).
		SetCountry(model.CountryFrance.String()).
		AddPreviousTanks(amx_50_120).
		Save(ctx)
	amx_50_120.Update().AddNextTanks(amx_50_b).Save(ctx)

	bc_25_t, _ := tx.Tank.Create().
		SetName("B-C 25 t").
		SetTier(10).
		SetTankClass(model.TankClassLightTank.String()).
		SetCountry(model.CountryFrance.String()).
		AddPreviousTanks(bc_25_t_ap).
		Save(ctx)
	bc_25_t_ap.Update().AddNextTanks(bc_25_t).Save(ctx)

	err = tx.Commit()
	if err != nil {
		panic(err)
	}
}
