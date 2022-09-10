## Details

### Output
This program outputs a cheat file of the format accepted by retroarch, specifically when running amberelec:
```
cheats = 1290
...
cheat893_desc = "wild pokemon modifier 2.0.3.2: MUNCHLAX"
cheat893_code = "000014D1+000A+8202404C+01F3"
cheat893_enable = false
...
```

You will want to place this folder into the default directory listed in your retroarch settings menu, or set a custom directory.
You will need to ssh/sftp into your instance to create/access the folders available in the partition used by retroarch.

After you've created/accessed your directories and placed your files through sftp, once booting a game, in the retroarch cheats menu, you will see all cheats listed as available and than can enable them and persist your settings if desired.

### Input
This accepts an input file of the following format:
```
<cheats global description ex. wild pokemon modifier>

code template - <ex. 82025804 YYYY>

#ex. comments require a # at the beginning of the line and can be placed on any line

Codes (<list or single>)
<your codes list or a single code>
----------
...
<repeat above layout>
```

Note that code template is optional for debuggign purposes.
Note the dashes are the delimiter between global cheat blocks.
Note the last cheat block should not have a delimiter.

See complete examples below:

single layout
```
infinite money

Codes (single)
820257BC 423F+820257BE 000F
----------
```

list layout
```
unlimited healing items

code template - 82025804 YYYY

Codes (list)
#tested and functional
82025804 000D - Potion
82025804 000E - Antidote
82025804 000F - Burn Heal
82025804 0010 - Ice Heal
82025804 0011 - Awakening
82025804 0012 - Parlyz Heal
82025804 0013 - Full Restore
82025804 0014 - Max Potion
82025804 0015 - Hyper Potion
82025804 0016 - Super Potion
82025804 0017 - Full Heal
82025804 0018 - Revive
82025804 0019 - Max Revive
----------
```
