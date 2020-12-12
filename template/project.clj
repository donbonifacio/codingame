(defproject donbonifacio/coding-challenge-template "0.1.0-SNAPSHOT"
  :description "FIXME: write description"
  :url "http://example.com/FIXME"
  :license {:name "Eclipse Public License"
            :url "http://www.eclipse.org/legal/epl-v10.html"}
  :dependencies [[org.clojure/clojure "1.8.0"]]

  :aliases {"autotest" ["trampoline" "with-profile" "+test,+test-deps" "test-refresh"]
            "test"  ["trampoline" "with-profile" "+test,+test-deps" "test"]}

  :main Solution

  :profiles {:uberjar {:aot :all}
             :test-deps {:dependencies [[org.clojure/tools.namespace "0.2.11"]]

                         :plugins [[com.jakemccrary/lein-test-refresh "0.15.0"]]}}

  :test-refresh {:quiet true
                 :changes-only true})
