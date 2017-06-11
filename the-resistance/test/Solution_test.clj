(ns Solution-test
  (:require [clojure.test :refer :all]
            [Solution :as solution]))

#_(deftest example-test
  (let [morse-sequence "......-...-..---.-----.-..-..-.."
        dictionary ["HELL"
                    "HELLO"
                    "OWORLD"
                    "WORLD"
                    "TEST"]]
    (is (= 2 (solution/number-of-possible-messages morse-sequence dictionary)))))

(deftest spaceless-message-test
  (is (= (solution/morse "HELLO") ".....-.-...-..---"))
  (is (= (solution/morse "HELLO" " ") ".... .- .-.. .-.. ---")))
