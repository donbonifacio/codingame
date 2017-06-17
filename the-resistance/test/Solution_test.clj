(ns Solution-test
  (:require [clojure.test :refer :all]
            [Solution :as solution]))

(def morse solution/morse)

(deftest example-test
  (let [morse-sequence "......-...-..---.-----.-..-..-.."
        dictionary ["HELL"
                    "HELLO"
                    "OWORLD"
                    "WORLD"
                    "TEST"]]
    (is (= 2 (solution/morse-counter morse-sequence 0 (map morse dictionary))))))

(deftest test-01
  (let [morse-sequence "-.-"
        dictionary ["A"
                    "B"
                    "C"
                    "HELLO"
                    "K"
                    "WORLD"]]
    (is (= 1 (solution/morse-counter morse-sequence 0 (map morse dictionary))))))

(deftest test-02
  (let [morse-sequence "--.-------.."
        dictionary ["GOD"
                    "GOOD"
                    "MORNING"
                    "G"
                    "HELLO"]]
    (is (= 1 (solution/morse-counter morse-sequence 0 (map morse dictionary))))))

(deftest test-03
  (let [morse-sequence "......-...-..---.-----.-..-..-.."
        dictionary ["HELL"
                    "HELLO"
                    "OWORLD"
                    "WORLD"
                    "TEST"]]
    (is (= 2 (solution/morse-counter morse-sequence 0 (map morse dictionary))))))

(deftest spaceless-message-test
  (is (= (morse "HELLO") "......-...-..---"))
  (is (= (morse "HELLO" " ") ".... . .-.. .-.. ---")))

(defn test-possible-n [expected-n text dictionary]
  (is (= expected-n
         (solution/morse-counter
           (morse text)
           0
           (map morse dictionary)))))

(deftest possible-words-test
  (testing "simple sequence"
    (test-possible-n 1 "HELLO" ["HELLO"])
    (test-possible-n 1 "HELLOWORLD" ["HELLO" "WORLD" "TEST"])
    (test-possible-n 1 "HELLOWORLDTEST" ["HELLO" "WORLD" "TEST"]))
  (testing "simple branch"
    (test-possible-n 2 "HELLOWORLD" ["HELLO" "WORLD" "HELL" "OWORLD"])))

(deftest word-match-test
  (is (= 4 (solution/word-match "HELLO" "HELL" 0)))
  (is (= -2 (solution/word-match "HELL" "HELLO" 0)))

  (is (= -1 (solution/word-match "HELLO" "WAZA" 0)))

  (is (= 5 (solution/word-match "HELLOWORLD" "HELLO" 0)))
  (is (= (count "HELLOWORLD") (solution/word-match "HELLOWORLD" "WORLD" 5))))

(deftest morse-counter-test
  (is (= 2 (solution/morse-counter (morse "HELLOWORLD")
                                   0
                                   (map morse ["HELLO" "WORLD"
                                               "HELL" "OWORLD"])))))

(deftest index-dictionary-test
  (let [index (solution/index-dictionary ["HELLO" "HELL"])]
    (is (not (get-in index [\H \E :word?])))
    (is (true? (get-in index [\H \E \L \L :word?])))
    (is (true? (get-in index [\H \E \L \L \O :word?])))))
