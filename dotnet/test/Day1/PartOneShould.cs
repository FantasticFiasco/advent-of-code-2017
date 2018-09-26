using System.Linq;
using Shouldly;
using Xunit;

namespace AdventOfCode.Day1
{
    public class PartOneShould
    {
        [Theory]
        [InlineData("1122", 3)]
        [InlineData("1111", 4)]
        [InlineData("1234", 0)]
        [InlineData("91212129", 9)]
        [InlineData(Input, 1069)]
        public void Satisfy(string sequence, int expected)
        {
            // Arrange
            var input = sequence
                .ToCharArray()
                .Select(digit => int.Parse(digit.ToString()))
                .ToArray();

            // Act
            var actual = PartOne.Solve(input);

            // Assert
            actual.ShouldBe(expected);
        }

        private const string Input =
            "3674367652242621474168763928218321697812856559411236481728359862138483975662842414677931192831838359723596864468766515959157341323361" +
            "6717112157752469191845757712928347624726438516211153946892241449523148419426259291788938621886334734497823163281389389853675932246734" +
            "1535638612338949526576258684154323161554872428137984257797435619875637349449628468652637227127686748382444443857685684898429898781636" +
            "5577184736265615337226594546412866841243924896693939876544617185514454428546351725874981373131436594737254881143464638159527317298246" +
            "6142248474238762554858654679415418693478512641864168398722199638775667744977941183772494538685398862344164521446115925528534491788728" +
            "4486684553495889724432953913853895517832894173498233833247484116891982193299966667522518155625227593745426529691476964196699145345867" +
            "3243691279851969772258679574637169733841671684221431339322858741339953471639498418394312337551781962283797279643116626464643289347855" +
            "7659387795573234889141897313158457637142238315327877493994933514112645586351127139429281675912366669475931711974332271368287413985682" +
            "3749431958864559278395739864645551416792919986459366836391625883759745494677676234639355618478695273833952782489523147921121131262312" +
            "4674275311974811382884391781254722449831984994751774562584481917597398684363662841496566446658217241919722769536849243335319923355887" +
            "2319529626825788288176275546566474824257336863977574347328469153319428883748696399544974133392589823343773897313173336568883385364166" +
            "3363623986366844598862839642422492289383832192555139964685869535196381115999352291152288375592427529259436536236829855763239294154454" +
            "4337818947278245495823234198662679118286164411297441823928648672265444214485117353875685964721876813457285833184954326616967274522139" +
            "1659363674921469481143686952478771714585793322926824623482923579986434741714167134346384551362664177865452895348948953472328966995731" +
            "1696725735556219395848721879993253223278933367366119297526132419352116642489615276877783719712596545412394717667144691222137933484144" +
            "77789271187324629397292446879752673";
    }
}
