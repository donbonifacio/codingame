(ns Solution-test
  (:require [clojure.test :refer :all]
            [Solution :as solution]))

(def morse solution/morse)

(deftest example-test
  (solution/clear-cache)
  (let [morse-sequence "......-...-..---.-----.-..-..-.."
        dictionary ["HELL"
                    "HELLO"
                    "OWORLD"
                    "WORLD"
                    "TEST"]]
    (is (= 2 (solution/indexed-morse-counter morse-sequence 0 (map morse dictionary))))))

(deftest test-01
  (solution/clear-cache)
  (let [morse-sequence "-.-"
        dictionary ["A"
                    "B"
                    "C"
                    "HELLO"
                    "K"
                    "WORLD"]]
    (is (= 1 (solution/indexed-morse-counter morse-sequence 0 (map morse dictionary))))))

(deftest test-02
  (solution/clear-cache)
  (let [morse-sequence "--.-------.."
        dictionary ["GOD"
                    "GOOD"
                    "MORNING"
                    "G"
                    "HELLO"]]
    (is (= 1 (solution/indexed-morse-counter morse-sequence 0 (map morse dictionary))))))

(deftest test-03
  (solution/clear-cache)
  (let [morse-sequence "......-...-..---.-----.-..-..-.."
        dictionary ["HELL"
                    "HELLO"
                    "OWORLD"
                    "WORLD"
                    "TEST"]]
    (is (= 2 (solution/indexed-morse-counter morse-sequence 0 (map morse dictionary))))))

(deftest test-net-01
  (solution/clear-cache)
  (let [morse-sequence "....----"
        dictionary ["E"
                    "EE"
                    "T"
                    "TT"]]
    (is (= 25 (solution/indexed-morse-counter morse-sequence 0 (map morse dictionary))))))

(deftest test-net-02
  (solution/clear-cache)
  (let [morse-sequence "........."
        dictionary ["E"
                    "EEE"]]
    (is (= 19 (solution/indexed-morse-counter morse-sequence 0 (map morse dictionary))))))

(deftest spaceless-message-test
  (is (= (morse "HELLO") "......-...-..---"))
  (is (= (morse "HELLO" " ") ".... . .-.. .-.. ---")))

(deftest index-dictionary-test
  (let [index (solution/index-dictionary ["HELLO" "HELL"])]
    (is (not (get-in index [\H \E :word?])))
    (is (true? (get-in index [\H \E \L \L :word?])))
    (is (true? (get-in index [\H \E \L \L \O :word?])))))
