package cards

import (
	"duel-masters/game/cards/dm01"
	"duel-masters/game/cards/dm02"
	"duel-masters/game/cards/dm03"
	"duel-masters/game/match"
)

// Sets is a map of pointers to the available card sets
var Sets = map[string]*map[string]match.CardConstructor{
	"dm-01": &DM01,
	"dm-02": &DM02,
	"dm-03": &DM03,
}

// DM01 is a map with all the card id's in the game and corresponding CardConstructor for dm01
var DM01 = map[string]match.CardConstructor{

	"57eeb3c3-2561-4841-a381-2e50d17533d1": dm01.AquaHulcus,
	"ecd1ae69-4f63-4e8d-a3f4-9a5c81f98a20": dm01.EmeraldGrass,
	"09b218fc-9c5a-48ef-9555-4908932271e9": dm01.AquaKnight,
	"4097a036-a775-4218-9a1d-f57ead85dda6": dm01.AquaSniper,
	"c43bc627-9e7a-4686-9d61-789425669b02": dm01.AquaSoldier,
	"9781089f-1aa9-4a75-b106-35e9d431e31d": dm01.AquaVehicle,
	"1d72eb3e-5185-449a-a16f-391bd2338343": dm01.BurningMane,
	"fcd0cb50-b687-4180-90a8-390aeb8705cc": dm01.FearFang,
	"10e0e90f-ad7d-4b69-98d5-f01525eb1cdd": dm01.SteelSmasher,
	"015fd6bb-37a9-45cf-bb6b-a5497412b880": dm01.BronzeArmTribe,
	"6663848d-035e-44b6-9d9f-7b236ea5bc43": dm01.GoldenWingStriker,
	"0e26fe1a-a9d1-4c78-80e9-7f4cc0e4c1c8": dm01.MightyShouter,
	"0b1e4f56-6342-46db-9faf-882fd1f1f179": dm01.ArtisanPicora,
	"983e72d7-3f4e-466d-a4e3-06552e392af2": dm01.NomadHeroGigio,
	"0cc5279e-0a26-41a8-a2a5-f7711120b772": dm01.LahPurificationEnforcer,
	"808ddd60-e8ca-49f0-9baa-57e632f85b28": dm01.RaylaTruthEnforcer,
	"91db2302-6794-4aa4-b17b-6637d356e9ac": dm01.AstrocometDragon,
	"0ffdcae3-9db2-401b-8a82-dfad707b83cd": dm01.BolshackDragon,
	"6cf85053-abaa-4577-b151-86123004980e": dm01.Draglide,
	"3b6e6c29-017d-41b9-bf93-186f7963723e": dm01.GatlingSkyterror,
	"1c5511be-7629-41c5-bf17-4bc810be5472": dm01.ScarletSkyterror,
	"a4adb373-0aec-4fff-997c-3820c7ec528d": dm01.DomeShell,
	"1ecb54a2-bcbf-4396-bf09-50dfe984e287": dm01.StormShell,
	"c761c174-87c3-4f4a-ab94-aa837c5ab587": dm01.TowerShell,
	"18e0e199-7827-4a4c-a37d-3acfa4e500d6": dm01.RoaringGreatHorn,
	"2aeae452-5630-4f86-b073-7e9dc07adc43": dm01.StampedingLonghorn,
	"84e1b416-c2d5-4ae1-aca0-025651c6aa58": dm01.TriHornShepherd,
	"3e2940f4-5654-4456-bfc2-fa5e43911cfb": dm01.KingCoral,
	"cd13f7c2-aa5e-43b8-8811-700f230a5de5": dm01.KingDepthcon,
	"f04feb7f-971f-4192-893a-46c23180233a": dm01.KingRippedHide,
	"596f5b72-2502-4120-81f9-9ff9a17271d8": dm01.CandyDrop,
	"a3cf18f0-b04f-45e9-97f7-2a2ead0a1787": dm01.FaerieChild,
	"3f331274-f5f8-42e7-9f28-ce637add34d4": dm01.MarineFlower,
	"ce48ff2c-ea9e-4c12-8629-028d2480b063": dm01.IllusionaryMerfolk,
	"4b021e6f-39cf-401e-89cf-f164f7c0a797": dm01.PhantomFish,
	"cfe9f5b8-2eeb-42c9-89ff-7e69734adc4d": dm01.RevolverFish,
	"70e6cc2c-c63d-4dd9-9b6e-0713fed174bb": dm01.SaucerHeadShark,
	"4c9acf76-cc52-44c3-9e39-613d744c63c5": dm01.PoisonousMushroom,
	"cc9762c3-515a-4734-a3fe-1e0c4c3b3d71": dm01.BoneAssassin,
	"4d3201e8-0d9b-481e-b8e3-86cb90058e20": dm01.BoneSpider,
	"ec46daa1-49ce-4b88-b2bc-e923672ad0f3": dm01.SkeletonSoldierTheDefiled,
	"90b2ed59-828c-4237-ac2e-b7008a02ad2e": dm01.WanderingBraineater,
	"5d3d7052-e5fa-4502-8d31-c72673232317": dm01.HanusaRadianceElemental,
	"6a4270cf-f3be-4c66-8b30-eb2c769065dc": dm01.IocantTheOracle,
	"25a2af16-cc42-4f4c-8c3d-59fb3a7ca74b": dm01.UrthPurifyingElemental,
	"c4839847-e393-47b0-b172-95531aa6d39e": dm01.Gigaberos,
	"5d73062e-acff-47e6-b49a-c0bb1a1762b5": dm01.Gigagiele,
	"6161e271-5294-4073-94d2-b9c06f9d8fa3": dm01.Gigargon,
	"dc1b51b3-52e7-4f1c-8770-515d4e1cb53d": dm01.DeathligerLionOfChaos,
	"07a0115e-797a-49d8-90bf-9ea6de39978d": dm01.ZagaanKnightOfDarkness,
	"7a6f1c82-a8ac-4646-b3e9-fb8592bdd0a4": dm01.Tropico,
	"b7d11c62-2ab3-439b-b147-ae29d34e9216": dm01.FreiVizierOfAir,
	"578ed21b-8ba5-42b2-b662-87a321ee0c7d": dm01.IereVizierOfBullets,
	"cf5eb3d3-e128-42db-bf1a-161d5dd4b972": dm01.LokVizierOfHunting,
	"15efe8b0-02c1-439b-8e7c-4548e74f5c33": dm01.MieleVizierOfLightning,
	"340ec79b-3a4e-4483-ac0e-5dd6b40eb4e1": dm01.ToelVizierOfLight,
	"d067285f-10ea-4666-99c8-bc23e27e3262": dm01.Meteosaur,
	"41c0664d-1969-487d-bde5-866127c1c49e": dm01.Stonesaur,
	"c1ebdda0-be88-4665-937e-2ef3ada8d378": dm01.DeathbladeBeetle,
	"43abeec5-0597-43b3-93cf-766b95d19b5b": dm01.ForestHornet,
	"b3ca1944-41a2-4939-ae85-1a73b1fe085f": dm01.RedEyeScorpion,
	"162f70fb-33f7-4436-a114-41f255c0ce7e": dm01.DarkRavenShadowOfGrief,
	"ea878730-fde0-4bd0-ad25-95e49f54a1b2": dm01.MaskedHorrorShadowOfScorn,
	"f16795cc-4378-4e36-b13a-19f9b932228c": dm01.NightMasterShadowOfDecay,
	"c5a869f4-a959-4667-a352-92df5369e0b9": dm01.DeadlyFighterBraidClaw,
	"f682051b-7cc3-4155-aa8b-eb3335b0435c": dm01.ExplosiveFighterUcarn,
	"c782edd9-34ef-47f5-8f16-af2c3b107a36": dm01.FireSweeperBurningHellion,
	"198ffce7-3d79-420e-9d9b-ebd6421adb6f": dm01.OnslaughterTriceps,
	"becd0856-fb8b-46fd-a950-b57cc5d17c70": dm01.SuperExplosiveVolcanodon,
	"616c146e-049f-4720-a225-0a189729ca79": dm01.ChiliasTheOracle,
	"7b58e8c2-0b1e-4ef5-812f-e667c2092c73": dm01.ReusolTheOracle,
	"d5d57060-ca58-48e1-8903-9b8362c92b0d": dm01.RubyGrass,
	"725a28b7-8c06-4691-93d8-1c6b0dacdba5": dm01.SenatineJadeTree,
	"ae66061e-6039-4dee-abf0-51169913bb35": dm01.ArmoredWalkerUrherion,
	"5370bad9-1260-455e-8120-ea89badc7eaf": dm01.BrawlerZyler,
	"ebd730e1-1099-41ec-a028-6ef1d4cf91b2": dm01.FatalAttackerHorvath,
	"af3bc221-1cc2-4f58-83ea-2673ac2c66c5": dm01.ImmortalBaronVorg,
	"f7dc24d2-2a84-46ff-9661-0b8418d68650": dm01.DiaNorkMoonlightGuardian,
	"39090f65-779c-46c9-856c-67303dd5605c": dm01.GranGureSpaceGuardian,
	"c05fe45d-690e-4856-bddb-5f46154e57e5": dm01.LaUraGigaSkyGuardian,
	"eccceb7c-834c-4bf9-b0cd-c2dc6fad3dbf": dm01.SzubsKinTwilightGuardian,
	"a7eceb07-4f6d-4b2b-8dba-7a3df8f803f7": dm01.StingerWorm,
	"edd6cffc-8c91-4682-b8af-64cfe823103b": dm01.SwampWorm,
	"a8503655-fdcb-48e2-bfb0-0ad3aae31f0e": dm01.RothusTheTraveler,
	"e2e5e1ef-c613-449a-8400-15581082501b": dm01.CoilingVines,
	"bee69327-ca6b-455c-b3dc-463fc3284b61": dm01.PoisonousDahlia,
	"bbc655b3-3676-4cda-9554-e2d465e20b99": dm01.ThornyMandra,
	"c971ff15-5735-4d61-bb16-c805130ca405": dm01.HunterFish,
	"446eaf96-36c8-4093-b4b2-e77e7afb6e3f": dm01.Seamine,
	"c9b98336-312d-4d08-8add-6b820a88815f": dm01.UnicornFish,
	"dbdbad44-6a62-4eff-b8f1-95f56588a13a": dm01.VampireSilphy,
	"f3ded71d-3cf9-415b-a9d2-b759ca0ce07b": dm01.BloodySquito,
	"dd9d1cc1-01cb-4bf9-80c4-821e3c449887": dm01.DarkClown,
	"7b22cc2c-3a4a-4f50-9e61-fb9646a762cd": dm01.AuraBlast,
	"7f225860-af37-47ac-9b36-1480872576b6": dm01.BrainSerum,
	"5cafc789-e730-4472-9a62-8b333b2691e6": dm01.BurningPower,
	"71d90484-c144-4dcf-ad8e-23e7e55f0f2e": dm01.ChaosStrike,
	"452aead9-2a65-46b1-84b2-383fe99ddc5f": dm01.CreepingPlague,
	"87a102b5-71fd-410a-a8f0-c35182217f08": dm01.CrimsonHammer,
	"2c8ded77-89f3-4625-aa3e-6576b83e0384": dm01.CrystalMemory,
	"6f2cc530-1228-4b03-9ec0-ba24f6a367bf": dm01.DarkReversal,
	"d1703c3b-8e49-4959-8322-ae11a7ca6632": dm01.DeathSmoke,
	"32acfe8b-90fc-4ba9-b6ad-7655c0abee12": dm01.DimensionGate,
	"ae797f95-54b1-48e9-9216-f315b39826bd": dm01.GhostTouch,
	"0ec572b0-ffaf-4abd-a540-ba26c98aacc5": dm01.HolyAwe,
	"95bfccf9-91cf-4ab9-8298-c95bc368bf0b": dm01.LaserWing,
	"35a9315c-2c08-46e0-b96b-daf3e8e996ce": dm01.MagmaGazer,
	"b12f1d66-46ee-49b9-878d-59cc3d515633": dm01.MoonlightFlash,
	"a2e11f7f-63f4-4357-9e77-5314576dff45": dm01.NaturalSnare,
	"858b2c1d-5507-46f0-8840-151ee59f760b": dm01.PangeasSong,
	"9430e127-ce64-4572-b386-f1ce3f50e94a": dm01.SolarRay,
	"c3389188-718d-45bb-8946-0a572d96916b": dm01.SonicWing,
	"be80f0e8-05b7-4914-916a-f24d5e616ea8": dm01.SpiralGate,
	"a3331db6-d3c2-4b3f-9fcc-f4aa9fc2bb52": dm01.Teleportation,
	"5883180e-d88c-4f24-b17c-f5a837420147": dm01.TerrorPit,
	"48c5c29b-2f4e-4a57-86b4-864c6f0dc124": dm01.TornadoFlame,
	"be8c0d0b-dcab-402c-8e7b-878e35bacca7": dm01.UltimateForce,
	"68d78fd4-db8a-43a6-8eb6-e1435cfc2959": dm01.VirtualTripwire,
	"c62208ec-efc8-4b08-bb01-6cd5251b0969": dm01.BlackFeatherShadowOfRage,
	"e2b992ee-91a3-49d3-8228-7be60a0b9ec5": dm01.WrithingBoneGhoul,
}

// DM02 is a map with all the card id's in the game and corresponding CardConstructor for dm02
var DM02 = map[string]match.CardConstructor{
	"40439f79-8f48-4e62-9009-cb06798ef7ac": dm02.EthelStarSeaElemental,
	"48ab3f2b-4ae3-41a4-ae6f-61b49c958bdb": dm02.BarkwhipTheSmasher,
	"0bea1262-311a-47b1-888d-dd065cfe3d7f": dm02.EngineerKipo,
	"0dca6f6c-c426-4c88-b283-043527f04bb3": dm02.FighterDualFang,
	"1eca6a24-9270-477f-a588-80859481ef94": dm02.SpiralGrass,
	"2c7e38e1-0546-47ab-9388-383d093405b2": dm02.FortressShell,
	"3f0fb8f6-d01e-4005-8340-b84584f50a2a": dm02.CrystalLancer,
	"4b715b5c-2e82-4686-9c9f-4ce1e5503621": dm02.BurstShot,
	"05d946f7-5977-4f51-8bca-ecb39845f1a2": dm02.BolzardDragon,
	"5cf64846-0eb2-4e8d-bf15-4ca573f96e58": dm02.KingNautilus,
	"5d095b28-262e-454d-96c7-9174ed83e3f6": dm02.LadiaBaleTheInspirational,
	"6e381955-231b-4e4e-a14b-82509a5e193b": dm02.RumblingTerahorn,
	"9fed2257-362f-43c7-b50e-5526ccf799aa": dm02.MiniTitanGett,
	"17ee5046-c3fd-4422-af14-c54a4be8d9a2": dm02.LogicCube,
	"39b8b1c0-bc1d-445b-b60d-91cabfe62fb5": dm02.DarkTitanMaginn,
	"41e8d5c2-bfeb-48bb-ab9e-aaee79852c89": dm02.CrystalPaladin,
	"60d8c6a6-20c1-425c-9ecc-b56981a70e21": dm02.ElfX,
	"66ac493f-b836-46c9-a28e-09e7bd040064": dm02.Gigastand,
	"66ee3b9e-fdbf-41c2-9363-5327572706f2": dm02.ChaosWorm,
	"84b7a5f2-cbfc-4d2d-a757-6685aa38c241": dm02.GrayBalloonShadowOfGreed,
	"94e5d2a3-d8b2-4fce-9899-61adb063dc60": dm02.SilverAxe,
	"96ecef8e-0485-418d-9d0d-a4169c3b70b7": dm02.FonchTheOracle,
	"98b1afbe-5a0a-461e-800c-34d02339f21f": dm02.ThoughtProbe,
	"215c4cfb-2a22-4ee1-b4ea-28ac24a1eeee": dm02.UltracideWorm,
	"733a4f35-7470-40e3-9cd7-479aa965bfbb": dm02.HorridWorm,
	"749d1209-4847-4401-8d70-8c95fd26c220": dm02.MagrisVizierOfMagnetism,
	"839f780b-067b-4b3c-9090-7cd614375da9": dm02.AmberPiercer,
	"1960f3c2-5321-4da1-be56-50a7e98cae0d": dm02.PhalEegaDawnGuardian,
	"4916f651-6169-4e8b-be5d-6a3237dc0a70": dm02.MarrowOozeTheTwister,
	"7956b4f5-b910-403d-b388-b67c837b7e99": dm02.ScissorEye,
	"9311b3ec-fe9e-45bd-a989-dc00aec3ca60": dm02.DogarnTheMarauder,
	"24353d06-89ef-4867-9513-485750d01e10": dm02.ArmoredCannonBalbaro,
	"99791f19-21ee-4ee3-a6b0-e26702c2a380": dm02.CriticalBlade,
	"5781986a-4ca8-48e1-a5a1-95e820455bce": dm02.HypersquidWalter,
	"6406411f-e9ae-4b55-8213-7e18c4cd9aee": dm02.StainedGlass,
	"9275747d-f2bb-4298-9b70-7075b17d1e0d": dm02.XenoMantis,
	"206c53d6-8068-4fec-bc5b-287de58eed5f": dm02.WynTheOracle,
	"a219c895-5822-458e-862f-6ade1529c42e": dm02.LeapingTornadoHorn,
	"aa1cca8a-3140-4180-9c9f-3f126dd16a68": dm02.ReconOperation,
	"ac636f90-215f-4934-bcd3-abfec69a5fda": dm02.AquaBouncer,
	"b85b0492-132c-46ac-85e6-c8b5b466598a": dm02.LagunaLightningEnforcer,
	"b691251e-8b9b-4436-8ca8-d9d3aba3ac72": dm02.ResoPacosClearSkyGuardian,
	"ba8c4d25-a352-4725-a3aa-6eb015f99d6c": dm02.Galsaur,
	"bf5ce598-f604-4c55-b0c9-9ed0e059aeec": dm02.ManaCrisis,
	"c125f786-e6d5-4477-8ab0-1e92d6eed348": dm02.CavalryGeneralCuratops,
	"cc5c643e-5120-421d-b26f-76be381dead7": dm02.RumbleGate,
	"cf536a4d-1512-44d8-b0e6-615d0725e3f0": dm02.MetalwingSkyterror,
	"d7e4a7da-6131-42e9-adf1-de0178549bd5": dm02.SilverFist,
	"e1e112d7-11e1-4f01-9c91-00a2b1626043": dm02.Bombersaur,
	"eac1bc57-bdf8-4629-86b3-9609d1bf2aba": dm02.ArmoredBlasterValdios,
	"ec3213d2-6914-4091-bf86-869354a36b15": dm02.PoisonWorm,
	"ecd61b34-85ca-4a91-b993-7a8c9976a2b5": dm02.Corile,
	"f4a364f5-d0e9-4777-b51e-6dc6e39b803c": dm02.AquaShooter,
	"f9b47778-c7d5-48da-aac3-1ddf8efbbbf1": dm02.PlasmaChaser,
	"fae765a5-fa73-4627-94d1-f74f2ed00792": dm02.LostSoul,
	"fcf2f484-c471-4a5d-bc41-1fcd56604d73": dm02.GeneralDarkFiend,
	"fd9f6d26-8fd8-43f2-84e5-b83493a1106e": dm02.LarbaGeerTheImmaculate,
	"ff33e31f-4f1c-4c07-b293-2e65a13e1077": dm02.RainbowStone,
	"ed3931f2-2b60-48d4-b82f-361d3a8c5bd7": dm02.EssenceElf,
	"a73bf1d3-b4a6-48cd-89d3-1983f160304d": dm02.DiamondCutter,
}

// DM03 is a map with all the card id's in the game and corresponding CardConstructor for dm03
var DM03 = map[string]match.CardConstructor{
	"70270aa3-ff24-476c-be22-5b9f48fc682a": dm03.MiarCometElemental,
	"9c7e3304-3aff-4362-a687-b5ca5333fe98": dm03.ChaosFish,
	"9df4d8ac-3c86-4c1f-b916-c9a384b0340f": dm03.GirielGhastlyWarrior,
	"41c2e4dc-460f-459a-b7cf-ef17b5c9a4eb": dm03.GarkagoDragon,
	"54652ec5-3bc3-4124-9c67-b05ba56def5f": dm03.EarthstompGiant,
	"4459f97b-8927-4231-a11d-f4f88175b71c": dm03.AlekSolidityEnforcer,
	"9d9468f8-3eb3-448c-8638-5d77c77f936b": dm03.AlessTheOracle,
	"b275fbf0-5355-45ec-b3a8-a956cf898ae6": dm03.BoomerangComet,
	"6b6ff714-0602-48da-9ee0-2578aa8ea32e": dm03.LenaVizierOfBrilliance,
	"96370266-e077-47c1-968f-0184bcb94227": dm03.LogicSphere,
	"b052351d-dd6c-4c40-a6d7-840e23398ff1": dm03.RaVuSeekerOfLightning,
	"eb46cc3f-75b9-4b50-851f-bf1fd3e76c06": dm03.RazaVegaThunderGuardian,
	"c49be3ab-da9b-4af8-8377-0595ca3160ce": dm03.SiegBaliculaTheIntense,
	"b9b1a83a-626c-471a-82ff-3e2372c9d838": dm03.SparkleFlower,
	"a097478c-af87-4316-b42d-888418d07919": dm03.SundropArmor,
	"7fdf0d01-a999-4f31-8a37-945396796998": dm03.UrPaleSeekerOfSunlight,
	"f76fd433-24ba-41e3-81eb-622cad7247e2": dm03.AnglerCluster,
	"23a97f92-d53d-4b9f-8e33-ca090bd30c1e": dm03.AquaDeformer,
	"d96cdbba-0e60-4d8e-9394-4830b11f559d": dm03.Emeral,
	"4852b823-7b96-4669-809b-edc14c042567": dm03.FloodValve,
	"39e0b4b6-eaa5-4673-9538-aa9d2c7fc5d8": dm03.KingNeptas,
	"2f2dae63-bb68-4d89-b7b1-48dd14eab40e": dm03.KingPonitas,
	"232a9fdd-e289-454f-9593-5b766c969a1e": dm03.LegendaryBynor,
	"e14514a5-592b-40ed-b693-7cb8f649d93a": dm03.LiquidScope,
	"c4f3db71-e9e7-4ab6-95fc-1174b4f42aca": dm03.PsychicShaper,
	"0f9a0c09-1d5b-4162-94ca-5399f8f37d63": dm03.Shtra,
	"fb3a5abd-c16a-4365-98c0-79753fdcd91d": dm03.StingerBall,
	"e4ad6b63-5c97-4002-9b0a-b38ba8b5312b": dm03.BaragaBladeOfGloom,
	"dae59bf4-cb23-40d0-aed2-bdbf953f73f5": dm03.BonePiercer,
	"826de4ef-a10d-410a-aed1-066783babbdc": dm03.EldritchPoison,
	"b3975c0b-2978-4b1a-8225-78d420ff941d": dm03.GamilKnightOfHatred,
	"21a9e35f-4219-490c-8963-3e1fee66c71b": dm03.GhastlyDrain,
	"03b16fba-71e7-4328-a605-4643bb5cfbf8": dm03.HangWormFetidLarva,
	"12d21399-f499-432b-bbb0-8ca1088b33c7": dm03.JackViperShadowofDoom,
	"ec9ebf80-3060-48cf-888a-bd00b61b3746": dm03.Mudman,
	"c08d8e5a-9f4b-4fe4-bf20-d5c7fca9ea15": dm03.Scratchclaw,
	"f773e1f2-37d0-45c2-a34f-63e91094aab1": dm03.SnakeAttack,
	"bfef8796-4214-4cfe-af86-4a8608a6fd7f": dm03.WailingShadowBelbetphlo,
	"b6109c08-3d01-486d-b544-5dfc11d341a0": dm03.ArmoredWarriorQuelos,
	"479b477f-f535-4564-ac0c-5f0aeaafe914": dm03.BabyZoppe,
	"d74a5d6f-a24e-4ecb-a163-7c696be8c06c": dm03.BlazeCannon,
	"e27ac147-3d7d-42d7-a3a4-2e3a1eccdb2c": dm03.BoltailDragon,
	"0f5a2f7e-74df-4ae5-ab0e-300b1761b417": dm03.ExplosiveDudeJoe,
	"caca44fe-8e47-4863-b1e9-7c7c09eb7ddc": dm03.Flametropus,
	"89bf69b3-25a7-4f9f-884a-fb9e7dd93efd": dm03.MuramasaDukeOfBlades,
	"d2cc5ca6-c0a3-4a83-935f-08d0e0878983": dm03.SearingWave,
	"cf6673ec-4bae-4891-9c80-e3e3a8269914": dm03.SnipStrikerBullraizer,
	"5646364b-4018-4556-9b20-265ef3fb372d": dm03.UberdragonJabaha,
	"0a5600c5-65a0-439d-b19c-4bf5ee036e32": dm03.VolcanicArrows,
	"9cff685f-3491-400b-b9d4-43c383193dc7": dm03.AuroraOfReversal,
	"22dae49f-f802-46d1-ad1e-400b980b80ab": dm03.DawnGiant,
	"e202fcb0-08a1-46ee-8654-66406ea436ce": dm03.Gigamantis,
	"5170634d-1574-4a06-b3bd-3c546fb9191e": dm03.ManaNexus,
	"f200f25c-479b-4381-967e-590bedd9dc90": dm03.MaskedPomegranate,
	"c3c09eac-d9b9-4eb6-8056-94afa4ba3927": dm03.PouchShell,
	"64371ab9-a8bc-4ec9-9b9b-f190c8ac1001": dm03.Psyshroom,
	"e8910059-0471-44ac-a573-a3fe8ed207d9": dm03.RagingDashHorn,
	"fc25a99c-7ef9-44bc-9204-ca84f6248116": dm03.RoarOfTheEarth,
	"8dab06d5-93c3-4ac7-abfb-3ce2e63d77e5": dm03.SniperMosquito,
	"84595f9f-c9c1-4d50-8e8f-29e5ef63bfbf": dm03.SwordButterfly,
}
