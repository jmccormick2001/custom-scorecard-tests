
# custom scorecard tests example

This example shows how someone can write custom scorecard tests
using the alpha version of operator-sdk scorecard.

## Example scorecard config.yaml

This example would have you defining the following scorecard
config.yaml entries:

```
- name: "customtest1"
  image: quay.io/jemccorm/custom-scorecard-tests
  entrypoint:
  - custom-scorecard-tests
  - customtest1
  labels:
    suite: custom
    test: customtest1
  description: an ISV custom test that does...
- name: "customtest2"
  entrypoint:
  - custom-scorecard-tests
  - customtest2
  image: quay.io/jemccorm/custom-scorecard-tests
  labels:
    suite: custom
    test: customtest2
```

## Example Execution Command

The *command-example* script shows how to run the
custom scorecard test using operator-sdk alpha scorecard.  It looks
like this:
```
operator-sdk alpha scorecard \
--bundle=bundle/ \
--selector=suite=custom \
--list=false \
-o json \
--wait-time=40s \
--skip-cleanup=false
```

Notice that we are using the selector flag to only run the custom
tests.   The scorecard configuration is found in the *bundle*
sub-directory.
