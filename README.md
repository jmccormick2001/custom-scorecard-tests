
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
