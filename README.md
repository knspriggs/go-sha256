```
17:50:04 kristianspriggs@Kristians-MacBook-Pro-8.local concurrent-sha256 master ? go run main.go
----- Correctness Tests ------
In:
e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855 matches expected
Took:  28.923us
-----
In: a
ca978112ca1bbdcafac231b39a23dc4da786eff8147c4e72b9807785afee48bb matches expected
Took:  4.216us
-----
In: ab
fb8e20fc2e4c3f248c60c39bd652f3c1347298bb977b8b4d5903b85055620603 matches expected
Took:  2.711us
-----
In: abc
ba7816bf8f01cfea414140de5dae2223b00361a396177a9cb410ff61f20015ad matches expected
Took:  2.647us
-----
In: abcd
88d4266fd4e6338d13b845fcf289579d209c897823b9217da3e161936f031589 matches expected
Took:  2.306us
-----
In: abcde
36bbe50ed96841d1443bcb670d6554fa34b761be67ec9c4a8ad2c0c44ca42c does not match 36bbe50ed96841d10443bcb670d6554f0a34b761be67ec9c4a8ad2c0c44ca42c
Took:  2.847us
-----
In: abcdef
bef57ec7f53a6d40beb640a780a639c83bc29ac8a9816f1fc6c5c6dcd93c4721 matches expected
Took:  3.294us
-----
In: abcdefg
7d1a54127b222502f5b79b5fb0803061152a44f92b37e23c6527baf665d4da9a matches expected
Took:  2.693us
-----
In: abcdefgh
9c56cc51b374c3ba189210d5b6d4bf57790d351c96c47c02190ecf1e430635ab matches expected
Took:  2.899us
-----
In: He who has a shady past knows that nice guys finish last.
d83d974521b211ec9eac46ead4b5ab1a301406266949f8037e28cab8247321c9 does not match 6dae5caa713a10ad04b46028bf6dad68837c581616a1589a265a11288d4bb5c4
Took:  4.288us
-----
In: I wouldn't marry him with a ten foot pole.
ae7a702a9509039ddbf29f0765e70d001177914b86459284dab8b348c2dce3f does not match ae7a702a9509039ddbf29f0765e70d0001177914b86459284dab8b348c2dce3f
Took:  2.844us
-----
In: Free! Free!/A trip/to Mars/for 900/empty jars/Burma Shave
802f984c15de3084529149098f906d6b9ed063980107080417d3e841a01d017 does not match 6748450b01c568586715291dfa3ee018da07d36bb7ea6f180c1af6270215c64f
Took:  2.044us
-----
In: The days of the digital watch are numbered.  -Tom Stoppard
f30483615bd8facc2c96851310bb97b7a01fafc5a3f9ddb46484b36c688f421e does not match 14b82014ad2b11f661b5ae6a99b75105c2ffac278cd071cd6c05832793635774
Took:  2.107us
-----
In: Nepal premier won't resign.
7102cfd76e2e324889eece5d6c41921b1e142a4ac5a2692be78803097f6a48d8 matches expected
Took:  2.316us
-----
In: For every action there is an equal and opposite government program.
23b1018cd81db1d67983c5f7417c44da9deb582459e378d7a068552ea649dc9f matches expected
Took:  14.818us
-----
In: His money is twice tainted: 'taint yours and 'taint mine.
56d6991f9336cb52d24096f4211a2f141a4c9fd3c14524b9c06097cc9780e6e does not match 8001f190dfb527261c4cfcab70c98e8097a7a1922129bc4096950e57c7999a5a
Took:  2.726us
-----
In: There is no reason for any individual to have a computer in their home. -Ken Olsen, 1977
8c87deb65505c3993eb24b7a150c4155e82eee6960cf0c3a8114ff736d69cad5 matches expected
Took:  3.107us
-----
In: It's a tiny change to the code and not completely disgusting. - Bob Manchek
bfb0a67a19cdec3646498b2ef751bddc41bba4b7f30081bb932aad214d16d7 does not match bfb0a67a19cdec3646498b2e0f751bddc41bba4b7f30081b0b932aad214d16d7
Took:  3.036us
-----
In: size:  a.out:  bad magic
7f9a0b9bf56332e19f5a0ec1ad9c1425a153da1c624868fda44561d6b74daf36 matches expected
Took:  2.554us
-----
In: The fugacity of a constituent in a mixture of gases at a given temperature is proportional to its mole fraction.  Lewis-Randall Rule
395585ce30617b62c80b93e8208ce866d4edc811a177fdb4b82d3911d8696423 matches expected
Took:  3.719us
-----
-------------------------------
--------- Speed Tests ---------
In:
In length:  0
My library:		3.523us
Standard Library:	11.283us
-----
In: a
In length:  1
My library:		2.519us
Standard Library:	1.559us
-----
In: ab
In length:  2
My library:		5.023us
Standard Library:	2.164us
-----
In: abc
In length:  3
My library:		2.819us
Standard Library:	1.25us
-----
In: abcd
In length:  4
My library:		2.917us
Standard Library:	1.436us
-----
In: abcde
In length:  5
My library:		2.363us
Standard Library:	1.229us
-----
In: abcdef
In length:  6
My library:		3.065us
Standard Library:	1.339us
-----
In: abcdefg
In length:  7
My library:		2.202us
Standard Library:	1.12us
-----
In: abcdefgh
In length:  8
My library:		2.387us
Standard Library:	1.253us
-----
In: He who has a shady past knows that nice guys finish last.
In length:  57
My library:		2.245us
Standard Library:	1.984us
-----
In: I wouldn't marry him with a ten foot pole.
In length:  42
My library:		3.645us
Standard Library:	1.707us
-----
In: Free! Free!/A trip/to Mars/for 900/empty jars/Burma Shave
In length:  57
My library:		2.516us
Standard Library:	4.774us
-----
In: The days of the digital watch are numbered.  -Tom Stoppard
In length:  58
My library:		3.329us
Standard Library:	2.222us
-----
In: Nepal premier won't resign.
In length:  27
My library:		2.797us
Standard Library:	1.749us
-----
In: For every action there is an equal and opposite government program.
In length:  67
My library:		4.763us
Standard Library:	1.787us
-----
In: His money is twice tainted: 'taint yours and 'taint mine.
In length:  57
My library:		2.008us
Standard Library:	1.897us
-----
In: There is no reason for any individual to have a computer in their home. -Ken Olsen, 1977
In length:  88
My library:		3.063us
Standard Library:	1.722us
-----
In: It's a tiny change to the code and not completely disgusting. - Bob Manchek
In length:  75
My library:		3.155us
Standard Library:	2.127us
-----
In: size:  a.out:  bad magic
In length:  24
My library:		2.174us
Standard Library:	1.13us
-----
In: The fugacity of a constituent in a mixture of gases at a given temperature is proportional to its mole fraction.  Lewis-Randall Rule
In length:  132
My library:		5.872us
Standard Library:	2.815us
-----
-------------------------------
```
