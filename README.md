# Useless-Fact-CLI-Tool

This command line tool was built in Go using the [Cobra library](https://github.com/spf13/cobra)

./uselessfact today - provides the useless fact of the day
./uselessfact random - provides a random useless fact

both commands allow the use of the language flag -l which can be used to determine which language the fact will be given in.

These facts are fetched from the uselessfacts.jsph.pl api. All languages supported by this api are also supported by the tool.
